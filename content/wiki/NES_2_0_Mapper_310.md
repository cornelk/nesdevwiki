---
title: "NES 2.0 Mapper 310"
url: "/wiki/NES_2.0_Mapper_310/"
description: "NES 2.0 Mapper 310 denotes the K-1053 circuit board, used for the 1200-in-1 New World Multi multicart. It is an extension of INES Mapper 015 that supports more than 1 MiB of PRG ROM, and 32 KiB…"
wikiPage: "NES_2_0_Mapper_310"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 310** denotes the **K-1053** circuit board, used for the *1200-in-1 New World Multi* multicart. It is an extension of [INES Mapper 015](../INES_Mapper_015/) that supports more than 1 MiB of PRG ROM, and 32 KiB rather than 8 KiB of CHR RAM.

# Banks {#Banks}

- CPU \$8000-\$FFFF: Switchable PRG-ROM banks (8/16/32 KiB)
- PPU \$0000-\$1FFF: Switchable 8 KiB (out of 32 KiB) of CHR-RAM that can be write-protected
- Nametable mirroring: Switchable horizontal/vertical settings

# <span id="Data_Latch_.28.248000-.24BFFF.2C_write.29"></span>Data Latch (\$8000-\$BFFF, write) {#Data_Latch_.28.248000-.24BFFF.2C_write.29}

    Mask: $C000

    D~7654 3210
    ---------
    pMPP PPPP
    ||++-++++- PRG A19..A14
    |+-------- Nametable mirroring
    +--------- PRG A13 if SS=2, ignored otherwise
     Power-on and reset value: All bits clear

# <span id="Address_and_Data_Latch_.28.24C000-.24FFFF.2C_write.29"></span>Address and Data Latch (\$C000-\$FFFF, write) {#Address_and_Data_Latch_.28.24C000-.24FFFF.2C_write.29}

    Mask: $C000

    A~FEDC BA98 7654 3210  D~7654 3210
      -------------------    ---------
      11.. .... .... PPSS    .... ..CC
                     ||||           ++- CHR A14..A13
                     ||++-------------- PRG banking mode
                     ||                  0: NROM-256 (PRG A14=CPU A14)
                     ||                  1: UNROM    (PRG A14..16=111 when CPU A14=1)
                     ||                  2: NROM-64  (PRG A13=p)
                     ||                  3: NROM-128
                     ++---------------- PRG A21..20
    Power-on and reset value: All bits clear

CHR-RAM is write-protected in PRG banking modes 0 and 3, and write-enabled in modes 1 and 2.
