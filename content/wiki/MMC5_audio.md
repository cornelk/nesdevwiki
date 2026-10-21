---
title: "MMC5 audio"
url: "/wiki/MMC5_audio/"
description: "Nintendo's MMC5 mapper provides extra sound output, consisting of two pulse wave channels and a PCM channel. The pulse wave channels behave almost identically to the native pulse channels in the NES…"
wikiPage: "MMC5_audio"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
---

Nintendo's [MMC5](../MMC5/) mapper provides extra sound output, consisting of two pulse wave channels and a PCM channel. The pulse wave channels behave almost identically to the native pulse channels in the [NES APU](../APU/).

The sound output of the square channels are equivalent in volume to the corresponding APU channels, but the polarity of all MMC5 channels is reversed compared to the APU.

The PCM channel is similarly equivalent in volume to the APU with equivalent input, and inverted, but using the extra high bit of input it can become twice as loud.

## <span id="Pulse_1_.28.245000-.245003.29"></span>Pulse 1 (\$5000-\$5003) {#Pulse_1_.28.245000-.245003.29}

- Write only
- Reset detection not used

These registers manipulate the MMC5's first [pulse wave](../APU_Pulse/) channel, which functions the same as to those found in the [NES APU](../APU/) except for the following differences:

- \$5001 is not implemented. The MMC5 pulse channels will not sweep, as they have no [sweep unit](../APU_Pulse/).
- Frequency values less than 8 do not silence the MMC5 pulse channels; they can output ultrasonic frequencies.
- Length counter operates twice as fast as the APU length counter (might be clocked at the envelope rate).
- MMC5 does not have an equivalent [frame sequencer](../APU_Frame_Counter/) (APU \$4017); envelope and length counter are fixed to a 240Hz update rate. Based on die images, the exact rate appears to be the M2 frequency divided by the simplified integer 7424, which gives 241.079Hz. However, a bench test has not been performed to confirm this 0.45% difference.

MMC5’s NTSC-only length counter lookup table doesn’t implement the lowest output bit, which is always 0. Accounting for this, it appears to have 1 difference versus the APU’s NTSC lookup table:

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>5-Bit<br />
Register<br />
Value</th>
<th>APU 7-Bit<br />
Count<br />
Value</th>
<th>MMC5 7-Bit<br />
Count<br />
Value</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>$0A</td>
<td>$0A</td>
</tr>
<tr>
<td>$01</td>
<td>$FE</td>
<td>$FE</td>
</tr>
<tr>
<td>$02</td>
<td>$14</td>
<td>$14</td>
</tr>
<tr>
<td>$03</td>
<td>$02</td>
<td>$02</td>
</tr>
<tr>
<td>$04</td>
<td>$28</td>
<td>$28</td>
</tr>
<tr>
<td>$05</td>
<td>$04</td>
<td>$04</td>
</tr>
<tr>
<td>$06</td>
<td>$50</td>
<td>$50</td>
</tr>
<tr>
<td>$07</td>
<td>$06</td>
<td>$06</td>
</tr>
<tr>
<td>$08</td>
<td>$A0</td>
<td>$A0</td>
</tr>
<tr>
<td>$09</td>
<td>$08</td>
<td>$08</td>
</tr>
<tr>
<td>$0A</td>
<td>$3C</td>
<td>$3C</td>
</tr>
<tr>
<td>$0B</td>
<td>$0A</td>
<td>$0A</td>
</tr>
<tr>
<td>$0C</td>
<td>$0E</td>
<td>$0E</td>
</tr>
<tr>
<td>$0D</td>
<td>$0C</td>
<td>$0C</td>
</tr>
<tr>
<td>$0E</td>
<td>$1A</td>
<td>$1A</td>
</tr>
<tr>
<td>$0F</td>
<td>$0E</td>
<td>$0E</td>
</tr>
<tr>
<td>$10</td>
<td>$0C</td>
<td>$0C</td>
</tr>
<tr>
<td>$11</td>
<td>$10</td>
<td>$10</td>
</tr>
<tr>
<td>$12</td>
<td>$18</td>
<td>$18</td>
</tr>
<tr>
<td>$13</td>
<td>$12</td>
<td>$12</td>
</tr>
<tr>
<td>$14</td>
<td>$30</td>
<td>$30</td>
</tr>
<tr>
<td>$15</td>
<td>$14</td>
<td>$14</td>
</tr>
<tr>
<td>$16</td>
<td>$60</td>
<td>$60</td>
</tr>
<tr>
<td>$17</td>
<td>$16</td>
<td>$16</td>
</tr>
<tr>
<td>$18</td>
<td>$C0</td>
<td>$C0</td>
</tr>
<tr>
<td>$19</td>
<td>$18</td>
<td>$18</td>
</tr>
<tr>
<td>$1A</td>
<td>$48</td>
<td>$48</td>
</tr>
<tr>
<td>$1B</td>
<td>$1A</td>
<td>$1A</td>
</tr>
<tr>
<td>$1C</td>
<td>$00</td>
<td style="background: #FFC0C0">$10</td>
</tr>
<tr>
<td>$1D</td>
<td>$1C</td>
<td>$1C</td>
</tr>
<tr>
<td>$1E</td>
<td>$20</td>
<td>$20</td>
</tr>
<tr>
<td>$1F</td>
<td>$1E</td>
<td>$1E</td>
</tr>
</tbody>
</table>

Other features such as the envelope and phase reset are the same as their APU counterparts.

## <span id="Pulse_2_.28.245004-.245007.29"></span>Pulse 2 (\$5004-\$5007) {#Pulse_2_.28.245004-.245007.29}

These registers manipulate the MMC5's second pulse channel.

## <span id="PCM_Mode.2FIRQ_.28.245010.29"></span>PCM Mode/IRQ (\$5010) {#PCM_Mode.2FIRQ_.28.245010.29}

- Read/Write
- Reset detection is used

### Write {#Write}

    7  bit  0
    ---- ----
    0xxx xxx0  Reset Value
    ---- ----
    Ixxx xxxM
    |       |
    |       +- Mode select (0 = write mode. 1 = read mode.)
    +--------- PCM IRQ enable (1 = enabled.)

### Read {#Read}

    7  bit  0
    ---- ----
    Ixxx xxxM  MMC5A default power-on read value = $01
    |       |
    |       +- Revision (0 = Letterless MMC5. 1 = MMC5A.)
    +--------- IRQ (0 = No IRQ triggered. 1 = IRQ was triggered.)

Reading \$5010 acknowledges the IRQ and clears this flag.

## <span id="Raw_PCM_.28.245011.29"></span>Raw PCM (\$5011) {#Raw_PCM_.28.245011.29}

- Write only
- Ignores direct writes in PCM read-mode
- Write-by-read writes to this register in PCM read-mode.
- Reset detection not used
- Power-on state effectively random

### Write {#Write_2}

    7  bit  0
    ---- ----
    WWWW WWWW
    |||| ||||
    ++++-++++- 8-bit PCM data

This functions similarly to the [NES APU](../APU/)'s register \$4011, except that all 8 bits are used.

*Shin 4 Nin Uchi Mahjong* is the only game to use the extra PCM channel (\$5011).

Writing \$00 to this register will have no effect on the output sound, and does not change the PCM counter.

### Pin 2 DAC Characteristic {#Pin_2_DAC_Characteristic}

Pin 2 no-load voltage very closely follows the equation:

Voltage = \[(DAC value / 255) \* (0.4 \* AVcc)\] + (0.1 \* AVcc)

[<img src="../../wiki-images/MMC5_DAC_Characteristic.png" width="847" height="610" />](../../wiki-images/MMC5_DAC_Characteristic.png)

The DAC output is very low impedance -- it holds its voltage true under load. To protect itself, the output also has a current limit. The current limit level depends on AVcc. MMC5A measured current limit in microAmps = (42.178 \* Avcc) - 36.808. It is unknown, but quite doubtful, that prolonged exposure to current limit mode could cause damage to the DAC or other parts of the MMC5. Since current limit mode would cause audible distortion, it should be considered a self-protection feature that should be avoided. In practice, with Avcc = 5V, the DAC output pin should drive 15kohms or higher to avoid current limit mode.

The same MMC5A chip, tested with repeated power cycles on different days of the week and different moon phases, etc, has been observed to power on with a DAC voltage equivalent to DAC value \$EF or \$FF. No other power-on values have been observed. This voltage is not affected by a reset detected when M2 stops toggling.

## PCM description {#PCM_description}

The MMC5's DAC is changed either by writing a value to \$5011 (in write mode) or reading a value from \$8000-BFFF (in read mode). If you try to assign a value of \$00, the DAC is not changed; an IRQ is generated instead and must be acknowledged by reading \$5010.

This functionality can be used in an unrolled loop to stream 8-bit PCM from ROM and terminate at \$00:

    ; write mode
    WriteModePCM:
      lda PCMData ; +4 cycles
      sta $5011 ; +4 cycles = 8 cycles per DAC update
      lda PCMData+1
      sta $5011
      ; ... until a $00 write fires an IRQ

    ; read mode
    ; appears to be intended for making the unrolled loop tighter (thus improving playback rate)
    ReadModePCM:
      lda PCMData ; 4 cycles per DAC update
      lda PCMData+1
      ; ... until a $00 read fires an IRQ

    ; must be within $8000-$BFFF for read mode
    ; should be in a switchable bank for code reuse
    PCMData:
     .byte $AA, $55
     ; ...
     .byte $00

### IRQ operation {#IRQ_operation}

(pseudocode)

    (On DAC write)
        if(value=0)
            irqTrip=1
        else
            irqTrip=0

    (On $5010 write)
        irqEnable=value.bit7

    (On $5010 read)
        value.bit7=(irqTrip AND irqEnable)
        irqTrip=0

    Cart IRQ line=(irqTrip AND irqEnable)

## <span id="Status_.28.245015.29"></span>Status (\$5015) {#Status_.28.245015.29}

- Read/Write
- Reset detection not used

This register is analogous to the [APU Status](../APU/) register found within the NES at \$4015, except that only the bottom 2 bits are used; being for the MMC5's two pulse channels.

## Hardware {#Hardware}

Expansion audio hardware configuration from HVC-ETROM-02:

                R3 15k   C4 100n
    From 2A03 ---/\/\/-----||------+           R4 10k
    (Cart.45)                      |   +--------/\/\/---------+
                                   |   |           _          |
                R2 15k   C3 100n   |   | (MMC5.1) | \         +--- To RF
    MMC5 DAC ----/\/\/-----||------+---+----------|-  \       |    (Cart.46)
    (MMC5.2)                       |              |    >------+
                                   |   intvref ---|+  / (MMC5.100)
                R1 15k   C2 100n   |              |_/
    MMC5 Pulse --/\/\/-----||------+
    (MMC5.3)

## References {#References}

- [Famicom expansion hardware recordings](https://forums.nesdev.org/viewtopic.php?p=90245#p90245) - forum thread with MMC5 test ROMs and recordings confirming various MMC5 audio features.
- [MMC5 volume test](https://forums.nesdev.org/viewtopic.php?p=124288#p124288) - hotswap test ROM for investigating MMC5 volume.
