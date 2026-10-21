---
title: "NES 2.0 Mapper 384"
url: "/wiki/NES_2.0_Mapper_384/"
description: "NES 2.0 Mapper 384 denotes the L1A16 circuit board, used on a 4-in-1 multicart containing Crisis Force, among other games. It uses a VRC4 clone in the VRC4e configuration (VRC4 A0=CPU A2, VRC4 A1=CPU…"
wikiPage: "NES_2_0_Mapper_384"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 384** denotes the **L1A16** circuit board, used on a *4-in-1* multicart containing *Crisis Force*, among other games. It uses a [VRC4](../VRC2_and_VRC4/) clone in the VRC4e configuration (VRC4 A0=CPU A2, VRC4 A1=CPU A3) and has 2 KiB of WRAM, mirrored once, in the CPU \$6000-\$6FFF area. The outer bank register overlays this area and contains a Lock bit to prevent Crisis Force's WRAM writes from changing the outer bank.

## <span id="Outer_Bank_Register_.28.246800-.246FFF.2C_write.29"></span>Outer Bank Register (\$6800-\$6FFF, write) {#Outer_Bank_Register_.28.246800-.246FFF.2C_write.29}

    Mask: $F800

    D~7654 3210
      ---------
      .... LbBB
           ||++- PRG/CHR A17/A18 (128 KiB outer bank)
           |+--- unused, but could be PRG/CHR A19 on a larger multicart
           +---- 1: Lock outer bank

    Power-on value: $00
