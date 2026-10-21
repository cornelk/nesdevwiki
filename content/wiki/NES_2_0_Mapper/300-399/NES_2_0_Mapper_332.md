---
title: "NES 2.0 Mapper 332"
url: "/wiki/NES_2.0_Mapper_332/"
description: "NES 2.0 mapper 332 is used for the Super 40-in-1 and Super 500-in-1 multicarts. Its UNIF board name is BMC-WS."
wikiPage: "NES_2_0_Mapper_332"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 mapper 332** is used for the *Super 40-in-1* and *Super 500-in-1* multicarts. Its UNIF board name is **BMC-WS**.

## <span id="PRG.2FMirroring_Register_.28.246000.29"></span>PRG/Mirroring Register (\$6000) {#PRG.2FMirroring_Register_.28.246000.29}

    Mask: probably $E001

    D~7654 3210
      ---------
      .H.M BPPp
       | | |+++- PRG A14..A16 for CPU $8000-$BFFF and $C000-$FFFF
       | | +---- Select PRG A14 Mode
       | |        0: PRG A14=CPU A14 (NROM-256)
       | |        1: PRG A14=p (NROM-128)
       | +------ Select nametable mirroring mode
       |          0: Vertical
       |          1: Horizontal
       +-------- Select PRG A17 and CHR A16
    Reset value: $00

## <span id="CHR.2FSolder_Pad_Register_.28.246001.29"></span>CHR/Solder Pad Register (\$6001) {#CHR.2FSolder_Pad_Register_.28.246001.29}

    Mask: probably $E001

    D~7654 3210
      ---------
      DDMM .Cdc
      ||||  |++- CHR A13..A14 for PPU $0000-$1FFF if MM=11
      ||||  +--- CHR A15
      ||++------ 00: CNROM-256 mode: CHR A13..A14 from CNROM latch
      ||         10: CNROM-128 mode: CHR A13 from CNROM latch, CHR A14 from d
      ||         11: NROM mode: CHR A13..A14 from dc
      ++-------- Solder pad check
    Reset value: $00

The solder pads connect the latch output directly to PRG-ROM /CE. The menu code, executing from RAM, sets either bit to check whether CPU \$8000-\$FFFF now returns open bus, and selects a menu with different game counts depending on the result.

## <span id="CNROM_latch_.28.248000-.24FFFF.29"></span>CNROM latch (\$8000-\$FFFF) {#CNROM_latch_.28.248000-.24FFFF.29}

    D~7654 3210
      ---------
      .... ..CC
             ++- CHR A13..A14 for PPU $0000-$1FFF if MM=00
    Reset value: $00
