package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cappyzawa/terraform-registry/internal/config"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/module"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/provider"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/wellknown"
	"github.com/go-chi/chi"
)

// Server describes the http server
type Server struct {
	server *http.Server
}

// NewServer initializes the http server
func NewServer(port string, c *config.Config, logger *log.Logger) *Server {
	r := chi.NewRouter()
	registerRoute(r, c, logger)

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: r,
		},
	}
}

func registerRoute(r *chi.Mux, c *config.Config, logger *log.Logger) {
	wh := wellknown.NewHandler(wellknown.Logger(logger))
	ph := provider.NewHandler(
		provider.Providers(c.Providers),
		provider.Logger(logger),
	)
	mh := module.NewHandler(
		module.Modules(c.Modules),
		module.Logger(logger),
	)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Get("/.well-known/terraform.json", wh.WellKnown)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/providers", func(r chi.Router) {
			r.Get("/{namespace}/{type}/versions", ph.Versions)
			r.Get("/{namespace}/{type}/{version}/download/{os}/{arch}", ph.Download)
		})
		r.Route("/modules", func(r chi.Router) {
			r.Get("/{namespace}/{name}/{provider}/versions", mh.Versions)
		})
	})
}

// Start starts the http server
func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http server ListenAndServe: %v", err)
	}

	return nil
}

// Stop stops the http server
func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("http server Shutdown: %v", err)
	}
	return nil
}
