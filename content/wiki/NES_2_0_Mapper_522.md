---
title: "NES 2.0 Mapper 522"
url: "/wiki/NES_2.0_Mapper_522/"
description: "NES 2.0 Mapper 522 is used for Whirlwind Manu's cartridge conversion of the Famicom Disk System game 風雲 少林拳 (Fūun Shōrinken). Its UNIF board name is UNL-LH10."
wikiPage: "NES_2_0_Mapper_522"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 522 is used for Whirlwind Manu's cartridge conversion of the Famicom Disk System game 風雲 少林拳 (Fūun Shōrinken). Its UNIF board name is **UNL-LH10**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB PRG-ROM, hard-wired to second-to-last bank
- CPU \$8000-\$9FFF: 8 KiB PRG-ROM, switchable
- CPU \$A000-\$BFFF: 8 KiB PRG-ROM, switchable
- CPU \$C000-\$DFFF: 8 KiB PRG-RAM
- CPU \$E000-\$FFFF: 8 KiB PRG-ROM, hard-wired to last bank
- PPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM

# Registers {#Registers}

## <span id=".248000-.249FFE_.28even.29"></span>\$8000-\$9FFE (even) {#.248000-.249FFE_.28even.29}

Select bank register number

    $0-$5: No function
    $6: Select 8 KiB PRG-ROM bank at CPU $8000-$9FFF
    $7: Select 8 KiB PRG-ROM bank at CPU $A000-$BFFF

## <span id=".248001-.249FFF_.28odd.29"></span>\$8001-\$9FFF (odd) {#.248001-.249FFF_.28odd.29}

Data for bank register selected by \$8000-\$9FFE (even).

# Notes {#Notes}

- Mirroring is hard-wired (to Vertical in Fūun Shōrinken's case).
- The register layout indicates that an MMC3 clone is being used, connected in a restricted manner.
