---
title: "Registers"
url: "/wiki/Registers/"
description: "Registers provide internal state storage for a device."
wikiPage: "Registers"
---

Registers provide internal state storage for a device.

These may store a number, like the accumulator register in the [CPU](../CPU/), or they may store other behaviour states, like the colour output of the [PPU](../PPU/). The various cartridge [mappers](../Mapper/) have registers that control ~~banking~~, as well as other devices (e.g. [IRQ](../IRQ/) timer).

A register might be entirely internal, or it might have a memory mapped address to provide external access via the CPU. Such registers are often referred to by this write address.

See:

- [CPU registers](../CPU_registers/)
- [PPU registers](../PPU_registers/)
- [APU registers](../APU/)
- [2A03 register map](../2A03/)
