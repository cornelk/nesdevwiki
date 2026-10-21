---
title: "NES 2.0 Mapper 382"
url: "/wiki/NES_2.0_Mapper_382/"
description: "NES 2.0 Mapper 382 denotes the 830928C circuit board, used on a 512 KiB 5-in-1 and a 1 MiB 9-in-1 multicart containing the BNROM game Journey to the West and Capcom/Konami UNROM games."
wikiPage: "NES_2_0_Mapper_382"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 382** denotes the **830928C** circuit board, used on a 512 KiB *5-in-1* and a 1 MiB *9-in-1* multicart containing the [BNROM](../INES_Mapper_034/) game *Journey to the West* and Capcom/Konami [UNROM](../UxROM/) games.

The written address sets the mode and outer bank number, the data latch the inner bank number. The data latch is subject to bus conflicts, making it difficult to write a desired inner bank value as long as the outer bank is not locked.

## <span id="Outer_Bank_register_.28.248000-.24FFFF.2C_write.29"></span>Outer Bank register (\$8000-\$FFFF, write) {#Outer_Bank_register_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210
      -------------------
      1... .... ..LM UPPP
                  || |+++- PRG A19..A17 (outer 128 KiB PRG-ROM bank)
                  || +---- 0: UNROM mode
                  ||       1: BNROM mode
                  |+------ 0: Vertical mirroring
                  |        1: Horizontal mirroring
                  +------- 1: Lock outer bank
    Power-on value: $00

## <span id="Inner_Bank_register_.28.248000-.24FFFF.2C_write.29"></span>Inner Bank register (\$8000-\$FFFF, write) {#Inner_Bank_register_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    D~7654 3210 (bus conflicts)
      ---------
      .... .ABB
            |++- In BNROM mode: PRG A16..A15 (inner 32 KiB PRG-ROM bank)
            +++- In UNROM mode: PRG A16..A14 (inner 16 KiB PRG-ROM bank
                                              at CPU $8000-$BFFF)
