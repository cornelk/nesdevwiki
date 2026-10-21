---
title: "J.Y. Company ASIC"
url: "/wiki/J.Y._Company_ASIC/"
description: "晶太 (Jīngtài, also known as J.Y. Company)'s proprietary ASIC is used for their later single-game cartridges as well as for most of their multicarts. Due to its latter purpose, it allows flexibly…"
wikiPage: "J_Y__Company_ASIC"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers using J.Y. Company ASIC"
    url: "/category/Mappers_using_J.Y._Company_ASIC/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with ROM nametables"
    url: "/category/Mappers_with_ROM_nametables/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
---

<div class="infobox"><div style="text-align: center"><b>J.Y Company ASIC</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>J.Y Company, Hummer Team </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=209">0 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>EL861226C, etc </td></tr><tr><td><b>Pinout</b></td><td>[[<a href="https://forums.nesdev.org/viewtopic.php?t=19017">J.Y Company ASIC pinout</a>]] </td></tr><tr><td><b>PRG ROM capacity</b></td><td>1024K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K, 16K, or 32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>1024K </td></tr><tr><td><b>CHR window</b></td><td>1K, 2K, 4K, or 8K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Arbitrary </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">209</a>, <a href="./">035</a>, <a href="./">090</a>, <a href="./">211</a>, <a href="../NES_2.0_Mapper_281/">281</a>, <a href="../NES_2.0_Mapper_282/">282</a>, <a href="../NES_2.0_Mapper_295/">295</a>, <a href="../NES_2.0_Mapper_358/">358</a></td></tr></tbody></table></div>

晶太 (Jīngtài, also known as J.Y. Company)'s proprietary ASIC is used for their later single-game cartridges as well as for most of their multicarts. Due to its latter purpose, it allows flexibly switching between several PRG and CHR banking modes, supports **ROM nametables**, and permits specifying the CIRAM bank for **each nametable separately** (Extended Mirroring), automatic [MMC4](../MMC4/)-like CHR bankswitching, and a very flexible IRQ counter.

It is part of several iNES/NES 2.0 mappers which differ in their PCB configuration:

- [INES Mapper 209](./) is the **standard implementation**, used for most single-game cartridges. The outer PRG bank size is 512 KiB.
- [INES Mapper 090](./) inhibits, via jumper, ROM nametables and Extended Mirroring that would otherwise be enabled via register.
- [INES Mapper 211](./) supposedly has ROM nametables and Extended Mirroring always enabled regardless of register content. No such PCB exists in real hardware; the mapper was defined when the function of register \$D001 bit 3 was not known, and is a mere duplicate of INES Mapper 209 with correct emulation.
- [INES Mapper 035](./) is like Mapper 209 but explicitly specifies 8 KiB of WRAM. In the presence of a valid NES 2.0 header, which has its own PRG-RAM field, INES Mapper 035 is an exact duplicate of INES Mappers 209.
- [NES 2.0 Mapper 281](../NES_2.0_Mapper_281/) and [NES 2.0 Mapper 282](../NES_2.0_Mapper_282/) specify two PCBs with an outer PRG bank size of 256 KiB but different bit assignments in the Outer Bank register.
- [NES 2.0 Mapper 295](../NES_2.0_Mapper_295/) specifies a PCB with an outer PRG bank size of 128 KiB.
- [NES 2.0 Mapper 358](../NES_2.0_Mapper_358/) specifies a PCB with an outer PRG bank size of 512 KiB selected in part by the CHR banking mode.
- In addition, these mappers also use the J.Y. Company ASIC:
  - [NES 2.0 Mapper 386](../NES_2.0_Mapper_386/)
  - [NES 2.0 Mapper 387](../NES_2.0_Mapper_387/)
  - [NES 2.0 Mapper 388](../NES_2.0_Mapper_388/)
  - [NES 2.0 Mapper 397](../NES_2.0_Mapper_397/)

# Banks {#Banks}

The J.Y Company ASIC provides several banking modes (along with other useful modes) similiar to [MMC5](../MMC5/).

## PRG mode 0 {#PRG_mode_0}

- CPU \$8000-\$FFFF: 32 KB switchable PRG ROM bank

## PRG mode 1 {#PRG_mode_1}

- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KB switchable PRG ROM bank

## <span id="PRG_mode_2.2C_3"></span>PRG mode 2, 3 {#PRG_mode_2.2C_3}

- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB switchable PRG ROM bank

## CHR mode 0 {#CHR_mode_0}

- PPU \$0000-\$1FFF: 8 KB switchable CHR bank

## CHR mode 1 {#CHR_mode_1}

- PPU \$0000-\$0FFF: 4 KB switchable CHR bank
- PPU \$1000-\$1FFF: 4 KB switchable CHR bank

## CHR mode 2 {#CHR_mode_2}

- PPU \$0000-\$07FF: 2 KB switchable CHR bank
- PPU \$0800-\$0FFF: 2 KB switchable CHR bank
- PPU \$1000-\$17FF: 2 KB switchable CHR bank
- PPU \$1800-\$1FFF: 2 KB switchable CHR bank

## CHR mode 3 {#CHR_mode_3}

- PPU \$0000-\$03FF: 1 KB switchable CHR bank
- PPU \$0400-\$07FF: 1 KB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KB switchable CHR bank
- PPU \$1000-\$13FF: 1 KB switchable CHR bank
- PPU \$1400-\$17FF: 1 KB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KB switchable CHR bank

# Registers {#Registers}

## Miscellaneous {#Miscellaneous}

### <span id="Board_Solder_Pad_Register_.28.245000.2C_.245400.2C_.245C00.2C_read.29"></span>Board Solder Pad Register (\$5000, \$5400, \$5C00, read) {#Board_Solder_Pad_Register_.28.245000.2C_.245400.2C_.245C00.2C_read.29}

Mask: \$FFFF

    7  bit  0
    ---- ----
    JJ.. ....
    ||
    ++-------- Jumper/solder pad configuration

Depending on the game, the jumper setting selects different title screens or game counts on multicarts. Apparently, the address of this register is fully decoded, as the SC-128 version of *Final Fight 3*, which otherwise happily randomizes the masked-off address bits, keeps A0-A11 constantly at zero to read from this register.

### <span id="Unsigned_8x8_to_16_Hardware_Multiplier_.28.245800-.245801.2C_read.2Fwrite.29"></span>Unsigned 8x8 to 16 Hardware Multiplier (\$5800-\$5801, read/write) {#Unsigned_8x8_to_16_Hardware_Multiplier_.28.245800-.245801.2C_read.2Fwrite.29}

Mask: \$F803

- \$5800 write: 8-bit Unsigned Multiplicand
- \$5801 write: 8-bit Unsigned Multiplier (and begin multiplying)
- \$5800 read: Unsigned 16-bit Product (low byte)
- \$5801 read: Unsigned 16-bit Product (high byte)

After writing the two operands to \$5800 and \$5801, the result can be read 8 M2 cycles later; if the registers are read from earlier than that, intermediate results are obtained.

### <span id="Accumulator_and_Test_Register_.28.245802-.245803.2C_read.2Fwrite.29"></span>Accumulator and Test Register (\$5802-\$5803, read/write) {#Accumulator_and_Test_Register_.28.245802-.245803.2C_read.2Fwrite.29}

Mask: \$F803

- \$5802 write: Increase accumulator by written value
- \$5803 write: Reset accumulator to 0 and set test register to written value
- \$5802 read: Read accumulator value
- \$5803 read: Read test register

## Banking Registers {#Banking_Registers}

### <span id="PRG-ROM_Bank_Select_.28.248000-.248003.2C_write.29"></span>PRG-ROM Bank Select (\$8000-\$8003, write) {#PRG-ROM_Bank_Select_.28.248000-.248003.2C_write.29}

Mask: \$F803. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing range \$8000-87FF).

In 8 KiB PRG banking mode:

- \$8000: Select 8 KiB PRG-ROM bank at CPU \$8000-\$9FFF
- \$8001: Select 8 KiB PRG-ROM bank at CPU \$A000-\$BFFF
- \$8002: Select 8 KiB PRG-ROM bank at CPU \$C000-\$DFFF
- \$8003: Select 8 KiB PRG-ROM bank at CPU \$E000-\$FFFF if switchable last bank is selected via \$D000 bit 2

In 16 KiB PRG banking mode:

- \$8001: Select 16 KiB PRG-ROM bank at CPU \$8000-\$BFFF
- \$8003: Select 16 KiB PRG-ROM bank at CPU \$C000-\$FFFF if switchable last bank is selected via \$D000 bit 2

In 32 KiB PRG banking mode:

- \$8003: Select 32 KiB PRG-ROM bank at CPU \$8000-\$FFFF if switchable last bank is selected via \$D000 bit 2

If \$D000 bit 2 is clear, the 8/16/32 KiB bank at \$E000/\$C000/\$8000-\$FFFF is hard-wired to the last bank within the outer bank. Furthermore, if register \$D000 bit 7 is set, \$8003 also selects an 8 KiB PRG-ROM bank at CPU \$6000-\$7FFF, which is

- \$8003 in 8 KiB banking mode;
- \$8003 SHL 1 OR 1 in 16 KiB banking mode;
- \$8003 SHL 2 OR 3 in 32 KiB banking mode.

The PRG-ROM bank selected via \$8000-\$8003 is always masked to the outer bank size (512/256/128 KiB, depending on PCB/mapper number). Accessing more than that requires manipulating the Outer Bank register at \$D003.

### <span id="CHR_Bank_Select_.28.249000-.249007.2C_.24A000-.24A007.2C_write.29"></span>CHR Bank Select (\$9000-\$9007, \$A000-\$A007, write) {#CHR_Bank_Select_.28.249000-.249007.2C_.24A000-.24A007.2C_write.29}

Mask: \$F807. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$9000-97FF and \$A000-A7FF).

In 8 KiB CHR banking mode:

- \$9000/\$A000: Select 8 KiB CHR bank at PPU \$0000-\$1FFF

In 4 KiB CHR banking mode with MMC4 mode off:

- \$9000/\$A000: Select 4 KiB CHR bank at PPU \$0000-\$0FFF
- \$9004/\$A004: Select 4 KiB CHR bank at PPU \$1000-\$1FFF

In 4 KiB CHR banking mode with MMC4 mode on:

- \$9000/\$A000 (LSB/MSB): Select 4 KiB CHR bank at PPU \$0000-\$0FFF when latch =0
- \$9002/\$A002 (LSB/MSB): Select 4 KiB CHR bank at PPU \$0000-\$0FFF when latch =1
- \$9004/\$A004 (LSB/MSB): Select 4 KiB CHR bank at PPU \$1000-\$1FFF when latch =0
- \$9006/\$A006 (LSB/MSB): Select 4 KiB CHR bank at PPU \$1000-\$1FFF when latch =1

In 2 KiB CHR banking mode:

- \$9000/\$A000 (LSB/MSB): Select 2 KiB CHR bank at PPU \$0000-\$07FF
- \$9002/\$A002 (LSB/MSB): Select 2 KiB CHR bank at PPU \$0800-\$0FFF
- \$9004/\$A004 (LSB/MSB): Select 2 KiB CHR bank at PPU \$1000-\$17FF
- \$9006/\$A006 (LSB/MSB): Select 2 KiB CHR bank at PPU \$1800-\$1FFF

In 1 KiB CHR banking mode:

- \$9000/\$A000 (LSB/MSB): Select 1 KiB CHR bank at PPU \$0000-\$03FF
- \$9001/\$A001 (LSB/MSB): Select 1 KiB CHR bank at PPU \$0400-\$07FF
- \$9002/\$A002 (LSB/MSB): Select 1 KiB CHR bank at PPU \$0800-\$0BFF
- \$9003/\$A003 (LSB/MSB): Select 1 KiB CHR bank at PPU \$0C00-\$0FFF
- \$9004/\$A004 (LSB/MSB): Select 1 KiB CHR bank at PPU \$1000-\$13FF
- \$9005/\$A005 (LSB/MSB): Select 1 KiB CHR bank at PPU \$1400-\$17FF
- \$9006/\$A006 (LSB/MSB): Select 1 KiB CHR bank at PPU \$1800-\$1BFF
- \$9007/\$A007 (LSB/MSB): Select 1 KiB CHR bank at PPU \$1C00-\$1FFF

The CHR bank selected via \$9000-\$A007 is always masked to the outer bank size (512/256/128 KiB, depending on PCB/mapper number and \$D003 value). Accessing more than that requires manipulating the Outer Bank register at \$D003.

If [MMC4](../MMC4/)-like bankswitching mode is enabled via \$D003 bit 7, the latch is set to 0 when the PPU reads from \$0FD8-\$0FDF/\$1FD8-\$1FDF and to 1 when the PPU reads from \$0FE8-\$0FEF/\$1FE8-\$1FEF.

### <span id="ROM_Nametable_.2F_CIRAM_Bank_Select_.28.24B000-.24B003.2C_.24B004-.24B007.2C_write.29"></span>ROM Nametable / CIRAM Bank Select (\$B000-\$B003, \$B004-\$B007, write) {#ROM_Nametable_.2F_CIRAM_Bank_Select_.28.24B000-.24B003.2C_.24B004-.24B007.2C_write.29}

Mask: \$F807. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$B000-B7FF).

These registers have no function if neither ROM nametables (\$D000 bit 5) nor Extended Mirroring (\$D001 bit 3) are enabled, or their function is supressed via a jumper on some PCB variants that are denoted by [iNES Mapper 090](./).

If Extended Mirroring is enabled (\$D001 bit 3 set):

- \$B000 bit 0: Select 1 KiB CIRAM bank at PPU \$2000-\$23FF
- \$B001 bit 0: Select 1 KiB CIRAM bank at PPU \$2400-\$27FF
- \$B002 bit 0: Select 1 KiB CIRAM bank at PPU \$2800-\$2BFF
- \$B003 bit 0: Select 1 KiB CIRAM bank at PPU \$2C00-\$2FFF

If ROM nametables are enabled (\$D000 bit 5 set) for all nametables (\$D000 bit 6 set):

- \$B000/\$B004 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2000-\$23FF
- \$B001/\$B005 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2400-\$27FF
- \$B002/\$B006 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2800-\$2BFF
- \$B003/\$B007 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2C00-\$2FFF

If ROM nametables are enabled (\$D000 bit 5 set) for selected nametables (\$D000 bit 6 clear):

- \$B000 bit 0: Select 1 KiB CIRAM bank at PPU \$2000-\$23FF if \$B000 bit 7 matches \$D002 bit 7
- \$B001 bit 0: Select 1 KiB CIRAM bank at PPU \$2400-\$27FF if \$B001 bit 7 matches \$D002 bit 7
- \$B002 bit 0: Select 1 KiB CIRAM bank at PPU \$2800-\$2BFF if \$B002 bit 7 matches \$D002 bit 7
- \$B003 bit 0: Select 1 KiB CIRAM bank at PPU \$2C00-\$2FFF if \$B003 bit 7 matches \$D002 bit 7
- \$B000/\$B004 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2000-\$23FF if \$B000 bit 7 differs from \$D002 bit 7
- \$B001/\$B005 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2400-\$27FF if \$B001 bit 7 differs from \$D002 bit 7
- \$B002/\$B006 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2800-\$2BFF if \$B002 bit 7 differs from \$D002 bit 7
- \$B003/\$B007 (LSB/MSB): Select 1 KiB CHR bank at PPU \$2C00-\$2FFF if \$B003 bit 7 differs from \$D002 bit 7

This means that enabling ROM nametables via \$D005 bit 5 and all \$B00x registers having bit 7 match \$D002 bit 7 effectively duplicates the functionality of Extended Mirroring (\$D001 bit 3).

## IRQ Registers and Operation {#IRQ_Registers_and_Operation}

1.  When counting is enabled and the counter is clocked from the source selected by \$C001 bits 0-1, the prescaler is increased/decreased (depending on \$C001 bits 6-7).
2.  If the AND-masked prescaler (according to \$C001 bit 2) wraps, meaning it is zero after increasing or matches the prescaler mask after decreasing, the IRQ counter is increased/decreased (likewise depending on \$C001 bits 6-7).
3.  If the IRQ counter wraps, meaning it is zero after increasing or \$FF after decreasing, an IRQ is generated. Disabling the IRQ via registers \$C000 or \$C002 acknowledges it.

Mask: \$F007. Writes are accepted even if address bit 11 is set.

- \$C000: IRQ Disable/Enable
- \$C001: IRQ Mode/Flags
- \$C002: IRQ Disable
- \$C003: IRQ Enable
- \$C004: Set prescaler value (\*)
- \$C005: Set counter value (\*)
- \$C006: Set XOR value
- \$C007: Configure unknown mode

(\*) The written value is XORed with the content of register \$C006 before being stored.

"Disabling" acknowledges a pending IRQ, inhibits counting, and resets the prescaler to zero. "Enabling" means that counting is no longer inhibited, and IRQs will be generated once prescaler and counter wrap.

### <span id="IRQ_Disable_.2F_Enable_.28.24C000.2C_write.29"></span>IRQ Disable / Enable (\$C000, write) {#IRQ_Disable_.2F_Enable_.28.24C000.2C_write.29}

    7  bit  0
    ---- ----
    .... ...E
            |
            +- Disable (0)/Enable (1) IRQ generation.

### <span id="IRQ_Mode_Select_.28.24C001.2C_write.29"></span>IRQ Mode Select (\$C001, write) {#IRQ_Mode_Select_.28.24C001.2C_write.29}

    7  bit  0
    ---- ----
    DD.. EPSS
    ||   ||||
    ||   ||++- Select IRQ source
    ||   ||     0: CPU M2 rise
    ||   ||     1: PPU A12 rise (unfiltered, eight per scanline)
    ||   ||     2: PPU reads (170 per scanline)
    ||   ||     3: CPU writes
    ||   |+--- Select prescaler mask
    ||   |      0: $FF
    ||   |      1: $07
    ||   +---- Disable (0) / Enable (1) $C007 mode.
    ++-------- Select counting direction
                0: Counting disabled
                1: Increase
                2: Decrease
                3: Counting disabled

### <span id="IRQ_Unknown_Mode_Configuration_.28.24C007.2C_write.29"></span>IRQ Unknown Mode Configuration (\$C007, write) {#IRQ_Unknown_Mode_Configuration_.28.24C007.2C_write.29}

Mask: \$F007. Writes are accepted even if address bit 11 is set.

This register is used together with register \$C001 bit 3. Its functionality is not known, and no known game uses it.

## Configuration {#Configuration}

### <span id="Mode_Select_.28.24D000.2C_write.29"></span>Mode Select (\$D000, write) {#Mode_Select_.28.24D000.2C_write.29}

Mask: \$F803. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$D000-D7FF).

    7  bit  0
    ---- ----
    6GRC CLPP
    |||| ||||
    |||| ||++- Select PRG-ROM banking mode
    |||| ||        0: 32 KiB banking
    |||| ||        1: 16 KiB banking
    |||| ||        2: 8 KiB banking
    |||| ||        3: 8 KiB banking, but with bank numbers bits 0-6 reversed
    |||| |+--- Select bank $8000-(32 KiB)/$C000-(16 KiB)/$E000-(8KiB)-$FFFF
    |||| |         0: Hard-wired to last bank
    |||| |         1: Switchable via $8003
    |||+-+---- Select CHR banking mode
    |||            0: 8 KiB banking
    |||            1: 4 KiB banking
    |||            2: 2 KiB banking
    |||            3: 1 KiB banking
    ||+------- Select ROM nametable status
    ||             0: ROM nametables disabled
    ||             1: ROM nametables enabled (for all or some nametables, depending on bit 6)
    |+-------- Select ROM nametable selection method if R=1, ignored otherwise
    |              0: ROM nametables selected separately by $B00x bit 7 XOR $D002 bit 7
    |              1: ROM nametables enabled globally for all nametables
    +--------- Select CPU $6000-$7FFF mapping
                   0: Map WRAM, if present, otherwise open bus
                   1: Map 8 KiB PRG-ROM bank selected via $8003
                   (appropriately shifted and right-filled with binary 1 in 16/32 KiB PRG modes)

### <span id="Nametable_Arrangement_.28.24D001.2C_write.29"></span>Nametable Arrangement (\$D001, write) {#Nametable_Arrangement_.28.24D001.2C_write.29}

Mask: \$F803. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$D000-D7FF).

    7  bit  0
    ---- ----
    .... E.MM
         | ||
         | ++- Select nametable arrangement type if E=0 and $D000 bit 5=0
         |         0: Horizontal (A10)
         |         1: Vertical (A11)
         |         2: One-screen, page 0 (GND)
         |         3: One-screen, page 1 (VCC)
         +---- Select Extended Mirroring
                   0: Disabled, use MM bits
                   1: Enabled, ignore MM bits, use $B000-$B003 bit 0

### <span id="PPU_Memory_Configuration_.28.24D002.2C_write.29"></span>PPU Memory Configuration (\$D002, write) {#PPU_Memory_Configuration_.28.24D002.2C_write.29}

Mask: \$F803. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$D000-D7FF).

    7  bit  0
    ---- ----
    RW.. ....
    ||
    |+-------- CHR write-enable
    |              0: disabled, CHR memory is write-protected
    |              1: enabled, CHR memory is write-enabled
    +--------- Define ROM nametable selection. 
               Functional only if ROM nametables are generally enabled, but individually selected ($D000 bit 5 set and $D000 bit 6 clear)
                   0: $B00x bit 7=0 selects CIRAM, $B00x bit 7=1 selects ROM nametable
                   1: $B00x bit 7=1 selects CIRAM, $B00x bit 7=0 selects ROM nametable
                   Note: x=0..3

Bit 6 only matters when using CHR-RAM.

### <span id="Outer_Bank_Select_.28.24D003.2C_write.29"></span>Outer Bank Select (\$D003, write) {#Outer_Bank_Select_.28.24D003.2C_write.29}

Mask: \$F803. This means that writes are ignored if address bit 11 is set (i.e. these registers extend only through addressing ranges \$D000-D7FF).

The exact meaning of this register depends on the PCB, specified by the mapper number.

    7  bit  0
    ---- ----
    A.MM MMMM
    | || ||||
    | ++-++++- PCB/Mapper-specific
    +--------- Select MMC4-like automatic CHR-ROM bankswitching mode
                   0: Disable
                   1: Enable (only meaningful in 4 KiB CHR-ROM banking mode)

#### <span id="iNES_Mappers_035.2C_090.2C_209.2C_211"></span>iNES Mappers [035](./), [090](./), [209](./), [211](./) {#iNES_Mappers_035.2C_090.2C_209.2C_211}

    7  bit  0
    ---- ----
    4.Lc cPPC
    | || ||||
    | || |||+- If L=0: Select 256 KiB outer CHR-ROM bank (CHR A18), ignored if L=1
    | || |++-- Select 512 KiB outer PRG-ROM bank (PRG A19-A20)
    | |+-+---- Select 512 KiB outer CHR-ROM bank (CHR A19-A20)
    | +------- Select outer CHR-ROM bank size
    |              0: Mask $900x/$A00x to 256 KiB, use C
    |              1: Mask $900x/$A00x to 512 KiB, ignore C
    +--------- Select MMC4-like automatic CHR-ROM bankswitching mode
                   0: Disable
                   1: Enable (only meaningful in 4 KiB CHR-ROM banking mode)

# Example Games {#Example_Games}

- Mapper 209 is the standard implementation.
  - *Mighty Morphin' Power Rangers III*, alternative title *Mighty Morphin' Power Rangers: The Movie IV* (uses MMC4-like mode for displaying title screen)
  - *Mike Tyson's Punch-Out!!*
  - *Shin Samurai Spirits 2*

<!-- -->

- Mapper 90 inhibits, via jumper, ROM nametables and Extended Mirroring that would otherwise be enabled via register.
  - *1998 Super 45-in-1* (JY-120 (A))
  - *Aladdin* (Hummer Team's port of Capcom's SNES game), alternative titles *Aladdin II*/*III* and *Popeye II: Travels in Persia*
  - *Contra Spirits*
  - *Final Fight 3*
  - *Mickey Mania 7*, alternative title *China Rabbit Baby*
  - *Mortal Kombat II Special*, also known as *Mortal Kombat 3 Special 56 Pepoles* (sic)
  - *Super Aladdin: The Return of Jaffar*
  - *Super Mario World*
  - *Tekken 2*, alternative title *Battle Arena Toushinden*

<!-- -->

- Mapper 211 supposedly has ROM nametables and Extended Mirroring always enabled regardless of register content. No such PCB exists in real hardware; the mapper was defined when the function of register [\$D001 bit 3](#Mirroring_Select_.28.24D001.29) was not known, and is a mere duplicate of Mapper 209 with correct emulation.
  - *Donkey Kong Country 4*, alternative title *The Jungle Book 2*
  - *Tiny Toon Adventures 6*, alternative title *Porky Pig & Daffy Duck*

<!-- -->

- Mapper 35 is a duplicate of Mapper 209 which explicitly specifies 8 KiB of WRAM. In the presence of a valid NES 2.0 header, which has its own PRG-RAM field, Mapper 35 is a duplicate of Mappers 209.
  - *Warioland II* (hack of *Hoshi no Kirby: Yume no Izumi no Monogatari*

# Notes {#Notes}

- This chip is only known to exist as a glob. The pinout can be found here from krzysiobal: <https://forums.nesdev.org/viewtopic.php?t=19017>
- Some boards do not connect CPU A11, changing the \$F800 address masks to \$F000, and effectively making the hardware multiplier at \$5800 inaccessible.
- See also: <https://forums.nesdev.org/viewtopic.php?t=15333>
