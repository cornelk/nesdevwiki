---
title: "NES 2.0 Mapper 295"
url: "/wiki/NES_2.0_Mapper_295/"
description: "NES 2.0 Mapper 295 is used for several multicarts from J.Y. Company:"
wikiPage: "NES_2_0_Mapper_295"
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

NES 2.0 Mapper 295 is used for several multicarts from J.Y. Company:

- *1996 Soccer 7-in-1 (JY-014B)*
- *1997 Super HIK 4-in-1 (JY-095)*
- *1997 Super HIK 7-in-1 (JY-096)*
- *1997 Super HIK 4-in-1 (JY-099)*
- *1997 Super 9-in-1 (JY-109)*
- *1997 Super 13-in-1 (JY-110)*

Its UNIF board name is **BMC-13in1JY110**, although FCEUX' support is non-working. It uses the [J.Y. Company ASIC](../J.Y._Company_ASIC/).

The same bits in register \$D003 select both PRG and CHR outer banks. It is almost identical to [NES 2.0 Mapper 281](../NES_2.0_Mapper_281/), except that the outer bank size is 128 KiB rather than 256 KiB.

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F003

    7654 3210
    ---- ----
    .... .CCC
          +++- Select 128 KiB outer PRG-ROM bank
          +++- Select 128 KiB outer CHR-ROM bank

PRG (\$8000-\$8FFF) and CHR (\$9000-\$AFFF) bank select register bits that select 128/256 KiB banks are masked off.

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).
