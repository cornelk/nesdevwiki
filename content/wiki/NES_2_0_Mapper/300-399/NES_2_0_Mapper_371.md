---
title: "NES 2.0 Mapper 371"
url: "/wiki/NES_2.0_Mapper_371/"
description: "NES 2.0 Mapper 371 denotes the Spanish PEC-586 home computer main cartridge, which differs from both the Chinese and Russian equivalents. It has two PRG-ROM chips, one with 64 KiB, one with 512 KiB,…"
wikiPage: "NES_2_0_Mapper_371"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

[<img src="../../wiki-images/PEC-586_Spanish_PCB_F.jpg" width="300" height="265" />](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)

[PCB front](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_F.jpg)[<img src="../../wiki-images/PEC-586_Spanish_PCB_B.jpg" width="300" height="256" />](../../wiki-images/PEC-586_Spanish_PCB_B.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_B.jpg)

[PCB back (mirrored, pin 31 on the left)](../../wiki-images/PEC-586_Spanish_PCB_B.jpg)

[](../../wiki-images/PEC-586_Spanish_PCB_B.jpg)

NES 2.0 Mapper 371 denotes the Spanish PEC-586 home computer main cartridge, which differs from both the Chinese and Russian equivalents. It has two PRG-ROM chips, one with 64 KiB, one with 512 KiB, for a total of 576 KiB. The first chip contains normal home computer applications such as a BASIC interpreter and a text editor, while the second chip contains a selection of common NROM games. Like all PEC-586 machines and cartridges, it adds a 1bpp video mode by appropriately mapping CHR-RAM, and requires the ~~PEC-586 Keyboard~~ to operate.

# Banks {#Banks}

Common:

- CPU \$6000-\$7FFF: 8 KiB of unbanked PRG-RAM
- PPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM

When the first (64 KiB) PRG-ROM chip is selected

- CPU \$8000-\$BFFF: Switchable 16 KiB PRG-ROM bank (#0-#3)
- CPU \$C000-\$FFFF: Fixed 16 KiB PRG-ROM bank \#3

When the second (512 KiB) PRG-ROM chip is selected:

- CPU \$8000-\$BFFF: Switchable 16 KiB PRG-ROM bank
- CPU \$C000-\$FFFF: Mirror of CPU \$8000-\$BFFF

# Registers {#Registers}

## <span id="Main_Register_.28.245000.2C_write.29"></span>Main Register (\$5000, write) {#Main_Register_.28.245000.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      1CCC PPPP
      |||| ++++- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF (PRG A14..A17)
      |+++------ Select PRG-ROM source
      |           0: First chip (64 KiB)
      |           5: Second chip (512 KiB)
      |           7: Expansion slot
      +--------- Select 1bpp CHR mode
                  0: Disabled
                  1: Enabled

Power-on value: \$00

## <span id="Auxiliary_Register_.28.245100.2C_write.29"></span>Auxiliary Register (\$5100, write) {#Auxiliary_Register_.28.245100.2C_write.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... ..MP
             |+- If first PRG-ROM chip selected: Tape Out data
             |   If second PRG-ROM chip selected: PRG A18 (256 KiB)
             +-- Select nametable mirroring type
                  0: Horizontal
                  1: Vertical

Power-on value: \$03

## <span id="Tape_In_.28.245500.2C_read.29"></span>Tape In (\$5500, read) {#Tape_In_.28.245500.2C_read.29}

    Mask: $FF00

    D~7654 3210
      ---------
      .... .T..
            +--- Tape In data

# Protection {#Protection}

When the second PRG-ROM chip is selected *and* PRG A18=0, the cartridge hardware will forcibly reset on attempts to read from CPU \$D100-\$D1FF. As the first 256 KiB bank of the second PRG-ROM chip only contains data for CHR-RAM as well as some code that is copied into and executed from PRG-RAM, and \$C000-\$FFFF is a mirror of \$8000-\$BFFF in that configuration, the \$D100-\$D1FF range is never read from during normal operation.

# 1bpp CHR mode {#1bpp_CHR_mode}

When register \$5000 bit 7 is set, CHR A3 is substituted with PPU A0 from the last nametable (not attribute table) read, and CHR A12 is substituted with PPU A9 from the last nametable (not attribute table) read. This creates a 1bpp all-points-addressable video mode that allows the 8 KiB CHR-RAM to hold just enough data for one screen.
