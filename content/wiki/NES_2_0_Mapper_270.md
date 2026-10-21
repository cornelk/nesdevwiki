---
title: "NES 2.0 Mapper 270"
url: "/wiki/NES_2.0_Mapper_270/"
description: "NES 2.0 Mapper 270 denotes OneBus console multicarts that use the consoles' universal input/output (UIO) register $412C to bankswitch higher-order PRG address lines or several PRG chips, and select…"
wikiPage: "NES_2_0_Mapper_270"
wikiCategories:
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
  - name: "VTxx mappers"
    url: "/category/VTxx_mappers/"
---

**NES 2.0 Mapper 270** denotes [OneBus](../NES_2.0_Mapper_256/) console multicarts that use the consoles' universal input/output (UIO) register \$412C to bankswitch higher-order PRG address lines or several PRG chips, and select CHR-RAM via \$4242.

# Submapper 0: Combination {#Submapper_0:_Combination}

Responds to all bank bits of submappers 1-3 at the same time. This is permissible, since the only bit that is used by several submappers has the same function in each of them.

# Submapper 1: Game Prince RS-16 {#Submapper_1:_Game_Prince_RS-16}

## <span id="Outer_PRG_Bank_.28.24412C.2C_write.29"></span>Outer PRG Bank (\$412C, write) {#Outer_PRG_Bank_.28.24412C.2C_write.29}

    D~7654 3210
      ---------
      .... ..A.
             +-- PRG/CHR A24

# Submapper 2: Family Pocket 638-in-1 {#Submapper_2:_Family_Pocket_638-in-1}

## <span id="Outer_PRG_Bank_.28.24412C.2C_write.29_2"></span>Outer PRG Bank (\$412C, write) {#Outer_PRG_Bank_.28.24412C.2C_write.29_2}

    D~7654 3210
      ---------
      .... ..AB
             |+- PRG/CHR A25
             +-- PRG/CHR A24

## <span id="Jumper_Cartridge_Detection_.28.24412C.2C_read.29"></span>Jumper Cartridge Detection (\$412C, read) {#Jumper_Cartridge_Detection_.28.24412C.2C_read.29}

    D~7654 3210
      ---------
      .... B...
           +---- 0: Jumper cartridge not inserted
                 1: Jumper cartridge inserted

## <span id="CHR-RAM_enable_.28.244242.2C_write.29"></span>CHR-RAM enable (\$4242, write) {#CHR-RAM_enable_.28.244242.2C_write.29}

    D~7654 3210
      ---------
      .... ...C
              +- 0: Use normal OneBus CHR banking
                 1: Use 8 KiB of unbanked CHR-RAM

# Submapper 3: Bittboy 300-in-1 {#Submapper_3:_Bittboy_300-in-1}

## <span id="Outer_PRG_Bank_.28.24412C.2C_write.29_3"></span>Outer PRG Bank (\$412C, write) {#Outer_PRG_Bank_.28.24412C.2C_write.29_3}

    D~7654 3210
      ---------
      .... .A..
            +--- PRG/CHR A24

# Notes {#Notes}

- Any \$412C bit can only be read from if the corresponding bit in register \$412B (UIO input/output selector) is set to "0" (input).
- Any \$412C bit can only be written to if the corresponding bit in register \$412B (UIO input/output selector) is set to "1" (output).
