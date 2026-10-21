---
title: "NES 2.0 Mapper 333"
url: "/wiki/NES_2.0_Mapper_333/"
description: "NES 2.0 mapper 333 is used for the New Star Super 8-in-1 multicart. Its UNIF board name are BMC-8-IN-1 and BMC-NEWSTAR-GRM070-8IN1."
wikiPage: "NES_2_0_Mapper_333"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 333 is used for the *New Star Super 8-in-1* multicart. Its UNIF board name are **BMC-8-IN-1** and **BMC-NEWSTAR-GRM070-8IN1**.

## <span id="Outer_Bank_and_PRG_Banking_Mode_Register_.28.249000-.249FFF.2C_.24B000-.24BFFF.2C_.24D000-.24DFFF.2C_.24F000-.24FFFF.2C_write.29"></span>Outer Bank and PRG Banking Mode Register (\$9000-\$9FFF, \$B000-\$BFFF, \$D000-\$DFFF, \$F000-\$FFFF, write) {#Outer_Bank_and_PRG_Banking_Mode_Register_.28.249000-.249FFF.2C_.24B000-.24BFFF.2C_.24D000-.24DFFF.2C_.24F000-.24FFFF.2C_write.29}

    Mask: $9000

    D~7654 3210
      ---------
      ...P OOII
         | ||++- Select inner 32 KiB PRG-ROM bank at CPU $8000-$FFFF if P=0
         | ++--- Select 128 KiB outer PRG-ROM and CHR-ROM bank
         +------ Select PRG Banking Mode
                  0: NROM-256
                  1: MMC3

## <span id="MMC3-compatible_registers_.28.248000-.248FFF.2C_.24A000-.24AFFF.2C_.24C000-.24CFFF.2C_.24E000-.24EFFF.2C_write.29"></span>MMC3-compatible registers (\$8000-\$8FFF, \$A000-\$AFFF, \$C000-\$CFFF, \$E000-\$EFFF, write) {#MMC3-compatible_registers_.28.248000-.248FFF.2C_.24A000-.24AFFF.2C_.24C000-.24CFFF.2C_.24E000-.24EFFF.2C_write.29}

    Mask: $F001

    See MMC3.
