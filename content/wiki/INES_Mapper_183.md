---
title: "INES Mapper 183"
url: "/wiki/INES_Mapper_183/"
description: "INES Mapper 183 was supposedly assigned to represent the pirate variant of Gimmick! known as Shui Guan Pipe. It uses a VRC4 clone in the VRC4e configuration (VRC4 A0=CPU A2, VRC4 A1=CPU A3), but…"
wikiPage: "INES_Mapper_183"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
---

***INES Mapper 183*** was supposedly assigned to represent the pirate variant of *Gimmick!* known as *Shui Guan Pipe*. It uses a [VRC4](../VRC2_and_VRC4/) clone in the VRC4e configuration (VRC4 A0=CPU A2, VRC4 A1=CPU A3), but additional circuits on the board adds an ability for mapping 8 KiB PRG-ROM bank at \$6000-\$7FFF.

## <span id="PRG_Select_.246000-.247FFF_.28Write.29"></span>PRG Select \$6000-\$7FFF (Write) {#PRG_Select_.246000-.247FFF_.28Write.29}

    A~[011. .... .... PPPP] $6000-$7FFF
                      ||||
                      ++++- Set 8 KiB PRG-ROM bank at $6000-$7FFF (only banks 0-15 can be swapped)

Note that address bits are used for specifying bank number, not data bus (writing to \$6002 would mean bank 2, \$6004 = bank 4 \$600A = bank 10, so on).

## Errata {#Errata}

- Early implementions of Mapper 183 has a register at \$A000 that switches 8 KiB PRG bank at \$C000-\$DFFF with the registers were masked with \$F80C, but this was not the case for actual cartridge.
  - Due to this, [Gimmick VRC4 Hack](../Gimmick_VRC4_Hack/) is previously thought to not be Mapper 183.
