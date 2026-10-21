---
title: "NES 2.0 Mapper 285"
url: "/wiki/NES_2.0_Mapper_285/"
description: "NES 2.0 Mapper 285 is used for the A65AS and JY-066 multicarts. The UNIF board name for both incompatible variants is BMC-A65AS. A second UNIF board name, BMC-GF-401CD refers to submapper 1. A…"
wikiPage: "NES_2_0_Mapper_285"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 285** is used for the *A65AS* and *JY-066* multicarts. The UNIF board name for both incompatible variants is **BMC-A65AS**. A second UNIF board name, **BMC-GF-401CD** refers to submapper 1. A later-discovered Russian bootleg cartridge of *Highway Star* (Русские Гонки) is assigned to mapper 285 submapper 1 as well, as it uses UNROM-banking while changing mirroring via latch bit 3.

# Banks {#Banks}

- CPU \$8000-\$BFFF: In 16 KiB mode: 16 KiB switchable inner bank, 128 KiB switchable outer bank
- CPU \$C000-\$FFFF: In 16 KiB mode: 16 KiB fixed inner bank 7, 128 KiB switchable outer bank
- CPU \$8000-\$FFFF: In 32 KiB mode: 32 KiB switchable bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# <span id="Data_Latch_.28CPU_.248000-.24FFFF.29"></span>Data Latch (CPU \$8000-\$FFFF) {#Data_Latch_.28CPU_.248000-.24FFFF.29}

## <span id="A65AS_.28Submapper_0.29"></span>A65AS (Submapper 0) {#A65AS_.28Submapper_0.29}

    Bit 7654 3210
        ---------
        SPMO OBBB
        |||| |+++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF in 16 KiB mode (bit 6=0)
        |||| |++-- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF in 32 KiB mode (bit 6=1)
        |||+-+---- Select 128 KiB outer PRG-ROM bank at CPU $8000-$FFFF in 16 KiB mode (bit 6=0)
        ||+------- Select nametable mirroring
        ||         0: Vertical (bit 7=0)/Single-screen, page 0 (bit 7=1)
        ||         1: Horizontal (bit 7=0)/Single-screen, page 1 (bit 7=1)
        |+-------- Select PRG banking mode
        |          0: 16 KiB
        |          1: 32 KiB
        +--------- Select nametable mirroring type
                   0: Horizontal/vertical
                   1: Single-screen

## <span id="JY-066_.28Submapper_1.29"></span>JY-066 (Submapper 1) {#JY-066_.28Submapper_1.29}

    Bit 7654 3210
        ---------
        SPMO HBBB
        |||| |+++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF in 16 KiB mode (bit 6=0)
        |||| |++-- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF in 32 KiB mode (bit 6=1)
        |||| +---- Select nametable mirroring if bit 7=0
        ||||       0: Vertical
        ||||       1: Horizontal
        ||++------ Select 128 KiB outer PRG-ROM bank at CPU $8000-$FFFF in 16 KiB mode (bit 6=0)
        ||+------- Select nametable mirroring if bit 7=1
        ||         0: Single-screen, page 0
        ||         1: Single-screen, page 1
        |+-------- Select PRG banking mode
        |          0: 16 KiB
        |          1: 32 KiB
        +--------- Select nametable mirroring type
                   0: Horizontal/vertical (selected via bit 3)
                   1: Single-screen (selected via bit 5)
