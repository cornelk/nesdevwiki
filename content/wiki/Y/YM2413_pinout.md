---
title: "YM2413 pinout"
url: "/wiki/YM2413_pinout/"
description: "Yamaha YM2413 or ??? K-663A: 18-pin 0.3\" PDIP (used in Mapper 515). Also used as a cheaper and similar-sounding substitute for the VRC7 in the TNS-HFX4."
wikiPage: "YM2413_pinout"
wikiCategories:
  - name: "Audio"
    url: "/category/Audio/"
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Yamaha YM2413 or ??? K-663A: 18-pin 0.3" PDIP (used in [Mapper 515](../NES_2.0_Mapper_515/)). Also used as a cheaper and similar-sounding substitute for the [VRC7](../VRC7_audio/) in the [TNS-HFX4](http://www2s.biglobe.ne.jp/~tns/nr26730801.html).

## Pinout {#Pinout}

    YM2413 / K-663A                     UM3567 / U3567 / HTC-3567 / 3567 clones
             .--\/--.                            .--\/--.
      GND -- |01  18| <> D1               GND -- |01  24| -- n/c
       D2 <> |02  17| <> D0                D2 <> |02  23| <> D1
       D3 <> |03  16| -- +5V               D3 <> |03  22| <> D0
       D4 <> |04  15| -> RHYTHM OUT        D4 <> |04  21| -- +5V
       D5 <> |05  14| -> MELODY OUT       n/c -- |05  20| -- GND
       D6 <> |06  13| <- /RESET            D5 <> |06  19| -> RHYTHM OUT
       D7 <> |07  12| <- /CE               D6 <> |07  18| -> MELODY OUT
      XIN -> |08  11| <- R/W               D7 <> |08  17| -- n/c
     XOUT <- |09  10| <- A0               GND -- |09  16| <- /RESET
             '------'                     XIN -> |10  15| <- /CE
                                         XOUT <- |11  14| <- R/W
                                          n/c -- |12  13| <- A0
                                                 '------'

## Signal descriptions {#Signal_descriptions}

- +5V, GND: power supply
- D0-D7: bi-directional data bus
- /RESET: also called /IC in datasheet, resets all chip registers to 0 when this pin is low
- /CE: chip select
- R/W: also called /WE in datasheet (low = write, high = read)
- A0: address (low) / data write (high) control
- RHYTHM OUT, MELODY OUT: analog audio outputs
- XIN, XOUT: clock in/out

Sources: [\[1\]](https://forums.nesdev.org/viewtopic.php?t=16124), [\[2\]](http://www.alldatasheet.pl/datasheet-pdf/pdf/1135013/YAMAHA/YM2413.html), [\[3\]](https://www.msxpro.com/fm.html)
