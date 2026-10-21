---
title: "BU3266 / BU3270 pinout"
url: "/wiki/BU3266_/_BU3270_pinout/"
description: "Custom Nintendo PIO chip found in the New Famicom (HVC-101). This replaces the U3 (139) and the U7/U8 (368)."
wikiPage: "BU3266___BU3270_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Custom Nintendo PIO chip found in the New Famicom (HVC-101). This replaces the U3 (139) and the U7/U8 (368).

              .---v---.  
    INV 3I -> |01   32| -- +5V
    INV 3O <- |02   31| -> INV 3B1
        M2 -> |03   30| -> INV 3B2
       A15 -> |04   29| -> /ROMSEL
       A14 -> |05   28| -> PPU /CE
       A13 -> |06   27| -> CPU RAM /CE
     P1 D0 -> |07   26| <- INV 2I
     P0 D0 -> |08   25| -> INV 2O
     P1 D1 -> |09   24| -> CPU D0
     P0 D1 -> |10   23| -> CPU D1
     P1 D2 -> |11   22| <- /INP0 ($4016 /RD)
     P0 D2 -> |12   21| <- /INP1 ($4017 /RD)
     P1 D3 -> |13   20| -> CPU D2
     P1 D4 -> |14   19| -> CPU D3
       GND -- |15   18| -> CPU D4
    INV 1I -> |16   17| -> INV 1O
              '-------'  

Notes:

- INV 1I/O (16-17) = Audio Inverter
- INV 2I/O (25-26) = PA13 Inverter
- INV 3I/O/B1/B1 (01,02,30,31) = unused (intended for CIC clock based on topology) [\[1\]](https://forums.nesdev.org/viewtopic.php?p=268199#p268199)

Sources:

- [nes_pio_pinout.txt](http://nesdev.org/nes_pio_pinout.txt)
- [BBS - AV Famicom JIO Chip Replacement Idea](https://forums.nesdev.org/viewtopic.php?t=16764)
