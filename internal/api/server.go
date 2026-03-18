package api

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "sysmc/config"
    "sysmc/internal/logger"
)

type Server struct {
    httpServer *http.Server
    logger     *logger.Logger
    config     *config.Config
}

func NewServer(cfg *config.Config, handler http.Handler, logger *logger.Logger) *Server {
    return &Server{
        httpServer: &http.Server{
            Addr:         cfg.ServerPort,
            Handler:      handler,
            ReadTimeout:  cfg.ReadTimeout,
            WriteTimeout: cfg.WriteTimeout,
        },
        logger: logger,
        config: cfg,
    }
}

func (s *Server) Run() error {
    // Graceful shutdown channel
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        s.logger.Info("Server started on", s.config.ServerPort)
        if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            s.logger.Error("Server error:", err)
        }
    }()

    <-quit
    s.logger.Info("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := s.httpServer.Shutdown(ctx); err != nil {
        s.logger.Error("Server forced to shutdown:", err)
        return err
    }

    s.logger.Info("Server exited gracefully")
    return nil
}