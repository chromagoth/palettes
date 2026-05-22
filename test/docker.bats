#!/usr/bin/env bats
# Container integration tests — requires buildah + image chromagoth:test
# Build: buildah bud --platform linux/amd64 -t chromagoth:test .

IMAGE="chromagoth:test"

setup_file() {
  # Build once per test file run
  if ! buildah inspect "$IMAGE" &>/dev/null; then
    buildah bud --platform linux/amd64 -t "$IMAGE" \
      "$(dirname "$BATS_TEST_DIRNAME")" >&2
  fi
}

run_container() {
  # Run a throwaway container and capture output
  buildah run --isolation=chroot \
    "$(buildah from --quiet "$IMAGE")" \
    -- /chromagoth "$@"
}

# ── image sanity ───────────────────────────────────────────────────────────

@test "image exists" {
  run buildah inspect "$IMAGE"
  [ "$status" -eq 0 ]
}

@test "entrypoint is /chromagoth" {
  run buildah inspect --format '{{ .OCIv1.Config.Entrypoint }}' "$IMAGE"
  [ "$status" -eq 0 ]
  [[ "$output" == *"/chromagoth"* ]]
}

# ── container preview ──────────────────────────────────────────────────────

@test "container: preview ascii runs" {
  ctr=$(buildah from --quiet "$IMAGE")
  run buildah run "$ctr" -- /chromagoth preview ascii
  buildah rm "$ctr" &>/dev/null
  [ "$status" -eq 0 ]
  [[ "$output" == *"Chromagoth"* ]]
}

@test "container: preview ascii --variant health runs" {
  ctr=$(buildah from --quiet "$IMAGE")
  run buildah run "$ctr" -- /chromagoth preview ascii --variant health
  buildah rm "$ctr" &>/dev/null
  [ "$status" -eq 0 ]
  [[ "$output" == *"Health"* ]]
}

# ── container render ───────────────────────────────────────────────────────

@test "container: render with mounted template" {
  TMPDIR="$(mktemp -d)"
  echo '/* {{ .variant }} */' > "$TMPDIR/t.css.tmpl"

  ctr=$(buildah from --quiet "$IMAGE")
  buildah copy "$ctr" "$TMPDIR/t.css.tmpl" /work/t.css.tmpl
  run buildah run "$ctr" -- /chromagoth render /work/t.css.tmpl -o /work/out.css
  buildah rm "$ctr" &>/dev/null
  rm -rf "$TMPDIR"

  [ "$status" -eq 0 ]
}

@test "container: unknown variant exits nonzero" {
  ctr=$(buildah from --quiet "$IMAGE")
  run buildah run "$ctr" -- /chromagoth preview ascii --variant nope
  buildah rm "$ctr" &>/dev/null
  [ "$status" -ne 0 ]
}
