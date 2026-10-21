---
title: "INES Mapper 187"
url: "/wiki/INES_Mapper_187/"
description: "iNES Mapper 187 is used for the 卡聖 (Kǎshèng) A98402 and similar boards, an MMC3-clone-bearing board with an NROM-like PRG-ROM register that can override the MMC3 clone's PRG-ROM bank, and…"
wikiPage: "INES_Mapper_187"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 187 is used for the 卡聖 (Kǎshèng) A98402 and similar boards, an [MMC3](../MMC3/)-clone-bearing board with an NROM-like PRG-ROM register that can override the MMC3 clone's PRG-ROM bank, and supports 512 KiB of CHR-ROM. Used by:

- *Street Fighter Zero 2* (Kǎshèng/Hummer Team, not *Street Fighter Zero 2 '97*)
- *The King of Fighters '96* (Kǎshèng/Hummer Team)

Only King of Fighters '96 has 512 KiB of CHR-ROM; CHR A18 (the second half of CHR-ROM) is directly selected by PPU A12, regardless of what the MMC3 is set to do. This means that the "left" pattern table always uses the first and the "right" pattern table always uses the second 256 KiB half of CHR-ROM.

# Registers {#Registers}

## <span id="NROM_Override_Register_.28.245000.29"></span>NROM Override Register (\$5000) {#NROM_Override_Register_.28.245000.29}

    Mask: $F001

    D~7654 3210
      ---------
      M.NB BBb.
      | |+-+++-- Select 16 KiB PRG-ROM bank at CPU
      | |        $8000-$BFFF and $C000-$FFFF
      | +------- 0: Do not replace bit 1 (NROM-128)
      |          1: Replace bit 1 with CPU A14 (NROM-256)
      +--------- 0: Use PRG bank from MMC3; ignore $5000 bits 1-4/5
                 1: Ignore PRG bank from MMC3; apply $5000 bits 1-4/5

## <span id="Protection_Read_.28.245000-.245FFF.29"></span>Protection Read (\$5000-\$5FFF) {#Protection_Read_.28.245000-.245FFF.29}

Mask: Unknown

The actual values that are returned are unknown; *The King of Fighters '96* reads from here and only expects bit 7 of the value being returned to be set.

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.29"></span>MMC3-compatible registers (\$8000-\$FFFF) {#MMC3-compatible_registers_.28.248000-.24FFFF.29}

Mask: \$E001, see [MMC3](../MMC3/).

# Errata {#Errata}

- GoodNES 3.23b has several [INES Mapper 121](../INES_Mapper_121/) and [INES Mapper 189](../INES_Mapper_189/) images erroneously set to 187.

# Similar Mappers {#Similar_Mappers}

- iNES Mappers [115](../INES_Mapper_115/) and [123](../INES_Mapper_123/) have similar NROM override registers, with Mapper 123's bits arranged a bit differently.
- iNES Mapper [121](../INES_Mapper_121/) has the same CHR A18 behavior, drops the NROM override register, and adds more complex protection.
- The licensed mapper [iNES Mapper 088](../INES_Mapper_088/) uses a similar setup to address larger CHR-ROM, but using the MMC3 predecessor [Namco 108](../Namcot_108_family_pinout/), and can only address 128 KiB of CHR as a result.
