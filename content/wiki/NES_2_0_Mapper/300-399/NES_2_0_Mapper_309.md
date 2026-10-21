---
title: "NES 2.0 Mapper 309"
url: "/wiki/NES_2.0_Mapper_309/"
description: "NES 2.0 Mapper 309 is used for Whirlwind Manu's ROM cartridge conversion of the Japanese Famicom Disk System game 愛戦士ニコル (Ai Senshi Nicol, cartridge code LH51). Its UNIF board name is…"
wikiPage: "NES_2_0_Mapper_309"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 309 is used for Whirlwind Manu's ROM cartridge conversion of the Japanese Famicom Disk System game *愛戦士ニコル* (*Ai Senshi Nicol*, cartridge code LH51). Its UNIF board name is **UNL-LH51**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB unbanked PRG-RAM
- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB fixed PRG-ROM bank \#13
- CPU \$C000-\$DFFF: 8 KiB fixed PRG-ROM bank \#14
- CPU \$E000-\$FFFF: 8 KiB fixed PRG-ROM bank \#15
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="PRG-ROM_Bank_.28.248000-9FFF.29"></span>PRG-ROM Bank (\$8000-9FFF) {#PRG-ROM_Bank_.28.248000-9FFF.29}

    Mask: probably $E000

    D~7654 3210
      ---------
      .... PPPP
           ++++- Select 8 KiB PRG-ROM bank number at CPU $8000-$9FFF

## <span id="Mirroring_.28.24E000-9FFF.29"></span>Mirroring (\$E000-9FFF) {#Mirroring_.28.24E000-9FFF.29}

    Mask: probably $E000

    D~7654 3210
      ---------
      .... M...
           +---- Select nametable mirroring type
                  0: Vertical
                  1: Horizontal

# Notes {#Notes}

- This particular conversion of the game has broken sound [even on real hardware](https://www.youtube.com/watch?v=Rgt9cWk0yh0). Other conversions of the game that run under [INES Mapper 042](../INES_Mapper_042/) do not suffer from this problem.
- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Whirlwind Manu cartridge did not. Given that the sound routine itself is broken however, doing such a thing may not be worthwhile.
