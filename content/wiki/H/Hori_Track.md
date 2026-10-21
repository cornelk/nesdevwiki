---
title: "Hori Track"
url: "/wiki/Hori_Track/"
description: "The Hori Track (TRK-7) is a Famicom trackball controller released by Hori. A prototype U.S. version was exhibited behind glass in Nintendo World, but it never reached stores."
wikiPage: "Hori_Track"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Pointing devices"
    url: "/category/Pointing_devices/"
---

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Hori-track.png"><img src="../../wiki-images/Hori-track.png" width="300" height="287" /></a>
<figcaption>Top view of trackball</figcaption>
</figure>

The Hori Track (**TRK-7**) is a [Famicom](../Family_Computer/) trackball controller released by Hori. A prototype U.S. version was [exhibited behind glass in Nintendo World](https://forums.nesdev.org/viewtopic.php?f=2&t=15184&p=183787#p183787), but it never reached stores.

Supported games:

- *Moero!! Pro Soccer*
- *Operation Wolf* (also supports the [zapper](../Zapper/))
- *Putt Putt Golf* (FDS)
- *Quarter Back Scramble*
- *US Championship V'Ball*

## Protocol {#Protocol}

Report byte 1 is the embedded [standard controller](../Standard_controller/).

Byte 2, MSB first:

    7654 3210
    |||| ++++- Axis 2, signed 4 bit, XOR with $F
    ++++------ Axis 1, signed 4 bit, XOR with $F

Byte 3, MSB first:

    7654 3210
    |||| ++++- Unknown (read and unused by games)
    ||++------ ID bits (1 or 2 depending on version)
    |+-------- Speed switch (0: Hi, 1: Lo)
    +--------- Rotation mode switch (0: R, 1: L)

In rotation mode L, Up on the Control Pad points up, axis 1 points down, and axis 2 points right. In rotation mode R, Up on the Control Pad points right, axis 1 points left, and axis 2 points down.

Unlike the rotation mode switch, the speed switch will alter the axis values before they are presented to the console.

The ID bits are in the same place as those of the [Four Score](../Four_player_adapters/) hub, which is also based on a Hori design. It is speculated that axis 2 can be used to distinguish them.

## See Also {#See_Also}

- [Hori Tracker](https://github.com/TakuikaNinja/HoriTracker) - Input display program.
- [Video demonstrating the Hori Track in *Operation Wolf*](https://youtu.be/GSIAxqK0UgE)

## References {#References}

- [Forum discussion](https://forums.nesdev.org/viewtopic.php?t=25358)
- [nocash's documentation](https://problemkaputt.de/everynes.htm#controllerstrackballandmouse)
