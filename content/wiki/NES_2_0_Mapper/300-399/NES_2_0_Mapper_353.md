---
title: "NES 2.0 Mapper 353"
url: "/wiki/NES_2.0_Mapper_353/"
description: "NES 2.0 Mapper 353 is used for the 92 Super Mario Family multicart, consisting of an MMC3 clone ASIC together with a PAL. The PCB code is 81-03-05-C."
wikiPage: "NES_2_0_Mapper_353"
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

NES 2.0 Mapper 353 is used for the *92 Super Mario Family* multicart, consisting of an [MMC3](../MMC3/) clone ASIC together with a PAL. The PCB code is **81-03-05-C**.

# <span id="Outer_Bank_Register_.28.248080-.24E080.29"></span>Outer Bank Register (\$8080-\$E080) {#Outer_Bank_Register_.28.248080-.24E080.29}

    Mask: Unknown, but must be more than just $8080.

    A~FEDC BA98 7654 3210
      -------------------
      1OO. .... 1... ....
       ++----------------- Select 256 KiB outer PRG-ROM bank (PRG A18-19)
                           and 128 KiB outer CHR-ROM bank (CHR A17-18)

    Power-up value: $00

# Notes {#Notes}

- In addition to the setting the outer PRG-ROM and CHR-ROM bank, the outer bank register also modifies other aspects:
  - Outer bank 0: Uses CHR A17 instead of the MMC3's Mirroring register, similar to [INES Mapper 118](../INES_Mapper_118/).
  - Outer bank 2: The outer PRG-ROM bank size is just 128 KiB, with CHR A17 serving as PRG A17 instead. Additionally, if CHR A17 is high, 8 KiB of unbanked CHR-RAM are used instead of CHR-ROM.
  - Outer bank 3: If CHR A17 is low, PRG A17 is replaced with CPU A14, so that CPU \$8000-\$BFFF are taken from the first and CPU \$C000-\$FFF are taken from the second 128 KiB half.
- The conversion of *Bio Miracle Bokutte Upa* on that multicart retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original multicart did not.
- The conversion of *Super Mario Bros. 2 (J)* is incomplete and will fail to work after World 4-4.
- Resetting the console always returns to the main menu (outer bank register reset to \$0) due to a circuit detecting interruptions in the M2 signal.

# See also {#See_also}

[PCB image and analysis](http://forums.nesdev.org/viewtopic.php?f=9&t=17524)
