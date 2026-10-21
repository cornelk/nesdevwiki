---
title: "Oeka Kids tablet"
url: "/wiki/Oeka_Kids_tablet/"
description: "The Oeka Kids tablet is a Famicom accessory from Bandai that resembles a drawing tablet, and plugs into the expansion port. There are two known games that use this peripheral, using iNES Mapper 096,…"
wikiPage: "Oeka_Kids_tablet"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Pointing devices"
    url: "/category/Pointing_devices/"
---

The Oeka Kids tablet is a Famicom accessory from Bandai that resembles a drawing tablet, and plugs into the expansion port. There are two known games that use this peripheral, using [iNES Mapper 096](../INES_Mapper_096/), both of which contain a variety of activities, such as painting, drawing lessons, hiragana lessons, and a variety of minigames such as mazes and sliding puzzles.

Very little information about this accessory is available, and only [basic reverse engineering](http://forums.nesdev.org/viewtopic.php?p=19454#19454) has been performed so far. As such, this information **may be incorrect**, but it seems to be acceptable for the two commercial games that use it and two test ROMs.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)<span id="cite_ref-little_things_oekakids_2-0"></span>[\[2\]](#cite_note-little_things_oekakids-2)</sup>

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxAS
           ||
           |+- Strobe (0 = Latch, 1 = Read mode)
           +-- Advance to next bit

Bits can be read only while S is 1. It advances to the next bit when S is 1 and A transitions from 0 to 1.

It works to write \$01 then \$00, read the [standard controllers](../Standard_controller/), and then write \$01 before reading the tablet.

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxx DSxx
         ||
         |+--- 1 if strobe is 0 or advance is 1, 0 otherwise
         +---- 0 if strobe is 0, inverted report bit otherwise

To read one bit, the program must turn on advance (\$4016=\$03), wait for \$4017 bit 2 to become 0, turn off advance (\$4016=\$01), and wait for \$4017 bit 2 to become 1. If the game does not wait for the tablet's slow microcontroller to acknowledge each edge, the data is incorrect. The MCU can take upwards of 200 cycles to acknowledge each edge, and reading all 18 bits takes about 67 scanlines.<sup><span id="cite_ref-little_things_oekakids_2-1"></span>[\[2\]](#cite_note-little_things_oekakids-2)<span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

The serial data is returned *most significant bit first*, and inverted (including the touch and click bits).

    $4017.3
       ^
       |  <-- <-- <-- <--
       XXXXXXXXYYYYYYYYBA
       ||||||||||||||||||
       |||||||||||||||||+- Click (gray "space bar" at bottom of tablet is held)
       ||||||||||||||||+-- Stylus is touching tablet
       ||||||||++++++++--- Stylus Y, scaled to 0-255
       ++++++++----------- Stylus X, scaled to 0-239

The fact that X and Y are scaled backwards is *not* an error - the games which use this mapper rescale the coordinates by multiplying X by 256/240 and multiplying Y by 240/256.

Note that the Stylus X and Y values are nonsense when the stylus is not touching the tablet.

Small Y coordinates (less than \$20 or so) correspond to regions under the bezel, with 7 demarcated touch regions. The exact bounding boxes have not yet been verified to be consistent across multiple different instances of the hardware.<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)</sup>

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) ["keytest" by Zorchenhimer](https://github.com/zorchenhimer/keytest/blob/master/tablet.asm)
2.  <span id="cite_note-little_things_oekakids-2"></span>↑ <sup>[2.0](#cite_ref-little_things_oekakids_2-0)</sup> <sup>[2.1](#cite_ref-little_things_oekakids_2-1)</sup> ["Oeka Kids test" by Damian Yerrick](https://github.com/pinobatch/little-things-nes/tree/master/oekakids)
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) ["Oeka Kids tablet test"](https://forums.nesdev.org/viewtopic.php?p=308617#p308617) by tepples et al. on NESdev
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) <https://forums.nesdev.org/viewtopic.php?p=299893#p299893>
