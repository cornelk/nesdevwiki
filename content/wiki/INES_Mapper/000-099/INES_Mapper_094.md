---
title: "INES Mapper 094"
url: "/wiki/INES_Mapper_094/"
description: "INES Mapper 094 is used to represent the board HVC-UN1ROM. Senjou no Ookami (Japanese version of Commando) is the only game known to utilize this mapper. It is very similar to UxROM, but the register…"
wikiPage: "INES_Mapper_094"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=94">iNES 094</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+%60%25%40%25%60&amp;unif=UN1ROM">UN1ROM</a></td></tr></tbody></table></div>

INES Mapper 094 is used to represent the board HVC-UN1ROM. *Senjou no Ookami* (Japanese version of *Commando*) is the only game known to utilize this mapper. It is very similar to [UxROM](../UxROM/), but the register is shifted by two bits.

## Overview {#Overview}

- PRG ROM size: 128 KB
- PRG ROM bank size: 16 KB
- PRG RAM: None
- CHR capacity: 8 KB RAM
- CHR bank size: Not bankswitched
- Nametable [mirroring](../Mirroring/): Solder pads select vertical or horizontal mirroring
- Subject to [bus conflicts](../Bus_conflict/): Yes

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxxP PPxx
       | ||
       +-++--- Select 16 KB PRG ROM bank for CPU $8000-$BFFF

## References {#References}

- NES mapper list by Disch [\[1\]](http://www.romhacking.net/documents/362/)
