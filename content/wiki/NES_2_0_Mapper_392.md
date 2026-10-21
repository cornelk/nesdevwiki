---
title: "NES 2.0 Mapper 392"
url: "/wiki/NES_2.0_Mapper_392/"
description: "NES 2.0 Mapper 392 denotes the 00202650 PCB, used on an 8-in-1 multicart. It uses an MMC3 clone with both CHR-ROM and CHR-RAM and an outer bank register that is implemented using a GAL16v8."
wikiPage: "NES_2_0_Mapper_392"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 392** denotes the **00202650** PCB, used on an 8-in-1 multicart. It uses an [MMC3](../MMC3/) clone with both CHR-ROM and CHR-RAM and an outer bank register that is implemented using a GAL16v8.

The circuit board mounts two PRG-ROM chips each holding 512 KiB of game data (U2/U3), a third PRG-ROM chip with 32 KiB of menu data (U1), as well as two 512 KiB CHR-ROM chips holding game graphics (U10/U4) and 8 KiB of CHR-RAM (U6). The NES 2.0 file holds these chips' data in the order U2-U3-U1-U10-U4.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    D~7654 3210
      ---------
      ...M .BBB
         |  +++- PRG/CHR A19..A17 in game mode
         +------ 0: Menu Mode
                 1: Game Mode, lock register

## Notes {#Notes}

- In menu mode, the 32 KiB PRG-ROM chip is mapped into CPU address space, and 8 KiB of unbanked CHR-RAM is mapped into PPU address space.
- In game mode, the 1 MiB PRG-ROM and 1 MiB CHR-ROM chips are mapped and banked by the MMC3 clone within the 128 KiB PRG/CHR window selected by the outer bank register.
- The outer bank register functions independently of the MMC3's WRAM register.
- See also: [NesDev forum post with krzysiobal's analysis](https://forums.nesdev.org/viewtopic.php?f=9&t=19390&p=243217)
