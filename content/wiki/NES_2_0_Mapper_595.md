---
title: "NES 2.0 Mapper 595"
url: "/wiki/NES_2.0_Mapper_595/"
description: "NES 2.0 Mapper 595 denotes the NES-4MROM-512 circuit board."
wikiPage: "NES_2_0_Mapper_595"
---

**NES 2.0 Mapper 595** denotes the **NES-4MROM-512** circuit board.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB of unbanked PRG-RAM
- CPU \$8000-\$BFFF: 16 KiB switchable window into 512 KiB of PRG-ROM
- CPU \$C000-\$FFFF: 16 KiB fixed window into the last 16 KiB of 512 KiB of PRG-ROM
- PPU \$0000-\$3FFF: 16 KiB of unbanked RAM (CHR-RAM and four nametables)

# Registers {#Registers}

## <span id="Flash_ROM_Write_.28.248000-.24BFFF.2C_write.29"></span>Flash ROM Write (\$8000-\$BFFF, write) {#Flash_ROM_Write_.28.248000-.24BFFF.2C_write.29}

The Flash ROM write sequence is identical to that of [UNROM-512](../UNROM_512/#Flash_ROM_Save), the chip is likewise an SST39SF040 with 4 KiB sectors. Writing anything here has no effect on the shift register.

## <span id="PRG-ROM_Bank_.28.24C000-.24FFFF.2C_write.29"></span>PRG-ROM Bank (\$C000-\$FFFF, write) {#PRG-ROM_Bank_.28.24C000-.24FFFF.2C_write.29}

    D~[.... ...B] (Address AND $C000) = $C000
               +-- Serial shift register data

Somewhat similarly to an [MMC1](../MMC1/), writing to this register shifts the written bit 0 into the highest bit of a five-bit shift register. Unlike MMC1, there is no reset bit, and the register is continuously updated on each write, so that after five writes, the register is guaranteed to hold a valid five-bit number even if a previous write sequence was interrupted. This implies that undesirable banks will be temporarily switched in between the first and fifth writes, meaning that the code performing the writes must execute either only from the fixed bank, or must be replicated in every possible switchable bank at the same address. There is only one register, which has this format:

      [PPPPP]
       +++++- PRG A18..A14 for CPU $8000-$BFFF
