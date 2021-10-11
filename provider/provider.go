package provider

import "github.com/cappyzawa/terraform-registry/config"

// Versions describes response for provider verions
type Versions struct {
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

// Download desribes response for downloading provider
type Download struct {
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
