---
title: "NES 2.0 Mapper 387"
url: "/wiki/NES_2.0_Mapper_387/"
description: "NES 2.0 Mapper 387 denotes J.Y. Company's YY850735C and YY850817C circuit boards, used for the multicarts:"
wikiPage: "NES_2_0_Mapper_387"
wikiCategories:
  - name: "Mappers using J.Y. Company ASIC"
    url: "/category/Mappers_using_J.Y._Company_ASIC/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 387** denotes J.Y. Company's **YY850735C** and **YY850817C** circuit boards, used for the multicarts:

- *1996 Super HiK 4-in-1 (JY-087)*
- *1996 Super HIK 4-in-1 (JY-089)*
- *1996 Super HIK 4-in-1 (JY-093)*
- *1996 Super HIK 4-in-1 (JY-094)*

It uses the [J.Y. Company ASIC](../J.Y._Company_ASIC/), has 512 MiB of PRG-ROM and 1 MiB of CHR-ROM. [NES 2.0 Mapper 386](../NES_2.0_Mapper_386/) is almost the same but changes the outer PRG bank size from 128 to 256 KiB.

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F003

    7654 3210
    ---- ----
    ..L? PCpc
      || |||+- CHR A18, ignored if L=1
      || ||+-- PRG A17
      || |+--- CHR A19
      || +---- PRG A18
      |+------ Always 1
      +------- Select outer CHR-ROM bank size
                0: Mask $900x/$A00x to 256 KiB, use c
                1: Mask $900x/$A00x to 512 KiB, ignore c   

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
