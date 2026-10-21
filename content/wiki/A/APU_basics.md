---
title: "APU basics"
url: "/wiki/APU_basics/"
description: "This tutorial covers basic usage of the APU's four waveform channels on an NTSC NES. It does not cover the DMC or more advanced usage. Any registers unrelated to basic operation are not even…"
wikiPage: "APU_basics"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

This tutorial covers basic usage of the APU's four waveform channels on an NTSC NES. It does not cover the DMC or more advanced usage. Any registers unrelated to basic operation are not even mentioned here. A simplified though fully usable model of the APU is presented, one that will serve many programmers.

## Channels {#Channels}

The APU has five channels: two pulse waves, triangle wave, noise, and DMC (sample playback). Only the first four are covered here.

The channel registers begin at \$4000, and each channel has four registers devoted to it. All but the triangle wave have 4-bit volume control (the triangle just has an un-mute flag).

|               |                   |
|---------------|-------------------|
| \$4000-\$4003 | First pulse wave  |
| \$4004-\$4007 | Second pulse wave |
| \$4008-\$400B | Triangle wave     |
| \$400C-\$400F | Noise             |

In register descriptions below, bits listed as - can have any value written to them, while bits listed as 1 **must** have a 1 written, otherwise other APU features will be enabled, causing the registers to behave differently than described here.

## Register initialization {#Register_initialization}

Before using the APU, first initialize all the registers to known values that silence all channels.

    init_apu:
            ; Init $4000-4013
            ldy #$13
    @loop:  lda @regs,y
            sta $4000,y
            dey
            bpl @loop

            ; We have to skip over $4014 (OAMDMA)
            lda #$0f
            sta $4015
            lda #$40
            sta $4017

            rts
    @regs:
            .byte $30,$08,$00,$00
            .byte $30,$08,$00,$00
            .byte $80,$00,$00,$00
            .byte $30,$00,$00,$00
            .byte $00,$00,$00,$00

The initialization above prepares the APU to a known state, ready to be used by the examples below. In particular, it disables hardware [sweep](../APU_Pulse/), [envelope](../APU_Envelope/), and [length](../APU_Length_Counter/), which this tutorial does not use.

## Pulse wave channels {#Pulse_wave_channels}

*Main article: [APU Pulse](../APU_Pulse/)*

There are two pulse wave channels, each with pitch, volume, and timbre controls.

<table class="tabular">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<tbody>
<tr>
<td>$4000</td>
<td>$4004</td>
<td><code>%DD11VVVV</code></td>
<td>Duty cycle and volume<br />
&#10;<p>DD: 00=12.5% 01=25% 10=50% 11=75%<br />
VVVV: 0000=silence 1111=maximum</p></td>
</tr>
<tr>
<td>$4002</td>
<td>$4006</td>
<td><code>%LLLLLLLL</code></td>
<td>Low 8 bits of raw period</td>
</tr>
<tr>
<td>$4003</td>
<td>$4007</td>
<td><code>%-----HHH</code></td>
<td>High 3 bits of raw period</td>
</tr>
</tbody>
</table>

To determine the raw period for a given frequency in Hz, use this formula (round the result to a whole number)::

raw period = 111860.8/frequency - 1

The following code plays a 400 Hz square wave (50% duty) at maximum volume:

    jsr init_apu

    lda #<279
    sta $4002

    lda #>279
    sta $4003

    lda #%10111111
    sta $4000

All parameters can be changed while the tone is playing. To fade a note out, for example, write to \$4000 or \$4004 with the lower 4 bits decreasing every few frames.

Note that writing to \$4003 and \$4007 resets the phase, which causes a slight pop. This is an issue when doing vibrato, for example, and beyond the scope of this article.

## Triangle wave channel {#Triangle_wave_channel}

*Main article: [APU Triangle](../APU_Triangle/)*

The triangle channel allows control over frequency and muting.

|        |             |                             |
|--------|-------------|-----------------------------|
| \$4008 | `%1U------` | Un-mute                     |
| \$400A | `%LLLLLLLL` | Low 8 bits of raw period    |
| \$400B | `%-----HHH` | High 3 bits of raw period   |
| \$4017 | `%1-------` | Apply un-muting immediately |

For any given period, the triangle channel's frequency is half that of the pulse channel, or a pitch one octave lower. To determine the raw period for a given frequency in Hz, use this formula (round the result to a whole number):

raw period = 55930.4/frequency - 1

The following code plays a 400 Hz triangle wave:

    jsr init_apu

    lda #<139
    sta $400A

    lda #>139
    sta $400B

    lda #%11000000
    sta $4008
    sta $4017

The raw period can be changed while the channel is playing.

To silence the wave, write %10000000 to \$4008 and then \$4017. Writing a raw period of 0 also silences the wave, but produces a pop, so it's not the preferred method.

## Noise channel {#Noise_channel}

*Main article: [APU Noise](../APU_Noise/)*

The noise channel allows control over frequency, volume, and timbre.

<table class="tabular">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<tbody>
<tr>
<td>$400C</td>
<td><code>%--11VVVV</code></td>
<td>Volume<br />
&#10;<p>VVVV: 0000=silence 1111=maximum</p></td>
</tr>
<tr>
<td>$400E</td>
<td><code>%T---PPPP</code></td>
<td>Tone mode enable, Period</td>
</tr>
</tbody>
</table>

The following code plays a tone-like noise at maximum volume:

    jsr init_apu

    lda #%10000101
    sta $400E

    lda #%00111111
    sta $400C

All parameters can be changed while the noise is playing.

## Playing a musical note {#Playing_a_musical_note}

To easily play a musical note, use the [APU period table](../APU_period_table/). The following code sets the first pulse wave's frequency based on the note number in the X register:

    ; Set first pulse channel's frequency to note code in X register
    lda periodTableHi,x
    sta $4003

    lda periodTableLo,x
    sta $4002

    ...

    ; NTSC period table generated by mktables.py
    periodTableLo:
      .byte $f1,$7f,$13,$ad,$4d,$f3,$9d,$4c,$00,$b8,$74,$34
      .byte $f8,$bf,$89,$56,$26,$f9,$ce,$a6,$80,$5c,$3a,$1a
      .byte $fb,$df,$c4,$ab,$93,$7c,$67,$52,$3f,$2d,$1c,$0c
      .byte $fd,$ef,$e1,$d5,$c9,$bd,$b3,$a9,$9f,$96,$8e,$86
      .byte $7e,$77,$70,$6a,$64,$5e,$59,$54,$4f,$4b,$46,$42
      .byte $3f,$3b,$38,$34,$31,$2f,$2c,$29,$27,$25,$23,$21
      .byte $1f,$1d,$1b,$1a,$18,$17,$15,$14

    periodTableHi:
      .byte $07,$07,$07,$06,$06,$05,$05,$05,$05,$04,$04,$04
      .byte $03,$03,$03,$03,$03,$02,$02,$02,$02,$02,$02,$02
      .byte $01,$01,$01,$01,$01,$01,$01,$01,$01,$01,$01,$01
      .byte $00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00
      .byte $00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00
      .byte $00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00,$00
      .byte $00,$00,$00,$00,$00,$00,$00,$00

The triangle plays an octave lower for the same raw period. There are two ways to compensate for this. One way is to halve the value from the above table to get the desired note:

    ; Set triangle frequency to note code in X register
    lda periodTableHi,x
    lsr a
    sta $400B

    lda periodTableLo,x
    ror a
    sta $400A

The other way is to read period values one octave later in the table:

    ; Set triangle frequency to note code in X register
    lda periodTableHi+12,x
    sta $400B

    lda periodTableLo+12,x
    sta $400A

The following full program plays pulse and triangle scales:

[apu_scale.s](http://blargg.8bitalley.com/parodius/nes-code/apu_scale.s)

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
<td><a href="../APU/">Registers</a> | Basics | <a href="../APU_period_table/">Period table</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Channels</strong></td>
<td><a href="../APU_Pulse/">Pulse 1 and Pulse 2</a> | <a href="../APU_Triangle/">Triangle</a> | <a href="../APU_Noise/">Noise</a> | <a href="../APU_DMC/">DMC</a></td>
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
