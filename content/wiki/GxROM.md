---
title: "GxROM"
url: "/wiki/GxROM/"
description: "The designation GxROM refers to Nintendo cartridge boards labeled NES-GNROM and NES-MHROM (and their HVC counterparts), which use discrete logic to provide up to four 32 KB banks of PRG ROM and up to…"
wikiPage: "GxROM"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>GxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, others </td></tr><tr><td><b>Boards</b></td><td>GNROM,MHROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128KiB (512KiB oversize) </td></tr><tr><td><b>PRG ROM window</b></td><td>32KiB </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>32KiB (128KiB oversize) </td></tr><tr><td><b>CHR window</b></td><td>8KiB </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">066</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=66">iNES 066</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=GNROM">GNROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=MHROM">MHROM</a></td></tr></tbody></table></div>

The designation **GxROM** refers to Nintendo cartridge boards labeled NES-GNROM and NES-MHROM (and their [HVC](../Family_Computer/) counterparts), which use discrete logic to provide up to four 32 KB banks of PRG ROM and up to four 8 KB banks of CHR ROM. The [iNES](../INES/) format assigns **mapper 66** to these boards.

The Jaleco board assigned to [iNES Mapper 140](../INES_Mapper_140/) is sometimes confused with GNROM, as they are very similar but with the bankswitch register in a different location.

Example games:

- *Doraemon*
- *Dragon Power*
- *Gumshoe*
- *Paris-Dakar Rally Special!*
- *Takahashi Meijin no Bugutte Honey*
- *Thunder & Lightning*
- *Super Mario Bros. + Duck Hunt* (MHROM)

## Board Types {#Board_Types}

The following GxROM boards are known to exist:

| Board | PRG ROM | CHR            |
|-------|---------|----------------|
| GNROM | 128 KB  | 32 KB ROM      |
| MHROM | 64 KB   | 16 / 32 KB ROM |

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank
- PPU \$0000-\$1FFF: 8 KB switchable CHR ROM bank

## Registers {#Registers}

### <span id="Bank_select_.28.248000-.24FFFF.29"></span>Bank select (\$8000-\$FFFF) {#Bank_select_.28.248000-.24FFFF.29}

    7  bit  0
    ---- ----
    xxPP xxCC
      ||   ||
      ||   ++- Select 8 KB CHR ROM bank for PPU $0000-$1FFF
      ++------ Select 32 KB PRG ROM bank for CPU $8000-$FFFF

Bit 5 is not used on MHROM, which supports only 64 KB PRG.

## Solder pad config {#Solder_pad_config}

- Horizontal mirroring : 'H' disconnected, 'V' connected.
- Vertical mirroring : 'H' connected, 'V' disconnected.

## Hardware {#Hardware}

The GNROM board contains a [74HC161](../74161/) binary counter used as a quad D latch (4-bit register) to select the current PRG and CHR banks. MHROM, on the other hand, was often a [glop-top](https://www.google.com/search?q=Chip-on-board), as it was used for pack-in games, such as the *Super Mario Bros./Duck Hunt* [multicart](../Multicart/), and needed to be very inexpensive to produce in huge quantities.

## Variants {#Variants}

Theoretically the bank select register could be implemented with a [74HC377](../74377/) octal D latch, allowing up to 512 KB of PRG ROM and 128 KB of CHR ROM. There are a [large number](../../category/GNROM-like_mappers/) of other variants on GNROM, where the bits or the writeable address were moved around. Notable variants include:

- [iNES Mapper 140](../INES_Mapper_140/) places the bank register in \$6000-\$7FFF instead of \$8000-\$FFFF. [iNES Mapper 240](../INES_Mapper_240/) moves it further at \$4020-\$5FFF.
- [NES 2.0 Mapper 271](../NES_2.0_Mapper_271/) adds mapper-controlled nametable arrangement, controlled by one of the PRG banking bit.
- The [Color Dreams](../Color_Dreams/) board leaves the port at \$8000-\$FFFF, swaps the nibbles, expands CHR by two bits, and adds two bits for charge pump control. [iNES Mapper 046](../INES_Mapper_046/) additionally has an outer bank register.
- [iNES Mapper 086](../INES_Mapper_086/) supports speech sample playback in addition to an additional bit of CHR banking.
- [iNES Mapper 070](../INES_Mapper_070/) can be seen as a mix of GxROM and [UxROM](../UxROM/), which combines GNROM-style bank register with UxROM-style 16 KiB PRG banking. [iNES Mapper 152](../INES_Mapper_152/) additionally has hardware-controlled single-screen mirroring.

## See also {#See_also}

- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\ information about mapper's initial state is inaccurate.
