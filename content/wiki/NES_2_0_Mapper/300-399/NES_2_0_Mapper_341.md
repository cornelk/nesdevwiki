---
title: "NES 2.0 Mapper 341"
url: "/wiki/NES_2.0_Mapper_341/"
description: "NES 2.0 mapper 341 is used for a simple 4-in-1 multicart. Its UNIF board name is BMC-TJ-03."
wikiPage: "NES_2_0_Mapper_341"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 341 is used for a simple *4-in-1* multicart. Its UNIF board name is **BMC-TJ-03**.

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1... .BBB .... ..M.
            |||        +-- Select nametable mirroring type
            |||             0: Vertical
            |||             1: Horizontal
            +++----------- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF and $C000-$FFFF
                           and 8 KiB CHR-ROM bank at PPU $0000-$1FFF
