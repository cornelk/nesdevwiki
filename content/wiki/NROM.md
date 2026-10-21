---
title: "NROM"
url: "/wiki/NROM/"
description: "The generic designation NROM refers to the Nintendo cartridge boards NES-NROM-128, NES-NROM-256, their HVC counterparts, and clone boards. The iNES format assigns mapper 0 to NROM."
wikiPage: "NROM"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>NROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo, others </td></tr><tr><td><b>Boards</b></td><td>NROM, HROM*, RROM, RTROM, SROM, STROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>16K or 32K </td></tr><tr><td><b>PRG ROM window</b></td><td>n/a </td></tr><tr><td><b>PRG RAM capacity</b></td><td>2K or 4K in <i>Family BASIC</i> only </td></tr><tr><td><b>PRG RAM window</b></td><td>n/a </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads (*V only) </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">000</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=0">iNES 000</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-NROM">NROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-HROM">HROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-RROM">RROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-RTROM">RTROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-SROM">SROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-STROM">STROM</a></td></tr></tbody></table></div>

The generic designation **NROM** refers to the Nintendo cartridge boards NES-NROM-128, NES-NROM-256, their [HVC](../Family_Computer/) counterparts, and clone boards. The [iNES](../INES/) format assigns **mapper 0** to NROM.

The suffixes 128 and 256 refer to kilobits by Nintendo's own designation; not kilobytes (as is the case with [UNROM 512](../UNROM_512/), for example).

Although mapper 0 does not have any submappers, other fields in the [NES 2.0](../NES_2.0/) header may be used to specify the following:

- (Battery-backed) PRG-RAM size, for *Family BASIC*
- PRG-ROM size in exponent-multiplier form, for *Galaxian*

## Banks {#Banks}

- CPU \$6000-\$7FFF: Unbanked PRG-RAM, mirrored as necessary to fill entire 8 KiB window, write protectable with an external switch. (*Family BASIC* only)
- CPU \$8000-\$BFFF: First 16 KiB of PRG-ROM.
- CPU \$C000-\$FFFF: Last 16 KiB of PRG-ROM (NROM-256) or mirror of \$8000-\$BFFF (NROM-128).
- PPU \$0000-\$1FFF: 8 KiB CHR-ROM.

All banks are fixed.

## Solder pad config {#Solder_pad_config}

- Vertical arrangement ("Horizontal mirroring") : 'H' disconnected, 'V' connected.
- Horizontal arrangement ("Vertical mirroring") : 'H' connected, 'V' disconnected.

## Registers {#Registers}

None. This normally has no mapping capabilities whatsoever, so the presence of [bus conflicts](../Bus_conflict/) are irrelevant from a software perspective.

Nevertheless, tile animation can be done by swapping between pattern tables \$0000 and \$1000, using [PPUCTRL](../PPU_registers/) bits 4-3 as a "poor man's [CNROM](../CNROM/)".

## Variants {#Variants}

These board variants do not affect emulation:

- [RROM and SROM](../Mask_ROM_pinout/#Variants) use different CHR-ROM pinouts.
- RTROM and STROM split the PRG-ROM into two 8KiB ROMs.
- HROM is an early variant without the solder pads - only the horizontal arrangement ("vertical mirroring") is available, as if the 'H' pad were selected.

A few cartridges require NES 2.0 features to accurately specify:

- *Family BASIC* contains 2KiB PRG-RAM (or 4KiB PRG-RAM for V3) accessible at \$6000-\$7FFF and [decoded with a 74HC20](../PRG_RAM_circuit/), that is backed with 2 AA batteries. The cartridge is equipped with a back up switch that, when enabled, disables all access to PRG-RAM in order to prevent possible data corruption when the Famicom is turned on or off. Most emulators will map 8KiB of PRG-RAM to \$6000-\$7FFF if the battery flag is set (for iNES compatibility) - the NES 2.0 PRG-RAM size field should be used to denote the exact size.
- [*Galaxian*](https://nescartdb.com/profile/view/1808/galaxian) uses a smaller 8KiB PRG-ROM chip. iNES dumps of this game are therefore all overdumps. NES 2.0 is capable of specifying this size using the exponent-multiplier form.

Some unlicensed games and modern homebrew programs use mapper 0 with 8KiB of CHR-RAM, which would not be supported on official NROM PCBs without rewiring ([BNROM](../INES_Mapper_034/) may be a better option for reproduction boards). Most emulators will do the right thing with mapper 0 + CHR-RAM, however.

[NROM-368](../NROM-368/) is a recent invention that allows addressing more PRG-ROM without bankswitching.

## See also {#See_also}

- [Programming NROM](../Programming_NROM/)
