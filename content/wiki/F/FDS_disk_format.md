---
title: "FDS disk format"
url: "/wiki/FDS_disk_format/"
description: "The Famicom Disk System uses disks that are a modified version of Mitsumi Quick Disk hardware, with a custom data format stored on the disk."
wikiPage: "FDS_disk_format"
---

The [Famicom Disk System](../Family_Computer_Disk_System/) uses disks that are a modified version of Mitsumi Quick Disk hardware, with a custom data format stored on the disk.

Each side can hold more than 64KB of data.

For archive and emulation, dumps of these disks are often preserved using the [FDS file format](../FDS_file_format/) (.FDS).

## FDS Disk Side format {#FDS_Disk_Side_format}

Each disk side must be structured into block as follows :

1, 2, 3, 4, 3, 4, ...., 3, 4

The 3, 4 pattern should be repeated once per file present on the disk.

Blocks type 1 and 2 represent information about the disk and how many files it has, and each block type 3 and 4 pair represents a single file. See *[Block format](#Block_format)* below.

### Gaps {#Gaps}

Physically on the disk, there are "gaps" of 0 recorded between blocks and before the start of the disk. Length of the gaps are as follows:

- Before the start of the disk : At least 26150 bits, 28300 typical.
- Gap between blocks : At least 480 bits, 976 bits typical.

Gaps are teminated by a single '1' bit. In terms of bytes, it would be \$80, as the data is stored in little endian format.

### CRCs {#CRCs}

At the end of each block, a 16-bit CRC is stored. This CRC is *not* calculated by the CPU executing the BIOS code, but by the [RP2C33](../RP2C33_pinout/) handling the disk drive signals. It will automatically set \$4030.D4 = 1 (leading to error code 27) if a block's CRC value does not match the calculated one during loading.

The CRC algorithm used is the common [CRC-16/KERMIT](https://reveng.sourceforge.io/crc-catalogue/16.htm#crc.cat.crc-16-kermit) algorithm. The '1' bit at the end of the gap is included in the calculation.

### True disk capacity {#True_disk_capacity}

The commonly used [FDS file format](../FDS_file_format/) is limited to 65500 bytes per side and does not contain CRCs or gaps. The [QD file format](../QD/) observed in rereleases on Nintendo platforms is limited to 65536 bytes per side and contains CRCs, but no gaps.

The actual disk capacity is somewhat larger, and there are a few variant disk formats that may have even more capacity.

## Block format {#Block_format}

### <span id="Disk_info_block_.28block_1.29"></span>Disk info block (block 1) {#Disk_info_block_.28block_1.29}

<table class="wikitable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Offset</th>
<th>Length (bytes)</th>
<th>Description</th>
<th>Details</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>1</td>
<td>Block code</td>
<td>Raw byte: $01</td>
</tr>
<tr>
<td>$01</td>
<td>14</td>
<td>Disk verification</td>
<td>Literal ASCII string: <code>*NINTENDO-HVC*</code><br />
Used by BIOS to verify legitimate disk image</td>
</tr>
<tr>
<td>$0F</td>
<td>1</td>
<td>Licensee code</td>
<td>See <a href="../Licensee_codes/">Licensee codes</a></td>
</tr>
<tr>
<td>$10</td>
<td>3</td>
<td>Game name</td>
<td>3-letter ASCII code per game (e.g. ZEL for <em>The Legend of Zelda</em>)</td>
</tr>
<tr>
<td>$13</td>
<td>1</td>
<td>Game type</td>
<td>$20 = <code>" "</code> — Normal disk<br />
$45 = <code>"E"</code> — Event (e.g. Japanese national <a href="http://ja.wikipedia.org/wiki/%E3%83%95%E3%82%A1%E3%83%9F%E3%83%AA%E3%83%BC%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%82%BF_%E3%83%87%E3%82%A3%E3%82%B9%E3%82%AF%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0#.E3.83.87.E3.82.A3.E3.82.B9.E3.82.AF.E3.83.95.E3.82.A1.E3.82.AF.E3.82.B9">DiskFax</a> tournaments)<br />
$4A = <code>"J"</code> — Unknown. May indicate <a href="../Family_Computer_Network_Adapter/">Family Computer Network Adapter</a><br />
$52 = <code>"R"</code> — Reduction in price via advertising</td>
</tr>
<tr>
<td>$14</td>
<td>1</td>
<td>Game version</td>
<td>Starts at $00, increments per revision</td>
</tr>
<tr>
<td>$15</td>
<td>1</td>
<td>Side number</td>
<td>$00 = Side A, $01 = Side B. Single-sided disks use $00</td>
</tr>
<tr>
<td>$16</td>
<td>1</td>
<td>Disk number</td>
<td>First disk is $00, second is $01, etc.</td>
</tr>
<tr>
<td>$17</td>
<td>1</td>
<td>Disk type (FMC)</td>
<td>$01 for FMC blue-disk releases, and $00 otherwise <sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup></td>
</tr>
<tr>
<td>$18</td>
<td>1</td>
<td>Unknown</td>
<td>$00 in all known games</td>
</tr>
<tr>
<td>$19</td>
<td>1</td>
<td>Boot read file code</td>
<td>Refers to the file code/file number to load upon boot/start-up</td>
</tr>
<tr>
<td>$1A</td>
<td>5</td>
<td>Unknown</td>
<td>Raw bytes: $FF $FF $FF $FF $FF</td>
</tr>
<tr>
<td>$1F</td>
<td>3</td>
<td>Manufacturing date</td>
<td>See <a href="#Date_format">Date format</a></td>
</tr>
<tr>
<td>$22</td>
<td>1</td>
<td>Country code</td>
<td>$49 = Japan</td>
</tr>
<tr>
<td>$23</td>
<td>1</td>
<td>Unknown</td>
<td>Raw byte: $61. Speculative: Region code?</td>
</tr>
<tr>
<td>$24</td>
<td>1</td>
<td>Unknown</td>
<td>Raw byte: $00. Speculative: Location/site?</td>
</tr>
<tr>
<td>$25</td>
<td>2</td>
<td>Unknown</td>
<td>Raw bytes: $00 $02</td>
</tr>
<tr>
<td>$27</td>
<td>5</td>
<td>Unknown</td>
<td>Speculative: some kind of game information representation?</td>
</tr>
<tr>
<td>$2C</td>
<td>3</td>
<td>"Rewritten disk" date</td>
<td>See <a href="#Date_format">Date format</a>. It's speculated this refers to the date the disk was formatted and rewritten by something like a <a href="http://www.nintendolife.com/news/2014/03/weirdness_this_footage_of_the_famicom_disk_writer_kiosk_is_a_bit_awesome">Disk Writer kiosk</a>.<br />
In the case of an original (non-copied) disk, this should be the same as <a href="#Manufacturing_date">Manufacturing date</a></td>
</tr>
<tr>
<td>$2F</td>
<td>1</td>
<td>Unknown</td>
<td></td>
</tr>
<tr>
<td>$30</td>
<td>1</td>
<td>Unknown</td>
<td>Raw byte: $80</td>
</tr>
<tr>
<td>$31</td>
<td>2</td>
<td>Disk Writer serial number</td>
<td></td>
</tr>
<tr>
<td>$33</td>
<td>1</td>
<td>Unknown</td>
<td>Raw byte: $07</td>
</tr>
<tr>
<td>$34</td>
<td>1</td>
<td>Disk rewrite count</td>
<td>Value stored in <a href="https://en.wikipedia.org/wiki/Binary-coded_decimal">BCD</a> format. $00 = Original (no copies).</td>
</tr>
<tr>
<td>$35</td>
<td>1</td>
<td>Actual disk side</td>
<td>$00 = Side A, $01 = Side B</td>
</tr>
<tr>
<td>$36</td>
<td>1</td>
<td>Disk type (other)</td>
<td>$00 = yellow disk, $FF = blue disk, $FE = prototype, sample, or internal-use (white or pink) disk. Some prototype disks have been observed with value $00, but this may indicate the field was simply not filled in.</td>
</tr>
<tr>
<td>$37</td>
<td>1</td>
<td>Disk version</td>
<td>Unknown how this differs from game version ($14). Disk version numbers indicate different software revisions. Speculation is that disk version incremented with each disk received from a licensee.</td>
</tr>
<tr>
<td>$38</td>
<td>2</td>
<td>CRC</td>
<td>(Omitted from .FDS files)</td>
</tr>
</tbody>
</table>

The `*NINTENDO-HVC*` string, stored in ASCII, proves that this is a FDS disk. If the string doesn't match, the BIOS will refuse to read the disk further.

If the FDS is started with a disk whose side number and disk number aren't both \$00, it will be prompted to insert the first disk side. However, some games make this number \$00, even for the second disk to make it bootable too.

All files with IDs smaller or equals to the *boot read file code* will be loaded when the game is booting.

#### Date format {#Date_format}

All bytes are stored in BCD format, in order of year, month, and day. The year is usually represented as the 1-indexed number of years since the start of the current [Japanese calendar](https://en.wikipedia.org/wiki/Japanese_calendar) era, either the Shōwa era (1926) or Heisei era (1989). Values observed are \$61-63 for 1986-1988 and \$01-03 for 1989-1991. Disk Writer kiosks continued to use Shōwa era dates beyond the end of the Shōwa era, so disks written by the service up to its discontinuation in 2003 may have year values as high as \$78.

A small number of games used the last two digits of the current year instead of a Japanese calendar date. Values observed are \$85-\$86 for 1985-1986.

### <span id="File_amount_block_.28block_2.29"></span>File amount block (block 2) {#File_amount_block_.28block_2.29}

This block contains the total number of files recorded on disk.

| Offset | Length (bytes) | Description | Details                   |
|--------|----------------|-------------|---------------------------|
| \$00   | 1              | Block code  | Raw byte: \$02            |
| \$01   | 1              | File Amount |                           |
| \$02   | 2              | CRC         | (Omitted from .FDS files) |

More files might exist on the disk, but the BIOS load routine will ignore them, those files are called "hidden" files. Some games have a simple copy protection this way: They have their own loading routine similar to the one from the BIOS but hard-code the file amount to a higher number, which will allow for loading hidden files. This also allows the game to load faster because the BIOS will stop reading the disk after the last non-hidden file.

### <span id="File_header_block_.28block_3.29"></span>File header block (block 3) {#File_header_block_.28block_3.29}

This block stores metadata for the file data block which follows it. Because the BIOS has to manually seek through the disk one byte at a time, it relies on the header block fields to correctly load or skip file data blocks.

<table class="wikitable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Offset</th>
<th>Length (bytes)</th>
<th>Description</th>
<th>Details</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>1</td>
<td>Block code</td>
<td>Raw byte: $03</td>
</tr>
<tr>
<td>$01</td>
<td>1</td>
<td>File Number</td>
<td></td>
</tr>
<tr>
<td>$02</td>
<td>1</td>
<td>File ID</td>
<td>The ID specified at disk-read function call.</td>
</tr>
<tr>
<td>$03</td>
<td>8</td>
<td>File Name</td>
<td>Typically in uppercase ASCII.</td>
</tr>
<tr>
<td>$0B</td>
<td>2</td>
<td>File Address</td>
<td>16-bit little-endian. The destination address when loading.</td>
</tr>
<tr>
<td>$0D</td>
<td>2</td>
<td>File Size</td>
<td>16-bit little-endian.</td>
</tr>
<tr>
<td>$0F</td>
<td>1</td>
<td>File Type</td>
<td>0: Program (PRG-RAM)<br />
1: Character (CHR-RAM)<br />
2: Nametable (VRAM)</td>
</tr>
<tr>
<td>$10</td>
<td>2</td>
<td>CRC</td>
<td>(Omitted from .FDS files)</td>
</tr>
</tbody>
</table>

File Numbers are generally in increasing order, first file is 0. The BIOS does not use file Numbers during reads, and file Numbers written when calling AppendFile/WriteFile are manually counted from the first file.

File IDs can be freely assigned, and this is the number which will decide which file is loaded from the disk (instead of the file number). For example, the parameter disks for *Quick Hunter* uses identical file Numbers for everything except the first file and only distinguishes them by file IDs. A file ID smaller or equal to than the *boot read file code* means the file is a boot file, and will be loaded on first start up.

The BIOS only writes to the PPU when the file type is not equal to 0, so types 2 and above are functionally identical to type 1.

### <span id="File_data_block_.28block_4.29"></span>File data block (block 4) {#File_data_block_.28block_4.29}

This block stores the actual file data.

| Offset | Length (bytes) | Description | Details                   |
|--------|----------------|-------------|---------------------------|
| \$00   | 1              | Block code  | Raw byte: \$04            |
| \$01   | --             | File data   |                           |
| --     | 2              | CRC         | (Omitted from .FDS files) |

### <span id="Test_data_block_.28block_5.29"></span>Test data block (block 5) {#Test_data_block_.28block_5.29}

| Offset | Length (bytes) | Description | Details |
|----|----|----|----|
| \$00 | 1 | Block code | Raw byte: \$05 |
| \$01 | -- | Test data | Repeating sequence of \$6D, \$B6, \$DB to fill the disk side. |
| -- | 2 | CRC | (Omitted from .FDS files) |

This block type and fragments of its data has been observed in retail disks. It appears to be test data used during the disk formatting process. *FMC Disk Card Checker* is capable of writing and verifying this block.

## Non-standard disk formats {#Non-standard_disk_formats}

Because the disk format is only enforced in software by the BIOS disk I/O routines, custom formats are possible through the use of low-level disk I/O routines.

Some unlicensed disks have been observed to use non-standard formats, often as a means of copy-protection:

- *Kosodate Gokko* stores its main program without any block indicator and loads it using the BIOS' disk IRQ handler. (data is prefixed with a \$12 byte)
- *Quick Hunter* uses a block type of 0 to store its encrypted main program. (data is XORed with \$C9) Note that common dumps of this disk lack the needed 0 byte.
- *Graphic Editor Hokusai* and *Jingorou* use a custom format (interleaved + encrypted PRG/CHR data) with fake file header/data blocks for their main programs.
- Game Doctor/Magic Card devices switch in a replacement BIOS to load [custom disk formats](../Game_Doctor/Magic_Card_FDS_Format/).
  - Block type 5 is used by some Magic Card 1M/2M/4M disks to store trainer data.

## See also {#See_also}

- [FDS file format](../FDS_file_format/) (**.FDS**)
- [QD file format](../QD/) (**.QD**)
- [FDS technical reference.txt](https://nesdev.org/FDS%20technical%20reference.txt) by Brad Taylor
- [Enri's Famicom Disk System page](http://cmpslv3.stars.ne.jp/Famic/Famdis.htm) (Japanese)
- [Enri's Famicom Disk System page](https://web.archive.org/web/20091023182159/http://www2.odn.ne.jp/~haf09260/Famic/Famdis.htm) (Japanese) (old/outdated)
- [Fantasy's FDS Maniacs page](http://park19.wakwak.com/~fantasy/fds/) (Japanese). Technical information is in the [ディスクシステム資料室](http://park19.wakwak.com/~fantasy/fds/document/document.htm) section.
- [FDS Study site](https://web.archive.org/web/20080515232015/http://fdsstudy.hp.infoseek.co.jp/) via archive.org (Japanese). The FDSStudy program can still be found [here](https://w.atwiki.jp/famicomall/pages/589.html).
- [fds-nori.txt](https://nesdev.org/fds-nori.txt) - FDS reference in Japanese by Nori
- [Forum post](https://forums.nesdev.org/viewtopic.php?p=194867#p194867): .fds format: Can checksums be heuristically detected? - Includes a CRC implementation in C.
- [Forum thread](https://forums.nesdev.org/viewtopic.php?t=26370): Block type 5 in non-Game-Doctor FDS disks
- [Forum thread](https://forums.nesdev.org/viewtopic.php?t=18777): Copy protection methods on FDS disks

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [FamicomWorld:](http://famicomworld.com/workshop/articles/product-codes-fmc-and-fsc/) FMC and FSC product codes
