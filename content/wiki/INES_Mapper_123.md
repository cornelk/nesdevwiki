---
title: "INES Mapper 123"
url: "/wiki/INES_Mapper_123/"
description: "iNES Mapper 123 is used for the 卡聖 (Kǎshèng) H2288 board, an MMC3-clone-bearing board with scrambled register indices, and an NROM-like PRG-ROM register that can override the MMC3 clone's…"
wikiPage: "INES_Mapper_123"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 123 is used for the 卡聖 (Kǎshèng) H2288 board, an [MMC3](../MMC3/)-clone-bearing board with scrambled register indices, and an NROM-like PRG-ROM register that can override the MMC3 clone's PRG-ROM bank. Its UNIF board name is **UNL-H2288**. Used by:

- *Mortal Kombat 3* (Kǎshèng/Hummer Team, not to be confused with *Mortal Kombat 3* from SuperGame)
- *Earthworm Jim 2* (Kǎshèng/Shin-Shin Electronics, not to be confused with *Earthworm Jim 2* from SuperGame)

# Registers {#Registers}

## <span id="NROM_Override_Register_.28.245800.29"></span>NROM Override Register (\$5800) {#NROM_Override_Register_.28.245800.29}

    Mask: $F800

    D~7654 3210
      ---------
      .M31 .2S0
       +++--+|+- Select 16 KiB PRG-ROM bank at CPU
       |     |   $8000-$BFFF and $C000-$FFFF (notice the bit order)
       |     +-- 0: Do not replace bit 0 (NROM-128)
       |         1: Replace bit 0 with CPU A14 (NROM-256)
       +-------- 0: Use PRG bank from MMC3; ignore $5800 bits 0-3/5
                 1: Ignore PRG bank from MMC3; apply $5800 bits 0-3/5

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.29"></span>MMC3-compatible registers (\$8000-\$FFFF) {#MMC3-compatible_registers_.28.248000-.24FFFF.29}

    $8000 index written   MMC3 register meant
    -------------------   -------------------
    0                     0
    1                     3
    2                     1
    3                     5
    4                     6
    5                     7
    6                     2
    7                     4

Bits 6 and 7 of the value written to \$8000 are unchanged. See [MMC3](../MMC3/) for further details.

# Similar Mappers {#Similar_Mappers}

- iNES Mappers [114](../INES_Mapper_114/) and [115](../INES_Mapper_115/) move the \$5800 register to \$6000 while rearranging its bits and add an outer CHR-ROM bank register at \$6001. Mapper 114 keeps the MMC3 index register scrambling (using the same pattern) and adds register address scrambling, while Mapper 115 drops all scrambling.
