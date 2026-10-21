---
title: "INES Mapper 103"
url: "/wiki/INES_Mapper_103/"
description: "iNES Mapper 103 describes a specific pirate port of the FDS version of Doki Doki Panic."
wikiPage: "INES_Mapper_103"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 103 describes a specific pirate port of the FDS version of Doki Doki Panic.

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB RAM, or 8 KiB controllable ROM bank.
- CPU \$8000-\$B7FF: 14 KiB ROM, starting at ROM offset 0x18000
- CPU \$B800-\$D7FF: 8 KiB RAM, or 8 KiB fixed ROM bank, starting at ROM offset 0x1B800
- CPU \$D800-\$FFFF: 10 KiB ROM, starting at ROM offset 0x1D800

Writes to regions where RAM can be mapped will always write to RAM, even if RAM isn't enabled for reading.

## Registers {#Registers}

Registers are not cleared on reset, and state is unknown on power-up. No bus conflicts.

### <span id="PRG_bank_.28.248000-.248FFF.29"></span>PRG bank (\$8000-\$8FFF) {#PRG_bank_.28.248000-.248FFF.29}

Mask: \$F000

    7  bit  0
    ---- ----
    .... PPPP
         ||||
         ++++-- select 8 KiB PRG bank at $6000

### <span id="Mirroring_control_.28.24E000-.24EFFF.29"></span>Mirroring control (\$E000-\$EFFF) {#Mirroring_control_.28.24E000-.24EFFF.29}

Mask: \$F000

    7  bit  0
    ---- ----
    .... m...
         |
         +----- 0=V (PPU A10), 1=H (PPU A11)

### <span id="RAM_disable_.28.24F000-.24FFFF.29"></span>RAM disable (\$F000-\$FFFF) {#RAM_disable_.28.24F000-.24FFFF.29}

Mask: \$F000

    7  bit  0
    ---- ----
    ...d ....
       |
       +------- 0=RAM available at $6000-$7FFF and $B800-$D7FF
                1=ROM instead

## Hardware {#Hardware}

This hardware involves three [7400](../7400/), two [7410](../7410/), a [7432](../7432/), a [7474](../7474/), a [74138](../74138/), a [74157](../74157/), and a [74161](../74161/).

## See also {#See_also}

- [Krzysiobal's reverse engineered documentation](https://forums.nesdev.org/viewtopic.php?f=9&t=18956)
