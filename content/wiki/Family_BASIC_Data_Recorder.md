---
title: "Family BASIC Data Recorder"
url: "/wiki/Family_BASIC_Data_Recorder/"
description: "The Family BASIC Data Recorder (HVC-008) is a device manufactured by Matsushita/Panasonic for Nintendo and was released sometime in 1984 only in Japan. The device serves as an addition to the Family…"
wikiPage: "Family_BASIC_Data_Recorder"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
  - name: "Hardware"
    url: "/category/Hardware/"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/Nintendo-Family-BASIC-Data-Recorder.jpg"><img src="../../wiki-images/Nintendo-Family-BASIC-Data-Recorder.jpg" width="250" height="250" /></a>
<figcaption>The Family BASIC Data Recorder device</figcaption>
</figure>

The **Family BASIC Data Recorder** (HVC-008) is a device manufactured by Matsushita/Panasonic for Nintendo and was released sometime in 1984 only in Japan. The device serves as an addition to the [Family BASIC Keyboard](../Family_BASIC_Keyboard/) to read and write data from BASIC programs created by users. Several Famicom titles also used it to save its own data. The device itself can be substituted with any compact cassette player device or other audio capable devices that plug into the 3.5mm monaural phone jacks.

Example Titles:

- *Family BASIC*
- *Family BASIC V3*
- *Arkanoid II*
- *Castle Excellent*
- *Excitebike*
- *Lode Runner*
- *Mach Rider*
- *Wrecking Crew*

## Hardware interface {#Hardware_interface}

The [Family BASIC Keyboard](../Family_BASIC_Keyboard/) provides two ⅛" (3.5mm) monaural phone jacks which can be plugged into the data recorder. Other [expansion port](../Expansion_port/#Famicom) devices may also provide this interface (e.g. Hori SD Station, ASCII Stick II Turbo, etc.). [This circuit](http://nesdev.org/tapedrv.PNG) will provide the same interface on the NES.

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xExS
          | |
          | +- 1-bit DAC audio to audio cassette
          +--- When 0, force audio readback to always read as binary 0 (5V)

The audio to the cassette recorder goes through a first-order highpass at 100Hz and is attenuated to 5mVPP at the input to the recorder.

### <span id="Output_.28.244016_read.29"></span>Output (\$4016 read) {#Output_.28.244016_read.29}

    7  bit  0
    ---- ----
    xxxx xxAx
           |
           +-- 1-bit ADC audio from audio cassette

Because of how magnetic tape works, playing back the tape will produce a signal that is the lowpassed derivative of the original. Then this audio from the cassette recorder goes through a highpass with corner frequency of 800Hz before being discretized. In simulation, square waves of frequency 600 Hz up to the bandwidth of the tape appear to be recovered by this processing.

For an emulator, just play back the same bit stream as it was emitted.

## Software {#Software}

It is not known whether Family BASIC uses [Kansas City Standard](https://en.wikipedia.org/wiki/Kansas_City_Standard) encoding, [Bell 103](https://en.wikipedia.org/wiki/Bell_103_modem) or [202](https://en.wikipedia.org/wiki/Bell_202_modem), or some other arbitrary home-grown convention for encoding the audio on the tape. Castle Excellent's recorder handling code mostly exists between \$8000 and \$80FE, and provides save games using the 1200 baud ('CUTS') and *bit-reversed* variant of Kansas City Standard. Wrecking Crew uses the [tape format](https://original.sharpmz.org/tapeprob01.htm) from the [Sharp MZ](https://en.wikipedia.org/wiki/Sharp_MZ) personal computer.

For homebrew use, Chris Covell wrote a set of KCS [-generating](http://www.chrismcovell.com/TapeDump_Controls.html) and [-decoding](http://forums.nesdev.org/viewtopic.php?p=111232#p111232) routines, which could easily be adapted to work with the data recorder.

## References {#References}

Reverse-engineered schematics by [Enri](http://cmpslv3.stars.ne.jp/):

- <http://cmpslv3.stars.ne.jp/Famic/Fambas.htm>
- Also available [here](http://atariusa.com/Famicom_Schematics/FC%20Family%20Basic%20Keyboard%20and%20Sound%20Circuit%20Schematic.png) and [here](http://atariusa.com/Famicom_Schematics/FC%20Keyboard%20Key%20Schematic.png)
- [ccovell identified Wrecking Crew's encoding](http://forums.nesdev.org/viewtopic.php?p=109242#p109242)
