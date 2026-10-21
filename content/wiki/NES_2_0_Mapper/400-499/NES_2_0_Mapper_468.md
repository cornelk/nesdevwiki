---
title: "NES 2.0 Mapper 468"
url: "/wiki/NES_2.0_Mapper_468/"
description: "NES 2.0 Mapper 468 denotes BlazePro multicarts based on the Altera MAX II CPLD. It can emulate sixteen mappers at once, selected via the MSB nibble at $5700. Different multicarts program the CPLD…"
wikiPage: "NES_2_0_Mapper_468"
wikiCategories:
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 468** denotes **BlazePro** multicarts based on the Altera MAX II CPLD. It can emulate sixteen mappers at once, selected via the MSB nibble at \$5700. Different multicarts program the CPLD with different sets of mappers, denoted via the NES 2.0 submapper field. The LSB nibble of \$5700 contains mapper-specific flags.

| Submapper | \$5700.4-7 | Mapper |
|----|----|----|
| 0 | 0 | [SxROM](../MMC1/) |
| 3 | 2 | [SxROM](../MMC1/) |
| 0 | 1 | [SUROM](../MMC1/) |
| 0 | 4 | [ANROM](../AxROM/)/[BNROM](../INES_Mapper_034/) |
| 1 | 4 | [ANROM](../AxROM/)/[BNROM](../INES_Mapper_034/) |
| 0 | 6 | [AOROM](../AxROM/) |
| 1 | 6 | [AOROM](../AxROM/) |
| 0 | 10 | [PNROM](../MMC2/) |
| 1 | 0 | [TxROM](../TxROM/), 256 KiB CHR |
| 1 | 1 | [TxROM](../TxROM/), 128 KiB CHR |
| 1 | 2 | [TxSROM](../INES_Mapper_118/) |
| 0 | 8 | [FxROM](../FxROM/) |
| 0 | 9 | [UNROM](../UxROM/), [IF-12 (Holy Diver)](../INES_Mapper_078/) |
| 0 | 11 | [UOROM](../UxROM/) |
| 1 | 7 | [UxROM](../UxROM/) |
| 3 | 7 | [UxROM](../UxROM/) |
| 0 | 5 | (C)NROM, [BF9097](../INES_Mapper_071/#Mirroring_.28.248000-.249FFF.29) |
| 1 | 5 | (C)NROM, [BF9097](../INES_Mapper_071/#Mirroring_.28.248000-.249FFF.29) |
| 0 | 12 | [GNROM](../GxROM/) |
| 0 | 13 | [Color Dreams](../Color_Dreams/) |
| 1 | 12 | [GNROM](../GxROM/) |
| 1 | 13 | [Color Dreams](../Color_Dreams/) |
| 2 | 0 | [VRC2](../VRC2_and_VRC4/) A0/A1 |
| 2 | 1 | [VRC4](../VRC2_and_VRC4/) A0/A1 |
| 2 | 2 | [VRC2](../VRC2_and_VRC4/) A1/A0 |
| 2 | 3 | [VRC4](../VRC2_and_VRC4/) A1/A0 |
| 3 | 0 | [VRC6](../VRC6/), 256 KiB CHR |
| 3 | 1 | [VRC6](../VRC6/), 128 KiB CHR |
| 4 | 0 | [VRC1](../VRC1/) |
| 4 | 1 | [VRC7](../VRC7/) |
| 4 | 4 | [VRC3](../VRC3/) |
| 0 | 14 | [Nanjing](../INES_Mapper_163/) |
| 1 | 14 | [Nanjing](../INES_Mapper_163/) |
| 0 | 7 | [SMB2J](../INES_Mapper_040/) |
| 5 | 0 | [FME-7](../Sunsoft_FME-7/) |
