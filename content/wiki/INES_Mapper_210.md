---
title: "INES Mapper 210"
url: "/wiki/INES_Mapper_210/"
description: "iNES Mapper 210 is used to denote Famicom boards bearing either the Namco 175 (Submapper 1) or Namco 340 (Submapper 2) ASIC. Both can be thought of as cost-reduced versions of the Namco 163:"
wikiPage: "INES_Mapper_210"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>Namco 175</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Namco </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=210">4 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>110F0B </td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V fixed </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>210 </td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>Namco 340</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Namco </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=210">8 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>CS000x </td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>none </td></tr><tr><td><b>PRG RAM window</b></td><td>none </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1 switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>210 </td></tr></tbody></table></div>

iNES Mapper 210 is used to denote Famicom boards bearing either the Namco 175 (Submapper 1) or Namco 340 (Submapper 2) ASIC. Both can be thought of as cost-reduced versions of the [Namco 163](../INES_Mapper_019/):

| Function          | N163                  | N175          | N340               |
|-------------------|-----------------------|---------------|--------------------|
| IRQ counter       | Yes                   | No            | No                 |
| ROM nametables    | Yes                   | No            | No                 |
| Optional WRAM     | Yes                   | Yes           | No                 |
| Expansion audio   | Yes                   | No            | No                 |
| ASIC-internal RAM | Yes                   | No            | No                 |
| Mirroring         | Selectable (extended) | Hardwired H/V | Selectable H/V/0/1 |

## Game list {#Game_list}

### <span id="Namco_175_.28Submapper_1.29"></span>Namco 175 (Submapper 1) {#Namco_175_.28Submapper_1.29}

| Name                                                | Battery | WRAM  | Mirroring |
|-----------------------------------------------------|---------|-------|-----------|
| Chibi Maruko-chan: Uki Uki Shopping                 | No      | No    | V         |
| Family Circuit '91                                  | Yes     | 2 KiB | V         |
| Famista '91                                         | No      | No    | V         |
| Heisei Tensai Bakabon                               | No      | No    | V         |
| Splatterworld: Rick to Kyoufu no Daiou (unreleased) | Yes     | 8 KiB | V         |

### <span id="Namco_340_.28Submapper_2.29"></span>Namco 340 (Submapper 2) {#Namco_340_.28Submapper_2.29}

- Dream Master
- Famista '92
- Famista '93
- Famista '94
- Splatterhouse: Wanpaku Graffiti
- Top Striker
- Wagyan Land 2
- Wagyan Land 3

## Banks {#Banks}

- CPU \$6000-\$7FFF: PRG RAM bank, mirrored if smaller than 8 KB (Namco 175 only)
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KB switchable CHR bank
- PPU \$0400-\$07FF: 1 KB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KB switchable CHR bank
- PPU \$1000-\$13FF: 1 KB switchable CHR bank
- PPU \$1400-\$17FF: 1 KB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KB switchable CHR bank

## Registers {#Registers}

The 175 and 340 have 12 registers within \$8000-\$FFFF. Each register occupies a range of \$800 bytes, so \$8000-\$87FF all refers to one register, \$8800-\$8FFF all refers to another register, and so on.

### <span id="CHR_Select_.28.248000-.24BFFF.29_w"></span>CHR Select (\$8000-\$BFFF) w {#CHR_Select_.28.248000-.24BFFF.29_w}

| Value CPU writes | Behavior                    |
|------------------|-----------------------------|
| \$00-\$FF        | Selects 1KB page of CHR-ROM |

| Write to CPU address | 1KB CHR bank affected |
|----------------------|-----------------------|
| \$8000-\$87FF        | \$0000-\$03FF         |
| \$8800-\$8FFF        | \$0400-\$07FF         |
| \$9000-\$97FF        | \$0800-\$0BFF         |
| \$9800-\$9FFF        | \$0C00-\$0FFF         |
| \$A000-\$A7FF        | \$1000-\$13FF         |
| \$A800-\$AFFF        | \$1400-\$17FF         |
| \$B000-\$B7FF        | \$1800-\$1BFF         |
| \$B800-\$BFFF        | \$1C00-\$1FFF         |

### <span id="External_PRG_RAM_enable_.28.24C000-.24C7FF.29_w_.28Namco_175_only.29"></span>External PRG RAM enable (\$C000-\$C7FF) w (Namco 175 only) {#External_PRG_RAM_enable_.28.24C000-.24C7FF.29_w_.28Namco_175_only.29}

    7  bit  0
    ---- ----
    .... ...E
            |
            +- 1: Enable PRG RAM

### <span id="PRG_Select_1_.28.24E000-.24E7FF.29_w"></span>PRG Select 1 (\$E000-\$E7FF) w {#PRG_Select_1_.28.24E000-.24E7FF.29_w}

    7  bit  0
    ---- ----
    MMPP PPPP
    |||| ||||
    ||++-++++- Select 8KB page of PRG-ROM at $8000
    ++-------- Namco 340 only: Select mirroring
                 0: One-screen A
                 1: Vertical
                 2: One-screen B
                 3: Horizontal

### <span id="PRG_Select_2_.28.24E800-.24EFFF.29_w"></span>PRG Select 2 (\$E800-\$EFFF) w {#PRG_Select_2_.28.24E800-.24EFFF.29_w}

    7  bit  0
    ---- ----
    ..PP PPPP
      || ||||
      ++-++++- Select 8KB page of PRG-ROM at $A000

### <span id="PRG_Select_3_.28.24F000-.24F7FF.29_w"></span>PRG Select 3 (\$F000-\$F7FF) w {#PRG_Select_3_.28.24F000-.24F7FF.29_w}

    7  bit  0
    ---- ----
    ..PP PPPP
      || ||||
      ++-++++- Select 8KB page of PRG-ROM at $C000

## Notes {#Notes}

- *Family Circuit '91* relies on its 2 KiB of PRG-RAM being correctly mirrored throughout the \$6000-\$7FFF address range, while *Splatterworld* relies on having a full 8 KiB in the same range.
  - Both of these games contain a hidden test activated by [shorting expansion port pins](../Namco_Serial_Checker_Interface/#SRAM_Checker_Variant), which only passes or reinitializes PRG-RAM contents correctly with the aforementioned PRG-RAM sizes.
- Many INES Mapper 210 games are incorrectly set to [INES Mapper 019](../INES_Mapper_019/).
- All the commercial Namco 175 games (including the unreleased Splatterworld) set the upper bits of \$E000 to match their respective hardwired nametable mirroring. Thus, in the absence of a NES 2.0 header, emulating this mapper as a 175 if the header battery bit is set and a 340 otherwise will work with all commercial games.
- *Splatterhouse* and *Famista '91* both perform a complete Namco 163 initialization on reset, writing to various registers that do not exist on the 175 or 340. *Splatterhouse* additionally writes to the 163 nametable select registers (\$C000-\$DFFF) every frame.

## See also {#See_also}

- [Naruko's blog's post on mirroring on the 340](http://d.hatena.ne.jp/na6ko/20110501/1304181038)
