---
title: "NES 2.0 Mapper 518"
url: "/wiki/NES_2.0_Mapper_518/"
description: "NES 2.0 Mapper 518 is used for several games and educational computer cartridges from Subor:"
wikiPage: "NES_2_0_Mapper_518"
---

**NES 2.0 Mapper 518** is used for several games and educational computer cartridges from Subor:

- *小霸王 Subor 999*
- *小霸王 Subor V*
- *跳舞天使: 動感 2000* (also known as *Dance 2000 12-in-1*)

Its UNIF board name is **UNL-DANCE2000**. It banks PRG-ROM in 16 or 32 KiB amounts, has 8 KiB WRAM, and 2x4 KiB of CHR-RAM that can be bankswitched according to the current nametable offset during rendering.

*小霸王 Subor V* was sold both as the built-in BIOS of the Subor SB-97 educational computer, and as a stand-alone cartridge; both versions have identical ROM content. The stand-alone cartridge is a plug-through cartridge to which an expansion cartridge can be connected. In its built-in SB-97 configuration, is additionally has access to a 3.5" floppy disk drive and controller, as well as 128 KiB of secondary PRG-RAM.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB of unbanked primary PRG-RAM
- CPU \$8000-\$BFFF: 16 KiB of PRG-ROM or PRG-RAM, switchable
- CPU \$C000-\$FFFF: 16 KiB of PRG-ROM, fixed to bank \#0
- PPU \$0000-\$0FFF: 4 KiB of CHR-RAM, switchable
- PPU \$0000-\$1FFF: 4 KiB of CHR-RAM, fixed to bank \#1

## <span id="PRG-ROM_Bank_Register_.28.245000.29"></span>PRG-ROM Bank Register (\$5000) {#PRG-ROM_Bank_Register_.28.245000.29}

    D~[REPP PPPP]
       ||++-++++- Select 16 or 32 KiB PRG bank at CPU $8000
       |+-------- 0: Bank comes from main cartridge
       |          1: Bank comes from expansion cartridge
       +--------- 0: Bank comes from ROM (main or expansion cartridge)
                  1: Bank comes from 128 KiB secondary PRG-RAM (SB-97 only)

In 16 KiB PRG banking mode, \$C000-\$FFFF always comes from main cartridge ROM bank \#0. ROM images of expansion cartridges contain the main cartridge data in their first 1 MiB, followed by the expansion cartridge's 1 MiB of data. This means that the 'E' bit needs no special treatment in emulators, as it merely maps to PRG A20 of the combined 2 MiB address space.

## <span id="Mode_Register_.28.245200.29"></span>Mode Register (\$5200) {#Mode_Register_.28.245200.29}

    D~[.... .SBM]
             ||+- Select nametable arrangement
             ||   0: Horizontal arrangement/vertical mirroring
             ||   1: Vertical arrangement/horizontal mirroring
             |+-- Select CHR-RAM auto-bankswitch
             |     0: disabled
             |     1: enabled
             +--- Select PRG-ROM bank size
                   0: 16 KiB PRG-ROM bank at CPU $8000-$BFFF, $C000-$FFFF fixed to first PRG-ROM bank
                   1: 32 KiB PRG-ROM bank at CPU $8000-$FFFF

If "CHR-RAM auto-bankswitch" is enabled, CHR A12 is substituted with the last CIRAM A10 (latched during the previous rising edge of PPU A13) during reads only from PPU \$0000-\$0FFF. This means that the tile data for the first/second nametable come from the first/second 4 KiB of CHR-RAM, with "second" being the right/lower nametable when horizontal/vertical arrangement is active. By setting the scroll position such that a nametable boundary occurs in the middle of a scanline, a horizontal (mid-scanline) bankswitch is possible.

## <span id="LPC_Speech_Chip_.28.245300.29"></span>LPC Speech Chip (\$5300) {#LPC_Speech_Chip_.28.245300.29}

Read:

    D~[R... EEEE]
       |    ++++- "Energy" field of current LPC frame being played
       +--------- 1=Chip ready to receive data

Write:

    D~[DDDD DDDD]
       ++++-++++- LPC-D6 speech data

The LPC chip is a customized Texas Instruments TSP50C04 that can operate in two modes: it can play back LPC data in the original TMS5220-PE format, used for English-language speech, as well as the newer LPC D6 format, used for Mandarin-language speech. Each speech message must be preceded by 8 bits denoting the codec to be used: 0Ah for LPC D6 data, 0Bh for TMS5220-PE data; these values are from the CPU point of view (the chip reverses the bit order internally). Refer to the [TMS5220 data sheet](https://www.sprow.co.uk/bbc/hardware/speech/tms5220.pdf), U.S. Patents [4,209,844](https://patents.google.com/patent/US4829573A/en), [4,331,836](https://patents.google.com/patent/US4331836/en) and [4,335,277](https://patents.google.com/patent/US4335277/en) as well as [MAME's source code](https://github.com/mamedev/mame/blob/master/src/devices/sound/tms5220.cpp) for details on the original TMS5220-PE LPC data format and decoding, and to the sample program in Appendix B of the [TSP50C0x/1 Family Design Manual](https://www.ti.com/lit/ml/spss011d/spss011d.pdf) as well as [VirtuaNES BBK's source code](https://gitee.com/fanoble/emulator-bbk/blob/sb2k/NES/ApuEX/LPC_D6_SYNTH.C) for information on decoding the LPC D6 data format.

Data must be written into the speech data register \$5300 while the chip signals Ready (\$5300.7=1) until the end of the message is reached, then a terminating \$0F byte must be sent before the chip will accept a new message. The end of the message is denoted by a variable-bit-length frame's "energy" field having a value of 1111 binary (15), which the chip conveniently returns on reads from register \$5300.

## <span id="Floppy_Disk_Controller_.28.2455xx-.2456xx.29"></span>Floppy Disk Controller (\$55xx-\$56xx) {#Floppy_Disk_Controller_.28.2455xx-.2456xx.29}

- \$5500 (write): Corresponds to a µPD765-compatible floppy disk controller's Configuration Control register.
- \$5501 (write): Corresponds to a µPD765-compatible floppy disk controller's Digital Output register.
- \$5505 (write): Corresponds to a µPD765-compatible floppy disk controller's Data Register.
- \$5604 (read): Corresponds to a µPD765-compatible floppy disk controller's Main Status register.
- \$5605 (read): Corresponds to a µPD765-compatible floppy disk controller's Data Register.

## Notes {#Notes}

- Note that these games require Dendy video timing and will produce graphical glitches with NTSC or regular PAL timing; *Subor V* will freeze with NTSC timing.
- [NES 2.0 Mapper 514](../NES_2.0_Mapper_514/) is a simpler form of this mapper with only 32 KiB PRG-ROM banking.
- VirtuaNES emulates this mapper as number 168.
