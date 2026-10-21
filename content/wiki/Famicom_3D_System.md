---
title: "Famicom 3D System"
url: "/wiki/Famicom_3D_System/"
description: "The Famicom 3D System is a Famicom expansion port device that uses active LCD shutter glasses to provide 3D visuals. The headset (HVC-031) plugs into a control box (HVC-032) using a 3.5mm connector.…"
wikiPage: "Famicom_3D_System"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/Famicom-3D-System.jpg"><img src="../../wiki-images/Famicom-3D-System.jpg" width="400" height="192" /></a>
<figcaption>Famicom 3D System</figcaption>
</figure>

The Famicom 3D System is a Famicom expansion port device that uses active LCD shutter glasses to provide 3D visuals. The headset (HVC-031) plugs into a control box (HVC-032) using a 3.5mm connector. The control box has two such ports, allowing up to two headsets, and passes through a Famicom expansion port so other controllers can be used. Compatible games manually control which eye sees which frame, alternating to give each eye a unique perspective that creates a 30 Hz stereoscopic image. While these glasses do not require a CRT TV to correctly function, they do require very low latency and some kind of impulse display such as black frame insertion.

The 3D System is supported by the following Famicom games:

- *Attack Animal Gakuen* (1987)
- *Cosmic Epsilon* (1989)
- *Falsion* (1987)
- *Famicom Grand Prix II: 3D Hot Rally* (1988)
- *Fuuun Shourin Ken: Ankoku no Maou* (1988)
- *Highway Star* (1987)
- *JJ: Tobidase Daisakusen Part II* (1987)
- *Tobidase Daisakusen* (1987)

In these games, the 3D mode can be toggled with the select button except in *Fuuun Shourin Ken: Ankoku no Maou*, where the mode is enabled by holding A while pressing start on the title screen.

Note that the 3D System was only released in Japan. Versions of these games released internationally (*3-D WorldRunner* and *Rad Racer*) instead achieve 3D with an anaglyph mode using passive red and cyan glasses.

### <span id="Output_.28.244016_write.29"></span>Output (\$4016 write) {#Output_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxFx
           |
           +--- 0: Next field of video is for left eye (left eye lens is transparent; right eye lens is opaque)
                1: Next field of video is for right eye

Because the Famicom only generated images at 60Hz, each resultant image will flicker at 30Hz.

### Hardware {#Hardware}

The glasses used here take the same signal as Sega's Master System 3D and several other contemporary active shutter systems. There is a 3.5mm headphone jack, with the same mapping as headphones (tip to sleeve: left eye; ring to sleeve: right eye).

The "active shutter" glasses themselves use this convention: 10V peak, 20V peak-to-peak, indicates that the corresponding eye should be opaque. 0V indicates that it should become transparent. No DC component is ever permitted: that will damage the liquid crystal displays.

The adapter itself (HVC-032) contains a JRC4558 dual opamp and an [proprietary logic chip made by Sharp labeled 'IX0043AE'](../SHARP_IX0043AE_pinout/). This IX0043AE emits a clock for other hardware to generate 10V, -10V, and somehow converts the above-described control into signals for the two eyes; specifics are not known.

### See Also {#See_Also}

- <https://forums.nesdev.org/viewtopic.php?t=20036> (Krzysiobal's reverse engineering of HVC-032)
- <https://nicole.express/2023/the-third-dimension-and-the-eight-bits.html>
- <http://cmpslv3.stars.ne.jp/Mark3/M33d.png> (Enri's schematic of Sega's 3-D glasses adapter)
