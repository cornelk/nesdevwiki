---
title: "NES 2.0 Mapper 412"
url: "/wiki/NES_2.0_Mapper_412/"
description: "NES 2.0 Mapper 412 denotes an \"MMC3 plus NROM\" multicart mapper, used for the Intellivision 10-in-1 Plug 'n Play 2nd Edition as well as assorted multicarts from Henggedianzi. It uses an MMC3 clone…"
wikiPage: "NES_2_0_Mapper_412"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 412** denotes an "MMC3 plus NROM" multicart mapper, used for the *Intellivision 10-in-1 Plug 'n Play 2nd Edition* as well as assorted multicarts from Henggedianzi. It uses an [MMC3](../MMC3/) clone with additional outer bank registers at \$6000-\$7FFF that are accessible via the MMC3's WRAM interface, meaning that WRAM must be enabled via \$A001 for them to be written to.

## <span id="8_KiB_CHR_Bank_Register_.28.246000.2C_write.29"></span>8 KiB CHR Bank Register (\$6000, write) {#8_KiB_CHR_Bank_Register_.28.246000.2C_write.29}

    D~[BBBB BB..]  Mask: $E003, reset value: 0
       ++++-++---- CHR A18..A13

Only applicable in NROM banking mode (\$6002.1=1).

## <span id="Mode_Register_.28.246001.2C_write.29"></span>Mode Register (\$6001, write) {#Mode_Register_.28.246001.2C_write.29}

    D~[CARQ DBPL]  Mask: $E003, reset value: 0
       |||| |||+-- Lock register $6001
       |||| ||+--- 0: PRG A18 from MMC3, 1: from B Bit
       |||| |+---- PRG A18 if Q=2
       |||| +----- CHR A18
       |||+------- 0: PRG A17 from MMC3, 1: from A Bit
       ||+-------- 0: CHR A17 from MMC3, 1: from C Bit
       |+--------- PRG A17 if P=1
       +---------- CHR A17 if R=1

## <span id="PRG_Bank_Register_.28.246002.2C_write.29"></span>PRG Bank Register (\$6002, write) {#PRG_Bank_Register_.28.246002.2C_write.29}

    D~[BBBB bQP.]  Mask: $E003, reset value: 0
       |||| ||+--- 0: MMC3 banking, 1: NROM PRG/CHR banking
       |||| |+---- 0: NROM-128, 1: NROM-256 banking, only if P=1
       |||| +----- PRG A14 if P=1 and Q=0
       ++++------- PRG A18..A15 if P=1
