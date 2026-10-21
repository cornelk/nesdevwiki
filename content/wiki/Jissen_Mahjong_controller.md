---
title: "Jissen Mahjong controller"
url: "/wiki/Jissen_Mahjong_controller/"
description: "This is a 21-button controller used by the following Capcom games:"
wikiPage: "Jissen_Mahjong_controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

This is a 21-button controller used by the following Capcom games:

- *Ide Yousuke Meijin no Jissen Mahjong*
- *Ide Yousuke Meijin no Jissen Mahjong 2*

## Button layout {#Button_layout}

<table class="wikitable" style="text-align:center">
<tbody>
<tr>
<td>A</td>
<td>B</td>
<td>C</td>
<td>D</td>
<td>E</td>
<td>F</td>
<td>G</td>
<td>H</td>
<td>I</td>
<td>J</td>
<td>K</td>
<td>L</td>
<td>M</td>
<td>Ⓝ</td>
</tr>
<tr>
<td colspan="5"></td>
<td>SEL</td>
<td>ST</td>
<td>カン</td>
<td>ポン</td>
<td>チー</td>
<td>リーチ</td>
<td>ロン</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>

The [カン (kan)](https://en.wikipedia.org/wiki/Japanese_Mahjong#Kan), [ポン (pon)](https://en.wikipedia.org/wiki/Japanese_Mahjong#Pon), [チー (chii)](https://en.wikipedia.org/wiki/Japanese_Mahjong#Ch.C4.AB), [リーチ (riichi)](https://en.wikipedia.org/wiki/Japanese_Mahjong#R.C4.ABchi) and [ロン (ron)](https://en.wikipedia.org/wiki/Japanese_Mahjong#Winning) buttons are Mahjong-related.

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xRRS
          |||
          ||+- Strobe
          |+-- Row selection (bit 0)
          +--- Row selection (bit 1)

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxx xxDx
           |
           +- Serial data

Reading the buttons is very similar to the [standard controller](../Standard_controller/), but one of three rows is selected for loading values *into* the shift register depending on the contents of the two Row selection pins while the Strobe bit is high.

After toggling the strobe bit from 1 to 0, the controller will return 8 bits worth of data (1 bit per read) based on the "row selection" bits.

    Row 3 returns: <empty>, ロン, リーチ, チー, ポン, カン, Start, Select
    Row 2 returns: H, G, F, E, D, C, B, A
    Row 1 returns: <empty>, <empty>, N, M, L, K, J, I
    Row 0 returns the bitwise OR of rows 1 and 2.

Like the standard controller, buttons return 1 when extant and held down, 0 otherwise.

The buttons are arranged as a 3x8 keyboard matrix similar to the Famicom keyboard. A variety of diodes, resistors, and an NPN BJT pull one of the rows low; the eight columns are then loaded into the same 4021 shift register as used on the standard controller. <sup><span id="cite_ref-reverse_1-0"></span>[\[1\]](#cite_note-reverse-1)</sup> The Serial Input to the 4021 is tied to ground; all subsequent reads beyond the first eight should return 0V (logic 1 due to the 74368 inside the Famicom).

Like the [Family BASIC Keyboard](../Family_BASIC_Keyboard/), there are no diodes to prevent [ghosting](https://en.wikipedia.org/wiki/Rollover_(key)#Key_jamming_and_ghosting), and pressing three buttons simultaneously could cause the incorrect detection of a 4th button press.

There are (at least) two hardware revisions of the PCB inside, [\[1\]](http://archive.is/r170z) but no behavioral differences between the two are known nor any believed to exist.

## References {#References}

1.  <span id="cite_note-reverse-1"></span>[↑](#cite_ref-reverse_1-0) [(?whose?) reverse engineered schematic](http://archive.is/k6jkR)

- [EveryNES§Jissen Mahjong controller specs](http://problemkaputt.de/everynes.htm#controllerskeypads)
- [Koi-Koi's pictures on the forum](https://forums.nesdev.org/viewtopic.php?p=245593#p245593)
