---
title: "INES Mapper 163"
url: "/wiki/INES_Mapper_163/"
description: "iNES Mapper 163 denotes the 南晶 (Nánjīng) FC-001 circuit board, used on most of their games, including:"
wikiPage: "INES_Mapper_163"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**iNES Mapper 163** denotes the 南晶 (Nánjīng) **FC-001** circuit board, used on most of their games, including:

- *牧场物语 - Harvest Moon* (NJ011)
- *水浒神兽* (Shuǐhǔ Shénshòu, NJ019)
- *暗黑破坏神 - Diablo* (NJ037)
- *轩辕剑外传 之 天之痕* (Xuānyuánjiàn Wàizhuàn zhī Tiānzhīhén, NJ045)
- *Final Fantasy IV - 最终幻想 4꞉ 光与暗 水晶纷争* (NJ098)
- *雷电皇 比卡丘传说* (NJ046)

A variant that mounts the same ADPCM chip used in [NES 2.0 Mapper 419](../NES_2.0_Mapper_419/) with the PCB code **NJ-YUYIN0106** (standing for 语音 Yǔyīn "voice") is denoted by **Submapper 1**, used only on the *真 Samurai Spirits - 武士道列传: 侍魂* (NJ-033) cartridge.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB unbanked PRG-RAM, battery-backed
- CPU \$8000-\$FFFF: 32 KiB switchable PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM bank, 4 KiB can be automatically switched
- Nametable mirroring: hard-wired

# Registers {#Registers}

All registers are initialized to \$00 on reset.

## <span id="PRG_Bank_Low.2FCHR-RAM_Switch_.28.245000.2C_write.29"></span>PRG Bank Low/CHR-RAM Switch (\$5000, write) {#PRG_Bank_Low.2FCHR-RAM_Switch_.28.245000.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      C... PPPP
      |    ++++- PRG A18..A15
      +--------- Enable automatic 4 KiB CHR-RAM switch: when PPU A13=0 (pattern table) ...
                  0: CHR A12=PPU A12 (disable auto-switch)
                  1: CHR A12=PPU A9 latched on last rise of PPU A13 (enable auto-switch)

- Automatic 4 KiB CHR-RAM switch means that the left pattern table is used for the the top half of any nametable, and the right pattern table for the bottom half of any nametable, regardless of the scroll position. This auto-switch behavior is similar to that of [mapper 96](../INES_Mapper_096/).
- Bits 0 and 1 are subject to the [Mode register](#Mode_.28.245300.2C_write.29).

## <span id="PRG_Bank_High_.28.245200.2C_write.29"></span>PRG Bank High (\$5200, write) {#PRG_Bank_High_.28.245200.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... ..PP
             ++- PRG A20..A19

- Bits 0 and 1 are subject to the [Mode register](#Mode_.28.245300.2C_write.29). 1 MiB games connect both ASIC PRG A19 and A20 outputs to ROM A19, effectively exempting this register from the bit-swap.

## <span id="Feedback_Write_.28.245100-.245101.2C_write.29"></span>Feedback Write (\$5100-\$5101, write) {#Feedback_Write_.28.245100-.245101.2C_write.29}

    Mask: $FF01

    D~7654 3210  A~FEDC BA98 7654 3210
      ---------    -------------------
      .... .F.E    .... .... .... ...A
            | |                      +- Action on write
            | +- A=0: Value latched
            |    A=1: Flip the latched F bit when E=1
            +--- A=0: Value latched
                 A=1: Ignored, previously-latched value either kept or flipped

- Bit 0 is subject to the [Mode register](#Mode_.28.245300.2C_write.29).

## <span id="Feedback_Read_.28.245500-.245501.2C_read.29"></span>Feedback Read (\$5500-\$5501, read) {#Feedback_Read_.28.245500-.245501.2C_read.29}

    Mask: $F300

    D~7654 3210
      ---------
      .... .F..
            +--- Inverted value of latched F bit

- *暗黑破坏神 - Diablo* (Ànhēi Pòhuàishén, NJ037) only checks that F is read back inverted.
- *真 Samurai Spirits - 魂之利刃* (NJ042) and *轩辕剑外传 之 天之痕* (Xuānyuánjiàn Wàizhuàn zhī Tiān zhī Hén, NJ045) check that F is flipped after a \$5101 write with E=1.

## <span id="Mode_.28.245300.2C_write.29"></span>Mode (\$5300, write) {#Mode_.28.245300.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .A?B
            | +- 0: Do not swap D0,D1 on writes to $5000-$5200.
            |    1: Swap D0,D1 on writes to $5000-$5200.
            +--- 0: PRG A15/A16=11b
                 1: PRG A15/A16 from $5000

- This register is not subject to bit-swapping, i.e. the register is not affected by its own bit 0.
- Because reset clears the A bit, games will boot in 32 KiB PRG-ROM bank \#3.

# Notes {#Notes}

- Mapper 163 cartridges should be dumped with \$5300=\$04 to obtain the correct PRG bank order. To verify that the cartridge is really mapper 163 ...
  - ... rule out [INES Mapper 162](../INES_Mapper_162/) by writing \$04 to \$5300 and verifying that \$5100.1 does not change the PRG bank;
  - ... rule out [NES 2.0 Mapper 558](../NES_2.0_Mapper_558/) by verifying that \$5000 value \$01 does not result in a different PRG bank between \$5300 values \$04 and \$05 (because the B bit is bit 1 in mapper 558).
- *暗黑破坏神 - Diablo* (NJ037) and its title screen hack *毁灭之神* (NJ078) mistakenly disable automatic 4 KiB CHR-RAM switching mid-screen during the ending cutscene, causing flickering on real hardware. *魔兽世界: 恶魔猎人* (NJ097) puts the 8x16 sprites into the wrong pattern table for automatic 4 KiB CHR-RAM switching, causing a glitched cursor on the title screen on real hardware.

# See also {#See_also}

- Similar mappers: [INES Mapper 162](../INES_Mapper_162/), [INES Mapper 164](../INES_Mapper_164/), [NES 2.0 Mapper 558](../NES_2.0_Mapper_558/)
