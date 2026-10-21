---
title: "INES Mapper 078"
url: "/wiki/INES_Mapper_078/"
description: "iNES Mapper 078 was used for two games: Holy Diver and Uchuusen - Cosmo Carrier."
wikiPage: "INES_Mapper_078"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

iNES Mapper 078 was used for two games: *Holy Diver* and *Uchuusen - Cosmo Carrier*.

## Overview {#Overview}

- PRG ROM size: 128 KiB
- PRG ROM bank size: 16 KiB
- PRG RAM: No
- CHR capacity: 128KiB ROM
- CHR bank size: 8 KiB
- Nametable mirroring: Mapper controlled
- Subject to [bus conflicts](../Bus_conflict/): Yes

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: Fixed to last 16KiB of PRG ROM
- PPU \$0000-\$1FFF: 8KiB switchable CHR ROM bank

## Registers {#Registers}

### <span id="Bank_Select_.28.248000-.24FFFF.29"></span>Bank Select (\$8000-\$FFFF) {#Bank_Select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    CCCC MPPP
    |||| ||||
    |||| |+++-- Select 16 KiB PRG ROM bank for CPU $8000-$BFFF
    |||| +----- Mirroring.  Holy Diver: 0 = H, 1 = V.  Cosmo Carrier: 0 = 1scA, 1 = 1scB.
    ++++------- Select 8KiB CHR ROM bank for PPU $0000-$1FFF

The two mirroring modes are not compatible, and trying to run either game with the other's hardware will produce major graphical glitching or lockups.

iNES1 ROM image headers often set the "alternative nametables" flag for Holy Diver and cleared it for Cosmo Carrier. NES2.0 uses submappers.

## Hardware {#Hardware}

In both cases, the hardware implementation is simply a [CNROM](../CNROM/) and a [UNROM](../UxROM/) put together. Two [74HC161](../74161/) are used to form an 8 bit latch, with 4 bits connecting the CHR ROM's high address lines, and 3 bits connecting to a [74HC32](../7432/), which provides the logic to fix the top bank of PRG ROM while swapping the first 16KiB. The last latched bit is connected directly to CIRAM A10 on Cosmo Carrier, as in [AxROM](../AxROM/). On Holy Diver, that bit is connected to a [74HC00](../7400/), which serves as a 2:1 mux of PPU A10 and PPU A11 into CIRAM A10 to provide the desired mirroring. Holy Diver also contains a 74HC245 to compensate for a slow CHR ROM chip.

For homebrew, both Nestopia-1.4.0 and FCEUX-2.1.5 default to the 1scA/1scB mirroring provided by Uchuusen. Emulators SHOULD support [NES 2.0 submappers](../NES_2.0_submappers/): 1 for Uchuusen and 3 for Holy Diver.

## See also {#See_also}

- [NES Mapper list](http://www.romhacking.net/documents/362/) by Disch.
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ Information on mapper's initial state is innacurate.
