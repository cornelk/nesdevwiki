---
title: "INES Mapper 188"
url: "/wiki/INES_Mapper_188/"
description: "iNES Mapper 188 describes the board used for Bandai's Karaoke Studio. This game, like Nantettatte!! Baseball additionally supports an external ROM that allows augmenting the original game."
wikiPage: "INES_Mapper_188"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
---

iNES Mapper 188 describes the board used for Bandai's Karaoke Studio. This game, like [Nantettatte!! Baseball](../INES_Mapper_068/#.24F000-.24FFFF:_PRG_bank_at_.248000) additionally supports an external ROM that allows augmenting the original game.

## Overview {#Overview}

- PRG ROM: Up to 256 KiB + Up to 256 KiB expansion cartridge
- PRG ROM bank size: 16 KiB
- PRG RAM: None
- CHR capacity: 8 KiB RAM
- CHR bank size: Not bankswitched
- Nametable [mirroring](../Mirroring/): Controlled by mapper
- Subject to [bus conflicts](../Bus_conflict/): Yes

The [mapper IC](../Bandai_M60001-0115P_pinout/) used here has Mitsubishi's part number for custom orders: M60001.

## Registers {#Registers}

### <span id=".246000-.247FFF:_Microphone_input"></span>\$6000-\$7FFF: Microphone input {#.246000-.247FFF:_Microphone_input}

The microphone for this game is permanently tethered to the cartridge, and contains two buttons labeled "A" (closer to the microphone itself) and "B" (closer to the cord).

Reads from this address provide the status of the microphone and its buttons:

    7  bit  0
    .... .MBA
    ---- ----
    |||| ||||
    |||| |||+- 0: A button is pressed
    |||| ||+-- 0: B button is pressed
    |||| |+--- 1-bit ADC microphone input
    ++++-+---- open bus

### <span id=".248000-.24FFFF:_Banking_control"></span>\$8000-\$FFFF: Banking control {#.248000-.24FFFF:_Banking_control}

    7  bit  0
    .LXR BBBB
    ---- ----
     ||| ||||
     ||| ++++- Select ROM bank mapped from $8000-$BFFF.
     ||+------ 0: Select external ROM, 1: Select internal ROM
     |+------- CIRAM A10 is connected to  0: PPU A10 ("vertical mirroring"), 1: PPU A11 ("horizontal mirroring")
     +-------- 1 bit latch, present but unused

The bank from \$C000-\$FFFF is fixed to the last bank of the internal ROM.

## Warning {#Warning}

The 256 KiB iNES dumps under this mapper are actually the 128KiB internal ROM followed by a 128KiB expansion ROM.

For the 128 KiB iNES dump (solely the internal ROM), several emulators implement logic that selects the last bank of the internal ROM when the game tries to select the first of the external ROM. There may be a fingerprint the game requires on being absent. Properly implementing open bus is more accurate.

## References {#References}

- Enri's schematic: <http://cmpslv3.stars.ne.jp/Famic/Famic.htm> (scroll down to " カラオケスタジオの回路図 ")
- Naruko's wiki: [カラオケスタジオ](http://seesaawiki.jp/famicomcartridge/d/%a5%ab%a5%e9%a5%aa%a5%b1%a5%b9%a5%bf%a5%b8%a5%aa)
- Nestopia and FCEUX source
