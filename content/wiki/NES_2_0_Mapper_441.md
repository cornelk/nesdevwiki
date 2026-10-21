---
title: "NES 2.0 Mapper 441"
url: "/wiki/NES_2.0_Mapper_441/"
description: "NES 2.0 Mapper 441 denotes the 841026C and 850335C multicart circuit boards. It is similar to NES 2.0 Mapper 391 in functionality, but the register bits arranged differently."
wikiPage: "NES_2_0_Mapper_441"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 441** denotes the **841026C** and **850335C** multicart circuit boards. It is similar to [NES 2.0 Mapper 391](../NES_2.0_Mapper_391/) in functionality, but the register bits arranged differently.

Known cartridges:

- *1996 Power Rangers V 4-in-1* (MK-412)

## <span id="Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank and Mode Register (\$6000-\$7FFF, write) {#Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29}

    D~[LcCC pNPP]
       |||| ||++-- PRG A18..A17
       |||| |+---- 1=NROM-256 mode
       |||| +----- 0=PRG A17 from MMC1, 1=PRG A17 from PP
       ||++------- CHR A18..A17
       |+--------- 0=CHR A17 from MMC1, 1=CHR A17 from CC
       +---------- 1=Lock Outer Bank Register

- WRAM must be enabled in \$A001.7 before writing to this register.
- NROM-256 mode forces MMC3's CPU A14 input to GND to make MMC3 bank registers 6 and 7 apply to the entire CPU \$8000-\$FFFF range, with PRG A13 and PRG A14 being replaced by CPU A13 and A14, respectively.

## MMC3-compatible registers {#MMC3-compatible_registers}

See [MMC3](../MMC3/).
