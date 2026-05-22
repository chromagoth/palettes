package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gitlab.com/chromagoth/palettes/internal/config"
	"gitlab.com/chromagoth/palettes/internal/palette"
	"gitlab.com/chromagoth/palettes/internal/preview"
	"gitlab.com/chromagoth/palettes/internal/render"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "chromagoth",
	Short: "Chromagoth theme CLI — build, render, preview",
}

// ── build ─────────────────────────────────────────────────────────────────

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Regenerate dist/ outputs (CSS, SCSS, LESS, JS, TS)",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: migrate generate.js outputs here
		fmt.Println("build: not yet implemented — run pnpm build for now")
		return nil
	},
}

// ── render ────────────────────────────────────────────────────────────────

var renderCmd = &cobra.Command{
	Use:   "render [template]",
	Short: "Render port templates against all palette variants",
	Long: `Render one or more Go templates against the embedded palette data.

With no argument, reads chromagoth.toml in the current directory and
renders all [[render]] entries defined there.

With a template argument, renders that single file. Output path is
derived from the template name (strip .tmpl) unless -o is set.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runRender,
}

var renderOutput string

func init() {
	renderCmd.Flags().StringVarP(&renderOutput, "output", "o", "", "output file path (single template only)")
	renderCmd.Flags().Bool("per-variant", false, "render one output file per palette variant")

	// Bind flags → viper so env vars also work:
	// CHROMAGOTH_RENDER_OUTPUT, CHROMAGOTH_RENDER_PER_VARIANT
	_ = viper.BindPFlag("render.output", renderCmd.Flags().Lookup("output"))
	_ = viper.BindPFlag("render.per_variant", renderCmd.Flags().Lookup("per-variant"))
}

func runRender(cmd *cobra.Command, args []string) error {
	palettes, err := palette.LoadAll()
	if err != nil {
		return fmt.Errorf("loading palettes: %w", err)
	}

	// Explicit template argument
	if len(args) == 1 {
		out := viper.GetString("render.output")
		if out == "" {
			// Strip .tmpl suffix as default output path
			out = args[0]
			if len(out) > 5 && out[len(out)-5:] == ".tmpl" {
				out = out[:len(out)-5]
			}
		}
		rc := render.RenderConfig{
			Template:   args[0],
			Output:     out,
			PerVariant: viper.GetBool("render.per_variant"),
		}
		return render.File(rc, palettes)
	}

	// No argument — read chromagoth.toml
	cfgPath := viper.GetString("config")
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return fmt.Errorf("no template given and %w", err)
	}

	fmt.Printf("port: %s\n", cfg.Port.Name)
	for _, rc := range cfg.Render {
		if err := render.File(rc, palettes); err != nil {
			return err
		}
	}
	return nil
}

// ── preview ───────────────────────────────────────────────────────────────

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview palettes (subcommands: ascii, build, serve)",
}

var previewVariant string

var previewAsciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "Print truecolor swatch table to stdout",
	RunE: func(cmd *cobra.Command, args []string) error {
		palettes, err := palette.LoadAll()
		if err != nil {
			return err
		}
		variant := viper.GetString("preview.variant")
		if variant != "" {
			for _, p := range palettes {
				if p.Variant == variant {
					preview.PrintOne(p)
					return nil
				}
			}
			return fmt.Errorf("variant %q not found", variant)
		}
		preview.PrintAll(palettes)
		return nil
	},
}

var previewBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate public/ assets for Pages deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: generate public/chromagoth.css and public/palettes.js
		fmt.Println("preview build: not yet implemented")
		return nil
	},
}

var previewServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve public/ locally over HTTP",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: net/http file server for public/
		fmt.Println("preview serve: not yet implemented")
		return nil
	},
}

// ── init / main ───────────────────────────────────────────────────────────

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: chromagoth.toml in CWD)")

	// preview ascii --variant
	previewAsciiCmd.Flags().StringVarP(&previewVariant, "variant", "v", "", "single variant to preview")
	_ = viper.BindPFlag("preview.variant", previewAsciiCmd.Flags().Lookup("variant"))

	previewCmd.AddCommand(previewAsciiCmd)
	previewCmd.AddCommand(previewBuildCmd)
	previewCmd.AddCommand(previewServeCmd)

	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(renderCmd)
	rootCmd.AddCommand(previewCmd)
}

func initConfig() {
	// Config file: flag > env > default
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("chromagoth")
		viper.SetConfigType("toml")
	}

	// Env vars: CHROMAGOTH_* prefix, auto-mapped
	viper.SetEnvPrefix("CHROMAGOTH")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// Read config if present; non-fatal if missing (not all commands need it)
	_ = viper.ReadInConfig()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
