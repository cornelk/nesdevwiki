---
title: "AxROM"
url: "/wiki/AxROM/"
description: "The generic designation AxROM refers to Nintendo cartridge boards NES-AMROM, NES-ANROM, NES-AN1ROM, NES-AOROM, their HVC counterparts, and clone boards. AxROM and compatible boards are implemented in…"
wikiPage: "AxROM"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>AxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, Rare, others </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=7">34 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>AMROM, ANROM, <br/>AN1ROM, AOROM, others </td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>PRG RAM window</b></td><td>n/a </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>1 page switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>AMROM/AOROM only </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">007</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=7">iNES 007</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-A%25ROM">AxROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=AMROM">AMROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=ANROM">ANROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=AN1ROM">AN1ROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=AOROM">AOROM</a></td></tr></tbody></table></div>

The generic designation **AxROM** refers to Nintendo cartridge boards **NES-AMROM**, **NES-ANROM**, **NES-AN1ROM**, **NES-AOROM**, their [HVC](../Family_Computer/) counterparts, and clone boards. AxROM and compatible boards are implemented in [iNES](../INES/) format with **iNES Mapper 7**.

- **[Submapper number](../NES_2.0_submappers/#002.2C_003.2C_007:_UxROM.2C_CNROM.2C_AxROM)** determines the bus conflict behavior of the game.
  - Submapper 0: Bus conflict behavior undefined, but emulators *shouldn't* emulate bus conflicts
  - Submapper 1: No bus conflicts
  - Submapper 2: AND-type bus conflicts

## Board types {#Board_types}

The following AxROM boards are known to exist:

| Board  | PRG ROM      | Bus conflicts         |
|--------|--------------|-----------------------|
| AMROM  | 128 KB       | Yes                   |
| ANROM  | 128 KB       | No                    |
| AN1ROM | 64 KB        | No                    |
| AOROM  | 128 / 256 KB | Depends on +CE wiring |

## Overview {#Overview}

- PRG ROM size: up to 256 KB
- PRG ROM bank size: 32 KB
- PRG RAM: None
- CHR capacity: 8 KB RAM
- CHR bank size: Not bankswitched
- Nametable [mirroring](../Mirroring/): Single-screen, mapper-selectable
- Subject to [bus conflicts](../Bus_conflict/): AMROM/AOROM only

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank

## Solder pad config {#Solder_pad_config}

No solder pad config is needed on the AxROM board family.

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxxM xPPP
       |  |||
       |  +++- Select 32 KB PRG ROM bank for CPU $8000-$FFFF
       +------ Select 1 KB VRAM page for all 4 nametables

## Hardware {#Hardware}

The AxROM boards contain a [74HC161](../74161/) binary counter used as a quad D latch (4-bit register). The ANROM and AN1ROM boards also contains a [74HC02](../7402/) which is used to disable the PRG ROM during writes, thus avoiding [bus conflicts](../Bus_conflict/).

## Various notes {#Various_notes}

On the AOROM board, special mask ROMs with an additional positive CE on pin 31 (which is connected to PRG R/W) can be used to prevent bus conflicts without an additional chip. Some 128 KB games were made with AOROM to save the cost of a 74HC02. It seems that only *Double Dare* and *Wheel of Fortune* employ this trick noticeably--that is, if emulated with bus conflicts enabled, the games will glitch. [The list of AxROM games in NesCartDB](https://nescartdb.com/search/advanced?ines_op=equal&ines=7) lacks quality coverage of the PCB backs for the AOROM games, so it is hard to determine yet which games may be wired this way.

It is likely that every retail AOROM game could be emulated correctly without emulating bus conflicts.

## Variants {#Variants}

- [BNROM](../INES_Mapper_034/) is the same as AMROM except it uses a fixed horizontal or vertical mirroring.
- [NES 2.0 Mapper 470](../NES_2.0_Mapper_470/) adds an outer bank register to AOROM.
- Some emulators allow bit 3 to be used to select up to 512 KB of PRG ROM for an oversized AxROM. In hardware this could be implemented by using an octal latch in place of the quad latch ([74HC377](../74377/)).
  - The pirate multicart and unlicensed music game *Hot Dance 2000* uses a 512 KB AxROM variant.
  - A ROM Hack of Battletoads expanded the ROM size to 512 KB.

## See also {#See_also}

- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
