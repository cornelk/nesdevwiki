---
title: "MMC3 pinout"
url: "/wiki/MMC3_pinout/"
description: "Nintendo MMC3: 44-pin 0.8mm pitch QFP (Canonically mapper 4)"
wikiPage: "MMC3_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Nintendo [MMC3](../MMC3/): 44-pin 0.8mm pitch QFP (Canonically [mapper 4](../MMC3/))

                                  ___
                                 /   \
                                /     \
                        n/c -- / 1  44 \ -> CHR A16 (r)
               (r) CHR A10 <- / 2    43 \ -> CHR A11 (r)
              (n) PPU A12 -> / 3   O  42 \ -> PRG RAM /WE (w)
             (n) PPU A11 -> / 4        41 \ -> PRG RAM +CE (w)
            (n) PPU A10 -> / 5          40 \ -- GND
                   GND -- / 6            39 \ <- CPU D3 (nrw)         Orientation:
          (r) CHR A13 <- / 7              38 \ <- CPU D2 (nrw)        ------------------
         (r) CHR A14 <- / 8                37 \ <- CPU D4 (nrw)          33        23
        (r) CHR A12 <- / 9                  36 \ <- CPU D1 (nrw)          |        |
     (n) CIRAM A10 <- / 10                   35 \ <- CPU D5 (nrw)        .----------.
      (r) CHR A15 <- / 11                     34 \ <- CPU D0 (nrw)    34-|          |-22
                    /        Nintendo MMC3        \                      | Nintendo |
                    \        Package QFP-44       /                      | MMC3B    |
      (r) CHR A17 <- \ 12     0.8mm pitch     33 / <- CPU D6 (nrw)    44-|o         |-12
          (n) /IRQ <- \ 13                   32 / <- CPU A0 (nrw)        \----------'
        (n) /ROMSEL -> \ 14                 31 / <- CPU D7 (nrw)          |        |
                 GND -- \ 15               30 / -> PRG RAM /CE (w)        1        11
                  n/c -- \ 16             29 / <- M2 (n)
               (n) R/W -> \ 17           28 / -- GND          Legend:
            (r) PRG A15 <- \ 18         27 / -- VCC           ------------------------------
             (r) PRG A13 <- \ 19       26 / -> PRG /CE (r)    --[MMC3]-- Power
              (n) CPU A14 -> \ 20     25 / -> PRG A17 (r)     ->[MMC3]<- MMC3 input
               (r) PRG A16 <- \ 21   24 / <- CPU A13 (n)      <-[MMC3]-> MMC3 output
                (r) PRG A14 <- \ 22 23 / -> PRG A18 (r)           n      NES connection
                                \     /                           r      ROM chip connection
                                 \   /                            w      RAM chip connection
                                  \ /
                                   V
    01, 16: both officially no connection. sometimes shorted to pin 02, 15 respectively

# Functional variations {#Functional_variations}

As with several other ASIC mappers, parts of the pinout are often repurposed:

## [iNES Mapper 037](../INES_Mapper_037/) and [iNES Mapper 047](../INES_Mapper_047/) {#iNES_Mapper_037_and_iNES_Mapper_047}

Mappers [37](../INES_Mapper_037/) and [47](../INES_Mapper_047/) connect pins 42 and 30 to a [74161](../74161/)

- Pin 30 (was PRG RAM /CE) is now 74'161 CLOCK
- Pin 34 (CPU D0) is additionally 74'161 D0
- In mapper 37, Pins 36, 38 (CPU D1, CPU D2) are additionally 74'161 D1, D2.
- Pin 42 (was PRG RAM /WE) is now 74'161 /LOAD

## TKSROM and TLSROM {#TKSROM_and_TLSROM}

[iNES Mapper 118](../INES_Mapper_118/) connects pin 12 to CIRAM A10, and pin 10 is n/c.

- Pin 10 (was CIRAM A10) is now no-connect
- Pin 12 (was CHR A17) is now CIRAM A10.

## [TQROM](../INES_Mapper_119/) {#TQROM}

[iNES Mapper 119](../INES_Mapper_119/) connects pin 44 to a [7432](../7432/) and to the CHR RAM's +CE pin.

- Pin 44 (was CHR A16) is now CHR RAM +CE.
- Pin 44 is also ORed with PPU A13 and that goes to CHR ROM -CE

## Acclaim [MC-ACC pinout](../MC-ACC_pinout/) {#Acclaim_MC-ACC_pinout}

[MC-ACC](../MMC3/#IRQ_Specifics) chip: (40-pin 0.6" DIP)

                   .--\/--.
            GND ?? |01  40| -- 5V
    (r) PRG A16 <- |02  39| -> PRG A18 (r)
    (r) PRG A15 <- |03  38| -> PRG A17 (r)
    (n) CPU A14 -> |04  37| -> PRG A14 (r)
    (n) CPU A13 -> |05  36| -> PRG A13 (r)
            n/c ?? |06  35| <- CPU R/W (n)
            n/c ?? |07  34| <- /ROMSEL (n)
            n/c ?? |08  33| -> PRG /OE (r)
    (nr) CPU A0 -> |09  32| <- CPU D7 (nr)
    (nr) CPU D0 -> |10  31| <- CPU D6 (nr)
    (nr) CPU D1 -> |11  30| <- CPU D5 (nr)
    (nr) CPU D2 -> |12  29| <- CPU D4 (nr)
    (r) CHR A16 <- |13  28| <- CPU D3 (nr)
    (r) CHR A15 <- |14  27| -> CIRAM A10 (n)
    (r) CHR A14 <- |15  26| -> CHR A17 (r)
    (r) CHR A13 <- |16  25| -> /IRQ (n)
    (r) CHR A12 <- |17  24| <- PPU A12 (n)
    (r) CHR A11 <- |18  23| <- PPU A11 (n)
    (r) CHR A10 <- |19  22| <- PPU A10 (n)
            GND -- |20  21| <- M2 (n)
                   '------'

    39: confirmed in [1]
    24: on the 55741 PCB, deglitched and shaped as MCACCA12IN = AND(PPUA12,AND(AND(AND(PPUA12)))). Kevtris's notes

Source: <https://forums.nesdev.org/viewtopic.php?p=16795#p16795>

# <span id="Pirate_versions_.28600_mil_40-pin_DIP_package.29"></span>Pirate versions (600 mil 40-pin DIP package) {#Pirate_versions_.28600_mil_40-pin_DIP_package.29}

                 .--\/--.                              .--\/--.                           .--\/--.
           M2 -> |01  40| -- VCC            /ROMSEL -> |01  40| -> WRAM /CE    CHR A13 <- |01  40| <- PPU A10
     WRAM /CE <- |02  39| -- NC             PRG /CE <- |02  39| -- VCC         CHR A14 <- |02  39| <- PPU A11
       CPU D7 -> |03  38| -> PRG /CE       WRAM /WE <- |03  38| -> WRAM +CE    CHR A12 <- |03  38| <- PPU A12
       CPU A0 -> |04  37| -> PRG A17        CPU A14 -> |04  37| <- CPU R/W   CIRAM A10 <- |04  37| -> CHR A10
       CPU D6 -> |05  36| <- CPU A13        CPU A13 -> |05  36| -> PRG A13     CHR A15 <- |05  36| -> CHR A16
       CPU D0 -> |06  35| -> PRG A18         CPU A0 -> |06  35| -> PRG A14     CHR A17 <- |06  35| -> CHR A11
       CPU D5 -> |07  34| -> PRG A14             M2 -> |07  34| -> PRG A15        /IRQ <- |07  34| -> WRAM /WE
       CPU D1 -> |08  33| -> PRG A16        PPU A12 -> |08  33| -> PRG A16     /ROMSEL -> |08  33| -> WRAM +CE
       CPU D4 -> |09  32| <- CPU A14           /IRQ <- |09  32| -> PRG A17         GND -- |09  32| -- GND
       CPU D2 -> |10  31| -> PRG A13      CIRAM A10 <- |10  31| -> PRG A18     CPU R/W -> |10  31| <- CPU D3
       CPU D3 -> |11  30| -> PRG A15        PPU A10 -> |11  30| -- NC          PRG A15 <- |11  30| <- CPU D2
     WRAM +CE <- |12  29| <- CPU R/W        PPU A11 -> |12  29| -> CHR A17     PRG A13 <- |12  29| <- CPU D4
     WRAM /WE <- |13  28| <- /ROMSEL         CPU D0 -> |13  28| -> CHR A16     CPU A14 -> |13  28| <- CPU D1
      CHR A11 <- |14  27| -> /IRQ            CPU D1 -> |14  27| -> CHR A15     PRG A16 <- |14  27| <- CPU D5
      CHR A16 <- |15  26| -> CHR A17         CPU D2 -> |15  26| -> CHR A14     PRG A14 <- |15  26| <- CPU D0
      CHR A10 <- |16  25| -> CHR A15         CPU D3 -> |16  25| -> CHR A13     PRG A18 <- |16  25| <- CPU D6
      PPU A12 -> |17  24| -> CIRAM A10       CPU D4 -> |17  24| -> CHR A12     CPU A13 -> |17  24| <- CPU A0
      PPU A11 -> |18  23| -> CHR A12         CPU D5 -> |18  23| -> CHR A11     PRG A17 <- |18  23| <- CPU D7
      PPU A10 -> |19  22| -> CHR A14         CPU D6 -> |19  22| -> CHR A10     PRG /CE <- |19  22| -> WRAM /CE
          GND -- |20  21| -> CHR A13            GND -- |20  21| <- CPU D7          VCC -- |20  21| <- M2
                 `------'                              `------'                           '------'
                AX5202P #1                      AX5202P #2 / MC-3 (NTDEC's?)                 88

                 .--\/--.                              .--\/--.                           .--\/--.             
      CHR A10 <- |01  40| -- VCC            PPU A10 -> |01  40| -- VCC              M2 -> |01  40| -- VCC   
      PPU A12 -> |02  39| -> CHR A16        PPU A11 -> |02  39| -> PRG A18?    CPU R/W -> |02  39| <- PPU A10
      PPU A11 -> |03  38| -> CHR A11        PPU A12 -> |03  38| -> PRG A17?    /ROMSEL -> |03  38| <- PPU A11
      PPU A10 -> |04  37| -> WRAM /WE       CHR A10 <- |04  37| -> PRG A16    WRAM +CE <- |04  37| <- PPU A12  
      CHR A13 <- |05  36| -> WRAM +CE       CHR A11 <- |05  36| -> PRG A15    WRAM /CE <- |05  36| <- CPU A0  
      CHR A14 <- |06  35| <- CPU D3         CHR A12 <- |06  35| -> PRG A14    WRAM /WE <- |06  35| <- CPU A13  
      CHR A12 <- |07  34| <- CPU D2         CHR A13 <- |07  34| -> PRG A13     PRG /CE <- |07  34| <- CPU A14  
    CIRAM A10 <- |08  33| <- CPU D4         CHR A14 <- |08  33| <- CPU D7       CPU D0 -> |08  33| -> /IRQ   
      CHR A15 <- |09  32| <- CPU D1         CHR A15 <- |09  32| <- CPU D6       CPU D1 -> |09  32| -> DELAYED M2
      CHR A17 <- |10  31| <- CPU D5         CHR A16 <- |10  31| <- CPU D5       CPU D2 -> |10  31| -> CIR A10  
         /IRQ -> |11  30| <- CPU D0        ?CHR A17 <- |11  30| <- CPU D4       CPU D3 -> |11  30| -> CHR A17  
      /ROMSEL -> |12  29| <- CPU D6         /ROMSEL -> |12  29| <- CPU D3       CPU D4 -> |12  29| -> CHR A15  
      CPU R/W -> |13  28| <- CPU A0         CPU R/W -> |13  28| <- CPU D2       CPU D5 -> |13  28| -> CHR A14  
      PRG A15 <- |14  27| <- CPU D7            /IRQ <- |14  27| <- CPU D1       CPU D6 -> |14  27| -> CHR A13  
      PRG A13 <- |15  26| -> WRAM /CE       CIR A10 <- |15  26| <- CPU D0       CPU D7 -> |15  26| -> CHR A12
      CPU A14 -> |16  25| <- M2             CPU A0  -> |16  25| ?? ?           PRG A13 <- |16  25| -> CHR A11  
      PRG A16 <- |17  24| -> PRG /CE        CPU A13 -> |17  24| -> PRG /CE     PRG A14 <- |17  24| -> CHR A10  
      PRG A14 <- |18  23| ??                CPU A14 -> |18  23| ?? ?           PRG A15 <- |18  23| -> CHR A16  
      PRG A18 <- |19  22| -> PRG A17             M2 -> |19  22| ?? ?           PRG A16 <- |19  22| -> PRG A18  
          GND -- |20  21| <- CPU A13            GND -- |20  21| ?? ?               GND -- |20  21| -> PRG A17  
                 '------'                              '------'                           '------'             
                   9112                                7903 9102                             T1

Notes:

- Those chips are fully compatible MMC3 clone ICs that can be found in many bootleg games (there might be minor differences like the PPU A12 edge on which the scanline counter is clocked).
- The connections to edge connector for single games are the same in both pirate and original version, although some pirate multicarts use those chips in non-standard way with different connections.
- There are reports that some lots of AX5202P contain a large number of factory seconds with the IRQ not working.
- AX5202P \#1 and AX5202P \#2 are confirmed to enable WRAM at \$6000-\$7FFF at power up but protect them from writes (during CPU write cycle to \$6000-\$7FFF when WRAM is protected, WRAM !CE and WRAM CE are not asserted). When RAM is disabled, open bus behavior is observed.
- NC seems to be not connected internally in both versions (multimeter diode test does not show any conducting voltage between NC and any other pins).
- AX5202P \#1 is the one you can still buy nowadays (it has AX5202P marking) (this applies to [MindKids' brand new DIP-40 AX5202Ps with AA6023 core inside](../SMD133_%28AA6023%29_pinout/) too, with just a few pin differences).
- AX5202P \#2 was found in at least one game - Doki Doki Yuuenchi bootleg (the chip does not have any markings).
- Boards with clone MMC3 IC sometimes use WRAM enable line for different purposes, usually for outer banking.

## References {#References}

- AX5202P \#1: [BBS](http://forums.nesdev.org/viewtopic.php?t=5280)
- AX5202P \#2: [BBS](https://forums.nesdev.org/viewtopic.php?f=9&t=15813&view=previous#p205854)
- 88 [BBS](https://forums.nesdev.org/viewtopic.php?t=16864)
- 9112 [BBS](http://forums.nesdev.org/viewtopic.php?p=156990#p156990)
- 7903 9102 [BBS](http://forums.nesdev.org/viewtopic.php?f=9&t=18602#p236010)
- T1 [\[2\]](http://forums.nesdev.org/viewtopic.php?f=9&t=23096)
