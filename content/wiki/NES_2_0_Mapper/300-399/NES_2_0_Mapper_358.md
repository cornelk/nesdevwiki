---
title: "NES 2.0 Mapper 358"
url: "/wiki/NES_2.0_Mapper_358/"
description: "NES 2.0 Mapper 358 denotes a variation of NES 2.0 Mapper 282 with partially overlapping PRG/CHR outer banks:"
wikiPage: "NES_2_0_Mapper_358"
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

NES 2.0 Mapper 358 denotes a variation of [NES 2.0 Mapper 282](../NES_2.0_Mapper_282/) with partially overlapping PRG/CHR outer banks:

- *1997 Super Game 7-in-1 (JY-016)* (PCB YY860606C, 512 kiB PRG-ROM)
- *1998 Super HiK 5-in-1 (JY-113)* (PCB YY860606C, 512 kiB PRG-ROM)
- *1998 Super HiK 6-in-1 (JY-117)* (unmarked PCB, 1024 KiB PRG-ROM)

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F803

    7654 3210
    ---- ----
    ..M. .ppC
      |   ||+-- CHR-ROM A18 if M=0, ignored if M=1
      |   ++--- PRG-ROM A18-A19
      |   +---- CHR-ROM A19
      +-------- Outer CHR bank size and selection
                 0: Mask to 256 KiB, use C
                 1: Mask to 512 KiB, ignore C

PRG (\$8000-\$8FFF) bank select register bits that select 256 KiB banks are masked off. Dumping the CHR-ROM does not work reliably in 1 KiB CHR-ROM banking mode.

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
