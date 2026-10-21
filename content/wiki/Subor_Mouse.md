---
title: "Subor Mouse"
url: "/wiki/Subor_Mouse/"
description: "The Subor Mouse is a third party accessory, used not only by Subor themselves for their later educational cartridges, but also by a number of third-party titles by other developers. It responds to…"
wikiPage: "Subor_Mouse"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Pointing devices"
    url: "/category/Pointing_devices/"
---

The **Subor Mouse** is a third party accessory, used not only by Subor themselves for their later educational cartridges, but also by a number of third-party titles by other developers. It responds to the same strobe signal as a [standard controller](../Standard_controller/), and the first eight bits returned are defined such that a standard controller's response will be that of a mouse that returns a movement amount of 1. As a result, the mouse pointer can be moved with a standard controller, and the movement of the Subor Mouse will be seen by non-mouse-aware programs as an equivalent standard controller response.

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Strobe

### <span id="Output_.28.244016_or_.244017_read.29"></span>Output (\$4016 or \$4017 read) {#Output_.28.244016_or_.244017_read.29}

    7  bit  0
    ---- ----
    xxxx xxxD
            |
            +- Serial data (bit 23->bit 22->...->bit 0)

### 24-bit response {#24-bit_response}

    222211111111110000000000
    321098765432109876543210
    ------------------------
    lrE.UDLR.XXXXXXX.YYYYYYY
    ||| |||| ||||||| +++++++- absolute amount of Y movement (0-127) if E=1
    ||| |||| +++++++--------- absolute amount of X movement (0-127) if E=1
    ||| |||+----------------- positive X movement detected (right)
    ||| ||+------------------ negative X movement detected (left)
    ||| |+------------------- positive Y movement detected (down)
    ||| +-------------------- negative Y movement detected (up)
    ||+---------------------- 1: movement amount greater than 1 or lesser than -1 detected, X/Y bits must be interpreted
    |+----------------------- 1: right mouse button pressed
    +------------------------ 1: left mouse button pressed

### Notes {#Notes}

- *GameStar Fun Educator꞉ 32-in-1* will only recognize mouse button presses if the E bit is set. This suggests that some models of the Subor Mouse may have that bit set all the time.
