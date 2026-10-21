---
title: "VRC1 pinout"
url: "/wiki/VRC1_pinout/"
description: "There isn't a consistent pinout for the VRC1 and its clones; here are two. The pins are mostly in the same order, but power and ground definitely move."
wikiPage: "VRC1_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

There isn't a consistent pinout for the [VRC1](../VRC1/) and its clones; here are two. The pins are *mostly* in the same order, but power and ground definitely move.

Anything with a ? next to it is uncertain; these pinouts were deduced solely based on images of the PCB.

VRC-1: 28-pin 0.6" DIP (canonically **mapper 75**)

Labeled "Konami 8626K7 007422 VRC 075" on PCB labeled "Konami 302114A":

                    .--\/--.
     (r) PRG A15 <- |01  28| -- +5V
     (r) PRG A14 <- |02  27| <- CPU A14 (f)
    (fr) CPU A12 -> |03  26| <- CPU A13 (f)
     (r) PRG A13 <- |04  25| <- PPU A12 (f)
     (r) PRG A16 <- |05  24| <- PPU A11 (fr)
     (fr) CPU D0 -> |06  23| <- PPU A10 (fr)
     (fr) CPU D1 -> |07  22| -> PRG ROM /CE (r)
     (fr) CPU D2 -> |08  21| <- R/W (f)
     (fr) CPU D3 -> |09  20| -> CHR /CE (r)
     (r) CHR A12 <- |10  19| <- PPU A13 (f)
     (r) CHR A14 <- |11  18| <- PPU /RD (f)
     (r) CHR A13 <- |12  17| -> CIRAM A10?
     (r) CHR A15 <- |13  16| <- /ROMSEL (f)
             GND -- |14  15| -> CHR A16 (r)
                    `------'

Labeled "VRC 007421 8805 Z04" on PCB labeled "JALECO-20":

                   .--\/--.
    (r) PRG A15 <- |01  28| <- CPU A14 (f)
       ?PRG A14 <- |02  27| <- CPU A13 (f)
       ?CPU A12 -> |03  26| <- /ROMSEL? 
       ?PRG A13 <- |04  25| <- PPU A12?        
    (r) PRG A16 <- |05  24| <- PPU A11?        
    (fr) CPU D0 -> |06  23| <- R/W (f)
            GND -- |07  22| -> PRG ROM /CE (r)
    (fr) CPU D1 -> |08  21| -- +5V
    (fr) CPU D2 -> |09  20| -> CHR /CE?
    (fr) CPU D3 -> |10  19| <- PPU A10?
       ?CHR A12 <- |11  18| <- PPU /RD (f)
    (r) CHR A14 <- |12  17| -> CIRAM A10 (f)
       ?CHR A13 <- |13  16| <- (CPU /ROMSEL OR R/W)?
    (r) CHR A15 <- |14  15| -> CHR A16 (r)
                   `------'

Other part numbers seen include:

- Konami 007506
- (NEC µP)D65005C075 (D65xxx = CMOS gate array; 005 = 1.5µ 320 gates; C = DIP28; 75 = 75th order)
