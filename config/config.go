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
	Namespace string    `yaml:"namespace"`
	Type      string    `yaml:"type"`
	Versions  []Version `yaml:"versions"`
}

// Version desribes provider version
type Version struct {
	Name      string     `yaml:"name"`
	Platforms []Platform `yaml:"platforms"`
	Source    Source     `yaml:"source"`
}

// Platform describes available platform for provider
type Platform struct {
	OS   string `yaml:"os" json:"os"`
	Arch string `yaml:"arch" json:"arch"`
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
