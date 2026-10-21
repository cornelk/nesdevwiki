---
title: "INES Mapper 152"
url: "/wiki/INES_Mapper_152/"
description: "The version of this board without mirroring control is mapper 70. One variant of mapper 78 is the same thing with the nibbles swapped."
wikiPage: "INES_Mapper_152"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
---

========================
     =  Mapper 152          =
     ========================

     Example Games:
     --------------------------
     Arkanoid 2 (J)
     Gegege no Kitarou 2

     Registers: (** BUS CONFLICTS **)
     --------------------------

       $8000-FFFF:  [MPPP CCCC]
         M = Mirroring:
             0 = 1ScA
             1 = 1ScB

         P = PRG Reg (16k @ $8000)
         C = CHR Reg (8k @ $0000)

     PRG Setup:
     --------------------------

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         |     $8000     |     { -1}     |
         +---------------+---------------+

The version of this board without mirroring control is [mapper 70](../INES_Mapper_070/). One variant of [mapper 78](../INES_Mapper_078/) is the same thing with the nibbles swapped.
