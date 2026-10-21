---
title: "VT02+ Video Modes"
url: "/wiki/VT02+_Video_Modes/"
description: "The concept of a video mode is rarely applied to the original NES/Famicom, as it would be considered to only have a single one. Things are more complex on the VTxx series: the VT02 has two, the VT03…"
wikiPage: "VT02__Video_Modes"
---

The concept of a video mode is rarely applied to the original NES/Famicom, as it would be considered to only have a single one. Things are more complex on the VTxx series: the VT02 has two, the VT03 has four, and VT16 and onward have six video modes defined by the number of bits per pixel (bpp), whether Address Extension is used, and whether the Video Data Bus has eight or sixteen bits. The following table shows which video modes are available starting on which console

<table class="wikitable">
<thead>
<tr>
<th>Bits per pixel</th>
<th>Address Extension</th>
<th colspan="2">Video Data Bus Width</th>
</tr>
</thead>
<tbody>
<tr>
<th>8 bit</th>
<th>16 bit</th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td rowspan="2">2</td>
<td>off</td>
<td>RP2C02+</td>
<td>-</td>
</tr>
<tr>
<td>on</td>
<td>VT02+</td>
<td>-</td>
</tr>
<tr>
<td rowspan="2">4</td>
<td>off</td>
<td>VT03</td>
<td>VT16+</td>
</tr>
<tr>
<td>on</td>
<td>VT03+</td>
<td>VT16+</td>
</tr>
</tbody>
</table>

## Bits per pixel {#Bits_per_pixel}

The number of bits per pixel (bpp) is selected separately for background and sprites.

For the background, the number of bits per pixel is selected by register \$2010 bit 1 (BK16EN), with a cleared bit indicating 2bpp, and a set bit indicating 4bpp. The two additional bits are always used as the upper two bits of a four bit color number.

For sprites, the number of bits per pixel is selected by register \$2010 bit 2 (SP16EN), with a cleared bit indicating 2bpp, and a set bit indicating 4bpp. In 4bpp mode, it can further be selected whether the additional two bits are to be used as the upper two bits of a four bit color number, or whether they are used to increase the number of horizontal pixels per sprite from eight to sixteen. In sixteen pixels mode, the lower two bits are used for the left eight-pixel half of the sprite, while the upper two bits are used for the right eight-pixel half of the sprite.

On the VT03, the choice between sixteen colors per sprite --- register \$2010 bit 0 (PIX16EN) cleared --- and sixteen pixels per sprite --- register \$2010 bit 0 (PIX16EN) set --- applies to all sprites. On VT16 and onward, the same applies if register \$2010 bit 5 (SPOPEN) is cleared; if register \$2010 bit 2 (SP16EN) and bit 5 (SPOPEN) are both set, the choice can be made for each sprite individually by clearing (sixteen colors) or setting (sixteen pixels) bit 4 of [byte 2](../PPU_OAM/#Byte_2) of that sprite's OAM data.

## Video data bus width {#Video_data_bus_width}

The original NES/Famicom had an eight-bit data bus for PPU accesses to CHR pattern data. So does the VT02, the VT03 even in 4bpp modes, and VT16 and onward if the sixteen bit video data bus is disabled.

Since twice the number of data must be fetched in 4bpp graphics modes, the VT03 increases the number of eight-bit fetches per time unit. Together with the fact that 4bpp mode is usually used with OneBus addressing that puts CPU and PPU data into the same memory chip, the speed requirements of that one memory chip increase. For this reason, VT16 and onward (and presumably the VT09) adds a sixteen bit data bus mode in which an alternative [CHR Pattern Data Layout](../VT02+_CHR_Pattern_Data_Layout/) is used that allows two bytes to be fetched in one sixteen-bit fetch, reducing the number of fetches per time unit. Accordingly, V.R. Technology's Application Notes name a maximum ROM access time of 70 ns when using the eight bit data bus mode and 120 ns when using the sixteen bit data bus mode.

The sixteen bit video data bus is enabled by setting register \$2010 bit 6 (V16BEN).

## Address Extension {#Address_Extension}

On the original NES/Famicom, and on VT02 and onward with Address Extension Disabled, both background and sprite tiles are indexed by an eight-bit number taken from the [nametable](../PPU_nametables/) for backgrounds and [byte 1](../PPU_OAM/#Byte_1) of each sprites' OAM data. Address Extension adds three bits to form an eleven bit tile number, which can be seen as performing an implicit bankswitch for each tile. VR Technology's VT03 Demonstration ROM uses Background Address Extension to show off the use of more than 256 [Hanzi](https://en.wikipedia.org/wiki/Chinese_characters) characters on a single page.

Address Extension can be enabled separately for background and sprite pattern data. If it is enabled for either background or sprites, then it is also active during \$2006/\$2007 transfers. Please refer to the [VT02+ CHR-ROM Bankswitching](../VT02+_CHR-ROM_Bankswitching/) article for information on the (complex) way Address Extension affects CHR bankswitching.

### Background Address Extension {#Background_Address_Extension}

Background Address Extension is enabled by setting register \$2010 bit 4 (BKEXTEN). The three bits of the Extended Video Address are derived as follows:

    Bit 210
    -------
        PAA
        |++- Attribute data bits 0 and 1
        +--- If register $2011 bit 0 (EVA12S) =0: register $2018 bit 3 (BKPAGE)
             If register $2011 bit 0 (EVA12S) =1: register $4106 bit 0 (HV)

Since the two bits from the attribute table are now used as part of the tile number, they are forced to zero when forming the final palette index of each pixel. Since this effectively reduces the number of available background colors by a factor of four, background address extension is only used in games using 4bpp modes.

### Sprite Address Extension {#Sprite_Address_Extension}

Sprite Address Extension is enabled by setting register \$2010 bit 3 (SPEXTEN). The three bits of the Extended Video Address are derived as follows:

    Bit 210
    -------
        AAA
        +++- Sprites' PPU OAM byte 2, bits 2-4

Note that on VT16 and onward with register \$2010 bit 2 (SP16EN) and bit 5 (SPOPEN) both set, each sprite's PPU OAM byte 2 bit 4 is also used to choose between sixteen colors and sixteen pixels.

### CPU Access Address Extension {#CPU_Access_Address_Extension}

While data is read or written via \$2007, if either Background or Sprite Address Axtension is active, the three bits of the Extended Video Address are derived as follows:

    Bit 210
    -------
        AAA
        +++- Register $2018 bits 0-2 (VRWB)
