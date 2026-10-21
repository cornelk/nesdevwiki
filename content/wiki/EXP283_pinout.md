---
title: "EXP283 pinout"
url: "/wiki/EXP283_pinout/"
description: "The first release of Based Loaded has a patcher IC labelled \"EXP283\" (shrink DIP42) that patches the game to make the NMI handler safely reentrant. Subsequent releases instead have the NMI handler…"
wikiPage: "EXP283_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

The first release of [Based Loaded](https://nescartdb.com/profile/view/745/bases-loaded) has a patcher IC labelled "EXP283" (shrink DIP42) that patches the game to make the NMI handler safely reentrant. Subsequent releases instead have the NMI handler detect reentry and return early.

It appears to do two things:

- Reads from `PRG ROM ADDRESS & $7FFF == $0180` replace the byte with \$60; this replaces a `lda $0573` with `lda $6073`
- Reads from \$6000-\$6FFF return MMC1 PRG A15 through A17 on D4-D6.

<!-- -->

                NC ?? | 1   42 | -- +5V
                NC ?? | 2   41 | <- MMC1 PRG A17
                NC ?? | 3   40 | <- MMC1 PRG A15
                NC ?? | 4   39 | <- MMC1 PRG A14
                NC ?? | 5   38 | <- CPU A12
                NC ?? | 6   37 | <- CPU A13
                NC ?? | 7   36 | <- CPU A7
                NC ?? | 8   35 | <- CPU A8
                NC ?? | 9   34 | <- CPU A6
    PRG RAM +CE IN -> | 10  33 | <- CPU A9
               GND ?? | 11  32 | <- CPU A5
            CPU D7 <> | 12  31 | <- CPU A11
            CPU D6 <> | 13  30 | <- CPU A4
            CPU D5 <> | 14  29 | <- MMC1 PRG A16
            CPU D4 <> | 15  28 | <- CPU A3
            CPU D3 <> | 16  27 | <- CPU A10
            CPU D2 <> | 17  26 | <- CPU A2
            CPU D1 <> | 18  25 | -> PRG ROM /CE OUT
            CPU D0 <> | 19  24 | <- CPU A1
                NC ?? | 20  23 | <- MMC1 PRG /CE IN
               GND -- | 21  22 | <- CPU A0
