---
title: "INES Mapper 107"
url: "/wiki/INES_Mapper_107/"
description: "INES Mapper 107 is assigned to the unlicensed game Magic Dragon released by Magicseries."
wikiPage: "INES_Mapper_107"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with overlapping PRG and CHR registers"
    url: "/category/Mappers_with_overlapping_PRG_and_CHR_registers/"
---

***INES Mapper 107*** is assigned to the unlicensed game *[Magic Dragon](http://community.fandom.com/wiki/c:bootleggames:Magic_Dragon)* released by Magicseries.

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    BBBB BBBb
    |||| ||||
    ++++ ++++- Select 8 KB CHR ROM bank for PPU $0000-$1FFF
    ++++ +++-- Select 32 KB PRG ROM bank for CPU $8000-$FFFF

Unusually, bits 1-7 is used to switch both CHR and PRG banks simultaneously.

## Notes {#Notes}

- It's currently unknown whether this mapper suffers from bus conflicts or not.
- FCEUX implements only the bottom 3 bits of the unified register.
