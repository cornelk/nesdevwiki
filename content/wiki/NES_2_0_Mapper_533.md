---
title: "NES 2.0 Mapper 533"
url: "/wiki/NES_2.0_Mapper_533/"
description: "NES 2.0 Mapper 533 is used for the Sachen 3014 board, used for the game 動動腦 II: 國中英文(一) (Dòngdòngnǎo II: Guózhōng Yīngwén (I), also known as Middle School English II, SA-003).…"
wikiPage: "NES_2_0_Mapper_533"
wikiCategories:
  - name: "CNROM-like mappers"
    url: "/category/CNROM-like_mappers/"
---

NES 2.0 Mapper 533 is used for the **Sachen 3014** board, used for the game *動動腦 II: 國中英文(一)* (Dòngdòngnǎo II: Guózhōng Yīngwén (I), also known as *Middle School English II*, SA-003). It's a [CNROM](../CNROM/)-like board with the added ability to read back the latch content for protection purposes.

# Banks {#Banks}

- CPU \$8000-\$FFFF: Fixed 32 KiB PRG-ROM bank
- PPU \$0000-\$1FFF: Switchable 8 KiB CHR-ROM bank
- Mirroring: hard-wired

# Registers {#Registers}

## <span id="CHR_Bank_Select_and_Latch_Write_.28.248000-.24FFFF.2C_write.29"></span>CHR Bank Select and Latch Write (\$8000-\$FFFF, write) {#CHR_Bank_Select_and_Latch_Write_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    D~7654 3210
      ---------
      LLLl ....
      ++++------ Set latched value that can be read back at $E000-$EFFF
         +------ Select 8 KiB CHR-ROM bank at PPU $0000-$1FFF

## <span id="Latch_Read_.28.24E000-.24EFFF.2C_Read.29"></span>Latch Read (\$E000-\$EFFF, Read) {#Latch_Read_.28.24E000-.24EFFF.2C_Read.29}

    Mask: $F000

    D~7654 3210
      ---------
      .... LLLL
      |||| ++++- Latched value previously written to D4-D7 of $8000-$FFFF
      ++++------ value from ROM

# Notes {#Notes}

- The board is subject to [bus conflicts](../Bus_conflict/) when writing or reading. The software reads the latch ten times, assuming that after enough tries the driver will overpower the ROM.

# See also {#See_also}

- [forum thread with pictures of PCB](https://forums.nesdev.org/viewtopic.php?t=17817)
