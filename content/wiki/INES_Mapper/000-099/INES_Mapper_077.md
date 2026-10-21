---
title: "INES Mapper 077"
url: "/wiki/INES_Mapper_077/"
description: "iNES Mapper 077 is a singleton allocated for the game Napoleon Senki."
wikiPage: "INES_Mapper_077"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR ROM and CHR RAM"
    url: "/category/Mappers_with_CHR_ROM_and_CHR_RAM/"
---

**iNES Mapper 077** is a singleton allocated for the game [Napoleon Senki](https://nescartdb.com/profile/view/2260/napoleon-senki).

     Registers: (** BUS CONFLICTS **)
     ---------------------------

       $8000-FFFF:  [CCCC PPPP]
         C = CHR Reg (2k @ $0000)
         P = PRG Reg (32k @ $8000)

     CHR Setup:
     ---------------------------
     CHR-RAM is fixed at $0800-$1FFF.  CHR-ROM is swappable at $0000:

        $0000-$07FF $0800-$0FFF $1000-$17FF $1800-$1FFF $2000-$27FF  $2800-$2FFF  $3000-$3FFF
       +-----------+-----------+-----------+-----------+-----------+-------------+----------+
       | $8000,ROM |  {1},RAM  |  {2},RAM  |  {3},RAM  |  {0},RAM  |Internal VRAM| Open Bus |
       +-----------+-----------+-----------+-----------+-----------+-------------+----------+

     When making an emulator, you do not need to care about the specific order of
      the CHR-RAM banks: just provide 10KiB from $0800-$2FFF.

## Hardware {#Hardware}

Two [74161s](../74161/) latch the 8 bits above.

A [74138](../74138/) divides the PPU's 16KiB of address space into eight 2KiB regions. One of these goes straight to the CHR ROM; another goes straight to the CIRAM.

A 74'21 (two 4-input AND gates) combines the other four outputs to generate a chip enable for the 8KiB RAM on the PPU.
