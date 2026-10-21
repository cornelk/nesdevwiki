---
title: "NES 2.0 Mapper 535"
url: "/wiki/NES_2.0_Mapper_535/"
description: "NES 2.0 Mapper 535 is used for Whirlwind Manu's cartridge conversion (cartrige code LH53) of the Famicom Disk System Game 謎の村雨城 (Nazo no Murasamejō). Its UNIF board name is UNL-LH53."
wikiPage: "NES_2_0_Mapper_535"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with fixed-timing cycle IRQs"
    url: "/category/Mappers_with_fixed-timing_cycle_IRQs/"
---

NES 2.0 Mapper 535 is used for Whirlwind Manu's cartridge conversion (cartrige code LH53) of the Famicom Disk System Game *謎の村雨城* (Nazo no Murasamejō). Its UNIF board name is **UNL-LH53**.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG-ROM bank
- CPU \$8000-\$FFFF: 32 KiB fixed PRG-ROM bank \#3, except where overlaid by PRG-RAM
- CPU \$B800-\$D7FF: 8 KiB fixed PRG-RAM bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="IRQ_register_.28.24E000-.24EFFF.2C_write.29"></span>IRQ register (\$E000-\$EFFF, write) {#IRQ_register_.28.24E000-.24EFFF.2C_write.29}

    Mask: probably $F000

    D~7654 3210
      ---------
      .... ..E.
             +-- 1= Enable counting

Writing to this register acknowledges a pending IRQ, and resets the counter to zero. When counting is enabled, the counter increases on every M2 cycle; upon reaching a value of 7560, an IRQ is generated (according to [FCEUX source code](https://github.com/TASVideos/fceux/blob/master/src/boards/lh53.cpp)).

## <span id="PRG_bank_register_.28.24F000-.24FFFF.2C_write.29"></span>PRG bank register (\$F000-\$FFFF, write) {#PRG_bank_register_.28.24F000-.24FFFF.2C_write.29}

Mask: probably \$F000

Sets the 8 KiB PRG-ROM bank at CPU \$6000-\$7FFF.

# Notes {#Notes}

- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Whirlwind Manu cartridge did not.
- Game progress cannot be saved.
