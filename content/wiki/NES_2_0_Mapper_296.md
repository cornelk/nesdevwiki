---
title: "NES 2.0 Mapper 296"
url: "/wiki/NES_2.0_Mapper_296/"
description: "NES 2.0 Mapper is used for consoles using VT3x hardware, which is mostly compatible hardware to the earlier VT0x and VT1x consoles, but adds an outer bank register and extended support for legacy…"
wikiPage: "NES_2_0_Mapper_296"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
  - name: "VTxx mappers"
    url: "/category/VTxx_mappers/"
---

NES 2.0 Mapper is used for consoles using VT3x hardware, which is mostly compatible hardware to the earlier [VT0x and VT1x](../VTxx/) consoles, but adds an outer bank register and extended support for legacy mappers. While the original OneBus hardware only has MMC3-compatible registers, this type can directly emulate CNROM, UNROM and MMC1 behavior as well.

A non-exhaustive list of notable consoles using this mapper:

- *FC Pocket RS-20*
- *dreamGEAR My Arcade Gamer V*
- *Retro Station NX-85*

# Banking registers {#Banking_registers}

## <span id="Mapper_Configuration_Register_.28.24411D.2C_write.29"></span>Mapper Configuration Register (\$411D, write) {#Mapper_Configuration_Register_.28.24411D.2C_write.29}

    7654 3210
    ---------
    .... .CMM
          |++- Select legacy mapper to emulate in the $8000-$FFFF range
          |     0: MMC3
          |     1: MMC1
          |     2: UNROM
          |     3: CNROM
          +--- Select CHR memory type
                0: CHR-ROM (banked)
                1: CHR-RAM (unbanked)

## <span id="Encryption_Control_.28.24411E.2C_write.29"></span>Encryption Control (\$411E, write) {#Encryption_Control_.28.24411E.2C_write.29}

    7654 3210
    ---------
    cc.. .C.C
    ||    +-+- 00: Disable CPU opcode encryption after the next JMP instruction
    ||         11: Enable CPU opcode encryption after the next JMP instruction
    ++-------- 00: Disable CHR data encryption
               11: Enable CHR data encryption

- When CPU encryption is active, CPU opcode bytes only are XORed with \$A1.
- When CHR data encryption is active, CHR data bytes' bits 0 and 4, 1 and 2, 3 and 7, and 5 and 6 are swapped.

## <span id="Outer_Bank_Register_.28.24412C.29"></span>Outer Bank Register (\$412C) {#Outer_Bank_Register_.28.24412C.29}

    7654 3210
    ---------
    .... cpCP
         |||+- PRG A25 (32 MiB)
         ||+-- CHR A25 (32 MiB)
         |+--- PRG A26 (64 MiB)
         +---- CHR A26 (64 MiB)

## <span id="Outer_Bank_Register_2_.28.24412E.29"></span>Outer Bank Register 2 (\$412E) {#Outer_Bank_Register_2_.28.24412E.29}

    7654 3210
    ---------
    .... ...B
            +- PRG/CHR A27 (128 MiB)

# Two-channel PCM expansion sound {#Two-channel_PCM_expansion_sound}

Expansion sound involves two simultaneous linear 8 bit unsigned PCM channels which just keep playing from their current address until byte \$FF is encountered. Note that the address (excluding bits 0-5) is programmed directly and not relative to other bankswitching registers except the 32 MiB outer bank. The playback rate is set using the normal [\$4010 register](../APU_DMC/) bits 0-3.

## <span id="Select_channel_0_and_set_volume_.28.244031.2C_write.29.2C_Select_channel_1_and_set_volume_.28.244032.2C_write.29"></span>Select channel 0 and set volume (\$4031, write), Select channel 1 and set volume (\$4032, write) {#Select_channel_0_and_set_volume_.28.244031.2C_write.29.2C_Select_channel_1_and_set_volume_.28.244032.2C_write.29}

Selects the channel to which the next writes to \$4012, \$4035 and \$4036 will apply, and sets its linear volume (0-255).

## <span id="Expansion_sound_and_channel_enable_.28.244033.2C_write.29"></span>Expansion sound and channel enable (\$4033, write) {#Expansion_sound_and_channel_enable_.28.244033.2C_write.29}

    7654 3210
    ---------
    E..0 1...
    |  | +---- 1=Enable channel 1
    |  +------ 1=Enable channel 0
    +--------- 1=Enable expansion sound

Note that the order of the channel bits is reversed from \$4014 (see below).

## <span id="Select_playback_address_bits_6-13_.28.244012.29.2C_bits_14-20_.28.244035.29.2C_bits_21-28_.28.244036.29"></span>Select playback address bits 6-13 (\$4012), bits 14-20 (\$4035), bits 21-28 (\$4036) {#Select_playback_address_bits_6-13_.28.244012.29.2C_bits_14-20_.28.244035.29.2C_bits_21-28_.28.244036.29}

These registers can only be accessed if expansion sound is enabled (via \$4033 bit 7, otherwise it \$4012 will go to the regular DPCM channel register), and if the selected channel has been *disabled* via the respective bit in \$4033, otherwise the write will be ignored.

## <span id="Channel_status_.28.244014.2C_read.29"></span>Channel status (\$4014, read) {#Channel_status_.28.244014.2C_read.29}

    7654 3210
    ---------
    .... ..10
           |+- 1=Channel 0 still playing
           +-- 1=Channel 1 still playing

These bits can be used to detect when a sample has finished playing (having encountered byte \$FF) in order to restart it. Once a sample has finished playing, these bits return zero, and the channels are no longer enabled. Note that the *FC Pocket RS-20* menu apparently confuses the two channel bits and continuously writes to restart the background music, which however does not actually occur because the hardware blocks accesss to the playback address registers while the sample is still playing.

# Palette {#Palette}

While the original [VT03+ Enhanced Palette](../VT03+_Enhanced_Palette/) uses four bits each for hue, saturation and luminance, the VT3x 12-bit palette simply directly encodes 4-bit RGB values:

    BA98 7654 3210
    --------------
    BBBB GGGG RRRR
    |||| |||| ++++- Red level (0-15)
    |||| ++++------ Green level (0-15)
    ++++----------- Blue level (0-15)

The *dreamGEAR My Arcade Gamer V* ROM contains a few games that were previously published on VT03 ROMs. When converting these games to VT3x, the developers apparently forgot to change a few colors to the new format, resulting in some elements using strange color combinations even on real hardware (e.g. the waterfall in Curly Monkey).
