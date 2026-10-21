---
title: "Yuxing Mouse"
url: "/wiki/Yuxing_Mouse/"
description: "The 裕兴 (Yùxìng) Mouse is a nearly bog-standard serial mouse that can be connected to either controller port. The only difference is that the X and Y axes are swapped compared to the normal…"
wikiPage: "Yuxing_Mouse"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Pointing devices"
    url: "/category/Pointing_devices/"
---

The **裕兴 (Yùxìng) Mouse** is a nearly bog-standard serial mouse that can be connected to either controller port. The only difference is that the X and Y axes are swapped compared to the normal serial mouse protocol, most likely for purposes of vendor lock-in.

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Inverted strobe

The mouse is read by first writing 0, then 1, to \$4016 -- the opposite of what is written to read a standard controller.

### <span id="Output_.28.244016_or_.244017_read.29"></span>Output (\$4016 or \$4017 read) {#Output_.28.244016_or_.244017_read.29}

    7  bit  0
    ---- ----
    xxxx xxxD
            |
            +- Inverted serial data

### Response {#Response}

The serial data consists of three nine-bit packets: one 0 start bit, seven data bits, one 1 stop bit.

    #1 [0 1LRXXYY 1]
        | ||||||| +- Stop bit
        | |||||++--- Y movement, bits 7..6
        | |||++----- X movement, bits 7..6
        | ||+------- 1: right mouse button pressed
        | |+-------- 0: left mouse button pressed
        | +--------- Synchronization, indicates first of three packets
        +----------- Start bit

    #2 [0 0YYYYYY 1]
        |  |||||| +- Stop bit
        |  ++++++--- Y movement, bits 5..0
        +----------- Start bit

    #3 [0 0XXXXXX 1]
        |  |||||| +- Stop bit
        |  ++++++--- X movement, bits 5..0
        +----------- Start bit
