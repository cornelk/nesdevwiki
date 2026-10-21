---
title: "NES 2.0 Mapper 403"
url: "/wiki/NES_2.0_Mapper_403/"
description: "NES 2.0 Mapper 403 denotes the 89433 circuit board with up to 1 MiB PRG-ROM and 32 KiB of CHR-RAM, bankable with 8 KiB granularity."
wikiPage: "NES_2_0_Mapper_403"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 403** denotes the **89433** circuit board with up to 1 MiB PRG-ROM and 32 KiB of CHR-RAM, bankable with 8 KiB granularity.

- *Tetris Family - 玩家 19-in-1 智瑟實典* (NO-1683)
- *Sachen Superpack* (versions A-C)

## <span id="PRG_bank_register_.28.244100.29"></span>PRG bank register (\$4100) {#PRG_bank_register_.28.244100.29}

    Mask: $E103

    D~[.PPP PPp.]
        +++-+++--- PRG A19..A14

    Power-on value: $00

## <span id="CHR_bank_register_.28.244101.2C_.248000-.24FFFF.29"></span>CHR bank register (\$4101, \$8000-\$FFFF) {#CHR_bank_register_.28.244101.2C_.248000-.24FFFF.29}

    Mask: $E103, $8000

    D~[.... ..CC]
              ++-- CHR A14..A13

    Power-on value: $00

The latch at \$8000-\$FFFF is enabled/disabled via the Mode register's C bit.

## <span id="Mode_register_.28.244102.29"></span>Mode register (\$4102) {#Mode_register_.28.244102.29}

    Mask: $E103

    D~[...H .C.N]
          |  | +- 0: PRG A14=CPU A14 (NROM-256)
          |  |    1: PRG A14=p (NROM-128)
          |  +--- 0: Latch at $8000-$FFFF disabled
          |       1: Latch at $8000-$FFFF enabled
          +------ 0: Vertical mirroring
                  1: Horizontal mirroring

    Power-on value: $00
