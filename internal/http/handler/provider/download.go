package provider

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/cappyzawa/terraform-registry/internal/config"
	"github.com/go-chi/chi"
)

// DownloadResponse desribes response for downloading provider
type DownloadResponse struct {
	Protocols []string `json:"protocols"`
	OS        string   `json:"os"`
	Arch      string   `json:"arch"`
	*DownloadMeta
}

// DownloadQuery desribes query for finding provider
type DownloadQuery struct {
	Namespace string
	Type      string
	Version   string
	OS        string
	Arch      string
}

// DownloadMeta desribes metadata for downloding
type DownloadMeta struct {
	Filename            string             `json:"filename"`
	DownloadURL         string             `json:"download_url"`
	ShasumsURL          string             `json:"shasums_url"`
	ShasumsSignatureURL string             `json:"shasums_signature_url"`
	Shasum              string             `json:"shasum"`
	SigningKeys         config.SigningKeys `json:"signing_keys"`
}

// DownloadHandler handles requests for downloading provider
type DownloadHandler struct {
	Providers []config.Provider
	Logger    *log.Logger
}

func (h *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := DownloadQuery{
		Namespace: chi.URLParam(r, "namespace"),
		Type:      chi.URLParam(r, "type"),
		Version:   chi.URLParam(r, "version"),
		OS:        chi.URLParam(r, "os"),
		Arch:      chi.URLParam(r, "arch"),
	}

	meta, ok := getDownloadMeta(h.Providers, &q)
	if !ok {
		h.Logger.Printf("provider: %s/%s does not exist", q.Namespace, q.Type)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res := DownloadResponse{
		Protocols:    []string{"5.3"},
		OS:           q.OS,
		Arch:         q.Arch,
		DownloadMeta: meta,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
	return
}

func replaceMeta(base string, query *DownloadQuery) string {
	replacedNS := strings.ReplaceAll(base, "{namespace}", query.Namespace)
	replacedType := strings.ReplaceAll(replacedNS, "{type}", query.Type)
	replacedVersion := strings.ReplaceAll(replacedType, "{version}", query.Version)
	replacedOS := strings.ReplaceAll(replacedVersion, "{os}", query.OS)
	return strings.ReplaceAll(replacedOS, "{arch}", query.Arch)
}

func getDownloadMeta(providers []config.Provider, query *DownloadQuery) (*DownloadMeta, bool) {
	for _, provider := range providers {
		if provider.Namespace == query.Namespace && provider.Type == query.Type {
			downloadURL := replaceMeta(provider.DownloadURLFmt, query)
			filename := filepath.Base(downloadURL)
			shasumURL := replaceMeta(provider.ShasumsURLFmt, query)
			shasumSigURL := replaceMeta(provider.ShasumsSignatureURLFmt, query)
			for _, v := range provider.Versions {
				if v.Name == query.Version {
					for _, a := range v.Assets {
						if a.OS == query.OS && a.Arch == query.Arch {
							return &DownloadMeta{
								DownloadURL:         downloadURL,
								Filename:            filename,
								ShasumsURL:          shasumURL,
								ShasumsSignatureURL: shasumSigURL,
								Shasum:              a.Shasum,
								SigningKeys:         provider.SigningKeys,
							}, true
						}
					}
				}
			}
		}
	}
	return nil, false
}
