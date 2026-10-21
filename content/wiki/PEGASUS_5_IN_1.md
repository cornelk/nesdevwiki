---
title: "PEGASUS 5 IN 1"
url: "/wiki/PEGASUS_5_IN_1/"
description: "PEGASUS 5 IN 1 (known in Poland as Złota Piątka - Golden Five) is an unlicenced Famiclone cartridge (sold in box with manual), which consists of the following 5 Codemasters' games (each of them was…"
wikiPage: "PEGASUS_5_IN_1"
---

[<img src="../../wiki-images/5-in-1_menu.png" width="256" height="223" alt="5-in-1 menu.png" />](../../wiki-images/5-in-1_menu.png)

[](../../wiki-images/5-in-1_menu.png)

[](../../wiki-images/5-in-1_menu.png)

[](../../wiki-images/5-in-1_menu.png)

[](../../wiki-images/5-in-1_menu.png)[<img src="../../wiki-images/5in1_2xROM_components.jpg" width="300" height="205" />](../../wiki-images/5in1_2xROM_components.jpg)

[](../../wiki-images/5in1_2xROM_components.jpg)

[2xROM version components](../../wiki-images/5in1_2xROM_components.jpg)

[](../../wiki-images/5in1_2xROM_components.jpg)

[](../../wiki-images/5in1_2xROM_components.jpg)

[](../../wiki-images/5in1_2xROM_components.jpg)[<img src="../../wiki-images/5in1_2xROM_bottom.jpg" width="300" height="168" />](../../wiki-images/5in1_2xROM_bottom.jpg)

[](../../wiki-images/5in1_2xROM_bottom.jpg)

[2xROM version bottom side](../../wiki-images/5in1_2xROM_bottom.jpg)

[](../../wiki-images/5in1_2xROM_bottom.jpg)

[](../../wiki-images/5in1_2xROM_bottom.jpg)

[](../../wiki-images/5in1_2xROM_bottom.jpg)[<img src="../../wiki-images/5in1_2xROM_sch.png" width="300" height="215" />](../../wiki-images/5in1_2xROM_sch.png)

[](../../wiki-images/5in1_2xROM_sch.png)

[2xROM version schematics](../../wiki-images/5in1_2xROM_sch.png)

[](../../wiki-images/5in1_2xROM_sch.png)

[](../../wiki-images/5in1_2xROM_sch.png)

[](../../wiki-images/5in1_2xROM_sch.png)[<img src="../../wiki-images/5in1_4xROM_components.jpg" width="300" height="225" />](../../wiki-images/5in1_4xROM_components.jpg)

[](../../wiki-images/5in1_4xROM_components.jpg)

[4xROM version components side](../../wiki-images/5in1_4xROM_components.jpg)

[](../../wiki-images/5in1_4xROM_components.jpg)

[](../../wiki-images/5in1_4xROM_components.jpg)

[](../../wiki-images/5in1_4xROM_components.jpg)[<img src="../../wiki-images/5in1_4xROM_top.jpg" width="300" height="217" />](../../wiki-images/5in1_4xROM_top.jpg)

[](../../wiki-images/5in1_4xROM_top.jpg)

[4xROM version top side](../../wiki-images/5in1_4xROM_top.jpg)

[](../../wiki-images/5in1_4xROM_top.jpg)

[](../../wiki-images/5in1_4xROM_top.jpg)

[](../../wiki-images/5in1_4xROM_top.jpg)[<img src="../../wiki-images/5in1_4xROM_bottom.jpg" width="300" height="215" />](../../wiki-images/5in1_4xROM_bottom.jpg)

[](../../wiki-images/5in1_4xROM_bottom.jpg)

[4xROM version bottom side](../../wiki-images/5in1_4xROM_bottom.jpg)

[](../../wiki-images/5in1_4xROM_bottom.jpg)

[](../../wiki-images/5in1_4xROM_bottom.jpg)

[](../../wiki-images/5in1_4xROM_bottom.jpg)[<img src="../../wiki-images/5in1_4xROM_sch.png" width="300" height="177" />](../../wiki-images/5in1_4xROM_sch.png)

[](../../wiki-images/5in1_4xROM_sch.png)

[4xROM version schematics](../../wiki-images/5in1_4xROM_sch.png)

[](../../wiki-images/5in1_4xROM_sch.png)

PEGASUS 5 IN 1 (known in Poland as Złota Piątka - Golden Five) is an unlicenced Famiclone cartridge (sold in box with manual), which consists of the following 5 Codemasters' games (each of them was also separately available as licensed game):

- Big Nose Freaks Out,
- Micro Machines,
- Fantastic Adventures of Dizzy (blue version),
- Ultimate Stuntman,
- Big Nose The Caveman.

Cartridge menu was encoded into Dizzy and Ultimate Stuntman was slightly modified (Codemaster's title screen was cut off and game was adapted to be playable on Dendy's clones, as the original ROM hangs due to bug in video system detection routine). Other games are the same like licensed version.

The cartridge PCB was available in two versions:

- 4xROM (3x 2 Mbit mask ROM + 4 Mbit mask KROM with additional 74138 decoder)
- 2xROM (2 Mbit mask ROM + 8 Mbit mask ROM)

First version is probably older due to earlier datamarks on chips. PAL chip is slightly different in both versions, although the ROM's content (after combining) and mapper description is the same

# Registers {#Registers}

OUTER & INNER registers are cleared on startup but not software reset

    $8000-$bfff: OUTER REGISTER
    [.... wPPP]
          ||||
          |+++- select 256KB outer PRG bank (one out of 5 games: 0=Dizzy, 1=BNFO, 2=US, 3=BNTC, 4=MM)
          +---- 1: protect OUTER register from further writes

    $c000-$ffff: INNER REGISTER
    [.... pppp]
          ||||
          ++++- select 16KB inner PRG bank 

# Memory map {#Memory_map}

- \$8000-\$bfff - switchable 16 kB PRG-ROM bank
- \$c000-\$ffff - fixed 16 kB PRG-ROM bank

<!-- -->

    $8000-$bfff | $c000-$ffff
    ------------+------------
    PPPpppp     |  PPP1111

# Trivia {#Trivia}

- PAL's marking on every PCB was removed to make rev-en harder,

<!-- -->

- The reset R-C circuit (which clears both registers at startup but does not affect them after console reset) is problematic, causing the cartridge to not start properly on every power-up or make this is a side-effect of the copy-protection descibed bellow,

<!-- -->

- This multicart uses something whis acts like copy protection:
  - When game is power up, all regs are cleared by R-C circuit, so the menu routine from Dizzy's ROM is started,
  - At some moment game check for magic value in RAM if it was run one time or not. At this time it was not but it marks that the first time hast just occured.
  - At CPU cycle 101077 (0.06sec after powerup) game writes \$0C (0b1100) to \$8927, which would normally result in switching to Micro Machines bank and protect outer register from further writes. However, 100n capacitor in RC circuit is charged by very high 39k resistance and at this time it hasn't charged yet so the register is still hold in reset and ignores writes. Then game compares \$0C with \$08 and because comparison fails, it jumps to reset (\$FFFC).
  - Now checking for first run fails so the game jumps to display menu routine.
