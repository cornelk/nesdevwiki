---
title: "Mask ROM pinout"
url: "/wiki/Mask_ROM_pinout/"
description: "Nintendo used by default JEDEC standard compatible pinouts for all their mask ROMs of 64 kBytes and below (but some particular boards might be exceptions !). Names [in brackets] applies when the…"
wikiPage: "Mask_ROM_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

## <span id="8kB_.2F_16kB_.2F_32kB_.2F_64kBytes_.2828pin.29_ROMs"></span>8kB / 16kB / 32kB / 64kBytes (28pin) ROMs {#8kB_.2F_16kB_.2F_32kB_.2F_64kBytes_.2828pin.29_ROMs}

Nintendo used by default JEDEC standard compatible pinouts for all their mask ROMs of 64 kBytes and below (but some particular boards might be exceptions !). Names \[in brackets\] applies when the corresponding address pin is unused. On boards where an adress pin is never used (for example, A15 is never used on [NROM](../NROM/) boards as the ROM can't be greater than 32k), what is in brackets connects to the unused pin.

      27C64/128/256/512 EPROM pinout

                  ---_---
     [+5V] A15 - |01   28| - +5V
           A12 - |02   27| - A14 [PGM]
           A7  - |03   26| - A13 [NC]
           A6  - |04   25| - A8
           A5  - |05   24| - A9
           A4  - |06   23| - A11
           A3  - |07   22| - /OE
           A2  - |08   21| - A10
           A1  - |09   20| - /CE
           A0  - |10   19| - D7
           D0  - |11   18| - D6
           D1  - |12   17| - D5
           D2  - |13   16| - D4
           GND - |14   15| - D3
                  -------

## <span id="128.2F256.2F512_KBytes_.2828.2F32pin.29_ROMs"></span>128/256/512 KBytes (28/32pin) ROMs {#128.2F256.2F512_KBytes_.2828.2F32pin.29_ROMs}

Nintendo used the standard pinout for (extremely rare) prototype boards intended to take a 27C010/020/040 EPROM. But retail Game Paks made by Nintendo have mask ROMs with a different pinout which is not compatible. This pinout, with reshuffled enable lines and higher address lines, allows a 32-pin hole to take a 28-pin 128 KiB PRG ROM in pin 3 to pin 30.

For games using 256 KiB and larger ROMs, other companies producing Famicom or NES boards used either epoxy blobs or standard EPROM pinout. But games with 128 KiB of PRG were more often in a 28-pin package than not.

On Nintendo's boards where PRG A18 is never used, pin 2 is connected with pin 22.

Nintendo's MMC5 boards use the same pinout for both PRG and CHR ROMs, and even 128 KB PRG ROMs are 32-pin so they can have a VCC pin.

```
        CHR ROM, PRG ROM, and 27C010/020/040/080 EPROM pinouts compared

 MMC5    CHR        PRG      EPROM                   EPROM        PRG   CHR   MMC5
                                         ---_---
 A17    A17 [+5V]  A17      [VPP] A19 - |01   32| - +5V
 A18    /OE        A18 [/CE]      A16 - |02   31| - A18 [PGM]    +5V    /CE2   /CE2
                                  A15 - |03   30| - A17 [NC]     +5V    +5V    A19
                                  A12 - |04   29| - A14
                                  A7  - |05   28| - A13
                                  A6  - |06   27| - A8
                                  A5  - |07   26| - A9
                                  A4  - |08   25| - A11
                                  A3  - |09   24| - /OE          A16    A16    A16
                                  A2  - |10   23| - A10
                                  A1  - |11   22| - /CE          /CE    GND    /CE
                                  A0  - |12   21| - D7
                                  D0  - |13   20| - D6
                                  D1  - |14   19| - D5
                                  D2  - |15   18| - D4
                                  GND - |16   17| - D3
                                         -------

PRG and CHR pins are listed only if they differ from EPROM.
```

## Variants {#Variants}

Here is a list of multiple variants of Nintendo's pinouts above. Only a couple of enable pins typically differs (which are shown in bold).

### <span id="Nintendo_RROM_CHR_ROM_pinout_-_8_KBytes_.2828pin.29"></span>Nintendo RROM CHR ROM pinout - 8 KBytes (28pin) {#Nintendo_RROM_CHR_ROM_pinout_-_8_KBytes_.2828pin.29}

This particular board is functionally identical to NROM but with a strange CHR pinout :

      RROM Non-standard 64-kbit CHR pinout
                 ---_---
          A7  - |01   28| - +5V
          A6  - |02   27| - A8
          A5  - |03   26| - A9
          A4  - |04   25| - A12
          A3  - |05   24| - /CE
          +5V - |06   23| - NC
          +5V - |07   22| - /OE
          A2  - |08   21| - A10
          A1  - |09   20| - A11
          A0  - |10   19| - D7
          D0  - |11   18| - D6
          D1  - |12   17| - D5
          D2  - |13   16| - D4
          GND - |14   15| - D3
                 -------

### <span id="Nintendo_SROM_CHR_ROM_pinout_-_8_KBytes_.2824pin.29"></span>Nintendo SROM CHR ROM pinout - 8 KBytes (24pin) {#Nintendo_SROM_CHR_ROM_pinout_-_8_KBytes_.2824pin.29}

This particular board is functionally identical to NROM but the CHR uses a 24-pin 8KB ROM with pinout very similar to the 27C32:

      SROM  23C62/64 JEDEC-Standard 64-kbit CHR pinout
                ---_---
          A7 - |01   24| - Vcc
          A6 - |02   23| - A8
          A5 - |03   22| - A9
          A4 - |04   21| - A12
          A3 - |05   20| - /OE
          A2 - |06   19| - A10
          A1 - |07   18| - A11
          A0 - |08   17| - D7
          D0 - |09   16| - D6
          D1 - |10   15| - D5
          D2 - |11   14| - D4
         Gnd - |12   13| - D3
                -------

With only one output enable, the board synthesizes the signal (/OE = PPU A13 + PPU /RD) using two transistors and a resistor.

### Nintendo STROM {#Nintendo_STROM}

This board uses three 23C62, 8 KBytes, 24 pin MASK ROMs (two as PRG-ROM and one as CHR-ROM) and a few discrete elements for an address decoder [\[1\]](https://forums.nesdev.org/viewtopic.php?f=9&t=3447&p=274110#p25434)

### <span id="Nintendo_AOROM_PRG_ROM_pinout_-_128.2F256.2FKBytes_.2832pin.29"></span>Nintendo AOROM PRG ROM pinout - 128/256/KBytes (32pin) {#Nintendo_AOROM_PRG_ROM_pinout_-_128.2F256.2FKBytes_.2832pin.29}

Very slight variant of the standard PRG-ROM pinout above, where an additional active high enable line is used to prevent bus conflicts.

                  ---_---
           A17 - |01   32| - +5V
           /CE - |02   31| - CE (R/W)
           A15 - |03   30| - +5V
           A12 - |04   29| - A14
           A7  - |05   28| - A13
           A6  - |06   27| - A8 
           A5  - |07   26| - A9
           A4  - |08   25| - A11
           A3  - |09   24| - A16
           A2  - |10   23| - A10
           A1  - |11   22| - /CE
           A0  - |12   21| - D7
           D0  - |13   20| - D6
           D1  - |14   19| - D5
           D2  - |15   18| - D4
          GND  - |16   17| - D3
                  -------

Japanese *After Burner* connects two such 128 kB mask roms in parallel (pin 2 = PPU /RD, pin 22 = SS4-CHR-/CE, pin31 = SS4-CHR-A17), which implies that one of them needs to have pin 31 active low and the other active high. [\[2\]](https://nescartdb.com/profile/view/3806/after-burner)

## Signal descriptions {#Signal_descriptions}

A0-A12  
address

D0-D7  
data

/CE, /OE  
The ROM will output data at address A on pins D only if all its CE (chip enable) and OE (output enable) pins are active (CE active high, /CE active low). /CE is sometimes called /CS (chip select). A memory responds faster to /OE than to /CE but draws less current while /CE is not asserted.

PGM, VPP  
Used only during EPROM programing / erasing process.

See [Cartridge connector](../Cartridge_connector/#Signal_Descriptions) for other signals descriptions.

## Converting a donor board {#Converting_a_donor_board}

To convert a cartridge board to accept EPROM:

1.  Compare the pinout of your EPROM to the mask ROM pinout expected by the board.
2.  Lift any pins whose signals differ.
3.  Solder short wires to the corresponding holes.
4.  Solder down the memory or socket with the pins that do not differ.
5.  Solder each lifted pin to the hole with the same signal.

## See also {#See_also}

- [NES EPROM Conversions](http://nesdev.org/NES%20EPROM%20Conversions.txt) by Drk. Instructions on how to modify certain boards to use EPROMs.
- [NES ROM pinouts](http://nesdev.org/NES%20ROM%20Pinouts.txt) by Drk. Covers all PRG, CHR, and RAM chips used in NES cartridges.
- [EPROM pinouts](http://nesdev.org/EPROM%20Pinouts.txt) by Drk.
- [Advanced MMC3 NES Reproduction Tutorial](http://callanbrown.com/index.php/advanced-mmc3-nes-reproduction) by Callan Brown
- [SNROM to SUROM](http://forums.nesdev.org/viewtopic.php?p=171375#p171375) by Ice Man
