---
title: "NES 2.0 Mapper 495"
url: "/wiki/NES_2.0_Mapper_495/"
description: "NES 2.0 Mapper 495 denotes the N-46 PCB, used by one bootleg of Mike Tyson's Punch-Out!! containing a discrete implementation of MMC2-like functionality. In contrast to other known discrete bootlegs…"
wikiPage: "NES_2_0_Mapper_495"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "Mappers triggering on reads"
    url: "/category/Mappers_triggering_on_reads/"
---

**NES 2.0 Mapper 495** denotes the **N-46** PCB, used by one bootleg of *Mike Tyson's Punch-Out!!* containing a discrete implementation of [MMC2](../MMC2/)-like functionality. In contrast to other known discrete bootlegs which maintain near-complete compatibility, this mapper is not directly compatible with MMC2.

## Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$0FFF: Two 4 KiB switchable CHR ROM banks
- PPU \$1000-\$1FFF: Two 4 KiB switchable CHR ROM banks

The two 4 KiB PPU banks each have two 4 KiB banks, which can be switched during rendering by using the special tiles \$FD or \$FE in either a sprite or the background. See [CHR banking](#CHR_banking) below for behavioral differences.

## Registers {#Registers}

### <span id="PRG_ROM_bank_select_.288000-.249FFF.2C_.24A000-.24BFFF.2C_.24C000-.24DFFF.29"></span>PRG ROM bank select (8000-\$9FFF, \$A000-\$BFFF, \$C000-\$DFFF) {#PRG_ROM_bank_select_.288000-.249FFF.2C_.24A000-.24BFFF.2C_.24C000-.24DFFF.29}

    7  bit  0
    ---- ----
    xxxx PPPP
         ||||
         ++++- Select 8 KiB PRG ROM bank

### <span id="CHR_ROM_.24FD.2F0000_bank_select.2FNametable_Arrangement_.28.24E000-.24F3FF.29"></span>CHR ROM \$FD/0000 bank select/Nametable Arrangement (\$E000-\$F3FF) {#CHR_ROM_.24FD.2F0000_bank_select.2FNametable_Arrangement_.28.24E000-.24F3FF.29}

    Mask: $E000

    7  bit  0
    ---- ----
    MMxC CCCC
    || | ||||
    || +-++++- Select 4 KiB CHR ROM bank for PPU $0000-$0FFF
    ||         used when latch 0 = $FD
    ++-------- Nametable Arrangement, see below

### <span id="CHR_ROM_.24FE.2F0000_bank_select.2FNametable_Arrangement_.28.24E400-.24F7FF.29"></span>CHR ROM \$FE/0000 bank select/Nametable Arrangement (\$E400-\$F7FF) {#CHR_ROM_.24FE.2F0000_bank_select.2FNametable_Arrangement_.28.24E400-.24F7FF.29}

    Mask: $E400

    7  bit  0
    ---- ----
    MMxC CCCC
    || | ||||
    || +-++++- Select 4 KiB CHR ROM bank for PPU $0000-$0FFF
    ||         used when latch 0 = $FE
    ++-------- Nametable Arrangement, see below

### <span id="CHR_ROM_.24FD.2F1000_bank_select_.28.24E800-.24FBFF.29"></span>CHR ROM \$FD/1000 bank select (\$E800-\$FBFF) {#CHR_ROM_.24FD.2F1000_bank_select_.28.24E800-.24FBFF.29}

    Mask: $E800

    7  bit  0
    ---- ----
    MMxC CCCC
    || | ||||
    || +-++++- Select 4 KiB CHR ROM bank for PPU $1000-$1FFF
    ||         used when latch 1 = $FD
    ++-------- $3000-$3FFF Mirrored Nametable Arrangement, see below

### <span id="CHR_ROM_.24FE.2F1000_bank_select_.28.24EC00-.24FFFF.29"></span>CHR ROM \$FE/1000 bank select (\$EC00-\$FFFF) {#CHR_ROM_.24FE.2F1000_bank_select_.28.24EC00-.24FFFF.29}

    Mask: $EC00

    7  bit  0
    ---- ----
    MMxC CCCC
    || | ||||
    || +-++++- Select 4 KiB CHR ROM bank for PPU $1000-$1FFF
    ||         used when latch 1 = $FE
    ++-------- $3000-$3FFF Mirrored Nametable Arrangement, see below

## CHR banking {#CHR_banking}

CHR banking functions similarly to MMC2. However, PPU A13, A3, A2, A1, and A0 are ignored. Therefore the latches will react to \$2FDx, \$2FEx, \$3FDx, and \$3FEx in addition to \$0FDx, \$0FEx, \$1FDx, and \$1FEx.

## Nametable Arrangement {#Nametable_Arrangement}

The upper 2 bits of \$E000/\$E400 are separate and yield the following nametable arrangements:

    MM
    00 = S0-S0-S0-S1 (lower right unique, or vertically-flipped L mirroring)
    01 = Vertical arrangement ("Horizontal mirroring")
    10 = Horizontal arrangement ("Vertical mirroring")
    11 = Single-screen, page 1 (1scB mirroring)

Note that which arrangement is used depends on which trigger tile was fetched most recently.

The upper 2 bits of \$E800/\$EC00 operate in the same way, except they affect the mirrored nametables in the normally unused \$3000-\$3FFF PPU region instead.

## References {#References}

- [Forum post:](https://forums.nesdev.org/viewtopic.php?p=301488#p301488) board analysis
