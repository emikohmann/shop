package application

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/internal/logger"
	"github.com/emikohmann/shop/backend/items-api/pkg/config"
	itemService "github.com/emikohmann/shop/backend/items-api/pkg/items"
	itemMetrics "github.com/emikohmann/shop/backend/items-api/pkg/items/metrics"
	itemQueues "github.com/emikohmann/shop/backend/items-api/pkg/items/queues"
	itemRepositories "github.com/emikohmann/shop/backend/items-api/pkg/items/repositories"
	transportHTTP "github.com/emikohmann/shop/backend/items-api/pkg/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type application struct {
	logger *logrus.Logger
	config *config.Config
	router *gin.Engine
}

type itemsRepository interface {
	GetItem(ctx context.Context, id int64) (itemService.Item, apierrors.APIError)
	SaveItem(ctx context.Context, item itemService.Item) apierrors.APIError
	UpdateItem(ctx context.Context, item itemService.Item) apierrors.APIError
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
}

type itemsMetrics interface {
	NotifyMetric(ctx context.Context, action itemService.Action)
}

type itemsQueue interface {
	PublishItemNotification(ctx context.Context, action itemService.Action, priority itemService.Priority, id int64) apierrors.APIError
}

type itemsService interface {
	GetItem(ctx context.Context, id int64) (itemService.Item, apierrors.APIError)
	SaveItem(ctx context.Context, item itemService.Item) (itemService.Item, apierrors.APIError)
	UpdateItem(ctx context.Context, item itemService.Item) (itemService.Item, apierrors.APIError)
	DeleteItem(ctx context.Context, id int64) apierrors.APIError
}

type metrics struct {
	itemsPrometheusMetrics itemsMetrics
}

type queues struct {
	itemsRabbitMQQueue itemsQueue
}

type repositories struct {
	itemsMongoDBRepository itemsRepository
}

type services struct {
	itemsService itemsService
}

type handlers struct {
	metricsHandler    gin.HandlerFunc
	getItemHandler    gin.HandlerFunc
	saveItemHandler   gin.HandlerFunc
	updateItemHandler gin.HandlerFunc
	deleteItemHandler gin.HandlerFunc
}

// NewApplication creates a new instance of the application
func NewApplication() (*application, error) {
	logger, err := buildLogger()
	if err != nil {
		return nil, err
	}

	config, err := buildConfig(logger)
	if err != nil {
		return nil, err
	}

	router, err := buildRouter(logger)
	if err != nil {
		return nil, err
	}

	repositories, err := buildRepositories(logger, config)
	if err != nil {
		return nil, err
	}

	metrics, err := buildMetrics(logger)
	if err != nil {
		return nil, err
	}

	queues, err := buildQueues(logger, config)
	if err != nil {
		return nil, err
	}

	services, err := buildServices(repositories, metrics, queues, logger)
	if err != nil {
		return nil, err
	}

	handlers, err := buildHandlers(logger, services)
	if err != nil {
		return nil, err
	}

	if err := mapRouter(logger, router, handlers); err != nil {
		return nil, err
	}

	return &application{
		logger: logger,
		config: config,
		router: router,
	}, nil
}

// buildLogger creates the instance for the logger
func buildLogger() (*logrus.Logger, error) {
	logger := logger.NewLogger(logrus.InfoLevel)
	logger.Info("Logger successfully initialized")
	return logger, nil
}

// buildConfig creates the instance for the config
func buildConfig(logger *logrus.Logger) (*config.Config, error) {
	config, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}
	logger.Info("Config successfully initialized")
	return config, nil
}

// buildRouter creates the instance for the router
func buildRouter(logger *logrus.Logger) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("error setting up HTTP trusted proxies: %w", err)
	}
	logger.Info("Router successfully initialized")
	return router, nil
}

// buildMetrics creates the instances for the metric collectors
func buildMetrics(logger *logrus.Logger) (metrics, error) {
	itemsPrometheusMetrics := itemMetrics.NewPrometheusMetrics(logger)
	logger.Info("Metrics successfully initialized")
	return metrics{
		itemsPrometheusMetrics: itemsPrometheusMetrics,
	}, nil
}

// buildQueues creates the instances for the queues
func buildQueues(logger *logrus.Logger, config *config.Config) (queues, error) {
	itemsRabbitMQQueue, err := itemQueues.NewItemsRabbitMQ(
		config.ItemsRabbitMQ.Host,
		config.ItemsRabbitMQ.Port,
		config.ItemsRabbitMQ.User,
		config.ItemsRabbitMQ.Password,
		config.ItemsRabbitMQ.QueueName,
		config.App.Name,
		logger,
	)
	if err != nil {
		return queues{}, fmt.Errorf("error initializing itemService RabbitMQ queue: %w", err)
	}
	logger.Info("Queues successfully initialized")
	return queues{
		itemsRabbitMQQueue: itemsRabbitMQQueue,
	}, nil
}

// buildRepositories creates the instances for the repositories
func buildRepositories(logger *logrus.Logger, config *config.Config) (repositories, error) {
	itemsMongoDBRepository, err := itemRepositories.NewItemsMongoDB(
		config.ItemsMongoDB.Host,
		config.ItemsMongoDB.Port,
		config.ItemsMongoDB.Database,
		config.ItemsMongoDB.Collection,
		logger)
	if err != nil {
		return repositories{}, fmt.Errorf("error initializing itemService MongoDB repository: %w", err)
	}
	logger.Info("Repositories successfully initialized")
	return repositories{
		itemsMongoDBRepository: itemsMongoDBRepository,
	}, nil
}

// buildServices creates the instances for the services
func buildServices(repositories repositories, metrics metrics, queues queues, logger *logrus.Logger) (services, error) {
	itemsService := itemService.NewService(
		repositories.itemsMongoDBRepository,
		metrics.itemsPrometheusMetrics,
		queues.itemsRabbitMQQueue,
		logger)
	logger.Info("Services successfully initialized")
	return services{
		itemsService: itemsService,
	}, nil
}

// buildHandlers creates the instances for the handlers
func buildHandlers(logger *logrus.Logger, services services) (handlers, error) {
	metricsHandler := transportHTTP.MetricsHandler(logger)
	getItemHandler := transportHTTP.GetItemHandler(services.itemsService, logger)
	saveItemHandler := transportHTTP.SaveItemHandler(services.itemsService, logger)
	updateItemHandler := transportHTTP.UpdateItemHandler(services.itemsService, logger)
	deleteItemHandler := transportHTTP.DeleteItemHandler(services.itemsService, logger)
	logger.Info("Handlers successfully initialized")
	return handlers{
		metricsHandler:    metricsHandler,
		getItemHandler:    getItemHandler,
		saveItemHandler:   saveItemHandler,
		updateItemHandler: updateItemHandler,
		deleteItemHandler: deleteItemHandler,
	}, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(logger *logrus.Logger, router *gin.Engine, handlers handlers) error {
	router.GET(transportHTTP.PathMetrics, handlers.metricsHandler)
	router.GET(transportHTTP.PathGetItem, handlers.getItemHandler)
	router.POST(transportHTTP.PathSaveItem, handlers.saveItemHandler)
	router.PUT(transportHTTP.PathUpdateItem, handlers.updateItemHandler)
	router.DELETE(transportHTTP.PathDeleteItem, handlers.deleteItemHandler)
	logger.Info("Router successfully mapped")
	return nil
}

// Run starts the application execution
func (app *application) Run() error {
	app.logger.Infof("Running application on :%d", app.config.HTTP.Port)
	if err := app.router.Run(fmt.Sprintf(":%d", app.config.HTTP.Port)); err != nil {
		return fmt.Errorf("error running HTTP server: %w", err)
	}
	return nil
}
