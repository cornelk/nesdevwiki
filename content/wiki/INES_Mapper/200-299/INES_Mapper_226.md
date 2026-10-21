---
title: "INES Mapper 226"
url: "/wiki/INES_Mapper_226/"
description: "The original describing document states the 2s bit of the register at $8001 disables CHRRAM writes. No games seem to rely on it, though. It also says that the mapper hardware is made of a 74LS74A,…"
wikiPage: "INES_Mapper_226"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

Here are Disch's original notes:  
     ========================
     =  Mapper 226          =
     ========================

     Example Games:
     --------------------------
     76-in-1
     Super 42-in-1

     Registers:
     ---------------------------

     Range, Mask:  $8000-FFFF, $8001

       $8000:  [PMOP PPPP]
          P = Low 6 bits of PRG Reg
          M = Mirroring (0=Horz, 1=Vert)
          O = PRG Mode

       $8001:  [.... ...H]
          H = high bit of PRG

     PRG Setup:
     ---------------------------

     Low 6 bits of the PRG Reg come from $8000, high bit comes from $8001

                    $8000   $A000   $C000   $E000  
                  +-------------------------------+
     PRG Mode 0:  |             <Reg>             |
                  +-------------------------------+
     PRG Mode 1:  |      Reg      |      Reg      |
                  +---------------+---------------+

The [original describing document](http://nesdev.org/226.txt) states the 2s bit of the register at \$8001 disables CHRRAM writes. No games seem to rely on it, though. It also says that the mapper hardware is made of a 74LS74A, 74LS273, [74LS139](../74139/), [74LS02](../7402/), and [74LS153](../74153/).

The multicart [reportedly relies](http://web.archive.org/web/20030602133825/http://www.nesplayer.com/pirates/76in1.htm) on the value of uninitialized memory to determine exactly how many games are available.

The multicart clears both registers on soft reset.

See also: [Repair of one such cartridge](http://forums.nesdev.org/viewtopic.php?t=5642)

At least three different games using this mapper exists:

- *42 in 1 (1 MB)*
- *63 in 1 (1.5 MB)*
- *76 in 1 (2 MB)*

They appear in different ROM package and count configurations:

- [<img src="../../wiki-images/42in1_1x1MB_chip.jpg" width="120" height="69" alt="42in1 1x1MB chip" />](../../wiki-images/42in1_1x1MB_chip.jpg)

  42in1 1x1MB chip

- [<img src="../../wiki-images/42in1_2x512kB_chip_v1.jpg" width="120" height="79" alt="42in1 2x512kB chip v1" />](../../wiki-images/42in1_2x512kB_chip_v1.jpg)

  42in1 2x512kB chip v1

- [<img src="../../wiki-images/42in1_2x512kB_chip_v2.jpg" width="120" height="92" alt="42in1 2x512kB chip v2" />](../../wiki-images/42in1_2x512kB_chip_v2.jpg)

  42in1 2x512kB chip v2

- [<img src="../../wiki-images/42in1_2x512kB_chip_v3.jpg" width="120" height="85" alt="42in1 2x512kB chip v3" />](../../wiki-images/42in1_2x512kB_chip_v3.jpg)

  42in1 2x512kB chip v3

- [<img src="../../wiki-images/63in1_3x512kb_chip.jpg" width="120" height="89" alt="63in1 3x512kb chip" />](../../wiki-images/63in1_3x512kb_chip.jpg)

  63in1 3x512kb chip

- [<img src="../../wiki-images/76in1_1x2MB_chip.jpg" width="120" height="69" alt="76in1 1x2MB chip" />](../../wiki-images/76in1_1x2MB_chip.jpg)

  76in1 1x2MB chip

- [<img src="../../wiki-images/76in1_1x2MB_epoxy_or_chip.jpg" width="120" height="72" alt="76in1 1x2MB epoxy or chip" />](../../wiki-images/76in1_1x2MB_epoxy_or_chip.jpg)

  76in1 1x2MB epoxy or chip

- [<img src="../../wiki-images/76in1_1x2MB_epoxy.jpg" width="120" height="63" alt="76in1 1x2MB epoxy" />](../../wiki-images/76in1_1x2MB_epoxy.jpg)

  76in1 1x2MB epoxy

- [<img src="../../wiki-images/76in1_2x1MB_chip_v1.jpg" width="120" height="68" alt="76in1 2x1MB chip v1" />](../../wiki-images/76in1_2x1MB_chip_v1.jpg)

  76in1 2x1MB chip v1

- [<img src="../../wiki-images/76in1_2x1MB_chip_v2.jpg" width="120" height="79" alt="76in1 2x1MB chip v2" />](../../wiki-images/76in1_2x1MB_chip_v2.jpg)

  76in1 2x1MB chip v2

- [<img src="../../wiki-images/76in1_2x1MB_epoxy.jpg" width="120" height="70" alt="76in1 2x1MB epoxy" />](../../wiki-images/76in1_2x1MB_epoxy.jpg)

  76in1 2x1MB epoxy

- [<img src="../../wiki-images/76in1_4x512kB_chip.jpg" width="120" height="90" alt="76in1 4x512kB chip" />](../../wiki-images/76in1_4x512kB_chip.jpg)

  76in1 4x512kB chip
