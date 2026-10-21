---
title: "NES 2.0 Mapper 369"
url: "/wiki/NES_2.0_Mapper_369/"
description: "NES 2.0 mapper 369 denotes the N49C-300 circuit board, used for the Super Mario Bros. Party multicart. Accordingly, it combines all the mappers used for the various Mario games (including SMB2J): 0,…"
wikiPage: "NES_2_0_Mapper_369"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with fixed-timing cycle IRQs"
    url: "/category/Mappers_with_fixed-timing_cycle_IRQs/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 mapper 369** denotes the **N49C-300** circuit board, used for the *Super Mario Bros. Party* multicart. Accordingly, it combines all the mappers used for the various Mario games (including SMB2J): [0](../NROM/), [4](../MMC3/) and [40](../INES_Mapper_040/) with an outer bank register. Its UNIF board name is **BMC-N49C-300**.

## <span id="Outer_Bank_register_.28.244100.2C_write.29"></span>Outer Bank register (\$4100, write) {#Outer_Bank_register_.28.244100.2C_write.29}

Mask: \$C100

The meaning of the individual bits is not understood yet. The following values are used:

    $4100  PRG offset  CHR offset  Mapper
    Value  (8K bank#)  (8K bank#)
    -------------------------------------
    $00    $00         $00         000: NROM (Mario Bros.)
    $01    $01         $01         000: NROM (Super Mario Bros.)
    $13    $08         $03         040: NTDEC 2722 (Super Mario Bros. 2J)
    $37    $10         $10         004: TLROM (Super Mario Bros. 2), PRG-ROM bank limited to 128 KiB
    $FF    $20         $20         004: TKROM (Super Mario Bros. 3), PRG-ROM bank limited to 256 KiB
