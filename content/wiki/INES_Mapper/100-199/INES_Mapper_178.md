---
title: "INES Mapper 178"
url: "/wiki/INES_Mapper_178/"
description: "iNES Mapper 178 is used by some games from Waixing (on a PCB named FS305), Nanjing (on a PCB named NJ0430), Jncota and Hengge Dianzi as well as some educational computers from GameStar, a.k.a.…"
wikiPage: "INES_Mapper_178"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with large PRG RAM"
    url: "/category/Mappers_with_large_PRG_RAM/"
---

**iNES Mapper 178** is used by some games from Waixing (on a PCB named **FS305**), Nanjing (on a PCB named **NJ0430**), Jncota and Hengge Dianzi as well as some educational computers from GameStar, a.k.a. Belsonic. The chipset is designed for using 8 KiB of chip-internal unbanked CHR-RAM, but can address external RAM/ROM as well, banked in 8 KiB amounts. Circuit boards described by iNES Mapper 178 connect external RAM as bankable PRG-RAM mapped to CPU \$6000-\$7FFF. The following submapper are defined:

- Submapper 0: No special features
- Submapper 1: Infrared Sensor that generates interrupt requests upon detecting movement
- Submapper 2: TMS5220-compatible LPC speech chip
- Submapper 3: Solder pad read-out

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG-RAM bank
- CPU \$8000-\$FFFF: 16/32 KiB PRG-ROM bank, switchable
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM
- Nametable mirroring: Switchable horizontal/vertical settings

# Registers {#Registers}

## <span id="Mode_Register_.28.244800.29"></span>Mode Register (\$4800) {#Mode_Register_.28.244800.29}

    7654 3210
    ---------
    .... .SSM
          ||+- Nametable mirroring
          ||    0: Vertical
          ||    1: Horizontal
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

## <span id="PRG-RAM_Bank_register_.28.244803.29"></span>PRG-RAM Bank register (\$4803) {#PRG-RAM_Bank_register_.28.244803.29}

    7654 3210
    ---------
    BBBB BBBB
    ++++-++++- PRG-RAM A20?..A13

PRG-RAM is mapped to CPU \$6000-\$7FFF and is banked in 8 KiB amounts.

## <span id="Infrared_Sensor_.28Submapper_1.29"></span>Infrared Sensor (Submapper 1) {#Infrared_Sensor_.28Submapper_1.29}

Some cartridges have an additional infrared sensor. Infrared receiving is working by having two registers:

     $5000-$5FFF:  [.... ...V]  (read only)
        V = read current bit from IR sensor
     $6000-$7FFF:  [I... ....]  (write only)
        I = set to activate interrupts by infrared signal

The I bit is set at power-on, i.e. without writing to the \$6000-\$7FFF register, interrupts will be *enabled*.

## <span id="TMS5220-compatible_LPC_Speech_Chip_.28Submapper_2.29"></span>TMS5220-compatible LPC Speech Chip (Submapper 2) {#TMS5220-compatible_LPC_Speech_Chip_.28Submapper_2.29}

The same speech chip as the one used in [Mapper 241](../INES_Mapper_241/#LPC_Speech_Chip), also with the register at CPU \$5000-\$5FFF.

## <span id="Solder_Pad_register_.28Submapper_3.29"></span>Solder Pad register (Submapper 3) {#Solder_Pad_register_.28Submapper_3.29}

After writing \$01 (or likely, any value with bit 0 set) to \$6801 (or likely, anywhere to CPU \$6000-\$7FFF), reading from PRG-ROM \$8000-\$FFFF returns the 2-bit solder pad value rather than the actual PRG-ROM content. This value is used for selecting one of four different multicart menu.

# Errata {#Errata}

- The commonly-available ROM images of
  - *宠物: 小精灵 IV (Chǒngwù: Xiǎo Jīnglíng IV)*, headerless CRC32 0xB0B13DBD
  - *Education Computer 32-in-1 (Game Star) \[!\]*, headerless CRC32 0xF834F634

will not work when emulated using the above description. The meaning of registers \$4801 and \$4802 is swapped, and an emulator must shift the value of \$4802 one bit to the left before ORing with \$4801. While it cannot be ruled out that these games use an alternative version of the hardware, the more likely explanation is that these are simply bad dumps (notwithstanding GoodNES 3.23b's \[!\] tag) where the dumper confused the two bank registers.

- Supposedly, the new PRG bank is only committed when \$4801 is written. Emulating this supposed behavior however will break *Ping Pong*, while commiting the PRG bank after any \$480x write does not seem to break anything else.
- Most of the educational computers using this mapper require Dendy video timing and will freeze at various points when run with NTSC timing.
- The Waixing FS305 circuit board connects PRG A14 directly to CPU A14, making all PRG banking modes except BNROM unusable. Since it writes the correct banking mode to \$4800, it does not require any special treatment.

# See also {#See_also}

- <https://forums.nesdev.org/viewtopic.php?t=16657>
