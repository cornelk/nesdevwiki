---
title: "NES 2.0 Mapper 323"
url: "/wiki/NES_2.0_Mapper_323/"
description: "NES 2.0 Mapper 323 is used for Farid's homebrew 8-in-1 SLROM multicart board. It consists of an MMC1 (clone) with WRAM replaced by an outer bank register. Its UNIF board name is FARID_SLROM_8-IN-1…"
wikiPage: "NES_2_0_Mapper_323"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 323 is used for Farid's homebrew 8-in-1 SLROM multicart board. It consists of an MMC1 (clone) with WRAM replaced by an outer bank register. Its UNIF board name is **FARID_SLROM_8-IN-1** (note underscores).

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    D~7654 3210
      ---------
      .OOO L...
       ||| +---- 1=Disable decoder (lock Outer Bank)
       +++------ Select 128 KiB outer PRG-ROM and CHR-ROM bank

WRAM must be enabled in the MMC1's registers for writes to be effective.

This register is cleared on both warm and cold boot.

## <span id="MMC1_registers_.28.248000-.24FFFF.29"></span>MMC1 registers (\$8000-\$FFFF) {#MMC1_registers_.28.248000-.24FFFF.29}

See [INES Mapper 001](../MMC1/).

## See also {#See_also}

[NesDev discussion](https://forums.nesdev.org/viewtopic.php?t=10632)
