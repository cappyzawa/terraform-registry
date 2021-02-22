package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/cappyzawa/terraform-registry/handler"
	"github.com/cappyzawa/terraform-registry/handler/module"
	"github.com/cappyzawa/terraform-registry/handler/provider"
	"github.com/go-chi/chi"
)

type cli struct {
	port       string
	configFile string
	pidFile    string
}

var (
	version = "dev"
)

func (c *cli) Run(args []string) {

	if len(args) != 0 && args[0] == "version" {
		print(version)
		return
	}

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
		if err := writePIDFile(c.pidFile); err != nil {
			log.Fatalf("Write PIDFILE(%s): %v", c.pidFile, err)
		}
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
		if err := deletePIDFile(c.pidFile); err != nil {
			log.Fatalf("Write PIDFILE(%s): %v", c.pidFile, err)
		}
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

func writePIDFile(path string) error {
	if path == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	pid := strconv.Itoa(os.Getpid())
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
		return err
	}
	log.Printf("Wrote PID file: %s", path)
	return nil
}

func deletePIDFile(path string) error {
	if path == "" {
		return nil
	}
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	log.Printf("Deleted PID file: %s", path)
	return nil
}

func main() {
	c := &cli{
		port:       os.Getenv("PORT"),
		configFile: os.Getenv("CONFIG_FILE"),
		pidFile:    os.Getenv("PID_FILE"),
	}
	c.Run(os.Args[1:])
}
