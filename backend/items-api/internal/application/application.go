package application

import (
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/internal/logger"
	"github.com/emikohmann/shop/backend/items-api/pkg/config"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	transportHTTP "github.com/emikohmann/shop/backend/items-api/pkg/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type application struct {
	logger *logrus.Logger
	config *config.Config
	router *gin.Engine
}

type itemsService interface {
	Get(id int64) (items.Item, apierrors.APIError)
}

type services struct {
	itemsService itemsService
}

type handlers struct {
	getItemHandler func(ctx *gin.Context)
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

	services, err := buildServices(logger)
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
	logger.Debug("Logger successfully initialized")
	return logger, nil
}

// buildConfig creates the instance for the config
func buildConfig(logger *logrus.Logger) (*config.Config, error) {
	config, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}
	logger.Debug("Config successfully initialized")
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
	logger.Debug("Router successfully initialized")
	return router, nil
}

// buildServices creates the instances for the services
func buildServices(logger *logrus.Logger) (services, error) {
	itemsService := items.NewService()
	logger.Debug("Services successfully initialized")
	return services{
		itemsService: itemsService,
	}, nil
}

// buildServices creates the instances for the handlers
func buildHandlers(logger *logrus.Logger, services services) (handlers, error) {
	getItemHandler := transportHTTP.GetItemHandler(services.itemsService)
	logger.Debug("Handlers successfully initialized")
	return handlers{
		getItemHandler: getItemHandler,
	}, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(logger *logrus.Logger, router *gin.Engine, handlers handlers) error {
	router.GET(transportHTTP.GetItem, handlers.getItemHandler)
	logger.Debug("Router successfully mapped")
	return nil
}

// Run starts the application execution
func (app *application) Run() error {
	app.logger.Info("Running application")
	if err := app.router.Run(fmt.Sprintf(":%d", app.config.HTTP.Port)); err != nil {
		return fmt.Errorf("error running HTTP server: %w", err)
	}
	return nil
}
