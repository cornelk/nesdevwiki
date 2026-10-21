---
title: "NES 2.0 Mapper 334"
url: "/wiki/NES_2.0_Mapper_334/"
description: "NES 2.0 mapper 334 is used for the 5/20-in-1 1993 Copyright multicart."
wikiPage: "NES_2_0_Mapper_334"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 334 is used for the *5/20-in-1 1993 Copyright* multicart.

## <span id="PRG-ROM_Bank_Register_.28.246000.2C_write.29"></span>PRG-ROM Bank Register (\$6000, write) {#PRG-ROM_Bank_Register_.28.246000.2C_write.29}

    Mask: probably $6003

    D~7654 3210
      ---------
      .... .PP.
            ++- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF

## <span id="Unknown_Register_.28.246001.2C_write.29"></span>Unknown Register (\$6001, write) {#Unknown_Register_.28.246001.2C_write.29}

Mask: probably \$6003

Unknown meaning; the menu writes \$01 to it.

## <span id="Jumper_Register_.28.246002.2C_read.29"></span>Jumper Register (\$6002, read) {#Jumper_Register_.28.246002.2C_read.29}

    Mask: probably $6003

    D~7654 3210
      ---------
      OOOO OOOJ
      |||| |||+- Jumper setting
      |||| |||    0: Show 20-in-1 menu (i.e. with repeats)
      |||| |||    1: Show 5-in-1 menu (i.e. without repeats)
      ++++-+++-- Open bus

The "hardware test" menu (seen when pressing SELECT+A+B during the menu) expects the jumper at zero.

## MMC3-compatible registers {#MMC3-compatible_registers}

Mask: \$E001

See [MMC3](../MMC3/).

## Notes {#Notes}

The MMC3 clone is only used for Mirroring and CHR-ROM banking; the PRG-ROM bank comes entirely from an additional register. As they use the MMC3 clones's WRAM interface, accessing the \$600x registers requires enabling and not write-protecting WRAM in the MMC's \$A001 register.
