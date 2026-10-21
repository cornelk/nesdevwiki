---
title: "APU Frame Counter"
url: "/wiki/APU_Frame_Counter/"
description: "The NES APU frame counter (or frame sequencer) generates low-frequency clocks for the channels and an optional 60 Hz interrupt. Despite the name being \"frame counter\", the clocks have nothing to do…"
wikiPage: "APU_Frame_Counter"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

The **[NES APU](../APU/) frame counter** (or **frame sequencer**) generates low-frequency clocks for the channels and an optional 60 Hz interrupt. Despite the name being "frame counter", the clocks have nothing to do with the [Vsync signal from the PPU](../PPU_registers/#Vblank_NMI).

The sequencer is clocked on every other CPU cycle, so 2 CPU cycles = 1 APU cycle. The sequencer keeps track of how many APU cycles have elapsed in total, and each step of the sequence will occur once that total has reached the indicated amount (with an additional delay of one CPU cycle for the quarter and half frame signals). Once the last step has executed, the count resets to 0 on the next APU cycle.

### <span id="Control_register_.28.244017.2C_write.29"></span>Control register (\$4017, write) {#Control_register_.28.244017.2C_write.29}

    7  bit  0
    ---- ----
    MI.. ....
    ||
    |+-------- Interrupt inhibit flag
    +--------- External synchronization mode
                   0: Freerun
                   1: Expect external synchronization

Writing to this register **resets the counter** to 0. If the write occurs *during* an APU cycle, the effects occur 3 CPU cycles after the write cycle, and if the write occurs *between* APU cycles, the effects occurs 4 CPU cycles after the write cycle.

When the write occurs with the bit 7 (`M` bit) set, both Quarter frame and Half frame clocks are also generated.

The **external synchronization mode** (`M` bit) exists to provide a way for synchronization with the external timing source; in NES' case, it's the [Vblank NMI](../PPU_registers/#Vblank_NMI). When the mode is disabled, the frame counter counts it own ≈60Hz (or ≈50Hz in 2A07) cycle that is *not* synchronized with the [Vsync signal from PPU](../PPU_registers/#Vblank_NMI). When it's enabled, it expects the CPU to manually synchronize the internal counter with the external timing source that the CPU uses. (See [below](#.244017.7_.3D_1_.28External_synchronization_mode.29))

### <span id=".244017.7_.3D_0_.28Freerun_mode.29"></span>\$4017.7 = 0 (Freerun mode) {#.244017.7_.3D_0_.28Freerun_mode.29}

<table class="tabular" style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th>Step</th>
<th colspan="2">APU cycles</th>
<th><a href="../APU_Envelope/">Envelopes</a> &amp; <a href="../APU_Triangle/">triangle's linear counter</a><br />
(Quarter frame)</th>
<th><a href="../APU_Length_Counter/">Length counters</a> &amp; <a href="../APU_Pulse/">sweep units</a><br />
(Half frame)</th>
<th>Interrupt flag</th>
</tr>
</thead>
<tbody>
<tr>
<th>NTSC</th>
<th>PAL</th>
<th></th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>1</td>
<td>3728, PUT</td>
<td>4156, PUT</td>
<td>Clock</td>
<td> </td>
<td> </td>
</tr>
<tr>
<td>2</td>
<td>7456, PUT</td>
<td>8313, PUT</td>
<td>Clock</td>
<td>Clock</td>
<td> </td>
</tr>
<tr>
<td>3</td>
<td>11185, PUT</td>
<td>12469, PUT</td>
<td>Clock</td>
<td> </td>
<td> </td>
</tr>
<tr>
<td rowspan="3">4</td>
<td>14914, GET</td>
<td>16626, GET</td>
<td> </td>
<td> </td>
<td>Set if interrupt inhibit flag (<a href="#Control_register_.28.244017.2C_write.29"><code>I</code> bit</a>) is clear</td>
</tr>
<tr>
<td>14914, PUT</td>
<td>16626, PUT</td>
<td>Clock</td>
<td>Clock</td>
<td>Set if interrupt inhibit flag (<a href="#Control_register_.28.244017.2C_write.29"><code>I</code> bit</a>) is clear</td>
</tr>
<tr>
<td>0 (14915), GET</td>
<td>0 (16627), GET</td>
<td> </td>
<td> </td>
<td>Set if interrupt inhibit flag (<a href="#Control_register_.28.244017.2C_write.29"><code>I</code> bit</a>) is clear</td>
</tr>
<tr>
<td></td>
<td colspan="2"> </td>
<td>NTSC: <em>240 Hz (approx.)</em><br />
PAL: <em>200 Hz (approx)</em></td>
<td>NTSC: <em>120 Hz (approx.)</em><br />
PAL: <em>100 Hz (approx)</em></td>
<td>NTSC: <em>60 Hz (approx.)</em><br />
PAL: <em>50 Hz (approx)</em></td>
</tr>
</tbody>
</table>

- In this mode, the interrupt flag is set every 29830 CPU cycles, which is slightly (0.166%) slower than the [29780.5 CPU cycles](../Cycle_reference_chart/) per NTSC PPU frame.
  - On the PAL RP2A07, this is every 33254 CPU cycles, still slower than the [33247.5 CPU cycles](../Cycle_reference_chart/) per PAL PPU frame but much closer (only off by 0.0196%).
- The interrupt is intended to have the software sound player synchronized with those hardware envelopes.
- The interrupt flag is connected to the CPU's [IRQ](../IRQ/) line, and can be acknowledged (cleared) by writing to [\$4017](#Control_register_.28.244017.2C_write.29) with interrupt inhibit flag set (`I` bit), or reading [\$4015](../APU/#APU_master_control_and_status_.28.244015.2C_read.2Fwrite.29) also clears the flag.

### <span id=".244017.7_.3D_1_.28External_synchronization_mode.29"></span>\$4017.7 = 1 (External synchronization mode) {#.244017.7_.3D_1_.28External_synchronization_mode.29}

<table class="tabular" style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th>Step</th>
<th colspan="2">APU cycles</th>
<th><a href="../APU_Envelope/">Envelopes</a> &amp; <a href="../APU_Triangle/">triangle's linear counter</a><br />
(Quarter frame)</th>
<th><a href="../APU_Length_Counter/">Length counters</a> &amp; <a href="../APU_Pulse/">sweep units</a><br />
(Half frame)</th>
<th>Interrupt flag</th>
</tr>
</thead>
<tbody>
<tr>
<th>NTSC</th>
<th>PAL</th>
<th></th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>1</td>
<td>3728, PUT</td>
<td>4156, PUT</td>
<td>Clock</td>
<td> </td>
<td> </td>
</tr>
<tr>
<td>2</td>
<td>7456, PUT</td>
<td>8313, PUT</td>
<td>Clock</td>
<td>Clock</td>
<td> </td>
</tr>
<tr>
<td>3</td>
<td>11185, PUT</td>
<td>12469, PUT</td>
<td>Clock</td>
<td> </td>
<td> </td>
</tr>
<tr>
<td>4</td>
<td>14914, PUT</td>
<td>16626, PUT</td>
<td> </td>
<td> </td>
<td> </td>
</tr>
<tr>
<td rowspan="2">5</td>
<td>18640, PUT</td>
<td>20782, PUT</td>
<td>Clock</td>
<td>Clock</td>
<td> </td>
</tr>
<tr>
<td>0 (18641), GET</td>
<td>0 (20783), GET</td>
<td> </td>
<td> </td>
<td> </td>
</tr>
<tr>
<td></td>
<td colspan="2"> </td>
<td>NTSC: <em>192 Hz (approx.), uneven timing</em><br />
PAL: <em>160 Hz (approx.), uneven timing</em></td>
<td>NTSC: <em>96 Hz (approx.), uneven timing</em><br />
PAL: <em>80 Hz (approx.), uneven timing</em></td>
<td>NTSC: <em>none</em><br />
PAL: <em>none</em></td>
</tr>
</tbody>
</table>

- In this mode, the interrupt flag is **never** set.
- This mode essentially gives the software about 7,454 CPU cycles (8,314 in 2A07) of gap to **reset the sequencer manually** before the counter wraps around to 0.
  - The synchronization is done by writing to [\$4017](#Control_register_.28.244017.2C_write.29) with the bit 7 (`M` bit) set at the (beginning of the) external timing source, which will reset the counter and also generate the both Quarter frame and Half frame clock.

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
<td><a href="../APU_Envelope/">Envelope</a> | <a href="../APU_Pulse/">Sweep</a> | Frame Counter | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | <a href="../APU_Mixer/">Mixer</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
