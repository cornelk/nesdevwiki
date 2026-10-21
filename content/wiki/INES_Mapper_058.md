---
title: "INES Mapper 058"
url: "/wiki/INES_Mapper_058/"
description: "iNES Mapper 058 is used for a number of simple NROM-/CNROM-based multicarts:"
wikiPage: "INES_Mapper_058"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 058** is used for a number of simple NROM-/CNROM-based multicarts:

- *21-in-1 (AS-5321)*
- *50-in-1 (WQ1806 B)*
- *55-in-1 (WQ2006 B)*
- *68-in-1 (HKX5268)*
- *75-in-1 (WQ1905 E)*
- *86-in-1 (AP-5486)*
- *92-in-1 (WQ1605 E)*
- *97-in-1 (WQ1708 B)*

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1... .... MSCC CPPP
                |||| |+++- PRG A16..A14
                ||++-+---- CHR A15..A13
                |+-------- PRG Mode
                |           0: NROM-256 (PRG A14=CPU A14)
                |           1: NROM-128
                +--------- Nametable mirroring
                            0: Vertical
                            1: Horizontal

## Notes {#Notes}

- [INES Mapper 213](../INES_Mapper_213/) is a duplicate.
