---
title: "Sunsoft 3 pinout"
url: "/wiki/Sunsoft_3_pinout/"
description: "Sunsoft-3: 42-pin 0.6\" shrink DIP (Canonically mapper 67)"
wikiPage: "Sunsoft_3_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Sunsoft-3: 42-pin 0.6" shrink DIP (Canonically [mapper 67](../INES_Mapper_067/))

                       .---\/---.
         delayed M2 -> | 01  42 | -- VCC                
       (f)       M2 -> | 02  41 | <- EXT /IRQ (n/c)
       (fr) CPU A11 -> | 03  40 | <- CPU D7  (fr)
       (fr) CPU A12 -> | 04  39 | <- CPU D6  (fr)
       (fr) CPU A13 -> | 05  38 | <- CPU D5  (fr)
       (f)  CPU A14 -> | 06  37 | <- CPU D4  (fr)
       (f)  /ROMSEL -> | 07  36 | <- CPU D3  (fr)
       (f)  CPU R/W -> | 08  35 | <- CPU D2  (fr)
    (f) PRG ROM /CE <- | 09  34 | <- CPU D1  (fr)
       (f)CIRAM A10 <- | 10  33 | <- CPU D0  (fr)
       (r)  CHR /CE <- | 11  32 | -> PRG RAM /CE  (n/c)
       (fr) PPU A10 -> | 12  31 | -> PRG RAM +CE  (n/c)
       (f)  PPU A11 -> | 13  30 | <- PPU A13 (f)       
       (f)  PPU A12 -> | 14  29 | <- PPU /RD (f)          
       (r)  CHR A16 <- | 15  28 | -> /IRQ    (f)    
       (r)  CHR A15 <- | 16  27 | -> PRG A14 (r) 
       (r)  CHR A14 <- | 17  26 | -> PRG A15 (r) 
       (r)  CHR A13 <- | 18  25 | -> PRG A16 (r) 
       (r)  CHR A12 <- | 19  24 | -> PRG A17 (n/c) 
       (r)  CHR A11 <- | 20  23 | -> /PRG A17 (n/c)      
            GND     -- | 21  22 | -> $f800.4 (n/c)
                       '--------'

Notes:

- 1 is delayed by roughly 80ns by an external RC
- 1 and 2 must both be high to access registers or drive the RAM enables, and a rising edge of this combined signal clocks the IRQ counter.
- 7 is CPU /A15 in the Vs. System
- 28 combines (bitwise AND) the internal IRQ with pin 41; as far as we know 28 is still push-pull

Source: \[[Krzysiobal](https://forums.nesdev.org/viewtopic.php?t=19759)\]
