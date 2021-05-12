package provider

import (
	"log"
	"os"

	"github.com/cappyzawa/terraform-registry/internal/config"
)

// Handler describes the handler for provider
type Handler struct {
	Providers []config.Provider
	Logger    *log.Logger
}

// Logger set logger to the Handler
func Logger(logger *log.Logger) func(*Handler) {
	return func(h *Handler) {
		h.Logger = logger
	}
}

// Providers set providers to the Handler
func Providers(providers []config.Provider) func(*Handler) {
	return func(h *Handler) {
		h.Providers = providers
	}
}

// NewHandler initializes the handler for providers
func NewHandler(options ...func(*Handler)) *Handler {
	h := &Handler{
		Providers: []config.Provider{},
		Logger:    log.New(os.Stderr, "", 0),
	}
	h.Logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	for _, option := range options {
		option(h)
	}
	return h
}
