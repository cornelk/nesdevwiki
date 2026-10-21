---
title: "Macro Winners Mouse"
url: "/wiki/Macro_Winners_Mouse/"
description: "A mouse that came with cartridges and keyboard famiclones manufactured by Macro Winners, also known as Belsonic, usually using their GameStar branding. GameStar-branded cartridges that did not come…"
wikiPage: "Macro_Winners_Mouse"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Pointing devices"
    url: "/category/Pointing_devices/"
---

A mouse that came with cartridges and keyboard famiclones manufactured by *Macro Winners*, also known as *Belsonic*, usually using their *GameStar* branding. *GameStar*-branded cartridges that did not come bundled with a keyboard famiclone tend to use the [Subor Mouse](../Subor_Mouse/) instead.

Example games:

- *Smart Genius*
- *Магистр Гений 2 Говорящий Картридж*
- *Мооспирв Обучающий Компьютер 2000* (a.k.a. Educational Computer 2000, though there are others with that title)

The *Mega Book* Mouse is virtually identical to the Macro Winners Mouse; the only difference is that the two bits that indicate the "byte number" are in a different bit position.

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Strobe

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxx xxxD
            |
            +- Serial data (bit 7->bit 6->...->bit 0)

The mouse returns either 1-byte or 3-byte responses based on the size of the x/y axis movement since the last read. If the movement value for both the X and Y axis is between -1 and 1 (inclusively), the mouse returns a 1-byte response. Otherwise, a 3-byte response is sent. In this case, the strobe bit must be toggled 3 times in a row to read the entire response (e.g read \$4017 8 times, toggle strobe bit, read 8 more times, etc.)

#### 1-byte response format {#1-byte_response_format}

Macro Winners:

    LRXXYY00
    ||||||||
    ||||||++- Always 0
    ||||++--- Y movement (0: no movement, 1 or 2 = went down 1 unit, 3 = went up 1 unit)
    ||++----- X movement (0: no movement, 1 or 2 = went right 1 unit, 3 = went left 1 unit)
    |+------- Right mouse button (1: pressed)
    +-------- Left mouse button (1: pressed)

The mouse button order is opposite that of the [Super NES Mouse](../Super_NES_Mouse/).

Mega Book:

    LR00XXYY
    ||||||||
    ||||||||
    ||||||++--- Y movement (0: no movement, 1 or 2 = went down 1 unit, 3 = went up 1 unit)
    ||||++----- X movement (0: no movement, 1 or 2 = went right 1 unit, 3 = went left 1 unit)
    ||++----- Always 0
    |+------- Right mouse button (1: pressed)
    +-------- Left mouse button (1: pressed)

#### 3-byte response format {#3-byte_response_format}

Macro Winners:

**Byte 1**

    LRSXTY01
    ||||||||
    ||||||++- Always 01
    |||||+--- Y movement (Bit 4)
    ||||+---- Y movement direction (1: up, 0: down)
    |||+----- X movement (Bit 4)
    ||+------ X movement direction (1: left, 0: right)
    |+------- Right mouse button (1: pressed)
    +-------- Left mouse button (1: pressed)

**Byte 2**

    --XXXX10
      ||||||
      ||||++- Always 10
      ++++--- X movement (Bits 0-3)

**Byte 3**

    --YYYY11
      ||||||
      ||||++- Always 11
      ++++--- Y movement (Bits 0-3)

Mega Book:

**Byte 1**

    LR01SXTY
    ||||||||
    ||||||||
    |||||||+--- Y movement (Bit 4)
    ||||||+---- Y movement direction (1: up, 0: down)
    |||||+----- X movement (Bit 4)
    ||||+------ X movement direction (1: left, 0: right)
    ||++----- Always 01
    |+------- Right mouse button (1: pressed)
    +-------- Left mouse button (1: pressed)

**Byte 2**

    --10XXXX
      ||||||
      ||||||
      ||++++--- X movement (Bits 0-3)
      ++------- Always 10

**Byte 3**

    --11YYYY
      ||||||
      ||||||
      ||++++--- Y movement (Bits 0-3)
      ++------- Always 11

## References {#References}

- [Nocash's EveryNES documentation](http://problemkaputt.de/everynes.htm#controllerstrackballandmouse) (incorrectly calls it the "Subor Mouse")
