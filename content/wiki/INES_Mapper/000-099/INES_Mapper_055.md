---
title: "INES Mapper 055"
url: "/wiki/INES_Mapper_055/"
description: "iNES Mapper 055 is used for the UNIF board BTL-MARIO1-MALEE2, used by several bootleg versions of the first Super Mario Brothers that add a great number of cheats:"
wikiPage: "INES_Mapper_055"
wikiCategories:
  - name: "Categories:"
    url: "/category/Categories_/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

[<img src="../../wiki-images/Geniusmali-fc-pcbf.jpg" width="300" height="211" alt="Geniusmali-fc-pcbf.jpg" />](../../wiki-images/Geniusmali-fc-pcbf.jpg)

[](../../wiki-images/Geniusmali-fc-pcbf.jpg)

[<img src="../../wiki-images/Geniusmali-fc-pcbb.jpg" width="300" height="198" alt="Geniusmali-fc-pcbb.jpg" />](../../wiki-images/Geniusmali-fc-pcbb.jpg)

[](../../wiki-images/Geniusmali-fc-pcbb.jpg)

[](../../wiki-images/Geniusmali-fc-pcbb.jpg)

iNES Mapper 055 is used for the UNIF board **BTL-MARIO1-MALEE2**, used by several bootleg versions of the first *Super Mario Brothers* that add a great number of cheats:

- *Fly Merio Bros.*
- *Super Mario Bros. Malee 2*
- *天才瑪琍* (*Tiāncái Mǎlí*, a.k.a. *Genius Mario*)

It adds 2 KiB of PRG-ROM and PRG-RAM each to a regular NROM-256 PCB, but has no bankswitching functionalities.

# Banks {#Banks}

- CPU \$6000-\$6FFF: 2 KiB PRG-ROM, mirrored once
- CPU \$7000-\$7FFF: 2 KiB PRG-RAM, mirrored once
- CPU \$8000-\$FFFF: 32 KiB PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB CHR-ROM

# ROM layout {#ROM_layout}

Nestopia expects Mapper 055 ROM images to first contain the 32 KiB of PRG-ROM at \$8000-\$FFFF, then the data for the 2 KiB of PRG-ROM at \$6000, appropriately padded or repeated to fit into the 16 KiB iNES bank size, for a total of three 16 KiB PRG-ROM banks.
