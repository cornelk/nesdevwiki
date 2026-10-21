---
title: "Sunsoft 5 pinout"
url: "/wiki/Sunsoft_5_pinout/"
description: "Sunsoft FME-7, 5A (and 5B): 44-pin 0.8mm PQFP"
wikiPage: "Sunsoft_5_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

[Sunsoft FME-7](../Sunsoft_FME-7/), 5A (and [5B](../Sunsoft_5B_audio/)): 44-pin 0.8mm PQFP

                                 ___
                                /   \
                               /     \
              (nrw) CPU D3 -> /01   44\ <- CPU D2 (nrw)
                      varies /02     43\ <- CPU D4 (nrw)
                     varies /03   O   42\ <- CPU D1 (nrw)
                 (n) M2 -> /04         41\ <- CPU D5 (nrw)
           (n) CPU A13 -> /05           40\ <- CPU D0 (nrw)
                  GND -- /06             39\ -- VCC                  Orientation:
         (n) CPU A14 -> /07               38\ <- CPU D6 (nrw)        ------------------
        (n) /ROMSEL -> /08                 37\ <- CPU D7 (nrw)          33        23
     (nrw) CPU R/W -> /09                   36\ -> PRG A13 (r)           |        |
         (n) /IRQ <- /10                     35\ -> PRG A16 (r)         .----------.
     (n) PPU /RD -> /11        Sunsoft        34\ -> PRG A14 (r)     34-|          |-22
                   /       FME-7 / 5A / 5B       \                      | Sunsoft  |
                   \       Package  QFP-44       /                      |FME7/5A/5B|
     (n) PPU A10 -> \12                       33/ -> PRG A15 (r)     44-|o         |-12
      (n) PPU A11 -> \13                     32/ -> PRG A17 (r)         \----------'
       (n) PPU A12 -> \14                   31/ -> PRG /CE (r)           |        |
        (n) PPU A13 -> \15                 30/ -> PRG RAM /CE (w)        1        11
         (r) CHR /CE <- \16               29/ -> PRG RAM +CE (w)
                  VCC -- \17             28/ -- GND          Legend:
                   varies \18           27/ varies          ------------------------------
     (nr) CHR/CIRAM A10 <- \19         26/ -> CHR A17 (r)    --[FME-7]-- Power
             (r) CHR A16 <- \20       25/ -> CHR A15 (r)     ->[FME-7]<- FME-7 input
              (r) CHR A11 <- \21     24/ -> CHR A14 (r)      <-[FME-7]-> FME-7 output
               (r) CHR A13 <- \22   23/ -> CHR A12 (r)            n      NES connection
                               \     /                            r      ROM chip connection
                                \   /                             w      RAM chip connection
                                 \ /
                                  V

    pin 2: Amplifier in (5B), no connect (5A), or PRG A18 (FME-7)
    pin 3: Amplifier out (5B), or no connect (5A and FME-7)
    pin 18: Audio /disable (5B) [1], unknown input (5A), /EXT IRQ (FME-7) - all three chip variations have internal pull-up
    pin 27: Audio out (5B), no connect (5A), or VCC (FME-7)

    On FME-7, the two ground supplies and three VCC supplies are shorted internally. Boards (like GRM-E301) that were used with the 5A or 5B did not connect pin 27 to VCC

    Expansion audio mixing circuit:[2]

                                         +                 +
               From S5B----/\/\/---+---||---+---/\/\/---+---||---To RF
               (S5B.27)    1K      |   1u   |   100K    |   1u   (Cart.46)
                                   |        |           |
               From 2A03---/\/\/---+        |           |
               (Cart.45)   10K              |           |
                                            |           |
               Amp. In----------------------+           |
               (S5B.2)                                  |
                                                        |
               Amp. Out---------------------------------+
               (S5B.3)
