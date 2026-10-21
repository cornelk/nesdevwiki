---
title: "NES 2.0 Mapper 529"
url: "/wiki/NES_2.0_Mapper_529/"
description: "NES 2.0 Mapper 529 denotes the YY0807, J-2148 and T-230 PCBs, used for particular bootlegs version of Datach Dragon Ball Z (T-230 PCB, with CHR-RAM) and Dragon Ball Z Gaiden (YY0807/J-2148 PCBs, with…"
wikiPage: "NES_2_0_Mapper_529"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

**NES 2.0 Mapper 529** denotes the **YY0807**, **J-2148** and **T-230** PCBs, used for particular bootlegs version of *Datach Dragon Ball Z* (**T-230** PCB, with CHR-RAM) and *Dragon Ball Z Gaiden* (**YY0807**/**J-2148** PCBs, with CHR-ROM). It's a [VRC4](../VRC2_and_VRC4/) clone (A2/A3, VRC4e) connected such that the \$800x register has no effect, and the \$A00x register selects a 16 KiB PRG-ROM bank, similar to the games' original Bandai mappers. Its UNIF board name is **UNL-T-230**.

Similarly to the Bandai mappers, it also supports saving data to serial EEPROM. The presence of EEPROM is denoted via the NES 2.0 header's PRG-NVRAM field indicating 256 byte of non-volatile save data. Unlike the original Bandai FCG boards that used an I²C interface, the bootleg variants use a 93C56 EEPROM via a Microwire interface, operated using the following registers:

## <span id="Read_Microwire_Interface_.28.245000.2C_read.29"></span>Read Microwire Interface (\$5000, read) {#Read_Microwire_Interface_.28.245000.2C_read.29}

    Mask: $F000

    D~[.... ...D]
               +- Serial Data Output from the 93C56 EEPROM

## <span id="Write_Microwrite_Interface_.28.248800.2C_write.29"></span>Write Microwrite Interface (\$8800, write) {#Write_Microwrite_Interface_.28.248800.2C_write.29}

    Mask: $8800

    D~[.... .ECD]
             ||+- Serial Data Input to 93C56 EEPROM
             |+-- Serial Clock to 93C56 EEPROM
             +--- Chip Select to 93C56 EEPROM

This register's address range overlaps with the VRC4's. The game writes to the \$FFFx address range to operate this register.

The 93C56's EEPROM's "ORG" pin is unconnected, indicating a 16-bit word length.

See also: [Board pictures and circuit diagrams](https://forums.nesdev.org/viewtopic.php?p=243032)
