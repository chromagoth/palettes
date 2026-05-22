package palette

import (
	"fmt"
	"strings"

	palettedata "gitlab.com/chromagoth/palettes/src"
	"gopkg.in/yaml.v3"
)

type Palette struct {
	Name    string            `yaml:"name"    json:"name"`
	Variant string            `yaml:"variant" json:"variant"`
	Mascot  string            `yaml:"mascot"  json:"mascot"`
	Style   string            `yaml:"style"   json:"style"`
	Vibe    string            `yaml:"vibe"    json:"vibe"`
	Dark    bool              `yaml:"dark"    json:"dark"`
	Status  string            `yaml:"status"  json:"status"`
	Colors  map[string]string `yaml:"colors"  json:"colors"`
}

// Slots defines the canonical slot order.
var Slots = []string{
	"ground", "veil", "field", "trace",
	"ash", "mist", "haze", "graphite",
	"circuit-lime", "powder-blush", "static-mint", "laser-blue",
	"cyber-pink", "ultraviolet", "amber-glow", "cherry-flux",
}

// LoadAll reads all non-WIP palettes from the embedded FS.
func LoadAll() ([]Palette, error) {
	entries, err := palettedata.FS.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("reading embedded palettes: %w", err)
	}

	var palettes []Palette
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".yaml") || strings.Contains(name, "-wip") {
			continue
		}
		data, err := palettedata.FS.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", name, err)
		}
		var p Palette
		if err := yaml.Unmarshal(data, &p); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", name, err)
		}
		palettes = append(palettes, p)
	}
	return palettes, nil
}

// TemplateContext builds the map passed to port templates for a single palette.
// Supports dot-notation access: {{ .surface.base }}, {{ .accent.primary }}, etc.
func (p Palette) TemplateContext() map[string]any {
	surface := map[string]string{
		"base":    p.Colors["ground"],
		"raised":  p.Colors["veil"],
		"overlay": p.Colors["field"],
		"border":  p.Colors["trace"],
	}
	text := map[string]string{
		"disabled": p.Colors["ash"],
		"subtle":   p.Colors["mist"],
		"dim":      p.Colors["haze"],
		"default":  p.Colors["graphite"],
	}
	accent := map[string]string{
		"primary":   p.Colors["cyber-pink"],
		"secondary": p.Colors["ultraviolet"],
		"link":      p.Colors["laser-blue"],
		"positive":  p.Colors["circuit-lime"],
		"info":      p.Colors["static-mint"],
		"caution":   p.Colors["powder-blush"],
		"warning":   p.Colors["amber-glow"],
		"danger":    p.Colors["cherry-flux"],
	}
	return map[string]any{
		"variant": p.Variant,
		"name":    p.Name,
		"mascot":  p.Mascot,
		"style":   p.Style,
		"vibe":    p.Vibe,
		"dark":    p.Dark,
		"status":  p.Status,
		"surface": surface,
		"text":    text,
		"accent":  accent,
		"colors":  p.Colors,
	}
}
