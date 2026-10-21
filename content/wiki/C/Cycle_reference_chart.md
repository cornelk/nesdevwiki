---
title: "Cycle reference chart"
url: "/wiki/Cycle_reference_chart/"
description: "The clock rate of various components in the NES differs between consoles in the USA and Europe due to the different television standards used (NTSC M vs. PAL B). The color encoding method used by the…"
wikiPage: "Cycle_reference_chart"
---

## Clock rates {#Clock_rates}

The **clock rate** of various components in the NES differs between consoles in the USA and Europe due to the different television standards used (NTSC M vs. PAL B). The color encoding method used by the NES (see [NTSC video](../NTSC_video/)) requires that the master clock frequency be six times that of the color subcarrier, but this frequency is about 24% higher on PAL than on NTSC. In addition, PAL has more scanlines per field and fewer fields per second than NTSC. Furthermore, the PAL CPU's master clock could have been divided by 15 to preserve the ratio between CPU and PPU speeds, but Nintendo chose to keep the [Johnson counter](https://en.wikipedia.org/wiki/Ring_counter) structure, which always has an even period, and divide by 16 instead.

So the main differences between the NTSC and PAL PPUs are depicted in the following table:

<table class="tabular" style="width:100%;">
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
</colgroup>
<thead>
<tr>
<th>Property</th>
<th>NTSC (2C02)</th>
<th>PAL (2C07)</th>
<th>"Dendy" PAL Famiclone</th>
<th>RGB (2C03/4/5)</th>
<th>Brazil Famiclone</th>
<th>Argentina Famiclone</th>
</tr>
</thead>
<tbody>
<tr>
<td>Master clock speed</td>
<td>21.477272 MHz ± 40 Hz<br />
236.25 MHz ÷ 11 by definition</td>
<td>26.601712 MHz ± 50 Hz<br />
26.6017125 MHz by definition</td>
<td>Like PAL</td>
<td>Like NTSC</td>
<td>21.453671 MHz<br />
3.067875 GHz ÷ 143 by definition</td>
<td>21.492338 MHz<br />
42.984675 MHz ÷ 2 by definition</td>
</tr>
<tr>
<td>CPU</td>
<td>Ricoh 2A03</td>
<td>Ricoh 2A07</td>
<td>UMC UA6527P</td>
<td>Ricoh 2A03</td>
<td colspan="2">UMC UA6527</td>
</tr>
<tr>
<td>CPU clock speed</td>
<td>21.47~ MHz ÷ 12 = 1.789773 MHz</td>
<td>26.60~ MHz ÷ 16 = 1.662607 MHz</td>
<td>26.60~ MHz ÷ 15 = 1.773448 MHz</td>
<td>Like NTSC</td>
<td>21.45~ MHz ÷ 12 = 1.787806 MHz</td>
<td>21.49~ MHz ÷ 12 = 1.791028 MHz</td>
</tr>
<tr>
<td><a href="../APU_Frame_Counter/">APU Frame Counter</a> rate</td>
<td>60.00 Hz</td>
<td>50.00 Hz<sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup></td>
<td>59.46 Hz<sup><span id="cite_ref-2"></span><a href="#cite_note-2">[2]</a></sup></td>
<td>Like NTSC</td>
<td>59.94 Hz</td>
<td>60.04 Hz</td>
</tr>
<tr>
<td>PPU</td>
<td>Ricoh 2C02</td>
<td>Ricoh 2C07</td>
<td>UMC UA6538</td>
<td>Ricoh 2C03, 2C04, 2C05</td>
<td>UMC UA6548</td>
<td>UMC UA6528P</td>
</tr>
<tr>
<td>Master clocks per PPU dot</td>
<td>4</td>
<td>5</td>
<td>5</td>
<td>4</td>
<td colspan="2" style="text-align: center;">4</td>
</tr>
<tr>
<td>PPU dots per CPU cycle</td>
<td>3</td>
<td>3.2</td>
<td colspan="4" style="text-align: center;">Like NTSC</td>
</tr>
<tr>
<td>CPU cycles per scanline</td>
<td>341 × 4÷12 = 113<sup>2</sup>⁄<sub>3</sub></td>
<td>341 × 5÷16 = 106<sup>9</sup>⁄<sub>16</sub></td>
<td>341 × 5÷15 = 113<sup>2</sup>⁄<sub>3</sub></td>
<td colspan="3" style="text-align: center;">Like NTSC</td>
</tr>
<tr>
<td>Height of picture</td>
<td>240 scanlines</td>
<td>239 scanlines</td>
<td>Like PAL</td>
<td>Like NTSC</td>
<td>thought to be like NTSC</td>
<td>thought to be like PAL</td>
</tr>
<tr>
<td>Nominal visible picture height (see <a href="../Overscan/">Overscan</a>)</td>
<td>224 scanlines</td>
<td>268 scanlines</td>
<td>Like PAL</td>
<td>Like NTSC</td>
<td>Like NTSC</td>
<td>Like PAL</td>
</tr>
<tr>
<td>"Post-render" blanking lines between end of picture and <a href="../NMI/">NMI</a></td>
<td>1 scanline</td>
<td>1 scanline</td>
<td>51 scanlines</td>
<td>1 scanline</td>
<td>1 scanline</td>
<td>51 scanlines</td>
</tr>
<tr>
<td>Length of vertical blanking after NMI</td>
<td>20 scanlines (≈ 2273 CPU cycles)</td>
<td>70 scanlines (≈ 7459 CPU cycles)</td>
<td colspan="4" style="text-align: center;">Like NTSC</td>
</tr>
<tr>
<td>Time during which <a href="../PPU_OAM/">OAM</a> can be written</td>
<td>Vertical or forced blanking</td>
<td>Only during first 24 scanlines after NMI (≈2557 CPU cycles)</td>
<td>Like NTSC</td>
<td colspan="3" style="text-align: center;">Like NTSC</td>
</tr>
<tr>
<td>"Pre-render" lines between vertical blanking and next picture</td>
<td colspan="6" style="text-align: center;">1 scanline</td>
</tr>
<tr>
<td>Total number of dots per frame</td>
<td>341 × 261 + 340.5 = 89341.5<br />
(pre-render line is one dot shorter in every odd frame)</td>
<td>341 × 312 = 106392</td>
<td>Like PAL</td>
<td>341 × 262 = 89342</td>
<td>unknown<br />
either like NTSC or RGB</td>
<td>Like PAL</td>
</tr>
<tr>
<td>Total number of CPU cycles per frame</td>
<td>89341.5 ÷ 3 = 29780.5</td>
<td>106392 ÷ 3.2 = 33247.5</td>
<td>106392 ÷ 3 = 35464</td>
<td>89342 ÷ 3 = 29780<sup>2</sup>⁄<sub>3</sub></td>
<td>unknown<br />
either like NTSC or RGB</td>
<td>Like PAL famiclone</td>
</tr>
<tr>
<td>Frame rate (vertical scan rate)</td>
<td>60.0988 Hz</td>
<td>50.0070 Hz</td>
<td>Like PAL</td>
<td>60.0985 Hz</td>
<td>60.03 Hz</td>
<td>50.5027 Hz</td>
</tr>
<tr>
<td>Color of top border</td>
<td colspan="6" style="text-align: center;">Always black ($0E)</td>
</tr>
<tr>
<td><a href="../NTSC_video/">Side and bottom borders</a></td>
<td><a href="../PPU_palettes/">Palette</a> entry at $3F00</td>
<td>Always black ($0E), <a href="../Overscan/#PAL">intruding on left and right 2 pixels and top 1 pixel of picture</a></td>
<td>Like PAL<sup><span id="cite_ref-3"></span><a href="#cite_note-3">[3]</a></sup></td>
<td>Like NTSC<sup><span id="cite_ref-4"></span><a href="#cite_note-4">[4]</a></sup></td>
<td>Unknown</td>
<td>Unknown, probably like PAL</td>
</tr>
<tr>
<td>Color emphasis<br />
(with correlating bit in <a href="../PPU_registers/">PPUMASK</a>)</td>
<td>Blue (D7), green (D6), red (D5)</td>
<td>Blue (D7), red (D6), green (D5)</td>
<td>Like PAL</td>
<td>Blue, green, red (full scale)</td>
<td colspan="2">Unknown</td>
</tr>
<tr>
<td>Other quirks</td>
<td>Early revisions cannot read back sprite or palette memory</td>
<td></td>
<td></td>
<td>Many</td>
<td colspan="2">poorly-researched</td>
</tr>
</tbody>
</table>

Some frequencies in the above table are rounded.

The 2C03, 2C04, and 2C05 PPUs were all found in Nintendo's [Vs. System](https://en.wikipedia.org/wiki/Nintendo_VS._System) and [PlayChoice-10](https://en.wikipedia.org/wiki/PlayChoice-10) (a.k.a. PC10 or PC-10) arcade systems. Famicom Titler, Famicom TVs, and RGB-modded NES consoles would use either the 2C03 or a 2C05 with glue logic to unswap \$2000 and \$2001. (Later RGB mods used a 2C02 in output mode and faked out all palette logic.)

The color emphasis bits on the PAL NES have their red and green bits in [PPUMASK](../PPU_registers/) swapped

The authentic NES sold in Brazil is an NTSC NES with an adapter board to turn the NTSC video into [PAL-M video](https://en.wikipedia.org/wiki/PAL-M), a variant of PAL using NTSC frequencies but PAL's color modulation.

Micro Genius is a clone of the Famicom, manufactured by TXC Corporation of Taiwan and sold under various brand names in the 50 Hz market.<sup><span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup> Among the best known brands is [Dendy](https://en.wikipedia.org/wiki/Dendy_(console)), distributed in Russia by Steepler, and the attention given by Russian reverse engineers to this clone has led to "Dendy" becoming a common name for all PAL Micro Genius-type famiclones. Its chipset (UA6527P+UA6538) is designed for compatibility with Famicom games, including games with CPU cycle counting mappers (e.g. [VRC4](../VRC2_and_VRC4/)) and games that use a cycle-timed NMI handler (e.g. *Balloon Fight*). This explains the faster CPU divider and longer post-render period vs. the authentic PAL NES.

To compensate for these differences, you can [detect the TV system](../Detect_TV_system/) at power-on.

## CPU cycle counts {#CPU_cycle_counts}

To make things easier for those programming on the NES, the below chart provides the number of CPU cycles that a particular PPU-oriented trait takes.

<table class="wikitable" style="width:35%">
<thead>
<tr>
<th>Property</th>
<th>NTSC</th>
<th>PAL</th>
<th>Dendy</th>
</tr>
</thead>
<tbody>
<tr>
<td>Scanline (341 pixels)</td>
<td>113<sup>2</sup>⁄<sub>3</sub></td>
<td>106<sup>9</sup>⁄<sub>16</sub></td>
<td>113<sup>2</sup>⁄<sub>3</sub></td>
</tr>
<tr>
<td>HBlank (85 pixels)</td>
<td>28<sup>1</sup>⁄<sub>3</sub></td>
<td>26<sup>9</sup>⁄<sub>16</sub></td>
<td>28<sup>1</sup>⁄<sub>3</sub></td>
</tr>
<tr>
<td>NMI to start of rendering</td>
<td>2273<sup>1</sup>⁄<sub>3</sub></td>
<td>7459<sup>3</sup>⁄<sub>8</sub></td>
<td>2273<sup>1</sup>⁄<sub>3</sub></td>
</tr>
<tr>
<td>Frame</td>
<td>29780.5*</td>
<td>33247.5</td>
<td>35464</td>
</tr>
<tr>
<td>PPU dots ÷ CPU cycles</td>
<td>3</td>
<td>3.2</td>
<td>3</td>
</tr>
<tr>
<td>OAM DMA</td>
<td colspan="3">513 (+1 if starting on CPU <a href="../DMA/#Cadence">get cycles</a>)</td>
</tr>
</tbody>
</table>

\* NTSC frame timing is 29780.5 cycles if rendering is enabled during the 20th scanline; 29780<sup>2</sup>⁄<sub>3</sub> otherwise.

## See also {#See_also}

- [PPU frame timing](../PPU_frame_timing/)
- [PPU rendering](../PPU_rendering/)
- [Cycle counting](../Cycle_counting/)

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) nesdev forum post by thefox: <http://forums.nesdev.org/viewtopic.php?p=160349#p160349>
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) nesdev forum post by Eugene.S: <http://forums.nesdev.org/viewtopic.php?p=174970#p174970>
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) nesdev forum post by Eugene.S: <https://forums.nesdev.org/viewtopic.php?p=173764#p173764>
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) nesdev forum post by lidnariq: <https://forums.nesdev.org/viewtopic.php?p=179705#p179705>
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) Post by feos at [TASVideos](http://tasvideos.org/forum/viewtopic.php?p=467169#467169) and [NESdev](https://forums.nesdev.org/viewtopic.php?p=216078#p216078)
