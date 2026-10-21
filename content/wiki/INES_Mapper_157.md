---
title: "INES Mapper 157"
url: "/wiki/INES_Mapper_157/"
description: "iNES Mapper 157 is used for games using the Datach Joint ROM system. This belongs in the Bandai FCG family of mappers. The Datach main unit is a device for insertion into the Famicom's cartridge port…"
wikiPage: "INES_Mapper_157"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

iNES Mapper 157 is used for games using the *Datach Joint ROM system*. This belongs in the [Bandai FCG](../Bandai_FCG_board/) family of mappers. The Datach main unit is a device for insertion into the Famicom's cartridge port that provides a barcode scanner, and that accepts its own custom game cartridges. It contains an **internal** 256-byte serial EEPROM (X24C02) that is shared among all Datach games. Graphics are stored in 8 KiB of unbanked CHR-RAM instead of bankswitched CHR-ROM.

One game, *Battle Rush: Build up Robot Tournament*, has an additional **external** 128-byte serial EEPROM (X24C01) on the game cartridge. As part of that game's functionality is to transfer save data between the Datach main unit's EEPROM and the game cartridge's EEPROM, emulators should save each EEPROM's data to separate files, one for the main unit's EEPROM, and one for the game cartridge's EEPROM.

The NES 2.0 header's PRG-NVRAM field will only denote whether the game cartridge has an additional 128-byte serial EEPROM (byte 10 value 0x10) or not (byte 10 value 0x00). The Datach main unit's serial EEPROM is not included in the NES 2.0 header's PRG-NVRAM field, as that EEPROM is not part of the game cartridge, but instead shared among all games using this mapper.

# Game List {#Game_List}

| Name | External EEPROM | [NES 2.0 Byte 10](../NES_2.0/#Byte_10_.28RAM_size.29) |
|----|----|----|
| *Battle Rush: Build up Robot Tournament* | X24C01 | 0x10 |
| *Crayon Shin-chan: Ora to Poi Poi* | \- | 0x00 |
| *Dragon Ball Z: Gekitou Tenkaichi Budoukai* | \- | 0x00 |
| *J-League Super Top Players* | \- | 0x00 |
| *SD Gundam Wars* | \- | 0x00 |
| *Ultraman Club: Spokon Fight!!* | \- | 0x00 |
| *Yuu Yuu Hakusho - Bakutou Ankoku Bujutsu-kai* | \- | 0x00 |

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="Read_Serial_EEPROM.2FBarcode_.28.246000-.247FFF_read.29"></span>Read Serial EEPROM/Barcode (\$6000-\$7FFF read) {#Read_Serial_EEPROM.2FBarcode_.28.246000-.247FFF_read.29}

    Mask: $E000

    7  bit  0
    ---- ----
    xxxE Bxxx
    |||| ||||
    +++|-|+++- Open bus
       | +---- Serial data out from Barcode scanner
       +------ Data out from I²C EEPROMs

- If there is both an internal and an external EEPROM, the Data out from both EEPROMs is ANDed (open drain circuit).

The barcode scanner just returns a sequence of light and dark periods from a one-dimesional barcode as the user moves it past the sensor. The game only expects [UPCs](https://en.wikipedia.org/wiki/Universal_Product_Code) to be scanned, so an emulator may wish to provide an interface to convert a series of digits into a UPC, similar to [FCEUX](https://github.com/TASVideos/fceux/blob/master/src/boards/bandai.cpp#L303).

## <span id="External_EEPROM_Clock_.28.248000-.248003_write.29"></span>External EEPROM Clock (\$8000-\$8003 write) {#External_EEPROM_Clock_.28.248000-.248003_write.29}

    Mask: $800F

    7  bit  0
    ---- ----
    xxxx Cxxx
         |   
         +----- External EEPROM's I²C SCL

- CHR A13 OUT is provided on the subcartridge header, and is used in *Battle Rush* as an external I²C clock.
- PPU rendering must either be disabled during traffic to the external EEPROM, or the same value must be written to all four registers.
- Because the ASIC's PA12 and PA13 inputs are grounded, only registers \$8000-\$8003 instead of \$8000-\$8007 are relevant.
- No CHR banking is available.

## <span id="PRG-ROM_Bank_Select_.28.248008_write.29"></span>PRG-ROM Bank Select (\$8008 write) {#PRG-ROM_Bank_Select_.28.248008_write.29}

## <span id="Nametable_Mirroring_Type_Select_.28.248009_write.29"></span>Nametable Mirroring Type Select (\$8009 write) {#Nametable_Mirroring_Type_Select_.28.248009_write.29}

## <span id="IRQ_Control_.28.24800A_write.29"></span>IRQ Control (\$800A write) {#IRQ_Control_.28.24800A_write.29}

## <span id="IRQ_Latch_.28.24800B-.24800C_write.29"></span>IRQ Latch (\$800B-\$800C write) {#IRQ_Latch_.28.24800B-.24800C_write.29}

These four registers function the same way as on [INES Mapper 016](../INES_Mapper_016/), submapper 5.

## <span id="EEPROM_Control_.28.24800D_write.29"></span>EEPROM Control (\$800D write) {#EEPROM_Control_.28.24800D_write.29}

    Mask: $800F

    7  bit  0
    ---- ----
    RDC. ....
    |||
    ||+-------- Internal EEPROM's I²C SCL
    |+--------- Both EEPROM's I²C SDA
    +---------- Both EEPROM's Direction bit (1=Enable Read)

- Please refer to generic I²C tutorials and the X24C01/24C02 datasheets on how to operate or emulate this register correctly.
- The device address of the internal X24C02 EEPROM is set to 000b (with the highest "device type" bits being 1010b), whereas the external X24C01 EEPROM responds to *any* I²C start condition. It is for this reason that there are two clock lines to address each EEPROM separately.
- Note that while the 24C02 serializes data bits from MSB to LSB (i.e. first bit 7 of any byte, then 6..0), Bandai persisted in believing that the X24C01 serializes data bits from LSB to MSB instead. The X24C01 is actually still a big-endian device, and multi-byte reads and writes increment the address as MSB-first.
