---
title: "Taito X1-005"
url: "/wiki/Taito_X1-005/"
description: "The Taito X1-005 is a ASIC mapper used in a few of their games. This mapper also contains 128 bytes of mapper-internal WRAM and can be battery backed-up. Board wriring variant for single screen…"
wikiPage: "Taito_X1_005"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>X1-005</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Taito </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=80">8 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>J9100240A, P3-034A, etc. </td></tr><tr><td><b>Pinout</b></td><td><a href="../Taito_X1-005_pinout/">Taito X1-005 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>128B </td></tr><tr><td><b>PRG RAM window</b></td><td>128B </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>2Kx2 + 1Kx4 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">080</a>, <a href="../INES_Mapper_207/">207</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TAITO-X1-005">Taito X1-005</a></td></tr></tbody></table></div>

The **Taito X1-005** is a [ASIC mapper](../../category/ASIC_mappers/) used in a few of their games. This mapper also contains 128 bytes of mapper-internal WRAM and can be battery backed-up. Board wriring variant for single screen nametable arrangement exists and is allocated at [iNES Mapper 207](../INES_Mapper_207/).

## Memory map {#Memory_map}

- CPU \$7F00-\$7F7F: 128 B internal RAM
- CPU \$7F80-\$7FFF: Mirror of \$7F00-\$7F7F
- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG-ROM bank
- CPU \$C000-\$DFFF: 8 KiB switahable PRG-ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG-ROM bank, fixed to the last bank
- PPU \$0000-\$07FF: 2 KiB switchable CHR bank
- PPU \$0800-\$0FFF: 2 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank

## Registers {#Registers}

The **Taito X1-005** provides total 11 registers for PRG bankings, CHR bankings, mirroring setup and RAM enable. All registers exist at \$7EFx and mirrored at \$7E7x because the CPU A7 line is ignored. All registers are write-only.

### <span id="CHR_bankings_.28.247EF0-.247EF5.29"></span>CHR bankings (\$7EF0-\$7EF5) {#CHR_bankings_.28.247EF0-.247EF5.29}

    7  bit  0
    ---- ----
    CCCC CCCc
    |||| ||||
    ++++-++++- Set CHR bank

Each 6 registers affect on:

- \$7EF0: Switch bank at PPU \$0000-\$07FF
- \$7EF1: Switch bank at PPU \$0800-\$0FFF
- \$7EF2: Switch bank at PPU \$1000-\$13FF
- \$7EF3: Switch bank at PPU \$1400-\$17FF
- \$7EF4: Switch bank at PPU \$1800-\$1BFF
- \$7EF5: Switch bank at PPU \$1C00-\$1FFF

Bit 0 of written bank number is ignored when switching 2K banks (\$7EF0, \$7EF1), just like other common mappers like [MMC3](../MMC3/), for example.

### <span id="Nametable_arrangement_.28.247EF6.29"></span>Nametable arrangement (\$7EF6) {#Nametable_arrangement_.28.247EF6.29}

    7  bit  0
    ---- ----
    .... ...M
            |
            +- Nametable arrangement (0: vertical (A11); 1: horizontal (A10))

This register is **not mirrored at \$7EF7**. This register only works in [iNES Mapper 080](./). [iNES Mapper 207](../INES_Mapper_207/) board handles mirroring differently and this register is ignored.

### <span id="Internal_RAM_permission_.28.247EF8.29"></span>Internal RAM permission (\$7EF8) {#Internal_RAM_permission_.28.247EF8.29}

    7  bit  0
    ---- ----
    RRRR RRRR
    |||| ||||
    ++++-++++- Enable read and write to internal RAM only when written value of this register is $A3

Writing \$A3 here enables reading and writing to internal RAM, other values disable.

This register is mirrored at \$7EF9. Writing to \$7EF9 should make the same result as writing to \$7EF8.

### <span id="PRG_bankings_.28.247EFA.2C_.247EFC.2C_.247EFE.29"></span>PRG bankings (\$7EFA, \$7EFC, \$7EFE) {#PRG_bankings_.28.247EFA.2C_.247EFC.2C_.247EFE.29}

    7  bit  0
    ---- ----
    ..PP PPPP
      || ||||
      ++-++++- Set PRG bank

Each 3 registers affect on:

- \$7EFA: Switch bank at CPU \$8000-\$9FFF
- \$7EFC: Switch bank at CPU \$A000-\$BFFF
- \$7EFE: Switch bank at CPU \$C000-\$DFFF

In addition, they are mirrored at right next respective addresses, i.e, \$7EFB is the mirror of \$7EFA, \$7EFD is \$7EFC's, so on.

## Notes {#Notes}

- Unknown-purpose \$7EF7 register
  - Currently unknown-purpose \$7EF7 register doesn't seem to doanything with the /RAMSEL of the [chip](../Taito_X1-005_pinout/). Did not affect any behaviors even when tested with [\$7EF8](#Internal_RAM_permission_.28.247EF8.29).
  - Reading from it returns [open-bus](../Open_bus_behavior/).
