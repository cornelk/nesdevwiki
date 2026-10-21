---
title: "RAMBO-1"
url: "/wiki/RAMBO-1/"
description: "The Tengen RAMBO-1 is an ASIC mapper, canonically designated as mapper 64. This mapper is basically Tengen's version of the MMC3, but with some extra features. The RAMBO-1 came as a 40-pin PDIP. A…"
wikiPage: "RAMBO_1"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>RAMBO-1</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Tengen </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=64">5 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>800032 </td></tr><tr><td><b>Pinout</b></td><td><a href="../Tengen_RAMBO-1_pinout/">Tengen RAMBO-1 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>3×8K + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 or 2Kx2 + 1Kx4 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">064</a>, <a href="../INES_Mapper_158/">158</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=64">iNES 064</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif=TENGEN-800032">800032</a></td></tr></tbody></table></div>

*For the mapper used in the game "Rambo", see [UxROM](../UxROM/).*

The **Tengen RAMBO-1** is an [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/), canonically designated as **mapper 64**. This mapper is basically Tengen's version of the [MMC3](../MMC3/), but with some extra features. The RAMBO-1 came as a [40-pin PDIP](../Tengen_RAMBO-1_pinout/). A variant with different mirroring control is [mapper 158](../INES_Mapper_158/).

Example games:

- *Klax*
- *Skull and Crossbones*
- *Shinobi*
- *Rolling Thunder*
- *Hard Drivin' (prototype)*

## Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU -- Three selectable configurations:
  1.  1 KiB switchable CHR banks at \$0000, \$0400, \$0800, \$0C00, \$1000, \$1400, \$1800, \$1C00
  2.  2 KiB switchable CHR banks at \$0000, \$0800; 1 KiB switchable CHR banks at \$1000, \$1400, \$1800, \$1C00
  3.  2 KiB switchable CHR banks at \$1000, \$1800; 1 KiB switchable CHR banks at \$0000, \$0400, \$0800, \$0C00

## Registers {#Registers}

The RAMBO-1 has four pairs of registers at \$8000-\$9FFF, \$A000-\$BFFF, \$C000-\$DFFF, and \$E000-\$FFFF - even addresses (\$8000, \$8002, etc.) select the low register and odd addresses (\$8001, \$8003, etc.) select the high register in each pair.

### <span id="Bank_select_.28.248000-.249FFE.2C_even.29"></span>Bank select (\$8000-\$9FFE, even) {#Bank_select_.28.248000-.249FFE.2C_even.29}

    7  bit  0
    ---- ----
    CPKx RRRR
    |||  ||||
    |||  ++++- Specify which bank register to update on next write to Bank Data register
    |||          0000: R0: Select 2 (K=0) or 1 (K=1) KiB CHR bank at PPU $0000 (or $1000)
    |||          0001: R1: Select 2 (K=0) or 1 (K=1) KiB CHR bank at PPU $0800 (or $1800)
    |||          0010: R2: Select 1 KiB CHR bank at PPU $1000-$13FF (or $0000-$03FF)
    |||          0011: R3: Select 1 KiB CHR bank at PPU $1400-$17FF (or $0400-$07FF)
    |||          0100: R4: Select 1 KiB CHR bank at PPU $1800-$1BFF (or $0800-$0BFF)
    |||          0101: R5: Select 1 KiB CHR bank at PPU $1C00-$1FFF (or $0C00-$0FFF)
    |||          0110: R6: Select 8 KiB PRG ROM bank at $8000-$9FFF (or $C000-$DFFF)
    |||          0111: R7: Select 8 KiB PRG ROM bank at $A000-$BFFF
    |||          1000: R8: If K=1, Select 1 KiB CHR bank at PPU $0400 (or $1400)
    |||          1001: R9: If K=1, Select 1 KiB CHR bank at PPU $0C00 (or $1C00)
    |||          1111: RF: Select 8 KiB PRG ROM bank at $C000-$DFFF (or $8000-$9FFF)
    ||+------- Full 1 KiB CHR bank mode  0: two 2 KiB banks at $0000-$0FFF (or $1000-$1FFF)
    ||                                   1: four 1 KiB banks at $0000-$0FFF (or $1000-$1FFF)
    |+-------- PRG ROM bank mode  0: $8000-$9FFF uses bank selected with R6
    |                                $A000-$BFFF uses bank selected with R7
    |                                $C000-$DFFF uses bank selected with RF
    |                             1: $8000-$9FFF uses bank selected with RF
    |                                $A000-$BFFF uses bank selected with R7
    |                                $C000-$DFFF uses bank selected with R6
    +--------- CHR A12 inversion  0: two 2 KiB banks (or four 1 KiB banks) at $0000-$0FFF
                                     four 1 KiB banks at $1000-$1FFF
                                  1: two 2 KiB banks (or four 1 KiB banks) at $1000-$1FFF
                                     four 1 KiB banks at $0000-$0FFF

#### CHR Banks {#CHR_Banks}

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>When $8000 &amp; $80</th>
<th>is $00</th>
<th>is $80</th>
</tr>
</thead>
<tbody>
<tr>
<th>PPU Bank</th>
<th colspan="2">Value of RAMBO-1 register</th>
</tr>
&#10;<tr>
<td>$0000-$03FF</td>
<td>R0</td>
<td>R2</td>
</tr>
<tr>
<td>$0400-$07FF</td>
<td>K=0: R0<br />
K=1: R8</td>
<td>R3</td>
</tr>
<tr>
<td>$0800-$0BFF</td>
<td>R1</td>
<td>R4</td>
</tr>
<tr>
<td>$0C00-$0FFF</td>
<td>K=0: R1<br />
K=1: R9</td>
<td>R5</td>
</tr>
<tr>
<td>$1000-$13FF</td>
<td>R2</td>
<td>R0</td>
</tr>
<tr>
<td>$1400-$17FF</td>
<td>R3</td>
<td>K=0: R0<br />
K=1: R8</td>
</tr>
<tr>
<td>$1800-$1BFF</td>
<td>R4</td>
<td>R1</td>
</tr>
<tr>
<td>$1C00-$1FFF</td>
<td>R5</td>
<td>K=0: R1<br />
K=1: R9</td>
</tr>
</tbody>
</table>

With 2KiB banks (where K=0), only even-numbered CHR banks can be selected. The lowest bit in the register is ignored and instead passes PPU A10 directly through to CHR A10.

#### PRG Banks {#PRG_Banks}

Because the values in R6, R7, RF, and \$8000 are unspecified at power on, the reset vector must point into \$E000-\$FFFF, and code must initialize these before jumping out of \$E000-\$FFFF.

<table class="wikitable">
<thead>
<tr>
<th>When $8000 &amp; $40</th>
<th>is $00</th>
<th>is $40</th>
</tr>
</thead>
<tbody>
<tr>
<th>CPU Bank</th>
<th colspan="2">Value of RAMBO-1 register</th>
</tr>
&#10;<tr>
<td>$8000-$9FFF</td>
<td>R6</td>
<td>RF</td>
</tr>
<tr>
<td>$A000-$BFFF</td>
<td>R7</td>
<td>R7</td>
</tr>
<tr>
<td>$C000-$DFFF</td>
<td>RF</td>
<td>R6</td>
</tr>
<tr>
<td>$E000-$FFFF</td>
<td>(-1)</td>
<td>(-1)</td>
</tr>
</tbody>
</table>

- (-1) : the last bank

### <span id="Bank_data_.28.248001-.249FFF.2C_odd.29"></span>Bank data (\$8001-\$9FFF, odd) {#Bank_data_.28.248001-.249FFF.2C_odd.29}

All eight bits are used for a new value for the bank based on last value written to Bank select register (as mentioned above)

### <span id="Mirroring_.28.24A000-.24BFFE.2C_even.29"></span>Mirroring (\$A000-\$BFFE, even) {#Mirroring_.28.24A000-.24BFFE.2C_even.29}

    7  bit  0
    ---- ----
    xxxx xxxM
            |
            +- Mirroring (0: vertical; 1: horizontal)

This applies to **mapper 64** only (see [Variants](#Variants) below).

### <span id="Unknown_.28.24A001-.24BFFF.2C_odd.29"></span>Unknown (\$A001-\$BFFF, odd) {#Unknown_.28.24A001-.24BFFF.2C_odd.29}

This register is either not implemented or has undiscovered functionality on RAMBO-1. The MMC3/MMC6 equivalent of this register controls external and internal PRG-RAM access. There is not presently any reason to believe that RAMBO-1 has any support for PRG-RAM.

## IRQ registers {#IRQ_registers}

### <span id="IRQ_latch_.28.24C000-.24DFFE.2C_even.29"></span>IRQ latch (\$C000-\$DFFE, even) {#IRQ_latch_.28.24C000-.24DFFE.2C_even.29}

All eight bits of this register specifies the IRQ counter reload value. When the IRQ counter is zero (or a reload is requested through \$C001), this value will be copied into the IRQ counter at the end of the current scanline.

### <span id="IRQ_mode_select_.2F_reload_.28.24C001-.24DFFF.2C_odd.29"></span>IRQ mode select / reload (\$C001-\$DFFF, odd) {#IRQ_mode_select_.2F_reload_.28.24C001-.24DFFF.2C_odd.29}

    7  bit  0
    ---- ----
    xxxx xxxM
            |
            +- IRQ mode select (0: Scanline Mode, 1: CPU Cycle Mode)

Writing to this register also clears the IRQ counter so that it will be reloaded at next clock, or the next scanline, depending on the selected mode. This also resets the prescaler in cycle mode, so the next clock will occur 4 cycles later.

### <span id="IRQ_acknowledge_.2F_disable_.28.24E000-.24FFFE.2C_even.29"></span>IRQ acknowledge / disable (\$E000-\$FFFE, even) {#IRQ_acknowledge_.2F_disable_.28.24E000-.24FFFE.2C_even.29}

Writing any value to this register will disable counter interrupts AND acknowledge any pending interrupts.

### <span id="IRQ_enable_.28.24E001-.24FFFF.2C_odd.29"></span>IRQ enable (\$E001-\$FFFF, odd) {#IRQ_enable_.28.24E001-.24FFFF.2C_odd.29}

Writing any value to this register will enable counter interrupts. Note that writing to this register does not acknowledge the IRQ if already set.

## IRQ counter operation {#IRQ_counter_operation}

There are two IRQ modes: PPU A12 mode (also known as *scanline mode*) and CPU cycle mode.

In *scanline mode*, the counter is clocked using a very similar method to that used by the [MMC3](../MMC3/) and follows the same restrictions. In comparison to the [MMC3](../MMC3/), the actual interrupt triggers slightly later. Specifically, it is delayed until [M2 falls twice after the PPU A12 rise](http://forums.nesdev.org/viewtopic.php?p=117323#p117323) that would have triggered the MMC3 interrupt.

In *CPU cycle mode*, the counter is clocked every 4 CPU cycles. The actual interrupt triggers [one M2 cycle later](http://forums.nesdev.org/viewtopic.php?p=117461#p117461) than one would naively expect.

Whichever IRQ mode is being used, the counter behaves the following way:

**When the IRQ is clocked by *scanline* or *CPU cycle* modes:**

- **IF** \$C001 was written to after previous clock:
  - reload IRQ counter with IRQ reload value; if non zero, this value is **ORed with 1** (see notes).
- **ELSE IF** IRQ counter is 0:
  - reload IRQ counter with IRQ reload value.
- **ELSE**
  - Decrement IRQ counter by 1.
- If IRQ counter is now 0 **AND** IRQs are enabled:
  - trigger IRQ after 4 CPU cycles.

**Notes:**

- It's still unknown how the extra kick works. *Klax* still requires **+1** to run perfectly.
- Most emulators run *Skull & Crossbones* with a glitched scanline on the "Continue" screen OR during the game (score bar).

Following these rules in scanline counter mode has demonstrated accurate behavior, but this needs to be refined and merged with the description above:

    On M2 falling edge:
       If PA12 = 0:
           If 4-bit counter < 16: 4-bit counter++.
           else keep the same value.
       If PA12 = 1:
           If 4-bit counter = 16: Clock the 8-bit counter.
           Always reset 4-bit counter to 0.

    On 8-bit counter clock:
       If value = 0:
           If IRQ is enabled, trigger IRQ (wait 1 more cycle before doing that though, regardless of PA12 next cycle...)
           Always reload counter with value in $C000.
       Else value--.

    On Write to $C001 (reset):
       Reload 8-bit counter with value in $C000
       Set 4-bit counter to 17.

## Variants {#Variants}

[Mapper 158](../INES_Mapper_158/), used for *Alien Syndrome*, has mirroring like [mapper 118](../INES_Mapper_118/) ([TLSROM](../TLSROM/)), where CIRAM A10 is connected to CHR A17, and bit 7 of each CHR bank mapped into PPU \$0000-\$0FFF controls which page of CIRAM is used for the corresponding nametable in \$2000-\$2FFF.

## Notes {#Notes}

- Games not using RAMBO-1's extended features (i.e. not setting the K bit, not using the CPU cycle IRQ mode, and with the PRG-ROM bank fixed at \$C000-\$DFFF like *Rolling Thunder*) are mostly compatible with a "standard" MMC3 (as seen in several MMC3-based bootleg multicarts), with minor IRQ-related problems like status bar shaking a little bit due to aforementioned subtle timing differences between RAMBO-1's scanline counter IRQ mode and MMC3.

## See also {#See_also}

- The unlicensed mapper [iNES Mapper 176](../INES_Mapper_176/)'s Extended MMC3 mode functions similar to this mapper, but with different bank register value and has WRAM support.
- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
