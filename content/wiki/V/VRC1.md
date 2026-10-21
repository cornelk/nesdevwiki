---
title: "VRC1"
url: "/wiki/VRC1/"
description: "The Konami VRC1 is an ASIC mapper. The iNES format assigns mapper 75 to VRC1."
wikiPage: "VRC1"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
---

<div class="infobox"><div style="text-align: center"><b>VRC1</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Konami </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=75">6 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>JF-20, JF-22, 4036, 302114A, 350459 </td></tr><tr><td><b>Pinout</b></td><td><a href="../VRC1_pinout/">VRC1 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>none </td></tr><tr><td><b>PRG RAM window</b></td><td>n/a </td></tr><tr><td><b>CHR capacity</b></td><td>128K </td></tr><tr><td><b>CHR window</b></td><td>4K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">075</a></td></tr></tbody></table></div>

The Konami VRC1 is an [ASIC](../../category/ASIC_mappers/) [mapper](../Mapper/). The [iNES](../INES/) format assigns **mapper 75** to VRC1.

The VRC1 was used in several Famicom games released by Konami and Jaleco:

- *Exciting Boxing*
- *Ganbare Goemon! Karakuri Douchuu*
- *Jajamaru Ninpouchou*
- *King Kong 2: Ikari no Megaton Punch*
- *Moero!! Junior Basket: Two on Two*
- *Tetsuwan Atom*

## Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$0FFF: 4 KiB switchable CHR bank
- PPU \$1000-\$1FFF: 4 KiB switchable CHR bank

## Registers {#Registers}

### <span id="PRG_Select_0_.28.248000-.248FFF.29"></span>PRG Select 0 (\$8000-\$8FFF) {#PRG_Select_0_.28.248000-.248FFF.29}

    7  bit  0
    ---------
    .... PPPP
         ||||
         ++++- Select 8 KB PRG ROM at $8000

### <span id="PRG_Select_1_.28.24A000-.24AFFF.29"></span>PRG Select 1 (\$A000-\$AFFF) {#PRG_Select_1_.28.24A000-.24AFFF.29}

    7  bit  0
    ---------
    .... PPPP
         ||||
         ++++- Select 8 KB PRG ROM at $A000

### <span id="PRG_Select_2_.28.24C000-.24CFFF.29"></span>PRG Select 2 (\$C000-\$CFFF) {#PRG_Select_2_.28.24C000-.24CFFF.29}

    7  bit  0
    ---------
    .... PPPP
         ||||
         ++++- Select 8 KB PRG ROM at $C000

### <span id="Mirroring_Control.2C_CHR_bits_.28.249000-.249FFF.29"></span>Mirroring Control, CHR bits (\$9000-\$9FFF) {#Mirroring_Control.2C_CHR_bits_.28.249000-.249FFF.29}

    7  bit  0
    ---------
    .... .BAM
          |||
          ||+- Mirroring  (0: Vertical; 1: Horizontal)
          |+-- High Bit of 4 KB CHR bank at PPU $0000
          +--- High Bit of 4 KB CHR bank at PPU $1000

The mirroring bit is ignored if the cartridge is wired for 4-screen VRAM, as is typical for [Vs. System](../Vs._System/) games using the VRC1.

### <span id="CHR_Select_0_.28.24E000-.24EFFF.29"></span>CHR Select 0 (\$E000-\$EFFF) {#CHR_Select_0_.28.24E000-.24EFFF.29}

    7  bit  0
    ---------
    .... CCCC
         ||||
         ++++- Low 4 bits of 4 KB CHR bank at PPU $0000

These bits combined with bit 1 of \$9000 make a 5-bit CHR selection.

### <span id="CHR_Select_1_.28.24F000-.24FFFF.29"></span>CHR Select 1 (\$F000-\$FFFF) {#CHR_Select_1_.28.24F000-.24FFFF.29}

    7  bit  0
    ---------
    .... CCCC
         ||||
         ++++- Low 4 bits of 4 KB CHR bank at PPU $1000

These bits combined with bit 2 of \$9000 make a 5-bit CHR selection.

## References {#References}

- [Disch's Mapper Notes](http://www.romhacking.net/documents/362/)
