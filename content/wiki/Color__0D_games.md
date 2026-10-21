---
title: "Color $0D games"
url: "/wiki/Color_$0D_games/"
description: "On an NES, the palette color $0D causes the signal to drop below the normal black level. This low voltage signal is sometimes mistaken by televisions for blanking signals, which can cause an unstable…"
wikiPage: "Color__0D_games"
---

On an NES, the palette color \$0D causes the signal to drop below the normal black level. This low voltage signal is sometimes mistaken by televisions for blanking signals, which can cause an unstable picture, or total picture loss on some devices. Other devices seem to process the signal without problem.

## Games {#Games}

<table class="wikitable sortable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="width: 18%">Game</th>
<th class="unsortable">Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>Contra 100 in 1</em></td>
<td>Uses both for background color and one of the sprites' color; to apply a patch, change values at those offsets in ROM from $0D to $0F:
<p>$388B, $388F, $3893, $3897, $389B, $389C, $389F, $38A3, $38A7</p></td>
</tr>
<tr>
<td><em>Contra 168 in 1</em></td>
<td>Uses both for background color and one of the sprites' color; to apply a patch, change values at those offsets in ROM from $0D to $0F:
<p>$2451, $2455, $2459, $245D, $2461, $2462, 2465, $2469, $246D</p></td>
</tr>
<tr>
<td><em>Bee 52</em></td>
<td></td>
</tr>
<tr>
<td><em>Castelian</em></td>
<td></td>
</tr>
<tr>
<td><em>Cybernoid</em></td>
<td>$0D is used as the background color.</td>
</tr>
<tr>
<td><em>The Fantastic Adventures of Dizzy</em></td>
<td></td>
</tr>
<tr>
<td><em>FireHawk</em> (USA)</td>
<td>Used in both the background and sprites. The European release (with NTSC video detection) corrects this.</td>
</tr>
<tr>
<td><em>Game Genie</em></td>
<td>The code entry screen uses it for its background.</td>
</tr>
<tr>
<td><em>The Immortal</em></td>
<td>Also uses all three de-emphasis bits to compensate for the the user cranking up the TV set's brightness so that regular black ($xE/$xF) can be used <a href="https://youtu.be/IMFoci7S2hs">as a first darker shade of gray, and color $1D as an even darker shade of gray, while color $0D is used as a black background color</a>.</td>
</tr>
<tr>
<td><em>Indiana Jones and the Last Crusade</em> (Taito)</td>
<td>Used as the background color in the motorcycle level.</td>
</tr>
<tr>
<td><em>Indiana Jones and the Last Crusade</em> (Ubisoft)</td>
<td></td>
</tr>
<tr>
<td><em>Maniac Mansion</em> (USA)</td>
<td></td>
</tr>
<tr>
<td><em>Micro Machines</em></td>
<td></td>
</tr>
<tr>
<td><em>MIG-29 Soviet Fighter</em></td>
<td></td>
</tr>
<tr>
<td><em>Quattro Sports</em></td>
<td></td>
</tr>
<tr>
<td><em>Quattro Adventures</em></td>
<td></td>
</tr>
<tr>
<td><em>Skate or Die 2</em></td>
<td>Used as the background color during the introduction cutscene sequence.</td>
</tr>
<tr>
<td><em>The Super Shinobi</em></td>
<td>Unlicensed clone of Shinobi III.</td>
</tr>
<tr>
<td><em>Teenage Mutant Ninja Turtles</em></td>
<td>Uses it for black outlines on sprites, the lack of large areas of this color mitigates the problem.</td>
</tr>
<tr>
<td><em>The Three Stooges</em></td>
<td>Uses $xD colors that are turned to $0D during fades.</td>
</tr>
</tbody>
</table>

## Effects {#Effects}

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Color_0d_-_contra_168.jpg"><img src="../../wiki-images/Color_0d_-_contra_168.jpg" width="300" height="217" /></a>
<figcaption>Modern LCD TV and modern famiclone (probably RS-35) showing distorded image in <em>Contra 168-in-1</em> cartridge</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Color_0d_-_super_robin_hood.jpg"><img src="../../wiki-images/Color_0d_-_super_robin_hood.jpg" width="300" height="225" /></a>
<figcaption>Different modern LCD TV with digital famiclone (connected via HDMI) shows $0D color to be displayed as gray in <em>Super Robin Hood</em> game</figcaption>
</figure>

<figure typeof="mw:File/Frame">
<a href="../../wiki-images/Color_0d_-_fantblue.gif"><img src="../../wiki-images/Color_0d_-_fantblue.gif" width="320" height="240" /></a>
<figcaption>Modern 29" CRT TV with IQ502 rev 2 (UM6561 nes-on-chip) famiclone shows periodic loss of v-sync when the paper unrolls at the beginning of <em>Fantastic Adventures of Dizzy</em> game. Note that this TV has no problems with displaying $0D colors with different famiclones, and even using this famiclone - this is the only moment in this game that the problem manifests, though this game uses this color all the time</figcaption>
</figure>

<figure typeof="mw:File/Frame">
<a href="../../wiki-images/Color_0d_-_ultimate_stuntman.gif"><img src="../../wiki-images/Color_0d_-_ultimate_stuntman.gif" width="320" height="240" /></a>
<figcaption>Same TV and console as above, No problems in the beginning screenshots, where $0d is used as universal background colors, while permanent v-sync loss during game, where $0d is used as blinking car's shadow color in <em>Ultimate Stuntman</em> game.</figcaption>
</figure>

Because the signal created by \$0D is outside the specifications for the video signal, there is a lot of variation in how display devices handle it. Here are some possible effects that may be seen when using \$0D:

- \$0D appears the same black as the other black colors (e.g. \$0F).
- \$0D appears slightly darker than other blacks.
- \$0D appears as gray.
- The device renormalizes the range when \$0D appears, slightly brightening all other colours while it is onscreen.
- Wobbly or distorted image from loss of horizontal blanking stability (either permanent or periodic)
- Loss of vertical blanking stability.
- Total loss of picture.

These effects are more likely to occur when color \$0D is used with the de-emphasis bits enabled, such as in *The Immortal*, as seen in these [example](https://youtu.be/3fhyX3HdVcg) [videos](https://youtu.be/4Kli09cSGOU).

On consoles with an RGB PPU like the Famicom Titler or [Sharp C1](../Sharp_C1/) TV, Color \$0D is simply a "normal" black palette entry identical to \$0F, so systems with an RGB PPU are immune to causing video output problems.

## Workarounds {#Workarounds}

- Patch the game code to replace all the \$0D writes with \$0F. This likely requires changes to many bytes in the ROM, so a Game Genie (limited to 3 changes) would not be sufficient. Instead, updated ROMs or a more capable pass-through patching device would likely be required
- If your TV has digital inputs (for example - HDMI), use RCA to HDMI adapter, whose analog to digital converter might cope better with the out-of-spec video signal
- Modify the console to use a different video output stage that adjusts the voltages so that the TV may accept it, such as by omitting the NPN transistor follower part of the video amplifier. However, the TV may ignore this voltage offset
- Change your TV, console and/or video adapter (if you are using one)

## Tests {#Tests}

- [Palette test ROM](http://forums.nesdev.org/viewtopic.php?f=3&t=13264) - Displays NES palette, and can toggle \$0D display.
- [NESPix](http://forums.nesdev.org/viewtopic.php?f=22&t=11255) - Native graphics editor that allows use of \$0D, and can test it in various visual arrangements.
- [240p test suite](http://forums.nesdev.org/viewtopic.php?f=22&t=13394) - TV testing program. Test cards with \$0D include PLUGE, SMPTE color bars, Solid color screen, and IRE. PLUGE also includes emphasized \$0D.

## See also {#See_also}

- [NTSC video](../NTSC_video/)
- [PPU palettes](../PPU_palettes/)
- [Colour-emphasis games](../Colour-emphasis_games/)
