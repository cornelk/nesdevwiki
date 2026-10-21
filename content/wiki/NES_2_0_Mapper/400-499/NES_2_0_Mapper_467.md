---
title: "NES 2.0 Mapper 467"
url: "/wiki/NES_2.0_Mapper_467/"
description: "NES 2.0 Mapper 467 denotes the 47-2 MMC3-based circuit board, used for a 72-in-1 multicart. It is effectively INES Mapper 197 submapper 0 with an outer bank register and an NROM mode."
wikiPage: "NES_2_0_Mapper_467"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 467** denotes the **47-2** [MMC3](../MMC3/)-based circuit board, used for a *72-in-1* multicart. It is effectively [INES Mapper 197](../INES_Mapper_197/) submapper 0 with an outer bank register and an NROM mode.

## <span id="Outer_Bank_Register_.28.249000.2C_write.29"></span>Outer Bank Register (\$9000, write) {#Outer_Bank_Register_.28.249000.2C_write.29}

    Mask: $F000?

    D~[MSNQ QPPp]
       ||++-++++- PRG A18..A14
       ||+------- PRG banking mode
       ||          0=NROM-128
       ||          1=MMC3-based
       |+-------- PRG (if N=1)/CHR banking mode
       |           0=PRG A14..A13 from MMC3, PRG A19..A15 from NQQPP
       |             CHR A12..A10=PPU A12..A10, CHR A18..A13 from MMC3 CHR A17..A12, CHR A19=0
       |           1=PRG A16..A13 from MMC3, PRG A19..A17 from NQQ
       |             CHR A10=PPU A10, CHR A18..A11 from MMC3 CHR A17..A10, CHR A19=1
       +--------- Nametable mirroring
                   0=Vertical
                   1=Horizontal

The MMC3 clone ignores the PRG/CHR inversion bits D6 and D7 in the pointer register \$8000, and its mirroring output is ignored in favor of the Outer Bank Register's.
