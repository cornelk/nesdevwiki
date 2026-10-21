---
title: "CPROM"
url: "/wiki/CPROM/"
description: "NES-CPROM (no HVC-CPROM known) is a particular Nintendo cartridge board which uses discrete logic to provide up to four 4 KB banks of CHR RAM. The iNES format assigns mapper 13 to this board. Only…"
wikiPage: "CPROM"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

NES-CPROM (no [HVC](../Family_Computer/)-CPROM known) is a particular Nintendo cartridge board which uses discrete logic to provide up to four 4 KB banks of CHR RAM. The [iNES](../INES/) format assigns **mapper 13** to this board. Only *[Videomation](https://en.wikipedia.org/wiki/Videomation)*, a paint program for NES, is known to use it.

## Overview {#Overview}

- PRG ROM size: 32 KiB (DIP-28 standard pinout)
- PRG ROM bank size: Not bankswitched
- PRG RAM: None
- CHR capacity: 16 KiB RAM (in two [8KiB RAMs](../6264_static_RAM/))
- CHR bank size: 4 KiB
- Nametable [mirroring](../Mirroring/): Vertical
- Subject to [bus conflicts](../Bus_conflict/): Yes

## Banks {#Banks}

- PPU \$0000-\$0FFF: 4 KiB fixed CHR RAM bank (first page)
- PPU \$1000-\$1FFF: 4 KiB swappable CHR RAM page

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxxx xxCC
           ||
           ++- Select 4 KiB CHR RAM bank for PPU $1000-$1FFF

## Hardware {#Hardware}

The CPROM board uses a [74HC161](../74161/) (4-bit latch) and a [74HC08](../7408/) (quad 2-input AND gate), with a [74HC04](../7404/) (hex inverter) to unify the two RAMs.

## See also {#See_also}

- [Kevtris's notes](http://kevtris.org/mappers/nes_discrete/NES_CPROM.html)
- The [NROM-368](../NROM-368/) proposal is also compatible with this board.
