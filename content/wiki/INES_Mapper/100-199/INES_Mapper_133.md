---
title: "INES Mapper 133"
url: "/wiki/INES_Mapper_133/"
description: "iNES Mapper 133 is used for two versions of the Sachen unlicensed game Jovial Race and the 60-pin Joy Van release of 盜帥 (Master Chu and the Drunkard Hu)."
wikiPage: "INES_Mapper_133"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 133 is used for two versions of the [Sachen](http://community.fandom.com/wiki/c:bootleggames:Sachen) unlicensed game *Jovial Race* and the 60-pin Joy Van release of *盜帥 (Master Chu and the Drunkard Hu)*.

## Original 60-pin release {#Original_60-pin_release}

The 60-pin releases of both games use the Sachen 3009 board, which wires a [Namco 108](../INES_Mapper_206/) (marked "ASLIC AX-24G") clone in an inscrutably complex way. There is no UNIF mapper designated for the Sachen 3009 board.

     Mask: $E001
     $8000: [..B. .B.B] - Select register for writing
     $8001: [...X XPCC] - Write value to register
                | ||||
                | ||++--- select 8 KiB CHR bank
                | |+----- select 32 KiB PRG bank
                +-+------ analog feedback path

Four of the registers produce an analog feedback path through the 108's CHR registers:

| XX bits | if this register was selected by the XX bits, switch to using contents of register |
|----|----|
| \$00 | \$04 |
| \$08 | \$24 |
| \$10 | \$01 |
| \$18 | \$21 |

If this feedback path doesn't settle on one specific register, then the selected banks will rapidly switch between all the registers specified by the feedback path, and the 108 clone will get hot.

## 72-pin release {#72-pin_release}

The 72-pin release of *Jovial Race* uses the Sachen 72008 board, a much simpler board that is similar but not identical to [INES Mapper 079](../NINA-003-006/). Its UNIF MAPR is **UNL-SA-72008**.

     Mask: $E100
     $4100: [.... .PCC] - Select 32 KiB PRG bank and 8 KiB CHR bank

The 72-pin releases of *Master Chu and the Drunkard Hu* on the other hand use [INES Mapper 011](../Color_Dreams/) for the Color Dreams release and [INES Mapper 079](../NINA-003-006/) for the Sachen release.

## Emulator compatibility {#Emulator_compatibility}

The 60-pin releases of *Jovial Race* and *盜帥* contain bankswitching code for both board variants, whereas the 72-pin release of *Jovial Race* only contains code for the \$4100-based latch variant. For this reason, and because the 60-pin release of *Jovial Race* was analyzed much later, emulators only emulate, and only need to emulate, the 72-pin variant of the board.

## See also {#See_also}

- <http://cah4e3.shedevr.org.ru/%5Blst%5D-sachen-mappers.txt>
- <https://forums.nesdev.org/viewtopic.php?p=210246#p210246>
