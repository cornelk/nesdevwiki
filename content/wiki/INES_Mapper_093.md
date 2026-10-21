---
title: "INES Mapper 093"
url: "/wiki/INES_Mapper_093/"
description: "This is part of a family of Sunsoft mappers used by Japanese games: (iNES Mapper 089, iNES Mapper 093, iNES Mapper 184)"
wikiPage: "INES_Mapper_093"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

This is part of a family of Sunsoft mappers used by Japanese games: ([iNES Mapper 089](../INES_Mapper_089/), iNES Mapper 093, [iNES Mapper 184](../INES_Mapper_184/))

This mapper represents the [Sunsoft-2 IC](../Sunsoft_2_pinout/) on the Sunsoft-3R board. It is used by [Shanghai](https://nescartdb.com/profile/view/3142/shanghai), as well as [Fantasy Zone](https://nescartdb.com/profile/view/3858/fantasy-zone) whose program is compatible with both this board and a different one that used the [Sunsoft-1 IC](../Sunsoft_1_pinout/).

[Mapper 89](../INES_Mapper_089/) uses the same Sunsoft-2 IC but on the Sunsoft-3 board and so differs in that it has mapper-controlled one-screen mirroring and CHR ROM banking. Note that the CHR bank lines are still hooked up to the RAM, explaining the RAM enable bit. For homebrew and hacks, games with CHR ROM should use 89's banking behavior with fixed mirroring.

Despite the mapper existing as a single IC, its functionality [is describable](../Sunsoft_2_pinout/) using just a [74377](../74377/) and a [7432](../7432/) and should probably be considered discrete logic.

Here is the documentation in [disch's original style](../INES_Mapper_DischDocs/):

     Registers **BUS CONFLICTS**:
     --------------------------

       $8000-FFFF:  [.PPP ...E]
         P = PRG Reg  (16k @ $8000)
         E = CHR RAM enable:
             0 = RAM is disabled; writes are ignored and reads are open bus
                 (like iNES Mapper 185 except no games use this)
             1 = normal operation.

     PRG Setup:
     --------------------------

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         |     $8000     |     { -1}     |
         +---------------+---------------+
