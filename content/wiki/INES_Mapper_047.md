---
title: "INES Mapper 047"
url: "/wiki/INES_Mapper_047/"
wikiPage: "INES_Mapper_047"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

Here are Disch's original notes:  
     ========================
     =  Mapper 047          =
     ========================

     Example Game:
     --------------------------
     Super Spike V'Ball + Nintendo World Cup

     Notes:
     ---------------------------

     Yet another MMC3 multicart.  See mapper 004 for info on MMC3.

     There is no PRG-RAM.  The multicart reg lies at $6000-7FFF, but is only writable when MMC3 PRG-RAM is enabled
     and writable (see $A001)

     Registers:
     ---------------------------

       $6000-7FFF:  [.... ...B]  Block select
       $8000-FFFF:  Same as MMC3 for selected block

     Each block has 128k PRG and 128k CHR.
