---
title: "NES 2.0 Mapper 364"
url: "/wiki/NES_2.0_Mapper_364/"
description: "NES 2.0 Mapper 364 is used for an early MMC3-based multicart PCB (JY830832C) from J.Y. Company, preceding the J.Y. Company ASIC. A compatible PCB has the code 9053."
wikiPage: "NES_2_0_Mapper_364"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 364** is used for an early MMC3-based multicart PCB (**JY830832C**) from J.Y. Company, preceding the [J.Y. Company ASIC](../J.Y._Company_ASIC/). A compatible PCB has the code **9053**.

- *1994 Super HiK 3-in-1 (JY-007)*
- *1996 Super HiK Gold Card (LT-953)*

## <span id="Outer_Bank_Register_.28.246000-.247FFF.29"></span>Outer Bank Register (\$6000-\$7FFF) {#Outer_Bank_Register_.28.246000-.247FFF.29}

    Mask: $E000

    D~7654 3210
      ---------
      .PSC ....
       ||+------ CHR A18
       |+------- PRG/CHR A17 mode
       |          0: Preserve MMC3's PRG/CHR A17
       |          1: Force PRG/CHR A17 to 0
       +-------- PRG A18

This register is accessible regardless of whether WRAM is enabled in the MMC3 clone's A001 register. It may be possible that there is another bit, unused on the only discovered cartridge using this mapper, that selects the bit value which PRG/CHR A17 is forced to when S=1.

## MMC3-compatible registers {#MMC3-compatible_registers}

    Mask: $E001

    See MMC3.

## See also {#See_also}

[NES 2.0 Mapper 361](../NES_2.0_Mapper_361/) describes a simpler variant with a fixed outer bank size.
