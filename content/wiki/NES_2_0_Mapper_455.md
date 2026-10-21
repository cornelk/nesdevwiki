---
title: "NES 2.0 Mapper 455"
url: "/wiki/NES_2.0_Mapper_455/"
description: "NES 2.0 Mapper 455 denotes the NTDEC N625836 MMC3-based multicart circuit PCB, used for a 9000000-in-1 multicart."
wikiPage: "NES_2_0_Mapper_455"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 455** denotes the NTDEC **N625836** [MMC3](../MMC3/)-based multicart circuit PCB, used for a *9000000-in-1* multicart.

## <span id="Extra_Register_.28.244100.2C_write.29"></span>Extra Register (\$4100, write) {#Extra_Register_.28.244100.2C_write.29}

    A~[010. .... ..Q. .qDC] D~[...P PPBA]
                   |   |||        | |||+- 0: MMC3 mode, 1: NROM mode
                   |   |||        | ||+-- 0: NROM-128 mode, 1: NROM-256 mode (if A=1)
                   |   |||        +-++--- PRG A16..A14 in NROM mode
                   |   ||+--------------- 0: 128 KiB MMC3 inner PRG bank, 1: 256 KiB MMC3 inner PRG bank
                   |   |+---------------- 0: 128 KiB MMC3 inner CHR bank, 1: 256 KiB MMC3 inner CHR bank
                   +---+----------------- PRG/CHR A18..A17
