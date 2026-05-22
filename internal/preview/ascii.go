package preview

import (
	"fmt"
	"strings"

	"gitlab.com/patrick.pfenning.92/chromagoth/palettes/internal/palette"
	"gitlab.com/patrick.pfenning.92/chromagoth/palettes/internal/render"
)

const (
	reset     = "\033[0m"
	bold      = "\033[1m"
	fgBlack   = "\033[38;2;0;0;0m"
	fgWhite   = "\033[38;2;255;255;255m"
)

func bgColor(r, g, b uint8) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func fgFor(r, g, b uint8) string {
	// Perceived luminance — pick black or white foreground for swatch label
	lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	if lum > 140 {
		return fgBlack
	}
	return fgWhite
}

func swatch(hex string) string {
	r, g, b, err := render.ParseHex(hex)
	if err != nil {
		return "    "
	}
	fg := fgFor(r, g, b)
	return bgColor(r, g, b) + fg + "    " + reset
}

// PrintAll renders a truecolor swatch table for every palette.
func PrintAll(palettes []palette.Palette) {
	for i, p := range palettes {
		if i > 0 {
			fmt.Println()
		}
		PrintOne(p)
	}
}

// PrintOne renders a truecolor swatch table for a single palette.
func PrintOne(p palette.Palette) {
	mode := "light"
	if p.Dark {
		mode = "dark"
	}
	fmt.Printf("%s%s%s  %s · %s\n", bold, p.Name, reset, p.Style, mode)
	fmt.Println(strings.Repeat("─", 42))

	base := palette.Slots[:8]
	accents := palette.Slots[8:]

	for _, slot := range base {
		hex := p.Colors[slot]
		fmt.Printf("  %-14s %s  %s\n", slot, swatch(hex), hex)
	}
	fmt.Println()
	for _, slot := range accents {
		hex := p.Colors[slot]
		fmt.Printf("  %-14s %s  %s\n", slot, swatch(hex), hex)
	}
}
