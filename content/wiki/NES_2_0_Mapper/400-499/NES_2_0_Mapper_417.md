---
title: "NES 2.0 Mapper 417"
url: "/wiki/NES_2.0_Mapper_417/"
description: "NES 2.0 Mapper 417 denotes the PCB of the \"Fine Studio\" bootleg copy of Batman: The Video Game (submapper 0) and of RoboCop (submapper 1)."
wikiPage: "NES_2_0_Mapper_417"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

**NES 2.0 Mapper 417** denotes the PCB of the "Fine Studio" bootleg copy of *Batman: The Video Game* (submapper 0) and of *RoboCop* (submapper 1).

# Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG-ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG-ROM bank
- CPU \$E000-\$FFFF: 8 KiB fixed PRG-ROM bank, last
- PPU \$0000-\$03FF: 1 KiB switchable CHR-ROM bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR-ROM bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR-ROM bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR-ROM bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR-ROM bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR-ROM bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR-ROM bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR-ROM bank
- PPU \$2000-\$23FF: 1 KiB switchable CIRAM bank
- PPU \$2400-\$27FF: 1 KiB switchable CIRAM bank
- PPU \$2800-\$2BFF: 1 KiB switchable CIRAM bank
- PPU \$2C00-\$2FFF: 1 KiB switchable CIRAM bank

# Registers {#Registers}

## PRG-ROM Bank Select {#PRG-ROM_Bank_Select}

    Mask: $8073
    Write to $8000: Select 8 KiB PRG-ROM bank at CPU $8000
    Write to $8001: Select 8 KiB PRG-ROM bank at CPU $A000
    Write to $8002: Select 8 KiB PRG-ROM bank at CPU $C000

## CHR-ROM Bank Select {#CHR-ROM_Bank_Select}

    Mask: $8073
    Write to $8010: Select 1 KiB CHR-ROM bank at PPU $0000
    Write to $8011: Select 1 KiB CHR-ROM bank at PPU $0400
    Write to $8012: Select 1 KiB CHR-ROM bank at PPU $0800
    Write to $8013: Select 1 KiB CHR-ROM bank at PPU $0C00
    Write to $8020: Select 1 KiB CHR-ROM bank at PPU $1000
    Write to $8021: Select 1 KiB CHR-ROM bank at PPU $1400
    Write to $8022: Select 1 KiB CHR-ROM bank at PPU $1800
    Write to $8023: Select 1 KiB CHR-ROM bank at PPU $1C00

Additional meaning for Submapper 1:

    Write to $8010: (bit 7) Select 1 KiB CIRAM bank at PPU $2000
    Write to $8011: (bit 7) Select 1 KiB CIRAM bank at PPU $2400
    Write to $8012: (bit 7) Select 1 KiB CIRAM bank at PPU $2800
    Write to $8013: (bit 7) Select 1 KiB CIRAM bank at PPU $2C00

## IRQ Control {#IRQ_Control}

    Mask: $8070
    Write to $8030: Reset counter and enable IRQ
    Write to $8040: Disable and acknowledge IRQ

The counter increases on every M2 cycle regardless of IRQ enable status. If enabled, an IRQ is generated upon 10-bit (submapper 0)/12-bit (submapper 1) counter overflow.

## <span id="CIRAM_Bank_Select_.28Submapper_0_only.29"></span>CIRAM Bank Select (Submapper 0 only) {#CIRAM_Bank_Select_.28Submapper_0_only.29}

    Mask: $8073
    Write to $8050: Select 1 KiB CIRAM bank at PPU $2000
    Write to $8051: Select 1 KiB CIRAM bank at PPU $2400
    Write to $8052: Select 1 KiB CIRAM bank at PPU $2800
    Write to $8053: Select 1 KiB CIRAM bank at PPU $2C00
