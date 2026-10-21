---
title: "INES Mapper 190"
url: "/wiki/INES_Mapper_190/"
description: "iNES Mapper 190 is assigned to the unlicensed game Magic Kid Googoo released by Zemina."
wikiPage: "INES_Mapper_190"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

***iNES Mapper 190*** is assigned to the unlicensed game *Magic Kid Googoo* released by Zemina.

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB RAM
- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the first bank
- PPU \$0000-\$07FF: 2 KiB switchable CHR ROM bank
- PPU \$0800-\$0FFF: 2 KiB switchable CHR ROM bank
- PPU \$1000-\$17FF: 2 KiB switchable CHR ROM bank
- PPU \$1800-\$1FFF: 2 KiB switchable CHR ROM bank

## Registers {#Registers}

### <span id="CHR-ROM_Bank_Select_.28.24A000-.24BFFF.29"></span>CHR-ROM Bank Select (\$A000-\$BFFF) {#CHR-ROM_Bank_Select_.28.24A000-.24BFFF.29}

    A~[101. .... .... ..bb] D~[CCCC CCCC]
                        ||     |||| ||||
                        ||     ++++-++++--- Select 2 KiB CHR bank ...
                        ++----------------- for PPU address starting at bb*$800

### <span id="PRG-ROM_Bank_Select_.28.248000-.249FFF.2C_.24C000-.24DFFF.29"></span>PRG-ROM Bank Select (\$8000-\$9FFF, \$C000-\$DFFF) {#PRG-ROM_Bank_Select_.28.248000-.249FFF.2C_.24C000-.24DFFF.29}

    A~[1P0. .... .... ....] D~[.... .PPP]
        |                            |||
        +----------------------------+++--- Select 16 KiB PRG bank at $8000-$BFFF
