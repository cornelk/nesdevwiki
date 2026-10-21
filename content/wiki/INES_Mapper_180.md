---
title: "INES Mapper 180"
url: "/wiki/INES_Mapper_180/"
description: "INES Mapper 180 is used to represent an unusual configuration of the UNROM board used only for the game Crazy Climber."
wikiPage: "INES_Mapper_180"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

<div class="infobox"><div style="text-align: center"><b>UxROM (inverted)</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nichibutsu </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=180">1 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>UNROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K/4096K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K fixed + 16K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>180, <a href="../UxROM/">002</a>, <a href="../INES_Mapper_094/">094</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=180">iNES 180</a></td></tr></tbody></table></div>

INES Mapper 180 is used to represent an unusual configuration of the [UNROM](../UxROM/) board used only for the game *Crazy Climber*.

## Notes {#Notes}

It is similar in principle to the more common [INES Mapper 002](../UxROM/) configuration but uses AND logic ([74HC08](../7408/)) instead of OR logic ([74HC32](../7432/)), producing a fixed first bank at \$8000 rather than a fixed last bank at \$C000.

Crazy Climber had a special "climber stick" add-on that mounted a stick on top of the Famicom's directional pad.

For a similar mapper with additional mirroring control, see [iNES Mapper 097](../INES_Mapper_097/).

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KB PRG ROM bank, fixed to the first bank
- CPU \$C000-\$FFFF: 16 KB switchable PRG ROM bank

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxxx xPPP
          |||
          +++-- Select 16 KB PRG ROM bank for CPU $C000-$FFFF

## See also {#See_also}

- [UxROM](../UxROM/)
- [Programming UNROM](../Programming_UNROM/)
- NES mapper list by Disch [\[1\]](http://www.romhacking.net/documents/362/)
