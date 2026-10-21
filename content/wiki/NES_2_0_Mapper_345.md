---
title: "NES 2.0 Mapper 345"
url: "/wiki/NES_2.0_Mapper_345/"
description: "NES 2.0 Mapper 345 is used for the New Star 6-in-1 Game Cartridge multicart. Unusually, it combines MMC3 and AxROM games. Its UNIF board name is BMC-L6IN1. CHR data comes from 8 KiB of MMC3-banked…"
wikiPage: "NES_2_0_Mapper_345"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 345 is used for the *New Star 6-in-1 Game Cartridge* multicart. Unusually, it combines [MMC3](../MMC3/) and [AxROM](../AxROM/) games. Its UNIF board name is **BMC-L6IN1**. CHR data comes from 8 KiB of MMC3-banked CHR-RAM.

## <span id="Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank and Mode Register (\$6000-\$7FFF, write) {#Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    D~7654 3210
      ---------
      OOMS mmPP
      |||| ||++- Select 32 KiB inner PRG-ROM bank in AxROM mode
      |||| ++--- Select PRG-ROM banking mode (0=AxROM, 1-3=MMC3)
      |||+------ Select 1 KiB CIRAM bank in one-screen mirroring mode
      ||+------- Select nametable mirroring type
      ||          0: from MMC3
      ||          1: one-screen mirroring
      ++-------- Select 128 KiB outer PRG-ROM bank in either mode

As it uses the MMC3 clones's WRAM interface, writing to the Outer Bank register requires enabling and not write-protecting WRAM in the MMC's \$A001 register.

## MMC3-compatible registers {#MMC3-compatible_registers}

    Mask: $E001

    See MMC3.
