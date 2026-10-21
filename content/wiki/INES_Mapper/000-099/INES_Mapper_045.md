---
title: "INES Mapper 045"
url: "/wiki/INES_Mapper_045/"
description: "iNES Mapper 045 denotes multicart PCBs using the GA23C ASIC in its standard configuration. It is an MMC3 clone with four outer bank registers."
wikiPage: "INES_Mapper_045"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

iNES Mapper 045 denotes multicart PCBs using the GA23C ASIC in its standard configuration. It is an [MMC3](../MMC3/) clone with four outer bank registers.

# Registers {#Registers}

The four outer bank registers are accessed at address \$6000. The first write goes to the first register, the second write to the second, and so on; the fifth write goes to the first register again. The outer bank registers overlay WRAM and function regardless of the MMC3's WRAM bits' setting.

## MMC3-compatible registers {#MMC3-compatible_registers}

- Bank select (\$8000-\$9FFE, even)
- Bank data (\$8001-\$9FFF, odd)
- Mirroring (\$A000-\$BFFE, even)
- PRG RAM protect (\$A001-\$BFFF, odd)
- IRQ latch (\$C000-\$DFFE, even)
- IRQ reload (\$C001-\$DFFF, odd)
- IRQ disable (\$E000-\$FFFE, even)
- IRQ enable (\$E001-\$FFFF, odd)

These registers function identically to the [MMC3](../MMC3/), with the PRG/CHR bank registers subject to masking by the outer bank registers.

## <span id="CHR-OR_LSB_.28.246000_.230.2C_write.29"></span>CHR-OR LSB (\$6000 \#0, write) {#CHR-OR_LSB_.28.246000_.230.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      CCCC CCCC
      ++++-++++- Select CHR A10-A17, OR'd with MMC3's CHR A10-A17 masked according to CHR-AND

## <span id="PRG-OR_LSB_.28.246000_.231.2C_write.29"></span>PRG-OR LSB (\$6000 \#1, write) {#PRG-OR_LSB_.28.246000_.231.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      ppPP PPPP
      ||++-++++- Select PRG A13-A18, OR'd with MMC3's PRG A13-A18 masked according to PRG-AND
      ++-------- Select PRG A19-A20

## <span id="CHR-AND.2C_CHR-OR.2FPRG-OR_MSB_.28.246000_.232.2C_write.29"></span>CHR-AND, CHR-OR/PRG-OR MSB (\$6000 \#2, write) {#CHR-AND.2C_CHR-OR.2FPRG-OR_MSB_.28.246000_.232.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      PPCC cccc
      |||| ++++- Select number of CHR bits from MMC3
      ||||       ($F: 256 KiB, $E: 128 KiB ... $7-$0: 1 KiB)
      ||++------ Select CHR A18-A19
      ++-------- Select CHR A20-A21 and PRG A21-22

## <span id="PRG-AND.2C_register_lock_.28.246000_.233.2C_write.29"></span>PRG-AND, register lock (\$6000 \#3, write) {#PRG-AND.2C_register_lock_.28.246000_.233.2C_write.29}

    Mask: $F001

    D~7654 3210
      ---------
      1LPP PPPP
       |++ ++++- Select PRG-AND mask (inverted)
       |||       ($00: 512 KiB, $20: 256 KiB...)
       +-------- 1=Lock outer bank registers

## <span id="Reset_outer_bank_registers_and_release_lock_.28.246001.2C_write.29"></span>Reset outer bank registers and release lock (\$6001, write) {#Reset_outer_bank_registers_and_release_lock_.28.246001.2C_write.29}

Writing any value to \$6001 (mask: \$F001) resets the outer bank registers as a soft reset would, clearing the "Lock" bit and making the next write to \$6000 go to register \#0.

## <span id="Menu_Selection_.231_.28.245000-.245FFF.2C_read.29"></span>Menu Selection \#1 (\$5000-\$5FFF, read) {#Menu_Selection_.231_.28.245000-.245FFF.2C_read.29}

The *超强年度新卡 (Super New Year Cart) 15-in-1* multicart has a DIP switch whose setting causes the menu program to select one of eight different menus.

    A~FEDC BA98 7654 3210  D~7654 3210
      -------------------    ---------
      0101 AAAA AAAA ....    .... ...D
           ++++-++++-----------------+- 1=If DIP switch is in the position selected by the corresponding A bit

This means that:

- If the dip switch is in the "0" position, reads from \$5010 return with D0 set while reads from \$5020, \$5040...\$5800 return with D0 clear.
- If the dip switch is in the "1" position, reads from \$5020 return with D0 set while reads from \$5010, \$5040...\$5800 return with D0 clear.
- If the dip switch is in the "7" position, reads from \$5800 return with D0 set while reads from \$5010, \$5020...\$5400 return with D0 clear.

## <span id="Menu_Selection_.232"></span>Menu Selection \#2 {#Menu_Selection_.232}

Some multicarts select between five different menus by connecting one of the higher address lines to PRG /CE. The menu code selects between menus by checking which of the higher address lines disables PRG-ROM when set.

    Selection 0: No address line connected to PRG /CE
    Selection 1: PRG A20 connected to PRG /CE
    Selection 2: CHR A20 connected to PRG /CE
    Selection 3: PRG A19 connected to PRG /CE
    Selection 4: CHR A19 connected to PRG /CE

Cartridges with 1 MiB of PRG-/CHR-ROM cannot use selections 3/4, cartridges with 2 MiB of more of PRG-/CHR-ROM cannot use this menu selection mechanism.

# Notes {#Notes}

- Some multicart menus write the same values to the outer bank registers multiple times. Experiments during dumping do not suggest that this is necessary.
- NES 2.0 Mappers [356](../NES_2.0_Mapper_356/), [372](../NES_2.0_Mapper_372/) and [373](../NES_2.0_Mapper_373/) repurpose bits from outer bank register \#2 for other things.
- [NES 2.0 Mapper 269](../NES_2.0_Mapper_269/) is an enhanced plug-and-play console variant of this mapper.
- [NES 2.0 Mapper 556](../NES_2.0_Mapper_556/) is similar to this mapper, but the register is moved to \$5000 and repurposed a bit in the second write to select MMC3/VRC4 mode.
