---
title: "INES Mapper 232"
url: "/wiki/INES_Mapper_232/"
wikiPage: "INES_Mapper_232"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

Here are Disch's original notes:  
     ========================
     =  Mapper 232          =
     ========================

     Example Games:
     --------------------------
     Quattro Adventure
     Quattro Sports
     Quattro Arcade

     Notes:
     --------------------------
     This is another Camerica/Codemasters mapper like 071.  Like 071, this mapper also involves a custom lockout
     defeat circuit which is mostly unimportant for emulation purposes.  Details will not be mentioned here, but
     are outlined in Kevtris' Camerica Mappers documentation.

     Registers:
     ---------------------------

      $8000-BFFF:   [...B B...]   PRG Block Select

      $C000-FFFF:   [.... ..PP]   PRG Page Select

     PRG Setup:
     ---------------------------

     Pages are taken from the 64k block currently selected by $8000.

           $8000   $A000   $C000   $E000  
         +---------------+---------------+
         |     $C000     |     { 3 }     |
         +---------------+---------------+
