---
title: "NES 2.0 Mapper 385"
url: "/wiki/NES_2.0_Mapper_385/"
description: "NES 2.0 Mapper 385 denotes the NTDEC 2779 circuit board, used for a simple 5-in-1 multicart with 8 KiB of unbanked CHR-RAM containing NROM-128 games. It is similar to INES Mapper 202 except that…"
wikiPage: "NES_2_0_Mapper_385"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 385** denotes the NTDEC **2779** circuit board, used for a simple *5-in-1* multicart with 8 KiB of unbanked CHR-RAM containing NROM-128 games. It is similar to [INES Mapper 202](../INES_Mapper_202/) except that there is no NROM-128/256 switch, and that it uses CHR-RAM rather than CHR-ROM.

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      .... .... .... BBBM
                     |||+- 0: Vertical, 1: Horizontal mirroring
                     +++-- PRG A14-A16 at CPU $8000-$BFFF and CPU $C000-$FFFF

## See also {#See_also}

[NESdev forum discussion](https://forums.nesdev.org/viewtopic.php?f=9&t=19198)
