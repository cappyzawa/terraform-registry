package provider

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// VersionResoponse describes response for provider verions
type VersionResoponse struct {
	Versions []Version `json:"versions"`
}

// Version desribes provider version
type Version struct {
	Name      string     `json:"version"`
	Protocols []string   `json:"protocols"`
	Platforms []Platform `json:"platforms"`
}

// Platform describes provider platform
type Platform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

// Versions handles requests for provider versions
func (h *Handler) Versions(w http.ResponseWriter, r *http.Request) {
	namespace := chi.URLParam(r, "namespace")
	_type := chi.URLParam(r, "type")
	var res VersionResoponse
	var versions []Version
	var exist bool
	for _, provider := range h.Providers {
		if provider.Namespace == namespace && _type == provider.Type {
			exist = true
			for _, v := range provider.Versions {
				var pfs []Platform
				for _, asset := range v.Assets {
					pf := Platform{
						OS:   asset.OS,
						Arch: asset.Arch,
					}
					pfs = append(pfs, pf)
				}
				version := Version{
					Name:      v.Name,
					Protocols: []string{"5.3"},
					Platforms: pfs,
				}
				versions = append(versions, version)
			}
		}
	}
	if !exist {
		h.Logger.Printf("provider: %s/%s does not found", namespace, _type)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res.Versions = versions
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
