package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/cappyzawa/terraform-registry/handler"
	"github.com/cappyzawa/terraform-registry/handler/provider"
	"github.com/go-chi/chi"
)

type cli struct {
	port       string
	configFile string
}

func (c *cli) Run() {
	if c.port == "" {
		c.port = "8080"
	}
	if c.configFile == "" {
		c.configFile = "./config.yaml"
	}

	config, err := config.Parse(c.configFile)
	if err != nil {
		log.Fatalf("parse config file: %v", err)
	}
	r := chi.NewRouter()
	registerRoute(r, config)

	server := http.Server{
		Addr:    ":" + c.port,
		Handler: r,
	}
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()
	select {
	case sig := <-sigCh:
		log.Printf("%s signal received server will shutdown soon...", sig)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}
}

func registerRoute(r *chi.Mux, c *config.Config) {
	wh := handler.WellKnwonHandler{}
	pvh := provider.VersionsHandler{
		Providers: c.Providers,
	}
	r.Get("/.wellknown/terraform.json", wh.ServeHTTP)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/providers", func(r chi.Router) {
			r.Get("/{namespace}/{type}/versions", pvh.ServeHTTP)
			r.Get("/{namespace}/{type}/{version}/download/{os}/{arch}", wh.ServeHTTP)
		})
		r.Route("/modules", func(r chi.Router) {
		})
	})
}

func main() {
	c := &cli{
		port:       os.Getenv("PORT"),
		configFile: os.Getenv("CONFIG_FILE"),
	}
	c.Run()
}
