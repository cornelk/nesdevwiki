---
title: "NES 2.0 Mapper 340"
url: "/wiki/NES_2.0_Mapper_340/"
description: "NES 2.0 mapper 340 is used for a 35-in-1 multicart. Its UNIF board name is BMC-K-3036."
wikiPage: "NES_2_0_Mapper_340"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 mapper 340 is used for a *35-in-1* multicart. Its UNIF board name is **BMC-K-3036**.

# Banks {#Banks}

- CPU \$8000-\$BFFF: Switchable 16 KiB inner PRG-ROM bank and switchable 128 KiB outer PRG-ROM bank
- CPU \$C000-\$FFFF: Fixed 16 KiB inner PRG-ROM bank \#7/Mirror of \$8000-\$BFFF and switchable 128 KiB outer PRG-ROM bank
- PPU \$0000-\$1FFF: Unbanked 8 KiB CHR-RAM

## <span id="Address.2FData_Latch_.28.248000-.24FFFF.29"></span>Address/Data Latch (\$8000-\$FFFF) {#Address.2FData_Latch_.28.248000-.24FFFF.29}

    Mask: $8000

    A~FEDC BA98 7654 3210   D~7654 3210
      -------------------     ---- ----
      .... .... ..MO OIII     .... .III
                  || ||||           +++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF in UNROM mode
                  || |+++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF/$C000-$FFFF in NROM-128 mode
                  |+-+|-|- Select 128 KiB outer PRG-ROM bank at CPU $8000-$FFFF
                  +---|-|- Select $C000-$FFFF behavior
                  |   | |   0: UNROM (Inner Bank forced to #7 at CPU $C000-$FFFF)
                  |   | |   1: NROM-128 (CPU $C000-$FFFF mirror of $8000-$BFFF)
                  +---+-+- Select nametable mirroring type
                            0x25: Horizontal
                            all others: Vertical
