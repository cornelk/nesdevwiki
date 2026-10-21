---
title: "NES 2.0 Mapper 430"
url: "/wiki/NES_2.0_Mapper_430/"
description: "NES 2.0 Mapper 430 denotes the MMC3-based 831031C/T-308 multicart circuit board."
wikiPage: "NES_2_0_Mapper_430"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 430** denotes the MMC3-based **831031C/T-308** multicart circuit board.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    A~FEDC BA98 7654 3210
      -------------------
      011. .... .... MNBb
                     ||++- PRG/CHR A18..A17
                     |+--- 0: CHR A17=MMC3 CHR A17 (256 KiB inner CHR bank)
                     |     1: CHR A17=b (128 KiB inner CHR bank)
                     +---- 0: Normal MMC3 PRG banking mode
                           1: GNROM-like PRG banking mode

The MMC3 can only select 128 KiB of PRG-ROM; the b bit always selects PRG A17.

GNROM mode means that the MMC3's CPU A14 input is held low, so MMC3 registers 6 and 7 apply to both \$8000/\$A000 and \$C000/\$E000,

As it uses the MMC3 clones's WRAM interface, writing to the outer bank register requires enabling and not write-protecting WRAM in the MMC's \$A001 register.

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.29"></span>MMC3-compatible registers (\$8000-\$FFFF) {#MMC3-compatible_registers_.28.248000-.24FFFF.29}

    Mask: $E001

    See MMC3.
