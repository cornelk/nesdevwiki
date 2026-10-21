---
title: "VTxx"
url: "/wiki/VTxx/"
description: "VTxx refers to V.R. Technology's \"NES-on-a-chip\" Famiclones, some of which having greatly extended capabilities. These capabilities are described in separate articles:"
wikiPage: "VTxx"
---

VTxx refers to V.R. Technology's "NES-on-a-chip" Famiclones, some of which having greatly extended capabilities. These capabilities are described in separate articles:

- [VT01 STN Palette](../VT01_STN_Palette/)
- [VT02+ PRG-ROM Bankswitching](../VT02+_PRG-ROM_Bankswitching/)
- [VT02+ Video Modes](../VT02+_Video_Modes/)
- [VT02+ CHR Pattern Data Layout](../VT02+_CHR_Pattern_Data_Layout/)
- [VT02+ CHR-ROM Bankswitching](../VT02+_CHR-ROM_Bankswitching/)
- [VT02+ MMC3 Compatibility Registers](../VT02+_MMC3_Compatibility_Registers/)
- [VT03+ Enhanced Palette](../VT03+_Enhanced_Palette/)
- [VT02+ Sound](../VT02+_Sound/)
- [VT02+ Registers](../VT02+_Registers/)
- ~~VT02+ Cartridge Connector~~

The progression of features as well as backwards-compatibility is contiguous until the VT09. The VT31 and VT368 are the lower- and higher-end enhancements to the VT09, are not compatible with each other, but are backwards-compatible to the VT09 except in the format of the enhanced palette and the absence of the second APU.

Subor uses a VT08 with standard NES functionality, and a similarly named VT06 was spotted in a handheld; its purpose or connection to VRT are unknown and unconfirmed. Jncota uses a VT38, based on a VT33.

The VT1682 and VT268 are not NES-compatible, even as their designs are inspired by it.

# Memory Map {#Memory_Map}

## <span id=".28Main.29_CPU_Memory_Map"></span>(Main) CPU Memory Map {#.28Main.29_CPU_Memory_Map}

    CPU $0000-$07FF: 2 KiB of internal RAM
    CPU $0800-$0FFF: VT01-VT03: Mirror of CPU $0000-$07FF, VT09+: Further 2 KiB of internal RAM
    CPU $1000-$1FFF: VT01-VT32: Mirror of CPU $0000-$0FFF, VT369: Embedded 4 KiB ROM
    CPU $2000-$2007: RP2C02-compatible PPU registers
    CPU $2010-$201F: VT02+: New PPU registers
    CPU $2020-$207F: VT369: Newer PPU registers
    CPU $3000-$3FFF: VT369: Mirror of PPU $2000-$2FFF, allowing direct nametable access
    CPU $4000-$401F: RP2A03-compatible APU, DMA and I/O registers
    CPU $4020-$403F: VT02+: New APU, DMA and I/O registers
    CPU $4100-$41FF: VT02+: New miscellaneous registers
    CPU $4800-$4FFF: VT369: 2 KiB of RAM shared with sound CPU
    CPU $5000-$53FF: VT369: Mirror of PPU $3C00-$3FFF, allowing direct palette access
    CPU $6000-$7FFF: Optional WRAM, if present on cartridge, VT369: Optional ROM bank
    CPU $8000-$FFFF (write): VT02+: When forwarding is enabled (register $410B bit 3 "FWEN" =0):
                             MMC3-compatible registers that are forwarded to $4101-$4108)
    CPU $8000-$FFFF (read): Four 8 KiB PRG-ROM banks.

## <span id="Sound_CPU_Memory_Map_.28VT369.29"></span>Sound CPU Memory Map (VT369) {#Sound_CPU_Memory_Map_.28VT369.29}

    CPU $0000-$07FF: 2 KiB of scratch RAM
    CPU $1800-$1FFF: 2 KiB of RAM shared with main CPU
    CPU $2000-$2FFF: Sound registers
    CPU $4000-$4FFF: Embedded 4 KiB ROM

## PPU Memory Map {#PPU_Memory_Map}

    PPU $0000-$1FFF: CHR-ROM, all modes: CHR pattern data, bit planes 0 and 1
    PPU $2000-$2FFF: NTRAM (2 KiB, selectable mirroring)
    PPU $3C00-$3EFF: VT369: CGRAM in 8bpp modes
    PPU $3F00-$3F1F: CGRAM in 2bpp modes, entries $3F10/$3F14/$3F18/$3F01 mirrors of $3F00/$3F04/$3F08/$3F0C.
                     VT03+: Also CGRAM in 4bpp modes, bits 0 to 5, of colors 00-1F; same mirroring of address bit 4 applies.
    PPU $3F20-$3F7F: VT03+: CGRAM in 4bpp modes, bits 0 to 5, of colors 20-7F; no mirroring of address bit 4.
    PPU $3F80-$3FFF: VT03+: CGRAM in 4bpp modes, bits 6 to 11, of colors 00-7F; no mirroring of address bit 4.
    PPU $4000-$5FFF: CHR-ROM, 4bpp modes: CHR pattern data, bit planes 2 and 3 (VT03+)

# List of VTxx consoles {#List_of_VTxx_consoles}

## VT01 {#VT01}

The VT01 is V.R. Technology's first Famiclone. Its only enhancement is the ability to drive an [STN display](https://en.wikipedia.org/wiki/Super-twisted_nematic_display) directly using a [modified palette](../VT01_STN_Palette/).

## VT02 {#VT02}

The first of the V.R. Technology Famiclones to add major enhancements:

- OneBus mode: PRG and CHR data come from the same address space. The cartridge connector's PPU Address lines are repurposed as higher-order Address lines.
- Integrated [MMC3-compatible](../VT02+_MMC3_Compatibility_Registers/) bankswitching by the console, active when in OneBus mode.
- Ability to use NTRAM as CHR-RAM.
- Choice of PPU A12 and HBLANK as a source for clocking the scanline interrupt.
- Automatic X/Y light gun position determination.
- Integrated RS232 interface.
- Direct NTSC/PAL and 50/60 Hz identification.
- Automatic per-attribute-tile background and per-sprite sprite bankswitching.
- DMA can be used to transfer data to PPU memory (via \$2007) in addition to the normal transferring of data to OAM memory (via \$2004).
- A second APU doubling the number of sound channels.
- Raw PCM output using full eight bits of resolution.
- DMA-driven raw PCM output.
- Data for DMA-driven DPCM/PCM output can be placed anywhere in CPU address space, not just within the \$C000-\$FFFF range.

## VT02S {#VT02S}

- Removes the second APU and the RS232 interface.
- Adds a mechanism to automatically turn off the system after a set amount of time.

## VT03 {#VT03}

Based on the original VT02. The VT03 greatly enhances the graphical capabilities of the original RP2C02:

- Sprites can be 16 pixels wide.
- Graphics can use four bits-per-pixel, so that together with two attribute data bits, sixty-four rather than 16 colors each for background and sprites may be chosen.
- Enhanced color palette with twelve bits rather than six bits per color.

## VT09 {#VT09}

Low-cost replacement for VT03. Knowledge of its functions was reverse engineered using the VT16 and VT18 datasheets.

- Internal CPU RAM is 4 KiB rather than just 2 KiB, removing the need for additional WRAM for many applications.
- 4 bpp pattern data can be arranged in an alternative pattern that allows a single 16-bit fetch to retrieve two planes' data, allowing the use of slower ROM chips.

## VT18 {#VT18}

Revision of the VT09 that offers an additional LCD controller.

## VT31 {#VT31}

Current major enhancement of the VT09 line.

- Hardware multiplier and divider.
- Enhanced palette format changed from the VT03's saturation-luminance-hue-based to a simpler 12-bit (4-4-4) RGB format.
- Second APU is removed.

## VT32 {#VT32}

Popular revision of the VT31 with a new two-channel (AD)PCM engine, allowing playback independently of normal PRG banking.

## VT368 {#VT368}

Final enhancement of the VT09. Strictly this is not part of the VTxx line and instead of the VTxxx line, but it is backwards-compatible with the VTxx line.

- Hardware multiplier and divider.
- 5.37 MHz CPU mode.
- 8 bits-per-plane video modes.
- 30 fps video mode for slow ROM chips.
- 16x 8-bits-per-plane 16x16 sprites per scanline.
- Enhanced palette format changed from the VT03's saturation-luminance-hue-based to a simpler 15-bit (5-5-5) RGB format.
- Additional sound CPU for arbitrarily-encoded multi-channel (AD)PCM playback.
- Second APU is removed.
- Additional *additive* outer bank register called "relative bank" that applies to all CPU/PPU ROM accesses.

## VT369 {#VT369}

Revision of the VT368; the additions that the chip was meant for are unconfirmed. Additions only confirmed to exist on this chip include:

- SD card support.

## VT389 {#VT389}

High-resolution version of the VT369, with a 512x480 video mode.

# References {#References}

- [Publicly-available datasheets](http://www.vrt.com.tw/datasheet.htm) on V.R. Technology's website
- [Archived copy of old V.R. Technology website](https://web.archive.org/web/20140501144450/http://www.vrt.com.tw/) containing VT31 information
