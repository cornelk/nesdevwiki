---
title: "NES 2.0 Mapper 545"
url: "/wiki/NES_2.0_Mapper_545/"
description: "NES 2.0 Mapper 545 denotes the ST-80 PCB, used by the 壽石組合 4-in-1 Super Combo 1216k multicart. It contains three 128K PRG-ROM+128K CHR-ROM TLROM games, plus one 256K+128K one (Super Mario…"
wikiPage: "NES_2_0_Mapper_545"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 545** denotes the **ST-80** PCB, used by the *壽石組合 4-in-1 Super Combo 1216k* multicart. It contains three 128K PRG-ROM+128K CHR-ROM [TLROM](../TxROM/) games, plus one 256K+128K one (Super Mario Bros. 3). It mounts two PRG-ROM chips, a 512 KiB one, and a 128 KiB one. The first, 512 KiB, one contains the PRG-ROM data for the first three games that have 128 KiB of PRG-ROM per game, plus the *second* 128 KiB half of the game that has 256 KiB of PRG-ROM. The second, 128 KiB, PRG-ROM chip contains the *first* 128 KiB half of the game that has 256 KiB of PRG-ROM. The .NES file contains the 512 KiB of PRG-ROM of the first chip, followed by the 128 KiB of the second chip, for a total of 640 KiB of PRG-ROM.

# Outer bank register {#Outer_bank_register}

    D~[.... DCBA] A~[1111 ...1 ..1. ....] (Address & $F120)==$F120
            ||++- PRG/CHR A18..A17 within first ROM chip
            ++--- 0: force CHR A16 high
            |     1: normal CHR A16 mode
            +---- 0: 128 KiB PRG mode
                  1: 256 KiB PRG mode: Use second ROM chip if MMC3 A17=0
    Power-on value: $00

The "force CHR A16 high" business is only used during the menu.
