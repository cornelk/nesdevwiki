---
title: "NES 2.0 Mapper 411"
url: "/wiki/NES_2.0_Mapper_411/"
description: "NES 2.0 Mapper 411 denotes the A88S-1 multicart circuit board. It is yet another MMC3-compatible chipset with outer bank registers and an NROM mode."
wikiPage: "NES_2_0_Mapper_411"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 411** denotes the **A88S-1** multicart circuit board. It is yet another [MMC3](../MMC3/)-compatible chipset with outer bank registers and an NROM mode.

- *1997 Super 7-in-1 (JY-201)*
- *1997 Super 6-in-1 (JY-202)*
- *1997 Super 7-in-1 (JY-203)*
- *1997 龍珠武鬥會 7-in-1 (JY-204)*
- *1997 Super 7-in-1 (JY-205)*
- *1997 Super 7-in-1 (JY-206)*

# Registers {#Registers}

## <span id="Register_1_.28.245FF0.2C_write.29"></span>Register 1 (\$5FF0, write) {#Register_1_.28.245FF0.2C_write.29}

    Mask: unknown, probably $F001

    D~7654 3210
      ---------
      1M.c BCnA
       | | |||+- PRG A14 in NROM-128 mode
       | | ||+-- NROM mode type
       | | ||     0: PRG A14=A (NROM-128)
       | | ||     1: PRG A14=CPU A14 (NROM-256)
       | | |+--- PRG A16 in NROM mode
       | | +---- PRG A15 in NROM mode
       | +------ CHR A18
       +-------- PRG banking mode (PRG A13-A16 source)
                  0: MMC3
                  1: NROM (bits A-C, CPU A13)
    Power-on value: $80

Yes, the PRG A15/A16 bits are really in that order.

## <span id="Register_2_.28.245FF1.2C_write.29"></span>Register 2 (\$5FF1, write) {#Register_2_.28.245FF1.2C_write.29}

    Mask: unknown, probably $F001

    D~7654 3210
      ---------
      1p.. PCM.
       |   ||+-- PRG A17/CHR A17 source
       |   ||     0: PRG A17=P, CHR A17=C (128 KiB inner PRG and CHR bank)
       |   ||     1: from MMC3 (256 KiB inner PRG and CHR bank)
       |   |+--- CHR A17 if M=0
       |   +---- PRG A17 if M=0
       +-------- PRG A18
    Power-on value: $82
