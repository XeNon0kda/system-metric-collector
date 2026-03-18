package api

import (
    "encoding/json"
    "net/http"

    "sysmc/internal/collector"
    "sysmc/internal/logger"
)

type MetricsHandler struct {
    collector collector.Collector
    logger    *logger.Logger
}

func NewMetricsHandler(c collector.Collector, l *logger.Logger) *MetricsHandler {
    return &MetricsHandler{
        collector: c,
        logger:    l,
    }
}

func (h *MetricsHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {
    metrics, err := h.collector.Collect(r.Context())
    if err != nil {
        h.logger.Error("failed to collect metrics:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(metrics); err != nil {
        h.logger.Error("failed to encode metrics:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}