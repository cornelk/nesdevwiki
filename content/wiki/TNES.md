---
title: "TNES"
url: "/wiki/TNES/"
description: "The TNES file format was discovered in Nintendo 3DS Virtual Console games. It fulfills much the same role as the iNES format, but it can also encapsulate QD images for FDS games and avoids…"
wikiPage: "TNES"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
---

The **TNES** file format was discovered in Nintendo 3DS Virtual Console games. It fulfills much the same role as the [iNES](../INES/) format, but it can also encapsulate [QD](../QD/) images for [FDS](../Family_Computer_Disk_System/) games and avoids [allegations of having misappropriated a format from the emulation scene](https://forums.nesdev.org/viewtopic.php?f=5&t=15403).

## TNES Header Format {#TNES_Header_Format}

| Offset | Contents                                                       |
|--------|----------------------------------------------------------------|
| 0-3    | "TNES" (\$54 \$4E \$45 \$53)                                   |
| 4      | Mapper (TNES codes, not iNES codes)                            |
| 5      | PRG ROM size in 8192 byte units                                |
| 6      | CHR ROM size in 8192 byte units                                |
| 7      | WRAM (0 for none, 1 for 8192 bytes)                            |
| 8      | Mirroring (0 mapper-controlled; 1 horizontal; 2 vertical)      |
| 9      | Battery (0 none; 1 battery present)                            |
| 10-11  | Used only for FDS images; purpose unknown, possibly side count |
| 12-15  | Zero                                                           |

### TNES Mapper Numbers {#TNES_Mapper_Numbers}

[Mapper](../Mapper/) numbers differ, and Nintendo's own ASIC mappers have a contiguous block:

<table class="wikitable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>TNES mapper</th>
<th>iNES mapper</th>
<th>Meaning</th>
<th>Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td><a href="../NROM/">0</a></td>
<td><a href="../NROM/">NROM</a></td>
<td></td>
</tr>
<tr>
<td>1</td>
<td><a href="../MMC1/">1</a></td>
<td>SxROM (<a href="../MMC1/">MMC1</a>)</td>
<td></td>
</tr>
<tr>
<td>2</td>
<td><a href="../MMC2/">9</a></td>
<td>PNROM (<a href="../MMC2/">MMC2</a>)</td>
<td></td>
</tr>
<tr>
<td>3</td>
<td><a href="../MMC3/">4</a></td>
<td>TxROM (<a href="../MMC3/">MMC3</a>)</td>
<td>And presumably HKROM (<a href="../MMC6/">MMC6</a>)</td>
</tr>
<tr>
<td>4</td>
<td><a href="../MMC4/">10</a></td>
<td>FxROM (<a href="../MMC4/">MMC4</a>)</td>
<td></td>
</tr>
<tr>
<td>5</td>
<td><a href="../MMC5/">5</a></td>
<td>ExROM (<a href="../MMC5/">MMC5</a>)</td>
<td></td>
</tr>
<tr>
<td>6</td>
<td><a href="../UxROM/">2</a></td>
<td><a href="../UxROM/">UxROM</a></td>
<td></td>
</tr>
<tr>
<td>7</td>
<td><a href="../CNROM/">3</a></td>
<td><a href="../CNROM/">CNROM</a></td>
<td></td>
</tr>
<tr>
<td>9</td>
<td><a href="../AxROM/">7</a></td>
<td><a href="../AxROM/">AxROM</a></td>
<td></td>
</tr>
<tr>
<td>31</td>
<td><a href="../INES_Mapper_086/">86</a></td>
<td>Jaleco JF-13</td>
<td></td>
</tr>
<tr>
<td>100</td>
<td></td>
<td><a href="../Family_Computer_Disk_System/">FDS</a></td>
<td>ROM image consists of a modified 8192 byte <a href="../FDS_BIOS/">BIOS</a><sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup> followed by <a href="../QD/">QD</a> sides.<br />
Size bytes 5-9 are zero, but 10-11 are used.</td>
</tr>
</tbody>
</table>

## References {#References}

- Based on [TNES format](https://pastebin.com/KLeWt2W3) by einstein95, via [TCRF forum post](http://jul.rustedlogic.net/thread.php?id=17316)
- [FDS .qd format](https://forums.nesdev.org/viewtopic.php?f=2&t=15792)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [TCRF: FDS BIOS modifications in 3DS VC](https://tcrf.net/Family_Computer_Disk_System#3DS_Virtual_Console)
