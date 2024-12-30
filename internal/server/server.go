package server

import (
    "encoding/json"
    "net/http"
    "sync"

    "gocache/internal/cache"
)

type Server struct {
    Cache cache.Cache
    mu    sync.Mutex
}

func NewServer(c cache.Cache) *Server {
    return &Server{Cache: c}
}

func (s *Server) Start(addr string) error {
    http.HandleFunc("/cache/get", s.handleGet)
    http.HandleFunc("/cache/set", s.handleSet)
    http.HandleFunc("/cache/delete", s.handleDelete)

    return http.ListenAndServe(addr, nil)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value, err := s.Cache.Get(key)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(value)
}

func (s *Server) handleSet(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Key   string `json:"key"`
        Value string `json:"value"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    s.Cache.Set(req.Key, req.Value)
    w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    s.Cache.Delete(key)
    w.WriteHeader(http.StatusNoContent)
}