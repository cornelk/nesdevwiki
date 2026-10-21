---
title: "NES 2.0 Mapper 338"
url: "/wiki/NES_2.0_Mapper_338/"
description: "NES 2.0 mapper 338 is used for a 16-in-1 and a 200/300/600/1000-in-1 multicart. Its UNIF board name is BMC-SA005-A."
wikiPage: "NES_2_0_Mapper_338"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 338 is used for a *16-in-1* and a *200/300/600/1000-in-1* multicart. Its UNIF board name is **BMC-SA005-A**.

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1... .... .... MBBB
                     ++++- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF and $C000-$FFFF
                     |     and 8 KiB CHR-ROM bank at PPU $0000-$1FFF
                     +---- Select nametable mirroring type
                            0: Horizontal
                            1: Vertical

## Similar mappers {#Similar_mappers}

[INES Mapper 200](../INES_Mapper_200/) is similar but with the mirroring bit's meaning flipped, and the mirroring bit being also used as a fourth bank bit.
