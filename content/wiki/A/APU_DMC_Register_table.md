---
title: "APU DMC/Register table"
url: "/wiki/APU_DMC/Register_table/"
description: "The rate (R bits) determines for how many CPU cycles happen between changes in the output level during automatic sample playback."
wikiPage: "APU_DMC_Register_table"
---

### <span id="Flags_and_sample_playback_rate_.28.244010.2C_write.29"></span>Flags and sample playback rate (\$4010, write) {#Flags_and_sample_playback_rate_.28.244010.2C_write.29}

    7  bit  0
    ---- ----
    IL.. RRRR
    ||   ||||
    ||   ++++- Playback rate index
    |+-------- Sample loop flag
    +--------- Enable DMC sample-end interrupt (hereinafter "DMC IRQ")

The rate (`R` bits) determines for how many CPU cycles happen between changes in the output level during [automatic sample playback](../#Sample_playback_unit_operation).

Setting the loop flag (`L` bit) inhibits DMC IRQ pending flag from being set.

#### Pitch table {#Pitch_table}

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

On RP2A07, rate \$4 and \$C appear to be incorrect with respect to their intended A-440 tuning scheme. [\[1\]](https://forums.nesdev.org/viewtopic.php?p=94079#p94079)

All "dendy timing" clone CPUs are known to use NTSC period table.

### <span id="Direct_counter_load_.28.244011.2C_write.29"></span>Direct counter load (\$4011, write) {#Direct_counter_load_.28.244011.2C_write.29}

    7  bit  0
    ---- ----
    .DDD DDDD
     ||| ||||
     +++-++++- Unsigned 7-bit counter value

Writing to this register directly sets the counter value (final output level). If the [automatic DMC playback unit](../#Sample_playback_unit_operation) clocks the counter at the same time, the output level is occasionally not changed properly. [\[2\]](https://forums.nesdev.org/viewtopic.php?p=104491#p104491)

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

The `D` bit returns 1 when the [bytes-remaining counter](../#Memory_reader) is not zero.

    7  bit  0 (write)
    ---- ----
    ...D NT21
       |
       +------ Start / Stop DMC sample playback

When writing to this register with the `D` bit *clear*, the [bytes-remaining counter](../#Memory_reader) is set to 0, consequently finishes the playback before the next sample byte is fetched. When the write is done with the bit *set* and if the [bytes-remaining counter](../#Memory_reader) is 0, the sample playback is restarted, by reloading the [address counter](../#Memory_reader) and [bytes-remaining counter](../#Memory_reader) based on [\$4012](#Sample_address_.28.244012.2C_write.29) and [\$4013](#Sample_length_.28.244013.2C_write.29).

Writing to this register also clears DMC IRQ pending flag.

For descriptions for other bits, see [APU#Register overview#APU master control and status (\$4015, read/write)](../../APU/#Register_overview.23APU_master_control_and_status_.28.244015.2C_read.2Fwrite.29).
