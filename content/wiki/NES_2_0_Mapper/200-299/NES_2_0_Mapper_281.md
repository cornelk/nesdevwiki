---
title: "NES 2.0 Mapper 281"
url: "/wiki/NES_2.0_Mapper_281/"
description: "NES 2.0 Mapper 281 denotes a number of compatible circuit boards from J.Y. Company:"
wikiPage: "NES_2_0_Mapper_281"
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

**NES 2.0 Mapper 281** denotes a number of compatible circuit boards from J.Y. Company:

- 512 KiB PRG-ROM, 512 KiB CHR-ROM: **YY841144C**, **YY850771C**, **YY851115C**, **YY860416C**, **YY860417C**, **YY860731C**
- 1 MiB PRG-ROM, 1 MiB CHR-ROM: **YY850719C**

and is used for the following multicarts:

- *1996 Super HiK 4-in-1 (JY-052)*
- *1997 Super HiK 4-in-1 (JY-052)*
- *1996 Super HiK 4-in-1 (JY-053)*
- *1996 Super HiK 4-in-1 (JY-054)*
- *1997 Super HiK 4-in-1 (JY-054)*
- *1996 Super HiK 4-in-1 (JY-055)*
- *1996 Power Rangers HiK 4-in-1 (JY-066)*
- *1997 Power Rangers HiK 4-in-1 (JY-066)*
- *1996 Super HiK 3-in-1 (JY-068)*
- *1996 Super HiK 3-in-1 (JY-080)*
- *1996 Super HiK 5-in-1 (JY-088)*

It uses the [J.Y. Company ASIC](../J.Y._Company_ASIC/).

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

Mask: \$F003

    7654 3210
    ---- ----
    .... ..BB
           ++- PRG/CHR A18/A19 (256 KiB outer PRG/CHR-ROM bank)

PRG (\$8000-\$8FFF) and CHR (\$9000-\$AFFF) bank select register bits that select 256 KiB banks are masked off.

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
