---
title: "NES 2.0 Mapper 339"
url: "/wiki/NES_2.0_Mapper_339/"
description: "NES 2.0 mapper 339 denotes several very similar MMC3+NROM multicarts, differentiated by submapper numbers that determine the address bits that select NROM-256 over NROM-128 PRG banking mode."
wikiPage: "NES_2_0_Mapper_339"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 mapper 339** denotes several very similar [MMC3](../MMC3/)+NROM multicarts, differentiated by submapper numbers that determine the address bits that select NROM-256 over NROM-128 PRG banking mode.

## <span id="Outer_Bank_.28.246000-.247FFF.2C_write.29"></span>Outer Bank (\$6000-\$7FFF, write) {#Outer_Bank_.28.246000-.247FFF.2C_write.29}

    A~[0110 .... P.MC Cbba] (Address AND $E000)==$6000, MMC3 $A001.7=1, $A001.6=0
                 | || |||+- PRG A14 in NROM-128 mode
                 | || |++-- PRG A16..A15 in NROM modes
                 | |+-+---- PRG/CHR A18..A17 in all modes
                 | +------- 0: NROM-128/-256 PRG mode: PRG A16..A14 from bits "bba", register remains writable
                 |          1: MMC3 PRG mode: PRG A16..A14 from MMC3 output, register write-protected
                 +--------- 0: Normal PRG-ROM addressing
                            1: PRG A1..A0 from solder pad for menu selection

## NROM-256 selection {#NROM-256_selection}

Every multicart using this mapper has its own method of selecting NROM-256 mode, in which PRG A14 comes directly from CPU A14 rather than the "a" bit, that is denoted by NES 2.0 submapper:

- **Submapper 0**: NROM-256 if (Address AND \$06) == \$06 (K-3006 PCB, UNIF MAPR **BMC-K-3006**)
- **Submapper 1**: NROM-256 if (Address AND \$04) != \$00 (unmarked PCB)
- **Submapper 2**: NROM-256 if (Address AND \$11) != \$00 (Realtec 8058 PCB)
- **Submapper 3**: NROM-256 if (Address AND \$18) != \$00 (K-3091/GN-16 PCB)
- **Submapper 4**: NROM-256 if (Address AND \$14) != \$00 (GR-002-31 PCB)

## See also {#See_also}

- [GR-002-31 PCB analysis](https://forums.nesdev.org/viewtopic.php?p=304532)
