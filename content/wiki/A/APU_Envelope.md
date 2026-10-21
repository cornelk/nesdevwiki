---
title: "APU Envelope"
url: "/wiki/APU_Envelope/"
description: "In a synthesizer, an envelope is the way a sound's parameter changes over time. The NES APU has an envelope generator that controls the volume in one of two ways: it can generate a decreasing saw…"
wikiPage: "APU_Envelope"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

In a synthesizer, an [envelope](https://en.wikipedia.org/wiki/ADSR_envelope) is the way a sound's parameter changes over time. The NES [APU](../APU/) has an envelope generator that controls the volume in one of two ways: it can generate a decreasing saw envelope (like a decay phase of an [ADSR](http://en.wikipedia.org/wiki/Synthesizer#ADSR_envelope)) with optional looping, or it can generate a constant volume that a more sophisticated software envelope generator can manipulate. Volume values are practically linear (see: [APU Mixer](../APU_Mixer/)).

Each volume envelope unit contains the following: start flag, [divider](../APU/), and decay level counter.

                                       Loop flag
                                            |
                   Start flag  +--------.   |   Constant volume
                               |        |   |        flag
                               v        v   v          |
    Quarter frame clock --> Divider --> Decay --> |    | 
                               ^        level     |    v
                               |                  | Select --> Envelope output
                               |                  |
            Envelope parameter +----------------> |                   

<table class="wikitable">
<thead>
<tr>
<th>Address</th>
<th>Bitfield</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>$4000</strong></td>
<td><code>ddLC.VVVV</code></td>
<td><strong><a href="../APU_Pulse/">Pulse channel 1</a></strong> duty and volume/envelope (write)</td>
</tr>
<tr>
<td><strong>$4004</strong></td>
<td><code>ddLC.VVVV</code></td>
<td><strong><a href="../APU_Pulse/">Pulse channel 2</a></strong> duty and volume/envelope (write)</td>
</tr>
<tr>
<td><strong>$400C</strong></td>
<td><code>--LC.VVVV</code></td>
<td><strong><a href="../APU_Noise/">Noise channel</a></strong> volume/envelope (write)</td>
</tr>
<tr>
<td>bit 5</td>
<td><code>--L- ----</code></td>
<td><a href="../APU_Length_Counter/">APU Length Counter</a> halt flag/envelope loop flag</td>
</tr>
<tr>
<td>bit 4</td>
<td><code>---C ----</code></td>
<td>Constant volume flag (0: use volume from envelope; 1: use constant volume)</td>
</tr>
<tr>
<td>bits 3-0</td>
<td><code>---- VVVV</code></td>
<td>Used as the volume in constant volume (C set) mode. Also used as the reload value for the envelope's divider (the period becomes V + 1 quarter frames).</td>
</tr>
<tr>
<td colspan="3"> </td>
</tr>
<tr>
<td><strong>$4003</strong></td>
<td><code>LLLL.Lttt</code></td>
<td><strong><a href="../APU_Pulse/">Pulse channel 1</a></strong> <a href="../APU_Length_Counter/">length counter load</a> and <a href="../APU/">timer</a> (write)</td>
</tr>
<tr>
<td><strong>$4007</strong></td>
<td><code>LLLL.Lttt</code></td>
<td><strong><a href="../APU_Pulse/">Pulse channel 2</a></strong> <a href="../APU_Length_Counter/">length counter load</a> and <a href="../APU/">timer</a> (write)</td>
</tr>
<tr>
<td><strong>$400F</strong></td>
<td><code>LLLL.L---</code></td>
<td><strong><a href="../APU_Noise/">Noise channel</a></strong> <a href="../APU_Length_Counter/">length counter load</a> (write)</td>
</tr>
<tr>
<td colspan="2">Side effects</td>
<td>Sets start flag</td>
</tr>
</tbody>
</table>

When clocked by the [frame counter](../APU_Frame_Counter/), one of two actions occurs: if the start flag is clear, the divider is clocked, otherwise the start flag is cleared, the decay level counter is loaded with 15, and the divider's period is immediately reloaded.

When the divider is clocked while at 0, it is loaded with V and clocks the decay level counter. Then one of two actions occurs: If the counter is non-zero, it is decremented, otherwise if the loop flag is set, the decay level counter is loaded with 15.

The envelope unit's volume output depends on the constant volume flag: if set, the envelope parameter directly sets the volume, otherwise the decay level is the current volume. The constant volume flag has no effect besides selecting the volume source; the decay level will still be updated when constant volume is selected.

Each of the three envelope units' output is fed through additional gates at the [sweep unit](../APU_Pulse/) (pulse only), waveform generator ([sequencer](../APU/#Glossary) or LFSR), and [length counter](../APU_Length_Counter/).

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
<td>Envelope | <a href="../APU_Pulse/">Sweep</a> | <a href="../APU_Frame_Counter/">Frame Counter</a> | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | <a href="../APU_Mixer/">Mixer</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
