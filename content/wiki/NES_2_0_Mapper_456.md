---
title: "NES 2.0 Mapper 456"
url: "/wiki/NES_2.0_Mapper_456/"
description: "NES 2.0 Mapper 456 denotes the MMC3-based Realtec K6C3001A multicart PCB, used for the (ST-401) Super 4-in-1 multicart. It is basically an oversized version of INES Mapper 047 with the outer bank…"
wikiPage: "NES_2_0_Mapper_456"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 456** denotes the [MMC3](../MMC3/)-based Realtec **K6C3001A** multicart PCB, used for the *(ST-401) Super 4-in-1* multicart. It is basically an oversized version of [INES Mapper 047](../INES_Mapper_047/) with the outer bank register moved from \$6000 to \$4100.

## <span id="Extra_Register_.28.244100.2C_write.29"></span>Extra Register (\$4100, write) {#Extra_Register_.28.244100.2C_write.29}

    Mask: $E100

    D~[.... ..BB]
              ++- PRG/CHR A18..A17
    Value on reset: $00
