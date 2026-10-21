---
title: "MMC5"
url: "/wiki/MMC5/"
description: "The Nintendo MMC5 is a mapper ASIC used in Nintendo's ExROM Game Pak boards. All MMC5 boards are assigned to mapper 5."
wikiPage: "MMC5"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers triggering on reads"
    url: "/category/Mappers_triggering_on_reads/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with large PRG RAM"
    url: "/category/Mappers_with_large_PRG_RAM/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>MMC5</b><br/><b>ExROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, Koei, others </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=5">15 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>EKROM, ELROM,<br/>ETROM, EWROM </td></tr><tr><td><b>Pinout</b></td><td><a href="../MMC5_pinout/">MMC5 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>1024K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K, 16K, or 32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>128K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K ($6000-$DFFF),<br/>16K (only $8000-$BFFF<br/>at PRG mode 1/2) </td></tr><tr><td><b>CHR capacity</b></td><td>1024K </td></tr><tr><td><b>CHR window</b></td><td>1K, 2K, 4K, or 8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>arbitrary, up to 3 source<br/>nametables (plus fill mode) </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td><a href="../MMC5_audio/">Yes</a></td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">005</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=5">iNES 005</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-E%25ROM">ExROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=EKROM">EKROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=ELROM">ELROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=ETROM">ETROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=EWROM">EWROM</a></td></tr></tbody></table></div>

The **Nintendo MMC5** is a [mapper](../Mapper/) [ASIC](../../category/ASIC_mappers/) used in Nintendo's [ExROM](../ExROM/) Game Pak boards. All MMC5 boards are assigned to **mapper 5**.

Example games:

- *Castlevania 3*
- *Just Breed*
- *Uncharted Waters*
- *Romance of the Three Kingdoms II*
- *Laser Invasion*
- *Metal Slader Glory*
- *Uchuu Keibitai SDF*
- *Shin 4 Nin Uchi Mahjong - Yakuman Tengoku*
- *Bandit Kings of Ancient China*
- *SimCity* (Prototype)

The first game to use this chip (*Nobunaga's Ambition II*) was released in February 1990. The [date codes on components on early released cartridges](http://bootgod.dyndns.org:7777/profile.php?id=3169) show that manufacturing had started at the end of 1989.

A later **[MMC5A](#MMC5A)** revision was introduced with a few extra features, but all released games do not rely on these features and are compatible with the original MMC5.

## Overview {#Overview}

The MMC5 is the most elaborate mapper ASIC Nintendo made for the NES and Famicom.

It supports a kitchen sink of features, including:

- 4 PRG ROM switching modes
- 4 CHR ROM switching modes
- Up to 128KB of WRAM, mappable not only at \$6000-\$7FFF but also within \$8000-\$DFFF
  - Supports either one chip (up to 128KB) or two chips (up to 32KB each)
- An 8 bit by 8 bit multiplier with a 16 bit result for performing quick calculations
- Scanline detection with counter and configurable IRQ
- Frame detection with status bit
- The ability to use different CHR banks for background and 8x16 sprites (allowing 256 unique 8x16 sprite tiles, independent of the background).
- 1024 bytes of on-chip memory, which can be used for 4 different purposes:
  - An extra general-use nametable
  - Attribute and tile index expansion - address 16384 background tiles at once, and allow each individual 8x8 tile to have its own palette setting
  - Vertical split-screen
  - Extra RAM for storing program variables
- Three extra sound channels
  - Two pulse channels, identical to those in the NES APU (except lacking pitch sweeps).
  - An 8-bit RAW PCM channel
- A 'fill mode' nametable, which can be instantly set to contain a specific tile in a specific color (useful for screen transitions)
- System reset detection
  - Triggered by a positive or negative gap in M2 of at least 11.2 usec.
  - Also triggered and latched by absence of AVcc.
  - After reapplying AVcc, another gap in M2 is sometimes necessary to clear the latch.
  - This feature resets some, but not all, states of the MMC5.
  - The PRG RAM +CE pin is a direct reflection of system reset detection state.
  - Reduces the need to hold the Reset Button while turning off the power in games with battery save.<sup><span id="cite_ref-WYGPNF_1-0"></span>[\[1\]](#cite_note-WYGPNF-1)</sup>

## Banks {#Banks}

The MMC5 provides 4 distinct banking modes for both PRG ROM and CHR ROM.

### PRG mode 0 {#PRG_mode_0}

- CPU \$6000-\$7FFF: 8 KB switchable PRG RAM bank
- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank

### PRG mode 1 {#PRG_mode_1}

- CPU \$6000-\$7FFF: 8 KB switchable PRG RAM bank
- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM/RAM bank
- CPU \$C000-\$FFFF: 16 KB switchable PRG ROM bank

### PRG mode 2 {#PRG_mode_2}

- CPU \$6000-\$7FFF: 8 KB switchable PRG RAM bank
- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM/RAM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM/RAM bank
- CPU \$E000-\$FFFF: 8 KB switchable PRG ROM bank

### PRG mode 3 {#PRG_mode_3}

- CPU \$6000-\$7FFF: 8 KB switchable PRG RAM bank
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM/RAM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM/RAM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM/RAM bank
- CPU \$E000-\$FFFF: 8 KB switchable PRG ROM bank

### CHR mode 0 {#CHR_mode_0}

- PPU \$0000-\$1FFF: 8 KB switchable CHR bank

### CHR mode 1 {#CHR_mode_1}

- PPU \$0000-\$0FFF: 4 KB switchable CHR bank
- PPU \$1000-\$1FFF: 4 KB switchable CHR bank

### CHR mode 2 {#CHR_mode_2}

- PPU \$0000-\$07FF: 2 KB switchable CHR bank
- PPU \$0800-\$0FFF: 2 KB switchable CHR bank
- PPU \$1000-\$17FF: 2 KB switchable CHR bank
- PPU \$1800-\$1FFF: 2 KB switchable CHR bank

### CHR mode 3 {#CHR_mode_3}

- PPU \$0000-\$03FF: 1 KB switchable CHR bank
- PPU \$0400-\$07FF: 1 KB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KB switchable CHR bank
- PPU \$1000-\$13FF: 1 KB switchable CHR bank
- PPU \$1400-\$17FF: 1 KB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KB switchable CHR bank

## Registers {#Registers}

### <span id="Sound_.28.245000-.245015.29"></span>Sound (\$5000-\$5015) {#Sound_.28.245000-.245015.29}

For details on sound operation, see [MMC5 audio](../MMC5_audio/)

### NES internal state monitoring {#NES_internal_state_monitoring}

All of these registers overlay various registers that are already used inside the NES, and are fully decoded. A game could write to a mirror of a PPU register to get the MMC5 out of sync, for example, writing to mirrors of \$2000 to make 8x16 sprites use the same banks with the nametable like other mappers (which is useful for [mapper hacks](../Glossary/)).

#### <span id="8x16_mode_enable_.28.242000_.3D_PPUCTRL.29"></span>8x16 mode enable (\$2000 = [PPUCTRL](../PPU_registers/)) {#8x16_mode_enable_.28.242000_.3D_PPUCTRL.29}

- Write only
- Reset detection not used

<!-- -->

    7  bit  0
    ---- ----
    xxZx xxxx
      |
      +------- Sprite size (0: 8x8 pixels; 1: 8x16 pixels)

Only when Z is set and at least one E bit is set does the MMC5 [draw 8x16 sprites from eight independent banks](#CHR_Bankswitching_.28.245120-.245130.29).<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

#### <span id="PPU_Data_Substitution_Enable_.28.242001_.3D_PPUMASK.29"></span>PPU Data Substitution Enable (\$2001 = [PPUMASK](../PPU_registers/)) {#PPU_Data_Substitution_Enable_.28.242001_.3D_PPUMASK.29}

- Write only
- Reset detection not used

<!-- -->

    7  bit  0
    ---------
    xxxE Exxx
       | |
       +-+---- 1,2,3: Substitutions enabled; 0: substitutions disabled

The MMC5 listens to writes to PPUMASK (\$2001). When it sees that both E bits are cleared, it disables its ability to make substitutions on the PPU data bus. This includes disabling:

- Independent bank 8x16 sprite mode
- Extended attribute mode
- Vertical split mode

The MMC5 only listens to the fully decoded address \$2001, so this can be tested by using the PPU’s mirrors of this register. For example, writing to register \$2009 will be heard by the PPU but not the MMC5. It is not clear if that could have a practical use beyond test purposes.

When the MMC5 sees \$00 written to \$2001, and then the PPU’s rendering gets enabled via a mirror of \$2001, the MMC5 still counts scanlines and can generate scanline interrupts even though it thinks \$2001 is still disabled. The transition from disabled to enabled resets the scanline counter.

Driving pin 92 low similarly disables substitutions. Because no scanline interrupts occur with pin 92 held low, it seems to hold the scanline counter in reset.

#### <span id="Scanline_Counter_Reset_.28.242002_.3D_PPUSTATUS.2C_read_only.29"></span>Scanline Counter Reset (\$2002 = [PPUSTATUS](../PPU_registers/), read only) {#Scanline_Counter_Reset_.28.242002_.3D_PPUSTATUS.2C_read_only.29}

- MMC5 stores the value read from the PPU
- Reset detection not used

<!-- -->

    7  bit  0
    ---------
    Exxx xxxx
    |
    +--------- 1: trigger scanline counter reset; 0: no effect

#### <span id="Unknown_.28.242005_.3D_PPUSCROLL.29"></span>Unknown (\$2005 = [PPUSCROLL](../PPU_registers/)) {#Unknown_.28.242005_.3D_PPUSCROLL.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---------
    xxxx x000  Reset Value
    ---- ----
    xxxx xYYY
          |||
          +++- CHR A2,A1,A0

The letterless MMC5 captures the low 3 bits of the second write to PPUSCROLL into a 3-bit counter. The value of the counter is output onto CHR A2,A1,A0 at certain times. It is not known how or when the counter does that. Tests using a mirror of PPUSCROLL on MMC5A were expected to cause rolled tiles, but that result was observed not to be true. It should be retested more carefully to confirm.

Second write tracking is resynced each time register \$2002 is read, in the same way that the PPU clears its ‘w’ address latch.

#### <span id="Unknown_.28.242006_.3D_PPUADDR.2C_MMC5A_only.29"></span>Unknown (\$2006 = [PPUADDR](../PPU_registers/), MMC5A only) {#Unknown_.28.242006_.3D_PPUADDR.2C_MMC5A_only.29}

Power analysis has detected that the MMC5A monitors writes here, and it has been confirmed by analyzing the MMC5A die. Though none of the values written to \$2006 are stored or used, it implements its own 2nd write tracking. The 2nd write tracking resyncs each time \$2002 is read, similar to the MMC5 and 5A’s \$2005 write tracking. When the 2nd write has been reached, it appears to wait until the next h-blank, then reload the 3-bit counter from the stored \$2005 fine vertical scroll value.

The purpose of this is not yet known.

#### <span id="Scanline_Counter_Reset_.28.244014_.3D_OAMDMA.29"></span>Scanline Counter Reset (\$4014 = [OAMDMA](../PPU_registers/)) {#Scanline_Counter_Reset_.28.244014_.3D_OAMDMA.29}

Writes to OAMDMA trigger a reset of the scanline counter. The value written is not captured by the MMC5.

### Configuration {#Configuration}

#### <span id="PRG_mode_.28.245100.29"></span>PRG mode (\$5100) {#PRG_mode_.28.245100.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xx11  Reset Value
    ---- ----
    xxxx xxPP
           ||
           ++- Select PRG banking mode

- 0 - One 32KB bank
- 1 - Two 16KB banks
- 2 - One 16KB bank (\$8000-\$BFFF) and two 8KB banks (\$C000-\$DFFF and \$E000-\$FFFF)
- 3 - Four 8KB banks

*Castlevania III* uses mode 2, which is similar to [VRC6](../VRC6/) PRG banking. All other games use mode 3. The Koei games never write to this register, apparently relying on the MMC5 defaulting to mode 3 at power on.

#### <span id="CHR_mode_.28.245101.29"></span>CHR mode (\$5101) {#CHR_mode_.28.245101.29}

- Write only
- Reset detection not used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xxCC
           ||
           ++- Select CHR banking mode

- 0 - 8KB CHR pages
- 1 - 4KB CHR pages
- 2 - 2KB CHR pages
- 3 - 1KB CHR pages

CHR mode always applies to sprite pattern fetches. However, CHR mode only applies to background pattern fetches during normal rendering. When extended attributes or vertical split are being rendered, the CHR bank size for background pattern fetches is fixed to 4kB.

CHR mode affects the mapping of the bits when writing to the 10-bit CHR bank registers \$5120-512B.

*Metal Slader Glory* uses 4KB CHR pages. All other games use 1KB pages.

#### <span id="PRG_RAM_Protect_1_.28.245102.29"></span>PRG RAM Protect 1 (\$5102) {#PRG_RAM_Protect_1_.28.245102.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xx01  Reset Value
    ---- ----
    xxxx xxWW
           ||
           ++- RAM protect 1

In order to enable writing to PRG RAM, this must be set to binary '10' (e.g. \$02).

#### <span id="PRG_RAM_Protect_2_.28.245103.29"></span>PRG RAM Protect 2 (\$5103) {#PRG_RAM_Protect_2_.28.245103.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xx10  Reset Value
    ---- ----
    xxxx xxWW
           ||
           ++- RAM protect 2

In order to enable writing to PRG RAM, this must be set to binary '01' (e.g. \$01).

#### <span id="Internal_extended_RAM_mode_.28.245104.29"></span>Internal extended RAM mode (\$5104) {#Internal_extended_RAM_mode_.28.245104.29}

- Write only
- Reset detection is used
- MMC5 additionally keeps track whether this register has ever been written since reset, purpose unknown.

<!-- -->

    7  bit  0
    ---- ----
    xxxx xx11  Reset Value
    ---- ----
    xxxx xxXX
           ||
           ++- Specify extended RAM usage

<table class="wikitable" style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th>$5104</th>
<th style="text-align: center;">CPU Access ($5C00-$5FFF)<br />
during blanking</th>
<th style="text-align: center;">CPU Access ($5C00-$5FFF)<br />
during rendering</th>
<th style="text-align: center;">PPUDATA Access ($2000-$2FFF)<sup>(1)</sup><br />
during blanking</th>
<th style="text-align: center;">Available as<br />
Nametable</th>
<th style="text-align: center;">Enable Extended<br />
Attribute Mode</th>
</tr>
</thead>
<tbody>
<tr>
<th>%00</th>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No<sup>(3)</sup></td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Write Only<sup>(4)</sup></td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Read/Write</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
</tr>
<tr>
<th>%01</th>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No<sup>(3)</sup></td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Write Only<sup>(4)</sup></td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Read/Write</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes<sup>2</sup></td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
</tr>
<tr>
<th>%10</th>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Read/Write</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Read/Write</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
</tr>
<tr>
<th>%11</th>
<td style="text-align: center;">Read Only</td>
<td style="text-align: center;">Read Only</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
</tr>
</tbody>
</table>

<sup>(1)</sup>When configured as a nametable in register \$5105, read and write access is possible through the PPU via registers \$2006/\$2007 when the PPU is not rendering. Based on register \$5105, any nametables mapped to extended RAM will support the corresponding PPU address ranges for reads and writes. For example, if NT4 is mapped to extended RAM, reads/writes to PPU \$2C00 will correspond with extended RAM \$5C00.

<sup>(2)</sup>Though it is possible to still assign the extended RAM as a nametable in the mode %01, you are going to have the same data used twice, once as extended attribute data, and once as nametable data, this does not seem to be a useful combination. In the other modes, the nametable will read as if it contains all \$00s.

<sup>(3)</sup>Counterintuitively, writes in these modes are only allowed when the PPU is rendering. If writes are attempted during V-blank, they may be ignored or cause a corruption at that memory address. In practice, temporarily switch to mode %10 if you wish to write during V-blank.

<sup>(4)</sup>Attempting to read extended RAM in those modes returns open bus.

##### Extended attributes {#Extended_attributes}

In mode %01, "Extended Attribute Mode", each byte of the MMC5's internal extended RAM is used to enhance the background tile at the corresponding nametable address. The extended attributes are 1-screen mirrored; in other words, they apply the same for all nametables.

Format of each extended attribute byte:

    7  bit  0
    ---- ----
    AACC CCCC
    |||| ||||
    ||++-++++- Select 4 KB CHR bank to use with specified tile
    ++-------- Select palette to use with specified tile

In extended attribute mode, CHR banking behaves differently than normal when fetching background tiles from pattern tables:

- CHR mode (register \$5101) is ignored. CHR banks are always 4KB in this mode.
- The values of the CHR banking registers \$5120-\$512B are also ignored.
- Bits 0-5 specified here are used for selecting a 4KB CHR bank on a per-tile basis.
- The two bits in \$5130 are used globally as CHR bank bits 6 and 7.
- Driving pin 92 low disables extended attribute mode. Extended attribute mode is also automatically disabled based on PPUMASK monitoring. (See section on \$2001 monitoring.) In these cases, the non-extended attribute table is used instead.

In other words, this works as if the nametable was extended from 8-bit to 14-bit tile offsets, with the ExRAM holding the upper 6-bits and the 2-bit palette value, while the nametable selected through \$5105 contains the lower 8 bits.

*Just Breed*, *Yakuman Tengoku*, and the Koei games use extended attributes continuously.

The CHR address bus is mapped as follows when the PPU is fetching background pattern data while rendering extended attributes:

| CHR Ax  | 4 KiB (only) |
|---------|--------------|
| **A19** | \$5130.1     |
| **A18** | \$5130.0     |
| **A17** | ExRAM D5     |
| **A16** | ExRAM D4     |
| **A15** | ExRAM D3     |
| **A14** | ExRAM D2     |
| **A13** | ExRAM D1     |
| **A12** | ExRAM D0     |
| **A11** | PPU A11      |
| **A10** | PPU A10      |

Note, the CHR address bus is completely unaffected by CHR mode in this case.

#### <span id="Nametable_mapping_.28.245105.29"></span>Nametable mapping (\$5105) {#Nametable_mapping_.28.245105.29}

- Write only
- Reset detection not used

<!-- -->

    7  bit  0
    ---- ----
    DDCC BBAA
    |||| ||||
    |||| ||++- Select nametable at PPU $2000-$23FF
    |||| ++--- Select nametable at PPU $2400-$27FF
    ||++------ Select nametable at PPU $2800-$2BFF
    ++-------- Select nametable at PPU $2C00-$2FFF

Nametable values:

- 0 - CIRAM page 0
- 1 - CIRAM page 1
- 2 - Internal extended RAM
  - When \$5104 is set to mode %10 or %11, the nametable will read as all zeros. This does not share functionality with fill mode.
- 3 - Fill-mode data
  - See registers \$5106 and \$5017

Examples:

| Mode                    | Value | NTD | NTC | NTB | NTA |
|-------------------------|-------|-----|-----|-----|-----|
| Vertical arrangement    | \$50  | %01 | %01 | %00 | %00 |
| Horizontal arrangement  | \$44  | %01 | %00 | %01 | %00 |
| Single-screen CIRAM 0   | \$00  | %00 | %00 | %00 | %00 |
| Single-screen CIRAM 1   | \$55  | %01 | %01 | %01 | %01 |
| Single-screen ExRAM     | \$AA  | %10 | %10 | %10 | %10 |
| Single-Screen Fill-mode | \$FF  | %11 | %11 | %11 | %11 |
| Diagonal                | \$14  | %00 | %01 | %01 | %00 |

#### <span id="Fill-mode_tile_.28.245106.29"></span>Fill-mode tile (\$5106) {#Fill-mode_tile_.28.245106.29}

- Write only
- Reset detection not used

When a nametable is mapped to fill-mode in register \$5105, all nametable fetches get replaced by the value of this register. Only the nametable is affected by fill mode. When the PPU later uses this information to fetch the corresponding tile from the pattern table, CHR banking is unaffected and continues to work normally. §

#### <span id="Fill-mode_color_.28.245107.29"></span>Fill-mode color (\$5107) {#Fill-mode_color_.28.245107.29}

- Write only
- Reset detection not used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xxAA
           ||
           ++- Specify background palette index to use for fill-mode nametable

When a nametable is mapped to fill-mode in register \$5105, and \$5104 is not in mode %01, all attribute table fetches get replaced by the value of this register. Each byte of the attribute table normally contains four 2-bit palette indexes. The two bits in this register are copied for all four indexes.

When \$5104 is in mode %01, extended attribute mode does apply for fill mode, and this register is ignored. However, if pin 92 is driven low in this mode, extended attribute mode becomes disabled and this register comes back into effect.

### <span id="PRG_Bankswitching_.28.245113-.245117.29"></span>PRG Bankswitching (\$5113-\$5117) {#PRG_Bankswitching_.28.245113-.245117.29}

- Write only
- Reset detection is used only for \$5117

In general, when the CPU accesses an address that corresponds to the range of a PRG bank designated by the present PRG mode, the bits of that PRG bank register are applied to the appropriate PRG address buses as follows:

    7  bit  0
    ---- ----
    1111 1111  Reset Value ($5117 only)
    ---- ----
    RAAA AaAA
    |||| ||||
    |||| |||+- PRG ROM/RAM A13
    |||| ||+-- PRG ROM/RAM A14
    |||| |+--- PRG ROM/RAM A15, also selecting between PRG RAM /CE 0 and 1
    |||| +---- PRG ROM/RAM A16
    |||+------ PRG ROM A17
    ||+------- PRG ROM A18
    |+-------- PRG ROM A19
    +--------- RAM/ROM toggle (0: RAM; 1: ROM) (registers $5114-$5116 only)

Bank register effective areas versus PRG mode:

<table class="wikitable">
<thead>
<tr>
<th></th>
<th colspan="4">CPU memory affected for each mode (see <a href="#PRG_mode_.28.245100.29">#PRG mode ($5100)</a>)</th>
</tr>
</thead>
<tbody>
<tr>
<th>CPU Address</th>
<th scope="col" style="width: 125px">Mode 3 Registers</th>
<th scope="col" style="width: 125px">Mode 2 Registers</th>
<th scope="col" style="width: 125px">Mode 1 Registers</th>
<th scope="col" style="width: 125px">Mode 0 Registers</th>
</tr>
&#10;<tr>
<td><strong>$6000-7FFF</strong></td>
<td>$5113 (RAM only)</td>
<td>$5113 (RAM only)</td>
<td>$5113 (RAM only)</td>
<td>$5113 (RAM only)</td>
</tr>
<tr>
<td><strong>$8000-9FFF</strong></td>
<td>$5114</td>
<td rowspan="2">$5115</td>
<td rowspan="2">$5115</td>
<td rowspan="4">$5117 (ROM only)</td>
</tr>
<tr>
<td><strong>$A000-BFFF</strong></td>
<td>$5115</td>
</tr>
<tr>
<td><strong>$C000-DFFF</strong></td>
<td>$5116</td>
<td>$5116</td>
<td rowspan="2">$5117 (ROM only)</td>
</tr>
<tr>
<td><strong>$E000-FFFF</strong></td>
<td>$5117 (ROM only)</td>
<td>$5117 (ROM only)</td>
</tr>
</tbody>
</table>

Register \$5113 always maps RAM and bits 7, 6, 5, and 4 are always ignored. Based on the letterless MMC5 die image, \$5117.7 is fully implemented, but testing has confirmed that the bit is prevented from ever entering RAM mode for an unknown reason. Since \$5117 appears to only map ROM, it is not possible to map the interrupt vectors to RAM in any mode. All known games have their reset vector in the last bank of PRG ROM, and the vector points to an address greater than or equal to \$E000. This tells us that \$5117 must have a reliable power-on value of \$FF.

When a bankswitch register controls an 8kByte CPU address range, register bits 6..0 correspond to PRG A19..A13.

In PRG modes where a register controls a 16kByte CPU address range, register bits 6..1 correspond to PRG A19..A14. Register bit 0 is ignored and instead CPU A13 directly controls PRG A13. For example, comparing mode 3 to mode 2 for CPU address range \$8000-BFFF. These are equivalent:

- In mode 3, write value \$90 to \$5114 and value \$91 to \$5115.
- In mode 2, write value \$90 to \$5115.

However, these are not equivalent:

- In mode 3, write value \$91 to \$5114 and value \$92 to \$5115.
- In mode 2, write value \$91 to \$5115.

The MMC5 ignores register bit 0 in the 16kByte bank.

Similarly, when register \$5117 controls a 32kByte CPU address range in mode 0, register bits 6..2 correspond to PRG A19..A15, register bits 1..0 are ignored, and CPU A14..A13 directly control PRG A14..A13

<span id="Separate_PRG-ROM_and_PRG-RAM_Address_Busses"></span>

#### Separate PRG-ROM and PRG-RAM Address Buses {#Separate_PRG-ROM_and_PRG-RAM_Address_Buses}

The MMC5 has separate sets of address pins for PRG-ROM and PRG-RAM. This is a concept that has proven difficult to understand and explain. It all stems from the fact that CPU A15 is not directly routed to the cartridge; the signal /ROMSEL is supplied instead. Though it is possible to figure out CPU A15 from /ROMSEL, it takes extra time. This reduces the setup time where the address bus is valid before enabling the RAM and ROM /CEs. To avoid this, the MMC5 *ignores* /ROMSEL (and therefore CPU A15) when it comes to the PRG address, and /ROMSEL only figures into the logic of PRG ROM and RAM /CEs. This effectively creates a mirror on the PRG address buses for CPU address range \$0000-7FFF and \$8000-FFFF. Specifically, this makes range \$6000-7FFF indistinguishable from \$E000-FFFF. Because the MMC5 wanted to have separately controllable mapping for those ranges without reducing the setup time, its solution was to make two separate PRG address buses.

The MMC5's PRG-ROM address pins follow this logic. Though the PRG-ROM address bus is decoded at *all* CPU addresses, the gray areas always have PRG ROM /CE disabled:

<table class="wikitable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>CPU Address</th>
<th scope="col" style="width: 125px">Mode 3 PRG-ROM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 2 PRG-ROM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 1 PRG-ROM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 0 PRG-ROM<br />
Address Source</th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>$0000-1FFF</strong></td>
<td style="background: #cccccc">$5114</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="4" style="background: #cccccc">$5117</td>
</tr>
<tr>
<td><strong>$2000-3FFF</strong></td>
<td style="background: #cccccc">$5115</td>
</tr>
<tr>
<td><strong>$4000-5FFF</strong></td>
<td style="background: #cccccc">$5116</td>
<td style="background: #cccccc">$5116</td>
<td rowspan="2" style="background: #cccccc">$5117</td>
</tr>
<tr>
<td><strong>$6000-7FFF</strong></td>
<td style="background: #cccccc">$5117</td>
<td style="background: #cccccc">$5117</td>
</tr>
<tr>
<td><strong>$8000-9FFF</strong></td>
<td style="background: #80C0FF">$5114</td>
<td rowspan="2" style="background: #80C0FF">$5115</td>
<td rowspan="2" style="background: #80C0FF">$5115</td>
<td rowspan="4" style="background: #80C0FF">$5117</td>
</tr>
<tr>
<td><strong>$A000-BFFF</strong></td>
<td style="background: #80C0FF">$5115</td>
</tr>
<tr>
<td><strong>$C000-DFFF</strong></td>
<td style="background: #80C0FF">$5116</td>
<td style="background: #80C0FF">$5116</td>
<td rowspan="2" style="background: #80C0FF">$5117</td>
</tr>
<tr>
<td><strong>$E000-FFFF</strong></td>
<td style="background: #80C0FF">$5117</td>
<td style="background: #80C0FF">$5117</td>
</tr>
</tbody>
</table>

The MMC5's PRG-RAM address pins follow this logic, again with PRG-RAM /CE always disabled in the gray areas:

<table class="wikitable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>CPU Address</th>
<th scope="col" style="width: 125px">Mode 3 PRG-RAM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 2 PRG-RAM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 1 PRG-RAM<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 0 PRG-RAM<br />
Address Source</th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>$0000-1FFF</strong></td>
<td style="background: #cccccc">$5114</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="3" style="background: #cccccc">$5113*</td>
</tr>
<tr>
<td><strong>$2000-3FFF</strong></td>
<td style="background: #cccccc">$5115</td>
</tr>
<tr>
<td><strong>$4000-5FFF</strong></td>
<td style="background: #cccccc">$5116</td>
<td style="background: #cccccc">$5116</td>
<td style="background: #cccccc">$5113*</td>
</tr>
<tr>
<td><strong>$6000-7FFF</strong></td>
<td style="background: #FFC0C0">$5113</td>
<td style="background: #FFC0C0">$5113</td>
<td style="background: #FFC0C0">$5113*</td>
<td style="background: #FFC0C0">$5113*</td>
</tr>
<tr>
<td><strong>$8000-9FFF</strong></td>
<td style="background: #FFC0C0">$5114</td>
<td rowspan="2" style="background: #FFC0C0">$5115</td>
<td rowspan="2" style="background: #FFC0C0">$5115</td>
<td rowspan="4" style="background: #cccccc">$5113*</td>
</tr>
<tr>
<td><strong>$A000-BFFF</strong></td>
<td style="background: #FFC0C0">$5115</td>
</tr>
<tr>
<td><strong>$C000-DFFF</strong></td>
<td style="background: #FFC0C0">$5116</td>
<td style="background: #FFC0C0">$5116</td>
<td rowspan="2" style="background: #cccccc">$5113*</td>
</tr>
<tr>
<td><strong>$E000-FFFF</strong></td>
<td style="background: #cccccc">$5113</td>
<td style="background: #cccccc">$5113</td>
</tr>
</tbody>
</table>

\*In these larger areas of \$5113, the selected 8kByte bank in \$5113 is mirrored to fill the region.

If we overlap these two tables, the original table reemerges. Pink always has /CE enabled for RAM (using \$5113), Blue always has /CE enabled for ROM (using \$5117), and Purple chooses which /CE using the register bit 7. Gray always has RAM /CE and ROM /CE disabled at all times.

<table class="wikitable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>CPU Address</th>
<th scope="col" style="width: 125px">Mode 3 PRG<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 2 PRG<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 1 PRG<br />
Address Source</th>
<th scope="col" style="width: 125px">Mode 0 PRG<br />
Address Source</th>
</tr>
</thead>
<tbody>
<tr>
<td><strong>$0000-1FFF</strong></td>
<td style="background: #cccccc">$5114</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="2" style="background: #cccccc">$5115</td>
<td rowspan="3" style="background: #cccccc">$5113/$5117</td>
</tr>
<tr>
<td><strong>$2000-3FFF</strong></td>
<td style="background: #cccccc">$5115</td>
</tr>
<tr>
<td><strong>$4000-5FFF</strong></td>
<td style="background: #cccccc">$5116</td>
<td style="background: #cccccc">$5116</td>
<td style="background: #cccccc">$5113/$5117</td>
</tr>
<tr>
<td><strong>$6000-7FFF</strong></td>
<td style="background: #FFC0C0">$5113/$5117</td>
<td style="background: #FFC0C0">$5113/$5117</td>
<td style="background: #FFC0C0">$5113/$5117</td>
<td style="background: #FFC0C0">$5113/$5117</td>
</tr>
<tr>
<td><strong>$8000-9FFF</strong></td>
<td style="background: #D0C0E0">$5114</td>
<td rowspan="2" style="background: #D0C0E0">$5115</td>
<td rowspan="2" style="background: #D0C0E0">$5115</td>
<td rowspan="4" style="background: #80C0FF">$5113/$5117</td>
</tr>
<tr>
<td><strong>$A000-BFFF</strong></td>
<td style="background: #D0C0E0">$5115</td>
</tr>
<tr>
<td><strong>$C000-DFFF</strong></td>
<td style="background: #D0C0E0">$5116</td>
<td style="background: #D0C0E0">$5116</td>
<td rowspan="2" style="background: #80C0FF">$5113/$5117</td>
</tr>
<tr>
<td><strong>$E000-FFFF</strong></td>
<td style="background: #80C0FF">$5113/$5117</td>
<td style="background: #80C0FF">$5113/$5117</td>
</tr>
</tbody>
</table>

#### PRG-RAM configurations {#PRG-RAM_configurations}

In [commercial configurations](../ExROM/), bits 0 and 1 select pages *within* an SRAM chip, and bit 2 selects between two separate SRAMs. 8K and 32K games have a single SRAM chip that will only be active when bit 2 is clear. 16K games instead have two chips, but only the first is battery backed.

<table class="wikitable">
<thead>
<tr>
<th>Configuration</th>
<th colspan="8">bank value &amp; 7</th>
</tr>
</thead>
<tbody>
<tr>
<th>0</th>
<th>1</th>
<th>2</th>
<th>3</th>
<th>4</th>
<th>5</th>
<th>6</th>
<th>7</th>
<th></th>
</tr>
&#10;<tr>
<th>ELROM 0K</th>
<td><a href="../Open_bus_behavior/">open bus</a></td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
</tr>
<tr>
<th>EKROM 8K (1 x 8K)</th>
<td>0:$0000</td>
<td>0:$0000</td>
<td>0:$0000</td>
<td>0:$0000</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
</tr>
<tr>
<th>ETROM 16K (2 x 8K)</th>
<td><strong>0</strong>:$0000</td>
<td><strong>0</strong>:$0000</td>
<td><strong>0</strong>:$0000</td>
<td><strong>0</strong>:$0000</td>
<td><strong>1</strong>:$0000</td>
<td><strong>1</strong>:$0000</td>
<td><strong>1</strong>:$0000</td>
<td><strong>1</strong>:$0000</td>
</tr>
<tr>
<th>EWROM 32K (1 x 32K)</th>
<td>0:$<strong>0000</strong></td>
<td>0:$<strong>2000</strong></td>
<td>0:$<strong>4000</strong></td>
<td>0:$<strong>6000</strong></td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
<td>open bus</td>
</tr>
<tr>
<th>Superset 64K (2 x 32K)</th>
<td><strong>0</strong>:$<strong>0000</strong></td>
<td><strong>0</strong>:$<strong>2000</strong></td>
<td><strong>0</strong>:$<strong>4000</strong></td>
<td><strong>0</strong>:$<strong>6000</strong></td>
<td><strong>1</strong>:$<strong>0000</strong></td>
<td><strong>1</strong>:$<strong>2000</strong></td>
<td><strong>1</strong>:$<strong>4000</strong></td>
<td><strong>1</strong>:$<strong>6000</strong></td>
</tr>
</tbody>
</table>

Since [iNES](../INES/) headers were lacking reliable PRG-RAM size information before [NES 2.0](../NES_2.0/), some emulators may have selected these behaviours through ROM CRC checks.

Because no ExROM game is known to write PRG-RAM with one bank value and then attempt to read back the same data with a different bank value, emulating the PRG-RAM as 64K at all times can be used as a compatible superset for all games.

Investigation of the [MMC5 pinout](../MMC5_pinout/) in 2018 revealed that bits 2 and 3 also control additional PRG-RAM address pins, which could theoretically have been used to select 16 banks of a single 128K SRAM, with bit 2 controlling PRG A15 directly rather than using the two chip select /CE outputs.

#### Other PRG-RAM notes {#Other_PRG-RAM_notes}

- *Bandit Kings of Ancient China* maps PRG-RAM to the CPU \$8000+ area and expects to be able to write to it through there. Failure to emulate this causes corruption when the background is restored on the world map.
- *Uncharted Waters* requires emulating bankswitching of PRG-RAM: it writes to PRG-RAM at one CPU address and expects to be able to read the same data back via a different CPU address.
- Games with 16K PRG-RAM only battery-save the first 8K.
- [List of MMC5 games which include a battery](http://bootgod.dyndns.org:7777/search.php?ines=5&battery=Yes&group=groupid).

### <span id="CHR_Bankswitching_.28.245120-.245130.29"></span>CHR Bankswitching (\$5120-\$5130) {#CHR_Bankswitching_.28.245120-.245130.29}

When using 8x8 sprites, only registers \$5120-\$5127 are used. Registers \$5128-\$512B are completely ignored.

When using 8x16 sprites, the PPU ignores the [sprite pattern table address](../PPU_registers/) and can select tiles from the entire 8 KiB of pattern tables, which on other mappers overlaps with the background pattern table. The MMC5 keeps track of whether the PPU is fetching background tiles or sprite tiles, and has new registers to specify independent banks for the background tiles even if they appear to be the same address from the PPU. This effectively creates a CHR window of 12 KiB, with up to eight 1 KiB banks of sprites available simultaneously. Registers \$5120-\$5127 specify banks for sprites, registers \$5128-\$512B apply to background tiles, and the last set of registers written to (either \$5120-\$5127 or \$5128-\$512B) will be used for I/O via [PPUDATA](../PPU_registers/) (\$2007). [\[1\]](https://forums.nesdev.org/viewtopic.php?p=193069#p193069) [\[2\]](http://forums.nesdev.org/viewtopic.php?p=194120#p194120) The MMC5 knows that sprite data is being fetched by counting the number of fetches since the last detected scanline start, similar to how it detects the position for the vertical split.

*Bandit Kings of Ancient China* and *Uchuu Keibitai SDF* have non-pattern data stored in CHR ROM, read out via \$2007.

The MMC5 is known to listen to the same address as the PPU to find out when to enable the 8x16 sprite mode; see [above](#8x16_mode_enable_1_.28.242000_.3D_PPUCTRL.29).

#### <span id="CHR_select_.245120-.24512B"></span>CHR select \$5120-\$512B {#CHR_select_.245120-.24512B}

- Write only
- Reset detection not used

<table class="wikitable">
<thead>
<tr>
<th></th>
<th colspan="4">PPU memory affected for each mode (see <a href="#CHR_mode_.28.245101.29">#CHR mode ($5101)</a>)</th>
</tr>
</thead>
<tbody>
<tr>
<th>Register</th>
<th>1 KiB</th>
<th>2 KiB</th>
<th>4 KiB</th>
<th>8 KiB</th>
</tr>
&#10;<tr>
<td><strong>$5120</strong></td>
<td>$0000-$03FF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5121</strong></td>
<td>$0400-$07FF</td>
<td>$0000-$07FF</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5122</strong></td>
<td>$0800-$0BFF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5123</strong></td>
<td>$0C00-$0FFF</td>
<td>$0800-$0FFF</td>
<td>$0000-$0FFF</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5124</strong></td>
<td>$1000-$13FF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5125</strong></td>
<td>$1400-$17FF</td>
<td>$1000-$17FF</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5126</strong></td>
<td>$1800-$1BFF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5127</strong></td>
<td>$1C00-$1FFF</td>
<td>$1800-$1FFF</td>
<td>$1000-$1FFF</td>
<td>$0000-$1FFF</td>
</tr>
<tr>
<td><strong>$5128</strong></td>
<td>$0000-$03FF and $1000-$13FF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$5129</strong></td>
<td>$0400-$07FF and $1400-$17FF</td>
<td>$0000-$07FF and $1000-$17FF</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$512A</strong></td>
<td>$0800-$0BFF and $1800-$1BFF</td>
<td>none</td>
<td>none</td>
<td>none</td>
</tr>
<tr>
<td><strong>$512B</strong></td>
<td>$0C00-$0FFF and $1C00-$1FFF</td>
<td>$0800-$0FFF and $1800-$1FFF</td>
<td>$0000-$0FFF and $1000-$1FFF</td>
<td>$0000-$1FFF</td>
</tr>
</tbody>
</table>

**Caution:** Unlike the MMC1 and unlike PRG banking on the MMC5, the banks are always indexed by the *currently selected size*. When using 2kb, 4kb or 8kb bank sizes, the registers hold bank index of that larger size, and lower bits are \*not\* ignored.

Technically, \$5120-512B are 10-bit registers inside the MMC5. When writing to these registers, all 10 bits are written. The source source of the bits depends on the active CHR mode.

| \$512x Bit | 1 KiB    | 2 KiB    | 4 KiB  | 8 KiB  |
|------------|----------|----------|--------|--------|
| **9**      | \$5130.1 | \$5130.0 | CPU D7 | CPU D6 |
| **8**      | \$5130.0 | CPU D7   | CPU D6 | CPU D5 |
| **7**      | CPU D7   | CPU D6   | CPU D5 | CPU D4 |
| **6**      | CPU D6   | CPU D5   | CPU D4 | CPU D3 |
| **5**      | CPU D5   | CPU D4   | CPU D3 | CPU D2 |
| **4**      | CPU D4   | CPU D3   | CPU D2 | CPU D1 |
| **3**      | CPU D3   | CPU D2   | CPU D1 | CPU D0 |
| **2**      | CPU D2   | CPU D1   | CPU D0 | ?      |
| **1**      | CPU D1   | CPU D0   | ?      | ?      |
| **0**      | CPU D0   | CPU D0   | CPU D0 | CPU D0 |

The mapping of the 10-bit bank registers to the CHR address bus also depends on CHR mode as follows:

| CHR Ax  | 1 KiB    | 2 KiB    | 4 KiB    | 8 KiB    |
|---------|----------|----------|----------|----------|
| **A19** | \$512x.9 | \$512x.9 | \$512x.9 | \$512x.9 |
| **A18** | \$512x.8 | \$512x.8 | \$512x.8 | \$512x.8 |
| **A17** | \$512x.7 | \$512x.7 | \$512x.7 | \$512x.7 |
| **A16** | \$512x.6 | \$512x.6 | \$512x.6 | \$512x.6 |
| **A15** | \$512x.5 | \$512x.5 | \$512x.5 | \$512x.5 |
| **A14** | \$512x.4 | \$512x.4 | \$512x.4 | \$512x.4 |
| **A13** | \$512x.3 | \$512x.3 | \$512x.3 | \$512x.3 |
| **A12** | \$512x.2 | \$512x.2 | \$512x.2 | PPU A12  |
| **A11** | \$512x.1 | \$512x.1 | PPU A11  | PPU A11  |
| **A10** | \$512x.0 | PPU A10  | PPU A10  | PPU A10  |

The gray boxes in the first table, which might be considered garbage data, correspond to the PPU Ax -\> CHR Ax passthrough bits of the larger size banks. When changing CHR modes, the contents of the 10-bit CHR bank registers (\$512x) are unaffected, and therefore the bits in the gray boxes can and would come into play.

#### <span id="Upper_CHR_Bank_bits_.28.245130.29"></span>Upper CHR Bank bits (\$5130) {#Upper_CHR_Bank_bits_.28.245130.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---- ----
    xxxx xx00  Reset Value
    ---- ----
    xxxx xxBB
           ||
           ++- Highest CHR bank bits

The \$5130 register has different uses in different situations.

##### Sprite Rendering and Normal Background Rendering {#Sprite_Rendering_and_Normal_Background_Rendering}

When the PPU is fetching sprite pattern data in any mode, or fetching background pattern data during normal rendering, the 10-bit CHR banks \$5120-512B are used. Depending on the CHR mode, writing to \$5120-\$512B can append one or both the bits from this register to the actual 10-bit bank value that gets stored. Regardless if \$5130 is changed thereafter, the 10-bit registers continue to store their value from the time they were written.

When the MMC5 is using 2KB/1KB CHR banks, only 512KB/256KB of CHR ROM can be selected using the previous registers. To access all 1024KB in those modes, first write the upper bit(s) to register \$5130 and then write the lower bits to \$5120-\$512B.

The only ExROM game with CHR ROM larger than 256KB is *Metal Slader Glory*, which uses 4KB CHR banks and does not use extended attributes. In other words, no official game relies on this register, and most don't even initialize it.

##### Extended Attribute Background Rendering {#Extended_Attribute_Background_Rendering}

When the PPU is fetching background pattern data with extended attribute mode active, the applicable 6-bit CHR bank value from ExRAM is used, and the real-time value of \$5130 is appended as the higher 2 bits. In other words, \$5130 selects in real-time which 256KB of CHR ROM is to be used for all extended attribute background tiles on the screen.

##### Vertical Split Background Rendering {#Vertical_Split_Background_Rendering}

When the PPU is fetching background pattern data for the vertical split region, this register has no effect at all whatsoever.

### Other Registers {#Other_Registers}

#### <span id="Vertical_Split_Mode_.28.245200.29"></span>Vertical Split Mode (\$5200) {#Vertical_Split_Mode_.28.245200.29}

- Write only
- Reset detection is used

<!-- -->

    7  bit  0
    ---- ----
    00x0 0000  Reset Value
    ---- ----
    ESxW WWWW
    || | ||||
    || +-++++- Specify vertical split threshold tile count
    |+-------- Specify vertical split region screen side (0:left; 1:right)
    +--------- Enable vertical split mode (0:disabled; 1:enabled)

*Uchuu Keibitai SDF* uses split screen mode during the intro, where it shows ship stats. *Bandit Kings of Ancient China* uses split screen mode during the ending sequence<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>.

MMC5 internal extended RAM is always used as the nametable in split screen mode. Extended RAM mode (\$5104) should be set to %00 or %01. RAM modes %10 or %11 disable split mode. Driving pin 92 low also disables split mode. Split mode is automatically disabled based on PPUMASK monitoring. (See section on \$2001 monitoring)

When using mode %01 with vertical split mode, the non-split region gains extended attributes, and the split region does not use extended attributes. As long as the non-split region remains horizontally scrolled to align with the left edge of a nametable, it is possible to share the extended RAM for both purposes. For example, a split region on the right side of the screen will read its nametable data and non-extended attribute data from the “right side” of the extended RAM, and the left side of the screen will read its nametable data from wherever it is assigned in \$5105, and get its extended attributes from the “left side” of the extended RAM. The two separate functions of extended RAM data would not overlap this way.

##### Method of Operation {#Method_of_Operation}

Based on nametable fetch count (horizontally) and scanline count (vertically), the MMC5 makes substitutions to the nametable data, attribute table data, and the lowest 3 address bits of the pattern data.

##### Horizontal Behavior {#Horizontal_Behavior}

For each visible scanline, the PPU fetches nametable data for 34 background tiles. The MMC5 locates the split threshold by counting these fetches and comparing the count to the count value stored in register \$5200. When the PPU is rendering in the split region, the MMC5 replaces the normal nametable data with the split region nametable data. Because the split threshold is located based on nametable fetch count and not based on the actual nametable address requested by the PPU, the location of the split is more of an absolute position on the screen, always positioned the same number of tiles from the left side of the screen. The PPU’s fine horizontal scrolling can cause the first tile rendered to be of fractional width, however the vertical split threshold will always occur at a full tile boundary, so the exact location of the split threshold is affected by fine horizontal scrolling, potentially deviating by up to 7 pixels this way.

Left Split:

- Tiles 0 to T-1 are the split region.
- Tiles T and on are rendered normally.

Right Split:

- Tiles 0 to T-1 are rendered normally.
- Tiles T and on are the split region.

There is no intended horizontal scrolling of any kind for the split region. Right-side split regions will always remain right-aligned with the right-hand side of the nametable, and left-side split regions will always remain left-aligned with the left-hand side of the nametable. Coarse horizontal scrolling can still be used for the non-split region.

##### Vertical Behavior {#Vertical_Behavior}

The MMC5 is able to have its own vertical origin in the split region because the nametable data that it substitutes is based on scanline count, which is always aligned to the top of the screen.

The MMC5 has a separate v-split scanline counter that resets to the vertical scrolling value in \$5201 each v-blank. It uses this to know what nametable data to substitute in the split region on each scanline. Though the PPU continues requesting nametable data corresponding to its normal vertical scrolling in the split region, the MMC5 is ignoring the nametable address requested in this case and directly substituting its own nametable data.

One additional caveat is fine vertical scrolling of the split region. For each tile on each scanline, the PPU reads from a nametable, and uses that information to fetch the pattern data for 1 row of pixels for that tile. The PPU normally keeps track of which row to fetch using PPU A0..2. However, if the MMC5’s fine vertical scrolling doesn’t match the PPU’s fine vertical scrolling, it won’t be fetching the correct row of pixels from the tile. This is why the MMC5 provides CHR A0..2 from the lowest 3 bits of the v-split scanline counter when rendering the split region. By default, Nintendo PCBs are wired in “CL Mode”, which prevents the correct fine vertical scrolling of the split region. (See [MMC5 pinout](../MMC5_pinout/)).

Vertical scrolling for the split region operates like normal vertical scrolling. 0-239 are valid scroll values, whereas 240-255 will display attribute table data as nametable data for the first few scanlines. The MMC5 does proper vertical mirroring of the split region nametable so that scrolling down properly rolls over to the top of the nametable, skipping the attribute table data that would naturally be located there.

#### <span id="Vertical_Split_Scroll_.28.245201.29"></span>Vertical Split Scroll (\$5201) {#Vertical_Split_Scroll_.28.245201.29}

- Write only
- Reset detection is used, value \$00

<!-- -->

     All eight bits specify the vertical scroll value to use in split region

MMC5 boards wired in "CL" mode should only use vertical scroll values whose bottom 3 bits match the [PPU](../PPU/)'s fine vertical scroll value. Using a mismatched value will cause tiles to seem to "roll" within themselves. In "SL" mode, any values can be used. (No existing games used the SL board configuration.)

Horizontal scrolling is not allowed within the split region.

#### <span id="Vertical_Split_Bank_.28.245202.29"></span>Vertical Split Bank (\$5202) {#Vertical_Split_Bank_.28.245202.29}

- Write only
- Reset detection not used

<!-- -->

     All eight bits select a 4 KB CHR bank at $0000-$0FFF and $1000-$1FFF while rendering the split region.

The CHR address bus is mapped as follows when the PPU is fetching background pattern data while rendering the vertical split region:

| CHR Ax  | 4 KiB (only) |
|---------|--------------|
| **A19** | \$5202.7     |
| **A18** | \$5202.6     |
| **A17** | \$5202.5     |
| **A16** | \$5202.4     |
| **A15** | \$5202.3     |
| **A14** | \$5202.2     |
| **A13** | \$5202.1     |
| **A12** | \$5202.0     |
| **A11** | PPU A11      |
| **A10** | PPU A10      |

Note, the CHR address bus is completely unaffected by CHR mode and \$5130 in this case.

#### <span id="IRQ_Scanline_Compare_Value_.28.245203.29"></span>IRQ Scanline Compare Value (\$5203) {#IRQ_Scanline_Compare_Value_.28.245203.29}

- Write only
- Reset detection not used

All eight bits specify the target scanline number at which to generate a scanline IRQ. Value \$00 is a special case that will not produce IRQ pending conditions, though it is possible to get an IRQ while this is set to \$00 (due to the pending flag being set already.) You will need to take additional measures to fully suppress the IRQ. See the detailed discussion.

#### <span id="Scanline_IRQ_Status_.28.245204.2C_read.2Fwrite.29"></span>Scanline IRQ Status (\$5204, read/write) {#Scanline_IRQ_Status_.28.245204.2C_read.2Fwrite.29}

- Read/Write
- Reset detection is used

##### Write {#Write}

    7  bit  0
    ---- ----
    0xxx xxxx  Reset Value
    ---- ----
    Exxx xxxx
    |
    +--------- Scanline IRQ Enable flag (1=enabled)

##### Read {#Read}

    7  bit  0
    ---- ----
    SVxx xxxJ  MMC5A default power-on value = $00
    ||      |
    ||      +- Abnormal M2 vs CPU A0 sequence detection flag
    |+-------- "In Frame" flag
    +--------- Scanline IRQ Pending flag

The Scanline IRQ Pending flag becomes set at any time that the internal scanline counter matches the value written to register \$5203. If the scanline IRQ is enabled, it will also generate /IRQ to the system.

The "In Frame" flag is set when the PPU is actively rendering visible scanlines and cleared when not rendering; for example, vertical blank scanlines. This flag does not clear for the short H-Blank period between visible scanlines. When pin 92 is driven low, this flag remains low at all times. Additionally, scanline IRQs no longer occur when pin 92 is driven low.

Bit 0 exists in both versions of MMC5. If an abnormal sequence of M2 vs CPU A0 occurs, there is a latch that stores this permanently until system reset detection. The letterless MMC5 additionally shows this latch on the CL3 pin. When no error has been detected, the bit is 0 and CL3 is high. Though the MMC5A has the same detection and status bit, no way to enable the function on the CL3 pin has been found. The function is not triggered by a CPU crash when triggered via JAM opcode. The intention of the latch is not known, but it may be involved in detecting when the player turns off the power without holding reset.

Any time this register is read, the Scanline IRQ Pending flag is automatically cleared (acknowledging the IRQ). Acknowledging the IRQ does not clear the scanline counter, and you may write a larger value to \$5203 within a scanline ISR in order to generate another IRQ on an additional scanline within the same frame.

For details, see [§ Scanline Detection and Scanline IRQ](#Scanline_Detection_and_Scanline_IRQ).

#### <span id="Unsigned_8x8_to_16_Multiplier_.28.245205.2C_.245206_read.2Fwrite.29"></span>Unsigned 8x8 to 16 Multiplier (\$5205, \$5206 read/write) {#Unsigned_8x8_to_16_Multiplier_.28.245205.2C_.245206_read.2Fwrite.29}

- Write only
- Reset detection not used

The unsigned 16-bit product is available to be read from these registers immediately after writing. All 65536 combinations of multiplicand and multiplier were tested and verified correct on MMC5A here[\[3\]](https://forums.nesdev.org/viewtopic.php?p=226537#p226537).

##### Write {#Write_2}

- \$5205 8-bit Unsigned Multiplicand
- \$5206 8-bit Unsigned Multiplier

##### Read {#Read_2}

- \$5205 Unsigned 16-bit Product (low byte)
- \$5206 Unsigned 16-bit Product (high byte)

### <span id="Internal_extended_RAM_.28.245C00-.245FFF.2C_read.2Fwrite.29"></span>Internal extended RAM (\$5C00-\$5FFF, read/write) {#Internal_extended_RAM_.28.245C00-.245FFF.2C_read.2Fwrite.29}

Refer to register \$5104 for special behaviors of the MMC5's 1KB internal extended RAM.

## MMC5A {#MMC5A}

The MMC5A was a later revision of MMC5 that included some extra features. Other than Just Breed mysteriously writing to \$5800, no game is known that uses these features. Therefore, these will probably tend to have poor support from emulators.

### MMC5A Registers {#MMC5A_Registers}

Registers \$5207, \$5208, \$5209, \$520A, and range \$5800-\$5BFF are present only in MMC5A.

#### <span id="CL3_.2F_SL3_Data_Direction_and_Output_Data_Source_.28MMC5A:_.245207_write_only.29"></span>CL3 / SL3 Data Direction and Output Data Source (MMC5A: \$5207 write only) {#CL3_.2F_SL3_Data_Direction_and_Output_Data_Source_.28MMC5A:_.245207_write_only.29}

    7  bit  0
    ---- ----
    11xx xx00  Reset Value
    ---- ----
    ABxx xxCD
    ||     ||
    ||     |+- MMC5.97 (CL3) Function (0 = I/O Control, 1 = $5800 read control
    ||     +-- MMC5.98 (SL3) Function (0 = I/O Control, 1 = $5800 write control
    |+-------- MMC5.97 (CL3) I/O Data Direction (0 = output, 1 = input)
    +--------- MMC5.98 (SL3) I/O Data Direction (0 = output, 1 = input)

When a function bit is set to 1, the pin becomes controlled by the applicable CPU writes or reads in the range \$5800-\$5BFF. The pin is normally driven high, and when the read or write occurs, it has a low pulse corresponding to !(M2). The intended purpose is unknown. It is theorized that it may attach a peripheral or additional RAM. When in this mode, none of the I/O configuration bits have any effect.

#### <span id="CL3_.2F_SL3_Status_.28MMC5A:_.245208_read.2Fwrite.29"></span>CL3 / SL3 Status (MMC5A: \$5208 read/write) {#CL3_.2F_SL3_Status_.28MMC5A:_.245208_read.2Fwrite.29}

##### Write {#Write_3}

    7  bit  0
    ---- ----
    00xx xxxx  Reset Value
    ---- ----
    ABxx xxxx
    ||
    |+-------- Value to be output on MMC5.97 pin (CL3) if it is set as an I/O output in $5207.
    +--------- Value to be output on MMC5.98 pin (SL3) if it is set as an I/O output in $5207.

Warning: The PCB may connect pins 97 and 98 directly to GND. Though not totally confirmed, setting them as output high while connected like this has appeared to break the output drivers of these pins. Also note that enabling the \$5800 function will have these pins driving high most of the time as well.

##### Read {#Read_3}

    7  bit  0
    ---- ----
    ABxx xxxx
    ||
    |+-------- Input value of MMC5.97 pin (CL3)
    +--------- Input value of MMC5.98 pin (SL3)

#### <span id="16-bit_Hardware_Timer_with_IRQ_.28MMC5A:_.245209_read.2Fwrite.2C_.24520A_write.29"></span>16-bit Hardware Timer with IRQ (MMC5A: \$5209 read/write, \$520A write) {#16-bit_Hardware_Timer_with_IRQ_.28MMC5A:_.245209_read.2Fwrite.2C_.24520A_write.29}

##### Read {#Read_4}

- \$5209

<!-- -->

    7  bit  0
    ---- ----
    Vxxx xxxx  MMC5A default power-on read value = $00
    |
    +--------- Hardware Timer IRQ Flag

##### Write {#Write_4}

- \$5209

<!-- -->

    7  bit  0
    ---- ----
    TTTT TTTT  MMC5A default power-on write value = $00
    |||| ||||
    ++++-++++- Timer count LSB

- \$520A

<!-- -->

    7  bit  0
    ---- ----
    TTTT TTTT  MMC5A default power-on write value = $00
    |||| ||||
    ++++-++++- Timer count MSB

Based on findings from krzysiobal: The timer automatically starts when writing any value to register \$5209, if the 16-bit timer value does not equal \$0000. For example, to write value \$0100, you would first write \$01 (MSB) to register \$520A, which does not start the timer. Then write \$00 (LSB) to register \$5209, which at that point will start the timer from value \$0100.

Each 8-bit value is written directly to an internal 16-bit counter that decrements each CPU cycle, specifically on the rising edge of M2. Additional writes while the timer is running will directly overwrite that portion of the counter. Reading register \$5209 while the timer is running reports \$00. The transition from counter value \$0001 to \$0000 generates an IRQ and sets the hardware timer IRQ flag. The timer stops at this point. Reading this register reports the IRQ flag, then automatically clears the IRQ and IRQ flag.

If the MMC5 detects a reset, it clears the timer if active, and it clears the IRQ and IRQ flag if set. Reset detection works by looking for a gap larger than about 11 usec on M2.

This register's IRQ operation is completely independent from register \$5204. Disabling interrupts through \$5204 has no effect on, and reading \$5204 does not report on, IRQs generated these registers.

#### <span id="Peripheral_Address_Range_.28MMC5A:_.245800-.245BFF.2C_read.2Fwrite.29"></span>Peripheral Address Range (MMC5A: \$5800-\$5BFF, read/write) {#Peripheral_Address_Range_.28MMC5A:_.245800-.245BFF.2C_read.2Fwrite.29}

Reads and writes in this address range are reflected on the CL3 and SL3 pins based on configuration in register \$5207. The purpose of this function is to connect an additional peripheral or dual port RAM for this address range. The MMC5A definitely does not have any built-in peripheral or RAM here.

Address \$5800 is written to by Just Breed. During each V-Blank, it writes value \$03, then \$01 to this address, reads and writes to PPU registers, then writes value \$00 to this address once complete. It is theorized that this was used as a debug signal to measure CPU usage / idle time during development.

## Scanline Detection and Scanline IRQ {#Scanline_Detection_and_Scanline_IRQ}

<figure class="mw-default-size mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/Mmc5_in_frame_status_bit.png"><img src="../../wiki-images/Mmc5_in_frame_status_bit.png" width="300" height="189" /></a>
<figcaption>MMC5 'in frame' status bit state diagram</figcaption>
</figure>

The MMC5 detects scanlines by first looking for three consecutive PPU reads from the same nametable address in the range \$2xxx. Once this has been seen, the next PPU read, regardless of address, is the point at which the scanline is detected. This works because the PPU does two matching dummy nametable byte reads at the end of each scanline, followed by a third matching nametable byte read at the beginning of the next scanline, followed by an attribute table byte read. So, the scanline gets detected when the PPU does the attribute table byte read, which is at PPU cycle 4.

Once that occurs, if the "in-frame" flag (register \$5204) was clear, it becomes set, and the internal 8-bit scanline counter is reset to zero; but if it was already set, the scanline counter is incremented, then compared against the value written to \$5203. If they match, the "irq pending" flag is set.

The IRQ pending flag is raised when the desired scanline is reached *regardless* of whether or not the scanline IRQ is enabled, i.e. even after a 0 was written to the scanline IRQ enable flag. However, an actual IRQ is only sent to the CPU if both the scanline IRQ enable flag and IRQ pending flag are set. A \$5203 value of \$00 is a special case where the comparison is never true. The MMC5's scanline IRQ occurs at PPU cycle 4, unlike the simpler scanline counter of the MMC3, which usually generates an IRQ around PPU cycle 260. See also [\[4\]](http://forums.nesdev.org/viewtopic.php?t=7653).

The "in-frame" flag is cleared when the PPU is no longer rendering. This is detected when 3 CPU cycles pass without a PPU read having occurred (PPU /RD has not been low during the last 3 M2 rises). The PPU does this in these conditions:

- the PPU begins the post-render scanline 240
- the PPU stops rendering because the user program wrote to CPU address \$2001 with bits 3 and 4 clear.
- note: the MMC5 may not listen directly to writes to \$2001 for this behavior, or it may be monitoring for particular transitions of \$2001 that have not yet been fully tested.

The "in frame" flag is cleared, scanline IRQ is automatically acknowledged, and the internal scanline counter is reset in any of these conditions:

- the V-blank NMI occurs, i.e. CPU reads the interrupt vector from addresses \$FFFA and \$FFFB
- the user program intentionally reads from CPU addresses \$FFFA or \$FFFB
- believed but not tested: the user program writes to \$4014 OAMDMA
- the 241st scanline is detected.
- when the MMC5 believes that \$2001 bits 3 and 4 are both clear and then either one of them gets set.

The scanline IRQ is acknowledged, but the "in frame" flag is not cleared and the scanline counter is not reset in any of these conditions:

- the user program reads register \$5204
- scanline 0 is detected

When system reset detection occurs, the only thing that happens is scanline IRQ becomes disabled. All other operation continues unaffected. These things happen any time that scanline IRQ becomes disabled:

- if /IRQ was low due to scanline IRQ pending flag, /IRQ returns high
- the scanline IRQ pending flag remains unaffected
- register \$5203 value remains unaffected
- scanline counter remains unaffected
- enabling the scanline IRQ will cause immediate /IRQ low if IRQ pending flag is set, but will not affect anything else

This means in pseudo-code:

    (On every PPU read -- PPU /RD falling edge)
    if address >= $2000 and address <= $2FFF and address == lastAddress
       matchCount := matchCount +1
       if matchCount == 2
         if inFrame == false
            inFrame := true
            scanline := 0
         else
            scanline := scanline +1
            if scanline == [$5203]
               irqPending := true
    else
       matchCount := 0
    lastAddress := address
    ppuIsReading := true

    (On every CPU cycle -- M2 rising edge)
    if ppuIsReading
       idleCount := 0
    else
       idleCount := idleCount +1      
       if idleCount == 3
          inFrame := false
          lastAddress := undefined
    ppuIsReading := false

    (On every CPU write)
    if address == $2001 and (value & $18) == 0
       inFrame := false
       lastAddress := undefined

    (On every CPU read)
    if address == $FFFA or address == $FFFB
       inFrame := false
       lastAddress := undefined

Please refer to the state diagrams on the right for a more formal description of the scanline and in-frame detection counters.

## Hardware {#Hardware}

The MMC5 exists in a 100-pin rectangular QFP package, see [MMC5 pinout](../MMC5_pinout/) for details.

MMC5 cartridge PCBs can be configured to different modes, see [ExROM](../ExROM/) for details.

At least two different versions of the MMC5 are known to exist: MMC5, and MMC5A. MMC5A has the addition of registers \$5207, \$5208, \$5209, and \$520A: SL3/CL3 control and hardware timer.

## References {#References}

1.  <span id="cite_note-WYGPNF-1"></span>[↑](#cite_ref-WYGPNF_1-0) "Why Your Game Paks Never Forget". *Nintendo Power* \#20 (March 1991), pp. 28-31.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [krzysiobal's RE notes](https://forums.nesdev.org/viewtopic.php?p=229375#p229375)
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) <https://forums.nesdev.org/viewtopic.php?f=2&t=12764>

## External links {#External_links}

- NES Mapper list by Disch [\[5\]](http://www.romhacking.net/documents/362/)
- Nintendo MMC5 by goroh, translated by Sgt. Bowhack [\[6\]](https://nesdev.org/mmc5-e.txt)
- Nintendo MMC5 Bankswitching by Kevin Horton [\[7\]](https://nesdev.org/mmc5_bank_switch.txt)
