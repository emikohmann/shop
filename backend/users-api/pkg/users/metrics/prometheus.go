package metrics

import (
    "context"
    prom "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "users-api/internal/logger"
    "users-api/pkg/users"
)

type prometheus struct {
    counters map[users.Action]prom.Counter
    logger   *logger.Logger
}

// NewPrometheusMetrics instances a new users' metric collector
func NewPrometheusMetrics(logger *logger.Logger) prometheus {
    counters := map[users.Action]prom.Counter{
        users.ActionGet: promauto.NewCounter(prom.CounterOpts{
            Name: "users_get",
            Help: "Counter for GET user operation",
        }),
        users.ActionList: promauto.NewCounter(prom.CounterOpts{
            Name: "users_list",
            Help: "Counter for LIST user operation",
        }),
        users.ActionSave: promauto.NewCounter(prom.CounterOpts{
            Name: "users_save",
            Help: "Counter for SAVE user operation",
        }),
        users.ActionUpdate: promauto.NewCounter(prom.CounterOpts{
            Name: "users_update",
            Help: "Counter for UPDATE user operation",
        }),
        users.ActionDelete: promauto.NewCounter(prom.CounterOpts{
            Name: "users_delete",
            Help: "Counter for DELETE user operation",
        }),
    }
    return prometheus{
        counters: counters,
        logger:   logger,
    }
}

// NotifyMetric increment corresponding users' operation counter in Prometheus
func (prometheus prometheus) NotifyMetric(ctx context.Context, action users.Action) {
    prometheus.counters[action].Inc()
}
