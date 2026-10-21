---
title: "INES Mapper 119"
url: "/wiki/INES_Mapper_119/"
description: "iNES Mapper 119 is used to designate the TQROM board, which uses the Nintendo MMC3 like other TxROM boards but uses the CHR bank number in such a way so as to allow both CHR ROM and CHR RAM to be…"
wikiPage: "INES_Mapper_119"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

iNES Mapper 119 is used to designate the **TQROM** board, which uses the [Nintendo MMC3](../MMC3/) like other [TxROM](../TxROM/) boards but uses the CHR bank number in such a way so as to allow both CHR ROM and CHR RAM to be used simultaneously.

Example games:

- *High Speed*
- *Pin-Bot*

## Registers {#Registers}

Other than the CHR bank register described below, see [MMC3](../MMC3/) for a full description.

### <span id="Bank_data_.28.248001-.249FFF.2C_odd.29"></span>Bank data (\$8001-\$9FFF, odd) {#Bank_data_.28.248001-.249FFF.2C_odd.29}

Bit 6 of the CHR bank number is used to switch between the CHR ROM chip and the CHR RAM chip.

    7  bit  0
    ---- ----
    xCDD DDDD
     ||| ||||
     |++-++++- New bank value, based on last value written to Bank select register
     |         0: Select 2 KB CHR bank at PPU $0000-$07FF (or $1000-$17FF);
     |         1: Select 2 KB CHR bank at PPU $0800-$0FFF (or $1800-$1FFF);
     |         2: Select 1 KB CHR bank at PPU $1000-$13FF (or $0000-$03FF);
     |         3: Select 1 KB CHR bank at PPU $1400-$17FF (or $0400-$07FF);
     |         4: Select 1 KB CHR bank at PPU $1800-$1BFF (or $0800-$0BFF);
     |         5: Select 1 KB CHR bank at PPU $1C00-$1FFF (or $0C00-$0FFF);
     |         6, 7: As standard MMC3
     +-------- Chip select (for CHR banks)
               0: Select CHR ROM; 1: Select CHR RAM

Nintendo *could* have used bit 7 instead of bit 6, allowing 128 KiB of CHR ROM instead of 64 KiB. However, using bit 7 would have precluded a hypothetical "TQSROM" board combining TQROM-style ROM/RAM mixing with [TLSROM](../TLSROM/)-style single-screen mirroring. [Rare](../Rare/) developed both games on this board, and also frequently used single-screen mirroring on the [AxROM](../AxROM/) boards.

## Hardware {#Hardware}

A [74HC32](../7432/) IC controls the CHR ROM's enable line. PPU A13 is ORed with CHR A16 in order to create a /CE signal for the CHR ROM. CHR A16 also directly connects the positive CE pin on the [6264](../6264_static_RAM/) SRAM used as CHR RAM.

For some totally unknown reason (board space?) official TQROM boards doesn't allow PRG ROMs greater than 128 KB. However, a third party board could expand this mapper up to 512 KB of PRG ROM, like any other MMC3 boards, and split a 128 KiB CHR ROM between banks 0-63 and banks 128-191.

Both TQROM games used an 8 KiB CHR RAM.

## See also {#See_also}

- [MMC3 with CHR ROM and CHR RAM](../MMC3_with_CHR_ROM_and_CHR_RAM/)

## References {#References}

- [NES-TQROM](http://kevtris.org/mappers/mmc3/NES_TQROM.html) by Kevin Horton
- [Disch's Mapper Notes](http://www.romhacking.net/documents/362/)
