#!/usr/bin/env bats
# CLI integration tests — requires the binary at test/chromagoth
# Build with: go build -o test/chromagoth ./cmd/chromagoth

BIN="$BATS_TEST_DIRNAME/chromagoth"

setup() {
  TMPDIR="$(mktemp -d)"
}

teardown() {
  rm -rf "$TMPDIR"
}

# ── binary sanity ──────────────────────────────────────────────────────────

@test "binary exists and is executable" {
  [ -x "$BIN" ]
}

@test "root command shows help" {
  run "$BIN" --help
  [ "$status" -eq 0 ]
  [[ "$output" == *"chromagoth"* ]]
}

# ── preview ascii ──────────────────────────────────────────────────────────

@test "preview ascii outputs all themes" {
  run "$BIN" preview ascii
  [ "$status" -eq 0 ]
  [[ "$output" == *"Chromagoth"* ]]
  [[ "$output" == *"ground"* ]]
  [[ "$output" == *"cyber-pink"* ]]
}

@test "preview ascii --variant cyber outputs single theme" {
  run "$BIN" preview ascii --variant cyber
  [ "$status" -eq 0 ]
  [[ "$output" == *"Cyber"* ]]
  [[ "$output" != *"Chromagoth Military"* ]]
}

@test "preview ascii --variant nonexistent exits nonzero" {
  run "$BIN" preview ascii --variant doesnotexist
  [ "$status" -ne 0 ]
  [[ "$output" == *"not found"* ]]
}

@test "preview ascii includes ash #6a6c70 in every theme" {
  run "$BIN" preview ascii
  [ "$status" -eq 0 ]
  # ash is fixed across all themes
  count=$(echo "$output" | grep -ci "6a6c70")
  [ "$count" -ge 11 ]
}

# ── render ─────────────────────────────────────────────────────────────────

@test "render explicit template produces output" {
  echo '/* {{ .variant }} */' > "$TMPDIR/t.css.tmpl"
  run "$BIN" render "$TMPDIR/t.css.tmpl" -o "$TMPDIR/out.css"
  [ "$status" -eq 0 ]
  [ -f "$TMPDIR/out.css" ]
}

@test "render per-variant produces one file per theme" {
  echo '/* {{ .variant }} */' > "$TMPDIR/t.tmpl"
  run "$BIN" render "$TMPDIR/t.tmpl" \
    --per-variant \
    -o "$TMPDIR/chromagoth-{{ .variant }}.css"
  [ "$status" -eq 0 ]
  count=$(ls "$TMPDIR"/chromagoth-*.css 2>/dev/null | wc -l)
  [ "$count" -ge 11 ]
}

@test "render with chromagoth.toml reads config" {
  echo '/* {{ .variant }} */' > "$TMPDIR/t.css.tmpl"
  cat > "$TMPDIR/chromagoth.toml" <<EOF
[port]
name = "test"

[[render]]
template = "t.css.tmpl"
output = "out.css"
per_variant = false
EOF
  cd "$TMPDIR" && run "$BIN" render
  [ "$status" -eq 0 ]
  [ -f "$TMPDIR/out.css" ]
}

@test "render with no template and no toml exits nonzero" {
  cd "$TMPDIR" && run "$BIN" render
  [ "$status" -ne 0 ]
}

@test "render template uses hex helper" {
  # per-variant mode: template receives single palette context
  echo '{{ hex .surface.base }}' > "$TMPDIR/t.tmpl"
  run "$BIN" render "$TMPDIR/t.tmpl" \
    --per-variant -o "$TMPDIR/chromagoth-{{ .variant }}.txt"
  [ "$status" -eq 0 ]
  first=$(ls "$TMPDIR"/chromagoth-*.txt | head -1)
  [[ "$(cat "$first")" == "#"* ]]
}

@test "render template uses rgb helper" {
  echo '{{ rgb .accent.primary }}' > "$TMPDIR/t.tmpl"
  run "$BIN" render "$TMPDIR/t.tmpl" \
    --per-variant -o "$TMPDIR/chromagoth-{{ .variant }}.txt"
  [ "$status" -eq 0 ]
  first=$(ls "$TMPDIR"/chromagoth-*.txt | head -1)
  [[ "$(cat "$first")" == "rgb("* ]]
}

@test "render template uses hsl helper" {
  echo '{{ hsl .accent.danger }}' > "$TMPDIR/t.tmpl"
  run "$BIN" render "$TMPDIR/t.tmpl" \
    --per-variant -o "$TMPDIR/chromagoth-{{ .variant }}.txt"
  [ "$status" -eq 0 ]
  first=$(ls "$TMPDIR"/chromagoth-*.txt | head -1)
  [[ "$(cat "$first")" == "hsl("* ]]
}

@test "render respects CHROMAGOTH_RENDER_OUTPUT env var" {
  echo '{{ range .palettes }}/* {{ .variant }} */{{ end }}' > "$TMPDIR/t.tmpl"
  CHROMAGOTH_RENDER_OUTPUT="$TMPDIR/env-out.css" \
    run "$BIN" render "$TMPDIR/t.tmpl"
  [ "$status" -eq 0 ]
  [ -f "$TMPDIR/env-out.css" ]
}
