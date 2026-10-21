---
title: "INES Mapper 113"
url: "/wiki/INES_Mapper_113/"
description: "iNES Mapper 113 denotes the HES NTD-8 circuit board, used to create multicarts made for a bunch of NINA-03/06 games."
wikiPage: "INES_Mapper_113"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>Mapper 113</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td><i>unknown</i></td></tr><tr><td><b>Boards</b></td><td><i>unknown</i></td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>128K </td></tr><tr><td><b>CHR window</b></td><td>8k </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, controlled by mapper </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>113 </td></tr></tbody></table></div>

iNES Mapper 113 denotes the HES **NTD-8** circuit board, used to create multicarts made for a bunch of [NINA-03/06](../NINA-003-006/) games.

Example Games:

- *HES 6-in-1*
- *Mind Blower Pak*
- *Total Funpak*

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank (max 256 KB)
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank (max 128 KB)

## Registers {#Registers}

### Control Register {#Control_Register}

This register is located at addresses:

    F        bit      0
    ---- ---- ---- ----
    010x xxx1 xxxx xxxx

That is, \$4100-\$41FF, \$4300-\$43FF, \$45xx, \$47xx, ..., \$5Dxx, and \$5Fxx.

    7  bit  0
    ---- ----
    MCPP PCCC
    |||| ||||
    |+||-|+++-- Select 8 KB CHR ROM bank for PPU $0000-$1FFF
    | ++-+----- Select 32 KB PRG ROM bank for CPU $8000-$FFFF
    +---------- Mirroring control (0 = Horizontal, 1 = Vertical)

Note the high bit of the CHR select is shifted away from the others.

## Variants {#Variants}

This mapper extends the [NINA-03/06](../NINA-003-006/) with an extra 4 bits of control register. The low 4 bits written behave the same, but the high 4 bits add [mirroring](../Mirroring/) control, 1 extra bit of CHR select, and 2 extra bits of PRG select.

## References {#References}

- [NES mapper list by Disch](http://www.romhacking.net/documents/362/)
