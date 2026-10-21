---
title: "NES 2.0 Mapper 348"
url: "/wiki/NES_2.0_Mapper_348/"
description: "NES 2.0 Mapper 348 is used for multicarts using 830118C-numbered PCBs such as 1994 New Series Red Pig 7-in-1. Its UNIF board name is BMC-830118C."
wikiPage: "NES_2_0_Mapper_348"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 348 is used for multicarts using *830118C*-numbered PCBs such as *1994 New Series Red Pig 7-in-1*. Its UNIF board name is **BMC-830118C**.

# Registers {#Registers}

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.2C_write.29"></span>MMC3-compatible registers (\$8000-\$FFFF, write) {#MMC3-compatible_registers_.28.248000-.24FFFF.2C_write.29}

These registers function identically to a normal [MMC3](../MMC3/), except in [GNROM mode](#GNROM-like_Mode).

## <span id="Outer_Bank_Register_.28.246800.2C_write.29"></span>Outer Bank Register (\$6800, write) {#Outer_Bank_Register_.28.246800.2C_write.29}

    Mask: probably $E000

    D~7654 3210
      ---------
      .... OO..
           ++--- Select 128 KiB outer PRG-ROM bank (i.e. PRG A17-A18)
                 and 128 KiB outer CHR-ROM bank (i.e. CHR A17-A18)

- This bank register functions as a substitute for WRAM and can only be written to if \$A001 bit 7 is set and \$A001 bit 6 is clear.
- The Outer Bank Register's PRG/CHR A17-A18 replace the MMC3's.

# GNROM-like Mode {#GNROM-like_Mode}

If PRG A17 and PRG A18 are both 1 (i.e. if \$6800 AND \$0C==\$0C), standard MMC3 PRG-ROM banking is replaced with an odd GNROM-ish mode, in which the MMC3's CPU A14 input is held low and PRG A14 is directly connected to CPU A14, resulting in the following bank layout:

- CPU \$8000-\$BFFF: MMC3 banks effective in the \$8000-\$BFFF range but as if their bank register values were AND'd with \$FD
- CPU \$C000-\$FFFF: MMC3 banks effective in the \$8000-\$BFFF range but as if their bank register values were OR'd with \$02

# Similar Mappers {#Similar_Mappers}

[NES 2.0 Mapper 315](../NES_2.0_Mapper_315/) is similar but with the Outer Bank Register bits being arranged slightly differently.
