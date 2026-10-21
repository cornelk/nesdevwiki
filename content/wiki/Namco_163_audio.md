---
title: "Namco 163 audio"
url: "/wiki/Namco_163_audio/"
description: "The Namco 163 offers up to 8 additional sound channels that play wavetable samples while the 175 and 340 do not. Each waveform can be of a configurable length, and each channel has linear volume…"
wikiPage: "Namco_163_audio"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
---

The [Namco 163](../INES_Mapper_019/) offers up to 8 additional sound channels that play wavetable samples while the [175 and 340](../INES_Mapper_210/) do not. Each waveform can be of a configurable length, and each channel has linear volume control. It has \$80 bytes of sound RAM shared by channel registers and wavetable samples; at least \$40 bytes are dedicated to samples, with more available if not all channels are used.

The chip is unable to clock every channel at once, so it cycles though channels, updating one every 15 CPU cycles. Because of this, the chip allows the game to configure the number of enabled channels. When fewer channels are enabled, the channels are clocked more often, allowing for higher tones with longer, more detailed waveforms. When more channels are enabled, clocking slows down since each channel has to wait its turn, resulting in lower tones and a high-pitched whining noise at the switching frequency. Most games using this IC used only 4 channels.

## Registers {#Registers}

### <span id="Sound_Enable_.28.24E000-E7FF.29"></span>Sound Enable (\$E000-E7FF) {#Sound_Enable_.28.24E000-E7FF.29}

    7  bit  0
    ---- ----
    .SPP PPPP
     ||| ||||
     |++-++++- Select 8KB page of PRG-ROM at $8000-$9FFF
     +-------- Disables sound if set.

Sound is enabled on the 163 by writing a clear bit 6 to this register (0 is sufficient).

### <span id="Address_Port_.28.24F800-.24FFFF.29"></span>Address Port (\$F800-\$FFFF) {#Address_Port_.28.24F800-.24FFFF.29}

    7  bit  0   (write only)
    ---- ----
    IAAA AAAA
    |||| ||||
    |+++-++++- Address
    +--------- Auto-increment

Writing to this register sets the internal address to write to or read from. If the 'I' bit is set, the address will increment on writes and reads to the Data Port (\$4800). Despite previous statements, it does **not** wrap, instead stopping at \$7F. <sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

### <span id="Data_Port_.28.244800-.244FFF.29"></span>Data Port (\$4800-\$4FFF) {#Data_Port_.28.244800-.244FFF.29}

    7  bit  0   (read / write)
    ---- ----
    DDDD DDDD
    |||| ||||
    ++++-++++- Data

This Port accesses the 163's internal \$80 bytes of sound RAM. Which of the \$80 bytes is determined by the Address register (\$F800). When read, the appropriate byte is returned. When written, the appropriate byte is stored.

**Caution:** Do not read from sound RAM while auto-increment is enabled and [DMC](../APU_DMC/) is playing. Otherwise, you are likely to miss bytes because of the [DMA multiple-read bug](../DMA/#Register_conflicts). Writing with auto-increment while DMC is playing behaves as expected.

This RAM is used primarily for wavetables. The sound channel control registers are also set by writing to certain addresses in sound RAM:

#### <span id="Sound_RAM_.2478_-_Low_Frequency"></span>Sound RAM \$78 - Low Frequency {#Sound_RAM_.2478_-_Low_Frequency}

    7  bit  0
    ---------
    FFFF FFFF
    |||| ||||
    ++++-++++- Low 8 bits of Frequency

#### <span id="Sound_RAM_.2479_-_Low_Phase"></span>Sound RAM \$79 - Low Phase {#Sound_RAM_.2479_-_Low_Phase}

    7  bit  0
    ---------
    PPPP PPPP
    |||| ||||
    ++++-++++- Low 8 bits of Phase

#### <span id="Sound_RAM_.247A_-_Mid_Frequency"></span>Sound RAM \$7A - Mid Frequency {#Sound_RAM_.247A_-_Mid_Frequency}

    7  bit  0
    ---------
    FFFF FFFF
    |||| ||||
    ++++-++++- Middle 8 bits of Frequency

#### <span id="Sound_RAM_.247B_-_Mid_Phase"></span>Sound RAM \$7B - Mid Phase {#Sound_RAM_.247B_-_Mid_Phase}

    7  bit  0
    ---------
    PPPP PPPP
    |||| ||||
    ++++-++++- Middle 8 bits of Phase

#### <span id="Sound_RAM_.247C_-_High_Frequency_and_Wave_Length"></span>Sound RAM \$7C - High Frequency and Wave Length {#Sound_RAM_.247C_-_High_Frequency_and_Wave_Length}

    7  bit  0
    ---------
    LLLL LLFF
    |||| ||||
    |||| ||++- High 2 bits of Frequency
    ++++-++--- Length of waveform ((64-L)*4 4-bit samples)

Equivalently, the waveform length = `256 - %LLLLLL00` samples.

The Namco 129 was never released with a game that used audio. However, on the Namco 129 this register instead only encodes length, supporting any length of waveform and not only multiple of 4 samples. In exchange, the high 2 bits of frequency can't be configured and are effectively always 0.

#### <span id="Sound_RAM_.247D_-_High_Phase"></span>Sound RAM \$7D - High Phase {#Sound_RAM_.247D_-_High_Phase}

    7  bit  0
    ---------
    PPPP PPPP
    |||| ||||
    ++++-++++- High 8 bits of Phase

The high byte of the 24-bit phase value directly determines the current sample position of the channel. Every time a channel is updated, the 18-bit frequency value is added to the 24-bit phase accumulator, which is stored in these three registers.

The phase registers may be written to immediately set the phase of the wave. It is good practice to set the frequency to 0 before doing so.

#### <span id="Sound_RAM_.247E_-_Wave_Address"></span>Sound RAM \$7E - Wave Address {#Sound_RAM_.247E_-_Wave_Address}

    7  bit  0
    ---------
    AAAA AAAA
    |||| ||||
    ++++-++++- Address of waveform (in 4-bit samples)

#### <span id="Sound_RAM_.247F_-_Volume"></span>Sound RAM \$7F - Volume {#Sound_RAM_.247F_-_Volume}

    7  bit  0
    ---------
    .CCC VVVV
     ||| ||||
     ||| ++++- Linear Volume
     +++------ Enabled Channels (1+C)

Note 'C' is available on register \$7F ONLY. Those bits have no effect in other registers.

- When C=0, only channel 8 enabled
- When C=1, channels 8 and 7 enabled
- When C=2, channels 8, 7, 6 enabled
- etc...

#### Other Channels {#Other_Channels}

Above Sound RAM register descriptions (\$78-\$7F) are for the 8th channel. The other 7 channels are accessed via the same pattern, but each 8 bytes before the last:

    Channel 8:  $78-$7F
    Channel 7:  $70-$77
    Channel 6:  $68-$6F
    Channel 5:  $60-$67
    Channel 4:  $58-$5F
    Channel 3:  $50-$57
    Channel 2:  $48-$4F
    Channel 1:  $40-$47

Again note that the 'C' bits in the final register is only available at address \$7F.

When channels are disabled, their registers are unused, and can be used for waveform data instead.

## Waveform {#Waveform}

Each enabled channel cycles through its waveform at a rate determined by the 18-bit frequency value 'F'. Each step in the waveform is 4-bits wide, and the number of steps is determined by the 'L' bits ((64-L)\*4). Two samples are stored to a byte, which is little-endian (unlike the Game Boy's wavetable channel).

The 'A' bits dictate where in the internal sound RAM the waveform starts. 'A' is the address in 4-bit samples, therefore a value of \$02 would be the low 4 bits of address \$01. A value of \$03 would be the high 4 bits of address \$01.

For a visual example, assume you have the following sound RAM:

    $00:    00 00 00 A8 DC EE FF FF EF DE AC 58 23 11 00 00
    $10:    10 21 53 00 00 00 00 00 00 00 00 00 00 00 00 00

And assume a channel has an 'A' value of \$06, and an 'L' value of \$38. That channel's waveform would be a sine wave, looking like the following:

    F -       *****
    E -     **     **
    D -    *         *
    C -   *           *
    B -
    A -  *             *
    9 - 
    8 - *               *
    7 - 
    6 -
    5 -                  *             *
    4 -
    3 -                   *           *
    2 -                    *         *
    1 -                     **     **
    0 -                       *****

## Channel Update {#Channel_Update}

Namco's 163 does not internally mix its channels. Instead, each channel is output one at a time. It takes exactly 15 CPU cycles to update and output one channel. When multiple channels are used it will cycle between them. With 6 or fewer channels, the time to update all channels is a rate faster than any audible frequency, and the difference between this serial output and mixing cannot be heard, but for 8 channels it creates a very loud and apparent noise at the update rate. For a Famicom through RF output, this noise is attenuated during demodulation (which performs a lowpass filter), but through A/V output to a TV that does not filter high frequencies, it can be very unpleasant. Only two games used all 8 channels: *King of Kings* and *Erika to Satoru no Yume Bouken*.

<table class="tabular">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>Channels<br />
Enabled</th>
<th>Update Rate<br />
(NTSC)</th>
<th>Update Rate<br />
(PAL)</th>
</tr>
</thead>
<tbody>
<tr>
<td>1</td>
<td>119.318 kHz</td>
<td>110.840 kHz</td>
</tr>
<tr>
<td>2</td>
<td>59.659 kHz</td>
<td>55.420 kHz</td>
</tr>
<tr>
<td>3</td>
<td>39.773 kHz</td>
<td>36.947 kHz</td>
</tr>
<tr>
<td>4</td>
<td>29.830 kHz</td>
<td>27.710 kHz</td>
</tr>
<tr>
<td>5</td>
<td>23.864 kHz</td>
<td>22.168 kHz</td>
</tr>
<tr>
<td>6</td>
<td>19.886 kHz</td>
<td>18.473 kHz</td>
</tr>
<tr>
<td>7</td>
<td>17.045 kHz</td>
<td>15.834 kHz</td>
</tr>
<tr>
<td>8</td>
<td>14.915 kHz</td>
<td>13.855 kHz</td>
</tr>
</tbody>
</table>

The following is a speculative version of a single channel update, occurring every 15 CPU cycles:

    * w[$80] = the 163's internal memory
    * sample(x) = (w[x/2] >> ((x&1)*4)) & $0F
    * phase = (w[$7D] << 16) + (w[$7B] << 8) + w[$79]
    * freq = ((w[$7C] & $03) << 16) + (w[$7A] << 8) + w[$78]
    * length = 256 - (w[$7C] & $FC)
    * offset = w[$7E]
    * volume = w[$7F] & $0F

    phase = (phase + freq) % (length << 16)
    output = (sample(((phase >> 16) + offset) & $FF) - 8) * volume

The ***output*** will be held until the next channel update. The 24-bit *phase* value will be stored back into *w*\[\$7D/\$7B/\$79\].

The sample value is biased by -8, meaning that a waveform value of 8 represents the centre voltage. This means that volume changes have no effect on a sample of 8, will tend negative if \<8 and positive if \>8.

## Frequency {#Frequency}

The wave position is driven by the high 8 bits of a 24-bit accumulator. Every 15 CPU clocks, one channel will add its 18-bit frequency to the accumulator. Because only one channel is updated per tick, the output frequency is thus divided by the number of channels enabled.

    f = wave frequency
    l = wave length
    c = number of channels
    p = 18-bit frequency value
    n = CPU clock rate (≈1789773 Hz)

    f = (n * p) / (15 * 65536 * l * c)

## Mixing {#Mixing}

The relative volume of the IC varies from game to game, unfortunately. The following samples have been recorded using various test programs.

<table class="tabular">
<thead>
<tr>
<th colspan="4">Difference between loudest APU square and loudest N163 square in 1-channel mode (dB)</th>
</tr>
</thead>
<tbody>
<tr>
<th>Game</th>
<th><a href="http://forums.nesdev.org/viewtopic.php?f=2&amp;t=16910#p222765">NewRisingSun</a></th>
<th><a href="http://forums.nesdev.org/viewtopic.php?f=2&amp;t=16910&amp;start=30#p222880">jrlepage</a></th>
<th><a href="http://forums.nesdev.org/viewtopic.php?f=2&amp;t=16910&amp;start=15#p222873">Rainwarrior</a></th>
</tr>
&#10;<tr>
<td>Final Lap</td>
<td>12.7</td>
<td>11.2</td>
<td></td>
</tr>
<tr>
<td>Sangokushi II: Haou no Tairiku</td>
<td>12.9</td>
<td></td>
<td></td>
</tr>
<tr>
<td>Megami Tensei II</td>
<td>13.0</td>
<td>11.9</td>
<td></td>
</tr>
<tr>
<td>Rolling Thunder</td>
<td>16.9</td>
<td>16.0</td>
<td>16.9</td>
</tr>
<tr>
<td>King of Kings</td>
<td>18.0</td>
<td>17.3</td>
<td></td>
</tr>
<tr>
<td>Mappy Kids</td>
<td>18.6</td>
<td></td>
<td></td>
</tr>
<tr>
<td>Erika to Satoru no Yume Bouken</td>
<td>18.8</td>
<td></td>
<td>18.9</td>
</tr>
<tr>
<td>Youkai Douchuu-ki</td>
<td>18.9</td>
<td></td>
<td></td>
</tr>
<tr>
<td>Sangokushi: Chuugen no Hasha</td>
<td>19.5</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>

Based on these measurements, the following submappers were allocated:

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th colspan="3"><a href="../INES_Mapper_019/Submapper_table/">INES Mapper 019 submapper table</a></th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Meaning</th>
<th>Note</th>
</tr>
&#10;<tr>
<td>0</td>
<td>Default</td>
<td>Expansion sound volume unspecified</td>
</tr>
<tr>
<td><em>1</em></td>
<td><em>Deprecated</em></td>
<td>Internal 128b RAM is battery backed, no external PRG-RAM is present.
<p>No expansion sound. (Equivalent to submapper 2 with 0 in PRG-NVRAM field.)</p></td>
</tr>
<tr>
<td>2</td>
<td>No expansion sound</td>
<td></td>
</tr>
<tr>
<td>3</td>
<td>N163 expansion sound: <strong>11.0-13.0 dB</strong> louder than NES APU</td>
<td></td>
</tr>
<tr>
<td>4</td>
<td>N163 expansion sound: <strong>16.0-17.0 dB</strong> louder than NES APU</td>
<td></td>
</tr>
<tr>
<td>5</td>
<td>N163 expansion sound: <strong>18.0-19.5 dB</strong> louder than NES APU</td>
<td></td>
</tr>
</tbody>
</table>

Because the high frequency generated by the channel cycling can be unpleasant, and emulation of high frequency audio can be difficult, it is often preferred to simply sum the channel outputs, and divide the output volume by the number of active channels. For 6 channels or more, where the switching frequency crosses the threshold of audibility, this approximation will become slightly too loud as it fails to compensate for the transferred energy.

## References {#References}

- [Namcot 163 family pinout](../Namcot_163_family_pinout/) - diagram of the chip
- [Namcot 106 Mapper Information](http://nesdev.org/namco106.txt) by Goroh, ZW4, nori
- [NES Music Format Spec](http://kevtris.org/nes/nsfspec.txt) by Kevin Horton, N106 info by Mamiya

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) <https://forums.nesdev.org/viewtopic.php?p=294940#p294940>
