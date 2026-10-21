---
title: "NES 2.0 Mapper 320"
url: "/wiki/NES_2.0_Mapper_320/"
description: "NES 2.0 Mapper 320 is used for the Super HiK 6-in-1 A-030 multicart. Basically UxROM with an address-latch-based outer bank register, its UNIF board name is BMC-830425C-4391T. Mirroring is hard-wired."
wikiPage: "NES_2_0_Mapper_320"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 320 is used for the *Super HiK 6-in-1 A-030* multicart. Basically [UxROM](../UxROM/) with an address-latch-based outer bank register, its UNIF board name is **BMC-830425C-4391T**. Mirroring is hard-wired.

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable inner and outer PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB fixed inner and switchable outer PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

## <span id="Outer_Bank_Register_.28.24F0E0-.24F0FF.2C_write.29"></span>Outer Bank Register (\$F0E0-\$F0FF, write) {#Outer_Bank_Register_.28.24F0E0-.24F0FF.2C_write.29}

    Mask: $FFE0

    A~FEDC BA98 7654 3210
      -------------------
      1111 0000 111M BBBB
                   | ++++- Select 128 KiB Outer Bank at CPU $8000-$FFFF
                   +------ Select Outer Bank Size
                            0: 256 KiB (UOROM), fixed inner bank #15 at CPU $C000-$FFFF
                            1: 128 KiB (UNROM), fixed inner bank #7 at CPU $C000-$FFFF

## <span id="Inner_Bank_Register_.28.248000-.24FFFF.2C_write.29"></span>Inner Bank Register (\$8000-\$FFFF, write) {#Inner_Bank_Register_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    D~7654 3210
      ---------
      .... pPPP
           |+++- Select 16 KiB Inner PRG-ROM bank at CPU $8000-$BFFF
           +---- Select 128 KiB Inner PRG-ROM bank at CPU $8000-$BFFF if Outer Bank bit 4=0
