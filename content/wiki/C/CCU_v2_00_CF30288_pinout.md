---
title: "CCU v2.00 CF30288 pinout"
url: "/wiki/CCU_v2.00_CF30288_pinout/"
description: "Camerica CCU: 28-pin 0.4\" shrink PDIP marked \"CCU_V2.00 CF30288 N X 9240 PHILLIPINES\" (Mappers 71 and 232)"
wikiPage: "CCU_v2_00_CF30288_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Camerica CCU: 28-pin 0.4" shrink PDIP marked "CCU_V2.00 CF30288 N X 9240 PHILLIPINES" (Mappers [71](../INES_Mapper_071/) and [232](../INES_Mapper_232/))

               .---v---.
     CPU M2 -> |01   28| ??
     CPU D4 -> |02   27| ??
     CPU D3 -> |03   26| ??
     CPU D2 -> |04   25| ??
     CPU D1 -> |05   24| -> ROM /OE
     CPU D0 -> |06   23| ??
        VCC -- |07   22| -- GND
     CPU A0 -> |08   21| -> CIC STUN
    CPU A13 -> |09   20| -> $C000.3 OR CPU A14
    CPU A14 -> |10   19| -> $8000.3
    /ROMSEL -> |11   18| -> $C000.2 OR CPU A14
    CPU R/W -> |12   17| -> $8000.4
            ?? |13   16| -> $C000.1 OR CPU A14
            ?? |14   15| -> $C000.0 OR CPU A14
               '-------'
            28-pin 0.4" shrink DIP

Notes:

- This chip is used in *Aladdin Deck Enhancer 1.1/2.0* and *Super Sports Challenge*
- CIC STUN pin is used only in *Aladdin Deck Enhancer 1.1*

<!-- -->

- Source: [\[1\]](http://forums.nesdev.org/viewtopic.php?f=9&t=19781&p=273471#p247213)
