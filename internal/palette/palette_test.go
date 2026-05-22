package palette

import (
	"testing"
)

func TestLoadAll(t *testing.T) {
	palettes, err := LoadAll()
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	if len(palettes) == 0 {
		t.Fatal("expected palettes, got none")
	}
	for _, p := range palettes {
		if p.Variant == "" {
			t.Errorf("palette %q: empty variant", p.Name)
		}
		if len(p.Colors) != 16 {
			t.Errorf("palette %q: expected 16 color slots, got %d", p.Variant, len(p.Colors))
		}
		ash, ok := p.Colors["ash"]
		if !ok {
			t.Errorf("palette %q: missing ash slot", p.Variant)
		} else if ash != "#6A6C70" && ash != "#6a6c70" {
			t.Errorf("palette %q: ash = %q, want #6A6C70", p.Variant, ash)
		}
	}
}

func TestLoadAll_NoWIP(t *testing.T) {
	palettes, err := LoadAll()
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	for _, p := range palettes {
		if p.Status == "wip" {
			t.Errorf("palette %q: WIP palette should not be loaded", p.Variant)
		}
	}
}

func TestTemplateContext(t *testing.T) {
	palettes, err := LoadAll()
	if err != nil {
		t.Fatalf("LoadAll: %v", err)
	}
	p := palettes[0]
	ctx := p.TemplateContext()

	for _, key := range []string{"variant", "name", "dark", "surface", "text", "accent", "colors"} {
		if _, ok := ctx[key]; !ok {
			t.Errorf("TemplateContext: missing key %q", key)
		}
	}

	surface, ok := ctx["surface"].(map[string]string)
	if !ok {
		t.Fatal("TemplateContext: surface is not map[string]string")
	}
	for _, role := range []string{"base", "raised", "overlay", "border"} {
		if v := surface[role]; v == "" {
			t.Errorf("TemplateContext: surface.%s is empty", role)
		}
	}

	accent, ok := ctx["accent"].(map[string]string)
	if !ok {
		t.Fatal("TemplateContext: accent is not map[string]string")
	}
	for _, role := range []string{"primary", "secondary", "link", "positive", "info", "caution", "warning", "danger"} {
		if v := accent[role]; v == "" {
			t.Errorf("TemplateContext: accent.%s is empty", role)
		}
	}
}

func TestSlots(t *testing.T) {
	if len(Slots) != 16 {
		t.Errorf("expected 16 slots, got %d", len(Slots))
	}
	if Slots[4] != "ash" {
		t.Errorf("slot[4] = %q, want ash", Slots[4])
	}
}
