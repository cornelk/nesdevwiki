---
title: "NES 2.0 Mapper 372"
url: "/wiki/NES_2.0_Mapper_372/"
description: "NES 2.0 Mapper 372 is used for a revision of the Rockman I-VI multicart (PCB ID SFC-12). It is INES Mapper 045 but with one bit of outer bank register #2 working as a CHR-ROM/RAM switch."
wikiPage: "NES_2_0_Mapper_372"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 372 is used for a revision of the Rockman I-VI multicart (PCB ID **SFC-12**). It is [INES Mapper 045](../INES_Mapper_045/) but with one bit of outer bank register \#2 working as a CHR-ROM/RAM switch.

## <span id="CHR-AND.2C_CHR-OR.2FPRG-OR_MSB.2C_CHR_source_.28.246000_.232.2C_write.29"></span>CHR-AND, CHR-OR/PRG-OR MSB, CHR source (\$6000 \#2, write) {#CHR-AND.2C_CHR-OR.2FPRG-OR_MSB.2C_CHR_source_.28.246000_.232.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      PPRC cccc
      |||| ++++- Select number of CHR bits from MMC3
      ||||       ($F: 256 KiB, $E: 128 KiB ... $7-$0: 1 KiB)
      |||+------ Select CHR A18
      ||+------- Select CHR-ROM (0)/CHR-RAM (1, unbanked)
      ++-------- Select CHR A20-A21 and PRG A21-22

Note that the CHR-RAM bit works in the opposite direction from [NES 2.0 Mapper 356](../NES_2.0_Mapper_356/).
