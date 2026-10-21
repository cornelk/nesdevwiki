---
title: "INES Mapper 231"
url: "/wiki/INES_Mapper_231/"
wikiPage: "INES_Mapper_231"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

Here are Disch's original notes:  
     ========================
     =  Mapper 231          =
     ========================

     Example Game:
     --------------------------
     20-in-1

     Registers:
     ---------------------------

      $8000-FFFF:     A~[.... .... M.LP PPP.]
         M = Mirroring (0=Vert, 1=Horz)
         L = Low bit of PRG
         P = High bits of PRG

     PRG Setup:
     ---------------------------

     Note that 'L' and 'P' bits make up the PRG reg, and the 'L' is the low bit.

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         | $8000 AND $1E |     $8000     |
         +---------------+---------------+
