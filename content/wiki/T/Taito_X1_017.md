---
title: "Taito X1-017"
url: "/wiki/Taito_X1-017/"
description: "The Taito X1-017 is a ASIC mapper used in a few of their games. It contains 5 KiB of internal WRAM, which can be battery backed-up."
wikiPage: "Taito_X1_017"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>X1-017</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Taito </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=82">5 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>J9100231A, P3-044B, etc </td></tr><tr><td><b>Pinout</b></td><td><a href="../Taito_X1-017_pinout/">Taito X1-017 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>5K </td></tr><tr><td><b>PRG RAM window</b></td><td>2K + 2K + 1K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>2Kx2 + 1Kx4 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">082</a>, <a href="./">552</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TAITO-X1-017">Taito X1-017</a></td></tr></tbody></table></div>

The **Taito X1-017** is a [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/) used in a few of their games. It contains 5 KiB of internal WRAM, which can be battery backed-up.

Comparatively strong pull-downs inside the ASIC mean that there is no functional open bus in games that use this hardware: all reads from locations that would be open bus read back as 0, including disabled RAM and both ASIC and 2A03 registers.

The [cartridge connector](../Cartridge_connector/)'s /IRQ line *is* connected to the mapper IC, but no commercial games used it. It's only recently been reverse-engineered.

[NES 2.0 Mapper 552](./) represents the actual way the mask ROM is connected and is thus the correct bank order, while [iNES Mapper 082](./) represents the bank order as it was understood before January 2020, when the mapper was reverse-engineered.

## Banks {#Banks}

- CPU \$6000-\$73FF: 5 KiB PRG RAM divided into 3 protection regions
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$07FF (or \$1000-\$17FF): 2 KB switchable CHR bank
- PPU \$0800-\$0FFF (or \$1800-\$1FFF): 2 KB switchable CHR bank
- PPU \$1000-\$13FF (or \$0000-\$03FF): 1 KB switchable CHR bank
- PPU \$1400-\$17FF (or \$0400-\$07FF): 1 KB switchable CHR bank
- PPU \$1800-\$1BFF (or \$0800-\$0BFF): 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF (or \$0C00-\$0FFF): 1 KB switchable CHR bank

## Registers {#Registers}

### <span id="CHR_bankings_.28.247EF0-.247EF5.29"></span>CHR bankings (\$7EF0-\$7EF5) {#CHR_bankings_.28.247EF0-.247EF5.29}

    7  bit  0
    ---- ----
    CCCC CCCc
    |||| ||||
    ++++-++++- Set CHR bank

Each 6 registers affect on:

- \$7EF0: Switch PPU \$0000-\$07FF or \$1000-\$17FF
- \$7EF1: Switch PPU \$0800-\$0FFF or \$1800-\$1FFF
- \$7EF2: Switch PPU \$1000-\$13FF or \$0000-\$03FF
- \$7EF3: Switch PPU \$1400-\$17FF or \$0400-\$07FF
- \$7EF4: Switch PPU \$1800-\$1BFF or \$0800-\$0BFF
- \$7EF5: Switch PPU \$1C00-\$1FFF or \$0C00-\$0FFF

Bit 0 of written bank number is ignored when switching 2K banks (\$7EF0, \$7EF1), just like other common mappers like [MMC3](../MMC3/) and/or [Taito X1-005](../Taito_X1-005/), for example.

### <span id="CHR_mode_.2F_Nametable_arrangement_.28.247EF6.29"></span>CHR mode / Nametable arrangement (\$7EF6) {#CHR_mode_.2F_Nametable_arrangement_.28.247EF6.29}

    7  bit  0
    ---- ----
    xxxx xxCM
           ||
           |+- Nametable arrangement (0: vertical (A11); 1: horizontal (A10))
           +-- CHR A12 inversion (0: two 2 KiB banks at $0000-$0FFF,
                                     four 1 KiB banks at $1000-$1FFF;
                                  1: two 2 KiB banks at $1000-$1FFF,
                                     four 1 KiB banks at $0000-$0FFF)

### <span id="PRG_RAM_enable_0_.28.247EF7.29"></span>PRG RAM enable 0 (\$7EF7) {#PRG_RAM_enable_0_.28.247EF7.29}

    7  bit  0
    ---- ----
    RRRR RRRR
    |||| ||||
    ++++-++++- Write $CA to enable 2 KiB internal RAM from $6000 to $67FF, write anything else to disable

### <span id="PRG_RAM_enable_1_.28.247EF8.29"></span>PRG RAM enable 1 (\$7EF8) {#PRG_RAM_enable_1_.28.247EF8.29}

    7  bit  0
    ---- ----
    RRRR RRRR
    |||| ||||
    ++++-++++- Write $69 to enable 2 KiB internal RAM from $6800 to $6FFF, write anything else to disable

### <span id="PRG_RAM_enable_2_.28.247EF9.29"></span>PRG RAM enable 2 (\$7EF9) {#PRG_RAM_enable_2_.28.247EF9.29}

    7  bit  0
    ---- ----
    RRRR RRRR
    |||| ||||
    ++++-++++- Write $84 to enable 1 KiB RAM from $7000 to $73FF, write anything else to disable

### <span id="PRG_Select_0.2F1.2F2_.28.247EFA.2C_.247EFB.2C_.247EFC.29"></span>PRG Select 0/1/2 (\$7EFA, \$7EFB, \$7EFC) {#PRG_Select_0.2F1.2F2_.28.247EFA.2C_.247EFB.2C_.247EFC.29}

Selects the 8 KiB PRG-ROM bank at CPU \$8000 (\$7EFA), \$A000 (\$7EFB) or \$C000 (\$7EFC).

[iNES Mapper 082](./):

    7  bit  0
    ---- ----
    xxDC BAxx
      || |+--- PRG A13
      || +---- PRG A14
      |+------ PRG A15
      +------- PRG A16

[iNES Mapper 082](./) does not support more than 128 KiB of PRG, due to its interpretation of the order of bits in the register. One translation assumes that bit 6 denotes PRG A17, but that bit is not connected to any ASIC pin, and therefore that translation cannot work on real hardware regardless of how the ASIC is connected to ROM.

[NES 2.0 Mapper 552](./):

    7  bit  0
    ---- ----
    xxAB CDxx
      || |||+- PRG A18
      || ||+-- PRG A17
      || |+--- PRG A16
      || +---- PRG A15
      |+------ PRG A14
      +------- PRG A13

In other words, [NES 2.0 Mapper 552](./) reflects the fact that the PRG mask ROM address lines A13-A18 are connected in reverse order.

### <span id="IRQ_Latch_.28.247EFD.29"></span>IRQ Latch (\$7EFD) {#IRQ_Latch_.28.247EFD.29}

The X1-017's IRQ functionality **was not used by the commercial games** and only reverse-engineered in January 2020. No emulators support the IRQ as of January 2020.

It effectively only provides the ability to schedule an IRQ up to 4112 cycles in the future, short enough that even a top status bar will require a "spacing" IRQ, and its 16x prescaler is too coarse for close raster effects.

\$7EFD is an eight-bit wide register, specifying the reload value for the IRQ. (see below)

### <span id="IRQ_Control_.28.247EFE.29"></span>IRQ Control (\$7EFE) {#IRQ_Control_.28.247EFE.29}

    7  bit  0
    ---- ----
    xxxx xMIC
          |||
          ||+- 1: Enable counting
          ||   0: Disable counting and reload counter, ([$7EFD]+2)*16 if [$7EFD] is nonzero; 17 if it's zero.
          |+-- 1: Enable asserting /IRQ, 0: don't.
          +--- Unknown. Counting only works if this bit is 0.

The IRQ will count down while counting is enabled. It's possible to generate multiple IRQs in a row by toggling the "I" bit above before acknowledging the interrupt.

### <span id="IRQ_Acknowledge_and_reload_.28.247EFF.29"></span>IRQ Acknowledge and reload (\$7EFF) {#IRQ_Acknowledge_and_reload_.28.247EFF.29}

Any write to this register will de-assert IRQ and reload the counter.

- if \[\$7EFD\] is nonzero, the value reloaded is (\[\$7EFD\]+1)\*16.
- if \[\$7EFD\] is zero, the value reloaded is 1.

## See Also {#See_Also}

- [Taito X1-005](../Taito_X1-005/)

## References {#References}

- BootGod mentions the RAM protect registers: [\[1\]](https://forums.nesdev.org/viewtopic.php?p=30165#p30165)
- Krzysiobal reverse-engineered the IRQ: [\[2\]](https://forums.nesdev.org/viewtopic.php?p=246602#p246602)
