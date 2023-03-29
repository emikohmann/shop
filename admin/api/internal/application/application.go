package application

import (
	"api/internal/apierrors"
	"api/internal/logger"
	"api/pkg/admin"
	"api/pkg/config"
	"api/pkg/docker"
	"api/pkg/static"
	"api/pkg/transport/http"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type application struct {
	logger *logger.Logger
	config *config.Config
	router *gin.Engine
}

type staticConfig interface {
	Reload() error
	Data() static.Data
}

type dockerClient interface {
	ListImages(ctx context.Context) ([]types.ImageSummary, error)
	ListContainers(ctx context.Context) ([]types.Container, error)
}

type adminService interface {
	ListServices(ctx context.Context) ([]admin.Service, apierrors.APIError)
	GetService(ctx context.Context, id string) (admin.Service, admin.DockerAdditionalInfo, apierrors.APIError)
}

type statics struct {
	staticConfig staticConfig
}

type clients struct {
	dockerClient dockerClient
}

type services struct {
	adminService adminService
}

type handlers struct {
	docsHandler         gin.HandlerFunc
	listServicesHandler gin.HandlerFunc
	getServiceHandler   gin.HandlerFunc
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

	statics, err := buildStatics(ctx, logger)
	if err != nil {
		return nil, err
	}

	clients, err := buildClients(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	services, err := buildServices(ctx, logger, statics.staticConfig, clients.dockerClient)
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

// buildStatic creates the static config for the service
func buildStatics(ctx context.Context, logger *logger.Logger) (statics, error) {
	staticConfig, err := static.NewStatic(logger)
	if err != nil {
		return statics{}, fmt.Errorf("error creating static config: %w", err)
	}
	logger.Info(ctx, "Statics successfully initialized")
	return statics{
		staticConfig: staticConfig,
	}, nil
}

// buildClients creates the client connections for the service
func buildClients(ctx context.Context, config *config.Config, logger *logger.Logger) (clients, error) {
	dockerClient, err := docker.NewDocker(ctx, config.Docker.APIVersion, logger)
	if err != nil {
		return clients{}, fmt.Errorf("error creating dockerClient client: %w", err)
	}
	logger.Info(ctx, "Clients successfully initialized")
	return clients{
		dockerClient: dockerClient,
	}, nil
}

// buildServices creates the instances for the services
func buildServices(ctx context.Context, logger *logger.Logger, static staticConfig, docker dockerClient) (services, error) {
	adminService := admin.NewService(ctx, logger, static, docker)
	logger.Info(ctx, "Services successfully initialized")
	return services{
		adminService: adminService,
	}, nil
}

// buildHandlers creates the instances for the handlers
func buildHandlers(ctx context.Context, logger *logger.Logger, services services) (handlers, error) {
	docsHandler := http.DocsHandler(logger)
	listServicesHandler := http.ListServicesHandler(ctx, services.adminService, logger)
	getServiceHandler := http.GetServiceHandler(ctx, services.adminService, logger)
	logger.Info(ctx, "Handlers successfully initialized")
	return handlers{
		docsHandler:         docsHandler,
		listServicesHandler: listServicesHandler,
		getServiceHandler:   getServiceHandler,
	}, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(ctx context.Context, logger *logger.Logger, router *gin.Engine, handlers handlers) error {
	middleware := http.Middleware(logger)
	router.GET(http.PathDocs, middleware, handlers.docsHandler)
	router.GET(http.PathListServices, middleware, handlers.listServicesHandler)
	router.GET(http.PathGetService, middleware, handlers.getServiceHandler)
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
