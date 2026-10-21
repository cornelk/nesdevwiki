---
title: "NES 2.0 Mapper 361"
url: "/wiki/NES_2.0_Mapper_361/"
description: "NES 2.0 Mapper 361 is used for an early MMC3-based multicart PCB (YY841101C) from J.Y. Company, preceding the J.Y. Company ASIC:"
wikiPage: "NES_2_0_Mapper_361"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 361 is used for an early MMC3-based multicart PCB (YY841101C) from J.Y. Company, preceding the [J.Y. Company ASIC](../J.Y._Company_ASIC/):

- *1994 Ball 4-in-1 Series (JY-009)*
- *1996 Super HiK 4-in-1 (JY-018)*
- *1995 Super HiK 4-in-1 (JY-019)*
- *Game 4-in-1 (OK-411)*

## <span id="Outer_Bank_Register_.28.247000.29"></span>Outer Bank Register (\$7000) {#Outer_Bank_Register_.28.247000.29}

    Mask: unknown

    D~7654 3210
      ---------
      OOOO ....
      ++++------ Select outer 128 KiB PRG-ROM and CHR-ROM bank

This register is accessible regardless of whether WRAM is enabled in the MMC3 clone's A001 register.

## MMC3-compatible registers {#MMC3-compatible_registers}

    Mask: $E001

    See MMC3.

## Note {#Note}

[FCEUX emulates this PCB under mapper number 205](https://github.com/TASVideos/fceux/blob/master/src/boards/mmc3.cpp) as "OK-411" together with the GN-45 PCB, both of which are in turn different from the [iNES Mapper 205](../INES_Mapper_205/) described in that mapper's wiki article.

## See also {#See_also}

[NES 2.0 Mapper 364](../NES_2.0_Mapper_364/) describes a slightly more complex variant supporting a selectable outer bank size.
