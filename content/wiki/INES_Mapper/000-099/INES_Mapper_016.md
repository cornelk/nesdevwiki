---
title: "INES Mapper 016"
url: "/wiki/INES_Mapper_016/"
description: "iNES Mapper 016 is used for some of the Bandai FCG boards, namely, boards with the FCG-1 ASIC that supports no EEPROM, and the LZ93D50 ASIC with no or 256 bytes of EEPROM."
wikiPage: "INES_Mapper_016"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

iNES Mapper 016 is used for some of the [Bandai FCG boards](../Bandai_FCG_board/), namely, boards with the [FCG-1 ASIC](../Bandai_FCG_pinout/) that supports no EEPROM, and the [LZ93D50 ASIC](../Bandai_LZ93D50_pinout/) with no or 256 bytes of EEPROM.

<table class="wikitable">
<thead>
<tr>
<th colspan="3"><a href="Submapper_table/">INES Mapper 016 submapper table</a></th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Meaning</th>
<th>Note</th>
</tr>
&#10;<tr>
<td>0</td>
<td>Unspecified</td>
<td>Emulate both FCG-1/2 and LZ93D50 chips in their respective CPU address ranges.</td>
</tr>
<tr>
<td><em>1</em></td>
<td><em>LZ93D50 with 128 byte serial EEPROM (24C01)</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_159/">INES Mapper 159</a> instead.</em></td>
</tr>
<tr>
<td><em>2</em></td>
<td><em>Datach Joint ROM System</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_157/">INES Mapper 157</a> instead.</em></td>
</tr>
<tr>
<td><em>3</em></td>
<td><em>8 KiB of WRAM instead of serial EEPROM</em></td>
<td><em>Deprecated, use <a href="../INES_Mapper_153/">INES Mapper 153</a> instead.</em></td>
</tr>
<tr>
<td>4</td>
<td>FCG-1/2</td>
<td>Responds only in the CPU $6000-$7FFF address range; IRQ counter is not latched.</td>
</tr>
<tr>
<td>5</td>
<td>LZ93D50 with no or 256-byte serial EEPROM (24C02)</td>
<td>Responds only in the CPU $8000-$FFFF address range; IRQ counter is latched.</td>
</tr>
</tbody>
</table>

# Game List {#Game_List}

| Name | Chip | EEPROM | NES 2.0 Submapper | [NES 2.0 Byte 10](../NES_2.0/#Byte_10_.28RAM_size.29) |
|----|----|----|----|----|
| *Akuma-kun: Makai no Wana* | FCG-2 | \- | 4 | \$00 |
| *Crayon Shin-chan: Ora to Poi Poi* | LZ93D50 | \- | 5 | \$00 |
| *Dragon Ball: Daimaou Fukkatsu* | FCG-1 | \- | 4 | \$00 |
| *Dragon Ball 3: Gokuu Den* | FCG-2 | \- | 4 | \$00 |
| *Dragon Ball Z II: Gekishin Freezer!!* | LZ93D50 | 24C02 | 5 | \$20 |
| *Dragon Ball Z III: Ressen Jinzou Ningen* | LZ93D50 | 24C02 | 5 | \$20 |
| *Dragon Ball Z Gaiden: Saiya-jin Zetsumetsu Keikaku* | LZ93D50 | 24C02 | 5 | \$20 |
| *Famicom Jump: Hero Retsuden* | FCG-2 | \- | 4 | \$00 |
| *Meimon! Dai-3 Yakyuu-bu* | FCG-1 | \- | 4 | \$00 |
| *Nishimura Kyoutarou Mystery: Blue Train Satsujin Jiken* | FCG-1 | \- | 4 | \$00 |
| *Rokudenashi Blues* | LZ93D50 | 24C02 | 5 | \$20 |
| *Sakigake!! Otoko Juku: Shippu 1-gou Sei* | FCG-1 | \- | 4 | \$00 |
| *SD Gundam Gaiden - Knight Gundam Monogatari 2: Hikari no Kishi* | LZ93D50 | 24C02 | 5 | \$20 |
| *SD Gundam Gaiden - Knight Gundam Monogatari 3: Densetsu no Kishidan* | LZ93D50 | 24C02 | 5 | \$20 |

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR ROM bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR ROM bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR ROM bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR ROM bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR ROM bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR ROM bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR ROM bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR ROM bank

# Registers {#Registers}

## <span id="Read_Serial_EEPROM_.28.246000-.247FFF_read.2C_Submapper_5_only.29"></span>Read Serial EEPROM (\$6000-\$7FFF read, **Submapper 5** only) {#Read_Serial_EEPROM_.28.246000-.247FFF_read.2C_Submapper_5_only.29}

    Mask: $E000

    7  bit  0
    ---- ----
    xxxE xxxx
    |||| ||||
    +++|-++++- Open bus
       +------ Data out from I²C EEPROM

## <span id="CHR-ROM_Bank_Select_.28.246000-.246007_write.2C_Submapper_4.3B_.248000-.248007_write.2C_Submapper_5.29"></span>CHR-ROM Bank Select (\$6000-\$6007 write, **Submapper 4**; \$8000-\$8007 write, **Submapper 5**) {#CHR-ROM_Bank_Select_.28.246000-.246007_write.2C_Submapper_4.3B_.248000-.248007_write.2C_Submapper_5.29}

    Mask: $E00F (Submapper 4), $800F (Submapper 5)

    7  bit  0
    ---- ----
    CCCC CCCC
    |||| ||||
    ++++-++++-- 1 KiB CHR-ROM bank number

- \$xxx0: Select 1 KiB CHR-ROM bank at PPU \$0000-\$03FF
- \$xxx1: Select 1 KiB CHR-ROM bank at PPU \$0400-\$07FF
- \$xxx2: Select 1 KiB CHR-ROM bank at PPU \$0800-\$0BFF
- \$xxx3: Select 1 KiB CHR-ROM bank at PPU \$0C00-\$0FFF
- \$xxx4: Select 1 KiB CHR-ROM bank at PPU \$1000-\$13FF
- \$xxx5: Select 1 KiB CHR-ROM bank at PPU \$1400-\$17FF
- \$xxx6: Select 1 KiB CHR-ROM bank at PPU \$1800-\$1BFF
- \$xxx7: Select 1 KiB CHR-ROM bank at PPU \$1C00-\$1FFF

## <span id="PRG-ROM_Bank_Select_.28.246008_write.2C_Submapper_4.3B_.248008_write.2C_Submapper_5.29"></span>PRG-ROM Bank Select (\$6008 write, **Submapper 4**; \$8008 write, **Submapper 5**) {#PRG-ROM_Bank_Select_.28.246008_write.2C_Submapper_4.3B_.248008_write.2C_Submapper_5.29}

    Mask: $E00F (Submapper 4), $800F (Submapper 5)

    7  bit  0
    ---- ----
    .... PPPP
         ||||
         ++++-- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF   

## <span id="Nametable_Mirroring_Type_Select_.28.246009_write.2C_Submapper_4.3B_.248009_write.2C_Submapper_5.29"></span>Nametable Mirroring Type Select (\$6009 write, **Submapper 4**; \$8009 write, **Submapper 5**) {#Nametable_Mirroring_Type_Select_.28.246009_write.2C_Submapper_4.3B_.248009_write.2C_Submapper_5.29}

    Mask: $E00F (Submapper 4), $800F (Submapper 5)

    7  bit  0
    ---- ----
    .... ..MM
           ||
           ++-- Select nametable mirroring type
                 0: Vertical
                 1: Horizontal
                 2: One-screen, page 0
                 3: One-screen, page 1

## <span id="IRQ_Control_.28.24600A_write.2C_Submapper_4.3B_.24800A_write.2C_Submapper_5.29"></span>IRQ Control (\$600A write, **Submapper 4**; \$800A write, **Submapper 5**) {#IRQ_Control_.28.24600A_write.2C_Submapper_4.3B_.24800A_write.2C_Submapper_5.29}

    Mask: $E00F (Submapper 4), $800F (Submapper 5)

    7  bit  0
    ---- ----
    .... ...C
            |
            +-- IRQ counter control
                 0: Counting disabled
                 1: Counting enabled

- Writing to this register acknowledges a pending IRQ.
- On the LZ93D50 (**Submapper 5**), writing to this register also copies the latch to the actual counter.
- If a write to this register enables counting while the counter is holding a value of zero, an IRQ is generated immediately.

## <span id="IRQ_Latch.2FCounter_.28.24600B-.24600C_write.2C_Submapper_4.3B_.24800B-.24800C_write.2C_Submapper_5.29"></span>IRQ Latch/Counter (\$600B-\$600C write, **Submapper 4**; \$800B-\$800C write, **Submapper 5**) {#IRQ_Latch.2FCounter_.28.24600B-.24600C_write.2C_Submapper_4.3B_.24800B-.24800C_write.2C_Submapper_5.29}

    Mask: $E00F (Submapper 4), $800F (Submapper 5)

       $C         $B
    7  bit  0  7  bit  0
    ---- ----  ---- ----
    CCCC CCCC  CCCC CCCC
    |||| ||||  |||| ||||
    ++++-++++--++++-++++-- Counter value (little-endian)

- If counting is enabled, the counter decreases on every M2 cycle. When it holds a value of zero, an IRQ is generated.
- On the FCG-1/2 (**Submapper 4**), writing to these two registers directly modifies the counter itself; all such games therefore disable counting before changing the counter value.
- On the LZ93D50 (**Submapper 5**), these registers instead modify a latch that will only be copied to the actual counter when register \$800A is written to.

## <span id="EEPROM_Control_.28.24800D_write.2C_Submapper_5_only.29"></span>EEPROM Control (\$800D write, **Submapper 5** only) {#EEPROM_Control_.28.24800D_write.2C_Submapper_5_only.29}

    Mask: $800F

    7  bit  0
    ---- ----
    RDC. ....
    |||
    ||+-------- I²C SCL
    |+--------- I²C SDA
    +---------- Direction bit (1=Enable Read)

- This register only exists on the LZ93D50 (Submapper 5), and only has an effect if a 24C02 EEPROM is present.
- Please refer to generic I²C tutorials and the 24C02 datasheet on how to operate or emulate this register correctly.

# Similar Mappers {#Similar_Mappers}

- [INES Mapper 153](../INES_Mapper_153/) replaces the serial EEPROM with 8 KiB of battery-backed WRAM, and uses CHR-RAM instead of CHR-ROM.
- [INES Mapper 157](../INES_Mapper_157/) adds a barcode reader, support for a second EEPROM, and uses CHR-RAM instead of CHR-ROM.
- [INES Mapper 159](../INES_Mapper_159/) replaces the 256 byte with an 128 byte serial EEPROM.

# See Also {#See_Also}

A discrete logic clone of the FCG-2 has been found and [reverse-engineered by Krzysiobal](https://forums.nesdev.org/viewtopic.php?p=243449#p243449)
