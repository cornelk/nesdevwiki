---
title: "Audio drivers"
url: "/wiki/Audio_drivers/"
description: "An audio driver or replayer is a program that runs on a game console to play music. An audio driver for NES reads music sequences and instrument definitions from the ROM and writes to the APU ports,…"
wikiPage: "Audio_drivers"
---

An **audio driver** or **replayer** is a program that runs on a game console to play music. An audio driver for NES reads music sequences and instrument definitions from the ROM and writes to the [APU](../APU/) ports, usually once per frame.

## Tracker replayers {#Tracker_replayers}

NSFs exported from FamiTracker and similar trackers contain replayers that can handle all tracker features, but they are not optimized for game use. They use a lot of ROM and RAM, CPU time, and can't handle sound effects. In addition, the embedded replayer may change in size/capability depending on the exported track's features (i.e., expansion audio support, bankswitching)

- [FamiTracker](http://famitracker.com/)
  - Forks (different features / behaviors)
    - [Dn-FamiTracker](https://github.com/Dn-Programming-Core-Management/Dn-FamiTracker)
    - [0CC-FamiTracker](https://github.com/HertzDevil/0CC-FT-NSF-Driver)

<!-- -->

- [Musetracker](https://fo.aspekt.fi/)

<!-- -->

- [Nerdtracker II](http://nesdev.org/nt2/)

<!-- -->

- [PPMCK/MML](http://ppmck.web.fc2.com/ppmck.html)

<!-- -->

- [Famistudio](https://famistudio.org/)

## Game replayers {#Game_replayers}

Replayers intended for use in games have limits on the instrument, effect, and expansion features they can use. Most can't handle expansions at all, as NES games in an unmodified NES cannot use expansion audio. They usually operate by translating a text export into assembly language data files that are included into the game's source code.

### Famistudio music engine {#Famistudio_music_engine}

[Famistudio music engine](http://forums.nesdev.org/viewtopic.php?f=2&t=20825&p=256016) by bleubleu

Famistudio's driver is originally of FamiTone2 genealogy, but has been massively reworked. Some of its key features is being highly configurable to cater for different needs, and extensive expansion audio support.

- Meant to go with its own music tool (Famistudio). Famitracker projects are importable.
- NESASM, CA65, ASM6 versions

Basic features:

- Pulse channels, triangle channel, noise channels
- C0-B7 note range (96 notes)
- Instruments with duty cycle, volume, pitch and arpeggio envelopes
- Absolute and relative pitch envelopes
- Looping sections in envelopes
- Release points for volume envelopes
- Ability to change the speed (FamiTracker tempo mode only)
- Ability to loop over a portion of the song
- Up to 64 instruments per export, an export can consist of as little as 1 song, or as many as 17
- Can enable/disable features to save ROM and RAM
- VRC6, VRC7, FDS, Sunsoft 5B, Namco 163 support (only one at a time)
- PAL/NTSC playback
- DPCM
- Sound effect support (configurable number of streams)
- Blaarg Smooth Vibrato technique to eliminate "pops" on square channels (incompatible with SFX at the moment)
- FamiTracker/FamiStudio tempo mode.
- Volume track
- Fine pitch track
- Slide notes
- Vibrato effect
- Arpeggio effect (not arpeggio envelopes in instruments, which are always available)

### FamiTone2 {#FamiTone2}

[FamiTone2](https://shiru.untergrund.net/code.shtml) by Shiru

- Notes limited to C-1 through D-6
- Instrument: volume, arpeggio, and pitch envelopes. Duty envelopes longer than 1 frame unsupported. Pitch limited to accumulated distance of 63 units in each direction. No release phase.
- No volume column
- DPCM for instrument 0 only
- Limit of 64 instruments, 17 songs, and 16 KiB of DPCM

Effects:

- `Fxx` (speed / tempo)
- `Dxx` (pattern cut / skip to next frame and start at xx)
- `Bxx` (loop / jump to frame)

Requirements:

- 3 bytes ZP
- 186 bytes other RAM
- 1636 bytes ROM (source: README)

### FamiTone 5.0 {#FamiTone_5.0}

[FamiTone 5.0](https://github.com/nesdoug/famitone5.0) by dougeff

Modification of FamiTone2 that adds

- Adds volume column support
- Note range A0 - B7
- Duty cycle envelopes
- Sound effects that are bigger than 256 bytes

and adds effects:

- `1xx` (portamento up)
- `2xx` (portamento down)
- `3xx` (glissando)
- `4xy` (vibrato)
- `Qxx` (slide up to note)
- `Rxx` (slide down to note)

### GGSound {#GGSound}

<https://github.com/gradualgames/ggsound> by Gradual Games

- NESASM3, ASM6, and ca65 syntax supported
- Note range: C0 - B7
- Instrument: volume, arpeggio, pitch, and duty envelopes, all looping
- SFX playback on up to two channels
- Pause / Unpause
- Note cuts
- 128 instruments (each with its own set of envelopes)
- 128 songs
- 128 SFX

Effects:

- `Bxx` (loop / jump to frame).

Difference vs famitracker: Loops each channel individually, so must be placed in all channels for proper song looping.

Requirements:

- 66 bytes ZP (57 without DPCM)
- 168 bytes other RAM (144 without DPCM)
- 3048 bytes ROM (out of date, needs test on current version)

### Lizard music engine {#Lizard_music_engine}

[Lizard music engine](http://forums.nesdev.org/viewtopic.php?f=2&t=18957&p=239567) by Rainwarrior

Features (subset of Famitracker):

- Volume column supported
- No DPCM
- No Hi-Pitch macros (but regular pitch works)
- Tempo fixed at 150 (use Speed instead)
- SFX on square 1 and/or noise channel

Effects:

- `Bxx` (loop / jump to frame)
- `D00` (Skip to next frame; no row support)
- `F0x` (just speed, not tempo)

Requirements:

- 22 bytes ZP
- 105 bytes other RAM
- 2152 bytes of code and tables
- ~1800 cycles per frame

### NSD.Lib {#NSD.Lib}

[NSD.Lib](https://github.com/Shaw02/nsdlib/) by S.W.

Focuses on providing as many features as possible while providing low CPU and memory usage (including features many other drivers do not support). It should be noted that the driver defaults to A=442 tuning, although a [tuning table generator](https://github.com/pinobatch/little-things-nes/tree/master/nsdl-tune) exists if desired.

Best known for its application in games like *8Bit Music Power*, *Astro Ninja Man*, and *Kira Kira Star Night*.

Features:

- Designed with low CPU usage and small code size in mind.
- Sequence data format optimized to minimize ROM space.
- Library functions callable from both Assembly and C (cc65 \_\_fastcall\_\_ convention).
- Can output either assembly source code for CA65 or NSF music files.
- Includes a C++ MML compiler that converts MML into sequence data usable by the driver.
- Provides a variety of effects and commands for expressive music creation.

Effects:

- Tempo (1-255)
- Note, Rest
- Octave (1-8)
- Volume (0-15)
- Detune(Pitch)
- Voice
- Envelope (Volume, Detune(Pitch), Note, Voice)
- Portamento
- Quantize
- Repeat
- Etc..

### Penguin {#Penguin}

[Penguin](https://github.com/pubby/penguin) by pubby

- Constant cycle count of 790, allowing use in a raster effect
- Plays sound effects without using additional cycles
- Famitracker Exported
- [Allegedly](https://forums.nesdev.org/viewtopic.php?p=205095#p205095) has similar features to FamiTone2, but also has duty envelopes
- No DPCM
- 12 bytes ZP
- 86 bytes other RAM
- Music data not particularly size-optimized
- Sound effects are expensive in terms of size
- Tempo fixed at 150 (use Speed instead)
- Minimum speed of 4

Effects:

- `D00` (terminate pattern)
- Sound effects support all effects

### <span id="PUF_.28NesFab.29"></span>PUF (NesFab) {#PUF_.28NesFab.29}

[PUF](https://github.com/pubby/nesfab) by pubby

This audio driver is included with the NesFab programming language, with built-in conversion support. It is likely too difficult to use outside of NesFab. The driver is based on the Penguin driver, also by pubby.

- Low cycle count
- Plays sound effects without using many additional cycles
- Famitracker Exported
- Has similar features to FamiTone2, but also has duty envelopes
- Supports DPCM
- Supports MMC5 and Rainbow (VRC6) expansion audio
- Music data mildly optimized
- Sound effects are expensive in terms of size
- Tempo fixed at 150 (use Speed instead)
- Supports any speed, but speeds below 4 have a processing cost

Effects:

- `D00` (terminate pattern)
- Sound effects support all effects

### Pently {#Pently}

[Pently](https://github.com/pinobatch/pently) by Damian Yerrick

- Notes limited to A-0 through C-7; changeable at build time
- Instrument: volume, duty, and (absolute) arpeggio envelopes. Pitch envelopes unsupported. The volume envelope can't loop, and the duty and arpeggio envelopes stop at the end of the volume envelope. No release phase.
- Volume column limited to 4 levels
- No DPCM
- Limit of 51 instruments, 25 drums, and 255 patterns
- tempo ("Rows per minute" tempo model allows runtime correction for NTSC, PAL, and Dendy)

Effects:

- `45x` (vibrato)
- `3xx` (portamento)
- `Sxx` grace note / delayed note
- `Gxx` grace note / delayed cut
- legato (change note pitch without restarting envelope)
- `0xy` (arpeggio with 1- or 2-tick scheme)

<!-- -->

- and loop (`Bxx`)
- Vibrato and portamento use "linear pitch" model similar to that of [0CC-FamiTracker](https://github.com/HertzDevil/0CC-FamiTracker)

Requirements:

- 32 bytes ZP
- 112 bytes other RAM
- 1918 bytes ROM
- Some effects are space-intensive and can be disabled at build time through a configuration file to reduce ROM and RAM size. [A feature set comparable to FamiTone2](https://forums.nesdev.org/viewtopic.php?p=205093#p205093) uses 1283 bytes of ROM.
- Compatible with ca65. In 2019, an experimental ASM6 port was added. Python 3 is used to preprocess the score and generate the RAM map.

The native input format ("Pently score") is inspired by MML and intended to be familiar to users of PPMCK or LilyPond. Conversion from FT text export is through [ft2pently](https://github.com/NovaSquirrel/ft2pently) by NovaSquirrel. As Pently score was originally intended for human writability, some features don't map well onto FamiTracker features, requiring manual configuration of ft2pently or manual editing of the Pently score it produces:

- Pattern start has row granularity, allowing a pattern such as a drum fill to replace the end of another pattern.
- Pattern length is not fixed, allowing long melodic patterns and short drum loops.
- Patterns are shared among channels. Use this with delayed pattern start for 2-channel echo.
- Patterns can be transposed. Use this for parallel fourths, fifths, or octaves, or for gear changes like that in the *Mermaids of Atlantis* soundtrack.
- Drum channel instead of noise channel. Drums are played as pairs of sound effects that interrupt a note played on the same channel, making Tim Follin-style triangle+noise drums somewhat easier than with the fixed arpeggio that one would use in FT.
- Envelope divided into an attack phase, similar to FT envelopes, and a decay phase with a constant duty and decreasing volume, similar to NerdTracker II envelopes.
- "Attack track" mode, where the attack phase of a note on MMC5 channel 1 temporarily overrides the note on pulse channel 1. Useful for staccato ostinatos and occasionally 1-channel echo.
- Automatic conversion doesn't use the newest features (linear portamento and 2-frame arpeggio) due to converter author's unfamiliarity with 0CC-FamiTracker.

### OFGS {#OFGS}

<http://offgao.net/ofgs/>

### Sabre {#Sabre}

<https://github.com/CutterCross/Sabre> by CutterCross

Features:

- CA65 and ASM6 syntax supported
- Note range: A0-B7 (full pulse/tri note range)
- Note cuts
- Speed and tempo
- DPCM supported for music (not for SFX)
- Instrument envelopes: Volume, arpeggio (absolute), pitch (relative) and duty.
- Loop points for all supported envelope types
- 63 instruments
- 255 unique envelopes total shared between instruments.
- 256 tracks
- 256 SFX
- 128 patterns per track
- 1 pattern per SFX
- NTSC, PAL, Dendy tempo & period adjustments
- Linear counter trill for Triangle channel
- Play / Pause / Stop Track routines
- Mute / Unmute channel routines

Effects supported:

- `Bxx` (loop / skip to frame)
- `C00` (halt / end song)
- `D00` (skip to next frame; specify row not supported)
- `Fxx` (speed / tempo)
- `Zxx` (set DPCM delta counter (\$4011); affects tri, noise and DPCM volume)

Requirements:

- 42 bytes ZP
- 121 bytes other RAM
- 1749 bytes ROM

Notable usage notes:

- Instruments share a common envelope set.
- Song format doesn't lay down redundant note period data if the instrument or length changes
- Loop/forward points should be in first channel, rather than all channels
