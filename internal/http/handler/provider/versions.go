package provider

import (
	"encoding/json"
	"net/http"

	p "github.com/cappyzawa/terraform-registry/provider"
	"github.com/go-chi/chi"
)

type (
	VersionResoponse = p.Versions
	Version          = p.Version
	Platform         = p.Platform
)

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
