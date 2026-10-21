---
title: "Port test controller"
url: "/wiki/Port_test_controller/"
description: "The port test controller is a test device with two standard controllers that also provides a fixed signature on D3 and D4 of both joypads via an NES-INP-01 board. The device plugs into both…"
wikiPage: "Port_test_controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The port test controller is a test device with two standard controllers that also provides a fixed signature on D3 and D4 of both joypads via an NES-INP-01 board. The device plugs into both controller ports using a single two-plug connector, like the [Four Score](../Four_player_adapters/). It was used in factories for validating NES controller port input.

## <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    .... ...S
            |
            +- Controller shift register strobe

This matches the normal strobe behavior used by the [standard controller](../Standard_controller/#Input_.28.244016_write.29) and applies to all shift registers in the device.

## <span id="Output_.28.244016.2F4017_read.29"></span>Output (\$4016/4017 read) {#Output_.28.244016.2F4017_read.29}

    $4016:
    7  bit  0
    ---- ----
    xxx4 3xxD
       | |  |
       | |  +- Joypad 1 standard controller status bit
       | +---- Joypad 1 D3 signature
       +------ Joypad 1 D4 signature

    $4017:
    7  bit  0
    ---- ----
    xxx4 3xxD
       | |  |
       | |  +- Joypad 2 standard controller status bit
       | +---- Joypad 2 D3 signature
       +------ Joypad 2 D4 signature

The results from the device's two standard controllers are returned in bit 0 of the corresponding register, as normal. Bits 3 and 4 of each register return a signature from a shift register. The 4 signatures use only 2 shift registers in total, with 1 shift register handling each joypad register. D4 taps the last bit, Q8, of the shift register as normal, while D3 taps Q6, causing it to skip the first two bits, returning now what D4 will return two reads later.

    Joypad 1 signature:
    0  - (Always 1)  <-- D4 starts here
    1  - (Always 0)
    2  - (Always 1)  <-- D3 starts here
    3  - (Always 0)
    4  - (Always 1)
    5  - (Always 0)
    6  - (Always 1)
    7  - (Always 0)

    8+ - (Always 1)

    Joypad 2 signature:
    0  - (Always 0)  <- D4 starts here
    1  - (Always 1)
    2  - (Always 0)  <- D3 starts here
    3  - (Always 1)
    4  - (Always 0)
    5  - (Always 1)
    6  - (Always 0)
    7  - (Always 1)

    8+ - (Always 0)

## References {#References}

- [Forum post:](https://forums.nesdev.org/viewtopic.php?f=2&t=8438) Port Test Cart - How does this work?
