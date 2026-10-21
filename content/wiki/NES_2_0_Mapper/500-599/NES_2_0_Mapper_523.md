---
title: "NES 2.0 Mapper 523"
url: "/wiki/NES_2.0_Mapper_523/"
description: "NES 2.0 Mapper 523 is used for the 晶科泰 (Jncota) re-release of the game 封神榜: 伏魔三太子 (Fēngshénbǎng: Fúmó Sān Tàizǐ), originally released by C\u0026E on INES Mapper 246. It is…"
wikiPage: "NES_2_0_Mapper_523"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

NES 2.0 Mapper 523 is used for the 晶科泰 (Jncota) re-release of the game 封神榜: 伏魔三太子 (Fēngshénbǎng: Fúmó Sān Tàizǐ), originally released by C&E on [INES Mapper 246](../INES_Mapper_246/). It is basically identical to the [Waixing FS005](../INES_Mapper_176/) board, except that ...

- Mirroring is hard-wired (to Horizontal in Fēngshénbǎng's case),
- CHR-ROM A10 is directly connected to PPU A10,
- CHR-ROM A11-A18 are connected to the MMC3 clone's CHR A10-A17 outputs,
- the MMC3 clone's PPU A10 input is connected to Ground.

With the MMC3 clone set to [Extended MMC3 Mode](../INES_Mapper_176/#MMC3-compatible_registers_.28.248000.2F.248001.2C_.24A000.2C_.24C000.2F.24C001.2C_.24E000.2F.24E001.29), the bank registers at \$8000 therefore change from their normal to the following meanings:

    $0: Select 2 KiB CHR bank at PPU $0000-$07FF (or $1000-$17FF)
    $1: Select 2 KiB CHR bank at PPU $0800-$0FFF (or $1800-$1FFF)
    $2: Select 2 KiB CHR bank at PPU $1000-$17FF (or $0000-$07FF)
    $3: No function
    $4: Select 2 KiB CHR bank at PPU $1800-$1FFF (or $0800-$0FFF)
    $5: No function
    $6: Select 8 KiB PRG ROM bank at $8000-$9FFF (or $C000-$DFFF)
    $7: Select 8 KiB PRG ROM bank at $A000-$BFFF
    $8: Select 8 KiB PRG ROM bank at $C000-$DFFF (or $8000-$9FFF)
    $9: Select 8 KiB PRG ROM bank at $E000-$FFFF
    $A: No function
    $B: No function

## See Also {#See_Also}

- The licensed mapper [iNES Mapper 076](../INES_Mapper_076/) addresses larger CHR-ROM in a similar way by rewiring the mapper IC to address 2KiB CHR bank.
