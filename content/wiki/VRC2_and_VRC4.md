---
title: "VRC2 and VRC4"
url: "/wiki/VRC2_and_VRC4/"
description: "The Konami VRC2 and Konami VRC4 are two related ASIC mappers in the VRC[1] family."
wikiPage: "VRC2_and_VRC4"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>VRC2, VRC4</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami </td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>VRC2a-c <p>VRC4a-f </p></td></tr><tr><td><b>Pinout</b></td><td><a href="../VRC2_pinout/">VRC2 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K + 8K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K (VRC2) <p>512K (VRC4) </p></td></tr><tr><td><b>CHR window</b></td><td>1K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, switchable (VRC2) <p>H, V, or 1, switchable (VRC4) </p></td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Only VRC4 </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">021</a>, <a href="./">022</a>, <a href="./">023</a>, <a href="./">025</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=KONAMI-VRC-2">VRC2</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=KONAMI-VRC-4">VRC4</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=21">iNES 021</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=22">iNES 022</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=23">iNES 023</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=25">iNES 025</a></td></tr></tbody></table></div>

The **Konami VRC2** and **Konami VRC4** are two related [ASIC](../../category/ASIC_mappers/) [mappers](../Mapper/) in the VRC<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> family.

Because the VRC2 is mostly a subset of the VRC4, relevant games are often emulated as VRC4 only. iNES mappers **21**, **22**, **23** and **25** implement various board permutations, and [NES 2.0 submappers](../NES_2.0_submappers/) may be used to disambiguate further.

Example games:

| Game                               | Variant | iNES mapper |
|------------------------------------|---------|-------------|
| *Ganbare Pennant Race*             | VRC2a   | 22          |
| *TwinBee 3*                        | VRC2a   | 22          |
| *Wai Wai World*                    | VRC2b   | 23          |
| *Ganbare Goemon Gaiden*            | VRC2c   | 25          |
| *Wai Wai World 2*                  | VRC4a   | 21          |
| *Bio Miracle Bokutte Upa*          | VRC4b   | 25          |
| *Gradius 2 (J)*                    | VRC4b   | 25          |
| *Ganbare Goemon Gaiden 2*          | VRC4c   | 21          |
| *Teenage Mutant Ninja Turtles (J)* | VRC4d   | 25          |
| *Boku Dracula-kun*                 | VRC4e   | 23          |
| *Tiny Toon Adventures (J)*         | VRC4e   | 23          |
| *World Hero (Unl)*                 | VRC4f   | 23          |

## Variations {#Variations}

Variations VRC2a/b/c and VRC4a/b/c/d/e/f are referring to different *PCB layouts*. Considering the mapper chip itself, all VRC2 behave the same way and all VRC4 behave the same way. This section describes how each variation presents itself to the NES CPU considering the different *PCBs*.

The VRC2 and VRC4 were used with several similar but incompatible boards. The primary difference between them was having the mapper address lines connected in different ways. In particular, two lines chosen from A0-A7 will be used to select registers.

<table class="wikitable sortable">
<thead>
<tr>
<th colspan="7">VRC2</th>
</tr>
</thead>
<tbody>
<tr>
<th>Nickname</th>
<th>PCB</th>
<th>A0</th>
<th>A1</th>
<th>Registers</th>
<th>iNES mapper</th>
<th>submapper</th>
</tr>
&#10;<tr>
<td>VRC2a</td>
<td>351618</td>
<td>A1</td>
<td>A0</td>
<td>$x000, $x002, $x001, $x003</td>
<td>22</td>
<td>0</td>
</tr>
<tr>
<td>VRC2b</td>
<td>many<a href="#VRC2bPCB">†</a></td>
<td>A0</td>
<td>A1</td>
<td>$x000, $x001, $x002, $x003</td>
<td>23</td>
<td>3</td>
</tr>
<tr>
<td>VRC2c</td>
<td>351948</td>
<td>A1</td>
<td>A0</td>
<td>$x000, $x002, $x001, $x003</td>
<td>25</td>
<td>3</td>
</tr>
</tbody>
</table>

<table class="wikitable sortable">
<thead>
<tr>
<th colspan="7">VRC4</th>
</tr>
</thead>
<tbody>
<tr>
<th>Nickname</th>
<th>PCB</th>
<th>A0</th>
<th>A1</th>
<th>Registers</th>
<th>iNES mapper</th>
<th>submapper</th>
</tr>
&#10;<tr>
<td>VRC4a</td>
<td>352398</td>
<td>A1</td>
<td>A2</td>
<td>$x000, $x002, $x004, $x006</td>
<td>21</td>
<td>1</td>
</tr>
<tr>
<td>VRC4b</td>
<td>351406</td>
<td>A1</td>
<td>A0</td>
<td>$x000, $x002, $x001, $x003</td>
<td>25</td>
<td>1</td>
</tr>
<tr>
<td>VRC4c</td>
<td>352889</td>
<td>A6</td>
<td>A7</td>
<td>$x000, $x040, $x080, $x0C0</td>
<td>21</td>
<td>2</td>
</tr>
<tr>
<td>VRC4d</td>
<td>352400</td>
<td>A3</td>
<td>A2</td>
<td>$x000, $x008, $x004, $x00C</td>
<td>25</td>
<td>2</td>
</tr>
<tr>
<td>VRC4e</td>
<td>352396</td>
<td>A2</td>
<td>A3</td>
<td>$x000, $x004, $x008, $x00C</td>
<td>23</td>
<td>2</td>
</tr>
<tr>
<td>VRC4f</td>
<td>-<a href="#VRC4fPCB">*</a></td>
<td>A0</td>
<td>A1</td>
<td>$x000, $x001, $x002, $x003</td>
<td>23</td>
<td>1</td>
</tr>
</tbody>
</table>

† <span id="VRC2bPCB"></span> The VRC2b PCBs include: 350603, 350636, 350926, 351179

\* <span id="VRC4fPCB"></span> There're no known official VRC4f boards being commercially produced, and this variant was only glanced at several pirate multicarts and single-game cartridges. As such, VRC4f should be considered as an unofficial variant.

### VRC2 vs VRC4 {#VRC2_vs_VRC4}

The VRC2 has:

- CHR limit of 256 KiB
- 1 bit latch at \$6000-\$6FFF
- Only horizontal or vertical nametable layout.

The VRC4 has:

- CHR limit of 512 KiB
- Additional 8 KiB PRG bank at \$C000
- Horizontal, vertical and 1-screen nametables.
- IRQ device

The VRC2a (mapper 22) additionally wires the CHR banking lines differently (see below).

Unlike VRC2's, the RAM decoding circuit that's part of the VRC4 itself only decodes RAM from \$6000-\$6FFF. For the [one game](https://nescartdb.com/profile/view/1573/ganbare-goemon-gaiden-2-tenka-no-zaihou) with 8 KiB of RAM, an external [circuit](../PRG_RAM_circuit/) was added.

### iNES mappers {#iNES_mappers}

iNES mappers 21, 23 and 25 each implements *two* address mappings at the same time:

- **21:** VRC4a, VRC4c
- **23:** VRC2b + VRC4f, VRC4e
- **25:** VRC2c + VRC4b, VRC4d

Because the address pairings do not overlap, and the games appear to use these registers in a well behaved manner, it is presumed sufficient for compatibility in most cases. Additionally, the VRC4 is always presumed for these three mappers.

With an [NES 2.0](../NES_2.0/) header, a [submapper](../NES_2.0_submappers/) may be used to specify which address mapping to use, and either VRC2 or VRC4.

## Banks {#Banks}

### Common {#Common}

- PPU \$0000-\$03FF: 1 KiB switchable CHR bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank

### VRC2 {#VRC2}

- CPU \$6000-\$6FFF: 1 bit latch, or
- CPU \$6000-\$7FFF: 8 KiB PRG RAM bank (optional)
- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last 16 KiB

### VRC4 {#VRC4}

- CPU \$6000-\$6FFF: 2 KiB PRG RAM bank (mirrored once) (optional), or
- CPU \$6000-\$7FFF: 8 KiB PRG RAM bank (via [extra circuit](https://nescartdb.com/profile/view/1573/ganbare-goemon-gaiden-2-tenka-no-zaihou)) (optional)
- CPU \$8000-\$9FFF (or \$C000-\$DFFF): 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF (or \$8000-\$9FFF): 8 KiB PRG ROM bank, fixed to the second-to-last bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank

## Registers {#Registers}

This page lists registers as they are in the VRC2b and VRC4f variants (iNES mapper 23).

### <span id="PRG_Swap_Mode.2FWRAM_control_.28.249002.29_VRC4"></span>PRG Swap Mode/WRAM control (\$9002) VRC4 {#PRG_Swap_Mode.2FWRAM_control_.28.249002.29_VRC4}

    7  bit  0
    ---------
    .... ..MW
           |+- WRAM Control
           +-- Swap Mode

This register is **VRC4 only**.

When 'W' is clear:

- the 8 KiB page at \$6000 is open bus, and WRAM content cannot be read nor written

When 'W' is set:

- the 8 KiB page at \$6000 is WRAM, and WRAM content can be read and written

When 'M' is clear:

- the 8 KiB page at \$8000 is controlled by the \$800x register
- the 8 KiB page at \$C000 is fixed to the second-to-last 8 KiB in the ROM

When 'M' is set:

- the 8 KiB page at \$8000 is fixed to the second-to-last 8 KiB in the ROM
- the 8 KiB page at \$C000 is controlled by the \$800x register

### <span id="PRG_Select_0_.28.248000.2C_.248001.2C_.248002.2C_.248003.29"></span>PRG Select 0 (\$8000, \$8001, \$8002, \$8003) {#PRG_Select_0_.28.248000.2C_.248001.2C_.248002.2C_.248003.29}

    7  bit  0
    ---------
    ...P PPPP
       | ||||
       +-++++- Select 8 KiB PRG bank at $8000 or $C000 depending on Swap Mode

VRC2 does not have a Swap Mode and the bank this register will switch is always the \$8000-\$9FFF bank.

### <span id="PRG_Select_1_.28.24A000.2C_.24A001.2C_.24A002.2C_.24A003.29"></span>PRG Select 1 (\$A000, \$A001, \$A002, \$A003) {#PRG_Select_1_.28.24A000.2C_.24A001.2C_.24A002.2C_.24A003.29}

    7  bit  0
    ---------
    ...P PPPP
       | ||||
       +-++++- Select 8 KiB PRG bank at $A000

### <span id="Nametable_arrangement_.28.249000.2C_.249001.2C_.249002.2C_.249003.29"></span>Nametable arrangement (\$9000, \$9001, \$9002, \$9003) {#Nametable_arrangement_.28.249000.2C_.249001.2C_.249002.2C_.249003.29}

    7  bit  0
    ---------
    .... ..MM
           ||
           ++- Mirroring (0: vertical; 1: horizontal; 2: one-screen, lower bank; 3: one-screen, upper bank)

VRC2 supports **only vertical or horizontal** nametables. Bit 1 is ignored.

VRC4 only has this register at **\$9000 only**. \$9002 is used to select PRG swap mode instead (see above).

VRC2-using games are usually well-behaved and only write 0 or 1 to this register, but [Konami Wai Wai World](https://nescartdb.com/profile/view/2274/konami-wai-wai-world) in one instance writes \$FF instead<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>.

### <span id="VRC4_External_Select_.28.249003.29"></span>VRC4 External Select (\$9003) {#VRC4_External_Select_.28.249003.29}

The VRC4 decodes writes to \$9003 and emits an active low signal. No Konami games made use of this, but some unlicensed games, assigned to other mappers, added extra hardware behind this address.

### <span id="CHR_Select_0_low.28.24B000.29.2C_high.28.24B001.29"></span>CHR Select 0 low(\$B000), high(\$B001) {#CHR_Select_0_low.28.24B000.29.2C_high.28.24B001.29}

      $B000        $B001
    7  bit  0    7  bit  0
    ---------    ---------
    .... LLLL    ...H HHHH
         ||||       | ||||
         ||||       +-++++- High 5 bits of 1 KiB CHR bank at PPU $0000
         ++++-------------- Low 4 bits

VRC2 only has 4 high bits of CHR select. \$B001 bit 4 is ignored.

On VRC2a (mapper 22), the low bit is ignored (right shift value by 1).

### <span id="CHR_Select_1_low.28.24B002.29.2C_high.28.24B003.29"></span>CHR Select 1 low(\$B002), high(\$B003) {#CHR_Select_1_low.28.24B002.29.2C_high.28.24B003.29}

      $B002        $B003
    7  bit  0    7  bit  0
    ---------    ---------
    .... LLLL    ...H HHHH
         ||||       | ||||
         ||||       +-++++- High 5 bits of 1 KiB CHR bank at PPU $0400
         ++++-------------- Low 4 bits

VRC2 only has 4 high bits of CHR select. \$B003 bit 4 is ignored.

On VRC2a (mapper 22), the low bit is ignored (right shift value by 1).

### <span id="CHR_Select_2.E2.80.A67_.28.24C000-.24EFFF.29"></span>CHR Select 2…7 (\$C000-\$EFFF) {#CHR_Select_2.E2.80.A67_.28.24C000-.24EFFF.29}

The other six CHR bank selects continue the pattern:

<table class="wikitable">
<thead>
<tr>
<th colspan="2">Write to CPU address</th>
<th>1KB CHR bank affected</th>
</tr>
</thead>
<tbody>
<tr>
<th>(low 4 bits)</th>
<th>(high 5 bits)</th>
<th></th>
</tr>
&#10;<tr>
<td>$C000</td>
<td>$C001</td>
<td>$0800-$0BFF</td>
</tr>
<tr>
<td>$C002</td>
<td>$C003</td>
<td>$0C00-$0FFF</td>
</tr>
<tr>
<td>$D000</td>
<td>$D001</td>
<td>$1000-$13FF</td>
</tr>
<tr>
<td>$D002</td>
<td>$D003</td>
<td>$1400-$17FF</td>
</tr>
<tr>
<td>$E000</td>
<td>$E001</td>
<td>$1800-$1BFF</td>
</tr>
<tr>
<td>$E002</td>
<td>$E003</td>
<td>$1C00-$1FFF</td>
</tr>
</tbody>
</table>

### <span id="IRQ_Control_.28.24F00x.29_VRC4"></span>IRQ Control (\$F00x) VRC4 {#IRQ_Control_.28.24F00x.29_VRC4}

    $F000:  IRQ Latch, low 4 bits
    $F001:  IRQ Latch, high 4 bits
    $F002:  IRQ Control
    $F003:  IRQ Acknowledge

Many VRC mappers use the same IRQ system. For details on IRQ operation, see [VRC IRQs](../VRC_IRQ/).

### <span id="Microwire_interface_.28.246000-.246FFF.29_VRC2"></span>Microwire interface (\$6000-\$6FFF) VRC2 {#Microwire_interface_.28.246000-.246FFF.29_VRC2}

This only exists on VRC2.

How it was supposed to work:

    7  bit  0
    ---------
    .... .SCD
          |||
          ||+- Data to EEPROM (write) or from EEPROM (read)
          |+-- Clock to EEPROM (write only)
          +--- Chip Select (write only)

How it works in practice:

    7  bit  0
    ---------
    .... ...L
            |
            +- 1-bit latch value (r/w)

Reads from \$6000-6FFF return open bus for the top 7 bits. Reads from \$7000-7FFF only ever return open bus.

The VRC2 was supposed to have shipped with a [Microwire](https://en.wikipedia.org/wiki/Microwire) interface for save games. However, Konami never used it, seemingly due to a defect in the VRC2.

Across all the VRC2-using boards, the *Data from EEPROM* pin has been connected to many different things. On **351618** ([22](./)) it's connected to ground. On **350603**, **350636**, and **351179** it floats. On **350926** it's connected to *Data to EEPROM*. On **351948** ([25](./)) it's connected to ground but extra hardware keeps the VRC2 from seeing reads from \$6000 so as to not conflict with PRG RAM.

Several games, including Contra (J) and Ganbare Goemon 2 (J), rely on the two *Data* pins being connected to each other, and so expect to be able to read the written value back. In these cases, the LSB agrees with the last value written and the upper seven bits are open bus, e.g. both LDA \$6100 and LDA \$6000 will return \$60\|latch. Returning neither open bus nor 0x00 will work, and these games will lock almost immediately after execution begins. Fortunately, no games ever rely on any other behavior.

Emulators that use the same VRC4 core (and its PRG RAM) for VRC2 emulation will have the effect simulated for them. However, only **351948** contains any RAM.

## Hardware {#Hardware}

The VRC2 and VRC4 exist in a [40-pin PDIP package](../VRC2_pinout/).

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [Forum thread](http://forums.nesdev.org/viewtopic.php?f=2&t=10611): Several Famicom Nes Documents in Japanese indicate that VRC stands for "Virtual ROM Controller"
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Forum thread](http://forums.nesdev.org/viewtopic.php?f=3&t=13473): incorrect mirroring in Wai Wai World, and VRC2 testing.

- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Konami Mapper](http://nesdev.org/konami-e.txt) by goroh, translated by Sgt. Bowhack
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
- The CHR registers really are 9 bits: <http://forums.nesdev.org/viewtopic.php?t=8569>
