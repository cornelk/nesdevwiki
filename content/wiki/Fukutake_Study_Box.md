---
title: "Fukutake Study Box"
url: "/wiki/Fukutake_Study_Box/"
description: "The StudyBox is a tape based educational game accessory for the Famicom released by Fukutake Shoten in April 1986. There are three known hardware revisions but only the last one has turned up outside…"
wikiPage: "Fukutake_Study_Box"
wikiCategories:
  - name: "Hardware"
    url: "/category/Hardware/"
---

<div class="infobox"><div style="text-align: center"><b>StudyBox</b></div><table class="oddf0 floatright"><tbody><tr><td><b>Company</b></td><td>Fukutake Shoten </td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>BIOS PRG ROM size</b></td><td>256KiB </td></tr><tr><td><b>PRG RAM capacity</b></td><td>16KiB &amp; 32 KiB </td></tr><tr><td><b>CHR capacity</b></td><td>8KiB </td></tr><tr><td><b>Tape Length</b></td><td>~30 Minutes </td></tr><tr><td><b>Mirroring</b></td><td>4 Fixed </td></tr><tr><td><b>Bus conflicts</b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>Yes </td></tr></tbody></table></div>

The **StudyBox** is a tape based educational game accessory for the Famicom released by Fukutake Shoten in April 1986. There are three known hardware revisions but only the last one has turned up outside of Japan. This page describes the third revision: model **SBX-01**.

- **Model 1000**: First model in a form factor similar to the [Family Computer Disk System](../Family_Computer_Disk_System/). Was leased out to users and never sold.
- **Model 1500**: A single-unit model. Not much is known about it.
- **SBX-01**: The most common model. Sold to customers directly during its lifespan.

The Model 1000 and SBX-01 contain different firmware different enough that tapes for one system will not work on the other. Some packet header data differs and the script opcodes are completely different.

There are four known main sets of tapes for the StudyBox: Enjoy English, Enjoy English+, Newton Land Science, and Newton Land Arithmetic. Each set is twelve lessons/tapes long and come with a corresponding workbook. There are also a handful of auxiliary tapes that don't seem to have an associated workbook.

## Hardware {#Hardware}

The StudyBox device contains a large 256KiB PRG ROM, extra PPU RAM (2KiB) for 4 screens of nametables, 8KiB CHR RAM, two 32KiB PRG RAMs, a Hitachi HG62E08R38 gate array, a Toshiba TMP42C60P5351 microcontroller, and on the actual cassette mechanism, a Fujitsu MB88505H microcontroller.

## Banks {#Banks}

- PPU \$0000-\$1FFF: 8KiB non-banked CHR-RAM
- PPU \$2800-\$2FFF: 2KiB non-banked PPU nametable RAM
- CPU \$4400-\$4FFF: 4KiB fixed to the first bank of RAM A (see below)
- CPU \$5000-\$5FFF: 16KiB switched RAM A
- CPU \$6000-\$7FFF: 32KiB switched RAM B
- CPU \$8000-\$BFFF: 256KiB Switched ROM
- CPU \$C000-\$FFFF: Fixed ROM bank 0

### <span id="Partial_Bank_at_.244400"></span>Partial Bank at \$4400 {#Partial_Bank_at_.244400}

The partial RAM A bank starting at \$4400 is aligned to \$4000. In other words, address \$4400 is the same as \$5400 if the \$5000 range has the first bank loaded. The data that would be mapped to \$4000-\$43FF is only accessible when mapped to the \$5000-\$5FFF range.

## Registers {#Registers}

### <span id="Data_From_Tape_.28.244200_Read.29"></span>Data From Tape (\$4200 Read) {#Data_From_Tape_.28.244200_Read.29}

       7  bit  0
       ---------
       DDDD DDDD
       |||| ||||
       ++++-++++- Data

An IRQ is fired when data is ready to be read from this register.

### <span id="Tape_Read_Status_.28.244201_Read.29"></span>Tape Read Status (\$4201 Read) {#Tape_Read_Status_.28.244201_Read.29}

       7  bit  0
       ---------
       D?S? xxxx
       |||| ||||
       |||| ++++- Unused
       |||+------ Unknown
       ||+------- Data region status
       |+-------- Unknown
       +--------- Decoder disabled

### <span id="Tape_Drive_Status_.28.244202_Read.29"></span>Tape Drive Status (\$4202 Read) {#Tape_Drive_Status_.28.244202_Read.29}

       7  bit  0
       ---------
       xRxx Pxxx
       |||| ||||
       |||| |+++- Unused
       |||| +---- Power Supply not connected
       ||++------ Unused
       |+-------- Shift ready for next bit / tape ready?
       +--------- Unused

The power supply bit is used, however it doesn't trip when the unit is unplugged. This is checked during the reset routine and will display an error message if set. Amusingly, the error message will display an image of the rear of the system that does not match the hardware.

### <span id="RAM_Page_Select_.28.244200_Write.29"></span>RAM Page Select (\$4200 Write) {#RAM_Page_Select_.28.244200_Write.29}

       7  bit  0
       ---------
       BBxx xxAA
       |||| ||||
       |||| ||++- Select RAM bank mapped to $5000-$5FFF
       ||++-++--- Unused
       ++-------- Select RAM bank mapped to $6000-$7FFF

### <span id="ROM_Page_Select_.28.244201_Write.29"></span>ROM Page Select (\$4201 Write) {#ROM_Page_Select_.28.244201_Write.29}

       7  bit  0
       ---------
       xxxx CCCC
       |||| ||||
       |||| ++++- Select ROM bank mapped to $8000-$BFFF
       ++++------ Unused

### <span id="Tape_Drive_Control_.28.244202_Write.29"></span>Tape Drive Control (\$4202 Write) {#Tape_Drive_Control_.28.244202_Write.29}

       7  bit  0
       ---------
       O?RC ?AID
       |||| ||||
       |||| |||+- Data decoding enable??
       |||| ||+-- IRQ enable
       |||| |+--- Tape audio enable
       |||| +---- Unknown
       |||+------ Pulse low to clock data bit
       ||+------- Pulse low to reset drive controller
       |+-------- Unknown
       +--------- Output data bit

#### Observed Commands {#Observed_Commands}

- \$00: Seek backwards to start of current page?
- \$01-\$07: Skip forward *N* pages
- \$41-\$47: Skip backward *N*-\$40 pages
- \$FF: ?? (skip to end of data block??)

## Tapes {#Tapes}

Programs were distributed on standard audio cassette tapes. These tapes are stereo, containing data on the right channel and recorded audio on the left. The tapes use a 4800 bits-per-second [Modified Frequency Modulation (MFM)](https://en.wikipedia.org/wiki/Modified_frequency_modulation) encoding.

There are no external physical controls for the tape drive. The drive is controlled entirely in software.

### Data Format {#Data_Format}

Data on the tapes are grouped into pages, each page consists of segments, and each segment is split into packets. There are six types of segments:

- Header
- Delay
- CHR Data
- Nametable/Sprite Data
- Script Data
- Padding

Header, delay, and padding segments are all a single packet. The header contains a non-unique, zero-based, Page number and a checksum. The page number in the header corresponds directly to the one-based page number entered on the title screen. Delay and padding segments are both one large packet and are used to synchronize the tape data with the audio track. Any time audio is playing the tape is either loading data or reading either a delay or padding packet.

CHR, Nametable, and Script data segments are split into 128-byte packets. Each packet contains a simple XOR checksum as the last byte of data.

### Script {#Script}

Each tape program is controlled by a script. The script a domain specific stack-based language. Reverse engineering work is still ongoing.

## See Also {#See_Also}

- [StudyBox file format](../StudyBox_file_format/)
- <http://cah4e3.shedevr.org.ru/dumping_2005.php>
- <https://github.com/TASEmulators/fceux/blob/master/src/boards/186.cpp>
- <https://tcrf.net/StudyBox>
- <https://github.com/SourMesen/Mesen/files/2920256/studybox.org.txt>
