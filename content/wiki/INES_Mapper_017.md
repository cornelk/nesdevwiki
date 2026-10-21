---
title: "INES Mapper 017"
url: "/wiki/INES_Mapper_017/"
description: "iNES Mapper 017 denotes ROM images that have been extracted from disk images for the Front Fareast Super Magic Card RAM cartridge. They represent games whose Doctor Header file denotes a Super Magic…"
wikiPage: "INES_Mapper_017"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
---

**iNES Mapper 017** denotes ROM images that have been extracted from disk images for the *Front Fareast Super Magic Card* [RAM cartridge](../RAM_cartridge/). They represent games whose [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Super_Magic_Card_disks) denotes a Super Magic Card disk (byte \$0 bit 7 set, byte \$7=\$AA). Refer to the [Super Magic Card](../Super_Magic_Card/) article for details on bankswitching. The Super Magic Card's registers are initialized to:

    ; Play mode, WRAM bank 0, 1 KiB CHR mode enabled
    $4500 = $47 

    ; PRG memory write-protected, two-screen mirroring
    $42FF = $20 | (verticalMirroring? 0x00: 0x10)

    ; 4M banking mode enabled
    $43FC = $00

    ; Initial PRG register content
    $4504 = Number of 8 KiB PRG banks -4
    $4505 = Number of 8 KiB PRG banks -3
    $4506 = Number of 8 KiB PRG banks -2
    $4507 = Number of 8 KiB PRG banks -1 (originally Doctor Header file byte $5)

The [iNES](../INES/) header may specify a [512-byte trainer](../INES/#Trainer) (corresponding to [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Super_Magic_Card_disks)'s byte \$0 bit 6 being set). The trainer must be loaded to an address originally denoted by the [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Super_Magic_Card_disks)'s byte \$2, and is here denoted by the NES 2.0 submapper. In its presence, instead of jumping to the game's reset vector on a hard reset, trainer offset +\$000 must be JMPed to.

    Submapper    Trainer load address
     0           $7000
     1           $5D00
     2           $5E00
     3           $5F00

Battery-saving of WRAM content is not supported by any Magic Card model. Hard-resetting a game while restoring previously-saved WRAM content in emulators interferes with the correct operation of the trainer's program.
