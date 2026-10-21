---
title: "INES Mapper 213"
url: "/wiki/INES_Mapper_213/"
description: "iNES Mapper 213 is a duplicate of INES Mapper 058, assigned to by:"
wikiPage: "INES_Mapper_213"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 213** is a duplicate of [INES Mapper 058](../INES_Mapper_058/), assigned to by:

- *9999999-in-1* (*Can You Feel the Love Tonight* menu music)
- *168-in-1* (*My Way* menu music)

[Pre-2018 versions of FCEUX](https://sourceforge.net/p/fceumm/code/186/tree//src/boards/addrlatch.c) as well as [current versions of Mesen](https://github.com/SourMesen/Mesen/blob/master/Core/Mapper213.h) emulate only the NROM-256 setting with hard-wired mirroring, while [current versions of FCEUX](https://github.com/TASVideos/fceux/blob/master/src/boards/addrlatch.cpp) implement a duplicate of mapper 58 with some logic for determining nametable mirroring that only works for one ROM (said *168-in-1*). Both ROM files run well as mapper 58.
