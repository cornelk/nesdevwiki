---
title: "NES 2.0 Mapper 525"
url: "/wiki/NES_2.0_Mapper_525/"
description: "NES 2.0 Mapper 525 is used for mapper hacks of VRC2-using games. Specific known games include Kaiser bootleg versions of Contra and 月風魔伝 (Getsu Fūma Den). Its UNIF board name is UNL-KS7021A."
wikiPage: "NES_2_0_Mapper_525"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
---

NES 2.0 Mapper 525 is used for mapper hacks of VRC2-using games. Specific known games include Kaiser bootleg versions of *Contra* and 月風魔伝 (Getsu Fūma Den). Its UNIF board name is **UNL-KS7021A**.

## Registers {#Registers}

Mask: \$F000

     $8000: [.... PPP.] - Select 16KiB PRG bank at $8000-$BFFF
     $9000: [.... ...M] - like VRC2, 0: Vertical Mirroring (PPU A10); 1: Horizontal Mirroring (PPU A11)

Mask: \$F007

- \$B000: Select 1 KiB CHR-ROM bank at PPU \$0000-\$03FF
- \$B001: Select 1 KiB CHR-ROM bank at PPU \$0400-\$07FF
- \$B002: Select 1 KiB CHR-ROM bank at PPU \$0800-\$0BFF
- \$B003: Select 1 KiB CHR-ROM bank at PPU \$0C00-\$0FFF
- \$B004: Select 1 KiB CHR-ROM bank at PPU \$1000-\$13FF
- \$B005: Select 1 KiB CHR-ROM bank at PPU \$1400-\$17FF
- \$B006: Select 1 KiB CHR-ROM bank at PPU \$1800-\$1BFF
- \$B007: Select 1 KiB CHR-ROM bank at PPU \$1C00-\$1FFF

## Hardware {#Hardware}

The 7021A PCB contains a [7400](../7400/), a [7427](../7427/), a [7432](../7432/), a [7474](../7474/), a [74139](../74139/), a [74157](../74157/), a [74174](../74174/), two ~~74189s~~, and the two ROMs.

## Notes {#Notes}

- Getsu Fūma Den PCB images and analysis: [BBS](https://forums.nesdev.org/viewtopic.php?t=21091)
