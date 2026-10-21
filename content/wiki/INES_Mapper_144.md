---
title: "INES Mapper 144"
url: "/wiki/INES_Mapper_144/"
description: "iNES Mapper 144, allocated for the game Death Race, describes a intentionally defective variant of the Color Dreams board (mapper 11)."
wikiPage: "INES_Mapper_144"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 144, allocated for the game [Death Race](https://nescartdb.com/profile/view/1223/death-race), describes a intentionally defective variant of the [Color Dreams](../Color_Dreams/) board (**mapper 11**).

This game's PCB (labelled **50282**) is almost identical to the [revision B Color Dreams boards](https://nescartdb.com/search/advanced?pcb=UNK-COLORDREAMS-REVB), but a 300Ω resistor was added between CPU D0 and the combination of mapper hardware and ROM. This addition means that *only* the ROM's *least significant bit* always wins bus conflicts.

Nestopia-1.4.0 implements this as `EffectiveData = (WrittenData & ROM[address]) | (ROM[address] & 1)`

FCEUX-2.1.5 works around this by ignoring writes to 0x8000

The most succinct description is `EffectiveData = (ROM[address] & (WrittenData|1))`
