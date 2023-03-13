package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"items-api/internal/apierrors"
	"items-api/internal/logger"
	"items-api/pkg/config"
	itemService "items-api/pkg/items"
	itemMetrics "items-api/pkg/items/metrics"
	itemQueues "items-api/pkg/items/queues"
	itemRepositories "items-api/pkg/items/repositories"
	transportHTTP "items-api/pkg/transport/http"
)

type application struct {
	logger *logger.Logger
	config *config.Config
	router *gin.Engine
}

type itemsRepository interface {
	GetItem(ctx context.Context, id int64) (itemService.Item, apierrors.APIError)
	ListItems(ctx context.Context, limit int, offset int) (itemService.ItemList, apierrors.APIError)
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
	ListItems(ctx context.Context, limit int, offset int) (itemService.ItemList, apierrors.APIError)
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
	docsHandler       gin.HandlerFunc
	metricsHandler    gin.HandlerFunc
	getItemHandler    gin.HandlerFunc
	listItemsHandler  gin.HandlerFunc
	saveItemHandler   gin.HandlerFunc
	updateItemHandler gin.HandlerFunc
	deleteItemHandler gin.HandlerFunc
}

// NewApplication creates a new instance of the application
func NewApplication(ctx context.Context) (*application, error) {
	logger, err := buildLogger(ctx)
	if err != nil {
		return nil, err
	}

	config, err := buildConfig(ctx, logger)
	if err != nil {
		return nil, err
	}

	router, err := buildRouter(ctx, logger)
	if err != nil {
		return nil, err
	}

	repositories, err := buildRepositories(ctx, logger, config)
	if err != nil {
		return nil, err
	}

	metrics, err := buildMetrics(ctx, logger)
	if err != nil {
		return nil, err
	}

	queues, err := buildQueues(ctx, logger, config)
	if err != nil {
		return nil, err
	}

	services, err := buildServices(ctx, repositories, metrics, queues, logger)
	if err != nil {
		return nil, err
	}

	handlers, err := buildHandlers(ctx, logger, services)
	if err != nil {
		return nil, err
	}

	if err := mapRouter(ctx, logger, router, handlers); err != nil {
		return nil, err
	}

	return &application{
		logger: logger,
		config: config,
		router: router,
	}, nil
}

// buildLogger creates the instance for the logger
func buildLogger(ctx context.Context) (*logger.Logger, error) {
	logger := logger.NewLogger(logrus.InfoLevel)
	logger.Info(ctx, "Logger successfully initialized")
	return logger, nil
}

// buildConfig creates the instance for the config
func buildConfig(ctx context.Context, logger *logger.Logger) (*config.Config, error) {
	config, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}
	logger.Info(ctx, "Config successfully initialized")
	return config, nil
}

// buildRouter creates the instance for the router
func buildRouter(ctx context.Context, logger *logger.Logger) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("error setting up HTTP trusted proxies: %w", err)
	}
	logger.Info(ctx, "Router successfully initialized")
	return router, nil
}

// buildMetrics creates the instances for the metric collectors
func buildMetrics(ctx context.Context, logger *logger.Logger) (metrics, error) {
	itemsPrometheusMetrics := itemMetrics.NewPrometheusMetrics(logger)
	logger.Info(ctx, "Metrics successfully initialized")
	return metrics{
		itemsPrometheusMetrics: itemsPrometheusMetrics,
	}, nil
}

// buildQueues creates the instances for the queues
func buildQueues(ctx context.Context, logger *logger.Logger, config *config.Config) (queues, error) {
	itemsRabbitMQQueue, err := itemQueues.NewItemsRabbitMQ(
		ctx,
		config.ItemsRabbitMQ.Host,
		config.ItemsRabbitMQ.Port,
		config.ItemsRabbitMQ.User,
		config.ItemsRabbitMQ.Password,
		config.ItemsRabbitMQ.QueueName,
		config.App.Name,
		logger,
	)
	if err != nil {
		return queues{}, fmt.Errorf("error initializing items RabbitMQ queue: %w", err)
	}
	logger.Info(ctx, "Queues successfully initialized")
	return queues{
		itemsRabbitMQQueue: itemsRabbitMQQueue,
	}, nil
}

// buildRepositories creates the instances for the repositories
func buildRepositories(ctx context.Context, logger *logger.Logger, config *config.Config) (repositories, error) {
	itemsMongoDBRepository, err := itemRepositories.NewItemsMongoDB(
		ctx,
		config.ItemsMongoDB.Host,
		config.ItemsMongoDB.Port,
		config.ItemsMongoDB.Database,
		config.ItemsMongoDB.Collection,
		logger)
	if err != nil {
		return repositories{}, fmt.Errorf("error initializing items MongoDB repository: %w", err)
	}
	logger.Info(ctx, "Repositories successfully initialized")
	return repositories{
		itemsMongoDBRepository: itemsMongoDBRepository,
	}, nil
}

// buildServices creates the instances for the services
func buildServices(ctx context.Context, repositories repositories, metrics metrics, queues queues, logger *logger.Logger) (services, error) {
	itemsService := itemService.NewService(
		repositories.itemsMongoDBRepository,
		metrics.itemsPrometheusMetrics,
		queues.itemsRabbitMQQueue,
		logger)
	logger.Info(ctx, "Services successfully initialized")
	return services{
		itemsService: itemsService,
	}, nil
}

// buildHandlers creates the instances for the handlers
func buildHandlers(ctx context.Context, logger *logger.Logger, services services) (handlers, error) {
	docsHandler := transportHTTP.DocsHandler(logger)
	metricsHandler := transportHTTP.MetricsHandler(logger)
	getItemHandler := transportHTTP.GetItemHandler(ctx, services.itemsService, logger)
	listItemsHandler := transportHTTP.ListItemsHandler(ctx, services.itemsService, logger)
	saveItemHandler := transportHTTP.SaveItemHandler(ctx, services.itemsService, logger)
	updateItemHandler := transportHTTP.UpdateItemHandler(ctx, services.itemsService, logger)
	deleteItemHandler := transportHTTP.DeleteItemHandler(ctx, services.itemsService, logger)
	logger.Info(ctx, "Handlers successfully initialized")
	return handlers{
		docsHandler:       docsHandler,
		metricsHandler:    metricsHandler,
		getItemHandler:    getItemHandler,
		listItemsHandler:  listItemsHandler,
		saveItemHandler:   saveItemHandler,
		updateItemHandler: updateItemHandler,
		deleteItemHandler: deleteItemHandler,
	}, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(ctx context.Context, logger *logger.Logger, router *gin.Engine, handlers handlers) error {
	middleware := transportHTTP.Middleware(logger)
	router.GET(transportHTTP.PathDocs, middleware, handlers.docsHandler)
	router.GET(transportHTTP.PathMetrics, middleware, handlers.metricsHandler)
	router.GET(transportHTTP.PathGetItem, middleware, handlers.getItemHandler)
	router.GET(transportHTTP.PathListItems, middleware, handlers.listItemsHandler)
	router.POST(transportHTTP.PathSaveItem, middleware, handlers.saveItemHandler)
	router.PUT(transportHTTP.PathUpdateItem, middleware, handlers.updateItemHandler)
	router.DELETE(transportHTTP.PathDeleteItem, middleware, handlers.deleteItemHandler)
	logger.Info(ctx, "Router successfully mapped")
	return nil
}

// Run starts the application execution
func (app *application) Run(ctx context.Context) error {
	app.logger.Infof(ctx, "Running application on :%d", app.config.HTTP.Port)
	if err := app.router.Run(fmt.Sprintf(":%d", app.config.HTTP.Port)); err != nil {
		return fmt.Errorf("error running HTTP server: %w", err)
	}
	return nil
}
