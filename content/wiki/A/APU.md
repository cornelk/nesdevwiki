---
title: "APU"
url: "/wiki/APU/"
description: "The NES APU (audio processing unit) is a programmable sound generator that is integrated inside RP2A03 and RP2A07. It provides total 5 individual sound channels."
wikiPage: "APU"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

The **NES APU (audio processing unit)** is a [programmable sound generator](https://en.wikipedia.org/wiki/Programmable_sound_generator) that is integrated inside [RP2A03 and RP2A07](../2A03/). It provides total 5 individual sound channels.

## Overview {#Overview}

*For a simple subset of APU functionality oriented toward music engine developers rather than emulator developers, see [APU basics](../APU_basics/).*

The APU includes five channels: two [pulse wave generators](#Pulse), a [triangle wave](#Triangle), [noise](#Noise), and a [delta modulation channel](#DMC) for playing DPCM samples.

Each channel has a variable-rate [timer](#Glossary) clocking a waveform generator, and various modulators driven by low-frequency clocks from the [frame counter](#Frame_Counter). The [DMC](#DMC) plays samples while the other channels play waveforms. Each sub-unit of a channel generally runs independently and in parallel to other units, and modification of a channel's parameter usually affects only one sub-unit and doesn't take effect until that unit's next internal cycle begins.

The read/write [status register](#Status_.28.244015.29) allows channels to be enabled and disabled, and their current [length counter status](#Length_Counter) to be queried.

The outputs from all the channels are combined using a [non-linear mixing](#Mixer) scheme.

### Notes {#Notes}

- This reference describes the *abstract* operation of the APU and is more suitable for NES software programmers. For accurate APU emulation development, individual channel pages have more details about the operations. Although the *exact* (logic transistor level) hardware implementation is not necessarily relevant to an emulator, but the [Visual 2A03](http://www.qmtpro.com/~nes/chipimages/visual2a03/) project can be used to determine detailed information about the hardware implementation.
- The [Famicom](../Family_Computer/) has an audio return loop on its [cartridge connector](../Cartridge_connector/#Pinout_of_60-pin_Famicom_consoles_and_cartridges) allowing extra audio from individual cartridges. See [Expansion audio](../../category/Expansion_audio/) for details on the audio produced by various [mappers](../Mapper/).
- For a basic usage guide to the APU, see [APU basics](../APU_basics/), or [Nerdy Nights: APU overview](../Nerdy_Nights__APU_overview/).
- The APU may have additional diagnostic features if CPU pin 30 is pulled high. See [diagram by Quietust](../../wiki-images/Apu_address.jpg), [CPU Test Mode](../CPU_Test_Mode/).

## Registers {#Registers}

| Registers | Channel | Units |
|----|----|----|
| \$4000–\$4003 | [**Pulse 1**](#Pulse) | [Timer](#Glossary), [Length counter](#Length_Counter), [Envelope](#Envelope), [Sweep](../APU_Pulse/) |
| \$4004–\$4007 | [**Pulse 2**](#Pulse) | [Timer](#Glossary), [Length counter](#Length_Counter), [Envelope](#Envelope), [Sweep](../APU_Pulse/) |
| \$4008–\$400B | [**Triangle**](#Triangle) | [Timer](#Glossary), [Length counter](#Length_Counter), [Linear counter](#Linear_Counter) |
| \$400C–\$400F | [**Noise**](#Noise) | [Timer](#Glossary), [Length counter](#Length_Counter), [Envelope](#Envelope), Linear feedback shift register |
| \$4010–\$4013 | [**DMC**](#DMC) | [Timer](#Glossary), [Memory reader](../APU_DMC/#Memory_reader), [Sample buffer](../APU_DMC/#Memory_reader), [Output unit](../APU_DMC/#Output_unit) |
| \$4015 | All | [Channel enable and status](#Status_.28.244015.29) |
| \$4017 | All | [Frame counter](#Frame_Counter) |

<span id="Pulse"></span>

### <span id="Pulse_.28.244000.E2.80.93.244007.29"></span>Pulse (\$4000–\$4007) {#Pulse_.28.244000.E2.80.93.244007.29}

- See [APU Pulse](../APU_Pulse/)

These two individual **pulse channels** produce a variable-width pulse wave, controlled by volume, [envelope](#Envelope), [length counter](#Length_Counter), and [sweep units](../APU_Pulse/).

|  |  |  |
|----|----|----|
| \$4000 / \$4004 | `DDLC VVVV` | Duty (**D**), [Envelope](#Envelope) loop / [Length counter](#Length_Counter) halt (**L**), Constant volume (**C**), Volume / [Envelope](#Envelope) divider (**V**) |
| \$4001 / \$4005 | `EPPP NSSS` | [Sweep unit](../APU_Pulse/): Enabled (**E**), Period (**P**), Negate (**N**), Shift (**S**) |
| \$4002 / \$4006 | `TTTT TTTT` | [Timer](#Glossary) low (**T**) |
| \$4003 / \$4007 | `LLLL LTTT` | [Length counter](#Length_Counter) load (**L**), [Timer](#Glossary) high (**T**), Restart [envelope](#Envelope) and pulse phase. |

- **\$4000**:
  - **D**: The width of the pulse is controlled by the duty bits in \$4000/\$4004
    - 0: 12.5% **`__^_______^_____`**
    - 1: 25.0% **`__^^______^^____`**
    - 2: 50.0% **`__^^^^____^^^^__`**
    - 3: 75.0% **`^^__^^^^^^__^^^^`**
  - **L**: 0 = One-shot, 1 = Infinite play.
    - If 0, the [length counter](#Length_Counter) will run and the [envelope](#Envelope) will not repeat when it ends. Because [length counter](#Length_Counter) runs in parallel with [envelope](#Envelope), the [length counter](#Length_Counter) should be loaded with a time longer than the length of the [envelope](#Envelope) to prevent it from being cut off early.
    - If 1, the [length counter](#Length_Counter) will be frozen at its current value, and the [envelope](#Envelope) will repeat forever. After reaching 0 in [envelope](#Envelope), it will restart at volume 15 at its next period.
    - The [length counter](#Length_Counter) and [envelope](#Envelope) units are clocked by the [frame counter](#Frame_Counter).
    - If the [length counter](#Length_Counter)'s current value is 0 the channel will be silenced whether or not this bit is set.
  - **C**: If **C** is set the volume will be a constant. The [envelope unit](#Envelope) still runs in background but it gets overrided by the current value of **V**. If **C** is clear, it uses the output from [envelope](#Envelope) instead, starting at volume 15 and lowering to 0 over time.
  - **V**: Sets the direct volume if constant, otherwise controls the rate which the [envelope](#Envelope) lowers.
- The frequency of the pulse channels is a division of the [CPU Clock](../Cycle_reference_chart/#Clock_rates) (*f<sub>CPU</sub>*; 1.789773MHz NTSC, 1.662607MHz PAL). The output frequency (*f*) of the generator can be determined by the 11-bit period value (*t*) written to \$4002–\$4003/\$4006–\$4007. If *t* \< 8, the corresponding pulse channel is silenced.
  - *f* = *f<sub>CPU</sub>* ÷ (16 × (*t* + 1))
  - *t* = (*f<sub>CPU</sub>* ÷ (16 × *f*)) − 1
- Writing to \$4003/\$4007 reloads the [length counter](#Length_Counter), restarts the [envelope](#Envelope), and resets the phase of the pulse generator. Because it resets the phase, music effects like vibrato should only write the low period register to avoid a phase reset click noise. At some pitches, particularly near A440, wide vibrato should normally be avoided (e.g. this flaw is heard throughout the Mega Man 2 ending).
- Registers \$4001/\$4005 control the [sweep unit](../APU_Pulse/). Enabling the [sweep](../APU_Pulse/) causes the pitch to constantly rise or fall. When the frequency reaches the end of the generator's range of output, the channel is silenced. See [APU Sweep#Muting](../APU_Pulse/#Muting) for details.
- The two pulse channels are combined in a nonlinear mix. See [mixer](#Mixer) below.

<span id="Triangle"></span>

### <span id="Triangle_.28.244008.E2.80.93.24400B.29"></span>Triangle (\$4008–\$400B) {#Triangle_.28.244008.E2.80.93.24400B.29}

- See [APU Triangle](../APU_Triangle/)

The **triangle channel** produces a quantized triangle wave. It has no volume control, but it has a [length counter](#Length_Counter) as well as a higher resolution [linear counter](#Linear_Counter) control (called "linear" since it uses the 7-bit value written to \$4008 directly instead of a lookup table like the [length counter](#Length_Counter)).

|  |  |  |
|----|----|----|
| \$4008 | `CRRR RRRR` | [Length counter](#Length_Counter) halt / [Linear counter](#Linear_Counter) control (**C**), [Linear counter](#Linear_Counter) load (**R**) |
| \$4009 | `---- ----` | Unused |
| \$400A | `TTTT TTTT` | [Timer](#Glossary) low (**T**) |
| \$400B | `LLLL LTTT` | [Length counter](#Length_Counter) load (**L**), [Timer](#Glossary) high (**T**), Set [linear counter](#Linear_Counter) reload flag |

- The triangle wave has 32 steps that output a 4-bit value.
- **C**: This bit controls both the [length counter](#Length_Counter) and [linear counter](#Linear_Counter) at the same time.
  - When the bit is set, this will stop the [length counter](#Length_Counter) in the same way as for the [pulse](#Pulse)/[noise](#Noise) channels.
  - When the bit is set, it prevents the [linear counter](#Linear_Counter)'s internal reload flag from clearing, which effectively halts it if \$400B is written after setting **C**.
  - The [linear counter](#Linear_Counter) silences the channel after a specified time with a resolution of ≈240Hz (or ≈200Hz in PAL). See [frame counter](#Frame_Counter) below.
  - Because both the [length](#Length_Counter) and [linear](#Linear_Counter) counters are be enabled at *the same time* and can halt the triangle *independently*, whichever has a longer setting is redundant.
- **R**: This reload value will be applied to the [linear counter](#Linear_Counter) on the next quarter [frame counter tick](#Frame_Counter), but only if its reload flag is set.
  - A write to \$400B is needed to raise the reload flag.
  - After a [frame counter](#Frame_Counter) tick applies the load value **R**, the reload flag will only be cleared if **C** is also clear, otherwise it will continually reload (i.e. halt).
- The pitch of the triangle channel is one octave below the [pulse channels](#Pulse) with an equivalent [timer](#Glossary) value (i.e. use the formula above but divide the resulting frequency by two).
- Silencing the triangle channel merely halts it. It will continue to output its last value rather than 0.
- There is no way to reset the triangle channel's phase.

<span id="Noise"></span>

### <span id="Noise_.28.24400C.E2.80.93.24400F.29"></span>Noise (\$400C–\$400F) {#Noise_.28.24400C.E2.80.93.24400F.29}

- See [APU Noise](../APU_Noise/)

The **noise channel** produces noise with a pseudo-random bit generator. It has a volume, [envelope](#Envelope), and [length counter](#Length_Counter) like the [pulse channels](#Pulse).

|  |  |  |
|----|----|----|
| \$400C | `--LC VVVV` | [Envelope](#Envelope) loop / [Length counter](#Length_Counter) halt (**L**), Constant volume (**C**), Volume / [Envelope](#Envelope) divider (**V**) |
| \$400D | `---- ----` | Unused |
| \$400E | `M--- PPPP` | Periodic noise mode (**M**), Noise period (**P**) |
| \$400F | `LLLL L---` | [Length counter](#Length_Counter) load (**L**) |

- The frequency of the noise is determined by a 4-bit value in \$400E, which loads a period from a lookup table (see [APU Noise](../APU_Noise/)).
  - Lower period value generates higher frequency noise and higher period value generates lower frequency noise.
- If **M** is set, the period of the random bit generation is drastically shortened, producing a buzzing tone (e.g. the metallic ding during *Solstice*'s gameplay). The actual timbre produced depends on whatever bits happen to be in the generator when it is switched to periodic, and is essentially random.

<span id="DMC"></span>

### <span id="DMC_.28.244010.E2.80.93.244013.29"></span>DMC (\$4010–\$4013) {#DMC_.28.244010.E2.80.93.244013.29}

- See [APU DMC](../APU_DMC/)

The **delta modulation channel (DMC)** provides automatic 1-bit [delta-encoded sample](https://en.wikipedia.org/wiki/Delta_modulation) playback unit with its output level also settable manually with dedicated register.

|  |  |  |
|----|----|----|
| \$4010 | `IL-- RRRR` | DMC sample-end IRQ enable (**I**), Loop (**L**), Sample rate (**R**) |
| \$4011 | `-DDD DDDD` | Load counter (**D**) |
| \$4012 | `AAAA AAAA` | Sample address (**A**) |
| \$4013 | `LLLL LLLL` | Sample length (**L**) |

- DPCM samples are stored as a stream of 1-bit deltas that control the 7-bit PCM counter that the channel outputs. A bit of 1 will increment the counter, 0 will decrement, and it will clamp rather than overflow if the 7-bit range is exceeded. It reads from the least-significant bit in the byte to the most-significant bit.
- DPCM samples may loop if the loop flag in \$4010 is set, and the DMC may be used to generate an [IRQ](../IRQ/) when the end of the sample is reached if its IRQ flag is set.
- The playback rate is controlled by register \$4010 with a 4-bit frequency index value (see [APU DMC](../APU_DMC/) for frequency lookup tables).
- DPCM samples must begin in the memory range \$C000–\$FFFF at an address set by register \$4012 (address = `$C000 + (A * 64)`).
- The length of the sample in bytes is set by register \$4013 (length = `(L * 16) + 1 bytes`).
- The playback control itself is done with [\$4015](#Status) register.

#### Other Uses {#Other_Uses}

- The \$4011 register can be used to play raw PCM samples by updating the counter value at a high frequency. Because this requires intensive use of the CPU, when used in games all other gameplay is usually halted to facilitate this.
- Because of the APU's [nonlinear mixing](#Mixer), a high value in the PCM counter reduces the volume of the triangle and noise channels. This is sometimes used to apply limited volume control to the triangle channel (e.g. *Super Mario Bros.* adjusts the counter between levels to accomplish this).
- The DMC's IRQ can be used as an [IRQ](../IRQ/)-based timer when the [mapper](../Mapper/) used does not contain one available. See [APU DMC#Usage of DMC for syncing to video](../APU_DMC/#Usage_of_DMC_for_syncing_to_video).

<span id="Status"></span>

### <span id="Master_Control_.2F_Status_.28.244015.29"></span>Master Control / Status (\$4015) {#Master_Control_.2F_Status_.28.244015.29}

The status register is used to enable and disable individual channels, control the [DMC](#DMC), and can read the status of [length counters](#Length_Counter) and APU interrupts.

|  |  |  |
|----|----|----|
| \$4015 write | `---D NT21` | Control [DMC](#DMC) (**D**), [Noise](#Noise) (**N**), [Triangle](#Triangle) (**T**), And [pulse channels](#Pulse) (**2**/**1**) |

- Writing a zero to any of the channel enable bits (NT21) will silence the specified channel by focing the [length counter](#Length_Counter) value to 0.
- If the DMC bit is clear, the bytes-remaining counter is set to 0 and the DMC will silence when it empties.
- If the DMC bit is set, the DMC sample will be restarted *only if* its bytes-remaining counter is 0. If there are bits remaining in the shift register, these will finish playing before the next sample is fetched.
- Writing to this register also clears the DMC sample-end interrupt flag regardless of the written value.
- [Power-up and reset](../CPU_power_up_state/#APU) have the effect of writing \$00, effeciently *silencing all channels* at reset.

|  |  |  |
|----|----|----|
| \$4015 read | `IF-D NT21` | DMC sample-end IRQ asserted (**I**), [Frame counter](#Frame_Counter) interrupt (**F**), [DMC](#DMC) sample playback active (**D**), [length counter](#Length_Counter) \> 0 (**N**/**T**/**2**/**1**) |

- **N**/**T**/**2**/**1** will return 1 if the corresponding [length counter](#Legnth_Counter) has not been halted through either expiring or a write of 0 to the corresponding bit. For the [triangle](#Triangle) channel, the status of the [linear counter](#Linear_Counter) is irrelevant.
- **D** will return 1 if the remaining sample byte is non-zero.
- Reading this register clears the frame counter interrupt flag (but not the DMC interrupt flag).
  - If an interrupt flag was set at the same moment of the read, it will read back as 1 but it will *not* be cleared.
- This register is internal to the [CPU](../CPU/) and so the external CPU data bus is disconnected when reading it. Therefore the returned value cannot be seen by the cartridge and so the value does not affect [open bus](../Open_bus_behavior/).
- Bit 5 is [open bus](../Open_bus_behavior/). Because the external bus is disconnected when reading \$4015, the open bus value comes from the last cycle that did not read \$4015.

<span id="Frame_Counter"></span>

### <span id="Frame_Counter_.28.244017.29"></span>Frame Counter (\$4017) {#Frame_Counter_.28.244017.29}

- See [APU Frame Counter](../APU_Frame_Counter/)

|  |  |  |
|----|----|----|
| \$4017 | `MI-- ----` | Mode (**M**, 0 = freerun, 1 = external synchronization mode), IRQ inhibit flag (**I**), Generate all clocks immediately as well if the M bit was written set. |

The **frame counter** provides reliable low frequency clocks (≈60/120/240Hz in NTSC or ≈50/100/200Hz in PAL) to [envelope](#Envelope), [sweep](#Sweep), [length counter](#Length_Counter) and [linear counter](#Linear_Counter) used in [pulse](#Pulse), [triangle](#Triangle) and [noise](#Noise) channels.

The following diagram illustrates the operation of the two modes, selected by bit 7 of \$4017:

    Each ticks occur in ≈240Hz rate (or ≈200Hz in PAL)
    Freerun    Ext. Sync.   Function
    ---------  -----------  -----------------------------
     - - - f    - - - - -    IRQ (if IRQ inhibit flag is clear)
     - l - l    - l - - l    Length counter and sweep
     e e e e    e e e - e    Envelope and linear counter

Despite the name being "Frame counter", the clocks generated have nothing to do with the [Vsync signal from the PPU](../PPU_registers/#Vblank_NMI) and it counts its counter by its own. The **external synchronization mode** exists to provide a way to actually synchronize the internal counter (and essentially all the APU hardware envelope units) with the [Vblank NMI](../PPU_registers/#Vblank_NMI).

When the mode is enabled via setting bit 7 in \$4017, it takes much longer before the fourth clock arrives than in freerun mode. Or in other words, this gives you about 7000 CPU cycles of time to *reset the sequence manually* before the counter wraps around by its own. Meaning, you can write to \$4017 with the mode bit set in the [Vblank NMI](../NMI/) to reliably synchronize the APU hardware envelope sequences with the actual frame timing.

For the [IRQ](../IRQ/) it generates, this interrupt is to synchronize the software music player program *to* the frame counter sequence (the opposite from external synchronization mode). When the sequencer mode is set to freerun mode (bit 7 of \$4017 clear), it can optionally generate an [interrupt](../IRQ/) which essentially is generated in every ≈60Hz (≈50Hz in PAL) cycles. Because the interrupt is generated by the frame counter itself, updating the music player here is synchronized with the APU hardware envelopes. Clearing the interrupt flag is done by writing to \$4017 with interrupt inhibit flag set, or it can be simply done by just *reading* [\$4015](#Status) once.

#### Length Counter {#Length_Counter}

- See [APU Length Counter](../APU_Length_Counter/)

The [pulse](#Pulse), [triangle](#Triangle), and [noise](#Noise) channels each have their own **length counter** unit that can automatically cut the channel in specified duration. It is clocked at ≈120Hz (≈100Hz in PAL) speed by the [frame counter](#Frame_Counter), and counts down to zero if counting is enabled. When the counter reaches zero, the channel is silenced. It can be loaded by writing a 5-bit value to the appropriate channel's length counter register, which will load an actual counter value from a lookup table.

<table class="wikitable sortable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th data-sort-type="number">Length counter index<br />
<a href="#Pulse">$4003</a>/<a href="#Pulse">$4007</a>/<a href="#Noise">$400F</a> bit 3-7</th>
<th data-sort-type="number">Number of ≈120Hz (≈100Hz in PAL)<br />
ticks until silencing the channel)</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>10</td>
</tr>
<tr>
<td>$01</td>
<td>254</td>
</tr>
<tr>
<td>$02</td>
<td>20</td>
</tr>
<tr>
<td>$03</td>
<td>2</td>
</tr>
<tr>
<td>$04</td>
<td>40</td>
</tr>
<tr>
<td>$05</td>
<td>4</td>
</tr>
<tr>
<td>$06</td>
<td>80</td>
</tr>
<tr>
<td>$07</td>
<td>6</td>
</tr>
<tr>
<td>$08</td>
<td>160</td>
</tr>
<tr>
<td>$09</td>
<td>8</td>
</tr>
<tr>
<td>$0A</td>
<td>60</td>
</tr>
<tr>
<td>$0B</td>
<td>10</td>
</tr>
<tr>
<td>$0C</td>
<td>14</td>
</tr>
<tr>
<td>$0D</td>
<td>12</td>
</tr>
<tr>
<td>$0E</td>
<td>26</td>
</tr>
<tr>
<td>$0F</td>
<td>14</td>
</tr>
<tr>
<td>$10</td>
<td>12</td>
</tr>
<tr>
<td>$11</td>
<td>16</td>
</tr>
<tr>
<td>$12</td>
<td>24</td>
</tr>
<tr>
<td>$13</td>
<td>18</td>
</tr>
<tr>
<td>$14</td>
<td>48</td>
</tr>
<tr>
<td>$15</td>
<td>20</td>
</tr>
<tr>
<td>$16</td>
<td>96</td>
</tr>
<tr>
<td>$17</td>
<td>22</td>
</tr>
<tr>
<td>$18</td>
<td>192</td>
</tr>
<tr>
<td>$19</td>
<td>24</td>
</tr>
<tr>
<td>$1A</td>
<td>72</td>
</tr>
<tr>
<td>$1B</td>
<td>26</td>
</tr>
<tr>
<td>$1C</td>
<td>16</td>
</tr>
<tr>
<td>$1D</td>
<td>28</td>
</tr>
<tr>
<td>$1E</td>
<td>32</td>
</tr>
<tr>
<td>$1F</td>
<td>30</td>
</tr>
</tbody>
</table>

#### Envelope {#Envelope}

- See [APU Envelope](../APU_Envelope/)

The [pulse channels](#Pulse) and [noise](#Noise) also contain **envelope generator**, which can automatically decrement the volume as the clock from [frame counter](#Frame_Counter) is fed. Additionally, the decay sequence can be looped via flag.

#### Linear Counter {#Linear_Counter}

- See [APU Triangle#Linear counter control (\$4008, write)](../APU_Triangle/#Linear_counter_control_.28.244008.2C_write.29)

The [triangle](#Triangle) channel additionally contains a **linear counter** unit which is clocked four times per [frame counter](#Frame_Counter) sequence (like the [envelope](#Envelope) of the other channels). It functions independently of the [length counter](#Length_Counter), and will also silence the [triangle](#Triangle) when it reaches zero.

## Conditions for channel output {#Conditions_for_channel_output}

The [pulse](#Pulse), [triangle](#Triangle) and [noise](#Noise) channels will play their corresponding waveforms (at either a constant volume or at a volume controlled by an [envelope](#Envelope) if available) only when (and in the model given here, *precisely* when) their [length counters](#Length_Counter) are all non-zero (this includes the [linear counter](#Linear_Counter) for the [triangle](#Triangle) channel). There are two exceptions for the [pulse channels](#Pulse), which can also be silenced either by having a frequency above a certain threshold, or by a [sweep](../APU_Pulse/) towards lower frequencies (longer periods) reaching the end of the range.

The [DMC](#DMC) channel always outputs the current value of its counter, regardless of the status of the [DMC enable bit](#Status) and whether the sample playback is going or not; the enable bit only controls automatic playback of delta-encoded samples (which is done through counter updates).

## Mixer {#Mixer}

- See [APU Mixer](../APU_Mixer/)

The APU audio output signal comes from two separate components. The [pulse channels](#Pulse) are output on one pin, and the [triangle](#Triangle)/[noise](#Noise)/[DMC](#DMC) are output on another, after which they are combined. Each of these channels has its own nonlinear DAC.

After combining the two output signals, the final signal may go through a lowpass and highpass filter. For instance, RF demodulation in televisions usually results in a strong lowpass. The NES' RCA output circuitry has a more mild lowpass filter.

## Glossary {#Glossary}

- All APU channels have some form of frequency control. The term **frequency** is used where larger register value(s) correspond with higher frequencies, and the term **period** is used where smaller register value(s) correspond with higher frequencies.
- In the block diagrams, a **gate** takes the input on the left and outputs it on the right, unless the control input on top tells the gate to ignore the input and always output 0.
- Some APU units use one or more of the following building blocks:
  - A **divider** outputs a clock periodically. It contains a period reload value, P, and a counter, that starts at P. When the divider is clocked, if the counter is currently 0, it is reloaded with P and generates an output clock, otherwise the counter is decremented. In other words, the divider's period is P + 1.
  - A divider can also be forced to reload its counter immediately (counter = P), but this *does not* output a clock. Similarly, changing a divider's period reload value *does not* affect the counter. Some counters offer no way to force a reload, but [setting P to 0](http://forums.nesdev.org/viewtopic.php?p=186129#p186129) at least synchronizes it to a known state once the current count expires.
  - A divider may be implemented as a down counter (5, 4, 3, …) or as a [linear feedback shift register](https://en.wikipedia.org/wiki/linear_feedback_shift_register) (LFSR). The dividers in the [pulse](#Pulse) and [triangle](#Triangle) channels are linear down-counters. The dividers for [noise](#Noise), [DMC](#DMC), and the [frame counter](#Frame_Counter) are implemented as LFSRs to save gates compared to the equivalent down counter.
  - A **sequencer** continuously loops over a sequence of values or events. When clocked, the next item in the sequence is generated. In this APU documentation, clocking a sequencer usually means either advancing to the next step in a waveform, or the event sequence of the [frame counter](../APU_Frame_Counter/) device.
  - A **timer** is used in each of the five channels to control the sound frequency. It contains a divider which is clocked by the [CPU clock](../Cycle_reference_chart/). The [triangle](#Triangle)'s timer is clocked on every CPU cycle, but the [pulse](#Pulse), [noise](#Noise), and [DMC](#DMC) timers are clocked only on every second CPU cycle and thus produce only even periods.

## References {#References}

- [Blargg's NES APU Reference](http://nesdev.org/apu_ref.txt)
- [Brad Taylor's 2A03 Technical Reference](http://nesdev.org/2A03%20technical%20reference.txt)

<table class="collapsible collapsed" style="width: 100%; border: 1px solid #000; padding: 0">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2" style="background-color: #CCF; font-size: 120%">V | <del>T</del>
APU</th>
</tr>
</thead>
<tbody>
<tr>
<td style="background-color: #DDF"><strong>Concepts</strong></td>
<td><a href="./">Registers</a> | <a href="../APU_basics/">Basics</a> | <a href="../APU_period_table/">Period table</a></td>
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
