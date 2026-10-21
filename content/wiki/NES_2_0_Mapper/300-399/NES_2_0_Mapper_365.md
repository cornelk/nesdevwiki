---
title: "NES 2.0 Mapper 365"
url: "/wiki/NES_2.0_Mapper_365/"
description: "NES 2.0 Mapper 365 is used for the Polish version, and probably other language versions, of the Asder PC-95 educational computer."
wikiPage: "NES_2_0_Mapper_365"
---

NES 2.0 Mapper 365 is used for the Polish version, and probably other language versions, of the Asder PC-95 educational computer.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB fixed PRG-RAM bank
- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG-ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG-ROM bank
- CPU \$E000-\$FFFF: 8 KiB switchable PRG-ROM bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR-RAM bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR-RAM bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR-RAM bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR-RAM bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR-RAM bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR-RAM bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR-RAM bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR-RAM bank

# Registers {#Registers}

## <span id="PRG-ROM_Bank_Select_.28.248000-.249FFF.2C_write.29"></span>PRG-ROM Bank Select (\$8000-\$9FFF, write) {#PRG-ROM_Bank_Select_.28.248000-.249FFF.2C_write.29}

    Mask: $E003

    $8000: Select 8 KiB PRG-ROM bank at CPU $8000-$9FFF
    $8001: Select 8 KiB PRG-ROM bank at CPU $A000-$BFFF
    $8002: Select 8 KiB PRG-ROM bank at CPU $C000-$DFFF
    $8003: Select 8 KiB PRG-ROM bank at CPU $E000-$FFFF. Written value is ORed with $01.

## <span id="CHR-RAM_Bank_Select_.28.24A000-.24BFFF.2C_write.29"></span>CHR-RAM Bank Select (\$A000-\$BFFF, write) {#CHR-RAM_Bank_Select_.28.24A000-.24BFFF.2C_write.29}

    Mask: $E007

    $A000: Select 1 KiB CHR-RAM bank at PPU $0000-$03FF
    $A001: Select 1 KiB CHR-RAM bank at PPU $0400-$07FF
    $A002: Select 1 KiB CHR-RAM bank at PPU $0800-$0BFF
    $A003: Select 1 KiB CHR-RAM bank at PPU $0C00-$0FFF
    $A004: Select 1 KiB CHR-RAM bank at PPU $1000-$13FF
    $A005: Select 1 KiB CHR-RAM bank at PPU $1400-$17FF
    $A006: Select 1 KiB CHR-RAM bank at PPU $1800-$1BFF
    $A007: Select 1 KiB CHR-RAM bank at PPU $1C00-$1FFF

## <span id="Unknown_Registers_.28.24C000-.24DFFF.2F.24E000-.24FFFF.2C_write.29"></span>Unknown Registers (\$C000-\$DFFF/\$E000-\$FFFF, write) {#Unknown_Registers_.28.24C000-.24DFFF.2F.24E000-.24FFFF.2C_write.29}

Mask: \$E001

One of them will select the nametable mirroring, but it is not known which one, and what the other register does.

## <span id="Unknown_Register_.28.244903.2C_read.29"></span>Unknown Register (\$4903, read) {#Unknown_Register_.28.244903.2C_read.29}

    Mask: unknown

    D~7654 3210
      ---------
      .... .S..
            +--- Status bit of unknown meaning

\$4903 seems to indicate the presence of a device of some kind and is only read when the computer's menu, or some of its individual applications, boot.

## <span id="Keyboard_Row_Select_.28.244904.2C_write.29"></span>Keyboard Row Select (\$4904, write) {#Keyboard_Row_Select_.28.244904.2C_write.29}

Mask: unknown

Selects the row of the keyboard from which the button state will be returned when reading from \$4906.

## <span id="Unknown_register_.28.244905.2C_write.29"></span>Unknown register (\$4905, write) {#Unknown_register_.28.244905.2C_write.29}

Mask: unknown

    D~7654 3210
      ---------
      ?... ..??

This register is written to by some of of the applications at the end of their NMI handler.

## <span id="Keyboard_Button_State_.28.244906.2C_read.29"></span>Keyboard Button State (\$4906, read) {#Keyboard_Button_State_.28.244906.2C_read.29}

Mask: unknown

Returns the button states of the keyboard row selected via \$4904. The mapping of keys to bits and rows matches that of the Subor keyboard.

## <span id="Unknown_Registers_.28.244910-.24491F.2C_read.2Fwrite.29"></span>Unknown Registers (\$4910-\$491F, read/write) {#Unknown_Registers_.28.244910-.24491F.2C_read.2Fwrite.29}

## <span id="Unknown_Registers_.28.244111.2C_write.3B_.245000-5003.2F.245080-.245083.2C_read.29"></span>Unknown Registers (\$4111, write; \$5000-5003/\$5080-\$5083, read) {#Unknown_Registers_.28.244111.2C_write.3B_.245000-5003.2F.245080-.245083.2C_read.29}

These registers are accessed when selecting "save" and "load" from the text editor ("Edytor tekstu").

## <span id="Printer_Port_Status_.28.244902.2C_read.29"></span>Printer Port Status (\$4902, read) {#Printer_Port_Status_.28.244902.2C_read.29}

    Mask: unknown

    D~7654 3210
      ---------
      B... ....
      +--------- 0=Ready, 1=Busy

## <span id="Printer_Port_Output_.28.244900.2F.244901.2C_write.29"></span>Printer Port Output (\$4900/\$4901, write) {#Printer_Port_Output_.28.244900.2F.244901.2C_write.29}

Mask: unknown

The text editor writes every byte to be printed both to \$4900 and \$4901.

# Note {#Note}

- The Arabic version of the Asder PC-95 Educational Computer most likely originally used this mapper as well. The publicly-available ROM image has obviously been hacked for [iNES Mapper 241](../INES_Mapper_241/) and crashes at several points.
- The Korean version uses different hardware and adds a speech chip, which are described by [NES 2.0 Mapper 531](../NES_2.0_Mapper_531/).
