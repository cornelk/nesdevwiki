---
title: "PPU"
url: "/wiki/PPU/"
description: "The NES PPU, or Picture Processing Unit, generates a composite video signal with 240 lines of pixels, designed to be received by a television. When the Famicom chipset was designed in the early…"
wikiPage: "PPU"
---

The NES PPU, or Picture Processing Unit, generates a composite video signal with 240 lines of pixels, designed to be received by a television. When the Famicom chipset was designed in the early 1980s, it was considered quite an advanced 2D picture generator for video games.

It has its own address space, which typically contains 10 kilobytes of memory: 8 kilobytes of ROM or RAM on the Game Pak (possibly more with one of the common [mappers](../Mapper/)) to store the shapes of background and sprite tiles, plus 2 kilobytes of RAM in the console to store a map or two. Two separate, smaller address spaces hold a palette, which controls which colors are associated to various indices, and OAM (Object Attribute Memory), which stores the position, orientation, shape, and color of the sprites, or independent moving objects. These are internal to the PPU itself, and while the palette is made of static memory, OAM uses dynamic memory (which will slowly decay if the PPU is not rendering).

### <span id="Programmer.27s_reference"></span>Programmer's reference {#Programmer.27s_reference}

- [Registers](../PPU_registers/)
- [Pattern tables](../PPU_pattern_tables/) (tile graphics for background and sprites)
- Background graphics
  - [Nametables](../PPU_nametables/)
  - [Attribute tables](../PPU_attribute_tables/)
- [OAM](../PPU_OAM/) (sprites)
- [Palettes](../PPU_palettes/)
- [Memory map](../PPU_memory_map/)

- **[All PPU documentation on one page](../PPU_programmer_reference/)** (printer-friendly)

### Hardware behaviors {#Hardware_behaviors}

- [Frame timing](../PPU_frame_timing/)
- [Power up state](../PPU_power_up_state/)
- [NMI](../NMI/)
- [Clock rate](../Cycle_reference_chart/#Clock_rates) and other NTSC/PAL/Dendy differences
- [NTSC video](../NTSC_video/)
- [Scrolling](../PPU_scrolling/)
- [Rendering](../PPU_rendering/)
- [Sprite evaluation](../PPU_sprite_evaluation/)
- [Sprite priority](../PPU_sprite_priority/)
- [PPU signals](../PPU_signals/)
- [Overscan](../Overscan/)
- [PPU pinout](../PPU_pinout/)
- [NTSC PPU frame timing diagram](../../wiki-images/Ppu.svg)
- [Visual 2C02](../Visual_2C02/): A hardware-level PPU simulator
- [List of known PPU versions and variants](../PPU_variants/)

### Notes {#Notes}

- The [NTSC video](../NTSC_video/) signal is made up of 262 scanlines, and 20 of those are spent in vblank state. After the program has received an NMI, it has about 2270 cycles to update the palette, sprites, and nametables as necessary before rendering begins.
- On NTSC systems, the PPU divides the master clock by 4 while the CPU uses the master clock divided by 12. Since both clocks are fed off the same master clock, this means that there are **exactly** three PPU ticks per CPU cycle, with no drifting over time (though the clock alignment might vary depending on when you press the Reset button).
- On PAL systems, the PPU divides the master clock by 5 while the CPU uses the master clock divided by 16. As a result, there are exactly 3.2 PPU ticks per CPU cycle.

### See also {#See_also}

- [2C02 technical reference](http://nesdev.org/2C02%20technical%20reference.TXT) by Brad Taylor. (Pretty old at this point; information on the wiki might be more up-to-date.)
