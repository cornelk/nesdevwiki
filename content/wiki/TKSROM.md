---
title: "TKSROM"
url: "/wiki/TKSROM/"
description: "TKSROM refers to a variant of the TKROM board with the VRAM mirroring signals routed through the CHR side of the Nintendo MMC3 a different way, as described in iNES Mapper 118. It is used only for…"
wikiPage: "TKSROM"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

**TKSROM** refers to a variant of the [TKROM](../TxROM/) board with the VRAM mirroring signals routed through the CHR side of the [Nintendo MMC3](../MMC3/) a different way, as described in [iNES Mapper 118](../INES_Mapper_118/). It is used only for the Japanese game *Ys III: Wanderers from Ys*.

## Overview {#Overview}

- PRG ROM size: 128, 256 or 512 KB (DIP-28/32 Nintendo pinout)
- PRG ROM bank size: 8 KB
- PRG RAM: 8 KB with optional battery
- CHR capacity: 128 KB ROM (DIP-32 Nintendo pinout)
- CHR bank size: 1 KB and 2 KB
- Nametable [mirroring](../Mirroring/): Controlled by mapper, but in a nonstandard way (see Mapper 118).
- Subject to [bus conflicts](../Bus_conflict/): No

## Solder Pad Config {#Solder_Pad_Config}

To be added.
