---
title: "NES 2.0 Mapper 520"
url: "/wiki/NES_2.0_Mapper_520/"
description: "NES 2.0 Mapper 520 is used for the 2-in-1 Datach Dragon Ball Z/Datach Yu Yu Hakusho multicart. It uses a VRC4 clone (in the VRC4e configuration, i.e. address 0x004s to A0 and address 0x008s to A1)…"
wikiPage: "NES_2_0_Mapper_520"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 520 is used for the *2-in-1 Datach Dragon Ball Z/Datach Yu Yu Hakusho* multicart. It uses a [VRC4](../VRC2_and_VRC4/) clone (in the VRC4e configuration, i.e. address 0x004s to A0 and address 0x008s to A1) using 8 KiB of banked CHR-RAM rather than CHR-ROM:

### <span id="CHR_Select_0_low_.28.24B000.29"></span>CHR Select 0 low (\$B000) {#CHR_Select_0_low_.28.24B000.29}

    7  bit  0
    ---------
    .... OLLL
         |+++- Select 1 KiB CHR-RAM bank at PPU $0000-$03FF
         +---- Select outer 256 KiB PRG-ROM bank when PPU $0000-$03FF is accessed

- The high CHR Select register (\$B004) is ignored.
- The other seven low CHR bank select registers continue the pattern for the other seven 1 KiB PPU banks.

## Notes {#Notes}

- The ROM image in GoodNES 3.23b is set to Mapper 23, even as the outer bank register is not part of that mapper's specification.
- As with [SOROM, SUROM and SXROM](../MMC1/#SOROM.2C_SUROM_and_SXROM), if the several CHR bank registers do not specify the same outer PRG ROM bank, then PRG ROM presumably would be bankswitched as the PPU renders.
