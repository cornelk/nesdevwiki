---
title: "Forbidden Four"
url: "/wiki/Forbidden_Four/"
description: "Double Crossing: The Forbidden Four is a multicart engine based on The Legend of Zelda and designed to accommodate the four hardest-to-obtain NES games in the first-generation Animal Crossing games."
wikiPage: "Forbidden_Four"
wikiCategories:
  - name: "Homebrew"
    url: "/category/Homebrew/"
---

<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/ForbiddenFour.png"><img src="../../wiki-images/ForbiddenFour.png" width="300" height="224" /></a>
<figcaption>Screenshot of menu</figcaption>
</figure>

**Double Crossing: The Forbidden Four** is a multicart engine based on *The Legend of Zelda* and designed to accommodate the four hardest-to-obtain NES games in the first-generation *[Animal Crossing](https://en.wikipedia.org/wiki/Animal_Crossing_(video_game))* games.

## Background {#Background}

*Animal Crossing*, a social simulator for Nintendo GameCube, included NES consoles with built-in games as collectible furniture items. The player could play these games through *AC'*s built-in emulator. There were several ways to receive games:

- Common games
  - Buy them from Tom Nook or Redd when he has them in stock
  - Enter a "universal code" created with a code generator in Tom Nook's store
- Uncommon games
  - A neighbor announces on the bulletin board that he has buried one of these in some part of town; use a shovel, find the star mark denoting a buried item, and dig it up
  - Enter a "universal code" created with a code generator in Tom Nook's store
- Rare games\*
  - Enter a code in the letter and send it to a neighbor
  - Connect a Game Boy Advance system and do a particular trading sequence on the island
  - Enter a trade code, given to a specific player name in a specific town name, in Tom Nook's store as the right player in the right town (most commonly for *[Punch-Out!!](../MMC2/)*)
  - Enter a "universal code" created with a code generator in Tom Nook's store
- Forbidden games
  - Receive one of these games at a Nintendo trade show
  - Scan a particular e-Reader card (for Mario Bros and Ice Climber)
  - Use Action Replay to change the ID number of a furni in your inventory

The forbidden games are *Mario Bros.*, *Ice Climber*, *Super Mario Bros.*, and *The Legend of Zelda*. The code for Tom Nook's store is specifically designed to reject universal codes that refer to the ID numbers of these games. These games are also completely missing from *Dōbutsu no Mori e+*, an upgraded version of the Japanese GameCube release.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

*Animal Crossing: Wild World* didn't have NES games because it was too soon after the release of Famicom Mini/Classic NES Series on GBA. Later games on Nintendo platforms (*City Folk* and *New Leaf*) didn't have NES games because they would compete with Virtual Console games from Wii Shop Channel/Nintendo eShop. *Pocket Camp* runs on hardware whose input device differs too much from that of the NES. NES furniture items would finally return (along with the addition of other Nintendo consoles) in the version 3.0 update for *New Horizons*, though playing them requires an active Nintendo Switch Online subscription.<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

\* None of them from [the developer of *Battletoads*](../Rare/)

## Instructions {#Instructions}

1.  Install 7-Zip. (Windows version on 7-Zip.org; under Ubuntu, `sudo apt-get install p7zip` to add .7z support to the File Roller archive manager.)
2.  Download and unpack [Forbidden Four](http://pics.pineight.com/nes/ForbiddenFour.7z).
3.  Put two [NROM](../NROM/)-128 games in [iNES](../INES/) format named "bros.nes" and "ice.nes" (not included) in the same folder as "menu.nes".
4.  Put one NROM-256 game named "smb1.nes" (not included) in the same folder.
5.  (Optional) Use a nametable editor to edit the names of the games in dxinglogo.nam and use CC65 to recompile the menu binary.
6.  Put The *Legend of Zelda* (or possibly other select games) named "loz.nes" in the same folder.
7.  (On systems other than Windows) Recompile the ROM builder 'buildf4'.
8.  Run "Build Multicart.bat", or on systems other than Windows, run tools/buildf4.

## How it works {#How_it_works}

Forbidden Four is a multicart engine that can take any four NES ROMs that fit the following criteria:

- Any NROM-128 game, called "Bros"
- A second NROM-128 game, called "Ice"
- Any NROM-256 game, called "SMB1"
- The *Legend of Zelda*, called "LoZ", or any other 1 Mbit [SNROM](../MMC1/) game with the same init code

Two further restrictions: The NROM code must not write to ROM \$8000-\$9FFF or \$E000-\$FFFF, as this will trigger an unexpected bankswitch on the [MMC1](../MMC1/) and probably crash the game. The NROM-128 games also must not depend on values read from the ROM [mirror](../Mirroring/) at \$8000-\$BFFF because this ROM mirror does not exist in the multicart.

This SNROM has 256 KiB of PRG ROM, 8 KiB of CHR RAM, and 8 KiB of battery-backed PRG RAM. The ROM of any MMC1 based game is divided into 16 KiB banks. There are sixteen such banks in Forbidden Four:

     0  LoZ PRG 0
     1  LoZ PRG 1
     2  LoZ PRG 2
     3  LoZ PRG 3
     4  LoZ PRG 4
     5  LoZ PRG 5
     6  LoZ PRG 6
     7  LoZ PRG 7
     8  SMB1 PRG 0
     9  SMB1 PRG 1
    10  Bros CHR and SMB1 CHR
    11  Bros PRG
    12  Ice CHR and menu CHR
    13  Ice PRG
    14  Menu PRG
    15  LoZ PRG 7 (duplicate for vectors)

The last bank of LoZ PRG is duplicated because the game expects to find it in both bank 7 and the MMC1's fixed bank.

The first 256 bytes of Menu contain

1.  a jump table with four entries,
2.  code to copy a CHR RAM loader and a jump table manager into RAM at \$0300, and
3.  the CHR RAM loader and jump table manager.

The four entries in the jump table specify which CHR page to load, which PRG page to load, how to set up the memory mapping and the mirroring, and where to put the program counter.

From the mapper init code of LoZ:

    setMMC1CTRL = $FF98
    setMMC1PRG = $FFAC
    $FF76:A9 0F     LDA #$0F
    $FF78:20 98 FF  JSR setMMC1CTRL
    ; Omitted code: set CHR bank to 0
    $FF90:A9 07     LDA #$07
    $FF92:20 AC FF  JSR setMMC1PRG
    $FF95:4C 40 E4  JMP $E440

So in order to intercept Zelda's boot sequence, we need to change the start bank and start address to point to our boot sequence at \$8020, and make sure it ends up at ctrl=\$0F bank=\$07 PC=\$E440 by the time Zelda starts.

    $FF91: 0E
    $FF96: 20
    $FF97: 80

This passes control to the first 256 bytes of Menu, which sets up CPU RAM and starts the menu using a mechanism similar to the jump table. This is done because the first 256 bytes of Menu are designed to run from \$8000-\$80FF, so that LoZ can pass control to it, while the rest of Menu runs from \$C100-\$FFFF, so that it can take over the vertical blank NMI for display timing.

## Further developments {#Further_developments}

If you are building this on top of an SNROM base game other than LoZ, you will need to change the addresses a bit in buildf4.c and then use GCC to rebuild buildf4.exe. If you have a patch to buildf4.c to support hooking other 1 Mbit SNROM games' init code, ~~talk~~ about it.

## Usage in *Animal Crossing* {#Usage_in_Animal_Crossing}

*AC* contains a blank NES/[Disk System](../Family_Computer_Disk_System/) furniture item that can be obtained through rare events (or more commonly with cheats), which seemingly does not function when interacted with. However, it was discovered several years after release that it attempts to load formatted NES ROMs or [QD](../QD/) images from the same memory card as town/save data.<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup> It is possible to encode games such as Forbidden Four into the required format, insert the data into a memory card, and have it loaded by the blank NES item in-game.<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)<span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup>

The original Nintendo 64 release of *Dōbutsu no Mori* was also discovered to contain a similar feature using bankswitched controller paks (used officially for a limited giveaway of *Ice Climber*), making game injections also possible in this version.<sup><span id="cite_ref-6"></span>[\[6\]](#cite_note-6)</sup>

## External Links {#External_Links}

- [Project thread](https://forums.nesdev.org/viewtopic.php?t=3087)

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) "[Forbidden Four](https://nookipedia.com/wiki/Forbidden_Four)" on Nookipedia
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) ["Animal Crossing: New Horizons – Nintendo Switch 2 Edition & Free Update – Announcement Trailer" by Nintendo of America](https://youtu.be/7nODHyw3lgs&t=580) (12 minutes) - "Nintendo crossover items" timestamped at 9:40.
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) "[NES (Animal Crossing)](https://nookipedia.com/wiki/Item:NES_(Animal_Crossing))" on Nookipedia
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) ["Injecting NES / Famicom Disk System Games into Animal Crossing" by Hunter R.](https://youtu.be/2w0U_evVEd4) (11 minutes)
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) ["Forbidden Four multicart & AccuracyCoin in Dōbutsu no Mori e+" by TakuikaNinja](https://youtu.be/rZZsNrOC4Bo) (7 minutes)
6.  <span id="cite_note-6"></span>[↑](#cite_ref-6) [Replicating Nintendo's Ice Climber Controller Pak for Doubutsu no Mori (N64)](https://hunter-r.com/posts/ice-climber-pak/)
