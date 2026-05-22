Sounds like the right split.

Use Codex for the deterministic parts:

```text
palettes/
  chromagoth-cyber.yaml
  chromagoth-exec.yaml
  chromagoth-tank.yaml
  chromagoth-retro.yaml
  chromagoth-pastel.yaml
  chromagoth-doll.yaml
  chromagoth-vamp.yaml
  chromagoth-synth.yaml
  chromagoth-bubble.yaml
  chromagoth-velvet.yaml
  chromagoth-haze.yaml
```

I’d finalize these before characters:

```text
1. Stable color slot names
2. Final hex values
3. Variant names
4. Light/dark classification where relevant
5. Accessibility/contrast checks
6. Exports: YAML, JSON, CSS vars, maybe terminal/theme formats
7. Short style notes per variant
8. Character prompt bible per mascot
```

Current variant lineup:

| Variant    | Style source           |
| ---------- | ---------------------- |
| **Cyber**  | Cybergoth / industrial |
| **Exec**   | Corpo goth             |
| **Tank**   | Military goth          |
| **Retro**  | Classic / trad goth    |
| **Pastel** | Pastel goth            |
| **Doll**   | Gothic lolita          |
| **Vamp**   | Vampiric goth          |
| **Synth**  | Health goth            |
| **Bubble** | Perky goth             |
| **Velvet** | Romantic goth          |
| **Haze**   | Soft goth              |

After the palettes are locked, character generation here will be much cleaner because each mascot can inherit the exact palette, vibe, material language, silhouette, and accessory rules.

==========================================================================================================

Toxic dropped for now. Added the four missing palettes:

| Color name       | Corpo Goth | Military Goth | Goth Lolita | Romantic Goth |
| ---------------- | ---------: | ------------: | ----------: | ------------: |
| **Ground**       |  `#0C0F13` |     `#0B0E0A` |   `#120E16` |     `#100A10` |
| **Veil**         |  `#141922` |     `#11170F` |   `#1D1722` |     `#1B1019` |
| **Field**        |  `#1E2630` |     `#1B2418` |   `#2A2130` |     `#291827` |
| **Trace**        |  `#2D3542` |     `#2C3826` |   `#3B3144` |     `#3B2637` |
| **Ash**          |  `#6A6C70` |     `#6A6C70` |   `#6A6C70` |     `#6A6C70` |
| **Mist**         |  `#9AA4B0` |     `#9AA68E` |   `#B2A8BB` |     `#B49AAD` |
| **Haze**         |  `#C8D0DC` |     `#C0CAAF` |   `#DED3E6` |     `#DCC1D0` |
| **Graphite**     |  `#F2F6FF` |     `#EFF6E8` |   `#FFF4FA` |     `#FFF0F7` |
| **Circuit Lime** |  `#A6E22E` |     `#A7D129` |   `#BFD86A` |     `#B8D66A` |
| **Powder Blush** |  `#D9919A` |     `#C8867A` |   `#FFC0C8` |     `#F0A0A8` |
| **Static Mint**  |  `#6FE7D6` |     `#78D7B6` |   `#A6E8DD` |     `#8FD9C6` |
| **Laser Blue**   |  `#4B7DFF` |     `#4F86C6` |   `#9AB9FF` |     `#7A9BFF` |
| **Cyber Pink**   |  `#D85AA7` |     `#C85C9A` |   `#F28BC8` |     `#E05A9F` |
| **Ultraviolet**  |  `#8F6EF6` |     `#8068C8` |   `#C9A7FF` |     `#A875E6` |
| **Amber Glow**   |  `#D6A14A` |     `#C99A3D` |   `#E8C16A` |     `#D6A85A` |
| **Cherry Flux**  |  `#D94B63` |     `#C94B4B` |   `#E86A86` |     `#D8456D` |

Current family:

| Style                  | Palette    |
| ---------------------- | ---------- |
| Cybergoth / Industrial | Neon Dark  |
| Perky Goth             | Neon Light |
| Soft Goth              | Soft Dark  |
| Pastel Goth            | Soft Light |
| Vampiric Goth          | Infrared   |
| Health Goth            | Cryo       |
| Corpo Goth             | new        |
| Military Goth          | new        |
| Goth Lolita            | new        |
| Romantic Goth          | new        |
-----

Good addition. **Chromagoth Retro** fills the “classic/trad goth” slot without competing with Cyber, Vamp, or Velvet.

Vibe: almost grayscale, old club flyers, black lipstick, vinyl, smoke, silver jewelry, cold photos, desaturated undertones.

| Color name       | Chromagoth Retro |
| ---------------- | ---------------: |
| **Ground**       |        `#0D0D0F` |
| **Veil**         |        `#151519` |
| **Field**        |        `#202126` |
| **Trace**        |        `#30323A` |
| **Ash**          |        `#6A6C70` |
| **Mist**         |        `#8D9096` |
| **Haze**         |        `#B8BAC0` |
| **Graphite**     |        `#ECEEF2` |
| **Circuit Lime** |        `#A7B08A` |
| **Powder Blush** |        `#B09A9E` |
| **Static Mint**  |        `#8FA9A3` |
| **Laser Blue**   |        `#8D9CB8` |
| **Cyber Pink**   |        `#AA8FA2` |
| **Ultraviolet**  |        `#978FAF` |
| **Amber Glow**   |        `#B0A078` |
| **Cherry Flux**  |        `#B08A90` |

Updated lineup:

| Style                  | Variant               |
| ---------------------- | --------------------- |
| Cybergoth / Industrial | **Chromagoth Cyber**  |
| Corpo Goth             | **Chromagoth Exec**   |
| Military Goth          | **Chromagoth Tank**   |
| Classic / Trad Goth    | **Chromagoth Retro**  |
| Pastel Goth            | **Chromagoth Pastel** |
| Goth Lolita            | **Chromagoth Doll**   |
| Vampiric Goth          | **Chromagoth Vamp**   |
| Health Goth            | **Chromagoth Synth**  |
| Perky Goth             | **Chromagoth Bubble** |
| Romantic Goth          | **Chromagoth Velvet** |
| Soft Goth              | **Chromagoth Haze**   |
-----

Sounds like the right split.

Use Codex for the deterministic parts:

```text
palettes/
  chromagoth-cyber.yaml
  chromagoth-exec.yaml
  chromagoth-tank.yaml
  chromagoth-retro.yaml
  chromagoth-pastel.yaml
  chromagoth-doll.yaml
  chromagoth-vamp.yaml
  chromagoth-synth.yaml
  chromagoth-bubble.yaml
  chromagoth-velvet.yaml
  chromagoth-haze.yaml
```

I’d finalize these before characters:

```text
1. Stable color slot names
2. Final hex values
3. Variant names
4. Light/dark classification where relevant
5. Accessibility/contrast checks
6. Exports: YAML, JSON, CSS vars, maybe terminal/theme formats
7. Short style notes per variant
8. Character prompt bible per mascot
```

Current variant lineup:

| Variant    | Style source           |
| ---------- | ---------------------- |
| **Cyber**  | Cybergoth / industrial |
| **Exec**   | Corpo goth             |
| **Tank**   | Military goth          |
| **Retro**  | Classic / trad goth    |
| **Pastel** | Pastel goth            |
| **Doll**   | Gothic lolita          |
| **Vamp**   | Vampiric goth          |
| **Synth**  | Health goth            |
| **Bubble** | Perky goth             |
| **Velvet** | Romantic goth          |
| **Haze**   | Soft goth              |

After the palettes are locked, character generation here will be much cleaner because each mascot can inherit the exact palette, vibe, material language, silhouette, and accessory rules.
-----

Fair. Use names that are abstract but still pleasant and theme-like.

| Role             | Neon Dark | Neon Light | Soft Dark | Soft Light |
| ---------------- | --------: | ---------: | --------: | ---------: |
| **Obsidian**     | `#0B0D14` |  `#FBFCFF` | `#11131A` |  `#FFFBFD` |
| **Eclipse**      | `#121725` |  `#F1F5FF` | `#191D27` |  `#F8F2F8` |
| **Carbon**       | `#1B2234` |  `#E2E8F5` | `#242A38` |  `#EDE6EE` |
| **Lattice**      | `#2A344C` |  `#C8D1E5` | `#333B4D` |  `#D9D1DC` |
| **Ash Graphite** | `#6A6C70` |  `#6A6C70` | `#6A6C70` |  `#6A6C70` |
| **Mist**         | `#95A0BA` |  `#4F5B74` | `#A3ABBD` |  `#655D70` |
| **Glyph**        | `#EDF2FF` |  `#171D2D` | `#F0F3FF` |  `#2C2534` |
| **Circuit Lime** | `#B5FF2E` |  `#88C60A` | `#C4E76B` |  `#9DC452` |
| **Powder Blush** | `#FF9F97` |  `#EF7F77` | `#FFC1BB` |  `#F5AAA5` |
| **Static Mint**  | `#67FFF2` |  `#00BFB5` | `#A0ECE4` |  `#73D9D0` |
| **Cyber Pink**   | `#FF4FA8` |  `#E63B8F` | `#EF8DC5` |  `#D96CA9` |
| **Ultraviolet**  | `#9B6DFF` |  `#6F48F8` | `#B6A1FF` |  `#917AE9` |
| **Amber Glow**   | `#FFBA4A` |  `#D88B00` | `#F2CA7D` |  `#D4A24C` |
| **Cherry Flux**  | `#FF5D74` |  `#D63B59` | `#EE8CA0` |  `#D16D80` |

Better than `Null / Drift / Trace / Vector`.

The four base names now read more like material/structure:

| Layer | Name         |
| ----- | ------------ |
| 1     | **Obsidian** |
| 2     | **Eclipse**  |
| 3     | **Carbon**   |
| 4     | **Lattice**  |

My only concern: **Obsidian** leans dark. For true light/dark neutrality, this alternative may be even cleaner:

| Layer | Name        |
| ----- | ----------- |
| 1     | **Origin**  |
| 2     | **Halo**    |
| 3     | **Matter**  |
| 4     | **Lattice** |
----


Toxic dropped for now. Added the four missing palettes:

| Color name       | Corpo Goth | Military Goth | Goth Lolita | Romantic Goth |
| ---------------- | ---------: | ------------: | ----------: | ------------: |
| **Ground**       |  `#0C0F13` |     `#0B0E0A` |   `#120E16` |     `#100A10` |
| **Veil**         |  `#141922` |     `#11170F` |   `#1D1722` |     `#1B1019` |
| **Field**        |  `#1E2630` |     `#1B2418` |   `#2A2130` |     `#291827` |
| **Trace**        |  `#2D3542` |     `#2C3826` |   `#3B3144` |     `#3B2637` |
| **Ash**          |  `#6A6C70` |     `#6A6C70` |   `#6A6C70` |     `#6A6C70` |
| **Mist**         |  `#9AA4B0` |     `#9AA68E` |   `#B2A8BB` |     `#B49AAD` |
| **Haze**         |  `#C8D0DC` |     `#C0CAAF` |   `#DED3E6` |     `#DCC1D0` |
| **Graphite**     |  `#F2F6FF` |     `#EFF6E8` |   `#FFF4FA` |     `#FFF0F7` |
| **Circuit Lime** |  `#A6E22E` |     `#A7D129` |   `#BFD86A` |     `#B8D66A` |
| **Powder Blush** |  `#D9919A` |     `#C8867A` |   `#FFC0C8` |     `#F0A0A8` |
| **Static Mint**  |  `#6FE7D6` |     `#78D7B6` |   `#A6E8DD` |     `#8FD9C6` |
| **Laser Blue**   |  `#4B7DFF` |     `#4F86C6` |   `#9AB9FF` |     `#7A9BFF` |
| **Cyber Pink**   |  `#D85AA7` |     `#C85C9A` |   `#F28BC8` |     `#E05A9F` |
| **Ultraviolet**  |  `#8F6EF6` |     `#8068C8` |   `#C9A7FF` |     `#A875E6` |
| **Amber Glow**   |  `#D6A14A` |     `#C99A3D` |   `#E8C16A` |     `#D6A85A` |
| **Cherry Flux**  |  `#D94B63` |     `#C94B4B` |   `#E86A86` |     `#D8456D` |

Current family:

| Style                  | Palette    |
| ---------------------- | ---------- |
| Cybergoth / Industrial | Neon Dark  |
| Perky Goth             | Neon Light |
| Soft Goth              | Soft Dark  |
| Pastel Goth            | Soft Light |
| Vampiric Goth          | Infrared   |
| Health Goth            | Cryo       |
| Corpo Goth             | new        |
| Military Goth          | new        |
| Goth Lolita            | new        |
| Romantic Goth          | new        |
