# Chromagoth Palettes

Source of truth for all theme data. YAML in `src/` → Go CLI generates CSS, SCSS, LESS, JS, and the preview assets. Also provides the `chromagoth` binary that port repos use to render their templates.

## Repo layout (full org)

```
gitlab.com/patrick.pfenning.92/chromagoth/
  palettes/      ← this repo: Go CLI + palette YAML + dist outputs
  css/           ← thin npm re-export of palettes dist
  site/          ← Zola SSG preview
  vscode/        ← port: template + generated output + marketplace CI
  alacritty/     ← port: template + generated output
  tailwind/      ← port: template + generated output + npm CI
  terminal/      ← port: template + generated output
  ...            ← one repo per port, named after the target
```

Each port repo contains its own template, committed generated output, and CI that publishes to its native registry/marketplace. The `chromagoth` binary (released from this repo) is the only shared contract.

## Go CLI

```
palettes/
  cmd/chromagoth/    # CLI entrypoint
  internal/
    palette/         # YAML loading, color types, semantic role mapping
    render/          # template execution, color helper funcs
  go.mod
  go.sum
```

### Commands

```bash
# Build
chromagoth build                      # regenerates dist/ (replaces generate.js)

# Port rendering — run from inside a port repo
chromagoth render <tmpl>              # renders template against all variants
chromagoth render chromagoth.css.tmpl # example: renders css port template

# Preview
chromagoth preview ascii              # truecolor swatch table for all themes
chromagoth preview ascii --variant cyber  # single theme
chromagoth preview build              # generates public/ assets for Pages deployment
chromagoth preview serve              # local HTTP server for public/
```

### Template context and color helpers

```
{{ .surface.base }}               # hex string via semantic role
{{ .accent.primary }}             # hex string via semantic role
{{ .colors.cyber-pink }}          # hex via raw slot name (avoid in port templates)
{{ rgb .accent.primary }}         # "rgb(255, 79, 168)"
{{ hsl .surface.border }}         # "hsl(228, 29%, 23%)"
{{ alpha .accent.secondary 0.5 }} # "rgba(155, 109, 255, 0.5)"
{{ hex .accent.primary }}         # "#ff4fa8" (explicit, same as default)
```

Port templates iterate over all variants — one template produces one output file per variant.

### Palette data — embedded in binary

YAML source files are compiled into the binary via `go:embed`:
```go
//go:embed src/*.yaml
var paletteFS embed.FS
```
Port repos need only the binary. No separate palette data files at runtime.
Upgrading palette data = releasing a new binary version.

### ASCII preview

Renders truecolor ANSI color blocks to stdout. Uses `\033[48;2;R;G;Bm    \033[0m` — no deps.
Works in any truecolor terminal. Light themes invert foreground text color for legibility.

```
Chromagoth Cyber — Neon Dark
  ground        ████  #0b0d14
  veil          ████  #121725
  ...
  ──────────────────────────
  circuit-lime  ████  #b5ff2e
  cyber-pink    ████  #ff4fa8
  ...
```

### CSS preview (GitLab Pages)

`chromagoth preview build` generates `public/chromagoth.css` and `public/palettes.js`.
`public/index.html` is a maintained static file — the CLI feeds it data, does not own the HTML.
`chromagoth preview serve` replaces `pnpm preview` / `scripts/serve.js`.

CI calls `chromagoth preview build` before the Pages deployment step.

## Retired once Go CLI is complete

- `scripts/generate.js`
- `scripts/serve.js`
- pnpm build / preview scripts (package.json scripts entry becomes a thin wrapper or is removed)

## Build & deploy

```bash
go run ./cmd/chromagoth build         # during development
chromagoth build                      # once installed / in CI
git add <files> && git commit -m "..." && git push -o "ci.variable=DEPLOY_PAGES=true"
```

## Palette system — 16 slots

**Base (8)** — form the tonal scale:
| Slot | Dark theme | Light theme |
|------|-----------|-------------|
| `ground` | darkest | lightest |
| `veil` | ↓ | ↓ |
| `field` | ↓ | ↓ |
| `trace` | ↓ | ↓ |
| `ash` | **#6A6C70 — FIXED, never change, universal across all themes** |
| `mist` | ↓ | ↓ |
| `haze` | ↓ | ↓ |
| `graphite` | lightest | darkest |

`dark: true` → ground is darkest, graphite is lightest.
`dark: false` → ground is lightest, graphite is darkest (inverted).

**Accents (8):**
`circuit-lime`, `powder-blush`, `static-mint`, `laser-blue`, `cyber-pink`, `ultraviolet`, `amber-glow`, `cherry-flux`

## YAML fields

```yaml
name: Chromagoth <Name>
variant: <slug>          # matches filename, used as data-theme value
mascot: <name>           # character codename for future mascot generation
style: <2-4 word title>  # evocative, like an album name — NOT literal
vibe: "<tweet>"          # abstract, punchy, 1-3 short sentences — not AI cringe
dark: true|false
status: draft|final|wip
```

**`style`** — abstract/evocative codename, not a material or literal description.
**`vibe`** — terse, poetic. Avoid AI filler. No latex/material references.

## Current themes

| Variant | Style | Dark | Status |
|---------|-------|------|--------|
| cyber | Neon Dark | ✓ | draft |
| health | Cryo | ✓ | **final** |
| lolita | Fragile Reign | | draft |
| military | Field Directive | ✓ | **final** |
| corpo | Corporate Séance | ✓ | draft |
| soft | Soft Dark | ✓ | draft |
| vampiric | Infrared | ✓ | draft |
| retro | Ghost Signal | ✓ | draft |
| pastel | Soft Light | | draft |
| perky | Neon Light | | draft |
| romantic | Velvet Confession | | draft |

`ash: "#6A6C70"` is identical in every single one of these. That is intentional and must never change.

## WIP workflow (palette iteration)

1. Create `src/chromagoth-{variant}-wip.yaml` alongside the canonical file
2. Set `variant: {variant}-wip`, `status: wip`
3. Run `pnpm build` and push with deploy flag — WIP appears in nav automatically
4. When approved: copy the `colors:` block into the canonical file, delete the WIP file, rebuild and push

## Preview (`public/index.html`)

- `NAV_ORDER` array controls sidebar order — update when adding new variants
- Nav uses roving tabindex: Tab enters/exits sidebar once, arrow keys navigate within it
- Arrow key navigation applies theme immediately on focus (no Enter needed)
- Format picker: HEX / RGB / HSL — click any hex value to copy
- Steckbrief card shows: name, dark/light badge, style, vibe, mascot, status
- `public/chromagoth.css` and `public/palettes.js` are generated — do not edit manually

## Semantic port layer

Sits between raw palette slots and any port target (IDE, app, terminal, web, Tailwind).
Every port template references semantic role names — never raw slot names directly.
The mapping is universal across all 11 variants; the underlying hex differs per theme, the role does not.

### Surfaces — depth hierarchy

| Role | Slot | Used for |
|------|------|----------|
| `surface.base` | `ground` | page bg, window bg, editor bg |
| `surface.raised` | `veil` | sidebar, panel, card, sheet |
| `surface.overlay` | `field` | dropdown, modal, tooltip, tab bar |
| `surface.border` | `trace` | dividers, input outlines, separators |

### Text — foreground hierarchy

| Role | Slot | Used for |
|------|------|----------|
| `text.disabled` | `ash` | **FIXED #6A6C70** — placeholder, disabled, inactive |
| `text.subtle` | `mist` | captions, secondary labels, code comments |
| `text.dim` | `haze` | tertiary text, breadcrumbs, metadata |
| `text.default` | `graphite` | body text, headings, primary content |

### Accents — functional meaning

| Role | Slot | Used for |
|------|------|----------|
| `accent.primary` | `cyber-pink` | main CTA, brand color, hover/active states |
| `accent.secondary` | `ultraviolet` | focus rings, badges, secondary brand |
| `accent.link` | `laser-blue` | hyperlinks, interactive inline text |
| `accent.positive` | `circuit-lime` | success, added diff, passing test, online |
| `accent.info` | `static-mint` | informational hint, neutral notice |
| `accent.caution` | `powder-blush` | soft warning, notice, attention |
| `accent.warning` | `amber-glow` | warning, caution, pending |
| `accent.danger` | `cherry-flux` | error, removed diff, destructive action |

### How roles map to each target type

**CSS / Tailwind** — CSS custom properties or Tailwind color config entries keyed by role name.
Hover/active variants are computed at the port level (e.g. `primary` + lightness shift), not defined in the palette.

**IDE chrome + syntax** (VS Code, Zed, JetBrains):
- Chrome: `base` → editor bg, `raised` → sidebar, `overlay` → tabs/dropdowns, `border` → borders, `secondary` → focus ring
- Syntax: `primary` → keywords/operators, `positive` → strings/literals, `link` → functions/methods, `secondary` → types/classes, `warning` → constants, `info` → builtins, `subtle` → comments, `danger` → errors/invalid

**Terminal ANSI 16:**
- black/bright black → `base` / `border`
- white/bright white → `dim` / `default`
- red → `danger`, green → `positive`, yellow → `warning`, blue → `link`, magenta → `primary`, cyan → `info`
- bright variants → same slots, slightly shifted by the port template

**Desktop apps** (GTK, Qt, Electron, Obsidian):
- Window/app background → `base`, widget/card background → `raised`, selected item → `overlay`, brand accent → `primary`, destructive → `danger`

### Port authoring rules

- Port templates reference role names (`accent.primary`, `surface.base`, etc.) not slot names (`cyber-pink`, `ground`)
- Expansion beyond 16 (e.g. VS Code's 50+ slots) is handled in the port template — map multiple VS Code keys to the same role, or compute variants (alpha, lightness) from the base role
- Per-variant overrides are not planned; the universal mapping is the contract

## Moodboard

Reference images live at `/home/p/code/chromagoth/moodboard/` (outside all git repos).
`sources.yaml` at the root of moodboard lists all original URLs keyed by theme name.
When downloading images, use the URL's hash/ID as the filename.
