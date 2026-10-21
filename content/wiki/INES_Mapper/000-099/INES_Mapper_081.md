---
title: "INES Mapper 081"
url: "/wiki/INES_Mapper_081/"
description: "This mapper is used on exactly one known game: Super Gun from NTDEC. Although the game's bankswitching code seems to suggest two separate data latches at $6000 and $8000-$FFFF (the latter with bus…"
wikiPage: "INES_Mapper_081"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

<div class="infobox"><div style="text-align: center"><b>N715021</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>NTDEC </td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>N715021 </td></tr><tr><td><b>PRG ROM capacity</b></td><td>64K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>32K </td></tr><tr><td><b>CHR window</b></td><td>8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>081 </td></tr></tbody></table></div>

This mapper is used on exactly one known game: Super Gun from NTDEC. Although the game's bankswitching code seems to suggest two separate data latches at \$6000 and \$8000-\$FFFF (the latter with bus conflicts), the board actually latches the four lowest address bits, and ignores the write to \$6000.

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank

## Registers {#Registers}

### <span id="PRG.2FCHR_Bank_Select_.28.248000-.24FFFF.29"></span>PRG/CHR Bank Select (\$8000-\$FFFF) {#PRG.2FCHR_Bank_Select_.28.248000-.24FFFF.29}

     CPU address bit#
    1111 1100 0000 0000
    5432 1098 7654 3210
    ---- ---- ---- ----
    1... .... .... PPCC
                   ||||
                   ||++------ Select 8 KB CHR ROM bank for PPU $0000-$1FFF
                   ++-------- Select 16 KB PRG ROM bank for CPU $8000-$BFFF

## Related {#Related}

[Discussion with PCB images and analysis as well as Kazzo dumping script](http://forums.nesdev.org/viewtopic.php?f=3&t=16412)
