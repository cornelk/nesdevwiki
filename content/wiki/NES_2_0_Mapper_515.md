---
title: "NES 2.0 Mapper 515"
url: "/wiki/NES_2.0_Mapper_515/"
description: "NES 2.0 Mapper 515 is used for 패밀리 노래방 Family Noraebang, a Korean karaoke cartridge that uses UNROM-style PRG banking, has provisions for an expansion cartridge, an ADC for Microphone…"
wikiPage: "NES_2_0_Mapper_515"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
---

***NES 2.0 Mapper 515*** is used for 패밀리 노래방 *Family Noraebang*, a Korean karaoke cartridge that uses UNROM-style PRG banking, has provisions for an expansion cartridge, an ADC for Microphone input, and a clone of the Yamaha [YM2413](https://en.wikipedia.org/wiki/Yamaha_YM2413) chip for expansion audio: the K-663A.

## Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank

## PRG-ROM bankswitching {#PRG-ROM_bankswitching}

    Mask $E00F:
      W 8000: [RBBB BBBB] - R:0-select external ROM; 1-select internal ROM; B-select bank at $8000-$BFFF
                            $C000-$FFFF is always last bank of internal ROM

The board does not seem to suffer from bus conflicts. The existing ROM images are arranged similarly to [INES Mapper 188](../INES_Mapper_188/): The last 1 MiB of PRG-ROM are the main cartridge, the first (total PRG-ROM size minus 1 MiB) are the expansion cartridge.

## Expansion sound {#Expansion_sound}

    Mask: $E003:
      W 6000: YM2413 index
      W 6001: YM2413 data
      R 6000: YM2413 read back test register contents

Note that the board does not allow the 2A03 audio to be mixed with the YM2413 output.

## Microphone input {#Microphone_input}

    Mask: $E003:
       R 6003: [D... ....] - SPI (ADC) data
    Mask not yet known:
       W C002: [S... ....] - SPI (ADC) chip select
       W C003: [K... ....] - SPI clock / ADC conversion clock

## See also {#See_also}

<http://forums.nesdev.org/viewtopic.php?f=9&t=16124>
