---
title: "Mappers with scanline IRQs"
url: "/category/Mappers_with_scanline_IRQs/"
description: "Mappers in this category have the ability to schedule an IRQ an arbitrary and continuously-adjustable number of scanlines in the future. This is triggered in response to a rendering signal such as…"
wikiPage: "Category_Mappers_with_scanline_IRQs"
type: category
wikiCategories:
  - name: "Mappers with IRQs"
    url: "/category/Mappers_with_IRQs/"
---

Mappers in this category have the ability to schedule an IRQ an arbitrary and continuously-adjustable number of scanlines in the future. This is triggered in response to a rendering signal such as [PPU A12](../../wiki/PPU_pinout/).

Mappers with configurable [cycle-based IRQs](../Mappers_with_cycle_IRQs/) can simulate a scanline-counter by waiting an equivalent number of CPU cycles, but have less robust horizontal timing, since they will fire relative to their start time, rather than relative to a rendering signal. The [VRC IRQ](../../wiki/VRC_IRQ/) has a convenient mode for this.

Mappers with a fixed delay (but still clocked by the PPU) should go in the [fixed-timing category](../Mappers_with_fixed-timing_scanline_IRQs/) instead.
