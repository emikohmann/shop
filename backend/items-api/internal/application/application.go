package application

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/internal/logger"
	"github.com/emikohmann/shop/backend/items-api/pkg/config"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	itemsRepositories "github.com/emikohmann/shop/backend/items-api/pkg/items/repositories"
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
	GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError)
	SaveItem(ctx context.Context, item items.Item) apierrors.APIError
	UpdateItem(ctx context.Context, item items.Item) apierrors.APIError
}

type itemsService interface {
	Get(ctx context.Context, id int64) (items.Item, apierrors.APIError)
	Save(ctx context.Context, item items.Item) apierrors.APIError
	Update(ctx context.Context, item items.Item) apierrors.APIError
}

type repositories struct {
	itemsMongoDBRepository itemsRepository
}

type services struct {
	itemsService itemsService
}

type handlers struct {
	getItemHandler    func(ctx *gin.Context)
	saveItemHandler   func(ctx *gin.Context)
	updateItemHandler func(ctx *gin.Context)
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

	services, err := buildServices(logger, repositories)
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
	logger := logger.NewLogger(logrus.DebugLevel)
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

// buildRepositories creates the instances for the repositories
func buildRepositories(logger *logrus.Logger, config *config.Config) (repositories, error) {
	itemsMongoDBRepository, err := itemsRepositories.NewItemsMongoDB(
		config.ItemsMongoDB.Host,
		config.ItemsMongoDB.Port,
		config.ItemsMongoDB.Database,
		config.ItemsMongoDB.Collection,
		logger)
	if err != nil {
		return repositories{}, fmt.Errorf("error initializing items MongoDB repository: %w", err)
	}
	logger.Info("Repositories successfully initialized")
	return repositories{
		itemsMongoDBRepository: itemsMongoDBRepository,
	}, nil
}

// buildServices creates the instances for the services
func buildServices(logger *logrus.Logger, repositories repositories) (services, error) {
	itemsService := items.NewService(repositories.itemsMongoDBRepository, logger)
	logger.Info("Services successfully initialized")
	return services{
		itemsService: itemsService,
	}, nil
}

// buildServices creates the instances for the handlers
func buildHandlers(logger *logrus.Logger, services services) (handlers, error) {
	getItemHandler := transportHTTP.GetItemHandler(services.itemsService, logger)
	saveItemHandler := transportHTTP.SaveItemHandler(services.itemsService, logger)
	updateItemHandler := transportHTTP.UpdateItemHandler(services.itemsService, logger)
	logger.Info("Handlers successfully initialized")
	return handlers{
		getItemHandler:    getItemHandler,
		saveItemHandler:   saveItemHandler,
		updateItemHandler: updateItemHandler,
	}, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(logger *logrus.Logger, router *gin.Engine, handlers handlers) error {
	router.GET(transportHTTP.PathGetItem, handlers.getItemHandler)
	router.POST(transportHTTP.PathSaveItem, handlers.saveItemHandler)
	router.PUT(transportHTTP.PathUpdateItem, handlers.updateItemHandler)
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
