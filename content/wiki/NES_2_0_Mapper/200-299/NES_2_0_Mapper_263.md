---
title: "NES 2.0 Mapper 263"
url: "/wiki/NES_2.0_Mapper_263/"
description: "NES 2.0 Mapper 263 is used for an MMC3 clone with scrambled mapper address and data bits. Its UNIF board name is UNL-KOF97. It is used for two games by Rex Soft:"
wikiPage: "NES_2_0_Mapper_263"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

NES 2.0 Mapper 263 is used for an [MMC3](../MMC3/) clone with scrambled mapper address and data bits. Its UNIF board name is **UNL-KOF97**. It is used for two games by Rex Soft:

- *Boogerman II: The Final Adventure*
- *The King of Fighters '97*

## Registers {#Registers}

A12 is used for the MMC3's A0:

Mask: \$F000

    Mapper 263    Regular MMC3
    address       address
    ----------    ------------
    8000          8000
    9000          8001
    A000          A000
    B000          A001
    C000          C000
    D000          C001
    E000          E000
    F000          E001

The data written to the MMC3 clone has the following bit order compared to a regular MMC3:

    Written bit 76543210
                --------
    Actual bit  76243051

Which means that to get from the written data to the actual data, the following operation must be performed:

    Val =(Val &0xD8) | ((Val &0x20) >>4) | ((Val &4) <<3) | ((Val &2) >>1) | ((Val &1) <<2);

## Note {#Note}

- Most versions of *The King of Fighters '97* are hacked for [Mapper 4](../MMC3/). Several of them also fail to disable the [frame IRQ](../APU_Frame_Counter/), freezing on emulators that accurately [power-on with the frame IRQ enabled](../CPU_power_up_state/).
