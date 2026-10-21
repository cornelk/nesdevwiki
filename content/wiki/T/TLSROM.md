---
title: "TLSROM"
url: "/wiki/TLSROM/"
description: "TLSROM refers to a variant of the TLROM board with the VRAM mirroring signals routed through the CHR side of the Nintendo MMC3 a different way, as described in iNES Mapper 118."
wikiPage: "TLSROM"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

**TLSROM** refers to a variant of the [TLROM](../TxROM/) board with the VRAM mirroring signals routed through the CHR side of the [Nintendo MMC3](../MMC3/) a different way, as described in [iNES Mapper 118](../INES_Mapper_118/).

## Overview {#Overview}

- PRG ROM size: 128, 256 or 512 KB (DIP-28/32 Nintendo pinout)
- PRG ROM bank size: 8 KB
- PRG RAM: None
- CHR capacity: 128 KB ROM (DIP-32 Nintendo pinout)
- CHR bank size: 1 KB and 2 KB
- Nametable [mirroring](../Mirroring/): Controlled by mapper, but in a non-standard way (see Mapper 118).
- Subject to [bus conflicts](../Bus_conflict/): No

## See also {#See_also}

- [Nintendo MMC3](../MMC3/)
