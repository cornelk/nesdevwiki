---
title: "PowerPak Menu"
url: "/wiki/PowerPak_Menu/"
description: "The PowerPak boot ROM and menu system use its own simple mapper. The PowerPak menu is split into a series of modules that reside at $0400-$07FF in the NES's built-in RAM and are loaded from the CF…"
wikiPage: "PowerPak_Menu"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
---

The [PowerPak](../PowerPak/) boot ROM and menu system use its own simple mapper. The PowerPak menu is split into a series of modules that reside at \$0400-\$07FF in the NES's built-in RAM and are loaded from the CF card. After booting, the module "I.MAP" is loaded and run.

See also: [PowerPak development resources](../PowerPak/#PowerPak_development_resources)

## Modules {#Modules}

Modules are switched between by writing the filename to fileEntry (usually a single character, calling ClearFindEntry beforehand if it's not already cleared), then calling CardLoadModule. Source for all modules is located [here](http://nespowerpak.com/powerpaksrc/).

- **A.MAP** (<u>A</u>uto) - checks a "saves" directory for a save with the same name as the ROM being loaded
- **B.MAP** (<u>B</u>attery) - save selection screen
- **D.MAP** (<u>D</u>isk) - [FDS](../Family_Computer_Disk_System/) save selection screen
- **E.MAP** (<u>E</u>rror) - displays "Bad File Header"
- **F.MAP** (<u>F</u>DS) - attempts to load an "FDSBIOS.BIN" from the PowerPak directory into WRAM, then starts an FDS game
- **G.MAP** (<u>G</u>ame Genie) - Game Genie .txt file selection screen
- **I.MAP** (<u>I</u>nit) - title screen
- **L.MAP** (<u>L</u>ast) - saves the game name to LASTNES.TXT in the PowerPak directory, also loads appropriate iNES mapper file based on header
- **M.MAP** (<u>M</u>usic) - loads an NSF and the NSF player, and starts the NSF
- **N.MAP** (<u>N</u>avigate) - ROM selection screen
- **O.MAP** (<u>O</u>ptions) - game options screen, with Game Genie code entering and save loading
- **Q.MAP** (<u>Q</u>uestion) - "save to CF card?" screen
- **S.MAP** (<u>S</u>ave) - writes a save to the CF card

## Banks {#Banks}

- PPU \$0000-\$1FFF: 8 KB switchable CHR RAM bank.
- CPU \$6000-\$7FFF: 8 KB switchable PRG RAM bank, selects from the 512 KB "PRG ROM" space.
- CPU \$8000-\$BFFF: 16 KB switchable PRG ROM bank, selects from boot ROM.
- CPU \$C000-\$FFFF: 16 KB PRG ROM bank, fixed to the last bank of boot ROM.

## Registers {#Registers}

### <span id="FPGADATA_.28.245000.29"></span>FPGADATA (\$5000) {#FPGADATA_.28.245000.29}

The boot ROM writes \$FF to this register. It tells the FPGA it is being reprogrammed.

### <span id="FPGAPROGRAM_.28.245800.29"></span>FPGAPROGRAM (\$5800) {#FPGAPROGRAM_.28.245800.29}

    7  bit  0
    ---- ----
    DDDD DDDD
    |||| ||||
    ++++-++++- Writes a byte of configuration to the FPGA.

### <span id="MAPPERWR_.28.248000.29"></span>MAPPERWR (\$8000) {#MAPPERWR_.28.248000.29}

    7  bit  0
    ---- ----
    .... ..DD
           ||
           ++- Selects the bank of boot ROM that appears at $8000-$BFFF.

Banks 0, 1 and 2 contain the initial FPGA configuration. Bank 3 contains 8KB of font and then 8KB of system routines. As the last 16KB are also accessible via the fixed bank, the \$8000-\$BFFF area isn't useful for menu software.

### <span id="PRGBANK_.28.244200.29"></span>PRGBANK (\$4200) {#PRGBANK_.28.244200.29}

    7  bit  0
    ---- ----
    WCDD DDDD
    |||| ||||
    ||++-++++- Selects the bank of PRG RAM that appears in $6000-$7FFF.
    |+-------- Enable CHR RAM writing if 1.
    +--------- Selects from 32KB of WRAM instead of main PRG RAM.

### <span id="CHRBANK_.28.244201.29"></span>CHRBANK (\$4201) {#CHRBANK_.28.244201.29}

    7  bit  0
    ---- ----
    ..DD DDDD
      || ||||
      ++-++++- Selects the bank of CHR RAM used.

## External Links {#External_Links}

- Collection of information and photos, source code for menu and loader: <http://nespowerpak.com/>
