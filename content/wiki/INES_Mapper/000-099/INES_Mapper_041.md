---
title: "INES Mapper 041"
url: "/wiki/INES_Mapper_041/"
description: "iNES Mapper 041 represents Caltron 6-in-1, a discrete logic multicart. It appears to have been designed to hold up to four each unmodified CNROM and NROM games (with NROMs' CHR in unused portions of…"
wikiPage: "INES_Mapper_041"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

iNES Mapper 041 represents Caltron 6-in-1, a discrete logic multicart. It appears to have been designed to hold up to four each unmodified [CNROM](../CNROM/) and [NROM](../NROM/) games (with NROMs' CHR in unused portions of the CNROM games).

## Overview {#Overview}

- PRG ROM size: 256 KiB
- PRG ROM bank size: 32 KiB
- PRG RAM: Impossible
- CHR capacity: 128 KiB ROM
- CHR bank size: 8 KiB inner / 32 KiB outer
- Nametable [mirroring](../Mirroring/): Mapper-controlled H or V
- Subject to [bus conflicts](../Bus_conflict/): partly

## Registers {#Registers}

Both registers are cleared on powerup, or while the reset button is pressed.

### <span id="Outer_Bank_Select_.28.246000-.2467FF.29"></span>Outer Bank Select (\$6000-\$67FF) {#Outer_Bank_Select_.28.246000-.2467FF.29}

    15   11   7  bit  0 (address lines)
    ---- ---- ---- ----
    0110 0xxx xxMC CEPP
                ││ ││││
                ││ │├┴┴── Select 32KiB PRG ROM bank for CPU $8000─$FFFF
                ││ ││
                ││ │└──── If high, enable Inner Bank Select
                │└─┴───── Select 32KiB outer CHR ROM bank
                └──────── Mirroring (0=Vertical, 1=Horizontal)

### <span id="Inner_CHR_Bank_Select_.28.248000-.24FFFF.29.2C_bus_conflicts"></span>Inner CHR Bank Select (\$8000-\$FFFF), bus conflicts {#Inner_CHR_Bank_Select_.28.248000-.24FFFF.29.2C_bus_conflicts}

    7  bit  0
    ---- ----
    xxZZ xxcc
      ││   ││
      ││   └┴── Select 8KiB inner bank of outer 32KiB CHR bank chosen above
      └┴─────── Extant but unused

Note that the Inner CHR Bank Select only can be written while the PRG ROM bank is 4, 5, 6, or 7. This means that execution must pass through one of the upper four PRG banks to choose the inner CHR bank before playing any games in the lower four PRG banks.
