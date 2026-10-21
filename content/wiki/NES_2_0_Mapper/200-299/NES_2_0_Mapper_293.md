---
title: "NES 2.0 Mapper 293"
url: "/wiki/NES_2.0_Mapper_293/"
description: "NES 2.0 Mapper 293 is used for the NewStar 12-in-1 and 76-in-1 multicarts. Both use 8 KiB of unbanked CHR-RAM rather than CHR-ROM."
wikiPage: "NES_2_0_Mapper_293"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 293 is used for the *NewStar 12-in-1* and *76-in-1* multicarts. Both use 8 KiB of unbanked CHR-RAM rather than CHR-ROM.

# Registers {#Registers}

## <span id="First_Banking_Register_.28.248000-.249FFF.2C_.24C000-.24DFFF.29"></span>First Banking Register (\$8000-\$9FFF, \$C000-\$DFFF) {#First_Banking_Register_.28.248000-.249FFF.2C_.24C000-.24DFFF.29}

    Mask: $A000

    7654 3210
    ---------
    .... PIII
         |+++- Select inner 16 KiB PRG-ROM bank
         +---- Bit 1 of the PRG banking mode number

## <span id="Second_Banking_Register_.28.248000-.24BFFF.29"></span>Second Banking Register (\$8000-\$BFFF) {#Second_Banking_Register_.28.248000-.24BFFF.29}

    Mask: $C000

    7654 3210
    ---------
    MP10 ...2
    ||++----+- Select outer 128 KiB PRG-ROM bank (note the bit distribution)
    |+-------- Bit 0 of the PRG banking mode number
    +--------- Select nametable mirroring type
                0: Vertical
                1: Horizontal

# PRG banking modes {#PRG_banking_modes}

- 0 (UNROM)

\$8000-\$BFFF: Switchable 16 KiB Inner PRG-ROM bank

\$C000-\$FFFF: Fixed 16 KiB Inner PRG-ROM bank \#7

- 1

\$8000-\$BFFF: Switchable 16 KiB Inner PRG-ROM bank AND \$FE

\$C000-\$FFFF: Fixed 16 KiB Inner PRG-ROM bank \#7

- 2 (NROM-128)

\$8000-\$BFFF: Switchable 16 KiB Inner PRG-ROM bank

\$C000-\$FFFF: Mirror of \$8000-\$BFFF

- 3 (NROM-256)

\$8000-\$FFFF: Switchable 32 KiB Inner PRG-ROM bank (bit 0 of inner bank number ignored)

# Notes {#Notes}

- The address ranges of the two banking registers do indeed partially overlap.
- Nestopia Plus! assigns this board to Mapper 281 instead.
