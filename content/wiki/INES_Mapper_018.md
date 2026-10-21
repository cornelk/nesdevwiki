---
title: "INES Mapper 018"
url: "/wiki/INES_Mapper_018/"
description: "iNES Mapper 018 represents the Jaleco SS 88006 mapper used for Magic John (Japanese version of Totally Rad) and about a dozen other games."
wikiPage: "INES_Mapper_018"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

<div class="infobox"><div style="text-align: center"><b>SS 88006</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Jaleco </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=18">15 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>JF-24, others </td></tr><tr><td><b>Pinout</b></td><td><a href="../Jaleco_SS_88006_pinout/">Jaleco SS 88006 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K + 8K + 8K + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>018 </td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=18">iNES 018</a></td></tr></tbody></table></div>

**iNES Mapper 018** represents the [Jaleco SS 88006 mapper](../Jaleco_SS_88006_pinout/) used for *Magic John* (Japanese version of *Totally Rad*) and about a dozen other games.

This mapper is connected only to A12-A14, A0-A1, and D0-D3, so the PRG bank and CHR bank numbers are split over two sequential addresses. Bits 0-3 are in the lower address of a pair (e.g. \$xxx0 or \$xxx2), and bits 4-7 are in the next higher address.

This mapper resembles a scrambled [VRC4](../VRC2_and_VRC4/).

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KB PRG RAM bank (optional)
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank

## Registers {#Registers}

### <span id="PRG_Select_0_low.28.248000.29.2C_high.28.248001.29"></span>PRG Select 0 low(\$8000), high(\$8001) {#PRG_Select_0_low.28.248000.29.2C_high.28.248001.29}

      $8000        $8001
    7  bit  0    7  bit  0
    ---------    ---------
    .... LLLL    .... ..HH
         ||||           ||
         ||||           ++- High 2 bits of 8 KiB PRG bank at CPU $8000
         ++++-------------- Low 4 bits

### <span id="PRG_Select_1_and_2_.28.248002-.249FFD.29"></span>PRG Select 1 and 2 (\$8002-\$9FFD) {#PRG_Select_1_and_2_.28.248002-.249FFD.29}

The other two PRG bank selects continue similarly:

<table class="wikitable">
<thead>
<tr>
<th colspan="2">Write to CPU address</th>
<th>8KB PRG bank affected</th>
</tr>
</thead>
<tbody>
<tr>
<th>(low 4 bits)</th>
<th>(high 2 bits)</th>
<th></th>
</tr>
&#10;<tr>
<td>$8002</td>
<td>$8003</td>
<td>$A000-$BFFF</td>
</tr>
<tr>
<td>$9000</td>
<td>$9001</td>
<td>$C000-$DFFF</td>
</tr>
</tbody>
</table>

### <span id="PRG_RAM_protect_.28.249002.29"></span>PRG RAM protect (\$9002) {#PRG_RAM_protect_.28.249002.29}

    7  bit  0
    ---------
    .... ..WR
           ||
           |+- PRG RAM chip enable (0: disable; 1: enable) (same as MMC3)
           +-- Write protection (0: deny writes; 1: allow writes) (opposite MMC3)

### <span id="CHR_Select_0_low.28.24A000.29.2C_high.28.24A001.29"></span>CHR Select 0 low(\$A000), high(\$A001) {#CHR_Select_0_low.28.24A000.29.2C_high.28.24A001.29}

      $A000        $A001
    7  bit  0    7  bit  0
    ---------    ---------
    .... LLLL    .... HHHH
         ||||         ||||
         ||||         ++++- High 4 bits of 1 KiB CHR bank at PPU $0000
         ++++-------------- Low 4 bits

### <span id="CHR_Selects_1.E2.80.A67_.28.24A002-.24DFFF.29"></span>CHR Selects 1…7 (\$A002-\$DFFF) {#CHR_Selects_1.E2.80.A67_.28.24A002-.24DFFF.29}

The other seven CHR bank selects continue similarly:

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
<th>(high 4 bits)</th>
<th></th>
</tr>
&#10;<tr>
<td>$A002</td>
<td>$A003</td>
<td>$0400-$07FF</td>
</tr>
<tr>
<td>$B000</td>
<td>$B001</td>
<td>$0800-$0BFF</td>
</tr>
<tr>
<td>$B002</td>
<td>$B003</td>
<td>$0C00-$0FFF</td>
</tr>
<tr>
<td>$C000</td>
<td>$C001</td>
<td>$1000-$13FF</td>
</tr>
<tr>
<td>$C002</td>
<td>$C003</td>
<td>$1400-$17FF</td>
</tr>
<tr>
<td>$D000</td>
<td>$D001</td>
<td>$1800-$1BFF</td>
</tr>
<tr>
<td>$D002</td>
<td>$D003</td>
<td>$1C00-$1FFF</td>
</tr>
</tbody>
</table>

### <span id="IRQ_reload_value_.28.24E000-.24EFFF.29"></span>IRQ reload value (\$E000-\$EFFF) {#IRQ_reload_value_.28.24E000-.24EFFF.29}

The four registers here represent a 16-bit reload value, split into four four-bit numbers, least significant four bits first.

### <span id="IRQ_reload_.28.24F000-.24FFFC.29"></span>IRQ reload (\$F000-\$FFFC) {#IRQ_reload_.28.24F000-.24FFFC.29}

Any write to this register will immediately reload the IRQ counter from the above reload value and acknowledge the IRQ.

### <span id="IRQ_counter_size_.28.24F001-.24FFFD.29"></span>IRQ counter size (\$F001-\$FFFD) {#IRQ_counter_size_.28.24F001-.24FFFD.29}

    7  bit  0
    ---------
    .... FETC
         ||||
         |||+- 1: Enable counting 
         ||+-- 1: Don't propagate counter borrow to bit 12; instead assert IRQ
         |+--- 1: Don't propagate counter borrow to bit 8; instead assert IRQ
         +---- 1: Don't propagate counter borrow to bit 4; instead assert IRQ

F overrides E overrides T. If none are set, the counter is 16 bits wide.

Writes to this register also acknowledge the IRQ.

When enabled, the counter counts down.

### <span id="Mirroring_Control_.28.24F002-.24FFFE.29"></span>Mirroring Control (\$F002-\$FFFE) {#Mirroring_Control_.28.24F002-.24FFFE.29}

    7  bit  0
    ---------
    .... ..MM
           ||
           ++-- 0: Horizontal (A11)
                1: Vertical (A10)
                2: 1scA (Ground)
                3: 1scB (Vcc)

### <span id="Expansion_sound_.28.24F003-.24FFFF.29"></span>Expansion sound (\$F003-\$FFFF) {#Expansion_sound_.28.24F003-.24FFFF.29}

Some games also have a µPD7755C or µPD7756C ADPCM sound IC. If so,

    7  bit  0
    ---------
    NNNN NNRS
    |||| ||||
    |||| |||+-- 1: Assert (drive low) RESET on the ADPCM sound IC. 0: deassert (drive high)
    |||| ||+--- 1: Assert START on the ADPCM sound IC. 0: deassert
    ++++-++---- Which sample to start, if START becomes asserted while RESET is deasserted.

This is the same audio playback IC used in other Jaleco mappers, with the same timing constraints. RESET must be asserted for at least 34 cycles to be assured of stopping playback. START must be de-asserted before the sample finishes playback, or else a random sample (depending on the current contents of the data bus) will start playing. START will be ignored until 23 cycles after RESET is deasserted.

## <span id="Disch.27s_Notes"></span>Disch's Notes {#Disch.27s_Notes}

     Here are Disch's original notes:
     ========================
     =  Mapper 018          =
     ========================

     Example Games:
     --------------------------
     The Lord of King
     Magic John
     Pizza Pop

     Registers:
     ---------------------------

     Range,Mask:   $8000-FFFF, $F003

       $800x,$900x:  [.... PPPP]  PRG Regs
       $A00x-$D00x:  [.... CCCC]  CHR Regs
       $E00x:  [.... IIII]  IRQ Reload value
       $F000:  [.... ....]  IRQ Reset
       $F001:  [.... SSSE]  IRQ Control
          S = Size of IRQ counter
          E = Enable

       $F002:  [.... ..MM]  Mirroring
          %00 = Horz
          %01 = Vert
          %10 = 1ScA
          %11 = 1ScB

     CHR Setup:
     ---------------------------

     Only low 4 bits of written value significant [.... CCCC]

     2 regs combined to get an 8-bit page number

     $x000 or $x002 are the low 4 bits
     $x001 or $x003 are the high 4 bits

           $0000   $0400   $0800   $0C00   $1000   $1400   $1800   $1C00 
         +-------+-------+-------+-------+-------+-------+-------+-------+
         |$A000+1|$A002+3|$B000+1|$B002+3|$C000+1|$C002+3|$D000+1|$D002+3|
         +-------+-------+-------+-------+-------+-------+-------+-------+

     PRG Setup:
     ---------------------------

     Same as CHR, $x000 low, $x001 high

           $8000   $A000   $C000   $E000 
         +-------+-------+-------+-------+
         |$8000+1|$8002+3|$9000+1| { -1} |
         +-------+-------+-------+-------+

     IRQ:
     ---------------------------

     16-bit IRQ Reload value is set via regs $E00x.  $E000 sets the low 4 bits, $E003 sets the high 4 bits.

     When enabled, the IRQ counter counts down every CPU cycle.  When it wraps, an IRQ is generated.

     The 'S' bits in the control reg determine the size of the IRQ counter.  It can be 4, 8, 12, or 16 bits wide:

       %000 = 16 bits wide
       %001 = 12 bits wide
       %01x = 8 bits wide
       %1xx = 4 bits wide

     If the counter is less than 16 bits, the high bits are not altered by IRQ counter clocking; they retain their
     value.

     Example:  if the IRQ counter contains $1232, and is in 4-bit mode, it counts like so:

       $1232
       $1231
       $1230
       $123F  <--- IRQ here
       $123E
        ...

     Any write to the reset reg ($F000) will copy the 16-bit reload value into the IRQ counter (full 16 bits are
     copied, regardless of current 'S' value).

     Any write to $F000 or $F001 will acknowledge the IRQ.

## See also {#See_also}

- [NES mapper list](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
- Sound IC D7756 [description](http://www.cpu-world.com/Support/7/77.html) and [datasheet](http://bootgod.dyndns.org:7777/downloads/UPD7755_UPD7756_UPD7757_UPD7758.pdf)
- Sound IC BBS topics [762](http://forums.nesdev.org/viewtopic.php?t=762) and [4030](http://forums.nesdev.org/viewtopic.php?p=32572#p32572)
