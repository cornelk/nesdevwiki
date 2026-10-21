---
title: "Comparison of Nintendo mappers"
url: "/wiki/Comparison_of_Nintendo_mappers/"
description: "This article compares the capabilities of several Nintendo mappers."
wikiPage: "Comparison_of_Nintendo_mappers"
---

This article compares the capabilities of several Nintendo mappers.

## Discrete logic {#Discrete_logic}

All of Nintendo's [discrete logic mappers](../../category/Discrete_logic_mappers/) have [bus conflicts](../Bus_conflict/) except some revisions of 7.

| iNES | Chips | Boards | Max PRG | PRG banks | Max CHR | CHR banks | Mirroring | PRG RAM? | IRQ |
|----|----|----|----|----|----|----|----|----|----|
| [0](../NROM/) | 0-1 | NROM | 32 |  | 8 |  | V/H hardwired | Japan only | None |
| [2](../UxROM/) | 2 | [UNROM, UOROM](../UxROM/) | 256 | 16 + 16F | 8 |  | V/H hardwired | No | None |
| [3](../CNROM/) | 1 | [CNROM](../CNROM/) | 32 |  | 32 | 8 | V/H hardwired | No | None |
| [7](../AxROM/) | 1-2 | [AxROM](../AxROM/) | 256 | 32 | 8 |  | 1 | No | None |
| [13](../CPROM/) | 3 | [CPROM](../CPROM/) (U) | 32 |  | 16 | 4 + 4F | V/H hardwired | No | None |
| [34](../INES_Mapper_034/) | 1 | [BNROM](../INES_Mapper_034/) | 256 | 32 | 8 |  | V/H hardwired | No | None |
| [66](../GxROM/) | 1 | [GNROM, MHROM](../GxROM/) | 128 | 32 | 32 | 8 | V/H hardwired | No | None |
| [94](../INES_Mapper_094/) | 2 | [UN1ROM](../UxROM/) | 256 | 16 + 16F | 8 |  | V/H hardwired | No | None |
| [180](../INES_Mapper_180/) | 2 | UNROM+74'08 | 256 | 16F + 16 | 8 |  | V/H hardwired | No | None |

## ASIC {#ASIC}

None of these [ASIC mappers](../../category/ASIC_mappers/) has bus conflicts.

<table class="tabular" style="width:100%;">
<colgroup>
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
<col style="width: 10%" />
</colgroup>
<thead>
<tr>
<th>iNES</th>
<th>Mapper</th>
<th>Boards</th>
<th>Max PRG ROM</th>
<th>PRG ROM banks</th>
<th>Max CHR</th>
<th>CHR banks</th>
<th>Mirroring</th>
<th>PRG RAM?</th>
<th>Assist</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="../MMC1/">1</a></td>
<td><a href="../MMC1/">MMC1</a></td>
<td><a href="../MMC1/">SGROM, SNROM, SUROM</a></td>
<td>512</td>
<td>16 + 16F; 32</td>
<td>8</td>
<td>4 + 4</td>
<td>V/H/1 switchable</td>
<td>Optional</td>
<td>None</td>
</tr>
<tr>
<td>1</td>
<td>MMC1</td>
<td>SKROM, SLROM</td>
<td>256</td>
<td>16 + 16F; 32</td>
<td>128</td>
<td>4 + 4</td>
<td>V/H/1 switchable</td>
<td>Optional</td>
<td>None</td>
</tr>
<tr>
<td><a href="../MMC3/">4</a></td>
<td><a href="../MMC3/">MMC3</a></td>
<td><a href="../TxROM/">TxROM</a></td>
<td>512</td>
<td>8 + 8 + 16F</td>
<td>256</td>
<td>2 + 2 + 1 + 1 + 1 + 1</td>
<td>V/H switchable</td>
<td>Optional</td>
<td>Scanline IRQ</td>
</tr>
<tr>
<td>4</td>
<td><a href="../MMC3/">MMC3</a></td>
<td>TGROM, TNROM</td>
<td>512</td>
<td>8 + 8 + 16F</td>
<td>8</td>
<td>2 + 2 + 1 + 1 + 1 + 1</td>
<td>V/H switchable</td>
<td>Japan only</td>
<td>Scanline IRQ</td>
</tr>
<tr>
<td><a href="../INES_Mapper_118/">118</a></td>
<td>MMC3</td>
<td><a href="../TKSROM/">TKSROM</a>, <a href="../TLSROM/">TLSROM</a></td>
<td>512</td>
<td>8 + 8 + 16F</td>
<td>128</td>
<td>2 + 2 + 1 + 1 + 1 + 1</td>
<td>Any switchable</td>
<td>Optional</td>
<td>Scanline IRQ</td>
</tr>
<tr>
<td><a href="../INES_Mapper_119/">119</a></td>
<td>MMC3</td>
<td>TQROM</td>
<td>128</td>
<td>8 + 8 + 16F</td>
<td>64 + 8</td>
<td>2 + 2 + 1 + 1 + 1 + 1</td>
<td>V/H switchable</td>
<td>No</td>
<td>Scanline IRQ</td>
</tr>
<tr>
<td>4</td>
<td><a href="../MMC6/">MMC6</a></td>
<td>HKROM</td>
<td>512</td>
<td>8 + 8 + 16F</td>
<td>256</td>
<td>2 + 2 + 1 + 1 + 1 + 1</td>
<td>V/H switchable</td>
<td>1KiB built-in; US and Europe only</td>
<td>Scanline IRQ</td>
</tr>
<tr>
<td><a href="../MMC5/">5</a></td>
<td><a href="../MMC5/">MMC5</a></td>
<td><a href="../ExROM/">ExROM</a></td>
<td>1024</td>
<td>32, 16 + 16, 16 + 8 + 8, 8 + 8 + 8 + 8</td>
<td>1024</td>
<td>BG: 4, 2 + 2, or 1 + 1 + 1 + 1<br />
Sprites: 8, 4 + 4, 2 * 4, or 1 * 8<br />
Split window: 4</td>
<td>Any</td>
<td>1 KiB built in; more optional</td>
<td>Scanline IRQ<br />
Smaller attribute areas<br />
Separate CHR bank per background tile<br />
Separate sets of CHR banks for sprites<br />
Vertical split screen<br />
8x8 multiplier</td>
</tr>
<tr>
<td><a href="../MMC2/">9</a></td>
<td><a href="../MMC2/">MMC2</a></td>
<td>PNROM</td>
<td>128</td>
<td>8 + 24F</td>
<td>128</td>
<td>4/4 + 4/4</td>
<td>V/H switchable</td>
<td>No</td>
<td>Tile triggered CHR bank switching</td>
</tr>
<tr>
<td><a href="../MMC4/">10</a></td>
<td><a href="../MMC4/">MMC4</a></td>
<td><a href="../FxROM/">FJROM, FKROM</a> (J)</td>
<td>256</td>
<td>16 + 16F</td>
<td>128</td>
<td>4/4 + 4/4</td>
<td>V/H switchable</td>
<td>Yes</td>
<td>Tile triggered CHR bank switching</td>
</tr>
</tbody>
</table>

## External links {#External_links}

- [Mapper wizard](http://pineight.com/nes/mapperwizard.html): Web-based decision tree for choosing a mapper for your NES project
