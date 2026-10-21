---
title: "CNROM"
url: "/wiki/CNROM/"
description: "CNROM is a discrete-logic circuit board providing up to four 8 KB banks of CHR-ROM. Two iNES mapper numbers denote its two configurations:"
wikiPage: "CNROM"
wikiCategories:
  - name: "CNROM-like mappers"
    url: "/category/CNROM-like_mappers/"
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>CNROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, others </td></tr><tr><td><b>Boards</b></td><td>CNROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>32K/128K (oversize Bandai PCB) </td></tr><tr><td><b>PRG ROM window</b></td><td>n/a </td></tr><tr><td><b>PRG RAM capacity</b></td><td>2 KiB (Namcot variant) </td></tr><tr><td><b>PRG RAM window</b></td><td>none </td></tr><tr><td><b>CHR capacity</b></td><td>32K </td></tr><tr><td><b>CHR window</b></td><td>8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">003</a>, <a href="./">185</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=3">iNES 003</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=185">iNES 185</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=CNROM">CNROM</a></td></tr></tbody></table></div>

**CNROM** is a discrete-logic circuit board providing up to four 8 KB banks of CHR-ROM. Two iNES mapper numbers denote its two configurations:

- **iNES Mapper 3** denotes the common usage mounting 16-32 KiB of CHR-ROM.
  - Submapper 0: Bus conflict behavior unknown
  - Submapper 1: No bus conflicts
  - Submapper 2: AND-type bus conflicts
- **iNES Mapper 185** denotes a special usage mounting only 8 KiB of CHR-ROM. CHR-ROM is disabled unless the correct bank number has been selected.
  - Submapper 0: CHR-ROM-enabling CS1/CS2 values unknown
  - Submapper 4: CS1/CS2=0 enables CHR-ROM, all other values disable CHR-ROM
  - Submapper 5: CS1/CS2=1 enables CHR-ROM, all other values disable CHR-ROM
  - Submapper 6: CS1/CS2=2 enables CHR-ROM, all other values disable CHR-ROM
  - Submapper 7: CS1/CS2=3 enables CHR-ROM, all other values disable CHR-ROM

Two Bandai boards extend the functionality of CNROM:

- [*Family Trainer: Jogging Race*](https://nescartdb.com/profile/view/4090/family-trainer-4-jogging-race) is a simple [oversized](#Oversize_mapper_3_with_up_to_128_KiB) version of CNROM and therefore is assigned to mapper 3 as well;
- [*Family Trainer: Aerobics Studio*](https://nescartdb.com/profile/view/3953/family-trainer-3-aerobics-studio) adds a [Mitsubishi M50805 speech chip](https://forums.nesdev.org/viewtopic.php?p=102300#p102300) using an [extra register](#Speech_Start.2FMessage_Select:_responds_to_writes_to_CPU_.246000-.247FFF). The chip is currently unemulated, but once it is, the speech ROM data will be included as NES 2.0 Misc. ROM data.

The Namco game [*Hayauchi Super Igo*](https://nescartdb.com/profile/view/3844/hayauchi-super-igo) adds 2 KiB of PRG-RAM, denoted using mapper 3 and the appropriate value in the header's PRG-RAM size field.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 2 KiB of PRG-RAM, mirrored three times (*Hayauchi Super Igo* only)
- CPU \$8000-\$FFFF: 32 KiB unbanked PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB switchable window into 32 KiB CHR-ROM
- Nametable arrangement: Fixed; solder pad selects between Horizontal and Vertical

# Registers {#Registers}

## <span id="Bank_Select:_responds_to_writes_to_CPU_.248000-.24FFFF"></span>Bank Select: responds to writes to CPU \$8000-\$FFFF {#Bank_Select:_responds_to_writes_to_CPU_.248000-.24FFFF}

### Regular mapper 3 with up to 32 KiB {#Regular_mapper_3_with_up_to_32_KiB}

    D~[..DC ..BA] A~[1... .... .... ....]
         ||   ||
         ||   ++- CHR A14..A13 (8 KiB bank)
         |+------ Output to Diode 2 (D2)
         +------- Output to Diode 1 (D1)

The original CNROM board is always subject to AND-type bus conflicts: the effective value is the value being written bitwise-AND'd with the PRG-ROM content at the address being written to. iNES Mapper 3's submapper indicates whether bus conflicts should actually be emulated. So far, the only .NES files requiring the absence of bus conflicts have been mapper hacks to mapper 3 from other mappers.

### Oversize mapper 3 with up to 128 KiB {#Oversize_mapper_3_with_up_to_128_KiB}

    D~[.... DCBA] A~[1... .... .... ....]
            ||||
            ++++- CHR A16..A13 (8 KiB bank)

Among the licensed NES/Famicom library, this configuration is only used by Bandai's *Family Trainer: Jogging Race*. Several unlicensed cartridges mount 64 KiB of CHR-ROM; although they work as oversize mapper 3 as well, they are canonically assigned to [INES Mapper 148](../INES_Mapper_148/) instead.

### Mapper 185 {#Mapper_185}

    D~[..DC ..BA] A~[1... .... .... ....]
         ||   ||
         ||   |+- Chip Select 2 (CS2)
         ||   +-- Chip Select 1 (CS1)      
         |+------ Output to Diode 2 (D2)
         +------- Output to Diode 1 (D1)

Mapper 185 always has AND-type bus conflicts. The submapper number denotes the correct Chip Select value that enables CHR-ROM; if another value is active, CHR-ROM is disabled, meaning that reading from the pattern tables returns [open bus](../Open_bus_behavior/#PPU_open_bus). Theoretically, this should return the LSB of the address read, but real-world behavior varies, and the earlier revision of Mighty Bomb Jack in fact relies on open bus at PPU address \$0000 being something other than \$00 (by means of a 10k pull-up resistor on the CHR ROM's D0 pin). If the correct Chip Select value is not known -- denoted by submapper 0, which applies to all .NES files without a [NES 2.0](../NES_2.0/) header, the simple heuristic "Disable CHR-ROM for the first two reads from \$2007 after a reset and then enable it" will work with all known games.

<table class="datatable sortable">
<thead>
<tr>
<th></th>
<th>bank</th>
<th>PPU addr</th>
<th>test</th>
<th>bank</th>
<th>PPU addr</th>
<th>test</th>
</tr>
</thead>
<tbody>
<tr>
<th>Game</th>
<th colspan="3">Incorrect bank</th>
<th colspan="3">Must work</th>
</tr>
&#10;<tr>
<td><a href="https://nescartdb.com/profile/view/1262">Bird Week</a></td>
<td>$F0</td>
<td>$1FF0</td>
<td>≠ $0C</td>
<td>$0F</td>
<td>$1FF0</td>
<td>= $0C</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/1264">B-Wings</a></td>
<td>$00</td>
<td>$0000</td>
<td>≠ $3C</td>
<td>$33</td>
<td>$0000</td>
<td>= $3C</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/1261">Mighty Bomb Jack (J, PRG0)</a></td>
<td>$00</td>
<td>$0000</td>
<td>≠ $00</td>
<td>$11</td>
<td colspan="2">untested</td>
</tr>
<tr>
<td>Mighty Bomb Jack (J, PRG1)</td>
<td>$00</td>
<td>$0001</td>
<td>≠ $3C</td>
<td>$11</td>
<td colspan="2">untested</td>
</tr>
<tr>
<td>Sansuu 1 Nen</td>
<td>$20</td>
<td>$000C</td>
<td>≠ $BC</td>
<td>$22</td>
<td>$000C</td>
<td>= $BC</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/1263">Sansuu 2 Nen</a></td>
<td>$20</td>
<td>$0003</td>
<td>≠ $42</td>
<td>$22</td>
<td>$0003</td>
<td>= $42</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/4061">Othello (J)</a></td>
<td>$20</td>
<td>$0006</td>
<td>≠ $3F</td>
<td>$22</td>
<td>$0006</td>
<td>= $3F</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/3791">Sansuu 3 Nen</a></td>
<td>$00</td>
<td>$0006</td>
<td>≠ $34</td>
<td>$2A</td>
<td>$0006</td>
<td>= $34</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/3592">Spy vs Spy (J)</a></td>
<td>$13</td>
<td>$1F20</td>
<td>≠ $55</td>
<td>$21</td>
<td colspan="2">untested</td>
</tr>
<tr>
<td><a href="https://nescartdb.com/profile/view/2332">Seicross (J)</a></td>
<td>$21</td>
<td>$0700…$0707</td>
<td>≠ $20,$60,$70,$70,$70,$40,$08,$38</td>
<td>$20</td>
<td colspan="2">untested</td>
</tr>
</tbody>
</table>

- The [North American](https://nescartdb.com/profile/view/202/mighty-bomb-jack) and [PAL](https://nescartdb.com/profile/view/2143/mighty-bomb-jack) versions of *Mighty Bomb Jack* mount 32 KiB of CHR-ROM instead and so use normal [mapper 3](./).
- The North American and PAL versions of *Spy vs Spy* use [NROM](../NROM/) instead, as does the North American versions of *Othello* and *Seicross*.

## <span id="Speech_Start.2FMessage_Select:_responds_to_writes_to_CPU_.246000-.247FFF"></span>Speech Start/Message Select: responds to writes to CPU \$6000-\$7FFF {#Speech_Start.2FMessage_Select:_responds_to_writes_to_CPU_.246000-.247FFF}

    D~[.D.. .CBA] A~[011. .... .... ....]
        |    +++- Message number (0-7)
        +-------- M5085 /SYNC signal

A falling edge of the D bit starts the selected speech line. This register only exists on the PCB of *Family Trainer: Aerobics Studio*.

# Hardware {#Hardware}

The 4-bit latch register is implemented using a [74HC161](../74161/) binary counter. On iNES Mapper 3, the latch's A and B outputs directly connect to the CHR-ROM chip's A14 and A13 inputs. On iNES Mapper 185, the particular ROM chip model (Sharp LH2367 or equivalent) changes the functionality of pins 26 (normally A13) and 27 (normally A14) into two programmable (during manufacturing) Chip Select inputs, effectively disabling CHR-ROM and tri-stating its data outputs when the one correct bank of four possible CHR-ROM banks is not selected, as a copy-protection mechanism.

## Security diodes {#Security_diodes}

As another copy protection mechanism, the CNROM circuit board has a spot for two diodes that produce additional bus conflicts to hinder cartridge dumping attempts. Diode 1 connects latch bit 'D' to CHR-ROM A10; Diode 2 connects latch bit 'C' to CHR-ROM A12. Each latch bit must be set such that the diode does *not* allow current to flow, because if it does, AND-type bus conflicts occur that may cause the wrong CHR-ROM A10 and A12 signals to be applied, depending on the relative output resistances of latch chip and console or dumping device. Modern Kazzo-like dumping devices have strong enough output drivers to always win these bus conflicts, but 1980s' dumping equipment as well as the NES PPU will lose these bus conflicts and produce an unusable readout. Each diode can be mounted either with the anode or the cathode facing the latch output. As a diode will allow current to flow if the anode-side voltage is significantly higher than the cathode-side voltage, the latch bit must be 0 if it faces the anode, and 1 if it faces the cathode. The security diodes were mounted by most Nintendo-manufactured Japanese CNROM games manufactured in 1986 as well as Bandai-manufactured CNROM games from 1986 and 1987. They were never used on North American or PAL CNROM games, even as the NES-CNROM board has an unpopulated spot for them until at least revision -02. Emulators need not emulate the functionality of security diodes to run games correctly.

## Solder Pad Configuration {#Solder_Pad_Configuration}

- Horizontal nametable arrangement ("vertical mirroring"): 'H' connected, 'V' disconnected.
- Vertical nametable arrangement ("horizontal mirroring"): 'H' disconnected, 'V' connected.
- 16 KiB PRG-ROM: 'SL' connected, 'CL' disconnected.
- 32 KiB PRG-ROM: 'SL' disconnected, 'CL' connected.

[*Hayauchi Super Igo*](https://nescartdb.com/profile/view/3844/hayauchi-super-igo) is a CNROM-like board with a 2 KiB SRAM mapped at \$6000, using a [74HC10](../7410/) as the address decoder.

Theoretically the bank select register could be implemented with a [74HC377](../74377/) octal D latch, allowing up to 2 MiB of CHR ROM.
