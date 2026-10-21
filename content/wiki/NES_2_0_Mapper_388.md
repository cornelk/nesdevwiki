---
title: "NES 2.0 Mapper 388"
url: "/wiki/NES_2.0_Mapper_388/"
description: "NES 2.0 Mapper 388 denotes J.Y. Company's YY841145C and YY850835C circuit boards, used for the multicarts:"
wikiPage: "NES_2_0_Mapper_388"
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

**NES 2.0 Mapper 388** denotes J.Y. Company's **YY841145C** and **YY850835C** circuit boards, used for the multicarts:

- *1996 Super HiK 4-in-1 (JY-056 earlier version)*
- *1996 Super HiK 5-in-1 (JY-056 later version)*

It uses the [J.Y. Company ASIC](../J.Y._Company_ASIC/), has 1 MiB of PRG-ROM and CHR-ROM each, and is similar to [NES 2.0 Mapper 386](../NES_2.0_Mapper_386/) but has the Outer Bank register bits in a different order. This particular PCB **inhibits** ROM nametables and Extended Mirroring that would otherwise be enabled via register, similar to [INES Mapper 090](../J.Y._Company_ASIC/).

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F003

    7654 3210
    ---- ----
    ..L. PpCc
      |  |||+- CHR A18, ignored if L=1
      |  ||+-- CHR A19
      |  |+--- PRG A18
      |  +---- PRG A19
      +------- Select outer CHR-ROM bank size
                0: Mask $900x/$A00x to 256 KiB, use c
                1: Mask $900x/$A00x to 512 KiB, ignore c   

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
