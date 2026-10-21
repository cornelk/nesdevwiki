---
title: "INES Mapper 042"
url: "/wiki/INES_Mapper_042/"
description: "iNES Mapper 042 was used for hacked FDS games converted to cartridge form. At least two examples are known: Ai Senshi Nicol and \"Mario Baby\" (really Bio Miracle Bokutte Upa). While they likely used…"
wikiPage: "INES_Mapper_042"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

iNES Mapper 042 was used for hacked FDS games converted to cartridge form. At least two examples are known: Ai Senshi Nicol and "Mario Baby" (really Bio Miracle Bokutte Upa). While they likely used different hardware, they can both be encompassed by a single mapper.

## Overview {#Overview}

- PRG ROM size: 128 KiB
- PRG ROM bank size: 8 KiB
- PRG RAM: No
- CHR capacity: 128KiB ROM or 8KiB RAM
- CHR bank size: 8 KiB
- Nametable mirroring: Mapper controlled
- Subject to [bus conflicts](../Bus_conflict/): ??

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8KiB switchable PRG ROM bank
- CPU \$8000-\$FFFF: Fixed to last 32 KiB of PRG ROM
- PPU \$0000-\$1FFF: 8KiB switchable CHR ROM bank or 8KiB fixed CHR RAM

## Registers {#Registers}

Register addresses should be masked with \$E003.

### <span id="CHR_Select_.28.248000.29"></span>CHR Select (\$8000) {#CHR_Select_.28.248000.29}

    7  bit  0
    ---- ----
    xxxx CCCC
         ||||
         ++++- Select 8 KB CHR ROM bank for PPU $0000-$1FFF

This register is ignored on Bio Miracle Bokutte Upa, which uses 8KiB CHR RAM instead.

### <span id="PRG_Select_.28.24E000.29"></span>PRG Select (\$E000) {#PRG_Select_.28.24E000.29}

    7  bit  0
    ---- ----
    xxxx PPPP
         ||||
         ++++- Select 8 KB PRG ROM bank for CPU $6000-$7FFF

### <span id="Mirroring_Control_.28.24E001.29"></span>Mirroring Control (\$E001) {#Mirroring_Control_.28.24E001.29}

    7  bit  0
    ---- ----
    xxxx Mxxx
         |
         +---- Mirroring (0: Vertical; 1: Horizontal)

### <span id="IRQ_Control_.28.24E002.29"></span>IRQ Control (\$E002) {#IRQ_Control_.28.24E002.29}

    7  bit  0
    ---- ----
    xxxx xxEx
           |
           +-- IRQ (0: Disable, Acknowledge, and Reset; 1: Enable)

The IRQ hardware counts M2 cycles and triggers after 24576 (\$6000) have elapsed. FCEUX's source implies that the IRQ hardware is a 15-bit counter, and IRQ is asserted while the two MSBs are set, i.e. in the absence of CPU involvement, IRQ will be be not asserted for 24576 cycles, then asserted for 8192 cycles, and repeat endlessly.

## See also {#See_also}

- [Pirate game "Mario Baby" Mapper \#42 Info](http://nesdev.org/42.txt) by The Mad Dumper.
