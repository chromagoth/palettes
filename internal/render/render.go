package render

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"gitlab.com/patrick.pfenning.92/chromagoth/palettes/internal/palette"
)

// ParseHex parses a #rrggbb hex string into r, g, b components.
func ParseHex(h string) (r, g, b uint8, err error) {
	h = strings.TrimPrefix(strings.ToLower(strings.TrimSpace(h)), "#")
	if len(h) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", h)
	}
	rv, err1 := strconv.ParseUint(h[0:2], 16, 8)
	gv, err2 := strconv.ParseUint(h[2:4], 16, 8)
	bv, err3 := strconv.ParseUint(h[4:6], 16, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %q", h)
	}
	return uint8(rv), uint8(gv), uint8(bv), nil
}

func toHSL(r, g, b uint8) (h, s, l float64) {
	rf, gf, bf := float64(r)/255, float64(g)/255, float64(b)/255
	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	l = (max + min) / 2
	if max == min {
		return 0, 0, l
	}
	d := max - min
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}
	switch max {
	case rf:
		h = (gf-bf)/d + map[bool]float64{true: 6}[gf < bf]
	case gf:
		h = (bf-rf)/d + 2
	case bf:
		h = (rf-gf)/d + 4
	}
	h /= 6
	return h * 360, s * 100, l * 100
}

// FuncMap returns the template helper functions available in port templates.
func FuncMap() template.FuncMap {
	mustParse := func(h string) (uint8, uint8, uint8) {
		r, g, b, err := ParseHex(h)
		if err != nil {
			return 0, 0, 0
		}
		return r, g, b
	}
	return template.FuncMap{
		"hex": func(h string) string {
			h = strings.TrimSpace(h)
			if !strings.HasPrefix(h, "#") {
				h = "#" + h
			}
			return strings.ToLower(h)
		},
		"rgb": func(h string) string {
			r, g, b := mustParse(h)
			return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		},
		"hsl": func(h string) string {
			r, g, b := mustParse(h)
			hv, s, l := toHSL(r, g, b)
			return fmt.Sprintf("hsl(%d, %d%%, %d%%)", int(math.Round(hv)), int(math.Round(s)), int(math.Round(l)))
		},
		"alpha": func(h string, a float64) string {
			r, g, b := mustParse(h)
			return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", r, g, b, a)
		},
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"toJSON": func(v any) (string, error) {
			b, err := json.MarshalIndent(v, "", "  ")
			return string(b), err
		},
	}
}

// RenderConfig describes one render job (mirrors chromagoth.toml [[render]] entry).
type RenderConfig struct {
	Template   string `toml:"template"`
	Output     string `toml:"output"`
	PerVariant bool   `toml:"per_variant"`
}

// File renders a single template file and writes output.
// If rc.PerVariant is true, one output file is written per palette variant;
// {{ .variant }} in rc.Output is substituted with the variant slug.
// Otherwise the template receives all palettes under {{ .palettes }}.
func File(rc RenderConfig, palettes []palette.Palette) error {
	tmplBytes, err := os.ReadFile(rc.Template)
	if err != nil {
		return fmt.Errorf("reading template %s: %w", rc.Template, err)
	}

	tmpl, err := template.New(filepath.Base(rc.Template)).Funcs(FuncMap()).Parse(string(tmplBytes))
	if err != nil {
		return fmt.Errorf("parsing template %s: %w", rc.Template, err)
	}

	if rc.PerVariant {
		return renderPerVariant(tmpl, rc.Output, palettes)
	}
	return renderAll(tmpl, rc.Output, palettes)
}

func renderPerVariant(tmpl *template.Template, outputPattern string, palettes []palette.Palette) error {
	for _, p := range palettes {
		ctx := p.TemplateContext()
		outPath := strings.ReplaceAll(outputPattern, "{{ .variant }}", p.Variant)
		if err := writeTemplate(tmpl, outPath, ctx); err != nil {
			return err
		}
	}
	return nil
}

func renderAll(tmpl *template.Template, outputPath string, palettes []palette.Palette) error {
	contexts := make([]map[string]any, len(palettes))
	for i, p := range palettes {
		contexts[i] = p.TemplateContext()
	}
	ctx := map[string]any{
		"palettes": contexts,
		"raw":      palettes,
		"slots":    palette.Slots,
	}
	return writeTemplate(tmpl, outputPath, ctx)
}

func writeTemplate(tmpl *template.Template, outPath string, ctx any) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return fmt.Errorf("creating output dir for %s: %w", outPath, err)
	}
	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("creating output file %s: %w", outPath, err)
	}
	defer f.Close()
	if err := tmpl.Execute(f, ctx); err != nil {
		return fmt.Errorf("executing template → %s: %w", outPath, err)
	}
	fmt.Printf("  ✓ %s\n", outPath)
	return nil
}
