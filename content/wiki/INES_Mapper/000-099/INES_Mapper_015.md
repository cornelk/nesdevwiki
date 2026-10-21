---
title: "INES Mapper 015"
url: "/wiki/INES_Mapper_015/"
description: "iNES Mapper 015 denotes the K-1029 and K-1030P circuit boards, used for the 100-in-1 Contra Function 16 and 168-in-1 New Contra Function 16 multicarts."
wikiPage: "INES_Mapper_015"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 015** denotes the **K-1029** and **K-1030P** circuit boards, used for the *100-in-1 Contra Function 16* and *168-in-1 New Contra Function 16* multicarts.

# Banks {#Banks}

- CPU \$8000-\$FFFF: Switchable PRG-ROM banks (8/16/32 KiB)
- PPU \$0000-\$1FFF: Unbanked 8 KiB of CHR-RAM that can be write-protected
- Nametable mirroring: Switchable horizontal/vertical settings

# <span id="Address_and_Data_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address and Data Latch (\$8000-\$FFFF, write) {#Address_and_Data_Latch_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    A~FEDC BA98 7654 3210  D~7654 3210
      -------------------    ---------
      1... .... .... ..SS    pMPP PPPP
                       ||    ||++-++++- PRG A19..A14
                       ||    |+-------- Nametable mirroring
                       ||    |           0: Vertical
                       ||    |           1: Horizontal
                       ||    +--------- PRG A13 if SS=2, ignored otherwise
                       ++-------------- PRG banking mode
                                         0: NROM-256 (PRG A14=CPU A14)
                                         1: UNROM    (PRG A14..16=111 when CPU A14=1)
                                         2: NROM-64  (PRG A13=p)
                                         3: NROM-128
    Power-on and reset value: All bits clear

CHR-RAM is write-protected in PRG banking modes 0 and 3, and write-enabled in modes 1 and 2.

# Use for mapper hacks {#Use_for_mapper_hacks}

The only known two cartridges using this mapper are the two listed multicarts. All other known ROM files claiming to use mapper 15 have been identified as being mapper hacks, usually from mappers [164](../INES_Mapper_164/) and [227](../INES_Mapper_227/). To run these mapper-hacked games, emulators must not enforce CHR-RAM write-protection even though an actual K-1029 PCB would, and must provide 8 KiB of PRG-RAM at CPU \$6000-\$7FFF even though an actual K-1029 PCB would not.

# See also {#See_also}

- [schematic RE'd](https://forums.nesdev.org/viewtopic.php?p=192828#p192828) by krzysiobal
