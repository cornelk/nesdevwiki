---
title: "NES 2.0 submappers"
url: "/wiki/NES_2.0_submappers/"
description: "Submapper is a term used in the NES 2.0 header for 4-bit codes designating functionally distinct variants of iNES mappers that cannot be distinguished by the memory size fields alone. Most emulators…"
wikiPage: "NES_2_0_submappers"
---

**Submapper** is a term used in the [NES 2.0](../NES_2.0/) header for 4-bit codes designating functionally distinct variants of iNES mappers that cannot be distinguished by the memory size fields alone. Most emulators using [iNES](../INES/) format distinguish these using CRC, SHA-1, or other hashes of the PRG ROM and CHR ROM, but this works only for games published prior to 1997, not for fan translations or ROM hacks, and not for new games on the same mapper.

Submapper 0 represents the default iNES behavour, so that backward compatibility is maintained with existing ROMs.

Submapper allocations that are listed as "deprecated" were assigned by kevtris' original proposal, but have no known use cases. The deprecation reserves these unused allocations to maintain continuity and compatibility.

**This document is a living specification. Proposals for new submappers should be made at:** [Proposals](Proposals/).

## [001](../MMC1/): [MMC1](../MMC1/) {#001:_MMC1}

Most [MMC1 boards](../MMC1/) are compatible with the standard mapper 1 behavior (submapper [0](#001:_0)).

Boards with CHR-RAM usually reuse the CHR banking lines to address other things.

- SUROM, SOROM, and SXROM implement extra PRG-ROM and PRG-RAM banking (*deprecated* submappers [1, 2, 4](#001:_1.2C_2.2C_4_.3D_SUROM.2C_SOROM.2C_SXROM)).
- SNROM implements a redundant PRG-RAM enable (no assigned submapper).

The variant MMC1A chip was assigned to [mapper 155](../MMC1/) (*deprecated* submapper [3](#001:_3)).

Most boards with 32k of PRG-ROM have no PRG banking: SEROM, SHROM, SH1ROM (submapper [5](#001:_5_Fixed_PRG)).

### 001: 0 {#001:_0}

Normal behavior.

### <span id="001:_1.2C_2.2C_4_.3D_SUROM.2C_SOROM.2C_SXROM"></span>001: 1, 2, 4 = SUROM, SOROM, SXROM {#001:_1.2C_2.2C_4_.3D_SUROM.2C_SOROM.2C_SXROM}

**Deprecated.**

These submappers are a redundancy. The difference between these boards and the "normal" MMC1 implementation is exclusively dependent on the sizes of CHR, PRG RAM, and PRG ROM. Because of this, the addition of PRG RAM to the NES 2.0 specification was enough to make these compatible with submapper 0.

These three boards used 8k CHR RAM, and reused the CHR banking bits to bank PRG ROM and RAM instead. The specific boards can be detected by these sizes, or simply emulated together in submapper 0:

    $A000 and $C000:
    4bit0
    -----
    EDCBA
    |||||
    ||||+- CHR A12
    |||+-- CHR A13 if CHR >= 16k
    ||+--- CHR A14 if CHR >= 32k; and PRG RAM A13 if PRG RAM = 32k (submapper 4)
    |+---- CHR A15 if CHR >= 64k; and PRG RAM A13 if PRG RAM = 16k (submapper 2)
    |                              or PRG RAM A14 if PRG RAM = 32k (submapper 4)
    +----- CHR A16 if CHR = 128k; and PRG ROM A18 if PRG ROM = 512k (submappers 1, 4)

The following games are on SUROM (submapper 1):

- *[Dragon Warrior III](https://nescartdb.com/profile/view/1273)*
- *[Dragon Warrior IV](https://nescartdb.com/profile/view/1276)* / *[Dragon Quest IV](https://nescartdb.com/profile/view/1526)*
- *[Ninjara Hoi!](https://nescartdb.com/profile/view/1531)*

SOROM (submapper 2):

- *[Genghis Khan](https://nescartdb.com/profile/view/919)* / *[Aoki Ookami to Shiroki Mejika: Genghis Khan](https://nescartdb.com/profile/view/1528)*
- *[Nobunaga's Ambition](https://nescartdb.com/profile/view/497)* / *[Nobunaga no Yabou: Zenkokuban](https://nescartdb.com/profile/view/1529)*
- *[Romance of the Three Kingdoms](https://nescartdb.com/profile/view/730)* / *[Sangokushi](https://nescartdb.com/profile/view/3170)*

SXROM (submapper 4):

- *[Final Fantasy I & II](https://nescartdb.com/profile/view/3642)*
- *[The Best Play Pro Yakyuu Special](https://nescartdb.com/profile/view/3757)*
- *[PR8](http://forums.nesdev.org/viewtopic.php?f=22&t=7231)*. If extra RAM is missing, it gives a ["DOES NOT WORK IN THIS EMULATOR" error](http://forums.nesdev.org/viewtopic.php?f=2&t=7629).

If any NES 2.0 ROMs are found using these deprecated submappers, the CHR, PRG RAM, and PRG ROM sizes must appropriately match to be a valid header.

### 001: 3 {#001:_3}

**Deprecated.**

This originally described a submapper that was already implemented as [iNES Mapper 155](../MMC1/).

### 001: 5 Fixed PRG {#001:_5_Fixed_PRG}

[SEROM](../MMC1/), [SHROM](../MMC1/), ~~SH1ROM~~ use a fixed 32k PRG ROM with no banking support. (This is distinct from ~~SIROM~~ which has 32k of bankable PRG ROM.)

PRG ROM A14 is connected directly to CPU A14 (and MMC1 A14 input) instead of MMC1 A14 output.

Existing games are compatible with submapper 0 if \$8000-BFFF is initialized to the low bank, and \$C000-FFFF is initialized to the high bank. These boards were used in several games: [SEROM](https://nescartdb.com/search/basic?keywords=serom&kwtype=pcb) [SHROM](https://nescartdb.com/search/basic?keywords=shrom&kwtype=pcb) [SH1ROM](https://nescartdb.com/search/basic?keywords=sh1rom&kwtype=pcb)

Test ROM:

- *[serom.zip](http://forums.nesdev.org/download/file.php?id=3753)* - validates non-existence of PRG banking control. ([forum post](http://forums.nesdev.org/viewtopic.php?p=153298#p153298))

### <span id="001:_6_.3D_2ME"></span>001: 6 = 2ME {#001:_6_.3D_2ME}

The 2ME board is a Famicom Network System MMC1B variant used by *JRA-PAT* revisions FCN027-04 and later. It supports 64 big-endian 16-bit words of EEPROM and up to 32 KiB of battery-backed PRG-RAM banked like SXROM. Because the Network System does not expose the PPU bus to the software card, all CHR-related outputs are repurposed to support this storage. See [MMC1 2ME](../MMC1/#2ME) for full details.

## <span id="002.2C_003.2C_007.2C_011:_UxROM.2C_CNROM.2C_AxROM.2C_Color_Dreams"></span>[002](../UxROM/), [003](../CNROM/), [007](../AxROM/), [011](../Color_Dreams/): [UxROM](../UxROM/), [CNROM](../CNROM/), [AxROM](../AxROM/), [Color Dreams](../Color_Dreams/) {#002.2C_003.2C_007.2C_011:_UxROM.2C_CNROM.2C_AxROM.2C_Color_Dreams}

Mappers 2, 3, 7, and 11 describe [discrete logic mappers](../../category/Discrete_logic_mappers/) that are usually subject to [bus conflicts](../Bus_conflict/). Most of these games are programmed in a way that does not rely on the bus conflict behaviour, but software bugs may expose the difference (e.g. *Cybernoid*<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>). Because these differences are obscure, the default iNES implementation for mappers 2 and 3 has been inconsistent across emulators, some with bus conflicts, some without. Mappers 7 and 11 are normally emulated without bus conflicts, the former because of a small number of games using the [ANROM](../AxROM/) board which rely on this, and the latter because of at least one prototype board.

Two submappers allow this ambiguity to be resolved, by selecting a specific behaviour:

0: Default iNES behaviour (Emulators should warn the user and/or enforce bus conflicts for mappers 2 and 3, and should warn the user and/or fail to enforce bus conflicts for mapper 7)  
1: Bus conflicts do not occur  
2: Bus conflicts occur, producing the bitwise AND of the written value and the value in ROM

Although all Nintendo-manufactured games using normal CNROM (mapper 3), normal UxROM (mapper 2), and inverted UxROM (mapper 180) had bus conflicts, apparently several unlicensed games require their absence, as does the updated version of *Donkey Kong* with the pie factory level<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>.

AxROM (mapper 7) is the only known licensed discrete logic mapper to unreliably come with [bus conflict](../Bus_conflict/) prevention circuitry. While no game documented in NesCartDB was released in one region on multiple board variants, several games did change boards when localized.

New ROMs are encouraged to specify submapper 2 to verify correct behaviour with bus conflicts.

CNROM with only 8 KiB of CHR-ROM (mapper 185) has a different set of submapper definitions.

### <span id="002:_0.2C_1.2C_2_UxROM"></span>002: 0, 1, 2 [UxROM](../UxROM/) {#002:_0.2C_1.2C_2_UxROM}

Test ROMs: ([forum post](https://forums.nesdev.org/viewtopic.php?p=157804#p157804))

- *2_test_0.nes* - Unspecified.
- *2_test_1.nes* - No bus conflicts.
- *2_test_2.nes* - AND bus conflicts.

### <span id="003:_0.2C_1.2C_2_CNROM"></span>003: 0, 1, 2 [CNROM](../CNROM/) {#003:_0.2C_1.2C_2_CNROM}

Test ROMs: ([forum post](https://forums.nesdev.org/viewtopic.php?p=154555#p154555))

- *3_test_0.nes* - Unspecified.
- *3_test_1.nes* - No bus conflicts.
- *3_test_2.nes* - AND bus conflicts.

### <span id="007:_0.2C_1.2C_2_AxROM"></span>007: 0, 1, 2 [AxROM](../AxROM/) {#007:_0.2C_1.2C_2_AxROM}

Test ROMs: ([forum post](https://forums.nesdev.org/viewtopic.php?p=157804#p157804))

- *7_test_0.nes* - Unspecified.
- *7_test_1.nes* - No bus conflicts.
- *7_test_2.nes* - AND bus conflicts.

## [004](../MMC3/): [MMC3](../MMC3/) {#004:_MMC3}

[iNES Mapper 004](../MMC3/) represents the most common boards using these four ICs: early [MMC3](../MMC3/), late MMC3, MC-ACC, and [MMC6](../MMC6/).

There are three known kinds of IRQ:

1.  NEC MMC3: IRQ is asserted on A12 rise, and loading the latch with 0 disables IRQ. At least one unlicensed game relies on this "old style" behavior.
2.  Sharp MMC3: IRQ is asserted on A12 rise, and loading the latch with 0 produces an IRQ on every scanline. MMC6 also behaves this way. Some later games rely on this "new style" behavior.
3.  MC-ACC: IRQ is asserted on A12 fall, typically four pixels later than Sharp MMC3. Interrupts can be produced every scanline, like the Sharp MMC3.<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

There are two known kinds of PRG RAM enable:

1.  MMC3: One set of enable bits controls the entire chip.
2.  MMC6: The first and second enables control the first and second half of PRG RAM, and an additional enable in bit 5 of \$8000 controls the whole PRG RAM.

The [TEROM and TFROM](../TxROM/) boards have two jumpers that can respectively disable IRQs and force hard-wired mirroring. It is believed that nothing was ever released that used them.

### 004: 0 Sharp MMC3 {#004:_0_Sharp_MMC3}

Normal.

- *Star Trek 25th Anniversary*

If the header is not using [NES 2.0](../NES_2.0/), an implementation without PRG-RAM write protection might be desired instead of a full MMC3. See: [iNES Mapper 004 and MMC6](../MMC3/#iNES_Mapper_004_and_MMC6)

### 004: 1 MMC6 {#004:_1_MMC6}

[MMC6](../MMC6/) has an alternative PRG-RAM enable and write protection scheme designed for its internal 1k PRG RAM. Uses Sharp MMC3 IRQ behavior.

- *StarTropics*
- *StarTropics 2*

### 004: 2 {#004:_2}

This describes MMC3C with hard-wired mirroring. No licensed games are known to require this, but a number of hacks and plug-and-play console extracts depend on it, and it is a valid configuration of the NES-TFROM circuit board.

### 004: 3 MC-ACC {#004:_3_MC-ACC}

MC-ACC

The MC-ACC is found in [13 second-source PCBs manufactured by Acclaim](https://nescartdb.com/search/basic?keywords=MC-ACC&kwtype=pcb):

- *Alien³*
- *George Foreman's KO Boxing*
- *The Incredible Crash Dummies*
- *Mickey's Safari in Letterland*
- *Roger Clemens' MVP Baseball*
- *Rollerblade Racer*
- *The Simpsons: Bart vs. The World*
- *The Simpsons: Bartman Meets Radioactive Man*
- *Spider-Man: Return of the Sinister Six*
- *T&C Surf Designs 2: Thrilla's Surfari*
- *T2: Terminator 2: Judgment Day*
- *WWF King of the Ring*
- *WWF WrestleMania: Steel Cage Challenge*

### 004: 4 NEC MMC3 {#004:_4_NEC_MMC3}

Loading the latch with 0 disables IRQ.

- SuperGame's *Aladdin* (Mapper 4 version)

### 004: 5 T9552 {#004:_5_T9552}

Indicates the presence of a [T9552](../T9552/#ROM_file_bank_order) scrambling chip.

## [016](../INES_Mapper_016/): Bandai FCG board {#016:_Bandai_FCG_board}

<table class="wikitable">
<thead>
<tr>
<th colspan="3"><a href="../INES_Mapper_016/Submapper_table/">INES Mapper 016 submapper table</a></th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Meaning</th>
<th>Note</th>
</tr>
&#10;<tr>
<td>0</td>
<td>Unspecified</td>
<td>Emulate both FCG-1/2 and LZ93D50 chips in their respective CPU address ranges.</td>
</tr>
<tr>
<td><em>1</em></td>
<td><em>LZ93D50 with 128 byte serial EEPROM (24C01)</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_159/">INES Mapper 159</a> instead.</em></td>
</tr>
<tr>
<td><em>2</em></td>
<td><em>Datach Joint ROM System</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_157/">INES Mapper 157</a> instead.</em></td>
</tr>
<tr>
<td><em>3</em></td>
<td><em>8 KiB of WRAM instead of serial EEPROM</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_153/">INES Mapper 153</a> instead.</em></td>
</tr>
<tr>
<td>4</td>
<td>FCG-1/2</td>
<td>Responds only in the CPU $6000-$7FFF address range; IRQ counter is not latched.</td>
</tr>
<tr>
<td>5</td>
<td>LZ93D50 with no or 256-byte serial EEPROM (24C02)</td>
<td>Responds only in the CPU $8000-$FFFF address range; IRQ counter is latched.</td>
</tr>
</tbody>
</table>

## <span id="019:_Namco_129.2F163"></span>[019](../INES_Mapper_019/): Namco 129/163 {#019:_Namco_129.2F163}

The NES 2.0 submapper is used to specify the mixing resistor that determines the relative volume of the expansion audio channels against the APU's audio channels:

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th colspan="3"><a href="../INES_Mapper_019/Submapper_table/">INES Mapper 019 submapper table</a></th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Meaning</th>
<th>Note</th>
</tr>
&#10;<tr>
<td>0</td>
<td>Default</td>
<td>Expansion sound volume unspecified</td>
</tr>
<tr>
<td><em>1</em></td>
<td><em>Deprecated</em></td>
<td>Internal 128b RAM is battery backed, no external PRG-RAM is present.
<p>No expansion sound. (Equivalent to submapper 2 with 0 in PRG-NVRAM field.)</p></td>
</tr>
<tr>
<td>2</td>
<td>No expansion sound</td>
<td></td>
</tr>
<tr>
<td>3</td>
<td>N163 expansion sound: <strong>11.0-13.0 dB</strong> louder than NES APU</td>
<td></td>
</tr>
<tr>
<td>4</td>
<td>N163 expansion sound: <strong>16.0-17.0 dB</strong> louder than NES APU</td>
<td></td>
</tr>
<tr>
<td>5</td>
<td>N163 expansion sound: <strong>18.0-19.5 dB</strong> louder than NES APU</td>
<td></td>
</tr>
</tbody>
</table>

## <span id="021.2C_023.2C_025:_VRC2_.2F_VRC4"></span>[021](../VRC2_and_VRC4/), [023](../VRC2_and_VRC4/), [025](../VRC2_and_VRC4/): [VRC2](../VRC2_and_VRC4/) / [VRC4](../VRC2_and_VRC4/) {#021.2C_023.2C_025:_VRC2_.2F_VRC4}

These three mappers collect various configurations of [VRC2](../VRC2_and_VRC4/) and [VRC4](../VRC2_and_VRC4/) boards.

VRC2 is mostly a subset of VRC4, with differences including:

- VRC2 has a serial EEPROM interface which no Konami game ever connected an actual EEPROM to, but one PCB is wired so that it functions as a single bit of memory mapped at \$6000-6FFF<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)</sup>
- VRC4 supports horizontal, vertical and one-screen nametable mirroring while VRC2 supports only horizontal and vertical mirroring<sup><span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup>
- VRC4 supports two PRG ROM banking modes, similar to MMC3's. VRC2 only has one PRG ROM banking mode.
- VRC4 has an interrupt device that VRC2 does not.

Additionally, different boards connect the address lines for the registers in various arrangements. The three iNES mappers 21, 23, and 25 each combine multiple boards whose addresses overlap. This is enough for game compatibility, but it creates an ugly combination that does not accurately describe the original hardware of either board.

### Submapper assignment {#Submapper_assignment}

Konami's [VRC2 and VRC4](../VRC2_and_VRC4/) mappers have several variations of how the board connects low CPU address lines among A7-A0 to the port select lines of the mapper. Mappers [21](../VRC2_and_VRC4/), [23](../VRC2_and_VRC4/) and [25](../VRC2_and_VRC4/) each combine two non-overlapping addressing variations, and neglect to specify VRC2 vs VRC4.

Mapper [22](../VRC2_and_VRC4/) is related, but only supports one variation (VRC2a) and has a different CHR banking arrangement than the others. It does not require submapper disambiguation.

There are three variations of the [VRC2](../VRC2_and_VRC4/) boards, and five of [VRC4](../VRC2_and_VRC4/):

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
<td>-</td>
<td>A0</td>
<td>A1</td>
<td>$x000, $x001, $x002, $x003</td>
<td>23</td>
<td>1</td>
</tr>
</tbody>
</table>

† <span id="VRC2bPCB"></span> The VRC2b PCBs include: 350603, 350636, 350926, 351179

This submapper assigment uses the following symmetrical arrangement:

- 0\. Both addressing
- 1\. Lower addressing, VRC4
- 2\. Higher addressing, VRC4
- 3\. Lower addressing, VRC2 *(no known examples for mapper 21)*
- 4\. *Higher addressing, VRC2 (no known examples)*

### <span id="021_.2F_023_.2F_025:_0"></span>021 / 023 / 025: 0 {#021_.2F_023_.2F_025:_0}

The default implementation acts as a VRC4 (mostly compatible superset of VRC2), and responds to register writes in one or more configurations simultaneously (supporting both boards at once).

### <span id="021_.2F_023_.2F_025:_1.2C_2_VRC4"></span>021 / 023 / 025: 1, 2 VRC4 {#021_.2F_023_.2F_025:_1.2C_2_VRC4}

These allocations each request a single specific addressing scheme for VRC4, rather than the combined version used by submapper 0.

- 021: 1 VRC4a
- 025: 1 VRC4b
- 021: 2 VRC4c
- 025: 2 VRC4d
- 023: 2 VRC4e
- 023: 1 VRC4f

### 023: 1 VRC4f {#023:_1_VRC4f}

Some unlicensed games used a second-source VRC4 clone with register addressing identical to VRC2b (the simplest contiguous in-order interpretation: 0,1,2,3). This was used in *World Hero*.

### 023: 3 VRC2b {#023:_3_VRC2b}

Games on PCB 350926 (and its chip-on-board equivalents) rely on a single bit of RAM mapped in the region from \$6000-\$6FFF. Prior to NES 2.0, these games were supported by implementing WRAM in this region (despite this board not having WRAM). This board was used in *[Contra (J)](https://nescartdb.com/profile/view/3986)*, *[Ganbare Goemon 2](https://nescartdb.com/profile/view/1568)* and *[Konami Wai Wai World](https://nescartdb.com/profile/view/2274)*.

There are other VRC2 boards with the same register addressing as 350926 but without this feature (\$6000-\$6FFF is effectively open bus on them), but no game relies on the *lack* of the single bit of memory, so a separate submapper has not been allocated for them. These boards were used in *[Dragon Scroll](https://nescartdb.com/profile/view/3802)*, *[Getsu Fuuma Den](https://nescartdb.com/profile/view/2276)*, and *[Jarinko Chie](https://nescartdb.com/profile/view/2278)*.

Since this submapper represents VRC2, it should not implement the VRC4-only one-screen mirroring, PRG ROM banking modes, or IRQ capabilities. Wai Wai World depends on the lack of one-screen mirroring; it will have corrupt backgrounds in the vertically-scrolling stages and the ending if emulated with VRC4 capabilities.

### 025: 3 VRC2c {#025:_3_VRC2c}

*[Ganbare Goemon Gaiden: Kieta Ougon Kiseru](https://nescartdb.com/profile/view/3823)* uses the same register addressing as VRC4b, but is VRC2, so this submapper should not implement VRC4-only capabilities.

### Test ROMs {#Test_ROMs}

- [VRC2/4 test program](https://forums.nesdev.org/viewtopic.php?p=203716#p203716)

## [032](../INES_Mapper_032/): Irem G101 {#032:_Irem_G101}

A variation of this mapper requires hardwired one-screen mirroring and entirely ignores writes to \$9000.

### 032: 0 {#032:_0}

Normal (H/V mapper-controlled mirroring)

### 032: 1 Major League {#032:_1_Major_League}

CIRAM A10 is tied high (fixed one-screen mirroring) and PRG banking style is fixed as 8+8+16F

- *[Major League](https://nescartdb.com/profile/view/3800)*

## <span id="034:_BNROM_.2F_NINA-001"></span>[034](../INES_Mapper_034/): [BNROM](../INES_Mapper_034/) / [NINA-001](../INES_Mapper_034/) {#034:_BNROM_.2F_NINA-001}

This iNES mapper unfortunately combines the unrelated [BNROM](../INES_Mapper_034/) and [NINA-001](../INES_Mapper_034/) mappers.

### 034: 0 {#034:_0}

Normal.

To disambiguate the two mappers, emulators have taken various approaches:

- The presense of CHR larger than 8 KiB unambiguously requires NINA-001, as BNROM has no CHR banking.
- The presence of CHR-RAM is taken to imply BNROM, because both extant BNROM games use CHR-RAM.
- CRC tests may be used to select a mapper for previously known ROMs.
- Implement both mappers simultaneously. This is compatible with existing games.

Selecting a single implementation based on CHR results in greater accuracy, since no game was ever intended for the combined definition.

New programs should not attempt to use a combined BNROM + NINA-001 mapper because this is not reliably available across emulators. Unusual combinations like NINA-001 with CHR-RAM are theoretically possible, but unlikely to be emulated consistently.

### 034: 1 NINA-001 {#034:_1_NINA-001}

[NINA-001](../INES_Mapper_034/) only.

Test ROM: ([forum post](https://forums.nesdev.org/viewtopic.php?p=153334#p153334))

- *34_test_1.nes* - NINA-001 with CHR ROM.

### 034: 2 BNROM {#034:_2_BNROM}

[BNROM](../INES_Mapper_034/) only.

Some unlicensed boards by Union Bond were a variation of BNROM that included PRG RAM. These may also use this submapper if PRG RAM is specified in the NES 2.0 header.

Test ROM: ([forum post](https://forums.nesdev.org/viewtopic.php?p=153334#p153334))

- *34_test_2.nes* - BNROM with CHR RAM.

## [068](../INES_Mapper_068/): Sunsoft 4 {#068:_Sunsoft_4}

In addition to its normal function, the Sunsoft 4 IC was used in *Nantettatte!! Baseball*, which allowed a second expansion cartridge to be plugged into it.

### 068: 0 {#068:_0}

0: Normal (max 256KiB PRG)

### 068: 1 Dual Cartridge System {#068:_1_Dual_Cartridge_System}

1: Sunsoft Dual Cartridge System a.k.a. NTB-ROM (max 128KiB PRG, licensing IC present, external option ROM of up to 128KiB should be selectable by a second menu)

- [*Nantettatte!! Baseball*](https://nescartdb.com/profile/view/2265)

## [071](../INES_Mapper_071/): Codemasters {#071:_Codemasters}

Some games use this with 1-screen [mirroring](../Mirroring/), where the mapper's mirroring control bit is wired directly to CIRAM A10. Others have hardwired horizontal or vertical mirroring.

Another variation of this mapper was used in the *Quattro* multicarts, but these have been reassigned to [mapper 232](../INES_Mapper_232/).

### 071: 0 {#071:_0}

Hardwired horizontal or vertical mirroring.

### 071: 1 Fire Hawk {#071:_1_Fire_Hawk}

Mapper controlled single-screen mirroring.

- *[Fire Hawk](https://nescartdb.com/profile/view/733)*

## <span id="078:_Cosmo_Carrier_.2F_Holy_Diver"></span>[078](../INES_Mapper_078/): Cosmo Carrier / Holy Diver {#078:_Cosmo_Carrier_.2F_Holy_Diver}

This mapper unfortunately combines two games with incompatible mirroring control.

One game uses this with 1-screen [mirroring](../Mirroring/), where the mapper's mirroring control bit is wired directly to CIRAM A10. The other can switch between horizontal and vertical mirroring, using a multiplexer between PPU A10 and PPU A11 whose output is sent to CIRAM A10.

Historically, pre-NES 2.0 ROM image headers set the 4-screen VRAM flag for Holy Diver and cleared it for Cosmo Carrier.

### 078: 0 {#078:_0}

Unspecified.

### 078: 1 Cosmo Carrier {#078:_1_Cosmo_Carrier}

Single-screen mirroring (nibble-swapped [mapper 152](../INES_Mapper_152/)).

- *[Uchuusen: Cosmo Carrier](https://nescartdb.com/profile/view/3968)*

### 078: 2 {#078:_2}

**Deprecated.**

This described a variation with fixed vertical mirroring, and WRAM. There is no known use case.

### 078: 3 Holy Diver {#078:_3_Holy_Diver}

Mapper-controlled H/V mirroring.

- *[Holy Diver](https://nescartdb.com/profile/view/4038)*

## [085](../VRC7/): Konami VRC7 {#085:_Konami_VRC7}

- VRC7b uses A3 to select registers (\$x008), used in Tiny Toon Adventures (Submapper 1).
- VRC7a uses A4 to select registers (\$x010), used in Lagrange Point (Submapper 2).

## [091](../INES_Mapper_091/): Super Fighter III {#091:_Super_Fighter_III}

- Submapper 0: YY830624C/JY830848C (J.Y. Company clone board)
- Submapper 1: EJ-006-1 (original *Super Fighter III* board).

## <span id="114:_Sugar_Softec.2FHosenkan.2C_fixed_MMC3_scrambling_pattern"></span>[114](../INES_Mapper_114/): Sugar Softec/Hosenkan, fixed MMC3 scrambling pattern {#114:_Sugar_Softec.2FHosenkan.2C_fixed_MMC3_scrambling_pattern}

Registers at \$6000 (and \$6001, although not used for effect by existing games)

- Submapper 0: The Lion King, Aladdin
- Submapper 1: Boogerman (different scrambling pattern)

## [178](../INES_Mapper_178/).1: Gameinis Infrared Sensor {#178.1:_Gameinis_Infrared_Sensor}

- Submapper 0: No infrared sensor present
- Submapper 1: Gameinis Infrared Sensor present

## [185](../CNROM/): CNROM used to selectively disable CHR-ROM {#185:_CNROM_used_to_selectively_disable_CHR-ROM}

A few NROM-like games were released on CNROM boards where all four bits of the latch were solely used as an anti-piracy measure. Games would check that CHR-ROM is disabled unless the latch holds the correct one of the four possible bank values. See this [thread](http://forums.nesdev.org/viewtopic.php?f=3&t=16700) for test ROMs and example implementation code.

### 185: 0 {#185:_0}

Correct bank value unknown. A heuristic that works with all known games that use this mapper is to return \$FF on the first two reads from CPU \$2007 after a reset.

### 185: 4 {#185:_4}

Enable CHR-ROM if bits 0..1 of the latch hold the value 0, otherwise disable CHR-ROM.

### 185: 5 {#185:_5}

Enable CHR-ROM if bits 0..1 of the latch hold the value 1, otherwise disable CHR-ROM.

### 185: 6 {#185:_6}

Enable CHR-ROM if bits 0..1 of the latch hold the value 2, otherwise disable CHR-ROM.

### 185: 7 {#185:_7}

Enable CHR-ROM if bits 0..1 of the latch hold the value 3, otherwise disable CHR-ROM.

## [206](../INES_Mapper_206/): Namco 118 PCB variants {#206:_Namco_118_PCB_variants}

### 206: 0 {#206:_0}

Namcot 118 ASIC with normal PRG banking

### 206: 1 {#206:_1}

Namcot **3407**, **3417** and **3451** PCBs with unbanked 32 KiB PRG-ROM.

- *Family Jockey*
- *Metro-Cross*
- *Sky Kid*
- *Super Chinese*
- *Super Sky Kid*
- *Valkyrieの冒険꞉ 時の鍵伝説*

## [210](../INES_Mapper_210/): Namco 175 and 340 {#210:_Namco_175_and_340}

[Mapper 210](../INES_Mapper_210/) doesn't distinguish between the 175's hardwired mirroring and 340's 1/H/V mirroring.

Also, previous confusion and compatibility code used by Namco when they were developing games means that many 175- and 340- using games are incorrectly tagged as [mapper 19](../INES_Mapper_019/).

### 210: 0 {#210:_0}

Namco 175 if the header specifies PRG RAM or if the battery flag is set in a NES 1.0 header; Namco 340 otherwise. Luckily none of the 175-using games actually depends on the hardwired mirroring; all of them write the correct mirroring bits to work on a 340.

### 210: 1 N175 {#210:_1_N175}

Namco 175. Hardwired mirroring, no IRQ.

- *Chibi Maruko-chan: Uki Uki Shopping*
- *Famista '91*
- *Family Circuit '91*
- *Heisei Tensai Bakabon / Genius Bakabon*
- *Splatterworld: Rick to Kyoufu no Daiou* (unreleased)

### 210: 2 N340 {#210:_2_N340}

Namco 340. 1/H/V mirroring, no IRQ, no internal or external RAM.

- *Dream Master*
- *Famista '92*
- *Famista '93*
- *Famista '94*
- *Splatterhouse: Wanpaku Graffiti*
- *Top Striker*
- *Wagyan Land 2*
- *Wagyan Land 3*

## <span id="215:_Sugar_Softec.2C_selectable_MMC3_scrambling_pattern"></span>[215](../INES_Mapper_215/): Sugar Softec, selectable MMC3 scrambling pattern {#215:_Sugar_Softec.2C_selectable_MMC3_scrambling_pattern}

Registers at \$5000-\$5007.

- Submapper 0: UNL-8237 (two non-overlapping outer bank bits for PRG/CHR each)
- Submapper 1: UNL-8237A (three overlapping outer bank bits for PRG/CHR each)

## [232](../INES_Mapper_232/): Quattro {#232:_Quattro}

Similar to \#71 above, with a separate register controlling which 64 KiB outer bank of the PRG ROM is used. This is used for the [Quattro multicarts](https://nescartdb.com/search/basic?keywords=quattro&kwtype=game).

The Aladdin Deck Enhancer version of these multicarts used a different banking scheme.

### 232: 0 {#232:_0}

0: Normal

### 232: 1 Aladdin Deck Enhancer {#232:_1_Aladdin_Deck_Enhancer}

Aladdin Deck Enhancer variation. Swap the bits of the outer bank number.

- *[Quattro Adventure](https://nescartdb.com/profile/view/553) (Aladdin Deck Enhancer)*
- *[Quattro Sports](https://nescartdb.com/profile/view/550) (Aladdin Deck Enhancer)*
- *Pegasus 4-in-1*
- *Super Adventure Quests* (European release of *Quattro Adventure*)

## [NES 2.0 Mapper 256](../NES_2.0_Mapper_256/) {#NES_2.0_Mapper_256}

Several plug-and-play consoles use a simple protection scheme intended to prevent interchanging the ROM data between different consoles. The protection scheme works by simply swapping the addresses of the bankswitch registers. Some patterns only change the VTxx-native register addresses, others apply to the MMC3 compatibility registers as well.

Other plug-and-play consoles, in particular those by Jungletac, scramble CPU opcode bytes (and only those) by swapping bits. These consoles power-on with opcode scrambling enabled, but allow it to be disabled --- for example, to run a pirated licensed NES game without modifying it --- by writing to a register.

<table class="wikitable">
<thead>
<tr>
<th>Submapper #</th>
<th>Name</th>
<th colspan="12">PPU bank affected by ...</th>
<th colspan="4">CPU bank affected by ...</th>
<th>CPU opcode bytes</th>
</tr>
</thead>
<tbody>
<tr>
<th>$2012</th>
<th>$2013</th>
<th>$2014</th>
<th>$2015</th>
<th>$2016</th>
<th>$2017</th>
<th>$8000.0</th>
<th>$8000.1</th>
<th>$8000.2</th>
<th>$8000.3</th>
<th>$8000.4</th>
<th>$8000.5</th>
<th>$4107</th>
<th>$4108</th>
<th>$8000.6</th>
<th>$8000.7</th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>Normal</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>1</td>
<td>Waixing VT03</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>2</td>
<td>Power Joy Supermax</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>none</td>
</tr>
<tr>
<td>3</td>
<td>Zechess/Hummer Team</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1000</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>4</td>
<td>Qishenglong</td>
<td>$1800</td>
<td>$0800</td>
<td>$1000</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>5</td>
<td>Waixing VT02</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>11</td>
<td>Vibes</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6; D5&lt;-&gt;D4 swapped, switched via $411C.1</td>
</tr>
<tr>
<td>12</td>
<td>Cheertone</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6</td>
</tr>
<tr>
<td>13</td>
<td>Taikee</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D1 and D4 swapped, switched via $4169</td>
</tr>
<tr>
<td>14</td>
<td>Karaoto</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D6 and D7 swapped, switched via $411C</td>
</tr>
<tr>
<td>15</td>
<td>Jungletac</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D5 and D6 swapped, switched via $4169</td>
</tr>
</tbody>
</table>

## <span id="268:_Coolboy.2FMindkids"></span>[268](../NES_2.0_Mapper_268/): Coolboy/Mindkids {#268:_Coolboy.2FMindkids}

- Submapper 0: Coolboy (outer bank registers at \$6000-\$6003)
- Submapper 1: Mindkids (outer bank registers \$5000-\$5003)

## References {#References}

- [Original submapper list by kevtris](http://blog.kevtris.org/blogfiles/nes/submappers.txt)
- [Other definitions by kevtris](http://blog.kevtris.org/blogfiles/nes/)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [NESDev post](http://forums.nesdev.org/viewtopic.php?t=7273) explaining Cybernoid's bugs.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Nintendoage thread](http://www.nintendoage.com/forum/messageview.cfm?catid=22&threadid=44613) discussing Donkey Kong Pie Factory.
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) [MC-ACC IRQ test results](http://forums.nesdev.org/viewtopic.php?p=128148#p128148)
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) <http://forums.nesdev.org/viewtopic.php?t=8274> VRC2 memory bit at \$6000-\$6FFF
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) <http://forums.nesdev.org/viewtopic.php?t=13473> VRC2 mirroring selection tested on hardware
