---
title: "NES 2.0 Mapper 337"
url: "/wiki/NES_2.0_Mapper_337/"
description: "NES 2.0 mapper 337 is used for a 12-in-1 Game Card multicart. Its UNIF board name is BMC-CTC-12IN1."
wikiPage: "NES_2_0_Mapper_337"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 337 is used for a *12-in-1 Game Card* multicart. Its UNIF board name is **BMC-CTC-12IN1**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB fixed PRG-ROM bank \#1
- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB switchable PRG-ROM bank
- PPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM

# Registers {#Registers}

## <span id="Outer_Bank_Register_.28.248000-.24BFFF.29"></span>Outer Bank Register (\$8000-\$BFFF) {#Outer_Bank_Register_.28.248000-.24BFFF.29}

    Mask: $C000

    D~7654 3210
      ---------
      PPMB B...
      |||+-+---- Select 128 KiB outer PRG-ROM bank
      ||+------- Select nametable mirroring type
      ||          0: Vertical
      ||          1: Horizontal
      ++-------- Select PRG-ROM banking mode
                  0: NROM-128
                  1: NROM-256 (inner bank bit 0 replaced with CPU A14)
                  2: UNROM (inner bank fixed to #7 at CPU $C000-$FFFF)
                  3: same as 2

CHR-RAM is write-protected in the NROM modes.

## <span id="Inner_Bank_Register_.28.24C000-.24FFFF.29"></span>Inner Bank Register (\$C000-\$FFFF) {#Inner_Bank_Register_.28.24C000-.24FFFF.29}

    Mask: $C000

    D~7654 3210    A~FEDC BA98 7654 3210
      ---------      -------------------
      .... .BBB      11W. .... .... ....
            |||        +----------------- 1=Write-protect CHR-RAM
            +++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF/$C000-$FFFF

# Similar Mappers {#Similar_Mappers}

[NES 2.0 Mapper 350](../NES_2.0_Mapper_350/) is a variant for larger ROM sizes of this mapper.
