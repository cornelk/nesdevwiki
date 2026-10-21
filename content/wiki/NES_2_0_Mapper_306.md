---
title: "NES 2.0 Mapper 306"
url: "/wiki/NES_2.0_Mapper_306/"
description: "NES 2.0 Mapper 306 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System game Exciting Basket (Double Dribble in North America and Europe). Its UNIF board name is…"
wikiPage: "NES_2_0_Mapper_306"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 306 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System game *Exciting Basket* (*Double Dribble* in North America and Europe). Its UNIF board name is **UNL-KS7016**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG-ROM bank
- CPU \$8000-\$FFFF: 32 KiB fixed PRG-ROM bank \#3
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="PRG-ROM_Bank_.28.24D903-.24D943.29"></span>PRG-ROM Bank (\$D903-\$D943) {#PRG-ROM_Bank_.28.24D903-.24D943.29}

    Mask: $D903 (according to FCEUX)

    A~FEDC BA98 7654 3210
      -------------------
      11.1 1..1 .Mpp PP11
                 ||| ++-- Bits 0-1 of 8 KiB PRG-ROM bank number at CPU $6000-$7FFF
                 |++----- Bits 2-3 of 8 KiB PRG-ROM bank number at CPU $6000-$7FFF if M=1
                 +------- 0=Ignore pp, switchable bank number is 10PP
                          1=Obey pp, switchable bank number is ppPP

# Notes {#Notes}

- Mirroring is hard-wired (to Vertical).
- Switchable bank numbers \$0C-\$0F are forced to be mirrors of \$0B.
- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Kaiser cartridge did not.
