package render

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gitlab.com/patrick.pfenning.92/chromagoth/palettes/internal/palette"
)

func TestParseHex(t *testing.T) {
	cases := []struct {
		input   string
		r, g, b uint8
		wantErr bool
	}{
		{"#ff4fa8", 255, 79, 168, false},
		{"#6a6c70", 106, 108, 112, false},
		{"#0b0d14", 11, 13, 20, false},
		{"FF4FA8", 255, 79, 168, false},
		{"#gggggg", 0, 0, 0, true},
		{"#fff", 0, 0, 0, true},
	}
	for _, c := range cases {
		r, g, b, err := ParseHex(c.input)
		if c.wantErr {
			if err == nil {
				t.Errorf("ParseHex(%q): expected error", c.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseHex(%q): %v", c.input, err)
			continue
		}
		if r != c.r || g != c.g || b != c.b {
			t.Errorf("ParseHex(%q) = (%d,%d,%d), want (%d,%d,%d)", c.input, r, g, b, c.r, c.g, c.b)
		}
	}
}

func TestFuncMap_rgb(t *testing.T) {
	fn := FuncMap()["rgb"].(func(string) string)
	if got := fn("#ff4fa8"); got != "rgb(255, 79, 168)" {
		t.Errorf("got %q, want %q", got, "rgb(255, 79, 168)")
	}
}

func TestFuncMap_hsl(t *testing.T) {
	fn := FuncMap()["hsl"].(func(string) string)
	got := fn("#6a6c70")
	if !strings.HasPrefix(got, "hsl(") || !strings.HasSuffix(got, ")") {
		t.Errorf("hsl(#6a6c70) = %q, expected hsl(...) format", got)
	}
}

func TestFuncMap_alpha(t *testing.T) {
	fn := FuncMap()["alpha"].(func(string, float64) string)
	if got := fn("#ff4fa8", 0.5); got != "rgba(255, 79, 168, 0.50)" {
		t.Errorf("got %q, want %q", got, "rgba(255, 79, 168, 0.50)")
	}
}

func TestFuncMap_hex(t *testing.T) {
	fn := FuncMap()["hex"].(func(string) string)
	cases := []struct{ in, want string }{
		{"#FF4FA8", "#ff4fa8"},
		{"FF4FA8", "#ff4fa8"},
		{"  #ff4fa8  ", "#ff4fa8"},
	}
	for _, c := range cases {
		if got := fn(c.in); got != c.want {
			t.Errorf("hex(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFile_PerVariant(t *testing.T) {
	palettes, err := palette.LoadAll()
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}

	dir := t.TempDir()
	tmplPath := filepath.Join(dir, "test.tmpl")
	if err := os.WriteFile(tmplPath, []byte(`/* {{ .variant }} {{ hex .surface.base }} */`), 0o644); err != nil {
		t.Fatal(err)
	}

	rc := RenderConfig{
		Template:   tmplPath,
		Output:     filepath.Join(dir, "chromagoth-{{ .variant }}.css"),
		PerVariant: true,
	}
	if err := File(rc, palettes); err != nil {
		t.Fatalf("File: %v", err)
	}

	for _, p := range palettes {
		out := filepath.Join(dir, "chromagoth-"+p.Variant+".css")
		data, err := os.ReadFile(out)
		if err != nil {
			t.Errorf("missing output for %q: %v", p.Variant, err)
			continue
		}
		if !strings.Contains(string(data), p.Variant) {
			t.Errorf("output for %q missing variant name", p.Variant)
		}
		if !strings.Contains(string(data), "#") {
			t.Errorf("output for %q missing hex color", p.Variant)
		}
	}
}

func TestFile_All(t *testing.T) {
	palettes, err := palette.LoadAll()
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}

	dir := t.TempDir()
	tmplPath := filepath.Join(dir, "all.tmpl")
	if err := os.WriteFile(tmplPath, []byte(`{{ range .palettes }}/* {{ .variant }} */{{ end }}`), 0o644); err != nil {
		t.Fatal(err)
	}

	rc := RenderConfig{
		Template:   tmplPath,
		Output:     filepath.Join(dir, "all.css"),
		PerVariant: false,
	}
	if err := File(rc, palettes); err != nil {
		t.Fatalf("File: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "all.css"))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range palettes {
		if !strings.Contains(string(data), p.Variant) {
			t.Errorf("all.css missing variant %q", p.Variant)
		}
	}
}
