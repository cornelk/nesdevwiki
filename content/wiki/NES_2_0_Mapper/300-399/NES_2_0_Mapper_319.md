---
title: "NES 2.0 Mapper 319"
url: "/wiki/NES_2.0_Mapper_319/"
description: "NES 2.0 Mapper 319 denotes the HP-898F and KD-7/9-E circuit boards. Its UNIF board name is BMC-HP898F (with a different bank order). Used for:"
wikiPage: "NES_2_0_Mapper_319"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 319 denotes the **HP-898F** and **KD-7/9-E** circuit boards. Its UNIF board name is **BMC-HP898F** ([with a different bank order](#UNIF_bank_order)). Used for:

- *Prima Soft 9999999-in-1*
- *Olympic 2000 1000000-in-1*
- *4-in-1 (0207)*
- *4-in-1 (0210)*
- *4-in-1 (KG-443)*
- *9999999-in-1 (KD-6037-B)*

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB switchable PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB switchable CHR-ROM bank

# Registers {#Registers}

## <span id="PRG-ROM_Bank_and_Mirroring_.28.246004.29"></span>PRG-ROM Bank and Mirroring (\$6004) {#PRG-ROM_Bank_and_Mirroring_.28.246004.29}

    Mask: $6004

    D~7654 3210
      ---------   
      MDaB C...
      |||| +---- PRG A15
      |||+------ PRG A16
      ||+------- PRG A14 if D=0
      |+-------- PRG A14 mode
      |           0: PRG A14=p (NROM-128)
      |           1: PRG A14=CPU A14 (NROM-256)
      +--------- Select nametable mirroring type
                  0: Horizontal
                  1: Vertical

## <span id="CHR-ROM_Bank_.28.246000.29"></span>CHR-ROM Bank (\$6000) {#CHR-ROM_Bank_.28.246000.29}

    Mask: $6004

    D~7654 3210
      ---------
      CCCC ....
      ||||
      ||||
      ++++------ Select 8 KiB CHR-ROM bank at PPU $0000-$1FFF

## <span id="Solder_Pad_Register_.28.245FF0.29"></span>Solder Pad Register (\$5FF0) {#Solder_Pad_Register_.28.245FF0.29}

    Mask: probably $F000

    D~7654 3210
      ---------
      .P.. ....
       +-------- Solder pad value

# UNIF bank order {#UNIF_bank_order}

The publicly-available UNIF ROM file of *Prima Soft 9999999-in-1* has the order of the 16 KiB PRG-ROM banks slightly mixed up, so that the PRG A14 mode bit operates on A16 instead of A14. To obtain the correct bank order, use UNIF 16 KiB PRG banks 0, 4, 1, 5, 2, 6, 3, 7.
