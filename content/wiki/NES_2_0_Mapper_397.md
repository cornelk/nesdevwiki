---
title: "NES 2.0 Mapper 397"
url: "/wiki/NES_2.0_Mapper_397/"
description: "NES 2.0 Mapper 295 is used for a single multicart from J.Y. Company:"
wikiPage: "NES_2_0_Mapper_397"
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

**NES 2.0 Mapper 295** is used for a single multicart from J.Y. Company:

- *1996 Soccer 6-in-1* (JY-082)

The same bits in register \$D003 select both PRG and CHR outer banks. It is almost identical to [NES 2.0 Mapper 295](../NES_2.0_Mapper_295/), except that PRG A17 selected via \$8000-\$8FFF is not masked off.

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F003

    7654 3210
    ---- ----
    .... .BBC
          +++- Select 128 KiB outer CHR-ROM bank (CHR A19..17)
          ++-- Select 256 KiB outer PRG-ROM bank (PRG A19..18)

PRG (\$8000-\$8FFF) and CHR (\$9000-\$AFFF) bank select register bits that select 256/128 KiB banks are masked off.

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
