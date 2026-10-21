---
title: "NES 2.0 Mapper 558"
url: "/wiki/NES_2.0_Mapper_558/"
description: "NES 2.0 Mapper 558 denotes the 燕城 (Yànchéng) YC-03-09 circuit board, used for the following games:"
wikiPage: "NES_2_0_Mapper_558"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**NES 2.0 Mapper 558** denotes the 燕城 (Yànchéng) **YC-03-09** circuit board, used for the following games:

- *口袋精靈: 水晶* (Pokémon: Crystal Edition)
- *大話三國* (Dàhuà Sānguó)

A compatible circuit board with an unknown PCB code is used for the original editions of a few games from 外星 (Wàixīng):

- *Pet Family - 宠物大家族: 部落纷争* (ES-1081)
- *Pet Evolve - 宠物进化史* (ES-1085)
- *口袋妖怪꞉ 鑽石版* (Pokémon: Diamond Edition, ES-1090)
- *口袋精靈: 紅* and its title screen hack *宠物: 红* (Pokémon: Red Edition, ES-1088)
- *数码暴龙 4: 水晶版* and its title screen hack *数码宝贝* (Digimon 4: Crystal Edition)
- *盟军敢死队 - Commandos*

The Yànchéng games save the game state to a 93C66 512-byte serial EEPROM, its "ORG" pin tied to GND, selecting a word size of 8 bits. The Wàixīng games save to 8 KiB battery-backed WRAM. The NES 2.0 Header's PRG-NVRAM field distinguishes between the two board variants.

# Banks {#Banks}

- CPU \$6000-\$7FFF: up to 8 KiB unbanked PRG-RAM, can be battery-backed
- CPU \$8000-\$FFFF: 32 KiB switchable PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM bank
- Nametable mirroring: hard-wired

# Registers {#Registers}

All registers are initialized to \$00 on reset.

## <span id="PRG-ROM_Bank_Low_.28.245000.2C_write.29"></span>PRG-ROM Bank Low (\$5000, write) {#PRG-ROM_Bank_Low_.28.245000.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... PPPP
           ++++- PRG A18..A15

- Note that if \$5300.0 is 0, the two lowest bits are swapped during writing.

## <span id="PRG-ROM_Bank_High_.28.245100.2C_write.29"></span>PRG-ROM Bank High (\$5100, write) {#PRG-ROM_Bank_High_.28.245100.2C_write.29}

    Mask: $FB00

    D~7654 3210
      ---------
      .... ..PP
             ++- PRG A20..A19

- Note that if \$5300.0 is 0, the two lowest bits are swapped during writing.

## <span id="Microwire_Interface_.28.245200.2C_write.29"></span>Microwire Interface (\$5200, write) {#Microwire_Interface_.28.245200.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .SDC
            ||+- 93C66 EEPROM CLK ($5300.0=0)/DAT ($5300.0=1) output
            |+-- 93C66 EEPROM DAT ($5300.0=0)/CLK ($5300.0=1) output
            +--- 93C66 EEPROM CS output

## <span id="Mode_.28.245300.2C_write.29"></span>Mode (\$5300, write) {#Mode_.28.245300.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .AB.
            |+-  0: Swap D0,D1 on all register writes to $5000-$5200.
            |    1: Leave D0,D1 unmodified on register writes to $5000-$5200.
            +--- 0: PRG A15/A16=1
                 1: PRG A15/A16 from $5000

- Writes to this register are *not* subject to D0/D1 bit-swapping, i.e. the register is not affected by its own bit 0.
- Because this register is initialized to \$00 on reset, thus clearing bit 2 which causes PRG A15/A16 to be held high, the game will boot in 32 KiB PRG-ROM bank \#3.
- Changing the value of \$5300.0 without writing to \$5000-\$5200 afterwards does not change the latched bits.

## <span id="Microwire_Interface_.28.245500.2C_read.29"></span>Microwire Interface (\$5500, read) {#Microwire_Interface_.28.245500.2C_read.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .A..
            |
            |
            +--- 93C66 EEPROM DAT input

Unlike [INES Mapper 164](../INES_Mapper_164/), the EEPROM input is *not* inverted.

# Notes {#Notes}

- CPU D0 and D1 are connected in reverse to the mapper blob compared to other circuit boards using this ASIC. This means that setting \$5300.0 to 0, not to 1, will effectively swap PRG A15/A16 and PRG A19/A20. Cartridges using this mapper must be dumped with \$5300=\$07 during dumping to obtain a readout that matches that of a desoldered ROM chip.
- The Wàixīng games using this mapper were previously set to 162/163 or 164, resulting in incompatibility either because of the different address of the PRG-ROM Bank High register (mapper 162/163) or the different startup bank (164).

# See also {#See_also}

- [INES Mapper 162](../INES_Mapper_162/)/163 swap the effective meaning of registers \$5100 and \$5200, the meaning of the D0/D1 bit swap, and use the microwire interface only for protection purposes.
- [INES Mapper 164](../INES_Mapper_164/) is an earlier implementation without D0/D1 swapping, optional UNROM-like bankswitching, and a 1 bpp CHR mode.
- PCB images and analysis
