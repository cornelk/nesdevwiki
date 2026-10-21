---
title: "Color Dreams"
url: "/wiki/Color_Dreams/"
description: "The Color Dreams Mapper was a mapper used by the Color Dreams company. The iNES format assigns mapper 11 to these boards."
wikiPage: "Color_Dreams"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>Color Dreams</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Color Dreams/Wisdom Tree </td></tr><tr><td><b>Boards</b></td><td>COLORDREAMS </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>128K </td></tr><tr><td><b>CHR window</b></td><td>8k </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">011</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=11">iNES 011</a></td></tr></tbody></table></div>

The **Color Dreams Mapper** was a mapper used by the **Color Dreams** company. The [iNES](../INES/) format assigns **mapper 11** to these boards.

The mapper is a board with [discrete logic](../../category/Discrete_logic_mappers/) that provides up to four 32 KB banks of PRG ROM, up to sixteen 8 KB banks of CHR ROM, and controls a charge pump to defeat the [CIC](../CIC_lockout_chip/) lockout chip.

**Color Dreams** was a company that developed and published unlicensed NES games. They also operated with the alternate labels **Wisdom Tree** which produced games with Christianity themes (primarily sold in Christian bookstores), and **Bunch Games** which produced lower budget games. American Game Cartridges also used the Color Dreams mapper for their port of Exidy's *Chiller*.

Example games:

- *Bible Adventures*
- *Chiller*
- *Crystal Mines*
- *Menace Beach*
- *Metal Fighter*

**Submappers** denote whether the circuit board has bus conflicts:

- Submapper 0: Bus conflict behavior unknown: default to "no bus conflicts" for safety
- Submapper 1: No bus conflicts: circuit board disables PRG-ROM when the CPU is writing
- Submapper 2: AND-type bus conflicts: the value applied during bank select is the written value AND the ROM content at the written address

All commercially-released games are subject to bus conflicts (submapper 2), while a few game prototypes have bus conflict prevention logic on their circuit boards, notably the *Free Fall* prototype.

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    CCCC LLPP
    |||| ||||
    |||| ||++- Select 32 KB PRG ROM bank for CPU $8000-$FFFF
    |||| ++--- Used for lockout defeat
    ++++------ Select 8 KB CHR ROM bank for PPU $0000-$1FFF

## Hardware {#Hardware}

The Color Dreams board contains a [74LS377](../74377/) octal D latch (8-bit register) to select the current PRG and CHR banks.

The lockout defeat is not necessary for emulator implementation. See [references](#References) below for more information.

## Variants {#Variants}

Nintendo's own [GxROM](../GxROM/) family of boards is nearly identical in function to the Color Dreams board.

A multicart mapper based on this mapper is [iNES Mapper 046](../INES_Mapper_046/), with an outer bank register at \$6000-\$7FFF.

Some variations of this board/mapper (e.g the one used in the prototype game "Free Fall") appear to be free of [bus conflicts](../Bus_conflict/) and will not work properly if bus conflicts are emulated.

## References {#References}

- [Color Dreams](http://kevtris.org/mappers/cdreams/) - documentation of mapper by Kevin Horton
- [Lockout](http://kevtris.org/mappers/lockout/) - documentation of lockout mechanisms by Kevtris
- [Games](http://bootgod.dyndns.org:7777/search.php?ines=11) - list of games using the Color Dreams mapper at NesCartDB
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
