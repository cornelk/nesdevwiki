---
title: "Turbo File"
url: "/wiki/Turbo_File/"
description: "The Turbo File (for Family Computer) is a Famicom expansion port peripheral developed by ASCII Corporation, and is used to store save data for certain games. Variants of the device were also produced…"
wikiPage: "Turbo_File"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
  - name: "Hardware"
    url: "/category/Hardware/"
---

The **Turbo File** (for [Family Computer](../Family_Computer/)) is a Famicom [expansion port](../Expansion_port/) peripheral developed by ASCII Corporation, and is used to store save data for certain games. Variants of the device were also produced for the Game Boy (Turbo File GB), Game Boy Advance (Turbo File Advance), and Super Famicom ([Turbo File Twin](https://snes.nesdev.org/wiki/Turbo_File_Twin)). Original Famicom Turbo File devices can be connected to a Super Famicom using an adapter unit, allowing certain SFC games to use the device in lieu of the Turbo File Twin.

The device exists in two iterations:

- The original model (AS-TFO2) containing 8K of battery-backed SRAM. This has two variants which connect to either the Famicom expansion port or the [ASCII Stick II Turbo's "OPTION" port](https://forums.nesdev.org/viewtopic.php?t=26282).
- The **Turbo File II** (AS-TF21) containing 32K of battery-backed SRAM (HM62256ALP-15), which is split into four 8K banks that can be manually switched to allow for four individual game saves.
  - The TFII also includes a proprietary cartridge slot which accepts "データROM" (Data ROM) packs. These packs have a switch that toggles between accessing the contents of the ROM (in which case the bank select switch is used to move between four 8K ROM banks) or the TFII's save RAM. Only two such packs were ever sold through direct mail order; both of which contain team data for games in the *Best Play Pro Yakyuu* series. ('89-Apr. and '90-Jun.)

Both devices have an LED that illuminates when connected to a Famicom that is powered on, as well as a write protect switch to prevent accidentally overwriting save data.

## Supported titles {#Supported_titles}

Famicom games that support the Turbo File include:

- *Best Keiba - Derby Stallion* (1991)
- *Best Play Pro Yakyuu* (1988)
- *Best Play Pro Yakyuu '90* (1990)
- *Best Play Pro Yakyuu II* (1990)
- *Best Play Pro Yakyuu Shin Data* (1988)
- *Best Play Pro Yakyuu Special* (1992)
- *Castle Excellent* (1986) - also supports the [Family BASIC Data Recorder](../Family_BASIC_Data_Recorder/). (Functionality removed from *Castlequest*, the US version)
- *Derby Stallion - Zenkoku Ban* (1992)
- *Downtown - Nekketsu Monogatari* (1989) & *Kunio 8-in-1* (Pirate multicart that contains the game as "HEROS STORY")
- *Dungeon Kid* (1990)
- *Famicom Shougi - Ryuuousen* (1991)
- *Fleet Commander* (1988)
- *Haja no Fuuin* (1987)
- *Itadaki Street - Watashi no Mise ni Yottette* (1990)
- *The Money Game 2 - Kabutochou no Kiseki* (1989)
- *Ninjara Hoi!* (1990)
- *Wizardry - Legacy of Llylgamyn* (1989)
- *Wizardry - Proving Grounds of the Mad Overlord* (1987)
- *Wizardry - The Knight of Diamonds* (1991)

Most titles bear a "TF" logo on the cartridge label signifying their compatibility.

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244017_Read.29"></span>Input (\$4017 Read) {#Input_.28.244017_Read.29}

    7  bit  0
    ---- ----
    xxxx xDxx
          |
          +---- Data input

### <span id="Output_.28.244016_Write.29"></span>Output (\$4016 Write) {#Output_.28.244016_Write.29}

    7  bit  0
    ---- ----
    xxxx xCRD
          |||
          ||+-- Data output
          |+--- 0: Reset to address $0000
          +---- Clock line (1: high, 0: low)

To reset to address \$0000:

1.  Write \$00 to \$4016.
2.  Write \$02 to \$4016.

To read a bit (and write it back unchanged):

1.  Read the old data from \$4017.D2: `old = ([$4017] & $04) >> 2`
2.  Write `$06 | old` to \$4016 (clock high).
3.  Write `$02 | old` to \$4016 (clock low).

To write a new bit:

1.  Write `$06 | new` to \$4016 (clock high).
2.  Write `$02 | new` to \$4016 (clock low).

The address automatically increments after 8 bits/clocks. To read/write a given address, the program must reset the address to \$0000 and seek to the desired address by performing dummy reads.

## Memory Setup and File Format {#Memory_Setup_and_File_Format}

The first byte of memory (offset \$0000) is unused, potentially out of fear that certain games with controller access might disrupt it, so a dummy byte is used to skip it after resetting the address. The rest of the space (\$0001-\$1FFF) is used to store save data, with files attached directly after each other. Invalid file IDs indicate the start of free memory.

The majority of Turbo File games utilize the following format for save files. Data is transferred **LSB** first:

    2   ID "AB" ($41,$42)
    2   Filesize (16+N+2) (including title and checksum)
    16  Title in ASCII (terminated by $00 or $01)
    N   Data Portion
    2   Checksum (all N bytes in Data Portion added together)

The exception to this, *Castle Excellent*, forgoes a filename and allocates a hardcoded size of 512 bytes (\$0000-\$01FF). Data is transferred **MSB** first:

    1   Don't care ;fixed at offset $0000
    2   ID $AA,$55 ;fixed at offset $0001..$0002
    509 Data Portion (Data, end code "BEDEUTUN", followed by some unused bytes)

Due to the bit-reversal, interpreting the above data using the "standard" format would result in \$??, \$55, \$AA, and so on.

Due to the hardcoded memory offset, *Castle Excellent* will destroy any other file that is located at the same address. Some later games for both the Famicom and Super Famicom (such as *Fleet Commander* and *Wizardry V*) are able to properly manage the *Castle Excellent* save file.

### File Deletion Support {#File_Deletion_Support}

File deletion features are uncommon in Turbo File games, primarily due to a lack of system RAM to perform file relocation. The following games have limited or full support for file deletion:

- *Derby Stallion - Zenkoku Ban* allows erasing files from the end of the filesystem to free up space for its save data. (Note: cannot handle *Castle Excellent* data)
- *Wizardry V - Heart of the Maelstrom* (SFC, JP version only) supports importing party members from files saved by the FC *Wizardry* games (via the adapter), and allows deleting individual files.

## See Also {#See_Also}

- [FDS-TurboFile-Dumper](https://github.com/TakuikaNinja/FDS-TurboFile-Dumper) - FDS program to dump/rewrite Turbo File contents via disk.
- [GitHub Gist](https://gist.github.com/TakuikaNinja/64ba846b38d211edc6c785e5b64406a6) - Python script for listing files within Turbo File saves.
- [Turbo File Twin](https://snes.nesdev.org/wiki/Turbo_File_Twin) - Super Famicom variant. Has a TFII mode which simulates a Turbo File II connected to an adapter.

## References {#References}

- nocash's NES documentation: <https://problemkaputt.de/everynes.htm#storageturbofile>
- enri's Famicom documentation: <http://cmpslv3.stars.ne.jp/Famic/Famic.htm> (scroll down to ターボファイル )
