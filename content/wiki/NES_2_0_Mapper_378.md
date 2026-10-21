---
title: "NES 2.0 Mapper 378"
url: "/wiki/NES_2.0_Mapper_378/"
description: "NES 2.0 Mapper 378 denotes a circuit board with unknown ID, used for an 8-in-1 512 KiB multicart containing one AxROM and two UNROM games."
wikiPage: "NES_2_0_Mapper_378"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 378** denotes a circuit board with unknown ID, used for an 8-in-1 512 KiB multicart containing one [AxROM](../AxROM/) and two [UNROM](../UxROM/) games.

## <span id="Outer_Bank_register_.28.248000-.24FFFF.2C_write.29"></span>Outer Bank register (\$8000-\$FFFF, write) {#Outer_Bank_register_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1... .... ..MS cbaa
                  || |+++- PRG A15-A17
                  || |+--- UNROM only: 0=Vertical, 1=Horizontal mirroring
                  || +---- UNROM only: PRG A14 when CPU A14=0
                  |+------ CIRAM A10 in AxROM mode
                  +------- PRG A18, also:
                            0: AxROM mode: PRG A14=CPU A14
                            1: UNROM mode: PRG A14=c if CPU A14=0,
                                           PRG A14..16=7 if CPU A14=1
