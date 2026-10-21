---
title: "INES Mapper 118"
url: "/wiki/INES_Mapper_118/"
description: "TxSROM is used to designate TKSROM and TLSROM boards, both of which use the Nintendo MMC3 in a nonstandard way. The only known difference between these boards and TKROM and TLROM is the mirroring…"
wikiPage: "INES_Mapper_118"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

**TxSROM** is used to designate [TKSROM](../TKSROM/) and [TLSROM](../TLSROM/) boards, both of which use the [Nintendo MMC3](../MMC3/) in a nonstandard way. The only known difference between these boards and [TKROM](../TxROM/) and [TLROM](../TxROM/) is the mirroring configuration. The CHR A17 line connects directly to CIRAM A10 line instead of MMC3's CIRAM A10 output, to compensate for the MMC3's lack of single-screen [mirroring](../Mirroring/). The [iNES](../INES/) format assigns **iNES Mapper 118** to these boards.

Example games:

- *Armadillo*
- *Goal! Two* (US/Europe Version)
- *NES Play Action Football*
- *Pro Sport Hockey*

## Registers {#Registers}

The behavior of these boards differs from that of a typical MMC3 board in the use of the upper CHR address line. This board relies on the fact that the MMC3's CHR bank circuit ignores A13 when calculating CHR A10-A17, responding to [nametable](../PPU_nametables/) fetches from \$2000-\$2FFF the same way as fetches from the first pattern table at \$0000-\$0FFF. This means that the 1KB/2KB banking scheme used for CHR bankswitching is active even during nametable fetches while the CHR ROM/RAM is disabled.

However, on these particular boards, the CHR bankswitching directly affects the [mirroring](../Mirroring/) mapping the nametable RAM. This allows programs to select which nametable is mapped to each slot, much like CHR banks are mapped to pattern table slots, in either two 2KB banks (allowing only single-screen or horizontal mirroring) or four 1KB banks (allowing all mirroring modes one can think of, because this is equal to the size of a nametable) at the price of mapping the 1KB CHR banks to the first pattern table by setting bit 7 of \$8000. If the IRQ counter is being used in a standard way, this involves having sprites bankswitched in 2KB banks and backgrounds in 1KB banks.

### <span id="Bank_data_.28.248001-.249FFF.2C_odd.29"></span>Bank data (\$8001-\$9FFF, odd) {#Bank_data_.28.248001-.249FFF.2C_odd.29}

    7  bit  0
    ---- ----
    MDDD DDDD
    |||| ||||
    |+++-++++- New bank value, based on last value written to Bank select register
    |          0: Select 2 KB CHR bank at PPU $0000-$07FF (or $1000-$17FF);
    |          1: Select 2 KB CHR bank at PPU $0800-$0FFF (or $1800-$1FFF);
    |          2: Select 1 KB CHR bank at PPU $1000-$13FF (or $0000-$03FF);
    |          3: Select 1 KB CHR bank at PPU $1400-$17FF (or $0400-$07FF);
    |          4: Select 1 KB CHR bank at PPU $1800-$1BFF (or $0800-$0BFF);
    |          5: Select 1 KB CHR bank at PPU $1C00-$1FFF (or $0C00-$0FFF);
    |          6, 7: as standard MMC3
    |
    +--------- Mirroring configuration, based on the last value
               written to Bank select register
               0: Select Nametable at PPU $2000-$27FF
               1: Select Nametable at PPU $2800-$2FFF
               Note : Those bits are ignored if corresponding CHR banks
               are mapped at $1000-$1FFF via $8000.

               2 : Select Nametable at PPU $2000-$23FF
               3 : Select Nametable at PPU $2400-$27FF
               4 : Select Nametable at PPU $2800-$2BFF
               5 : Select Nametable at PPU $2C00-$2FFF
               Note : Those bits are ignored if corresponding CHR banks
               are mapped at $1000-$1FFF via $8000.

### <span id="Mirroring_.28.24A000-.24BFFE.2C_even.29"></span>Mirroring (\$A000-\$BFFE, even) {#Mirroring_.28.24A000-.24BFFE.2C_even.29}

    7  bit  0
    ---- ----
    xxxx xxxM
            |
            +- Mirroring
               This bit is bypassed by the configuration described above, so writing here has no effect.

Note: In theory, the CHR limitation is 256 KB like all [MMC3](../MMC3/) boards. But because CHR A17 has another usage, having a CHR greater than 128 KB would require very careful CHR ROM layout because CHR bankswitching and mirroring will be linked through the same selection bits. Probably for this reason, official Nintendo TLSROM boards doesn't allow for 256 KB CHR-ROMs. However, a mapper 118 game that uses a third party MMC3/board, using 1-screen mirroring could draw the playfield with the lower 128 KB of CHR ROM and the lower nametable, and draw the status bar and menus with the upper 128 KB of CHR ROM and the upper nametable. Sprite tile banks could go in whatever space remains in either or both halves.

## Variants {#Variants}

Mappers [158](../INES_Mapper_158/), [207](../INES_Mapper_207/) and [527](../NES_2.0_Mapper_527/) do the exact same thing but with Tengen's [RAMBO-1](../RAMBO-1/), Taito's [X1-005](../Taito_X1-005/) and Konami's [VRC2](../VRC2_and_VRC4/), respectively.

[Mapper 95](../INES_Mapper_095/) is almost the same thing, but with the MMC3 predecessor [Namco 108](../INES_Mapper_206/) instead of an actual MMC3, and with the nametable select bit on A15 instead of A17 because the Namco 108 family can only address six CHR bank bits.

## References {#References}

- [NES Mapper list](http://www.romhacking.net/documents/362/) by Disch.
