---
title: "NES 2.0 Mapper 394"
url: "/wiki/NES_2.0_Mapper_394/"
description: "NES 2.0 Mapper 394 denotes Realtec's HSK007 multicart PCB. Unusually, it can be set to function either as an MMC3 clone or as a J.Y. Company ASIC. Two submapper-distinguished variants exist that…"
wikiPage: "NES_2_0_Mapper_394"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 394** denotes Realtec's **HSK007** multicart PCB. Unusually, it can be set to function either as an [MMC3](../MMC3/) clone or as a [J.Y. Company ASIC](../J.Y._Company_ASIC/). Two submapper-distinguished variants exist that differ in whether PRG and CHR A18 share the same register bit.

## <span id="Mode_Register_.28.245009.2C_write.29"></span>Mode Register (\$5009, write) {#Mode_Register_.28.245009.2C_write.29}

    Mask: unknown, but A3 must be set

    D~7654 3210
      ---------
      ...M MCBA
         | |||+- PRG A18
         | |||+- CHR A18 (submapper 0)
         | ||+-- PRG A19
         | ||+-- CHR A19
         | |+--- CHR A18 (submapper 1)
         +-+---- Mapper Mode
                 0: BNROM
                 1: MMC3
                 2: J.Y. Company ASIC
                 3: Same as 2

    Power-on value: $0F

## <span id="Outer_Bank_Register_.28.24500B.2C_write.29"></span>Outer Bank Register (\$500B, write) {#Outer_Bank_Register_.28.24500B.2C_write.29}

    Mask: unknown, but A3 must be set

    D~7654 3210
      ---------
      cC.p PBB.
      || | |++-- PRG A16..A15 in BNROM mode
      || | +---- PRG A17 if p=0
      || +------ PRG A17 in MMC3 mode
      ||          0: From Outer Bank Register (128 KiB inner bank)
      ||          1: From MMC3 (256 KiB inner bank)
      |+-------- CHR A17 if c=0
      +--------- CHR A17 in MMC3 mode
                  0: From Outer Bank Register (128 KiB inner bank)
                  1: From MMC3 (256 KiB inner bank)

    Power-on value: $90

In J.Y. Company ASIC Mode, PRG/CHR A17 always come from the J.Y. Company ASIC.
