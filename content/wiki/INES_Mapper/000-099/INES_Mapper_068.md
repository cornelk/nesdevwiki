---
title: "INES Mapper 068"
url: "/wiki/INES_Mapper_068/"
description: "iNES Mapper 068 denotes PCBs using the Sunsoft-4 mapper IC. In the US it was only used in the game After Burner. It has the unusual ability to map CHR ROM into the part of the PPU's address space…"
wikiPage: "INES_Mapper_068"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

iNES Mapper 068 denotes PCBs using the [Sunsoft-4 mapper IC](../Sunsoft_4_pinout/). In the US it was only used in the game *After Burner*. It has the unusual ability to map CHR ROM into the part of the PPU's address space used for nametables.

Example games:

- *After Burner*
- *Maharaja (J)*
- *Nantettatte!! Baseball (J)*

## Overview {#Overview}

- PRG ROM size: Up to 256 KiB
- PRG ROM bank size: 16 KiB
- PRG RAM: Up to 8 KiB
- CHR capacity: Up to 256 KiB ROM
- CHR bank size: 2 KiB
- Nametable [mirroring](../Mirroring/): Controlled by mapper
- Subject to [bus conflicts](../Bus_conflict/): No

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB PRG RAM bank
- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last internal bank
- PPU \$0000-\$07FF: 2 KiB switchable CHR bank
- PPU \$0800-\$0FFF: 2 KiB switchable CHR bank
- PPU \$1000-\$17FF: 2 KiB switchable CHR bank
- PPU \$1800-\$1FFF: 2 KiB switchable CHR bank
- PPU \$2000-\$23FF (and \$2400 or \$2800): 1 KiB switchable CHR bank
- PPU \$2C00-\$2FFF (and \$2800 or \$2400): 1 KiB switchable CHR bank

## Registers {#Registers}

### <span id=".246000-.247FFF:_Licensing_IC"></span>\$6000-\$7FFF: Licensing IC {#.246000-.247FFF:_Licensing_IC}

Only for the game *Nantettatte!! Baseball*, when WRAM is disabled (\$F000.4 is 0), writes to this register reset a timer in the [external option ROM cartridge](../Nantettatte%21%21_Baseball_pinout/). The external ROM is only readable while banked into \$8000-\$BFFF and the timer has not expired; once it has, reads from the external ROM should return open bus.

The game verifies that the timer expires after 107516 to 107575 M2 cycles (most likely exactly 107520=1024·105) by reading a signature from the end of the external ROM.

While WRAM is disabled, reads from \$6000-\$7FFF return open bus for *all* games. (The [licensing IC](../Sunsoft_6_pinout/) isn't readable because it is connected to only M2, WRAM +CE, WRAM /CE, and R/W. None of the data pins connect.)

### <span id=".248000-.248FFF:_CHR_Pattern_Table_ROM_bank_0"></span>\$8000-\$8FFF: CHR Pattern Table ROM bank 0 {#.248000-.248FFF:_CHR_Pattern_Table_ROM_bank_0}

Map a 2 KiB CHR ROM bank at PPU address \$0000-\$07FF.

### <span id=".249000-.249FFF:_CHR_Pattern_Table_ROM_bank_1"></span>\$9000-\$9FFF: CHR Pattern Table ROM bank 1 {#.249000-.249FFF:_CHR_Pattern_Table_ROM_bank_1}

Map a 2 KiB CHR ROM bank at PPU address \$0800-\$0FFF.

### <span id=".24A000-.24AFFF:_CHR_Pattern_Table_ROM_bank_2"></span>\$A000-\$AFFF: CHR Pattern Table ROM bank 2 {#.24A000-.24AFFF:_CHR_Pattern_Table_ROM_bank_2}

Map a 2 KiB CHR ROM bank at PPU address \$1000-\$17FF.

### <span id=".24B000-.24BFFF:_CHR_Pattern_Table_ROM_bank_3"></span>\$B000-\$BFFF: CHR Pattern Table ROM bank 3 {#.24B000-.24BFFF:_CHR_Pattern_Table_ROM_bank_3}

Map a 2 KiB CHR ROM bank at PPU address \$1800-\$1FFF.

### <span id=".24C000-.24CFFF:_CHR_Nametable_ROM_bank_0"></span>\$C000-\$CFFF: CHR Nametable ROM bank 0 {#.24C000-.24CFFF:_CHR_Nametable_ROM_bank_0}

Map a 1 KiB CHR ROM bank where CIRAM \$000-\$3FF would normally be mapped.

- This applies only when \$E000.4 = 1 for ROM nametable mode.
- This occupies the lower nametable and pattern table.
- The PPU address range of this bank follows the mirroring selection in \$E000.
- Only D6-D0 are used; D7 is ignored and treated as 1, so nametables must be in the last 128 KiB of CHR ROM.

### <span id=".24D000-.24DFFF:_CHR_Nametable_ROM_bank_1"></span>\$D000-\$DFFF: CHR Nametable ROM bank 1 {#.24D000-.24DFFF:_CHR_Nametable_ROM_bank_1}

Map a 1 KiB CHR ROM bank where CIRAM \$400-\$7FF would normally be mapped.

- This applies only when \$E000.4 = 1 for ROM nametable mode.
- This occupies the upper nametable and pattern table.
- The PPU address range of this bank follows the mirroring selection in \$E000.
- Only D6-D0 are used; D7 is ignored and treated as 1.

### <span id=".24E000-.24EFFF:_Nametable_control"></span>\$E000-\$EFFF: Nametable control {#.24E000-.24EFFF:_Nametable_control}

    7654 3210
       |   ||
       |   ++- Nametable Mirroring Behavior (See table below.)
       +------ CIRAM / CHR-ROM Select for Nametables (PPU $2000-$2FFF):
                 0 = CIRAM
                 1 = CHR-ROM (Uses ROM bank selection registers $C000, $D000.)

#### Nametable Mirroring Behavior {#Nametable_Mirroring_Behavior}

| Mode | \$E000.(1:0) | PPU \$2000-\$23FF | PPU \$2400-\$27FF | PPU \$2800-\$2BFF | PPU \$2C00-\$2FFF |
|----|----|----|----|----|----|
| Vertical | %00 | Low NT | High NT | Low NT | High NT |
| Horizontal | %01 | Low NT | Low NT | High NT | High NT |
| Single-screen NT0 | %10 | Low NT | Low NT | Low NT | Low NT |
| Single-screen NT1 | %11 | High NT | High NT | High NT | High NT |

Note: Nametable [mirroring](../Mirroring/) works the same way in both CIRAM and CHR-ROM modes.

### <span id=".24F000-.24FFFF:_PRG-ROM_bank"></span>\$F000-\$FFFF: PRG-ROM bank {#.24F000-.24FFFF:_PRG-ROM_bank}

Map a 16 KiB PRG ROM bank at CPU address \$8000-\$BFFF.

    7  bit  0
    ---- ----
    ...E BBBB
       | ||||
       | ++++- Select 16 KiB PRG banked into $8000-$BFFF
       +------ 1:Enable PRG RAM = WRAM +CS2

*Nantettatte!! Baseball* repurposes some of the bits:

    7  bit  0
    ---- ----
    ...E RBBB
       | ||||
       | |+++- Select 16 KiB PRG banked into $8000-$BFFF
       | +----   1: select internal ROM
       |         0: select external ROM
       +------ 1: Enable PRG RAM = WRAM +CS2
               0: Enable licensing verification

Note that although the external ROM could be up to 128KiB, both known subcartridges contain 16KiB of data, doubled to fill a 32KiB EPROM, mirrored across the bottom 128KiB.

The fixed bank is always from the internal ROM.

## Hardware {#Hardware}

The US release of *After Burner* has CHR ROM split into [two 28-pin chips](http://bootgod.dyndns.org:7777/profile.php?id=326). The Japanese releases exist both as [two 32-pin 128KiB CHR ROMs with extra enables](http://bootgod.dyndns.org:7777/profile.php?id=3806) and as [one 32-pin 256KiB ROM](http://forums.nesdev.org/viewtopic.php?p=106818#p106818).

## References {#References}

- [Goroh's Sunsoft mapper doc](http://nesdev.org/sunsoft.txt)
- [Naruko's notes in the forum](http://forums.nesdev.org/viewtopic.php?t=9744)
- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
