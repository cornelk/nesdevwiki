---
title: "NES reference guide"
url: "/wiki/NES_reference_guide/"
description: "This page lists all major hardware reference divided by categories in their simplest form. From this list, you can drill down to a more specific section of the selected category."
wikiPage: "NES_reference_guide"
---

This page lists all major hardware reference divided by categories in their simplest form. From this list, you can drill down to a more specific section of the selected category.

### General {#General}

- [Glossary](../Glossary/)
- [Cartridge and mappers' history](../Cartridge_and_mappers%27_history/)

### Hardware reference {#Hardware_reference}

- [2A03](../2A03/): [APU](../APU/), [CPU](../CPU/)
  - [Instruction reference](../Instruction_reference/)
- [PPU](../PPU/)
- [Input devices](../Input_devices/)
- [Mappers](../Mapper/)
- [Pinout](../Hardware_pinout/)
- [CE Decoder](../CE_Decoder/)
  - ~~Example flow~~
  - ~~Another example~~
- [Cycle reference chart](../Cycle_reference_chart/) for various PPU and frame-related timing details; includes CPU cycle counts
- [Cartridge board](../Cartridge_board_reference/)
- [RF Famicom wiring diagram](../../wiki-images/Neswires.jpg) (many parts also applicable to NES)
- [Errata](../Errata/): Hardware bugs and quirks
- [Famicom cartridge dimensions](../Famicom_cartridge_dimensions/)
- Console5 Tech wiki has schematics and information: [Famicom](https://wiki.console5.com/wiki/Famicom), [NES-001](https://wiki.console5.com/wiki/Nintendo_NES-001), [NES-101](https://wiki.console5.com/wiki/Nintendo_NES-101)

#### Miscellaneous {#Miscellaneous}

- [Family Computer](../Family_Computer/) - the original Japanese console
- [Family Computer Disk System](../Family_Computer_Disk_System/)
  - [FDS BIOS](../FDS_BIOS/)
  - [FDS disk format](../FDS_disk_format/)
- [Vs. System](../Vs._System/)
- [FamicomBox](../FamicomBox/)
- [CIC lockout chip](../CIC_lockout_chip/)
- [Myths](../Myths/)
- [PRG RAM circuit](../PRG_RAM_circuit/)
- [Implementing Mappers In Hardware](../Implementing_Mappers_In_Hardware/)
- [Tutorial on reading circuits in Visual 6502/2A03/2C02](../Visual_circuit_tutorial/)
- [RF modulator board](../MOD_RF/)
- [Nintendo NES-001](http://console5.com/wiki/Nintendo_NES-001) on Console5
- [Demo Vision](../Demo_Vision/)
- [Fonts](../Fonts/) - fonts used on official devices.
- [List of games with significant regional differences](../List_of_games_with_significant_regional_differences/)
- [Nintendo header](../Nintendo_header/) - a metadata header found in ~33% of licensed NES games.
- [Enri's Homepage: Famicom](http://cmpslv3.stars.ne.jp/Famic/Famic.htm) (Japanese) - transcription of some official Famicom documentation, and other notes.
- [NES Music Ripping Guide](https://nesdev.org/nesaudio.zip) by Chris Covell
- [Battery holder](../Battery_holder/) - guide for adding a battery holder to an NES cartridge
- [NES Classic Controller for Wii](../NES_Classic_Controller_for_Wii/) - guide for adapting an NES controller to Wii
- [Cartridge dumper and flasher devices](../../category/Dumper_and_Flasher_devices/)

### File format reference {#File_format_reference}

- [iNES](../INES/)
  - [NES 2.0](../NES_2.0/)
- [UNIF](../UNIF/)
- [NSF](../NSF/)
- [FDS](../FDS_file_format/)
- [QD](../QD/)
- IPS
  - [description](https://datacrystal.romhacking.net/wiki/Patch)
  - [file format specification](http://zerosoft.zophar.net/ips.htm)
  - [limitations and pitfalls](http://justsolve.archiveteam.org/wiki/IPS_(binary_patch_format))
  - [caveats and extensions](https://web.archive.org/web/20120310064326/http://romhack.wikia.com/wiki/IPS)

### Emulation reference {#Emulation_reference}

- List of NES [emulators](../Emulators/)
- [Emulator tests](../Emulator_tests/)
- [Game bugs](../Game_bugs/) - games that display buggy behavior on the actual hardware
- [Tricky-to-emulate games](../Tricky-to-emulate_games/)
- [Sprite overflow games](../Sprite_overflow_games/) - games which use the [sprite overflow bug](../PPU_sprite_evaluation/#Sprite_overflow_bug) of [OAM](../PPU_OAM/)
- [Colour-emphasis games](../Colour-emphasis_games/) - games which make use of the colour emphasis bits of \$2001
- [Colour \$0D games](../Color_$0D_games/) - games which use the infra-black colour \$0D.
- [Expansion audio games](../List_of_games_with_expansion_audio/) - Famicom games that use extra audio hardware.

### Notes {#Notes}

- All content refer to the NTSC system unless otherwise specified

### External links {#External_links}

- [NESdev.org archive](https://www.nesdev.org/archive.html) - Archive of the old NESdev webpage with many historical documents and links.
- [NES-001 (frontloader)](http://console5.com/wiki/Nintendo_NES-001) and [NES-101 (toploader)](http://console5.com/wiki/Nintendo_NES-101) Control Deck PCB maps on Console5 TechWiki
