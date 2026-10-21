---
title: "INES Mapper 020"
url: "/wiki/INES_Mapper_020/"
description: "Mapper 20 is defined as reserved for the Famicom Disk System (FDS) in official iNES specs. The mapper number is mostly intended for being internally used by emulators (to indicate to themselves that…"
wikiPage: "INES_Mapper_020"
wikiCategories:
  - name: "Bad iNES Mappers"
    url: "/category/Bad_iNES_Mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

Mapper 20 is defined as reserved for the Famicom Disk System ([FDS](../Family_Computer_Disk_System/)) in official [iNES](../INES/) specs. The mapper number is mostly intended for being internally used by emulators (to indicate to themselves that they are emulating the FDS hardware).

Actual FDS game disk images normally use the [FDS file format](../FDS_file_format/). (`.fds`)

The mapper number could be theoretically used for the [FDS BIOS](../FDS_BIOS/) ROM image, but this isn't done in practice. The BIOS is typically stored as a raw 8 KiB dump named `disksys.rom` (without an iNES header and thus no mapper number), though it is sometimes dumped as [NROM](../NROM/) via cartridge dumping hardware.
