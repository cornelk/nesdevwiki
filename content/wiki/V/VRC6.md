---
title: "VRC6"
url: "/wiki/VRC6/"
description: "Konami's VRC6 ASIC mapper comes in two variants:"
wikiPage: "VRC6"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
---

<div class="infobox"><div style="text-align: center"><b>VRC6</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=24">3 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>351951, 351949A </td></tr><tr><td><b>Pinout</b></td><td><a href="../VRC6_pinout/">VRC6 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 8K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td><a href="../VRC6_audio/">Yes</a></td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">024</a>, <a href="./">026</a></td></tr></tbody></table></div>

Konami's **VRC6** [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/) comes in two variants:

- **VRC6a** - **iNES Mapper 024** used for *Akumajou Densetsu* (Konami PCB 351951).
- **VRC6b** - **iNES Mapper 026** used for *Madara* and *Esper Dream 2* (Konami PCB 351949A).

The A0 and A1 lines are swapped between the two variants. The registers described on this page are for mapper 24, but for mapper 26 the register addresses must be adjusted (\$x001 becomes \$x002 and vice versa).

See [VRC6 pinout](../VRC6_pinout/) for the chip pinout.

## Banks {#Banks}

### CPU {#CPU}

- \$6000-\$7FFF: 8 KB PRG-RAM bank, fixed
- \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank

### PPU {#PPU}

- \$0000-\$03FF: 1 KB switchable CHR ROM bank
- \$0400-\$07FF: 1 KB switchable CHR ROM bank
- \$0800-\$0BFF: 1 KB switchable CHR ROM bank
- \$0C00-\$0FFF: 1 KB switchable CHR ROM bank
- \$1000-\$13FF: 1 KB switchable CHR ROM bank
- \$1400-\$17FF: 1 KB switchable CHR ROM bank
- \$1800-\$1BFF: 1 KB switchable CHR ROM bank
- \$1C00-\$1FFF: 1 KB switchable CHR ROM bank

## Registers {#Registers}

Only address lines 0, 1, and 12-15 are used for registers, therefore mirrors can be found by ANDing the address with \$F003 (\$DE6A mirrors \$D002)

The addresses described here are for mapper 24. The registers for mapper 26 can be found by swapping bits 0 and 1 of the address.

        variant   lines     registers                       Mapper Number
        =================================================================
        VRC6a:    A0, A1    $x000, $x001, $x002, $x003      024
        VRC6b:    A1, A0    $x000, $x002, $x001, $x003      026

### <span id="16k_PRG_Select_.28.248000-.248003.29"></span>16k PRG Select (\$8000-\$8003) {#16k_PRG_Select_.28.248000-.248003.29}

    7  bit  0
    ---------
    .... PPPP
         ||||
         ++++- Select 16 KB PRG ROM bank at $8000-$BFFF

### <span id="8k_PRG_Select_.28.24C000-.24C003.29"></span>8k PRG Select (\$C000-\$C003) {#8k_PRG_Select_.28.24C000-.24C003.29}

    7  bit  0
    ---------
    ...P PPPP
       | ||||
       +-++++- Select 8 KB PRG ROM bank at $C000-$DFFF

### <span id="PPU_Banking_Style_.28.24B003.29"></span>PPU Banking Style (\$B003) {#PPU_Banking_Style_.28.24B003.29}

    7  bit  0
    ---------
    W.PN MMDD
    | || ||||
    | || ||++- PPU banking mode; see below
    | || ++--- Mirroring varies by banking mode, see below
    | |+------ 1: Nametables come from CHRROM, 0: Nametables come from CIRAM
    | +------- CHR A10 is 1: subject to further rules 0: according to the latched value
    +--------- PRG RAM enable

The VRC6 supports the use of a larger RAM to provide more nametables. However, the three commercial VRC6 games neither provided extra nametable RAM, nor used ROM nametables. As a result these games only ever write the values \$20, \$24, \$28, \$2C, \$A0, \$A4, \$A8, and \$AC to this register.

CIRAM A10 is always connected to CHR A10, and bit 5 affects the behaviour of this signal (see below for details). The commercial games always left this bit set.

### <span id="CHR_Select_0.E2.80.A67_.28.24Dxxx.2C_.24Exxx.29"></span>CHR Select 0…7 (\$Dxxx, \$Exxx) {#CHR_Select_0.E2.80.A67_.28.24Dxxx.2C_.24Exxx.29}

For brevity, we refer to the registers at \$D000 through \$D003 and \$E000 through \$E003 as R0 through R7.

The lower 3 bits of the \$B003 register affect where the registers are used:

<table class="wikitable">
<thead>
<tr>
<th>[$B003] &amp; $03 →</th>
<th style="width: 2em">0</th>
<th style="width: 2em">1</th>
<th>2 or 3</th>
</tr>
</thead>
<tbody>
<tr>
<th>CHR bank</th>
<th colspan="3">Register used</th>
</tr>
&#10;<tr>
<td>$0000-$03FF</td>
<td>R0</td>
<td rowspan="2">R0</td>
<td>R0</td>
</tr>
<tr>
<td>$0400-$07FF</td>
<td>R1</td>
<td>R1</td>
</tr>
<tr>
<td>$0800-$0BFF</td>
<td>R2</td>
<td rowspan="2">R1</td>
<td>R2</td>
</tr>
<tr>
<td>$0C00-$0FFF</td>
<td>R3</td>
<td>R3</td>
</tr>
<tr>
<td>$1000-$13FF</td>
<td>R4</td>
<td rowspan="2">R2</td>
<td rowspan="2">R4</td>
</tr>
<tr>
<td>$1400-$17FF</td>
<td>R5</td>
</tr>
<tr>
<td>$1800-$1BFF</td>
<td>R6</td>
<td rowspan="2">R3</td>
<td rowspan="2">R5</td>
</tr>
<tr>
<td>$1C00-$1FFF</td>
<td>R7</td>
</tr>
</tbody>
</table>

<table class="wikitable">
<thead>
<tr>
<th>[$B003] &amp; $07 →</th>
<th>0, 6, or 7</th>
<th style="width: 4em">1 or 5</th>
<th>2, 3, or 4</th>
</tr>
</thead>
<tbody>
<tr>
<th>Nametable bank</th>
<th colspan="3">Register used</th>
</tr>
&#10;<tr>
<td>$2000-$23FF</td>
<td>R6</td>
<td>R4</td>
<td>R6</td>
</tr>
<tr>
<td>$2400-$27FF</td>
<td>R6</td>
<td>R5</td>
<td>R7</td>
</tr>
<tr>
<td>$2800-$2BFF</td>
<td>R7</td>
<td>R6</td>
<td>R6</td>
</tr>
<tr>
<td>$2C00-$2FFF</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
</tr>
</tbody>
</table>

When bit 5 of \$B003 is clear, CHR/CIRAM A10 will be controlled directly by the register LSB, causing 2 KiB banks to have duplicate 1 KiB halves. Existing Konami games did not use this configuration. This was intended for a different PCB (never used) where PPU A10 directly controls CHR A10 instead, permitting 512 KiB of CHR-ROM.

When bit 5 of \$B003 is set, 2 KiB pattern table banks pass PPU A10 through (ignoring the LSB of the register). Nametables apply different rules at the same time: see below.

#### Mirroring {#Mirroring}

With \$B003:5 set, there are four different modes corresponding to \$B003 & 3 that each have 4 nametable mirroring settings. Only mode 0 was used by Konami's commercial games.

##### Mode 0 {#Mode_0}

| \[\$B003\] & \$F | \$0      | \$4        | \$8        | \$C        |
|------------------|----------|------------|------------|------------|
| Mirroring        | vertical | horizontal | 1-screen A | 1-screen B |
| CIRAM A10        | PPU A10  | PPU A11    | Ground (0) | Vcc (1)    |

This mode was not intended for use with ROM nametables (\$B003:4 set), because it overrides the LSB of the nametable registers with the signal intended for CIRAM A10. Because R6 and R7 are already in use to control the pattern banks, this is not very suitable if combined with ROM nametables ([Mode 3](#Mode_3) is designed for that instead).

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
<th>[$B003] &amp; $F</th>
<th>$0</th>
<th>$4</th>
<th>$8</th>
<th>$C</th>
</tr>
</thead>
<tbody>
<tr>
<th>Mirroring</th>
<td>horizontal pairs<br />
2 KiB spread</td>
<td>vertical pairs<br />
2 KiB spread</td>
<td>horizontal mirroring<br />
1 KiB (even only)</td>
<td>vertical mirroring<br />
1 KiB (odd only)</td>
</tr>
<tr>
<th>$2000-$23FF</th>
<td>R6 even</td>
<td>R6 even</td>
<td>R6 even</td>
<td>R6 odd</td>
</tr>
<tr>
<th>$2400-$27FF</th>
<td>R6 odd</td>
<td>R7 even</td>
<td>R6 even</td>
<td>R7 odd</td>
</tr>
<tr>
<th>$2800-$2BFF</th>
<td>R7 even</td>
<td>R6 odd</td>
<td>R7 even</td>
<td>R6 odd</td>
</tr>
<tr>
<th>$2C00-$2FFF</th>
<td>R7 odd</td>
<td>R7 odd</td>
<td>R7 even</td>
<td>R7 odd</td>
</tr>
<tr>
<th>Mode 3 equivalent</th>
<td>$7</td>
<td>$3</td>
<td>$F</td>
<td>$B</td>
</tr>
</tbody>
</table>

##### Mode 1 {#Mode_1}

Mode 1 uses R4-R7 to directly control 4-screen mapping. This is very effective with ROM nametables (\$B003:4 set), but the LSB of each register still applies when using CIRAM.

<table class="wikitable">
<thead>
<tr>
<th>[$B003] &amp; $F</th>
<th>$1</th>
<th>$5</th>
<th>$9</th>
<th>$D</th>
</tr>
</thead>
<tbody>
<tr>
<th>Mirroring</th>
<td colspan="4">4-screen</td>
</tr>
<tr>
<th>$2000-$23FF</th>
<td colspan="4">R4</td>
</tr>
<tr>
<th>$2400-$27FF</th>
<td colspan="4">R5</td>
</tr>
<tr>
<th>$2800-$2BFF</th>
<td colspan="4">R6</td>
</tr>
<tr>
<th>$2C00-$2FFF</th>
<td colspan="4">R7</td>
</tr>
</tbody>
</table>

##### Mode 2 {#Mode_2}

Mode 2 uses R6-R7 to select two ROM pages for horizontal or vertical mirroring. If using CIRAM, the LSB of each register applies.

| \[\$B003\] & \$F | \$2      | \$6        | \$A      | \$E        |
|------------------|----------|------------|----------|------------|
| Mirroring        | vertical | horizontal | vertical | horizontal |
| \$2000-\$23FF    | R6       | R6         | R6       | R6         |
| \$2400-\$27FF    | R7       | R6         | R7       | R6         |
| \$2800-\$2BFF    | R6       | R7         | R6       | R7         |
| \$2C00-\$2FFF    | R7       | R7         | R7       | R7         |

##### Mode 3 {#Mode_3}

This mode is intended for ROM nametables (\$B003:4 set) but the nametable banking is extended to 2 KiB pages by forcing the LSB as a 0 (even) or 1 (odd) in different configurations.

This can be used to spread 2 adjacent ROM pages each across a pair of nametables.

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
<th>[$B003] &amp; $F</th>
<th>$3</th>
<th>$7</th>
<th>$B</th>
<th>$F</th>
</tr>
</thead>
<tbody>
<tr>
<th>Mirroring</th>
<td>vertical pairs<br />
2 KiB spread</td>
<td>horizontal pairs<br />
2 KiB spread</td>
<td>vertical mirroring<br />
1 KiB (odd only)</td>
<td>horizontal mirroring<br />
1 KiB (even only)</td>
</tr>
<tr>
<th>$2000-$23FF</th>
<td>R6 even</td>
<td>R6 even</td>
<td>R6 odd</td>
<td>R6 even</td>
</tr>
<tr>
<th>$2400-$27FF</th>
<td>R7 even</td>
<td>R6 odd</td>
<td>R7 odd</td>
<td>R6 even</td>
</tr>
<tr>
<th>$2800-$2BFF</th>
<td>R6 odd</td>
<td>R7 even</td>
<td>R6 odd</td>
<td>R7 even</td>
</tr>
<tr>
<th>$2C00-$2FFF</th>
<td>R7 odd</td>
<td>R7 odd</td>
<td>R7 odd</td>
<td>R7 even</td>
</tr>
<tr>
<th>Mode 0 equivalent</th>
<td>$4</td>
<td>$0</td>
<td>$C</td>
<td>$8</td>
</tr>
</tbody>
</table>

The nametable configurations are redundant with [Mode 0](#Mode_0), but in a different order; however, the reuse of R6 and R7 as pattern banks makes mode 0 ineffective for ROM nametables.

##### Other {#Other}

With \$B003:5 clear, CHR/CIRAM A10 always directly uses the LSB of the register used to map that range, rather than having it overridden for mirroring control. (See nametable bank table above.)

This setting was intended to be used with a different PCB that would have connected PPU A10 directly to CHR A10, enabling 512 KiB CHR-ROM (with mode 1 having 2 KiB banks). The actual behaviour here is a leftover consequence of being run on the wrong board:

<table class="wikitable">
<thead>
<tr>
<th>[$B003] &amp; $F</th>
<th>$0</th>
<th>$1</th>
<th>$2</th>
<th>$3</th>
<th>$4</th>
<th>$5</th>
<th>$6</th>
<th>$7</th>
<th>$8</th>
<th>$9</th>
<th>$A</th>
<th>$B</th>
<th>$C</th>
<th>$D</th>
<th>$E</th>
<th>$F</th>
</tr>
</thead>
<tbody>
<tr>
<th>Mirroring</th>
<td>H</td>
<td>4</td>
<td colspan="3">V</td>
<td>4</td>
<td colspan="3">H</td>
<td>4</td>
<td colspan="3">V</td>
<td>4</td>
<td colspan="2">H</td>
</tr>
<tr>
<th>$2000-$23FF</th>
<td>R6</td>
<td>R4</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R4</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R4</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R4</td>
<td>R6</td>
<td>R6</td>
</tr>
<tr>
<th>$2400-$27FF</th>
<td>R6</td>
<td>R5</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R5</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R5</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R5</td>
<td>R6</td>
<td>R6</td>
</tr>
<tr>
<th>$2800-$2BFF</th>
<td>R7</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R6</td>
<td>R7</td>
<td>R7</td>
</tr>
<tr>
<th>$2C00-$2FFF</th>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
<td>R7</td>
</tr>
</tbody>
</table>

Using ROM nametables on the hypothetical board mentioned above would transform all of the "H" layouts into "horizontal pairs 2 KiB spread", and all of the "V" and "4" layouts would have disjoint 256 KiB halves for the left two and right two nametables.

### <span id="IRQ_control_.28.24F00x.29"></span>IRQ control (\$F00x) {#IRQ_control_.28.24F00x.29}

For details on IRQ operation, see [VRC IRQs](../VRC_IRQ/). Many VRC mappers use the same IRQ system.

The VRC6 IRQ can be used to count either CPU cycles, or scanlines as a multiple of CPU cycles.

           7  bit  0
           ---------
    $F000: LLLL LLLL - IRQ Latch
    $F001: .... .MEA - IRQ Control
    $F002: .... .... - IRQ Acknowledge

- L - reload value for latch
- M - mode (1=cycle, 0=scanline)
- E - enable IRQ
- A - acknowledge bit

### <span id="Sound_.28.24900x.2C_.24A00x.2C_.24B000-.24B002.29"></span>Sound (\$900x, \$A00x, \$B000-\$B002) {#Sound_.28.24900x.2C_.24A00x.2C_.24B000-.24B002.29}

For details on sound information, see [VRC6 audio](../VRC6_audio/).

## Revisions {#Revisions}

Three revisions of the VRC6 chip are known to exist:

- 053328 - pin 1 is GND
- 053329 - pin 1 appears to be an input
- 053330 - based on the SLA7340 CMOS gate array, pin 1 is GND

## References {#References}

- Official VRC6 documentation: <http://www.assemblergames.com/forums/showthread.php?48390-Several-Famicom%28NES%29-misc-dev-documents-from-Nintendo-and-Konami>
- Naruko's explanation of the \$B003 register according to the VRC6 documentation: <http://forums.nesdev.org/viewtopic.php?f=11&t=10628>
- [VRCVI Chip Info](http://nesdev.org/vrcvi.txt) by Kevin Horton (audio only)
- [VRCVI](http://nesdev.org/vrc6-j.txt) by goroh (sound info is inaccurate)
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
- [VRC6 mirroring test ROM](https://forums.nesdev.org/viewtopic.php?t=11028)
