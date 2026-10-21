---
title: "Cartridge and mappers' history"
url: "/wiki/Cartridge_and_mappers'_history/"
description: "Note: this list is not a citeable source for release dates; all dates either come from NEScartDB or Wikipedia."
wikiPage: "Cartridge_and_mappers__history"
---

## 1983 {#1983}

1983 July 15th  
The Nintendo Family Computer ([Famicom](../Family_Computer/)) is released in Japan. The three launch games are *Donkey Kong*, *Donkey Kong Jr.* and *Popeye*. During the first year, only Nintendo released games for the system. All of them had 16 KiB PRG ROM and 8 KiB CHR ROM.

## 1984 {#1984}

1984 June 21th  
*Family BASIC* is released. It is the first special cart Nintendo made for the system: instead of being a game it allows the user to program the system using the BASIC programming language. It's also the first cart that uses 32 KiB of PRG. Additionally it provides 2 KiB of battery-backed PRG RAM. It is the first cart (for a long while) to use either PRG RAM or battery backup.

1984 July 20th  
Hudson's *Lode Runner* and *Nuts & Milk* are the first third-party games released for the Famicom in Japan. The cartridges were manufactured by Nintendo.

1984 November 2nd/8th  
Namco's *Pac Man* and *Xevious* appear to be the first releases for the Famicom that were entirely manufactured by a third party. *Xevious* appears to be the first 32 KiB PRG game for the system (as *Family BASIC* isn't a game). Because of the lack of any lockout chip, Namco and many other companies (Bandai, Irem, Jaleco, Konami, Sunsoft, Taito) will follow this trend and make their own cartridges instead of having Nintendo make them for them.

## 1985 {#1985}

1985 September 13th  
Nintendo's *Super Mario Bros.* is released in Japan. It is one of the first (if not the first) Famicom games that is not merely a port of another arcade or computer game. This will be a major factor in the console's success.

1985 September 27th  
Jaleco released *City Connection*, using [some circuitry](../INES_Mapper_087/) to allow 16 KiB of CHR ROM, switching between two 8 KiB banks. This is the first Famicom game to use hardware other than ROMs (not counting *Family BASIC*). Mappers were born.

1985 October 18th  
The Nintendo Entertainment System (NES) is released in the United States. The launch games are *10-Yard Fight*, *Baseball*, *Clu Clu Land*, *Donkey Kong Jr. Math*, *Duck Hunt*, *Excitebike*, *Golf*, *Gyromite*, *Hogan's Alley*, *Ice Climber*, *Kung Fu*, *Mach Rider*, *Pinball*, *Stack-Up*, *Super Mario Bros.*, *Tennis*, *Wild Gunman*, and *Wrecking Crew*. Some of these contained 60-pin Famicom cartridge PCBs connected to an internal Famicom-to-NES adapter, while others used 72-pin cartridge PCBs designed specifically for the NES. This list already includes some third-party games but this time the [lockout chip](../CIC_lockout_chip/) inside the NES prevented them from building their own cartridges.

1985 November 16th  
Konami's *Hyper Olympic Gentaiban!*, a limited-release sequel to a game released in June, is the second game to use the CHR ROM switch circuit first seen in *City Connection*.

## 1986 {#1986}

1986 February 21st  
Nintendo released the [Famicom Disk System](../Family_Computer_Disk_System/) (FDS) accessory, which vastly improved the capabilities of the console (as almost no games were yet using mappers). One disk can hold 128 KB of data (64 KB per side), and was cheaper to produce. Additionally it provided extra sound and the disks could hold save data.

1986 April  
Bandai, Konami and Nintendo appear to start producing [CNROM](../CNROM/) games almost simultaneously. It's a cost-saving improvement over Jaleco's circuit that uses only one chip instead of two, at the cost of having [bus conflicts](../Bus_conflict/).

1986 April 17th  
Sunsoft's *Atlantis no Nazo* appears to be the first cartridge that uses a [dedicated mapper chip](../INES_Mapper_184/) (as opposed to 74 series logic chips). Other companies follow this trend promptly.

1986 June 6th  
Nintendo's *Gumshoe* (released in USA) is the first [GNROM](../GxROM/) game. It's the first game to bank PRG ROM, have 128 KiB of PRG ROM, and to support banking both PRG and CHR.

1986 June 13th  
Capcom's *Maikamura* (a.k.a. *Ghosts 'n' Goblins*) appears to be the first [UNROM](../UxROM/) game. It is the first cartridge game to use CHR RAM.

1986 June 20th  
Namco's *Super Chinese* appears to be the first game that uses a complex [ASIC](../../category/ASIC_mappers/) mapper. [Namco's 108](../INES_Mapper_206/) is the MMC3's predecessor.

1986 July 30th  
Konami's [VRC1](../VRC1/) (the second ASIC mapper) is first used in *Ganbare Goemon! Karakuri Douchuu*. After the FDS, the VRC1 is the first mapper that supports changing the nametable layout ("[mirroring](../Mirroring/)") at runtime.

1986 September 1st  
The NES is released in PAL-B regions.

1986 December 15th  
Irem's *Mashou* is the first game published to use a [BNROM](../INES_Mapper_034/) board. It (along with the Western localized version *Deadly Towers*) would turn out to be the only licensed game on that board.

## 1987 {#1987}

1987 April 14th  
Seta's *Morita Shougi* appears to be the first game to use the Nintendo [MMC1](../MMC1/) mapper. It's also the first game cartridge with 8 KiB of PRG RAM and battery backup.

1987 May 15th  
The NES is released in PAL-A regions.

1987 June 30th  
Taito's [X1-005](../Taito_X1-005/) is first used in *Mirai Shinwa Jarvas*. The X1-005 is the first mapper to provide a small amount of extra RAM on-die, which was almost always used for saving.

1987 July 7th  
Konami's [VRC2](../VRC2_and_VRC4/) is first used in *Getsufuu Maden*. The VRC2 is the first mapper to support eight independent CHR banks.

1987 September 25th  
Konami's [VRC3](../VRC3/) is first and only used in *Salamander*. After the FDS, the VRC3 is the first mapper that can generate [IRQs](../IRQ/).

1987 September  
Nintendo's [MMC2](../MMC2/) is used in *Mike Tyson's Punch-out!!*. It's the first mapper to support hardware trickery to bypass the 256-tile background limitation.

1987 December  
Rare's *Wizards & Warriors* is the first game published to use an [ANROM](../AxROM/) board.

1987 December 4th  
Namco's *Star Wars* is the first game published to use their [129](../INES_Mapper_019/) mapper IC, which supported ROM nametables and expansion audio. (However, this game used neither.)

## 1988 {#1988}

1988 March 18th  
Irem released *Napoleon Senki*, the first game to have [four screens of nametables](../Mirroring/#4-Screen), and to contain both CHR RAM and CHR ROM.

1988 August 12th  
Namco's *Final Lap* is the first cartridge game (i.e. not FDS) to use [expansion audio](../../category/Expansion_audio/), via its [Namco 163](../Namco_163_audio/) mapper.

1988 September 27th  
Seta's *8 Eyes* appears to be the first game to use the Nintendo [MMC3](../MMC3/) mapper.

## 1990 {#1990}

1990 February 3rd  
Koei's *Nobunaga no Yabou: Sengoku Gunyuuden* (Japanese version of *Nobunaga's Ambition II*) was the first game to use the Nintendo [MMC5](../MMC5/) mapper. However, the writing "(c) 1989 Nintendo" and the date code on MMC5 chips suggest the chip already being available in late 1989.

1990 April  
Rare's *[Pin•Bot](../INES_Mapper_119/)* is the first game that can bankswitch CHR-RAM and CHR-ROM independently.

1990 August 11th  
Bandai's *[SD Gundam Gaiden: Knight Gundam Monogatari](../INES_Mapper_159/)* is the first game to save games in EEPROM.

## 1991 {#1991}

1991 August 30th  
HAL Laboratory released *Metal Slader Glory*, the largest licensed game with 512KB of PRG ROM and CHR ROM.

## 1994 {#1994}

1994 June 24th  
Hudson Soft's *Takahasi Meijin no Boukenjima (Adventure Islands) IV* is the last licensed game released for the Famicom.

1994 December  
Nintendo's *Wario's Woods* is the last licensed game for the NES in USA, and the only NES game to receive an [ESRB](https://en.wikipedia.org/wiki/Entertainment_Software_Ratings_Board) rating.

## 1995 {#1995}

1995 May 25  
*The Lion King* (PAL only) is the last licensed game for the NES.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

## 2003 {#2003}

2003 September 25th  
Nintendo discontinued all manufacturing and support for the Famicom.

**Note:** this list is not a citeable source for release dates; all dates either come from [NEScartDB](https://nescartdb.com) or Wikipedia.

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [The Lion King (video game)](https://en.wikipedia.org/wiki/The_Lion_King_(video_game)) – Wikipedia
