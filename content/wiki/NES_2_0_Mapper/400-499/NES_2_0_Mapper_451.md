---
title: "NES 2.0 Mapper 451"
url: "/wiki/NES_2.0_Mapper_451/"
description: "NES 2.0 Mapper 451 is used for the homebrew game Haratyler HG/MP. It is basically a homebrew TLROM-like circuit board that implements the MMC3 register's in an unusual fashion, and saves the high…"
wikiPage: "NES_2_0_Mapper_451"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with flash save"
    url: "/category/Mappers_with_flash_save/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**NES 2.0 Mapper 451** is used for the homebrew game *Haratyler HG/MP*. It is basically a homebrew TLROM-like circuit board that implements the MMC3 register's in an unusual fashion, and saves the high score to flash ROM. The game executes the AMIC A29040B flash ROM chip's "Software ID" command; if the manufacturer and model ID do not match the expected values (\$37 and \$86), or it detects that WRAM exists, the game replaces the first level with an unwinnable boss fight. *Haratyler MP* additionally provides MP3 playback from a built-in microSD card, similar to [GTMP3](../GTROM/#GTMP3).

# Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB fixed PRG ROM bank \$00
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank \$10+x
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank \$20+x
- CPU \$E000-\$FFFF: 8 KiB fixed PRG ROM bank \$30
- PPU \$0000-\$1FFF: 8 KiB switchable CHR RAM bank from 16 KiB total

# Registers {#Registers}

## <span id="MP3_Playback_Controller_.28.246000-.247FFF.29"></span>MP3 Playback Controller (\$6000-\$7FFF) {#MP3_Playback_Controller_.28.246000-.247FFF.29}

Sending a byte to the MP3 Playback Controller:

- Initiate the transfer by reading from \$6003 and wait for 1000 CPU cycles
- Send eight bits, MSB-\>LSB, by sending a 1 bit by \*not\* reading from \$6003, a 0 bit by \*reading\* from \$6003 \*and\* reading a status byte, and in both cases, by waiting for 1000 CPU cycles
- Terminate the transfer by one final read from \$6003.

The following commands are used:

- \$00: Return eight status bytes. Only the upper 4 bits are used; the lower four bits are always 0.
  - 1st: Unknown, ignored
  - 2nd: Must be \$A, checked by Haratyler MP
  - 3rd: Must be \$2, checked by Haratyler MP
  - 4th: Current equalizer setting (0-5)
  - 5th: Current volume (0-30) bits 0-3
  - 6th: Current volume (0-30) bit 4
  - 7th: Current folder (0-15)
  - 8th: Unknown, ignored
- \$01-\$BF?: Play song 1-191?, named 001.mp3-191.mp3, from current SD card folder
- \$C0-\$DE: Set volume to 0-30
- \$E0-\$E5: Set EQ to 0-5
- \$F0-\$FC?: Set folder number minus one. The folders must be named 01, 02, ..., 13 for values \$0..\$C.
- \$FD: Unkonwn, possibly "load status bytes with actual settings"?
- \$FF: Silence/reset.

The circuit board says "FC-ARDUINO MP3 KIT", and although the technical specifications are very similar to the [DFPlayer Mini MP3 Player](https://www.dfrobot.com/product-1121.html?srsltid=AfmBOoqMGB50dDCBha6DEKBAd4pd0pcvWO6b4Pg6xrz1AjMFrwPc4Qx3), the communications protocol and commands are much simpler.

## <span id="Mirroring_Register_.28.24A000-.24BFFF.29"></span>Mirroring Register (\$A000-\$BFFF) {#Mirroring_Register_.28.24A000-.24BFFF.29}

    A~[101. .... .... ...M]
                         +- 0: Vertical, 1: Horizontal

## <span id="IRQ_Register_.28.24C000-.24DFFF.29"></span>IRQ Register (\$C000-\$DFFF) {#IRQ_Register_.28.24C000-.24DFFF.29}

    A~[110. .... VVVV VVVV]

A write to this register amounts to the following writes on a normal MMC3:

- \$C000=VVVVVVVV -1
- \$C001=0
- If V=\$FF: \$E000, otherwise \$E001

## <span id="Bank_Register_.28.24E000-.24FFFF.29"></span>Bank Register (\$E000-\$FFFF) {#Bank_Register_.28.24E000-.24FFFF.29}

    A~[111. .... .... ..BA]
                        |+- PRG/CHR A13
                        +-- PRG A16

This means effectively:

    BA    CPU $A000   CPU $C000   PPU $0000
     0        $10         $20        0
     1        $11         $21        1
     2        $18         $28        0
     3        $19         $29        1

Together with the fixed banks, this means that only PRG ROM banks \$00/\$10/\$11/\$18/\$19/\$20/\$21/\$28/\$29/\$30 are accessible. The rest serve as padding for the flash ROM chip's 64 KiB sector size.
