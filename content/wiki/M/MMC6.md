---
title: "MMC6"
url: "/wiki/MMC6/"
description: "The Nintendo MMC6 is a mapper ASIC used in Nintendo's NES-HKROM Game Pak board. This board, along with most common TxROM boards (which use the Nintendo MMC3) are assigned to iNES Mapper 004. The…"
wikiPage: "MMC6"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>MMC6</b><br/><b>HxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=4">2 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>HKROM </td></tr><tr><td><b>Pinout</b></td><td><a href="../MMC6_pinout/">MMC6 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>1K </td></tr><tr><td><b>PRG RAM window</b></td><td>1K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1K + 2K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="../MMC3/">004</a></td></tr></tbody></table></div>

The **Nintendo MMC6** is a [mapper](../Mapper/) [ASIC](../../category/ASIC_mappers/) used in Nintendo's NES-HKROM Game Pak board. This board, along with most common [TxROM](../TxROM/) boards (which use the [Nintendo MMC3](../MMC3/)) are assigned to [iNES Mapper 004](../MMC3/). The MMC3C and the MMC6 are alike, except the MMC6 has 1 KB of internal PRG RAM with different write/enable controls. This page only explains the differences, see [MMC3](../MMC3/) for full details on the rest of the mapper.

The [NES 2.0 submapper 004:1](../NES_2.0_submappers/#004:_1_MMC6) was assigned to disambiguate MMC6 from the MMC3 mapper it shares an iNES mapper with.

This chip first appeared in December 1990. A total of two commercial games were released using the MMC6:

- StarTropics
- Zoda's Revenge: StarTropics II

## Banks {#Banks}

- CPU \$7000-\$7FFF: 1 KB PRG RAM, [mirrored](../Mirroring/)
- CPU \$8000-\$9FFF (or \$C000-\$DFFF): 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF (or \$8000-\$9FFF): 8 KB PRG ROM bank, fixed to the second-last bank
- CPU \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$07FF (or \$1000-\$17FF): 2 KB switchable CHR bank
- PPU \$0800-\$0FFF (or \$1800-\$1FFF): 2 KB switchable CHR bank
- PPU \$1000-\$13FF (or \$0000-\$03FF): 1 KB switchable CHR bank
- PPU \$1400-\$17FF (or \$0400-\$07FF): 1 KB switchable CHR bank
- PPU \$1800-\$1BFF (or \$0800-\$0BFF): 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF (or \$0C00-\$0FFF): 1 KB switchable CHR bank

## Registers {#Registers}

The MMC6 has 4 pairs of registers at \$8000-\$9FFF, \$A000-\$BFFF, \$C000-\$DFFF, and \$E000-\$FFFF - even addresses (\$8000, \$8002, etc.) select the low register and odd addresses (\$8001, \$8003, etc.) select the high register in each pair. Only \$8000 and \$A001 are covered here. For the rest of the registers, see [MMC3](../MMC3/).

### <span id="Bank_select_.28.248000-.249FFE.2C_even.29"></span>Bank select (\$8000-\$9FFE, even) {#Bank_select_.28.248000-.249FFE.2C_even.29}

    7  bit  0
    ---- ----
    CPMx xRRR
    |||   |||
    |||   +++- Specify which bank register to update on next write to Bank Data register
    ||+------- PRG RAM enable
    |+-------- PRG ROM bank configuration (0: $8000-$9FFF swappable, $C000-$DFFF fixed to second-last bank;
    |                                      1: $C000-$DFFF swappable, $8000-$9FFF fixed to second-last bank)
    +--------- CHR ROM bank configuration (0: two 2 KB banks at $0000-$0FFF, four 1 KB banks at $1000-$1FFF;
                                           1: four 1 KB banks at $0000-$0FFF, two 2 KB banks at $1000-$1FFF)

### <span id="PRG_RAM_protect_.28.24A001-.24BFFF.2C_odd.29"></span>PRG RAM protect (\$A001-\$BFFF, odd) {#PRG_RAM_protect_.28.24A001-.24BFFF.2C_odd.29}

    7  bit  0
    ---- ----
    HhLl xxxx
    ||||
    |||+------ Enable writing to RAM at $7000-$71FF
    ||+------- Enable reading RAM at $7000-$71FF
    |+-------- Enable writing to RAM at $7200-$73FF
    +--------- Enable reading RAM at $7200-$73FF

## Hardware {#Hardware}

The MMC6 exists in a [64-pin TQFP package](../MMC6_pinout/). At least two revisions exist<sup>\[***citation needed***\]</sup>, though only MMC6B has been observed.

The PRG RAM protect bits control mapping of two 512B banks of RAM. If neither bank is enabled for reading, the \$7000-\$7FFF area is [open bus](../Open_bus_behavior/). If only one bank is enabled for reading, the other reads back as zero. The write-enable bits only have effect if that bank is enabled for reading, otherwise the bank is not writable. Both banks may be enabled for reading (and optionally writing) at the same time.

When PRG RAM is disabled via \$8000, the mapper continuously sets \$A001 to \$00, and so all writes to \$A001 are ignored.
