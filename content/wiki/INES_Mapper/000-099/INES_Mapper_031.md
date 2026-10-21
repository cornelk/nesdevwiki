---
title: "INES Mapper 031"
url: "/wiki/INES_Mapper_031/"
description: "iNES Mapper 031 represents a mapper created to facilitate cartridge compilations of NSF music. It implements a common subset of the features used by NSFs."
wikiPage: "INES_Mapper_031"
wikiCategories:
  - name: "CPLD mappers"
    url: "/category/CPLD_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
---

<div class="infobox"><div style="text-align: center"><b>NSF subset</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>varies (homebrew) </td></tr><tr><td><b>Complexity</b></td><td>CPLD </td></tr><tr><td><b>Boards</b></td><td>unknown (homebrew) </td></tr><tr><td><b>PRG ROM capacity</b></td><td>1024K </td></tr><tr><td><b>PRG ROM window</b></td><td>4K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Not applicable </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>031 </td></tr></tbody></table></div>

**iNES Mapper 031** represents a mapper created to facilitate cartridge compilations of [NSF](../NSF/) music. It implements a common subset of the features used by NSFs.

PRG-ROM is bankswitched in 8 x 4 kB banks from \$8000-FFFF. These are controlled by registers at \$5FF8-\$5FFF like the NSF mapper. The high bank at \$F000-FFFF is initialized to the last bank at power-on.

There is no CHR banking, so it is recommended to use 8 kB CHR-RAM with this mapper.

As with [BNROM](../INES_Mapper_034/) and [UxROM](../UxROM/), there is no mirroring, CHR bank, or IRQ control; this mapper has hardwired H or V mirroring.

Examples:

- *2A03 Puritans* ([ROM](http://rainwarrior.ca/projects/nes/2a03puritans.html))
- *Famicompo Pico* ([ROMs](http://rainwarrior.ca/projects/nes/pico.html))
- *RNDM* ([info](http://megaranmusic.com/album/rndm))
- *EZNSF* ([forum post](https://forums.nesdev.org/viewtopic.php?t=15204))
- *ZENSF* ([forum post](https://forums.nesdev.org/viewtopic.php?f=6&t=17698))
- *Test ROMs* ([forum post](https://forums.nesdev.org/viewtopic.php?f=3&t=13120))

## Registers {#Registers}

### <span id="PRG_bank_select_.245000-.245FFF"></span>PRG bank select \$5000-\$5FFF {#PRG_bank_select_.245000-.245FFF}

    address              data
    15      bit       0  7  bit  0
    -------------------  ---------
    0101 .... .... .AAA  PPPP PPPP
                    |||  |||| ||||
                    |||  ++++-++++- Select 4 kB PRG ROM bank at slot specified by write address.
                    +++------------ Specify 4 kB bank slot at: $8000 + (AAA * $1000)

The canonical write position for these registers is \$5FF8-\$5FFF, as used in [NSFs](../NSF/).

At power on, the register at \$5FFF is set to \$FF. Startup code should be placed in the last bank. There is no change to this register on reset.

## References {#References}

- [Excerpt from readme.txt describing the mapper](http://forums.nesdev.org/viewtopic.php?p=128137#p128137)

## Emulator support {#Emulator_support}

- [BizHawk](http://tasvideos.org/BizHawk.html) 1.11.2
- [FCEUX](http://www.fceux.com) 2.2.3
- [MAME](http://mamedev.org/) 0.162
- [Mesen](https://www.mesen.ca/) 0.9.0
- [Nintendulator](https://www.qmtpro.com/~nes/nintendulator/) 0.975
- [OpenEmu](https://openemu.org/) (FCEUX core) 2.0.3
- [puNES](http://forums.nesdev.org/viewtopic.php?t=6928) 0.84

Cartridges:

- [Everdrive N8](../Everdrive_N8/) OS V14
- [InfiniteNESLives](http://www.infiniteneslives.com/nessupplies.php) Mapper 31 board
