package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config desribes configuration for registry
type Config struct {
	Providers []Provider `yaml:"providers"`
}

// Provider desribes config for provider
type Provider struct {
	Namespace              string      `yaml:"namespace"`
	Type                   string      `yaml:"type"`
	Versions               []Version   `yaml:"versions"`
	DownloadURLFmt         string      `yaml:"download_url_fmt"`
	ShasumsURLFmt          string      `yaml:"shasums_url_fmt"`
	ShasumsSignatureURLFmt string      `yaml:"shasums_signature_url_fmt"`
	SigningKeys            SigningKeys `yaml:"signing_keys"`
}

// SigningKeys desribes signingkeys
type SigningKeys struct {
	GpgPublicKeys []GpgPublicKey `yaml:"gpg_public_keys" json:"gpg_public_keys"`
}

// GpgPublicKey desribes gpg public key
type GpgPublicKey struct {
	KeyID          string `yaml:"key_id" json:"key_id"`
	ASCIIArmor     string `yaml:"ascii_armor" json:"ascii_armor"`
	TrustSignature string `yaml:"trust_signature" json:"trust_signature"`
	Source         string `yaml:"source" json:"source"`
	SourceURL      string `yaml:"source_url" json:"source_url"`
}

// Version desribes provider version
type Version struct {
	Name   string  `yaml:"name"`
	Assets []Asset `yaml:"assets"`
	Source Source  `yaml:"source"`
}

// Asset describes available platform for provider
type Asset struct {
	OS     string `yaml:"os" json:"os"`
	Arch   string `yaml:"arch" json:"arch"`
	Shasum string `yaml:"shasum" json:"shasum"`
}

// Source desribes source for provider
type Source struct {
	DownloadURL string `yaml:"download_url"`
}

// Parse parse yaml file to go struct
func Parse(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var c Config
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
