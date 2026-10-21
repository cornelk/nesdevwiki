---
title: "NES 2.0 Mapper 446"
url: "/wiki/NES_2.0_Mapper_446/"
description: "NES 2.0 Mapper 446 denotes Mindkids' SMD172B_FPGA circuit board. It is used for some large multicarts from Pixel Games, as well as licensed rereleases from Retro-Bit and a handful of homebrew games:"
wikiPage: "NES_2_0_Mapper_446"
wikiCategories:
  - name: "Mappers with flash save"
    url: "/category/Mappers_with_flash_save/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 446** denotes Mindkids' **SMD172B_FPGA** circuit board. It is used for some large multicarts from *Pixel Games*, as well as licensed rereleases from *Retro-Bit* and a handful of homebrew games:

- *Jay and Silent Bob: Mall Brawl*
- *Blazing Rangers* / *炎のRangerman*
- *Hammerin' Harry 2: Dan the Red Strikes Back* / *Daiku no Gen San 2* (Retro-Bit rerelease)

# Supervisor Registers {#Supervisor_Registers}

## <span id="Mapper_Selection_.28.245000.2C_write.29"></span>Mapper Selection (\$5000, write) {#Mapper_Selection_.28.245000.2C_write.29}

    D~[L..M MMMM]
       |  +-++++-- Mapper for the FPGA to emulate
       +---------- 1=Lock supervisor registers until reset

The set of mappers that the FPGA can emulate differs across cartridges and is selected by the NES 2.0 Submapper.

| NES 2.0 Submapper | \$5000.0-4 | Mapper                                 |
|-------------------|------------|----------------------------------------|
| 0                 | 01         | [SKROM](../MMC1/)                      |
| 1                 | 00         | [UNROM](../UxROM/)                     |
| 1                 | 01         | [TKROM](../TxROM/)                     |
| 1                 | 02         | [BNROM](../INES_Mapper_034/)           |
| 1                 | 03         | [CNROM](../CNROM/)                     |
| 1                 | 04         | [ANROM](../AxROM/)                     |
| 1                 | 05         | [SKROM](../MMC1/)                      |
| 1                 | 06         | [SNROM](../MMC1/)                      |
| 1                 | 07         | [SUROM](../MMC1/)                      |
| 1                 | 08         | [GNROM](../GxROM/)                     |
| 1                 | 09         | [PNROM](../MMC2/)                      |
| 1                 | 0A         | [HKROM](../MMC6/)                      |
| 1                 | 0B         | [INES Mapper 152](../INES_Mapper_152/) |
| 1                 | 0E         | [TKSROM](../TKSROM/)                   |
| 1                 | 0F         | [INES Mapper 189](../INES_Mapper_189/) |
| 1                 | 10         | [INES Mapper 026](../VRC6/)            |
| 1                 | 12         | [INES Mapper 022](../VRC2_and_VRC4/)   |
| 1                 | 15         | [INES Mapper 025](../VRC2_and_VRC4/)   |
| 1                 | 18         | [INES Mapper 023](../VRC2_and_VRC4/)   |
| 1                 | 1A         | [VRC1](../VRC1/)                       |
| 2                 | 00         | [UNROM](../UxROM/)                     |
| 2                 | 01         | [TKROM](../TxROM/)                     |
| 2                 | 02         | [BNROM](../INES_Mapper_034/)           |
| 2                 | 03         | [CNROM](../CNROM/)                     |
| 2                 | 04         | [ANROM](../AxROM/)                     |
| 2                 | 05         | [SKROM](../MMC1/)                      |
| 2                 | 06         | [SNROM](../MMC1/)                      |
| 2                 | 07         | [SUROM](../MMC1/)                      |
| 2                 | 08         | [GNROM](../GxROM/)                     |
| 2                 | 09         | [TLROM](../TxROM/)                     |
| 2                 | 0A         | [HKROM](../MMC6/)                      |
| 2                 | 0B         | [INES Mapper 152](../INES_Mapper_152/) |
| 2                 | 0D         | [TKROM](../TxROM/)                     |
| 2                 | 0E         | [TKSROM](../TKSROM/)                   |
| 2                 | 0F         | [INES Mapper 189](../INES_Mapper_189/) |
| 2                 | 10         | [INES Mapper 026](../VRC6/)            |
| 2                 | 12         | [INES Mapper 022](../VRC2_and_VRC4/)   |
| 2                 | 15         | [INES Mapper 025](../VRC2_and_VRC4/)   |
| 2                 | 18         | [INES Mapper 023](../VRC2_and_VRC4/)   |
| 2                 | 1A         | [VRC1](../VRC1/)                       |
| 3                 | 01         | [H3001](../INES_Mapper_065/)           |

## <span id="Outer_8_KiB_PRG-ROM_Bank_LSB_.28.245001.2C_write.29"></span>Outer 8 KiB PRG-ROM Bank LSB (\$5001, write) {#Outer_8_KiB_PRG-ROM_Bank_LSB_.28.245001.2C_write.29}

    D~[BBBB BBBB]
       ++++-++++- PRG A20..A13

## <span id="Outer_8_KiB_PRG-ROM_Bank_MSB_.28.245002.2C_write.29"></span>Outer 8 KiB PRG-ROM Bank MSB (\$5002, write) {#Outer_8_KiB_PRG-ROM_Bank_MSB_.28.245002.2C_write.29}

    D~[BBBB BBBB]
       ++++-++++- PRG A28..A21

## <span id="Inner_8_KiB_PRG-ROM_Bank_Mask_.28.245003.2C_write.29"></span>Inner 8 KiB PRG-ROM Bank Mask (\$5003, write) {#Inner_8_KiB_PRG-ROM_Bank_Mask_.28.245003.2C_write.29}

    D~[MMMM MMMM]
       ++++-++++- PRG A20..A13 mask

The register has an inverted meaning in Submapper 2 vs. Submappers 0-1. In Submappers 0-1, an "1" bit means that the Outer Bank bit is used. In Submapper 2, an "1" bit means that the Inner Bank bit is used.

## <span id="Flags_.28.245004.2C_write.29"></span>Flags (\$5004, write) {#Flags_.28.245004.2C_write.29}

    D~[..CC ...M]
         ||    +- Mirroring (V=1)
         ++------ Inner CHR-RAM address mask
                  0x: 256 KiB
                  10: 128 KiB
                  11: 32 KiB

The mirroring bit has no effect when selecting an FPGA mapper with software-selectable mirroring.

## <span id="CHR_RAM_protect_.28.245005.2C_write.29"></span>CHR RAM protect (\$5005, write) {#CHR_RAM_protect_.28.245005.2C_write.29}

    D~[.... .P..]
             +- 1: Write-protect CHR-RAM

## <span id="Outer_8_KiB_CHR-RAM_Bank_.28.245006.2C_write.29"></span>Outer 8 KiB CHR-RAM Bank (\$5006, write) {#Outer_8_KiB_CHR-RAM_Bank_.28.245006.2C_write.29}

    D~[...B BBBB]
          +-++++- CHR A17..A13
