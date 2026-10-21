---
title: "NES 2.0 Mapper 393"
url: "/wiki/NES_2.0_Mapper_393/"
description: "NES 2.0 Mapper 393 denotes the 820720C multicart PCB."
wikiPage: "NES_2_0_Mapper_393"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 393** denotes the **820720C** multicart PCB.

# Registers {#Registers}

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.2C_write.29"></span>MMC3-compatible registers (\$8000-\$FFFF, write) {#MMC3-compatible_registers_.28.248000-.24FFFF.2C_write.29}

These registers function identically to a normal [MMC3](../MMC3/), except in UNROM/BNROM mode.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    A~FEDC BA98 7654 3210
      --------- ---------
      011. .... ..MM CPPP
                  || |+++- PRG A19..A17
                  || +---- CHR Mode
                  ||        0: CHR-ROM with MMC3 banking
                  ||        1: 8 KiB unbanked CHR-RAM
                  ++------ PRG Mode
                            0: MMC3 (PRG A13-A16)
                            1: same as 0
                            2: BNROM (PRG A15-A16 from MMC3 R6,
                               PRG A13-A14=CPU A13-A14)
                            3: UNROM (PRG A14-A16 from $8000-$FFFF latch,
                               PRG A13=CPU A13) 

Selected by the MMC3's PRG-RAM /CE, WRAM must be enabled and writable in the MMC3's A001 register for Outer Bank Register writes to be recognized.
