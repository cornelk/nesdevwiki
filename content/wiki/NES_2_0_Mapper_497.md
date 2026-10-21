---
title: "NES 2.0 Mapper 497"
url: "/wiki/NES_2.0_Mapper_497/"
description: "NES 2.0 Mapper 497 denotes the 小霸王 (Xiǎobàwáng, a.k.a. Subor) LOGO cartridge's circuit board."
wikiPage: "NES_2_0_Mapper_497"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**NES 2.0 Mapper 497** denotes the 小霸王 (Xiǎobàwáng, a.k.a. Subor) LOGO cartridge's circuit board.

# Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KiB unbanked PRG-ROM
- PPU \$0000-\$0FFF: 4 KiB PRG-RAM window into 32 KiB of CHR-RAM, fixed at bank \#0
- PPU \$1000-\$1FFF: 4 KiB PRG-RAM window into 32 KiB of CHR-RAM, switchable

# Register {#Register}

## CHR-RAM Bankswitching Mode {#CHR-RAM_Bankswitching_Mode}

    D~[.... ..BA] A~[1... .... .... ....] ($8000-$FFFF, no bus conflicts)
              ++- PPU $1000-$1FFF bank
                   0,1: set to bank #1
                   2:   automatically switch to PPU A9..8 latched at rising edge of PPU A13, +2
                   3:   like 2, but when PPU A9..8 is 3, set to bank #1

# Effective CHR-RAM banks {#Effective_CHR-RAM_banks}

| PPU \$1000-\$1FFF switching mode | Last nametable fetch address (excluding attribute table fetches) | 4 KiB CHR-RAM bank at PPU \$0000 | 4 KiB CHR-RAM bank at PPU \$1000 |
|----|----|----|----|
| 0 | \- | 0 | 1 |
| 1 | \- | 0 | 1 |
| 2 | 0xx | 0 | 2 |
| 2 | 1xx | 0 | 3 |
| 2 | 2xx | 0 | 4 |
| 2 | 3xx | 0 | 5 |
| 3 | 0xx | 0 | 2 |
| 3 | 1xx | 0 | 3 |
| 3 | 2xx | 0 | 4 |
| 3 | 3xx | 0 | 1 |

Mode 3 is used to facilitate a split screen with graphics on top and text in the bottom six character rows.
