---
title: "NES 2.0 Mapper 400"
url: "/wiki/NES_2.0_Mapper_400/"
description: "NES 2.0 Mapper 400 denotes retroUSB's 8-BIT XMAS 2017 homebrew cartridge. It's basically a multicart version of the \"Christmas lights\" variant of UNROM 512 (mapper 30 submapper 4) with the addition…"
wikiPage: "NES_2_0_Mapper_400"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 400** denotes retroUSB's *8-BIT XMAS 2017* homebrew cartridge. It's basically a multicart version of the "Christmas lights" variant of [UNROM 512](../UNROM_512/) (mapper 30 submapper 4) with the addition of a TFT LCD in place of a cartridge label.

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable inner and outer window \#1 into 4 MiB of PRG-ROM
- CPU \$C000-\$FFFF: 16 KiB fixed inner and switchable outer window \#2 into 4 MiB of PRG-ROM, inner bank fixed at \#7
- PPU \$0000-\$1FFF: 8 KiB switchable window into 32 KiB of CHR-RAM
- PPU \$2000-\$3FFF: Switchable horizontally-arranged/vertically-arranged console-internal nametables

# Registers {#Registers}

## <span id="Inner_PRG-ROM_bank.2C_CHR-RAM_bank"></span>Inner PRG-ROM bank, CHR-RAM bank {#Inner_PRG-ROM_bank.2C_CHR-RAM_bank}

    D~[.CC. .PPP] A~[11.. .... .... ....] ($C000-$FFFF, no bus conflicts)
        ||   |||
        ||   +++- Inner 16 KiB PRG-ROM bank# at CPU $8000-$BFFF (PRG A16..14)
        ++------- 8 KiB CHR-RAM bank#

## Lower Outer PRG-ROM bank {#Lower_Outer_PRG-ROM_bank}

    D~[PPPP P...] A~[0111 0??? ??00 0000] ($7000)
       ++++-+----- Outer 128 KiB PRG-ROM bank# at CPU $8000-$BFFF (PRG A21..17)

The lower address bits are checked as well, because *8-Bit XMAS 2008* performs APU writes to \$7001-\$7015, oddly enough, which must be ignored.

## <span id="Upper_Outer_PRG-ROM_bank.2C_Nametable_arrangement"></span>Upper Outer PRG-ROM bank, Nametable arrangement {#Upper_Outer_PRG-ROM_bank.2C_Nametable_arrangement}

    D~[PPPP P...] A~[0111 1???? ????] ($7800)
       ++++-+----- Outer 128 KiB PRG-ROM bank# at CPU $C000-$FFFF (PRG A21..17)
         +-------- Nametable Arrangement
                    0: Horizontal arrangement/Vertical mirroring
                    1: Vertical arrangement/Horizontal mirroring

## Cartridge LEDs {#Cartridge_LEDs}

    D~[BYGR bygr] A~[10.. .... .... ....] ($8000-$BFFF, no bus conflicts)
       |||| |||+- Red 1, Red 2
       |||| ||+-- Green 1, Green 2
       |||| |+--- Yellow 1, Yellow 2
       |||| +---- Blue 1, Blue 2
       |||+------ Red 3, Red, 4
       ||+------- Green 3, Green 4
       |+-------- Yellow 3, Yellow 4
       +--------- Blue 3, Blue 4

LEDs are lit when the respective bit is clear and dark when the respective bit is set. Arrangement of LEDs and their respective register bit values:

    +-----------------+  +-----------------+
    |01             20|  |R1             G3|
    |04             80|  |Y1             B3|
    |08             40|  |B1             Y3|
    |02             10|  |G1             R3|
    |01             20|  |R2             G4|
    |04             80|  |Y2             B4|
    |   02 08 40 10   |  |   G2 B2 Y4 R4   |
    +--+           +--+  +--+           +--+
       +-----------+        +-----------+

## LCD Controller {#LCD_Controller}

**Command port**

    D~[.... ..DC] A~[0110 1... .... ....] ($6800-$6FFF)
              |+-- ILI9341 CSX pin
              +--- ILI9341 D/CX pin

**Data port**

    D~[DDDD DDDD] A~[???? ???? ???? ????] ($00FF)

The menu program writes \$00 to \$6800, then a command byte to \$00FF, followed by \$02 to \$6800 and a series of data bytes to \$00FF. It finishes by writing \$03 to \$6800. Refer to the ILI9341's data sheet for a list of commands. The cartridge drives the LCD controller in 8080 MCU 8-bit bus interface mode. For simple emulation purposes, only watch for command byte \$2C with 163,840 data bytes. The data consists of a 240x320 bitmap (with padding) in 16-bit big endian format, with a pixel represented as 5-6-5 RGB, bit order RRRR'RGGG'GGGB'BBBB. Note that horizontally, the data is stored from right to left.

# References {#References}

[ILI9341 data sheet](https://cdn-shop.adafruit.com/datasheets/ILI9341.pdf)
