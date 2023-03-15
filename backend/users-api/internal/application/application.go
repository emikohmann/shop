package application

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/config"
    transportHTTP "users-api/pkg/transport/http"
    userService "users-api/pkg/users"
    userMetrics "users-api/pkg/users/metrics"
    userQueues "users-api/pkg/users/queues"
    userRepositories "users-api/pkg/users/repositories"
)

type application struct {
    logger *logger.Logger
    config *config.Config
    router *gin.Engine
}

type usersRepository interface {
    GetUser(ctx context.Context, id int64) (userService.User, apierrors.APIError)
    ListUsers(ctx context.Context, limit int, offset int) (userService.UserList, apierrors.APIError)
    SaveUser(ctx context.Context, user userService.User) apierrors.APIError
    UpdateUser(ctx context.Context, user userService.User) apierrors.APIError
    DeleteUser(ctx context.Context, id int64) apierrors.APIError
}

type usersMetrics interface {
    NotifyMetric(ctx context.Context, action userService.Action)
}

type usersQueue interface {
    PublishUserNotification(ctx context.Context, action userService.Action, priority userService.Priority, id int64) apierrors.APIError
}

type usersService interface {
    GetUser(ctx context.Context, id int64) (userService.User, apierrors.APIError)
    ListUsers(ctx context.Context, limit int, offset int) (userService.UserList, apierrors.APIError)
    SaveUser(ctx context.Context, user userService.User) (userService.User, apierrors.APIError)
    UpdateUser(ctx context.Context, user userService.User) (userService.User, apierrors.APIError)
    DeleteUser(ctx context.Context, id int64) apierrors.APIError
}

type metrics struct {
    usersPrometheusMetrics usersMetrics
}

type queues struct {
    usersRabbitMQQueue usersQueue
}

type repositories struct {
    usersMySQLRepository usersRepository
}

type services struct {
    usersService usersService
}

type handlers struct {
    docsHandler       gin.HandlerFunc
    metricsHandler    gin.HandlerFunc
    getUserHandler    gin.HandlerFunc
    listUsersHandler  gin.HandlerFunc
    saveUserHandler   gin.HandlerFunc
    updateUserHandler gin.HandlerFunc
    deleteUserHandler gin.HandlerFunc
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
    usersPrometheusMetrics := userMetrics.NewPrometheusMetrics(logger)
    logger.Info(ctx, "Metrics successfully initialized")
    return metrics{
        usersPrometheusMetrics: usersPrometheusMetrics,
    }, nil
}

// buildQueues creates the instances for the queues
func buildQueues(ctx context.Context, logger *logger.Logger, config *config.Config) (queues, error) {
    usersRabbitMQQueue, err := userQueues.NewUsersRabbitMQ(
        ctx,
        config.UsersRabbitMQ.Host,
        config.UsersRabbitMQ.Port,
        config.UsersRabbitMQ.User,
        config.UsersRabbitMQ.Password,
        config.UsersRabbitMQ.QueueName,
        config.App.Name,
        logger,
    )
    if err != nil {
        return queues{}, fmt.Errorf("error initializing users RabbitMQ queue: %w", err)
    }
    logger.Info(ctx, "Queues successfully initialized")
    return queues{
        usersRabbitMQQueue: usersRabbitMQQueue,
    }, nil
}

// buildRepositories creates the instances for the repositories
func buildRepositories(ctx context.Context, logger *logger.Logger, config *config.Config) (repositories, error) {
    usersMySQLRepository, err := userRepositories.NewUsersMySQL(
        ctx,
        config.UsersMySQL.Host,
        config.UsersMySQL.Port,
        config.UsersMySQL.Database,
        config.UsersMySQL.User,
        config.UsersMySQL.Password,
        config.UsersMySQL.Table,
        logger)
    if err != nil {
        return repositories{}, fmt.Errorf("error initializing users MySQL repository: %w", err)
    }
    logger.Info(ctx, "Repositories successfully initialized")
    return repositories{
        usersMySQLRepository: usersMySQLRepository,
    }, nil
}

// buildServices creates the instances for the services
func buildServices(ctx context.Context, repositories repositories, metrics metrics, queues queues, logger *logger.Logger) (services, error) {
    usersService := userService.NewService(
        repositories.usersMySQLRepository,
        metrics.usersPrometheusMetrics,
        queues.usersRabbitMQQueue,
        logger)
    logger.Info(ctx, "Services successfully initialized")
    return services{
        usersService: usersService,
    }, nil
}

// buildHandlers creates the instances for the handlers
func buildHandlers(ctx context.Context, logger *logger.Logger, services services) (handlers, error) {
    docsHandler := transportHTTP.DocsHandler(logger)
    metricsHandler := transportHTTP.MetricsHandler(logger)
    getUserHandler := transportHTTP.GetUserHandler(ctx, services.usersService, logger)
    listUsersHandler := transportHTTP.ListUsersHandler(ctx, services.usersService, logger)
    saveUserHandler := transportHTTP.SaveUserHandler(ctx, services.usersService, logger)
    updateUserHandler := transportHTTP.UpdateUserHandler(ctx, services.usersService, logger)
    deleteUserHandler := transportHTTP.DeleteUserHandler(ctx, services.usersService, logger)
    logger.Info(ctx, "Handlers successfully initialized")
    return handlers{
        docsHandler:       docsHandler,
        metricsHandler:    metricsHandler,
        getUserHandler:    getUserHandler,
        listUsersHandler:  listUsersHandler,
        saveUserHandler:   saveUserHandler,
        updateUserHandler: updateUserHandler,
        deleteUserHandler: deleteUserHandler,
    }, nil
}

// mapRouter creates the connections between the router and the handlers
func mapRouter(ctx context.Context, logger *logger.Logger, router *gin.Engine, handlers handlers) error {
    middleware := transportHTTP.Middleware(logger)
    router.GET(transportHTTP.PathDocs, middleware, handlers.docsHandler)
    router.GET(transportHTTP.PathMetrics, middleware, handlers.metricsHandler)
    router.GET(transportHTTP.PathGetUser, middleware, handlers.getUserHandler)
    router.GET(transportHTTP.PathListUsers, middleware, handlers.listUsersHandler)
    router.POST(transportHTTP.PathSaveUser, middleware, handlers.saveUserHandler)
    router.PUT(transportHTTP.PathUpdateUser, middleware, handlers.updateUserHandler)
    router.DELETE(transportHTTP.PathDeleteUser, middleware, handlers.deleteUserHandler)
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
