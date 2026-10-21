---
title: "NES 2.0 Mapper 493"
url: "/wiki/NES_2.0_Mapper_493/"
description: "NES 2.0 Mapper denotes the prototype board used for the Maxivision 30 Super Games prototype multicart."
wikiPage: "NES_2_0_Mapper_493"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper** denotes the prototype board used for the *Maxivision 30 Super Games* prototype multicart.

# Overview {#Overview}

- CPU \$8000-\$FFFF: 32 KiB switchable window into 1 MiB PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB switchable window into 1 MiB CHR-ROM
- Nametable arrangement: Hard-wired to Horizontal ("Vertical mirroring")

# Registers {#Registers}

There are two writable registers in the CPU \$6000-\$FFFF address range.

**Outer PRG/CHR Bank Register**

    D~[.... BBBB] A~[011. .... .... ....] (Address &0xE000)==0x6000
            ++++- PRG/CHR A19..A16 (64 KiB outer PRG/CHR bank)

**Inner PRG/CHR Bank Register**

    D~[.CCC ...P] A~[1... .... .... ....] (Address &0x8000)==0x8000, bus conflicts
        |||    +- PRG A15 (32 KiB inner PRG bank)
        +++------ CHR A15..A13 (8 KiB inner CHR bank)

# Comparison to other mappers {#Comparison_to_other_mappers}

- It is [INES Mapper 011](../Color_Dreams/) restricted to 64 KiB PRG/CHR sizes per game, with an additional outer bank register.
- It is similar to [INES Mapper 046](../INES_Mapper_046/), with the difference that the same bits control both the outer PRG and outer CHR banks.
