---
title: "INES Mapper 125"
url: "/wiki/INES_Mapper_125/"
description: "iNES Mapper 125 is used for Whirlwind Manu's ROM cartridge conversion of the Japanese Famicom Disk System game モンティのドキドキ大脱走 (Monty no Doki Doki Daisassō, Monty on the Run,…"
wikiPage: "INES_Mapper_125"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

iNES Mapper 125 is used for Whirlwind Manu's ROM cartridge conversion of the Japanese Famicom Disk System game *モンティのドキドキ大脱走* (*Monty no Doki Doki Daisassō*, Monty on the Run, cartridge code LH32). Its UNIF board name is **UNL-LH32**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG-ROM bank
- CPU \$8000-\$9FFF: 8 KiB fixed PRG-ROM bank \#12
- CPU \$A000-\$BFFF: 8 KiB fixed PRG-ROM bank \#13
- CPU \$C000-\$DFFF: 8 KiB unbanked PRG-RAM bank
- CPU \$E000-\$FFFF: 8 KiB fixed PRG-ROM bank \#15
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="PRG-ROM_Bank_.28.246000-.247FFF.29"></span>PRG-ROM Bank (\$6000-\$7FFF) {#PRG-ROM_Bank_.28.246000-.247FFF.29}

    Mask: probably $E000

    D~7654 3210
      ---------
      .... PPPP
           ++++- Select 8 KiB PRG-ROM bank number at CPU $6000-$7FFF

# Notes {#Notes}

- Mirroring is hard-wired (to Vertical).
- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Whirlwind Manu cartridge did not.
