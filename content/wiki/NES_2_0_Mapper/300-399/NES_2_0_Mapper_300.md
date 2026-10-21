---
title: "NES 2.0 Mapper 300"
url: "/wiki/NES_2.0_Mapper_300/"
description: "NES 2.0 Mapper 300 is used for Golden 190-in-1 multicart. Its UNIF board name is BMC-190in1. It is similar to INES Mapper 214 except that mirroring is selectable"
wikiPage: "NES_2_0_Mapper_300"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 300 is used for *Golden 190-in-1* multicart. Its UNIF board name is **BMC-190in1**. It is similar to [INES Mapper 214](../INES_Mapper_214/) except that mirroring is selectable

## <span id="Address_latch_.28.248000-.24FFFF.29"></span>Address latch (\$8000-\$FFFF) {#Address_latch_.28.248000-.24FFFF.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      .... .... ...G GG.M
                   | || +- Select nametable mirroring type
                   | ||     0: Vertical
                   | ||     1: Horizontal
                   +-++--- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF mirrored at $C000-$FFFF
                           and 8 KiB CHR-ROM bank at PPU $0000-$1FFF
