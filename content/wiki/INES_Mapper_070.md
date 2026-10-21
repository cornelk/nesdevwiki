---
title: "INES Mapper 070"
url: "/wiki/INES_Mapper_070/"
description: "The version of this board with 1-screen mapper-controlled mirroring is mapper 152, and the version without a fixed bank at $C000-$FFFF is oversize GxROM. Before Mapper 152 was defined, the version…"
wikiPage: "INES_Mapper_070"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

========================
     =  Mapper 070          =
     ========================

     Example Games:
     --------------------------
     Family Trainer - Manhattan Police
     Family Trainer - Meiro Daisakusen
     Kamen Rider Club
     Space Shadow

     Notes:
     ---------------------------
     Many of these games use the Family Trainer Mat as an input device.

     Registers: **BUS CONFLICTS**
     --------------------------
       $8000-FFFF:  [PPPP CCCC]
         P = Selects 16k PRG @ $8000
         C = Selects 8k CHR @ $0000

     PRG Setup:
     --------------------------

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         |     $8000     |     { -1}     |
         +---------------+---------------+

The version of this board with 1-screen mapper-controlled mirroring is [mapper 152](../INES_Mapper_152/), and the version without a fixed bank at \$C000-\$FFFF is oversize [GxROM](../GxROM/). Before Mapper 152 was defined, the version with 1-screen mirroring had been denoted by setting the "alternative nametables" bit in the iNES header (according to the FWNES98E.TXT file of the fwNES emulator version 0.302b).
