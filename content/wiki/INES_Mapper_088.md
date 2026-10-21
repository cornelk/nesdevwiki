---
title: "INES Mapper 088"
url: "/wiki/INES_Mapper_088/"
description: "iNES Mapper 088 is the same as iNES Mapper 206 with the following exception:"
wikiPage: "INES_Mapper_088"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 088 is the same as [iNES Mapper 206](../INES_Mapper_206/) with the following exception:

CHR support is increased to 128 KiB by connecting PPU's A12 line to the CHR ROM's A16 line. Consequently, \$0000-\$0FFF can only address CHR from the first 64 KiB, \$1000-\$1FFF can only address CHR from the second 64 KiB half of CHR-ROM.

A possible way to implement this would be to mask the 1 KiB CHR-ROM bank output from the mapper by ANDing with \$3F, and then OR it with \$40 for Namco 108 registers 2, 3, 4, and 5.

If the CHR ROM is 64 KiB or smaller, it is identical to mapper 206.

There are three games known to use this mapper:

- [Quinty (J)](https://nescartdb.com/profile/view/1576/quinty)
- [Namcot Mahjong 3 - Mahjong Tengoku](https://nescartdb.com/profile/view/1828/namco-mahjong-iii-mahjong-tengoku)
- [Dragon Spirit - Aratanaru Densetsu](https://nescartdb.com/profile/view/2307/dragon-spirit-aratanaru-densetsu)

## See also {#See_also}

- Unlicensed mappers [iNES Mapper 121](../INES_Mapper_121/) and [iNES Mapper 187](../INES_Mapper_187/) uses a similar setup to address larger CHR-ROM, but using a full-fledged MMC3 clone, and can address 512 KiB of CHR as a result.
