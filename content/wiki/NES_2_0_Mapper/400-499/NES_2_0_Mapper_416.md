---
title: "NES 2.0 Mapper 416"
url: "/wiki/NES_2.0_Mapper_416/"
description: "NES 2.0 Mapper 416 denotes the PCB used for a 4-in-1 multicart that includes the N-32 cartridge conversion of SMB2J."
wikiPage: "NES_2_0_Mapper_416"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with fixed-timing cycle IRQs"
    url: "/category/Mappers_with_fixed-timing_cycle_IRQs/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 416** denotes the PCB used for a 4-in-1 multicart that includes the [N-32 cartridge conversion](../INES_Mapper_050/) of SMB2J.

# Banks {#Banks}

- CPU \$6000-\$7FFF: Fixed 8 KiB PRG-ROM bank \#7
- CPU \$8000-\$FFFF: Switchable 8/16/32 KiB PRG-ROM bank(s)
- PPU \$0000-\$1FFF: 8 KiB switchable CHR-ROM bank

# Registers {#Registers}

## Data Latch {#Data_Latch}

    Write to $8000 (mask $8000):
    D~[BPA. Mcc.]
       |||  |++--- CHR A14..A13
       |||  |+---- 1=Horizontal, 0=Vertical mirroring
       |||  +----- PRG A16, also select SMB2J (0)/Normal (1) mode
       ||+-------- PRG A14 in Normal mode
       +|--------- PRG A15 in Normal mode
       ++--------- 00: NROM-064 banking (PRG A13=0, PRG A14=A)
                   01: NROM-128 banking (PRG A13=CPU A13, PRG A14=A)
                   1x: NROM-256 banking (PRG A13=CPU A13, PGR A14=CPU A14)

## <span id="PRG_Bank_Select_.28SMB2J_mode_only.29"></span>PRG Bank Select (SMB2J mode only) {#PRG_Bank_Select_.28SMB2J_mode_only.29}

    Write to $4020 (mask unknown)
    D~[.... .BAC]
             ||+- PRG A15
             ++-- PRG A14..13

## <span id="IRQ_Control_.28SMB2J_mode_only.29"></span>IRQ Control (SMB2J mode only) {#IRQ_Control_.28SMB2J_mode_only.29}

    Write to $4120 (mask unknown)
    D~[.... ...E]
               +- 1=Counting, generate IRQ on 12-bit counter overflow
                  0=Counter reset, acknowledge/disable IRQ
