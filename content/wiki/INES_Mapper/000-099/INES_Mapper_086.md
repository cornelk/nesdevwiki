---
title: "INES Mapper 086"
url: "/wiki/INES_Mapper_086/"
description: "iNES Mapper 086 designates Jaleco's JF-13 board used for the Red and Black releases of Moero!! Pro Yakyuu, the Japanese version of Jaleco's Bases Loaded."
wikiPage: "INES_Mapper_086"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

**iNES Mapper 086** designates Jaleco's JF-13 board used for the Red and Black releases of *Moero!! Pro Yakyuu*, the Japanese version of Jaleco's *Bases Loaded*.

## Hardware {#Hardware}

This board has 128 KiB PRG ROM, 64 KiB CHR ROM, no PRG RAM (the mapper ports are mapped in \$6000-\$7FFF instead, but that means no bus conflicts either), a [74HC139](../74139/), two [74HC174s](../74174/), and an additional ASIC "D7756C 148" by NEC that plays recorded speech compressed with ADPCM. An additional set of sample files is needed to emulate the sound chip.

This board is subject to the same design flaw as [iNES Mapper 038](../INES_Mapper_038/): accidental copies of the registers are present at \$E000-\$FFFF because /ROMSEL is high when M2 is low, not just when A15 is low.

Half the 74139 decodes when R/W is low and A14 and /ROMSEL are high. That then chains into the other half, which decodes writes to \$6xxx (and \$Exxx) and \$7xxx (and \$Fxxx).

## Registers {#Registers}

### <span id="PRG_and_CHR_bank_.28.246000-.246FFF.29"></span>PRG and CHR bank (\$6000-\$6FFF) {#PRG_and_CHR_bank_.28.246000-.246FFF.29}

The PRG and CHR banking for this mapper is similar to [GNROM](../GxROM/) or [Color Dreams](../Color_Dreams/).

    7  bit  0
    .CPP ..CC
     |||   ||
     |++------ Select 32 KiB PRG bank at $8000
     +-----++- Select 8 KiB CHR bank at $0000

### <span id="Audio_control_.28.247000-.247FFF.29"></span>Audio control (\$7000-\$7FFF) {#Audio_control_.28.247000-.247FFF.29}

    7  bit  0
    ..RP TTTT
      || ||||
      || ++++- Track number
      |+------ While 0 and no sound is currently playing, start (or restart) sound specified in lower 4 bits
      +------- If 0 for at least 19µs=34cycles, the µPD7756C is reset and sound playback stops

Track numbers used in the game:

|     |                     |
|-----|---------------------|
| 0   | Strike              |
| 1   | Ball                |
| 2   | Time                |
| 3   | Out                 |
| 4   | Safe                |
| 5   | Foul ball           |
| 6   | Fair ball           |
| 7   | Batter out          |
| 8   | Play ball           |
| 9   | Ball four           |
| 10  | Home run            |
| 11  | New pitcher         |
| 12  | Ouch (hit by pitch) |
| 13  | Fool                |
| 14  | blow sound?         |
| 15  | shout of joy?       |

## References {#References}

- D7756 [description](http://www.cpu-world.com/Support/7/77.html) and [datasheet](http://bootgod.dyndns.org:7777/downloads/UPD7755_UPD7756_UPD7757_UPD7758.pdf)
- BBS topics [762](http://forums.nesdev.org/viewtopic.php?t=762) and [4030](http://forums.nesdev.org/viewtopic.php?p=32572#p32572)
