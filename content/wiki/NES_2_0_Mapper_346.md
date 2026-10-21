---
title: "NES 2.0 Mapper 346"
url: "/wiki/NES_2.0_Mapper_346/"
description: "NES 2.0 Mapper 346 is used for Kaiser's cartridge conversion of the FDS game Zanac. Its UNIF board name is UNL-KS7012."
wikiPage: "NES_2_0_Mapper_346"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

NES 2.0 Mapper 346 is used for Kaiser's cartridge conversion of the FDS game *Zanac*. Its UNIF board name is **UNL-KS7012**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: Unbanked 8 KiB of PRG-RAM
- CPU \$8000-\$FFFF: Selectable 32 KiB PRG-ROM bank, defaults to \#1
- PPU \$0000-\$1FFF: Unbanked 8 KiB of CHR-RAM

Mirroring is hard-wired.

# Registers {#Registers}

The actual address masks are unknown.

- CPU \$E0A0: Select 32 KiB PRG-ROM bank \#0 at CPU \$8000-\$FFFF.
- CPU \$EE36: Select 32 KiB PRG-ROM bank \#1 at CPU \$8000-\$FFFF.
