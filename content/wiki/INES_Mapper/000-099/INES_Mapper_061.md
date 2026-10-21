---
title: "INES Mapper 061"
url: "/wiki/INES_Mapper_061/"
description: "iNES Mapper 61 is used for three different multicart PCBs:"
wikiPage: "INES_Mapper_061"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**iNES Mapper 61** is used for three different multicart PCBs:

- **GS-2017**: 8 KiB of unbanked CHR-RAM (submapper irrelevant): *方塊外傳 9合1 - Tetris Family 9-in-1*
- NTDEC **0324**: 128 KiB of banked CHR-ROM (**submapper 0**): *HQ 高品質合卡 15-in-1*
- NTDEC **BS-N032**: 256 KiB of banked CHR-ROM (**submapper 1**): *32-in-1*

## <span id="Address_Latch_.28.248000-.24FFFF.2C_write.29"></span>Address Latch (\$8000-\$FFFF, write) {#Address_Latch_.28.248000-.24FFFF.2C_write.29}

    A~[1... CCCC McpN PPPP]
            |||| |||| ++++- PRG A18..A15
            |||| |||+------ 0: PRG A14=CPU A14 (NROM-256)
            |||| |||        1: PRG A14=p (NROM-128)
            |||| ||+------- PRG A14 if M=1
            |||  |+-------- CHR A13 (submapper 1 only)
            |||| +--------- Mirroring, 0=Vertical, 1=Horizontal
            ++++----------- CHR A16..A13 (submapper 0)
                            CHR A17..A14 (submapper 1)
