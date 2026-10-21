---
title: "NES 2.0 Mapper 362"
url: "/wiki/NES_2.0_Mapper_362/"
description: "NES 2.0 Mapper 362 is used for an early VRC4-clone-based multicart PCB (830506C) from J.Y. Company, preceding the J.Y. Company ASIC:"
wikiPage: "NES_2_0_Mapper_362"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 362 is used for an early VRC4-clone-based multicart PCB (830506C) from J.Y. Company, preceding the [J.Y. Company ASIC](../J.Y._Company_ASIC/):

- *1995 Super HiK 4-in-1 (JY-005)*

The [VRC4](../VRC2_and_VRC4/) clone is wired in the VRC4f configuration, i.e. address 0x001s to A0 and address 0x002s to A1. As each individual game on that multicart only has 128 KiB of PRG/CHR-ROM, the two highest CHR-ROM bank bits are repurposed as outer bank bits:

### <span id="CHR_Select_0_high_.28.24B001.29"></span>CHR Select 0 high (\$B001) {#CHR_Select_0_high_.28.24B001.29}

    7  bit  0
    ---------
    ...O OLLL
       | |+++- Select D4..D6 KiB of 1 KiB CHR-ROM bank at PPU $0000-$03FF
       +-+---- Select outer 128 KiB PRG/CHR-ROM bank when PPU $0000-$03FF is accessed

- The high CHR Select register (\$B004) is ignored.
- The other seven high CHR bank select registers continue the pattern for the other seven 1 KiB PPU banks.

## Notes {#Notes}

- As with [SOROM, SUROM and SXROM](../MMC1/#SOROM.2C_SUROM_and_SXROM), if the several CHR bank registers do not specify the same outer bank, then PRG-ROM will would be bankswitched as the PPU renders.
- As with [NES 2.0 Mapper 298](../NES_2.0_Mapper_298/), the VRC4 clone's F003 register only acknowledges IRQs but does not move the 'A' control bit to the 'E' control bit.
