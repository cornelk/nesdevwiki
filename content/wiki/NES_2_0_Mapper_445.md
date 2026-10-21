---
title: "NES 2.0 Mapper 445"
url: "/wiki/NES_2.0_Mapper_445/"
description: "NES 2.0 Mapper 445 denotes the DG574B MMC3-compatible multicart circuit board."
wikiPage: "NES_2_0_Mapper_445"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 445** denotes the **DG574B** [MMC3](../MMC3/)-compatible multicart circuit board.

## <span id="Outer_PRG_Bank_.28.245000.2C_write.29"></span>Outer PRG Bank (\$5000, write) {#Outer_PRG_Bank_.28.245000.2C_write.29}

    D~[.PPP ....]  Address mask: $F003
        +++------- PRG A19..A17

## <span id="Outer_CHR_Bank_.28.245001.2C_write.29"></span>Outer CHR Bank (\$5001, write) {#Outer_CHR_Bank_.28.245001.2C_write.29}

    D~[.CCC ....]  Address mask: $F003
        +++------- CHR A19..A17

## <span id="Mode_.28.245002.2C_write.29"></span>Mode (\$5002, write) {#Mode_.28.245002.2C_write.29}

    D~[...? c.?p]  Address Mask: $F003
            |  +- PRG A17 Mode
            |      0: from MMC3
            |      1: from Outer PRG Bank
            +---- CHR A17 Mode
                   0: from MMC3

## <span id="Lock_.28.245003.2C_write.29"></span>Lock (\$5003, write) {#Lock_.28.245003.2C_write.29}

    D~[..L. ....]  Address Mask: $F003
         +------- 1: Lock registers $5000-$5003

## MMC3-compatible registers {#MMC3-compatible_registers}

See [MMC3](../MMC3/).
