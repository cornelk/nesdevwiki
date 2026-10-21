---
title: "NES 2.0 Mapper 303"
url: "/wiki/NES_2.0_Mapper_303/"
description: "NES 2.0 Mapper 303 is used for Kaiser's ROM cartridge conversion of the Famicom Disk System game アルマナの奇跡 (Almana no Kiseki). Its UNIF board name is UNL-KS7017."
wikiPage: "NES_2_0_Mapper_303"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 303 is used for Kaiser's ROM cartridge conversion of the Famicom Disk System game *アルマナの奇跡* (Almana no Kiseki). Its UNIF board name is **UNL-KS7017**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB unbanked PRG-RAM
- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB fixed PRG-ROM bank \#2
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="PRG_Bank_Latch_.28.244A00-.244AFF.29"></span>PRG Bank Latch (\$4A00-\$4AFF) {#PRG_Bank_Latch_.28.244A00-.244AFF.29}

    Mask: probably $FF00

    A~FEDC BA98 7654 3210
      -------------------
      0100 1010 .P.. PP..
                 +---++-- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF

The new PRG bank number is not applied until register \$5100 is written to.

## <span id="PRG_Bank_Latch_Commit_.28.245100.29"></span>PRG Bank Latch Commit (\$5100) {#PRG_Bank_Latch_Commit_.28.245100.29}

    Mask: probably $FF00

## <span id="Mirroring.2C_IRQ_.28.244020.2C_.244021.2C_.244025_bit_3:_write.2C_.244030_bit_0:_read.29"></span>Mirroring, IRQ (\$4020, \$4021, \$4025 bit 3: write, \$4030 bit 0: read) {#Mirroring.2C_IRQ_.28.244020.2C_.244021.2C_.244025_bit_3:_write.2C_.244030_bit_0:_read.29}

Identical to their respective equivalents on the [Famicom Disk System](../Family_Computer_Disk_System/).

# Note {#Note}

The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Kaiser cartridge did not.
