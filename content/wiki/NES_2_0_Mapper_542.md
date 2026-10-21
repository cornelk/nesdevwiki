---
title: "NES 2.0 Mapper 542"
url: "/wiki/NES_2.0_Mapper_542/"
description: "NES 2.0 Mapper 542 denotes the unmarked PCB used in 毛泽东诞辰: 一百周年纪念 1893-1993 (Máozédōng Dànchén꞉ Yībǎi Zhōunián Jìniàn 1893-1993, Mao Zedong's Birthday: 100th…"
wikiPage: "NES_2_0_Mapper_542"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

NES 2.0 Mapper 542 denotes the unmarked PCB used in *毛泽东诞辰: 一百周年纪念 1893-1993* (Máozédōng Dànchén꞉ Yībǎi Zhōunián Jìniàn 1893-1993, Mao Zedong's Birthday: 100th Anniversary 1893-1993). The game is a modification of *英雄傳 - World Hero*, which is assigned either to mapper [23.1](../VRC2_and_VRC4/) or to [27](../INES_Mapper_027/).

It uses the same VRC4 clone (A0/A1), increases the PRG-ROM size from the original's 128 KiB to 256 KiB, keeps the 512 KiB of CHR-ROM, additionally maps another fixed part of PRG-ROM to \$6000-\$7FFF, and most interestingly, can map CIRAM into CHR address space.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB PRG-ROM bank, fixed to bank \$0F
- CPU \$8000-\$9FFF (or \$C000-\$DFFF): 8 KiB switchable PRG ROM bank (like VRC4)
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank (like VRC4)
- CPU \$C000-\$DFFF (or \$8000-\$9FFF): 8 KiB PRG ROM bank, fixed to the second-last bank (like VRC4)
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank (like VRC4, \$1F in this case)
- PPU \$0000-\$03FF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank (like VRC4) *or* 1 KiB CIRAM bank \#1
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank (like VRC4)
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank (like VRC4)

# Registers {#Registers}

## <span id="Enable_CIRAM_mapping_to_PPU_.240C00-.240FFF_.28.24D800.2C_write.29"></span>Enable CIRAM mapping to PPU \$0C00-\$0FFF (\$D800, write) {#Enable_CIRAM_mapping_to_PPU_.240C00-.240FFF_.28.24D800.2C_write.29}

Mask: \$F800

Once address \$D800 is written to, PPU \$0C00-\$0FFF is mapped to CIRAM bank 1 instead, ignoring the respective VRC4 register. The game uses this CIRAM as pseudo-CHR-RAM to animate the screen showing the next opponent.

## <span id="Disable_CIRAM_mapping_to_PPU_.240C00-.240FFF_.28.24E800.2C_write.29"></span>Disable CIRAM mapping to PPU \$0C00-\$0FFF (\$E800, write) {#Disable_CIRAM_mapping_to_PPU_.240C00-.240FFF_.28.24E800.2C_write.29}

Mask: \$F800

Once address \$E800 is written to, PPU \$0C00-\$0FFF is no longer mapped to CIRAM bank 1, and the respective VRC4 register applies again.

## <span id="VRC4_registers_.28.248000-.24FFFF.29"></span>VRC4 registers (\$8000-\$FFFF) {#VRC4_registers_.28.248000-.24FFFF.29}

Same as on [VRC4](../VRC2_and_VRC4/).
