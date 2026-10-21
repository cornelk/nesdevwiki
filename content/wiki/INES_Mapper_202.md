---
title: "INES Mapper 202"
url: "/wiki/INES_Mapper_202/"
description: "Implemented in BizHawk as of commit f72875b584abd79201ad9d3d0131fa5afcf0c429"
wikiPage: "INES_Mapper_202"
wikiCategories:
  - name: "Categories:"
    url: "/category/Categories_/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

Implemented in BizHawk as of commit f72875b584abd79201ad9d3d0131fa5afcf0c429

Used for the 150-in-1 pirate cart (which was shipped with dual sloat SP60 famiclone console).

[<img src="../../wiki-images/150in1_top.jpg" width="300" height="132" />](../../wiki-images/150in1_top.jpg)

[](../../wiki-images/150in1_top.jpg)

[150-in-1 PCB top side](../../wiki-images/150in1_top.jpg)

[](../../wiki-images/150in1_top.jpg)

[](../../wiki-images/150in1_top.jpg)

[](../../wiki-images/150in1_top.jpg)[<img src="../../wiki-images/150in1_bottom.jpg" width="300" height="137" />](../../wiki-images/150in1_bottom.jpg)

[](../../wiki-images/150in1_bottom.jpg)

[150-in-1 PCB bottom side](../../wiki-images/150in1_bottom.jpg)

[](../../wiki-images/150in1_bottom.jpg)

[](../../wiki-images/150in1_bottom.jpg)

[](../../wiki-images/150in1_bottom.jpg)[<img src="../../wiki-images/150in1_schematics.png" width="300" height="264" />](../../wiki-images/150in1_schematics.png)

[](../../wiki-images/150in1_schematics.png)

[150-in-1 schematics](../../wiki-images/150in1_schematics.png)

[](../../wiki-images/150in1_schematics.png)

     Registers:
     $8000-FFFF: A~[.... .... .... O..O]
                   [.... .... .... RRRM]

     O = PRG Mode, if 3, will use 32 PRG mode, else 16
     R = Page registers, used for both CHR 8k banks, and PRG banks (in both 16k and 32k modes)
     M = Mirroring, 0 = Vertical, 1 = Horizontal

     CHR Setup:

                        $0000   $0400   $0800   $0C00   $1000   $1400   $1800   $1C00 
                       +-------------------------------+-------------------------------+
     CHR-ROM:          |                             $8000                             |
                       +-------------------------------+-------------------------------+

     PRG Setup:
     ------------------------------
     32k Mode:

     $8000     $A000     $C000     $E000
     +-------------------------------------+
     |               <$8000>               |
     +-------------------------------------+

     16k Mode:

     In 16k mode both regions are mapped to the same PRG bank

     $8000     $A000     $C000     $E000     
     +--------------------------------------+
     |       $8000       |      $8000       |
     +--------------------------------------+
