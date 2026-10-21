---
title: "INES Mapper 240"
url: "/wiki/INES_Mapper_240/"
description: "This is two similar PCBs with subtly different addresses for the register."
wikiPage: "INES_Mapper_240"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
---

Here are Disch's original notes:  
     ========================
     =  Mapper 240          =
     ========================

     Example Games:
     --------------------------
     Jing Ke Xin Zhuan
     Sheng Huo Lie Zhuan

     Registers:
     --------------------------
       $4020-5FFF:  [PPPP CCCC]
         P = Selects 32k PRG @ $8000
         C = Selects 8k CHR @ $0000

This is two similar PCBs with subtly different addresses for the register.

Jing Ke Xin Zhuan (荊軻新傳) decodes the register as: (ADDRESS & \$E800) == \$4800

Sheng Huo Lie Zhuan (聖火列傳) decodes the register as: (ADDRESS & \$E100) == \$4100

Both boards add PRG RAM at the usual \$6000
