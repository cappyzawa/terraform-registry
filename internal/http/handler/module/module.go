package module

import (
	"log"
	"os"

	"github.com/cappyzawa/terraform-registry/internal/config"
)

// Handler describes the handler for module
type Handler struct {
	Modules []config.Module
	Logger  *log.Logger
}

// Logger set logger to the Handler
func Logger(logger *log.Logger) func(*Handler) {
	return func(h *Handler) {
		h.Logger = logger
	}
}

// Modules set modules to the Handler
func Modules(modules []config.Module) func(*Handler) {
	return func(h *Handler) {
		h.Modules = modules
	}
}

func NewHandler(options ...func(*Handler)) *Handler {
	h := &Handler{
		Modules: []config.Module{},
		Logger:  log.New(os.Stderr, "", 0),
	}
	h.Logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	for _, option := range options {
		option(h)
	}
	return h
}
