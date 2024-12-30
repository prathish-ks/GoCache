package main

import (
    "log"
    "net/http"
    "gocache/internal/config"
    "gocache/internal/server"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    srv := server.NewServer(cfg)
    log.Printf("Starting server on %s", cfg.ServerAddress)

    if err := http.ListenAndServe(cfg.ServerAddress, srv.Router); err != nil {
        log.Fatalf("could not start server: %v", err)
    }
}