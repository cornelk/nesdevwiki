---
title: "INES Mapper 101"
url: "/wiki/INES_Mapper_101/"
description: "This mapper was allocated for a naïve misdumping of the JF-10 board that contains Urusei Yatsura - Lum no Wedding Bell, which assumed the bank bits were reversed. Otherwise, it is entirely identical…"
wikiPage: "INES_Mapper_101"
wikiCategories:
  - name: "Bad iNES Mappers"
    url: "/category/Bad_iNES_Mappers/"
  - name: "CNROM-like mappers"
    url: "/category/CNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

This mapper was allocated for a naïve misdumping of the [JF-10](https://nescartdb.com/profile/view/1547) board that contains *Urusei Yatsura - Lum no Wedding Bell*, which assumed the bank bits were reversed. Otherwise, it is entirely identical to [J87](../INES_Mapper_087/).

## Registers {#Registers}

### <span id="Bank_Select_.28.246000-.247FFF.29"></span>Bank Select (\$6000-\$7FFF) {#Bank_Select_.28.246000-.247FFF.29}

This will select an 8 KB CHR ROM bank for PPU \$0000-\$1FFF.

    7  bit  0
    ---- ----
    CCCC CCHL
    |||| ||||
    |||| |||+- (Normal) Low CHR bit
    |||| ||+-- (Normal) High CHR bit
    ++++-++--- (Oversize only) Extra CHR bits

This is the same as J87 but with the bits in order, and so has a better defined [oversize](../Oversize/) extension; or the same as [CNROM](../CNROM/) without bus conflicts.
