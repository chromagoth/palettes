package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"gitlab.com/patrick.pfenning.92/chromagoth/palettes/internal/render"
)

const DefaultFile = "chromagoth.toml"

// Port describes the port-level metadata in chromagoth.toml.
type Port struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
}

// Config is the full chromagoth.toml structure used in port repos.
type Config struct {
	Port   Port                  `toml:"port"`
	Render []render.RenderConfig `toml:"render"`
}

// Load reads a chromagoth.toml file. Path defaults to DefaultFile if empty.
func Load(path string) (*Config, error) {
	if path == "" {
		path = DefaultFile
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}
	var cfg Config
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}
	return &cfg, nil
}
