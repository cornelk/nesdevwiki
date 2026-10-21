---
title: "INES Mapper 199"
url: "/wiki/INES_Mapper_199/"
description: "iNES Mapper 199 is used for the Waixing FS309 circuit board, a TNROM-compatible board with 32 KiB of battery-backed PRG RAM, of which 12 KiB are mapped into CPU address space from $5000-$7FFF. Used…"
wikiPage: "INES_Mapper_199"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**iNES Mapper 199** is used for the Waixing **FS309** circuit board, a [TNROM](../TxROM/)-compatible board with 32 KiB of battery-backed PRG RAM, of which 12 KiB are mapped into CPU address space from \$5000-\$7FFF. Used by the following games:

- 成吉思汗 (Chéngjísīhán), Waixing's Chinese translation of Koei's *Genghis Khan*
- 風色幻想 (Fēngsè Huànxiǎng) / 湯姆歷險記中文版 (Tāngmǔ Lìxiǎn jì Zhōngwén bǎn) / 汤姆寻宝记 (Tāngmǔ Xúnbǎo jì); Waixing's Chinese translation of *Square no Tom Sawyer*
- 吞食天地: 三国外传 (Tūnshí Tiāndì - Sānguó Wàizhuàn): Waixing's re-release of the Chinese translation of Capcom's 天地を喰らう II: 諸葛孔明伝 (Destiny of an Emperor II). Unlike the [original 1994 Xiangfeng cartoon release](../INES_Mapper_198/), this release uses one conventional 1 MiB of PRG-ROM.

The MMC3 clone used by this board supports more than 512 KiB PRG-ROM, necessary for that last game in the list. The 8 KiB of CHR RAM are unbanked, and games do not initialize the MMC3's CHR bank registers to a contiguous 8 KiB block!

Some Chinese emulators furthermore treat mapper 199 as a duplicate of [INES Mapper 176](../INES_Mapper_176/). Such games can be detected by them writing \$02 to \$5FF3 to enable Mapper 176's Extended MMC3 Mode and should be reassigned to 176.
