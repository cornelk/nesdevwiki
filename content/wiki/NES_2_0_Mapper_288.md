---
title: "NES 2.0 Mapper 288"
url: "/wiki/NES_2.0_Mapper_288/"
description: "NES 2.0 Mapper 288 is used for two GKCX1 21-in-1 multicarts. Oddly, GoodNES 3.23b has them set to INES Mapper 133 even though it is entirely incompatible with that mapper."
wikiPage: "NES_2_0_Mapper_288"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 288 is used for two GKCX1 21-in-1 multicarts. Oddly, GoodNES 3.23b has them set to [INES Mapper 133](../INES_Mapper_133/) even though it is entirely incompatible with that mapper.

# Address latch {#Address_latch}

    A~1... .... ..MP PCCC
                  || |+++- Select 8 KiB CHR-ROM bank at PPU $0000-$1FFF
                  |+-+---- Select 32 KiB CHR-ROM bank at CPU $8000-$FFFF
                  -------- 0=V mirroring, 1=H mirroring
