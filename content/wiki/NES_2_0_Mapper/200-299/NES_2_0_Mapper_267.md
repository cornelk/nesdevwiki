---
title: "NES 2.0 Mapper 267"
url: "/wiki/NES_2.0_Mapper_267/"
description: "NES 2.0 Mapper 267 is used for the 8-in-1 JY-119 multicart. It consists of an MMC3 clone with an outer bank register at $6000."
wikiPage: "NES_2_0_Mapper_267"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 267 is used for the 8-in-1 JY-119 multicart. It consists of an [MMC3](../MMC3/) clone with an outer bank register at \$6000.

## MMC3-compatible registers {#MMC3-compatible_registers}

Mask: \$E001

- \$8000, \$8001, \$A000, \$A001, \$C000, \$C001, \$E000, \$E001: As normal [MMC3](../MMC3/). Inner PRG banks are limited to 256 KiB, Inner CHR banks to 128 KiB.

## <span id="Outer_Bank_.28.246000.29"></span>Outer Bank (\$6000) {#Outer_Bank_.28.246000.29}

Mask: Unknown

    7654 3210
    ---- ----
    UUB. .BB.
    ||+---++-- Select 256 KiB outer PRG-ROM and 128 KiB outer CHR-ROM bank
    ++-------- Unknown, cleared when the menu is shown, both set once a game has been selected.

## Notes {#Notes}

- The initial value of this register must be \$00 for the multicart menu to be shown.
- The cart seems to run just fine when ignoring the two unknown bits.
