package main

import (
    "log"
    "gocache/internal/config"
    "gocache/internal/server"
)

func main() {
    cfg := config.LoadConfig()

    srv := server.New(cfg)
    log.Printf("Starting server on %s", cfg.Port)

    if err := srv.ListenAndServe(); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}