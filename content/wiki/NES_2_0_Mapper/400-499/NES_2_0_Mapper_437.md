---
title: "NES 2.0 Mapper 437"
url: "/wiki/NES_2.0_Mapper_437/"
description: "NES 2.0 Mapper 437 denotes the NTDEC TH2348 multicart circuit board for eight UNROM games with an additional outer bank register."
wikiPage: "NES_2_0_Mapper_437"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 437** denotes the NTDEC **TH2348** multicart circuit board for eight [UNROM](../UxROM/) games with an additional outer bank register.

## <span id="Outer_Bank_register_.28.245000-.245FFF.2C_write.29"></span>Outer Bank register (\$5000-\$5FFF, write) {#Outer_Bank_register_.28.245000-.245FFF.2C_write.29}

    A~[0101 .... .... MOOO]
                      |+++- PRG A19..A17
                      +---- Nametable mirroring
                            0: CIRAM A10=PPU A10 (vertical)
                            1: CIRAM A10=PPU A11 (horizontal)
