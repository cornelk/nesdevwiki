---
title: "INES Mapper 038"
url: "/wiki/INES_Mapper_038/"
description: "iNES Mapper 038 represents the board used for Bit Corp.'s Crime Busters. It is another mapper almost identical to GNROM (66) except the bits and writeable port moved around."
wikiPage: "INES_Mapper_038"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

<div class="infobox"><div style="text-align: center"><b>iNES Mapper 038</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Bit Corp. </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=38">1 in NesCartDB</a></td></tr><tr><td><b>Boards</b></td><td>UNL-PCI556 </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>32K </td></tr><tr><td><b>CHR window</b></td><td>8k </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>V(hardwired) </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>038 </td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=38">iNES 038</a></td></tr></tbody></table></div>

iNES Mapper 038 represents the board used for [Bit Corp.](http://community.fandom.com/wiki/c:bootleggames:Bit_Corp.)'s [Crime Busters](https://nescartdb.com/profile/view/4103/crime-busters). It is another mapper almost identical to [GNROM](../GxROM/) (**66**) except the bits and writeable port moved around.

## Registers {#Registers}

### <span id="Bank_Select_.28.247000-.247FFF.29"></span>Bank Select (\$7000-\$7FFF) {#Bank_Select_.28.247000-.247FFF.29}

    7  bit  0
    ---- ----
    xxxx CCPP
         ||||
         ||++- Select 32 KB PRG ROM bank for CPU $8000-$FFFF
         ++--- Select 8 KB CHR ROM bank for PPU $0000-$1FFF

## Hardware {#Hardware}

Writes to \$F000-\$FFFF are very likely to also trigger a bankswitch, but well behaved code has no reason to. The [74138](../74138/) used by this board uses /ROMSEL and R/W on its enable inputs, but /ROMSEL is 1 while M2 is low, not just while M2 is high and and A15 is low. As a result, before M2 is high, if the 2A03's A14, A13, and A12 lines are 1 and R/W is 0, the [74161](../74161/)'s /LOAD line will still load from the data bus.
