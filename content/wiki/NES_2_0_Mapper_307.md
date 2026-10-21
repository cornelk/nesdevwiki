---
title: "NES 2.0 Mapper 307"
url: "/wiki/NES_2.0_Mapper_307/"
description: "NES 2.0 Mapper 307 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of Metroid. Its UNIF board name is UNL-KS7037."
wikiPage: "NES_2_0_Mapper_307"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 307 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of *Metroid*. Its UNIF board name is **UNL-KS7037**.

# Banks {#Banks}

- CPU \$6000-\$6FFF: 4 KiB fixed PRG-RAM bank \#0
- CPU \$7000-\$7FFF: 4 KiB fixed PRG-ROM bank \#15
- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$AFFF: 4 KiB fixed PRG-ROM bank \#28
- CPU \$B000-\$BFFF: 4 KiB fixed PRG-RAM bank \#1
- CPU \$C000-\$DFFF: 8 KiB switchable PRG-ROM bank
- CPU \$E000-\$FFFF: 8 KiB fixed PRG-ROM bank \#15
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id=".248000-.249FFE_.28even.29"></span>\$8000-\$9FFE (even) {#.248000-.249FFE_.28even.29}

Select bank register number

    $0-$1: No function
    $2: Bit 0 selects 1 KiB CIRAM bank at PPU $2000-$23FF
    $3: Bit 0 selects 1 KiB CIRAM bank at PPU $2800-$2BFF
    $4: Bit 0 selects 1 KiB CIRAM bank at PPU $2400-$27FF
    $5: Bit 0 selects 1 KiB CIRAM bank at PPU $2C00-$2FFF
    $6: Select 8 KiB PRG-ROM bank at CPU $8000-$9FFF
    $7: Select 8 KiB PRG-ROM bank at CPU $C000-$DFFF

## <span id=".248001-.249FFF_.28odd.29"></span>\$8001-\$9FFF (odd) {#.248001-.249FFF_.28odd.29}

Data for bank register selected by \$8000-\$9FFE (even).

# Notes {#Notes}

- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Kaiser cartridge did not.
- Game progress cannot be saved; in fact, the save menu has been removed.
- The register layout indicates that an MMC3 or N118 clone is being used, connected in an unusual manner.
- Reverse-engineered schematic: [forum](https://forums.nesdev.org/viewtopic.php?p=242196#p242196)
