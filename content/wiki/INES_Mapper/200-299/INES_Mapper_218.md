---
title: "INES Mapper 218"
url: "/wiki/INES_Mapper_218/"
description: "iNES Mapper 218 is used by the homebrew game Magic Floor. Excluding the CIC, the cartridge board only contains a single PRG-ROM chip (up to 32KiB)."
wikiPage: "INES_Mapper_218"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

**iNES Mapper 218** is used by the homebrew game *Magic Floor*. Excluding the [CIC](../CIC_lockout_chip/), the cartridge board only contains a single PRG-[ROM](../ROM/) chip (up to 32KiB).

There is no [CHR-ROM or CHR-RAM](../CHR_ROM_vs._CHR_RAM/) on the cartridge. Instead, the console's internal 2KiB nametable RAM is mapped as CHR-RAM. This RAM is permanently selected (CIRAM /CE wired to GND), and can be used in four configurations by wiring CIRAM A10 to one of the PPU A10..A13 address lines:

| CIRAM A10 connection | Effect on [PPU nametables](../PPU_nametables/) | Effect on [PPU pattern tables](../PPU_pattern_tables/) | [iNES Flags 6](../INES/#Flags_6) | [UNIF](../UNIF/) "MIRR" bit 7-0 |
|----|----|----|----|----|
| PPU A10 | Two-screen, **Horizontal arrangement** ("Vertical mirroring") | 2KiB (128 tiles) CHR-RAM, also mapped to nametable data | \$A1 (0001) | \$01 |
| PPU A11 | Two-screen, **Vertical arrangement** ("Horizontal mirroring") | 2KiB (128 tiles) CHR-RAM, also mapped to nametable data | \$A0 (0000) | \$00 |
| PPU A12 | Single-screen arrangement | **BLK0**: 1KiB (64 tiles) CHR-RAM, "swappable" via [PPUCTRL](../PPU_registers/) | \$A8 (1000) | \$02 |
| PPU A13 | Single-screen arrangement | **BLK1**: 1KiB (64 tiles) CHR-RAM | \$A9 (1001) | \$03 |

Notes:

- The "alternative nametable layout" (bit 3) in [Flags 6](../INES/#Flags_6) of the iNES header is used to denote the two Single-screen configurations.
- If reads/writes to mirrored addresses are avoided, software for this mapper can also be made compatible with [AxROM](../AxROM/) (\$A8,\$A9), [BxROM](../INES_Mapper_034/) (\$A0,A1), [NROM](../NROM/) + CHR-RAM, or other suitable mappers. *Magic Floor* itself is well behaved enough to be emulated as AxROM or NROM + CHR-RAM with vertical mirroring.

## Software {#Software}

Games:

- [*Magic Floor*](https://problemkaputt.de/magicflr.htm) (\$A8)
- [*Starfight*](https://problemkaputt.de/starfigh.htm) (\$A8)

Tests:

- [OAM corruption stress test ROM](https://forums.nesdev.org/viewtopic.php?p=232238#p232238) (\$A0)
- [Coredump tool v1.7](https://forums.nesdev.org/viewtopic.php?p=292020&sid=34b6a67d955fbe3625158d12f209effd#p292020) (\$A8)

Miscellaneous:

- [Port of Apple 1 Woz Monitor](https://github.com/bitcores/nesmon) (\$A9)
- [Port of Microsoft BASIC](https://github.com/bitcores/msbasic) (\$A9)

## References {#References}

- [nesdev forum, Post subject: Single Chip Cartridge](https://forums.nesdev.org/viewtopic.php?t=9342)
- [nesdev forum, Post subject: Using the NES's PPU RAM as CHR RAM?](https://forums.nesdev.org/viewtopic.php?t=5156)
