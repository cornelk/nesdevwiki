---
title: "INES Mapper 192"
url: "/wiki/INES_Mapper_192/"
description: "iNES Mapper 192 is used to designate the Waixing FS308 circuit board, which is similar in capability to TQROM (iNES Mapper 119). It uses a clone of the MMC3 but uses the CHR bank number in such a way…"
wikiPage: "INES_Mapper_192"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
---

**iNES Mapper 192** is used to designate the Waixing **FS308** circuit board, which is similar in capability to TQROM ([iNES Mapper 119](../INES_Mapper_119/)). It uses a clone of the [MMC3](../MMC3/) but uses the CHR bank number in such a way so as to allow [both CHR ROM and CHR RAM to be used simultaneously](../MMC3_with_CHR_ROM_and_CHR_RAM/) for displaying Chinese characters.

BizHawk supports this mapper as of r3973; FCEUX and FCEU-MM have supported them for a long time but definitely as of 26-nov-2012, although CRC checks for m192\<-\>m074 may be lacking.

## Registers {#Registers}

Registers resemble [MMC3](../MMC3/), except CHR ROM banks 8, 9, 10, and 11 are replaced with CHR RAM. Other (unknown) inaccuracies compared to an authentic MMC3 may be present. This is effectively the same as [INES Mapper 074](../INES_Mapper_074/) but with twice as much CHR RAM. If the [NES 2.0](../NES_2.0/) header is present, 74 and 192 should be emulated identically, as NES 2.0 specifies CHR RAM size elsewhere.

     Example Game:
     --------------------------
     Ying Lie Qun Xia Zhuan

     Notes:
     --------------------------
     This mapper is a modified MMC3 (or is based on MMC3?).

     In addition to any CHR-ROM present, there is also an additional 4k of CHR-RAM which is selectable.

     CHR Pages $08-$0B are CHR-RAM, other pages are CHR-ROM.

Mappers [252](../INES_Mapper_252/) and [253](../INES_Mapper_253/), also used for Waixing's localizations, are similar but use a VRC4.

## See also {#See_also}

- [MMC3 with CHR ROM and CHR RAM](../MMC3_with_CHR_ROM_and_CHR_RAM/)
