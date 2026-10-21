---
title: "NES 2.0 Mapper 432"
url: "/wiki/NES_2.0_Mapper_432/"
description: "NES 2.0 Mapper 432 denotes a series of MMC3-based Realtec multicart circuit boards."
wikiPage: "NES_2_0_Mapper_432"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

**NES 2.0 Mapper 432** denotes a series of MMC3-based Realtec multicart circuit boards.

- Submapper 0: Realtec 8043, 8086, 8090, GN-30C PCB
- Submapper 1: Realtec 8023 PCB
- Submapper 2: Realtec 8085 PCB
- Submapper 3: Realtec 8145 PCB
- Submapper 4: 940702 PCB

## <span id="Solder_Pad_Enable_Register_.28.246000-.247FFE.2C_even.29"></span>Solder Pad Enable Register (\$6000-\$7FFE, even) {#Solder_Pad_Enable_Register_.28.246000-.247FFE.2C_even.29}

    Mask: $E001

    D~7654 3210
      ---------
      .... ...S
              +- 1=Replace PRG-ROM with solder pad value
                   in the CPU $8000-$FFFF address range

- Submapper 1 selects the solder pad via 6001.5 instead, and this register has no function.
- Submapper 4 just disables PRG-ROM when S=1 and the solder pad is connected.

## <span id="Outer_Bank_Register_.28.246001-.247FFF.2C_odd.29"></span>Outer Bank Register (\$6001-\$7FFF, odd) {#Outer_Bank_Register_.28.246001-.247FFF.2C_odd.29}

    Mask: $E001

    D~7654 3210
      ---------
      pNQP CmMb
      |||| |||+- PRG/CHR A17 if M=1
      |||| ||+-- PRG A17 mode
      |||| ||     0: from MMC3 (256 KiB inner bank)
      |||| ||     1: from b (128 KiB inner bank)
      |||| |+--- CHR A17 mode
      |||| |      0: from MMC3 (256 KiB inner bank)
      |||| |      1: from b (128 KiB inner bank)
      |||| +---- CHR A18
      |||+------ PRG A18
      ||+------- Submapper-specific functionality
      ||         Submapper 0: PRG+CHR A19
      ||         Submapper 1: Read solder pad
      ||         Submapper 2: NROM-256 mode
      ||         Submapper 3: 512 KiB MMC3 CHR mode (like INES Mapper 197 Submapper 2)
      |+-------- NROM-128 mode (1=enabled)
      +--------- Submapper 0/1: NROM-256 mode
                 Submapper 2/3: Lock Outer Bank register

As it uses the MMC3 clones's WRAM interface, writing to the outer bank register requires enabling and not write-protecting WRAM in the MMC's \$A001 register. NROM-128 mode forces the MMC3 clone's CPU A14 input low, causing MMC3 register 6 to select both the bank at \$8000 and \$C000, and register 7 to select both the bank at \$A000 and \$E000.

## <span id="MMC3-compatible_registers_.28.248000-.24FFFF.29"></span>MMC3-compatible registers (\$8000-\$FFFF) {#MMC3-compatible_registers_.28.248000-.24FFFF.29}

Mask: \$E001

See [MMC3](../MMC3/).
