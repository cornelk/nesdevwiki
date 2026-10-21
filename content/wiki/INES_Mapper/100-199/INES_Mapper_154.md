---
title: "INES Mapper 154"
url: "/wiki/INES_Mapper_154/"
description: "iNES Mapper 154 represents NAMCOT-3453, a board used only for the game Devil Man."
wikiPage: "INES_Mapper_154"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 154 represents **NAMCOT-3453**, a board used only for the game [Devil Man](https://nescartdb.com/profile/view/3837/devil-man).

It is identical to [Mapper 88](../INES_Mapper_088/), but with the addition of a single bit allowing for one-screen nametable control:

     $8000-$9FFF: [..xx xxxx]: See mapper 206 documentation
     $8000-$FFFF: [.N.. ....]
       N = Nametables
         0 = 1ScA
         1 = 1ScB

Note that the nametable control bit is present over the entire 32kB range; unlike the associated Namco 108 which is only present at \$8000-\$9FFF.

Like mapper 88, if the CHR ROM is larger than 65536 bytes, CHR ROM A16 is connected to PPU A12. This means content in the left pattern table (\$0xxx) comes from the first 64kB of CHR-ROM and content in the right pattern table (\$1xxx) comes from the last 64kB of CHR-ROM.
