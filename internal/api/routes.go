package api

import (
    "net/http"

    "sysmc/internal/logger"
)

func SetupRoutes(metricsHandler *MetricsHandler, logger *logger.Logger) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /api/metrics", metricsHandler.GetMetrics)

    fileServer := http.FileServer(http.Dir("./web/static"))
    mux.Handle("/", fileServer)

    handler := loggingMiddleware(logger, mux)
    handler = recoveryMiddleware(logger, handler)

    return handler
}

func loggingMiddleware(logger *logger.Logger, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logger.Info(r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func recoveryMiddleware(logger *logger.Logger, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                logger.Error("panic recovered:", err)
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}