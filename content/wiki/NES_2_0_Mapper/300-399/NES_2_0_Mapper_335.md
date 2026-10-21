---
title: "NES 2.0 Mapper 335"
url: "/wiki/NES_2.0_Mapper_335/"
description: "NES 2.0 mapper 335 is used for a 10-in-1 multicart. Its UNIF board name is BMC-CTC-09."
wikiPage: "NES_2_0_Mapper_335"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 335 is used for a *10-in-1* multicart. Its UNIF board name is **BMC-CTC-09**.

## <span id="PRG-ROM_and_Mirroring_Register_.28.24C000-.24FFFF.29"></span>PRG-ROM and Mirroring Register (\$C000-\$FFFF) {#PRG-ROM_and_Mirroring_Register_.28.24C000-.24FFFF.29}

    Mask: probably $C000

    D~7654 3210
      ---------
      ..MB pPPP
        || |+++- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF
        || +---- Select 16 KiB PRG-ROM bank if B=1
        |+------ Select PRG A14 Mode
        |         0: PRG A14=CPU A14 (NROM-256)
        |         1: PRG A14=p (NROM-128)
        +------- Select nametable mirroring mode
                  0: Vertical
                  1: Horizontal         

## <span id="CHR-ROM_Register_.28.248000-.24BFFF.29"></span>CHR-ROM Register (\$8000-\$BFFF) {#CHR-ROM_Register_.28.248000-.24BFFF.29}

    Mask: probably $C000

    D~7654 3210
      ---------
      .... CCCC
           ++++- Select 8 KiB CHR-ROM bank at PPU $0000-$1FFF
