---
title: "PPU signals"
url: "/wiki/PPU_signals/"
description: "The PPU has many internal signals for controlling its various actions during the frame. Each signal can control many things in the chip, and each of these can have a different amount of delay between…"
wikiPage: "PPU_signals"
---

The PPU has many internal signals for controlling its various actions during the frame. Each signal can control many things in the chip, and each of these can have a different amount of delay between the signal being generated and being used. These delays can be critical for correct operation.

## H decoder {#H_decoder}

These signals are all generated from the current dot number. Some of these signals are combinations of multiple H decoder signals, and most of them are disabled during vblank (scanlines 240-260) and fblank. For two of them, they are disabled specifically during scanlines 240-261, but since this is so uncommon, it is called out as needed rather than in a dedicated column.

<table class="wikitable sortable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Dots</th>
<th>Disabled during vblank/fblank</th>
<th>Names</th>
<th class="unsortable">Purpose</th>
</tr>
</thead>
<tbody>
<tr>
<td>dot % 8 == 0..1</td>
<td>Yes</td>
<td>Breaks: /F_NT</td>
<td>Pattern address register control
<ul>
<li><strong>(Delay: 1 dot) Pattern address register</strong>: During 2nd half even dots, this produces a signal that is used in the next half dot to latch data into the register used to generate pattern addresses. This signal latches OAM buffer bit 0 into the bit 12 register, bits 7-0 of the PPU data bus (delayed 0.5 dots) and OAM buffer bytes into the bits 11-4 registers, and bits 3-0 of the OAM in-range result (delayed 0.5 dots) into bits 3-0. The address register selects which data to use for output based on the fetch sprites (256-319) signal.</li>
<li><strong>(Delay: 1 dot) Sprite fetch</strong>: During 2nd half even dots, this latches the OAM buffer value's in-range result to be used while loading sprite shifters.</li>
</ul>
<p>Commentary: This signal is used for creating pattern addresses for both background and sprites. The latching is delayed half a dot to allow the OAM tile byte to be loaded into the OAM buffer. Much of the data used for this register is latched, but the live inputs are used for the pattern table selection, sprite size, and background fine y.</p></td>
</tr>
<tr>
<td>dot % 8 == 2..3<br />
(0..255, 320..335)</td>
<td>Yes</td>
<td>Breaks: F_AT</td>
<td>Background attributes control
<ul>
<li><strong>(Delay: 0.5 dots) VRAM address mux</strong>: When a pattern address is not selected, this causes an attributes address to be used instead of a nametable address.</li>
<li><strong>(Delay: 1 dot) Background rendering</strong>: During 2nd half even dots, this transfers the PPU data bus value into the attributes register.</li>
</ul>
<p>Commentary: This signal is relevant for palette corruption. Palette corruption happens when the palette address changes during 2nd half dots, which normally happens from rendering toggles. The output from the VRAM address mux is used to select between the rendering pipeline output and the low 5 bits of v for the palette RAM address. The low 4 bits are not synchronized and can change any time. Rendering state changes the addresses sent to the VRAM address mux, allowing the mux's output to immediately change and switch the low 4 bits of the palette RAM address. Rendering state controls bits 13-12 of the attributes address sent to the mux, making them either a fixed %10 or bits 13-12 of v. If rendering turns off during an attributes fetch, this mutated address can cause corruption by changing the selected palette RAM address even if v itself wasn't pointing into palette RAM; this happens any time v &amp; $3C00 == $3C00. At other times, the nametable address (which becomes v) becomes selected when rendering disables, so v has to point into palette RAM to trigger corruption (v &amp; $3F00 == $3F00).</p></td>
</tr>
<tr>
<td>dot % 8 == 4..5<br />
(0..255, 320..335)</td>
<td>Yes</td>
<td>Breaks: F_TA</td>
<td>Background low pattern register control
<ul>
<li><strong>(Delay: 1 dot) Background rendering</strong>: During 2nd half even dots, this transfers the PPU data bus value into the low pattern bits register.</li>
</ul></td>
</tr>
<tr>
<td>dot % 8 == 6..7<br />
(0..255, 320..335)</td>
<td>Yes</td>
<td>Breaks: F_TB</td>
<td>Background sliver commit
<ul>
<li><strong>(Delay: 1 dot) Background rendering</strong>: During 2nd half even dots, this increments coarse x and transfers data from the PPU data bus and pattern and attributes registers into the background shift registers.</li>
</ul></td>
</tr>
<tr>
<td>0..7</td>
<td>No</td>
<td>Breaks: CLIP_O / CLIP_B</td>
<td>Clipping enable
<ul>
<li><strong>(Delay: 2 dots) Background rendering</strong>: This signal permits the left column background masking bit in <a href="../PPU_registers/#PPUCTRL">PPUCTRL</a> to clear the pixels coming out of the background shift registers.</li>
<li><strong>(Delay: 1.5 dots) Sprite rendering</strong>: This signal permits the left column sprite masking bit in <a href="../PPU_registers/#PPUCTRL">PPUCTRL</a> to clear the pixels coming out of the sprite shifters.</li>
</ul>
<p>Commentary: This same clipping mechanism is used to block pixel output during not-visible pixels and when rendering components are disabled. This signal is actually a combination of a 0-7 or 256-263 signal and a !(0-255 on scanlines 0-239) signal, but only the 0-7 part of it can actually have any impact on the screen.</p></td>
</tr>
<tr>
<td>0..63</td>
<td>Yes</td>
<td>Breaks: I_OAM2</td>
<td>OAM2 init
<ul>
<li><strong>(Delay: 1 dot) OAM2 init</strong>: This overrides the input to the OAM buffer with $FF instead of the value read from OAM (so that $FF can be written to OAM2). Note that there's an additional 1 dot of delay before the OAM buffer latches its input (it latches during 2nd half dots and is normally read during 1st half dots), so the $FF exists in the buffer on dots 2-65.</li>
<li><strong>(Delay: 1.5 dots) OAM2 init</strong>: This suppresses OAM1ADDR increments that normally happen automatically during the 1st half visible odd dots.</li>
<li><strong>(Delay: 1.5 dots) OAM2 init</strong>: During 1st half odd dots, this increments OAM2ADDR.</li>
<li><strong>(Delay: 1 dot) Sprite evaluation</strong>: This forces the OAM in-range input to the sprite evaluation in-range circuit to 0 (out of range).</li>
<li><strong>(Delay: 1 dot) Sprite evaluation</strong>: This continuously clears the eval_done signal (which, when set, forces the sprite evaluation version of the OAM in-range result to 0 and turns OAM2 writes into reads).</li>
</ul></td>
</tr>
<tr>
<td>0..255<br />
(scanlines 0-239)</td>
<td>Yes</td>
<td>Breaks: /VIS</td>
<td>Handle visible dots
<ul>
<li><strong>(Delay: 1.5 dots) Background rendering</strong>: Outside this range, background pixels are clipped (forced to transparent).</li>
<li><strong>(Delay: 2.5 dots) Sprite rendering</strong>: Outside this range, sprite pixels are clipped (forced to transparent).</li>
<li><strong>(Delay: 1.5 dots) Sprite evaluation</strong>: During 1st half odd dots, this causes OAM1ADDR to increment unless suppressed by OAM2 init.</li>
<li><strong>(Delay: 1 dot) Sprite evaluation</strong>: This signal causes odd dots to use OAM1ADDR to address OAM. Otherwise, all dots during rendering use OAM2ADDR.</li>
<li><strong>(Delay: 1 dot) Sprite evaluation</strong>: When this signal is false, it forces the OAM in-range input to the sprite evaluation in-range circuit to 0 (out of range).</li>
<li><strong>(Delay: 1 dot) Sprite evaluation / Sprite fetch</strong>: When this signal is false, the automatic writes to OAM on even dots (for writing to OAM2 during OAM2 init and sprite evaluation) are instead reads.</li>
<li><strong>(Delay: 1 dot) Sprite rendering</strong>: During 2nd half odd dots, this allows sprite shifters to output-and-shift pixels for a dot if they are not counting.</li>
<li><strong>(Delay: 1 dot) Sprite 0 hit</strong>: When this signal is false, the sprite 0 hit flag cannot be set.</li>
</ul>
<p>Commentary: Most H decoder signals are enabled during the pre-render scanline, but this one is not. This produces unusual behavior for sprite rendering during the pre-render scanline in that it disables OAM2 init and sprite evaluation, but not sprite fetch. This usually (but not always) causes scanline 0 to be sprite-free.</p></td>
</tr>
<tr>
<td>0..255, 320..335</td>
<td>Yes</td>
<td>Breaks: /FO</td>
<td>Handle background
<ul>
<li><strong>(Delay: 1 dot) Background rendering</strong>: During 2nd half dots, this clocks the background shift registers.</li>
</ul>
<p>Commentary: This is the signal that limits the 'Every n..n+1' signals for background handling.</p></td>
</tr>
<tr>
<td>256..319</td>
<td>Yes</td>
<td>Breaks: OBJ_READ</td>
<td>Sprite fetch
<ul>
<li><strong>(Delay: 1 dot) Sprite evaluation</strong>: This continuously resets OAM1ADDR ($2003).</li>
<li><strong>(Delay: 1 dot) Sprite 0 hit</strong>: During 1st half dots, this enables the sprite_0_on_this_scanline flip-flop so it can store the current output of the sprite_0_on_next_scanline flip-flop.</li>
<li><strong>(Delay: 1 dot) Sprite fetch</strong>: During 1st half dots, this increments OAM2ADDR (if it's not being otherwise suppressed) if bit 2 of the dot (delayed 0.5 dots) is 0. Because this sprite fetch signal also changes during 1st half dots, this creates asynchronous timing at the start of dot 321, where the signal goes from 1 to 0 while OAM2ADDR is using it to increment. In practice, this reliably causes an extra increment.</li>
<li><strong>(Delay: 1 dot) Sprite fetch</strong>: This enables a demux that takes the low 3 bits of the dot number (delayed 1 dot) and produces 4 signals that tell the sprite shifters when to load each piece of data. The signals from the demux are latched during 2nd half dots.</li>
<li><strong>(Delay: 1 dot) Pattern address register</strong>: During 2nd half dots, this selects between background and sprite data for the bits of the pattern address register.</li>
</ul>
<p>Commentary: The extra OAM2ADDR increment that occurs on dot 321 is critical for software because it provides a wide window during sprite idle where rendering can be safely turned off without setting up the conditions for OAM corruption later on. Games rely on this behavior.</p></td>
</tr>
<tr>
<td>63, 255, 339</td>
<td>Yes</td>
<td>Breaks: /EVAL</td>
<td>OAM2ADDR reset
<ul>
<li><strong>(Delay: 1.5 dots) OAM2 init / Sprite evaluation / Sprite fetch</strong>: During 1st half dots, this clears OAM2ADDR and oam2_overflow.</li>
<li><strong>(Delay: 1.5 dots) Sprite evaluation / Sprite fetch</strong>: This prevents OAM2ADDR from being able to increment.</li>
</ul></td>
</tr>
<tr>
<td>65</td>
<td>Yes</td>
<td>Breaks: S_EV</td>
<td>Sprite 0 detection
<ul>
<li><strong>(Delay: 1 dot) Sprite 0 hit</strong>: During 1st half dots, this enables the sprite_0_on_next_scanline flip-flop so it can store the sprite evaluation version of the OAM in-range result.</li>
</ul></td>
</tr>
<tr>
<td>255</td>
<td>Yes</td>
<td>Breaks: E_EV</td>
<td>Increment v
<ul>
<li><strong>(Delay: 1 dot) Background rendering</strong>: During 2nd half dots, this increments the y component of v.</li>
<li><strong>(Delay: 2 dots) Background rendering</strong>: During 2nd half dots, this copies the coarse x component of t to v.</li>
</ul></td>
</tr>
<tr>
<td>339</td>
<td>Yes</td>
<td>Breaks: 0_HPOS</td>
<td>Start sprite shifter counters
<ul>
<li><strong>(Delay: 2.5 dots) Sprite rendering</strong>: This sets an SR latch in each sprite shifter that makes the shifter count instead of output-and-shift pixels. When the counter expires, the latch is cleared.</li>
</ul>
<p>Commentary: When the last dot of pre-render is skipped, this signal arrives at the sprite shifters one dot too late, so every shifter outputs its first pixel at X=0 and then they start counting on the next dot, outputting the remaining 7 pixels at their usual time.</p></td>
</tr>
<tr>
<td>339</td>
<td>Yes</td>
<td>Breaks: /HPLA_5</td>
<td>End of scanline (odd frames)
<ul>
<li><strong>(Delay: 0 dots) Frame end</strong>: This signal, along with a scanline 261 (delayed 1 dot) signal, allows an even/odd frame toggle to act as another input to the same end-of-scanline circuit as dot 340.</li>
</ul></td>
</tr>
<tr>
<td>340</td>
<td>No</td>
<td>Breaks: HPLA_23</td>
<td>End of scanline
<ul>
<li><strong>(Delay: 0 dots) Scanline end</strong>: This signal increments the scanline counter.</li>
<li><strong>(Delay: 0.5 dots) Scanline end / Frame end</strong>: This signal clears the dot counter. If the scanline is 261 (delayed 0.5 dots), it also clears the scanline counter.</li>
</ul></td>
</tr>
</tbody>
</table>
