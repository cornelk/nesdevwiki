---
title: "UNIF"
url: "/wiki/UNIF/"
description: "UNIF (Universal NES Image Format) is an alternative format for holding NES and Famicom ROM images."
wikiPage: "UNIF"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
---

**UNIF** (Universal NES Image Format) is an alternative format for holding NES and Famicom ROM images.

Its motivation was to offer more description of the mapper board than the popular [iNES 1.0](../INES/) format, but it suffered from other limiting constraints and a lack of popularity. The format is considered deprecated, replaced by the [NES 2.0](../NES_2.0/) revision of the iNES format, which better addresses the issues it had hoped to solve.

There are a small number of game rips that currently [only exist as UNIF](../UNIF_to_NES_2.0_Mapping/#So-far_unassigned_UNIF_board_names).

## History {#History}

The originator of the UNIF format, Tennessee Carmel-Veilleux, publically abandoned the format in December of 2008, and had its [website](http://web.archive.org/web/20080827175203/http://www.parodius.com/~veilleux/) deleted.

Since 2011 the UNIF standard has been maintained as part of the [libunif project](http://github.com/eteran/libunif), a library for loading and creating UNIF files.

UNIF is currently considered a deprecated standard. Further updates to UNIF are unlikely, and it is recommended that new games or rips requiring special mapper specifications use the [NES 2.0](../NES_2.0/) format instead of attempting to extend UNIF.

## Format {#Format}

The suggested file extension for UNIF is `.unf` or `.unif`.

("LE16" and "LE32" below denote little-endian 16/32-bit integers.)

UNIF images start with a 32-byte header:

| Offset | Length (bytes) | Value |
|----|----|----|
| 0 | 4 | "UNIF" |
| 4 | 4 | LE32, minimum version number required to parse all chunks in file |
| 8 | 24 | all nulls |

The header is followed by any number of Type+Length+Value blocks:

| Offset | Length (bytes) | Value                           |
|--------|----------------|---------------------------------|
| 0      | 4              | Type, varies, defined below     |
| 4      | 4              | LE32, length                    |
| 8      | length         | content encoding varies by type |

### Types {#Types}

The following Types are known:

<table class="wikitable sortable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>Type</th>
<th>Length</th>
<th>Minimum version required</th>
<th>Encoding</th>
<th>Content meaning</th>
</tr>
</thead>
<tbody>
<tr>
<td>MAPR</td>
<td>variable</td>
<td>1</td>
<td>null-terminated UTF-8</td>
<td>A unique human-readable identifier specifying the exact hardware used; <strong>not</strong> an iNES mapper number, and <strong>not</strong> a full text description of the mapper; required</td>
</tr>
<tr>
<td>PRG<em>n</em></td>
<td>variable, usually power of two</td>
<td>4</td>
<td>raw</td>
<td>the contents of the <em>n</em>th PRG ROM; at least PRG0 is required; <em>n</em> is in hexadecimal</td>
</tr>
<tr>
<td>CHR<em>n</em></td>
<td>variable, usually power of two</td>
<td>4</td>
<td>raw</td>
<td>the contents of the <em>n</em>th CHR ROM</td>
</tr>
<tr>
<td>PCK<em>n</em></td>
<td>4</td>
<td>5</td>
<td>LE32</td>
<td>the CRC-32 of the <em>n</em>th PRG ROM</td>
</tr>
<tr>
<td>CCK<em>n</em></td>
<td>4</td>
<td>5</td>
<td>LE32</td>
<td>the CRC-32 of the <em>n</em>th CHR ROM</td>
</tr>
<tr>
<td>NAME</td>
<td>variable</td>
<td>1</td>
<td>null-terminated UTF-8</td>
<td>the name of the game</td>
</tr>
<tr>
<td>WRTR</td>
<td>variable</td>
<td>unknown</td>
<td>null-terminated UTF-8</td>
<td>unofficial, invalid. The name of the dumping software. Should be in a DINF chunk instead</td>
</tr>
<tr>
<td>READ</td>
<td>variable</td>
<td>1</td>
<td>null-terminated UTF-8</td>
<td>comments about the game, especially licensing information for homebrew</td>
</tr>
<tr>
<td>DINF</td>
<td>204</td>
<td>2</td>
<td>special</td>
<td>Dumping information
<table class="tabular">
<thead>
<tr>
<th style="text-align: right;">Offset</th>
<th style="text-align: right;">Length (bytes)</th>
<th>Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: right;">0</td>
<td style="text-align: right;">100</td>
<td>null-terminated UTF-8 dumper name</td>
</tr>
<tr>
<td style="text-align: right;">100</td>
<td style="text-align: right;">1</td>
<td>day-of-month of dump</td>
</tr>
<tr>
<td style="text-align: right;">101</td>
<td style="text-align: right;">1</td>
<td>month-of-year of dump</td>
</tr>
<tr>
<td style="text-align: right;">102</td>
<td style="text-align: right;">2</td>
<td>LE16, year of dump</td>
</tr>
<tr>
<td style="text-align: right;">104</td>
<td style="text-align: right;">100</td>
<td>null-terminated UTF-8 the name of the dumping software or mechanism</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>TVCI</td>
<td>1</td>
<td>6</td>
<td>byte</td>
<td>TV standard compatibility information-
<table>
<tbody>
<tr>
<td style="text-align: right;">0.</td>
<td>NTSC only</td>
</tr>
<tr>
<td style="text-align: right;">1.</td>
<td>PAL only</td>
</tr>
<tr>
<td style="text-align: right;">2.</td>
<td>compatible with both</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>CTRL</td>
<td>1</td>
<td>7</td>
<td>byte</td>
<td>Controllers usable by this game (bitmask)
<table>
<tbody>
<tr>
<td style="text-align: right;">1.</td>
<td><a href="../Standard_controller/">Standard controller</a></td>
</tr>
<tr>
<td style="text-align: right;">2.</td>
<td><a href="../Zapper/">Zapper</a></td>
</tr>
<tr>
<td style="text-align: right;">4.</td>
<td>R.O.B.</td>
</tr>
<tr>
<td style="text-align: right;">8.</td>
<td><a href="../Arkanoid_controller/">Arkanoid controller</a></td>
</tr>
<tr>
<td style="text-align: right;">16.</td>
<td><a href="../Power_Pad/">Power Pad</a></td>
</tr>
<tr>
<td style="text-align: right;">32.</td>
<td><a href="../Four_player_adapters/">Four Score</a></td>
</tr>
<tr>
<td style="text-align: right;">64.</td>
<td>expansion (leave cleared)</td>
</tr>
<tr>
<td style="text-align: right;">128.</td>
<td>expansion (leave cleared)</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>BATR</td>
<td>1</td>
<td>5</td>
<td>byte</td>
<td>Boolean specifying whether the RAM is battery-backed.</td>
</tr>
<tr>
<td>VROR</td>
<td>1</td>
<td>5</td>
<td>byte</td>
<td>"If this chunk is present, then the CHR-ROM area will be considered as RAM even if ROM is present."</td>
</tr>
<tr>
<td>MIRR</td>
<td>1</td>
<td>5</td>
<td>byte</td>
<td><a href="../Mirroring/">What CIRAM A10 is connected to</a>:
<table>
<tbody>
<tr>
<td style="text-align: right;">0.</td>
<td>PPU A11 (horizontal mirroring)</td>
</tr>
<tr>
<td style="text-align: right;">1.</td>
<td>PPU A10 (vertical mirroring)</td>
</tr>
<tr>
<td style="text-align: right;">2.</td>
<td>Ground (one-screen A)</td>
</tr>
<tr>
<td style="text-align: right;">3.</td>
<td>Vcc (one-screen B)</td>
</tr>
<tr>
<td style="text-align: right;">4.</td>
<td>Extra memory has been added (four-screen)</td>
</tr>
<tr>
<td style="text-align: right;">5.</td>
<td>Mapper-controlled</td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table>

## Shortcomings {#Shortcomings}

Prior to 2013, no encoding was specified for any of the fields; 7-bit-clean ASCII was assumed, making NAME inadequate for the vast majority of non-US games. In the first quarter of 2013, [UTF-8 became the encoding](http://forums.nesdev.org/viewtopic.php?f=3&t=9883).

Chunks can come in any order, so conventional patching tools cannot work without going through an "unpacked" intermediary stage.

MAPR chunks are nominally supposed to use the text on the PCB, such as "NES-SNROM". However, some games have no identifying text on the PCB at all. Other games have only symbols in [Japanese](http://bootgod.dyndns.org:7777/profile.php?id=1762) or Chinese. Sometimes the [same PCB](../UxROM/) can have [different incompatible behavior](../INES_Mapper_180/), depending on how things are populated or what things are jumpered. The workaround has been to add extra text in the MAPR chunk (in the *Crazy Climber* case, ["HVC-UNROM+74HC08"](http://bootgod.dyndns.org:7777/profile.php?id=3869)).

There is no ability to specify PRG RAM outside of the MAPR chunk. Two games using VRC4 ([Gradius II](http://bootgod.dyndns.org:7777/profile.php?id=1565) and [Bio Miracle Bokutte Upa](http://bootgod.dyndns.org:7777/profile.php?id=3988)) use the exact same PCB, but the former adds 2KiB PRG RAM and the latter adds none.

For greater emulator compatibility, people sometimes use already-known-supported MAPR chunks to get something that's "close enough", rather than specifying a new MAPR for not-necessarily-identical behavior.

BATR chunks do not disambiguate which RAM is battery-backed, if more than one is present.

It's not clear exactly what VROR is supposed to mean—"Do not throw an error if this MAPR normally has CHR ROM but there are no CHR*n* chunks, just give me 8KiB of CHR RAM"? "All the data I gave you for CHR-ROM, that was actually RAM, make it writeable"? As such, Nestopia, Nintendulator, and FCEUX just ignore it.

CTRL chunks do not specify which controller should be plugged into which port, nor Famicom-only controllers, nor Super NES controllers plugged into a Famiclone or through an adapter (such as the [12-key controller](../Standard_controller/) or the [mouse](../Mouse/)).

There is no way to fully describe PlayChoice 10 or [Vs. System](../Vs._System/) games.

## References {#References}

- [libunif project on Github](http://github.com/eteran/libunif)
- [Last published version of the standard](http://raw.githubusercontent.com/eteran/libunif/master/UNIF_current.txt)

## See also {#See_also}

- [UNIF to NES 2.0 Mapping](../UNIF_to_NES_2.0_Mapping/)
