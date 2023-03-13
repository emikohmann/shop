package metrics

import (
    "context"
    prom "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "items-api/internal/logger"
    "items-api/pkg/items"
)

type prometheus struct {
    counters map[items.Action]prom.Counter
    logger   *logger.Logger
}

// NewPrometheusMetrics instances a new items' metric collector
func NewPrometheusMetrics(logger *logger.Logger) prometheus {
    counters := map[items.Action]prom.Counter{
        items.ActionGet: promauto.NewCounter(prom.CounterOpts{
            Name: "items_get",
            Help: "Counter for GET item operation",
        }),
        items.ActionList: promauto.NewCounter(prom.CounterOpts{
            Name: "items_list",
            Help: "Counter for LIST item operation",
        }),
        items.ActionSave: promauto.NewCounter(prom.CounterOpts{
            Name: "items_save",
            Help: "Counter for SAVE item operation",
        }),
        items.ActionUpdate: promauto.NewCounter(prom.CounterOpts{
            Name: "items_update",
            Help: "Counter for UPDATE item operation",
        }),
        items.ActionDelete: promauto.NewCounter(prom.CounterOpts{
            Name: "items_delete",
            Help: "Counter for DELETE item operation",
        }),
    }
    return prometheus{
        counters: counters,
        logger:   logger,
    }
}

// NotifyMetric increment corresponding items' operation counter in Prometheus
func (prometheus prometheus) NotifyMetric(ctx context.Context, action items.Action) {
    prometheus.counters[action].Inc()
}
