---
title: "Irem G-101 pinout"
url: "/wiki/Irem_G-101_pinout/"
description: "Irem G-101: 52-pin 0.6\" high-density PDIP (canonically iNES Mapper 032)"
wikiPage: "Irem_G_101_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Irem G-101: 52-pin 0.6" high-density PDIP (canonically [iNES Mapper 032](../INES_Mapper_032/))

                     .--\/--.
              GND -- | 1  52| -- +5V
           (f) M2 -> | 2  51| -> CIRAM A10 (f)
          (f) R/W -> | 3  50| ??
      (f) /ROMSEL -> | 4  49| -> PRG /CE (r)
      (f) PPU A13 -> | 5  48| ??
      (f) PPU RD# -> | 6  47| ??
      (f) CPU A14 -> | 7  46| ??
      (f) CPU A13 -> | 8  45| -> CHR /OE (r)
     (fr) CPU A12 -> | 9  44| ??
      (fr) CPU A2 -> |10  43| -> (always high)
      (fr) CPU A1 -> |11  42| -> (always high)
      (fr) CPU A0 -> |12  41| -> PRG A17 (r)
      (f) PPU A12 -> |13  40| -> PRG A16 (r)
      (f) PPU A11 -> |14  39| -> PRG A15 (r)
      (f) PPU A10 -> |15  38| -> PRG A14 (r)
      (fr) CPU D7 -> |16  37| -> PRG A13 (r)
      (fr) CPU D6 -> |17  36| ??
      (fr) CPU D5 -> |18  35| -> CHR A17 (f)
      (fr) CPU D4 -> |19  34| -> CHR A16 (f)
      (fr) CPU D3 -> |20  33| -> CHR A15 (f)
      (fr) CPU D2 -> |21  32| -> CHR A14 (f)
      (fr) CPU D1 -> |22  31| -> CHR A13 (f)
      (fr) CPU D0 -> |23  30| -> CHR A12 (f)
                ? ?? |24  29| -> CHR A11 (f)
             mode -> |25  28| -> CHR A10 (f)
              GND -- |26  27| -- +5V
                     '------'

- mode tied to +5V: PRG banking mode selectable by software
- mode tied to gnd: PRG banking mode fixed to 8+8+16F (as in *Major League*)
- For the game *Major League*, CIRAM A10 is not connected to the G-101, and is instead simply tied to +5V

Reference: <http://w.livedoor.jp/famicomcartridge/d/Irem%20G-101>
