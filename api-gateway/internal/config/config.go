package config

import (
	"github.com/BurntSushi/toml"
)

// Service holds one gateway backend definition.
type Service struct {
	Name string `toml:"name"`
	URL  string `toml:"url"`
}

// Config is the top‑level TOML structure.
type Config struct {
	Services []Service `toml:"service"`
}

// Load reads and parses the TOML file at path.
func Load(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// MustLoad panics on error.
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	return cfg
}
