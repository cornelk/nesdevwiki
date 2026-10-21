---
title: "NES 2.0 Mapper 443"
url: "/wiki/NES_2.0_Mapper_443/"
description: "NES 2.0 Mapper 443 denotes the NC3000M multicart circuit board."
wikiPage: "NES_2_0_Mapper_443"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 443** denotes the **NC3000M** multicart circuit board.

Known cartridges:

- *3000-in-1*

## <span id="Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank and Mode Register (\$6000-\$7FFF, write) {#Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    A~FEDC BA98 7654 3210
      -------------------
      011. .... ...P MMpC
                   | |||+- CHR A18
                   | ||+-- PRG A18
                   | ++--- Select PRG-ROM mode
                   |        0: MMC3 mode
                   |        1: NROM-256 mode (PRG A13..14=CPU A13..14, PRG A15..16=MMC3 A15..16)
                   |        2: Read solder pads' value in D0/D1 (leftmost pads in picture)
                   |        3: NROM-128 mode (PRG A13=CPU A13, PRG A14..16=MMC3 A14..16)
                   +------ PRG A17

- WRAM must be enabled in \$A001.7 before writing to this register.
- The inner PRG bank is restricted to 128 KiB.
- The inner CHR bank is unrestricted to 256 KiB.
- NROM mode forces MMC3's CPU A13 and A14 inputs to GND to make MMC3 bank register 6 apply to the entire CPU \$8000-\$FFFF range.

## MMC3-compatible registers {#MMC3-compatible_registers}

Mask: \$E001

See [MMC3](../MMC3/).
