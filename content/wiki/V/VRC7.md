---
title: "VRC7"
url: "/wiki/VRC7/"
description: "The Konami VRC7 is an ASIC mapper. It is assigned to iNES Mapper 085."
wikiPage: "VRC7"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

<div class="infobox"><div style="text-align: center"><b>VRC7</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=85">2 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>352402, 353429 </td></tr><tr><td><b>Pinout</b></td><td><a href="../VRC7_pinout/">VRC7 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K </td></tr><tr><td><b>PRG ROM window</b></td><td>8Kx3 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K fixed </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td><a href="../VRC7_audio/">VRC7a only</a></td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">085</a></td></tr></tbody></table></div>

The **Konami VRC7** is an [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/). It is assigned to **iNES Mapper 085**.

Example games:

- *Lagrange Point*
- *Tiny Toon Adventures 2*

See [VRC7 pinout](../VRC7_pinout/) for chip pinout.

## Banks {#Banks}

### CPU {#CPU}

- \$6000-\$7FFF: 8 KB PRG-RAM bank, fixed
- \$8000-\$9FFF: 8 KB switchable PRG-ROM bank
- \$A000-\$BFFF: 8 KB switchable PRG-ROM bank
- \$C000-\$DFFF: 8 KB switchable PRG-ROM bank
- \$E000-\$FFFF: 8 KB PRG-ROM bank, fixed to the last bank

### PPU {#PPU}

- \$0000-\$03FF: 1 KB switchable CHR-ROM bank
- \$0400-\$07FF: 1 KB switchable CHR-ROM bank
- \$0800-\$0BFF: 1 KB switchable CHR-ROM bank
- \$0C00-\$0FFF: 1 KB switchable CHR-ROM bank
- \$1000-\$13FF: 1 KB switchable CHR-ROM bank
- \$1400-\$17FF: 1 KB switchable CHR-ROM bank
- \$1800-\$1BFF: 1 KB switchable CHR-ROM bank
- \$1C00-\$1FFF: 1 KB switchable CHR-ROM bank

If CHR-RAM is used instead of CHR-ROM, the banking feature is still functional.

## Registers {#Registers}

There are two board variations:

- VRC7b uses A3 to select registers (\$x008), used in *Tiny Toon Adventures* (**Submapper 1**).
- VRC7a uses A4 to select registers (\$x010), used in *Lagrange Point* (**Submapper 2**).

Although A5 is wired for sound registers on both (\$x028, \$x030), in the former board the ceramic resonator necessary for the sound hardware to work is missing.

### <span id="PRG_Select_0_.28.248000.29"></span>PRG Select 0 (\$8000) {#PRG_Select_0_.28.248000.29}

    7  bit  0
    ---------
    ..PP PPPP
      || ||||
      ++-++++- Select 8 KB PRG ROM at $8000

### <span id="PRG_Select_1_.28.248010.2C_.248008.29"></span>PRG Select 1 (\$8010, \$8008) {#PRG_Select_1_.28.248010.2C_.248008.29}

    7  bit  0
    ---------
    ..PP PPPP
      || ||||
      ++-++++- Select 8 KB PRG ROM at $A000

### <span id="PRG_Select_2_.28.249000.29"></span>PRG Select 2 (\$9000) {#PRG_Select_2_.28.249000.29}

    7  bit  0
    ---------
    ..PP PPPP
      || ||||
      ++-++++- Select 8 KB PRG ROM at $C000

### <span id="CHR_Select_0.E2.80.A67_.28.24A000.E2.80.A6.24DFFF.29"></span>CHR Select 0…7 (\$A000…\$DFFF) {#CHR_Select_0.E2.80.A67_.28.24A000.E2.80.A6.24DFFF.29}

| Write to CPU address | 1KB CHR bank affected |
|----------------------|-----------------------|
| \$A000               | \$0000-\$03FF         |
| \$A008 or \$A010     | \$0400-\$07FF         |
| \$B000               | \$0800-\$0BFF         |
| \$B008 or \$B010     | \$0C00-\$0FFF         |
| \$C000               | \$1000-\$13FF         |
| \$C008 or \$C010     | \$1400-\$17FF         |
| \$D000               | \$1800-\$1BFF         |
| \$D008 or \$D010     | \$1C00-\$1FFF         |

### <span id="Mirroring_Control_.28.24E000.29"></span>Mirroring Control (\$E000) {#Mirroring_Control_.28.24E000.29}

    7  bit  0
    ---------
    RS.. ..MM
    ||     ||
    ||     ++- Mirroring (0: vertical; 1: horizontal;
    ||                        2: one-screen, lower bank; 3: one-screen, upper bank)
    |+-------- Silence expansion sound if set
    +--------- WRAM enable (1: enable WRAM, 0: protect)

### <span id="IRQ_Control_.28.24E008_-_.24F010.29"></span>IRQ Control (\$E008 - \$F010) {#IRQ_Control_.28.24E008_-_.24F010.29}

    $E008, $E010:  IRQ Latch
           $F000:  IRQ Control
    $F008, $F010:  IRQ Acknowledge

Many VRC mappers use the same IRQ system. For details on IRQ operation, see [VRC IRQs](../VRC_IRQ/).

## <span id="Sound_.28.249010.2C_.249030.29"></span>Sound (\$9010, \$9030) {#Sound_.28.249010.2C_.249030.29}

For details on sound information, see [VRC7 audio](../VRC7_audio/).

## See also {#See_also}

- [Konami VRC-VII Chip Info](http://nesdev.org/vrcvii.txt) by Kevin Horton.
- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Submapper test ROMs](https://forums.nesdev.org/viewtopic.php?f=3&t=17573#p223145)
- [siliconpr0n VRC7 decapsulation photographs](https://siliconpr0n.org/archive/doku.php?id=digshadow:konami:vrc_vii_053982)
