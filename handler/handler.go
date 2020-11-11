package handler

import (
	"encoding/json"
	"net/http"
)

// WellKnownResponse describes response for wellknwon request
type WellKnownResponse struct {
	Providers string `json:"providers.v1"`
	Modules   string `json:"modules.v1"`
}

// WellKnownHandler handles requests for wellknwon
type WellKnownHandler struct {
}

func (h *WellKnownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := WellKnownResponse{
		Providers: "/v1/providers/",
		Modules:   "/v1/modules/",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
