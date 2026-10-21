---
title: "NES 2.0 Mapper 327"
url: "/wiki/NES_2.0_Mapper_327/"
description: "NES 2.0 mapper 327 is used for a 6-in-1 multicart. Its UNIF board name is BMC-10-24-C-A1."
wikiPage: "NES_2_0_Mapper_327"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 327 is used for a *6-in-1* multicart. Its UNIF board name is **BMC-10-24-C-A1**.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    A~FEDC BA98 7654 3210
      -------------------
      0110 .... ..CT POOO
                  || |+++- Select 128 KiB PRG-ROM and CHR-ROM outer bank
                  || +---- Select MMC3 inner PRG-ROM bank address mask
                  ||        0: 128 KiB
                  ||        1: 256 KiB
                  |+------ Select CHR memory type
                  |         0: CHR-ROM
                  |         1: CHR-RAM
                  +------- Select MMC3 inner CHR-ROM bank address mask
                            0: 128 KiB
                            1: 256 KiB

As it uses the MMC3 clones's WRAM interface, writing to the Outer Bank register requires enabling and not write-protecting WRAM in the MMC's \$A001 register. Note that there are also 8 KiB of WRAM in the same address range. The outer bank register becomes locked against further writes once an outer bank other than 0 is selected, and is only unlocked on power-up.

## MMC3-compatible registers {#MMC3-compatible_registers}

    Mask: $E001

    See MMC3.
