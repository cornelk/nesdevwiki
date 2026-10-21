---
title: "NES 2.0 Mapper 297"
url: "/wiki/NES_2.0_Mapper_297/"
description: "NES 2.0 Mapper 297 denotes the TXC 01-22110-000 board, used for the multicart 2-in-1 Uzi Lightgun (MGC-002). Its hardware can function as an MMC1 or as INES Mapper 070."
wikiPage: "NES_2_0_Mapper_297"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 297 denotes the TXC 01-22110-000 board, used for the multicart *2-in-1 Uzi Lightgun* (MGC-002). Its hardware can function as an [MMC1](../MMC1/) or as [INES Mapper 070](../INES_Mapper_070/).

# Registers {#Registers}

## <span id="Mode_register_.28.244120.29"></span>Mode register (\$4120) {#Mode_register_.28.244120.29}

Mask: \$E100

    7654 3210
    ---------
    .... ..AM
           |+- Mapper mode and outer 128 KiB PRG/CHR-ROM bank
           |   0: Mapper 70, Latch at $8000-$FFFF, PRG/CHR A17=0
           |   1: Mapper 1, MMC1-compatible registers at $8000-$FFFF, PRG/CHR A17=1
           +-- PRG A16 in Mapper 70 mode

## <span id="Data_latch_in_Mapper_70_mode_.28.248000-.24FFFF.29"></span>Data latch in Mapper 70 mode (\$8000-\$FFFF) {#Data_latch_in_Mapper_70_mode_.28.248000-.24FFFF.29}

See [INES Mapper 070](../INES_Mapper_070/).

## <span id="MMC1_registers_in_Mapper_1_mode_.28.248000-.24FFFF.29"></span>MMC1 registers in Mapper 1 mode (\$8000-\$FFFF) {#MMC1_registers_in_Mapper_1_mode_.28.248000-.24FFFF.29}

See [INES Mapper 001](../MMC1/).

# Nametable mirroring {#Nametable_mirroring}

Mirroring is fixed to Vertical by way of an additional wire connecting CIRAM A10 to PA10.

# See also {#See_also}

[PCB image and analysis](http://forums.nesdev.org/viewtopic.php?f=9&t=17067)
