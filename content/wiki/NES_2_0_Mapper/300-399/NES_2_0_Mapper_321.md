---
title: "NES 2.0 Mapper 321"
url: "/wiki/NES_2.0_Mapper_321/"
description: "NES 2.0 Mapper 321 denotes the 820310 MMC3-based multicart circuit board, UNIF MAPR BMC-820310-C. It is basically NES 2.0 Mapper 287 looking at differently-ordered address bits."
wikiPage: "NES_2_0_Mapper_321"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 321** denotes the **820310** [MMC3](../MMC3/)-based multicart circuit board, [UNIF](../UNIF/) MAPR **BMC-820310-C**. It is basically [NES 2.0 Mapper 287](../NES_2.0_Mapper_287/) looking at differently-ordered address bits.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    A~[011. .... ..PP MB..], A001.7=1, A001.6=0
                   || |+---- PRG/CHR A17
                   || +----- 0=PRG A16..A13 from MMC3
                   ||        1=PRG A16..A15 from PP,
                   ||          PRG A14..A13=CPU A14..A13
                   ||          (NROM-256-like PRG banking)
                   ++------- PRG A16..A15 is M=1
