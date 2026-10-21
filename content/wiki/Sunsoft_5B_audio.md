---
title: "Sunsoft 5B audio"
url: "/wiki/Sunsoft_5B_audio/"
description: "The Sunsoft 5B is a superset of the Sunsoft FME-7. The IC is very close to FME-7, noticeably includes an extra audio. This audio hardware was only used in one game, Gimmick!. Because this game did…"
wikiPage: "Sunsoft_5B_audio"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
---

The **Sunsoft 5B** is a superset of the [Sunsoft FME-7](../Sunsoft_FME-7/). The IC is very close to FME-7, noticeably includes an extra audio. This audio hardware was only used in one game, **[Gimmick!](https://nescartdb.com/profile/view/2695/gimmick)**. Because this game did not use many features of the chip (e.g. noise, envelope), its features are often only partially implemented by emulators.

The audio hardware is a type of [Yamaha YM2149F SSG](../YM2149F_pinout/), which is itself a variant of the [General Instrument AY-3-8910 PSG](../YM2149F_pinout/).

## Registers {#Registers}

The two audio registers on the 5B correspond to writes to the internal YM2149F registers, first selecting the internal register through \$C000, then writing to it through \$E000.

### <span id="Audio_Register_Select_.28.24C000-.24DFFF.29"></span>Audio Register Select (\$C000-\$DFFF) {#Audio_Register_Select_.28.24C000-.24DFFF.29}

    7  bit  0
    ---- ----
    DDDD RRRR
    |||| ++++- The 4-bit internal register to select for use with $E000
    ++++------ Disable writes to $E000 if nonzero (like the original YM2149F)

### <span id="Audio_Register_Write_.28.24E000-.24FFFF.29"></span>Audio Register Write (\$E000-\$FFFF) {#Audio_Register_Write_.28.24E000-.24FFFF.29}

    7  bit  0
    ---- ----
    VVVV VVVV
    ++++-++++- The 8-bit value to write to the internal register selected with $C000

### Internal audio registers {#Internal_audio_registers}

The YM2149F has 16 internal audio registers, selected with \$C000 and written to with \$E000.

| Register | Bitfield | Description |
|----|----|----|
| \$00 | `LLLL LLLL` | Channel A low period |
| \$01 | `---- HHHH` | Channel A high period |
| \$02 | `LLLL LLLL` | Channel B low period |
| \$03 | `---- HHHH` | Channel B high period |
| \$04 | `LLLL LLLL` | Channel C low period |
| \$05 | `---- HHHH` | Channel C high period |
| \$06 | `---P PPPP` | Noise period |
| \$07 | `--CB Acba` | Noise disable on channels C/B/A, Tone disable on channels c/b/a |
| \$08 | `---E VVVV` | Channel A envelope enable (E), volume (V) |
| \$09 | `---E VVVV` | Channel B envelope enable (E), volume (V) |
| \$0A | `---E VVVV` | Channel C envelope enable (E), volume (V) |
| \$0B | `LLLL LLLL` | Envelope low period |
| \$0C | `HHHH HHHH` | Envelope high period |
| \$0D | `---- CAaH` | Envelope reset and shape: continue (C), attack (A), alternate (a), hold (H) |
| \$0E | `---- ----` | I/O port A (unused) |
| \$0F | `---- ----` | I/O port B (unused) |

## Sound {#Sound}

There are three channels that output a square wave tone. In addition there is one noise generator, and one envelope generator, both of which may be shared by any of the three channels.

The 5B's audio is driven by the [CPU clock](../Cycle_reference_chart/) (1.789773 MHz). It operates equivalent to a YM2149F with its SEL pin held low (see [datasheet](#References)). This causes the tone and noise channels to operate at half the speed of an AY-3-8190 with an equivalent clock.

To use an AY-3-8910 as a substitute, you would need an external divider to reduce the clock speed by half.

Unlike the [2A03](../APU_Pulse/) and [VRC6](../VRC6_audio/) pulse channels' frequency formulas, the formula for 5B does not add 1 to the period. A period value of 0 appears to produce the same result as a period value of 1, for tone<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>, noise and envelope<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>.

Correct behaviour can be implemented as a counter that counts up on every 16th clock cycle until it is equal to or greater than the period register, at which point the output flips and the counter resets to 0. (This means that shortening the period can cause an immediate flip if the phase counter is already past the new new period value.<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>)

None of the various generators (tone, noise, envelope) can be halted. Setting them to volume or disabling their output does not affect their internal continued operation<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)</sup>.

### Tone {#Tone}

The tone generators produce a square wave with a period controlled by the CPU clock and the 12-bit period value in registers \$00-05.

- *Frequency* = *Clock* / (32 \* *Period*)
- *Period* = *Clock* / (32 \* *Frequency*)

This means that the high/low state of the square wave is toggled every 16 clocks.

Register \$07 controls the mixing of tone and noise components of each channel. A bit of 0 enables the noise/tone on the specified channel, and a bit of 1 disables it. If both bits are 1, the channel outputs a constant signal at the specified volume. If both bits are 0, the result is the logical and of noise and tone.

If bit 4 of registers \$08-\$0A is set, the volume of the channel is controlled by the envelope generator (see below). Otherwise, it is controlled by the 4-bit value in bits 3-0 of the same register. Volume 0 is completely silent<sup><span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup>.

### Noise {#Noise}

The noise generator produces a 1-bit random output with a period controlled by the CPU clock and the 5-bit period value in register \$06.

- *Frequency* = *Clock* / (32 \* *Period*)
- *Period* = *Clock* / (32 \* *Frequency*)

This produces a new random bit every 32 clocks.

It is implemented as a 17-bit linear feedback shift register with taps at bits 16 and 13.<sup><span id="cite_ref-6"></span>[\[6\]](#cite_note-6)</sup>

### Envelope {#Envelope}

The envelope produces a ramp that can be directed up or down, or to oscillate by various shape parameters.

It has a 5-bit series of output levels. Its maximum level 31 corresponds to the 4-bit volume register level 15<sup><span id="cite_ref-7"></span>[\[7\]](#cite_note-7)</sup>. Every 2 down from maximum is equivalent to 1 step down in the 4-bit volume register, and finally envelope levels 0 and 1 are both equivalent to volume 0 (silent). The others fall in between at 1.5db logarithmic steps.

|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|----|
| 5-bit envelope | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
| 4-bit volume | 0 | 0 |  | 1 |  | 2 |  | 3 |  | 4 |  | 5 |  | 6 |  | 7 |  | 8 |  | 9 |  | 10 |  | 11 |  | 12 |  | 13 |  | 14 |  | 15 |

#### Period {#Period}

The ramp has a frequency controlled by the CPU clock and the 16-bit period value in registers \$0B-0C. Note this formula is the frequency of a single step of the ramp.

- *Frequency* = *Clock* / (16 \* *Period*)
- *Period* = *Clock* / (16 \* *Frequency*)

The 5B divides each ramp into 32 steps, so for continued ("sawtooth") envelope shapes the resulting frequency will be 1/32 of the step frequency, and for the continued alternating ("triangle") envelope shapes it will sound at 1/64 of the step frequency. Thus the overall period for the continued envelope has a factor of 512 (or 1024 if alternating), rather than just 16.

Because the envelope is primarily intended for low (sub-audio) frequencies, its pitch control is not as accurate in audio frequency ranges as the tone channels.

#### Shape {#Shape}

Writing register \$0D resets the envelope<sup><span id="cite_ref-8"></span>[\[8\]](#cite_note-8)</sup> and chooses its shape. The shape has four parameters: continue, attack, alternate, and hold.

- Continue specifies whether the envelope continues to oscillate after the attack. If it is 0, the alternate and hold parameters have no effect.

<!-- -->

- Attack specifies whether the attack goes from high to low (0) or low to high (1).

<!-- -->

- Alternate specifies whether the signal continues to alternate up and down after the attack. If combined with hold it provides an immediate flip after the attack followed by the hold.

<!-- -->

- Hold specifies that the value shall be held after the attack. If combined with alternate, the value at the end of the attack will be immediately flipped before holding.

| Value       | Continue | Attack | Alternate | Hold | Shape      |
|-------------|----------|--------|-----------|------|------------|
| \$00 - \$03 | 0        | 0      | x         | x    | `\_______` |
| \$04 - \$07 | 0        | 1      | x         | x    | `/_______` |
| \$08        | 1        | 0      | 0         | 0    | `\\\\\\\\` |
| \$09        | 1        | 0      | 0         | 1    | `\_______` |
| \$0A        | 1        | 0      | 1         | 0    | `\/\/\/\/` |
| \$0B        | 1        | 0      | 1         | 1    | `\¯¯¯¯¯¯¯` |
| \$0C        | 1        | 1      | 0         | 0    | `////////` |
| \$0D        | 1        | 1      | 0         | 1    | `/¯¯¯¯¯¯¯` |
| \$0E        | 1        | 1      | 1         | 0    | `/\/\/\/\` |
| \$0F        | 1        | 1      | 1         | 1    | `/_______` |

### Output {#Output}

The tone channels each produce a 5-bit signal which is then converted to analog with a logarithmic DAC. Note that the least significant bit cannot be controlled by the volume register, it is only used by the YM2149F's double-resolution envelope generator. The logarithmic curve increases by 1.5 decibels for each step in the 5-bit signal. This can easily be implemented as a lookup table.

Some emulator implementations that are based on the AY-3-8190 instead treat it as a 4-bit signal with a 3dB per step curve. Since the only extant 5B game does not use the envelope, the difference is unimportant unless accuracy is desired for homebrew 5B work.

The three output channels are mixed together linearly. The output is mixed with the 2A03 and amplified. It is very loud compared to other audio expansion carts.

The amplifier becomes nonlinear at higher amplitudes, and includes some filtering.<sup><span id="cite_ref-9"></span>[\[9\]](#cite_note-9)</sup>

## References {#References}

- YM2149 datasheet: <http://pdf1.alldatasheet.com/datasheet-pdf/view/103366/ETC/YM2149.html>
- GI AY-3-8910 datasheet: <http://www.speccy.org/hardware/datasheet/ay38910.pdf>

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [Period 0 verification for tone/noise](https://forums.nesdev.org/viewtopic.php?p=236701#p236701)
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Period 0 verification for envelope](https://forums.nesdev.org/viewtopic.php?p=236672#p236672)
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) [Phase counting verification](https://forums.nesdev.org/viewtopic.php?p=236734#p236734)
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) [No halt verification](https://forums.nesdev.org/viewtopic.php?p=236745#p236745)
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) [Volume 0 silence verification](https://forums.nesdev.org/viewtopic.php?p=236734#p236734)
6.  <span id="cite_note-6"></span>[↑](#cite_ref-6) [LFSR verification](https://forums.nesdev.org/viewtopic.php?p=236703#p236703)
7.  <span id="cite_note-7"></span>[↑](#cite_ref-7) [Envelope-Volume equivalence verification](https://forums.nesdev.org/viewtopic.php?p=236734#p236734)
8.  <span id="cite_note-8"></span>[↑](#cite_ref-8) [Envelope phase reset verification](https://forums.nesdev.org/viewtopic.php?p=236672#p236672)
9.  <span id="cite_note-9"></span>[↑](#cite_ref-9) [Sunsoft 5B amplifier investigation](https://forums.nesdev.org/viewtopic.php?f=2&t=17745) (ongoing)
