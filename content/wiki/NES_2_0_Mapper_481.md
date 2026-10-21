---
title: "NES 2.0 Mapper 481"
url: "/wiki/NES_2.0_Mapper_481/"
description: "NES 2.0 Mapper 481 denotes the Subor 045N PCB, used on at least one educational computer cartridge with two games on it. It's basically 2x128 KiB of UNROM, switched by an outer bank bit."
wikiPage: "NES_2_0_Mapper_481"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 481** denotes the Subor **045N** PCB, used on at least one educational computer cartridge with two games on it. It's basically 2x128 KiB of [UNROM](../UxROM/), switched by an outer bank bit.

- *小霸王 2合1꞉ 仓库世家 & 动脑筋*

## Banks {#Banks}

- CPU 8000-BFFF: Switchable 16 KiB inner bank
- CPU C000-FFFF: Fixed 16 KiB inner bank 7
- PPU 0000-1FFF: 8 KiB unbanked CHR-RAM
- Mirroring: hard-wired to horizontal or vertical

## <span id="Bank_register_.28.248000-.24FFFF.2C_write.29"></span>Bank register (\$8000-\$FFFF, write) {#Bank_register_.28.248000-.24FFFF.2C_write.29}

    D~[O... .BBB], Address mask: $8000
       |     +++- PRG A16..A14 if CPU A14=0
       |          7 if CPU A14=1
       +--------- PRG A17 regardless of CPU A14

## Notes {#Notes}

The common dump of *小霸王 2合1꞉ 仓库世家 & 动脑筋* is an overdump for FCEUX' idiosyncratic implementation of [INES Mapper 241](../INES_Mapper_241/).
