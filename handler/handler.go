package handler

import (
	"encoding/json"
	"net/http"
)

// WellKnwonResponse describes response for wellknwon request
type WellKnwonResponse struct {
	Providers string `json:"providers.v1"`
	Modules   string `json:"modules.v1"`
}

// WellKnwonHandler handles requests for wellknwon
type WellKnwonHandler struct {
}

func (h *WellKnwonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res WellKnwonResponse
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
