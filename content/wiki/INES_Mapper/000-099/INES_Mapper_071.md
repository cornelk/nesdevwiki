---
title: "INES Mapper 071"
url: "/wiki/INES_Mapper_071/"
description: "iNES Mapper 071 is assigned to games developed by Codemasters and published by Camerica. It's largely a clone of UNROM, and Camerica games were initially emulated under iNES Mapper 002 before 071 was…"
wikiPage: "INES_Mapper_071"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

**iNES Mapper 071** is assigned to games developed by Codemasters and published by Camerica. It's largely a clone of [UNROM](../UxROM/), and Camerica games were initially emulated under [iNES Mapper 002](../UxROM/) before 071 was assigned.

Example Games:

- *MiG 29 - Soviet Fighter*
- *Fire Hawk*
- *The Fantastic Adventures of Dizzy*
- *Bee 52*

## Overview {#Overview}

- PRG ROM size: 128 or 256 KiB
- PRG ROM bank size: 16 KiB inner; ( 64 KiB outer in [iNES Mapper 232](../INES_Mapper_232/) )
- PRG RAM: None
- CHR capacity: 8 KiB RAM
- CHR bank size: Not bankswitched
- Nametable [mirroring](../Mirroring/): Hardwired H or V (most) or mapper controlled 1-screen (Fire Hawk)
- Subject to [bus conflicts](../Bus_conflict/): No

## Banks {#Banks}

- \$8000-\$BFFF: Switchable
- \$C000-\$FFFF: Fixed to last bank
- \$8000-\$FFFF: In the "Quattro" games, A16 and A17 are controlled by a separate register, allowing 64 KiB super-banks.

## Registers {#Registers}

### <span id="Outer_bank_select_.28.248000-.24BFFF.29_.28Only_on_iNES_Mapper_232.29"></span>Outer bank select (\$8000-\$BFFF) (Only on [iNES Mapper 232](../INES_Mapper_232/)) {#Outer_bank_select_.28.248000-.24BFFF.29_.28Only_on_iNES_Mapper_232.29}

This register is present only on the Quattro carts, which use the [BF9096 IC](../Camerica_BF909x_pinout/), and on the CCU in the Aladdin adapter (per [KH](http://kevtris.org/mappers/camerica/aladdin.html))

    7  bit  0
    ---- ----
    xxxB Bxxx
       | |
       +-+--- Select 64 KiB PRG ROM bank for CPU $8000-$FFFF

### <span id="Mirroring_.28.248000-.249FFF.29"></span>Mirroring (\$8000-\$9FFF) {#Mirroring_.28.248000-.249FFF.29}

This register is present only on *Fire Hawk*, which uses the [BF9097 IC](../Camerica_BF909x_pinout/). [A submapper](../NES_2.0_submappers/#071:_1_Fire_Hawk) has been allocated for this difference.

*Fire Hawk* only writes this register at the address \$9000, and other games like *Micro Machines* and *Ultimate Stuntman* write \$00 to \$8000 on startup. For compatibility without using a submapper, FCEUX begins all games with fixed mirroring, and applies single screen mirroring only once \$9000-9FFF is written, ignoring writes to \$8000-8FFF.

    7  bit  0
    ---- ----
    xxxM xxxx
       |
       +----- Select 1 KiB CIRAM bank for PPU $2000-$2FFF

### <span id="Bank_select_.28.24C000-.24FFFF.29"></span>Bank select (\$C000-\$FFFF) {#Bank_select_.28.24C000-.24FFFF.29}

This register is present on all games: [BF9093](../Camerica_BF909x_pinout/) (most common), BF9096, BF9097, and CCU. The number of bits available vary: 4 for the BF9093, 3 for the BF9097, and 2 for the BF9096.

    7  bit  0
    ---- ----
    xxxx PPPP
         ||||
         ++++- Select 16 KiB PRG ROM bank for CPU $8000-$BFFF

### <span id="CIC_stun_control_.28.24E000-.24FFFF.29"></span>CIC stun control (\$E000-\$FFFF) {#CIC_stun_control_.28.24E000-.24FFFF.29}

This register is present on all games. A0 controls a latch going to the [CIC stun circuit](http://kevtris.org/mappers/lockout/camerica.html).

## Notes {#Notes}

This mapper covers several Camerica/Codemasters boards. Most boards have fixed vertical or horizontal mirroring, but the *Fire Hawk* board differs in that it has mapper controlled 1-screen mirroring. These variants can be distinguished by [NES 2.0 submappers](../NES_2.0_submappers/).

Some of these games are like *Battletoads* in that they're very [picky](../Tricky-to-emulate_games/) about timing and use some seldom used aspects of the NES, despite the mapper itself being as simple as the discrete logic mappers.

In particular:

- *Bee 52* uses the sprite overflow flag (\$2002.5)
- *MiG 29* uses DMC IRQs, and is VERY PICKY about their timing. If your DMC IRQ timing isn't spot on (or at least very close), this game will glitch a lot.

This mapper also involves a custom lockout defeat circuit which is mostly unimportant for emulation purposes. Details will not be mentioned here, but are outlined in Kevtris' Camerica Mappers documentation.

*Fire Hawk* does some strange timing code when changing the mirroring mode. It is unknown whether or not any special timing is required.

## See also {#See_also}

- [Camerica Mappers by Kevin Horton](http://kevtris.org/mappers/camerica/index.html)
- [NES Mapper list](http://www.romhacking.net/documents/362/) by Disch.
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ Information on mapper's initial state is innacurate.
