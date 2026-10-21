---
title: "MMC4"
url: "/wiki/MMC4/"
description: "The Nintendo MMC4 is an ASIC mapper, used on the FxROM board set. The iNES format assigns mapper 10 to these boards. The chip first appeared in August 1988."
wikiPage: "MMC4"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers triggering on reads"
    url: "/category/Mappers_triggering_on_reads/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>MMC4</b><br/><b>FxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=10">3 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>FJROM, FKROM </td></tr><tr><td><b>Pinout</b></td><td><a href="../MMC4_pinout/">MMC4 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>128K </td></tr><tr><td><b>CHR window</b></td><td>4K + 4K (triggered) </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">010</a></td></tr></tbody></table></div>

The **Nintendo MMC4** is an [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/), used on the [FxROM](../FxROM/) board set. The [iNES](../INES/) format assigns **mapper 10** to these boards. The chip first appeared in August 1988.

Nintendo's MMC2, used in PxROM boards, is a similar mapper with 8 KB switchable PRG ROM banks, a 24 KB fixed PRG ROM bank, no PRG RAM, and a slightly different behaviour in auto-switching on the left (low) pattern table. This page only explains the differences, see [MMC2](../MMC2/) for full details on the rest of the mapper.

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KB fixed PRG RAM bank
- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$0FFF: Two 4 KB switchable CHR ROM banks
- PPU \$1000-\$1FFF: Two 4 KB switchable CHR ROM banks

When the PPU reads from specific tiles in the pattern table during rendering, the MMC4 sets a latch that tells it to use a different 4 KB bank number. On the background layer, this has the effect of setting a different bank for all tiles to the right of a given tile, virtually increasing the tile count limit from 256 to 512 without monopolising the CPU.

- PPU reads \$0FD8 through \$0FDF: latch 0 is set to \$FD
- PPU reads \$0FE8 through \$0FEF: latch 0 is set to \$FE
- PPU reads \$1FD8 through \$1FDF: latch 1 is set to \$FD
- PPU reads \$1FE8 through \$1FEF: latch 1 is set to \$FE

## Registers {#Registers}

The MMC4 has 6 registers at \$A000-\$AFFF, \$B000-\$BFFF, \$C000-\$CFFF, \$D000-\$DFFF, \$E000-\$EFFF and \$F000-\$FFFF. Only \$A000-\$AFFF is covered here. For the rest of the registers, see [MMC2](../MMC2/).

### <span id="PRG_ROM_bank_select_.28.24A000-.24AFFF.29"></span>PRG ROM bank select (\$A000-\$AFFF) {#PRG_ROM_bank_select_.28.24A000-.24AFFF.29}

    7  bit  0
    ---- ----
    xxxx PPPP
         ||||
         ++++- Select 16 KB PRG ROM bank for CPU $8000-$BFFF

## Hardware {#Hardware}

The MMC4 is implemented in a 44-pin TQFP package: [MMC4 pinout](../MMC4_pinout/)

Only one revision is known to exist.

## See also {#See_also}

- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Nintendo MMC4](http://nesdev.org/mmc4.txt) (author unknown)
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
