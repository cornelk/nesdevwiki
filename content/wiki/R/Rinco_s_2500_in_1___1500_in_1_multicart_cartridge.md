---
title: "Rinco's 2500-in-1 / 1500-in-1 multicart cartridge"
url: "/wiki/Rinco's_2500-in-1_/_1500-in-1_multicart_cartridge/"
description: "This multicart cartridge was placed in Rinco Thompsonic's second slot cartridge and shipped with it. It consists of 1 MB mask ROM, 32 kB EEPROM, PAL16L8, latches and some gates. It allows to run the…"
wikiPage: "Rinco_s_2500_in_1___1500_in_1_multicart_cartridge"
wikiCategories:
  - name: "Mappers"
    url: "/category/Mappers/"
---

This multicart cartridge was placed in Rinco Thompsonic's second slot cartridge and shipped with it. It consists of 1 MB mask ROM, 32 kB EEPROM, PAL16L8, latches and some gates. It allows to run the following mappers: 16 kB NROM, 32 kB NROM, 128kB UNROM.

[<img src="../../../wiki-images/2500-in-1_top.jpg" width="300" height="234" />](../../../wiki-images/2500-in-1_top.jpg)

[](../../../wiki-images/2500-in-1_top.jpg)

[2500-in-1 top](../../../wiki-images/2500-in-1_top.jpg)

[](../../../wiki-images/2500-in-1_top.jpg)

[](../../../wiki-images/2500-in-1_top.jpg)

[](../../../wiki-images/2500-in-1_top.jpg)[<img src="../../../wiki-images/2500-in-1_bottom.jpg" width="300" height="242" />](../../../wiki-images/2500-in-1_bottom.jpg)

[](../../../wiki-images/2500-in-1_bottom.jpg)

[2500-in-1 bottom](../../../wiki-images/2500-in-1_bottom.jpg)

[](../../../wiki-images/2500-in-1_bottom.jpg)

[](../../../wiki-images/2500-in-1_bottom.jpg)

[](../../../wiki-images/2500-in-1_bottom.jpg)[<img src="../../../wiki-images/2500-in-1_title_screen.jpg" width="300" height="246" />](../../../wiki-images/2500-in-1_title_screen.jpg)

[](../../../wiki-images/2500-in-1_title_screen.jpg)

[2500-in-1 title screen](../../../wiki-images/2500-in-1_title_screen.jpg)

[](../../../wiki-images/2500-in-1_title_screen.jpg)

[](../../../wiki-images/2500-in-1_title_screen.jpg)

[](../../../wiki-images/2500-in-1_title_screen.jpg)[<img src="../../../wiki-images/2500-in-1_schematic.png" width="300" height="198" />](../../../wiki-images/2500-in-1_schematic.png)

[](../../../wiki-images/2500-in-1_schematic.png)

[2500-in-1 schematic](../../../wiki-images/2500-in-1_schematic.png)

[](../../../wiki-images/2500-in-1_schematic.png)

To be exact, it consists of the following games:

- 1\. contra
- 2\. tetris 2
- 3\. tank
- 4\. b-wings
- 5\. super mario 1
- 6\. 1942
- 7\. ninja III
- 8\. dig dug II
- 9\. binary land
- 10\. bomber man
- 11\. ice climber
- 12\. galaxian
- 13\. road fighter
- 14\. hogans alley
- 15\. wild gunman
- 16\. duck shoot
- 17\. antarctic adventure
- 18\. dig dug
- 19\. urban champion
- 20\. baloon fight
- 21\. clu clu clan jss
- 22\. sky destroyer
- 23\. popeye
- 24\. wrecking crew
- 25\. raid on bay her
- 26\. milk nuts
- 27\. devil world
- 28\. chack pop
- 29\. macross
- 30\. mario bros

## CPU Banks {#CPU_Banks}

\$8000-\$BFFF: 16 kB PRG bank

\$C000-\$FFFF: 16 kB PRG bank

or

\$8000-\$FFFF: 32 kB PRG bank

If PRG bank = 0 or PRG bank = 7 (PRG-A16=1 and PRG-A15=1 and PRG-A14=1) then EEPROM is used, otherwise MaskRom is used.

## PPU Banks {#PPU_Banks}

\$0000-\$1FFF: 8 kB CHR RAM (not banked), but in some modes it can be protected against writes.

## Mirroring {#Mirroring}

H or V (set by software).

## Subject to bus conflicts {#Subject_to_bus_conflicts}

Yes.

## Registers {#Registers}

\$8000-\$FFFF: one global configuration register (bits of written address is latched, data is irrelevant)

    15                0
    ---- ---- ---- ----
    1... ..pP wPPQ QRmM
           || |||| ||||
           || |||| |||+--- PRG size (0=16 kB, 1=32 kB)
           || |||| ||+---- mirroring (0=vertical, 1=horizontal)
           || |||+-++----- Inner PRG bank at $8000
           |+--++--------- 128 KiB Outer PRG bank
           |  +----------- UNROM mode (0=enabled) / CHR-RAM write protection (0=disabled, 1=enabled)
           +-------------- Inner PRG bank at $C000 in UNROM mode (note, single bit but used three times)

| M,w \\ PRG bank | \$8000-\$bfff | \$c000-\$ffff | resembles               |
|-----------------|---------------|---------------|-------------------------|
| M=0 and w=0     | PPPQQR        | PPPppp        | UNROM                   |
| M=0 and w=1     | PPPQQR        | PPPQQR        | 16K NROM                |
| M=1 and w=0     | PPPQQ0        | PPPppp        | (defective combination) |
| M=1 and w=1     | PPPQQ0        | PPPQQ1        | 32K NROM                |

## Quirks {#Quirks}

- In most cartridges where reset detection is needed, an R-C-D circuit is present and connected to any latches' clear inputs. As this cartridge uses current-based 74LS logic, only a capacitor and diode are needed.

<!-- -->

- The total PRG capacity is not 1 MiB + 32 KiB but only 1 MiB: the 32 KiB EPROM is a patch that replaces banks 0 and 7. It is believed that those banks contains menu data (list of games) and authors of the cartridge were unhappy with the original menu from mask ROM and they wanted to change it.

<!-- -->

- There are at least two variants of this cartridge with different markings on mask ROM and EPROM (other logic, including PAL is the same):
  - 1500-in-1 (mask ROM: CT-121H, EPROM: 91150)
  - 2500-in-1 (mask ROM: CT-120H, EPROM: 91250)
