package provider

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cappyzawa/terraform-registry/config"
	"github.com/go-chi/chi"
)

// VersionResoponse describes response for provider verions
type VersionResoponse struct {
	Versions []Version `json:"versions"`
}

// Version desribes provider version
type Version struct {
	Name      string            `json:"version"`
	Protocols []string          `json:"protocols"`
	Platforms []config.Platform `json:"platforms"`
}

// VersionsHandler handles request for provider versions
type VersionsHandler struct {
	Providers []config.Provider
}

func (h *VersionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	_type := chi.URLParam(r, "type")
	var res VersionResoponse
	var versions []Version
	var exist bool
	for _, provider := range h.Providers {
		if provider.Namespace == namespace && _type == provider.Type {
			exist = true
			for _, v := range provider.Versions {
				version := Version{
					Name:      v.Name,
					Protocols: []string{"5.3"},
					Platforms: v.Platforms,
				}
				versions = append(versions, version)
			}
		}
	}
	if !exist {
		log.Printf("provider: %s/%s does not found", namespace, _type)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res.Versions = versions
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
