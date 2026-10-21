---
title: "NES 2.0 Mapper 380"
url: "/wiki/NES_2.0_Mapper_380/"
description: "NES 2.0 Mapper 380 denotes the 970630C (submapper 0) and KN-35A (submapper 1) multicart circuit board mounting 256-512 KiB of PRG-ROM plus 8 KiB of unbanked CHR-RAM."
wikiPage: "NES_2_0_Mapper_380"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 380** denotes the **970630C** (submapper 0) and **KN-35A** (submapper 1) multicart circuit board mounting 256-512 KiB of PRG-ROM plus 8 KiB of unbanked CHR-RAM.

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    [A~1... ..Om CRQP PpMS]
              || |||| |||+-0: PRG A14=CPU A14
              || |||| |||  1: PRG A14=p
              || |||| ||+- 0: Vertical mirroring
              || |||| ||   1: Horizontal mirroring
              || |+++-++-- PRG A18..A14
              || +-------- 0: CHR-RAM writeable
              ||           1: CHR-RAM write-protected
              |+---------- Submapper 0:
              |            0: PRG A3..A0=CPU A3..A0
              |            1: PRG A3..A0=Solder pad 3-0
              |            Submapper 1:
              |            0: PRG A17=Q if CPU A14=1
              |            1: PRG A17=1 if CPU A14=1 (UOROM)
              +----------- 0: When CPU A14=1: PRG A16..14=111 (UNROM)
                           1: When CPU A14=1: PRG A16..14=PPp
    Power-on value: 0

Effective meaning:

    Bit 9   Bit 8   Bit 0   Meaning
    $200s   $100s   $001s
     (O)     (m)     (S)
      0       0       ?     Switchable inner 16 KiB bank PPp at CPU $8000-$BFFF, fixed inner bank #7 at CPU $C000-$FFFF (UNROM)
      0       1       ?     Switchable inner 16 KiB bank QPPp at CPU $8000-$BFFF, fixed inner bank #15 at CPU $C000-$FFFF (UOROM, submapper 1 only)
      1       ?       0     Switchable 32 KiB inner bank PP at CPU $8000-$FFFF (NROM-256)
      1       ?       1     Switchable 16 KiB inner bank PPp at CPU $8000-$BFFF, mirrored at CPU $C000-$FFFF (NROM-128)

- When the *m* bit is set, PRG A3-A0 are replaced with the values of four solder pads, which when the menu code reads particular ROM locations effectively selects one of up to 16 menus with different game counts (submapper 0 only).
- Because all bits are cleared on reset, CPU \$8000-\$BFFF is set to 16 KiB bank \#0, and \$C000-\$FFFF is set to 16 KiB bank \#7 on reset.

## Similar mappers {#Similar_mappers}

- [INES Mapper 227](../INES_Mapper_227/) and [INES Mapper 242](../INES_Mapper_242/) are incompatible variants with similar functionality.
