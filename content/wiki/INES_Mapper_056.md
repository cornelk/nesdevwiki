---
title: "INES Mapper 056"
url: "/wiki/INES_Mapper_056/"
description: "iNES Mapper 056 reportedly represents a specific unlicensed reproduction of Super Mario Bros. 3"
wikiPage: "INES_Mapper_056"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

iNES Mapper 056 reportedly represents a specific unlicensed reproduction of Super Mario Bros. 3

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KB PRG RAM bank (not battery backed!)
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB semi-fixed PRG ROM bank (banks 15 or 31)
- PPU \$0000-\$03FF: 1 KB switchable CHR bank
- PPU \$0400-\$07FF: 1 KB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KB switchable CHR bank
- PPU \$1000-\$13FF: 1 KB switchable CHR bank
- PPU \$1400-\$17FF: 1 KB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KB switchable CHR bank

## Registers {#Registers}

### <span id="IRQ_.28.248000-.24DFFF.29"></span>IRQ (\$8000-\$DFFF) {#IRQ_.28.248000-.24DFFF.29}

The IRQ functionality on this board is implemented by the KS202 ASIC on it. That IC was used in isolation in [iNES Mapper 142](../INES_Mapper_142/), and this mapper's IRQ will behave identically.

### <span id="Bank_select_.28.24E000-.24EFFF.29"></span>Bank select (\$E000-\$EFFF) {#Bank_select_.28.24E000-.24EFFF.29}

Select one of three bank registers to update on next write to \$F000.

    7  bit  0
    ---- ----
    .... .RRR
          |||
          +++- Specify which bank register to update on next write to Bank Data register
               1: Select 8 KB PRG bank at CPU $8000-$9FFF
               2: Select 8 KB PRG bank at CPU $A000-$BFFF
               3: Select 8 KB PRG bank at CPU $C000-$DFFF
               4,6: Writes to $F000 have no effect
               0,5,7: unknown

### <span id="Bank_data_.28.24F000-.24FFFF.29"></span>Bank data (\$F000-\$FFFF) {#Bank_data_.28.24F000-.24FFFF.29}

Note that this register is logically superimposed on the next thirteen registers.

    7  bit  0
    ---- ----
    .... DDDD
         ||||
         ++++- New bank value, based on last value written to Bank select register (mentioned above)

### <span id="PRG_bank_upper_bits_.28.24F000-.24F3FF.29"></span>PRG bank upper bits (\$F000-\$F3FF) {#PRG_bank_upper_bits_.28.24F000-.24F3FF.29}

Note that these registers are superimposed on the Bank data register.

Mask: \$FC03

    $F000: [...P ....] - select PRG ROM A17 during reads from $8000-$9FFF
    $F001: [...P ....] - same, for reads from $A000-$BFFF
    $F002: [...P ....] - same, for reads from $C000-$DFFF
    $F003: [...P ....] - same, for reads from $E000-$FFFF

The game relies on these four registers powering up holding 1.

The original hardware dummies out the upper three bits, but they could have been connected.

The above six registers are intended be used together as

     LDX #2
     STX $E000
     LDA bankA000
     STA $F001

### <span id="Mirroring_control_.28.24F800-.24FBFF.29"></span>Mirroring control (\$F800-\$FBFF) {#Mirroring_control_.28.24F800-.24FBFF.29}

Note that this register is superimposed on the Bank data register.

Mask: \$FC00

    $F800: [.... ...M] - 0: H (PPU A11), 1: V (PPU A10)

### <span id="CHR_banking_control_.28.24FC00-.24FFFF.29"></span>CHR banking control (\$FC00-\$FFFF) {#CHR_banking_control_.28.24FC00-.24FFFF.29}

Note that these registers are superimposed on the Bank data register.

Mask: \$FC07

     $FC00 [.CCCCCCC] - select 1kB CHR for $0000-$03ff
     $FC01 [.CCCCCCC] - select 1kB CHR for $0400-$07ff
     $FC02 [.CCCCCCC] - select 1kB CHR for $0800-$0bff
     $FC03 [.CCCCCCC] - select 1kB CHR for $0c00-$0fff
     $FC04 [.CCCCCCC] - select 1kB CHR for $1000-$13ff
     $FC05 [.CCCCCCC] - select 1kB CHR for $1400-$17ff
     $FC06 [.CCCCCCC] - select 1kB CHR for $1800-$1bff
     $FC07 [.CCCCCCC] - select 1kB CHR for $1c00-$1fff

The original hardware dummies out the top bit, but it could have been connected.

## See also {#See_also}

- [krzysiobal's reverse-engineering notes](https://forums.nesdev.org/viewtopic.php?f=9&t=19314)
- [nocash's documentation](https://problemkaputt.de/everynes.htm#mapper56piratesmb3)
