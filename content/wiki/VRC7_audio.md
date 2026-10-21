---
title: "VRC7 audio"
url: "/wiki/VRC7_audio/"
description: "The Konami VRC7, in addition to being a mapper chip, also produces 6 channels of 2-operator FM Synthesis Audio. It is a derivative of the Yamaha YM2413 OPLL, implementing a subset of its features and…"
wikiPage: "VRC7_audio"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
---

The [Konami VRC7](../VRC7/), in addition to being a mapper chip, also produces 6 channels of 2-operator FM Synthesis Audio. It is a derivative of the Yamaha YM2413 OPLL, implementing a subset of its features and containing a custom fixed patch set.

VRC7 audio was only used in one game, *Lagrange Point*. The chip also appears in *Tiny Toon Adventures 2*, but this cart does not use the audio, and its board lacks required additional audio mixing circuitry.

## Registers {#Registers}

### <span id="Audio_Reset_.28.24E000.29"></span>Audio Reset (\$E000) {#Audio_Reset_.28.24E000.29}

    7  bit  0
    ---- ----
    RSxx xxMM
     |
     +-------- Reset expansion sound

Setting this bit will silence the expansion audio and clear its registers (including tremolo LFO state, but not including vibrato LFO state).<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

Writes to \$9010 and \$9030 are disregarded while this bit is set.

Other bits in this register control mirroring and WRAM. See [VRC7: Mirroring Control (\$E000)](../VRC7/#Mirroring_Control_.28.24E000.29).

### <span id="Audio_Register_Select_.28.249010.29"></span>Audio Register Select (\$9010) {#Audio_Register_Select_.28.249010.29}

    7  bit  0
    ---- ----
    VVVV VVVV
    |||| ||||
    ++++-++++- The 8-bit internal register to select for use with $9030

This register is write-only.

After writing to this register, the program must not write to \$9030 (or \$9010 again) for at least 6 CPU clock cycles while the VRC7 internally sets up the address.

### <span id="Audio_Register_Write_.28.249030.29"></span>Audio Register Write (\$9030) {#Audio_Register_Write_.28.249030.29}

    7  bit  0
    ---- ----
    DDDD DDDD
    |||| ||||
    ++++-++++- The 8-bit value to write to the internal register selected with $9010

This register is write-only.

After writing to this register, the program must not write to \$9010 (or \$9030 again) for at least 42 CPU clock cycles while the VRC7 internally handles the write.

Addresses \$9010 and \$9030 differ in bit A5. This signal corresponds to A0 (pin 10) of a YM2413.

### Register Write Delay {#Register_Write_Delay}

Lagrange Point uses the following delay routine after writes to \$9010 and \$9030 to ensure enough time passes between writes:

    wait_9030:          ; JSR to this label immediately after writing to $9030
        STX $82         ; stores X temporarily (address is arbitrary)
        LDX #$08
    @wait_loop:
        DEX
        BNE @wait_loop
        LDX $82         ; restores X
    wait_9010:          ; JSR to this label immediately after writing to $9010
        RTS

### Internal Audio Registers {#Internal_Audio_Registers}

The VRCVII appears to have 26 internal registers. Registers \$00-\$07 define a custom patch that can be played on any channel set to use instrument \$0. Registers \$10-\$15, \$20-25, and \$30-35 control 6 channels for FM synthesis.

Other register values are ignored, except \$0F which is a test register with a few special functions.

#### Custom Patch {#Custom_Patch}

| Register | Bitfield | Description |
|----|----|----|
| \$00 | `TVSK MMMM` | Modulator tremolo (T), vibrato (V), sustain (S), key rate scaling (K), multiplier (M) |
| \$01 | `TVSK MMMM` | Carrier tremolo (T), vibrato (V), sustain (S), key rate scaling (K), multiplier (M) |
| \$02 | `KKOO OOOO` | Modulator key level scaling (K), output level (O) |
| \$03 | `KK-Q WFFF` | Carrier key level scaling (K), unused (-), carrier waveform (Q), modulator waveform (W), feedback (F) |
| \$04 | `AAAA DDDD` | Modulator attack (A), decay (D) |
| \$05 | `AAAA DDDD` | Carrier attack (A), decay (D) |
| \$06 | `SSSS RRRR` | Modulator sustain (S), release (R) |
| \$07 | `SSSS RRRR` | Carrier sustain (S), release (R) |

The patch defines a 2-operator FM unit with a single modulator and carrier. The carrier produces the output tone, and the output of the modulator modulates its frequency. The patch has the following parameters:

- **\$00/\$01 T** Tremolo applies amplitude modulation at a predefined rate.
- **\$00/\$01 V** Vibrato applies pitch modulation at a predefined rate.
- **\$00/\$01 S** Sustain determines whether the operator uses the sustain section of the envelope or not.
- **\$00/\$01 K** Key rate scaling adjusts the ADSR envelope speed, faster for high frequencies, slower for low ones.
- **\$00/\$01 MMMM** Multiplier is a multiplier on the operator's frequency according to a lookup table (see below).
- **\$02/\$03 KK** Key level scaling attenuates the operator at higher frequencies \$0 = none, \$3 = most.
- **\$02 OOOOOO** Modulator output level, this value reduces the modulator volume in 0.75db increments.
- **\$03 Q/W** Operator waveform, 0 = sine, 1 = half-wave rectified sine (where sine values less than 0 are clipped to 0).
- **\$03 FFF** Feedback applied to modulator according to a lookup table (see below).
- **\$04/\$05 AAAA** Attack is the speed of the attack fade in after key-on. \$0 = halt, \$1 = slowest, \$F = fastest.
- **\$04/\$05 DDDD** Decay is the speed of the decay fade to sustain after attack. \$0 = halt, \$1 = slowest, \$F = fastest.
- **\$06/\$07 SSSS** Sustain is the attenuation after decay, in 3db increments. \$0 = highest volume, \$F = lowest.
- **\$06/\$07 RRRR** Release is the speed of the release fade to silent after sustain. \$0 = halt, \$1 = slowest, \$F = fastest.

If a note is released before the attack or decay finishes, release begins from the current volume level. If the sustain bit is not set in **\$00/\$01 S**, release begins immediately after decay.

If the sustain bit is set in the channel control register **\$2X S** (see *Channels* section below), the release value in the patch is ignored and replaced with \$5.

The modulator's sustain bit (**\$00 S**) also disables the release section of its envelope. If its sustain bit is set, the Attack, Decay, and Sustain portions of the envelope are used, but when the note is released the modulator will continue to sustain while the carrier releases. The carrier does not behave this way: its envelope always enters release when the note is released.

| \$00/\$01 MMMM | \$0 | \$1 | \$2 | \$3 | \$4 | \$5 | \$6 | \$7 | \$8 | \$9 | \$A | \$B | \$C | \$D | \$E | \$F |
|----------------|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| Multiplier     | 1/2 | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 10  | 12  | 12  | 15  | 15  |

| \$03 FFF         | \$0 | \$1  | \$2 | \$3 | \$4 | \$5 | \$6 | \$7 |
|------------------|-----|------|-----|-----|-----|-----|-----|-----|
| Modulation Index | 0   | π/16 | π/8 | π/4 | π/2 | π   | 2π  | 4π  |

#### Channels {#Channels}

| Register | Bitfield | Description |
|----|----|----|
| \$10-\$15 | `LLLL LLLL` | Channel low 8 bits of frequency |
| \$20-\$25 | `--ST OOOH` | Channel sustain (S), trigger (T), octave (O), high bit of frequency (H) |
| \$30-\$35 | `IIII VVVV` | Channel instrument (I), volume (V) |

Each channel **X** is controlled by three registers at **\$1X**, **\$2X**, and **\$3X**.

The 8 bits of **\$1X** with a 9th bit from bit 0 of **\$2X** create a 9-bit frequency value (*freq*). This is combined with a 3-bit octave value from **\$2X** (*octave*) to define the output frequency (*F*):

         49716 Hz * freq
    F = -----------------
         2^(19 - octave)

The VRC7 is clocked by an external ceramic oscillator running at 3.58 MHz (roughly twice the NTSC NES CPU [clock rate](../Cycle_reference_chart/#Clock_Rates), but not synchronized with it), and it takes 72 internal clock cycles to update all of its channels, which means each channel is updated with a frequency of 49716 Hz (or roughly 36 NES CPU clocks). During these 72 cycles, the channels are output in series rather than mixed (like [Namco's 163](../Namco_163_audio/)), but because the frequency is so high it is not audibly different from mixed output.

Writing to register **\$2X** can begin a key-on or key-off event, depending on the value in the trigger bit (T). If the trigger bit changes from 0 to 1 (key-on), a new note begins. If it changes from 1 to 0 (key-off), it will begin the release portion of its envelope which will eventually silence the channel (unless release is \$0). If the trigger bit does not change, no new key-on or key-off will be generated.

Using the sustain bit (S) in **\$2X** overrides the normal release value for the patch with a value of \$5.

Register **\$3X** selects the instrument patch to use, and chooses a volume. Note that volume value is inverted: \$F is the lowest volume and \$0 is the highest. There is no silent volume value; its output scale is logarithmic in 3db increments.

#### <span id="Test_Register_.240F"></span>Test Register \$0F {#Test_Register_.240F}

At \$0F is an additional test register that responds to the low 4 bits.<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

- Bit 0: The envelope generators are replaced with constant 0 output (full volume) for both modulator and carrier. The envelopes are still running while their output is overridden.
- Bit 1: Hold LFO phase at zero. This halts, disables, and resets both the tremolo and vibrato LFO.
- Bit 2: Holds and resets waveform phase to zero. The envelopes are not halted, though the output will be silent.
- Bit 3: Update tremolo and vibrato LFOs every sample instead of once every several samples. (Tremolo is 64x faster, and vibrato is 1024x faster.)

## Internal patch set {#Internal_patch_set}

There are 16 different instrument patches available on the VRC7. With the exception of instrument \$0, which can be controlled by registers \$00-\$07 (see above), these are hardwired into the chip and cannot be altered.

Exact values for the fixed patch were dumped using the VRC7 rewired into a debug mode<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>:

       | 00 01 02 03 04 05 06 07 | Name
    ---+-------------------------+----------------
     0 | -- -- -- -- -- -- -- -- | (Custom Patch)
     1 | 03 21 05 06 E8 81 42 27 | Buzzy Bell
     2 | 13 41 14 0D D8 F6 23 12 | Guitar
     3 | 11 11 08 08 FA B2 20 12 | Wurly
     4 | 31 61 0C 07 A8 64 61 27 | Flute
     5 | 32 21 1E 06 E1 76 01 28 | Clarinet
     6 | 02 01 06 00 A3 E2 F4 F4 | Synth
     7 | 21 61 1D 07 82 81 11 07 | Trumpet
     8 | 23 21 22 17 A2 72 01 17 | Organ
     9 | 35 11 25 00 40 73 72 01 | Bells
     A | B5 01 0F 0F A8 A5 51 02 | Vibes
     B | 17 C1 24 07 F8 F8 22 12 | Vibraphone
     C | 71 23 11 06 65 74 18 16 | Tutti
     D | 01 02 D3 05 C9 95 03 02 | Fretless
     E | 61 63 0C 00 94 C0 33 F6 | Synth Bass
     F | 21 72 0D 00 C1 D5 56 06 | Sweep

The VRC7 instrument ROM dump also shows 3 drum patches. It is believed that these additional patches are an artifact from the YM2413 and are not playable on the VRC7. Curiously, byte \$07 of the snare drum (\$68) differs from YM2413 (\$48):

       | 01 01 18 0F DF F8 6A 6D | Bass Drum
       | 01 01 00 00 C8 D8 A7 68 | Snare Drum / Hi-Hat
       | 05 01 00 00 F8 AA 59 55 | Tom / Top Cymbal

## Debug Mode {#Debug_Mode}

Pin 15 of the VRC7 is an active-low input that enables a debug mode. In debug mode, it is possible to dump instrument patches from internal ROM, read and write the DAC value, and access serial data similar to the Yamaha YM2413 and other related synthesizers. The VRC7 uses the normal CPU data bus for reading and writing data in this mode, but uses a modified pinout to select the data. The modified pinout is not compatible with the normal Famicom connections in a cartridge. Since the /Debug pin will be tied high in a cartridge, this mode is not normally accessible without physical modification.

Pins 2 and 47 select a high-level mode, while pins 4, 16, and 28 change meaning depending on which mode is selected.

| Mode | Pin 15 | Pin 47 | Pin 2 | Pin 28 | Pin 4 | Pin 16 |  |
|----|----|----|----|----|----|----|----|
| VRC7 | /DEBUG | CPU M2 | PPU A13 | CPU A13 | CPU R/W | CPU A5 |  |
| Normal usage | 1 | x | x | x | x | x |  |
| YM2413 | /DEBUG | MODE0 | MODE1 | /CS | /WR | A0 | Data bus |
| No operation | 0 | 1 | 0 | 1 | x | x |  |
| Write Address | 0 | 1 | 0 | 0 | 0 | 0 |  |
| Write Data | 0 | 1 | 0 | 0 | 0 | 1 |  |
| Read Serial | 0 | 1 | 0 | 0 | 1 | 0 | `---- --DF` |
| No operation | 0 | 1 | 0 | 0 | 1 | 1 |  |
| DAC Test | /DEBUG | MODE0 | MODE1 | /CS | /WR | A0 | Data bus |
| No operation | 0 | 0 | 0 | 1 | x | x |  |
| Override DAC | 0 | 0 | 0 | 0 | 0 | x | Lower 8 bits |
| No operation | 0 | 0 | 0 | 0 | 1 | x |  |
| DAC Query | /DEBUG | MODE0 | MODE1 | /CS | /WR | A0 | Data bus |
| Read DAC | 0 | 1 | 1 | x | x | x | Upper 8 bits |
| Instrument Query | /DEBUG | MODE0 | MODE1 | /A2 | /A1 | /A0 | Data bus |
| Read InstROM 0 | 0 | 0 | 1 | 1 | 1 | 1 | `SSSS RRRR` |
| Read InstROM 1 | 0 | 0 | 1 | 1 | 1 | 0 | `AAAA DDDD` |
| Read InstROM 2 | 0 | 0 | 1 | 1 | 0 | 1 | `MMMM --KK` |
| Read InstROM 3 | 0 | 0 | 1 | 1 | 0 | 0 | `-FFF TVSK` |
| Read InstROM 4 | 0 | 0 | 1 | 0 | 1 | 1 | `OOOO OOQW` |
| Read 0 | 0 | 0 | 1 | 0 | 1 | 0 | ([legend](#Custom_Patch)) |
| Read 0 | 0 | 0 | 1 | 0 | 0 | 1 |  |
| Read 0 | 0 | 0 | 1 | 0 | 0 | 0 |  |

VRC7 Debug Mode Inputs

When /DEBUG (pin 15) is low, CPU A12 (pin 27) acts as an active-low Reset signal (instead of the active-high signal in register \$E000) for the YM2413, meaning that it must be high at all times for the audio circuit to be driven.

In "Read Serial" mode, data is returned on the CPU data bus as follows:

- CPU D0: lower 9 bits of phase and envelope amplitude, as seen in the normal self-test mode of the YM2413
- CPU D1: DAC value
- CPU D2-D7: logic 0

In "Override DAC" mode, the lower 8 bits of the DAC are instead driven with the contents of the external data bus. The DAC goes back to normal as soon as this mode is exited.

In "Read InstROM" mode, the current operator's parameters are returned, using the same legend as used above in [§Custom Patch](#Custom_Patch). Any bit marked with "-" reads as 0.

#### <span id="Rhythm_Register_.240E"></span>Rhythm Register \$0E {#Rhythm_Register_.240E}

In normal operation, [the "rhythm mode" bit in register \$0E](https://www.smspower.org/maxim/Documents/YM2413ApplicationManual#iii17) is treated as though it were always enabled, resulting only six audible FM channels. The VRC7 has no rhythm DAC, so the 5 rhythm channels are always inaudible.

In debug mode, the rhythm mode bit can be disabled, permitting 9 channels of OPLL audio. (It can also be enabled, but without the rhythm DAC, there is almost no utility in doing so)

## Differences from OPLL {#Differences_from_OPLL}

The synthesis core is related to the Yamaha YM2413 OPLL, which is itself a cost-reduced version of the YM3182 OPL2 chip made popular by AdLib and SoundBlaster sound cards.

- Register layout is the same.
- VRC7 has 6 channels, OPLL has 9.
- VRC7 has no rhythm channels, OPLL does (the last 3 channels are either FM or Rhythm on OPLL).
- VRC7 built-in instruments are not the same as OPLL instruments.
- VRC7 has no readily-accessible status register, under normal circumstances it is write-only; OPLL has an undocumented, 2-bit 'internal state' register.
- VRC7 has one DAC and output pin for audio, multiplexed for all 6 channels; OPLL has two DACs and output pins, one for FM and one for Rhythm.

A recording comparing the fixed patches of VRC7(left) and YM2413(right) is available here: [File:VRC7 and YM2413 on FAMICOM LP intro.ogg](../../wiki-images/VRC7_and_YM2413_on_FAMICOM_LP_intro.ogg)

## References {#References}

- VRC7 chip info by Kevin Horton: <http://kevtris.org/nes/vrcvii.txt>
- YM2413 datasheets: <http://www.datasheetarchive.com/YM2413-datasheet.html>
- VRC7 audio test program: <http://zzo38computer.org/nes_program/vrc7test.zip>
- VRC7 patch set revisions 2012, 2017, 2018: <http://forums.nesdev.org/viewtopic.php?f=6&t=9141>
- Discussion of modulator sustain: <http://famitracker.com/forum/posts.php?id=6804>

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [Forum post:](https://forums.nesdev.org/viewtopic.php?p=224897#p224897) VRC7 and 5B amplifier investigation, and function of silence bit.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [VRC7 test register described](https://forums.nesdev.org/viewtopic.php?p=236371#p236371)
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) [siliconpr0n: VRC7 Instrument ROM Dump](https://siliconpr0n.org/archive/doku.php?id=vendor:yamaha:opl2#vrc7_instrument_rom_dump)

## External links {#External_links}

- [Programmer's guide to the YM2413 from SMSPower](http://www.smspower.org/maxim/Documents/YM2413ApplicationManual)
- [How logsin and exp tables work in OPL2](https://docs.google.com/Doc?id=dd8kqn9f_13cqjkf4gp)
- [Yamaha FM internal bit depths](http://gendev.spritesmind.net/forum/viewtopic.php?t=386)
- [VRC7 discussion on NESdev BBS](http://forums.nesdev.org/viewtopic.php?f=3&t=9102)
- [siliconpr0n VRC7 decapsulation photographs](https://siliconprawn.org/map/konami/vrc-vii-053982/mcmaster_mz_mit20x/)
