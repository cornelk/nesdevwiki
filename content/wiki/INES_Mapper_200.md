---
title: "INES Mapper 200"
url: "/wiki/INES_Mapper_200/"
description: "iNES Mapper 200 denotes the circuit boards of simple address-latch-based NROM-128 multicarts, i.e. CPU $8000-$BFFF is always a mirror of CPU $C000-$FFFF."
wikiPage: "INES_Mapper_200"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 200** denotes the circuit boards of simple address-latch-based NROM-128 multicarts, i.e. CPU \$8000-\$BFFF is always a mirror of CPU \$C000-\$FFFF.

The mapper number is used for two variants that differ in the bit that selects nametable mirroring.

PCB that use this mapper :

- **MG109** for the *[1993 Super 50 in 1 Game](../1993_Super_50_in_1_Game/)* multicart

# <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

## Submapper 0 {#Submapper_0}

    A~[1... .... .... bBBB]
                      |+++- PRG A16..A14, CHR A15..A13
                      +---- PRG A17, CHR A16, Mirroring:
                             0: Vertical
                             1: Horizontal

## Submapper 1 {#Submapper_1}

    A~[1... .... .... .bBB]
                       |++- PRG A15..A14, CHR A14..A13
                       +--- PRG A16, CHR A15, Mirroring:
                             0: Vertical
                             1: Horizontal

# Notes {#Notes}

- Some menus select bank numbers beyond the PRG ROM size, expecting to read back a solder pad value to select one of several displayed game counts.
- [NES 2.0 Mapper 338](../NES_2.0_Mapper_338/) flips the meaning of the mirroring bit.
