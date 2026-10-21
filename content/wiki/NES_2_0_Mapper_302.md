---
title: "NES 2.0 Mapper 302"
url: "/wiki/NES_2.0_Mapper_302/"
description: "NES 2.0 Mapper 302 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of Gyruss. Its UNIF board name is UNL-KS7057."
wikiPage: "NES_2_0_Mapper_302"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

NES 2.0 Mapper 302 is used for Kaiser's ROM cartridge conversion of the Japanese Famicom Disk System version of *Gyruss*. Its UNIF board name is **UNL-KS7057**.

# Banks {#Banks}

- CPU \$6000-\$67FF: 2 KiB switchable PRG-ROM bank
- CPU \$6800-\$6FFF: 2 KiB switchable PRG-ROM bank
- CPU \$7000-\$77FF: 2 KiB switchable PRG-ROM bank
- CPU \$7800-\$7FFF: 2 KiB switchable PRG-ROM bank
- CPU \$8000-\$87FF: 2 KiB switchable PRG-ROM bank
- CPU \$8800-\$8FFF: 2 KiB switchable PRG-ROM bank
- CPU \$9000-\$97FF: 2 KiB switchable PRG-ROM bank
- CPU \$9800-\$9FFF: 2 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB fixed PRG-ROM bank \#\$0D
- CPU \$C000-\$FFFF: 16 KiB fixed PRG-ROM bank \#\$07
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="Mirroring_select_.28.248000-.249FFF.29"></span>Mirroring select (\$8000-\$9FFF) {#Mirroring_select_.28.248000-.249FFF.29}

    Mask: $E000

    D~7654 3210
      ---------
      .... ...M
              +- Select nametable mirroring type
                 0: Horizontal
                 1: Vertical

## <span id="PRG-ROM_Bank_.28.24B000-.24E003.29"></span>PRG-ROM Bank (\$B000-\$E003) {#PRG-ROM_Bank_.28.24B000-.24E003.29}

    Mask: $F003

    A~FEDC BA98 7654 3210
      -------------------
      1RRR .... .... ..RN
       |||             |+- 0: Select low nibble
       |||             |   1: Select high nibble
       +++-------------+-- Select register number
                           $B000/$B001: Set 2 KiB PRG-ROM bank at CPU $8000-$87FF
                           $B002/$B003: Set 2 KiB PRG-ROM bank at CPU $8800-$8FFF
                           $C000/$C001: Set 2 KiB PRG-ROM bank at CPU $9000-$97FF
                           $C002/$C003: Set 2 KiB PRG-ROM bank at CPU $9800-$9FFF
                           $D000/$D001: Set 2 KiB PRG-ROM bank at CPU $6000-$67FF
                           $D002/$D003: Set 2 KiB PRG-ROM bank at CPU $6800-$6FFF
                           $E000/$E001: Set 2 KiB PRG-ROM bank at CPU $7000-$77FF
                           $E002/$E003: Set 2 KiB PRG-ROM bank at CPU $7800-$7FFF

# See Also {#See_Also}

[Krzysiobal's reverse-engineering notes of KS204](https://forums.nesdev.org/viewtopic.php?t=19377) [Krzysiobal's reverse-engineering notes of KS7057](https://forums.nesdev.org/viewtopic.php?t=19310)
