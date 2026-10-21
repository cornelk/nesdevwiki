---
title: "VRC3"
url: "/wiki/VRC3/"
description: "The Konami VRC3 is an ASIC mapper apparently used only for the game Salamander. The iNES format assigns mapper 73 to VRC3."
wikiPage: "VRC3"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

<div class="infobox"><div style="text-align: center"><b>VRC3</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=73">1 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>350685 </td></tr><tr><td><b>Pinout</b></td><td><a href="../VRC3_pinout/">VRC3 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K, fixed </td></tr><tr><td><b>CHR capacity</b></td><td>8K RAM </td></tr><tr><td><b>CHR window</b></td><td>8K, fixed </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">073</a></td></tr></tbody></table></div>

The Konami VRC3 is an [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/) apparently used only for the game [Salamander](https://nescartdb.com/profile/view/1783/salamander). The [iNES](../INES/) format assigns **mapper 73** to VRC3.

## Banks {#Banks}

- CPU \$6000-\$7FFF: optional 8 KiB PRG RAM bank
- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last bank

## Registers {#Registers}

### <span id="PRG_Select_.28.24F000-.24FFFF.29"></span>PRG Select (\$F000-\$FFFF) {#PRG_Select_.28.24F000-.24FFFF.29}

    7  bit  0
    ---------
    .... .PPP
          |||
          +++- Select 16 KB PRG ROM at $8000

### <span id="IRQ_Latch_0_.28.248000-.248FFF.29"></span>IRQ Latch 0 (\$8000-\$8FFF) {#IRQ_Latch_0_.28.248000-.248FFF.29}

    7  bit  0
    ---------
    .... LLLL
         ||||
         ++++- Bits 0-3 of IRQ latch.

### <span id="IRQ_Latch_1_.28.249000-.249FFF.29"></span>IRQ Latch 1 (\$9000-\$9FFF) {#IRQ_Latch_1_.28.249000-.249FFF.29}

    7  bit  0
    ---------
    .... LLLL
         ||||
         ++++- Bits 4-7 of IRQ latch.

### <span id="IRQ_Latch_2_.28.24A000-.24AFFF.29"></span>IRQ Latch 2 (\$A000-\$AFFF) {#IRQ_Latch_2_.28.24A000-.24AFFF.29}

    7  bit  0
    ---------
    .... LLLL
         ||||
         ++++- Bits 8-11 of IRQ latch.

### <span id="IRQ_Latch_3_.28.24B000-.24BFFF.29"></span>IRQ Latch 3 (\$B000-\$BFFF) {#IRQ_Latch_3_.28.24B000-.24BFFF.29}

    7  bit  0
    ---------
    .... LLLL
         ||||
         ++++- Bits 12-15 of IRQ latch.

### <span id="IRQ_Control_.28.24C000-.24CFFF.29"></span>IRQ Control (\$C000-\$CFFF) {#IRQ_Control_.28.24C000-.24CFFF.29}

    7  bit  0
    ---------
    .... .MEA
          |||
          ||+- IRQ Enable on Acknowledgement (see $D000)
          |+-- IRQ Enable (1=enabled)
          +--- IRQ Mode (1=8-bit counter, 0=16-bit counter)

Any write to this register will acknowledge the pending IRQ. If this register is written to with the **E** bit set, the 16-bit IRQ counter is reloaded with the 16-bit Latch value (set by writes to \$8000-\$BFFF).

### <span id="IRQ_Acknowledge_.28.24D000-.24DFFF.29"></span>IRQ Acknowledge (\$D000-\$DFFF) {#IRQ_Acknowledge_.28.24D000-.24DFFF.29}

    7  bit  0
    ---------
    .... ....

Any write to this register will acknowledge the pending IRQ. In addition, the **A** control bit moves to the **E** control bit, enabling or disabling IRQs. Writes to this register do not affect the current state of the IRQ counter.

## IRQ Counter Operation {#IRQ_Counter_Operation}

VRC3 IRQs operate differently from other VRCs. The counter is 16 bits instead of 8 bits, and there is no scanline mode, only CPU cycle mode. Other aspects, however, are very similar.

IRQs on this mapper are generated through a 16-bit counter, which is incremented each CPU cycle (if enabled). When it overflows from \$FFFF, an IRQ is generated, and the counter is reloaded with the 16-bit latch value.

When the IRQ mode bit (**M**) is set, the upper 8 bits of the counter are ignored (and never incremented) and IRQs are generated when the lower 8 bits overflow from \$FF. When the IRQ is triggered, only the low 8 bits are reloaded from the latch value. Reloading with a write to \$C000 with **E** set will always reload all 16 bits, regardless of mode.

When IRQs are disabled (**E** control bit clear), the IRQ counter does not increment.

Writing to the acknowledge register at \$D000 automatically moves control bit **A** into control bit **E**, allowing the IRQ to be easily re-enabled while acknowledging.

## References {#References}

- [Disch's Mapper Notes](http://www.romhacking.net/documents/362/)
