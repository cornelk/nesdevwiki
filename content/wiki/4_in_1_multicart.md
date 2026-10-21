---
title: "4-in-1 multicart"
url: "/wiki/4-in-1_multicart/"
wikiPage: "4_in_1_multicart"
wikiCategories:
  - name: "Mappers"
    url: "/category/Mappers/"
---

4-in-1 multicart

[<img src="../../wiki-images/4-in-1_menu.png" width="256" height="223" />](../../wiki-images/4-in-1_menu.png)

[](../../wiki-images/4-in-1_menu.png)

[4-in-1 menu](../../wiki-images/4-in-1_menu.png)

[](../../wiki-images/4-in-1_menu.png)

[](../../wiki-images/4-in-1_menu.png)

[](../../wiki-images/4-in-1_menu.png)[<img src="../../wiki-images/4-in-1_top.jpg" width="300" height="159" />](../../wiki-images/4-in-1_top.jpg)

[](../../wiki-images/4-in-1_top.jpg)

[4-in-1 top](../../wiki-images/4-in-1_top.jpg)

[](../../wiki-images/4-in-1_top.jpg)

[](../../wiki-images/4-in-1_top.jpg)

[](../../wiki-images/4-in-1_top.jpg)[<img src="../../wiki-images/4-in-1_bottom.jpg" width="300" height="160" />](../../wiki-images/4-in-1_bottom.jpg)

[](../../wiki-images/4-in-1_bottom.jpg)

[4-in-1 bottom](../../wiki-images/4-in-1_bottom.jpg)

[](../../wiki-images/4-in-1_bottom.jpg)

[](../../wiki-images/4-in-1_bottom.jpg)

[](../../wiki-images/4-in-1_bottom.jpg)[<img src="../../wiki-images/4-in-1_schematic.png" width="300" height="187" />](../../wiki-images/4-in-1_schematic.png)

[](../../wiki-images/4-in-1_schematic.png)

[4-in-1 schematic](../../wiki-images/4-in-1_schematic.png)

[](../../wiki-images/4-in-1_schematic.png)

This cartridge contains 4 CNROM games:

- Puzzle
- F-15 City War
- Volley Ball
- PokeBlock

## Banks {#Banks}

- CPU \$8000-\$ffff: 32 KiB switchable PRG ROM bank
- PPU \$0000-\$1fff: 8 KiB switchable CHR ROM bank

## Mirroring {#Mirroring}

H/V (selected by software)

## Subject to bus conflicts {#Subject_to_bus_conflicts}

Yes

## Registers {#Registers}

Mask: \$8000

    $8000: [..pP CCCC]
              || ||||
              || ++++- select 8kB CHR-ROM bank at $0000-$1fff
              ++------ select 32 kB PRG-ROM bank at $8000-$ffff
              +------- select mirroring (0=H, 1=V)

## Quirks {#Quirks}

- Multicart register overlaps CNROM register and multicart register is not protected from writes after game is selected, so all games' bankswitching routines had to be patched
- The register is not initialized on powerup or reset, so all games' reset vector had to be patched
- DIP28 1 Mibit mask roms are used both for PRG and CHR (pin 22, which is normally used as !OE in DIP28 EEPROMS is utilized as A16)
- This is basically identical to [GxROM](../GxROM/) but one bit also determines nametable mirroring as well as PRG bank
