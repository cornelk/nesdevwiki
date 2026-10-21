---
title: "INES Mapper 140"
url: "/wiki/INES_Mapper_140/"
description: "Mapper 140 represents Jaleco's JF-11 and JF-14 boards, which act much like GNROM (66) except that the writable port is moved down to $6000-$7FFF."
wikiPage: "INES_Mapper_140"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

Mapper 140 represents Jaleco's [JF-11](https://nescartdb.com/search/advanced?pcb=JF-11) and [JF-14](https://nescartdb.com/search/advanced?pcb=JF-14) boards, which act much like [GNROM](../GxROM/) (**66**) except that the writable port is moved down to \$6000-\$7FFF.

     Here are Disch's original notes:  
     ========================
     =  Mapper 140          =
     ========================

     Example Game:
     --------------------------
     Bio Senshi Dan - Increaser Tono Tatakai

     Notes:
     ---------------------------
     Regs lie at $6000-7FFF, so there's no SRAM

     Registers:
     --------------------------
       $6000-7FFF:  [..PP CCCC]
         P = Selects 32k PRG @ $8000
         C = Selects 8k CHR @ $0000
