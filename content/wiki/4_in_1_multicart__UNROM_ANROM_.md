---
title: "4-in-1 multicart (UNROM/ANROM)"
url: "/wiki/4-in-1_multicart_(UNROM/ANROM)/"
wikiPage: "4_in_1_multicart__UNROM_ANROM_"
wikiCategories:
  - name: "Categories:"
    url: "/category/Categories_/"
  - name: "Mappers"
    url: "/category/Mappers/"
---

[<img src="../../../wiki-images/4-in-1_unrom%2Canrom_menu.png" width="256" height="223" />](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)

[Menu](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_menu.png)[<img src="../../../wiki-images/4-in-1_unrom%2Canrom_components.jpg" width="300" height="147" />](../../../wiki-images/4-in-1_unrom%2Canrom_components.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_components.jpg)

[Components](../../../wiki-images/4-in-1_unrom%2Canrom_components.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_components.jpg)

PRG: 512 kiB ROM

CHR: 8 kiB RAM

Bus conflicts: no

    gG | GAME
    ---+----------------
    00 | Trog
    10 | Duck Tales I
    01 | Marble Madness
    11 | Captain Skyhawk

## Registers {#Registers}

### <span id=".246000-.247FFF:_UNROM.2FANROM_banking_type_.26_game_number"></span>\$6000-\$7FFF: UNROM/ANROM banking type & game number {#.246000-.247FFF:_UNROM.2FANROM_banking_type_.26_game_number}

    7  bit  0
    ---- ----
    TgGP
    ||||
    |||+------------ protection:    0 - off, 1 - contents of this register cannot be changed until console reset
    |++------------- selects one of four 128 kB games; g is lower address line: PRG-A18 = G, PRG-A17=g 
    +--------------- banking style: 0 - UNROM: $8000-$bfff 16 kB switchable CBA bank, $c000-$ffff fixed to last bank, PPU vertical mirroring
                                    1 - ANROM: $8000-$ffff 32 kB switchable BA bank, PPU single screen mirroring

### <span id=".248000-.24FFFF:_bank_select_.28when_in_UNROM_mode.29"></span>\$8000-\$FFFF: bank select (when in UNROM mode) {#.248000-.24FFFF:_bank_select_.28when_in_UNROM_mode.29}

    7  bit  0
    ---- ----
          CBA
          |||
          +++------------ UNROM bank (when in UNROM MODE)

### <span id=".248000-.24FFFF:_bank_select_.28when_in_ANROM_mode.29"></span>\$8000-\$FFFF: bank select (when in ANROM mode) {#.248000-.24FFFF:_bank_select_.28when_in_ANROM_mode.29}

    7  bit  0
    ---- ----
       M   BA
       |   ||
       |   ++------------ ANROM bank (when in ANROM MODE)
       +----------------- one screen mirroring

[<img src="../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg" width="300" height="142" />](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)

[Top layer](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_top.jpg)[<img src="../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg" width="300" height="143" />](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)

[Bottom layer](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)

[](../../../wiki-images/4-in-1_unrom%2Canrom_bottom.jpg)[<img src="../../../wiki-images/4-in-1_unrom%2Canrom_schematics.png" width="300" height="209" />](../../../wiki-images/4-in-1_unrom%2Canrom_schematics.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_schematics.png)

[Schematics](../../../wiki-images/4-in-1_unrom%2Canrom_schematics.png)

[](../../../wiki-images/4-in-1_unrom%2Canrom_schematics.png)
