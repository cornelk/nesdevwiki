---
title: "Program compatibility"
url: "/wiki/Program_compatibility/"
description: "Homebrew development is as subject to bugs as old software was, but many suffer from additional compatibility problems with the NES hardware due to being tested on emulators exclusively. It is often…"
wikiPage: "Program_compatibility"
---

*This page describes defects in homebrew games. For defects in games prior to 1996, see [Game bugs](../Game_bugs/).*

Homebrew development is as subject to bugs as old software was, but many suffer from additional [compatibility problems](../Accuracy/) with the NES hardware due to being tested on emulators exclusively. It is often the case that early homebrew ROMs will not run correctly on hardware. Since the release of the [PowerPak](../PowerPak/) in 2008, and general improvement in emulators over time, these problems have become less frequent.

For a partial list of homebrew projects, see: [Projects](../Projects/)

## Common compatibility problems {#Common_compatibility_problems}

This is a list of issues that commonly appear due to lack of popular support in emulators, or sometimes even hardware limitations of flash-carts like the [PowerPak](../PowerPak/). Many emulators are merely trying to be compatible with existing games, rather than accurately reflecting the hardware, making them inadequate for verifying the correctness of new software.

- Use of DPCM samples causes a [conflict](../APU_DMC/) with controller and PPU reads during sample playback.
- Writes to [PPU registers](../PPU_registers/) outside of [VBlank](../The_frame_and_NMIs/#VBlank.2C_Rendering_Time.2C_and_NMIs) with rendering enabled conflicts with internal use of the PPU address, causing graphical corruption.
- Failure to initialize registers not guaranteed by the [CPU power up state](../CPU_power_up_state/) or the [PPU power up state](../PPU_power_up_state/).
- Failure to initialize RAM, nametables, or [mapper](../Mapper/) registers. (Few mappers have guaranteed power up states.)
- Failure to delay use of the PPU after power/reset to avoid conflicts with its [warm up state](../PPU_power_up_state/).

## References {#References}
