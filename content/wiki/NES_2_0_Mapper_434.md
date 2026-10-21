---
title: "NES 2.0 Mapper 434"
url: "/wiki/NES_2.0_Mapper_434/"
description: "NES 2.0 Mapper 434 denotes the S-009 PCB, used by multicarts consisting of several UNROM games."
wikiPage: "NES_2_0_Mapper_434"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 434** denotes the **S-009** PCB, used by multicarts consisting of several UNROM games.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.29"></span>Outer Bank Register (\$6000-\$7FFF) {#Outer_Bank_Register_.28.246000-.247FFF.29}

    Mask: $E000

    D~7654 3210
      ---------
      ..M. .PPP
        |   +++- PRG A19..A17
        +------- 0: Horizontal mirroring
                 1: Vertical mirroring

## <span id="Inner_Bank_Register_.28.248000-.24FFFF.29"></span>Inner Bank Register (\$8000-\$FFFF) {#Inner_Bank_Register_.28.248000-.24FFFF.29}

    Mask: $8000

    D~7654 3210 (bus conflicts)
      ---------
      .... .PPP
            +++- PRG A16..A14 if CPU A14=0
