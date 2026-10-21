---
title: "NES 2.0 Mapper 290"
url: "/wiki/NES_2.0_Mapper_290/"
description: "NES 2.0 Mapper 290 is used for the Asder 20-in-1 multicart. Its UNIF board name is BMC-NTD-03."
wikiPage: "NES_2_0_Mapper_290"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 290 is used for the *Asder 20-in-1* multicart. Its UNIF board name is **BMC-NTD-03**.

## Address latch {#Address_latch}

    Mask: $8000

    $8000: A~[1PPP PMCC Sp.. .CCC]
               ||| |||| ||    |||
               ||| ||++-------+++- Select 8 KiB CHR-ROM bank at PPU $0000-$1FFF
               ||| ||   |+-------- Select 16 KiB PRG-ROM bank at CPU $8000/$C000 if S=1
               ||| ||   +--------- Select PRG-ROM bank size, 0=32 KiB, 1=16 KiB
               ||| |+------------- Select nametable mirroring, 0=vertical, 1=horizontal 
               +++-+-------------- Select 32 KiB PRG-ROM bank at CPU $8000

## Notes {#Notes}

- The game *Benico* on that multicart has graphical glitches during attract mode on real hardware.
