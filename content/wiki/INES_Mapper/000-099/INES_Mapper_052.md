---
title: "INES Mapper 052"
url: "/wiki/INES_Mapper_052/"
description: "iNES Mapper 052 denotes the Realtec 8213 and similar MMC3-based multicart circuit boards."
wikiPage: "INES_Mapper_052"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 052** denotes the **Realtec 8213** and similar [MMC3](../MMC3/)-based multicart circuit boards.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    D~[LTCc SBPp]
       |||| ||++-- PRG A18..A17
       |||| |+---- PRG/CHR A19
       |||| +----- PRG A17 mode: 0=from MMC3, 1=from p
       ||++------- CHR A18..A17
       |+--------- CHR A17 mode: 0=from MMC3, 1=from c
       +---------- 1=Lock Outer Bank register until next reset
    Value on reset: $00

The MMC3's WRAM interface must be enabled and writeable in [MMC3 register \$A001](../MMC3/#PRG_RAM_protect_.28.24A001-.24BFFF.2C_odd.29). The Outer Bank Register overlaps any actual PRG RAM that may be present.

## <span id="Mixed_CHR_ROM.2FCHR_RAM_variants"></span>Mixed CHR ROM/CHR RAM variants {#Mixed_CHR_ROM.2FCHR_RAM_variants}

Nestopia Plus! has defined **Submapper 13** to denote a PCB variant with 512 KiB of each PRG ROM and CHR ROM that can switch between regular CHR ROM and 8 KiB unbanked CHR RAM. CHR RAM is selected when PRG A17 and PRG A18 are both =1.

Nestopia Plus! further uses Submapper 13 to denote the *(AB-128) Well 8-in-1* multicart with 1 MiB of each PRG ROM and CHR ROM, which is incompatible to the above definition, and which is hereby reassigned to **Submapper 14**. Its actual PCB name is **AB892**, and its Outer Bank Register has the following form:

    D~[LTRc SBBp]
       |||| |||+-- PRG A18..A17
       |||| |++--- PRG/CHR A19..A18
       |||| +----- PRG A17 mode: 0=from MMC3, 1=from p
       |||+------- CHR A17
       ||+-------- 0=CHR-ROM, 1=CHR RAM
       |+--------- CHR A17 mode: 0=from MMC3, 1=from c
       +---------- 1=Lock Outer Bank register until next reset

At least one cart (AB-134) re-uses this PCB but leaves the CHR RAM spot unpopulated, still using the combined bit for PRG/CHR A18, but rendering the "R" bit 5 meaningless. The correct NES 2.0 header for this would be submapper 14 but with zero CHR RAM.

## See also {#See_also}

- ["Mario Party" 7 in 1](http://nesdev.org/Mari7in1.txt) by The Mad Dumper.
- [Forum thread](http://forums.nesdev.org/viewtopic.php?f=9&t=10703) about mapper 052 by FARID
