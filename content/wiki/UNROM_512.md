---
title: "UNROM 512"
url: "/wiki/UNROM_512/"
description: "UNROM 512 is a discrete-logic board made by RetroUSB as an extension of UNROM with up to 512kB of PRG ROM, bankable CHR RAM, an option for mapper-controlled single-screen nametables, as well as a…"
wikiPage: "UNROM_512"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with flash save"
    url: "/category/Mappers_with_flash_save/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>UNROM 512</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>RetroUSB, InfiniteNESLives, Broke Studio </td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>UNROM-512-8, UNROM-512-16, UNROM-512-32, INL-D-RAM, UNROM-512-F </td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K or 512K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 16K fixed ($C000) </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>32K </td></tr><tr><td><b>CHR window</b></td><td>8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, 1, or 4 by solder pad. 1 permits choosing which half. H/V-switchable variant also exists. </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes (in non-flashable config), else No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">030</a></td></tr></tbody></table></div>

**UNROM 512** is a discrete-logic board made by RetroUSB as an extension of [UNROM](../UxROM/) with up to 512kB of PRG ROM, bankable CHR RAM, an option for mapper-controlled single-screen nametables, as well as a self-flashable configuration for rewriting PRG. Since its initial creation InfiniteNESLives has replicated its design and also added an optional four-screen nametables variation.

The [iNES](../INES/) format assigns **iNES Mapper 030** to UNROM 512. The [UNIF](../UNIF/) names for this board are **UNROM-512-8**, **UNROM-512-16** and **UNROM-512-32**, depending on how much CHR RAM is present. The following submappers are assigned:

- **Submapper 0**: Original definition. The Battery bit must be specified if the game relies on bus conflicts *not* being present, even if the game does not use the self-flashing capability; without the Battery bit set, bus conflicts are assumed to be present.
- **Submapper 1**: No bus conflicts. The Battery bit is only set if the game makes use of the self-flashing capability.
- **Submapper 2**: Bus conflicts present. Since bus conflicts interfere with the self-flashing procedure, the Battery bit cannot be set with submapper 2.
- **Submapper 3**: No bus conflicts, mapper-controlled horizontal/vertical nametable arrangement. Only used for *Mega Man II (30th Anniversary Edition)*.
- **Submapper 4**: No bus conflicts, an additional register controls LEDs inside the cartridge shell.

This means that when converting from a [NES 2.0](../NES_2.0/)-headered file specifying submapper 1 to [iNES](../INES/) which does not support submappers, the Battery bit must be set. [NES 2.0](../NES_2.0/) may specify the CHR-RAM size from 8 to 32 kB, but the default for iNES should to be use 32 kB.

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable window into 512 KiB of PRG-ROM
- CPU \$C000-\$FFFF: 16 KiB fixed window into the last 16 KiB of PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB switchable window into 32 KiB of CHR-RAM
- PPU \$2000-\$3FFF: Hard-wired horizontally-arranged/vertically-arranged/1-screen console-internal nametables (submappers 0,1,2,4), switchable horizontally-arranged/vertically-arranged console-internal nametables (submapper 3)

Or, in the "4-screen with cartridge VRAM" arrangement:

- PPU \$2000-\$23FF: Cartridge VRAM Nametable 0
- PPU \$2400-\$27FF: Cartridge VRAM Nametable 1
- PPU \$2800-\$2BFF: Cartridge VRAM Nametable 2
- PPU \$2C00-\$2FFF: Cartridge VRAM Nametable 3
- PPU \$3000-\$3EFF: Scratch RAM
- PPU \$3F00-\$3FFF: Palette RAM (32 bytes), mirrored seven times

The cartridge VRAM for the "4-screen with cartridge VRAM" comes from the last 8 KiB of CHR-RAM.

# Nametable Configuration {#Nametable_Configuration}

For **submappers 0-2** and **4**, the nametable arrangement bits in byte 6 of the [iNES](../INES/) header select one of four configurations of nametables:

- `%....0..0` - Vertical arrangement
- `%....0..1` - Horizontal arrangement
- `%....1..0` - 1-Screen, switchable
- `%....1..1` - 4-Screen, cartridge VRAM

For **submapper 3**, horizontal versus vertical nametable arrangement is determined by bit 7 of the [latch register](#PRG-ROM_bank.2C_CHR-RAM_bank.2C_Nametable_arrangement). The iNES header's mapper bits are ignored and should be set to 0.

# Registers {#Registers}

## <span id="PRG-ROM_bank.2C_CHR-RAM_bank.2C_Nametable_arrangement"></span>PRG-ROM bank, CHR-RAM bank, Nametable arrangement {#PRG-ROM_bank.2C_CHR-RAM_bank.2C_Nametable_arrangement}

    D~[NCCP PPPP] A~[1... .... .... ....] ($8000-$FFFF) (submapper 0 w/o battery bit; submapper 2)
       |||| ||||  A~[11.. .... .... ....] ($C000-$FFFF) (submapper 0 with battery bit; submappers 1,3,4)
       |||+-++++- 16 KiB PRG-ROM bank# at CPU $8000-$BFFF
       |++------- 8 KiB CHR-RAM bank#
       +--------- Nametable arrangement
                  Submapper 0,1,2,4 with 1-screen mirroring denoted in iNES header:
                   0: Lower bank (CIRAM A10=0)
                   1: Upper bank (CIRAM A10=1)
                  Submapper 3:
                   0: Vertical arrangement/Horizontal mirroring (CIRAM A10=PPU A11)
                   1: Horizonal arrangement/Vertical mirroring (CIRAM A10=PPU A10)

## Flash ROM Save {#Flash_ROM_Save}

    D~[.... ....] A~[10.. .... .... ....] ($8000-$BFFF) (submappers 0,1,4 with battery bit)

Write a sequence of bytes here for writing to flash ROM. The sequence is as follows (This code must exist in RAM):

    Erase 4KB Flash Sector:
    $C000:$01, $9555:$AA
    $C000:$00, $AAAA:$55
    $C000:$01, $9555:$80
    $C000:$01, $9555:$AA
    $C000:$00, $AAAA:$55
    $C000:BANK, ADDR:$30 (Where BANK is $00-1F, ADDR is $8000,$9000,$A000,$B000)

Read the written location twice until you get \$FF twice.

Write a byte:

    $C000:$01, $9555:$AA
    $C000:$00, $AAAA:$55
    $C000:$01, $9555:$A0
    $C000:BANK, ADDR:DATA (Where BANK is $00-1F, ADDR is $8000-$BFFF, DATA is $00-$FF)

Read the written location twice until you get DATA twice.

## Cartridge LEDs {#Cartridge_LEDs}

    D~[BYGR bygr] A~[10.. .... .... ....] ($8000-$BFFF) (submapper 4)
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

# Software support {#Software_support}

Supported in FCEUX as of r3071, Supported in BizHawk as of r6322

Supplemental mappers exist for flash carts as well:

- PowerPak: [forum post](https://forums.nesdev.org/viewtopic.php?p=236715#p236715) (limited flash save)
- Everdrive N8: [forum post](https://forums.nesdev.org/viewtopic.php?p=235457#p235457) (no flash save)

The PowerPak flash save capability lacks the sector erase and chip erase capabilities, due to the underlying hardware being RAM at the time of write. As long as the game writes every byte it intends to read back, this is a compatible implementation.

The above implementations currently only recognize submapper 0. Mesen supports the full implementation as of 2025-06-16.

# Secondary use for converted .WXN files {#Secondary_use_for_converted_.WXN_files}

In GoodNES 3.23b, a few ROMs using the Waixing FS005 board ([INES Mapper 176](../INES_Mapper_176/)) are also set to Mapper 30. These ROMs were converted from Waixing's proprietary WXN/MFC format, which has a 16 byte header that is somewhat similar but not the same as the [INES](../INES/) header, and when interpreting this header's data as if it were an iNES header, the resulting mapper number would be 30. The proper mapper for these games is [176 submapper 2](../INES_Mapper_176/).

# See Also {#See_Also}

- [Programming UNROM](../Programming_UNROM/) - For working around bus conflicts.
- [GTROM](../GTROM/) - A very similar homebrew mapper, but with 32KiB PRG banking.

# References {#References}

- [INES Mapper DischDocs](../INES_Mapper_DischDocs/)
- Data sheet for [SST39SF040](http://www.microchip.com/wwwproducts/Devices.aspx?product=SST39SF040) flash ROM chip
- Example flash writing code: [Flash.asm](https://www.nesdev.org/w/images/default/f/fa/Flash.asm.txt)
