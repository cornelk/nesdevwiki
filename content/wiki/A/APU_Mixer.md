---
title: "APU Mixer"
url: "/wiki/APU_Mixer/"
description: "The NES APU mixer takes the channel outputs and converts them to an analog audio signal. Each channel has its own internal digital-to-analog convertor (DAC), implemented in a way that causes…"
wikiPage: "APU_Mixer"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

The [NES APU](../APU/) mixer takes the channel outputs and converts them to an analog audio signal. Each channel has its own internal digital-to-analog convertor (DAC), implemented in a way that causes non-linearity and interaction between channels, so calculation of the resulting amplitude is somewhat involved. In particular, games such as *Super Mario Bros.* and *StarTropics* use the DMC level (\$4011) as a crude volume control for the triangle and noise channels.

The following formula<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> calculates the approximate audio output level within the range of 0.0 to 1.0. It is the sum of two sub-groupings of the channels:

    output = pulse_out + tnd_out

                                95.88
    pulse_out = ------------------------------------
                 (8128 / (pulse1 + pulse2)) + 100

                                           159.79
    tnd_out = -------------------------------------------------------------
                                        1
               ----------------------------------------------------- + 100
                (triangle / 8227) + (noise / 12241) + (dmc / 22638)

The values for [pulse1](../APU_Pulse/), [pulse2](../APU_Pulse/), [triangle](../APU_Triangle/), [noise](../APU_Noise/), and [dmc](../APU_DMC/) are the output values for the corresponding channel. The dmc value ranges from 0 to 127 and the others range from 0 to 15. When the values for one of the groups are all zero, the result for that group should be treated as zero rather than undefined due to the division by 0 that otherwise results.

Faster but less accurate approximations are also possible: using an efficient [lookup table](#Lookup_Table), or even rougher with a [linear formula](#Linear_Approximation).

The NES hardware follows the DACs with a [surprisingly involved circuit](https://archive.nesdev.org/NESAudio.gif) that adds several low-pass and high-pass filters:

- A first-order high-pass filter at 90 Hz
- Another first-order high-pass filter at 440 Hz
- A first-order low-pass filter at 14 kHz

See also:

- [blargg's data and analysis, and lidnariq's matching analog components to filters](http://forums.nesdev.org/viewtopic.php?p=44255#p44255)

The Famicom hardware instead ONLY specifies a first-order high-pass filter at 37 Hz, followed by the unknown (and varying) properties of the RF modulator and demodulator.

## Emulation {#Emulation}

The NES APU Mixer can be efficiently emulated using a lookup table or a less-accurate linear approximation.

### Lookup Table {#Lookup_Table}

The APU mixer formulas can be efficiently implemented using two lookup tables: a 31-entry table for the two [pulse channels](../APU_Pulse/) and a 203-entry table for the remaining channels (due to the approximation of tnd_out, the numerators are adjusted slightly to preserve the normalized output range).

        output = pulse_out + tnd_out

        pulse_table [n] = 95.52 / (8128.0 / n + 100)

        pulse_out = pulse_table [pulse1 + pulse2]

The tnd_out table is approximated (within 4%) by using a base unit close to the [DMC's](../APU_DMC/) DAC.

        tnd_table [n] = 163.67 / (24329.0 / n + 100)

        tnd_out = tnd_table [3 * triangle + 2 * noise + dmc]

### Linear Approximation {#Linear_Approximation}

A linear approximation can also be used, which results in slightly louder DMC samples, but otherwise fairly accurate operation since the wave channels use a small portion of the transfer curve. The overall volume will be reduced due to the headroom required by the DMC approximation.

        output = pulse_out + tnd_out

        pulse_out = 0.00752 * (pulse1 + pulse2)

        tnd_out = 0.00851 * triangle + 0.00494 * noise + 0.00335 * dmc

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [apu_ref.txt](https://www.nesdev.org/apu_ref.txt) by blargg

<table class="collapsible collapsed" style="width: 100%; border: 1px solid #000; padding: 0">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2" style="background-color: #CCF; font-size: 120%">V | <del>T</del>
<a href="../APU/">APU</a></th>
</tr>
</thead>
<tbody>
<tr>
<td style="background-color: #DDF"><strong>Concepts</strong></td>
<td><a href="../APU/">Registers</a> | <a href="../APU_basics/">Basics</a> | <a href="../APU_period_table/">Period table</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Channels</strong></td>
<td><a href="../APU_Pulse/">Pulse 1 and Pulse 2</a> | <a href="../APU_Triangle/">Triangle</a> | <a href="../APU_Noise/">Noise</a> | <a href="../APU_DMC/">DMC</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Components</strong></td>
<td><a href="../APU_Envelope/">Envelope</a> | <a href="../APU_Pulse/">Sweep</a> | <a href="../APU_Frame_Counter/">Frame Counter</a> | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | Mixer</td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
