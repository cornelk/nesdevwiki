---
title: "NES 2.0 Mapper 314"
url: "/wiki/NES_2.0_Mapper_314/"
description: "NES 2.0 Mapper 314 denotes the PB011003-3x1 PCB, used for the 32-in-1 Y2K, 64-in-1 Y2K and 76-in-1 Y2K multicarts. Its UNIF board name is BMC-64in1NoRepeat. It exists in both CHR-ROM and CHR-RAM…"
wikiPage: "NES_2_0_Mapper_314"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 314** denotes the **PB011003-3x1** PCB, used for the *32-in-1 Y2K*, *64-in-1 Y2K* and *76-in-1 Y2K* multicarts. Its UNIF board name is **BMC-64in1NoRepeat**. It exists in both CHR-ROM and CHR-RAM variants.

## <span id="Miscellaneous_register_.28.245000.2C_write.29"></span>Miscellaneous register (\$5000, write) {#Miscellaneous_register_.28.245000.2C_write.29}

    Mask: Unknown, probably $F003

    D~7654 3210
      ---------
      U.M. .CCP
      | |   ||+- PRG A21
      | |   ++- Bits 0-1 of 8 KiB CHR bank at PPU $0000-$1FFF
      | +------ Select nametable mirroring type
      |          0: Vertical
      |          1: Horizontal
      +-------- 0: UNROM
                1: NROM

    Power-on value: $80

## <span id="PRG-ROM_register_.28.245001.2C_write.29"></span>PRG-ROM register (\$5001, write) {#PRG-ROM_register_.28.245001.2C_write.29}

    Mask: Unknown, probably $F003

    D~7654 3210
      ---------
      MpPP PPPP
      ||++-++++- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF (PRG A20..A15)
      |+-------- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF and
      |          CPU $C000-$FFFF if M==0
      +--------- Select PRG banking mode
                  0: NROM-128
                  1: NROM-256

    Power-on value: $43

## <span id="CHR_mode_.28.245002.2C_write.29"></span>CHR mode (\$5002, write) {#CHR_mode_.28.245002.2C_write.29}

    Mask: Unknown, probably $F003

    D~7654 3210
      ---------
      .... CCCC
           ++++- Bits 2-5 of 8 KiB CHR bank at PPU $0000-$1FFF

    Power-on value: $00

## <span id="UNROM_Latch_.28.248000-.24FFFF.2C_write.29"></span>UNROM Latch (\$8000-\$FFFF, write) {#UNROM_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    D~7654 3210
      ---------
      .... .CCC
            +++- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF

    Power-on value: $00
