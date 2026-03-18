package main

import (
    "sysmc/config"
    "sysmc/internal/api"
    "sysmc/internal/collector"
    "sysmc/internal/logger"
)

func main() {
    cfg := config.Load()

    log := logger.New()

    coll := collector.NewSystemCollector()

    metricsHandler := api.NewMetricsHandler(coll, log)

    handler := api.SetupRoutes(metricsHandler, log)

    srv := api.NewServer(cfg, handler, log)
    if err := srv.Run(); err != nil {
        log.Error("Server failed:", err)
    }
}