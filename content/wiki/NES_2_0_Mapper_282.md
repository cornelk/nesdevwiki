---
title: "NES 2.0 Mapper 282"
url: "/wiki/NES_2.0_Mapper_282/"
description: "NES 2.0 Mapper 282 is used by at least three J.Y. multicarts:"
wikiPage: "NES_2_0_Mapper_282"
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

NES 2.0 Mapper 282 is used by at least three J.Y. multicarts:

- *1997 Super HIK 18-in-1 (JY-101)*
- *1997 Super HIK 21-in-1 (JY-105)*
- *1998 Super HIK 5-in-1 (JY-114)*
- *Super 25-in-1 Final Fight (SC-128)*

Each of them have 1 MiB each of PRG-ROM and CHR-ROM. They use the [J.Y. Company ASIC](../J.Y._Company_ASIC/). Register \$D003 controls 256 KiB outer banks.

## <span id="Outer_Bank_Register_.28.24D003.29"></span>Outer Bank Register (\$D003) {#Outer_Bank_Register_.28.24D003.29}

    Mask: $F003

    7654 3210
    ---- ----
    ..LM CPPc
      || |||+- Select 256 KiB outer CHR-ROM bank (CHR A18), ignored if L=1
      || |++-- Select 256/512 KiB outer PRG-ROM bank
      || +---- If bit 4==1: Select 512 KiB outer CHR-ROM bank
      |+------ 0: Ignore bit 3
      |        1: Obey bit 3
      +------- Select outer CHR-ROM bank size
                0: Mask $900x/$A00x to 256 KiB, use c
                1: Mask $900x/$A00x to 512 KiB, ignore c

PRG (\$8000-\$8FFF) and CHR (\$9000-\$AFFF) bank select register bits that select 256 KiB banks (L=0)/512 KiB banks (L=1) are masked off.

For a description of all other registers, see [J.Y. Company ASIC](../J.Y._Company_ASIC/).

# Trivia {#Trivia}

Some of the multicarts show different version of menu according to what \$5000 is returned from the mapper (the value that is read is masked with \$C0, so only bits 7 & 6 matters.):

    1997 Super HIK 18-in-1 (JY-101)
     $00   -> 18 in 1
     $40   -> 900 in 1
     $80   -> 1600 in 1
     $C0   -> 6000000 in 1

    1997 Super HIK 21-in-1 (JY-105)
     $00   -> 21 in 1
     $40   -> 350 in 1
     $80   -> 1200 in 1
     $C0   -> 7000000 in 1
