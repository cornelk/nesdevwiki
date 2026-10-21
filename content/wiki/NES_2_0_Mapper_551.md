---
title: "NES 2.0 Mapper 551"
url: "/wiki/NES_2.0_Mapper_551/"
description: "NES 2.0 Mapper 551 is a variant of INES Mapper 178 used for several Jncota re-releases of RPG games that were originally published by C\u0026E and Supertone:"
wikiPage: "NES_2_0_Mapper_551"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
---

**NES 2.0 Mapper 551** is a variant of [INES Mapper 178](../INES_Mapper_178/) used for several Jncota re-releases of RPG games that were originally published by C&E and Supertone:

- *聖火列傳 (Shènghuǒ Lièzhuàn)* and its title screen hack *侠客传奇 (Xiákè Chuánqí)*, originally released for [INES Mapper 240](../INES_Mapper_240/)
- *荊軻新傳 (Jīngkē Xīnzhuàn)* and its title screen hack *战国风云 (Zhànguó Fēngyún)*, originally released for [INES Mapper 240](../INES_Mapper_240/)

Compared to [INES Mapper 178](../INES_Mapper_178/), mirroring is hard-wired, and the chipset's internal CHR-RAM is not used in favor of CHR-ROM.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB unbanked PRG-RAM
- CPU \$8000-\$FFFF: 16/32 KiB PRG-ROM bank, switchable
- PPU \$0000-\$1FFF: 8 KiB switchable CHR-ROM bank
- Nametable mirroring: Hard-wired

# Registers {#Registers}

## <span id="Mode_Register_.28.244800.29"></span>Mode Register (\$4800) {#Mode_Register_.28.244800.29}

    7654 3210
    ---------
    .... .SS.
          ||
          ||
          ||
          ++-- PRG banking mode
                0: NROM-256/BNROM (PRG A14=CPU A14)
                1: UNROM (PRG A14..16=111b if CPU A14=1)
                2: NROM-128
                3: UNROM but with bit 0 of "fixed" bank selectable
                   (PRG A15..16=11b if CPU A14=1)

## <span id="Low_PRG_Bank_Register_.28.244801.29"></span>Low PRG Bank Register (\$4801) {#Low_PRG_Bank_Register_.28.244801.29}

    7654 3210
    ---------
    .... .LLL
          +++- PRG A16..A14

This can be considered an inner bank register for UNROM mode.

## <span id="High_PRG_Bank_Register_.28.244802.29"></span>High PRG Bank Register (\$4802) {#High_PRG_Bank_Register_.28.244802.29}

    7654 3210
    ---------
    HHHH HHHH
    ++++-++++- PRG A24?..A17

This can be considered an outer bank register for UNROM mode.

## <span id="CHR-ROM_Bank_register_.28.244803.29"></span>CHR-ROM Bank register (\$4803) {#CHR-ROM_Bank_register_.28.244803.29}

    7654 3210
    ---------
    BBBB BBBB
    ++++-++++- CHR-ROM A20?..A13
