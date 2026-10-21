---
title: "NES 2.0 Mapper 294"
url: "/wiki/NES_2.0_Mapper_294/"
description: "NES 2.0 Mapper 294 denotes the 63-1601 multicart PCB, used for the 強捧 16-in-1 (1682) multicart consisting of 2 MiB worth of UNROM games."
wikiPage: "NES_2_0_Mapper_294"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 294** denotes the **63-1601** multicart PCB, used for the *強捧 16-in-1 (1682)* multicart consisting of 2 MiB worth of UNROM games.

## <span id="Outer_Bank_Register_.28.244100.2C_write.29"></span>Outer Bank Register (\$4100, write) {#Outer_Bank_Register_.28.244100.2C_write.29}

    Mask: $C100
    D~[...M PPPP]
          | ++++- PRG A20..A17
          +------ Mirroring
                   0=Vertical
                   1=Horizontal
