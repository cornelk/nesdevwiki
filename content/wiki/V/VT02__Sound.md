---
title: "VT02+ Sound"
url: "/wiki/VT02+_Sound/"
description: "The VT02+ Famiclone consoles include a second APU, doubling the number of pulse, triangle and noise channels. Furthermore the single PCM/DPCM channel can use DMA for raw PCM as well, and fetch data…"
wikiPage: "VT02__Sound"
---

The [VT02+ Famiclone consoles](../VTxx/) include a second [APU](../APU/), doubling the number of pulse, triangle and noise channels. Furthermore the single PCM/DPCM channel can use DMA for raw PCM as well, and fetch data from the entire CPU address range.

# Address Map {#Address_Map}

    $4000-$4003   Pulse 1
    $4004-$4007   Pulse 2
    $4008-$400B   Triangle 1
    $400C-$400F   Noise 1
    $4010-$4013   DMC
    $4015         Channel enable and length counter status for $4000-$4013
    $4020-$4023   Pulse 3
    $4024-$4027   Pulse 4
    $4028-$402B   Triangle 2
    $402C-$402F   Noise 2
    $4030         DPCM/Raw PCM selection, output control
    $4031         Write 8 bit raw PCM data
    $4035         Channel enable and length counter status for $4020-$402F

# Registers {#Registers}

The meaning of registers \$4020-\$402F resembles that of registers \$4000-\$400F. Note that all V.R. Technology Famiclones reverse the 25% and 50% duty cycles, even as V.R. Technology's official emulator EmuVT does not.

## <span id="DPCM.2FRaw_PCM_selection.2C_output_control_.28.244030.29"></span>DPCM/Raw PCM selection, output control (\$4030) {#DPCM.2FRaw_PCM_selection.2C_output_control_.28.244030.29}

    7654 3210
    ---- ----
    LlMR EeAA
    |||| |||+- APU DMA address bit 14 (inverted)
    |||| ||+-- APU DMA address bit 15 (inverted)
    |||| |+--- 0: Enable output from first APU ($4000-$400F, DPCM or raw PCM written to $4011)
    |||| |     1: Disable output from first APU ($4000-$400F, DPCM or raw PCM written to $4011)
    |||| +---- 0: Disable output from second APU ($4020-$402F, raw PCM written to $4031 or via DMA)
    ||||       1: Enable output from second APU ($4020-$402F, raw PCM written to $4031 or via DMA)
    |||+------ 0: APU DMA uses DPCM
    |||        1: APU DMA uses raw PCM
    ||+------- 0: Maximum DMA transfer length is 4,081 bytes
    ||         1: Maximum DMA transfer length is 4,096 bytes
    |+-------- 0: Enable loudness of first APU ($4000-$400F, DPCM or raw PCM written to $4011)
    |          1: Disable loudness of first APU ($4000-$400F, DPCM or raw PCM written to $4011)
    +--------- 0: Enable loudness of second APU ($4020-$402F, raw PCM written to $4031 or via DMA)
               1: Disable loudness of second APU ($4020-$402F, raw PCM written to $4031 or via DMA)

Note that the meaning of bit 2 is the exact opposite of bit 3. The direction of the bits' effect is chosen so that writing \$00 will result in perfect NES/Famicom-compatible behavior: DMA address A14/A15 set so that transfers are limited to \$C000-\$FFFF, output enabled from the first and disabled from the second APU, DMA uses DPCM, maximum length is 4,081 bytes, and both APUs loud.

Playing raw PCM via DMA uses the same sampling rates as DPCM. Note that when playing raw PCM data via DMA, the output of the second APU must be enabled first.

## <span id="Write_8_bit_raw_PCM_data_.28.244031.29"></span>Write 8 bit raw PCM data (\$4031) {#Write_8_bit_raw_PCM_data_.28.244031.29}

Identical to register \$4011, except that the resolution is eight rather than seven bits.

## <span id="Channel_enable_and_length_counter_status_for_second_APU_channels_.28.244035.29"></span>Channel enable and length counter status for second APU channels (\$4035) {#Channel_enable_and_length_counter_status_for_second_APU_channels_.28.244035.29}

This register is the equivalent of \$4015, only that it applies to the second APU's channels.

Write:

    7654 3210
    ---- ----
    .... NT43
         |||+- Enable Pulse channel 3
         ||+-- Enable Pulse channel 4
         |+--- Enable Triangle channel 2
         +---- Enable Noise channel 2

Read:

    7654 3210
    ---- ----
    .... NT43
         |||+- 1=Pulse channel 3's length counter >0
         ||+-- 1=Pulse channel 4's length counter >0
         |+--- 1=Triangle channel 2's length counter >0
         +---- 1=Noise channel 2's length counter >0

## Links {#Links}

- [Santa Claus](https://forums.nesdev.org/viewtopic.php?p=225447#p225447) NSF rip
- [100 In 1 (101 In 1) Arcade Action II (AT-103)](https://forums.nesdev.org/viewtopic.php?p=242753#p242753) NSF rip
