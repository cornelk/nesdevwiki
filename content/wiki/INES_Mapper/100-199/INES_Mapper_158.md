---
title: "INES Mapper 158"
url: "/wiki/INES_Mapper_158/"
description: "iNES Mapper 158 represents Tengen's board 800037, used for Alien Syndrome."
wikiPage: "INES_Mapper_158"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 158 represents Tengen's board **800037**, used for [Alien Syndrome](https://nescartdb.com/profile/view/325/alien-syndrome).

It modifies [RAMBO-1](../RAMBO-1/) (**mapper 64**) in the exact same way that [TLSROM](../TLSROM/) ([118](../INES_Mapper_118/)) differs from [MMC3](../MMC3/) ([4](../MMC3/)), by which we mean the RAMBO-1's CHR A17 output is connected to CIRAM A10.

Nestopia-1.4.0's implementation looks incomplete; it seems to ignore [bank registers 8 and 9, as well as the K bit](../RAMBO-1/#Bank_select_.28.248000-.249FFE.2C_even.29).
