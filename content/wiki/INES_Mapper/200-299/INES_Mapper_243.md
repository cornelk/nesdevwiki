---
title: "INES Mapper 243"
url: "/wiki/INES_Mapper_243/"
description: "iNES Mapper 243 denotes the Sachen SA-020A circuit board. Using an eight-register ASIC with a fake \"74LS374N\" marking, it supports up to 128 KiB PRG-ROM, and 128 KiB of CHR-ROM. It is used for only…"
wikiPage: "INES_Mapper_243"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
---

**iNES Mapper 243** denotes the Sachen **SA-020A** circuit board. Using an eight-register ASIC with a fake "[74LS374N](../Sachen_74LS374N_pinout/)" marking, it supports up to 128 KiB PRG-ROM, and 128 KiB of CHR-ROM. It is used for only one game, *美女拳 - Honey Peach* (SA-006).

# Banks {#Banks}

- CPU \$8000-\$FFFF: switchable 32 KiB PRG-ROM bank
- PPU \$0000-\$1FFF: switchable 8 KiB CHR-ROM bank
- Nametable mirroring: switchable H/V/Vertically-flipped L/One-screen

# Registers {#Registers}

## <span id="Register_Index_.28.244100.2C_write.29"></span>Register Index (\$4100, write) {#Register_Index_.28.244100.2C_write.29}

    Mask: $C101

    D~7654 3210
      ---------
      .... .RRR
            +++- Select register number (Rx)

## <span id="Register_Data_.28.244101.2C_read.2Fwrite.29"></span>Register Data (\$4101, read/write) {#Register_Data_.28.244101.2C_read.2Fwrite.29}

    Mask: $C101

    D~7654 3210
      ---------
      .... .RRR
            +++- Register data

    D~7654 3210
      ---------
      .... ...A  R2: CHR A13
      .... ...B  R4: CHR A14
      .... ..PP  R5: PRG A16..A15
      .... ..DC  R6: CHR A16..A15
      ...  .MM.  R7: Nametable mirroring
                  0: S0-S0-S0-S1 (lower right unique, or vertically-flipped L)
                  1: Horizontal
                  2: Vertical
                  3: Single-screen, page 1

Registers 0, 1, and 3 have no external effect. All three bits in all eight registers are fully implemented and can be read from and written to.

# Errata {#Errata}

The **SA-150** PCB, denoted by [INES Mapper 150](../INES_Mapper_150/), connects the same ASIC differently, changing the meaning of the CHR-bank-related register bits.

# See also {#See_also}

- [Box, Cart, and PCB picture of *Honey Peach*](https://www.flickr.com/photos/153392699@N08/sets/72157682682439086)
