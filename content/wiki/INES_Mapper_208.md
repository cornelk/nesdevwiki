---
title: "INES Mapper 208"
url: "/wiki/INES_Mapper_208/"
description: "This mapper is used on exactly one known game: 快打傳説 Street Fighter IV. It uses MMC3-compatible registers for CHR ROM bankswitching and IRQ control, and additional registers for PRG ROM…"
wikiPage: "INES_Mapper_208"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

This mapper is used on exactly one known game: [快打傳説 Street Fighter IV](http://bootleggames.wikia.com/wiki/Street_Fighter_IV). It uses [MMC3](../MMC3/)-compatible registers for CHR ROM bankswitching and IRQ control, and additional registers for PRG ROM bankswitching, mirroring control, and protection.

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank
- PPU \$0000-\$07FF (or \$1000-\$17FF): 2 KB switchable CHR bank (same as MMC3)
- PPU \$0800-\$0FFF (or \$1800-\$1FFF): 2 KB switchable CHR bank (same as MMC3)
- PPU \$1000-\$13FF (or \$0000-\$03FF): 1 KB switchable CHR bank (same as MMC3)
- PPU \$1400-\$17FF (or \$0400-\$07FF): 1 KB switchable CHR bank (same as MMC3)
- PPU \$1800-\$1BFF (or \$0800-\$0BFF): 1 KB switchable CHR bank (same as MMC3)
- PPU \$1C00-\$1FFF (or \$0C00-\$0FFF): 1 KB switchable CHR bank (same as MMC3)

## Registers {#Registers}

This description is based on [FCEUX's implementation](https://github.com/asfdfdfd/fceux/blob/master/src/boards/208.cpp), except for the additional information on the mirroring control.

### <span id="PRG_Bank_and_Mirroring_.28.244800-.244FFF.2C_.246800-.246FFF.2C_write-only.29"></span>PRG Bank and Mirroring (\$4800-\$4FFF, \$6800-\$6FFF, write-only) {#PRG_Bank_and_Mirroring_.28.244800-.244FFF.2C_.246800-.246FFF.2C_write-only.29}

    7  bit  0
    ---- ----
    ..MP ...P
      ||    |
      |+----+- Select 32KB page of PRG-ROM at $8000
      +------- Select mirroring
                 0: Vertical
                 1: Horizontal

The initial value of this register is \$11.

### <span id="Protection_Table_Index_.28.245000-.2457FF.2C_write-only.29"></span>Protection Table Index (\$5000-\$57FF, write-only) {#Protection_Table_Index_.28.245000-.2457FF.2C_write-only.29}

Specifies the index into the 256-byte table that will be used for the next write to \$5800-\$5FFF.

### <span id="Protection_Data_Access_.28.245800-.245FFF.2C_read.29"></span>Protection Data Access (\$5800-\$5FFF, read) {#Protection_Data_Access_.28.245800-.245FFF.2C_read.29}

Returns the content of one of the four protection registers determined by the two lower bits of the address being read. The correct result is checked for by executing code at \$5800, which in the case of proper emulation of the protection scheme will result in the execution of a JMP opcode to a desired address.

### <span id="Protection_Data_Access_.28.245800-.245FFF.2C_write.29"></span>Protection Data Access (\$5800-\$5FFF, write) {#Protection_Data_Access_.28.245800-.245FFF.2C_write.29}

Writes to one of the four protection registers determined by the two lower bits of the address being written to. The value that is written to the protection register is the CPU data XORed with the table entry indexed by the last write to \$5000-\$57FF. The table holds the following values:

    static uint8 lut[256] = {
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x49, 0x19, 0x09, 0x59, 0x49, 0x19, 0x09,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x51, 0x41, 0x11, 0x01, 0x51, 0x41, 0x11, 0x01,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x49, 0x19, 0x09, 0x59, 0x49, 0x19, 0x09,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x51, 0x41, 0x11, 0x01, 0x51, 0x41, 0x11, 0x01,
        0x00, 0x10, 0x40, 0x50, 0x00, 0x10, 0x40, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x08, 0x18, 0x48, 0x58, 0x08, 0x18, 0x48, 0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x00, 0x10, 0x40, 0x50, 0x00, 0x10, 0x40, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x08, 0x18, 0x48, 0x58, 0x08, 0x18, 0x48, 0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x58, 0x48, 0x18, 0x08, 0x58, 0x48, 0x18, 0x08,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x50, 0x40, 0x10, 0x00, 0x50, 0x40, 0x10, 0x00,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x58, 0x48, 0x18, 0x08, 0x58, 0x48, 0x18, 0x08,
        0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x59, 0x50, 0x40, 0x10, 0x00, 0x50, 0x40, 0x10, 0x00,
        0x01, 0x11, 0x41, 0x51, 0x01, 0x11, 0x41, 0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x09, 0x19, 0x49, 0x59, 0x09, 0x19, 0x49, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x01, 0x11, 0x41, 0x51, 0x01, 0x11, 0x41, 0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
        0x09, 0x19, 0x49, 0x59, 0x09, 0x19, 0x49, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
    };

### <span id="IRQ_latch_.28.24C000-.24DFFE.2C_even.29"></span>IRQ latch (\$C000-\$DFFE, even) {#IRQ_latch_.28.24C000-.24DFFE.2C_even.29}

### <span id="IRQ_reload_.28.24C001-.24DFFF.2C_odd.29"></span>IRQ reload (\$C001-\$DFFF, odd) {#IRQ_reload_.28.24C001-.24DFFF.2C_odd.29}

### <span id="IRQ_disable_.28.24E000-.24FFFE.2C_even.29"></span>IRQ disable (\$E000-\$FFFE, even) {#IRQ_disable_.28.24E000-.24FFFE.2C_even.29}

### <span id="IRQ_enable_.28.24E001-.24FFFF.2C_odd.29"></span>IRQ enable (\$E001-\$FFFF, odd) {#IRQ_enable_.28.24E001-.24FFFF.2C_odd.29}

These registers are identical to the respective [MMC3](../MMC3/) registers.

## Submapper 1 {#Submapper_1}

A version of the game that removes the "Gouder Co. Ltd." copyright notice from the title screen uses a different MMC3-like mapper that is incompatible to the description above. That version of the game is inexplicably labelled "Mortal Kombat (JJ-01) (Ch) \[!\]" in GoodNES v3.23b. [The Nestopia Plus! author has assigned this different mapper to submapper number 1 of mapper 208](http://nestopia.gamemw.com/index.php/34-nestopia-plus-1-4-0-10-r272).

There are no additional registers, even as the ROM image retains some leftovers of code writing to \$5800. Instead of switching PRG ROM banks MMC3-style, there is only one 32k PRG ROM bank at CPU \$8000. Its value is determined by MMC3 bank register 6; the actual bank number is the bank register value SHR 2. CHR ROM bankswitching, mirroring and IRQ control are performed using the normal MMC3 registers.

## Related {#Related}

This is reminiscent of [iNES Mapper 189](../INES_Mapper_189/).
