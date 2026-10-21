---
title: "NES 2.0 Mapper 376"
url: "/wiki/NES_2.0_Mapper_376/"
description: "NES 2.0 Mapper 376 denotes J.Y. Company's YY841155C circuit board, used for one one known multicart:"
wikiPage: "NES_2_0_Mapper_376"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 376** denotes J.Y. Company's **YY841155C** circuit board, used for one one known multicart:

- *1996 大金鋼專輯 8-in-1* (JY-041)

Similarly to [NES 2.0 Mapper 361](../NES_2.0_Mapper_361/), it uses an [MMC3](../MMC3/) clone with two outer bank registers at \$7000/\$7001. This variant adds an NROM PRG mode.

## <span id="Outer_Bank_register_.230_.28.247000.2C_write.29"></span>Outer Bank register \#0 (\$7000, write) {#Outer_Bank_register_.230_.28.247000.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      UaV. .bbb
      |||   +++- In NROM mode: PRG A14..16
      |+|------- PRG/CHR A17
      +-+------- UV  PRG banking mode
                 --------------------
                 10  NROM-128
                 11  NROM-256
                 0x  MMC3

Regardless of the PRG banking mode, CHR-ROM A10..A16 will come from the MMC3.

## <span id="Outer_Bank_register_.231_.28.247001.2C_write.29"></span>Outer Bank register \#1 (\$7001, write) {#Outer_Bank_register_.231_.28.247001.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      .... ...B
              +- PRG/CHR A18
