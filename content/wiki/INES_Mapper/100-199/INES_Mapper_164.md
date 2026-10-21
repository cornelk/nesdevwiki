---
title: "INES Mapper 164"
url: "/wiki/INES_Mapper_164/"
description: "iNES Mapper 164 denotes the ROM chipset of the 东达 (Dōngdá) PEC-9588 Pyramid Educational Computer that was later also used on the 燕城 (Yànchéng) cy2000-3 circuit board, used for the…"
wikiPage: "INES_Mapper_164"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**iNES Mapper 164** denotes the ROM chipset of the 东达 (Dōngdá) *PEC-9588 Pyramid Educational Computer* that was later also used on the 燕城 (Yànchéng) **cy2000-3** circuit board, used for the following games:

- *Final Fantasy 太空戰士 V* (G-003)
- *大話西游* (Dàhuà Xīyóu, G-004)
- *櫻桃小丸子* (Chibi Maruko-chan, G-005)
- *岳飛傳* (Yuèfēi Zhuàn, G-006)
- *Dark Seed - 黑暗之蛊* (G-009)
- *口袋精靈: 金* (Pokémon: Gold Edition, not to be confused with several similarly-named ports of the same name)

The circuit board has address lines for 8 KiB of PRG-RAM, though all known carts mount either no or a 2 KiB RAM chip which is then mirrored three times across the respective address range. PRG-RAM is never battery-backed; all games save the game state to a 93C66 512-byte serial EEPROM. The EEPROM's "ORG" pin is tied to GND, selecting a word size of 8 bits. There is also a spot for a second EEPROM chip that is unpopulated on all known games.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 0-8 KiB unbanked PRG-RAM, not battery-backed
- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB semi-fixed PRG-ROM bank
- CPU \$8000-\$FFFF: *alternatively:* 32 KiB swichable PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

All registers are initialized to \$00 on reset.

## <span id="PRG-ROM_Bank_Low.2F1_bpp_Mode_Switch_.28.245000.2C_write.29"></span>PRG-ROM Bank Low/1 bpp Mode Switch (\$5000, write) {#PRG-ROM_Bank_Low.2F1_bpp_Mode_Switch_.28.245000.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      CSQM PPPp
      ||+|-++++- PRG A18..A14 if M=0
      || | ++++- PRG A18..A15 if M=1
      || +------ PRG banking mode
      ||          0: PRG A14..A18=QPPPp when CPU A14=0 (UxROM, 16 KiB switchable bank)
      ||             PRG A14..A18=11111 when CPU A14=1 and S=0 (fixed bank=1F)
      ||             PRG A14..A18=111p0 when CPU A14=1 and S=1 (fixed bank=1C or 1E)
      ||          1: PRG A14=CPU A14, PRG A15..A18=PPPp (BxROM, 32 KiB switchable bank)
      ||         Also selects nametable mirroring:
      ||          0: Forced vertical mirroring
      ||          1: Mirroring selected by $5300
      |+-------- See 'M' bit description
      +--------- 1 bpp video mode: when PPU A13=0 (pattern table) ...
                  0: CHR A3=PPU A3, CHR A12=PPU A12 (disable 1 bpp mode)
                  1: CHR A3=PPU A0, CHR A12=PPU A9, both latched on
                     last rise of PPU A13 (enable 1 bpp mode)

- 1 bpp mode is an all points-addressable video mode that allows the 8 KiB CHR-RAM to hold just enough data for an entire screen.
- Because this register is initialized to \$00 on reset, thus clearing bit 4 which causes UxROM mode, the game's reset vectors will be at the end of 16 KiB bank \#31.

## <span id="PRG-ROM_Bank_High_.28.245100.2C_write.29"></span>PRG-ROM Bank High (\$5100, write) {#PRG-ROM_Bank_High_.28.245100.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... ..PP
             ++- PRG A20..A19

## <span id="Microwire_Interface_.28.245200.2C_write.29"></span>Microwire Interface (\$5200, write) {#Microwire_Interface_.28.245200.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .T.S .C.D
       | |  | +- 93C66 EEPROM DAT output
       | |  +--- 93C66 EEPROM CLK output
       | +------ 93C66 EEPROM #1 CS output
       +-------- 93C66 EEPROM #2 CS output

## <span id="Mirroring_.28.245300.2C_write.29"></span>Mirroring (\$5300, write) {#Mirroring_.28.245300.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      M... ....
      +--------- Nametable mirroring if $5000.4=1
                  0=Horizontal
                  1=Vertical

## <span id="Microwire_Interface_.28.245500.2C_read.29"></span>Microwire Interface (\$5500, read) {#Microwire_Interface_.28.245500.2C_read.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .A..
            |
            |
            +--- 93C66 EEPROM DAT input (inverted)

# See also {#See_also}

- [NES 2.0 Mapper 558](../NES_2.0_Mapper_558/) is a later version which drops UxROM mode and connects the microwire interface to EEPROM differently.
- [PCB pictures and circuit diagram](https://forums.nesdev.org/viewtopic.php?p=257116)
