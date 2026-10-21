---
title: "INES Mapper 191"
url: "/wiki/INES_Mapper_191/"
description: "iNES Mapper 191 denotes the circuit board used for Xianfeng Cartoon's Chinese translation of Technos' すごろ Quest: Diceの戦士たち, entitled 赌神 (Dǔshén). Having increased the PRG-ROM…"
wikiPage: "INES_Mapper_191"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
---

**iNES Mapper 191** denotes the circuit board used for [Xianfeng Cartoon's](https://bootleggames.fandom.com/wiki/Xianfeng_Cartoon) Chinese translation of Technos' *すごろ Quest: Diceの戦士たち*, entitled *赌神* (Dǔshén). Having increased the PRG-ROM and CHR-ROM size from 128+128 KiB to 256+256 KiB without fully using it for the translation, Xianfeng Cartoon filled up the remaining space with three bonus games, making the cartridge a 4-in-1 multicart, with the main feature game using a dual "CHR-ROM plus 2 KiB of CHR-RAM" configuration similar to [TQROM](../INES_Mapper_119/).

Because the multicart functionality was only discovered recently, and because the common dump trims CHR-ROM to 128 KiB, removing the CHR pattern data of the three bonus games, **Submapper 0** disables the outer bank register and always keeps the CHR-ROM/CHR-RAM switch enabled, while **Submapper 1** enables the outer bank register and allows proper emulation. For [iNES](../INES/)-headered files without submapper information, use submapper 1 when the CHR-ROM size is 256 KiB and submapper 0 when the CHR-ROM size is 128 KiB. Running the full 256+256 KiB ROM file as submapper 0 skips the menu and boots directly into the main feature game.

## <span id="Outer_Bank_Register_.28.2490AA.2C_write.29"></span>Outer Bank Register (\$90AA, write) {#Outer_Bank_Register_.28.2490AA.2C_write.29}

    D~[.... ..GG] A~[1001 .... 1010 1010] (Address AND $F0FF)=$90AA
              ++- PRG A15..A14 when CPU A14=1

    Power-on value: 0

The outer bank register defines the behavior of a PAL/GAL that affects two aspects:

- PRG A15..A14 of the fixed banks, so in the 256 KiB PRG-ROM, the fixed 8 kiB banks are bank#:
  - Value 0: \$18 and \$19
  - Value 1: \$1A and \$1B
  - Value 2: \$1C and \$1D
  - Value 3: \$1E and \$1F
- CHR banking:
  - Values 0,1,2: CHR-ROM A17=1, CHR-RAM disabled
  - Value 3: CHR-ROM A17=0, MMC3's CHR A17 output selects 128 KiB CHR-ROM (0) or 2 KiB CHR-RAM (1)

Because the power-on value is 0, the reset vector for the 4-in-1 menu is at PRG-ROM offset \$33FFC.

## See also {#See_also}

- [MMC3 with CHR ROM and CHR RAM](../MMC3_with_CHR_ROM_and_CHR_RAM/)
