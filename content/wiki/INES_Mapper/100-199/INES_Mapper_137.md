---
title: "INES Mapper 137"
url: "/wiki/INES_Mapper_137/"
description: "iNES Mapper 137 is used to represent the Sachen 8259D board. It is similar to the Sachen 8259A-C boards, but has some important differences. The information here comes primarily from FCEU-mm and may…"
wikiPage: "INES_Mapper_137"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 137 is used to represent the Sachen 8259D board. It is similar to the [Sachen 8259](../Sachen_8259/)A-C boards, but has some important differences. The information here comes primarily from FCEU-mm and may have some inaccuracies.

## Overview {#Overview}

- PRG ROM size: up to 128 KiB
- PRG ROM bank size: 32 KiB
- PRG RAM: No
- CHR capacity: up to 32KiB ROM
- CHR bank size: 1 KiB / 4KiB
- Nametable mirroring: Mapper controlled
- Subject to [bus conflicts](../Bus_conflict/): No

## Variants {#Variants}

There is only one known 8259D game, *The Great Wall*. It has 32KiB PRG ROM and 32KiB CHR ROM. It has been seen in the wild as iNES 137 and UNIF "UNL-Sachen-8259D".

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32KiB switchable PRG ROM bank
- PPU \$0000-\$03FF: 1KiB switchable CHR ROM bank
- PPU \$0400-\$07FF: 1KiB switchable CHR ROM bank
- PPU \$0800-\$0BFF: 1KiB switchable CHR ROM bank
- PPU \$0C00-\$0FFF: 1KiB switchable CHR ROM bank
- PPU \$1000-\$1FFF: fixed to last 4KiB of CHR ROM

## Registers {#Registers}

Registers should be masked with \$C101.

### <span id="Register_Select_.28.244100.29"></span>Register Select (\$4100) {#Register_Select_.28.244100.29}

    7  bit  0
    ---- ----
    .... .rrr
          |||
          +++- Select register number for next data write

### <span id="Register_Data_.28.244101.29"></span>Register Data (\$4101) {#Register_Data_.28.244101.29}

    7  bit  0
    ---- ----
    .... .ddd
          |||
          +++- Output data to register selected by $4100.

### <span id="CHR_Select_.28Internal_0-3.29"></span>CHR Select (Internal 0-3) {#CHR_Select_.28Internal_0-3.29}

    7  bit  0
    ---- ----
    .... .ccc
          |||
          +++- Select lower 3 bits of 1KiB CHR block for PPU $0000/$0400/$0800/$0C00

### <span id="CHR_Top_Bits_.28Internal_4.29"></span>CHR Top Bits (Internal 4) {#CHR_Top_Bits_.28Internal_4.29}

    7  bit  0
    ---- ----
    .... .DCB
          |||
          ||+- Select bit #4 for CHR register 1
          |+-- Select bit #4 for CHR register 2
          +--- Select bit #4 for CHR register 3

### <span id="PRG_Bank_.28Internal_5.29"></span>PRG Bank (Internal 5) {#PRG_Bank_.28Internal_5.29}

    7  bit  0
    ---- ----
    .... .ppp
          |||
          +++- Select 32KiB PRG bank for CPU $8000

### <span id="CHR_3_Bit_.233_.28Internal_6.29"></span>CHR 3 Bit \#3 (Internal 6) {#CHR_3_Bit_.233_.28Internal_6.29}

    7  bit  0
    ---- ----
    .... ...E
            |
            +- Select bit #3 for CHR register 3

Putting all of this together, the 4 CHR registers have the following bit layouts:

    7  bit  0
    ---- ----
    ...0 0ccc  CHR register 0
    ...B 0ccc  CHR register 1
    ...C 0ccc  CHR register 2
    ...D Eccc  CHR register 3

So only CHR register 3 can access all possible locations in CHR ROM.

### <span id="Mode_and_Mirroring_Select_.28Internal_7.29"></span>Mode and Mirroring Select (Internal 7) {#Mode_and_Mirroring_Select_.28Internal_7.29}

The only game only writes 0x2 to this register (i.e. mm=1, s=0) and wants vertical mirroring/horizontal layout. Underdocumentation has caused someone to set the "4-screen" bit in the header for *The Great Wall*, but this is unnecessary and erroneous.

The below is deduction, and differs from the normal Sachen 8259 behavior, because some rewiring is necessary for the reduced bank size.

    7  bit  0
    ---- ----
    .... .mms
          |||
          ||+- Enable "simple" mode. (mirroring is fixed to H, and banks become weird)
          ++-- Select mirroring (0 = H, 1 = V, 2 = (0,1,1,1), 3 = 1scA)

The upper two bits of the address are still controlled by the mapper when in "simple" mode, even though the three LSBs now all come from register 0. "Simple" mode doesn't affect the behavior of the external 74-series ICs.

## See also {#See_also}

- <http://cah4e3.shedevr.org.ru/%5Blst%5D-sachen-mappers.txt>
