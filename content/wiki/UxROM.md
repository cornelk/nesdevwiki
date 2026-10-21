---
title: "UxROM"
url: "/wiki/UxROM/"
description: "The generic designation UxROM refers to the Nintendo cartridge boards NES-UNROM, NES-UOROM, their HVC counterparts, HVC-UN1ROM, and clone boards."
wikiPage: "UxROM"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>UxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, others </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=2">155 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>UNROM, UOROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K/4096K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes/No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">002</a>, <a href="../INES_Mapper_094/">094</a>, <a href="../INES_Mapper_180/">180</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=2">iNES 002</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-U%25ROM">UxROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=UNROM">UNROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=UOROM">UOROM</a></td></tr></tbody></table></div>

The generic designation *UxROM* refers to the Nintendo cartridge boards *NES-UNROM*, *NES-UOROM*, their [HVC](../Family_Computer/) counterparts, [HVC-UN1ROM](../INES_Mapper_094/), and clone boards.

- **[iNES](../INES/) Mapper 002** is the implementation of the most common usage of UxROM compatible boards, described in this article.
  - Submapper 0: Bus conflict behavior unknown
  - Submapper 1: No bus conflicts
  - Submapper 2: AND-type bus conflicts
- [iNES Mapper 094](../INES_Mapper_094/) describes UN1ROM, used only in *Senjou no Ookami*.
- [iNES Mapper 180](../INES_Mapper_180/) describes a reconfiguration of UNROM used only in *Crazy Climber*.

Example games:

- *Mega Man*
- *Castlevania*
- *Contra*
- *Duck Tales*
- *Metal Gear*

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank

## Solder pad config {#Solder_pad_config}

- Horizontal mirroring : 'H' disconnected, 'V' connected.
- Vertical mirroring : 'H' connected, 'V' disconnected.

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxxx pPPP
         ||||
         ++++- Select 16 KB PRG ROM bank for CPU $8000-$BFFF
              (UNROM uses bits 2-0; UOROM uses bits 3-0)

Emulator implementations of **iNES mapper 2** treat this as a full 8-bit bank select register, without bus conflicts. This allows the mapper to be used for similar boards that are compatible.

To make use of all 8-bits for a 4 MB PRG ROM, an [NES 2.0](../NES_2.0/) header must be used ([iNES](../INES/) can only effectively go to 2 MB).

The original UxROM boards used by Nintendo were subject to bus conflicts, and the relevant games all work around this in software. Some emulators (notably FCEUX) will have bus conflicts by default, but others have none. [NES 2.0 submappers](../NES_2.0_submappers/#002.2C_003.2C_007:_UxROM.2C_CNROM.2C_AxROM) were assigned to accurately specify whether the game should be emulated with bus conflicts.

## Hardware {#Hardware}

The UNROM, UN1ROM, and UOROM boards contain a [74HC161](../74161/) binary counter used as a quad D latch (4-bit register) and a [74HC32](../7432/) quad 2-input OR gate to make one bank always visible.

The circuit behaves as if it were as follows:

          /PRGSEL               A14  A13-A0
           |                      |      |
           |                      |      |
           | D2-D0-.       ,------'      |
           |       |     . |             |
           | ,-----+--.  |`+.            |
           | |Register+--+0  `.          |
           | `--------'  |    |__.       |
           |       |     |    |  |       |
           | R/W --'   7-+1  ,'  |       |
           |             | ,'    |       |
           |             |'      |       |
           |                     |       |
          ,+---------------------+-------+-----.
          |/CE              A16-A14  A13-A0    |
          |         128K by 8 bit ROM     D7-D0+-- to 2A03 data bus
          |                                    |
          `------------------------------------'

(This diagram is for UNROM. UOROM has one more bit in the register, multiplexer output, and address input A17, and the multiplexer's other input is \$F. Homebrew boards capable of 512 KiB have one more of each, and the multiplexer's input is \$1F.)

The quad OR gate here acts as a multiplexer. A [74HC02](../7402/) quad NOR gate can be used instead if the banks are stored in reverse order in the ROM. If the program is 128 KiB or smaller, the 7402 way leaves one NOR gate free to invert R/W into /OE to avoid [bus conflicts](../Bus_conflict/).[\[1\]](http://forums.nesdev.org/viewtopic.php?p=111686#p111686)

If an actual multiplexer ([74HC157](../74157/) quad 2:1) is cheaper than an OR gate, a third-party UxROM-compatible board can use that instead of the 74HC32, as [kyuusaku suggested](http://forums.nesdev.org/viewtopic.php?p=93516#p93516).

## Variants {#Variants}

- The [mapper](../INES_Mapper_071/) used in Codemasters games published by Camerica extends UxROM with [CIC](../CIC_lockout_chip/) defeat circuitry, in addition being devoid of bus conflicts.
- Nintendo's [HVC-UN1ROM](../INES_Mapper_094/) board moves the bankswitching bits within the byte.
- *Crazy Climber* replaces the [74HC32](../7432/) quad-OR gate by a [74HC08](../7408/) quad-AND gate, so that the first bank is fixed at \$8000-\$BFFF and the switchable bank is present at \$C000-\$FFFF. This configuration is assigned to [iNES Mapper 180](../INES_Mapper_180/), which uses the same UNROM PCB.
- With an 8-bit latch ([74HC377](../74377/) or an additional 74HC161) and an additional 74HC32 to control A18-A21, a third-party board implementing this mapper can switch 4 MiB of PRG ROM.
  - Battle Kid 2: Mountain of Torment implements a 512kB UxROM mapper.
  - [UNROM 512](../UNROM_512/) implements a superset of 512kB UxROM, with additional CHR-RAM banking and the possibility of flash save.
- The homebrew game *Alwa's Awakening* adds 8 KiB of battery-backed PRG-RAM in the usual \$6000-\$7FFF address range. A few Subor educational cartridges have non-battery-backed PRG-RAM in the same range as well. Emulators should support this PRG-RAM at least in the presence of a NES 2.0 header that explicitly specifies it.

## See also {#See_also}

- [Programming UNROM](../Programming_UNROM/)

## External links {#External_links}

- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
- NES mapper list by Disch [\[2\]](https://romhack.ing/database/content/entry/ktNw5JQBNs8FWu0CfZdn/nes-mapper-list)
- [Converting UNROM to UOROM](https://forums.nesdev.org/viewtopic.php?p=179581#p179581)
