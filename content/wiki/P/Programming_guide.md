---
title: "Programming guide"
url: "/wiki/Programming_guide/"
description: "This programming guide is the complement to the NES reference guide. You will find more in-depth documentation about specific topics, tutorials on the most common problems encountered, a list of…"
wikiPage: "Programming_guide"
---

This programming guide is the complement to the [NES reference guide](../NES_reference_guide/). You will find more in-depth documentation about specific topics, tutorials on the most common problems encountered, a list of existing tools that will assist you during you development phase and more. If you're new to NES development, you should give a look to the "[getting started](#Getting_started)" section first.

### Getting started {#Getting_started}

- [Before the basics](../Before_the_basics/)
- [Commissioning](../Commissioning/) - suggestions for a non-programmer seeking help
- [Programming Basics](../Programming_Basics/)
- [Installing CC65](../Installing_CC65/)

### General reference {#General_reference}

- [Registers](../Registers/)
- [Waiting for the PPU to power on](../PPU_power_up_state/#Best_practice)
- [CHR-ROM vs CHR-RAM](../CHR_ROM_vs._CHR_RAM/) and how to switch an NROM project from the former to the latter
- [Emulation Libraries](../Emulation_libraries/)
- [Catch-up](../Catch-up/) technique used by emulators
- [Flash Cartridges](../../category/Flash_Cartridge/) ([PowerPak](../PowerPak/) / [Everdrive N8](../Everdrive_N8/) / [Everdrive N8 Pro](../Everdrive_N8_Pro/) / [KrzysioCart](../KrzysioCart/))
- [Music](../Music/): [Expansion audio](../../category/Expansion_audio/)
- [Tools](../Tools/) — includes assemblers, disassemblers, graphics, music/audio, and emulator-related utilities

### Tutorials {#Tutorials}

- [CPU basics](../CPU/)
- [PPU reference](../PPU/)
- [APU basics](../APU_basics/)
- [Compression](../Compression/)
  - [Fixed Bit Length Encoding](../Fixed_Bit_Length_Encoding/)
- [Programming mappers](../Programming_mappers/)
  - [NROM](../Programming_NROM/)
  - [UxROM](../Programming_UNROM/) and similar
  - [MMC1](../Programming_MMC1/)
  - [MMC3](../Programming_MMC3/)
- Programming Techniques
  - [6502 assembly optimisations](../6502_assembly_optimisations/)
  - [Cycle counting](../Cycle_counting/)
  - [RTS Trick](../RTS_Trick/)
  - [Jump Table](../Jump_table/)
  - [Scanning Tables](../Scanning_tables/)
  - [Scanning Large Tables](../Scanning_large_tables/)
  - [Synthetic Instructions](../Synthetic_instructions/)
  - [Programming with unofficial opcodes](../Programming_with_unofficial_opcodes/)
  - [Pointer tables](../Pointer_table/)
  - [Multibyte constant](../Multibyte_constant/)
- [Limitations](../Limitations/)
- [Emulation Tutorials](../Emulation_tutorials/)
- [Sample RAM map](../Sample_RAM_map/)
- [PPU scrolling](../PPU_scrolling/)
  - How to scroll an infinite map
  - How to scroll with less artifact
  - How to split the screen for a status bar
  - How to do parallax scrolling with the scroll register
- [Sprite size](../Sprite_size/)
- [How to make raster effects](http://story.monoroch.net/nes_raster.html) by Bregalad
- [The frame and NMIs](../The_frame_and_NMIs/) by Disch
  - [NMI thread](../NMI_thread/): Make your status bar rock-solid
- [Nerdy Nights sound](../Nerdy_Nights_sound/)
- [Palette change mid frame](../Palette_change_mid_frame/)
- [Don't hardcode OAM addresses](../Don%27t_hardcode_OAM_addresses/)
- [Interrupt forwarding](../Interrupt_forwarding/)
- [Sprite cel streaming](../Sprite_cel_streaming/)
- [Drawing terrain](../Drawing_terrain/)
- [Releasing on modern platforms](../Releasing_on_modern_platforms/)

### Libraries and samples {#Libraries_and_samples}

#### System {#System}

- [Init code](../Init_code/)

#### Input {#Input}

- [Controller reading](../Controller_reading/)
- [Controller reading code](../Controller_reading_code/)

#### Graphics {#Graphics}

- [Detecting video standard](../Detect_TV_system/)
- [Placeholder graphics](../Placeholder_graphics/)
- [RLE decompression code for CHR](http://forums.nesdev.org/viewtopic.php?p=142703#p142703) by koitsu

#### Math {#Math}

- [Multiplication by a constant integer](../Multiplication_by_a_constant_integer/)
- [Fast signed multiply](../Fast_signed_multiply/) - Fast table driven multiplication
- [8-bit Multiply](../8-bit_Multiply/) - Multiplies two 8-bit integers to a 16-bit result
- [8-bit Divide](../8-bit_Divide/) - Divide two 8-bit integers to a 16-bit result
- [Division by a constant integer](../Division_by_a_constant_integer/)
- [Divide by 3](../Divide_by_3/)
- [16-bit BCD](../16-bit_BCD/) - An efficient 16-bit binary to decimal converter
- [Base 100](../Base_100/) - An alternate method of storing numbers, to simplify BCD conversion
- [Random number generator](../Random_number_generator/)

#### Audio {#Audio}

- [APU period table](../APU_period_table/)
- [Audio drivers](../Audio_drivers/)
- [Nerdtracker player in NESASM](../Nerdtracker_player_in_NESASM/)

#### Misc {#Misc}

- [Skipping the FDS license screen](https://forums.nesdev.org/viewtopic.php?p=194826#p194826)
- [CRC32 checksum calculation](../Calculate_CRC32/)

#### Examples {#Examples}

- [Projects](../Projects/) - includes many useful open-source examples.
- [Nerdy Nights Out (aka: Learning 6502)](https://nerdy-nights.nes.science/) by BunnyBoy
- [Programming NES Games in C](https://shiru.untergrund.net/articles/programming_nes_games_in_c.htm) by Shiru
- [How to Program an NES game in C](https://nesdoug.com/) by dougeff
- [Simple game and editing guide (in C)](https://nes-starter-kit.nes.science)
- [Project template](https://pineight.com/nes/#template) by tepples
- [Minimal NES example using ca65](https://forums.nesdev.org/viewtopic.php?t=11151) by rainwarrior

### Reverse engineered techniques {#Reverse_engineered_techniques}

- [Bad Apple](../Bad_Apple/) video compression
- [Codemasters tile compression](https://forums.nesdev.org/viewtopic.php?t=5860)
- [Big Bird's Hide and Speak sample compression](https://forums.nesdev.org/viewtopic.php?t=8675)
- [Battletoads text compression](https://forums.nesdev.org/viewtopic.php?t=8609) (Huffman coding)
- [Rad Racer rendering](https://forums.nesdev.org/viewtopic.php?t=8588)

### Useful reference materials {#Useful_reference_materials}

#### Books {#Books}

- Cady, Fredrick M. (1997). *Microcontrollers and Microcomputers: Principles of Software and Hardware Engineering*. New York: Oxford University Press. ISBN 0-19-511008-0.
- Eyes, David; Lichty, Ron (1986). *Programming the 65816 including the 6502, 65C02, and 65802*. New York: Brady Books/Prentice Hall Press (Simon & Schuster, Inc). ISBN 0-89303-789-3. Rights later purchased by Western Design Center.
- Fernandez, Judi N. (1983). *6502 Assembly Language Programming (Self-teaching Guides)*. John Wiley & Sons Inc. ISBN 978-0471861201.
- Leventhal, Lance A. (1986). *6502 Assembly Language Programming 2nd Edition*. Osborne/McGraw-Hill. ISBN 0-07-881216-X.
- Zaks, Rodnay (1983). *Programming the 6502 (Fourth Edition)*. Sybex, Inc. ISBN 0-89588-135-7.
- Zaks, Rodnay (1982). *Advanced 6502 Programming*. Sybex, Inc. ISBN 0-89588-089-X.
- Zaks, Rodnay (1980). *Programming the 6502 (Third Edition)*. Sybex, Inc. ISBN 0-89588-046-6.

#### Online {#Online}

- [Programming the 65816 (including the 6502, 65C02, and 65802)](http://wdc65xx.com/Programming-Manual/) by Western Design Center
  1.  [Programmingthe65816_ProgManual.pdf](http://web.archive.org/web/20170321153712id_/http://www.westerndesigncenter.com/wdc/documentation/Programmingthe65816_ProgManual.pdf) — 54MBytes, created 2015/03/27.
      - This is a full scan of the original Eyes/Lichty book using full images for each page, with the added bonus of OCR being applied so that text is searchable and copy-paste-able. Accurate/reliable given the nature of the scan.
        - [backup copy](https://drive.google.com/file/d/1cMzhGjAZPkxlkJxS7c6cRtRO53KC9v1s/view?usp=sharing) via Koitsu (primary)
        - [backup copy](https://jdc.koitsu.org/lj/Programmingthe65816_ProgManual.pdf) via Koitsu (secondary)
  2.  [Programmanual.pdf](https://www.nesdev.org/w/images/default/7/76/Programmanual.pdf) (a.k.a. WDC65C816_Program_Manual.pdf) — 1.7MBytes, created 2007/04/16.
      - This is a "pure text" version of the Eyes/Lichty book, however it is rife with very dangerous typos (wrong opcodes, etc.) in many areas, and contains a large number of formatting/layout mistakes not present in the original book. Bill Mensch (6502/65C02/65816 architect/designer) [confirmed these mistakes](http://forums.nesdev.org/viewtopic.php?p=142722#p142722), stating WDC is in the process of fixing them. If at all possible, use the newer PDF mentioned above.
      - It is suspected that WDC has pulled this PDF entirely from [their website](http://65xx.com/Products/Programming_and_Hardware-Manuals/), instead advocating purchasing a "redone" version of the Lichty/Eyes book via Amazon. Kindle and paperback versions are available, but I (Koitsu) have not reviewed it.

### Links {#Links}

- [The infamous 6502.txt](http://www.zophar.net/documents/6502/6502-txt.html): allegedly complete info on the 6502 instructions
- [Machine Language for Beginners](http://www.atariarchives.org/mlb/): a 1983 introduction to 6502 assembly language
