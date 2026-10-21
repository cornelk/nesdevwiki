---
title: "NES 2.0 Mapper 305"
url: "/wiki/NES_2.0_Mapper_305/"
description: "NES 2.0 Mapper 305 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of ドラキュラII 呪いの封印 (Dracula II: Noroi no Fūin). Its UNIF board name is…"
wikiPage: "NES_2_0_Mapper_305"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 305 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of *ドラキュラII 呪いの封印* (Dracula II: Noroi no Fūin). Its UNIF board name is **KS7031** (without any prefix).

# Banks {#Banks}

- CPU \$6000-\$67FF: 2 KiB switchable PRG-ROM bank
- CPU \$6800-\$6FFF: 2 KiB switchable PRG-ROM bank
- CPU \$7000-\$77FF: 2 KiB switchable PRG-ROM bank
- CPU \$7800-\$7FFF: 2 KiB switchable PRG-ROM bank
- CPU \$8000-\$87FF: 2 KiB fixed PRG-ROM bank \#15
- CPU \$8800-\$8FFF: 2 KiB fixed PRG-ROM bank \#14
- CPU \$9000-\$97FF: 2 KiB fixed PRG-ROM bank \#13
- CPU \$9800-\$9FFF: 2 KiB fixed PRG-ROM bank \#12
- CPU \$A000-\$A7FF: 2 KiB fixed PRG-ROM bank \#11
- CPU \$A800-\$AFFF: 2 KiB fixed PRG-ROM bank \#10
- CPU \$B000-\$B7FF: 2 KiB fixed PRG-ROM bank \#9
- CPU \$B800-\$BFFF: 2 KiB fixed PRG-ROM bank \#8
- CPU \$C000-\$C7FF: 2 KiB fixed PRG-ROM bank \#7
- CPU \$C800-\$CFFF: 2 KiB fixed PRG-ROM bank \#6
- CPU \$D000-\$D7FF: 2 KiB fixed PRG-ROM bank \#5
- CPU \$D800-\$DFFF: 2 KiB fixed PRG-ROM bank \#4
- CPU \$E000-\$E7FF: 2 KiB fixed PRG-ROM bank \#3
- CPU \$E800-\$EFFF: 2 KiB fixed PRG-ROM bank \#2
- CPU \$F000-\$F7FF: 2 KiB fixed PRG-ROM bank \#1
- CPU \$F800-\$FFFF: 2 KiB fixed PRG-ROM bank \#0
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="PRG-ROM_Bank_.28.248000-.24FFFF.29"></span>PRG-ROM Bank (\$8000-\$FFFF) {#PRG-ROM_Bank_.28.248000-.24FFFF.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1..R R... .... ....
         + +-------------- Select register number
                           $8000: Set 2 KiB PRG-ROM bank at CPU $6000-$67FF
                           $8800: Set 2 KiB PRG-ROM bank at CPU $6800-$6FFF
                           $9000: Set 2 KiB PRG-ROM bank at CPU $7000-$77FF
                           $9800: Set 2 KiB PRG-ROM bank at CPU $7800-$7FFF

# Notes {#Notes}

- Mirroring is hard-wired to Vertical. One of the two common UNIF ROM images of this game has no MIRR chunk.
- Save data is not retained, even as one of the two common UNIF ROM images has a BATT chunk.
- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Kaiser cartridge did not.
