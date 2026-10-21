---
title: "INES Mapper 121"
url: "/wiki/INES_Mapper_121/"
description: "iNES Mapper 121 is used for the 卡聖 (Kǎshèng) A9711 and A9713 MMC3-clone-bearing protected boards."
wikiPage: "INES_Mapper_121"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 121 is used for the 卡聖 (Kǎshèng) A9711 and A9713 [MMC3](../MMC3/)-clone-bearing protected boards.

A9711 (256 KiB PRG-ROM):

- *MK6* (title screen hack of [*幽☆遊☆白書 Final*](../INES_Mapper_115/))
- *Sonic & Knuckles 5* (hack of *Somari*, originally released by 哥德 (Gouder) on [INES Mapper 116](../INES_Mapper_116/))
- *Sonic 3D Blast 6* (hack of *Somari*, originally released by 哥德 (Gouder) on [INES Mapper 116](../INES_Mapper_116/))
- *Street Fighter Zero 2 '97* (not [*Street Fighter Zero 2*](../INES_Mapper_187/))
- *Super Real Bout 97* (hack of *餓狼伝説 Special*, originally released by 哥德 (Gouder) on INES Mapper 116)
- *The Panda Prince/Super Lion King 2/熊貓太子* (*Xióngmāo Tàizǐ*)
- *Ultimate Mortal Kombat 3* (hack of [*Mortal Kombat 3*](../INES_Mapper_123/))
- *Ultimate Mortal Kombat 4* (hack of [*Mortal Kombat 3*](../INES_Mapper_123/))
- *幽☆遊☆白書 97 V* (*Yū Yū Hakusho 97 V*, title screen hack of [*幽☆遊☆白書 Final*](../INES_Mapper_115/))

Only *Street Fighter Zero 2 '97* has 512 KiB of CHR-ROM; CHR A18 (the second half of CHR-ROM) is directly selected by PPU A12, regardless of what the MMC3 is set to do. This means that the "left" pattern table always uses the first and the "right" pattern table always uses the second 256 KiB half of CHR-ROM.

A9713 (512 KiB PRG-ROM):

- *3-in-1: Earthworm Jim 2+4, The Lion King*

# Registers {#Registers}

## <span id="Protection_Array_Read_.28.245000-.245FFF.29"></span>Protection Array Read (\$5000-\$5FFF) {#Protection_Array_Read_.28.245000-.245FFF.29}

Mask: probably \$F000

Returns the value from the four-byte protection array (\$83, \$83, \$42, \$00) indexed by the previous write to \$5000-\$5FFF.

## <span id="Protection_Array_Index_.28.245000-.245FFF.29"></span>Protection Array Index (\$5000-\$5FFF) {#Protection_Array_Index_.28.245000-.245FFF.29}

    Mask: probably $F000

    D~7654 3210
      ---------
      .... ..II
             ++- Select index into protection array on next $5000-$5FFF read

[FCEUX' source code](https://github.com/asfdfdfd/fceux/blob/master/src/boards/121.cpp) claims that address bit 8 switches between two protection arrays, the second one being \$00, \$02, \$02, \$03.

## <span id="Outer_Bank_Register_.28.245180.2C_A9713_only.29"></span>Outer Bank Register (\$5180, A9713 only) {#Outer_Bank_Register_.28.245180.2C_A9713_only.29}

    Mask: $F180

    D~7654 3210
      ---------
      M... ....
      +--------- Select outer 256 KiB PRG-ROM and CHR-ROM bank

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.29"></span>MMC3-compatible registers (\$8000-\$FFFF) {#MMC3-compatible_registers_.28.248000-.24FFFF.29}

Mask: \$E001, see [MMC3](../MMC3/).

## <span id="Protection_Latch_.28.248001.29"></span>Protection Latch (\$8001) {#Protection_Latch_.28.248001.29}

Mask: \$E003

If the previous write to the Protection Index register (\$8003) selected indices \$26, \$28 or \$2A, the value written to \$8001 --- with the lower six bits reversed -- is immediately applied. Otherwise, it will only be applied upon the next write to the \$8003. Regardless of the \$8003 setting, the MMC3 clone will always receive the write as well.

## <span id="Protection_Index_.28.248003.29"></span>Protection Index (\$8003) {#Protection_Index_.28.248003.29}

Mask: \$E003

Bits 0-5 select a protection index register. If a valid index has been selected, the value previously written to \$8001 --- with the lower six bits reversed --- is applied. If an invalid index has been selected, the \$A000/\$C000/\$E000 banks selected via \$8003 expire, and normal MMC3 PRG-ROM banking is restored.

- Index \$26: Select 8 KiB PRG-ROM bank at CPU \$E000-\$FFFF using current bit-reversed \$8001 value, and apply subsequent writes to \$8001 to this index as well.
- Index \$28: Select 8 KiB PRG-ROM bank at CPU \$C000-\$DFFF using current bit-reversed \$8001 value, and apply subsequent writes to \$8001 to this index as well.
- Index \$2A: Select 8 KiB PRG-ROM bank at CPU \$A000-\$BFFF using current bit-reversed \$8001 value, and apply subsequent writes to \$8001 to this index as well.
- Index \$2C: Select 8 KiB PRG-ROM bank at CPU \$E000-\$FFFF using current bit-reversed \$8001 value *if it is not zero*, but do not apply subsequent writes to \$8001 to this index.
- Index \$2F: Do nothing, but keep the \$A000/\$C000/\$E000 banks previously selected via \$8003.
- Indices \$20/\$29/\$2B/\$3C/\$3F: Select 8 KiB PRG-ROM bank at CPU \$E000-\$FFFF using current bit-reversed \$8001 value, but do not apply subsequent writes to \$8001 to this index.
- All other indices: Restore normal MMC3 PRG-ROM banking.

Writing to \$8003 also has the side effect of writing the same value to the MMC clone's \$8000 register.

# Errata {#Errata}

- GoodNES 3.23b has several INES Mapper 121 images erroneously set to 187.

# Similar Mappers {#Similar_Mappers}

- iNES Mapper [187](../INES_Mapper_187/) has the same CHR A18 behavior as the A9711 board, adds an NROM override register, and drops the complex protection.
- The licensed mapper [iNES Mapper 088](../INES_Mapper_088/) uses a similar setup to address larger CHR-ROM, but using the MMC3 predecessor [Namco 108](../Namcot_108_family_pinout/), and can only address 128 KiB of CHR as a result.
