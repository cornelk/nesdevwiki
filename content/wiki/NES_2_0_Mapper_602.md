---
title: "NES 2.0 Mapper 602"
url: "/wiki/NES_2.0_Mapper_602/"
description: "NES 2.0 Mapper 602 denotes the Subor SB-2000 UM6576-based educational computer. The first 512 KiB of its PRG address space are PRG-ROM, the second 512 KiB are PRG-DRAM. There are 256 KiB of CHR-RAM,…"
wikiPage: "NES_2_0_Mapper_602"
wikiCategories:
  - name: "UM6576/UM6578 mappers"
    url: "/category/UM6576/UM6578_mappers/"
---

**NES 2.0 Mapper 602** denotes the Subor SB-2000 UM6576-based educational computer. The first 512 KiB of its PRG address space are PRG-ROM, the second 512 KiB are PRG-DRAM. There are 256 KiB of CHR-RAM, selectable via an outer bank register in 32 KiB amounts.

# PRG address space layout {#PRG_address_space_layout}

- First 512 KiB (4 KiB PRG bank# \$00-\$7F): PRG-ROM
- Second 512 KiB (4 KiB PRG bank# \$80-\$FF): DRAM

# Extra Registers {#Extra_Registers}

## Floppy Disk Controller {#Floppy_Disk_Controller}

- \$4202: Digital Output Register (R)
- \$4204: Main Status Register (R)
- \$4205: Data Register (R/W)
- \$4207: Digital Input Register (R), Configuration Control Register (W)

The floppy disk controller is Intel 82077-compatible.

## Miscellaneous {#Miscellaneous}

- \$4300: Outer 32 KiB CHR-RAM bank (out of 256 KiB)
- \$4301: Reset register: Writing \$55 resets the UM6576 to normal NES-compatible mode. Used to run NES games from floppy disk.
- \$4302: LPC speech chip, same bit assignents as in [mapper 518](../NES_2.0_Mapper_518/#LPC_Speech_Chip_.28.245300.29).
