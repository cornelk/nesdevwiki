---
title: "MMC4 pinout"
url: "/wiki/MMC4_pinout/"
wikiPage: "MMC4_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Nintendo [MMC4](../MMC4/): 44-pin TQFP

                               ___
                              /   \
                             /     \
                       ? ?? / 1  44 \ <- CPU R/W (n)
            (n) /ROMSEL -> / 2    43 \ <- CPU D0 (nrw)
           (n) PPU /RD -> / 3   O  42 \ <- CPU D1 (nrw)
           (n) PPU A3 -> / 4        41 \ <- CPU D2 (nrw)
                   ? ?? / 5          40 \ <- CPU D3 (nrw)         Orientation:
                GND -- / 6            39 \ <- CPU D4 (nrw)        ------------------
        (n) PPU A4 -> / 7              38 \ -> PRG /CE (r)           33        23
       (n) PPU A5 -> / 8                37 \ -> PRG A17 (r)           |        |
      (n) PPU A6 -> / 9                  36 \ -> PRG A16 (r)         .----------.
     (n) PPU A7 -> / 10                   35 \ -> PRG A15 (r)     34-|          |-22
             ? ?? / 11                     34 \ -> PRG A14 (r)       | Nintendo |
                 /        Nintendo MMC-4       \                     | MMC-4    |
                 \        Package QFP-44       /                  44-|o         |-12
    (n) PPU A8 -> \ 12     0.8mm pitch     33 / -- Tied to Pin 32    \----------'
     (n) PPU A9 -> \ 13                   32 / <- CPU A12 (n)         |        |
     (n) PPU A10 -> \ 14                 31 / <- CPU A13 (n)          1        11
      (n) PPU A11 -> \ 15               30 / <- CPU A14 (n)
       (n) PPU A12 -> \ 16             29 / <- M2 (n)             Legend:
        (n) PPU A13 -> \ 17           28 / -- GND                 ------------------------------
                 GND -- \ 18         27 / -- VCC                  --[MMC4]-- Power
          (r) CHR A12 <- \ 19       26 / -> CIRAM A10 (n)          ->[MMC4]<- MMC4 input
           (r) CHR A13 <- \ 20     25 / -> PRG RAM +CE (w)        <-[MMC4]-> MMC4 Output
            (r) CHR A14 <- \ 21   24 / -> CHR A16 (r)             ??[MMC4]?? Unknown
             (r) CHR A15 <- \ 22 23 / -> PRG RAM /CE (w)              n      NES connection
                             \     /                                  r      ROM chip connection
                              \   /                                   w      RAM chip connection
                               \ /
                                V

- Source: [INL on the forum](http://forums.nesdev.org/viewtopic.php?p=100951#p100951)
