package wellknown

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Handler describes the handler
type Handler struct {
	Logger *log.Logger
}

// New initilize the handler
func NewHandler(options ...func(*Handler)) *Handler {
	h := &Handler{
		Logger: log.New(os.Stderr, "", 0),
	}
	h.Logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	for _, option := range options {
		option(h)
	}
	return h
}

// Logger set logger to the Handler
func Logger(logger *log.Logger) func(*Handler) {
	return func(h *Handler) {
		h.Logger = logger
	}
}

// WellKnownResponse describes response for wellknwon request
type WellKnownResponse struct {
	Providers string `json:"providers.v1"`
	Modules   string `json:"modules.v1"`
}

// WellKnown handles requests for wellknown endpoint
func (h *Handler) WellKnown(w http.ResponseWriter, r *http.Request) {
	res := WellKnownResponse{
		Providers: "/v1/providers/",
		Modules:   "/v1/modules/",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
