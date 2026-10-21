---
title: "NES 2.0 Mapper 415"
url: "/wiki/NES_2.0_Mapper_415/"
description: "NES 2.0 Mapper 415 denotes the 0353 circuit board for a cartridge conversion of the FDS game Roger Rabbit, renamed to Lucky Rabbit."
wikiPage: "NES_2_0_Mapper_415"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
---

**NES 2.0 Mapper 415** denotes the **0353** circuit board for a cartridge conversion of the FDS game *Roger Rabbit*, renamed to *Lucky Rabbit*.

# Banks {#Banks}

- CPU \$6000-\$7FFF: Switchable 8 KiB PRG-ROM bank
- CPU \$8000-\$FFFF: Fixed last 32 KiB PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Data latch {#Data_latch}

    Mask: $8000
    D~[...M PPPP]
          | ++++- PRG A16..A13 at CPU $6000-$7FFF
          +------ 0=Vertical, 1=Horizontal mirroring
