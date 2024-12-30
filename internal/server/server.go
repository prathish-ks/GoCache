package server

import (
    "gocache/internal/cache"
    "gocache/internal/config"
    "net/http"
)

type Server struct {
    cache *cache.MemoryCache
}

func New(cfg *config.Config) *http.Server {
    cache := cache.NewMemoryCache()
    srv := &Server{cache: cache}

    mux := http.NewServeMux()
    mux.HandleFunc("/set", srv.handleSet)
    mux.HandleFunc("/get", srv.handleGet)
    mux.HandleFunc("/delete", srv.handleDelete)

    return &http.Server{
        Addr:    ":" + cfg.Port,
        Handler: mux,
    }
}

func (s *Server) handleSet(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value := r.URL.Query().Get("value")
    err := s.cache.Set(key, value)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value, err := s.cache.Get(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Write([]byte(value.(string)))
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    err := s.cache.Delete(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}