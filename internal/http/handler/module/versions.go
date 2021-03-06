package module

import (
	"net/http"
)

// VersionsResponse describes response for versions of target module
type VersionsResponse struct {
	Modules []Module `json:"modules"`
}

// Module describes module
type Module struct {
	Source   string    `json:"source"`
	Versions []Version `json:"versions"`
}

// Version describes module version
type Version struct {
	Name       string `json:"version"`
	SubModules []Info `json:"submodules"`
	Root       Info   `json:"root"`
}

// Info describes submodule of module
type Info struct {
	Path         string     `json:"path"`
	Providers    []Provider `json:"providers"`
	Dependencies []string   `json:"dependencies"`
}

// Provider describes provider used in module
type Provider struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Versions handles requests for module versions
func (h *Handler) Versions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("not implemention"))
}
