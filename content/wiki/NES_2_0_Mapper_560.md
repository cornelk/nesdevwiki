---
title: "NES 2.0 Mapper 560"
url: "/wiki/NES_2.0_Mapper_560/"
description: "NES 2.0 Mapper 560 denotes the C/E BASIC educational cartridge PCB. Its main feature is to address 128 KiB with per-tile bankswitching using 14-bit tile numbers. Two variants exist:"
wikiPage: "NES_2_0_Mapper_560"
---

**NES 2.0 Mapper 560** denotes the **C/E BASIC** educational cartridge PCB. Its main feature is to address 128 KiB with per-tile bankswitching using 14-bit tile numbers. Two variants exist:

- **Submapper 0**: 鸿达 (Hongda)
- **Submapper 1**: 小霸王 (Subor)

# Banks {#Banks}

- CPU \$6000-\$7FFF: Unbanked 8 KiB of PRG-RAM
- CPU \$8000-\$FFFF: Switchable 32 KiB window into 64 KiB of PRG-ROM
- PPU \$0000-\$0FFF: Switchable 4 KiB window into 128 KiB of CHR-ROM (two modes)
- PPU \$1000-\$1FFF: Mirror of PPU \$0000-\$0FFF

Nametable arrangement: horizontal (vertical mirroring)

# <span id="Banking_Modes_.28Submapper_0.29"></span>Banking Modes (Submapper 0) {#Banking_Modes_.28Submapper_0.29}

## Mode 0 {#Mode_0}

This mode is selected for programs lifted from Nintendo's Family BASIC cartridge. This is more-or-less normal PPU CHR mapping, with the exception that PPU A3 is connected to CHR A16, and the other bits shifted accordingly. Because CHR A15..12 are set to 0, the CHR data must come from the beginning of CHR-ROM for bit plane 0, and from the beginning of the second 64 KiB half for bit plane 1.

    During pattern table fetches,
               CHR A..  1 1111 1100 0000 0000
                        6 5432 1098 7654 3210
               ------------------------------
    come from: constant . 0000 .... .... ....
               PPU A..  0 .... 1110 0000 0000
                        3 .... 2109 8765 4210

## Mode 1 {#Mode_1}

This mode is selected for original programs using Chinese characters. Tile numbers are extended from 8 to 14 bits in length; the low eight bits are written to the name table with A10=0, the high eight bits are written to the name table with A10=1; these writes are stored by the cartridge in shadow RAM so it can load the high six bits from shadow RAM when the PPU reads the low eight bits from CIRAM. The upper two bits of shadow RAM data are ignored.

    During pattern table fetches,
               CHR A..      1 1111 1100 0000 0000
                            6 5432 1098 7654 3210
               ----------------------------------
    come from: PPU A..      . .... .110 0000 0000
                            . .... .109 8765 4210
               Tile MSB D.. 5 4321 0

PPU A2..0 represent the scanline number (as usual), PPU A3 is ignored (so pixel colors are limited to values 0 or 3), PPU A11..4 represent the low eight bits of the tile number (as usual), Tile MSB D5..0 represent the MSB bits of the tile number latched from Shadow RAM when the PPU reads CIRAM.

# <span id="Banking_Modes_.28Submapper_1.29"></span>Banking Modes (Submapper 1) {#Banking_Modes_.28Submapper_1.29}

## Mode 0 {#Mode_0_2}

In Subor's variant, mode 0 uses normal PPU operation with pattern data coming from the beginning of CHR-ROM.

## Mode 1 {#Mode_1_2}

Subor's variant is similar to Hongda's variant except that the bit shift is a bit different, and more importantly, the upper two bits of shadow RAM that are written to CIRAM with A10=1 are not ignored but select the bitplanes in which the character data appears:

    D~[01TT TTTT]
       ||++-++++-- Tile MSB
       |+-- 1: Return actual pattern data if PPU A8=1, 0: Return $00 if PPU A8=1
       +--- 1: Return actual pattern data if PPU A8=0, 0: Return $00 if PPU A8=0
    If neither bit 6 nor bit 7 are set, shadow RAM is ignored and regular PPU operation is used, just like mode 0!

    During pattern table fetches,
               CHR A..      1 1111 1100 0000 0000
                            6 5432 1098 7654 3210
               ----------------------------------
    come from: PPU A..      . .... 1100 0000 .000
                            . .... 1098 7654 .210
               Tile MSB D.. 4 3210 .... .... 5...

PPU A2..0 represent the scanline number (as usual), PPU A3 is ignored (so pixel colors are decided by the upper two bits of the Shadow RAM byte), PPU A11..4 represent the low eight bits of the tile number (as usual), Tile MSB D5..0 represent the MSB bits of the tile number latched from Shadow RAM when the PPU reads CIRAM.

# <span id="Register:_Mode_and_PRG-ROM_bank_.28.248000-.24FFFF.2C_write.29"></span>Register: Mode and PRG-ROM bank (\$8000-\$FFFF, write) {#Register:_Mode_and_PRG-ROM_bank_.28.248000-.24FFFF.2C_write.29}

Any write to the PRG-ROM region switches between the two 32 KiB PRG-ROM banks and the two CHR banking modes, regardless of value and address. On reset, bank/mode 1 (0-based) is selected.

# See also {#See_also}

[PCB picture and analysis of submapper 0](https://forums.nesdev.org/viewtopic.php?t=23032)
