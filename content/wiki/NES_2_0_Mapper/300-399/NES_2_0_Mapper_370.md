---
title: "NES 2.0 Mapper 370"
url: "/wiki/NES_2.0_Mapper_370/"
description: "NES 2.0 Mapper 370 denotes the F600 circuit board, used for the Golden Mario Party II - Around the World 6-in-1 multicart. It uses an MMC3 clone with an outer bank register that also selects TKROM…"
wikiPage: "NES_2_0_Mapper_370"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 370** denotes the **F600** circuit board, used for the *Golden Mario Party II - Around the World 6-in-1* multicart. It uses an [MMC3](../MMC3/) clone with an outer bank register that also selects [TKROM](../TxROM/) and [TLSROM](../TLSROM/) configurations (the latter for a Mario-themed *Armadillo* hack).

## <span id="Outer_Bank_register_.28.245000-.245FFF.2C_write.29"></span>Outer Bank register (\$5000-\$5FFF, write) {#Outer_Bank_register_.28.245000-.245FFF.2C_write.29}

    Mask: $F000

    A~FEDC BA98 7654 3210
      -------------------
      0101 .... ..pP PcCC
                  || |+++- CHR A17-A19
                  || |+--- CHR Mask: 0=128K, 1=256 K
                  ++-+---- PRG A17-A19
                  +------- PRG Mask: 0=128K, 1=256 K

[TLSROM](../TLSROM/) nametable mirroring is used when cCC is 001, otherwise normal [MMC3](../MMC3/) mirroring is used.

## <span id="Solder_Pad_register_.28.245000-5FFF.2C_read.29"></span>Solder Pad register (\$5000-5FFF, read) {#Solder_Pad_register_.28.245000-5FFF.2C_read.29}

    Mask: $F000

    D~7654 3210
      ---------
      V... ....
      +--------- Solder pad setting

The solder pad setting selects one of two different menus.
