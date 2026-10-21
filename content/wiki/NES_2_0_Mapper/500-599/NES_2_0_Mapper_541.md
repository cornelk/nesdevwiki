---
title: "NES 2.0 Mapper 541"
url: "/wiki/NES_2.0_Mapper_541/"
description: "NES 2.0 Mapper 541 denotes the Korean 리틀컴 (LittleCom) 160-in-1 multicart."
wikiPage: "NES_2_0_Mapper_541"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 541 denotes the Korean *리틀컴 (LittleCom) 160-in-1* multicart.

## <span id="Address_latch_.28.24C000-.24FFFF.2C_write.29"></span>Address latch (\$C000-\$FFFF, write) {#Address_latch_.28.24C000-.24FFFF.2C_write.29}

    Mask: $C000

    A~FEDC BA98 7654 3210
      -------------------
      11.. .... PPPP PpBM
                |||| |||+- Nametable mirroring
                |||| |||    0: Horizontal
                |||| |||    1: Vertical
                |||| ||+-- Set PRG A14 behavior
                |||| ||     0: PRG A14=CPU A14 (NROM-256)
                |||| ||     1: PRG A14=p (NROM-128)
                ++++-++--- Set 16 KiB PRG-ROM bank (PRG A14..A19)
                           at CPU $8000-$BFFF/$C000-$FFFF,
                           p replaced with CPU A14 if B=0

## Notes {#Notes}

- CHR pattern data is stored in 8 KiB of unbanked CHR-RAM.
- Some games write to the unmapped CPU address range \$8000-\$BFFF. These writes have no effect.
- The mirroring bit works in the opposite way compared to most mappers.
