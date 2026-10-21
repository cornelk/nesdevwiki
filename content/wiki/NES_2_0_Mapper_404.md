---
title: "NES 2.0 Mapper 404"
url: "/wiki/NES_2.0_Mapper_404/"
description: "NES 2.0 Mapper 404 denotes the JY012005 MMC1-based multicart circuit board, used only for the 1998 Super HiK 8-in-1 (JY-021B) multicart."
wikiPage: "NES_2_0_Mapper_404"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 404** denotes the **JY012005** [MMC1](../MMC1/)-based multicart circuit board, used only for the *1998 Super HiK 8-in-1 (JY-021B)* multicart.

## <span id="Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank and Mode Register (\$6000-\$7FFF, write) {#Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    D~7654 3210
      ---------
      LM.. ?BBb
      ||   |+++- PRG/CHR A19..A17
      ||   +---- Always set to 1 by the menu
      |+-------- 0: PRG A17=MMC1 A17 (256 KiB inner PRG bank)
      |          1: PRG A17=b        (128 KiB inner PRG bank)
      +--------- 1: Lock outer bank register

    Power-on value: $00

## <span id="MMC1-compatible_registers_.28.248000-.24FFFF.29"></span>MMC1-compatible registers (\$8000-\$FFFF) {#MMC1-compatible_registers_.28.248000-.24FFFF.29}

Mask: \$E000

See [MMC1](../MMC1/).
