---
title: "INES Mapper 106"
url: "/wiki/INES_Mapper_106/"
description: "This board, made entirely from discrete chips, is used by Super Mario Bros 3 bootleg. The game was modified (with copyright information cut out) and ported from MMC3."
wikiPage: "INES_Mapper_106"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

This board, made entirely from discrete chips, is used by Super Mario Bros 3 bootleg. The game was modified (with copyright information cut out) and ported from MMC3.

[<img src="../../wiki-images/SMB3_m106_components.jpg" width="300" height="272" alt="SMB3 m106 components.jpg" />](../../wiki-images/SMB3_m106_components.jpg)

[](../../wiki-images/SMB3_m106_components.jpg)

[](../../wiki-images/SMB3_m106_components.jpg)

[](../../wiki-images/SMB3_m106_components.jpg)

[](../../wiki-images/SMB3_m106_components.jpg)[<img src="../../wiki-images/SMB3_m106_top.jpg" width="300" height="268" alt="SMB3 m106 top.jpg" />](../../wiki-images/SMB3_m106_top.jpg)

[](../../wiki-images/SMB3_m106_top.jpg)

[](../../wiki-images/SMB3_m106_top.jpg)

[](../../wiki-images/SMB3_m106_top.jpg)

[](../../wiki-images/SMB3_m106_top.jpg)[<img src="../../wiki-images/SMB3_m106_bottom.jpg" width="300" height="272" alt="SMB3 m106 bottom.jpg" />](../../wiki-images/SMB3_m106_bottom.jpg)

[](../../wiki-images/SMB3_m106_bottom.jpg)

[](../../wiki-images/SMB3_m106_bottom.jpg)

[](../../wiki-images/SMB3_m106_bottom.jpg)

[](../../wiki-images/SMB3_m106_bottom.jpg)[<img src="../../wiki-images/SMB3_m106_schematics.png" width="300" height="142" alt="SMB3 m106 schematics.png" />](../../wiki-images/SMB3_m106_schematics.png)

[](../../wiki-images/SMB3_m106_schematics.png)

[](../../wiki-images/SMB3_m106_schematics.png)

## Memory map {#Memory_map}

- PRG-ROM: 2x128 KiB
- CHR-ROM: 128 KiB
- PRG-RAM: 8 KiB

- CPU \$6000-\$7FFF: 8 KB PRG RAM bank
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank0
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank1
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank2
- CPU \$E000-\$FFFF: 8 KB switchable PRG ROM bank3

- PPU \$0000-\$03FF 1 KB switchable CHR bank
- PPU \$0400-\$07FF 1 KB switchable CHR bank
- PPU \$0800-\$0BFF 1 KB switchable CHR bank
- PPU \$0C00-\$0FFF 1 KB switchable CHR bank
- PPU \$1000-\$13FF 1 KB switchable CHR bank
- PPU \$1400-\$17FF 1 KB switchable CHR bank
- PPU \$1800-\$1BFF 1 KB switchable CHR bank
- PPU \$1C00-\$1FFF 1 KB switchable CHR bank

Subject to bus conflicts: no

## Registers {#Registers}

There are 16 registers at \$8000-\$800f. Only A3-A0 is taken into account (mask \$800F)

### <span id="CHR_banking_.28.248000-.248007.29"></span>CHR banking (\$8000-\$8007) {#CHR_banking_.28.248000-.248007.29}

    [.CCC CCC0] $8000
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $0000-$03ff

    [.CCC CCC1] $8001
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $0400-$07ff

    [.CCC CCC0] $8002
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $0800-$0bff

    [.CCC CCC1] $8003
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $0c00-$0fff

    [.CCC CCCC] $8004
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $1000-$13ff

    [.CCC CCCC] $8005
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $1400-$17ff

    [.CCC CCCC] $8006
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $1800-$1bff

    [.CCC CCCC] $8007
      ||| ||||
      +++-++++- 1 kB CHR ROM bank at $1c00-$1fff

### <span id="PRG_banking_.28.248008-.24800b.29"></span>PRG banking (\$8008-\$800b) {#PRG_banking_.28.248008-.24800b.29}

    [.... PPPP] $8008
          ||||
          ++++- 8 kB PRG ROM bank at $8000-$9fff (from ROM 5602)

    [...R PPPP] $8009
        | ||||
        | ++++- 8 kB PRG ROM bank at $a000-$bfff
        +------ ROM select at $a000-$bfff (0: 5601, 1: 5602)

    [...R PPPP] $800a
        | ||||
        | ++++- 8 kB PRG ROM bank at $c000-$dfff
        +------ ROM select at $c000-$dfff (0: 5601, 1: 5602)

    [.... PPPP] $800b
          ||||
          ++++- 8 kB PRG ROM bank at $e000-$ffff (from ROM 5602)

Note: if you combine 128K ROM 5601 with 128K ROM 5602 into one 256K ROM, add 16 to \$8008/\$800b and treat RPPPP from \$8009/\$800a as 5-bit bank number.

### <span id="Mirroring_.28.24800c.29"></span>Mirroring (\$800c) {#Mirroring_.28.24800c.29}

    [.... ...M] $800c
             |
             +- mirroring (0: V, 1: H)

### <span id="Counter_and_interrupts_.28.24800d-.24800f.29"></span>Counter and interrupts (\$800d-\$800f) {#Counter_and_interrupts_.28.24800d-.24800f.29}

There is one 16 bit counter, which increments at end of each CPU cycle (on falling edges of M2). The counter will count every cycle that it's not equal to \$FFFF: it cannot be stopped otherwise. While the counter is \$FFFF, if the IRQ is enabled, /IRQ is asserted.

    [.... ....] $800d - writing any value resets counter to 0 and disables interrupts.
                        This necessarily acknowledges a pending interrupt.

    [DDDD DDDD] $800e
     |||| ||||
     ++++-++++- set low byte of counter to this value

    [DDDD DDDD] $800f
     |||| ||||
     ++++-++++- set high byte of counter to this value and enable interrupts.

## Trivia {#Trivia}

- FCEUX implementation is not 100% accurate with hardware. Here are differences:
  - none of the registers are cleared or initialized to a known value at startup,
    - The 74LS670, used in this mapper and [iNES Mapper 246](../INES_Mapper_246/), seems to reliably power up containing all "\$F". m[246](../INES_Mapper_246/) relies on this; this game does not expect registers to be at any known state (tested on FCEUX).
  - CHR registers does not take into account bit 7,
  - writing to \$800d does not stop counter, but starts it and interrupt firing is disabled.
- This bootleg was created for NTSC console. FCEUX implementation of this mapper causes the game to crash immediatelly on title screen in PAL mode (where interrupt comes too early). Improving mapper in the way that real hardware works makes the game work in NTSC/PAL mode, but the game in PAL is almost unplayable, because second nametable, containing status bar, is switched too early, causing half of screen to display garbage.
