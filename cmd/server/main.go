package main

import (
    "sysmc/config"
    "sysmc/internal/api"
    "sysmc/internal/collector"
    "sysmc/internal/logger"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Initialize logger
    log := logger.New()

    // Initialize collector
    coll := collector.NewSystemCollector()

    // Initialize metrics handler
    metricsHandler := api.NewMetricsHandler(coll, log)

    // Setup routes
    handler := api.SetupRoutes(metricsHandler, log)

    // Create and run server
    srv := api.NewServer(cfg, handler, log)
    if err := srv.Run(); err != nil {
        log.Error("Server failed:", err)
    }
}