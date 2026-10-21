---
title: "INES Mapper 089"
url: "/wiki/INES_Mapper_089/"
description: "This is part of a family of Sunsoft mappers used largely by Japanese games: (iNES Mapper 089, iNES Mapper 093, iNES Mapper 184)"
wikiPage: "INES_Mapper_089"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

This is part of a family of Sunsoft mappers used largely by Japanese games: (iNES Mapper 089, [iNES Mapper 093](../INES_Mapper_093/), [iNES Mapper 184](../INES_Mapper_184/))

This mapper is used by Tenka no Goikenban: Mito Koumon (J). This game uses the [Sunsoft-2](../Sunsoft_2_pinout/) chip on the Sunsoft-3 board, and so works differently than [Mapper 93](../INES_Mapper_093/)

Here is the documentation in [disch's original style](../INES_Mapper_DischDocs/):

     Registers **BUS CONFLICTS**:
     --------------------------

       $8000-FFFF:  [CPPP MCCC]
         C = Select 8k CHR @ $0000
         P = Select 16k PRG @ $8000
         M = Mirroring:
           0 = 1ScA
           1 = 1ScB

     PRG Setup:
     --------------------------

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         |     $8000     |     { -1}     |
         +---------------+---------------+
