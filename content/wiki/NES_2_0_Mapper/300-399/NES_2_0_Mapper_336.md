---
title: "NES 2.0 Mapper 336"
url: "/wiki/NES_2.0_Mapper_336/"
description: "NES 2.0 mapper 336 is used for multicarts containing several UNROM games. Its UNIF board name is BMC-K-3046."
wikiPage: "NES_2_0_Mapper_336"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 mapper 336** is used for multicarts containing several UNROM games. Its UNIF board name is **BMC-K-3046**.

**Submapper 0:**

- 11-in-1 (K-3046)
- 2-in-1 (YH-217)
- 4-in-1 (YH-4109)

**Submapper 1 (M1K30 PCB):**

- Top Gun III 2-in-1

**Submapper 2:**

- 2-in-1 (YH-218)

# Banks {#Banks}

- CPU \$8000-\$BFFF: Switchable 16 KiB inner PRG-ROM bank and switchable 128 KiB outer PRG-ROM bank
- CPU \$C000-\$FFFF: Fixed 16 KiB inner PRG-ROM bank \#7 and switchable 128 KiB outer PRG-ROM bank
- PPU \$0000-\$1FFF: Unbanked 8 KiB CHR-RAM

## <span id="Data_Latch_.28.24C000-.24FFFF.29"></span>Data Latch (\$C000-\$FFFF) {#Data_Latch_.28.24C000-.24FFFF.29}

    Mask: $8000

    D~7654 3210
      ---------
      ..OO OIII
        || |+++- Select 16 KiB inner PRG-ROM bank at CPU $8000-$BFFF
        ++-+---- Select 128 KiB outer PRG-ROM bank at CPU $8000-$FFFF
        |  +---- Mirroring (Submapper 2), 1=H, 0=V
        +------- Mirroring (Submappers 0/1), 1=H, 0=V

## Note {#Note}

**Submappers 0** and **2** have no bus conflicts. **Submapper 1** does: D0-D2 are affected by normal AND-type bus conflicts, while D3 has a resistor in the signal path to ensure that the ROM bit always "wins", similar to [INES Mapper 144](../INES_Mapper_144/). This behavior is necessary for proper function.
