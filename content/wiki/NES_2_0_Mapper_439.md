---
title: "NES 2.0 Mapper 439"
url: "/wiki/NES_2.0_Mapper_439/"
description: "NES 2.0 Mapper 439 denotes the YS2309 UxROM-like multicart PCB. Its UNIF MAPRs are BMC-DS-07 and BMC-K86B. All registers set to $00 on reset."
wikiPage: "NES_2_0_Mapper_439"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 439** denotes the **YS2309** UxROM-like multicart PCB. Its UNIF MAPRs are **BMC-DS-07** and **BMC-K86B**. All registers set to \$00 on reset.

## <span id="Latch_Register_.28.248000-.24FFFF.2C_write.29"></span>Latch Register (\$8000-\$FFFF, write) {#Latch_Register_.28.248000-.24FFFF.2C_write.29}

    D~[M..PP PPPP]  Address mask: $8000
       |  ++-++++- PRG A19..A14 when CPU A14=0,
       |           $3F when CPU A14=1
       +---------- Nametable mirroring, 0=Vertical, 1=Horizontal

This is basically an UxROM-like latch register for up to 1 MiB of PRG ROM, plus software-switchable mirroring.

## <span id="Outer_Bank_Register_.28.246000.2C_write.29"></span>Outer Bank Register (\$6000, write) {#Outer_Bank_Register_.28.246000.2C_write.29}

    D~[.PPP ....]  Address mask: $E001
        +++------ PRG A19..A17 if enabled via $6001

## <span id="Mask_Register_.28.246001.2C_write.29"></span>Mask Register (\$6001, write) {#Mask_Register_.28.246001.2C_write.29}

    D~[MPPP ....]  Address mask: $E001
       |||+------- 0: PRG A17 from Latch, 1: PRG A17 from $6000
       ||+-------- 0: PRG A18 from Latch, 1: PRG A18 from $6000
       |+--------- 0: PRG A19 from Latch, 1: PRG A19 from $6000
       +---------- 1: Protect Mirroring bit in Latch from being changed
