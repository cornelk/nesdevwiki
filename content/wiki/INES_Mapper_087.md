---
title: "INES Mapper 087"
url: "/wiki/INES_Mapper_087/"
description: "J87 (Also known as Mapper 87) refers to the first \"mapper\" created for the Famicom. Despite this, no games with this mapper ever made it overseas to USA NES owners. iNES Mapper 101 is the same board,…"
wikiPage: "INES_Mapper_087"
wikiCategories:
  - name: "CNROM-like mappers"
    url: "/category/CNROM-like_mappers/"
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

<div class="infobox"><div style="text-align: center"><b>J87</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami, Jaleco, Taito </td></tr><tr><td><b>Boards</b></td><td>JF-0x, JF-10 (Jaleco), FC-009 (Taito), MF4034/30xxA (Konami) </td></tr><tr><td><b>PRG ROM capacity</b></td><td>32K </td></tr><tr><td><b>PRG ROM window</b></td><td>n/a </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>16k (most games), 32K (JF-09/JF-10 only) </td></tr><tr><td><b>CHR window</b></td><td>8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>087, <a href="../INES_Mapper_101/">101</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=87">iNES 087</a></td></tr></tbody></table></div>

**J87** (Also known as Mapper 87) refers to the first "mapper" created for the Famicom. Despite this, no games with this mapper ever made it overseas to USA NES owners. [iNES Mapper 101](../INES_Mapper_101/) is the same board, but for a dump that naïvely reversed the CHR bit-order. [CNROM](../CNROM/) is a later cost-reduced version of this mapper, by cutting down on one logic chip, at the cost of having bus conflicts, and also has the CHR bank order reversed as well. Most games are capped at 16k CHR, however, the JF-09 and JF-10 boards contain an additional trace that allows the design to address 32k. All games with this mapper are emulated as the 32k variant.

The name "J87" was given to this mapper due to its historical significance, and creation by Jaleco. Otherwise, no "canonical" name was given to this mapper. Jaleco would later upgrade this mapper to become [iNES Mapper 140](../INES_Mapper_140/).

## Banks {#Banks}

- CPU \$8000-\$FFFF: 16 KB PRG ROM, fixed (if 16 KB PRG ROM used, then this is the same as \$C000-\$FFFF)
- CPU \$C000-\$FFFF: 16 KB PRG ROM, fixed
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank

## Registers {#Registers}

### <span id="Bank_Select_.28.246000-.247FFF.29"></span>Bank Select (\$6000-\$7FFF) {#Bank_Select_.28.246000-.247FFF.29}

This will select an 8 KB CHR ROM bank for PPU \$0000-\$1FFF.

    7  bit  0
    ---- ----
    xxxx xxLH
           ||
           |+- High CHR bit
           +-- Low CHR bit
