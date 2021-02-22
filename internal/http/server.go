package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cappyzawa/terraform-registry/internal/config"
	"github.com/cappyzawa/terraform-registry/internal/http/handler"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/module"
	"github.com/cappyzawa/terraform-registry/internal/http/handler/provider"
	"github.com/go-chi/chi"
)

// Server describes the http server
type Server struct {
	server *http.Server
}

// NewServer initializes the http server
func NewServer(port int, c *config.Config) *Server {
	r := chi.NewRouter()
	registerRoute(r, c)

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
		},
	}
}

func registerRoute(r *chi.Mux, c *config.Config) {
	wh := handler.WellKnownHandler{}
	pvh := provider.VersionsHandler{
		Providers: c.Providers,
	}
	pdh := provider.DownloadHandler{
		Providers: c.Providers,
	}
	mvh := module.VersionsHandler{
		Modules: c.Modules,
	}
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Get("/.well-known/terraform.json", wh.ServeHTTP)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/providers", func(r chi.Router) {
			r.Get("/{namespace}/{type}/versions", pvh.ServeHTTP)
			r.Get("/{namespace}/{type}/{version}/download/{os}/{arch}", pdh.ServeHTTP)
		})
		r.Route("/modules", func(r chi.Router) {
			r.Get("/{namespace}/{name}/{provider}/versions", mvh.ServeHTTP)
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
