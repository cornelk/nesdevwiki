---
title: "INES Mapper 032"
url: "/wiki/INES_Mapper_032/"
description: "iNES Mapper 032 represents games using Irem's G-101 mapper IC. It came in a 52-pin DIP package."
wikiPage: "INES_Mapper_032"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

iNES Mapper 032 represents games using Irem's G-101 mapper IC. It came in a [52-pin DIP](../Irem_G-101_pinout/) package.

Example Games:

- Image Fight
- Major League
- Kaiketsu Yanchamaru 2
- Paaman - Enban o Torikaese!!

## Banks {#Banks}

- CPU \$8000-\$9FFF (or \$C000-\$DFFF): 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF (or \$8000-\$9FFF): 8 KiB PRG ROM bank, fixed to the second-to-last bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank

## Registers {#Registers}

### <span id="PRG_Select_0_.28.248000-.248FFF.29"></span>PRG Select 0 (\$8000-\$8FFF) {#PRG_Select_0_.28.248000-.248FFF.29}

Mask: \$F000

    7  bit  0
    ---------
    ...P PPPP
       | ||||
       +-++++- Select 8 KB PRG ROM at $8000-$9FFF or $C000-$DFFF

### <span id="Mirroring_Control.2C_PRG_Mode_.28.249000-.249FFF.29"></span>Mirroring Control, PRG Mode (\$9000-\$9FFF) {#Mirroring_Control.2C_PRG_Mode_.28.249000-.249FFF.29}

Mask: \$F000

    7  bit  0
    ---------
    .... ..PM
           ||
           |+- Mirroring  (0: Vertical; 1: Horizontal)
           +-- PRG Mode

When 'P' is clear:

- the 8 KiB page at \$8000 is controlled by the \$800x register
- the 8 KiB page at \$C000 is fixed to the second-to-last 8 KiB in the ROM

When 'P' is set:

- the 8 KiB page at \$8000 is fixed to the second-to-last 8 KiB in the ROM
- the 8 KiB page at \$C000 is controlled by the \$800x register

### <span id="PRG_Select_1_.28.24A000-.24AFFF.29"></span>PRG Select 1 (\$A000-\$AFFF) {#PRG_Select_1_.28.24A000-.24AFFF.29}

Mask: \$F000

    7  bit  0
    ---------
    ...P PPPP
       | ||||
       +-++++- Select 8 KB PRG ROM at $A000-$BFFF

### <span id="CHR-ROM_Bank_Select_.28.24B000-.24BFFF.29"></span>CHR-ROM Bank Select (\$B000-\$BFFF) {#CHR-ROM_Bank_Select_.28.24B000-.24BFFF.29}

Mask: \$F007

    7  bit  0
    ---------
    CCCC CCCC
    |||| ||||
    ++++-++++- Select 1 KiB CHR bank

| Write to CPU addresses | 1KB CHR bank affected |
|------------------------|-----------------------|
| \$B000                 | \$0000-\$03FF         |
| \$B001                 | \$0400-\$07FF         |
| \$B002                 | \$0800-\$0BFF         |
| \$B003                 | \$0C00-\$0FFF         |
| \$B004                 | \$1000-\$13FF         |
| \$B005                 | \$1400-\$17FF         |
| \$B006                 | \$1800-\$1BFF         |
| \$B007                 | \$1C00-\$1FFF         |

## Notes {#Notes}

- Major League wants hardwired 1-screen mirroring. (CIRAM A10 is tied to +5V on this game). Additionally, the register at \$9000 is entirely disabled: the game can only request "PRG mode 0". An [NES 2.0](../NES_2.0/) [submapper](../NES_2.0_submappers/#iNES_Mapper_032_.2F_Irem_G101) has been assigned for this difference, otherwise you'll have to use a hash check.
- Naruko says that only the game "愛先生の OSHIETE 私の星" uses PRG mode 1, which is why it was documented incorrectly for so long.

## See also {#See_also}

- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
- Naruko's notes: [\[1\]](http://w.livedoor.jp/famicomcartridge/d/Irem%20G-101)
