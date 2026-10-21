---
title: "SNES controller"
url: "/wiki/SNES_controller/"
description: "The Super NES Controller (SHVC-005, SNS-005, SNS-102) is very similar to the NES's standard controller, with a collection of digital inputs that are latched and read in series. With a suitable…"
wikiPage: "SNES_controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Super NES"
    url: "/category/Super_NES/"
---

[<img src="../../wiki-images/SNES-Controller-Flat.jpg" width="400" height="195" />](../../wiki-images/SNES-Controller-Flat.jpg)

[](../../wiki-images/SNES-Controller-Flat.jpg)

[Standard SNES controller](../../wiki-images/SNES-Controller-Flat.jpg)

[](../../wiki-images/SNES-Controller-Flat.jpg)

[The **Super NES Controller** (SHVC-005, SNS-005, SNS-102) is very similar to the NES's](../../wiki-images/SNES-Controller-Flat.jpg) [standard controller](../Standard_controller/), with a collection of digital inputs that are latched and read in series. With a [suitable adapter](../Controller_port_pinout/), reading the SNES controller is like reading an NES controller with 4 extra buttons.

The 16 bit report will read, in order:

     0 - B
     1 - Y
     2 - Select
     3 - Start
     4 - Up
     5 - Down
     6 - Left
     7 - Right

     8 - A
     9 - X
    10 - L
    11 - R
    12 - 0
    13 - 0
    14 - 0
    15 - 0

Note that the first 8 values map directly to the original NES controller's 8 inputs (SNES **B** = NES **A**, and SNES **Y** = NES **B**). The [Virtual Boy controller](../Virtual_Boy_controller/) likewise returns button states in nearly the same order, with the right Control Pad corresponding to ABXY on the Super NES.

The last 4 bits will always read 0 from a standard SNES controller. Other values here may indicate other devices,<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> such as a [mouse](../Mouse/) or a Virtual Boy controller.

After the 16 bit report, subsequent bits will read as 1.

## Games {#Games}

An incomplete list of NES games with special support for SNES controllers:

- *Spook-O-Tron*
- *Nova the Squirrel*<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>
- *Full Quiet*<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

## See also {#See_also}

- [Controller port pinout](../Controller_port_pinout/)
- [Controller reading](../Controller_reading/)

## References {#References}

- [Standard controller](https://snes.nesdev.org/wiki/Standard_controller) at SNESdev Wiki

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [FullSNES:](http://problemkaputt.de/fullsnes.htm#snescontrollershardwareidcodes) SNES Controller Hardware ID Codes
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Nova the Squirrel (WIP)](https://forums.nesdev.org/viewtopic.php?p=198552#p198552) - forum post discussing SNES controller support.
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) [Tweet by E.C. Myers](https://twitter.com/ecmyers/status/1600574149092589569)
