---
title: "APU DMC"
url: "/wiki/APU_DMC/"
description: "The NES APU's delta modulation channel (DMC) provides automatic 1-bit delta-encoded sample playback unit that can fetch a sample byte from the same CPU memory area by its own, or the 7-bit counter…"
wikiPage: "APU_DMC"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

The **[NES APU's](../APU/) delta modulation channel (DMC)** provides [automatic](#Sample_playback_unit_operation) 1-bit [delta-encoded sample](https://en.wikipedia.org/wiki/Delta_modulation) playback unit that can fetch a sample byte from the same [CPU memory area](../CPU_memory_map/) by its own, or the 7-bit counter (final output level) can be [directly loaded](#Direct_counter_load_.28.244011.2C_write.29), allowing flexible manual sample playback via CPU.

                             Timer        Direct Load
                               |               |
                               v               v
    Reader ---> Buffer ---> Shifter ---> 7-bit Counter ---> (to mixer)

## Registers {#Registers}

### <span id="Flags_and_sample_playback_rate_.28.244010.2C_write.29"></span>Flags and sample playback rate (\$4010, write) {#Flags_and_sample_playback_rate_.28.244010.2C_write.29}

    7  bit  0
    ---- ----
    IL.. RRRR
    ||   ||||
    ||   ++++- Playback rate index
    |+-------- Sample loop flag
    +--------- Enable DMC sample-end interrupt (hereinafter "DMC IRQ")

The rate (`R` bits) determines for how many CPU cycles happen between changes in the output level during [automatic sample playback](#Sample_playback_unit_operation).

Setting the loop flag (`L` bit) inhibits DMC IRQ pending flag from being set.

#### Pitch table {#Pitch_table}

Periods in this table are in CPU cycles.

<table class="wikitable">
<tbody>
<tr>
<td></td>
<td colspan="6" data-bgcolor="FFDDDD">NTSC</td>
<td colspan="6" data-bgcolor="DDDDFF">PAL</td>
</tr>
<tr>
<td>$4010.0-3</td>
<td>Period</td>
<td>Frequency</td>
<td>Note (raw)</td>
<td>Note (1-byte loop)</td>
<td>Note (17-byte loop)</td>
<td>Note (33-byte loop)</td>
<td>Period</td>
<td>Frequency</td>
<td>Note (raw)</td>
<td>Note (1-byte loop)</td>
<td>Note (17-byte loop)</td>
<td>Note (33-byte loop)</td>
</tr>
<tr>
<td>$0</td>
<td data-bgcolor="FFDDDD">428</td>
<td>4181.71 Hz</td>
<td>C-8 -2c</td>
<td>C-5 -2c</td>
<td>B-0 -7c</td>
<td>infrasound</td>
<td data-bgcolor="DDDDFF">398</td>
<td>4177.40 Hz</td>
<td>C-8 -4c</td>
<td>C-5 -4c</td>
<td>B-0 -9c</td>
<td>infrasound</td>
</tr>
<tr>
<td>$1</td>
<td data-bgcolor="FFDDDD">380</td>
<td>4709.93 Hz</td>
<td>D-8 +4c</td>
<td>D-5 +4c</td>
<td>C#1 -1c</td>
<td>infrasound</td>
<td data-bgcolor="DDDDFF">354</td>
<td>4696.63 Hz</td>
<td>D-8 -1c</td>
<td>D-5 -1c</td>
<td>C#1 -6c</td>
<td>infrasound</td>
</tr>
<tr>
<td>$2</td>
<td data-bgcolor="FFDDDD">340</td>
<td>5264.04 Hz</td>
<td>E-8 -3c</td>
<td>E-5 -3c</td>
<td>Eb1 -8c</td>
<td>infrasound</td>
<td data-bgcolor="DDDDFF">316</td>
<td>5261.41 Hz</td>
<td>E-8 -4c</td>
<td>E-5 -4c</td>
<td>Eb1 -9c</td>
<td>infrasound</td>
</tr>
<tr>
<td>$3</td>
<td data-bgcolor="FFDDDD">320</td>
<td>5593.04 Hz</td>
<td>F-8 +2c</td>
<td>F-5 +2c</td>
<td>E-1 -3c</td>
<td>E-0 +48c</td>
<td data-bgcolor="DDDDFF">298</td>
<td>5579.22 Hz</td>
<td>F-8 -3c</td>
<td>F-5 -3c</td>
<td>E-1 -8c</td>
<td>E-0 +44c</td>
</tr>
<tr>
<td>$4</td>
<td data-bgcolor="FFDDDD">286</td>
<td>6257.95 Hz</td>
<td>G-8 -4c</td>
<td>G-5 -4c</td>
<td>F#1 -9c</td>
<td>F#0 +43c</td>
<td data-bgcolor="DDDDFF">276</td>
<td>6023.94 Hz</td>
<td>G-8 -70c</td>
<td>G-5 -70c</td>
<td>F#1 -75c</td>
<td>F#0 -23c</td>
</tr>
<tr>
<td>$5</td>
<td data-bgcolor="FFDDDD">254</td>
<td>7046.35 Hz</td>
<td>A-8 +2c</td>
<td>A-5 +2c</td>
<td>Ab1 -3c</td>
<td>Ab0 +48c</td>
<td data-bgcolor="DDDDFF">236</td>
<td>7044.94 Hz</td>
<td>A-8 +1c</td>
<td>A-5 +1c</td>
<td>Ab1 -4c</td>
<td>Ab0 +48c</td>
</tr>
<tr>
<td>$6</td>
<td data-bgcolor="FFDDDD">226</td>
<td>7919.35 Hz</td>
<td>B-8 +4c</td>
<td>B-5 +4c</td>
<td>Bb1 -1c</td>
<td>Bb0 +50c</td>
<td data-bgcolor="DDDDFF">210</td>
<td>7917.18 Hz</td>
<td>B-8 +3c</td>
<td>B-5 +3c</td>
<td>Bb1 -2c</td>
<td>Bb0 +50c</td>
</tr>
<tr>
<td>$7</td>
<td data-bgcolor="FFDDDD">214</td>
<td>8363.42 Hz</td>
<td>C-9 -2c</td>
<td>C-6 -2c</td>
<td>B-1 -7c</td>
<td>B-0 +45c</td>
<td data-bgcolor="DDDDFF">198</td>
<td>8397.01 Hz</td>
<td>C-9 +5c</td>
<td>C-6 +5c</td>
<td>B-1 +0c</td>
<td>B-0 +52c</td>
</tr>
<tr>
<td>$8</td>
<td data-bgcolor="FFDDDD">190</td>
<td>9419.86 Hz</td>
<td>D-9 +4c</td>
<td>D-6 +4c</td>
<td>C#2 -1c</td>
<td>C#1 +51c</td>
<td data-bgcolor="DDDDFF">176</td>
<td>9446.63 Hz</td>
<td>D-9 +9c</td>
<td>D-9 +9c</td>
<td>C#2 +4c</td>
<td>C#1 +56c</td>
</tr>
<tr>
<td>$9</td>
<td data-bgcolor="FFDDDD">160</td>
<td>11186.1 Hz</td>
<td>F-9 +2c</td>
<td>F-6 +2c</td>
<td>E-2 -3c</td>
<td>E-1 +48c</td>
<td data-bgcolor="DDDDFF">148</td>
<td>11233.8 Hz</td>
<td>F-9 +9c</td>
<td>F-6 +9c</td>
<td>E-2 +4c</td>
<td>E-1 +56c</td>
</tr>
<tr>
<td>$A</td>
<td data-bgcolor="FFDDDD">142</td>
<td>12604.0 Hz</td>
<td>G-9 +8c</td>
<td>G-6 +8c</td>
<td>F#2 +3c</td>
<td>F#1 +55c</td>
<td data-bgcolor="DDDDFF">132</td>
<td>12595.5 Hz</td>
<td>G-9 +7c</td>
<td>G-6 +7c</td>
<td>F#2 -2c</td>
<td>G-1 +54c</td>
</tr>
<tr>
<td>$B</td>
<td data-bgcolor="FFDDDD">128</td>
<td>13982.6 Hz</td>
<td>A-9 -12c</td>
<td>A-6 -12c</td>
<td>Ab2 -17c</td>
<td>Ab1 +35c</td>
<td data-bgcolor="DDDDFF">118</td>
<td>14089.9 Hz</td>
<td>A-9 +1c</td>
<td>A-6 +1c</td>
<td>Ab2 -4c</td>
<td>Ab1 +48c</td>
</tr>
<tr>
<td>$C</td>
<td data-bgcolor="FFDDDD">106</td>
<td>16884.6 Hz</td>
<td>C-10 +14c</td>
<td>C-7 +14c</td>
<td>B-2 +10c</td>
<td>B-1 +61c</td>
<td data-bgcolor="DDDDFF">98</td>
<td>16965.4 Hz</td>
<td>C-10 +23c</td>
<td>C-7 +23c</td>
<td>B-2 +18c</td>
<td>B-1 +69c</td>
</tr>
<tr>
<td>$D</td>
<td data-bgcolor="FFDDDD">84</td>
<td>21306.8 Hz</td>
<td>E-10 +17c</td>
<td>E-7 +17c</td>
<td>Eb3 +12c</td>
<td>Eb2 +64c</td>
<td data-bgcolor="DDDDFF">78</td>
<td>21315.5 Hz</td>
<td>E-10 +18c</td>
<td>E-7 +18c</td>
<td>Eb3 +13c</td>
<td>Eb2 +65c</td>
</tr>
<tr>
<td>$E</td>
<td data-bgcolor="FFDDDD">72</td>
<td>24858.0 Hz</td>
<td>G-10 -16c</td>
<td>G-7 -16c</td>
<td>F#3 -21c</td>
<td>F#2 +31c</td>
<td data-bgcolor="DDDDFF">66</td>
<td>25191.0 Hz</td>
<td>G-10 +7c</td>
<td>G-7 +7c</td>
<td>F#3 +2c</td>
<td>F#2 +54c</td>
</tr>
<tr>
<td>$F</td>
<td data-bgcolor="FFDDDD">54</td>
<td>33143.9 Hz</td>
<td>C-11 -18c</td>
<td>C-8 -18c</td>
<td>B-3 -23c</td>
<td>B-2 +29c</td>
<td data-bgcolor="DDDDFF">50</td>
<td>33252.1 Hz</td>
<td>C-11 -12c</td>
<td>C-8 -12c</td>
<td>B-3 -17c</td>
<td>B-2 +34c</td>
</tr>
</tbody>
</table>

(Deviation from note is given in cents, which are defined as 1/100th of a semitone.)

On RP2A07, the pitches at \$4 and \$C appear to be incorrect with respect to their intended A-440 tuning scheme. [\[1\]](https://forums.nesdev.org/viewtopic.php?p=94079#p94079)

All "Dendy timing" clone CPUs are known to use NTSC period table.

### <span id="Direct_counter_load_.28.244011.2C_write.29"></span>Direct counter load (\$4011, write) {#Direct_counter_load_.28.244011.2C_write.29}

    7  bit  0
    ---- ----
    .DDD DDDD
     ||| ||||
     +++-++++- Unsigned 7-bit counter value

Writing to this register directly sets the counter value (final output level). If the [automatic DMC playback unit](#Sample_playback_unit_operation) clocks the counter at the same time, the output level is occasionally not changed properly. [\[2\]](https://forums.nesdev.org/viewtopic.php?p=104491#p104491)

The change still occurs even when the [automatic DMC playback unit](#Sample_playback_unit_operation) is deactivated.

### <span id="Sample_address_.28.244012.2C_write.29"></span>Sample address (\$4012, write) {#Sample_address_.28.244012.2C_write.29}

    7  bit  0
    ---- ----
    AAAA AAAA
    |||| ||||
    ++++-++++- Sample address = $C000 + (A * 64)

This register sets reload value. Retrigerring the sample playback reloads the value back from this register.

### <span id="Sample_length_.28.244013.2C_write.29"></span>Sample length (\$4013, write) {#Sample_length_.28.244013.2C_write.29}

    7  bit  0
    ---- ----
    LLLL LLLL
    |||| ||||
    ++++-++++- Sample length = (L * 16) + 1 bytes

This register sets reload value. Retrigerring the sample playback reloads the value back from this register.

### <span id="Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29"></span>Sample playback control, status (\$4015, read/write) {#Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29}

    7  bit  0 (read)
    ---- ----
    IF.D NT21
    |  |
    |  +------ Sample is on playback flag
    +--------- DMC IRQ asserted

The `D` bit returns 1 when the [bytes-remaining counter](#Memory_reader) is not zero.

    7  bit  0 (write)
    ---- ----
    ...D NT21
       |
       +------ Start / Stop DMC sample playback

When writing to this register with the `D` bit *clear*, the [bytes-remaining counter](#Memory_reader) is set to 0, consequently finishes the playback before the next sample byte is fetched. When the write is done with the bit *set* and if the [bytes-remaining counter](#Memory_reader) is 0, the sample playback is restarted, by reloading the [address counter](#Memory_reader) and [bytes-remaining counter](#Memory_reader) based on [\$4012](#Sample_address_.28.244012.2C_write.29) and [\$4013](#Sample_length_.28.244013.2C_write.29).

Writing to this register also clears DMC IRQ pending flag.

For descriptions for other bits, see [APU#Status](../APU/#Status).

## Sample playback unit operation {#Sample_playback_unit_operation}

The output level, or the current output of the 7-bit counter is sent to the [mixer](../APU_Mixer/) regardless the channel is [enabled or not](#Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29). It can be updated via [\$4011 writes](#Direct_counter_load_.28.244011.2C_write.29) *and* the automatic [delta-encoded sample](https://en.wikipedia.org/wiki/Delta_modulation) playback unit.

The automatic sample playback unit consists of two *seperate* units that operate in parallel, a [memory reader](#Memory_reader) and an [output unit](#Output_unit).

### Output unit {#Output_unit}

The **output unit** take the sample byte from the sample buffer (filled by [memory reader](#Memory_reader)) and update the output level accordingly. It includes an 8-bit **shift register** that is loaded with the content from sample buffer, shifted as the clocking happens.

The output unit runs *regardless* the sample playback is happening or not. When [timer](../APU/#Glossary) generates clock, following occur:

1.  The least significant bit of the shift register detemines whether to up or down count the output level:
    - If it is 0 and the current output level is higher than 1, decrement the output level by 2.
    - If it is 1 and the current output level is below 126, increment the output level by 2.
    - Notice that the difference per a sample bit is **2** and not **1**.
2.  The shift register shifts the bits to *right*.
3.  If all eight bits in the shift register is spent, whether the sample buffer is empty or not determines if the new byte cycle starts or the output unit is inactive for the next eight [timer](../APU/#Glossary) clocks.
    - If the sample buffer is empty: *do not update the output level* for the next eight [timer](../APU/#Glossary) clocks.
    - Otherwise if the sample buffer is filled: The shift register is filled with the content of the sample buffer, sample buffer becomes empty state and the byte cycle repeats.
      - Emptying the sample buffer essentially means the [memory reader](#Memory_reader) will immediately fill the sample buffer back if the current sample isn't finished.

### Memory reader {#Memory_reader}

The **memory reader** includes a [DMA device](../DMA/#DMC_DMA) that reads the target sample byte from the [CPU memory area](../CPU_memory_map/) by its own.

With these internal components:

- An **address counter** that points at the CPU address of the sample byte to read.
- A **bytes-remaining counter** that counts the number of bytes that are left to read.
- An 8-bit **sample buffer** that either temporary holds the current sample byte read with the DMA or is in empty *state*.

**(Re)starting a sample** reloads the address counter and the bytes-remaining counter based on [\$4012](#Sample_address_.28.244012.2C_write.29) and [\$4013](#Sample_length_.28.244013.2C_write.29).

The sample byte fetching process occurs whenever the sample buffer is empty *and* the bytes-remaining counter isn't 0. Which only means it occurs when one of following happens:

- The playback has started with [\$4015 write](#Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29) when sample buffer was empty.
- The playback has restarted because of the [loop flag (`L` bit)](#Flags_and_sample_playback_rate_.28.244010.2C_write.29).
- The [output unit](#Output_unit) emptied the sample buffer to fill the shift register and the sample hasn't finished yet.

When the conditions are met, following happen:

1.  [DMC DMA](../DMA/#DMC_DMA) halts the CPU for 1-4 cycles to read the target byte, and fill the sample buffer with the byte read.
2.  The address counter is incremented. If it exceeds \$FFFF, it is wrapped around \$8000.
3.  The bytes-remaining counter is decremented. If it reaches 0-
    - and if the [loop flag (`L` bit)](#Flags_and_sample_playback_rate_.28.244010.2C_write.29) is set, the sample is restarted. (reloads the address counter and the bytes-remaining counter based on [\$4012](#Sample_address_.28.244012.2C_write.29) and [\$4013](#Sample_length_.28.244013.2C_write.29))
    - Otherwise if it reaches 0 and if the [sample-end interrupt (`I` bit)](#Flags_and_sample_playback_rate_.28.244010.2C_write.29) is enabled, DMC IRQ pending flag is set.
      - At any time, if the interrupt flag is set, the [CPU's IRQ line](../IRQ/) is continuously asserted until the interrupt flag is cleared. It can be clared by writing any value to [\$4015](#Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29) or by clearing the [sample-end interrupt flag (`I` bit) in \$4010](#Flags_and_sample_playback_rate_.28.244010.2C_write.29).

## Conflict with controller and PPU read {#Conflict_with_controller_and_PPU_read}

On the 2A03 CPU used in the NTSC NES and Famicom, if a new sample byte is fetched from memory at the same time the program is reading a register that has read side effects (such as [controllers via \$4016 or \$4017](../Controller_reading/#.244016_.2F_.244017_Read), [PPUDATA (\$2007)](../PPU_registers/#PPUDATA), or the [Vblank flag](../PPU_registers/#Vblank_flag) in [PPUSTATUS (\$2002)](../PPU_registers/#PPUSTATUS), that register will see multiple extra reads that result in data loss or corruption. This is mostly a problem for controllers, and special [controller reading code](../Controller_reading_code/#DPCM_Safety) is necessary to work around this.

This problem is fixed on the 2A07 series of PAL CPUs by forcing DMA to halt on the first cycle of an instruction, which fetches an instruction opcode and thus isn't normally reading from registers.

See [DMA#DMC DMA](../DMA/#DMC_DMA) for a thorough explanation of how DMA works and [how it conflicts with register reads](../DMA/#Register_conflicts).

## Usage of DMC for syncing to video {#Usage_of_DMC_for_syncing_to_video}

The DMC can generate [CPU interrupt](../IRQ/) when it loaded the last sample byte to the [sample buffer](#Memory_reader), and can play short enough sample that is *shorter than a video frame*. Using this, the DMC channel can be used to perform mid-frame video stuffs that usually required [sprite zero hit flag](../PPU_registers/#Sprite_0_hit_flag), [sprite overflow flag](../PPU_registers/#Sprite_overflow_flag) or external [mapper](../Mapper/) timer IRQ.

The basic concept is about essentially playing a short silent sample at the fixed point of the screen (like at the very beginning of the [PPU NMI](../NMI/), however because how the [playback unit](#Output_unit) works, the sample playback **does not begin immediately** after you toggle the bit in [\$4015](#Sample_playback_control.2C_status_.28.244015.2C_read.2Fwrite.29), and it can take maximum of 8 timer clocks until the [shift register](#Output_unit) is ready.

The easiest and straightfoward way to work around of the issue is to **measure the CPU cycles until it actually begins the playback and use that to cancel out the variable delay inside the generated IRQ**.

At a fixed point in video rendering, a dummy single-byte sample at rate \$F is started. Due to a hardware quirk†, the sample needs to be started three times in a row like this:

    sei
    lda #$10 
    sta $4015 
    sta $4015 
    sta $4015 
    cli

- The amount of cycles before a DMC IRQ happens is then measured (either using an actual IRQ, or by polling \$4015).
  - At rate \$F, there are 54 CPU cycles between clocks, so there are 432 CPU cycles between boundaries.
- The main sample that will be used for the timing is then started (please refer to the table below to have sample lengths for various waiting times)
- When the main IRQ happens, the measurement from before is retrieved, and a timing loop with variable delay is used. In order to synchronize with Vblank, after a DMC IRQ we should wait 432 CPU cycles minus the time we measured.

†**Note:** The hardware quirk mentioned above deals with how DMC IRQs are generated. Basically, the IRQ is generated when the last **byte** of the sample is **read**, **not** when the last *sample* of the sample *plays*. The sample buffer sometimes has enough time to empty itself between writes to \$4015, meaning your next write to \$4015 will trigger an immediate IRQ. [Fortunately, writing to \$4015 three times will avoid this issue.](https://forums.nesdev.org/viewtopic.php?p=66712#p66712)

Still using Vblank as an example, the measurement tells how far into the 8-clock boundary Vblank occurred, and by delaying after a DMC IRQ, we perform a raster effect at the same point within the 8-clock boundary, aligning it with Vblank. By performing this same method each frame, the raster effect will have a reasonably stable timing to it. As a bonus, since mostly using ~~IRQs~~ are being used, the CPU is free to do something else, instead of waiting in a timed loop.

It's possible to use more than one [IRQ](../IRQ/) per frame - but the *measurement* part needs to be done at the *same time* within each frame, before the usage of any [IRQ](../IRQ/).

Only a single split-point per [IRQ](../IRQ/) is possible, with the shortest [IRQ](../IRQ/) being 3.8 scanlines. For split points closer than this amount, timed code has to be used.

In order to remain silent, samples should be made up of all \$00 or \$FF bytes, and \$00 or \$7F should have been previously written to [\$4011](#Direct_counter_load_.28.244011.2C_write.29). Otherwise, audio will unintentionally be created. This *is* a sound channel, after all.

### Timing table {#Timing_table}

This table converts sample length in scanline length (all values are rounded to the higher integer).

    NTSC               Rate 
    Length              $0    $1   $2   $3   $4   $5   $6   $7   $8   $9   $a   $b   $c   $d   $e   $f 
    ---------------------------------------------------------------------------------------------------- 
    1-byte (8 bits)     31    27   24   23   21   18   16   16   14   12   10   10   8    6    6    4 
    17-byte (136 bits)  **    **   **   **   **   **   **   **   228  192  170  154  127  101  87   65 
    33-byte (264 bits)  **    **   **   **   **   **   **   **   **   **   **   **   **   196  168  126 
    49-byte (392 bits)  **    **   **   **   **   **   **   **   **   **   **   **   **   **   **   187 

    PAL                Rate 
    Length              $0    $1   $2   $3   $4   $5   $6   $7   $8   $9   $a   $b   $c   $d   $e   $f 
    ---------------------------------------------------------------------------------------------------- 
    1-byte (8 bits)     30    27   24   23   21   18   16   15   14   12   10   9    8    6    5    4 
    17-byte (136 bits)  **    **   **   **   **   **   **   **   225  189  169  151  126  100  85   64 
    33-byte (264 bits)  **    **   **   **   **   **   **   **   **   **   **   **   **   194  164  124 
    49-byte (392 bits)  **    **   **   **   **   **   **   **   **   **   **   **   **   **   **   184

### Number of scanlines to wait table {#Number_of_scanlines_to_wait_table}

This table gives the best sample length and frequency combinations for all possible scanlines interval to wait. They are best because they are where the CPU will have to kill the least time. However, it's still possible to use options to wait for fewer lines and kill more time during the interrupt before the video effect.

Because a PAL interrupt will always happen about the same time or a bit sooner than a NTSC interrupt, the NTSC table will be used to set the "best" setting here :

    Scanlines  Best opt. for IRQ 

    1-3        Timed code 
    4-5        Length $0, rate $f 
    6-7        Length $0, rate $d 
    8-9        Length $0, rate $c 
    10-11      Length $0, rate $a 
    12-13      Length $0, rate $9 
    14-15      Length $0, rate $8 
    16-17      Length $0, rate $6 
    18-20      Length $0, rate $5 
    21-22      Length $0, rate $4 
    23         Length $0, rate $3 
    24-26      Length $0, rate $2 
    27-30      Length $0, rate $1 
    31-64      Length $0, rate $0 
    65-86      Length $1, rate $f 
    87-100     Length $1, rate $e 
    101-125    Length $1, rate $d 
    126        Length $2, rate $f 
    127-153    Length $1, rate $c 
    154-167    Length $1, rate $b 
    168-169    Length $2, rate $e 
    170-186    Length $1, rate $a 
    187-191    Length $3, rate $f 
    192-195    Length $1, rate $9 
    196-227    Length $2, rate $d 
    228-239    Length $1, rate $8

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
<td><a href="../APU_Pulse/">Pulse 1 and Pulse 2</a> | <a href="../APU_Triangle/">Triangle</a> | <a href="../APU_Noise/">Noise</a> | DMC</td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Components</strong></td>
<td><a href="../APU_Envelope/">Envelope</a> | <a href="../APU_Pulse/">Sweep</a> | <a href="../APU_Frame_Counter/">Frame Counter</a> | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | <a href="../APU_Mixer/">Mixer</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
