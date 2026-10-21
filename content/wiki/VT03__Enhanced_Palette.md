---
title: "VT03+ Enhanced Palette"
url: "/wiki/VT03+_Enhanced_Palette/"
description: "On the VT03 and later consoles, CGRAM is extended from the original RP2C02's 32 bytes (disregarding mirrored bytes) to 256 bytes."
wikiPage: "VT03__Enhanced_Palette"
---

On the VT03 and later consoles, CGRAM is extended from the [original RP2C02](../PPU_palettes/)'s 32 bytes (disregarding mirrored bytes) to 256 bytes.

Entries \$3F10/\$3F14/\$3F18/\$3F1C are still mirrors of \$3F00/\$3F04/\$3F08/\$3F0C, but all the others are not. Using a 4bpp video mode (bits 0 to 3), together with a bit selecting background or sprite index (bit 4) and two attribute bits (bits 5 and 6) yields a total of 128 palette indices, disregarding mirrored bytes.

## Regular 2C02 colors {#Regular_2C02_colors}

If register \$2010 bit 7 (COLCOMP) is cleared, the RP2C02's standard colors are available, even if 4bpp are used. The greater number of palette indices then allows for greater freedom in combining different colors. This method is used by Jungletac's games. Only the 128 bytes from \$3F00-\$3F7F are used in this case. The color numbers have the same format as on the RP2C02:

    7654 3210
    ---- ----
    ..LL HHHH
      || ++++- Hue number. 0=Bright gray (upper oscillation bound), 
      ||                   1-12=Colored hues (blue->pink->red->yellow->green),
      ||                   13: Dark gray (lower oscillation bound),
      ||                   14-15: black.
      ++------ Combined luminance and saturation (0: darkest, 3: brightest).
               Colors are saturated for levels 0-2 and pastel-like for level 3.

## Extended colors {#Extended_colors}

If register \$2010 bit 7 (COLCOMP) is set, then an extended palette becomes available, where each color has twelve instead of six bits. \$3F00-\$3F7F hold the lower six bits of the color number, \$3F80-\$3FFF hold the upper six bits of the color number. These extended color numbers separate and expand luminance and saturation to four bits each:

    BA98 7654 3210
    ---- ---- ----
    SSSS LLLL HHHH
    |||| |||| ++++- Hue number. Same hues as when COLCOMP=0 (exceptions apply).
    |||| ++++------ Luminance (0: darkest, F: brightest).
    ++++----------- Saturation (0: completely desaturated grays, F: maximum saturation).

## Inverted extended color numbers {#Inverted_extended_color_numbers}

The scheme for the extended color numbers as described in V.R. Technology's datasheet is not applied consistently throughout the entire value range. As saturation increases, more and more colors use an alternative, or inverted, addressing scheme. More precisely, with each increasing level of saturation, the number of luminance levels for which all hues are inverted increases by one, as seen in the following chart:

    Saturation Inverted Luminance Levels   0123456789ABCDEF
    0       -               ................
    1       0                               X...............
    2       0 F                             X..............X
    3       0,1 F                           XX.............X
    4       0,1 E,F                         XX............XX
    5       0,1,2 E,F                       XXX...........XX
    6       0,1,2 D,E,F                     XXX..........XXX
    7       0,1,2,3 D,E,F                   XXXX.........XXX
    8       0,1,2,3 C,D,E,F                 XXXX........XXXX
    9       0,1,2,3,4 C,D,E,F               XXXXX.......XXXX
    A       0,1,2,3,4 B,C,D,E,F             XXXXX......XXXXX
    B       0,1,2,3,4,5 B,C,D,E,F           XXXXXX.....XXXXX
    C       0,1,2,3,4,5 A,B,C,D,E,F         XXXXXX....XXXXXX
    D       0,1,2,3,4,5,6 A,B,C,D,E,F       XXXXXXX...XXXXXX
    E       0,1,2,3,4,5,6 9,A,B,C,D,E,F     XXXXXXX..XXXXXXX
    F       0,1,2,3,4,5,6,7 9,A,B,C,D,E,F   XXXXXXXX.XXXXXXX

It is not known if this is by design or a mistake during chip design, but all VT03 games make use of these inverted color numbers as well as the non-inverted ones. The pattern can be described as follows:

    BOOL inverted = (LuminanceValue< (SaturationValue+1) >>1 || LuminanceValue> 15 -(ChromaValue >>1));

Inversion means that

- the saturation value is replaced with 16 minus the original value;
- the luminance value is replaced with (LuminanceValue -8) AND \$0F;
- the hue value is replaced with:

<!-- -->

    Original Hue   0   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
    Replaced Hue   D   7   8   9   A   B   C   1   2   3   4   5   6   0   E   F

## Palette Chart {#Palette_Chart}

The following table shows the palette that is hard-coded into the EmuVT emulator, which it saves to HSL2RGB.TAB every time it is run:

<table class="wikitable">
<thead>
<tr>
<th colspan="16" style="text-align: center;">Saturation level $0</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$000</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$001</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$002</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$003</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$004</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$005</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$006</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$007</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$008</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$009</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$00A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$00B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$00C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$00D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$00E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$00F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$010</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$011</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$012</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$013</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$014</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$015</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$016</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$017</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$018</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$019</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$01A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$01B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$01C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$01D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$01E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$01F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$020</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$021</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$022</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$023</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$024</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$025</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$026</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$027</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$028</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$029</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$02A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$02B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$02C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$02D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$02E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$02F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$030</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$031</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$032</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$033</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$034</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$035</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$036</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$037</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$038</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$039</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$03F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$040</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$041</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$042</td>
<td style="text-align: center; border: 0px; background-color: #000001; width: 32px; height: 32px; color: #fff;">$043</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$044</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$045</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$046</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$047</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$048</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$049</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$04A</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$04B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$04C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$04D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$04E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$04F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1B1B1B; width: 32px; height: 32px; color: #fff;">$050</td>
<td style="text-align: center; border: 0px; background-color: #1B1A1B; width: 32px; height: 32px; color: #fff;">$051</td>
<td style="text-align: center; border: 0px; background-color: #1A1B17; width: 32px; height: 32px; color: #fff;">$052</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1A; width: 32px; height: 32px; color: #fff;">$053</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1A; width: 32px; height: 32px; color: #fff;">$054</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1C; width: 32px; height: 32px; color: #fff;">$055</td>
<td style="text-align: center; border: 0px; background-color: #1C1B19; width: 32px; height: 32px; color: #fff;">$056</td>
<td style="text-align: center; border: 0px; background-color: #1A1C1B; width: 32px; height: 32px; color: #fff;">$057</td>
<td style="text-align: center; border: 0px; background-color: #1B1B1B; width: 32px; height: 32px; color: #fff;">$058</td>
<td style="text-align: center; border: 0px; background-color: #1A1C16; width: 32px; height: 32px; color: #fff;">$059</td>
<td style="text-align: center; border: 0px; background-color: #1A1B15; width: 32px; height: 32px; color: #fff;">$05A</td>
<td style="text-align: center; border: 0px; background-color: #1A1921; width: 32px; height: 32px; color: #fff;">$05B</td>
<td style="text-align: center; border: 0px; background-color: #1A1B16; width: 32px; height: 32px; color: #fff;">$05C</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1A; width: 32px; height: 32px; color: #fff;">$05D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$05E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$05F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #353630; width: 32px; height: 32px; color: #fff;">$060</td>
<td style="text-align: center; border: 0px; background-color: #373435; width: 32px; height: 32px; color: #fff;">$061</td>
<td style="text-align: center; border: 0px; background-color: #353531; width: 32px; height: 32px; color: #fff;">$062</td>
<td style="text-align: center; border: 0px; background-color: #343433; width: 32px; height: 32px; color: #fff;">$063</td>
<td style="text-align: center; border: 0px; background-color: #343434; width: 32px; height: 32px; color: #fff;">$064</td>
<td style="text-align: center; border: 0px; background-color: #363432; width: 32px; height: 32px; color: #fff;">$065</td>
<td style="text-align: center; border: 0px; background-color: #353535; width: 32px; height: 32px; color: #fff;">$066</td>
<td style="text-align: center; border: 0px; background-color: #333436; width: 32px; height: 32px; color: #fff;">$067</td>
<td style="text-align: center; border: 0px; background-color: #363634; width: 32px; height: 32px; color: #fff;">$068</td>
<td style="text-align: center; border: 0px; background-color: #353536; width: 32px; height: 32px; color: #fff;">$069</td>
<td style="text-align: center; border: 0px; background-color: #353436; width: 32px; height: 32px; color: #fff;">$06A</td>
<td style="text-align: center; border: 0px; background-color: #343434; width: 32px; height: 32px; color: #fff;">$06B</td>
<td style="text-align: center; border: 0px; background-color: #353533; width: 32px; height: 32px; color: #fff;">$06C</td>
<td style="text-align: center; border: 0px; background-color: #353535; width: 32px; height: 32px; color: #fff;">$06D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$06E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$06F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #515151; width: 32px; height: 32px; color: #fff;">$070</td>
<td style="text-align: center; border: 0px; background-color: #524F4F; width: 32px; height: 32px; color: #fff;">$071</td>
<td style="text-align: center; border: 0px; background-color: #514F54; width: 32px; height: 32px; color: #fff;">$072</td>
<td style="text-align: center; border: 0px; background-color: #4F504E; width: 32px; height: 32px; color: #fff;">$073</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$074</td>
<td style="text-align: center; border: 0px; background-color: #50514C; width: 32px; height: 32px; color: #fff;">$075</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$076</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$077</td>
<td style="text-align: center; border: 0px; background-color: #505251; width: 32px; height: 32px; color: #fff;">$078</td>
<td style="text-align: center; border: 0px; background-color: #51524C; width: 32px; height: 32px; color: #fff;">$079</td>
<td style="text-align: center; border: 0px; background-color: #505052; width: 32px; height: 32px; color: #fff;">$07A</td>
<td style="text-align: center; border: 0px; background-color: #51514F; width: 32px; height: 32px; color: #fff;">$07B</td>
<td style="text-align: center; border: 0px; background-color: #51514F; width: 32px; height: 32px; color: #fff;">$07C</td>
<td style="text-align: center; border: 0px; background-color: #515151; width: 32px; height: 32px; color: #fff;">$07D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$07E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$07F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6B6A6E; width: 32px; height: 32px; color: #fff;">$080</td>
<td style="text-align: center; border: 0px; background-color: #6A6B66; width: 32px; height: 32px; color: #fff;">$081</td>
<td style="text-align: center; border: 0px; background-color: #696C64; width: 32px; height: 32px; color: #fff;">$082</td>
<td style="text-align: center; border: 0px; background-color: #696A69; width: 32px; height: 32px; color: #fff;">$083</td>
<td style="text-align: center; border: 0px; background-color: #6A6A68; width: 32px; height: 32px; color: #fff;">$084</td>
<td style="text-align: center; border: 0px; background-color: #686A69; width: 32px; height: 32px; color: #fff;">$085</td>
<td style="text-align: center; border: 0px; background-color: #6A6A6A; width: 32px; height: 32px; color: #fff;">$086</td>
<td style="text-align: center; border: 0px; background-color: #696A6C; width: 32px; height: 32px; color: #fff;">$087</td>
<td style="text-align: center; border: 0px; background-color: #6B6B69; width: 32px; height: 32px; color: #fff;">$088</td>
<td style="text-align: center; border: 0px; background-color: #6B6A6A; width: 32px; height: 32px; color: #fff;">$089</td>
<td style="text-align: center; border: 0px; background-color: #6A6A6A; width: 32px; height: 32px; color: #fff;">$08A</td>
<td style="text-align: center; border: 0px; background-color: #6B6A6C; width: 32px; height: 32px; color: #fff;">$08B</td>
<td style="text-align: center; border: 0px; background-color: #6A6A67; width: 32px; height: 32px; color: #fff;">$08C</td>
<td style="text-align: center; border: 0px; background-color: #6A6B66; width: 32px; height: 32px; color: #fff;">$08D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$08E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$08F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #868684; width: 32px; height: 32px; color: #000;">$090</td>
<td style="text-align: center; border: 0px; background-color: #848685; width: 32px; height: 32px; color: #000;">$091</td>
<td style="text-align: center; border: 0px; background-color: #858585; width: 32px; height: 32px; color: #000;">$092</td>
<td style="text-align: center; border: 0px; background-color: #868586; width: 32px; height: 32px; color: #000;">$093</td>
<td style="text-align: center; border: 0px; background-color: #858585; width: 32px; height: 32px; color: #000;">$094</td>
<td style="text-align: center; border: 0px; background-color: #858786; width: 32px; height: 32px; color: #000;">$095</td>
<td style="text-align: center; border: 0px; background-color: #858585; width: 32px; height: 32px; color: #000;">$096</td>
<td style="text-align: center; border: 0px; background-color: #848685; width: 32px; height: 32px; color: #000;">$097</td>
<td style="text-align: center; border: 0px; background-color: #858585; width: 32px; height: 32px; color: #000;">$098</td>
<td style="text-align: center; border: 0px; background-color: #858687; width: 32px; height: 32px; color: #000;">$099</td>
<td style="text-align: center; border: 0px; background-color: #868588; width: 32px; height: 32px; color: #000;">$09A</td>
<td style="text-align: center; border: 0px; background-color: #858586; width: 32px; height: 32px; color: #000;">$09B</td>
<td style="text-align: center; border: 0px; background-color: #858588; width: 32px; height: 32px; color: #000;">$09C</td>
<td style="text-align: center; border: 0px; background-color: #858583; width: 32px; height: 32px; color: #000;">$09D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$09E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$09F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9EA09D; width: 32px; height: 32px; color: #000;">$0A0</td>
<td style="text-align: center; border: 0px; background-color: #A09E9E; width: 32px; height: 32px; color: #000;">$0A1</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9D; width: 32px; height: 32px; color: #000;">$0A2</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9F; width: 32px; height: 32px; color: #000;">$0A3</td>
<td style="text-align: center; border: 0px; background-color: #9F9F9F; width: 32px; height: 32px; color: #000;">$0A4</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9C; width: 32px; height: 32px; color: #000;">$0A5</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9C; width: 32px; height: 32px; color: #000;">$0A6</td>
<td style="text-align: center; border: 0px; background-color: #9EA09F; width: 32px; height: 32px; color: #000;">$0A7</td>
<td style="text-align: center; border: 0px; background-color: #9F9FA1; width: 32px; height: 32px; color: #000;">$0A8</td>
<td style="text-align: center; border: 0px; background-color: #A09F9A; width: 32px; height: 32px; color: #000;">$0A9</td>
<td style="text-align: center; border: 0px; background-color: #9EA09D; width: 32px; height: 32px; color: #000;">$0AA</td>
<td style="text-align: center; border: 0px; background-color: #9E9F9D; width: 32px; height: 32px; color: #000;">$0AB</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9E; width: 32px; height: 32px; color: #000;">$0AC</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9B; width: 32px; height: 32px; color: #000;">$0AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$0B0</td>
<td style="text-align: center; border: 0px; background-color: #B8B9B5; width: 32px; height: 32px; color: #000;">$0B1</td>
<td style="text-align: center; border: 0px; background-color: #B8B9BA; width: 32px; height: 32px; color: #000;">$0B2</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B7; width: 32px; height: 32px; color: #000;">$0B3</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$0B4</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B6; width: 32px; height: 32px; color: #000;">$0B5</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$0B6</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$0B7</td>
<td style="text-align: center; border: 0px; background-color: #BABABA; width: 32px; height: 32px; color: #000;">$0B8</td>
<td style="text-align: center; border: 0px; background-color: #B9B9BA; width: 32px; height: 32px; color: #000;">$0B9</td>
<td style="text-align: center; border: 0px; background-color: #B8B8BC; width: 32px; height: 32px; color: #000;">$0BA</td>
<td style="text-align: center; border: 0px; background-color: #B9BAB2; width: 32px; height: 32px; color: #000;">$0BB</td>
<td style="text-align: center; border: 0px; background-color: #B8B9BB; width: 32px; height: 32px; color: #000;">$0BC</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$0BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D5D1D5; width: 32px; height: 32px; color: #000;">$0C0</td>
<td style="text-align: center; border: 0px; background-color: #D2D2CF; width: 32px; height: 32px; color: #000;">$0C1</td>
<td style="text-align: center; border: 0px; background-color: #D3D1D2; width: 32px; height: 32px; color: #000;">$0C2</td>
<td style="text-align: center; border: 0px; background-color: #D2D2CE; width: 32px; height: 32px; color: #000;">$0C3</td>
<td style="text-align: center; border: 0px; background-color: #D2D2D2; width: 32px; height: 32px; color: #000;">$0C4</td>
<td style="text-align: center; border: 0px; background-color: #D2D2CF; width: 32px; height: 32px; color: #000;">$0C5</td>
<td style="text-align: center; border: 0px; background-color: #D2D2D2; width: 32px; height: 32px; color: #000;">$0C6</td>
<td style="text-align: center; border: 0px; background-color: #D2D2D0; width: 32px; height: 32px; color: #000;">$0C7</td>
<td style="text-align: center; border: 0px; background-color: #D4D4D1; width: 32px; height: 32px; color: #000;">$0C8</td>
<td style="text-align: center; border: 0px; background-color: #D3D3D0; width: 32px; height: 32px; color: #000;">$0C9</td>
<td style="text-align: center; border: 0px; background-color: #D1D3D3; width: 32px; height: 32px; color: #000;">$0CA</td>
<td style="text-align: center; border: 0px; background-color: #D3D2D4; width: 32px; height: 32px; color: #000;">$0CB</td>
<td style="text-align: center; border: 0px; background-color: #D3D4CE; width: 32px; height: 32px; color: #000;">$0CC</td>
<td style="text-align: center; border: 0px; background-color: #D3D2D5; width: 32px; height: 32px; color: #000;">$0CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ECECEC; width: 32px; height: 32px; color: #000;">$0D0</td>
<td style="text-align: center; border: 0px; background-color: #E9EDEB; width: 32px; height: 32px; color: #000;">$0D1</td>
<td style="text-align: center; border: 0px; background-color: #ECEAED; width: 32px; height: 32px; color: #000;">$0D2</td>
<td style="text-align: center; border: 0px; background-color: #EBEAEA; width: 32px; height: 32px; color: #000;">$0D3</td>
<td style="text-align: center; border: 0px; background-color: #EBEBED; width: 32px; height: 32px; color: #000;">$0D4</td>
<td style="text-align: center; border: 0px; background-color: #EAEAEC; width: 32px; height: 32px; color: #000;">$0D5</td>
<td style="text-align: center; border: 0px; background-color: #EAECE9; width: 32px; height: 32px; color: #000;">$0D6</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$0D7</td>
<td style="text-align: center; border: 0px; background-color: #ECEBEF; width: 32px; height: 32px; color: #000;">$0D8</td>
<td style="text-align: center; border: 0px; background-color: #EBECE9; width: 32px; height: 32px; color: #000;">$0D9</td>
<td style="text-align: center; border: 0px; background-color: #EBEAEF; width: 32px; height: 32px; color: #000;">$0DA</td>
<td style="text-align: center; border: 0px; background-color: #EBEBE8; width: 32px; height: 32px; color: #000;">$0DB</td>
<td style="text-align: center; border: 0px; background-color: #EBECE6; width: 32px; height: 32px; color: #000;">$0DC</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$0DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E0</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFE; width: 32px; height: 32px; color: #000;">$0E1</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFE; width: 32px; height: 32px; color: #000;">$0E2</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFE; width: 32px; height: 32px; color: #000;">$0E3</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E4</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFD; width: 32px; height: 32px; color: #000;">$0E5</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E6</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E7</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E8</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0E9</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0EA</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0EB</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0EC</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F0</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F1</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F2</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F3</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F4</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F5</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F6</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F7</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F8</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0F9</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0FA</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0FB</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0FC</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$0FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$0FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $1</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$100</td>
<td style="text-align: center; border: 0px; background-color: #FE5700; width: 32px; height: 32px; color: #fff;">$101</td>
<td style="text-align: center; border: 0px; background-color: #A88D01; width: 32px; height: 32px; color: #000;">$102</td>
<td style="text-align: center; border: 0px; background-color: #42C100; width: 32px; height: 32px; color: #000;">$103</td>
<td style="text-align: center; border: 0px; background-color: #01F500; width: 32px; height: 32px; color: #000;">$104</td>
<td style="text-align: center; border: 0px; background-color: #00F407; width: 32px; height: 32px; color: #000;">$105</td>
<td style="text-align: center; border: 0px; background-color: #00C4ED; width: 32px; height: 32px; color: #000;">$106</td>
<td style="text-align: center; border: 0px; background-color: #0194FD; width: 32px; height: 32px; color: #fff;">$107</td>
<td style="text-align: center; border: 0px; background-color: #485CFD; width: 32px; height: 32px; color: #fff;">$108</td>
<td style="text-align: center; border: 0px; background-color: #A52DFD; width: 32px; height: 32px; color: #fff;">$109</td>
<td style="text-align: center; border: 0px; background-color: #FF00FE; width: 32px; height: 32px; color: #fff;">$10A</td>
<td style="text-align: center; border: 0px; background-color: #FD00E7; width: 32px; height: 32px; color: #fff;">$10B</td>
<td style="text-align: center; border: 0px; background-color: #FF2C01; width: 32px; height: 32px; color: #fff;">$10C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$10D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$10E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$10F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$110</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$111</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$112</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$113</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$114</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$115</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$116</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$117</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$118</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$119</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$11A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$11B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$11C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$11D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$11E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$11F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$120</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$121</td>
<td style="text-align: center; border: 0px; background-color: #010004; width: 32px; height: 32px; color: #fff;">$122</td>
<td style="text-align: center; border: 0px; background-color: #000003; width: 32px; height: 32px; color: #fff;">$123</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$124</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$125</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$126</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$127</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$128</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$129</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$12A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$12B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$12C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$12D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$12E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$12F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$130</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$131</td>
<td style="text-align: center; border: 0px; background-color: #000106; width: 32px; height: 32px; color: #fff;">$132</td>
<td style="text-align: center; border: 0px; background-color: #00000B; width: 32px; height: 32px; color: #fff;">$133</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$134</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$135</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$136</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$137</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$138</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$139</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$13F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0D0C10; width: 32px; height: 32px; color: #fff;">$140</td>
<td style="text-align: center; border: 0px; background-color: #00020F; width: 32px; height: 32px; color: #fff;">$141</td>
<td style="text-align: center; border: 0px; background-color: #000019; width: 32px; height: 32px; color: #fff;">$142</td>
<td style="text-align: center; border: 0px; background-color: #04001A; width: 32px; height: 32px; color: #fff;">$143</td>
<td style="text-align: center; border: 0px; background-color: #0C0011; width: 32px; height: 32px; color: #fff;">$144</td>
<td style="text-align: center; border: 0px; background-color: #0C0006; width: 32px; height: 32px; color: #fff;">$145</td>
<td style="text-align: center; border: 0px; background-color: #0E0000; width: 32px; height: 32px; color: #fff;">$146</td>
<td style="text-align: center; border: 0px; background-color: #0B0000; width: 32px; height: 32px; color: #fff;">$147</td>
<td style="text-align: center; border: 0px; background-color: #020200; width: 32px; height: 32px; color: #fff;">$148</td>
<td style="text-align: center; border: 0px; background-color: #010600; width: 32px; height: 32px; color: #fff;">$149</td>
<td style="text-align: center; border: 0px; background-color: #000804; width: 32px; height: 32px; color: #fff;">$14A</td>
<td style="text-align: center; border: 0px; background-color: #010800; width: 32px; height: 32px; color: #fff;">$14B</td>
<td style="text-align: center; border: 0px; background-color: #000407; width: 32px; height: 32px; color: #fff;">$14C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$14D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$14E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$14F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #272725; width: 32px; height: 32px; color: #fff;">$150</td>
<td style="text-align: center; border: 0px; background-color: #0D1E2A; width: 32px; height: 32px; color: #fff;">$151</td>
<td style="text-align: center; border: 0px; background-color: #171638; width: 32px; height: 32px; color: #fff;">$152</td>
<td style="text-align: center; border: 0px; background-color: #1D1335; width: 32px; height: 32px; color: #fff;">$153</td>
<td style="text-align: center; border: 0px; background-color: #27102D; width: 32px; height: 32px; color: #fff;">$154</td>
<td style="text-align: center; border: 0px; background-color: #291025; width: 32px; height: 32px; color: #fff;">$155</td>
<td style="text-align: center; border: 0px; background-color: #2A1510; width: 32px; height: 32px; color: #fff;">$156</td>
<td style="text-align: center; border: 0px; background-color: #281809; width: 32px; height: 32px; color: #fff;">$157</td>
<td style="text-align: center; border: 0px; background-color: #1D1E00; width: 32px; height: 32px; color: #fff;">$158</td>
<td style="text-align: center; border: 0px; background-color: #182300; width: 32px; height: 32px; color: #fff;">$159</td>
<td style="text-align: center; border: 0px; background-color: #0F2408; width: 32px; height: 32px; color: #fff;">$15A</td>
<td style="text-align: center; border: 0px; background-color: #0C2410; width: 32px; height: 32px; color: #fff;">$15B</td>
<td style="text-align: center; border: 0px; background-color: #0C1F25; width: 32px; height: 32px; color: #fff;">$15C</td>
<td style="text-align: center; border: 0px; background-color: #0C0C0C; width: 32px; height: 32px; color: #fff;">$15D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$15E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$15F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$160</td>
<td style="text-align: center; border: 0px; background-color: #293A48; width: 32px; height: 32px; color: #fff;">$161</td>
<td style="text-align: center; border: 0px; background-color: #323255; width: 32px; height: 32px; color: #fff;">$162</td>
<td style="text-align: center; border: 0px; background-color: #392F52; width: 32px; height: 32px; color: #fff;">$163</td>
<td style="text-align: center; border: 0px; background-color: #432C49; width: 32px; height: 32px; color: #fff;">$164</td>
<td style="text-align: center; border: 0px; background-color: #452C41; width: 32px; height: 32px; color: #fff;">$165</td>
<td style="text-align: center; border: 0px; background-color: #45302B; width: 32px; height: 32px; color: #fff;">$166</td>
<td style="text-align: center; border: 0px; background-color: #433324; width: 32px; height: 32px; color: #fff;">$167</td>
<td style="text-align: center; border: 0px; background-color: #393A1A; width: 32px; height: 32px; color: #fff;">$168</td>
<td style="text-align: center; border: 0px; background-color: #333E1B; width: 32px; height: 32px; color: #fff;">$169</td>
<td style="text-align: center; border: 0px; background-color: #2A3F23; width: 32px; height: 32px; color: #fff;">$16A</td>
<td style="text-align: center; border: 0px; background-color: #27402C; width: 32px; height: 32px; color: #fff;">$16B</td>
<td style="text-align: center; border: 0px; background-color: #273C41; width: 32px; height: 32px; color: #fff;">$16C</td>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$16D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$16E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$16F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5E60; width: 32px; height: 32px; color: #fff;">$170</td>
<td style="text-align: center; border: 0px; background-color: #435463; width: 32px; height: 32px; color: #fff;">$171</td>
<td style="text-align: center; border: 0px; background-color: #4D4C6E; width: 32px; height: 32px; color: #fff;">$172</td>
<td style="text-align: center; border: 0px; background-color: #53496C; width: 32px; height: 32px; color: #fff;">$173</td>
<td style="text-align: center; border: 0px; background-color: #5E4763; width: 32px; height: 32px; color: #fff;">$174</td>
<td style="text-align: center; border: 0px; background-color: #5F465B; width: 32px; height: 32px; color: #fff;">$175</td>
<td style="text-align: center; border: 0px; background-color: #5F4A45; width: 32px; height: 32px; color: #fff;">$176</td>
<td style="text-align: center; border: 0px; background-color: #5E4E3E; width: 32px; height: 32px; color: #fff;">$177</td>
<td style="text-align: center; border: 0px; background-color: #545535; width: 32px; height: 32px; color: #fff;">$178</td>
<td style="text-align: center; border: 0px; background-color: #4D5733; width: 32px; height: 32px; color: #fff;">$179</td>
<td style="text-align: center; border: 0px; background-color: #435A3F; width: 32px; height: 32px; color: #fff;">$17A</td>
<td style="text-align: center; border: 0px; background-color: #425A49; width: 32px; height: 32px; color: #fff;">$17B</td>
<td style="text-align: center; border: 0px; background-color: #40545B; width: 32px; height: 32px; color: #fff;">$17C</td>
<td style="text-align: center; border: 0px; background-color: #424242; width: 32px; height: 32px; color: #fff;">$17D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$17E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$17F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #777777; width: 32px; height: 32px; color: #fff;">$180</td>
<td style="text-align: center; border: 0px; background-color: #5E6E80; width: 32px; height: 32px; color: #fff;">$181</td>
<td style="text-align: center; border: 0px; background-color: #686887; width: 32px; height: 32px; color: #fff;">$182</td>
<td style="text-align: center; border: 0px; background-color: #6E6488; width: 32px; height: 32px; color: #fff;">$183</td>
<td style="text-align: center; border: 0px; background-color: #79627F; width: 32px; height: 32px; color: #fff;">$184</td>
<td style="text-align: center; border: 0px; background-color: #7A6174; width: 32px; height: 32px; color: #fff;">$185</td>
<td style="text-align: center; border: 0px; background-color: #7A6761; width: 32px; height: 32px; color: #fff;">$186</td>
<td style="text-align: center; border: 0px; background-color: #786859; width: 32px; height: 32px; color: #fff;">$187</td>
<td style="text-align: center; border: 0px; background-color: #6F7050; width: 32px; height: 32px; color: #fff;">$188</td>
<td style="text-align: center; border: 0px; background-color: #68724D; width: 32px; height: 32px; color: #fff;">$189</td>
<td style="text-align: center; border: 0px; background-color: #5E7559; width: 32px; height: 32px; color: #fff;">$18A</td>
<td style="text-align: center; border: 0px; background-color: #5C7464; width: 32px; height: 32px; color: #fff;">$18B</td>
<td style="text-align: center; border: 0px; background-color: #5D7076; width: 32px; height: 32px; color: #fff;">$18C</td>
<td style="text-align: center; border: 0px; background-color: #5D5F5E; width: 32px; height: 32px; color: #fff;">$18D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$18E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$18F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #919491; width: 32px; height: 32px; color: #000;">$190</td>
<td style="text-align: center; border: 0px; background-color: #788899; width: 32px; height: 32px; color: #000;">$191</td>
<td style="text-align: center; border: 0px; background-color: #8182A2; width: 32px; height: 32px; color: #000;">$192</td>
<td style="text-align: center; border: 0px; background-color: #897FA2; width: 32px; height: 32px; color: #000;">$193</td>
<td style="text-align: center; border: 0px; background-color: #937C99; width: 32px; height: 32px; color: #000;">$194</td>
<td style="text-align: center; border: 0px; background-color: #957D8D; width: 32px; height: 32px; color: #000;">$195</td>
<td style="text-align: center; border: 0px; background-color: #947F7A; width: 32px; height: 32px; color: #000;">$196</td>
<td style="text-align: center; border: 0px; background-color: #938373; width: 32px; height: 32px; color: #000;">$197</td>
<td style="text-align: center; border: 0px; background-color: #898A68; width: 32px; height: 32px; color: #000;">$198</td>
<td style="text-align: center; border: 0px; background-color: #828D6B; width: 32px; height: 32px; color: #000;">$199</td>
<td style="text-align: center; border: 0px; background-color: #789071; width: 32px; height: 32px; color: #000;">$19A</td>
<td style="text-align: center; border: 0px; background-color: #77907B; width: 32px; height: 32px; color: #000;">$19B</td>
<td style="text-align: center; border: 0px; background-color: #768B90; width: 32px; height: 32px; color: #000;">$19C</td>
<td style="text-align: center; border: 0px; background-color: #777777; width: 32px; height: 32px; color: #fff;">$19D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$19E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$19F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACACA9; width: 32px; height: 32px; color: #000;">$1A0</td>
<td style="text-align: center; border: 0px; background-color: #92A1B1; width: 32px; height: 32px; color: #000;">$1A1</td>
<td style="text-align: center; border: 0px; background-color: #9C9CBB; width: 32px; height: 32px; color: #000;">$1A2</td>
<td style="text-align: center; border: 0px; background-color: #A298BA; width: 32px; height: 32px; color: #000;">$1A3</td>
<td style="text-align: center; border: 0px; background-color: #AC95B1; width: 32px; height: 32px; color: #000;">$1A4</td>
<td style="text-align: center; border: 0px; background-color: #AD96A8; width: 32px; height: 32px; color: #000;">$1A5</td>
<td style="text-align: center; border: 0px; background-color: #AF9A95; width: 32px; height: 32px; color: #000;">$1A6</td>
<td style="text-align: center; border: 0px; background-color: #AC9C8D; width: 32px; height: 32px; color: #000;">$1A7</td>
<td style="text-align: center; border: 0px; background-color: #A2A383; width: 32px; height: 32px; color: #000;">$1A8</td>
<td style="text-align: center; border: 0px; background-color: #9DA784; width: 32px; height: 32px; color: #000;">$1A9</td>
<td style="text-align: center; border: 0px; background-color: #93AA8F; width: 32px; height: 32px; color: #000;">$1AA</td>
<td style="text-align: center; border: 0px; background-color: #91A997; width: 32px; height: 32px; color: #000;">$1AB</td>
<td style="text-align: center; border: 0px; background-color: #91A4AA; width: 32px; height: 32px; color: #000;">$1AC</td>
<td style="text-align: center; border: 0px; background-color: #929292; width: 32px; height: 32px; color: #000;">$1AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C6C5C9; width: 32px; height: 32px; color: #000;">$1B0</td>
<td style="text-align: center; border: 0px; background-color: #ACBCCC; width: 32px; height: 32px; color: #000;">$1B1</td>
<td style="text-align: center; border: 0px; background-color: #B5B5D5; width: 32px; height: 32px; color: #000;">$1B2</td>
<td style="text-align: center; border: 0px; background-color: #BCB2D5; width: 32px; height: 32px; color: #000;">$1B3</td>
<td style="text-align: center; border: 0px; background-color: #C6AFCB; width: 32px; height: 32px; color: #000;">$1B4</td>
<td style="text-align: center; border: 0px; background-color: #C7B0C4; width: 32px; height: 32px; color: #000;">$1B5</td>
<td style="text-align: center; border: 0px; background-color: #C7B4AE; width: 32px; height: 32px; color: #000;">$1B6</td>
<td style="text-align: center; border: 0px; background-color: #C6B6A7; width: 32px; height: 32px; color: #000;">$1B7</td>
<td style="text-align: center; border: 0px; background-color: #BDBE9E; width: 32px; height: 32px; color: #000;">$1B8</td>
<td style="text-align: center; border: 0px; background-color: #B6C09C; width: 32px; height: 32px; color: #000;">$1B9</td>
<td style="text-align: center; border: 0px; background-color: #ADC3A7; width: 32px; height: 32px; color: #000;">$1BA</td>
<td style="text-align: center; border: 0px; background-color: #ABC3AE; width: 32px; height: 32px; color: #000;">$1BB</td>
<td style="text-align: center; border: 0px; background-color: #ABBEC4; width: 32px; height: 32px; color: #000;">$1BC</td>
<td style="text-align: center; border: 0px; background-color: #ABADAC; width: 32px; height: 32px; color: #000;">$1BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DFDFDD; width: 32px; height: 32px; color: #000;">$1C0</td>
<td style="text-align: center; border: 0px; background-color: #C4D4E4; width: 32px; height: 32px; color: #000;">$1C1</td>
<td style="text-align: center; border: 0px; background-color: #CECFEE; width: 32px; height: 32px; color: #000;">$1C2</td>
<td style="text-align: center; border: 0px; background-color: #D6CBED; width: 32px; height: 32px; color: #000;">$1C3</td>
<td style="text-align: center; border: 0px; background-color: #DDC8E5; width: 32px; height: 32px; color: #000;">$1C4</td>
<td style="text-align: center; border: 0px; background-color: #E1C8DB; width: 32px; height: 32px; color: #000;">$1C5</td>
<td style="text-align: center; border: 0px; background-color: #E1CCC8; width: 32px; height: 32px; color: #000;">$1C6</td>
<td style="text-align: center; border: 0px; background-color: #DFCFC0; width: 32px; height: 32px; color: #000;">$1C7</td>
<td style="text-align: center; border: 0px; background-color: #D7D6B7; width: 32px; height: 32px; color: #000;">$1C8</td>
<td style="text-align: center; border: 0px; background-color: #D0DAB7; width: 32px; height: 32px; color: #000;">$1C9</td>
<td style="text-align: center; border: 0px; background-color: #C6DDC1; width: 32px; height: 32px; color: #000;">$1CA</td>
<td style="text-align: center; border: 0px; background-color: #C5DCC8; width: 32px; height: 32px; color: #000;">$1CB</td>
<td style="text-align: center; border: 0px; background-color: #C4D9DE; width: 32px; height: 32px; color: #000;">$1CC</td>
<td style="text-align: center; border: 0px; background-color: #C4C6CA; width: 32px; height: 32px; color: #000;">$1CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F7F7F7; width: 32px; height: 32px; color: #000;">$1D0</td>
<td style="text-align: center; border: 0px; background-color: #E0EDFD; width: 32px; height: 32px; color: #000;">$1D1</td>
<td style="text-align: center; border: 0px; background-color: #E7E7FE; width: 32px; height: 32px; color: #000;">$1D2</td>
<td style="text-align: center; border: 0px; background-color: #EEE4FD; width: 32px; height: 32px; color: #000;">$1D3</td>
<td style="text-align: center; border: 0px; background-color: #F6E2FD; width: 32px; height: 32px; color: #000;">$1D4</td>
<td style="text-align: center; border: 0px; background-color: #F9E0F3; width: 32px; height: 32px; color: #000;">$1D5</td>
<td style="text-align: center; border: 0px; background-color: #F9E5E3; width: 32px; height: 32px; color: #000;">$1D6</td>
<td style="text-align: center; border: 0px; background-color: #F8E8D9; width: 32px; height: 32px; color: #000;">$1D7</td>
<td style="text-align: center; border: 0px; background-color: #EFF0D0; width: 32px; height: 32px; color: #000;">$1D8</td>
<td style="text-align: center; border: 0px; background-color: #E8F2D0; width: 32px; height: 32px; color: #000;">$1D9</td>
<td style="text-align: center; border: 0px; background-color: #DFF5D9; width: 32px; height: 32px; color: #000;">$1DA</td>
<td style="text-align: center; border: 0px; background-color: #DEF5E4; width: 32px; height: 32px; color: #000;">$1DB</td>
<td style="text-align: center; border: 0px; background-color: #DDF0F6; width: 32px; height: 32px; color: #000;">$1DC</td>
<td style="text-align: center; border: 0px; background-color: #DFDDDE; width: 32px; height: 32px; color: #000;">$1DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1E0</td>
<td style="text-align: center; border: 0px; background-color: #F7FEFE; width: 32px; height: 32px; color: #000;">$1E1</td>
<td style="text-align: center; border: 0px; background-color: #FEFEFE; width: 32px; height: 32px; color: #000;">$1E2</td>
<td style="text-align: center; border: 0px; background-color: #FEFCFF; width: 32px; height: 32px; color: #000;">$1E3</td>
<td style="text-align: center; border: 0px; background-color: #FFFAFF; width: 32px; height: 32px; color: #000;">$1E4</td>
<td style="text-align: center; border: 0px; background-color: #FEF9FF; width: 32px; height: 32px; color: #000;">$1E5</td>
<td style="text-align: center; border: 0px; background-color: #FEFDF9; width: 32px; height: 32px; color: #000;">$1E6</td>
<td style="text-align: center; border: 0px; background-color: #FEFFF1; width: 32px; height: 32px; color: #000;">$1E7</td>
<td style="text-align: center; border: 0px; background-color: #FEFFEE; width: 32px; height: 32px; color: #000;">$1E8</td>
<td style="text-align: center; border: 0px; background-color: #FDFEEB; width: 32px; height: 32px; color: #000;">$1E9</td>
<td style="text-align: center; border: 0px; background-color: #F8FEF4; width: 32px; height: 32px; color: #000;">$1EA</td>
<td style="text-align: center; border: 0px; background-color: #F6FEFD; width: 32px; height: 32px; color: #000;">$1EB</td>
<td style="text-align: center; border: 0px; background-color: #F4FFFF; width: 32px; height: 32px; color: #000;">$1EC</td>
<td style="text-align: center; border: 0px; background-color: #F7F7F5; width: 32px; height: 32px; color: #000;">$1ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F0</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F1</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F2</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F3</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F4</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F5</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F6</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1F7</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFD; width: 32px; height: 32px; color: #000;">$1F8</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFE; width: 32px; height: 32px; color: #000;">$1F9</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1FA</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1FB</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1FC</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$1FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$1FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $2</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$200</td>
<td style="text-align: center; border: 0px; background-color: #FF5700; width: 32px; height: 32px; color: #fff;">$201</td>
<td style="text-align: center; border: 0px; background-color: #A78C01; width: 32px; height: 32px; color: #000;">$202</td>
<td style="text-align: center; border: 0px; background-color: #45BE01; width: 32px; height: 32px; color: #000;">$203</td>
<td style="text-align: center; border: 0px; background-color: #01F000; width: 32px; height: 32px; color: #000;">$204</td>
<td style="text-align: center; border: 0px; background-color: #00F007; width: 32px; height: 32px; color: #000;">$205</td>
<td style="text-align: center; border: 0px; background-color: #00C2DE; width: 32px; height: 32px; color: #000;">$206</td>
<td style="text-align: center; border: 0px; background-color: #0194FD; width: 32px; height: 32px; color: #fff;">$207</td>
<td style="text-align: center; border: 0px; background-color: #4C57FE; width: 32px; height: 32px; color: #fff;">$208</td>
<td style="text-align: center; border: 0px; background-color: #A32CFD; width: 32px; height: 32px; color: #fff;">$209</td>
<td style="text-align: center; border: 0px; background-color: #FF00FE; width: 32px; height: 32px; color: #fff;">$20A</td>
<td style="text-align: center; border: 0px; background-color: #FF00E2; width: 32px; height: 32px; color: #fff;">$20B</td>
<td style="text-align: center; border: 0px; background-color: #FF2A09; width: 32px; height: 32px; color: #fff;">$20C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$20D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$20E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$20F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$210</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$211</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$212</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$213</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$214</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$215</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$216</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$217</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$218</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$219</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21D</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21E</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$21F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$220</td>
<td style="text-align: center; border: 0px; background-color: #00000F; width: 32px; height: 32px; color: #fff;">$221</td>
<td style="text-align: center; border: 0px; background-color: #00001C; width: 32px; height: 32px; color: #fff;">$222</td>
<td style="text-align: center; border: 0px; background-color: #00001B; width: 32px; height: 32px; color: #fff;">$223</td>
<td style="text-align: center; border: 0px; background-color: #00000D; width: 32px; height: 32px; color: #fff;">$224</td>
<td style="text-align: center; border: 0px; background-color: #060000; width: 32px; height: 32px; color: #fff;">$225</td>
<td style="text-align: center; border: 0px; background-color: #060000; width: 32px; height: 32px; color: #fff;">$226</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$227</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$228</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$229</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$22A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$22B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$22C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$22D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$22E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$22F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$230</td>
<td style="text-align: center; border: 0px; background-color: #010015; width: 32px; height: 32px; color: #fff;">$231</td>
<td style="text-align: center; border: 0px; background-color: #000023; width: 32px; height: 32px; color: #fff;">$232</td>
<td style="text-align: center; border: 0px; background-color: #000020; width: 32px; height: 32px; color: #fff;">$233</td>
<td style="text-align: center; border: 0px; background-color: #030013; width: 32px; height: 32px; color: #fff;">$234</td>
<td style="text-align: center; border: 0px; background-color: #0A0001; width: 32px; height: 32px; color: #fff;">$235</td>
<td style="text-align: center; border: 0px; background-color: #0A0000; width: 32px; height: 32px; color: #fff;">$236</td>
<td style="text-align: center; border: 0px; background-color: #020100; width: 32px; height: 32px; color: #fff;">$237</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$238</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$239</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23A</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23B</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$23F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1B1B1B; width: 32px; height: 32px; color: #fff;">$240</td>
<td style="text-align: center; border: 0px; background-color: #000327; width: 32px; height: 32px; color: #fff;">$241</td>
<td style="text-align: center; border: 0px; background-color: #010033; width: 32px; height: 32px; color: #fff;">$242</td>
<td style="text-align: center; border: 0px; background-color: #070032; width: 32px; height: 32px; color: #fff;">$243</td>
<td style="text-align: center; border: 0px; background-color: #160026; width: 32px; height: 32px; color: #fff;">$244</td>
<td style="text-align: center; border: 0px; background-color: #1B000F; width: 32px; height: 32px; color: #fff;">$245</td>
<td style="text-align: center; border: 0px; background-color: #1C0001; width: 32px; height: 32px; color: #fff;">$246</td>
<td style="text-align: center; border: 0px; background-color: #160002; width: 32px; height: 32px; color: #fff;">$247</td>
<td style="text-align: center; border: 0px; background-color: #070601; width: 32px; height: 32px; color: #fff;">$248</td>
<td style="text-align: center; border: 0px; background-color: #000D00; width: 32px; height: 32px; color: #fff;">$249</td>
<td style="text-align: center; border: 0px; background-color: #001200; width: 32px; height: 32px; color: #fff;">$24A</td>
<td style="text-align: center; border: 0px; background-color: #001201; width: 32px; height: 32px; color: #fff;">$24B</td>
<td style="text-align: center; border: 0px; background-color: #000B11; width: 32px; height: 32px; color: #fff;">$24C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$24D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$24E</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$24F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #333635; width: 32px; height: 32px; color: #fff;">$250</td>
<td style="text-align: center; border: 0px; background-color: #041E3F; width: 32px; height: 32px; color: #fff;">$251</td>
<td style="text-align: center; border: 0px; background-color: #131350; width: 32px; height: 32px; color: #fff;">$252</td>
<td style="text-align: center; border: 0px; background-color: #220C4E; width: 32px; height: 32px; color: #fff;">$253</td>
<td style="text-align: center; border: 0px; background-color: #31093E; width: 32px; height: 32px; color: #fff;">$254</td>
<td style="text-align: center; border: 0px; background-color: #36092A; width: 32px; height: 32px; color: #fff;">$255</td>
<td style="text-align: center; border: 0px; background-color: #37100B; width: 32px; height: 32px; color: #fff;">$256</td>
<td style="text-align: center; border: 0px; background-color: #311700; width: 32px; height: 32px; color: #fff;">$257</td>
<td style="text-align: center; border: 0px; background-color: #222201; width: 32px; height: 32px; color: #fff;">$258</td>
<td style="text-align: center; border: 0px; background-color: #132801; width: 32px; height: 32px; color: #fff;">$259</td>
<td style="text-align: center; border: 0px; background-color: #042E00; width: 32px; height: 32px; color: #fff;">$25A</td>
<td style="text-align: center; border: 0px; background-color: #002B0D; width: 32px; height: 32px; color: #fff;">$25B</td>
<td style="text-align: center; border: 0px; background-color: #00252B; width: 32px; height: 32px; color: #fff;">$25C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$25D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$25E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$25F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #4F5251; width: 32px; height: 32px; color: #fff;">$260</td>
<td style="text-align: center; border: 0px; background-color: #1F3B5D; width: 32px; height: 32px; color: #fff;">$261</td>
<td style="text-align: center; border: 0px; background-color: #2F2F6C; width: 32px; height: 32px; color: #fff;">$262</td>
<td style="text-align: center; border: 0px; background-color: #3E2869; width: 32px; height: 32px; color: #fff;">$263</td>
<td style="text-align: center; border: 0px; background-color: #4C2459; width: 32px; height: 32px; color: #fff;">$264</td>
<td style="text-align: center; border: 0px; background-color: #532448; width: 32px; height: 32px; color: #fff;">$265</td>
<td style="text-align: center; border: 0px; background-color: #532A24; width: 32px; height: 32px; color: #fff;">$266</td>
<td style="text-align: center; border: 0px; background-color: #4C3113; width: 32px; height: 32px; color: #fff;">$267</td>
<td style="text-align: center; border: 0px; background-color: #3C3C00; width: 32px; height: 32px; color: #fff;">$268</td>
<td style="text-align: center; border: 0px; background-color: #2F4202; width: 32px; height: 32px; color: #fff;">$269</td>
<td style="text-align: center; border: 0px; background-color: #1E4912; width: 32px; height: 32px; color: #fff;">$26A</td>
<td style="text-align: center; border: 0px; background-color: #1A4725; width: 32px; height: 32px; color: #fff;">$26B</td>
<td style="text-align: center; border: 0px; background-color: #194047; width: 32px; height: 32px; color: #fff;">$26C</td>
<td style="text-align: center; border: 0px; background-color: #1A1B16; width: 32px; height: 32px; color: #fff;">$26D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$26E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$26F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6A6C6B; width: 32px; height: 32px; color: #fff;">$270</td>
<td style="text-align: center; border: 0px; background-color: #395577; width: 32px; height: 32px; color: #fff;">$271</td>
<td style="text-align: center; border: 0px; background-color: #484B81; width: 32px; height: 32px; color: #fff;">$272</td>
<td style="text-align: center; border: 0px; background-color: #584384; width: 32px; height: 32px; color: #fff;">$273</td>
<td style="text-align: center; border: 0px; background-color: #673E76; width: 32px; height: 32px; color: #fff;">$274</td>
<td style="text-align: center; border: 0px; background-color: #6D3E62; width: 32px; height: 32px; color: #fff;">$275</td>
<td style="text-align: center; border: 0px; background-color: #6D463F; width: 32px; height: 32px; color: #fff;">$276</td>
<td style="text-align: center; border: 0px; background-color: #694C2E; width: 32px; height: 32px; color: #fff;">$277</td>
<td style="text-align: center; border: 0px; background-color: #5A581D; width: 32px; height: 32px; color: #fff;">$278</td>
<td style="text-align: center; border: 0px; background-color: #4C601B; width: 32px; height: 32px; color: #fff;">$279</td>
<td style="text-align: center; border: 0px; background-color: #39642E; width: 32px; height: 32px; color: #fff;">$27A</td>
<td style="text-align: center; border: 0px; background-color: #33623F; width: 32px; height: 32px; color: #fff;">$27B</td>
<td style="text-align: center; border: 0px; background-color: #335C62; width: 32px; height: 32px; color: #fff;">$27C</td>
<td style="text-align: center; border: 0px; background-color: #353533; width: 32px; height: 32px; color: #fff;">$27D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$27E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$27F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #858786; width: 32px; height: 32px; color: #000;">$280</td>
<td style="text-align: center; border: 0px; background-color: #547292; width: 32px; height: 32px; color: #fff;">$281</td>
<td style="text-align: center; border: 0px; background-color: #6466A1; width: 32px; height: 32px; color: #fff;">$282</td>
<td style="text-align: center; border: 0px; background-color: #745DA1; width: 32px; height: 32px; color: #fff;">$283</td>
<td style="text-align: center; border: 0px; background-color: #825991; width: 32px; height: 32px; color: #fff;">$284</td>
<td style="text-align: center; border: 0px; background-color: #895A7E; width: 32px; height: 32px; color: #fff;">$285</td>
<td style="text-align: center; border: 0px; background-color: #89605C; width: 32px; height: 32px; color: #fff;">$286</td>
<td style="text-align: center; border: 0px; background-color: #846749; width: 32px; height: 32px; color: #fff;">$287</td>
<td style="text-align: center; border: 0px; background-color: #747235; width: 32px; height: 32px; color: #fff;">$288</td>
<td style="text-align: center; border: 0px; background-color: #667937; width: 32px; height: 32px; color: #fff;">$289</td>
<td style="text-align: center; border: 0px; background-color: #547F47; width: 32px; height: 32px; color: #fff;">$28A</td>
<td style="text-align: center; border: 0px; background-color: #4F7F59; width: 32px; height: 32px; color: #fff;">$28B</td>
<td style="text-align: center; border: 0px; background-color: #4F787E; width: 32px; height: 32px; color: #fff;">$28C</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$28D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$28E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$28F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9F9F9F; width: 32px; height: 32px; color: #000;">$290</td>
<td style="text-align: center; border: 0px; background-color: #6E8AAA; width: 32px; height: 32px; color: #000;">$291</td>
<td style="text-align: center; border: 0px; background-color: #7D7FBC; width: 32px; height: 32px; color: #000;">$292</td>
<td style="text-align: center; border: 0px; background-color: #8E77BA; width: 32px; height: 32px; color: #000;">$293</td>
<td style="text-align: center; border: 0px; background-color: #9B72AA; width: 32px; height: 32px; color: #000;">$294</td>
<td style="text-align: center; border: 0px; background-color: #A27397; width: 32px; height: 32px; color: #000;">$295</td>
<td style="text-align: center; border: 0px; background-color: #A27B74; width: 32px; height: 32px; color: #000;">$296</td>
<td style="text-align: center; border: 0px; background-color: #9D8062; width: 32px; height: 32px; color: #000;">$297</td>
<td style="text-align: center; border: 0px; background-color: #8A8E51; width: 32px; height: 32px; color: #000;">$298</td>
<td style="text-align: center; border: 0px; background-color: #80934E; width: 32px; height: 32px; color: #000;">$299</td>
<td style="text-align: center; border: 0px; background-color: #6D9863; width: 32px; height: 32px; color: #000;">$29A</td>
<td style="text-align: center; border: 0px; background-color: #699971; width: 32px; height: 32px; color: #000;">$29B</td>
<td style="text-align: center; border: 0px; background-color: #699097; width: 32px; height: 32px; color: #000;">$29C</td>
<td style="text-align: center; border: 0px; background-color: #6A6A68; width: 32px; height: 32px; color: #fff;">$29D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$29E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$29F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$2A0</td>
<td style="text-align: center; border: 0px; background-color: #88A4C6; width: 32px; height: 32px; color: #000;">$2A1</td>
<td style="text-align: center; border: 0px; background-color: #9799D4; width: 32px; height: 32px; color: #000;">$2A2</td>
<td style="text-align: center; border: 0px; background-color: #A792D2; width: 32px; height: 32px; color: #000;">$2A3</td>
<td style="text-align: center; border: 0px; background-color: #B68EC3; width: 32px; height: 32px; color: #000;">$2A4</td>
<td style="text-align: center; border: 0px; background-color: #BA8EB1; width: 32px; height: 32px; color: #000;">$2A5</td>
<td style="text-align: center; border: 0px; background-color: #BD948E; width: 32px; height: 32px; color: #000;">$2A6</td>
<td style="text-align: center; border: 0px; background-color: #B69B7E; width: 32px; height: 32px; color: #000;">$2A7</td>
<td style="text-align: center; border: 0px; background-color: #A9A76E; width: 32px; height: 32px; color: #000;">$2A8</td>
<td style="text-align: center; border: 0px; background-color: #9BAE6C; width: 32px; height: 32px; color: #000;">$2A9</td>
<td style="text-align: center; border: 0px; background-color: #89B37B; width: 32px; height: 32px; color: #000;">$2AA</td>
<td style="text-align: center; border: 0px; background-color: #84B28E; width: 32px; height: 32px; color: #000;">$2AB</td>
<td style="text-align: center; border: 0px; background-color: #83AAAF; width: 32px; height: 32px; color: #000;">$2AC</td>
<td style="text-align: center; border: 0px; background-color: #868686; width: 32px; height: 32px; color: #000;">$2AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D3D3D4; width: 32px; height: 32px; color: #000;">$2B0</td>
<td style="text-align: center; border: 0px; background-color: #A3BEDE; width: 32px; height: 32px; color: #000;">$2B1</td>
<td style="text-align: center; border: 0px; background-color: #B1B3ED; width: 32px; height: 32px; color: #000;">$2B2</td>
<td style="text-align: center; border: 0px; background-color: #C0ADEC; width: 32px; height: 32px; color: #000;">$2B3</td>
<td style="text-align: center; border: 0px; background-color: #CFA6DE; width: 32px; height: 32px; color: #000;">$2B4</td>
<td style="text-align: center; border: 0px; background-color: #D5A8C7; width: 32px; height: 32px; color: #000;">$2B5</td>
<td style="text-align: center; border: 0px; background-color: #D6AFAA; width: 32px; height: 32px; color: #000;">$2B6</td>
<td style="text-align: center; border: 0px; background-color: #CFB496; width: 32px; height: 32px; color: #000;">$2B7</td>
<td style="text-align: center; border: 0px; background-color: #C2C286; width: 32px; height: 32px; color: #000;">$2B8</td>
<td style="text-align: center; border: 0px; background-color: #B4C885; width: 32px; height: 32px; color: #000;">$2B9</td>
<td style="text-align: center; border: 0px; background-color: #A4CD93; width: 32px; height: 32px; color: #000;">$2BA</td>
<td style="text-align: center; border: 0px; background-color: #9DCCA9; width: 32px; height: 32px; color: #000;">$2BB</td>
<td style="text-align: center; border: 0px; background-color: #9EC5CC; width: 32px; height: 32px; color: #000;">$2BC</td>
<td style="text-align: center; border: 0px; background-color: #A0A0A0; width: 32px; height: 32px; color: #000;">$2BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$2C0</td>
<td style="text-align: center; border: 0px; background-color: #BCD7F5; width: 32px; height: 32px; color: #000;">$2C1</td>
<td style="text-align: center; border: 0px; background-color: #CBCBFE; width: 32px; height: 32px; color: #000;">$2C2</td>
<td style="text-align: center; border: 0px; background-color: #DAC5FF; width: 32px; height: 32px; color: #000;">$2C3</td>
<td style="text-align: center; border: 0px; background-color: #E8C1F6; width: 32px; height: 32px; color: #000;">$2C4</td>
<td style="text-align: center; border: 0px; background-color: #EDC1E4; width: 32px; height: 32px; color: #000;">$2C5</td>
<td style="text-align: center; border: 0px; background-color: #EFC8C3; width: 32px; height: 32px; color: #000;">$2C6</td>
<td style="text-align: center; border: 0px; background-color: #E8CDB0; width: 32px; height: 32px; color: #000;">$2C7</td>
<td style="text-align: center; border: 0px; background-color: #DAD9A0; width: 32px; height: 32px; color: #000;">$2C8</td>
<td style="text-align: center; border: 0px; background-color: #CCDFA0; width: 32px; height: 32px; color: #000;">$2C9</td>
<td style="text-align: center; border: 0px; background-color: #BDE5B1; width: 32px; height: 32px; color: #000;">$2CA</td>
<td style="text-align: center; border: 0px; background-color: #B7E4C3; width: 32px; height: 32px; color: #000;">$2CB</td>
<td style="text-align: center; border: 0px; background-color: #B6DDE2; width: 32px; height: 32px; color: #000;">$2CC</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$2CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$2D0</td>
<td style="text-align: center; border: 0px; background-color: #D6EEFE; width: 32px; height: 32px; color: #000;">$2D1</td>
<td style="text-align: center; border: 0px; background-color: #E4E4FF; width: 32px; height: 32px; color: #000;">$2D2</td>
<td style="text-align: center; border: 0px; background-color: #F1DEFE; width: 32px; height: 32px; color: #000;">$2D3</td>
<td style="text-align: center; border: 0px; background-color: #FEDAFE; width: 32px; height: 32px; color: #000;">$2D4</td>
<td style="text-align: center; border: 0px; background-color: #FFDAF9; width: 32px; height: 32px; color: #000;">$2D5</td>
<td style="text-align: center; border: 0px; background-color: #FFE0DD; width: 32px; height: 32px; color: #000;">$2D6</td>
<td style="text-align: center; border: 0px; background-color: #FFE7CD; width: 32px; height: 32px; color: #000;">$2D7</td>
<td style="text-align: center; border: 0px; background-color: #F3F1C0; width: 32px; height: 32px; color: #000;">$2D8</td>
<td style="text-align: center; border: 0px; background-color: #E6F5BE; width: 32px; height: 32px; color: #000;">$2D9</td>
<td style="text-align: center; border: 0px; background-color: #D5FBCD; width: 32px; height: 32px; color: #000;">$2DA</td>
<td style="text-align: center; border: 0px; background-color: #D1FADA; width: 32px; height: 32px; color: #000;">$2DB</td>
<td style="text-align: center; border: 0px; background-color: #CFF4FA; width: 32px; height: 32px; color: #000;">$2DC</td>
<td style="text-align: center; border: 0px; background-color: #D2D2D2; width: 32px; height: 32px; color: #000;">$2DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$2E0</td>
<td style="text-align: center; border: 0px; background-color: #EFFFFE; width: 32px; height: 32px; color: #000;">$2E1</td>
<td style="text-align: center; border: 0px; background-color: #FBFCFE; width: 32px; height: 32px; color: #000;">$2E2</td>
<td style="text-align: center; border: 0px; background-color: #FEF7FE; width: 32px; height: 32px; color: #000;">$2E3</td>
<td style="text-align: center; border: 0px; background-color: #FFF4FE; width: 32px; height: 32px; color: #000;">$2E4</td>
<td style="text-align: center; border: 0px; background-color: #FEF2FF; width: 32px; height: 32px; color: #000;">$2E5</td>
<td style="text-align: center; border: 0px; background-color: #FEF9F5; width: 32px; height: 32px; color: #000;">$2E6</td>
<td style="text-align: center; border: 0px; background-color: #FEFEE6; width: 32px; height: 32px; color: #000;">$2E7</td>
<td style="text-align: center; border: 0px; background-color: #FFFFD9; width: 32px; height: 32px; color: #000;">$2E8</td>
<td style="text-align: center; border: 0px; background-color: #FDFEDA; width: 32px; height: 32px; color: #000;">$2E9</td>
<td style="text-align: center; border: 0px; background-color: #EFFFE5; width: 32px; height: 32px; color: #000;">$2EA</td>
<td style="text-align: center; border: 0px; background-color: #EBFFF6; width: 32px; height: 32px; color: #000;">$2EB</td>
<td style="text-align: center; border: 0px; background-color: #E9FFFD; width: 32px; height: 32px; color: #000;">$2EC</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEA; width: 32px; height: 32px; color: #000;">$2ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$2F0</td>
<td style="text-align: center; border: 0px; background-color: #F53D00; width: 32px; height: 32px; color: #fff;">$2F1</td>
<td style="text-align: center; border: 0px; background-color: #887501; width: 32px; height: 32px; color: #fff;">$2F2</td>
<td style="text-align: center; border: 0px; background-color: #27A400; width: 32px; height: 32px; color: #fff;">$2F3</td>
<td style="text-align: center; border: 0px; background-color: #01D701; width: 32px; height: 32px; color: #fff;">$2F4</td>
<td style="text-align: center; border: 0px; background-color: #00D600; width: 32px; height: 32px; color: #fff;">$2F5</td>
<td style="text-align: center; border: 0px; background-color: #00A4C7; width: 32px; height: 32px; color: #fff;">$2F6</td>
<td style="text-align: center; border: 0px; background-color: #0079FD; width: 32px; height: 32px; color: #fff;">$2F7</td>
<td style="text-align: center; border: 0px; background-color: #2F3DFD; width: 32px; height: 32px; color: #fff;">$2F8</td>
<td style="text-align: center; border: 0px; background-color: #8611FE; width: 32px; height: 32px; color: #fff;">$2F9</td>
<td style="text-align: center; border: 0px; background-color: #FA00FE; width: 32px; height: 32px; color: #fff;">$2FA</td>
<td style="text-align: center; border: 0px; background-color: #FE00D0; width: 32px; height: 32px; color: #fff;">$2FB</td>
<td style="text-align: center; border: 0px; background-color: #FF0E00; width: 32px; height: 32px; color: #fff;">$2FC</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$2FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$2FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $3</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$300</td>
<td style="text-align: center; border: 0px; background-color: #FE5900; width: 32px; height: 32px; color: #000;">$301</td>
<td style="text-align: center; border: 0px; background-color: #A18E01; width: 32px; height: 32px; color: #000;">$302</td>
<td style="text-align: center; border: 0px; background-color: #45BC00; width: 32px; height: 32px; color: #000;">$303</td>
<td style="text-align: center; border: 0px; background-color: #01E800; width: 32px; height: 32px; color: #000;">$304</td>
<td style="text-align: center; border: 0px; background-color: #01E80C; width: 32px; height: 32px; color: #000;">$305</td>
<td style="text-align: center; border: 0px; background-color: #00BAD9; width: 32px; height: 32px; color: #000;">$306</td>
<td style="text-align: center; border: 0px; background-color: #008FFE; width: 32px; height: 32px; color: #fff;">$307</td>
<td style="text-align: center; border: 0px; background-color: #5053FD; width: 32px; height: 32px; color: #fff;">$308</td>
<td style="text-align: center; border: 0px; background-color: #A42AFE; width: 32px; height: 32px; color: #fff;">$309</td>
<td style="text-align: center; border: 0px; background-color: #FF00FE; width: 32px; height: 32px; color: #fff;">$30A</td>
<td style="text-align: center; border: 0px; background-color: #FF01DB; width: 32px; height: 32px; color: #fff;">$30B</td>
<td style="text-align: center; border: 0px; background-color: #FE2D07; width: 32px; height: 32px; color: #fff;">$30C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$30D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$30E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$30F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$310</td>
<td style="text-align: center; border: 0px; background-color: #FE7200; width: 32px; height: 32px; color: #000;">$311</td>
<td style="text-align: center; border: 0px; background-color: #BBA700; width: 32px; height: 32px; color: #000;">$312</td>
<td style="text-align: center; border: 0px; background-color: #5ED500; width: 32px; height: 32px; color: #000;">$313</td>
<td style="text-align: center; border: 0px; background-color: #00FE00; width: 32px; height: 32px; color: #000;">$314</td>
<td style="text-align: center; border: 0px; background-color: #00FE2A; width: 32px; height: 32px; color: #000;">$315</td>
<td style="text-align: center; border: 0px; background-color: #00D3F1; width: 32px; height: 32px; color: #000;">$316</td>
<td style="text-align: center; border: 0px; background-color: #01A8FE; width: 32px; height: 32px; color: #fff;">$317</td>
<td style="text-align: center; border: 0px; background-color: #6072FD; width: 32px; height: 32px; color: #fff;">$318</td>
<td style="text-align: center; border: 0px; background-color: #AF4AFD; width: 32px; height: 32px; color: #fff;">$319</td>
<td style="text-align: center; border: 0px; background-color: #FE18FE; width: 32px; height: 32px; color: #fff;">$31A</td>
<td style="text-align: center; border: 0px; background-color: #FF1AF9; width: 32px; height: 32px; color: #fff;">$31B</td>
<td style="text-align: center; border: 0px; background-color: #FF462B; width: 32px; height: 32px; color: #fff;">$31C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$31D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$31E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$31F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$320</td>
<td style="text-align: center; border: 0px; background-color: #00001E; width: 32px; height: 32px; color: #fff;">$321</td>
<td style="text-align: center; border: 0px; background-color: #000032; width: 32px; height: 32px; color: #fff;">$322</td>
<td style="text-align: center; border: 0px; background-color: #000034; width: 32px; height: 32px; color: #fff;">$323</td>
<td style="text-align: center; border: 0px; background-color: #0B0020; width: 32px; height: 32px; color: #fff;">$324</td>
<td style="text-align: center; border: 0px; background-color: #130004; width: 32px; height: 32px; color: #fff;">$325</td>
<td style="text-align: center; border: 0px; background-color: #130101; width: 32px; height: 32px; color: #fff;">$326</td>
<td style="text-align: center; border: 0px; background-color: #0B0000; width: 32px; height: 32px; color: #fff;">$327</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$328</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$329</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$32A</td>
<td style="text-align: center; border: 0px; background-color: #000400; width: 32px; height: 32px; color: #fff;">$32B</td>
<td style="text-align: center; border: 0px; background-color: #000004; width: 32px; height: 32px; color: #fff;">$32C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$32D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$32E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$32F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0D0D0D; width: 32px; height: 32px; color: #fff;">$330</td>
<td style="text-align: center; border: 0px; background-color: #000022; width: 32px; height: 32px; color: #fff;">$331</td>
<td style="text-align: center; border: 0px; background-color: #00003E; width: 32px; height: 32px; color: #fff;">$332</td>
<td style="text-align: center; border: 0px; background-color: #01003A; width: 32px; height: 32px; color: #fff;">$333</td>
<td style="text-align: center; border: 0px; background-color: #0F0023; width: 32px; height: 32px; color: #fff;">$334</td>
<td style="text-align: center; border: 0px; background-color: #160005; width: 32px; height: 32px; color: #fff;">$335</td>
<td style="text-align: center; border: 0px; background-color: #160100; width: 32px; height: 32px; color: #fff;">$336</td>
<td style="text-align: center; border: 0px; background-color: #0C0000; width: 32px; height: 32px; color: #fff;">$337</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$338</td>
<td style="text-align: center; border: 0px; background-color: #000101; width: 32px; height: 32px; color: #fff;">$339</td>
<td style="text-align: center; border: 0px; background-color: #000800; width: 32px; height: 32px; color: #fff;">$33A</td>
<td style="text-align: center; border: 0px; background-color: #010600; width: 32px; height: 32px; color: #fff;">$33B</td>
<td style="text-align: center; border: 0px; background-color: #000005; width: 32px; height: 32px; color: #fff;">$33C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$33D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$33E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$33F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$340</td>
<td style="text-align: center; border: 0px; background-color: #00053C; width: 32px; height: 32px; color: #fff;">$341</td>
<td style="text-align: center; border: 0px; background-color: #000051; width: 32px; height: 32px; color: #fff;">$342</td>
<td style="text-align: center; border: 0px; background-color: #0B004D; width: 32px; height: 32px; color: #fff;">$343</td>
<td style="text-align: center; border: 0px; background-color: #210039; width: 32px; height: 32px; color: #fff;">$344</td>
<td style="text-align: center; border: 0px; background-color: #2A0017; width: 32px; height: 32px; color: #fff;">$345</td>
<td style="text-align: center; border: 0px; background-color: #2A0001; width: 32px; height: 32px; color: #fff;">$346</td>
<td style="text-align: center; border: 0px; background-color: #200100; width: 32px; height: 32px; color: #fff;">$347</td>
<td style="text-align: center; border: 0px; background-color: #080800; width: 32px; height: 32px; color: #fff;">$348</td>
<td style="text-align: center; border: 0px; background-color: #001400; width: 32px; height: 32px; color: #fff;">$349</td>
<td style="text-align: center; border: 0px; background-color: #001B00; width: 32px; height: 32px; color: #fff;">$34A</td>
<td style="text-align: center; border: 0px; background-color: #001A00; width: 32px; height: 32px; color: #fff;">$34B</td>
<td style="text-align: center; border: 0px; background-color: #001017; width: 32px; height: 32px; color: #fff;">$34C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$34D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$34E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$34F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #424242; width: 32px; height: 32px; color: #fff;">$350</td>
<td style="text-align: center; border: 0px; background-color: #012151; width: 32px; height: 32px; color: #fff;">$351</td>
<td style="text-align: center; border: 0px; background-color: #10116A; width: 32px; height: 32px; color: #fff;">$352</td>
<td style="text-align: center; border: 0px; background-color: #260668; width: 32px; height: 32px; color: #fff;">$353</td>
<td style="text-align: center; border: 0px; background-color: #3B0050; width: 32px; height: 32px; color: #fff;">$354</td>
<td style="text-align: center; border: 0px; background-color: #450134; width: 32px; height: 32px; color: #fff;">$355</td>
<td style="text-align: center; border: 0px; background-color: #450A04; width: 32px; height: 32px; color: #fff;">$356</td>
<td style="text-align: center; border: 0px; background-color: #3B1500; width: 32px; height: 32px; color: #fff;">$357</td>
<td style="text-align: center; border: 0px; background-color: #252501; width: 32px; height: 32px; color: #fff;">$358</td>
<td style="text-align: center; border: 0px; background-color: #112E00; width: 32px; height: 32px; color: #fff;">$359</td>
<td style="text-align: center; border: 0px; background-color: #013600; width: 32px; height: 32px; color: #fff;">$35A</td>
<td style="text-align: center; border: 0px; background-color: #003602; width: 32px; height: 32px; color: #fff;">$35B</td>
<td style="text-align: center; border: 0px; background-color: #002A37; width: 32px; height: 32px; color: #fff;">$35C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$35D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$35E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$35F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$360</td>
<td style="text-align: center; border: 0px; background-color: #153B6F; width: 32px; height: 32px; color: #fff;">$361</td>
<td style="text-align: center; border: 0px; background-color: #2B2D84; width: 32px; height: 32px; color: #fff;">$362</td>
<td style="text-align: center; border: 0px; background-color: #412083; width: 32px; height: 32px; color: #fff;">$363</td>
<td style="text-align: center; border: 0px; background-color: #581A6F; width: 32px; height: 32px; color: #fff;">$364</td>
<td style="text-align: center; border: 0px; background-color: #601B51; width: 32px; height: 32px; color: #fff;">$365</td>
<td style="text-align: center; border: 0px; background-color: #612620; width: 32px; height: 32px; color: #fff;">$366</td>
<td style="text-align: center; border: 0px; background-color: #593004; width: 32px; height: 32px; color: #fff;">$367</td>
<td style="text-align: center; border: 0px; background-color: #424101; width: 32px; height: 32px; color: #fff;">$368</td>
<td style="text-align: center; border: 0px; background-color: #2C4B00; width: 32px; height: 32px; color: #fff;">$369</td>
<td style="text-align: center; border: 0px; background-color: #145202; width: 32px; height: 32px; color: #fff;">$36A</td>
<td style="text-align: center; border: 0px; background-color: #0C521C; width: 32px; height: 32px; color: #fff;">$36B</td>
<td style="text-align: center; border: 0px; background-color: #0B474F; width: 32px; height: 32px; color: #fff;">$36C</td>
<td style="text-align: center; border: 0px; background-color: #0E0B0C; width: 32px; height: 32px; color: #fff;">$36D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$36E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$36F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #777775; width: 32px; height: 32px; color: #fff;">$370</td>
<td style="text-align: center; border: 0px; background-color: #30588C; width: 32px; height: 32px; color: #fff;">$371</td>
<td style="text-align: center; border: 0px; background-color: #4648A0; width: 32px; height: 32px; color: #fff;">$372</td>
<td style="text-align: center; border: 0px; background-color: #5C3D9E; width: 32px; height: 32px; color: #fff;">$373</td>
<td style="text-align: center; border: 0px; background-color: #74368B; width: 32px; height: 32px; color: #fff;">$374</td>
<td style="text-align: center; border: 0px; background-color: #7B366C; width: 32px; height: 32px; color: #fff;">$375</td>
<td style="text-align: center; border: 0px; background-color: #7F4037; width: 32px; height: 32px; color: #fff;">$376</td>
<td style="text-align: center; border: 0px; background-color: #744B1F; width: 32px; height: 32px; color: #fff;">$377</td>
<td style="text-align: center; border: 0px; background-color: #5D5C02; width: 32px; height: 32px; color: #fff;">$378</td>
<td style="text-align: center; border: 0px; background-color: #486602; width: 32px; height: 32px; color: #fff;">$379</td>
<td style="text-align: center; border: 0px; background-color: #2F6C1C; width: 32px; height: 32px; color: #fff;">$37A</td>
<td style="text-align: center; border: 0px; background-color: #286D38; width: 32px; height: 32px; color: #fff;">$37B</td>
<td style="text-align: center; border: 0px; background-color: #27616C; width: 32px; height: 32px; color: #fff;">$37C</td>
<td style="text-align: center; border: 0px; background-color: #272928; width: 32px; height: 32px; color: #fff;">$37D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$37E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$37F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #949293; width: 32px; height: 32px; color: #000;">$380</td>
<td style="text-align: center; border: 0px; background-color: #4B72A6; width: 32px; height: 32px; color: #fff;">$381</td>
<td style="text-align: center; border: 0px; background-color: #6063BC; width: 32px; height: 32px; color: #fff;">$382</td>
<td style="text-align: center; border: 0px; background-color: #7756BB; width: 32px; height: 32px; color: #fff;">$383</td>
<td style="text-align: center; border: 0px; background-color: #8E51A1; width: 32px; height: 32px; color: #fff;">$384</td>
<td style="text-align: center; border: 0px; background-color: #995087; width: 32px; height: 32px; color: #fff;">$385</td>
<td style="text-align: center; border: 0px; background-color: #9A5B52; width: 32px; height: 32px; color: #fff;">$386</td>
<td style="text-align: center; border: 0px; background-color: #8F6437; width: 32px; height: 32px; color: #fff;">$387</td>
<td style="text-align: center; border: 0px; background-color: #79781C; width: 32px; height: 32px; color: #fff;">$388</td>
<td style="text-align: center; border: 0px; background-color: #63821C; width: 32px; height: 32px; color: #fff;">$389</td>
<td style="text-align: center; border: 0px; background-color: #498935; width: 32px; height: 32px; color: #fff;">$38A</td>
<td style="text-align: center; border: 0px; background-color: #428856; width: 32px; height: 32px; color: #fff;">$38B</td>
<td style="text-align: center; border: 0px; background-color: #3D7D86; width: 32px; height: 32px; color: #fff;">$38C</td>
<td style="text-align: center; border: 0px; background-color: #41433E; width: 32px; height: 32px; color: #fff;">$38D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$38E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$38F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACACAF; width: 32px; height: 32px; color: #000;">$390</td>
<td style="text-align: center; border: 0px; background-color: #668DBD; width: 32px; height: 32px; color: #000;">$391</td>
<td style="text-align: center; border: 0px; background-color: #7B7ED2; width: 32px; height: 32px; color: #000;">$392</td>
<td style="text-align: center; border: 0px; background-color: #9171D4; width: 32px; height: 32px; color: #000;">$393</td>
<td style="text-align: center; border: 0px; background-color: #A76CBC; width: 32px; height: 32px; color: #000;">$394</td>
<td style="text-align: center; border: 0px; background-color: #B06BA1; width: 32px; height: 32px; color: #000;">$395</td>
<td style="text-align: center; border: 0px; background-color: #B17670; width: 32px; height: 32px; color: #000;">$396</td>
<td style="text-align: center; border: 0px; background-color: #A87F52; width: 32px; height: 32px; color: #000;">$397</td>
<td style="text-align: center; border: 0px; background-color: #929238; width: 32px; height: 32px; color: #000;">$398</td>
<td style="text-align: center; border: 0px; background-color: #7D9D36; width: 32px; height: 32px; color: #000;">$399</td>
<td style="text-align: center; border: 0px; background-color: #64A354; width: 32px; height: 32px; color: #000;">$39A</td>
<td style="text-align: center; border: 0px; background-color: #5CA269; width: 32px; height: 32px; color: #000;">$39B</td>
<td style="text-align: center; border: 0px; background-color: #5B97A1; width: 32px; height: 32px; color: #000;">$39C</td>
<td style="text-align: center; border: 0px; background-color: #5E5E60; width: 32px; height: 32px; color: #fff;">$39D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$39E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$39F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C6C6C8; width: 32px; height: 32px; color: #000;">$3A0</td>
<td style="text-align: center; border: 0px; background-color: #82A6D6; width: 32px; height: 32px; color: #000;">$3A1</td>
<td style="text-align: center; border: 0px; background-color: #9497ED; width: 32px; height: 32px; color: #000;">$3A2</td>
<td style="text-align: center; border: 0px; background-color: #AC8CEF; width: 32px; height: 32px; color: #000;">$3A3</td>
<td style="text-align: center; border: 0px; background-color: #C085D5; width: 32px; height: 32px; color: #000;">$3A4</td>
<td style="text-align: center; border: 0px; background-color: #C984BA; width: 32px; height: 32px; color: #000;">$3A5</td>
<td style="text-align: center; border: 0px; background-color: #CA9085; width: 32px; height: 32px; color: #000;">$3A6</td>
<td style="text-align: center; border: 0px; background-color: #C2996D; width: 32px; height: 32px; color: #000;">$3A7</td>
<td style="text-align: center; border: 0px; background-color: #AFAB52; width: 32px; height: 32px; color: #000;">$3A8</td>
<td style="text-align: center; border: 0px; background-color: #97B551; width: 32px; height: 32px; color: #000;">$3A9</td>
<td style="text-align: center; border: 0px; background-color: #7FBC6C; width: 32px; height: 32px; color: #000;">$3AA</td>
<td style="text-align: center; border: 0px; background-color: #76BB87; width: 32px; height: 32px; color: #000;">$3AB</td>
<td style="text-align: center; border: 0px; background-color: #76B0BB; width: 32px; height: 32px; color: #000;">$3AC</td>
<td style="text-align: center; border: 0px; background-color: #777777; width: 32px; height: 32px; color: #fff;">$3AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DFDFE1; width: 32px; height: 32px; color: #000;">$3B0</td>
<td style="text-align: center; border: 0px; background-color: #9CC0F0; width: 32px; height: 32px; color: #000;">$3B1</td>
<td style="text-align: center; border: 0px; background-color: #AEB1FE; width: 32px; height: 32px; color: #000;">$3B2</td>
<td style="text-align: center; border: 0px; background-color: #C4A6FC; width: 32px; height: 32px; color: #000;">$3B3</td>
<td style="text-align: center; border: 0px; background-color: #DA9FEE; width: 32px; height: 32px; color: #000;">$3B4</td>
<td style="text-align: center; border: 0px; background-color: #E4A0D2; width: 32px; height: 32px; color: #000;">$3B5</td>
<td style="text-align: center; border: 0px; background-color: #E4ABA4; width: 32px; height: 32px; color: #000;">$3B6</td>
<td style="text-align: center; border: 0px; background-color: #DBB588; width: 32px; height: 32px; color: #000;">$3B7</td>
<td style="text-align: center; border: 0px; background-color: #C5C56D; width: 32px; height: 32px; color: #000;">$3B8</td>
<td style="text-align: center; border: 0px; background-color: #B1CE70; width: 32px; height: 32px; color: #000;">$3B9</td>
<td style="text-align: center; border: 0px; background-color: #99D783; width: 32px; height: 32px; color: #000;">$3BA</td>
<td style="text-align: center; border: 0px; background-color: #91D5A3; width: 32px; height: 32px; color: #000;">$3BB</td>
<td style="text-align: center; border: 0px; background-color: #8FCAD2; width: 32px; height: 32px; color: #000;">$3BC</td>
<td style="text-align: center; border: 0px; background-color: #929292; width: 32px; height: 32px; color: #000;">$3BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F7F8F4; width: 32px; height: 32px; color: #000;">$3C0</td>
<td style="text-align: center; border: 0px; background-color: #B5D8FE; width: 32px; height: 32px; color: #000;">$3C1</td>
<td style="text-align: center; border: 0px; background-color: #C8CAFF; width: 32px; height: 32px; color: #000;">$3C2</td>
<td style="text-align: center; border: 0px; background-color: #DEC0FE; width: 32px; height: 32px; color: #000;">$3C3</td>
<td style="text-align: center; border: 0px; background-color: #F5B9FF; width: 32px; height: 32px; color: #000;">$3C4</td>
<td style="text-align: center; border: 0px; background-color: #FBB9E7; width: 32px; height: 32px; color: #000;">$3C5</td>
<td style="text-align: center; border: 0px; background-color: #FCC3BC; width: 32px; height: 32px; color: #000;">$3C6</td>
<td style="text-align: center; border: 0px; background-color: #F3CCA1; width: 32px; height: 32px; color: #000;">$3C7</td>
<td style="text-align: center; border: 0px; background-color: #DEDE88; width: 32px; height: 32px; color: #000;">$3C8</td>
<td style="text-align: center; border: 0px; background-color: #C9E78C; width: 32px; height: 32px; color: #000;">$3C9</td>
<td style="text-align: center; border: 0px; background-color: #B3EDA1; width: 32px; height: 32px; color: #000;">$3CA</td>
<td style="text-align: center; border: 0px; background-color: #ABECC3; width: 32px; height: 32px; color: #000;">$3CB</td>
<td style="text-align: center; border: 0px; background-color: #AAE3EC; width: 32px; height: 32px; color: #000;">$3CC</td>
<td style="text-align: center; border: 0px; background-color: #ACACAC; width: 32px; height: 32px; color: #000;">$3CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$3D0</td>
<td style="text-align: center; border: 0px; background-color: #CCF0FF; width: 32px; height: 32px; color: #000;">$3D1</td>
<td style="text-align: center; border: 0px; background-color: #E1E2FD; width: 32px; height: 32px; color: #000;">$3D2</td>
<td style="text-align: center; border: 0px; background-color: #F5D9FE; width: 32px; height: 32px; color: #000;">$3D3</td>
<td style="text-align: center; border: 0px; background-color: #FED3FF; width: 32px; height: 32px; color: #000;">$3D4</td>
<td style="text-align: center; border: 0px; background-color: #FED3FF; width: 32px; height: 32px; color: #000;">$3D5</td>
<td style="text-align: center; border: 0px; background-color: #FEDCD6; width: 32px; height: 32px; color: #000;">$3D6</td>
<td style="text-align: center; border: 0px; background-color: #FFE5BA; width: 32px; height: 32px; color: #000;">$3D7</td>
<td style="text-align: center; border: 0px; background-color: #F8F5A8; width: 32px; height: 32px; color: #000;">$3D8</td>
<td style="text-align: center; border: 0px; background-color: #E4FDA7; width: 32px; height: 32px; color: #000;">$3D9</td>
<td style="text-align: center; border: 0px; background-color: #CDFFB8; width: 32px; height: 32px; color: #000;">$3DA</td>
<td style="text-align: center; border: 0px; background-color: #C6FED0; width: 32px; height: 32px; color: #000;">$3DB</td>
<td style="text-align: center; border: 0px; background-color: #C4F9FD; width: 32px; height: 32px; color: #000;">$3DC</td>
<td style="text-align: center; border: 0px; background-color: #C5C5C3; width: 32px; height: 32px; color: #000;">$3DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$3E0</td>
<td style="text-align: center; border: 0px; background-color: #E6FFFE; width: 32px; height: 32px; color: #000;">$3E1</td>
<td style="text-align: center; border: 0px; background-color: #F7FBFE; width: 32px; height: 32px; color: #000;">$3E2</td>
<td style="text-align: center; border: 0px; background-color: #FEF1FE; width: 32px; height: 32px; color: #000;">$3E3</td>
<td style="text-align: center; border: 0px; background-color: #FFECFF; width: 32px; height: 32px; color: #000;">$3E4</td>
<td style="text-align: center; border: 0px; background-color: #FFECFF; width: 32px; height: 32px; color: #000;">$3E5</td>
<td style="text-align: center; border: 0px; background-color: #FFF5EC; width: 32px; height: 32px; color: #000;">$3E6</td>
<td style="text-align: center; border: 0px; background-color: #FFFBD8; width: 32px; height: 32px; color: #000;">$3E7</td>
<td style="text-align: center; border: 0px; background-color: #FFFFC3; width: 32px; height: 32px; color: #000;">$3E8</td>
<td style="text-align: center; border: 0px; background-color: #FAFFC4; width: 32px; height: 32px; color: #000;">$3E9</td>
<td style="text-align: center; border: 0px; background-color: #E7FED6; width: 32px; height: 32px; color: #000;">$3EA</td>
<td style="text-align: center; border: 0px; background-color: #E1FEEF; width: 32px; height: 32px; color: #000;">$3EB</td>
<td style="text-align: center; border: 0px; background-color: #DFFEFF; width: 32px; height: 32px; color: #000;">$3EC</td>
<td style="text-align: center; border: 0px; background-color: #DFDFDC; width: 32px; height: 32px; color: #000;">$3ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$3F0</td>
<td style="text-align: center; border: 0px; background-color: #E73F00; width: 32px; height: 32px; color: #fff;">$3F1</td>
<td style="text-align: center; border: 0px; background-color: #887201; width: 32px; height: 32px; color: #fff;">$3F2</td>
<td style="text-align: center; border: 0px; background-color: #2BA100; width: 32px; height: 32px; color: #fff;">$3F3</td>
<td style="text-align: center; border: 0px; background-color: #00CB00; width: 32px; height: 32px; color: #fff;">$3F4</td>
<td style="text-align: center; border: 0px; background-color: #00CB00; width: 32px; height: 32px; color: #fff;">$3F5</td>
<td style="text-align: center; border: 0px; background-color: #009FBD; width: 32px; height: 32px; color: #fff;">$3F6</td>
<td style="text-align: center; border: 0px; background-color: #0176FE; width: 32px; height: 32px; color: #fff;">$3F7</td>
<td style="text-align: center; border: 0px; background-color: #2D3EFD; width: 32px; height: 32px; color: #fff;">$3F8</td>
<td style="text-align: center; border: 0px; background-color: #8213FD; width: 32px; height: 32px; color: #fff;">$3F9</td>
<td style="text-align: center; border: 0px; background-color: #EE00FE; width: 32px; height: 32px; color: #fff;">$3FA</td>
<td style="text-align: center; border: 0px; background-color: #FE00C2; width: 32px; height: 32px; color: #fff;">$3FB</td>
<td style="text-align: center; border: 0px; background-color: #FE1401; width: 32px; height: 32px; color: #fff;">$3FC</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$3FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$3FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $4</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$400</td>
<td style="text-align: center; border: 0px; background-color: #F65800; width: 32px; height: 32px; color: #fff;">$401</td>
<td style="text-align: center; border: 0px; background-color: #9C8E01; width: 32px; height: 32px; color: #000;">$402</td>
<td style="text-align: center; border: 0px; background-color: #47B800; width: 32px; height: 32px; color: #000;">$403</td>
<td style="text-align: center; border: 0px; background-color: #00DC00; width: 32px; height: 32px; color: #000;">$404</td>
<td style="text-align: center; border: 0px; background-color: #02DC14; width: 32px; height: 32px; color: #000;">$405</td>
<td style="text-align: center; border: 0px; background-color: #00B2D0; width: 32px; height: 32px; color: #000;">$406</td>
<td style="text-align: center; border: 0px; background-color: #008CFE; width: 32px; height: 32px; color: #fff;">$407</td>
<td style="text-align: center; border: 0px; background-color: #4D53FE; width: 32px; height: 32px; color: #fff;">$408</td>
<td style="text-align: center; border: 0px; background-color: #972DFE; width: 32px; height: 32px; color: #fff;">$409</td>
<td style="text-align: center; border: 0px; background-color: #FC02FE; width: 32px; height: 32px; color: #fff;">$40A</td>
<td style="text-align: center; border: 0px; background-color: #FD08D0; width: 32px; height: 32px; color: #fff;">$40B</td>
<td style="text-align: center; border: 0px; background-color: #FF3012; width: 32px; height: 32px; color: #fff;">$40C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$40D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$40E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$40F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$410</td>
<td style="text-align: center; border: 0px; background-color: #FF7300; width: 32px; height: 32px; color: #000;">$411</td>
<td style="text-align: center; border: 0px; background-color: #B7A800; width: 32px; height: 32px; color: #000;">$412</td>
<td style="text-align: center; border: 0px; background-color: #61D300; width: 32px; height: 32px; color: #000;">$413</td>
<td style="text-align: center; border: 0px; background-color: #0BF701; width: 32px; height: 32px; color: #000;">$414</td>
<td style="text-align: center; border: 0px; background-color: #00F530; width: 32px; height: 32px; color: #000;">$415</td>
<td style="text-align: center; border: 0px; background-color: #00CCEB; width: 32px; height: 32px; color: #000;">$416</td>
<td style="text-align: center; border: 0px; background-color: #09A5FE; width: 32px; height: 32px; color: #000;">$417</td>
<td style="text-align: center; border: 0px; background-color: #636EFE; width: 32px; height: 32px; color: #fff;">$418</td>
<td style="text-align: center; border: 0px; background-color: #B048FD; width: 32px; height: 32px; color: #fff;">$419</td>
<td style="text-align: center; border: 0px; background-color: #FD1FFD; width: 32px; height: 32px; color: #fff;">$41A</td>
<td style="text-align: center; border: 0px; background-color: #FE22F0; width: 32px; height: 32px; color: #fff;">$41B</td>
<td style="text-align: center; border: 0px; background-color: #FF4B30; width: 32px; height: 32px; color: #fff;">$41C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$41D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$41E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$41F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$420</td>
<td style="text-align: center; border: 0px; background-color: #000031; width: 32px; height: 32px; color: #fff;">$421</td>
<td style="text-align: center; border: 0px; background-color: #00004A; width: 32px; height: 32px; color: #fff;">$422</td>
<td style="text-align: center; border: 0px; background-color: #00004B; width: 32px; height: 32px; color: #fff;">$423</td>
<td style="text-align: center; border: 0px; background-color: #15002D; width: 32px; height: 32px; color: #fff;">$424</td>
<td style="text-align: center; border: 0px; background-color: #20000D; width: 32px; height: 32px; color: #fff;">$425</td>
<td style="text-align: center; border: 0px; background-color: #1F0000; width: 32px; height: 32px; color: #fff;">$426</td>
<td style="text-align: center; border: 0px; background-color: #140001; width: 32px; height: 32px; color: #fff;">$427</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$428</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$429</td>
<td style="text-align: center; border: 0px; background-color: #000E00; width: 32px; height: 32px; color: #fff;">$42A</td>
<td style="text-align: center; border: 0px; background-color: #000D00; width: 32px; height: 32px; color: #fff;">$42B</td>
<td style="text-align: center; border: 0px; background-color: #01010D; width: 32px; height: 32px; color: #fff;">$42C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$42D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$42E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$42F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1B1A1F; width: 32px; height: 32px; color: #fff;">$430</td>
<td style="text-align: center; border: 0px; background-color: #000031; width: 32px; height: 32px; color: #fff;">$431</td>
<td style="text-align: center; border: 0px; background-color: #000054; width: 32px; height: 32px; color: #fff;">$432</td>
<td style="text-align: center; border: 0px; background-color: #00004E; width: 32px; height: 32px; color: #fff;">$433</td>
<td style="text-align: center; border: 0px; background-color: #180032; width: 32px; height: 32px; color: #fff;">$434</td>
<td style="text-align: center; border: 0px; background-color: #24000C; width: 32px; height: 32px; color: #fff;">$435</td>
<td style="text-align: center; border: 0px; background-color: #230000; width: 32px; height: 32px; color: #fff;">$436</td>
<td style="text-align: center; border: 0px; background-color: #180000; width: 32px; height: 32px; color: #fff;">$437</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$438</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$439</td>
<td style="text-align: center; border: 0px; background-color: #001100; width: 32px; height: 32px; color: #fff;">$43A</td>
<td style="text-align: center; border: 0px; background-color: #010F01; width: 32px; height: 32px; color: #fff;">$43B</td>
<td style="text-align: center; border: 0px; background-color: #00020E; width: 32px; height: 32px; color: #fff;">$43C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$43D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$43E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$43F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #363636; width: 32px; height: 32px; color: #fff;">$440</td>
<td style="text-align: center; border: 0px; background-color: #000945; width: 32px; height: 32px; color: #fff;">$441</td>
<td style="text-align: center; border: 0px; background-color: #01006A; width: 32px; height: 32px; color: #fff;">$442</td>
<td style="text-align: center; border: 0px; background-color: #0F0065; width: 32px; height: 32px; color: #fff;">$443</td>
<td style="text-align: center; border: 0px; background-color: #2B0047; width: 32px; height: 32px; color: #fff;">$444</td>
<td style="text-align: center; border: 0px; background-color: #370020; width: 32px; height: 32px; color: #fff;">$445</td>
<td style="text-align: center; border: 0px; background-color: #380000; width: 32px; height: 32px; color: #fff;">$446</td>
<td style="text-align: center; border: 0px; background-color: #2A0100; width: 32px; height: 32px; color: #fff;">$447</td>
<td style="text-align: center; border: 0px; background-color: #0D0E01; width: 32px; height: 32px; color: #fff;">$448</td>
<td style="text-align: center; border: 0px; background-color: #011B00; width: 32px; height: 32px; color: #fff;">$449</td>
<td style="text-align: center; border: 0px; background-color: #002500; width: 32px; height: 32px; color: #fff;">$44A</td>
<td style="text-align: center; border: 0px; background-color: #002200; width: 32px; height: 32px; color: #fff;">$44B</td>
<td style="text-align: center; border: 0px; background-color: #011520; width: 32px; height: 32px; color: #fff;">$44C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$44D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$44E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$44F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$450</td>
<td style="text-align: center; border: 0px; background-color: #002361; width: 32px; height: 32px; color: #fff;">$451</td>
<td style="text-align: center; border: 0px; background-color: #0C0F7A; width: 32px; height: 32px; color: #fff;">$452</td>
<td style="text-align: center; border: 0px; background-color: #29017F; width: 32px; height: 32px; color: #fff;">$453</td>
<td style="text-align: center; border: 0px; background-color: #460062; width: 32px; height: 32px; color: #fff;">$454</td>
<td style="text-align: center; border: 0px; background-color: #54003C; width: 32px; height: 32px; color: #fff;">$455</td>
<td style="text-align: center; border: 0px; background-color: #500500; width: 32px; height: 32px; color: #fff;">$456</td>
<td style="text-align: center; border: 0px; background-color: #471200; width: 32px; height: 32px; color: #fff;">$457</td>
<td style="text-align: center; border: 0px; background-color: #282901; width: 32px; height: 32px; color: #fff;">$458</td>
<td style="text-align: center; border: 0px; background-color: #0D3500; width: 32px; height: 32px; color: #fff;">$459</td>
<td style="text-align: center; border: 0px; background-color: #004100; width: 32px; height: 32px; color: #fff;">$45A</td>
<td style="text-align: center; border: 0px; background-color: #003E00; width: 32px; height: 32px; color: #fff;">$45B</td>
<td style="text-align: center; border: 0px; background-color: #00303C; width: 32px; height: 32px; color: #fff;">$45C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$45D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$45E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$45F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #696967; width: 32px; height: 32px; color: #fff;">$460</td>
<td style="text-align: center; border: 0px; background-color: #093F7F; width: 32px; height: 32px; color: #fff;">$461</td>
<td style="text-align: center; border: 0px; background-color: #2829A1; width: 32px; height: 32px; color: #fff;">$462</td>
<td style="text-align: center; border: 0px; background-color: #441A9A; width: 32px; height: 32px; color: #fff;">$463</td>
<td style="text-align: center; border: 0px; background-color: #63117F; width: 32px; height: 32px; color: #fff;">$464</td>
<td style="text-align: center; border: 0px; background-color: #701155; width: 32px; height: 32px; color: #fff;">$465</td>
<td style="text-align: center; border: 0px; background-color: #702017; width: 32px; height: 32px; color: #fff;">$466</td>
<td style="text-align: center; border: 0px; background-color: #612C00; width: 32px; height: 32px; color: #fff;">$467</td>
<td style="text-align: center; border: 0px; background-color: #454500; width: 32px; height: 32px; color: #fff;">$468</td>
<td style="text-align: center; border: 0px; background-color: #2A5202; width: 32px; height: 32px; color: #fff;">$469</td>
<td style="text-align: center; border: 0px; background-color: #075D00; width: 32px; height: 32px; color: #fff;">$46A</td>
<td style="text-align: center; border: 0px; background-color: #005B13; width: 32px; height: 32px; color: #fff;">$46B</td>
<td style="text-align: center; border: 0px; background-color: #014E58; width: 32px; height: 32px; color: #fff;">$46C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$46D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$46E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$46F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #848484; width: 32px; height: 32px; color: #000;">$470</td>
<td style="text-align: center; border: 0px; background-color: #275A9B; width: 32px; height: 32px; color: #fff;">$471</td>
<td style="text-align: center; border: 0px; background-color: #4247BC; width: 32px; height: 32px; color: #fff;">$472</td>
<td style="text-align: center; border: 0px; background-color: #6037B7; width: 32px; height: 32px; color: #fff;">$473</td>
<td style="text-align: center; border: 0px; background-color: #7F2D9C; width: 32px; height: 32px; color: #fff;">$474</td>
<td style="text-align: center; border: 0px; background-color: #8C2C73; width: 32px; height: 32px; color: #fff;">$475</td>
<td style="text-align: center; border: 0px; background-color: #8C3C33; width: 32px; height: 32px; color: #fff;">$476</td>
<td style="text-align: center; border: 0px; background-color: #804909; width: 32px; height: 32px; color: #fff;">$477</td>
<td style="text-align: center; border: 0px; background-color: #5E6000; width: 32px; height: 32px; color: #fff;">$478</td>
<td style="text-align: center; border: 0px; background-color: #466D01; width: 32px; height: 32px; color: #fff;">$479</td>
<td style="text-align: center; border: 0px; background-color: #24770A; width: 32px; height: 32px; color: #fff;">$47A</td>
<td style="text-align: center; border: 0px; background-color: #177730; width: 32px; height: 32px; color: #fff;">$47B</td>
<td style="text-align: center; border: 0px; background-color: #166873; width: 32px; height: 32px; color: #fff;">$47C</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1A; width: 32px; height: 32px; color: #fff;">$47D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$47E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$47F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9F9F9C; width: 32px; height: 32px; color: #000;">$480</td>
<td style="text-align: center; border: 0px; background-color: #3F74B9; width: 32px; height: 32px; color: #fff;">$481</td>
<td style="text-align: center; border: 0px; background-color: #5D5FD5; width: 32px; height: 32px; color: #fff;">$482</td>
<td style="text-align: center; border: 0px; background-color: #7B51D6; width: 32px; height: 32px; color: #fff;">$483</td>
<td style="text-align: center; border: 0px; background-color: #9A48B7; width: 32px; height: 32px; color: #fff;">$484</td>
<td style="text-align: center; border: 0px; background-color: #A6468D; width: 32px; height: 32px; color: #fff;">$485</td>
<td style="text-align: center; border: 0px; background-color: #A6564B; width: 32px; height: 32px; color: #fff;">$486</td>
<td style="text-align: center; border: 0px; background-color: #9A6225; width: 32px; height: 32px; color: #fff;">$487</td>
<td style="text-align: center; border: 0px; background-color: #7B7B01; width: 32px; height: 32px; color: #fff;">$488</td>
<td style="text-align: center; border: 0px; background-color: #5E8801; width: 32px; height: 32px; color: #fff;">$489</td>
<td style="text-align: center; border: 0px; background-color: #3D9226; width: 32px; height: 32px; color: #fff;">$48A</td>
<td style="text-align: center; border: 0px; background-color: #32924B; width: 32px; height: 32px; color: #fff;">$48B</td>
<td style="text-align: center; border: 0px; background-color: #2E828D; width: 32px; height: 32px; color: #fff;">$48C</td>
<td style="text-align: center; border: 0px; background-color: #343436; width: 32px; height: 32px; color: #fff;">$48D</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$48E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$48F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B8B8B6; width: 32px; height: 32px; color: #000;">$490</td>
<td style="text-align: center; border: 0px; background-color: #5C8FD0; width: 32px; height: 32px; color: #000;">$491</td>
<td style="text-align: center; border: 0px; background-color: #797BF0; width: 32px; height: 32px; color: #000;">$492</td>
<td style="text-align: center; border: 0px; background-color: #956CED; width: 32px; height: 32px; color: #000;">$493</td>
<td style="text-align: center; border: 0px; background-color: #B264D1; width: 32px; height: 32px; color: #000;">$494</td>
<td style="text-align: center; border: 0px; background-color: #BF63A6; width: 32px; height: 32px; color: #000;">$495</td>
<td style="text-align: center; border: 0px; background-color: #C17168; width: 32px; height: 32px; color: #000;">$496</td>
<td style="text-align: center; border: 0px; background-color: #B47E42; width: 32px; height: 32px; color: #000;">$497</td>
<td style="text-align: center; border: 0px; background-color: #96961E; width: 32px; height: 32px; color: #000;">$498</td>
<td style="text-align: center; border: 0px; background-color: #7AA31E; width: 32px; height: 32px; color: #000;">$499</td>
<td style="text-align: center; border: 0px; background-color: #5AAC3E; width: 32px; height: 32px; color: #000;">$49A</td>
<td style="text-align: center; border: 0px; background-color: #4EA967; width: 32px; height: 32px; color: #000;">$49B</td>
<td style="text-align: center; border: 0px; background-color: #4B9DA8; width: 32px; height: 32px; color: #000;">$49C</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$49D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$49E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$49F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D0D3D2; width: 32px; height: 32px; color: #000;">$4A0</td>
<td style="text-align: center; border: 0px; background-color: #76A9EA; width: 32px; height: 32px; color: #000;">$4A1</td>
<td style="text-align: center; border: 0px; background-color: #9294FD; width: 32px; height: 32px; color: #000;">$4A2</td>
<td style="text-align: center; border: 0px; background-color: #AF86FE; width: 32px; height: 32px; color: #000;">$4A3</td>
<td style="text-align: center; border: 0px; background-color: #CE7CEA; width: 32px; height: 32px; color: #000;">$4A4</td>
<td style="text-align: center; border: 0px; background-color: #D97DBE; width: 32px; height: 32px; color: #000;">$4A5</td>
<td style="text-align: center; border: 0px; background-color: #D98D80; width: 32px; height: 32px; color: #000;">$4A6</td>
<td style="text-align: center; border: 0px; background-color: #CE985C; width: 32px; height: 32px; color: #000;">$4A7</td>
<td style="text-align: center; border: 0px; background-color: #AAB239; width: 32px; height: 32px; color: #000;">$4A8</td>
<td style="text-align: center; border: 0px; background-color: #91BE38; width: 32px; height: 32px; color: #000;">$4A9</td>
<td style="text-align: center; border: 0px; background-color: #72C65E; width: 32px; height: 32px; color: #000;">$4AA</td>
<td style="text-align: center; border: 0px; background-color: #65C586; width: 32px; height: 32px; color: #000;">$4AB</td>
<td style="text-align: center; border: 0px; background-color: #65B6C8; width: 32px; height: 32px; color: #000;">$4AC</td>
<td style="text-align: center; border: 0px; background-color: #696A6C; width: 32px; height: 32px; color: #fff;">$4AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBED; width: 32px; height: 32px; color: #000;">$4B0</td>
<td style="text-align: center; border: 0px; background-color: #8FC3FD; width: 32px; height: 32px; color: #000;">$4B1</td>
<td style="text-align: center; border: 0px; background-color: #ABAEFE; width: 32px; height: 32px; color: #000;">$4B2</td>
<td style="text-align: center; border: 0px; background-color: #C79FFE; width: 32px; height: 32px; color: #000;">$4B3</td>
<td style="text-align: center; border: 0px; background-color: #E897FE; width: 32px; height: 32px; color: #000;">$4B4</td>
<td style="text-align: center; border: 0px; background-color: #F198DA; width: 32px; height: 32px; color: #000;">$4B5</td>
<td style="text-align: center; border: 0px; background-color: #F2A59B; width: 32px; height: 32px; color: #000;">$4B6</td>
<td style="text-align: center; border: 0px; background-color: #E5B378; width: 32px; height: 32px; color: #000;">$4B7</td>
<td style="text-align: center; border: 0px; background-color: #C9CA55; width: 32px; height: 32px; color: #000;">$4B8</td>
<td style="text-align: center; border: 0px; background-color: #AED755; width: 32px; height: 32px; color: #000;">$4B9</td>
<td style="text-align: center; border: 0px; background-color: #8DDF7A; width: 32px; height: 32px; color: #000;">$4BA</td>
<td style="text-align: center; border: 0px; background-color: #82DD9D; width: 32px; height: 32px; color: #000;">$4BB</td>
<td style="text-align: center; border: 0px; background-color: #82D0DD; width: 32px; height: 32px; color: #000;">$4BC</td>
<td style="text-align: center; border: 0px; background-color: #848682; width: 32px; height: 32px; color: #000;">$4BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$4C0</td>
<td style="text-align: center; border: 0px; background-color: #A9DAFE; width: 32px; height: 32px; color: #000;">$4C1</td>
<td style="text-align: center; border: 0px; background-color: #C5C8FF; width: 32px; height: 32px; color: #000;">$4C2</td>
<td style="text-align: center; border: 0px; background-color: #E1B9FE; width: 32px; height: 32px; color: #000;">$4C3</td>
<td style="text-align: center; border: 0px; background-color: #FDB0FE; width: 32px; height: 32px; color: #000;">$4C4</td>
<td style="text-align: center; border: 0px; background-color: #FFB1F2; width: 32px; height: 32px; color: #000;">$4C5</td>
<td style="text-align: center; border: 0px; background-color: #FEBEB4; width: 32px; height: 32px; color: #000;">$4C6</td>
<td style="text-align: center; border: 0px; background-color: #FCC992; width: 32px; height: 32px; color: #000;">$4C7</td>
<td style="text-align: center; border: 0px; background-color: #DEE170; width: 32px; height: 32px; color: #000;">$4C8</td>
<td style="text-align: center; border: 0px; background-color: #C6ED70; width: 32px; height: 32px; color: #000;">$4C9</td>
<td style="text-align: center; border: 0px; background-color: #A8F58F; width: 32px; height: 32px; color: #000;">$4CA</td>
<td style="text-align: center; border: 0px; background-color: #9DF3B5; width: 32px; height: 32px; color: #000;">$4CB</td>
<td style="text-align: center; border: 0px; background-color: #9CE6F1; width: 32px; height: 32px; color: #000;">$4CC</td>
<td style="text-align: center; border: 0px; background-color: #9F9FA1; width: 32px; height: 32px; color: #000;">$4CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$4D0</td>
<td style="text-align: center; border: 0px; background-color: #C6F0FE; width: 32px; height: 32px; color: #000;">$4D1</td>
<td style="text-align: center; border: 0px; background-color: #DEDFFE; width: 32px; height: 32px; color: #000;">$4D2</td>
<td style="text-align: center; border: 0px; background-color: #F8D3FE; width: 32px; height: 32px; color: #000;">$4D3</td>
<td style="text-align: center; border: 0px; background-color: #FDCBFE; width: 32px; height: 32px; color: #000;">$4D4</td>
<td style="text-align: center; border: 0px; background-color: #FDCBFE; width: 32px; height: 32px; color: #000;">$4D5</td>
<td style="text-align: center; border: 0px; background-color: #FFD9D0; width: 32px; height: 32px; color: #000;">$4D6</td>
<td style="text-align: center; border: 0px; background-color: #FEE3AE; width: 32px; height: 32px; color: #000;">$4D7</td>
<td style="text-align: center; border: 0px; background-color: #F9F794; width: 32px; height: 32px; color: #000;">$4D8</td>
<td style="text-align: center; border: 0px; background-color: #DFFF8D; width: 32px; height: 32px; color: #000;">$4D9</td>
<td style="text-align: center; border: 0px; background-color: #C4FDAE; width: 32px; height: 32px; color: #000;">$4DA</td>
<td style="text-align: center; border: 0px; background-color: #BBFFCF; width: 32px; height: 32px; color: #000;">$4DB</td>
<td style="text-align: center; border: 0px; background-color: #B7FDFD; width: 32px; height: 32px; color: #000;">$4DC</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B8; width: 32px; height: 32px; color: #000;">$4DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$4E0</td>
<td style="text-align: center; border: 0px; background-color: #BF2301; width: 32px; height: 32px; color: #fff;">$4E1</td>
<td style="text-align: center; border: 0px; background-color: #685800; width: 32px; height: 32px; color: #fff;">$4E2</td>
<td style="text-align: center; border: 0px; background-color: #128300; width: 32px; height: 32px; color: #fff;">$4E3</td>
<td style="text-align: center; border: 0px; background-color: #00A600; width: 32px; height: 32px; color: #fff;">$4E4</td>
<td style="text-align: center; border: 0px; background-color: #00A600; width: 32px; height: 32px; color: #fff;">$4E5</td>
<td style="text-align: center; border: 0px; background-color: #007D99; width: 32px; height: 32px; color: #fff;">$4E6</td>
<td style="text-align: center; border: 0px; background-color: #0157FE; width: 32px; height: 32px; color: #fff;">$4E7</td>
<td style="text-align: center; border: 0px; background-color: #161CFE; width: 32px; height: 32px; color: #fff;">$4E8</td>
<td style="text-align: center; border: 0px; background-color: #6100FE; width: 32px; height: 32px; color: #fff;">$4E9</td>
<td style="text-align: center; border: 0px; background-color: #C500FE; width: 32px; height: 32px; color: #fff;">$4EA</td>
<td style="text-align: center; border: 0px; background-color: #E8009F; width: 32px; height: 32px; color: #fff;">$4EB</td>
<td style="text-align: center; border: 0px; background-color: #EB0001; width: 32px; height: 32px; color: #fff;">$4EC</td>
<td style="text-align: center; border: 0px; background-color: #D1D3D2; width: 32px; height: 32px; color: #000;">$4ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$4F0</td>
<td style="text-align: center; border: 0px; background-color: #DC3E02; width: 32px; height: 32px; color: #fff;">$4F1</td>
<td style="text-align: center; border: 0px; background-color: #847401; width: 32px; height: 32px; color: #fff;">$4F2</td>
<td style="text-align: center; border: 0px; background-color: #2F9F00; width: 32px; height: 32px; color: #fff;">$4F3</td>
<td style="text-align: center; border: 0px; background-color: #00C100; width: 32px; height: 32px; color: #fff;">$4F4</td>
<td style="text-align: center; border: 0px; background-color: #00C100; width: 32px; height: 32px; color: #fff;">$4F5</td>
<td style="text-align: center; border: 0px; background-color: #0099B4; width: 32px; height: 32px; color: #fff;">$4F6</td>
<td style="text-align: center; border: 0px; background-color: #0074FE; width: 32px; height: 32px; color: #fff;">$4F7</td>
<td style="text-align: center; border: 0px; background-color: #303BFE; width: 32px; height: 32px; color: #fff;">$4F8</td>
<td style="text-align: center; border: 0px; background-color: #7B15FE; width: 32px; height: 32px; color: #fff;">$4F9</td>
<td style="text-align: center; border: 0px; background-color: #E000FE; width: 32px; height: 32px; color: #fff;">$4FA</td>
<td style="text-align: center; border: 0px; background-color: #FE02B8; width: 32px; height: 32px; color: #fff;">$4FB</td>
<td style="text-align: center; border: 0px; background-color: #FE1700; width: 32px; height: 32px; color: #fff;">$4FC</td>
<td style="text-align: center; border: 0px; background-color: #EBEDEA; width: 32px; height: 32px; color: #000;">$4FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$4FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $5</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$500</td>
<td style="text-align: center; border: 0px; background-color: #EB5900; width: 32px; height: 32px; color: #fff;">$501</td>
<td style="text-align: center; border: 0px; background-color: #988E01; width: 32px; height: 32px; color: #000;">$502</td>
<td style="text-align: center; border: 0px; background-color: #49B600; width: 32px; height: 32px; color: #000;">$503</td>
<td style="text-align: center; border: 0px; background-color: #00D400; width: 32px; height: 32px; color: #fff;">$504</td>
<td style="text-align: center; border: 0px; background-color: #00D419; width: 32px; height: 32px; color: #fff;">$505</td>
<td style="text-align: center; border: 0px; background-color: #00AEC7; width: 32px; height: 32px; color: #fff;">$506</td>
<td style="text-align: center; border: 0px; background-color: #008BFD; width: 32px; height: 32px; color: #fff;">$507</td>
<td style="text-align: center; border: 0px; background-color: #524FFF; width: 32px; height: 32px; color: #fff;">$508</td>
<td style="text-align: center; border: 0px; background-color: #952CFD; width: 32px; height: 32px; color: #fff;">$509</td>
<td style="text-align: center; border: 0px; background-color: #F00BFE; width: 32px; height: 32px; color: #fff;">$50A</td>
<td style="text-align: center; border: 0px; background-color: #FE0FC5; width: 32px; height: 32px; color: #fff;">$50B</td>
<td style="text-align: center; border: 0px; background-color: #FF3614; width: 32px; height: 32px; color: #fff;">$50C</td>
<td style="text-align: center; border: 0px; background-color: #F8F8F8; width: 32px; height: 32px; color: #000;">$50D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$50E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$50F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$510</td>
<td style="text-align: center; border: 0px; background-color: #FF7400; width: 32px; height: 32px; color: #000;">$511</td>
<td style="text-align: center; border: 0px; background-color: #B1AA00; width: 32px; height: 32px; color: #000;">$512</td>
<td style="text-align: center; border: 0px; background-color: #64D001; width: 32px; height: 32px; color: #000;">$513</td>
<td style="text-align: center; border: 0px; background-color: #12EA00; width: 32px; height: 32px; color: #000;">$514</td>
<td style="text-align: center; border: 0px; background-color: #01EA36; width: 32px; height: 32px; color: #000;">$515</td>
<td style="text-align: center; border: 0px; background-color: #00C5E0; width: 32px; height: 32px; color: #000;">$516</td>
<td style="text-align: center; border: 0px; background-color: #12A3FE; width: 32px; height: 32px; color: #000;">$517</td>
<td style="text-align: center; border: 0px; background-color: #666CFD; width: 32px; height: 32px; color: #fff;">$518</td>
<td style="text-align: center; border: 0px; background-color: #AB49FD; width: 32px; height: 32px; color: #fff;">$519</td>
<td style="text-align: center; border: 0px; background-color: #FE28FE; width: 32px; height: 32px; color: #000;">$51A</td>
<td style="text-align: center; border: 0px; background-color: #FF2AE2; width: 32px; height: 32px; color: #fff;">$51B</td>
<td style="text-align: center; border: 0px; background-color: #FE5138; width: 32px; height: 32px; color: #000;">$51C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$51D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$51E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$51F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0C0C0C; width: 32px; height: 32px; color: #fff;">$520</td>
<td style="text-align: center; border: 0px; background-color: #FE8E01; width: 32px; height: 32px; color: #000;">$521</td>
<td style="text-align: center; border: 0px; background-color: #CAC200; width: 32px; height: 32px; color: #000;">$522</td>
<td style="text-align: center; border: 0px; background-color: #7EE800; width: 32px; height: 32px; color: #000;">$523</td>
<td style="text-align: center; border: 0px; background-color: #31FE01; width: 32px; height: 32px; color: #000;">$524</td>
<td style="text-align: center; border: 0px; background-color: #10FE50; width: 32px; height: 32px; color: #000;">$525</td>
<td style="text-align: center; border: 0px; background-color: #0EDFF6; width: 32px; height: 32px; color: #000;">$526</td>
<td style="text-align: center; border: 0px; background-color: #2DBEFE; width: 32px; height: 32px; color: #000;">$527</td>
<td style="text-align: center; border: 0px; background-color: #8484FE; width: 32px; height: 32px; color: #000;">$528</td>
<td style="text-align: center; border: 0px; background-color: #C761FD; width: 32px; height: 32px; color: #000;">$529</td>
<td style="text-align: center; border: 0px; background-color: #FE43FE; width: 32px; height: 32px; color: #000;">$52A</td>
<td style="text-align: center; border: 0px; background-color: #FE46FB; width: 32px; height: 32px; color: #000;">$52B</td>
<td style="text-align: center; border: 0px; background-color: #FF6A4F; width: 32px; height: 32px; color: #000;">$52C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$52D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$52E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$52F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #292927; width: 32px; height: 32px; color: #fff;">$530</td>
<td style="text-align: center; border: 0px; background-color: #000043; width: 32px; height: 32px; color: #fff;">$531</td>
<td style="text-align: center; border: 0px; background-color: #00006A; width: 32px; height: 32px; color: #fff;">$532</td>
<td style="text-align: center; border: 0px; background-color: #000166; width: 32px; height: 32px; color: #fff;">$533</td>
<td style="text-align: center; border: 0px; background-color: #210144; width: 32px; height: 32px; color: #fff;">$534</td>
<td style="text-align: center; border: 0px; background-color: #310016; width: 32px; height: 32px; color: #fff;">$535</td>
<td style="text-align: center; border: 0px; background-color: #300000; width: 32px; height: 32px; color: #fff;">$536</td>
<td style="text-align: center; border: 0px; background-color: #220000; width: 32px; height: 32px; color: #fff;">$537</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$538</td>
<td style="text-align: center; border: 0px; background-color: #000E01; width: 32px; height: 32px; color: #fff;">$539</td>
<td style="text-align: center; border: 0px; background-color: #001900; width: 32px; height: 32px; color: #fff;">$53A</td>
<td style="text-align: center; border: 0px; background-color: #011700; width: 32px; height: 32px; color: #fff;">$53B</td>
<td style="text-align: center; border: 0px; background-color: #000913; width: 32px; height: 32px; color: #fff;">$53C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$53D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$53E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$53F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$540</td>
<td style="text-align: center; border: 0px; background-color: #000B5C; width: 32px; height: 32px; color: #fff;">$541</td>
<td style="text-align: center; border: 0px; background-color: #000083; width: 32px; height: 32px; color: #fff;">$542</td>
<td style="text-align: center; border: 0px; background-color: #13007C; width: 32px; height: 32px; color: #fff;">$543</td>
<td style="text-align: center; border: 0px; background-color: #380059; width: 32px; height: 32px; color: #fff;">$544</td>
<td style="text-align: center; border: 0px; background-color: #450026; width: 32px; height: 32px; color: #fff;">$545</td>
<td style="text-align: center; border: 0px; background-color: #460000; width: 32px; height: 32px; color: #fff;">$546</td>
<td style="text-align: center; border: 0px; background-color: #360000; width: 32px; height: 32px; color: #fff;">$547</td>
<td style="text-align: center; border: 0px; background-color: #111200; width: 32px; height: 32px; color: #fff;">$548</td>
<td style="text-align: center; border: 0px; background-color: #002201; width: 32px; height: 32px; color: #fff;">$549</td>
<td style="text-align: center; border: 0px; background-color: #002E00; width: 32px; height: 32px; color: #fff;">$54A</td>
<td style="text-align: center; border: 0px; background-color: #002D00; width: 32px; height: 32px; color: #fff;">$54B</td>
<td style="text-align: center; border: 0px; background-color: #001B27; width: 32px; height: 32px; color: #fff;">$54C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$54D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$54E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$54F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #605D5E; width: 32px; height: 32px; color: #fff;">$550</td>
<td style="text-align: center; border: 0px; background-color: #012679; width: 32px; height: 32px; color: #fff;">$551</td>
<td style="text-align: center; border: 0px; background-color: #0C0C9D; width: 32px; height: 32px; color: #fff;">$552</td>
<td style="text-align: center; border: 0px; background-color: #2E009C; width: 32px; height: 32px; color: #fff;">$553</td>
<td style="text-align: center; border: 0px; background-color: #540074; width: 32px; height: 32px; color: #fff;">$554</td>
<td style="text-align: center; border: 0px; background-color: #610043; width: 32px; height: 32px; color: #fff;">$555</td>
<td style="text-align: center; border: 0px; background-color: #640000; width: 32px; height: 32px; color: #fff;">$556</td>
<td style="text-align: center; border: 0px; background-color: #531000; width: 32px; height: 32px; color: #fff;">$557</td>
<td style="text-align: center; border: 0px; background-color: #2C2D01; width: 32px; height: 32px; color: #fff;">$558</td>
<td style="text-align: center; border: 0px; background-color: #0D3E00; width: 32px; height: 32px; color: #fff;">$559</td>
<td style="text-align: center; border: 0px; background-color: #004A00; width: 32px; height: 32px; color: #fff;">$55A</td>
<td style="text-align: center; border: 0px; background-color: #004900; width: 32px; height: 32px; color: #fff;">$55B</td>
<td style="text-align: center; border: 0px; background-color: #013741; width: 32px; height: 32px; color: #fff;">$55C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$55D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$55E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$55F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #797977; width: 32px; height: 32px; color: #fff;">$560</td>
<td style="text-align: center; border: 0px; background-color: #014394; width: 32px; height: 32px; color: #fff;">$561</td>
<td style="text-align: center; border: 0px; background-color: #2626B8; width: 32px; height: 32px; color: #fff;">$562</td>
<td style="text-align: center; border: 0px; background-color: #4A17B8; width: 32px; height: 32px; color: #fff;">$563</td>
<td style="text-align: center; border: 0px; background-color: #6F0B92; width: 32px; height: 32px; color: #fff;">$564</td>
<td style="text-align: center; border: 0px; background-color: #7F0B60; width: 32px; height: 32px; color: #fff;">$565</td>
<td style="text-align: center; border: 0px; background-color: #7F1B0C; width: 32px; height: 32px; color: #fff;">$566</td>
<td style="text-align: center; border: 0px; background-color: #6F2C00; width: 32px; height: 32px; color: #fff;">$567</td>
<td style="text-align: center; border: 0px; background-color: #484800; width: 32px; height: 32px; color: #fff;">$568</td>
<td style="text-align: center; border: 0px; background-color: #285900; width: 32px; height: 32px; color: #fff;">$569</td>
<td style="text-align: center; border: 0px; background-color: #006600; width: 32px; height: 32px; color: #fff;">$56A</td>
<td style="text-align: center; border: 0px; background-color: #00650D; width: 32px; height: 32px; color: #fff;">$56B</td>
<td style="text-align: center; border: 0px; background-color: #00545E; width: 32px; height: 32px; color: #fff;">$56C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$56D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$56E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$56F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$570</td>
<td style="text-align: center; border: 0px; background-color: #1B5EB1; width: 32px; height: 32px; color: #fff;">$571</td>
<td style="text-align: center; border: 0px; background-color: #4243D4; width: 32px; height: 32px; color: #fff;">$572</td>
<td style="text-align: center; border: 0px; background-color: #6530D5; width: 32px; height: 32px; color: #fff;">$573</td>
<td style="text-align: center; border: 0px; background-color: #8B27AE; width: 32px; height: 32px; color: #fff;">$574</td>
<td style="text-align: center; border: 0px; background-color: #9A247A; width: 32px; height: 32px; color: #fff;">$575</td>
<td style="text-align: center; border: 0px; background-color: #993828; width: 32px; height: 32px; color: #fff;">$576</td>
<td style="text-align: center; border: 0px; background-color: #8A4700; width: 32px; height: 32px; color: #fff;">$577</td>
<td style="text-align: center; border: 0px; background-color: #646500; width: 32px; height: 32px; color: #fff;">$578</td>
<td style="text-align: center; border: 0px; background-color: #437401; width: 32px; height: 32px; color: #fff;">$579</td>
<td style="text-align: center; border: 0px; background-color: #188101; width: 32px; height: 32px; color: #fff;">$57A</td>
<td style="text-align: center; border: 0px; background-color: #0A7F2A; width: 32px; height: 32px; color: #fff;">$57B</td>
<td style="text-align: center; border: 0px; background-color: #096F7D; width: 32px; height: 32px; color: #fff;">$57C</td>
<td style="text-align: center; border: 0px; background-color: #0C0D08; width: 32px; height: 32px; color: #fff;">$57D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$57E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$57F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ABADAA; width: 32px; height: 32px; color: #000;">$580</td>
<td style="text-align: center; border: 0px; background-color: #3778C9; width: 32px; height: 32px; color: #fff;">$581</td>
<td style="text-align: center; border: 0px; background-color: #5C5DEE; width: 32px; height: 32px; color: #fff;">$582</td>
<td style="text-align: center; border: 0px; background-color: #7F4CF0; width: 32px; height: 32px; color: #fff;">$583</td>
<td style="text-align: center; border: 0px; background-color: #A440C7; width: 32px; height: 32px; color: #fff;">$584</td>
<td style="text-align: center; border: 0px; background-color: #B44095; width: 32px; height: 32px; color: #fff;">$585</td>
<td style="text-align: center; border: 0px; background-color: #B45343; width: 32px; height: 32px; color: #fff;">$586</td>
<td style="text-align: center; border: 0px; background-color: #A56318; width: 32px; height: 32px; color: #fff;">$587</td>
<td style="text-align: center; border: 0px; background-color: #7F8000; width: 32px; height: 32px; color: #fff;">$588</td>
<td style="text-align: center; border: 0px; background-color: #5F9000; width: 32px; height: 32px; color: #fff;">$589</td>
<td style="text-align: center; border: 0px; background-color: #329D11; width: 32px; height: 32px; color: #fff;">$58A</td>
<td style="text-align: center; border: 0px; background-color: #249B41; width: 32px; height: 32px; color: #fff;">$58B</td>
<td style="text-align: center; border: 0px; background-color: #228A97; width: 32px; height: 32px; color: #fff;">$58C</td>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$58D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$58E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$58F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C5C5C5; width: 32px; height: 32px; color: #000;">$590</td>
<td style="text-align: center; border: 0px; background-color: #5093E2; width: 32px; height: 32px; color: #000;">$591</td>
<td style="text-align: center; border: 0px; background-color: #7577FD; width: 32px; height: 32px; color: #000;">$592</td>
<td style="text-align: center; border: 0px; background-color: #9A65FD; width: 32px; height: 32px; color: #000;">$593</td>
<td style="text-align: center; border: 0px; background-color: #BF5BE2; width: 32px; height: 32px; color: #000;">$594</td>
<td style="text-align: center; border: 0px; background-color: #CE5AAE; width: 32px; height: 32px; color: #000;">$595</td>
<td style="text-align: center; border: 0px; background-color: #D16C62; width: 32px; height: 32px; color: #000;">$596</td>
<td style="text-align: center; border: 0px; background-color: #BF7D33; width: 32px; height: 32px; color: #000;">$597</td>
<td style="text-align: center; border: 0px; background-color: #9B9B05; width: 32px; height: 32px; color: #000;">$598</td>
<td style="text-align: center; border: 0px; background-color: #79AA05; width: 32px; height: 32px; color: #000;">$599</td>
<td style="text-align: center; border: 0px; background-color: #4DB72C; width: 32px; height: 32px; color: #000;">$59A</td>
<td style="text-align: center; border: 0px; background-color: #3FB55E; width: 32px; height: 32px; color: #000;">$59B</td>
<td style="text-align: center; border: 0px; background-color: #3DA5AF; width: 32px; height: 32px; color: #000;">$59C</td>
<td style="text-align: center; border: 0px; background-color: #424242; width: 32px; height: 32px; color: #fff;">$59D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$59E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$59F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DFDFDF; width: 32px; height: 32px; color: #000;">$5A0</td>
<td style="text-align: center; border: 0px; background-color: #6CADFD; width: 32px; height: 32px; color: #000;">$5A1</td>
<td style="text-align: center; border: 0px; background-color: #9093FE; width: 32px; height: 32px; color: #000;">$5A2</td>
<td style="text-align: center; border: 0px; background-color: #B581FE; width: 32px; height: 32px; color: #000;">$5A3</td>
<td style="text-align: center; border: 0px; background-color: #D775FC; width: 32px; height: 32px; color: #000;">$5A4</td>
<td style="text-align: center; border: 0px; background-color: #EA76CB; width: 32px; height: 32px; color: #000;">$5A5</td>
<td style="text-align: center; border: 0px; background-color: #EA887D; width: 32px; height: 32px; color: #000;">$5A6</td>
<td style="text-align: center; border: 0px; background-color: #DB954D; width: 32px; height: 32px; color: #000;">$5A7</td>
<td style="text-align: center; border: 0px; background-color: #B4B722; width: 32px; height: 32px; color: #000;">$5A8</td>
<td style="text-align: center; border: 0px; background-color: #94C522; width: 32px; height: 32px; color: #000;">$5A9</td>
<td style="text-align: center; border: 0px; background-color: #6ACF4E; width: 32px; height: 32px; color: #000;">$5AA</td>
<td style="text-align: center; border: 0px; background-color: #5ACF77; width: 32px; height: 32px; color: #000;">$5AB</td>
<td style="text-align: center; border: 0px; background-color: #58BECA; width: 32px; height: 32px; color: #000;">$5AC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$5AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F5F8F7; width: 32px; height: 32px; color: #000;">$5B0</td>
<td style="text-align: center; border: 0px; background-color: #84C6FE; width: 32px; height: 32px; color: #000;">$5B1</td>
<td style="text-align: center; border: 0px; background-color: #A9ABFE; width: 32px; height: 32px; color: #000;">$5B2</td>
<td style="text-align: center; border: 0px; background-color: #CC99FE; width: 32px; height: 32px; color: #000;">$5B3</td>
<td style="text-align: center; border: 0px; background-color: #F191FF; width: 32px; height: 32px; color: #000;">$5B4</td>
<td style="text-align: center; border: 0px; background-color: #FD90E3; width: 32px; height: 32px; color: #000;">$5B5</td>
<td style="text-align: center; border: 0px; background-color: #FEA096; width: 32px; height: 32px; color: #000;">$5B6</td>
<td style="text-align: center; border: 0px; background-color: #F2AF68; width: 32px; height: 32px; color: #000;">$5B7</td>
<td style="text-align: center; border: 0px; background-color: #CBCF3A; width: 32px; height: 32px; color: #000;">$5B8</td>
<td style="text-align: center; border: 0px; background-color: #ACDD3D; width: 32px; height: 32px; color: #000;">$5B9</td>
<td style="text-align: center; border: 0px; background-color: #82E964; width: 32px; height: 32px; color: #000;">$5BA</td>
<td style="text-align: center; border: 0px; background-color: #74E895; width: 32px; height: 32px; color: #000;">$5BB</td>
<td style="text-align: center; border: 0px; background-color: #71D7E5; width: 32px; height: 32px; color: #000;">$5BC</td>
<td style="text-align: center; border: 0px; background-color: #777777; width: 32px; height: 32px; color: #fff;">$5BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$5C0</td>
<td style="text-align: center; border: 0px; background-color: #A0DDFE; width: 32px; height: 32px; color: #000;">$5C1</td>
<td style="text-align: center; border: 0px; background-color: #C2C4FD; width: 32px; height: 32px; color: #000;">$5C2</td>
<td style="text-align: center; border: 0px; background-color: #E4B4FE; width: 32px; height: 32px; color: #000;">$5C3</td>
<td style="text-align: center; border: 0px; background-color: #FFABFF; width: 32px; height: 32px; color: #000;">$5C4</td>
<td style="text-align: center; border: 0px; background-color: #FFABF9; width: 32px; height: 32px; color: #000;">$5C5</td>
<td style="text-align: center; border: 0px; background-color: #FFBBB0; width: 32px; height: 32px; color: #000;">$5C6</td>
<td style="text-align: center; border: 0px; background-color: #FFCA84; width: 32px; height: 32px; color: #000;">$5C7</td>
<td style="text-align: center; border: 0px; background-color: #E5E45E; width: 32px; height: 32px; color: #000;">$5C8</td>
<td style="text-align: center; border: 0px; background-color: #C5F35C; width: 32px; height: 32px; color: #000;">$5C9</td>
<td style="text-align: center; border: 0px; background-color: #9FFD84; width: 32px; height: 32px; color: #000;">$5CA</td>
<td style="text-align: center; border: 0px; background-color: #92FDAF; width: 32px; height: 32px; color: #000;">$5CB</td>
<td style="text-align: center; border: 0px; background-color: #90EEF5; width: 32px; height: 32px; color: #000;">$5CC</td>
<td style="text-align: center; border: 0px; background-color: #939391; width: 32px; height: 32px; color: #000;">$5CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$5D0</td>
<td style="text-align: center; border: 0px; background-color: #BCF4FF; width: 32px; height: 32px; color: #000;">$5D1</td>
<td style="text-align: center; border: 0px; background-color: #DCDCFE; width: 32px; height: 32px; color: #000;">$5D2</td>
<td style="text-align: center; border: 0px; background-color: #FCCEFE; width: 32px; height: 32px; color: #000;">$5D3</td>
<td style="text-align: center; border: 0px; background-color: #FFC4FE; width: 32px; height: 32px; color: #000;">$5D4</td>
<td style="text-align: center; border: 0px; background-color: #FFC3FF; width: 32px; height: 32px; color: #000;">$5D5</td>
<td style="text-align: center; border: 0px; background-color: #FFD3C8; width: 32px; height: 32px; color: #000;">$5D6</td>
<td style="text-align: center; border: 0px; background-color: #FEE2A0; width: 32px; height: 32px; color: #000;">$5D7</td>
<td style="text-align: center; border: 0px; background-color: #FDFC7A; width: 32px; height: 32px; color: #000;">$5D8</td>
<td style="text-align: center; border: 0px; background-color: #DEFF7A; width: 32px; height: 32px; color: #000;">$5D9</td>
<td style="text-align: center; border: 0px; background-color: #BAFE9D; width: 32px; height: 32px; color: #000;">$5DA</td>
<td style="text-align: center; border: 0px; background-color: #AEFEC9; width: 32px; height: 32px; color: #000;">$5DB</td>
<td style="text-align: center; border: 0px; background-color: #ABFFFF; width: 32px; height: 32px; color: #000;">$5DC</td>
<td style="text-align: center; border: 0px; background-color: #AAADAC; width: 32px; height: 32px; color: #000;">$5DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$5E0</td>
<td style="text-align: center; border: 0px; background-color: #B52401; width: 32px; height: 32px; color: #fff;">$5E1</td>
<td style="text-align: center; border: 0px; background-color: #605900; width: 32px; height: 32px; color: #fff;">$5E2</td>
<td style="text-align: center; border: 0px; background-color: #118100; width: 32px; height: 32px; color: #fff;">$5E3</td>
<td style="text-align: center; border: 0px; background-color: #009D00; width: 32px; height: 32px; color: #fff;">$5E4</td>
<td style="text-align: center; border: 0px; background-color: #009D00; width: 32px; height: 32px; color: #fff;">$5E5</td>
<td style="text-align: center; border: 0px; background-color: #007696; width: 32px; height: 32px; color: #fff;">$5E6</td>
<td style="text-align: center; border: 0px; background-color: #0153FB; width: 32px; height: 32px; color: #fff;">$5E7</td>
<td style="text-align: center; border: 0px; background-color: #121BFF; width: 32px; height: 32px; color: #fff;">$5E8</td>
<td style="text-align: center; border: 0px; background-color: #5B01FE; width: 32px; height: 32px; color: #fff;">$5E9</td>
<td style="text-align: center; border: 0px; background-color: #B900FE; width: 32px; height: 32px; color: #fff;">$5EA</td>
<td style="text-align: center; border: 0px; background-color: #D90097; width: 32px; height: 32px; color: #fff;">$5EB</td>
<td style="text-align: center; border: 0px; background-color: #DF0100; width: 32px; height: 32px; color: #fff;">$5EC</td>
<td style="text-align: center; border: 0px; background-color: #C6C7C2; width: 32px; height: 32px; color: #000;">$5ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$5F0</td>
<td style="text-align: center; border: 0px; background-color: #D24000; width: 32px; height: 32px; color: #fff;">$5F1</td>
<td style="text-align: center; border: 0px; background-color: #7F7400; width: 32px; height: 32px; color: #fff;">$5F2</td>
<td style="text-align: center; border: 0px; background-color: #2C9D00; width: 32px; height: 32px; color: #fff;">$5F3</td>
<td style="text-align: center; border: 0px; background-color: #00B900; width: 32px; height: 32px; color: #fff;">$5F4</td>
<td style="text-align: center; border: 0px; background-color: #00B800; width: 32px; height: 32px; color: #fff;">$5F5</td>
<td style="text-align: center; border: 0px; background-color: #0192AD; width: 32px; height: 32px; color: #fff;">$5F6</td>
<td style="text-align: center; border: 0px; background-color: #006EFE; width: 32px; height: 32px; color: #fff;">$5F7</td>
<td style="text-align: center; border: 0px; background-color: #3635FD; width: 32px; height: 32px; color: #fff;">$5F8</td>
<td style="text-align: center; border: 0px; background-color: #7D12FD; width: 32px; height: 32px; color: #fff;">$5F9</td>
<td style="text-align: center; border: 0px; background-color: #DA00FE; width: 32px; height: 32px; color: #fff;">$5FA</td>
<td style="text-align: center; border: 0px; background-color: #F900B1; width: 32px; height: 32px; color: #fff;">$5FB</td>
<td style="text-align: center; border: 0px; background-color: #FB1A01; width: 32px; height: 32px; color: #fff;">$5FC</td>
<td style="text-align: center; border: 0px; background-color: #DFDEE3; width: 32px; height: 32px; color: #000;">$5FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$5FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $6</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$600</td>
<td style="text-align: center; border: 0px; background-color: #E15B00; width: 32px; height: 32px; color: #fff;">$601</td>
<td style="text-align: center; border: 0px; background-color: #948E00; width: 32px; height: 32px; color: #fff;">$602</td>
<td style="text-align: center; border: 0px; background-color: #4CB400; width: 32px; height: 32px; color: #000;">$603</td>
<td style="text-align: center; border: 0px; background-color: #02CB01; width: 32px; height: 32px; color: #fff;">$604</td>
<td style="text-align: center; border: 0px; background-color: #00CB21; width: 32px; height: 32px; color: #fff;">$605</td>
<td style="text-align: center; border: 0px; background-color: #01A7BD; width: 32px; height: 32px; color: #fff;">$606</td>
<td style="text-align: center; border: 0px; background-color: #0186FE; width: 32px; height: 32px; color: #fff;">$607</td>
<td style="text-align: center; border: 0px; background-color: #514EFE; width: 32px; height: 32px; color: #fff;">$608</td>
<td style="text-align: center; border: 0px; background-color: #8F2FFD; width: 32px; height: 32px; color: #fff;">$609</td>
<td style="text-align: center; border: 0px; background-color: #E611FD; width: 32px; height: 32px; color: #fff;">$60A</td>
<td style="text-align: center; border: 0px; background-color: #FE13C4; width: 32px; height: 32px; color: #fff;">$60B</td>
<td style="text-align: center; border: 0px; background-color: #FE381B; width: 32px; height: 32px; color: #fff;">$60C</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$60D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$60E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$60F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$610</td>
<td style="text-align: center; border: 0px; background-color: #F97500; width: 32px; height: 32px; color: #000;">$611</td>
<td style="text-align: center; border: 0px; background-color: #ADA901; width: 32px; height: 32px; color: #000;">$612</td>
<td style="text-align: center; border: 0px; background-color: #66CD01; width: 32px; height: 32px; color: #000;">$613</td>
<td style="text-align: center; border: 0px; background-color: #1BE400; width: 32px; height: 32px; color: #000;">$614</td>
<td style="text-align: center; border: 0px; background-color: #00E33A; width: 32px; height: 32px; color: #000;">$615</td>
<td style="text-align: center; border: 0px; background-color: #00BFD9; width: 32px; height: 32px; color: #000;">$616</td>
<td style="text-align: center; border: 0px; background-color: #1AA0FE; width: 32px; height: 32px; color: #000;">$617</td>
<td style="text-align: center; border: 0px; background-color: #656AFE; width: 32px; height: 32px; color: #fff;">$618</td>
<td style="text-align: center; border: 0px; background-color: #A748FC; width: 32px; height: 32px; color: #fff;">$619</td>
<td style="text-align: center; border: 0px; background-color: #FB2DFE; width: 32px; height: 32px; color: #000;">$61A</td>
<td style="text-align: center; border: 0px; background-color: #FF31DF; width: 32px; height: 32px; color: #000;">$61B</td>
<td style="text-align: center; border: 0px; background-color: #FE523D; width: 32px; height: 32px; color: #000;">$61C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$61D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$61E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$61F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1A1A18; width: 32px; height: 32px; color: #fff;">$620</td>
<td style="text-align: center; border: 0px; background-color: #FE8F00; width: 32px; height: 32px; color: #000;">$621</td>
<td style="text-align: center; border: 0px; background-color: #C5C200; width: 32px; height: 32px; color: #000;">$622</td>
<td style="text-align: center; border: 0px; background-color: #83E400; width: 32px; height: 32px; color: #000;">$623</td>
<td style="text-align: center; border: 0px; background-color: #3BF901; width: 32px; height: 32px; color: #000;">$624</td>
<td style="text-align: center; border: 0px; background-color: #1CFA57; width: 32px; height: 32px; color: #000;">$625</td>
<td style="text-align: center; border: 0px; background-color: #1AD8F0; width: 32px; height: 32px; color: #000;">$626</td>
<td style="text-align: center; border: 0px; background-color: #37BAFD; width: 32px; height: 32px; color: #000;">$627</td>
<td style="text-align: center; border: 0px; background-color: #8681FE; width: 32px; height: 32px; color: #000;">$628</td>
<td style="text-align: center; border: 0px; background-color: #C362FE; width: 32px; height: 32px; color: #000;">$629</td>
<td style="text-align: center; border: 0px; background-color: #FD4BFE; width: 32px; height: 32px; color: #000;">$62A</td>
<td style="text-align: center; border: 0px; background-color: #FE4EEE; width: 32px; height: 32px; color: #000;">$62B</td>
<td style="text-align: center; border: 0px; background-color: #FE7250; width: 32px; height: 32px; color: #000;">$62C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$62D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$62E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$62F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #363636; width: 32px; height: 32px; color: #fff;">$630</td>
<td style="text-align: center; border: 0px; background-color: #000152; width: 32px; height: 32px; color: #fff;">$631</td>
<td style="text-align: center; border: 0px; background-color: #00007E; width: 32px; height: 32px; color: #fff;">$632</td>
<td style="text-align: center; border: 0px; background-color: #02017B; width: 32px; height: 32px; color: #fff;">$633</td>
<td style="text-align: center; border: 0px; background-color: #2B0052; width: 32px; height: 32px; color: #fff;">$634</td>
<td style="text-align: center; border: 0px; background-color: #3F001B; width: 32px; height: 32px; color: #fff;">$635</td>
<td style="text-align: center; border: 0px; background-color: #3C0000; width: 32px; height: 32px; color: #fff;">$636</td>
<td style="text-align: center; border: 0px; background-color: #2A0001; width: 32px; height: 32px; color: #fff;">$637</td>
<td style="text-align: center; border: 0px; background-color: #010100; width: 32px; height: 32px; color: #fff;">$638</td>
<td style="text-align: center; border: 0px; background-color: #001201; width: 32px; height: 32px; color: #fff;">$639</td>
<td style="text-align: center; border: 0px; background-color: #002100; width: 32px; height: 32px; color: #fff;">$63A</td>
<td style="text-align: center; border: 0px; background-color: #001F00; width: 32px; height: 32px; color: #fff;">$63B</td>
<td style="text-align: center; border: 0px; background-color: #000E1B; width: 32px; height: 32px; color: #fff;">$63C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$63D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$63E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$63F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #525252; width: 32px; height: 32px; color: #fff;">$640</td>
<td style="text-align: center; border: 0px; background-color: #000D72; width: 32px; height: 32px; color: #fff;">$641</td>
<td style="text-align: center; border: 0px; background-color: #000095; width: 32px; height: 32px; color: #fff;">$642</td>
<td style="text-align: center; border: 0px; background-color: #180099; width: 32px; height: 32px; color: #fff;">$643</td>
<td style="text-align: center; border: 0px; background-color: #43006A; width: 32px; height: 32px; color: #fff;">$644</td>
<td style="text-align: center; border: 0px; background-color: #540031; width: 32px; height: 32px; color: #fff;">$645</td>
<td style="text-align: center; border: 0px; background-color: #540000; width: 32px; height: 32px; color: #fff;">$646</td>
<td style="text-align: center; border: 0px; background-color: #440001; width: 32px; height: 32px; color: #fff;">$647</td>
<td style="text-align: center; border: 0px; background-color: #151701; width: 32px; height: 32px; color: #fff;">$648</td>
<td style="text-align: center; border: 0px; background-color: #002B02; width: 32px; height: 32px; color: #fff;">$649</td>
<td style="text-align: center; border: 0px; background-color: #003800; width: 32px; height: 32px; color: #fff;">$64A</td>
<td style="text-align: center; border: 0px; background-color: #003600; width: 32px; height: 32px; color: #fff;">$64B</td>
<td style="text-align: center; border: 0px; background-color: #002132; width: 32px; height: 32px; color: #fff;">$64C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$64D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$64E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$64F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6B6B6B; width: 32px; height: 32px; color: #fff;">$650</td>
<td style="text-align: center; border: 0px; background-color: #012989; width: 32px; height: 32px; color: #fff;">$651</td>
<td style="text-align: center; border: 0px; background-color: #080AB5; width: 32px; height: 32px; color: #fff;">$652</td>
<td style="text-align: center; border: 0px; background-color: #3100B5; width: 32px; height: 32px; color: #fff;">$653</td>
<td style="text-align: center; border: 0px; background-color: #5D0087; width: 32px; height: 32px; color: #fff;">$654</td>
<td style="text-align: center; border: 0px; background-color: #6F004B; width: 32px; height: 32px; color: #fff;">$655</td>
<td style="text-align: center; border: 0px; background-color: #6E0000; width: 32px; height: 32px; color: #fff;">$656</td>
<td style="text-align: center; border: 0px; background-color: #5D1101; width: 32px; height: 32px; color: #fff;">$657</td>
<td style="text-align: center; border: 0px; background-color: #2F3100; width: 32px; height: 32px; color: #fff;">$658</td>
<td style="text-align: center; border: 0px; background-color: #094500; width: 32px; height: 32px; color: #fff;">$659</td>
<td style="text-align: center; border: 0px; background-color: #005300; width: 32px; height: 32px; color: #fff;">$65A</td>
<td style="text-align: center; border: 0px; background-color: #005100; width: 32px; height: 32px; color: #fff;">$65B</td>
<td style="text-align: center; border: 0px; background-color: #013C4A; width: 32px; height: 32px; color: #fff;">$65C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$65D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$65E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$65F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #868686; width: 32px; height: 32px; color: #000;">$660</td>
<td style="text-align: center; border: 0px; background-color: #0044A4; width: 32px; height: 32px; color: #fff;">$661</td>
<td style="text-align: center; border: 0px; background-color: #2124D5; width: 32px; height: 32px; color: #fff;">$662</td>
<td style="text-align: center; border: 0px; background-color: #4D10D3; width: 32px; height: 32px; color: #fff;">$663</td>
<td style="text-align: center; border: 0px; background-color: #7903A5; width: 32px; height: 32px; color: #fff;">$664</td>
<td style="text-align: center; border: 0px; background-color: #8E0169; width: 32px; height: 32px; color: #fff;">$665</td>
<td style="text-align: center; border: 0px; background-color: #8D1709; width: 32px; height: 32px; color: #fff;">$666</td>
<td style="text-align: center; border: 0px; background-color: #7A2B00; width: 32px; height: 32px; color: #fff;">$667</td>
<td style="text-align: center; border: 0px; background-color: #4C4D00; width: 32px; height: 32px; color: #fff;">$668</td>
<td style="text-align: center; border: 0px; background-color: #266202; width: 32px; height: 32px; color: #fff;">$669</td>
<td style="text-align: center; border: 0px; background-color: #007101; width: 32px; height: 32px; color: #fff;">$66A</td>
<td style="text-align: center; border: 0px; background-color: #006F06; width: 32px; height: 32px; color: #fff;">$66B</td>
<td style="text-align: center; border: 0px; background-color: #015A6C; width: 32px; height: 32px; color: #fff;">$66C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$66D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$66E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$66F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9F9F9F; width: 32px; height: 32px; color: #000;">$670</td>
<td style="text-align: center; border: 0px; background-color: #1260C1; width: 32px; height: 32px; color: #fff;">$671</td>
<td style="text-align: center; border: 0px; background-color: #3D41EA; width: 32px; height: 32px; color: #fff;">$672</td>
<td style="text-align: center; border: 0px; background-color: #682BED; width: 32px; height: 32px; color: #fff;">$673</td>
<td style="text-align: center; border: 0px; background-color: #931FBE; width: 32px; height: 32px; color: #fff;">$674</td>
<td style="text-align: center; border: 0px; background-color: #A81D84; width: 32px; height: 32px; color: #fff;">$675</td>
<td style="text-align: center; border: 0px; background-color: #A83224; width: 32px; height: 32px; color: #fff;">$676</td>
<td style="text-align: center; border: 0px; background-color: #944500; width: 32px; height: 32px; color: #fff;">$677</td>
<td style="text-align: center; border: 0px; background-color: #676900; width: 32px; height: 32px; color: #fff;">$678</td>
<td style="text-align: center; border: 0px; background-color: #417B01; width: 32px; height: 32px; color: #fff;">$679</td>
<td style="text-align: center; border: 0px; background-color: #0F8A01; width: 32px; height: 32px; color: #fff;">$67A</td>
<td style="text-align: center; border: 0px; background-color: #008723; width: 32px; height: 32px; color: #fff;">$67B</td>
<td style="text-align: center; border: 0px; background-color: #007581; width: 32px; height: 32px; color: #fff;">$67C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$67D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$67E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$67F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #BABABA; width: 32px; height: 32px; color: #000;">$680</td>
<td style="text-align: center; border: 0px; background-color: #2D7CDC; width: 32px; height: 32px; color: #fff;">$681</td>
<td style="text-align: center; border: 0px; background-color: #5A5BFE; width: 32px; height: 32px; color: #fff;">$682</td>
<td style="text-align: center; border: 0px; background-color: #8446FD; width: 32px; height: 32px; color: #fff;">$683</td>
<td style="text-align: center; border: 0px; background-color: #AF39D9; width: 32px; height: 32px; color: #fff;">$684</td>
<td style="text-align: center; border: 0px; background-color: #C3389D; width: 32px; height: 32px; color: #fff;">$685</td>
<td style="text-align: center; border: 0px; background-color: #C44F3E; width: 32px; height: 32px; color: #fff;">$686</td>
<td style="text-align: center; border: 0px; background-color: #AF6208; width: 32px; height: 32px; color: #fff;">$687</td>
<td style="text-align: center; border: 0px; background-color: #838500; width: 32px; height: 32px; color: #fff;">$688</td>
<td style="text-align: center; border: 0px; background-color: #5D9800; width: 32px; height: 32px; color: #fff;">$689</td>
<td style="text-align: center; border: 0px; background-color: #2CA603; width: 32px; height: 32px; color: #fff;">$68A</td>
<td style="text-align: center; border: 0px; background-color: #18A43E; width: 32px; height: 32px; color: #fff;">$68B</td>
<td style="text-align: center; border: 0px; background-color: #16909F; width: 32px; height: 32px; color: #fff;">$68C</td>
<td style="text-align: center; border: 0px; background-color: #1A1A1C; width: 32px; height: 32px; color: #fff;">$68D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$68E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$68F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D2D2D4; width: 32px; height: 32px; color: #000;">$690</td>
<td style="text-align: center; border: 0px; background-color: #4596F6; width: 32px; height: 32px; color: #000;">$691</td>
<td style="text-align: center; border: 0px; background-color: #7476FF; width: 32px; height: 32px; color: #000;">$692</td>
<td style="text-align: center; border: 0px; background-color: #9D60FE; width: 32px; height: 32px; color: #000;">$693</td>
<td style="text-align: center; border: 0px; background-color: #C953F5; width: 32px; height: 32px; color: #000;">$694</td>
<td style="text-align: center; border: 0px; background-color: #DD52B9; width: 32px; height: 32px; color: #000;">$695</td>
<td style="text-align: center; border: 0px; background-color: #E0685A; width: 32px; height: 32px; color: #000;">$696</td>
<td style="text-align: center; border: 0px; background-color: #CB7B22; width: 32px; height: 32px; color: #000;">$697</td>
<td style="text-align: center; border: 0px; background-color: #9EA000; width: 32px; height: 32px; color: #000;">$698</td>
<td style="text-align: center; border: 0px; background-color: #76B200; width: 32px; height: 32px; color: #000;">$699</td>
<td style="text-align: center; border: 0px; background-color: #45C01B; width: 32px; height: 32px; color: #000;">$69A</td>
<td style="text-align: center; border: 0px; background-color: #32BF54; width: 32px; height: 32px; color: #000;">$69B</td>
<td style="text-align: center; border: 0px; background-color: #2FABB5; width: 32px; height: 32px; color: #000;">$69C</td>
<td style="text-align: center; border: 0px; background-color: #353535; width: 32px; height: 32px; color: #fff;">$69D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$69E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$69F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$6A0</td>
<td style="text-align: center; border: 0px; background-color: #61B0FE; width: 32px; height: 32px; color: #000;">$6A1</td>
<td style="text-align: center; border: 0px; background-color: #8D90FD; width: 32px; height: 32px; color: #000;">$6A2</td>
<td style="text-align: center; border: 0px; background-color: #B77CFD; width: 32px; height: 32px; color: #000;">$6A3</td>
<td style="text-align: center; border: 0px; background-color: #E16EFF; width: 32px; height: 32px; color: #000;">$6A4</td>
<td style="text-align: center; border: 0px; background-color: #F56ED1; width: 32px; height: 32px; color: #000;">$6A5</td>
<td style="text-align: center; border: 0px; background-color: #F48473; width: 32px; height: 32px; color: #000;">$6A6</td>
<td style="text-align: center; border: 0px; background-color: #E39540; width: 32px; height: 32px; color: #000;">$6A7</td>
<td style="text-align: center; border: 0px; background-color: #B8B708; width: 32px; height: 32px; color: #000;">$6A8</td>
<td style="text-align: center; border: 0px; background-color: #92CA07; width: 32px; height: 32px; color: #000;">$6A9</td>
<td style="text-align: center; border: 0px; background-color: #60D935; width: 32px; height: 32px; color: #000;">$6AA</td>
<td style="text-align: center; border: 0px; background-color: #4DD96D; width: 32px; height: 32px; color: #000;">$6AB</td>
<td style="text-align: center; border: 0px; background-color: #4CC4D2; width: 32px; height: 32px; color: #000;">$6AC</td>
<td style="text-align: center; border: 0px; background-color: #515151; width: 32px; height: 32px; color: #fff;">$6AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$6B0</td>
<td style="text-align: center; border: 0px; background-color: #7CC8FF; width: 32px; height: 32px; color: #000;">$6B1</td>
<td style="text-align: center; border: 0px; background-color: #A6A9FD; width: 32px; height: 32px; color: #000;">$6B2</td>
<td style="text-align: center; border: 0px; background-color: #CE95FE; width: 32px; height: 32px; color: #000;">$6B3</td>
<td style="text-align: center; border: 0px; background-color: #F88AFD; width: 32px; height: 32px; color: #000;">$6B4</td>
<td style="text-align: center; border: 0px; background-color: #FF88E9; width: 32px; height: 32px; color: #000;">$6B5</td>
<td style="text-align: center; border: 0px; background-color: #FF9C8D; width: 32px; height: 32px; color: #000;">$6B6</td>
<td style="text-align: center; border: 0px; background-color: #FBAF59; width: 32px; height: 32px; color: #000;">$6B7</td>
<td style="text-align: center; border: 0px; background-color: #CED125; width: 32px; height: 32px; color: #000;">$6B8</td>
<td style="text-align: center; border: 0px; background-color: #A7E422; width: 32px; height: 32px; color: #000;">$6B9</td>
<td style="text-align: center; border: 0px; background-color: #79F155; width: 32px; height: 32px; color: #000;">$6BA</td>
<td style="text-align: center; border: 0px; background-color: #67EF8E; width: 32px; height: 32px; color: #000;">$6BB</td>
<td style="text-align: center; border: 0px; background-color: #69D9EE; width: 32px; height: 32px; color: #000;">$6BC</td>
<td style="text-align: center; border: 0px; background-color: #696B67; width: 32px; height: 32px; color: #fff;">$6BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$6C0</td>
<td style="text-align: center; border: 0px; background-color: #96E0FF; width: 32px; height: 32px; color: #000;">$6C1</td>
<td style="text-align: center; border: 0px; background-color: #C1C3FE; width: 32px; height: 32px; color: #000;">$6C2</td>
<td style="text-align: center; border: 0px; background-color: #E6B0FE; width: 32px; height: 32px; color: #000;">$6C3</td>
<td style="text-align: center; border: 0px; background-color: #FEA4FE; width: 32px; height: 32px; color: #000;">$6C4</td>
<td style="text-align: center; border: 0px; background-color: #FEA4FE; width: 32px; height: 32px; color: #000;">$6C5</td>
<td style="text-align: center; border: 0px; background-color: #FEB7A9; width: 32px; height: 32px; color: #000;">$6C6</td>
<td style="text-align: center; border: 0px; background-color: #FEC875; width: 32px; height: 32px; color: #000;">$6C7</td>
<td style="text-align: center; border: 0px; background-color: #E6E943; width: 32px; height: 32px; color: #000;">$6C8</td>
<td style="text-align: center; border: 0px; background-color: #C2F843; width: 32px; height: 32px; color: #000;">$6C9</td>
<td style="text-align: center; border: 0px; background-color: #96FE73; width: 32px; height: 32px; color: #000;">$6CA</td>
<td style="text-align: center; border: 0px; background-color: #85FEA5; width: 32px; height: 32px; color: #000;">$6CB</td>
<td style="text-align: center; border: 0px; background-color: #85F1FE; width: 32px; height: 32px; color: #000;">$6CC</td>
<td style="text-align: center; border: 0px; background-color: #858487; width: 32px; height: 32px; color: #000;">$6CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$6D0</td>
<td style="text-align: center; border: 0px; background-color: #890800; width: 32px; height: 32px; color: #fff;">$6D1</td>
<td style="text-align: center; border: 0px; background-color: #413B01; width: 32px; height: 32px; color: #fff;">$6D2</td>
<td style="text-align: center; border: 0px; background-color: #005F00; width: 32px; height: 32px; color: #fff;">$6D3</td>
<td style="text-align: center; border: 0px; background-color: #017301; width: 32px; height: 32px; color: #fff;">$6D4</td>
<td style="text-align: center; border: 0px; background-color: #007502; width: 32px; height: 32px; color: #fff;">$6D5</td>
<td style="text-align: center; border: 0px; background-color: #015369; width: 32px; height: 32px; color: #fff;">$6D6</td>
<td style="text-align: center; border: 0px; background-color: #0134C3; width: 32px; height: 32px; color: #fff;">$6D7</td>
<td style="text-align: center; border: 0px; background-color: #0100FE; width: 32px; height: 32px; color: #fff;">$6D8</td>
<td style="text-align: center; border: 0px; background-color: #3B00FE; width: 32px; height: 32px; color: #fff;">$6D9</td>
<td style="text-align: center; border: 0px; background-color: #8E00D1; width: 32px; height: 32px; color: #fff;">$6DA</td>
<td style="text-align: center; border: 0px; background-color: #A9006F; width: 32px; height: 32px; color: #fff;">$6DB</td>
<td style="text-align: center; border: 0px; background-color: #AE0001; width: 32px; height: 32px; color: #fff;">$6DC</td>
<td style="text-align: center; border: 0px; background-color: #9EA09F; width: 32px; height: 32px; color: #000;">$6DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$6E0</td>
<td style="text-align: center; border: 0px; background-color: #A62300; width: 32px; height: 32px; color: #fff;">$6E1</td>
<td style="text-align: center; border: 0px; background-color: #5E5900; width: 32px; height: 32px; color: #fff;">$6E2</td>
<td style="text-align: center; border: 0px; background-color: #167C00; width: 32px; height: 32px; color: #fff;">$6E3</td>
<td style="text-align: center; border: 0px; background-color: #009100; width: 32px; height: 32px; color: #fff;">$6E4</td>
<td style="text-align: center; border: 0px; background-color: #009201; width: 32px; height: 32px; color: #fff;">$6E5</td>
<td style="text-align: center; border: 0px; background-color: #006F83; width: 32px; height: 32px; color: #fff;">$6E6</td>
<td style="text-align: center; border: 0px; background-color: #0050E1; width: 32px; height: 32px; color: #fff;">$6E7</td>
<td style="text-align: center; border: 0px; background-color: #1917FF; width: 32px; height: 32px; color: #fff;">$6E8</td>
<td style="text-align: center; border: 0px; background-color: #5800FF; width: 32px; height: 32px; color: #fff;">$6E9</td>
<td style="text-align: center; border: 0px; background-color: #AD00EA; width: 32px; height: 32px; color: #fff;">$6EA</td>
<td style="text-align: center; border: 0px; background-color: #C9008E; width: 32px; height: 32px; color: #fff;">$6EB</td>
<td style="text-align: center; border: 0px; background-color: #CD0300; width: 32px; height: 32px; color: #fff;">$6EC</td>
<td style="text-align: center; border: 0px; background-color: #B8BAB9; width: 32px; height: 32px; color: #000;">$6ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6F0</td>
<td style="text-align: center; border: 0px; background-color: #C44000; width: 32px; height: 32px; color: #fff;">$6F1</td>
<td style="text-align: center; border: 0px; background-color: #797400; width: 32px; height: 32px; color: #fff;">$6F2</td>
<td style="text-align: center; border: 0px; background-color: #2F9A01; width: 32px; height: 32px; color: #fff;">$6F3</td>
<td style="text-align: center; border: 0px; background-color: #00AE01; width: 32px; height: 32px; color: #fff;">$6F4</td>
<td style="text-align: center; border: 0px; background-color: #01AD0A; width: 32px; height: 32px; color: #fff;">$6F5</td>
<td style="text-align: center; border: 0px; background-color: #008BA6; width: 32px; height: 32px; color: #fff;">$6F6</td>
<td style="text-align: center; border: 0px; background-color: #016BFD; width: 32px; height: 32px; color: #fff;">$6F7</td>
<td style="text-align: center; border: 0px; background-color: #3134FD; width: 32px; height: 32px; color: #fff;">$6F8</td>
<td style="text-align: center; border: 0px; background-color: #7413FE; width: 32px; height: 32px; color: #fff;">$6F9</td>
<td style="text-align: center; border: 0px; background-color: #CA00FE; width: 32px; height: 32px; color: #fff;">$6FA</td>
<td style="text-align: center; border: 0px; background-color: #E600A9; width: 32px; height: 32px; color: #fff;">$6FB</td>
<td style="text-align: center; border: 0px; background-color: #EB1D05; width: 32px; height: 32px; color: #fff;">$6FC</td>
<td style="text-align: center; border: 0px; background-color: #D3D3D3; width: 32px; height: 32px; color: #000;">$6FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$6FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $7</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$700</td>
<td style="text-align: center; border: 0px; background-color: #D75A00; width: 32px; height: 32px; color: #fff;">$701</td>
<td style="text-align: center; border: 0px; background-color: #918C01; width: 32px; height: 32px; color: #fff;">$702</td>
<td style="text-align: center; border: 0px; background-color: #4EAE00; width: 32px; height: 32px; color: #fff;">$703</td>
<td style="text-align: center; border: 0px; background-color: #09C101; width: 32px; height: 32px; color: #fff;">$704</td>
<td style="text-align: center; border: 0px; background-color: #00C126; width: 32px; height: 32px; color: #fff;">$705</td>
<td style="text-align: center; border: 0px; background-color: #00A0B6; width: 32px; height: 32px; color: #fff;">$706</td>
<td style="text-align: center; border: 0px; background-color: #0884FD; width: 32px; height: 32px; color: #fff;">$707</td>
<td style="text-align: center; border: 0px; background-color: #4D4EFE; width: 32px; height: 32px; color: #fff;">$708</td>
<td style="text-align: center; border: 0px; background-color: #8C30FE; width: 32px; height: 32px; color: #fff;">$709</td>
<td style="text-align: center; border: 0px; background-color: #D91AFE; width: 32px; height: 32px; color: #fff;">$70A</td>
<td style="text-align: center; border: 0px; background-color: #F51BBA; width: 32px; height: 32px; color: #fff;">$70B</td>
<td style="text-align: center; border: 0px; background-color: #F93B2A; width: 32px; height: 32px; color: #fff;">$70C</td>
<td style="text-align: center; border: 0px; background-color: #DFE0DB; width: 32px; height: 32px; color: #000;">$70D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$70E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$70F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0C0C0C; width: 32px; height: 32px; color: #fff;">$710</td>
<td style="text-align: center; border: 0px; background-color: #E97900; width: 32px; height: 32px; color: #000;">$711</td>
<td style="text-align: center; border: 0px; background-color: #A9A601; width: 32px; height: 32px; color: #000;">$712</td>
<td style="text-align: center; border: 0px; background-color: #69C501; width: 32px; height: 32px; color: #000;">$713</td>
<td style="text-align: center; border: 0px; background-color: #26DA00; width: 32px; height: 32px; color: #000;">$714</td>
<td style="text-align: center; border: 0px; background-color: #0AD943; width: 32px; height: 32px; color: #000;">$715</td>
<td style="text-align: center; border: 0px; background-color: #0ABAD3; width: 32px; height: 32px; color: #000;">$716</td>
<td style="text-align: center; border: 0px; background-color: #279BFD; width: 32px; height: 32px; color: #000;">$717</td>
<td style="text-align: center; border: 0px; background-color: #6D68FD; width: 32px; height: 32px; color: #fff;">$718</td>
<td style="text-align: center; border: 0px; background-color: #AB4BFE; width: 32px; height: 32px; color: #fff;">$719</td>
<td style="text-align: center; border: 0px; background-color: #F534FE; width: 32px; height: 32px; color: #000;">$71A</td>
<td style="text-align: center; border: 0px; background-color: #FE37D0; width: 32px; height: 32px; color: #000;">$71B</td>
<td style="text-align: center; border: 0px; background-color: #FF583A; width: 32px; height: 32px; color: #000;">$71C</td>
<td style="text-align: center; border: 0px; background-color: #F8F7FC; width: 32px; height: 32px; color: #000;">$71D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$71E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$71F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #2A2823; width: 32px; height: 32px; color: #fff;">$720</td>
<td style="text-align: center; border: 0px; background-color: #FF9103; width: 32px; height: 32px; color: #000;">$721</td>
<td style="text-align: center; border: 0px; background-color: #C3C000; width: 32px; height: 32px; color: #000;">$722</td>
<td style="text-align: center; border: 0px; background-color: #83DF00; width: 32px; height: 32px; color: #000;">$723</td>
<td style="text-align: center; border: 0px; background-color: #41F203; width: 32px; height: 32px; color: #000;">$724</td>
<td style="text-align: center; border: 0px; background-color: #27F35E; width: 32px; height: 32px; color: #000;">$725</td>
<td style="text-align: center; border: 0px; background-color: #24D4EC; width: 32px; height: 32px; color: #000;">$726</td>
<td style="text-align: center; border: 0px; background-color: #41B7FE; width: 32px; height: 32px; color: #000;">$727</td>
<td style="text-align: center; border: 0px; background-color: #8C80FE; width: 32px; height: 32px; color: #000;">$728</td>
<td style="text-align: center; border: 0px; background-color: #C865FE; width: 32px; height: 32px; color: #000;">$729</td>
<td style="text-align: center; border: 0px; background-color: #FE52FE; width: 32px; height: 32px; color: #000;">$72A</td>
<td style="text-align: center; border: 0px; background-color: #FD55E0; width: 32px; height: 32px; color: #000;">$72B</td>
<td style="text-align: center; border: 0px; background-color: #FF7651; width: 32px; height: 32px; color: #000;">$72C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$72D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$72E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$72F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #424443; width: 32px; height: 32px; color: #fff;">$730</td>
<td style="text-align: center; border: 0px; background-color: #FFAB1D; width: 32px; height: 32px; color: #000;">$731</td>
<td style="text-align: center; border: 0px; background-color: #DDD800; width: 32px; height: 32px; color: #000;">$732</td>
<td style="text-align: center; border: 0px; background-color: #9DF701; width: 32px; height: 32px; color: #000;">$733</td>
<td style="text-align: center; border: 0px; background-color: #5DFE1F; width: 32px; height: 32px; color: #000;">$734</td>
<td style="text-align: center; border: 0px; background-color: #43FF77; width: 32px; height: 32px; color: #000;">$735</td>
<td style="text-align: center; border: 0px; background-color: #40EBFD; width: 32px; height: 32px; color: #000;">$736</td>
<td style="text-align: center; border: 0px; background-color: #5DCFFE; width: 32px; height: 32px; color: #000;">$737</td>
<td style="text-align: center; border: 0px; background-color: #A09EFE; width: 32px; height: 32px; color: #000;">$738</td>
<td style="text-align: center; border: 0px; background-color: #D683FE; width: 32px; height: 32px; color: #000;">$739</td>
<td style="text-align: center; border: 0px; background-color: #FE6EFE; width: 32px; height: 32px; color: #000;">$73A</td>
<td style="text-align: center; border: 0px; background-color: #FF70FE; width: 32px; height: 32px; color: #000;">$73B</td>
<td style="text-align: center; border: 0px; background-color: #FF8E75; width: 32px; height: 32px; color: #000;">$73C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$73D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$73E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$73F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5D63; width: 32px; height: 32px; color: #fff;">$740</td>
<td style="text-align: center; border: 0px; background-color: #00127C; width: 32px; height: 32px; color: #fff;">$741</td>
<td style="text-align: center; border: 0px; background-color: #0000B6; width: 32px; height: 32px; color: #fff;">$742</td>
<td style="text-align: center; border: 0px; background-color: #1D00B2; width: 32px; height: 32px; color: #fff;">$743</td>
<td style="text-align: center; border: 0px; background-color: #4F0080; width: 32px; height: 32px; color: #fff;">$744</td>
<td style="text-align: center; border: 0px; background-color: #630039; width: 32px; height: 32px; color: #fff;">$745</td>
<td style="text-align: center; border: 0px; background-color: #640000; width: 32px; height: 32px; color: #fff;">$746</td>
<td style="text-align: center; border: 0px; background-color: #4F0101; width: 32px; height: 32px; color: #fff;">$747</td>
<td style="text-align: center; border: 0px; background-color: #1C1B01; width: 32px; height: 32px; color: #fff;">$748</td>
<td style="text-align: center; border: 0px; background-color: #003201; width: 32px; height: 32px; color: #fff;">$749</td>
<td style="text-align: center; border: 0px; background-color: #014301; width: 32px; height: 32px; color: #fff;">$74A</td>
<td style="text-align: center; border: 0px; background-color: #004102; width: 32px; height: 32px; color: #fff;">$74B</td>
<td style="text-align: center; border: 0px; background-color: #012837; width: 32px; height: 32px; color: #fff;">$74C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$74D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$74E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$74F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #797979; width: 32px; height: 32px; color: #fff;">$750</td>
<td style="text-align: center; border: 0px; background-color: #012D98; width: 32px; height: 32px; color: #fff;">$751</td>
<td style="text-align: center; border: 0px; background-color: #0609C9; width: 32px; height: 32px; color: #fff;">$752</td>
<td style="text-align: center; border: 0px; background-color: #3700CD; width: 32px; height: 32px; color: #fff;">$753</td>
<td style="text-align: center; border: 0px; background-color: #690099; width: 32px; height: 32px; color: #fff;">$754</td>
<td style="text-align: center; border: 0px; background-color: #7E0054; width: 32px; height: 32px; color: #fff;">$755</td>
<td style="text-align: center; border: 0px; background-color: #7E0001; width: 32px; height: 32px; color: #fff;">$756</td>
<td style="text-align: center; border: 0px; background-color: #691000; width: 32px; height: 32px; color: #fff;">$757</td>
<td style="text-align: center; border: 0px; background-color: #343700; width: 32px; height: 32px; color: #fff;">$758</td>
<td style="text-align: center; border: 0px; background-color: #094C00; width: 32px; height: 32px; color: #fff;">$759</td>
<td style="text-align: center; border: 0px; background-color: #005E00; width: 32px; height: 32px; color: #fff;">$75A</td>
<td style="text-align: center; border: 0px; background-color: #005A00; width: 32px; height: 32px; color: #fff;">$75B</td>
<td style="text-align: center; border: 0px; background-color: #004352; width: 32px; height: 32px; color: #fff;">$75C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$75D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$75E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$75F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$760</td>
<td style="text-align: center; border: 0px; background-color: #0047BB; width: 32px; height: 32px; color: #fff;">$761</td>
<td style="text-align: center; border: 0px; background-color: #2124EC; width: 32px; height: 32px; color: #fff;">$762</td>
<td style="text-align: center; border: 0px; background-color: #530AEA; width: 32px; height: 32px; color: #fff;">$763</td>
<td style="text-align: center; border: 0px; background-color: #8700B7; width: 32px; height: 32px; color: #fff;">$764</td>
<td style="text-align: center; border: 0px; background-color: #9C006E; width: 32px; height: 32px; color: #fff;">$765</td>
<td style="text-align: center; border: 0px; background-color: #9B1400; width: 32px; height: 32px; color: #fff;">$766</td>
<td style="text-align: center; border: 0px; background-color: #872A00; width: 32px; height: 32px; color: #fff;">$767</td>
<td style="text-align: center; border: 0px; background-color: #525100; width: 32px; height: 32px; color: #fff;">$768</td>
<td style="text-align: center; border: 0px; background-color: #246801; width: 32px; height: 32px; color: #fff;">$769</td>
<td style="text-align: center; border: 0px; background-color: #007900; width: 32px; height: 32px; color: #fff;">$76A</td>
<td style="text-align: center; border: 0px; background-color: #007800; width: 32px; height: 32px; color: #fff;">$76B</td>
<td style="text-align: center; border: 0px; background-color: #005F71; width: 32px; height: 32px; color: #fff;">$76C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$76D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$76E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$76F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACABAF; width: 32px; height: 32px; color: #000;">$770</td>
<td style="text-align: center; border: 0px; background-color: #0563D0; width: 32px; height: 32px; color: #fff;">$771</td>
<td style="text-align: center; border: 0px; background-color: #3B3CFE; width: 32px; height: 32px; color: #fff;">$772</td>
<td style="text-align: center; border: 0px; background-color: #7023FE; width: 32px; height: 32px; color: #fff;">$773</td>
<td style="text-align: center; border: 0px; background-color: #A015CF; width: 32px; height: 32px; color: #fff;">$774</td>
<td style="text-align: center; border: 0px; background-color: #B7168B; width: 32px; height: 32px; color: #fff;">$775</td>
<td style="text-align: center; border: 0px; background-color: #B62E19; width: 32px; height: 32px; color: #fff;">$776</td>
<td style="text-align: center; border: 0px; background-color: #A14401; width: 32px; height: 32px; color: #fff;">$777</td>
<td style="text-align: center; border: 0px; background-color: #6D6C00; width: 32px; height: 32px; color: #fff;">$778</td>
<td style="text-align: center; border: 0px; background-color: #3E8300; width: 32px; height: 32px; color: #fff;">$779</td>
<td style="text-align: center; border: 0px; background-color: #069400; width: 32px; height: 32px; color: #fff;">$77A</td>
<td style="text-align: center; border: 0px; background-color: #00931B; width: 32px; height: 32px; color: #fff;">$77B</td>
<td style="text-align: center; border: 0px; background-color: #007B8B; width: 32px; height: 32px; color: #fff;">$77C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$77D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$77E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$77F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C5C5C5; width: 32px; height: 32px; color: #000;">$780</td>
<td style="text-align: center; border: 0px; background-color: #227EEF; width: 32px; height: 32px; color: #fff;">$781</td>
<td style="text-align: center; border: 0px; background-color: #5459FE; width: 32px; height: 32px; color: #fff;">$782</td>
<td style="text-align: center; border: 0px; background-color: #883FFD; width: 32px; height: 32px; color: #fff;">$783</td>
<td style="text-align: center; border: 0px; background-color: #BD31EB; width: 32px; height: 32px; color: #fff;">$784</td>
<td style="text-align: center; border: 0px; background-color: #D031A8; width: 32px; height: 32px; color: #fff;">$785</td>
<td style="text-align: center; border: 0px; background-color: #D34A37; width: 32px; height: 32px; color: #fff;">$786</td>
<td style="text-align: center; border: 0px; background-color: #BC5F00; width: 32px; height: 32px; color: #fff;">$787</td>
<td style="text-align: center; border: 0px; background-color: #888900; width: 32px; height: 32px; color: #fff;">$788</td>
<td style="text-align: center; border: 0px; background-color: #5B9E01; width: 32px; height: 32px; color: #fff;">$789</td>
<td style="text-align: center; border: 0px; background-color: #20AF00; width: 32px; height: 32px; color: #fff;">$78A</td>
<td style="text-align: center; border: 0px; background-color: #0BAF35; width: 32px; height: 32px; color: #fff;">$78B</td>
<td style="text-align: center; border: 0px; background-color: #0895A6; width: 32px; height: 32px; color: #fff;">$78C</td>
<td style="text-align: center; border: 0px; background-color: #090E0C; width: 32px; height: 32px; color: #fff;">$78D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$78E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$78F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DFDFDD; width: 32px; height: 32px; color: #000;">$790</td>
<td style="text-align: center; border: 0px; background-color: #3D9AFE; width: 32px; height: 32px; color: #000;">$791</td>
<td style="text-align: center; border: 0px; background-color: #6F74FE; width: 32px; height: 32px; color: #000;">$792</td>
<td style="text-align: center; border: 0px; background-color: #A55BFE; width: 32px; height: 32px; color: #000;">$793</td>
<td style="text-align: center; border: 0px; background-color: #D74CFE; width: 32px; height: 32px; color: #000;">$794</td>
<td style="text-align: center; border: 0px; background-color: #EC4DC4; width: 32px; height: 32px; color: #000;">$795</td>
<td style="text-align: center; border: 0px; background-color: #EE6554; width: 32px; height: 32px; color: #000;">$796</td>
<td style="text-align: center; border: 0px; background-color: #D67B13; width: 32px; height: 32px; color: #000;">$797</td>
<td style="text-align: center; border: 0px; background-color: #A2A400; width: 32px; height: 32px; color: #000;">$798</td>
<td style="text-align: center; border: 0px; background-color: #75BB00; width: 32px; height: 32px; color: #000;">$799</td>
<td style="text-align: center; border: 0px; background-color: #3BCB0E; width: 32px; height: 32px; color: #000;">$79A</td>
<td style="text-align: center; border: 0px; background-color: #25C951; width: 32px; height: 32px; color: #000;">$79B</td>
<td style="text-align: center; border: 0px; background-color: #24B1C5; width: 32px; height: 32px; color: #000;">$79C</td>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$79D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$79E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$79F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F6F6F6; width: 32px; height: 32px; color: #000;">$7A0</td>
<td style="text-align: center; border: 0px; background-color: #57B4FE; width: 32px; height: 32px; color: #000;">$7A1</td>
<td style="text-align: center; border: 0px; background-color: #898FFC; width: 32px; height: 32px; color: #000;">$7A2</td>
<td style="text-align: center; border: 0px; background-color: #BD75FD; width: 32px; height: 32px; color: #000;">$7A3</td>
<td style="text-align: center; border: 0px; background-color: #F067FF; width: 32px; height: 32px; color: #000;">$7A4</td>
<td style="text-align: center; border: 0px; background-color: #FF68DC; width: 32px; height: 32px; color: #000;">$7A5</td>
<td style="text-align: center; border: 0px; background-color: #FE7F6E; width: 32px; height: 32px; color: #000;">$7A6</td>
<td style="text-align: center; border: 0px; background-color: #F0942D; width: 32px; height: 32px; color: #000;">$7A7</td>
<td style="text-align: center; border: 0px; background-color: #BCBD00; width: 32px; height: 32px; color: #000;">$7A8</td>
<td style="text-align: center; border: 0px; background-color: #8CD400; width: 32px; height: 32px; color: #000;">$7A9</td>
<td style="text-align: center; border: 0px; background-color: #56E329; width: 32px; height: 32px; color: #000;">$7AA</td>
<td style="text-align: center; border: 0px; background-color: #42E16C; width: 32px; height: 32px; color: #000;">$7AB</td>
<td style="text-align: center; border: 0px; background-color: #3ECADB; width: 32px; height: 32px; color: #000;">$7AC</td>
<td style="text-align: center; border: 0px; background-color: #424242; width: 32px; height: 32px; color: #fff;">$7AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$7B0</td>
<td style="text-align: center; border: 0px; background-color: #74CCFF; width: 32px; height: 32px; color: #000;">$7B1</td>
<td style="text-align: center; border: 0px; background-color: #A5A9FF; width: 32px; height: 32px; color: #000;">$7B2</td>
<td style="text-align: center; border: 0px; background-color: #D68FFE; width: 32px; height: 32px; color: #000;">$7B3</td>
<td style="text-align: center; border: 0px; background-color: #FF82FF; width: 32px; height: 32px; color: #000;">$7B4</td>
<td style="text-align: center; border: 0px; background-color: #FE82F2; width: 32px; height: 32px; color: #000;">$7B5</td>
<td style="text-align: center; border: 0px; background-color: #FF998A; width: 32px; height: 32px; color: #000;">$7B6</td>
<td style="text-align: center; border: 0px; background-color: #FEAE4B; width: 32px; height: 32px; color: #000;">$7B7</td>
<td style="text-align: center; border: 0px; background-color: #D3D50F; width: 32px; height: 32px; color: #000;">$7B8</td>
<td style="text-align: center; border: 0px; background-color: #ABEA10; width: 32px; height: 32px; color: #000;">$7B9</td>
<td style="text-align: center; border: 0px; background-color: #73FA48; width: 32px; height: 32px; color: #000;">$7BA</td>
<td style="text-align: center; border: 0px; background-color: #60F888; width: 32px; height: 32px; color: #000;">$7BB</td>
<td style="text-align: center; border: 0px; background-color: #5EE2F1; width: 32px; height: 32px; color: #000;">$7BC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$7BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$7C0</td>
<td style="text-align: center; border: 0px; background-color: #8DE3FE; width: 32px; height: 32px; color: #000;">$7C1</td>
<td style="text-align: center; border: 0px; background-color: #BEC2FF; width: 32px; height: 32px; color: #000;">$7C2</td>
<td style="text-align: center; border: 0px; background-color: #EBA9FF; width: 32px; height: 32px; color: #000;">$7C3</td>
<td style="text-align: center; border: 0px; background-color: #FE9CFE; width: 32px; height: 32px; color: #000;">$7C4</td>
<td style="text-align: center; border: 0px; background-color: #FF9BFE; width: 32px; height: 32px; color: #000;">$7C5</td>
<td style="text-align: center; border: 0px; background-color: #FFB3A5; width: 32px; height: 32px; color: #000;">$7C6</td>
<td style="text-align: center; border: 0px; background-color: #FEC76A; width: 32px; height: 32px; color: #000;">$7C7</td>
<td style="text-align: center; border: 0px; background-color: #E9EC2B; width: 32px; height: 32px; color: #000;">$7C8</td>
<td style="text-align: center; border: 0px; background-color: #BEFE2E; width: 32px; height: 32px; color: #000;">$7C9</td>
<td style="text-align: center; border: 0px; background-color: #8AFE66; width: 32px; height: 32px; color: #000;">$7CA</td>
<td style="text-align: center; border: 0px; background-color: #7BFEA5; width: 32px; height: 32px; color: #000;">$7CB</td>
<td style="text-align: center; border: 0px; background-color: #78F7FE; width: 32px; height: 32px; color: #000;">$7CC</td>
<td style="text-align: center; border: 0px; background-color: #777775; width: 32px; height: 32px; color: #fff;">$7CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$7D0</td>
<td style="text-align: center; border: 0px; background-color: #7F0C00; width: 32px; height: 32px; color: #fff;">$7D1</td>
<td style="text-align: center; border: 0px; background-color: #3F3A01; width: 32px; height: 32px; color: #fff;">$7D2</td>
<td style="text-align: center; border: 0px; background-color: #015A00; width: 32px; height: 32px; color: #fff;">$7D3</td>
<td style="text-align: center; border: 0px; background-color: #006D00; width: 32px; height: 32px; color: #fff;">$7D4</td>
<td style="text-align: center; border: 0px; background-color: #006D00; width: 32px; height: 32px; color: #fff;">$7D5</td>
<td style="text-align: center; border: 0px; background-color: #014D64; width: 32px; height: 32px; color: #fff;">$7D6</td>
<td style="text-align: center; border: 0px; background-color: #0032B7; width: 32px; height: 32px; color: #fff;">$7D7</td>
<td style="text-align: center; border: 0px; background-color: #0100FE; width: 32px; height: 32px; color: #fff;">$7D8</td>
<td style="text-align: center; border: 0px; background-color: #3D00FE; width: 32px; height: 32px; color: #fff;">$7D9</td>
<td style="text-align: center; border: 0px; background-color: #8700BD; width: 32px; height: 32px; color: #fff;">$7DA</td>
<td style="text-align: center; border: 0px; background-color: #9F0068; width: 32px; height: 32px; color: #fff;">$7DB</td>
<td style="text-align: center; border: 0px; background-color: #A20000; width: 32px; height: 32px; color: #fff;">$7DC</td>
<td style="text-align: center; border: 0px; background-color: #929292; width: 32px; height: 32px; color: #000;">$7DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$7E0</td>
<td style="text-align: center; border: 0px; background-color: #9B2501; width: 32px; height: 32px; color: #fff;">$7E1</td>
<td style="text-align: center; border: 0px; background-color: #595400; width: 32px; height: 32px; color: #fff;">$7E2</td>
<td style="text-align: center; border: 0px; background-color: #197500; width: 32px; height: 32px; color: #fff;">$7E3</td>
<td style="text-align: center; border: 0px; background-color: #008500; width: 32px; height: 32px; color: #fff;">$7E4</td>
<td style="text-align: center; border: 0px; background-color: #008700; width: 32px; height: 32px; color: #fff;">$7E5</td>
<td style="text-align: center; border: 0px; background-color: #00697E; width: 32px; height: 32px; color: #fff;">$7E6</td>
<td style="text-align: center; border: 0px; background-color: #004CD2; width: 32px; height: 32px; color: #fff;">$7E7</td>
<td style="text-align: center; border: 0px; background-color: #1C1BFE; width: 32px; height: 32px; color: #fff;">$7E8</td>
<td style="text-align: center; border: 0px; background-color: #5300FE; width: 32px; height: 32px; color: #fff;">$7E9</td>
<td style="text-align: center; border: 0px; background-color: #A000DE; width: 32px; height: 32px; color: #fff;">$7EA</td>
<td style="text-align: center; border: 0px; background-color: #B90083; width: 32px; height: 32px; color: #fff;">$7EB</td>
<td style="text-align: center; border: 0px; background-color: #BF0800; width: 32px; height: 32px; color: #fff;">$7EC</td>
<td style="text-align: center; border: 0px; background-color: #ADADAB; width: 32px; height: 32px; color: #000;">$7ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7F0</td>
<td style="text-align: center; border: 0px; background-color: #B64102; width: 32px; height: 32px; color: #fff;">$7F1</td>
<td style="text-align: center; border: 0px; background-color: #747100; width: 32px; height: 32px; color: #fff;">$7F2</td>
<td style="text-align: center; border: 0px; background-color: #349101; width: 32px; height: 32px; color: #fff;">$7F3</td>
<td style="text-align: center; border: 0px; background-color: #00A300; width: 32px; height: 32px; color: #fff;">$7F4</td>
<td style="text-align: center; border: 0px; background-color: #00A30F; width: 32px; height: 32px; color: #fff;">$7F5</td>
<td style="text-align: center; border: 0px; background-color: #01839A; width: 32px; height: 32px; color: #fff;">$7F6</td>
<td style="text-align: center; border: 0px; background-color: #0068EE; width: 32px; height: 32px; color: #fff;">$7F7</td>
<td style="text-align: center; border: 0px; background-color: #3534FE; width: 32px; height: 32px; color: #fff;">$7F8</td>
<td style="text-align: center; border: 0px; background-color: #7017FD; width: 32px; height: 32px; color: #fff;">$7F9</td>
<td style="text-align: center; border: 0px; background-color: #BC02F2; width: 32px; height: 32px; color: #fff;">$7FA</td>
<td style="text-align: center; border: 0px; background-color: #D6029E; width: 32px; height: 32px; color: #fff;">$7FB</td>
<td style="text-align: center; border: 0px; background-color: #DC210D; width: 32px; height: 32px; color: #fff;">$7FC</td>
<td style="text-align: center; border: 0px; background-color: #C6C6C6; width: 32px; height: 32px; color: #000;">$7FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$7FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $8</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$800</td>
<td style="text-align: center; border: 0px; background-color: #C85C01; width: 32px; height: 32px; color: #fff;">$801</td>
<td style="text-align: center; border: 0px; background-color: #8B8901; width: 32px; height: 32px; color: #fff;">$802</td>
<td style="text-align: center; border: 0px; background-color: #51A401; width: 32px; height: 32px; color: #fff;">$803</td>
<td style="text-align: center; border: 0px; background-color: #17B601; width: 32px; height: 32px; color: #fff;">$804</td>
<td style="text-align: center; border: 0px; background-color: #00B630; width: 32px; height: 32px; color: #fff;">$805</td>
<td style="text-align: center; border: 0px; background-color: #009AAE; width: 32px; height: 32px; color: #fff;">$806</td>
<td style="text-align: center; border: 0px; background-color: #1680FA; width: 32px; height: 32px; color: #fff;">$807</td>
<td style="text-align: center; border: 0px; background-color: #5453FE; width: 32px; height: 32px; color: #fff;">$808</td>
<td style="text-align: center; border: 0px; background-color: #8738FE; width: 32px; height: 32px; color: #fff;">$809</td>
<td style="text-align: center; border: 0px; background-color: #CC25FD; width: 32px; height: 32px; color: #fff;">$80A</td>
<td style="text-align: center; border: 0px; background-color: #E226B1; width: 32px; height: 32px; color: #fff;">$80B</td>
<td style="text-align: center; border: 0px; background-color: #E7422F; width: 32px; height: 32px; color: #fff;">$80C</td>
<td style="text-align: center; border: 0px; background-color: #D4D2D3; width: 32px; height: 32px; color: #000;">$80D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$80E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$80F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1A1A1A; width: 32px; height: 32px; color: #fff;">$810</td>
<td style="text-align: center; border: 0px; background-color: #E27700; width: 32px; height: 32px; color: #000;">$811</td>
<td style="text-align: center; border: 0px; background-color: #A7A201; width: 32px; height: 32px; color: #000;">$812</td>
<td style="text-align: center; border: 0px; background-color: #6CBF00; width: 32px; height: 32px; color: #000;">$813</td>
<td style="text-align: center; border: 0px; background-color: #32D000; width: 32px; height: 32px; color: #000;">$814</td>
<td style="text-align: center; border: 0px; background-color: #1AD048; width: 32px; height: 32px; color: #000;">$815</td>
<td style="text-align: center; border: 0px; background-color: #17B5C1; width: 32px; height: 32px; color: #000;">$816</td>
<td style="text-align: center; border: 0px; background-color: #2F9BFF; width: 32px; height: 32px; color: #000;">$817</td>
<td style="text-align: center; border: 0px; background-color: #6E6CFE; width: 32px; height: 32px; color: #fff;">$818</td>
<td style="text-align: center; border: 0px; background-color: #A153FE; width: 32px; height: 32px; color: #fff;">$819</td>
<td style="text-align: center; border: 0px; background-color: #E140FE; width: 32px; height: 32px; color: #000;">$81A</td>
<td style="text-align: center; border: 0px; background-color: #FB41CA; width: 32px; height: 32px; color: #000;">$81B</td>
<td style="text-align: center; border: 0px; background-color: #FE5D4B; width: 32px; height: 32px; color: #000;">$81C</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$81D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$81E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$81F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #353535; width: 32px; height: 32px; color: #fff;">$820</td>
<td style="text-align: center; border: 0px; background-color: #F89014; width: 32px; height: 32px; color: #000;">$821</td>
<td style="text-align: center; border: 0px; background-color: #BDB900; width: 32px; height: 32px; color: #000;">$822</td>
<td style="text-align: center; border: 0px; background-color: #86D601; width: 32px; height: 32px; color: #000;">$823</td>
<td style="text-align: center; border: 0px; background-color: #4DE715; width: 32px; height: 32px; color: #000;">$824</td>
<td style="text-align: center; border: 0px; background-color: #33E762; width: 32px; height: 32px; color: #000;">$825</td>
<td style="text-align: center; border: 0px; background-color: #32CDDF; width: 32px; height: 32px; color: #000;">$826</td>
<td style="text-align: center; border: 0px; background-color: #4CB3FE; width: 32px; height: 32px; color: #000;">$827</td>
<td style="text-align: center; border: 0px; background-color: #8686FE; width: 32px; height: 32px; color: #000;">$828</td>
<td style="text-align: center; border: 0px; background-color: #BC6DFE; width: 32px; height: 32px; color: #000;">$829</td>
<td style="text-align: center; border: 0px; background-color: #FC59FD; width: 32px; height: 32px; color: #000;">$82A</td>
<td style="text-align: center; border: 0px; background-color: #FE5BE5; width: 32px; height: 32px; color: #000;">$82B</td>
<td style="text-align: center; border: 0px; background-color: #FF7764; width: 32px; height: 32px; color: #000;">$82C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$82D</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$82E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$82F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$830</td>
<td style="text-align: center; border: 0px; background-color: #FEAB33; width: 32px; height: 32px; color: #000;">$831</td>
<td style="text-align: center; border: 0px; background-color: #D7D400; width: 32px; height: 32px; color: #000;">$832</td>
<td style="text-align: center; border: 0px; background-color: #A1EE00; width: 32px; height: 32px; color: #000;">$833</td>
<td style="text-align: center; border: 0px; background-color: #69FE33; width: 32px; height: 32px; color: #000;">$834</td>
<td style="text-align: center; border: 0px; background-color: #53FD80; width: 32px; height: 32px; color: #000;">$835</td>
<td style="text-align: center; border: 0px; background-color: #51E4F8; width: 32px; height: 32px; color: #000;">$836</td>
<td style="text-align: center; border: 0px; background-color: #67CCFF; width: 32px; height: 32px; color: #000;">$837</td>
<td style="text-align: center; border: 0px; background-color: #A79EFE; width: 32px; height: 32px; color: #000;">$838</td>
<td style="text-align: center; border: 0px; background-color: #D986FE; width: 32px; height: 32px; color: #000;">$839</td>
<td style="text-align: center; border: 0px; background-color: #FD75FE; width: 32px; height: 32px; color: #000;">$83A</td>
<td style="text-align: center; border: 0px; background-color: #FE77F8; width: 32px; height: 32px; color: #000;">$83B</td>
<td style="text-align: center; border: 0px; background-color: #FF937C; width: 32px; height: 32px; color: #000;">$83C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$83D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$83E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$83F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6B6A6E; width: 32px; height: 32px; color: #fff;">$840</td>
<td style="text-align: center; border: 0px; background-color: #011597; width: 32px; height: 32px; color: #fff;">$841</td>
<td style="text-align: center; border: 0px; background-color: #0000C8; width: 32px; height: 32px; color: #fff;">$842</td>
<td style="text-align: center; border: 0px; background-color: #2100C9; width: 32px; height: 32px; color: #fff;">$843</td>
<td style="text-align: center; border: 0px; background-color: #590090; width: 32px; height: 32px; color: #fff;">$844</td>
<td style="text-align: center; border: 0px; background-color: #710044; width: 32px; height: 32px; color: #fff;">$845</td>
<td style="text-align: center; border: 0px; background-color: #720000; width: 32px; height: 32px; color: #fff;">$846</td>
<td style="text-align: center; border: 0px; background-color: #590000; width: 32px; height: 32px; color: #fff;">$847</td>
<td style="text-align: center; border: 0px; background-color: #1E1F00; width: 32px; height: 32px; color: #fff;">$848</td>
<td style="text-align: center; border: 0px; background-color: #003700; width: 32px; height: 32px; color: #fff;">$849</td>
<td style="text-align: center; border: 0px; background-color: #004B00; width: 32px; height: 32px; color: #fff;">$84A</td>
<td style="text-align: center; border: 0px; background-color: #004900; width: 32px; height: 32px; color: #fff;">$84B</td>
<td style="text-align: center; border: 0px; background-color: #002E42; width: 32px; height: 32px; color: #fff;">$84C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$84D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$84E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$84F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #868688; width: 32px; height: 32px; color: #000;">$850</td>
<td style="text-align: center; border: 0px; background-color: #0032AA; width: 32px; height: 32px; color: #fff;">$851</td>
<td style="text-align: center; border: 0px; background-color: #0507F0; width: 32px; height: 32px; color: #fff;">$852</td>
<td style="text-align: center; border: 0px; background-color: #3C00E9; width: 32px; height: 32px; color: #fff;">$853</td>
<td style="text-align: center; border: 0px; background-color: #7600AC; width: 32px; height: 32px; color: #fff;">$854</td>
<td style="text-align: center; border: 0px; background-color: #8D005A; width: 32px; height: 32px; color: #fff;">$855</td>
<td style="text-align: center; border: 0px; background-color: #8E0000; width: 32px; height: 32px; color: #fff;">$856</td>
<td style="text-align: center; border: 0px; background-color: #780E00; width: 32px; height: 32px; color: #fff;">$857</td>
<td style="text-align: center; border: 0px; background-color: #3C3B02; width: 32px; height: 32px; color: #fff;">$858</td>
<td style="text-align: center; border: 0px; background-color: #095500; width: 32px; height: 32px; color: #fff;">$859</td>
<td style="text-align: center; border: 0px; background-color: #006A01; width: 32px; height: 32px; color: #fff;">$85A</td>
<td style="text-align: center; border: 0px; background-color: #006800; width: 32px; height: 32px; color: #fff;">$85B</td>
<td style="text-align: center; border: 0px; background-color: #014D5D; width: 32px; height: 32px; color: #fff;">$85C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$85D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$85E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$85F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #A09FA3; width: 32px; height: 32px; color: #000;">$860</td>
<td style="text-align: center; border: 0px; background-color: #014BCD; width: 32px; height: 32px; color: #fff;">$861</td>
<td style="text-align: center; border: 0px; background-color: #1C20FC; width: 32px; height: 32px; color: #fff;">$862</td>
<td style="text-align: center; border: 0px; background-color: #5604FD; width: 32px; height: 32px; color: #fff;">$863</td>
<td style="text-align: center; border: 0px; background-color: #9200CC; width: 32px; height: 32px; color: #fff;">$864</td>
<td style="text-align: center; border: 0px; background-color: #AD007B; width: 32px; height: 32px; color: #fff;">$865</td>
<td style="text-align: center; border: 0px; background-color: #AB0F00; width: 32px; height: 32px; color: #fff;">$866</td>
<td style="text-align: center; border: 0px; background-color: #922801; width: 32px; height: 32px; color: #fff;">$867</td>
<td style="text-align: center; border: 0px; background-color: #575702; width: 32px; height: 32px; color: #fff;">$868</td>
<td style="text-align: center; border: 0px; background-color: #227001; width: 32px; height: 32px; color: #fff;">$869</td>
<td style="text-align: center; border: 0px; background-color: #008401; width: 32px; height: 32px; color: #fff;">$86A</td>
<td style="text-align: center; border: 0px; background-color: #008201; width: 32px; height: 32px; color: #fff;">$86B</td>
<td style="text-align: center; border: 0px; background-color: #00677A; width: 32px; height: 32px; color: #fff;">$86C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$86D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$86E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$86F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B8BAB7; width: 32px; height: 32px; color: #000;">$870</td>
<td style="text-align: center; border: 0px; background-color: #0067E9; width: 32px; height: 32px; color: #fff;">$871</td>
<td style="text-align: center; border: 0px; background-color: #383CFD; width: 32px; height: 32px; color: #fff;">$872</td>
<td style="text-align: center; border: 0px; background-color: #731EFE; width: 32px; height: 32px; color: #fff;">$873</td>
<td style="text-align: center; border: 0px; background-color: #AF0DEE; width: 32px; height: 32px; color: #fff;">$874</td>
<td style="text-align: center; border: 0px; background-color: #C90D95; width: 32px; height: 32px; color: #fff;">$875</td>
<td style="text-align: center; border: 0px; background-color: #C92814; width: 32px; height: 32px; color: #fff;">$876</td>
<td style="text-align: center; border: 0px; background-color: #B14401; width: 32px; height: 32px; color: #fff;">$877</td>
<td style="text-align: center; border: 0px; background-color: #6E7501; width: 32px; height: 32px; color: #fff;">$878</td>
<td style="text-align: center; border: 0px; background-color: #3A8D01; width: 32px; height: 32px; color: #fff;">$879</td>
<td style="text-align: center; border: 0px; background-color: #00A001; width: 32px; height: 32px; color: #fff;">$87A</td>
<td style="text-align: center; border: 0px; background-color: #009F15; width: 32px; height: 32px; color: #fff;">$87B</td>
<td style="text-align: center; border: 0px; background-color: #00829C; width: 32px; height: 32px; color: #fff;">$87C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$87D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$87E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$87F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D3D3D1; width: 32px; height: 32px; color: #000;">$880</td>
<td style="text-align: center; border: 0px; background-color: #1781FC; width: 32px; height: 32px; color: #fff;">$881</td>
<td style="text-align: center; border: 0px; background-color: #5258FE; width: 32px; height: 32px; color: #fff;">$882</td>
<td style="text-align: center; border: 0px; background-color: #8D3AFE; width: 32px; height: 32px; color: #fff;">$883</td>
<td style="text-align: center; border: 0px; background-color: #C928FE; width: 32px; height: 32px; color: #fff;">$884</td>
<td style="text-align: center; border: 0px; background-color: #E327B1; width: 32px; height: 32px; color: #fff;">$885</td>
<td style="text-align: center; border: 0px; background-color: #E4432F; width: 32px; height: 32px; color: #fff;">$886</td>
<td style="text-align: center; border: 0px; background-color: #CB5E01; width: 32px; height: 32px; color: #fff;">$887</td>
<td style="text-align: center; border: 0px; background-color: #8D8E00; width: 32px; height: 32px; color: #fff;">$888</td>
<td style="text-align: center; border: 0px; background-color: #56A902; width: 32px; height: 32px; color: #fff;">$889</td>
<td style="text-align: center; border: 0px; background-color: #13BD01; width: 32px; height: 32px; color: #fff;">$88A</td>
<td style="text-align: center; border: 0px; background-color: #00B931; width: 32px; height: 32px; color: #fff;">$88B</td>
<td style="text-align: center; border: 0px; background-color: #009FAF; width: 32px; height: 32px; color: #fff;">$88C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$88D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$88E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$88F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBED; width: 32px; height: 32px; color: #000;">$890</td>
<td style="text-align: center; border: 0px; background-color: #329CFC; width: 32px; height: 32px; color: #000;">$891</td>
<td style="text-align: center; border: 0px; background-color: #6F71FE; width: 32px; height: 32px; color: #000;">$892</td>
<td style="text-align: center; border: 0px; background-color: #A855FD; width: 32px; height: 32px; color: #000;">$893</td>
<td style="text-align: center; border: 0px; background-color: #E444FE; width: 32px; height: 32px; color: #000;">$894</td>
<td style="text-align: center; border: 0px; background-color: #FD43CA; width: 32px; height: 32px; color: #000;">$895</td>
<td style="text-align: center; border: 0px; background-color: #FC5F4C; width: 32px; height: 32px; color: #000;">$896</td>
<td style="text-align: center; border: 0px; background-color: #E37900; width: 32px; height: 32px; color: #000;">$897</td>
<td style="text-align: center; border: 0px; background-color: #A9A800; width: 32px; height: 32px; color: #000;">$898</td>
<td style="text-align: center; border: 0px; background-color: #70C300; width: 32px; height: 32px; color: #000;">$899</td>
<td style="text-align: center; border: 0px; background-color: #2ED500; width: 32px; height: 32px; color: #000;">$89A</td>
<td style="text-align: center; border: 0px; background-color: #18D449; width: 32px; height: 32px; color: #000;">$89B</td>
<td style="text-align: center; border: 0px; background-color: #13B9C8; width: 32px; height: 32px; color: #000;">$89C</td>
<td style="text-align: center; border: 0px; background-color: #1A191E; width: 32px; height: 32px; color: #fff;">$89D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$89E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$89F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$8A0</td>
<td style="text-align: center; border: 0px; background-color: #4EB5FE; width: 32px; height: 32px; color: #000;">$8A1</td>
<td style="text-align: center; border: 0px; background-color: #898BFF; width: 32px; height: 32px; color: #000;">$8A2</td>
<td style="text-align: center; border: 0px; background-color: #C36EFF; width: 32px; height: 32px; color: #000;">$8A3</td>
<td style="text-align: center; border: 0px; background-color: #FA5EFE; width: 32px; height: 32px; color: #000;">$8A4</td>
<td style="text-align: center; border: 0px; background-color: #FF5EE0; width: 32px; height: 32px; color: #000;">$8A5</td>
<td style="text-align: center; border: 0px; background-color: #FF7965; width: 32px; height: 32px; color: #000;">$8A6</td>
<td style="text-align: center; border: 0px; background-color: #FC9319; width: 32px; height: 32px; color: #000;">$8A7</td>
<td style="text-align: center; border: 0px; background-color: #C4BF00; width: 32px; height: 32px; color: #000;">$8A8</td>
<td style="text-align: center; border: 0px; background-color: #8DDA00; width: 32px; height: 32px; color: #000;">$8A9</td>
<td style="text-align: center; border: 0px; background-color: #4BED13; width: 32px; height: 32px; color: #000;">$8AA</td>
<td style="text-align: center; border: 0px; background-color: #33EC60; width: 32px; height: 32px; color: #000;">$8AB</td>
<td style="text-align: center; border: 0px; background-color: #2FD1DD; width: 32px; height: 32px; color: #000;">$8AC</td>
<td style="text-align: center; border: 0px; background-color: #35362F; width: 32px; height: 32px; color: #fff;">$8AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8AE</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$8AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$8B0</td>
<td style="text-align: center; border: 0px; background-color: #68CEFF; width: 32px; height: 32px; color: #000;">$8B1</td>
<td style="text-align: center; border: 0px; background-color: #A0A7FE; width: 32px; height: 32px; color: #000;">$8B2</td>
<td style="text-align: center; border: 0px; background-color: #D889FE; width: 32px; height: 32px; color: #000;">$8B3</td>
<td style="text-align: center; border: 0px; background-color: #FD7AFE; width: 32px; height: 32px; color: #000;">$8B4</td>
<td style="text-align: center; border: 0px; background-color: #FF79FC; width: 32px; height: 32px; color: #000;">$8B5</td>
<td style="text-align: center; border: 0px; background-color: #FE9381; width: 32px; height: 32px; color: #000;">$8B6</td>
<td style="text-align: center; border: 0px; background-color: #FEAB39; width: 32px; height: 32px; color: #000;">$8B7</td>
<td style="text-align: center; border: 0px; background-color: #D9DA00; width: 32px; height: 32px; color: #000;">$8B8</td>
<td style="text-align: center; border: 0px; background-color: #A8F300; width: 32px; height: 32px; color: #000;">$8B9</td>
<td style="text-align: center; border: 0px; background-color: #67FE36; width: 32px; height: 32px; color: #000;">$8BA</td>
<td style="text-align: center; border: 0px; background-color: #50FE7A; width: 32px; height: 32px; color: #000;">$8BB</td>
<td style="text-align: center; border: 0px; background-color: #4CEAF6; width: 32px; height: 32px; color: #000;">$8BC</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$8BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$8C0</td>
<td style="text-align: center; border: 0px; background-color: #570000; width: 32px; height: 32px; color: #fff;">$8C1</td>
<td style="text-align: center; border: 0px; background-color: #1E1900; width: 32px; height: 32px; color: #fff;">$8C2</td>
<td style="text-align: center; border: 0px; background-color: #013600; width: 32px; height: 32px; color: #fff;">$8C3</td>
<td style="text-align: center; border: 0px; background-color: #004500; width: 32px; height: 32px; color: #fff;">$8C4</td>
<td style="text-align: center; border: 0px; background-color: #004600; width: 32px; height: 32px; color: #fff;">$8C5</td>
<td style="text-align: center; border: 0px; background-color: #012A3E; width: 32px; height: 32px; color: #fff;">$8C6</td>
<td style="text-align: center; border: 0px; background-color: #001185; width: 32px; height: 32px; color: #fff;">$8C7</td>
<td style="text-align: center; border: 0px; background-color: #0000CD; width: 32px; height: 32px; color: #fff;">$8C8</td>
<td style="text-align: center; border: 0px; background-color: #1801CB; width: 32px; height: 32px; color: #fff;">$8C9</td>
<td style="text-align: center; border: 0px; background-color: #59008B; width: 32px; height: 32px; color: #fff;">$8CA</td>
<td style="text-align: center; border: 0px; background-color: #6E003A; width: 32px; height: 32px; color: #fff;">$8CB</td>
<td style="text-align: center; border: 0px; background-color: #710100; width: 32px; height: 32px; color: #fff;">$8CC</td>
<td style="text-align: center; border: 0px; background-color: #6A6B66; width: 32px; height: 32px; color: #fff;">$8CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$8D0</td>
<td style="text-align: center; border: 0px; background-color: #740C00; width: 32px; height: 32px; color: #fff;">$8D1</td>
<td style="text-align: center; border: 0px; background-color: #393500; width: 32px; height: 32px; color: #fff;">$8D2</td>
<td style="text-align: center; border: 0px; background-color: #015200; width: 32px; height: 32px; color: #fff;">$8D3</td>
<td style="text-align: center; border: 0px; background-color: #006300; width: 32px; height: 32px; color: #fff;">$8D4</td>
<td style="text-align: center; border: 0px; background-color: #006300; width: 32px; height: 32px; color: #fff;">$8D5</td>
<td style="text-align: center; border: 0px; background-color: #00465C; width: 32px; height: 32px; color: #fff;">$8D6</td>
<td style="text-align: center; border: 0px; background-color: #002EA2; width: 32px; height: 32px; color: #fff;">$8D7</td>
<td style="text-align: center; border: 0px; background-color: #0301EF; width: 32px; height: 32px; color: #fff;">$8D8</td>
<td style="text-align: center; border: 0px; background-color: #3601ED; width: 32px; height: 32px; color: #fff;">$8D9</td>
<td style="text-align: center; border: 0px; background-color: #7A00AA; width: 32px; height: 32px; color: #fff;">$8DA</td>
<td style="text-align: center; border: 0px; background-color: #910061; width: 32px; height: 32px; color: #fff;">$8DB</td>
<td style="text-align: center; border: 0px; background-color: #920000; width: 32px; height: 32px; color: #fff;">$8DC</td>
<td style="text-align: center; border: 0px; background-color: #868686; width: 32px; height: 32px; color: #000;">$8DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8E0</td>
<td style="text-align: center; border: 0px; background-color: #912600; width: 32px; height: 32px; color: #fff;">$8E1</td>
<td style="text-align: center; border: 0px; background-color: #555200; width: 32px; height: 32px; color: #fff;">$8E2</td>
<td style="text-align: center; border: 0px; background-color: #1B6F00; width: 32px; height: 32px; color: #fff;">$8E3</td>
<td style="text-align: center; border: 0px; background-color: #008000; width: 32px; height: 32px; color: #fff;">$8E4</td>
<td style="text-align: center; border: 0px; background-color: #008000; width: 32px; height: 32px; color: #fff;">$8E5</td>
<td style="text-align: center; border: 0px; background-color: #006376; width: 32px; height: 32px; color: #fff;">$8E6</td>
<td style="text-align: center; border: 0px; background-color: #0148C1; width: 32px; height: 32px; color: #fff;">$8E7</td>
<td style="text-align: center; border: 0px; background-color: #1C1BFF; width: 32px; height: 32px; color: #fff;">$8E8</td>
<td style="text-align: center; border: 0px; background-color: #5200FF; width: 32px; height: 32px; color: #fff;">$8E9</td>
<td style="text-align: center; border: 0px; background-color: #9500C6; width: 32px; height: 32px; color: #fff;">$8EA</td>
<td style="text-align: center; border: 0px; background-color: #AC0078; width: 32px; height: 32px; color: #fff;">$8EB</td>
<td style="text-align: center; border: 0px; background-color: #AF0B00; width: 32px; height: 32px; color: #fff;">$8EC</td>
<td style="text-align: center; border: 0px; background-color: #A0A09D; width: 32px; height: 32px; color: #000;">$8ED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8F0</td>
<td style="text-align: center; border: 0px; background-color: #AE4101; width: 32px; height: 32px; color: #fff;">$8F1</td>
<td style="text-align: center; border: 0px; background-color: #706D01; width: 32px; height: 32px; color: #fff;">$8F2</td>
<td style="text-align: center; border: 0px; background-color: #368800; width: 32px; height: 32px; color: #fff;">$8F3</td>
<td style="text-align: center; border: 0px; background-color: #009B00; width: 32px; height: 32px; color: #fff;">$8F4</td>
<td style="text-align: center; border: 0px; background-color: #009B0F; width: 32px; height: 32px; color: #fff;">$8F5</td>
<td style="text-align: center; border: 0px; background-color: #007F90; width: 32px; height: 32px; color: #fff;">$8F6</td>
<td style="text-align: center; border: 0px; background-color: #0065DE; width: 32px; height: 32px; color: #fff;">$8F7</td>
<td style="text-align: center; border: 0px; background-color: #3736FC; width: 32px; height: 32px; color: #fff;">$8F8</td>
<td style="text-align: center; border: 0px; background-color: #6B1CFE; width: 32px; height: 32px; color: #fff;">$8F9</td>
<td style="text-align: center; border: 0px; background-color: #AE09E1; width: 32px; height: 32px; color: #fff;">$8FA</td>
<td style="text-align: center; border: 0px; background-color: #C60C91; width: 32px; height: 32px; color: #fff;">$8FB</td>
<td style="text-align: center; border: 0px; background-color: #C72614; width: 32px; height: 32px; color: #fff;">$8FC</td>
<td style="text-align: center; border: 0px; background-color: #B9BAB5; width: 32px; height: 32px; color: #000;">$8FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$8FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $9</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0E0B0C; width: 32px; height: 32px; color: #fff;">$900</td>
<td style="text-align: center; border: 0px; background-color: #BB5F00; width: 32px; height: 32px; color: #fff;">$901</td>
<td style="text-align: center; border: 0px; background-color: #878401; width: 32px; height: 32px; color: #fff;">$902</td>
<td style="text-align: center; border: 0px; background-color: #529E00; width: 32px; height: 32px; color: #fff;">$903</td>
<td style="text-align: center; border: 0px; background-color: #1FAB00; width: 32px; height: 32px; color: #fff;">$904</td>
<td style="text-align: center; border: 0px; background-color: #0AAB37; width: 32px; height: 32px; color: #fff;">$905</td>
<td style="text-align: center; border: 0px; background-color: #0A92A5; width: 32px; height: 32px; color: #fff;">$906</td>
<td style="text-align: center; border: 0px; background-color: #207DE6; width: 32px; height: 32px; color: #fff;">$907</td>
<td style="text-align: center; border: 0px; background-color: #5754FD; width: 32px; height: 32px; color: #fff;">$908</td>
<td style="text-align: center; border: 0px; background-color: #833FFE; width: 32px; height: 32px; color: #fff;">$909</td>
<td style="text-align: center; border: 0px; background-color: #BD2DEA; width: 32px; height: 32px; color: #fff;">$90A</td>
<td style="text-align: center; border: 0px; background-color: #D22EA6; width: 32px; height: 32px; color: #fff;">$90B</td>
<td style="text-align: center; border: 0px; background-color: #D44734; width: 32px; height: 32px; color: #fff;">$90C</td>
<td style="text-align: center; border: 0px; background-color: #C6C6C6; width: 32px; height: 32px; color: #000;">$90D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$90E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$90F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #272727; width: 32px; height: 32px; color: #fff;">$910</td>
<td style="text-align: center; border: 0px; background-color: #D57908; width: 32px; height: 32px; color: #000;">$911</td>
<td style="text-align: center; border: 0px; background-color: #A29E01; width: 32px; height: 32px; color: #000;">$912</td>
<td style="text-align: center; border: 0px; background-color: #6EB700; width: 32px; height: 32px; color: #000;">$913</td>
<td style="text-align: center; border: 0px; background-color: #3BC609; width: 32px; height: 32px; color: #000;">$914</td>
<td style="text-align: center; border: 0px; background-color: #27C652; width: 32px; height: 32px; color: #000;">$915</td>
<td style="text-align: center; border: 0px; background-color: #26AEBC; width: 32px; height: 32px; color: #000;">$916</td>
<td style="text-align: center; border: 0px; background-color: #3B98FE; width: 32px; height: 32px; color: #000;">$917</td>
<td style="text-align: center; border: 0px; background-color: #726FFE; width: 32px; height: 32px; color: #000;">$918</td>
<td style="text-align: center; border: 0px; background-color: #9D59FE; width: 32px; height: 32px; color: #000;">$919</td>
<td style="text-align: center; border: 0px; background-color: #D848FF; width: 32px; height: 32px; color: #000;">$91A</td>
<td style="text-align: center; border: 0px; background-color: #ED4ABF; width: 32px; height: 32px; color: #000;">$91B</td>
<td style="text-align: center; border: 0px; background-color: #EF6250; width: 32px; height: 32px; color: #000;">$91C</td>
<td style="text-align: center; border: 0px; background-color: #DEDFE3; width: 32px; height: 32px; color: #000;">$91D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$91E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$91F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #424242; width: 32px; height: 32px; color: #fff;">$920</td>
<td style="text-align: center; border: 0px; background-color: #EE9227; width: 32px; height: 32px; color: #000;">$921</td>
<td style="text-align: center; border: 0px; background-color: #BBB701; width: 32px; height: 32px; color: #000;">$922</td>
<td style="text-align: center; border: 0px; background-color: #8AD000; width: 32px; height: 32px; color: #000;">$923</td>
<td style="text-align: center; border: 0px; background-color: #57DF25; width: 32px; height: 32px; color: #000;">$924</td>
<td style="text-align: center; border: 0px; background-color: #41DE6B; width: 32px; height: 32px; color: #000;">$925</td>
<td style="text-align: center; border: 0px; background-color: #43C6DA; width: 32px; height: 32px; color: #000;">$926</td>
<td style="text-align: center; border: 0px; background-color: #56B2FE; width: 32px; height: 32px; color: #000;">$927</td>
<td style="text-align: center; border: 0px; background-color: #8E87FF; width: 32px; height: 32px; color: #000;">$928</td>
<td style="text-align: center; border: 0px; background-color: #B973FF; width: 32px; height: 32px; color: #000;">$929</td>
<td style="text-align: center; border: 0px; background-color: #F362FE; width: 32px; height: 32px; color: #000;">$92A</td>
<td style="text-align: center; border: 0px; background-color: #FF64D9; width: 32px; height: 32px; color: #000;">$92B</td>
<td style="text-align: center; border: 0px; background-color: #FE7D6A; width: 32px; height: 32px; color: #000;">$92C</td>
<td style="text-align: center; border: 0px; background-color: #F6F6F6; width: 32px; height: 32px; color: #000;">$92D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$92E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$92F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$930</td>
<td style="text-align: center; border: 0px; background-color: #FFAC41; width: 32px; height: 32px; color: #000;">$931</td>
<td style="text-align: center; border: 0px; background-color: #D3CF0E; width: 32px; height: 32px; color: #000;">$932</td>
<td style="text-align: center; border: 0px; background-color: #A0E90C; width: 32px; height: 32px; color: #000;">$933</td>
<td style="text-align: center; border: 0px; background-color: #71F63F; width: 32px; height: 32px; color: #000;">$934</td>
<td style="text-align: center; border: 0px; background-color: #5EF687; width: 32px; height: 32px; color: #000;">$935</td>
<td style="text-align: center; border: 0px; background-color: #5DDFEF; width: 32px; height: 32px; color: #000;">$936</td>
<td style="text-align: center; border: 0px; background-color: #6FC9FE; width: 32px; height: 32px; color: #000;">$937</td>
<td style="text-align: center; border: 0px; background-color: #A9A2FE; width: 32px; height: 32px; color: #000;">$938</td>
<td style="text-align: center; border: 0px; background-color: #D58CFE; width: 32px; height: 32px; color: #000;">$939</td>
<td style="text-align: center; border: 0px; background-color: #FE7EFE; width: 32px; height: 32px; color: #000;">$93A</td>
<td style="text-align: center; border: 0px; background-color: #FF81EC; width: 32px; height: 32px; color: #000;">$93B</td>
<td style="text-align: center; border: 0px; background-color: #FE9883; width: 32px; height: 32px; color: #000;">$93C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$93D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$93E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$93F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #797775; width: 32px; height: 32px; color: #fff;">$940</td>
<td style="text-align: center; border: 0px; background-color: #FFC55D; width: 32px; height: 32px; color: #000;">$941</td>
<td style="text-align: center; border: 0px; background-color: #ECE72B; width: 32px; height: 32px; color: #000;">$942</td>
<td style="text-align: center; border: 0px; background-color: #BBFD2A; width: 32px; height: 32px; color: #000;">$943</td>
<td style="text-align: center; border: 0px; background-color: #8BFF5E; width: 32px; height: 32px; color: #000;">$944</td>
<td style="text-align: center; border: 0px; background-color: #77FE9F; width: 32px; height: 32px; color: #000;">$945</td>
<td style="text-align: center; border: 0px; background-color: #77F6FF; width: 32px; height: 32px; color: #000;">$946</td>
<td style="text-align: center; border: 0px; background-color: #8AE1FD; width: 32px; height: 32px; color: #000;">$947</td>
<td style="text-align: center; border: 0px; background-color: #BEBBFD; width: 32px; height: 32px; color: #000;">$948</td>
<td style="text-align: center; border: 0px; background-color: #E6A7FE; width: 32px; height: 32px; color: #000;">$949</td>
<td style="text-align: center; border: 0px; background-color: #FE96FE; width: 32px; height: 32px; color: #000;">$94A</td>
<td style="text-align: center; border: 0px; background-color: #FF99FD; width: 32px; height: 32px; color: #000;">$94B</td>
<td style="text-align: center; border: 0px; background-color: #FEB19B; width: 32px; height: 32px; color: #000;">$94C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$94D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$94E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$94F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939296; width: 32px; height: 32px; color: #000;">$950</td>
<td style="text-align: center; border: 0px; background-color: #0033C6; width: 32px; height: 32px; color: #fff;">$951</td>
<td style="text-align: center; border: 0px; background-color: #0205FE; width: 32px; height: 32px; color: #fff;">$952</td>
<td style="text-align: center; border: 0px; background-color: #4300FE; width: 32px; height: 32px; color: #fff;">$953</td>
<td style="text-align: center; border: 0px; background-color: #8300C2; width: 32px; height: 32px; color: #fff;">$954</td>
<td style="text-align: center; border: 0px; background-color: #9F0064; width: 32px; height: 32px; color: #fff;">$955</td>
<td style="text-align: center; border: 0px; background-color: #A00000; width: 32px; height: 32px; color: #fff;">$956</td>
<td style="text-align: center; border: 0px; background-color: #840E00; width: 32px; height: 32px; color: #fff;">$957</td>
<td style="text-align: center; border: 0px; background-color: #404000; width: 32px; height: 32px; color: #fff;">$958</td>
<td style="text-align: center; border: 0px; background-color: #065D00; width: 32px; height: 32px; color: #fff;">$959</td>
<td style="text-align: center; border: 0px; background-color: #017300; width: 32px; height: 32px; color: #fff;">$95A</td>
<td style="text-align: center; border: 0px; background-color: #007001; width: 32px; height: 32px; color: #fff;">$95B</td>
<td style="text-align: center; border: 0px; background-color: #005268; width: 32px; height: 32px; color: #fff;">$95C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$95D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$95E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$95F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACACAC; width: 32px; height: 32px; color: #000;">$960</td>
<td style="text-align: center; border: 0px; background-color: #0050DD; width: 32px; height: 32px; color: #fff;">$961</td>
<td style="text-align: center; border: 0px; background-color: #1C1FFD; width: 32px; height: 32px; color: #fff;">$962</td>
<td style="text-align: center; border: 0px; background-color: #5B00FD; width: 32px; height: 32px; color: #fff;">$963</td>
<td style="text-align: center; border: 0px; background-color: #A000E1; width: 32px; height: 32px; color: #fff;">$964</td>
<td style="text-align: center; border: 0px; background-color: #BB0085; width: 32px; height: 32px; color: #fff;">$965</td>
<td style="text-align: center; border: 0px; background-color: #BC0B01; width: 32px; height: 32px; color: #fff;">$966</td>
<td style="text-align: center; border: 0px; background-color: #A02901; width: 32px; height: 32px; color: #fff;">$967</td>
<td style="text-align: center; border: 0px; background-color: #595D00; width: 32px; height: 32px; color: #fff;">$968</td>
<td style="text-align: center; border: 0px; background-color: #1F7800; width: 32px; height: 32px; color: #fff;">$969</td>
<td style="text-align: center; border: 0px; background-color: #008E00; width: 32px; height: 32px; color: #fff;">$96A</td>
<td style="text-align: center; border: 0px; background-color: #008A00; width: 32px; height: 32px; color: #fff;">$96B</td>
<td style="text-align: center; border: 0px; background-color: #016B83; width: 32px; height: 32px; color: #fff;">$96C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$96D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$96E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$96F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C7C4CA; width: 32px; height: 32px; color: #000;">$970</td>
<td style="text-align: center; border: 0px; background-color: #0068FA; width: 32px; height: 32px; color: #fff;">$971</td>
<td style="text-align: center; border: 0px; background-color: #3739FE; width: 32px; height: 32px; color: #fff;">$972</td>
<td style="text-align: center; border: 0px; background-color: #751AFD; width: 32px; height: 32px; color: #fff;">$973</td>
<td style="text-align: center; border: 0px; background-color: #B908F7; width: 32px; height: 32px; color: #fff;">$974</td>
<td style="text-align: center; border: 0px; background-color: #D4089B; width: 32px; height: 32px; color: #fff;">$975</td>
<td style="text-align: center; border: 0px; background-color: #D5270E; width: 32px; height: 32px; color: #fff;">$976</td>
<td style="text-align: center; border: 0px; background-color: #B84301; width: 32px; height: 32px; color: #fff;">$977</td>
<td style="text-align: center; border: 0px; background-color: #767700; width: 32px; height: 32px; color: #fff;">$978</td>
<td style="text-align: center; border: 0px; background-color: #3A9401; width: 32px; height: 32px; color: #fff;">$979</td>
<td style="text-align: center; border: 0px; background-color: #01A900; width: 32px; height: 32px; color: #fff;">$97A</td>
<td style="text-align: center; border: 0px; background-color: #00A70D; width: 32px; height: 32px; color: #fff;">$97B</td>
<td style="text-align: center; border: 0px; background-color: #00889E; width: 32px; height: 32px; color: #fff;">$97C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$97D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$97E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$97F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DDE0DF; width: 32px; height: 32px; color: #000;">$980</td>
<td style="text-align: center; border: 0px; background-color: #0E84FE; width: 32px; height: 32px; color: #fff;">$981</td>
<td style="text-align: center; border: 0px; background-color: #5255FE; width: 32px; height: 32px; color: #fff;">$982</td>
<td style="text-align: center; border: 0px; background-color: #9136FC; width: 32px; height: 32px; color: #fff;">$983</td>
<td style="text-align: center; border: 0px; background-color: #D523FE; width: 32px; height: 32px; color: #fff;">$984</td>
<td style="text-align: center; border: 0px; background-color: #EF22BB; width: 32px; height: 32px; color: #fff;">$985</td>
<td style="text-align: center; border: 0px; background-color: #F0412C; width: 32px; height: 32px; color: #fff;">$986</td>
<td style="text-align: center; border: 0px; background-color: #D35E00; width: 32px; height: 32px; color: #fff;">$987</td>
<td style="text-align: center; border: 0px; background-color: #909200; width: 32px; height: 32px; color: #000;">$988</td>
<td style="text-align: center; border: 0px; background-color: #57B000; width: 32px; height: 32px; color: #000;">$989</td>
<td style="text-align: center; border: 0px; background-color: #0BC701; width: 32px; height: 32px; color: #fff;">$98A</td>
<td style="text-align: center; border: 0px; background-color: #00C327; width: 32px; height: 32px; color: #fff;">$98B</td>
<td style="text-align: center; border: 0px; background-color: #00A4B8; width: 32px; height: 32px; color: #fff;">$98C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$98D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$98E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$98F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F8F7FB; width: 32px; height: 32px; color: #000;">$990</td>
<td style="text-align: center; border: 0px; background-color: #2D9DFF; width: 32px; height: 32px; color: #000;">$991</td>
<td style="text-align: center; border: 0px; background-color: #6B70FE; width: 32px; height: 32px; color: #fff;">$992</td>
<td style="text-align: center; border: 0px; background-color: #AC4FFE; width: 32px; height: 32px; color: #fff;">$993</td>
<td style="text-align: center; border: 0px; background-color: #ED3EFD; width: 32px; height: 32px; color: #000;">$994</td>
<td style="text-align: center; border: 0px; background-color: #FE3ED2; width: 32px; height: 32px; color: #000;">$995</td>
<td style="text-align: center; border: 0px; background-color: #FE5C46; width: 32px; height: 32px; color: #000;">$996</td>
<td style="text-align: center; border: 0px; background-color: #ED7900; width: 32px; height: 32px; color: #000;">$997</td>
<td style="text-align: center; border: 0px; background-color: #ACAC00; width: 32px; height: 32px; color: #000;">$998</td>
<td style="text-align: center; border: 0px; background-color: #70CA00; width: 32px; height: 32px; color: #000;">$999</td>
<td style="text-align: center; border: 0px; background-color: #26E000; width: 32px; height: 32px; color: #000;">$99A</td>
<td style="text-align: center; border: 0px; background-color: #0ADF40; width: 32px; height: 32px; color: #000;">$99B</td>
<td style="text-align: center; border: 0px; background-color: #07BED4; width: 32px; height: 32px; color: #000;">$99C</td>
<td style="text-align: center; border: 0px; background-color: #0B0D0C; width: 32px; height: 32px; color: #fff;">$99D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$99E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$99F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$9A0</td>
<td style="text-align: center; border: 0px; background-color: #48B8FF; width: 32px; height: 32px; color: #000;">$9A1</td>
<td style="text-align: center; border: 0px; background-color: #888AFE; width: 32px; height: 32px; color: #000;">$9A2</td>
<td style="text-align: center; border: 0px; background-color: #C56BFF; width: 32px; height: 32px; color: #000;">$9A3</td>
<td style="text-align: center; border: 0px; background-color: #FF59FF; width: 32px; height: 32px; color: #000;">$9A4</td>
<td style="text-align: center; border: 0px; background-color: #FF58EB; width: 32px; height: 32px; color: #000;">$9A5</td>
<td style="text-align: center; border: 0px; background-color: #FF7761; width: 32px; height: 32px; color: #000;">$9A6</td>
<td style="text-align: center; border: 0px; background-color: #FE9311; width: 32px; height: 32px; color: #000;">$9A7</td>
<td style="text-align: center; border: 0px; background-color: #C5C600; width: 32px; height: 32px; color: #000;">$9A8</td>
<td style="text-align: center; border: 0px; background-color: #8CE000; width: 32px; height: 32px; color: #000;">$9A9</td>
<td style="text-align: center; border: 0px; background-color: #45F50B; width: 32px; height: 32px; color: #000;">$9AA</td>
<td style="text-align: center; border: 0px; background-color: #2BF360; width: 32px; height: 32px; color: #000;">$9AB</td>
<td style="text-align: center; border: 0px; background-color: #26D6EC; width: 32px; height: 32px; color: #000;">$9AC</td>
<td style="text-align: center; border: 0px; background-color: #27282C; width: 32px; height: 32px; color: #fff;">$9AD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9AE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9AF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$9B0</td>
<td style="text-align: center; border: 0px; background-color: #63D0FE; width: 32px; height: 32px; color: #000;">$9B1</td>
<td style="text-align: center; border: 0px; background-color: #A2A4FE; width: 32px; height: 32px; color: #000;">$9B2</td>
<td style="text-align: center; border: 0px; background-color: #DE86FE; width: 32px; height: 32px; color: #000;">$9B3</td>
<td style="text-align: center; border: 0px; background-color: #FF73FE; width: 32px; height: 32px; color: #000;">$9B4</td>
<td style="text-align: center; border: 0px; background-color: #FF73FF; width: 32px; height: 32px; color: #000;">$9B5</td>
<td style="text-align: center; border: 0px; background-color: #FE917C; width: 32px; height: 32px; color: #000;">$9B6</td>
<td style="text-align: center; border: 0px; background-color: #FEAD2D; width: 32px; height: 32px; color: #000;">$9B7</td>
<td style="text-align: center; border: 0px; background-color: #DBDD00; width: 32px; height: 32px; color: #000;">$9B8</td>
<td style="text-align: center; border: 0px; background-color: #A4F901; width: 32px; height: 32px; color: #000;">$9B9</td>
<td style="text-align: center; border: 0px; background-color: #5EFE25; width: 32px; height: 32px; color: #000;">$9BA</td>
<td style="text-align: center; border: 0px; background-color: #44FE78; width: 32px; height: 32px; color: #000;">$9BB</td>
<td style="text-align: center; border: 0px; background-color: #41EEFE; width: 32px; height: 32px; color: #000;">$9BC</td>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$9BD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9BE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9BF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$9C0</td>
<td style="text-align: center; border: 0px; background-color: #4D0000; width: 32px; height: 32px; color: #fff;">$9C1</td>
<td style="text-align: center; border: 0px; background-color: #191700; width: 32px; height: 32px; color: #fff;">$9C2</td>
<td style="text-align: center; border: 0px; background-color: #002F00; width: 32px; height: 32px; color: #fff;">$9C3</td>
<td style="text-align: center; border: 0px; background-color: #003D01; width: 32px; height: 32px; color: #fff;">$9C4</td>
<td style="text-align: center; border: 0px; background-color: #003C00; width: 32px; height: 32px; color: #fff;">$9C5</td>
<td style="text-align: center; border: 0px; background-color: #00243A; width: 32px; height: 32px; color: #fff;">$9C6</td>
<td style="text-align: center; border: 0px; background-color: #000F76; width: 32px; height: 32px; color: #fff;">$9C7</td>
<td style="text-align: center; border: 0px; background-color: #0000B2; width: 32px; height: 32px; color: #fff;">$9C8</td>
<td style="text-align: center; border: 0px; background-color: #1400B2; width: 32px; height: 32px; color: #fff;">$9C9</td>
<td style="text-align: center; border: 0px; background-color: #4E007A; width: 32px; height: 32px; color: #fff;">$9CA</td>
<td style="text-align: center; border: 0px; background-color: #610039; width: 32px; height: 32px; color: #fff;">$9CB</td>
<td style="text-align: center; border: 0px; background-color: #640000; width: 32px; height: 32px; color: #fff;">$9CC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E61; width: 32px; height: 32px; color: #fff;">$9CD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9CE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9CF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$9D0</td>
<td style="text-align: center; border: 0px; background-color: #660C01; width: 32px; height: 32px; color: #fff;">$9D1</td>
<td style="text-align: center; border: 0px; background-color: #343100; width: 32px; height: 32px; color: #fff;">$9D2</td>
<td style="text-align: center; border: 0px; background-color: #024800; width: 32px; height: 32px; color: #fff;">$9D3</td>
<td style="text-align: center; border: 0px; background-color: #005700; width: 32px; height: 32px; color: #fff;">$9D4</td>
<td style="text-align: center; border: 0px; background-color: #005700; width: 32px; height: 32px; color: #fff;">$9D5</td>
<td style="text-align: center; border: 0px; background-color: #004151; width: 32px; height: 32px; color: #fff;">$9D6</td>
<td style="text-align: center; border: 0px; background-color: #002992; width: 32px; height: 32px; color: #fff;">$9D7</td>
<td style="text-align: center; border: 0px; background-color: #0603D6; width: 32px; height: 32px; color: #fff;">$9D8</td>
<td style="text-align: center; border: 0px; background-color: #3100D1; width: 32px; height: 32px; color: #fff;">$9D9</td>
<td style="text-align: center; border: 0px; background-color: #6C0095; width: 32px; height: 32px; color: #fff;">$9DA</td>
<td style="text-align: center; border: 0px; background-color: #800055; width: 32px; height: 32px; color: #fff;">$9DB</td>
<td style="text-align: center; border: 0px; background-color: #820000; width: 32px; height: 32px; color: #fff;">$9DC</td>
<td style="text-align: center; border: 0px; background-color: #787878; width: 32px; height: 32px; color: #fff;">$9DD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9DE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9DF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9E0</td>
<td style="text-align: center; border: 0px; background-color: #852700; width: 32px; height: 32px; color: #fff;">$9E1</td>
<td style="text-align: center; border: 0px; background-color: #514F00; width: 32px; height: 32px; color: #fff;">$9E2</td>
<td style="text-align: center; border: 0px; background-color: #1C6700; width: 32px; height: 32px; color: #fff;">$9E3</td>
<td style="text-align: center; border: 0px; background-color: #007500; width: 32px; height: 32px; color: #fff;">$9E4</td>
<td style="text-align: center; border: 0px; background-color: #007501; width: 32px; height: 32px; color: #fff;">$9E5</td>
<td style="text-align: center; border: 0px; background-color: #005D6F; width: 32px; height: 32px; color: #fff;">$9E6</td>
<td style="text-align: center; border: 0px; background-color: #0046B0; width: 32px; height: 32px; color: #fff;">$9E7</td>
<td style="text-align: center; border: 0px; background-color: #1F1EF5; width: 32px; height: 32px; color: #fff;">$9E8</td>
<td style="text-align: center; border: 0px; background-color: #4E06F5; width: 32px; height: 32px; color: #fff;">$9E9</td>
<td style="text-align: center; border: 0px; background-color: #8A00B7; width: 32px; height: 32px; color: #fff;">$9EA</td>
<td style="text-align: center; border: 0px; background-color: #9E0071; width: 32px; height: 32px; color: #fff;">$9EB</td>
<td style="text-align: center; border: 0px; background-color: #A01100; width: 32px; height: 32px; color: #fff;">$9EC</td>
<td style="text-align: center; border: 0px; background-color: #939391; width: 32px; height: 32px; color: #000;">$9ED</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$9EE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9EF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9F0</td>
<td style="text-align: center; border: 0px; background-color: #A14200; width: 32px; height: 32px; color: #fff;">$9F1</td>
<td style="text-align: center; border: 0px; background-color: #6C6901; width: 32px; height: 32px; color: #fff;">$9F2</td>
<td style="text-align: center; border: 0px; background-color: #388300; width: 32px; height: 32px; color: #fff;">$9F3</td>
<td style="text-align: center; border: 0px; background-color: #059001; width: 32px; height: 32px; color: #fff;">$9F4</td>
<td style="text-align: center; border: 0px; background-color: #00911A; width: 32px; height: 32px; color: #fff;">$9F5</td>
<td style="text-align: center; border: 0px; background-color: #00798A; width: 32px; height: 32px; color: #fff;">$9F6</td>
<td style="text-align: center; border: 0px; background-color: #0460CD; width: 32px; height: 32px; color: #fff;">$9F7</td>
<td style="text-align: center; border: 0px; background-color: #3939FD; width: 32px; height: 32px; color: #fff;">$9F8</td>
<td style="text-align: center; border: 0px; background-color: #6921FE; width: 32px; height: 32px; color: #fff;">$9F9</td>
<td style="text-align: center; border: 0px; background-color: #A411D3; width: 32px; height: 32px; color: #fff;">$9FA</td>
<td style="text-align: center; border: 0px; background-color: #B8128C; width: 32px; height: 32px; color: #fff;">$9FB</td>
<td style="text-align: center; border: 0px; background-color: #BB291A; width: 32px; height: 32px; color: #fff;">$9FC</td>
<td style="text-align: center; border: 0px; background-color: #ACADA8; width: 32px; height: 32px; color: #000;">$9FD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9FE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$9FF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $A</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1A1C1B; width: 32px; height: 32px; color: #fff;">$A00</td>
<td style="text-align: center; border: 0px; background-color: #B16000; width: 32px; height: 32px; color: #fff;">$A01</td>
<td style="text-align: center; border: 0px; background-color: #838001; width: 32px; height: 32px; color: #fff;">$A02</td>
<td style="text-align: center; border: 0px; background-color: #589701; width: 32px; height: 32px; color: #fff;">$A03</td>
<td style="text-align: center; border: 0px; background-color: #2AA300; width: 32px; height: 32px; color: #fff;">$A04</td>
<td style="text-align: center; border: 0px; background-color: #15A43B; width: 32px; height: 32px; color: #fff;">$A05</td>
<td style="text-align: center; border: 0px; background-color: #158F9C; width: 32px; height: 32px; color: #fff;">$A06</td>
<td style="text-align: center; border: 0px; background-color: #2879D5; width: 32px; height: 32px; color: #fff;">$A07</td>
<td style="text-align: center; border: 0px; background-color: #5857FD; width: 32px; height: 32px; color: #fff;">$A08</td>
<td style="text-align: center; border: 0px; background-color: #8043FD; width: 32px; height: 32px; color: #fff;">$A09</td>
<td style="text-align: center; border: 0px; background-color: #B333D5; width: 32px; height: 32px; color: #fff;">$A0A</td>
<td style="text-align: center; border: 0px; background-color: #C635A0; width: 32px; height: 32px; color: #fff;">$A0B</td>
<td style="text-align: center; border: 0px; background-color: #C84A3C; width: 32px; height: 32px; color: #fff;">$A0C</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$A0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #343437; width: 32px; height: 32px; color: #fff;">$A10</td>
<td style="text-align: center; border: 0px; background-color: #C87A18; width: 32px; height: 32px; color: #000;">$A11</td>
<td style="text-align: center; border: 0px; background-color: #9D9801; width: 32px; height: 32px; color: #000;">$A12</td>
<td style="text-align: center; border: 0px; background-color: #70AF00; width: 32px; height: 32px; color: #000;">$A13</td>
<td style="text-align: center; border: 0px; background-color: #44BD18; width: 32px; height: 32px; color: #000;">$A14</td>
<td style="text-align: center; border: 0px; background-color: #30BD55; width: 32px; height: 32px; color: #000;">$A15</td>
<td style="text-align: center; border: 0px; background-color: #2FA7B5; width: 32px; height: 32px; color: #000;">$A16</td>
<td style="text-align: center; border: 0px; background-color: #4495F0; width: 32px; height: 32px; color: #000;">$A17</td>
<td style="text-align: center; border: 0px; background-color: #7370FD; width: 32px; height: 32px; color: #000;">$A18</td>
<td style="text-align: center; border: 0px; background-color: #9A5EFE; width: 32px; height: 32px; color: #000;">$A19</td>
<td style="text-align: center; border: 0px; background-color: #CC4FF2; width: 32px; height: 32px; color: #000;">$A1A</td>
<td style="text-align: center; border: 0px; background-color: #DF51B5; width: 32px; height: 32px; color: #000;">$A1B</td>
<td style="text-align: center; border: 0px; background-color: #E36653; width: 32px; height: 32px; color: #000;">$A1C</td>
<td style="text-align: center; border: 0px; background-color: #D2D3CD; width: 32px; height: 32px; color: #000;">$A1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #51504E; width: 32px; height: 32px; color: #fff;">$A20</td>
<td style="text-align: center; border: 0px; background-color: #E39535; width: 32px; height: 32px; color: #000;">$A21</td>
<td style="text-align: center; border: 0px; background-color: #B5B306; width: 32px; height: 32px; color: #000;">$A22</td>
<td style="text-align: center; border: 0px; background-color: #8BC70B; width: 32px; height: 32px; color: #000;">$A23</td>
<td style="text-align: center; border: 0px; background-color: #60D536; width: 32px; height: 32px; color: #000;">$A24</td>
<td style="text-align: center; border: 0px; background-color: #4CD772; width: 32px; height: 32px; color: #000;">$A25</td>
<td style="text-align: center; border: 0px; background-color: #4FC0D0; width: 32px; height: 32px; color: #000;">$A26</td>
<td style="text-align: center; border: 0px; background-color: #60AEFE; width: 32px; height: 32px; color: #000;">$A27</td>
<td style="text-align: center; border: 0px; background-color: #8F8BFF; width: 32px; height: 32px; color: #000;">$A28</td>
<td style="text-align: center; border: 0px; background-color: #B27BFE; width: 32px; height: 32px; color: #000;">$A29</td>
<td style="text-align: center; border: 0px; background-color: #E56BFE; width: 32px; height: 32px; color: #000;">$A2A</td>
<td style="text-align: center; border: 0px; background-color: #F96CCF; width: 32px; height: 32px; color: #000;">$A2B</td>
<td style="text-align: center; border: 0px; background-color: #FB806F; width: 32px; height: 32px; color: #000;">$A2C</td>
<td style="text-align: center; border: 0px; background-color: #EBEAEF; width: 32px; height: 32px; color: #000;">$A2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6A6A6A; width: 32px; height: 32px; color: #fff;">$A30</td>
<td style="text-align: center; border: 0px; background-color: #F8AD52; width: 32px; height: 32px; color: #000;">$A31</td>
<td style="text-align: center; border: 0px; background-color: #CECB24; width: 32px; height: 32px; color: #000;">$A32</td>
<td style="text-align: center; border: 0px; background-color: #A4E026; width: 32px; height: 32px; color: #000;">$A33</td>
<td style="text-align: center; border: 0px; background-color: #79EC51; width: 32px; height: 32px; color: #000;">$A34</td>
<td style="text-align: center; border: 0px; background-color: #69ED8A; width: 32px; height: 32px; color: #000;">$A35</td>
<td style="text-align: center; border: 0px; background-color: #68D8E4; width: 32px; height: 32px; color: #000;">$A36</td>
<td style="text-align: center; border: 0px; background-color: #79C6FE; width: 32px; height: 32px; color: #000;">$A37</td>
<td style="text-align: center; border: 0px; background-color: #A8A5FF; width: 32px; height: 32px; color: #000;">$A38</td>
<td style="text-align: center; border: 0px; background-color: #C993FF; width: 32px; height: 32px; color: #000;">$A39</td>
<td style="text-align: center; border: 0px; background-color: #FD84FE; width: 32px; height: 32px; color: #000;">$A3A</td>
<td style="text-align: center; border: 0px; background-color: #FF85E8; width: 32px; height: 32px; color: #000;">$A3B</td>
<td style="text-align: center; border: 0px; background-color: #FF9A8C; width: 32px; height: 32px; color: #000;">$A3C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$A3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #858585; width: 32px; height: 32px; color: #000;">$A40</td>
<td style="text-align: center; border: 0px; background-color: #FEC76C; width: 32px; height: 32px; color: #000;">$A41</td>
<td style="text-align: center; border: 0px; background-color: #E6E43F; width: 32px; height: 32px; color: #000;">$A42</td>
<td style="text-align: center; border: 0px; background-color: #BDF942; width: 32px; height: 32px; color: #000;">$A43</td>
<td style="text-align: center; border: 0px; background-color: #93FF6E; width: 32px; height: 32px; color: #000;">$A44</td>
<td style="text-align: center; border: 0px; background-color: #84FFA5; width: 32px; height: 32px; color: #000;">$A45</td>
<td style="text-align: center; border: 0px; background-color: #85F0FF; width: 32px; height: 32px; color: #000;">$A46</td>
<td style="text-align: center; border: 0px; background-color: #95DEFD; width: 32px; height: 32px; color: #000;">$A47</td>
<td style="text-align: center; border: 0px; background-color: #C0BEFE; width: 32px; height: 32px; color: #000;">$A48</td>
<td style="text-align: center; border: 0px; background-color: #E4ADFF; width: 32px; height: 32px; color: #000;">$A49</td>
<td style="text-align: center; border: 0px; background-color: #FE9FFE; width: 32px; height: 32px; color: #000;">$A4A</td>
<td style="text-align: center; border: 0px; background-color: #FEA1FC; width: 32px; height: 32px; color: #000;">$A4B</td>
<td style="text-align: center; border: 0px; background-color: #FEB5A3; width: 32px; height: 32px; color: #000;">$A4C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$A4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #A0A09E; width: 32px; height: 32px; color: #000;">$A50</td>
<td style="text-align: center; border: 0px; background-color: #0035D8; width: 32px; height: 32px; color: #fff;">$A51</td>
<td style="text-align: center; border: 0px; background-color: #0203FD; width: 32px; height: 32px; color: #fff;">$A52</td>
<td style="text-align: center; border: 0px; background-color: #4600FF; width: 32px; height: 32px; color: #fff;">$A53</td>
<td style="text-align: center; border: 0px; background-color: #8E00D2; width: 32px; height: 32px; color: #fff;">$A54</td>
<td style="text-align: center; border: 0px; background-color: #AC006E; width: 32px; height: 32px; color: #fff;">$A55</td>
<td style="text-align: center; border: 0px; background-color: #AC0000; width: 32px; height: 32px; color: #fff;">$A56</td>
<td style="text-align: center; border: 0px; background-color: #8F0D01; width: 32px; height: 32px; color: #fff;">$A57</td>
<td style="text-align: center; border: 0px; background-color: #444301; width: 32px; height: 32px; color: #fff;">$A58</td>
<td style="text-align: center; border: 0px; background-color: #036300; width: 32px; height: 32px; color: #fff;">$A59</td>
<td style="text-align: center; border: 0px; background-color: #007E00; width: 32px; height: 32px; color: #fff;">$A5A</td>
<td style="text-align: center; border: 0px; background-color: #007900; width: 32px; height: 32px; color: #fff;">$A5B</td>
<td style="text-align: center; border: 0px; background-color: #005772; width: 32px; height: 32px; color: #fff;">$A5C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$A5D</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$A5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #BCB9BA; width: 32px; height: 32px; color: #000;">$A60</td>
<td style="text-align: center; border: 0px; background-color: #0051F4; width: 32px; height: 32px; color: #fff;">$A61</td>
<td style="text-align: center; border: 0px; background-color: #191EFC; width: 32px; height: 32px; color: #fff;">$A62</td>
<td style="text-align: center; border: 0px; background-color: #6100FE; width: 32px; height: 32px; color: #fff;">$A63</td>
<td style="text-align: center; border: 0px; background-color: #AC01F0; width: 32px; height: 32px; color: #fff;">$A64</td>
<td style="text-align: center; border: 0px; background-color: #C90089; width: 32px; height: 32px; color: #fff;">$A65</td>
<td style="text-align: center; border: 0px; background-color: #C90800; width: 32px; height: 32px; color: #fff;">$A66</td>
<td style="text-align: center; border: 0px; background-color: #AC2900; width: 32px; height: 32px; color: #fff;">$A67</td>
<td style="text-align: center; border: 0px; background-color: #5E6000; width: 32px; height: 32px; color: #fff;">$A68</td>
<td style="text-align: center; border: 0px; background-color: #1C8101; width: 32px; height: 32px; color: #fff;">$A69</td>
<td style="text-align: center; border: 0px; background-color: #009B00; width: 32px; height: 32px; color: #fff;">$A6A</td>
<td style="text-align: center; border: 0px; background-color: #009800; width: 32px; height: 32px; color: #fff;">$A6B</td>
<td style="text-align: center; border: 0px; background-color: #00758E; width: 32px; height: 32px; color: #fff;">$A6C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$A6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A6E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D3D3D5; width: 32px; height: 32px; color: #000;">$A70</td>
<td style="text-align: center; border: 0px; background-color: #016DFE; width: 32px; height: 32px; color: #fff;">$A71</td>
<td style="text-align: center; border: 0px; background-color: #313BFC; width: 32px; height: 32px; color: #fff;">$A72</td>
<td style="text-align: center; border: 0px; background-color: #7C16FD; width: 32px; height: 32px; color: #fff;">$A73</td>
<td style="text-align: center; border: 0px; background-color: #C801FE; width: 32px; height: 32px; color: #fff;">$A74</td>
<td style="text-align: center; border: 0px; background-color: #E600AA; width: 32px; height: 32px; color: #fff;">$A75</td>
<td style="text-align: center; border: 0px; background-color: #E9220B; width: 32px; height: 32px; color: #fff;">$A76</td>
<td style="text-align: center; border: 0px; background-color: #CB4101; width: 32px; height: 32px; color: #fff;">$A77</td>
<td style="text-align: center; border: 0px; background-color: #7F7A00; width: 32px; height: 32px; color: #fff;">$A78</td>
<td style="text-align: center; border: 0px; background-color: #3C9901; width: 32px; height: 32px; color: #fff;">$A79</td>
<td style="text-align: center; border: 0px; background-color: #01B601; width: 32px; height: 32px; color: #fff;">$A7A</td>
<td style="text-align: center; border: 0px; background-color: #00B403; width: 32px; height: 32px; color: #fff;">$A7B</td>
<td style="text-align: center; border: 0px; background-color: #0091A6; width: 32px; height: 32px; color: #fff;">$A7C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$A7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBECE8; width: 32px; height: 32px; color: #000;">$A80</td>
<td style="text-align: center; border: 0px; background-color: #0688FF; width: 32px; height: 32px; color: #fff;">$A81</td>
<td style="text-align: center; border: 0px; background-color: #5154FE; width: 32px; height: 32px; color: #fff;">$A82</td>
<td style="text-align: center; border: 0px; background-color: #9531FD; width: 32px; height: 32px; color: #fff;">$A83</td>
<td style="text-align: center; border: 0px; background-color: #E11CFF; width: 32px; height: 32px; color: #fff;">$A84</td>
<td style="text-align: center; border: 0px; background-color: #FD1AC1; width: 32px; height: 32px; color: #fff;">$A85</td>
<td style="text-align: center; border: 0px; background-color: #FD3D24; width: 32px; height: 32px; color: #fff;">$A86</td>
<td style="text-align: center; border: 0px; background-color: #E05E01; width: 32px; height: 32px; color: #fff;">$A87</td>
<td style="text-align: center; border: 0px; background-color: #949700; width: 32px; height: 32px; color: #000;">$A88</td>
<td style="text-align: center; border: 0px; background-color: #54B600; width: 32px; height: 32px; color: #000;">$A89</td>
<td style="text-align: center; border: 0px; background-color: #00D001; width: 32px; height: 32px; color: #fff;">$A8A</td>
<td style="text-align: center; border: 0px; background-color: #00CE24; width: 32px; height: 32px; color: #fff;">$A8B</td>
<td style="text-align: center; border: 0px; background-color: #00ABC4; width: 32px; height: 32px; color: #fff;">$A8C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A8D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A8E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$A90</td>
<td style="text-align: center; border: 0px; background-color: #21A0FE; width: 32px; height: 32px; color: #000;">$A91</td>
<td style="text-align: center; border: 0px; background-color: #6B6CFD; width: 32px; height: 32px; color: #fff;">$A92</td>
<td style="text-align: center; border: 0px; background-color: #AF4AFD; width: 32px; height: 32px; color: #fff;">$A93</td>
<td style="text-align: center; border: 0px; background-color: #F936FF; width: 32px; height: 32px; color: #000;">$A94</td>
<td style="text-align: center; border: 0px; background-color: #FF35D6; width: 32px; height: 32px; color: #000;">$A95</td>
<td style="text-align: center; border: 0px; background-color: #FF583D; width: 32px; height: 32px; color: #000;">$A96</td>
<td style="text-align: center; border: 0px; background-color: #F97801; width: 32px; height: 32px; color: #000;">$A97</td>
<td style="text-align: center; border: 0px; background-color: #B0B000; width: 32px; height: 32px; color: #000;">$A98</td>
<td style="text-align: center; border: 0px; background-color: #6BD100; width: 32px; height: 32px; color: #000;">$A99</td>
<td style="text-align: center; border: 0px; background-color: #1CE700; width: 32px; height: 32px; color: #000;">$A9A</td>
<td style="text-align: center; border: 0px; background-color: #00E444; width: 32px; height: 32px; color: #000;">$A9B</td>
<td style="text-align: center; border: 0px; background-color: #00C2DE; width: 32px; height: 32px; color: #000;">$A9C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$A9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$AA0</td>
<td style="text-align: center; border: 0px; background-color: #3EBBFE; width: 32px; height: 32px; color: #000;">$AA1</td>
<td style="text-align: center; border: 0px; background-color: #8688FD; width: 32px; height: 32px; color: #000;">$AA2</td>
<td style="text-align: center; border: 0px; background-color: #C867FE; width: 32px; height: 32px; color: #000;">$AA3</td>
<td style="text-align: center; border: 0px; background-color: #FD51FD; width: 32px; height: 32px; color: #000;">$AA4</td>
<td style="text-align: center; border: 0px; background-color: #FF52F1; width: 32px; height: 32px; color: #000;">$AA5</td>
<td style="text-align: center; border: 0px; background-color: #FF735C; width: 32px; height: 32px; color: #000;">$AA6</td>
<td style="text-align: center; border: 0px; background-color: #FE9302; width: 32px; height: 32px; color: #000;">$AA7</td>
<td style="text-align: center; border: 0px; background-color: #C7CA00; width: 32px; height: 32px; color: #000;">$AA8</td>
<td style="text-align: center; border: 0px; background-color: #8AE801; width: 32px; height: 32px; color: #000;">$AA9</td>
<td style="text-align: center; border: 0px; background-color: #3AFE01; width: 32px; height: 32px; color: #000;">$AAA</td>
<td style="text-align: center; border: 0px; background-color: #1CFD57; width: 32px; height: 32px; color: #000;">$AAB</td>
<td style="text-align: center; border: 0px; background-color: #18DDF0; width: 32px; height: 32px; color: #000;">$AAC</td>
<td style="text-align: center; border: 0px; background-color: #191919; width: 32px; height: 32px; color: #fff;">$AAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AAE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$AB0</td>
<td style="text-align: center; border: 0px; background-color: #2C0000; width: 32px; height: 32px; color: #fff;">$AB1</td>
<td style="text-align: center; border: 0px; background-color: #020100; width: 32px; height: 32px; color: #fff;">$AB2</td>
<td style="text-align: center; border: 0px; background-color: #011400; width: 32px; height: 32px; color: #fff;">$AB3</td>
<td style="text-align: center; border: 0px; background-color: #012000; width: 32px; height: 32px; color: #fff;">$AB4</td>
<td style="text-align: center; border: 0px; background-color: #001F00; width: 32px; height: 32px; color: #fff;">$AB5</td>
<td style="text-align: center; border: 0px; background-color: #000D1A; width: 32px; height: 32px; color: #fff;">$AB6</td>
<td style="text-align: center; border: 0px; background-color: #00004F; width: 32px; height: 32px; color: #fff;">$AB7</td>
<td style="text-align: center; border: 0px; background-color: #010080; width: 32px; height: 32px; color: #fff;">$AB8</td>
<td style="text-align: center; border: 0px; background-color: #010084; width: 32px; height: 32px; color: #fff;">$AB9</td>
<td style="text-align: center; border: 0px; background-color: #2E0051; width: 32px; height: 32px; color: #fff;">$ABA</td>
<td style="text-align: center; border: 0px; background-color: #3C001D; width: 32px; height: 32px; color: #fff;">$ABB</td>
<td style="text-align: center; border: 0px; background-color: #400001; width: 32px; height: 32px; color: #fff;">$ABC</td>
<td style="text-align: center; border: 0px; background-color: #353736; width: 32px; height: 32px; color: #fff;">$ABD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ABE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ABF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$AC0</td>
<td style="text-align: center; border: 0px; background-color: #3F0001; width: 32px; height: 32px; color: #fff;">$AC1</td>
<td style="text-align: center; border: 0px; background-color: #131200; width: 32px; height: 32px; color: #fff;">$AC2</td>
<td style="text-align: center; border: 0px; background-color: #002700; width: 32px; height: 32px; color: #fff;">$AC3</td>
<td style="text-align: center; border: 0px; background-color: #013300; width: 32px; height: 32px; color: #fff;">$AC4</td>
<td style="text-align: center; border: 0px; background-color: #003500; width: 32px; height: 32px; color: #fff;">$AC5</td>
<td style="text-align: center; border: 0px; background-color: #00202D; width: 32px; height: 32px; color: #fff;">$AC6</td>
<td style="text-align: center; border: 0px; background-color: #010D63; width: 32px; height: 32px; color: #fff;">$AC7</td>
<td style="text-align: center; border: 0px; background-color: #000096; width: 32px; height: 32px; color: #fff;">$AC8</td>
<td style="text-align: center; border: 0px; background-color: #100097; width: 32px; height: 32px; color: #fff;">$AC9</td>
<td style="text-align: center; border: 0px; background-color: #420067; width: 32px; height: 32px; color: #fff;">$ACA</td>
<td style="text-align: center; border: 0px; background-color: #52002F; width: 32px; height: 32px; color: #fff;">$ACB</td>
<td style="text-align: center; border: 0px; background-color: #570100; width: 32px; height: 32px; color: #fff;">$ACC</td>
<td style="text-align: center; border: 0px; background-color: #51524D; width: 32px; height: 32px; color: #fff;">$ACD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ACE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ACF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AD0</td>
<td style="text-align: center; border: 0px; background-color: #5C0D01; width: 32px; height: 32px; color: #fff;">$AD1</td>
<td style="text-align: center; border: 0px; background-color: #2F2D00; width: 32px; height: 32px; color: #fff;">$AD2</td>
<td style="text-align: center; border: 0px; background-color: #044200; width: 32px; height: 32px; color: #fff;">$AD3</td>
<td style="text-align: center; border: 0px; background-color: #004E00; width: 32px; height: 32px; color: #fff;">$AD4</td>
<td style="text-align: center; border: 0px; background-color: #004F00; width: 32px; height: 32px; color: #fff;">$AD5</td>
<td style="text-align: center; border: 0px; background-color: #00394A; width: 32px; height: 32px; color: #fff;">$AD6</td>
<td style="text-align: center; border: 0px; background-color: #002880; width: 32px; height: 32px; color: #fff;">$AD7</td>
<td style="text-align: center; border: 0px; background-color: #0904B9; width: 32px; height: 32px; color: #fff;">$AD8</td>
<td style="text-align: center; border: 0px; background-color: #2D01B3; width: 32px; height: 32px; color: #fff;">$AD9</td>
<td style="text-align: center; border: 0px; background-color: #61007D; width: 32px; height: 32px; color: #fff;">$ADA</td>
<td style="text-align: center; border: 0px; background-color: #72004C; width: 32px; height: 32px; color: #fff;">$ADB</td>
<td style="text-align: center; border: 0px; background-color: #720000; width: 32px; height: 32px; color: #fff;">$ADC</td>
<td style="text-align: center; border: 0px; background-color: #696B68; width: 32px; height: 32px; color: #fff;">$ADD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ADE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ADF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AE0</td>
<td style="text-align: center; border: 0px; background-color: #782901; width: 32px; height: 32px; color: #fff;">$AE1</td>
<td style="text-align: center; border: 0px; background-color: #4C4A00; width: 32px; height: 32px; color: #fff;">$AE2</td>
<td style="text-align: center; border: 0px; background-color: #1F5F00; width: 32px; height: 32px; color: #fff;">$AE3</td>
<td style="text-align: center; border: 0px; background-color: #006C00; width: 32px; height: 32px; color: #fff;">$AE4</td>
<td style="text-align: center; border: 0px; background-color: #006D04; width: 32px; height: 32px; color: #fff;">$AE5</td>
<td style="text-align: center; border: 0px; background-color: #005768; width: 32px; height: 32px; color: #fff;">$AE6</td>
<td style="text-align: center; border: 0px; background-color: #00429D; width: 32px; height: 32px; color: #fff;">$AE7</td>
<td style="text-align: center; border: 0px; background-color: #231FD8; width: 32px; height: 32px; color: #fff;">$AE8</td>
<td style="text-align: center; border: 0px; background-color: #4B0BDB; width: 32px; height: 32px; color: #fff;">$AE9</td>
<td style="text-align: center; border: 0px; background-color: #7F00A4; width: 32px; height: 32px; color: #fff;">$AEA</td>
<td style="text-align: center; border: 0px; background-color: #8E0068; width: 32px; height: 32px; color: #fff;">$AEB</td>
<td style="text-align: center; border: 0px; background-color: #901504; width: 32px; height: 32px; color: #fff;">$AEC</td>
<td style="text-align: center; border: 0px; background-color: #858583; width: 32px; height: 32px; color: #000;">$AED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AF0</td>
<td style="text-align: center; border: 0px; background-color: #944401; width: 32px; height: 32px; color: #fff;">$AF1</td>
<td style="text-align: center; border: 0px; background-color: #666400; width: 32px; height: 32px; color: #fff;">$AF2</td>
<td style="text-align: center; border: 0px; background-color: #3C7A00; width: 32px; height: 32px; color: #fff;">$AF3</td>
<td style="text-align: center; border: 0px; background-color: #0E8600; width: 32px; height: 32px; color: #fff;">$AF4</td>
<td style="text-align: center; border: 0px; background-color: #008820; width: 32px; height: 32px; color: #fff;">$AF5</td>
<td style="text-align: center; border: 0px; background-color: #007181; width: 32px; height: 32px; color: #fff;">$AF6</td>
<td style="text-align: center; border: 0px; background-color: #0E5DB9; width: 32px; height: 32px; color: #fff;">$AF7</td>
<td style="text-align: center; border: 0px; background-color: #3F3BF2; width: 32px; height: 32px; color: #fff;">$AF8</td>
<td style="text-align: center; border: 0px; background-color: #6628F2; width: 32px; height: 32px; color: #fff;">$AF9</td>
<td style="text-align: center; border: 0px; background-color: #9719BD; width: 32px; height: 32px; color: #fff;">$AFA</td>
<td style="text-align: center; border: 0px; background-color: #A91A83; width: 32px; height: 32px; color: #fff;">$AFB</td>
<td style="text-align: center; border: 0px; background-color: #AB3021; width: 32px; height: 32px; color: #fff;">$AFC</td>
<td style="text-align: center; border: 0px; background-color: #A0A0A0; width: 32px; height: 32px; color: #000;">$AFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$AFF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $B</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #272821; width: 32px; height: 32px; color: #fff;">$B00</td>
<td style="text-align: center; border: 0px; background-color: #A6610E; width: 32px; height: 32px; color: #fff;">$B01</td>
<td style="text-align: center; border: 0px; background-color: #7E7D00; width: 32px; height: 32px; color: #fff;">$B02</td>
<td style="text-align: center; border: 0px; background-color: #598F00; width: 32px; height: 32px; color: #fff;">$B03</td>
<td style="text-align: center; border: 0px; background-color: #329811; width: 32px; height: 32px; color: #fff;">$B04</td>
<td style="text-align: center; border: 0px; background-color: #239941; width: 32px; height: 32px; color: #fff;">$B05</td>
<td style="text-align: center; border: 0px; background-color: #228896; width: 32px; height: 32px; color: #fff;">$B06</td>
<td style="text-align: center; border: 0px; background-color: #3078C4; width: 32px; height: 32px; color: #fff;">$B07</td>
<td style="text-align: center; border: 0px; background-color: #5D59F4; width: 32px; height: 32px; color: #fff;">$B08</td>
<td style="text-align: center; border: 0px; background-color: #7C49F5; width: 32px; height: 32px; color: #fff;">$B09</td>
<td style="text-align: center; border: 0px; background-color: #A73CC7; width: 32px; height: 32px; color: #fff;">$B0A</td>
<td style="text-align: center; border: 0px; background-color: #B63E96; width: 32px; height: 32px; color: #fff;">$B0B</td>
<td style="text-align: center; border: 0px; background-color: #B85043; width: 32px; height: 32px; color: #fff;">$B0C</td>
<td style="text-align: center; border: 0px; background-color: #ADAEA9; width: 32px; height: 32px; color: #000;">$B0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #434341; width: 32px; height: 32px; color: #fff;">$B10</td>
<td style="text-align: center; border: 0px; background-color: #BF7B28; width: 32px; height: 32px; color: #000;">$B11</td>
<td style="text-align: center; border: 0px; background-color: #989804; width: 32px; height: 32px; color: #000;">$B12</td>
<td style="text-align: center; border: 0px; background-color: #75A801; width: 32px; height: 32px; color: #000;">$B13</td>
<td style="text-align: center; border: 0px; background-color: #4DB329; width: 32px; height: 32px; color: #000;">$B14</td>
<td style="text-align: center; border: 0px; background-color: #3EB45C; width: 32px; height: 32px; color: #000;">$B15</td>
<td style="text-align: center; border: 0px; background-color: #3CA4B1; width: 32px; height: 32px; color: #000;">$B16</td>
<td style="text-align: center; border: 0px; background-color: #4D90DE; width: 32px; height: 32px; color: #000;">$B17</td>
<td style="text-align: center; border: 0px; background-color: #7774FF; width: 32px; height: 32px; color: #000;">$B18</td>
<td style="text-align: center; border: 0px; background-color: #9863FE; width: 32px; height: 32px; color: #000;">$B19</td>
<td style="text-align: center; border: 0px; background-color: #C156E5; width: 32px; height: 32px; color: #000;">$B1A</td>
<td style="text-align: center; border: 0px; background-color: #D059AD; width: 32px; height: 32px; color: #000;">$B1B</td>
<td style="text-align: center; border: 0px; background-color: #D06B5C; width: 32px; height: 32px; color: #000;">$B1C</td>
<td style="text-align: center; border: 0px; background-color: #C5C5C5; width: 32px; height: 32px; color: #000;">$B1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5D5D5B; width: 32px; height: 32px; color: #fff;">$B20</td>
<td style="text-align: center; border: 0px; background-color: #D89745; width: 32px; height: 32px; color: #000;">$B21</td>
<td style="text-align: center; border: 0px; background-color: #B2B220; width: 32px; height: 32px; color: #000;">$B22</td>
<td style="text-align: center; border: 0px; background-color: #90C222; width: 32px; height: 32px; color: #000;">$B23</td>
<td style="text-align: center; border: 0px; background-color: #69CF47; width: 32px; height: 32px; color: #000;">$B24</td>
<td style="text-align: center; border: 0px; background-color: #5ACE7B; width: 32px; height: 32px; color: #000;">$B25</td>
<td style="text-align: center; border: 0px; background-color: #59BECA; width: 32px; height: 32px; color: #000;">$B26</td>
<td style="text-align: center; border: 0px; background-color: #6AACFA; width: 32px; height: 32px; color: #000;">$B27</td>
<td style="text-align: center; border: 0px; background-color: #9090FE; width: 32px; height: 32px; color: #000;">$B28</td>
<td style="text-align: center; border: 0px; background-color: #B180FE; width: 32px; height: 32px; color: #000;">$B29</td>
<td style="text-align: center; border: 0px; background-color: #DB74FA; width: 32px; height: 32px; color: #000;">$B2A</td>
<td style="text-align: center; border: 0px; background-color: #E776C5; width: 32px; height: 32px; color: #000;">$B2B</td>
<td style="text-align: center; border: 0px; background-color: #EA857B; width: 32px; height: 32px; color: #000;">$B2C</td>
<td style="text-align: center; border: 0px; background-color: #DFDEE1; width: 32px; height: 32px; color: #000;">$B2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #7A7778; width: 32px; height: 32px; color: #fff;">$B30</td>
<td style="text-align: center; border: 0px; background-color: #EDAF61; width: 32px; height: 32px; color: #000;">$B31</td>
<td style="text-align: center; border: 0px; background-color: #CCC83F; width: 32px; height: 32px; color: #000;">$B32</td>
<td style="text-align: center; border: 0px; background-color: #A8D93D; width: 32px; height: 32px; color: #000;">$B33</td>
<td style="text-align: center; border: 0px; background-color: #84E561; width: 32px; height: 32px; color: #000;">$B34</td>
<td style="text-align: center; border: 0px; background-color: #75E593; width: 32px; height: 32px; color: #000;">$B35</td>
<td style="text-align: center; border: 0px; background-color: #76D3E2; width: 32px; height: 32px; color: #000;">$B36</td>
<td style="text-align: center; border: 0px; background-color: #81C3FE; width: 32px; height: 32px; color: #000;">$B37</td>
<td style="text-align: center; border: 0px; background-color: #A7AAFE; width: 32px; height: 32px; color: #000;">$B38</td>
<td style="text-align: center; border: 0px; background-color: #C59AFE; width: 32px; height: 32px; color: #000;">$B39</td>
<td style="text-align: center; border: 0px; background-color: #EF8CFF; width: 32px; height: 32px; color: #000;">$B3A</td>
<td style="text-align: center; border: 0px; background-color: #FF8CE6; width: 32px; height: 32px; color: #000;">$B3B</td>
<td style="text-align: center; border: 0px; background-color: #FE9C98; width: 32px; height: 32px; color: #000;">$B3C</td>
<td style="text-align: center; border: 0px; background-color: #F8F8F8; width: 32px; height: 32px; color: #000;">$B3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$B40</td>
<td style="text-align: center; border: 0px; background-color: #FEC97F; width: 32px; height: 32px; color: #000;">$B41</td>
<td style="text-align: center; border: 0px; background-color: #E4E05F; width: 32px; height: 32px; color: #000;">$B42</td>
<td style="text-align: center; border: 0px; background-color: #C2F25D; width: 32px; height: 32px; color: #000;">$B43</td>
<td style="text-align: center; border: 0px; background-color: #A0FA7E; width: 32px; height: 32px; color: #000;">$B44</td>
<td style="text-align: center; border: 0px; background-color: #93FBAE; width: 32px; height: 32px; color: #000;">$B45</td>
<td style="text-align: center; border: 0px; background-color: #92ECF7; width: 32px; height: 32px; color: #000;">$B46</td>
<td style="text-align: center; border: 0px; background-color: #9FDCFF; width: 32px; height: 32px; color: #000;">$B47</td>
<td style="text-align: center; border: 0px; background-color: #C3C2FE; width: 32px; height: 32px; color: #000;">$B48</td>
<td style="text-align: center; border: 0px; background-color: #DFB4FE; width: 32px; height: 32px; color: #000;">$B49</td>
<td style="text-align: center; border: 0px; background-color: #FEA8FE; width: 32px; height: 32px; color: #000;">$B4A</td>
<td style="text-align: center; border: 0px; background-color: #FEA9FB; width: 32px; height: 32px; color: #000;">$B4B</td>
<td style="text-align: center; border: 0px; background-color: #FFB9AD; width: 32px; height: 32px; color: #000;">$B4C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$B4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACACAC; width: 32px; height: 32px; color: #000;">$B50</td>
<td style="text-align: center; border: 0px; background-color: #FEE098; width: 32px; height: 32px; color: #000;">$B51</td>
<td style="text-align: center; border: 0px; background-color: #FBF877; width: 32px; height: 32px; color: #000;">$B52</td>
<td style="text-align: center; border: 0px; background-color: #D9FE76; width: 32px; height: 32px; color: #000;">$B53</td>
<td style="text-align: center; border: 0px; background-color: #BAFE99; width: 32px; height: 32px; color: #000;">$B54</td>
<td style="text-align: center; border: 0px; background-color: #ACFFC9; width: 32px; height: 32px; color: #000;">$B55</td>
<td style="text-align: center; border: 0px; background-color: #AAFFFF; width: 32px; height: 32px; color: #000;">$B56</td>
<td style="text-align: center; border: 0px; background-color: #B8F2FE; width: 32px; height: 32px; color: #000;">$B57</td>
<td style="text-align: center; border: 0px; background-color: #DCDAFF; width: 32px; height: 32px; color: #000;">$B58</td>
<td style="text-align: center; border: 0px; background-color: #F7CCFE; width: 32px; height: 32px; color: #000;">$B59</td>
<td style="text-align: center; border: 0px; background-color: #FEC2FE; width: 32px; height: 32px; color: #000;">$B5A</td>
<td style="text-align: center; border: 0px; background-color: #FEC3FE; width: 32px; height: 32px; color: #000;">$B5B</td>
<td style="text-align: center; border: 0px; background-color: #FFD2C6; width: 32px; height: 32px; color: #000;">$B5C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$B5D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C8C4C8; width: 32px; height: 32px; color: #000;">$B60</td>
<td style="text-align: center; border: 0px; background-color: #0054FF; width: 32px; height: 32px; color: #fff;">$B61</td>
<td style="text-align: center; border: 0px; background-color: #1821FD; width: 32px; height: 32px; color: #fff;">$B62</td>
<td style="text-align: center; border: 0px; background-color: #6500FE; width: 32px; height: 32px; color: #fff;">$B63</td>
<td style="text-align: center; border: 0px; background-color: #B800FF; width: 32px; height: 32px; color: #fff;">$B64</td>
<td style="text-align: center; border: 0px; background-color: #DA0095; width: 32px; height: 32px; color: #fff;">$B65</td>
<td style="text-align: center; border: 0px; background-color: #D90400; width: 32px; height: 32px; color: #fff;">$B66</td>
<td style="text-align: center; border: 0px; background-color: #BB2701; width: 32px; height: 32px; color: #fff;">$B67</td>
<td style="text-align: center; border: 0px; background-color: #636001; width: 32px; height: 32px; color: #fff;">$B68</td>
<td style="text-align: center; border: 0px; background-color: #1E8201; width: 32px; height: 32px; color: #fff;">$B69</td>
<td style="text-align: center; border: 0px; background-color: #00A700; width: 32px; height: 32px; color: #fff;">$B6A</td>
<td style="text-align: center; border: 0px; background-color: #00A101; width: 32px; height: 32px; color: #fff;">$B6B</td>
<td style="text-align: center; border: 0px; background-color: #007C95; width: 32px; height: 32px; color: #fff;">$B6C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$B6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B6E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #E1DEDF; width: 32px; height: 32px; color: #000;">$B70</td>
<td style="text-align: center; border: 0px; background-color: #0171FF; width: 32px; height: 32px; color: #fff;">$B71</td>
<td style="text-align: center; border: 0px; background-color: #303EFE; width: 32px; height: 32px; color: #fff;">$B72</td>
<td style="text-align: center; border: 0px; background-color: #8117FE; width: 32px; height: 32px; color: #fff;">$B73</td>
<td style="text-align: center; border: 0px; background-color: #D600FE; width: 32px; height: 32px; color: #fff;">$B74</td>
<td style="text-align: center; border: 0px; background-color: #F700B4; width: 32px; height: 32px; color: #fff;">$B75</td>
<td style="text-align: center; border: 0px; background-color: #F71D04; width: 32px; height: 32px; color: #fff;">$B76</td>
<td style="text-align: center; border: 0px; background-color: #D74102; width: 32px; height: 32px; color: #fff;">$B77</td>
<td style="text-align: center; border: 0px; background-color: #817A00; width: 32px; height: 32px; color: #fff;">$B78</td>
<td style="text-align: center; border: 0px; background-color: #389D01; width: 32px; height: 32px; color: #fff;">$B79</td>
<td style="text-align: center; border: 0px; background-color: #00C101; width: 32px; height: 32px; color: #fff;">$B7A</td>
<td style="text-align: center; border: 0px; background-color: #00BF03; width: 32px; height: 32px; color: #fff;">$B7B</td>
<td style="text-align: center; border: 0px; background-color: #0099B3; width: 32px; height: 32px; color: #fff;">$B7C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$B7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F7F8FC; width: 32px; height: 32px; color: #000;">$B80</td>
<td style="text-align: center; border: 0px; background-color: #008CFF; width: 32px; height: 32px; color: #fff;">$B81</td>
<td style="text-align: center; border: 0px; background-color: #4A58FE; width: 32px; height: 32px; color: #fff;">$B82</td>
<td style="text-align: center; border: 0px; background-color: #9D30FD; width: 32px; height: 32px; color: #fff;">$B83</td>
<td style="text-align: center; border: 0px; background-color: #EF12FF; width: 32px; height: 32px; color: #fff;">$B84</td>
<td style="text-align: center; border: 0px; background-color: #FF10CD; width: 32px; height: 32px; color: #fff;">$B85</td>
<td style="text-align: center; border: 0px; background-color: #FF391D; width: 32px; height: 32px; color: #fff;">$B86</td>
<td style="text-align: center; border: 0px; background-color: #F15C00; width: 32px; height: 32px; color: #fff;">$B87</td>
<td style="text-align: center; border: 0px; background-color: #9B9500; width: 32px; height: 32px; color: #000;">$B88</td>
<td style="text-align: center; border: 0px; background-color: #51B902; width: 32px; height: 32px; color: #000;">$B89</td>
<td style="text-align: center; border: 0px; background-color: #01DC01; width: 32px; height: 32px; color: #000;">$B8A</td>
<td style="text-align: center; border: 0px; background-color: #00DA1A; width: 32px; height: 32px; color: #000;">$B8B</td>
<td style="text-align: center; border: 0px; background-color: #00B4CC; width: 32px; height: 32px; color: #000;">$B8C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B8D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B8E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$B8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$B90</td>
<td style="text-align: center; border: 0px; background-color: #14A6FE; width: 32px; height: 32px; color: #000;">$B91</td>
<td style="text-align: center; border: 0px; background-color: #6671FE; width: 32px; height: 32px; color: #fff;">$B92</td>
<td style="text-align: center; border: 0px; background-color: #B44BFE; width: 32px; height: 32px; color: #fff;">$B93</td>
<td style="text-align: center; border: 0px; background-color: #FE2CFF; width: 32px; height: 32px; color: #000;">$B94</td>
<td style="text-align: center; border: 0px; background-color: #FF2DE6; width: 32px; height: 32px; color: #000;">$B95</td>
<td style="text-align: center; border: 0px; background-color: #FF543A; width: 32px; height: 32px; color: #000;">$B96</td>
<td style="text-align: center; border: 0px; background-color: #FF7601; width: 32px; height: 32px; color: #000;">$B97</td>
<td style="text-align: center; border: 0px; background-color: #B5AE00; width: 32px; height: 32px; color: #000;">$B98</td>
<td style="text-align: center; border: 0px; background-color: #6CD200; width: 32px; height: 32px; color: #000;">$B99</td>
<td style="text-align: center; border: 0px; background-color: #11F301; width: 32px; height: 32px; color: #000;">$B9A</td>
<td style="text-align: center; border: 0px; background-color: #00EF37; width: 32px; height: 32px; color: #000;">$B9B</td>
<td style="text-align: center; border: 0px; background-color: #01CAE4; width: 32px; height: 32px; color: #000;">$B9C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$B9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$BA0</td>
<td style="text-align: center; border: 0px; background-color: #32BEFC; width: 32px; height: 32px; color: #000;">$BA1</td>
<td style="text-align: center; border: 0px; background-color: #818AFE; width: 32px; height: 32px; color: #000;">$BA2</td>
<td style="text-align: center; border: 0px; background-color: #CD63FE; width: 32px; height: 32px; color: #000;">$BA3</td>
<td style="text-align: center; border: 0px; background-color: #FD4AFE; width: 32px; height: 32px; color: #000;">$BA4</td>
<td style="text-align: center; border: 0px; background-color: #FF48FC; width: 32px; height: 32px; color: #000;">$BA5</td>
<td style="text-align: center; border: 0px; background-color: #FF6F53; width: 32px; height: 32px; color: #000;">$BA6</td>
<td style="text-align: center; border: 0px; background-color: #FF9101; width: 32px; height: 32px; color: #000;">$BA7</td>
<td style="text-align: center; border: 0px; background-color: #CBCA00; width: 32px; height: 32px; color: #000;">$BA8</td>
<td style="text-align: center; border: 0px; background-color: #87EC00; width: 32px; height: 32px; color: #000;">$BA9</td>
<td style="text-align: center; border: 0px; background-color: #2DFF00; width: 32px; height: 32px; color: #000;">$BAA</td>
<td style="text-align: center; border: 0px; background-color: #0DFE52; width: 32px; height: 32px; color: #000;">$BAB</td>
<td style="text-align: center; border: 0px; background-color: #09E3FE; width: 32px; height: 32px; color: #000;">$BAC</td>
<td style="text-align: center; border: 0px; background-color: #0B0B0D; width: 32px; height: 32px; color: #fff;">$BAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BAE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$BB0</td>
<td style="text-align: center; border: 0px; background-color: #220000; width: 32px; height: 32px; color: #fff;">$BB1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BB2</td>
<td style="text-align: center; border: 0px; background-color: #000E00; width: 32px; height: 32px; color: #fff;">$BB3</td>
<td style="text-align: center; border: 0px; background-color: #011901; width: 32px; height: 32px; color: #fff;">$BB4</td>
<td style="text-align: center; border: 0px; background-color: #001800; width: 32px; height: 32px; color: #fff;">$BB5</td>
<td style="text-align: center; border: 0px; background-color: #010713; width: 32px; height: 32px; color: #fff;">$BB6</td>
<td style="text-align: center; border: 0px; background-color: #000040; width: 32px; height: 32px; color: #fff;">$BB7</td>
<td style="text-align: center; border: 0px; background-color: #00006C; width: 32px; height: 32px; color: #fff;">$BB8</td>
<td style="text-align: center; border: 0px; background-color: #00006C; width: 32px; height: 32px; color: #fff;">$BB9</td>
<td style="text-align: center; border: 0px; background-color: #260042; width: 32px; height: 32px; color: #fff;">$BBA</td>
<td style="text-align: center; border: 0px; background-color: #330013; width: 32px; height: 32px; color: #fff;">$BBB</td>
<td style="text-align: center; border: 0px; background-color: #320001; width: 32px; height: 32px; color: #fff;">$BBC</td>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$BBD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BBE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BBF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$BC0</td>
<td style="text-align: center; border: 0px; background-color: #350002; width: 32px; height: 32px; color: #fff;">$BC1</td>
<td style="text-align: center; border: 0px; background-color: #110E00; width: 32px; height: 32px; color: #fff;">$BC2</td>
<td style="text-align: center; border: 0px; background-color: #001F00; width: 32px; height: 32px; color: #fff;">$BC3</td>
<td style="text-align: center; border: 0px; background-color: #002B00; width: 32px; height: 32px; color: #fff;">$BC4</td>
<td style="text-align: center; border: 0px; background-color: #002B00; width: 32px; height: 32px; color: #fff;">$BC5</td>
<td style="text-align: center; border: 0px; background-color: #011828; width: 32px; height: 32px; color: #fff;">$BC6</td>
<td style="text-align: center; border: 0px; background-color: #000A51; width: 32px; height: 32px; color: #fff;">$BC7</td>
<td style="text-align: center; border: 0px; background-color: #00007E; width: 32px; height: 32px; color: #fff;">$BC8</td>
<td style="text-align: center; border: 0px; background-color: #0E007E; width: 32px; height: 32px; color: #fff;">$BC9</td>
<td style="text-align: center; border: 0px; background-color: #350053; width: 32px; height: 32px; color: #fff;">$BCA</td>
<td style="text-align: center; border: 0px; background-color: #420025; width: 32px; height: 32px; color: #fff;">$BCB</td>
<td style="text-align: center; border: 0px; background-color: #460002; width: 32px; height: 32px; color: #fff;">$BCC</td>
<td style="text-align: center; border: 0px; background-color: #41443F; width: 32px; height: 32px; color: #fff;">$BCD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BCE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BCF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BD0</td>
<td style="text-align: center; border: 0px; background-color: #521000; width: 32px; height: 32px; color: #fff;">$BD1</td>
<td style="text-align: center; border: 0px; background-color: #2D2A01; width: 32px; height: 32px; color: #fff;">$BD2</td>
<td style="text-align: center; border: 0px; background-color: #073B00; width: 32px; height: 32px; color: #fff;">$BD3</td>
<td style="text-align: center; border: 0px; background-color: #004800; width: 32px; height: 32px; color: #fff;">$BD4</td>
<td style="text-align: center; border: 0px; background-color: #004800; width: 32px; height: 32px; color: #fff;">$BD5</td>
<td style="text-align: center; border: 0px; background-color: #003442; width: 32px; height: 32px; color: #fff;">$BD6</td>
<td style="text-align: center; border: 0px; background-color: #012770; width: 32px; height: 32px; color: #fff;">$BD7</td>
<td style="text-align: center; border: 0px; background-color: #0C08A2; width: 32px; height: 32px; color: #fff;">$BD8</td>
<td style="text-align: center; border: 0px; background-color: #2A009E; width: 32px; height: 32px; color: #fff;">$BD9</td>
<td style="text-align: center; border: 0px; background-color: #560071; width: 32px; height: 32px; color: #fff;">$BDA</td>
<td style="text-align: center; border: 0px; background-color: #630041; width: 32px; height: 32px; color: #fff;">$BDB</td>
<td style="text-align: center; border: 0px; background-color: #660000; width: 32px; height: 32px; color: #fff;">$BDC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E60; width: 32px; height: 32px; color: #fff;">$BDD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BDE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BDF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BE0</td>
<td style="text-align: center; border: 0px; background-color: #6D2B01; width: 32px; height: 32px; color: #fff;">$BE1</td>
<td style="text-align: center; border: 0px; background-color: #484600; width: 32px; height: 32px; color: #fff;">$BE2</td>
<td style="text-align: center; border: 0px; background-color: #235800; width: 32px; height: 32px; color: #fff;">$BE3</td>
<td style="text-align: center; border: 0px; background-color: #006401; width: 32px; height: 32px; color: #fff;">$BE4</td>
<td style="text-align: center; border: 0px; background-color: #00640E; width: 32px; height: 32px; color: #fff;">$BE5</td>
<td style="text-align: center; border: 0px; background-color: #00525E; width: 32px; height: 32px; color: #fff;">$BE6</td>
<td style="text-align: center; border: 0px; background-color: #00418F; width: 32px; height: 32px; color: #fff;">$BE7</td>
<td style="text-align: center; border: 0px; background-color: #2824BF; width: 32px; height: 32px; color: #fff;">$BE8</td>
<td style="text-align: center; border: 0px; background-color: #4514C2; width: 32px; height: 32px; color: #fff;">$BE9</td>
<td style="text-align: center; border: 0px; background-color: #720694; width: 32px; height: 32px; color: #fff;">$BEA</td>
<td style="text-align: center; border: 0px; background-color: #7F0962; width: 32px; height: 32px; color: #fff;">$BEB</td>
<td style="text-align: center; border: 0px; background-color: #81190C; width: 32px; height: 32px; color: #fff;">$BEC</td>
<td style="text-align: center; border: 0px; background-color: #787878; width: 32px; height: 32px; color: #fff;">$BED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0D0D0D; width: 32px; height: 32px; color: #fff;">$BF0</td>
<td style="text-align: center; border: 0px; background-color: #8D4601; width: 32px; height: 32px; color: #fff;">$BF1</td>
<td style="text-align: center; border: 0px; background-color: #656201; width: 32px; height: 32px; color: #fff;">$BF2</td>
<td style="text-align: center; border: 0px; background-color: #3F7400; width: 32px; height: 32px; color: #fff;">$BF3</td>
<td style="text-align: center; border: 0px; background-color: #167F00; width: 32px; height: 32px; color: #fff;">$BF4</td>
<td style="text-align: center; border: 0px; background-color: #087E26; width: 32px; height: 32px; color: #fff;">$BF5</td>
<td style="text-align: center; border: 0px; background-color: #086E7C; width: 32px; height: 32px; color: #fff;">$BF6</td>
<td style="text-align: center; border: 0px; background-color: #1A5DAB; width: 32px; height: 32px; color: #fff;">$BF7</td>
<td style="text-align: center; border: 0px; background-color: #4140DA; width: 32px; height: 32px; color: #fff;">$BF8</td>
<td style="text-align: center; border: 0px; background-color: #622EDA; width: 32px; height: 32px; color: #fff;">$BF9</td>
<td style="text-align: center; border: 0px; background-color: #8D23AD; width: 32px; height: 32px; color: #fff;">$BFA</td>
<td style="text-align: center; border: 0px; background-color: #9C247E; width: 32px; height: 32px; color: #fff;">$BFB</td>
<td style="text-align: center; border: 0px; background-color: #9B3628; width: 32px; height: 32px; color: #fff;">$BFC</td>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$BFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$BFF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $C</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #363636; width: 32px; height: 32px; color: #fff;">$C00</td>
<td style="text-align: center; border: 0px; background-color: #98621E; width: 32px; height: 32px; color: #fff;">$C01</td>
<td style="text-align: center; border: 0px; background-color: #7A7702; width: 32px; height: 32px; color: #fff;">$C02</td>
<td style="text-align: center; border: 0px; background-color: #5C8600; width: 32px; height: 32px; color: #fff;">$C03</td>
<td style="text-align: center; border: 0px; background-color: #3F8F20; width: 32px; height: 32px; color: #fff;">$C04</td>
<td style="text-align: center; border: 0px; background-color: #309049; width: 32px; height: 32px; color: #fff;">$C05</td>
<td style="text-align: center; border: 0px; background-color: #2F818C; width: 32px; height: 32px; color: #fff;">$C06</td>
<td style="text-align: center; border: 0px; background-color: #3D74B4; width: 32px; height: 32px; color: #fff;">$C07</td>
<td style="text-align: center; border: 0px; background-color: #5D5DD9; width: 32px; height: 32px; color: #fff;">$C08</td>
<td style="text-align: center; border: 0px; background-color: #774ED9; width: 32px; height: 32px; color: #fff;">$C09</td>
<td style="text-align: center; border: 0px; background-color: #9B44B4; width: 32px; height: 32px; color: #fff;">$C0A</td>
<td style="text-align: center; border: 0px; background-color: #A6458B; width: 32px; height: 32px; color: #fff;">$C0B</td>
<td style="text-align: center; border: 0px; background-color: #A7524D; width: 32px; height: 32px; color: #fff;">$C0C</td>
<td style="text-align: center; border: 0px; background-color: #9E9E9E; width: 32px; height: 32px; color: #000;">$C0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #50504E; width: 32px; height: 32px; color: #fff;">$C10</td>
<td style="text-align: center; border: 0px; background-color: #B37D3B; width: 32px; height: 32px; color: #000;">$C11</td>
<td style="text-align: center; border: 0px; background-color: #95931C; width: 32px; height: 32px; color: #000;">$C12</td>
<td style="text-align: center; border: 0px; background-color: #77A21C; width: 32px; height: 32px; color: #000;">$C13</td>
<td style="text-align: center; border: 0px; background-color: #5AAA3B; width: 32px; height: 32px; color: #000;">$C14</td>
<td style="text-align: center; border: 0px; background-color: #4CAC65; width: 32px; height: 32px; color: #000;">$C15</td>
<td style="text-align: center; border: 0px; background-color: #4A9CA7; width: 32px; height: 32px; color: #000;">$C16</td>
<td style="text-align: center; border: 0px; background-color: #5890CB; width: 32px; height: 32px; color: #000;">$C17</td>
<td style="text-align: center; border: 0px; background-color: #7777F1; width: 32px; height: 32px; color: #000;">$C18</td>
<td style="text-align: center; border: 0px; background-color: #936BF0; width: 32px; height: 32px; color: #000;">$C19</td>
<td style="text-align: center; border: 0px; background-color: #B460D1; width: 32px; height: 32px; color: #000;">$C1A</td>
<td style="text-align: center; border: 0px; background-color: #C260AB; width: 32px; height: 32px; color: #000;">$C1B</td>
<td style="text-align: center; border: 0px; background-color: #C27065; width: 32px; height: 32px; color: #000;">$C1C</td>
<td style="text-align: center; border: 0px; background-color: #B7B9BB; width: 32px; height: 32px; color: #000;">$C1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #696967; width: 32px; height: 32px; color: #fff;">$C20</td>
<td style="text-align: center; border: 0px; background-color: #CC9755; width: 32px; height: 32px; color: #000;">$C21</td>
<td style="text-align: center; border: 0px; background-color: #AFAC39; width: 32px; height: 32px; color: #000;">$C22</td>
<td style="text-align: center; border: 0px; background-color: #91BB35; width: 32px; height: 32px; color: #000;">$C23</td>
<td style="text-align: center; border: 0px; background-color: #73C257; width: 32px; height: 32px; color: #000;">$C24</td>
<td style="text-align: center; border: 0px; background-color: #68C37E; width: 32px; height: 32px; color: #000;">$C25</td>
<td style="text-align: center; border: 0px; background-color: #64B7BF; width: 32px; height: 32px; color: #000;">$C26</td>
<td style="text-align: center; border: 0px; background-color: #71A9E4; width: 32px; height: 32px; color: #000;">$C27</td>
<td style="text-align: center; border: 0px; background-color: #8F90FC; width: 32px; height: 32px; color: #000;">$C28</td>
<td style="text-align: center; border: 0px; background-color: #AC84FE; width: 32px; height: 32px; color: #000;">$C29</td>
<td style="text-align: center; border: 0px; background-color: #CE7AE9; width: 32px; height: 32px; color: #000;">$C2A</td>
<td style="text-align: center; border: 0px; background-color: #DA7AC2; width: 32px; height: 32px; color: #000;">$C2B</td>
<td style="text-align: center; border: 0px; background-color: #DC8A7F; width: 32px; height: 32px; color: #000;">$C2C</td>
<td style="text-align: center; border: 0px; background-color: #D0D3D2; width: 32px; height: 32px; color: #000;">$C2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #868485; width: 32px; height: 32px; color: #000;">$C30</td>
<td style="text-align: center; border: 0px; background-color: #E5B16F; width: 32px; height: 32px; color: #000;">$C31</td>
<td style="text-align: center; border: 0px; background-color: #C7C652; width: 32px; height: 32px; color: #000;">$C32</td>
<td style="text-align: center; border: 0px; background-color: #ABD453; width: 32px; height: 32px; color: #000;">$C33</td>
<td style="text-align: center; border: 0px; background-color: #8DDC71; width: 32px; height: 32px; color: #000;">$C34</td>
<td style="text-align: center; border: 0px; background-color: #82DD98; width: 32px; height: 32px; color: #000;">$C35</td>
<td style="text-align: center; border: 0px; background-color: #80CED8; width: 32px; height: 32px; color: #000;">$C36</td>
<td style="text-align: center; border: 0px; background-color: #8EC2FC; width: 32px; height: 32px; color: #000;">$C37</td>
<td style="text-align: center; border: 0px; background-color: #ABACFE; width: 32px; height: 32px; color: #000;">$C38</td>
<td style="text-align: center; border: 0px; background-color: #C69EFF; width: 32px; height: 32px; color: #000;">$C39</td>
<td style="text-align: center; border: 0px; background-color: #E795FC; width: 32px; height: 32px; color: #000;">$C3A</td>
<td style="text-align: center; border: 0px; background-color: #F296D9; width: 32px; height: 32px; color: #000;">$C3B</td>
<td style="text-align: center; border: 0px; background-color: #F4A49B; width: 32px; height: 32px; color: #000;">$C3C</td>
<td style="text-align: center; border: 0px; background-color: #EBEBE9; width: 32px; height: 32px; color: #000;">$C3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9F9F9D; width: 32px; height: 32px; color: #000;">$C40</td>
<td style="text-align: center; border: 0px; background-color: #FBC88F; width: 32px; height: 32px; color: #000;">$C41</td>
<td style="text-align: center; border: 0px; background-color: #DFDD71; width: 32px; height: 32px; color: #000;">$C42</td>
<td style="text-align: center; border: 0px; background-color: #C4EA71; width: 32px; height: 32px; color: #000;">$C43</td>
<td style="text-align: center; border: 0px; background-color: #A8F38E; width: 32px; height: 32px; color: #000;">$C44</td>
<td style="text-align: center; border: 0px; background-color: #9CF3AF; width: 32px; height: 32px; color: #000;">$C45</td>
<td style="text-align: center; border: 0px; background-color: #9DE6EF; width: 32px; height: 32px; color: #000;">$C46</td>
<td style="text-align: center; border: 0px; background-color: #A6DAFE; width: 32px; height: 32px; color: #000;">$C47</td>
<td style="text-align: center; border: 0px; background-color: #C4C6FF; width: 32px; height: 32px; color: #000;">$C48</td>
<td style="text-align: center; border: 0px; background-color: #E0B8FE; width: 32px; height: 32px; color: #000;">$C49</td>
<td style="text-align: center; border: 0px; background-color: #FDAEFF; width: 32px; height: 32px; color: #000;">$C4A</td>
<td style="text-align: center; border: 0px; background-color: #FFAFF3; width: 32px; height: 32px; color: #000;">$C4B</td>
<td style="text-align: center; border: 0px; background-color: #FFBCB3; width: 32px; height: 32px; color: #000;">$C4C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$C4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$C50</td>
<td style="text-align: center; border: 0px; background-color: #FEE2AB; width: 32px; height: 32px; color: #000;">$C51</td>
<td style="text-align: center; border: 0px; background-color: #F7F590; width: 32px; height: 32px; color: #000;">$C52</td>
<td style="text-align: center; border: 0px; background-color: #DDFF91; width: 32px; height: 32px; color: #000;">$C53</td>
<td style="text-align: center; border: 0px; background-color: #C4FDAC; width: 32px; height: 32px; color: #000;">$C54</td>
<td style="text-align: center; border: 0px; background-color: #BAFFD0; width: 32px; height: 32px; color: #000;">$C55</td>
<td style="text-align: center; border: 0px; background-color: #B9FDFE; width: 32px; height: 32px; color: #000;">$C56</td>
<td style="text-align: center; border: 0px; background-color: #C4F1FF; width: 32px; height: 32px; color: #000;">$C57</td>
<td style="text-align: center; border: 0px; background-color: #DEDEFF; width: 32px; height: 32px; color: #000;">$C58</td>
<td style="text-align: center; border: 0px; background-color: #F6D1FE; width: 32px; height: 32px; color: #000;">$C59</td>
<td style="text-align: center; border: 0px; background-color: #FFC8FF; width: 32px; height: 32px; color: #000;">$C5A</td>
<td style="text-align: center; border: 0px; background-color: #FEC9FE; width: 32px; height: 32px; color: #000;">$C5B</td>
<td style="text-align: center; border: 0px; background-color: #FDD7CA; width: 32px; height: 32px; color: #000;">$C5C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$C5D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D3D3D3; width: 32px; height: 32px; color: #000;">$C60</td>
<td style="text-align: center; border: 0px; background-color: #015AFE; width: 32px; height: 32px; color: #fff;">$C61</td>
<td style="text-align: center; border: 0px; background-color: #1727FC; width: 32px; height: 32px; color: #fff;">$C62</td>
<td style="text-align: center; border: 0px; background-color: #6E00FE; width: 32px; height: 32px; color: #fff;">$C63</td>
<td style="text-align: center; border: 0px; background-color: #C500FF; width: 32px; height: 32px; color: #fff;">$C64</td>
<td style="text-align: center; border: 0px; background-color: #EA00A2; width: 32px; height: 32px; color: #fff;">$C65</td>
<td style="text-align: center; border: 0px; background-color: #EB0100; width: 32px; height: 32px; color: #fff;">$C66</td>
<td style="text-align: center; border: 0px; background-color: #C62700; width: 32px; height: 32px; color: #fff;">$C67</td>
<td style="text-align: center; border: 0px; background-color: #676302; width: 32px; height: 32px; color: #fff;">$C68</td>
<td style="text-align: center; border: 0px; background-color: #178A01; width: 32px; height: 32px; color: #fff;">$C69</td>
<td style="text-align: center; border: 0px; background-color: #00B200; width: 32px; height: 32px; color: #fff;">$C6A</td>
<td style="text-align: center; border: 0px; background-color: #01AD02; width: 32px; height: 32px; color: #fff;">$C6B</td>
<td style="text-align: center; border: 0px; background-color: #0084A1; width: 32px; height: 32px; color: #fff;">$C6C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$C6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C6E</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$C6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$C70</td>
<td style="text-align: center; border: 0px; background-color: #0175FE; width: 32px; height: 32px; color: #fff;">$C71</td>
<td style="text-align: center; border: 0px; background-color: #3141FD; width: 32px; height: 32px; color: #fff;">$C72</td>
<td style="text-align: center; border: 0px; background-color: #8616FC; width: 32px; height: 32px; color: #fff;">$C73</td>
<td style="text-align: center; border: 0px; background-color: #E000FE; width: 32px; height: 32px; color: #fff;">$C74</td>
<td style="text-align: center; border: 0px; background-color: #FF00BC; width: 32px; height: 32px; color: #fff;">$C75</td>
<td style="text-align: center; border: 0px; background-color: #FF1B00; width: 32px; height: 32px; color: #fff;">$C76</td>
<td style="text-align: center; border: 0px; background-color: #E04301; width: 32px; height: 32px; color: #fff;">$C77</td>
<td style="text-align: center; border: 0px; background-color: #857B00; width: 32px; height: 32px; color: #fff;">$C78</td>
<td style="text-align: center; border: 0px; background-color: #37A001; width: 32px; height: 32px; color: #fff;">$C79</td>
<td style="text-align: center; border: 0px; background-color: #00C901; width: 32px; height: 32px; color: #fff;">$C7A</td>
<td style="text-align: center; border: 0px; background-color: #00C601; width: 32px; height: 32px; color: #fff;">$C7B</td>
<td style="text-align: center; border: 0px; background-color: #009CBA; width: 32px; height: 32px; color: #fff;">$C7C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$C7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$C80</td>
<td style="text-align: center; border: 0px; background-color: #008DFE; width: 32px; height: 32px; color: #fff;">$C81</td>
<td style="text-align: center; border: 0px; background-color: #4A59FE; width: 32px; height: 32px; color: #fff;">$C82</td>
<td style="text-align: center; border: 0px; background-color: #9F30FD; width: 32px; height: 32px; color: #fff;">$C83</td>
<td style="text-align: center; border: 0px; background-color: #F70EFE; width: 32px; height: 32px; color: #fff;">$C84</td>
<td style="text-align: center; border: 0px; background-color: #FF0CD0; width: 32px; height: 32px; color: #fff;">$C85</td>
<td style="text-align: center; border: 0px; background-color: #FF3615; width: 32px; height: 32px; color: #fff;">$C86</td>
<td style="text-align: center; border: 0px; background-color: #F95B01; width: 32px; height: 32px; color: #fff;">$C87</td>
<td style="text-align: center; border: 0px; background-color: #9E9701; width: 32px; height: 32px; color: #000;">$C88</td>
<td style="text-align: center; border: 0px; background-color: #4FBC00; width: 32px; height: 32px; color: #000;">$C89</td>
<td style="text-align: center; border: 0px; background-color: #01E401; width: 32px; height: 32px; color: #000;">$C8A</td>
<td style="text-align: center; border: 0px; background-color: #00E01A; width: 32px; height: 32px; color: #000;">$C8B</td>
<td style="text-align: center; border: 0px; background-color: #00B6D6; width: 32px; height: 32px; color: #000;">$C8C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$C8D</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$C8E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$C90</td>
<td style="text-align: center; border: 0px; background-color: #0FA9FE; width: 32px; height: 32px; color: #000;">$C91</td>
<td style="text-align: center; border: 0px; background-color: #6574FE; width: 32px; height: 32px; color: #fff;">$C92</td>
<td style="text-align: center; border: 0px; background-color: #B94AFE; width: 32px; height: 32px; color: #fff;">$C93</td>
<td style="text-align: center; border: 0px; background-color: #FE27FE; width: 32px; height: 32px; color: #fff;">$C94</td>
<td style="text-align: center; border: 0px; background-color: #FF29ED; width: 32px; height: 32px; color: #fff;">$C95</td>
<td style="text-align: center; border: 0px; background-color: #FE4F37; width: 32px; height: 32px; color: #000;">$C96</td>
<td style="text-align: center; border: 0px; background-color: #FE7600; width: 32px; height: 32px; color: #000;">$C97</td>
<td style="text-align: center; border: 0px; background-color: #BAAE00; width: 32px; height: 32px; color: #000;">$C98</td>
<td style="text-align: center; border: 0px; background-color: #6CD500; width: 32px; height: 32px; color: #000;">$C99</td>
<td style="text-align: center; border: 0px; background-color: #07FD00; width: 32px; height: 32px; color: #000;">$C9A</td>
<td style="text-align: center; border: 0px; background-color: #00FB2C; width: 32px; height: 32px; color: #000;">$C9B</td>
<td style="text-align: center; border: 0px; background-color: #00D3EE; width: 32px; height: 32px; color: #000;">$C9C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$C9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$CA0</td>
<td style="text-align: center; border: 0px; background-color: #160000; width: 32px; height: 32px; color: #fff;">$CA1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CA2</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$CA3</td>
<td style="text-align: center; border: 0px; background-color: #000E01; width: 32px; height: 32px; color: #fff;">$CA4</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$CA5</td>
<td style="text-align: center; border: 0px; background-color: #01010D; width: 32px; height: 32px; color: #fff;">$CA6</td>
<td style="text-align: center; border: 0px; background-color: #01002A; width: 32px; height: 32px; color: #fff;">$CA7</td>
<td style="text-align: center; border: 0px; background-color: #00014C; width: 32px; height: 32px; color: #fff;">$CA8</td>
<td style="text-align: center; border: 0px; background-color: #00004E; width: 32px; height: 32px; color: #fff;">$CA9</td>
<td style="text-align: center; border: 0px; background-color: #16002B; width: 32px; height: 32px; color: #fff;">$CAA</td>
<td style="text-align: center; border: 0px; background-color: #20000D; width: 32px; height: 32px; color: #fff;">$CAB</td>
<td style="text-align: center; border: 0px; background-color: #1F0000; width: 32px; height: 32px; color: #fff;">$CAC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CAE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$CB0</td>
<td style="text-align: center; border: 0px; background-color: #170000; width: 32px; height: 32px; color: #fff;">$CB1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CB2</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$CB3</td>
<td style="text-align: center; border: 0px; background-color: #001000; width: 32px; height: 32px; color: #fff;">$CB4</td>
<td style="text-align: center; border: 0px; background-color: #000F02; width: 32px; height: 32px; color: #fff;">$CB5</td>
<td style="text-align: center; border: 0px; background-color: #00020E; width: 32px; height: 32px; color: #fff;">$CB6</td>
<td style="text-align: center; border: 0px; background-color: #01002E; width: 32px; height: 32px; color: #fff;">$CB7</td>
<td style="text-align: center; border: 0px; background-color: #000052; width: 32px; height: 32px; color: #fff;">$CB8</td>
<td style="text-align: center; border: 0px; background-color: #000050; width: 32px; height: 32px; color: #fff;">$CB9</td>
<td style="text-align: center; border: 0px; background-color: #1B0033; width: 32px; height: 32px; color: #fff;">$CBA</td>
<td style="text-align: center; border: 0px; background-color: #24000C; width: 32px; height: 32px; color: #fff;">$CBB</td>
<td style="text-align: center; border: 0px; background-color: #240000; width: 32px; height: 32px; color: #fff;">$CBC</td>
<td style="text-align: center; border: 0px; background-color: #1B1B19; width: 32px; height: 32px; color: #fff;">$CBD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CBE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CBF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CC0</td>
<td style="text-align: center; border: 0px; background-color: #290001; width: 32px; height: 32px; color: #fff;">$CC1</td>
<td style="text-align: center; border: 0px; background-color: #0D0B00; width: 32px; height: 32px; color: #fff;">$CC2</td>
<td style="text-align: center; border: 0px; background-color: #001900; width: 32px; height: 32px; color: #fff;">$CC3</td>
<td style="text-align: center; border: 0px; background-color: #002201; width: 32px; height: 32px; color: #fff;">$CC4</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$CC5</td>
<td style="text-align: center; border: 0px; background-color: #00141F; width: 32px; height: 32px; color: #fff;">$CC6</td>
<td style="text-align: center; border: 0px; background-color: #000644; width: 32px; height: 32px; color: #fff;">$CC7</td>
<td style="text-align: center; border: 0px; background-color: #000064; width: 32px; height: 32px; color: #fff;">$CC8</td>
<td style="text-align: center; border: 0px; background-color: #0B0065; width: 32px; height: 32px; color: #fff;">$CC9</td>
<td style="text-align: center; border: 0px; background-color: #2C0043; width: 32px; height: 32px; color: #fff;">$CCA</td>
<td style="text-align: center; border: 0px; background-color: #370021; width: 32px; height: 32px; color: #fff;">$CCB</td>
<td style="text-align: center; border: 0px; background-color: #380000; width: 32px; height: 32px; color: #fff;">$CCC</td>
<td style="text-align: center; border: 0px; background-color: #353537; width: 32px; height: 32px; color: #fff;">$CCD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CCE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CCF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CD0</td>
<td style="text-align: center; border: 0px; background-color: #451300; width: 32px; height: 32px; color: #fff;">$CD1</td>
<td style="text-align: center; border: 0px; background-color: #2A2600; width: 32px; height: 32px; color: #fff;">$CD2</td>
<td style="text-align: center; border: 0px; background-color: #0A3500; width: 32px; height: 32px; color: #fff;">$CD3</td>
<td style="text-align: center; border: 0px; background-color: #003D01; width: 32px; height: 32px; color: #fff;">$CD4</td>
<td style="text-align: center; border: 0px; background-color: #003E00; width: 32px; height: 32px; color: #fff;">$CD5</td>
<td style="text-align: center; border: 0px; background-color: #00303C; width: 32px; height: 32px; color: #fff;">$CD6</td>
<td style="text-align: center; border: 0px; background-color: #012360; width: 32px; height: 32px; color: #fff;">$CD7</td>
<td style="text-align: center; border: 0px; background-color: #0D0E84; width: 32px; height: 32px; color: #fff;">$CD8</td>
<td style="text-align: center; border: 0px; background-color: #280087; width: 32px; height: 32px; color: #fff;">$CD9</td>
<td style="text-align: center; border: 0px; background-color: #470064; width: 32px; height: 32px; color: #fff;">$CDA</td>
<td style="text-align: center; border: 0px; background-color: #54003B; width: 32px; height: 32px; color: #fff;">$CDB</td>
<td style="text-align: center; border: 0px; background-color: #550400; width: 32px; height: 32px; color: #fff;">$CDC</td>
<td style="text-align: center; border: 0px; background-color: #505050; width: 32px; height: 32px; color: #fff;">$CDD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CDE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CDF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CE0</td>
<td style="text-align: center; border: 0px; background-color: #602D01; width: 32px; height: 32px; color: #fff;">$CE1</td>
<td style="text-align: center; border: 0px; background-color: #444100; width: 32px; height: 32px; color: #fff;">$CE2</td>
<td style="text-align: center; border: 0px; background-color: #265000; width: 32px; height: 32px; color: #fff;">$CE3</td>
<td style="text-align: center; border: 0px; background-color: #095900; width: 32px; height: 32px; color: #fff;">$CE4</td>
<td style="text-align: center; border: 0px; background-color: #005915; width: 32px; height: 32px; color: #fff;">$CE5</td>
<td style="text-align: center; border: 0px; background-color: #004A55; width: 32px; height: 32px; color: #fff;">$CE6</td>
<td style="text-align: center; border: 0px; background-color: #0A3E7A; width: 32px; height: 32px; color: #fff;">$CE7</td>
<td style="text-align: center; border: 0px; background-color: #2826A1; width: 32px; height: 32px; color: #fff;">$CE8</td>
<td style="text-align: center; border: 0px; background-color: #431B9A; width: 32px; height: 32px; color: #fff;">$CE9</td>
<td style="text-align: center; border: 0px; background-color: #63107A; width: 32px; height: 32px; color: #fff;">$CEA</td>
<td style="text-align: center; border: 0px; background-color: #6E1157; width: 32px; height: 32px; color: #fff;">$CEB</td>
<td style="text-align: center; border: 0px; background-color: #711F14; width: 32px; height: 32px; color: #fff;">$CEC</td>
<td style="text-align: center; border: 0px; background-color: #6A696E; width: 32px; height: 32px; color: #fff;">$CED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1A1C19; width: 32px; height: 32px; color: #fff;">$CF0</td>
<td style="text-align: center; border: 0px; background-color: #7F4806; width: 32px; height: 32px; color: #fff;">$CF1</td>
<td style="text-align: center; border: 0px; background-color: #615E00; width: 32px; height: 32px; color: #fff;">$CF2</td>
<td style="text-align: center; border: 0px; background-color: #436E01; width: 32px; height: 32px; color: #fff;">$CF3</td>
<td style="text-align: center; border: 0px; background-color: #267607; width: 32px; height: 32px; color: #fff;">$CF4</td>
<td style="text-align: center; border: 0px; background-color: #177730; width: 32px; height: 32px; color: #fff;">$CF5</td>
<td style="text-align: center; border: 0px; background-color: #156772; width: 32px; height: 32px; color: #fff;">$CF6</td>
<td style="text-align: center; border: 0px; background-color: #225999; width: 32px; height: 32px; color: #fff;">$CF7</td>
<td style="text-align: center; border: 0px; background-color: #4242BC; width: 32px; height: 32px; color: #fff;">$CF8</td>
<td style="text-align: center; border: 0px; background-color: #5F35BE; width: 32px; height: 32px; color: #fff;">$CF9</td>
<td style="text-align: center; border: 0px; background-color: #802B9B; width: 32px; height: 32px; color: #fff;">$CFA</td>
<td style="text-align: center; border: 0px; background-color: #8A2E75; width: 32px; height: 32px; color: #fff;">$CFB</td>
<td style="text-align: center; border: 0px; background-color: #8D3B30; width: 32px; height: 32px; color: #fff;">$CFC</td>
<td style="text-align: center; border: 0px; background-color: #868683; width: 32px; height: 32px; color: #000;">$CFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$CFF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $D</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$D00</td>
<td style="text-align: center; border: 0px; background-color: #8E6632; width: 32px; height: 32px; color: #fff;">$D01</td>
<td style="text-align: center; border: 0px; background-color: #78751E; width: 32px; height: 32px; color: #fff;">$D02</td>
<td style="text-align: center; border: 0px; background-color: #608219; width: 32px; height: 32px; color: #fff;">$D03</td>
<td style="text-align: center; border: 0px; background-color: #498932; width: 32px; height: 32px; color: #fff;">$D04</td>
<td style="text-align: center; border: 0px; background-color: #418751; width: 32px; height: 32px; color: #fff;">$D05</td>
<td style="text-align: center; border: 0px; background-color: #407C83; width: 32px; height: 32px; color: #fff;">$D06</td>
<td style="text-align: center; border: 0px; background-color: #4973A5; width: 32px; height: 32px; color: #fff;">$D07</td>
<td style="text-align: center; border: 0px; background-color: #6061BE; width: 32px; height: 32px; color: #fff;">$D08</td>
<td style="text-align: center; border: 0px; background-color: #7756C0; width: 32px; height: 32px; color: #fff;">$D09</td>
<td style="text-align: center; border: 0px; background-color: #904EA5; width: 32px; height: 32px; color: #fff;">$D0A</td>
<td style="text-align: center; border: 0px; background-color: #9A4F87; width: 32px; height: 32px; color: #fff;">$D0B</td>
<td style="text-align: center; border: 0px; background-color: #9B5B51; width: 32px; height: 32px; color: #fff;">$D0C</td>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$D0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$D10</td>
<td style="text-align: center; border: 0px; background-color: #A8804E; width: 32px; height: 32px; color: #000;">$D11</td>
<td style="text-align: center; border: 0px; background-color: #938E38; width: 32px; height: 32px; color: #000;">$D12</td>
<td style="text-align: center; border: 0px; background-color: #7A9B37; width: 32px; height: 32px; color: #000;">$D13</td>
<td style="text-align: center; border: 0px; background-color: #63A34C; width: 32px; height: 32px; color: #000;">$D14</td>
<td style="text-align: center; border: 0px; background-color: #5CA16B; width: 32px; height: 32px; color: #000;">$D15</td>
<td style="text-align: center; border: 0px; background-color: #5A969F; width: 32px; height: 32px; color: #000;">$D16</td>
<td style="text-align: center; border: 0px; background-color: #648DBA; width: 32px; height: 32px; color: #000;">$D17</td>
<td style="text-align: center; border: 0px; background-color: #7C7DD7; width: 32px; height: 32px; color: #000;">$D18</td>
<td style="text-align: center; border: 0px; background-color: #9172D7; width: 32px; height: 32px; color: #000;">$D19</td>
<td style="text-align: center; border: 0px; background-color: #A969BF; width: 32px; height: 32px; color: #000;">$D1A</td>
<td style="text-align: center; border: 0px; background-color: #B36AA2; width: 32px; height: 32px; color: #000;">$D1B</td>
<td style="text-align: center; border: 0px; background-color: #B4756F; width: 32px; height: 32px; color: #000;">$D1C</td>
<td style="text-align: center; border: 0px; background-color: #ACABB0; width: 32px; height: 32px; color: #000;">$D1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #77767A; width: 32px; height: 32px; color: #fff;">$D20</td>
<td style="text-align: center; border: 0px; background-color: #C09968; width: 32px; height: 32px; color: #000;">$D21</td>
<td style="text-align: center; border: 0px; background-color: #ACA84F; width: 32px; height: 32px; color: #000;">$D22</td>
<td style="text-align: center; border: 0px; background-color: #94B44F; width: 32px; height: 32px; color: #000;">$D23</td>
<td style="text-align: center; border: 0px; background-color: #7CBA67; width: 32px; height: 32px; color: #000;">$D24</td>
<td style="text-align: center; border: 0px; background-color: #75B988; width: 32px; height: 32px; color: #000;">$D25</td>
<td style="text-align: center; border: 0px; background-color: #74AEB9; width: 32px; height: 32px; color: #000;">$D26</td>
<td style="text-align: center; border: 0px; background-color: #7CA5D1; width: 32px; height: 32px; color: #000;">$D27</td>
<td style="text-align: center; border: 0px; background-color: #9594F1; width: 32px; height: 32px; color: #000;">$D28</td>
<td style="text-align: center; border: 0px; background-color: #AA8AEB; width: 32px; height: 32px; color: #000;">$D29</td>
<td style="text-align: center; border: 0px; background-color: #C383D6; width: 32px; height: 32px; color: #000;">$D2A</td>
<td style="text-align: center; border: 0px; background-color: #CA85B8; width: 32px; height: 32px; color: #000;">$D2B</td>
<td style="text-align: center; border: 0px; background-color: #CD8E85; width: 32px; height: 32px; color: #000;">$D2C</td>
<td style="text-align: center; border: 0px; background-color: #C5C5C3; width: 32px; height: 32px; color: #000;">$D2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939393; width: 32px; height: 32px; color: #000;">$D30</td>
<td style="text-align: center; border: 0px; background-color: #DAB284; width: 32px; height: 32px; color: #000;">$D31</td>
<td style="text-align: center; border: 0px; background-color: #C4C268; width: 32px; height: 32px; color: #000;">$D32</td>
<td style="text-align: center; border: 0px; background-color: #ADCD6D; width: 32px; height: 32px; color: #000;">$D33</td>
<td style="text-align: center; border: 0px; background-color: #9AD384; width: 32px; height: 32px; color: #000;">$D34</td>
<td style="text-align: center; border: 0px; background-color: #8ED2A1; width: 32px; height: 32px; color: #000;">$D35</td>
<td style="text-align: center; border: 0px; background-color: #8EC9D1; width: 32px; height: 32px; color: #000;">$D36</td>
<td style="text-align: center; border: 0px; background-color: #96BFEB; width: 32px; height: 32px; color: #000;">$D37</td>
<td style="text-align: center; border: 0px; background-color: #AFB0FD; width: 32px; height: 32px; color: #000;">$D38</td>
<td style="text-align: center; border: 0px; background-color: #C1A6FD; width: 32px; height: 32px; color: #000;">$D39</td>
<td style="text-align: center; border: 0px; background-color: #DB9EEE; width: 32px; height: 32px; color: #000;">$D3A</td>
<td style="text-align: center; border: 0px; background-color: #E39ECE; width: 32px; height: 32px; color: #000;">$D3B</td>
<td style="text-align: center; border: 0px; background-color: #E7A8A2; width: 32px; height: 32px; color: #000;">$D3C</td>
<td style="text-align: center; border: 0px; background-color: #DFDEE3; width: 32px; height: 32px; color: #000;">$D3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #AEABAA; width: 32px; height: 32px; color: #000;">$D40</td>
<td style="text-align: center; border: 0px; background-color: #F0CB9D; width: 32px; height: 32px; color: #000;">$D41</td>
<td style="text-align: center; border: 0px; background-color: #DDDB86; width: 32px; height: 32px; color: #000;">$D42</td>
<td style="text-align: center; border: 0px; background-color: #C6E587; width: 32px; height: 32px; color: #000;">$D43</td>
<td style="text-align: center; border: 0px; background-color: #B2EB9E; width: 32px; height: 32px; color: #000;">$D44</td>
<td style="text-align: center; border: 0px; background-color: #A9EBBB; width: 32px; height: 32px; color: #000;">$D45</td>
<td style="text-align: center; border: 0px; background-color: #A8E1EA; width: 32px; height: 32px; color: #000;">$D46</td>
<td style="text-align: center; border: 0px; background-color: #B1D7FE; width: 32px; height: 32px; color: #000;">$D47</td>
<td style="text-align: center; border: 0px; background-color: #C8C9FF; width: 32px; height: 32px; color: #000;">$D48</td>
<td style="text-align: center; border: 0px; background-color: #DCBFFF; width: 32px; height: 32px; color: #000;">$D49</td>
<td style="text-align: center; border: 0px; background-color: #F5B7FF; width: 32px; height: 32px; color: #000;">$D4A</td>
<td style="text-align: center; border: 0px; background-color: #FBB8EB; width: 32px; height: 32px; color: #000;">$D4B</td>
<td style="text-align: center; border: 0px; background-color: #FCC3BD; width: 32px; height: 32px; color: #000;">$D4C</td>
<td style="text-align: center; border: 0px; background-color: #F8F8FA; width: 32px; height: 32px; color: #000;">$D4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C6C6C4; width: 32px; height: 32px; color: #000;">$D50</td>
<td style="text-align: center; border: 0px; background-color: #FDE6BA; width: 32px; height: 32px; color: #000;">$D51</td>
<td style="text-align: center; border: 0px; background-color: #F5F3A6; width: 32px; height: 32px; color: #000;">$D52</td>
<td style="text-align: center; border: 0px; background-color: #DFFDA6; width: 32px; height: 32px; color: #000;">$D53</td>
<td style="text-align: center; border: 0px; background-color: #CDFFB8; width: 32px; height: 32px; color: #000;">$D54</td>
<td style="text-align: center; border: 0px; background-color: #C5FFD6; width: 32px; height: 32px; color: #000;">$D55</td>
<td style="text-align: center; border: 0px; background-color: #C4F9FF; width: 32px; height: 32px; color: #000;">$D56</td>
<td style="text-align: center; border: 0px; background-color: #CAEFFE; width: 32px; height: 32px; color: #000;">$D57</td>
<td style="text-align: center; border: 0px; background-color: #E2E1FF; width: 32px; height: 32px; color: #000;">$D58</td>
<td style="text-align: center; border: 0px; background-color: #F2D8FF; width: 32px; height: 32px; color: #000;">$D59</td>
<td style="text-align: center; border: 0px; background-color: #FFD1FF; width: 32px; height: 32px; color: #000;">$D5A</td>
<td style="text-align: center; border: 0px; background-color: #FED2FF; width: 32px; height: 32px; color: #000;">$D5B</td>
<td style="text-align: center; border: 0px; background-color: #FEDCD3; width: 32px; height: 32px; color: #000;">$D5C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$D5D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DFDFDF; width: 32px; height: 32px; color: #000;">$D60</td>
<td style="text-align: center; border: 0px; background-color: #FDFCD7; width: 32px; height: 32px; color: #000;">$D61</td>
<td style="text-align: center; border: 0px; background-color: #FFFEC5; width: 32px; height: 32px; color: #000;">$D62</td>
<td style="text-align: center; border: 0px; background-color: #F8FEC4; width: 32px; height: 32px; color: #000;">$D63</td>
<td style="text-align: center; border: 0px; background-color: #E6FFD7; width: 32px; height: 32px; color: #000;">$D64</td>
<td style="text-align: center; border: 0px; background-color: #DFFFEF; width: 32px; height: 32px; color: #000;">$D65</td>
<td style="text-align: center; border: 0px; background-color: #DFFEFF; width: 32px; height: 32px; color: #000;">$D66</td>
<td style="text-align: center; border: 0px; background-color: #E7FFFF; width: 32px; height: 32px; color: #000;">$D67</td>
<td style="text-align: center; border: 0px; background-color: #F9FAFE; width: 32px; height: 32px; color: #000;">$D68</td>
<td style="text-align: center; border: 0px; background-color: #FEF1FF; width: 32px; height: 32px; color: #000;">$D69</td>
<td style="text-align: center; border: 0px; background-color: #FEECFF; width: 32px; height: 32px; color: #000;">$D6A</td>
<td style="text-align: center; border: 0px; background-color: #FFECFE; width: 32px; height: 32px; color: #000;">$D6B</td>
<td style="text-align: center; border: 0px; background-color: #FFF3F2; width: 32px; height: 32px; color: #000;">$D6C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$D6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D6E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F8F8F8; width: 32px; height: 32px; color: #000;">$D70</td>
<td style="text-align: center; border: 0px; background-color: #0279FF; width: 32px; height: 32px; color: #fff;">$D71</td>
<td style="text-align: center; border: 0px; background-color: #2F44FE; width: 32px; height: 32px; color: #fff;">$D72</td>
<td style="text-align: center; border: 0px; background-color: #8C17FD; width: 32px; height: 32px; color: #fff;">$D73</td>
<td style="text-align: center; border: 0px; background-color: #EB00FF; width: 32px; height: 32px; color: #fff;">$D74</td>
<td style="text-align: center; border: 0px; background-color: #FF00C6; width: 32px; height: 32px; color: #fff;">$D75</td>
<td style="text-align: center; border: 0px; background-color: #FE1600; width: 32px; height: 32px; color: #fff;">$D76</td>
<td style="text-align: center; border: 0px; background-color: #EF4201; width: 32px; height: 32px; color: #fff;">$D77</td>
<td style="text-align: center; border: 0px; background-color: #8E7900; width: 32px; height: 32px; color: #fff;">$D78</td>
<td style="text-align: center; border: 0px; background-color: #39A501; width: 32px; height: 32px; color: #fff;">$D79</td>
<td style="text-align: center; border: 0px; background-color: #01D802; width: 32px; height: 32px; color: #fff;">$D7A</td>
<td style="text-align: center; border: 0px; background-color: #00D500; width: 32px; height: 32px; color: #fff;">$D7B</td>
<td style="text-align: center; border: 0px; background-color: #00A8C7; width: 32px; height: 32px; color: #fff;">$D7C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$D7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$D80</td>
<td style="text-align: center; border: 0px; background-color: #0093FF; width: 32px; height: 32px; color: #fff;">$D81</td>
<td style="text-align: center; border: 0px; background-color: #485FFE; width: 32px; height: 32px; color: #fff;">$D82</td>
<td style="text-align: center; border: 0px; background-color: #A730FD; width: 32px; height: 32px; color: #fff;">$D83</td>
<td style="text-align: center; border: 0px; background-color: #FE03FE; width: 32px; height: 32px; color: #fff;">$D84</td>
<td style="text-align: center; border: 0px; background-color: #FE03DF; width: 32px; height: 32px; color: #fff;">$D85</td>
<td style="text-align: center; border: 0px; background-color: #FF3111; width: 32px; height: 32px; color: #fff;">$D86</td>
<td style="text-align: center; border: 0px; background-color: #FE5A00; width: 32px; height: 32px; color: #000;">$D87</td>
<td style="text-align: center; border: 0px; background-color: #A59500; width: 32px; height: 32px; color: #000;">$D88</td>
<td style="text-align: center; border: 0px; background-color: #50BE00; width: 32px; height: 32px; color: #000;">$D89</td>
<td style="text-align: center; border: 0px; background-color: #01F100; width: 32px; height: 32px; color: #000;">$D8A</td>
<td style="text-align: center; border: 0px; background-color: #01ED0E; width: 32px; height: 32px; color: #000;">$D8B</td>
<td style="text-align: center; border: 0px; background-color: #00C1DC; width: 32px; height: 32px; color: #000;">$D8C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$D8D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D8E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$D90</td>
<td style="text-align: center; border: 0px; background-color: #07ABFE; width: 32px; height: 32px; color: #000;">$D91</td>
<td style="text-align: center; border: 0px; background-color: #6476FE; width: 32px; height: 32px; color: #000;">$D92</td>
<td style="text-align: center; border: 0px; background-color: #BD4AFE; width: 32px; height: 32px; color: #000;">$D93</td>
<td style="text-align: center; border: 0px; background-color: #FE20FD; width: 32px; height: 32px; color: #fff;">$D94</td>
<td style="text-align: center; border: 0px; background-color: #FF20F2; width: 32px; height: 32px; color: #fff;">$D95</td>
<td style="text-align: center; border: 0px; background-color: #FE4C2E; width: 32px; height: 32px; color: #fff;">$D96</td>
<td style="text-align: center; border: 0px; background-color: #FF7500; width: 32px; height: 32px; color: #000;">$D97</td>
<td style="text-align: center; border: 0px; background-color: #BDB001; width: 32px; height: 32px; color: #000;">$D98</td>
<td style="text-align: center; border: 0px; background-color: #68D900; width: 32px; height: 32px; color: #000;">$D99</td>
<td style="text-align: center; border: 0px; background-color: #00FF01; width: 32px; height: 32px; color: #000;">$D9A</td>
<td style="text-align: center; border: 0px; background-color: #00FF26; width: 32px; height: 32px; color: #000;">$D9B</td>
<td style="text-align: center; border: 0px; background-color: #00D9F7; width: 32px; height: 32px; color: #000;">$D9C</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$D9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$DA0</td>
<td style="text-align: center; border: 0px; background-color: #0C0000; width: 32px; height: 32px; color: #fff;">$DA1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DA2</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$DA3</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$DA4</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$DA5</td>
<td style="text-align: center; border: 0px; background-color: #000004; width: 32px; height: 32px; color: #fff;">$DA6</td>
<td style="text-align: center; border: 0px; background-color: #00001C; width: 32px; height: 32px; color: #fff;">$DA7</td>
<td style="text-align: center; border: 0px; background-color: #000034; width: 32px; height: 32px; color: #fff;">$DA8</td>
<td style="text-align: center; border: 0px; background-color: #000032; width: 32px; height: 32px; color: #fff;">$DA9</td>
<td style="text-align: center; border: 0px; background-color: #0B0021; width: 32px; height: 32px; color: #fff;">$DAA</td>
<td style="text-align: center; border: 0px; background-color: #140105; width: 32px; height: 32px; color: #fff;">$DAB</td>
<td style="text-align: center; border: 0px; background-color: #130101; width: 32px; height: 32px; color: #fff;">$DAC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DAE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$DB0</td>
<td style="text-align: center; border: 0px; background-color: #0E0000; width: 32px; height: 32px; color: #fff;">$DB1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DB2</td>
<td style="text-align: center; border: 0px; background-color: #000200; width: 32px; height: 32px; color: #fff;">$DB3</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$DB4</td>
<td style="text-align: center; border: 0px; background-color: #000700; width: 32px; height: 32px; color: #fff;">$DB5</td>
<td style="text-align: center; border: 0px; background-color: #000004; width: 32px; height: 32px; color: #fff;">$DB6</td>
<td style="text-align: center; border: 0px; background-color: #000120; width: 32px; height: 32px; color: #fff;">$DB7</td>
<td style="text-align: center; border: 0px; background-color: #000038; width: 32px; height: 32px; color: #fff;">$DB8</td>
<td style="text-align: center; border: 0px; background-color: #000037; width: 32px; height: 32px; color: #fff;">$DB9</td>
<td style="text-align: center; border: 0px; background-color: #0F0020; width: 32px; height: 32px; color: #fff;">$DBA</td>
<td style="text-align: center; border: 0px; background-color: #160006; width: 32px; height: 32px; color: #fff;">$DBB</td>
<td style="text-align: center; border: 0px; background-color: #160100; width: 32px; height: 32px; color: #fff;">$DBC</td>
<td style="text-align: center; border: 0px; background-color: #0D0D0D; width: 32px; height: 32px; color: #fff;">$DBD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DBE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DBF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DC0</td>
<td style="text-align: center; border: 0px; background-color: #1F0000; width: 32px; height: 32px; color: #fff;">$DC1</td>
<td style="text-align: center; border: 0px; background-color: #0B0800; width: 32px; height: 32px; color: #fff;">$DC2</td>
<td style="text-align: center; border: 0px; background-color: #011300; width: 32px; height: 32px; color: #fff;">$DC3</td>
<td style="text-align: center; border: 0px; background-color: #001A00; width: 32px; height: 32px; color: #fff;">$DC4</td>
<td style="text-align: center; border: 0px; background-color: #001A00; width: 32px; height: 32px; color: #fff;">$DC5</td>
<td style="text-align: center; border: 0px; background-color: #001017; width: 32px; height: 32px; color: #fff;">$DC6</td>
<td style="text-align: center; border: 0px; background-color: #000634; width: 32px; height: 32px; color: #fff;">$DC7</td>
<td style="text-align: center; border: 0px; background-color: #00014C; width: 32px; height: 32px; color: #fff;">$DC8</td>
<td style="text-align: center; border: 0px; background-color: #070050; width: 32px; height: 32px; color: #fff;">$DC9</td>
<td style="text-align: center; border: 0px; background-color: #210038; width: 32px; height: 32px; color: #fff;">$DCA</td>
<td style="text-align: center; border: 0px; background-color: #2A0017; width: 32px; height: 32px; color: #fff;">$DCB</td>
<td style="text-align: center; border: 0px; background-color: #2A0001; width: 32px; height: 32px; color: #fff;">$DCC</td>
<td style="text-align: center; border: 0px; background-color: #282828; width: 32px; height: 32px; color: #fff;">$DCD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DCE</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$DCF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DD0</td>
<td style="text-align: center; border: 0px; background-color: #391400; width: 32px; height: 32px; color: #fff;">$DD1</td>
<td style="text-align: center; border: 0px; background-color: #262200; width: 32px; height: 32px; color: #fff;">$DD2</td>
<td style="text-align: center; border: 0px; background-color: #0E2E00; width: 32px; height: 32px; color: #fff;">$DD3</td>
<td style="text-align: center; border: 0px; background-color: #003500; width: 32px; height: 32px; color: #fff;">$DD4</td>
<td style="text-align: center; border: 0px; background-color: #003400; width: 32px; height: 32px; color: #fff;">$DD5</td>
<td style="text-align: center; border: 0px; background-color: #022A32; width: 32px; height: 32px; color: #fff;">$DD6</td>
<td style="text-align: center; border: 0px; background-color: #00204D; width: 32px; height: 32px; color: #fff;">$DD7</td>
<td style="text-align: center; border: 0px; background-color: #0E1067; width: 32px; height: 32px; color: #fff;">$DD8</td>
<td style="text-align: center; border: 0px; background-color: #240564; width: 32px; height: 32px; color: #fff;">$DD9</td>
<td style="text-align: center; border: 0px; background-color: #3D004F; width: 32px; height: 32px; color: #fff;">$DDA</td>
<td style="text-align: center; border: 0px; background-color: #450133; width: 32px; height: 32px; color: #fff;">$DDB</td>
<td style="text-align: center; border: 0px; background-color: #450A04; width: 32px; height: 32px; color: #fff;">$DDC</td>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$DDD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DDE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DDF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0B0B0B; width: 32px; height: 32px; color: #fff;">$DE0</td>
<td style="text-align: center; border: 0px; background-color: #552F00; width: 32px; height: 32px; color: #fff;">$DE1</td>
<td style="text-align: center; border: 0px; background-color: #433F00; width: 32px; height: 32px; color: #fff;">$DE2</td>
<td style="text-align: center; border: 0px; background-color: #2A4A01; width: 32px; height: 32px; color: #fff;">$DE3</td>
<td style="text-align: center; border: 0px; background-color: #135002; width: 32px; height: 32px; color: #fff;">$DE4</td>
<td style="text-align: center; border: 0px; background-color: #0A501E; width: 32px; height: 32px; color: #fff;">$DE5</td>
<td style="text-align: center; border: 0px; background-color: #0A454E; width: 32px; height: 32px; color: #fff;">$DE6</td>
<td style="text-align: center; border: 0px; background-color: #123D68; width: 32px; height: 32px; color: #fff;">$DE7</td>
<td style="text-align: center; border: 0px; background-color: #2B2C88; width: 32px; height: 32px; color: #fff;">$DE8</td>
<td style="text-align: center; border: 0px; background-color: #3F2185; width: 32px; height: 32px; color: #fff;">$DE9</td>
<td style="text-align: center; border: 0px; background-color: #58186E; width: 32px; height: 32px; color: #fff;">$DEA</td>
<td style="text-align: center; border: 0px; background-color: #621952; width: 32px; height: 32px; color: #fff;">$DEB</td>
<td style="text-align: center; border: 0px; background-color: #63241B; width: 32px; height: 32px; color: #fff;">$DEC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E60; width: 32px; height: 32px; color: #fff;">$DED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #282826; width: 32px; height: 32px; color: #fff;">$DF0</td>
<td style="text-align: center; border: 0px; background-color: #724B19; width: 32px; height: 32px; color: #fff;">$DF1</td>
<td style="text-align: center; border: 0px; background-color: #5C5A02; width: 32px; height: 32px; color: #fff;">$DF2</td>
<td style="text-align: center; border: 0px; background-color: #456601; width: 32px; height: 32px; color: #fff;">$DF3</td>
<td style="text-align: center; border: 0px; background-color: #2E6C19; width: 32px; height: 32px; color: #fff;">$DF4</td>
<td style="text-align: center; border: 0px; background-color: #266C37; width: 32px; height: 32px; color: #fff;">$DF5</td>
<td style="text-align: center; border: 0px; background-color: #26606B; width: 32px; height: 32px; color: #fff;">$DF6</td>
<td style="text-align: center; border: 0px; background-color: #2E5785; width: 32px; height: 32px; color: #fff;">$DF7</td>
<td style="text-align: center; border: 0px; background-color: #4646A0; width: 32px; height: 32px; color: #fff;">$DF8</td>
<td style="text-align: center; border: 0px; background-color: #5C3CA2; width: 32px; height: 32px; color: #fff;">$DF9</td>
<td style="text-align: center; border: 0px; background-color: #73338C; width: 32px; height: 32px; color: #fff;">$DFA</td>
<td style="text-align: center; border: 0px; background-color: #7E3569; width: 32px; height: 32px; color: #fff;">$DFB</td>
<td style="text-align: center; border: 0px; background-color: #7F4038; width: 32px; height: 32px; color: #fff;">$DFC</td>
<td style="text-align: center; border: 0px; background-color: #777978; width: 32px; height: 32px; color: #fff;">$DFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$DFF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $E</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #4F504C; width: 32px; height: 32px; color: #fff;">$E00</td>
<td style="text-align: center; border: 0px; background-color: #836743; width: 32px; height: 32px; color: #fff;">$E01</td>
<td style="text-align: center; border: 0px; background-color: #737134; width: 32px; height: 32px; color: #fff;">$E02</td>
<td style="text-align: center; border: 0px; background-color: #627A33; width: 32px; height: 32px; color: #fff;">$E03</td>
<td style="text-align: center; border: 0px; background-color: #527F44; width: 32px; height: 32px; color: #fff;">$E04</td>
<td style="text-align: center; border: 0px; background-color: #4C7E59; width: 32px; height: 32px; color: #fff;">$E05</td>
<td style="text-align: center; border: 0px; background-color: #4C777D; width: 32px; height: 32px; color: #fff;">$E06</td>
<td style="text-align: center; border: 0px; background-color: #527190; width: 32px; height: 32px; color: #fff;">$E07</td>
<td style="text-align: center; border: 0px; background-color: #6364A4; width: 32px; height: 32px; color: #fff;">$E08</td>
<td style="text-align: center; border: 0px; background-color: #725CA6; width: 32px; height: 32px; color: #fff;">$E09</td>
<td style="text-align: center; border: 0px; background-color: #855891; width: 32px; height: 32px; color: #fff;">$E0A</td>
<td style="text-align: center; border: 0px; background-color: #8A587D; width: 32px; height: 32px; color: #fff;">$E0B</td>
<td style="text-align: center; border: 0px; background-color: #89605A; width: 32px; height: 32px; color: #fff;">$E0C</td>
<td style="text-align: center; border: 0px; background-color: #868686; width: 32px; height: 32px; color: #000;">$E0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #6B6B69; width: 32px; height: 32px; color: #fff;">$E10</td>
<td style="text-align: center; border: 0px; background-color: #9B8160; width: 32px; height: 32px; color: #000;">$E11</td>
<td style="text-align: center; border: 0px; background-color: #8D8B52; width: 32px; height: 32px; color: #000;">$E12</td>
<td style="text-align: center; border: 0px; background-color: #7D9450; width: 32px; height: 32px; color: #000;">$E13</td>
<td style="text-align: center; border: 0px; background-color: #6C975F; width: 32px; height: 32px; color: #000;">$E14</td>
<td style="text-align: center; border: 0px; background-color: #689773; width: 32px; height: 32px; color: #000;">$E15</td>
<td style="text-align: center; border: 0px; background-color: #679096; width: 32px; height: 32px; color: #000;">$E16</td>
<td style="text-align: center; border: 0px; background-color: #6C89A7; width: 32px; height: 32px; color: #000;">$E17</td>
<td style="text-align: center; border: 0px; background-color: #7D7FBA; width: 32px; height: 32px; color: #000;">$E18</td>
<td style="text-align: center; border: 0px; background-color: #8B76BC; width: 32px; height: 32px; color: #000;">$E19</td>
<td style="text-align: center; border: 0px; background-color: #9D71B0; width: 32px; height: 32px; color: #000;">$E1A</td>
<td style="text-align: center; border: 0px; background-color: #A2729A; width: 32px; height: 32px; color: #000;">$E1B</td>
<td style="text-align: center; border: 0px; background-color: #A37A74; width: 32px; height: 32px; color: #000;">$E1C</td>
<td style="text-align: center; border: 0px; background-color: #9E9EA0; width: 32px; height: 32px; color: #000;">$E1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #848484; width: 32px; height: 32px; color: #000;">$E20</td>
<td style="text-align: center; border: 0px; background-color: #B79B7A; width: 32px; height: 32px; color: #000;">$E21</td>
<td style="text-align: center; border: 0px; background-color: #A7A66A; width: 32px; height: 32px; color: #000;">$E22</td>
<td style="text-align: center; border: 0px; background-color: #97AD6C; width: 32px; height: 32px; color: #000;">$E23</td>
<td style="text-align: center; border: 0px; background-color: #87B27A; width: 32px; height: 32px; color: #000;">$E24</td>
<td style="text-align: center; border: 0px; background-color: #82B18D; width: 32px; height: 32px; color: #000;">$E25</td>
<td style="text-align: center; border: 0px; background-color: #82A9B0; width: 32px; height: 32px; color: #000;">$E26</td>
<td style="text-align: center; border: 0px; background-color: #87A4C2; width: 32px; height: 32px; color: #000;">$E27</td>
<td style="text-align: center; border: 0px; background-color: #9799D4; width: 32px; height: 32px; color: #000;">$E28</td>
<td style="text-align: center; border: 0px; background-color: #A791D9; width: 32px; height: 32px; color: #000;">$E29</td>
<td style="text-align: center; border: 0px; background-color: #B88DC3; width: 32px; height: 32px; color: #000;">$E2A</td>
<td style="text-align: center; border: 0px; background-color: #BB8DB2; width: 32px; height: 32px; color: #000;">$E2B</td>
<td style="text-align: center; border: 0px; background-color: #BC938D; width: 32px; height: 32px; color: #000;">$E2C</td>
<td style="text-align: center; border: 0px; background-color: #B9B9B6; width: 32px; height: 32px; color: #000;">$E2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #9D9F9E; width: 32px; height: 32px; color: #000;">$E30</td>
<td style="text-align: center; border: 0px; background-color: #D0B493; width: 32px; height: 32px; color: #000;">$E31</td>
<td style="text-align: center; border: 0px; background-color: #C0BF81; width: 32px; height: 32px; color: #000;">$E32</td>
<td style="text-align: center; border: 0px; background-color: #B1C783; width: 32px; height: 32px; color: #000;">$E33</td>
<td style="text-align: center; border: 0px; background-color: #A0CB93; width: 32px; height: 32px; color: #000;">$E34</td>
<td style="text-align: center; border: 0px; background-color: #9BCAA6; width: 32px; height: 32px; color: #000;">$E35</td>
<td style="text-align: center; border: 0px; background-color: #9BC2C9; width: 32px; height: 32px; color: #000;">$E36</td>
<td style="text-align: center; border: 0px; background-color: #A0BDDB; width: 32px; height: 32px; color: #000;">$E37</td>
<td style="text-align: center; border: 0px; background-color: #B0B2ED; width: 32px; height: 32px; color: #000;">$E38</td>
<td style="text-align: center; border: 0px; background-color: #BFAAED; width: 32px; height: 32px; color: #000;">$E39</td>
<td style="text-align: center; border: 0px; background-color: #D1A5E1; width: 32px; height: 32px; color: #000;">$E3A</td>
<td style="text-align: center; border: 0px; background-color: #D5A7CB; width: 32px; height: 32px; color: #000;">$E3B</td>
<td style="text-align: center; border: 0px; background-color: #D6ADA7; width: 32px; height: 32px; color: #000;">$E3C</td>
<td style="text-align: center; border: 0px; background-color: #D1D1D1; width: 32px; height: 32px; color: #000;">$E3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #B9B9B9; width: 32px; height: 32px; color: #000;">$E40</td>
<td style="text-align: center; border: 0px; background-color: #E8CDAF; width: 32px; height: 32px; color: #000;">$E41</td>
<td style="text-align: center; border: 0px; background-color: #D8D8A0; width: 32px; height: 32px; color: #000;">$E42</td>
<td style="text-align: center; border: 0px; background-color: #CBDE9E; width: 32px; height: 32px; color: #000;">$E43</td>
<td style="text-align: center; border: 0px; background-color: #BDE4AF; width: 32px; height: 32px; color: #000;">$E44</td>
<td style="text-align: center; border: 0px; background-color: #B6E3C2; width: 32px; height: 32px; color: #000;">$E45</td>
<td style="text-align: center; border: 0px; background-color: #B7DCE2; width: 32px; height: 32px; color: #000;">$E46</td>
<td style="text-align: center; border: 0px; background-color: #BDD6F4; width: 32px; height: 32px; color: #000;">$E47</td>
<td style="text-align: center; border: 0px; background-color: #CCCCFF; width: 32px; height: 32px; color: #000;">$E48</td>
<td style="text-align: center; border: 0px; background-color: #D8C5FF; width: 32px; height: 32px; color: #000;">$E49</td>
<td style="text-align: center; border: 0px; background-color: #E9C1F6; width: 32px; height: 32px; color: #000;">$E4A</td>
<td style="text-align: center; border: 0px; background-color: #EEC1E3; width: 32px; height: 32px; color: #000;">$E4B</td>
<td style="text-align: center; border: 0px; background-color: #F0C7C3; width: 32px; height: 32px; color: #000;">$E4C</td>
<td style="text-align: center; border: 0px; background-color: #EBEBEB; width: 32px; height: 32px; color: #000;">$E4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #D3D3D3; width: 32px; height: 32px; color: #000;">$E50</td>
<td style="text-align: center; border: 0px; background-color: #FEE7C7; width: 32px; height: 32px; color: #000;">$E51</td>
<td style="text-align: center; border: 0px; background-color: #F1EFBB; width: 32px; height: 32px; color: #000;">$E52</td>
<td style="text-align: center; border: 0px; background-color: #E4F6BA; width: 32px; height: 32px; color: #000;">$E53</td>
<td style="text-align: center; border: 0px; background-color: #D8FAC8; width: 32px; height: 32px; color: #000;">$E54</td>
<td style="text-align: center; border: 0px; background-color: #D0FCDB; width: 32px; height: 32px; color: #000;">$E55</td>
<td style="text-align: center; border: 0px; background-color: #D0F3F9; width: 32px; height: 32px; color: #000;">$E56</td>
<td style="text-align: center; border: 0px; background-color: #D6EDFD; width: 32px; height: 32px; color: #000;">$E57</td>
<td style="text-align: center; border: 0px; background-color: #E7E4FF; width: 32px; height: 32px; color: #000;">$E58</td>
<td style="text-align: center; border: 0px; background-color: #F1DDFF; width: 32px; height: 32px; color: #000;">$E59</td>
<td style="text-align: center; border: 0px; background-color: #FEDAFE; width: 32px; height: 32px; color: #000;">$E5A</td>
<td style="text-align: center; border: 0px; background-color: #FEDAF7; width: 32px; height: 32px; color: #000;">$E5B</td>
<td style="text-align: center; border: 0px; background-color: #FFE1D7; width: 32px; height: 32px; color: #000;">$E5C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$E5D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #EBEBE9; width: 32px; height: 32px; color: #000;">$E60</td>
<td style="text-align: center; border: 0px; background-color: #FEFDE4; width: 32px; height: 32px; color: #000;">$E61</td>
<td style="text-align: center; border: 0px; background-color: #FEFFD7; width: 32px; height: 32px; color: #000;">$E62</td>
<td style="text-align: center; border: 0px; background-color: #FBFED7; width: 32px; height: 32px; color: #000;">$E63</td>
<td style="text-align: center; border: 0px; background-color: #EFFFE5; width: 32px; height: 32px; color: #000;">$E64</td>
<td style="text-align: center; border: 0px; background-color: #E9FFF3; width: 32px; height: 32px; color: #000;">$E65</td>
<td style="text-align: center; border: 0px; background-color: #E9FFFD; width: 32px; height: 32px; color: #000;">$E66</td>
<td style="text-align: center; border: 0px; background-color: #EFFFFF; width: 32px; height: 32px; color: #000;">$E67</td>
<td style="text-align: center; border: 0px; background-color: #FDFDFE; width: 32px; height: 32px; color: #000;">$E68</td>
<td style="text-align: center; border: 0px; background-color: #FDF7FE; width: 32px; height: 32px; color: #000;">$E69</td>
<td style="text-align: center; border: 0px; background-color: #FEF2FF; width: 32px; height: 32px; color: #000;">$E6A</td>
<td style="text-align: center; border: 0px; background-color: #FEF3FD; width: 32px; height: 32px; color: #000;">$E6B</td>
<td style="text-align: center; border: 0px; background-color: #FEF9F4; width: 32px; height: 32px; color: #000;">$E6C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$E6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E6E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$E70</td>
<td style="text-align: center; border: 0px; background-color: #017BFF; width: 32px; height: 32px; color: #fff;">$E71</td>
<td style="text-align: center; border: 0px; background-color: #2C47FE; width: 32px; height: 32px; color: #fff;">$E72</td>
<td style="text-align: center; border: 0px; background-color: #8E17FD; width: 32px; height: 32px; color: #fff;">$E73</td>
<td style="text-align: center; border: 0px; background-color: #F400FF; width: 32px; height: 32px; color: #fff;">$E74</td>
<td style="text-align: center; border: 0px; background-color: #FF00CD; width: 32px; height: 32px; color: #fff;">$E75</td>
<td style="text-align: center; border: 0px; background-color: #FF1402; width: 32px; height: 32px; color: #fff;">$E76</td>
<td style="text-align: center; border: 0px; background-color: #F74001; width: 32px; height: 32px; color: #fff;">$E77</td>
<td style="text-align: center; border: 0px; background-color: #8C7D01; width: 32px; height: 32px; color: #fff;">$E78</td>
<td style="text-align: center; border: 0px; background-color: #34A601; width: 32px; height: 32px; color: #fff;">$E79</td>
<td style="text-align: center; border: 0px; background-color: #01DE01; width: 32px; height: 32px; color: #000;">$E7A</td>
<td style="text-align: center; border: 0px; background-color: #00D902; width: 32px; height: 32px; color: #fff;">$E7B</td>
<td style="text-align: center; border: 0px; background-color: #00ABC9; width: 32px; height: 32px; color: #fff;">$E7C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$E80</td>
<td style="text-align: center; border: 0px; background-color: #0097FF; width: 32px; height: 32px; color: #fff;">$E81</td>
<td style="text-align: center; border: 0px; background-color: #4462FE; width: 32px; height: 32px; color: #fff;">$E82</td>
<td style="text-align: center; border: 0px; background-color: #A632FC; width: 32px; height: 32px; color: #fff;">$E83</td>
<td style="text-align: center; border: 0px; background-color: #FF01FD; width: 32px; height: 32px; color: #fff;">$E84</td>
<td style="text-align: center; border: 0px; background-color: #FE00EA; width: 32px; height: 32px; color: #fff;">$E85</td>
<td style="text-align: center; border: 0px; background-color: #FE2D12; width: 32px; height: 32px; color: #fff;">$E86</td>
<td style="text-align: center; border: 0px; background-color: #FE5900; width: 32px; height: 32px; color: #000;">$E87</td>
<td style="text-align: center; border: 0px; background-color: #A19900; width: 32px; height: 32px; color: #000;">$E88</td>
<td style="text-align: center; border: 0px; background-color: #47C600; width: 32px; height: 32px; color: #000;">$E89</td>
<td style="text-align: center; border: 0px; background-color: #01F801; width: 32px; height: 32px; color: #000;">$E8A</td>
<td style="text-align: center; border: 0px; background-color: #01F40C; width: 32px; height: 32px; color: #000;">$E8B</td>
<td style="text-align: center; border: 0px; background-color: #00C3E9; width: 32px; height: 32px; color: #000;">$E8C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E8D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E8E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E90</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E91</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E92</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E93</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E94</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E95</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E96</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E97</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E98</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E99</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E9A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E9B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E9C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$E9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$E9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$EA0</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA2</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA3</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA4</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA5</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EA6</td>
<td style="text-align: center; border: 0px; background-color: #01000E; width: 32px; height: 32px; color: #fff;">$EA7</td>
<td style="text-align: center; border: 0px; background-color: #00001C; width: 32px; height: 32px; color: #fff;">$EA8</td>
<td style="text-align: center; border: 0px; background-color: #00001F; width: 32px; height: 32px; color: #fff;">$EA9</td>
<td style="text-align: center; border: 0px; background-color: #02000D; width: 32px; height: 32px; color: #fff;">$EAA</td>
<td style="text-align: center; border: 0px; background-color: #060000; width: 32px; height: 32px; color: #fff;">$EAB</td>
<td style="text-align: center; border: 0px; background-color: #060000; width: 32px; height: 32px; color: #fff;">$EAC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EAE</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$EAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EB0</td>
<td style="text-align: center; border: 0px; background-color: #030000; width: 32px; height: 32px; color: #fff;">$EB1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EB2</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EB3</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EB4</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$EB5</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EB6</td>
<td style="text-align: center; border: 0px; background-color: #010010; width: 32px; height: 32px; color: #fff;">$EB7</td>
<td style="text-align: center; border: 0px; background-color: #000022; width: 32px; height: 32px; color: #fff;">$EB8</td>
<td style="text-align: center; border: 0px; background-color: #000020; width: 32px; height: 32px; color: #fff;">$EB9</td>
<td style="text-align: center; border: 0px; background-color: #040012; width: 32px; height: 32px; color: #fff;">$EBA</td>
<td style="text-align: center; border: 0px; background-color: #090000; width: 32px; height: 32px; color: #fff;">$EBB</td>
<td style="text-align: center; border: 0px; background-color: #0A0000; width: 32px; height: 32px; color: #fff;">$EBC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EBD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EBE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EBF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EC0</td>
<td style="text-align: center; border: 0px; background-color: #140001; width: 32px; height: 32px; color: #fff;">$EC1</td>
<td style="text-align: center; border: 0px; background-color: #060500; width: 32px; height: 32px; color: #fff;">$EC2</td>
<td style="text-align: center; border: 0px; background-color: #000D01; width: 32px; height: 32px; color: #fff;">$EC3</td>
<td style="text-align: center; border: 0px; background-color: #001100; width: 32px; height: 32px; color: #fff;">$EC4</td>
<td style="text-align: center; border: 0px; background-color: #001200; width: 32px; height: 32px; color: #fff;">$EC5</td>
<td style="text-align: center; border: 0px; background-color: #000A10; width: 32px; height: 32px; color: #fff;">$EC6</td>
<td style="text-align: center; border: 0px; background-color: #000326; width: 32px; height: 32px; color: #fff;">$EC7</td>
<td style="text-align: center; border: 0px; background-color: #000031; width: 32px; height: 32px; color: #fff;">$EC8</td>
<td style="text-align: center; border: 0px; background-color: #050032; width: 32px; height: 32px; color: #fff;">$EC9</td>
<td style="text-align: center; border: 0px; background-color: #160022; width: 32px; height: 32px; color: #fff;">$ECA</td>
<td style="text-align: center; border: 0px; background-color: #1B0010; width: 32px; height: 32px; color: #fff;">$ECB</td>
<td style="text-align: center; border: 0px; background-color: #1C0001; width: 32px; height: 32px; color: #fff;">$ECC</td>
<td style="text-align: center; border: 0px; background-color: #1B1B19; width: 32px; height: 32px; color: #fff;">$ECD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ECE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ECF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$ED0</td>
<td style="text-align: center; border: 0px; background-color: #301500; width: 32px; height: 32px; color: #fff;">$ED1</td>
<td style="text-align: center; border: 0px; background-color: #202000; width: 32px; height: 32px; color: #fff;">$ED2</td>
<td style="text-align: center; border: 0px; background-color: #122700; width: 32px; height: 32px; color: #fff;">$ED3</td>
<td style="text-align: center; border: 0px; background-color: #032C00; width: 32px; height: 32px; color: #fff;">$ED4</td>
<td style="text-align: center; border: 0px; background-color: #002C09; width: 32px; height: 32px; color: #fff;">$ED5</td>
<td style="text-align: center; border: 0px; background-color: #01242A; width: 32px; height: 32px; color: #fff;">$ED6</td>
<td style="text-align: center; border: 0px; background-color: #041F3D; width: 32px; height: 32px; color: #fff;">$ED7</td>
<td style="text-align: center; border: 0px; background-color: #14134D; width: 32px; height: 32px; color: #fff;">$ED8</td>
<td style="text-align: center; border: 0px; background-color: #210B53; width: 32px; height: 32px; color: #fff;">$ED9</td>
<td style="text-align: center; border: 0px; background-color: #310741; width: 32px; height: 32px; color: #fff;">$EDA</td>
<td style="text-align: center; border: 0px; background-color: #36082C; width: 32px; height: 32px; color: #fff;">$EDB</td>
<td style="text-align: center; border: 0px; background-color: #370F06; width: 32px; height: 32px; color: #fff;">$EDC</td>
<td style="text-align: center; border: 0px; background-color: #333534; width: 32px; height: 32px; color: #fff;">$EDD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EDE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EDF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #1B1B1B; width: 32px; height: 32px; color: #fff;">$EE0</td>
<td style="text-align: center; border: 0px; background-color: #4C3013; width: 32px; height: 32px; color: #fff;">$EE1</td>
<td style="text-align: center; border: 0px; background-color: #3D3C01; width: 32px; height: 32px; color: #fff;">$EE2</td>
<td style="text-align: center; border: 0px; background-color: #2E4302; width: 32px; height: 32px; color: #fff;">$EE3</td>
<td style="text-align: center; border: 0px; background-color: #1E470F; width: 32px; height: 32px; color: #fff;">$EE4</td>
<td style="text-align: center; border: 0px; background-color: #184723; width: 32px; height: 32px; color: #fff;">$EE5</td>
<td style="text-align: center; border: 0px; background-color: #194047; width: 32px; height: 32px; color: #fff;">$EE6</td>
<td style="text-align: center; border: 0px; background-color: #1D3A58; width: 32px; height: 32px; color: #fff;">$EE7</td>
<td style="text-align: center; border: 0px; background-color: #2F2E6A; width: 32px; height: 32px; color: #fff;">$EE8</td>
<td style="text-align: center; border: 0px; background-color: #3C276A; width: 32px; height: 32px; color: #fff;">$EE9</td>
<td style="text-align: center; border: 0px; background-color: #4D2255; width: 32px; height: 32px; color: #fff;">$EEA</td>
<td style="text-align: center; border: 0px; background-color: #522346; width: 32px; height: 32px; color: #fff;">$EEB</td>
<td style="text-align: center; border: 0px; background-color: #532A26; width: 32px; height: 32px; color: #fff;">$EEC</td>
<td style="text-align: center; border: 0px; background-color: #4F5150; width: 32px; height: 32px; color: #fff;">$EED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #343434; width: 32px; height: 32px; color: #fff;">$EF0</td>
<td style="text-align: center; border: 0px; background-color: #684C2B; width: 32px; height: 32px; color: #fff;">$EF1</td>
<td style="text-align: center; border: 0px; background-color: #58561A; width: 32px; height: 32px; color: #fff;">$EF2</td>
<td style="text-align: center; border: 0px; background-color: #485E1B; width: 32px; height: 32px; color: #fff;">$EF3</td>
<td style="text-align: center; border: 0px; background-color: #3A632B; width: 32px; height: 32px; color: #fff;">$EF4</td>
<td style="text-align: center; border: 0px; background-color: #33623E; width: 32px; height: 32px; color: #fff;">$EF5</td>
<td style="text-align: center; border: 0px; background-color: #335C60; width: 32px; height: 32px; color: #fff;">$EF6</td>
<td style="text-align: center; border: 0px; background-color: #375675; width: 32px; height: 32px; color: #fff;">$EF7</td>
<td style="text-align: center; border: 0px; background-color: #4A4889; width: 32px; height: 32px; color: #fff;">$EF8</td>
<td style="text-align: center; border: 0px; background-color: #574187; width: 32px; height: 32px; color: #fff;">$EF9</td>
<td style="text-align: center; border: 0px; background-color: #683D73; width: 32px; height: 32px; color: #fff;">$EFA</td>
<td style="text-align: center; border: 0px; background-color: #6D3E60; width: 32px; height: 32px; color: #fff;">$EFB</td>
<td style="text-align: center; border: 0px; background-color: #6E453F; width: 32px; height: 32px; color: #fff;">$EFC</td>
<td style="text-align: center; border: 0px; background-color: #6A6B70; width: 32px; height: 32px; color: #fff;">$EFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$EFF</td>
</tr>
<tr>
<td colspan="16" style="text-align: center;">Saturation level $F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #5E5F5B; width: 32px; height: 32px; color: #fff;">$F00</td>
<td style="text-align: center; border: 0px; background-color: #786857; width: 32px; height: 32px; color: #fff;">$F01</td>
<td style="text-align: center; border: 0px; background-color: #6F6F4F; width: 32px; height: 32px; color: #fff;">$F02</td>
<td style="text-align: center; border: 0px; background-color: #677350; width: 32px; height: 32px; color: #fff;">$F03</td>
<td style="text-align: center; border: 0px; background-color: #5E7559; width: 32px; height: 32px; color: #fff;">$F04</td>
<td style="text-align: center; border: 0px; background-color: #5C7563; width: 32px; height: 32px; color: #fff;">$F05</td>
<td style="text-align: center; border: 0px; background-color: #5C7176; width: 32px; height: 32px; color: #fff;">$F06</td>
<td style="text-align: center; border: 0px; background-color: #5F6F7F; width: 32px; height: 32px; color: #fff;">$F07</td>
<td style="text-align: center; border: 0px; background-color: #696888; width: 32px; height: 32px; color: #fff;">$F08</td>
<td style="text-align: center; border: 0px; background-color: #6E6489; width: 32px; height: 32px; color: #fff;">$F09</td>
<td style="text-align: center; border: 0px; background-color: #79627E; width: 32px; height: 32px; color: #fff;">$F0A</td>
<td style="text-align: center; border: 0px; background-color: #7B6277; width: 32px; height: 32px; color: #fff;">$F0B</td>
<td style="text-align: center; border: 0px; background-color: #7B6661; width: 32px; height: 32px; color: #fff;">$F0C</td>
<td style="text-align: center; border: 0px; background-color: #787878; width: 32px; height: 32px; color: #fff;">$F0D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F0E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F0F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #777777; width: 32px; height: 32px; color: #fff;">$F10</td>
<td style="text-align: center; border: 0px; background-color: #928370; width: 32px; height: 32px; color: #000;">$F11</td>
<td style="text-align: center; border: 0px; background-color: #898967; width: 32px; height: 32px; color: #000;">$F12</td>
<td style="text-align: center; border: 0px; background-color: #818C68; width: 32px; height: 32px; color: #000;">$F13</td>
<td style="text-align: center; border: 0px; background-color: #788F72; width: 32px; height: 32px; color: #000;">$F14</td>
<td style="text-align: center; border: 0px; background-color: #74907A; width: 32px; height: 32px; color: #000;">$F15</td>
<td style="text-align: center; border: 0px; background-color: #768B90; width: 32px; height: 32px; color: #000;">$F16</td>
<td style="text-align: center; border: 0px; background-color: #778898; width: 32px; height: 32px; color: #000;">$F17</td>
<td style="text-align: center; border: 0px; background-color: #8382A4; width: 32px; height: 32px; color: #000;">$F18</td>
<td style="text-align: center; border: 0px; background-color: #887EA1; width: 32px; height: 32px; color: #000;">$F19</td>
<td style="text-align: center; border: 0px; background-color: #927B99; width: 32px; height: 32px; color: #000;">$F1A</td>
<td style="text-align: center; border: 0px; background-color: #947B90; width: 32px; height: 32px; color: #000;">$F1B</td>
<td style="text-align: center; border: 0px; background-color: #958178; width: 32px; height: 32px; color: #000;">$F1C</td>
<td style="text-align: center; border: 0px; background-color: #929292; width: 32px; height: 32px; color: #000;">$F1D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F1E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F1F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #939391; width: 32px; height: 32px; color: #000;">$F20</td>
<td style="text-align: center; border: 0px; background-color: #AD9C8F; width: 32px; height: 32px; color: #000;">$F21</td>
<td style="text-align: center; border: 0px; background-color: #A2A385; width: 32px; height: 32px; color: #000;">$F22</td>
<td style="text-align: center; border: 0px; background-color: #9BA684; width: 32px; height: 32px; color: #000;">$F23</td>
<td style="text-align: center; border: 0px; background-color: #92A98D; width: 32px; height: 32px; color: #000;">$F24</td>
<td style="text-align: center; border: 0px; background-color: #8FA797; width: 32px; height: 32px; color: #000;">$F25</td>
<td style="text-align: center; border: 0px; background-color: #8FA4A7; width: 32px; height: 32px; color: #000;">$F26</td>
<td style="text-align: center; border: 0px; background-color: #92A2B1; width: 32px; height: 32px; color: #000;">$F27</td>
<td style="text-align: center; border: 0px; background-color: #9D9CBB; width: 32px; height: 32px; color: #000;">$F28</td>
<td style="text-align: center; border: 0px; background-color: #A298BA; width: 32px; height: 32px; color: #000;">$F29</td>
<td style="text-align: center; border: 0px; background-color: #AD96B2; width: 32px; height: 32px; color: #000;">$F2A</td>
<td style="text-align: center; border: 0px; background-color: #AE96A9; width: 32px; height: 32px; color: #000;">$F2B</td>
<td style="text-align: center; border: 0px; background-color: #AE9B95; width: 32px; height: 32px; color: #000;">$F2C</td>
<td style="text-align: center; border: 0px; background-color: #AAADAA; width: 32px; height: 32px; color: #000;">$F2D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F2E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F2F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #ACACAE; width: 32px; height: 32px; color: #000;">$F30</td>
<td style="text-align: center; border: 0px; background-color: #C5B5A5; width: 32px; height: 32px; color: #000;">$F31</td>
<td style="text-align: center; border: 0px; background-color: #BBBC9C; width: 32px; height: 32px; color: #000;">$F32</td>
<td style="text-align: center; border: 0px; background-color: #B5C09D; width: 32px; height: 32px; color: #000;">$F33</td>
<td style="text-align: center; border: 0px; background-color: #ABC2A5; width: 32px; height: 32px; color: #000;">$F34</td>
<td style="text-align: center; border: 0px; background-color: #A9C1B1; width: 32px; height: 32px; color: #000;">$F35</td>
<td style="text-align: center; border: 0px; background-color: #A9BEC3; width: 32px; height: 32px; color: #000;">$F36</td>
<td style="text-align: center; border: 0px; background-color: #ABBBCA; width: 32px; height: 32px; color: #000;">$F37</td>
<td style="text-align: center; border: 0px; background-color: #B6B5D4; width: 32px; height: 32px; color: #000;">$F38</td>
<td style="text-align: center; border: 0px; background-color: #BCB1D5; width: 32px; height: 32px; color: #000;">$F39</td>
<td style="text-align: center; border: 0px; background-color: #C6AFCA; width: 32px; height: 32px; color: #000;">$F3A</td>
<td style="text-align: center; border: 0px; background-color: #C7B0C3; width: 32px; height: 32px; color: #000;">$F3B</td>
<td style="text-align: center; border: 0px; background-color: #C7B4AE; width: 32px; height: 32px; color: #000;">$F3C</td>
<td style="text-align: center; border: 0px; background-color: #C6C6C6; width: 32px; height: 32px; color: #000;">$F3D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F3E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F3F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #C7C7C7; width: 32px; height: 32px; color: #000;">$F40</td>
<td style="text-align: center; border: 0px; background-color: #E0D0C0; width: 32px; height: 32px; color: #000;">$F41</td>
<td style="text-align: center; border: 0px; background-color: #D6D6B6; width: 32px; height: 32px; color: #000;">$F42</td>
<td style="text-align: center; border: 0px; background-color: #CFD9B6; width: 32px; height: 32px; color: #000;">$F43</td>
<td style="text-align: center; border: 0px; background-color: #C5DCBF; width: 32px; height: 32px; color: #000;">$F44</td>
<td style="text-align: center; border: 0px; background-color: #C3DBCB; width: 32px; height: 32px; color: #000;">$F45</td>
<td style="text-align: center; border: 0px; background-color: #C3D8DB; width: 32px; height: 32px; color: #000;">$F46</td>
<td style="text-align: center; border: 0px; background-color: #C6D6E5; width: 32px; height: 32px; color: #000;">$F47</td>
<td style="text-align: center; border: 0px; background-color: #D0CFEE; width: 32px; height: 32px; color: #000;">$F48</td>
<td style="text-align: center; border: 0px; background-color: #D6CCEE; width: 32px; height: 32px; color: #000;">$F49</td>
<td style="text-align: center; border: 0px; background-color: #E0C9E3; width: 32px; height: 32px; color: #000;">$F4A</td>
<td style="text-align: center; border: 0px; background-color: #E1CADC; width: 32px; height: 32px; color: #000;">$F4B</td>
<td style="text-align: center; border: 0px; background-color: #E1CEC6; width: 32px; height: 32px; color: #000;">$F4C</td>
<td style="text-align: center; border: 0px; background-color: #DEDEDC; width: 32px; height: 32px; color: #000;">$F4D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F4E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F4F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #DEDEDE; width: 32px; height: 32px; color: #000;">$F50</td>
<td style="text-align: center; border: 0px; background-color: #F7E7D8; width: 32px; height: 32px; color: #000;">$F51</td>
<td style="text-align: center; border: 0px; background-color: #EDEECF; width: 32px; height: 32px; color: #000;">$F52</td>
<td style="text-align: center; border: 0px; background-color: #E6F1CF; width: 32px; height: 32px; color: #000;">$F53</td>
<td style="text-align: center; border: 0px; background-color: #DFF3D8; width: 32px; height: 32px; color: #000;">$F54</td>
<td style="text-align: center; border: 0px; background-color: #DCF3E3; width: 32px; height: 32px; color: #000;">$F55</td>
<td style="text-align: center; border: 0px; background-color: #DCEFF3; width: 32px; height: 32px; color: #000;">$F56</td>
<td style="text-align: center; border: 0px; background-color: #DDEDFA; width: 32px; height: 32px; color: #000;">$F57</td>
<td style="text-align: center; border: 0px; background-color: #EAE8FE; width: 32px; height: 32px; color: #000;">$F58</td>
<td style="text-align: center; border: 0px; background-color: #EEE6FE; width: 32px; height: 32px; color: #000;">$F59</td>
<td style="text-align: center; border: 0px; background-color: #F7E0FB; width: 32px; height: 32px; color: #000;">$F5A</td>
<td style="text-align: center; border: 0px; background-color: #F7E2F3; width: 32px; height: 32px; color: #000;">$F5B</td>
<td style="text-align: center; border: 0px; background-color: #F8E6E4; width: 32px; height: 32px; color: #000;">$F5C</td>
<td style="text-align: center; border: 0px; background-color: #F3F8F8; width: 32px; height: 32px; color: #000;">$F5D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F5E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F5F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #F7F7F7; width: 32px; height: 32px; color: #000;">$F60</td>
<td style="text-align: center; border: 0px; background-color: #FEFEF2; width: 32px; height: 32px; color: #000;">$F61</td>
<td style="text-align: center; border: 0px; background-color: #FFFFEC; width: 32px; height: 32px; color: #000;">$F62</td>
<td style="text-align: center; border: 0px; background-color: #FDFEEC; width: 32px; height: 32px; color: #000;">$F63</td>
<td style="text-align: center; border: 0px; background-color: #F9FFF1; width: 32px; height: 32px; color: #000;">$F64</td>
<td style="text-align: center; border: 0px; background-color: #F3FFF9; width: 32px; height: 32px; color: #000;">$F65</td>
<td style="text-align: center; border: 0px; background-color: #F4FFFF; width: 32px; height: 32px; color: #000;">$F66</td>
<td style="text-align: center; border: 0px; background-color: #F6FEFF; width: 32px; height: 32px; color: #000;">$F67</td>
<td style="text-align: center; border: 0px; background-color: #FEFEFE; width: 32px; height: 32px; color: #000;">$F68</td>
<td style="text-align: center; border: 0px; background-color: #FEFCFF; width: 32px; height: 32px; color: #000;">$F69</td>
<td style="text-align: center; border: 0px; background-color: #FDFAFE; width: 32px; height: 32px; color: #000;">$F6A</td>
<td style="text-align: center; border: 0px; background-color: #FFF9FE; width: 32px; height: 32px; color: #000;">$F6B</td>
<td style="text-align: center; border: 0px; background-color: #FEFDF9; width: 32px; height: 32px; color: #000;">$F6C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F6D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F6E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F6F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F70</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F71</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFE; width: 32px; height: 32px; color: #000;">$F72</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFC; width: 32px; height: 32px; color: #000;">$F73</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F74</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F75</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F76</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F77</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F78</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F79</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F7A</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F7B</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F7C</td>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F7D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F7E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F7F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #FFFFFF; width: 32px; height: 32px; color: #000;">$F80</td>
<td style="text-align: center; border: 0px; background-color: #009BFD; width: 32px; height: 32px; color: #fff;">$F81</td>
<td style="text-align: center; border: 0px; background-color: #4665FE; width: 32px; height: 32px; color: #fff;">$F82</td>
<td style="text-align: center; border: 0px; background-color: #AD32FD; width: 32px; height: 32px; color: #fff;">$F83</td>
<td style="text-align: center; border: 0px; background-color: #FF00FE; width: 32px; height: 32px; color: #fff;">$F84</td>
<td style="text-align: center; border: 0px; background-color: #FE00EF; width: 32px; height: 32px; color: #fff;">$F85</td>
<td style="text-align: center; border: 0px; background-color: #FE2E0B; width: 32px; height: 32px; color: #fff;">$F86</td>
<td style="text-align: center; border: 0px; background-color: #FE5C02; width: 32px; height: 32px; color: #000;">$F87</td>
<td style="text-align: center; border: 0px; background-color: #A69900; width: 32px; height: 32px; color: #000;">$F88</td>
<td style="text-align: center; border: 0px; background-color: #4EC400; width: 32px; height: 32px; color: #000;">$F89</td>
<td style="text-align: center; border: 0px; background-color: #00FD02; width: 32px; height: 32px; color: #000;">$F8A</td>
<td style="text-align: center; border: 0px; background-color: #00F809; width: 32px; height: 32px; color: #000;">$F8B</td>
<td style="text-align: center; border: 0px; background-color: #00C5F0; width: 32px; height: 32px; color: #000;">$F8C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F8D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F8E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F8F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F90</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F91</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F92</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F93</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F94</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F95</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F96</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F97</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F98</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F99</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F9A</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F9B</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F9C</td>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$F9D</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F9E</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$F9F</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #010101; width: 32px; height: 32px; color: #fff;">$FA0</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA2</td>
<td style="text-align: center; border: 0px; background-color: #000001; width: 32px; height: 32px; color: #fff;">$FA3</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA4</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA5</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA6</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FA7</td>
<td style="text-align: center; border: 0px; background-color: #010008; width: 32px; height: 32px; color: #fff;">$FA8</td>
<td style="text-align: center; border: 0px; background-color: #000005; width: 32px; height: 32px; color: #fff;">$FA9</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FAA</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FAB</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FAC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FAD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FAE</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$FAF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB0</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB1</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB2</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB3</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB4</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB5</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB6</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FB7</td>
<td style="text-align: center; border: 0px; background-color: #000106; width: 32px; height: 32px; color: #fff;">$FB8</td>
<td style="text-align: center; border: 0px; background-color: #00000A; width: 32px; height: 32px; color: #fff;">$FB9</td>
<td style="text-align: center; border: 0px; background-color: #000001; width: 32px; height: 32px; color: #fff;">$FBA</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FBB</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FBC</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FBD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FBE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FBF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FC0</td>
<td style="text-align: center; border: 0px; background-color: #0A0001; width: 32px; height: 32px; color: #fff;">$FC1</td>
<td style="text-align: center; border: 0px; background-color: #020102; width: 32px; height: 32px; color: #fff;">$FC2</td>
<td style="text-align: center; border: 0px; background-color: #000500; width: 32px; height: 32px; color: #fff;">$FC3</td>
<td style="text-align: center; border: 0px; background-color: #010800; width: 32px; height: 32px; color: #fff;">$FC4</td>
<td style="text-align: center; border: 0px; background-color: #010800; width: 32px; height: 32px; color: #fff;">$FC5</td>
<td style="text-align: center; border: 0px; background-color: #000308; width: 32px; height: 32px; color: #fff;">$FC6</td>
<td style="text-align: center; border: 0px; background-color: #000110; width: 32px; height: 32px; color: #fff;">$FC7</td>
<td style="text-align: center; border: 0px; background-color: #000018; width: 32px; height: 32px; color: #fff;">$FC8</td>
<td style="text-align: center; border: 0px; background-color: #02001D; width: 32px; height: 32px; color: #fff;">$FC9</td>
<td style="text-align: center; border: 0px; background-color: #0C0010; width: 32px; height: 32px; color: #fff;">$FCA</td>
<td style="text-align: center; border: 0px; background-color: #0E0006; width: 32px; height: 32px; color: #fff;">$FCB</td>
<td style="text-align: center; border: 0px; background-color: #0E0000; width: 32px; height: 32px; color: #fff;">$FCC</td>
<td style="text-align: center; border: 0px; background-color: #0C0D08; width: 32px; height: 32px; color: #fff;">$FCD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FCE</td>
<td style="text-align: center; border: 0px; background-color: #000002; width: 32px; height: 32px; color: #fff;">$FCF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #0C0C0C; width: 32px; height: 32px; color: #fff;">$FD0</td>
<td style="text-align: center; border: 0px; background-color: #281708; width: 32px; height: 32px; color: #fff;">$FD1</td>
<td style="text-align: center; border: 0px; background-color: #1E1E01; width: 32px; height: 32px; color: #fff;">$FD2</td>
<td style="text-align: center; border: 0px; background-color: #162002; width: 32px; height: 32px; color: #fff;">$FD3</td>
<td style="text-align: center; border: 0px; background-color: #0E2508; width: 32px; height: 32px; color: #fff;">$FD4</td>
<td style="text-align: center; border: 0px; background-color: #0C2510; width: 32px; height: 32px; color: #fff;">$FD5</td>
<td style="text-align: center; border: 0px; background-color: #0C2126; width: 32px; height: 32px; color: #fff;">$FD6</td>
<td style="text-align: center; border: 0px; background-color: #0E1E2D; width: 32px; height: 32px; color: #fff;">$FD7</td>
<td style="text-align: center; border: 0px; background-color: #171837; width: 32px; height: 32px; color: #fff;">$FD8</td>
<td style="text-align: center; border: 0px; background-color: #1E1433; width: 32px; height: 32px; color: #fff;">$FD9</td>
<td style="text-align: center; border: 0px; background-color: #26102D; width: 32px; height: 32px; color: #fff;">$FDA</td>
<td style="text-align: center; border: 0px; background-color: #291025; width: 32px; height: 32px; color: #fff;">$FDB</td>
<td style="text-align: center; border: 0px; background-color: #29170D; width: 32px; height: 32px; color: #fff;">$FDC</td>
<td style="text-align: center; border: 0px; background-color: #272727; width: 32px; height: 32px; color: #fff;">$FDD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FDE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FDF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #272727; width: 32px; height: 32px; color: #fff;">$FE0</td>
<td style="text-align: center; border: 0px; background-color: #423222; width: 32px; height: 32px; color: #fff;">$FE1</td>
<td style="text-align: center; border: 0px; background-color: #393A1A; width: 32px; height: 32px; color: #fff;">$FE2</td>
<td style="text-align: center; border: 0px; background-color: #313D1A; width: 32px; height: 32px; color: #fff;">$FE3</td>
<td style="text-align: center; border: 0px; background-color: #283F22; width: 32px; height: 32px; color: #fff;">$FE4</td>
<td style="text-align: center; border: 0px; background-color: #263F2A; width: 32px; height: 32px; color: #fff;">$FE5</td>
<td style="text-align: center; border: 0px; background-color: #263B40; width: 32px; height: 32px; color: #fff;">$FE6</td>
<td style="text-align: center; border: 0px; background-color: #293949; width: 32px; height: 32px; color: #fff;">$FE7</td>
<td style="text-align: center; border: 0px; background-color: #323352; width: 32px; height: 32px; color: #fff;">$FE8</td>
<td style="text-align: center; border: 0px; background-color: #382F51; width: 32px; height: 32px; color: #fff;">$FE9</td>
<td style="text-align: center; border: 0px; background-color: #422B4C; width: 32px; height: 32px; color: #fff;">$FEA</td>
<td style="text-align: center; border: 0px; background-color: #452B3F; width: 32px; height: 32px; color: #fff;">$FEB</td>
<td style="text-align: center; border: 0px; background-color: #44312B; width: 32px; height: 32px; color: #fff;">$FEC</td>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$FED</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FEE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FEF</td>
</tr>
<tr>
<td style="text-align: center; border: 0px; background-color: #434343; width: 32px; height: 32px; color: #fff;">$FF0</td>
<td style="text-align: center; border: 0px; background-color: #5C4D3A; width: 32px; height: 32px; color: #fff;">$FF1</td>
<td style="text-align: center; border: 0px; background-color: #545337; width: 32px; height: 32px; color: #fff;">$FF2</td>
<td style="text-align: center; border: 0px; background-color: #4C5834; width: 32px; height: 32px; color: #fff;">$FF3</td>
<td style="text-align: center; border: 0px; background-color: #435A3D; width: 32px; height: 32px; color: #fff;">$FF4</td>
<td style="text-align: center; border: 0px; background-color: #405944; width: 32px; height: 32px; color: #fff;">$FF5</td>
<td style="text-align: center; border: 0px; background-color: #415658; width: 32px; height: 32px; color: #fff;">$FF6</td>
<td style="text-align: center; border: 0px; background-color: #425262; width: 32px; height: 32px; color: #fff;">$FF7</td>
<td style="text-align: center; border: 0px; background-color: #4E4D6D; width: 32px; height: 32px; color: #fff;">$FF8</td>
<td style="text-align: center; border: 0px; background-color: #53496B; width: 32px; height: 32px; color: #fff;">$FF9</td>
<td style="text-align: center; border: 0px; background-color: #5C4562; width: 32px; height: 32px; color: #fff;">$FFA</td>
<td style="text-align: center; border: 0px; background-color: #5F465D; width: 32px; height: 32px; color: #fff;">$FFB</td>
<td style="text-align: center; border: 0px; background-color: #5E4B45; width: 32px; height: 32px; color: #fff;">$FFC</td>
<td style="text-align: center; border: 0px; background-color: #5E5E5E; width: 32px; height: 32px; color: #fff;">$FFD</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FFE</td>
<td style="text-align: center; border: 0px; background-color: #000000; width: 32px; height: 32px; color: #fff;">$FFF</td>
</tr>
</tbody>
</table>
