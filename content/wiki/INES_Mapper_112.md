---
title: "INES Mapper 112"
url: "/wiki/INES_Mapper_112/"
description: "iNES Mapper 112 is quite similar to iNES Mapper 206, which resembles an simplified MMC3; however, the registers appears to be scrambled a bit."
wikiPage: "INES_Mapper_112"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

***iNES Mapper 112*** is quite similar to [iNES Mapper 206](../INES_Mapper_206/), which resembles an simplified [MMC3](../MMC3/); however, the registers appears to be scrambled a bit.

Example Games:

- Huang Di
- San Guo Zhi - Qun Xiong Zheng Ba

## Registers {#Registers}

Mask: \$E001

### <span id="Bank_select_.28.248000.29"></span>Bank select (\$8000) {#Bank_select_.28.248000.29}

    7  bit  0
    ---- ----
    xxxx xRRR
          |||
          +++- Specify which bank register to update on next write to Bank Data register
               0: Select 8 KB PRG ROM bank at $8000-$9FFF
               1: Select 8 KB PRG ROM bank at $A000-$BFFF
               2: Select 2 KB CHR bank at PPU $0000-$07FF
               3: Select 2 KB CHR bank at PPU $0800-$0FFF
               4: Select 1 KB CHR bank at PPU $1000-$13FF
               5: Select 1 KB CHR bank at PPU $1400-$17FF
               6: Select 1 KB CHR bank at PPU $1800-$1BFF
               7: Select 1 KB CHR bank at PPU $1C00-$1FFF

### <span id="Bank_data_.28.24A000.29"></span>Bank data (\$A000) {#Bank_data_.28.24A000.29}

    7  bit  0
    ---- ----
    DDDD DDDD
    |||| ||||
    ++++-++++- New bank value, based on last value written to bank select register (mentioned above)

### <span id="Nametable_arrangement_.28.24E000.29"></span>Nametable arrangement (\$E000) {#Nametable_arrangement_.28.24E000.29}

    7  bit  0
    ---- ----
    xxxx xxxN
            |
            +- Nametable arrangement (0: horizontal (A10); 1: vertical (A11))
