---
title: "APU Pulse"
url: "/wiki/APU_Pulse/"
description: "The NES APU pulse (square) channels generate a pulse wave with variable pulse width."
wikiPage: "APU_Pulse"
wikiCategories:
  - name: "APU"
    url: "/category/APU/"
---

The **[NES APU](../APU/) pulse (square) channels** generate a pulse wave with variable pulse width.

                  Sweep Unit ---> Timer
                       |            |
                       |            |
                       |            v 
                       |        Sequencer   Length Counter
                       |            |             |
                       |            |             |
                       v            v             v
    Envelope -------> Gate -----> Gate -------> Gate ---> (to mixer)

## Registers {#Registers}

Register [\$4000](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29), [\$4001](#Sweep_control_.28.244001_.2F_.244005.2C_write.29), [\$4002](#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29) and [\$4003](#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29) correspond to the first Pulse channel, [\$4004](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29), [\$4005](#Sweep_control_.28.244001_.2F_.244005.2C_write.29), [\$4006](#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29), [\$4007](#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29) correspond to the second Pulse channel.

### <span id="Pulse_settings_.28.244000_.2F_.244004.2C_write.29"></span>Pulse settings (\$4000 / \$4004, write) {#Pulse_settings_.28.244000_.2F_.244004.2C_write.29}

    7  bit  0
    ---- ----
    DDLC VVVV
    |||| ||||
    |||| ++++- When C is set: channel volume value
    ||||       When C is clear: envelope divider reload value
    |||+------ Constant volume flag (override envelope with V)
    ||+------- Length counter halt flag and envelope loop flag
    ++-------- Pulse width
                   0 = 12.5%
                   1 = 25.0%
                   2 = 50.0%
                   3 = 75.0%

### <span id="Sweep_control_.28.244001_.2F_.244005.2C_write.29"></span>Sweep control (\$4001 / \$4005, write) {#Sweep_control_.28.244001_.2F_.244005.2C_write.29}

    7  bit  0
    ---- ----
    EPPP NSSS
    |||| ||||
    |||| |+++- Bit shift count
    |||| +---- Negate flag
    ||||           0: add to period, sweeping toward lower frequencies
    ||||           1: subtract from period, sweeping toward higher frequencies
    |+++------ Clock rate
    +--------- Sweep enable flag

Writing to this register also sets sweep reload flag (the write does not have immediate effect, see below).

Each of the two pulse channels contain **sweep units**, which can automatically change the period by up or down without needing CPU's help to calculate it.

#### Sweep unit operation {#Sweep_unit_operation}

Whenever the current period changes (either from the period register writes or update from the sweep unit), the **target period** is also calculated immediately, in this way:

1.  The **change amount** is calculated by shifting the current period's bits right by the shift count ([`S` bits](#Sweep_control_.28.244001_.2F_.244005.2C_write.29)). Becomes negative if the negate flag ([`N` bit](#Sweep_control_.28.244001_.2F_.244005.2C_write.29)) is set.
2.  The **target period** is the *sum of the current period and the change amount*, clamped to zero if this sum is negative.
    - The two pulse channels have their adders' carry inputs wired differently, which produces different results when each channel's change amount is made negative:
      - Pulse 1 adds the [ones' complement](https://en.wikipedia.org/wiki/Ones%27_complement) (*−c − 1*). Making 20 negative produces a change amount of −21.
      - Pulse 2 adds the [two's complement](https://en.wikipedia.org/wiki/Two%27s_complement) (*−c*). Making 20 negative produces a change amount of −20.

- If the **target period** is larger than \$7FF, it **silences the channel**.
  - This muting happens regardless of the state of the sweep unit. Meaning, **even if sweep unit isn't enabled and the current period itself does't exceed \$7FF**, the sweep unit mutes the channel if the calculated result overflows \$7FF. Thus to fully disable the sweep unit, a program must additionally turn on the negate flag ([`N` bit](#Sweep_control_.28.244001_.2F_.244005.2C_write.29)), such as by writing \$08. This ensures that the target period is not greater than the current period and therefore not greater than \$7FF.

When the [frame counter](../APU_Frame_Counter/) generates a half-frame clock (at ≈120 (NTSC) or ≈96Hz (PAL)), following occur:

1.  If-
    - the sweep clock counter *is* zero,
    - the sweep is enabled (via [`E` flag](#Sweep_control_.28.244001_.2F_.244005.2C_write.29)),
    - the shift count ([`S` bits](#Sweep_control_.28.244001_.2F_.244005.2C_write.29)) is nonzero
    - *and* the sweep unit is *not* muting the channel, the pulse's current period is updated (set) to the **target period**. (see above)
2.  If the sweep clock counter *is* zero *or* if the sweep reload flag is set, the counter is reloaded to [`P`](#Sweep_control_.28.244001_.2F_.244005.2C_write.29) and the sweep reload flag is cleared. Otherwise, the counter is decremented.

### <span id="Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29"></span>Timer period lower 8 bits (\$4002 / \$4006, write) {#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29}

    7  bit  0
    ---- ----
    TTTT TTTT
    |||| ||||
    ++++-++++- Lower 8 bits of timer period value

### <span id="Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29"></span>Length counter load value and timer period upper 3 bits (\$4003 / \$4007, write) {#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29}

    7  bit  0
    ---- ----
    LLLL LTTT
    |||| ||||
    |||| |+++- Upper 3 bits of timer period value
    ++++-+---- Length counter load value

Writing to this register also resets the sequencer index (i.e., the pulse phase) to the first value, restart [envelope](../APU_Envelope/).

## Sequencer operation {#Sequencer_operation}

The sequencer is clocked by a [timer](../APU/#Glossary) where the period is the current 11-bit period value *plus one*. Given "`t`" as the timer value, this timer is updated every APU cycles (i.e., every second CPU cycles), and counts `t, t-1, ..., 0, t, t-1, ...`, clocking the waveform generator when it goes *from 0 to t*.

Since the period of the timer is *t+1* APU cycles and the sequencer has 8 steps, the period of the waveform is *8\*(t+1)* APU cycles, or equivalently *16\*(t+1)* CPU cycles.

With:

- f<sub>CPU</sub> = the [clock rate](../Cycle_reference_chart/#Clock_rates) of the CPU
- tval = the current period value
- f<sub>pulse</sub> = the frequency of the wave generated by this channel

Resulting:

- f<sub>pulse</sub> = f<sub>CPU</sub>/(16\*(tval+1))
- tval = f<sub>CPU</sub>/(16\*f<sub>pulse</sub>) - 1

**Note:** A period of *t \< 8*, either set explicitly or via a [sweep](#Sweep_control_.28.244001_.2F_.244005.2C_write.29) period update, **silences the corresponding pulse channel**.

#### Duty Cycle Sequences {#Duty_Cycle_Sequences}

| Duty | Output waveform                 |
|------|---------------------------------|
| 0    | `0 1 0 0 0 0 0 0` (12.5%)       |
| 1    | `0 1 1 0 0 0 0 0` (25%)         |
| 2    | `0 1 1 1 1 0 0 0` (50%)         |
| 3    | `1 0 0 1 1 1 1 1` (25% negated) |

Notice that a few [Famiclone](../CPU_variants/) units have swapped APU duty cycles, works as if the meaning of the two [`V` bits](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29) are swapped each other.

#### Implementation details {#Implementation_details}

The reason for the odd output from the sequencer is that the counter is initialized to zero but counts *downward* rather than upward. Thus it reads the sequence lookup table in the order 0, 7, 6, 5, 4, 3, 2, 1.

| Duty | Sequence lookup table | Output waveform                 |
|------|-----------------------|---------------------------------|
| 0    | `0 0 0 0 0 0 0 1`     | `0 1 0 0 0 0 0 0` (12.5%)       |
| 1    | `0 0 0 0 0 0 1 1`     | `0 1 1 0 0 0 0 0` (25%)         |
| 2    | `0 0 0 0 1 1 1 1`     | `0 1 1 1 1 0 0 0` (50%)         |
| 3    | `1 1 1 1 1 1 0 0`     | `1 0 0 1 1 1 1 1` (25% negated) |

## Mixer output {#Mixer_output}

The [mixer](../APU_Mixer/) receives the pulse channel's current [envelope volume](../APU_Envelope/) (lower 4 bits from [\$4000 / \$4004](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29)) except when

- the sequencer output is zero, or
- overflow from the [sweep unit's adder](#Sweep_unit_operation) is silencing the channel, or
- the [length counter](../APU_Length_Counter/) *is* zero, or
- the timer value is less than eight (*t\<8*, noted [above](#Sequencer_operation)).

If any of the above are true, then the pulse channel sends zero (silence) to the mixer.

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
<td>Pulse 1 and Pulse 2 | <a href="../APU_Triangle/">Triangle</a> | <a href="../APU_Noise/">Noise</a> | <a href="../APU_DMC/">DMC</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Components</strong></td>
<td><a href="../APU_Envelope/">Envelope</a> | <a href="./">Sweep</a> | <a href="../APU_Frame_Counter/">Frame Counter</a> | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | <a href="../APU_Mixer/">Mixer</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
