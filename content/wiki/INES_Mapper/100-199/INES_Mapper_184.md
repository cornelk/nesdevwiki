---
title: "INES Mapper 184"
url: "/wiki/INES_Mapper_184/"
description: "This mapper is for a few games which use the Sunsoft-1 mapper on the Sunsoft-K PCB."
wikiPage: "INES_Mapper_184"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

This mapper is for a few games which use the [Sunsoft-1 mapper](../Sunsoft_1_pinout/) on the Sunsoft-K PCB.

- Atlantis no Nazo
- 東海道五十三次
- The Wing of Madoola

It does not include *Fantasy Zone* (J), despite using the Sunsoft-1 chip; see below.

Despite the mapper existing as a single IC, its functionality is describable using a [74139](../74139/), a [74174](../74174/), and a [74157](../74157/), and so should probably be considered discrete logic.

fwNES, a NES emulator for DOS, originally assigned mapper **122** for this mapper.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

## Register {#Register}

### <span id="CHR_ROM_bank_selects_.28.246000-.247FFF.29"></span>CHR ROM bank selects (\$6000-\$7FFF) {#CHR_ROM_bank_selects_.28.246000-.247FFF.29}

    7  bit  0
    .1HH .LLL
     |||  |||
     |||  +++- Select 4k CHR bank from $0000-0FFF
     +++------ Select 4k CHR bank from $1000-1FFF; only banks 4-7 selectable; 
                 existing software is well behaved and writes to this with the $40s bit set

This mapper precludes iNES-default RAM because the register is mapped into \$6000-7FFF.

## Fantasy Zone {#Fantasy_Zone}

*Fantasy Zone* contains code that's compatible with both the SUNSOFT-4 PCB described here, and a not-yet-found release on the SUNSOFT-3R PCB with a SUNSOFT-2 ASIC. Because SUNSOFT-3R was already allocated to [Mapper 093](../INES_Mapper_093/), that is the preferred way to support this game.

The SUNSOFT-4 PCB instead provides 8 KiB of unbanked CHR-RAM and repurposes the CHR lines for PRG banking instead:

    7  bit  0
    .... .PPP
          |||
          +++- Selects 16k PRG @ $8000-BFFF

Bank at \$C000-\$FFFF is fixed to the last 16k.

You do not need to support this, nor is there any utility in doing so.

This distinction is similar to what happened with [mapper 133](../INES_Mapper_133/)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) FWNES98E.TXT, fwNES 0.302b, 1998/11/16, by FanWen Yang
