---
title: "INES Mapper 100"
url: "/wiki/INES_Mapper_100/"
description: "The assignment of iNES Mapper 100 depends on emulators."
wikiPage: "INES_Mapper_100"
wikiCategories:
  - name: "Bad iNES Mappers"
    url: "/category/Bad_iNES_Mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

The assignment of **iNES Mapper 100** depends on emulators.

## Nesticle mapper hacks {#Nesticle_mapper_hacks}

This mapper was assigned by fwNES<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> and NintendulatorNRS to denote [mapper hacks](../Glossary/) for the [MMC3](../MMC3/) emulation of Nesticle, an influential NES emulator for MS-DOS in the late 1990s.

Nesticle did not directly support many of the custom ASICs by Irem, Konami and Namco, but did emulate the MMC3 in an incorrect way that made it useful to run modified ROM images. Most of them use 512-byte trainers that must be loaded to CPU address \$7000.

### <span id="Bankswitch_Index_Register_.28.248000.29.2FData_Register_.28.248001.29"></span>Bankswitch Index Register (\$8000)/Data Register (\$8001) {#Bankswitch_Index_Register_.28.248000.29.2FData_Register_.28.248001.29}

    Mask: $E001 (nominally)

    $00: 1 KiB CHR-ROM bank at PPU $0000-$03FF := Value AND $FE
         1 KiB CHR-ROM bank at PPU $0400-$07FF := Value OR $01
    $01: 1 KiB CHR-ROM bank at PPU $0800-$0BFF := Value AND $FE
         1 KiB CHR-ROM bank at PPU $0C00-$0FFF := Value OR $01
    $02: 1 KiB CHR-ROM bank at PPU $1000-$13FF := Value
    $03: 1 KiB CHR-ROM bank at PPU $1400-$17FF := Value
    $04: 1 KiB CHR-ROM bank at PPU $1800-$1BFF := Value
    $05: 1 KiB CHR-ROM bank at PPU $1C00-$1FFF := Value
    $06: 8 KiB PRG-ROM bank at CPU $8000-$9FFF := Value
    $07: 8 KiB PRG-ROM bank at CPU $A000-$BFFF := Value
    $46: 8 KiB PRG-ROM bank at CPU $C000-$DFFF := Value
    $47: 8 KiB PRG-ROM bank at CPU $A000-$BFFF := Value
    $80: 1 KiB CHR-ROM bank at PPU $1000-$13FF := Value AND $FE
         1 KiB CHR-ROM bank at PPU $1400-$17FF := Value OR $01
    $81: 1 KiB CHR-ROM bank at PPU $1800-$1BFF := Value AND $FE
         1 KiB CHR-ROM bank at PPU $1C00-$1FFF := Value OR $01
    $82: 1 KiB CHR-ROM bank at PPU $0000-$03FF := Value
    $83: 1 KiB CHR-ROM bank at PPU $0400-$07FF := Value
    $84: 1 KiB CHR-ROM bank at PPU $0800-$0BFF := Value
    $85: 1 KiB CHR-ROM bank at PPU $0C00-$0FFF := Value

On an actual MMC3, changing bits 6 and 7 in the index register would immediately apply the current content of all seven bank registers to their flipped meaning. Nesticle only applies these bits on the affected bank of the next \$8001 write, effectively making eight independent 1 KiB CHR-ROM banks and three independent 8 KiB PRG-ROM banks available:

- Registers \$82-\$85 to set four 1 KiB CHR-ROM banks at CPU \$0000-\$0FFF
- Registers \$02-\$05 to set four 1 KiB CHR-ROM banks at CPU \$1000-\$1FFF
- Registers \$06, \$07 and \$46 to set three 8 KiB PRG-ROM banks at \$8000-\$DFFF

### See Also {#See_Also}

- [RAMBO-1](../RAMBO-1/) and [iNES Mapper 176](../INES_Mapper_176/)'s Extended MMC3 mode offers eight 1 KiB CHR-ROM banks and more than two 8 KiB PRG-ROM banks natively, and thus can be seen as the hardware analogue of this emulator-specific mapper.

## Other assignments {#Other_assignments}

- Nestopia uses iNES mapper 100 as a magic number for [UNIF](../UNIF/)-only boards.
- The original Nintendulator uses this as a general-purpose debugging mapper, allowing PRG/CHR banks to be selected arbitrarily using a control dialog.

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) FWNES98E.TXT, fwNES 0.302b, 1998/11/16, by FanWen Yang
