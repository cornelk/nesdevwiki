---
title: "NES 2.0 Mapper 429"
url: "/wiki/NES_2.0_Mapper_429/"
description: "NES 2.0 Mapper 429 denotes both the Milowork FCFC1 flash cartridge (in A mode) and several similar multicart PCBs."
wikiPage: "NES_2_0_Mapper_429"
---

NES 2.0 Mapper 429 denotes both the Milowork FCFC1 flash cartridge (in A mode) and several similar multicart PCBs.

## Overview {#Overview}

- PRG ROM size: 512 KB (or 1024 KB, FCFC1 DIY)
- PRG ROM bank size: 32 KB
- PRG RAM: None
- CHR capacity: 32 KB RAM, optional battery-backup or FeRAM
- CHR bank size: 8 KB
- Nametable [mirroring](../Mirroring/): vertical or horizontal selected by switch, or fixed (or Single-screen, FCFC1 DIY)
- Subject to [bus conflicts](../Bus_conflict/): No (submapper 0-1), Yes (submapper 2)

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

#### Regular {#Regular}

    7  bit  0
    ---- ----
    xpPP PPCC
     ||| ||||
     ||| ||++- Select 8 KB CHR RAM bank for PPU $0000-$1FFF, initial bank 0
     +++-++--- Select 32 KB PRG ROM bank for CPU $8000-$FFFF, initial bank 1

#### <span id="Submapper_.231"></span>Submapper \#1 {#Submapper_.231}

    7  bit  0
    ---- ----
    MpPP PPCC
    |||| ||||
    |||| ||++- Select 8 KB CHR RAM bank for PPU $0000-$1FFF, initial bank 0
    |+++-++--- Select 32 KB PRG ROM bank for CPU $8000-$FFFF, initial bank 1
    +--------- Same as AxROM

#### <span id="Submapper_.232_.28810831_PCB.29"></span>Submapper \#2 (810831 PCB) {#Submapper_.232_.28810831_PCB.29}

    7  bit  0
    ---- ----
    .... BACC
         ||||
         ||++- CHR A14..A13
         ++--- CHR A15=A*B; meaning 1 if A=1 and B=1, otherwise 0
         ++--- Select 32 KB PRG ROM bank for CPU $8000-$FFFF, initial bank 0

## FCFC1 Details {#FCFC1_Details}

- The A mode of FCFC1 has compatibility with [CNROM](../CNROM/)/[GxROM](../GxROM/), but the compatibility is fragile, some games will switch back to the page where the firmware is located, causing problems such as crashes.
- The B mode of FCFC1 is a [BNROM](../INES_Mapper_034/)/[AxROM](../AxROM/) with initial bank 4 and no [bus conflict](../Bus_conflict/), and has compatibility with [NROM](../NROM/)/[UNROM](../UxROM/).
- To simplify the circuit implementation, the actual location of the page in ROM is shuffled.
- It is recommended to save the preserved data to a page other than 0 of CHR RAM, as this avoids being overwritten by the kazzo programming process, and it can also make the bank selection play a role similar to the WRAM lock bit.
- Some FCFC1 sold as game carriers do not have switches, or the switches do not have the W position. These FCFC1 are not compatible with kazzo.

| Register value causes error | CNROM (A) | GxROM (A) | NROM (B) (Except using CHR RAM) |
|----|----|----|----|
| Firmware 2.1 and before | xxx1xx | xxx1xx | xxx1xx |
| Firmware 2.2 and after | xx01xx | xx01xx | xx01xx |
