---
title: "INES Mapper 019"
url: "/wiki/INES_Mapper_019/"
description: "iNES Mapper 019 is used to denote Famicom boards bearing the Namco 129 and Namco 163 ASICs. These chips contain 128 bytes of internal RAM that can be used either for expansion audio or, together with…"
wikiPage: "INES_Mapper_019"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with CHR ROM and CHR RAM"
    url: "/category/Mappers_with_CHR_ROM_and_CHR_RAM/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
---

<div class="infobox"><div style="text-align: center"><b>Namco 129, 163</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Namco </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=19">20 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>11xF0, 111F0S </td></tr><tr><td><b>Pinout</b></td><td><a href="../Namcot_163_family_pinout/">Namcot 163 family pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 (PT) + 1Kx4 (NT) </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>arbitrary, up to 226 source nametables </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>CPU cycle counter </td></tr><tr><td><b>Audio</b></td><td><a href="../Namco_163_audio/">Yes</a></td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>019 </td></tr></tbody></table></div>

**iNES Mapper 019** is used to denote Famicom boards bearing the **Namco 129** and **Namco 163** ASICs. These chips contain 128 bytes of internal RAM that can be used either for [expansion audio](../Namco_163_audio/) or, together with a battery, for 128 bytes of save RAM. Games may also come with the regular 8 KiB of additional battery-backed WRAM mapped into CPU address space at \$6000-\$7FFF.

In [NES 2.0](../NES_2.0/), the 128 bytes of internal chip RAM are not included in the PRG-RAM and PRG-NVRAM fields. NES 2.0 Mapper 19 ROM images therefore either have:

- no battery and no WRAM: battery bit clear, PRG-RAM/PRG-NVRAM both set to zero.
- a battery but no WRAM: battery bit set, PRG-RAM/PRG-NVRAM both set to zero. The game writes its save data into the 128 byte internal RAM.
- a battery and WRAM: battery bit set, PRG-RAM set to zero and PRG-NVRAM set to 8192.

The NES 2.0 submapper is used to specify the mixing resistor that determines the relative volume of the expansion audio channels against the [APU](../APU/)'s audio channels:

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th colspan="3"><a href="Submapper_table/">INES Mapper 019 submapper table</a></th>
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

## Game list {#Game_list}

| Name | Chip | Battery | WRAM | Expansion Audio | NES 2.0 Submapper |
|----|----|----|----|----|----|
| *Battle Fleet* | 163 | Yes | No | No | 2 |
| *Dragon Ninja* | 163 | No | No | No | 2 |
| *Digital Devil Story: Megami Tensei II* | 163 | Yes | Yes | Yes | 3 |
| *Dokuganryuu Masamune* | 163 | Yes | No | No | 2 |
| *Erika to Satoru no Yume Bouken* | 163 | No | No | Yes | 5 |
| *Famista '90* | 163 | Yes | No | No | 2 |
| *Final Lap* | 163 | No | No | Yes | 3 |
| *Hydlide 3: Yami kara no Houmonsha* | 163 | Yes | No | No | 2 |
| *Juvei Quest* | 163 | Yes | Yes | No | 2 |
| *Kaijuu Monogatari* | 163 | Yes | No | No | 2 |
| *King of Kings* | 163 | Yes | Yes | Yes | 5 |
| *Mappy Kids* | 163 | No | No | Yes | 5 |
| *Mindseeker* | 163 | Yes | No | No | 2 |
| *Namco Classic* | 163 | No | No | No | 2 |
| *Namco Classic II* | 163 | No | No | Yes | 3 |
| *Rolling Thunder* | 163 | No | No | Yes | 4 |
| *Sangokushi: Chuugen no Hasha* | 163 | Yes | Yes | Yes | 5 |
| *Sangokushi II: Haou no Tairiku* | 163 | Yes | Yes | Yes | 3 |
| *Star Wars* | 129 | No | No | No | 2 |
| *Youkai Douchuuki* | 163 | No | No | Yes | 5 |

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB PRG RAM bank (optional)
- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank
- PPU \$2000-\$23FF: 1 KiB switchable CHR bank
- PPU \$2400-\$27FF: 1 KiB switchable CHR bank
- PPU \$2800-\$2BFF: 1 KiB switchable CHR bank
- PPU \$2C00-\$2FFF: 1 KiB switchable CHR bank

These ASICs have the unusual ability to select the internal 2 KiB nametable RAM as a CHR bank page, allowing it to be used as CHR RAM in combination with the existing CHR ROM.

## Registers {#Registers}

The 163 has 19 registers within \$4800-\$5FFF and \$8000-\$FFFF. Each register occupies a range of \$800 bytes, so \$4800-\$4FFF all refers to one register, \$5000-\$57FF all refers to another register, and so on.

### <span id="Chip_RAM_Data_Port_.28.244800-.244FFF.2C_read.2Fwrite.29"></span>Chip RAM Data Port (\$4800-\$4FFF, read/write) {#Chip_RAM_Data_Port_.28.244800-.244FFF.2C_read.2Fwrite.29}

See [Namco 163 audio](../Namco_163_audio/).

### <span id="IRQ_Counter_.28low.29_.28.245000-.2457FF.2C_read.2Fwrite.29"></span>IRQ Counter (low) (\$5000-\$57FF, read/write) {#IRQ_Counter_.28low.29_.28.245000-.2457FF.2C_read.2Fwrite.29}

    7  bit  0
    ---- ----
    LLLL LLLL
    |||| ||||
    ++++-++++- Low 8 bits of IRQ counter

Writes to this register (but not reads) will acknowledge pending IRQs.

### <span id="IRQ_Counter_.28high.29_.2F_IRQ_Enable_.28.245800-.245FFF.2C_read.2Fwrite.29"></span>IRQ Counter (high) / IRQ Enable (\$5800-\$5FFF, read/write) {#IRQ_Counter_.28high.29_.2F_IRQ_Enable_.28.245800-.245FFF.2C_read.2Fwrite.29}

    7  bit  0
    ---- ----
    EHHH HHHH
    |||| ||||
    |+++-++++- High 7 bits of IRQ counter
    +--------- IRQ Enable: (0: disabled; 1: enabled)

Writes to this register (but not reads) will acknowledge pending IRQs.

### <span id="CHR_and_NT_Select_.28.248000-.24DFFF.2C_write.29"></span>CHR and NT Select (\$8000-\$DFFF, write) {#CHR_and_NT_Select_.28.248000-.24DFFF.2C_write.29}

| Value CPU writes | Behavior |
|----|----|
| \$00-\$DF | Always selects 1 KiB page of CHR-ROM |
| \$E0-\$FF | If enabled by bit in \$E800, use the NES's internal nametables (even values for A, odd values for B) |

| Write to CPU address | 1 KiB CHR bank affected | Values ≥ \$E0 denote NES NTRAM if |
|----|----|----|
| \$8000-\$87FF | \$0000-\$03FF | \$E800.6 = 0 |
| \$8800-\$8FFF | \$0400-\$07FF | \$E800.6 = 0 |
| \$9000-\$97FF | \$0800-\$0BFF | \$E800.6 = 0 |
| \$9800-\$9FFF | \$0C00-\$0FFF | \$E800.6 = 0 |
| \$A000-\$A7FF | \$1000-\$13FF | \$E800.7 = 0 |
| \$A800-\$AFFF | \$1400-\$17FF | \$E800.7 = 0 |
| \$B000-\$B7FF | \$1800-\$1BFF | \$E800.7 = 0 |
| \$B800-\$BFFF | \$1C00-\$1FFF | \$E800.7 = 0 |
| \$C000-\$C7FF | \$2000-\$23FF | always |
| \$C800-\$CFFF | \$2400-\$27FF | always |
| \$D000-\$D7FF | \$2800-\$2BFF | always |
| \$D800-\$DFFF | \$2C00-\$2FFF | always |

It is believed, but untested, that a game could add a normal SRAM and use it in lieu of the nametable RAM; if so, a game would be able to get 4-screen mirroring and many more pages of CHR-RAM.

### <span id="PRG_Select_1_.28.24E000-.24E7FF.2C_write.29"></span>PRG Select 1 (\$E000-\$E7FF, write) {#PRG_Select_1_.28.24E000-.24E7FF.2C_write.29}

    7  bit  0
    ---- ----
    AMPP PPPP
    |||| ||||
    ||++-++++- Select 8 KiB PRG ROM bank at $8000-$9FFF
    |+-------- Disable sound if set
    +--------- Pin 22 (open collector) reflects the inverse of this value, unchanged by the address bus inputs.

If not battery backed, the internal sound RAM powers on in a random state, and similarly the N163 may power on with sound enabled. A short burst of sound due to these random settings may be heard at first power-on, before the game software manages to disable it.

There is a report that pin 22 might not be /\$E000.7, neither open collector output, based on [decapping](https://www.qmtpro.com/~nes/chipimages/visualn163/). Further test is required before confirming it fully.

### <span id="PRG_Select_2_.2F_CHR-RAM_Enable_.28.24E800-.24EFFF.2C_write.29"></span>PRG Select 2 / CHR-RAM Enable (\$E800-\$EFFF, write) {#PRG_Select_2_.2F_CHR-RAM_Enable_.28.24E800-.24EFFF.2C_write.29}

    7  bit  0
    ---- ----
    HLPP PPPP
    |||| ||||
    ||++-++++- Select 8 KiB PRG ROM bank at $A000-$BFFF
    |+-------- Disable CHR-RAM at $0000-$0FFF
    |              0: Pages $E0-$FF use NT RAM as CHR-RAM
    |              1: Pages $E0-$FF are the last $20 banks of CHR-ROM
    +--------- Disable CHR-RAM at $1000-$1FFF
                   0: Pages $E0-$FF use NT RAM as CHR-RAM
                   1: Pages $E0-$FF are the last $20 banks of CHR-ROM

### <span id="PRG_Select_3_.28.24F000-.24F7FF.2C_write.29"></span>PRG Select 3 (\$F000-\$F7FF, write) {#PRG_Select_3_.28.24F000-.24F7FF.2C_write.29}

    7  bit  0
    ---- ----
    CDPP PPPP
    |||| ||||
    ||++-++++- Select 8 KiB PRG ROM bank at $C000-$DFFF
    |+-------- Pin 44 output for PPU $2000-$3FFF
    +--------- Pin 44 output for PPU $0000-$0FFF

For PPU \$1000-\$1FFF, pin 44 always outputs 1.

Writing \$3F here (selecting the final bank and disabling both pin 44 bits) replaces the CHR bank outputs with debugging information for the internal audio state, varying according to the state of PPU A10-A12.

### <span id="Write_Protect_for_External_RAM_and_Chip_RAM_Address_Port_.28.24F800-.24FFFF.2C_write.29"></span>Write Protect for External RAM and Chip RAM Address Port (\$F800-\$FFFF, write) {#Write_Protect_for_External_RAM_and_Chip_RAM_Address_Port_.28.24F800-.24FFFF.2C_write.29}

    7  bit  0
    ---- ----
    KKKK DCBA
    |||| ||||
    |||| |||+- 1: Write-protect 2 KiB window of external RAM from $6000-$67FF (0: write enable)
    |||| ||+-- 1: Write-protect 2 KiB window of external RAM from $6800-$6FFF (0: write enable)
    |||| |+--- 1: Write-protect 2 KiB window of external RAM from $7000-$77FF (0: write enable)
    |||| +---- 1: Write-protect 2 KiB window of external RAM from $7800-$7FFF (0: write enable)
    ++++------ Additionally the upper nybble must be equal to b0100 to enable writes

Any value outside of the range \$40-\$4E will cause all PRG RAM to be read-only.

This register is also used as address port of the ASICs' internal RAM. See [Namco 163 audio](../Namco_163_audio/).

## IRQ Operation {#IRQ_Operation}

The IRQ is a 15-bit CPU cycle **up-counter**. \$5000 and \$5800 provide *direct* access to the counter itself (i.e., this isn't a reload value). Games can read and write to these registers in realtime.

The IRQ counter increments on each CPU cycle while enabled (\$5800.D7 == 1). Upon reaching \$7FFF, an IRQ is fired and the counter will stop counting.

Writes to \$5000 or \$5800 will acknowledge any pending IRQs, but not reads from these registers. Scroll/CHR bankswitch splits will frequently fail and cause flickering in games such as *Namco Classic II* if emulated incorrectly.

## Notes {#Notes}

- Many older programs incorrectly identify this mapper by the name Namco 106. Some sources use the name Namcot instead of Namco, since this was the name of Namco's console game division from 1984 to 1995.
- The N163 supports 8 KiB of PRG-RAM but also contains 128 bytes of internal memory. If there is a battery, then both will be battery backed. The internal memory is battery backed for several games that have no additional PRG-RAM.
- If it is not battery backed, the internal RAM state will be undefined at power-on<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>.
- The only known difference between the Namco 129 and 163 is the 129's subtly different design of expansion audio. The only known game (Star Wars) using the Namco 129 can also be found in later runs with a Namco 163 ASIC.
- An IC labelled "160" has been found with a copy of *Dokuganryuu Masamune*: [\[1\]](https://forums.nesdev.org/viewtopic.php?p=248291#p248291)
- Many [INES Mapper 210](../INES_Mapper_210/) games are incorrectly set to Mapper 019.
- *Dokuganryuu Masamune* is often thought as having WRAM. Tests have indicated however that the fourth glob is not SRAM but a [protection circuit for the chip-internal RAM](http://forums.nesdev.org/viewtopic.php?t=7727#p111097).

## CHR Memory Configurations {#CHR_Memory_Configurations}

All commercial-era titles only come with CHR-ROM. There is interest in the homebrew commmunity however in creating cartridges with both CHR-ROM and CHR-RAM. The following configurations are plausible and should be emulated as follows:

- Only CHR-ROM: CHR banks \$00-\$DF (or all if \$E800.6/7=1) go to CHR-ROM, and \$E0-\$FF go to CIRAM (normal commercial-era configuration).
- Only CHR-RAM: CHR banks \$00-\$DF (or all if \$E800.6/7=1) go to CHR-RAM, and \$E0-\$FF go to CIRAM.
- CHR-ROM of any size, plus up to 32 KiB of CHR-RAM: CHR banks \$00-\$DF (or all if \$E800.6/7=1) go to CHR-ROM, and \$E0-\$FF go to CHR-RAM.
- CHR-ROM plus more than 32 KiB of CHR-RAM are invalid.

There is no trivial means, short of manually decoding the CHR output address lines, to distinguish three types of CHR memory, CHR-ROM + CHR-RAM + CIRAM.

## See also {#See_also}

- [Namcot 106](http://nesdev.org/namco106.txt) by goroh, fix by ZW4 and nori, english translation by nori.
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.

All of the below are in Japanese:

- [Namco mapper](http://nesdev.org/namco.txt) by goroh
- [Naruko's post to the forum](http://forums.nesdev.org/viewtopic.php?p=77795#77795)
- [Wiki on same](http://unagi.sourceforge.jp/cgi-bin/hiki/hiki.cgi?163_namco)
- [Naruko's blog's post on the write protect register at \$F800](http://d.hatena.ne.jp/na6ko/20110430/1304099059)

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [N163 Sound RAM Initialisation Behaviour](https://forums.nesdev.org/viewtopic.php?p=285135)
