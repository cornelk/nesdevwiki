---
title: "FxROM"
url: "/wiki/FxROM/"
description: "The generic designation FxROM refers to cartridge boards made by Nintendo that use the Nintendo MMC4 mapper."
wikiPage: "FxROM"
wikiCategories:
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

The generic designation **FxROM** refers to cartridge boards made by Nintendo that use the [Nintendo MMC4](../MMC4/) mapper.

## Board Types {#Board_Types}

The following FxROM boards are known to exist:

| Board | PRG ROM      | PRG RAM | CHR                 |
|-------|--------------|---------|---------------------|
| FJROM | 128 / 256 KB | 8 KB    | 16 / 32 / 64 KB ROM |
| FKROM | 128 / 256 KB | 8 KB    | 128 KB ROM          |

## Overview {#Overview}

- PRG ROM size: up to 256 KB
- PRG ROM bank size: 16 KB
- PRG RAM: 8 KB plus optional battery
- CHR capacity: up to 128 KB ROM
- CHR bank size: 4 KB
- Nametable [mirroring](../Mirroring/): Controlled by mapper
- Subject to [bus conflicts](../Bus_conflict/): No

## Solder Pad config {#Solder_Pad_config}

### Battery backup {#Battery_backup}

- PRG RAM with data retention : 'SL' disconnected, D1, D2, R1, R2, R3 present.
- PRG RAM without data retention : 'SL' connected, D1, D2, R1, R2, R3 not present.

On the FKROM board, battery backup appears to be compulsory unless there is a hidden solder pad below the MM1026 chip.

## Hardware {#Hardware}

The FKROM board uses a [Mitsumi MM1026](http://www.mitsumi.co.jp/latest/Catalog/ic/common/stop/MM1026_1245_1080_1134_e.pdf) chip dedicated to safe battery backup. It's the only Nintendo board to use a specific chip dedicated to safe battery backup.

## See also {#See_also}

- [Nintendo MMC4](../MMC4/)
