---
title: "SMD133 (AA6023) pinout"
url: "/wiki/SMD133_(AA6023)_pinout/"
description: "SMD133 (AA6023B): LQFP-48 0.5mm pitch (canonically NES 2.0 Mapper 268 submapper 2/3)"
wikiPage: "SMD133__AA6023__pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

SMD133 (AA6023B): LQFP-48 0.5mm pitch (canonically [NES 2.0 Mapper 268](../NES_2.0_Mapper_268/) submapper 2/3)

                                                ________
                                    PPU A11 -> / 1   48 \ <- PPU A12
                                   PPU A10 -> / 2     47 \ <- CPU A12
                                      VCC -- / 3       46 \ -> CHR A10
                                     GND -- / 4         45 \ -> CHR A16
                                 CPU A1 -> / 5           44 \ -> CHR A11
    (PRG A22, may vary) "PX0" $xxx1.1  <- / 6             43 \ -> $xxx0.5 (CHR A19 in submappers 2/3, varies in others)
                              CHR A13 <- / 7               42 \ -> $xxx0.4 (CHR A18, varies)
                             CHR A14 <- / 8                 41 \ <- CPU D3
                            CHR A12 <- / 9     SMD133        40 \ <- CPU D2
                  ("HV") CIRAM A10 <- / 10     (AA6023)       39 \ <- CPU D4
                          CHR A15 <- / 11                      38 \ <- CPU D1
                         CHR A17 <- / 12                        37 \ <- CPU D5
                         PRG A20 <- \ 13                        36 / <- CPU D0
                             /IRQ <- \ 14                      35 / -> $xxx1.2 "PX1" (PRG A21, may vary)
                           /ROMSEL -> \ 15                    34 / <- CPU D6
                            CPU R/W -> \ 16                  33 / <- CPU A0
                             PRG A15 <- \ 17                32 / <- CPU D7
                              PRG A13 <- \ 18              31 / -> PRG RAM /CE ("EXT CSB")
                               CPU A14 -> \ 19            30 / <- M2 ("CK18")
                           K5K7 or K5K6 -> \ 20          29 / -> VFL
                                 PRG A16 <- \ 21        28 / -- VDD
                                  PRG A14 <- \ 22      27 / -> PRG A19
                                   PRG A18 <- \ 23    26 / -> PRG /CE ("POEB")
                                    CPU A13 -> \ 24  25 / -> PRG A17
                                                --------
    Legend:
    ->[SMD]<- Input
    <-[SMD]-> Output
    --[SMD]-- Power

    pin 20: 
        low / floating: supervisor registers at $7xxx or $6xxx depending on silicon revision
        high: supervisor registers at $5xxx

"VFL" is reportedly a diode from VDD, originally intended for powering 3V flash. (Don't do that, it's not a good enough voltage regulator).

Names in "" are from Mindkids's original documentation.

Source: [\[1\]](https://aliexpress.com/item/1005004245220736.html)

The AA6023 is also provided as in [DIP-40 package](https://ko.aliexpress.com/item/1005007598009710.html), mostly compatible with original AX5202P but a few pins got different purpose.

                .--\/--.            
          M2 -> |01  40| -- VCC
    WRAM /CE <- |02  39| <- CPU A1
      CPU D7 -> |03  38| -> PRG /CE
      CPU A0 -> |04  37| -> PRG A17
      CPU D6 -> |05  36| <- CPU A13
      CPU D0 -> |06  35| -> PRG A18
      CPU D5 -> |07  34| -> PRG A14
      CPU D1 -> |08  33| -> PRG A16
      CPU D4 -> |09  32| <- CPU A14
      CPU D2 -> |10  31| -> PRG A13
      CPU D3 -> |11  30| -> PRG A15
     CHR A18 <- |12  29| <- CPU R/W
     CHR A19 <- |13  28| <- /ROMSEL
     CHR A11 <- |14  27| -> /IRQ
     CHR A16 <- |15  26| -> CHR A17
     CHR A10 <- |16  25| -> CHR A15
     PPU A12 -> |17  24| -> CIRAM A10
     PPU A11 -> |18  23| -> CHR A12
     PPU A10 -> |19  22| -> CHR A14
         GND -- |20  21| -> CHR A13
                `------'
             AX5202P-AA6023

    pin 39 has internal pull-down
