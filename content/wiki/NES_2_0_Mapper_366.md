---
title: "NES 2.0 Mapper 366"
url: "/wiki/NES_2.0_Mapper_366/"
description: "NES 2.0 Mapper 366 is used for an MMC3-based multicart PCB (GN-45):"
wikiPage: "NES_2_0_Mapper_366"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 366 is used for an MMC3-based multicart PCB (GN-45):

- *4-in-1 (K-3131GS)*
- *4-in-1 (K-3131SS)*

## <span id="Outer_Bank_Register_.28.246800.29"></span>Outer Bank Register (\$6800) {#Outer_Bank_Register_.28.246800.29}

    Mask: unknown, probably $E0F0

    A~FEDC BA98 7654 3210
      -------------------
      0110 1000 LOOO ....
                |+++------ Select outer 128 KiB PRG-ROM and CHR-ROM bank
                +--------- 1=Lock outer bank

This register is accessible regardless of whether WRAM is enabled in the MMC3 clone's A001 register. It overlays the 8 KiB of WRAM that both multicarts have, making the Lock bit necessary to prevent the outer bank from changing once the game writes to WRAM.

## MMC3-compatible registers {#MMC3-compatible_registers}

    Mask: $E001

    See MMC3.

## Note {#Note}

[FCEUX emulates this PCB under mapper number 205](https://github.com/TASVideos/fceux/blob/master/src/boards/mmc3.cpp) together with "OK-411" ([NES 2.0 Mapper 361](../NES_2.0_Mapper_361/)), both of which are in turn different from the [iNES Mapper 205](../INES_Mapper_205/) described in that mapper's wiki article.

## See also {#See_also}

[NES 2.0 Mapper 361](../NES_2.0_Mapper_361/) describes a simpler variant that latches the data bus instead of the address bus.
