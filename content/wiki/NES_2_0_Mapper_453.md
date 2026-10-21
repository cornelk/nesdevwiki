---
title: "NES 2.0 Mapper 453"
url: "/wiki/NES_2.0_Mapper_453/"
description: "NES 2.0 Mapper 453 denotes the Realtec 8042 multicart circuit board, used for the 3-in-1 1995 World Competitive Champion Edition (DG-002) multicart. It supports both AOROM and UNROM banking modes."
wikiPage: "NES_2_0_Mapper_453"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 453** denotes the **Realtec 8042** multicart circuit board, used for the *3-in-1 1995 World Competitive Champion Edition (DG-002)* multicart. It supports both AOROM and UNROM banking modes.

## <span id="Data_Latch_.28.248000-.24FFFF.29.2C_write"></span>Data Latch (\$8000-\$FFFF), write {#Data_Latch_.28.248000-.24FFFF.29.2C_write}

    D~[.PpM .BBB]
        |||  +++- PRG A16..A14 in UNROM mode,
        |||       PRG A17..A15 in AOROM mode
        ||+------ Mirroring
        ||         0=V, 1=H in UNROM mode,
        ||         CIRAM A10 in AOROM mode
        ++------- Lock Pp if either bit is set
        |+------- PRG A17 in UNROM mode
        +-------- PRG A18
        +-------- PRG banking mode
                   0=UNROM, 1=AOROM
