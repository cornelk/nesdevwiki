---
title: "INES Mapper 083"
url: "/wiki/INES_Mapper_083/"
description: "iNES Mapper 83 and NES 2.0 Mapper 264 are used for games by Cony, also known as Yoko. The distinction between mappers 83 and 264 as well as NES 2.0 Submappers denote different printed circuit boards…"
wikiPage: "INES_Mapper_083"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with large PRG RAM"
    url: "/category/Mappers_with_large_PRG_RAM/"
---

**iNES Mapper 83** and **NES 2.0 Mapper 264** are used for games by Cony, also known as Yoko. The distinction between mappers 83 and 264 as well as NES 2.0 Submappers denote different printed circuit boards that all incorporate the same chip-on-board ASIC.

- **Mapper 83.0**: Standard single-game variant, 256 KiB of CHR-ROM maximum: *Street Fighter II Pro* and variants, *Fatal Fury 2*.
- **Mapper 83.1**: **NC-24/82-07-17J** PCB. Large CHR single-game variant, 512 KiB of CHR-ROM maximum: *World Heroes 2'*, *Super Blaster VII Turbo 28*.
- **Mapper 264**: **NC-14/82-05-26J** PCB. Large CHR 2-in-1 multicart variant: *Mortal Kombat II/V Pro*, *Master Fighter VI*. UNIF MAPR: **UNL-YOKO**.
- **Mapper 83.2**: **N-C88-1603/83-08-19J** PCB. 4-in-1 256 KiB PRG+256 KiB CHR+32 KiB PRG-NVRAM multicart variant: *Dragon Ball Party*.
- **Mapper 83.3**: **82-08-24J** PCB. 4-in-1 128 KiB PRG+128 KiB+no PRG-RAM multicart variant: *1994 超強 20合1*.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable window into 32 KiB PRG-NVRAM (Mapper **83.2**)/into PRG-ROM (all others)
- CPU \$8000-\$FFFF. 2x16 KiB or 4x8 KiB switchable windows into PRG-ROM
- PPU \$0000-\$1FFF: 4x2 KiB (Mappers **83.1** and **264**)/8x1 KiB (all others) switchable windows into CHR-ROM

# Registers {#Registers}

Only mapper 83's addresses given. Mapper 264 connects CPU A11.A10 to the ASIC's A9..A8 inputs, so all addresses change accordingly:

    Mapper 83  Mappper 264
    $5100      $5400 
    $8100      $8400
    $8200      $8800
    $8300      $8C00
    $8318      $8C18

## <span id="PRG_Base_.28.248000.2C_write.29"></span>PRG Base (\$8000, write) {#PRG_Base_.28.248000.2C_write.29}

    D~[PPPP PPPP] (Address AND $8300)=$8000  Mapper 83.0, 83.1, 264
       ++++-++++-- PRG A21..A14

    D~[RRQQ PPPP] (Address AND $8300)=$8000  Mapper 83.2
       |||| ++++-- PRG A17..A14
       ||++------- PRG/CHR A19..A18
       ++--------- PRG-RAM A14..A13

    D~[CCPP PPPP] (Address AND $8300)=$8000  Mapper 83.3
       ||++ ++++-- PRG A19..A14
       ++--------- CHR-ROM A19..A18

The bits from the PRG Base register are modified according to the [PRG banking mode](#Mode_Register_.28.248100.2C_write.29).

## <span id="Mode_Register_.28.248100.2C_write.29"></span>Mode Register (\$8100, write) {#Mode_Register_.28.248100.2C_write.29}

    D~[CD6P P.NN] (Address AND $8300)=$8100
       |||| | ++-- Nametable arrangement
       |||| |       0: Horizontal (Vertical mirroring)
       |||| |       1: Vertical (Horizontal mirroring)
       |||| |       2: Single screen 0
       |||| |       3: Single screen 1
       |||+-+----- PRG Base bits modification (PRG banking mode)
       |||          0: PRG A17/A16..A14=1 if CPU A14=1 (UxROM)
       |||          1: None (NROM-128)
       |||          2: PRG A17/A16..A13 from 4x8 KiB PRG registers
       |||          3: Same as 2
       ||+-------- 1: Map 8 KiB PRG-ROM bank to CPU $6000
       |+--------- 0: IRQ counter increments, 1: IRQ counter decrements
       +---------- 1: IRQ counting enabled on next reload

- PRG Base A17 is only modified on mappers **83.0**, **83.1** and **83.2**, and not on **83.3** and **264**. In other words, the former have a 256 KiB inner PRG bank, and the latter a 128 KiB inner PRG bank.
- The **6** bit has no effect on mapper **83.2**, as the circuit board hardware around the ASIC forces CPU \$6000-\$7FFF to PRG-RAM instead.

## <span id="IRQ_Counter_.28.248200-.248201.2C_write.29"></span>IRQ Counter (\$8200-\$8201, write) {#IRQ_Counter_.28.248200-.248201.2C_write.29}

    D~[LLLL LLLL] (Address AND $8301)=$8200
       ++++-++++- IRQ Counter LSB
    D~[MMMM MMMM] (Address AND $8301)=$8201
       ++++-++++- IRQ Counter MSB

Writing directly changes the current counter value. Writing to the LSB acknowledges a pending IRQ; writing to the MSB enables counting if bit **C** in the [Mode register](#Mode_Register_.28.248100.2C_write.29) was previously set.

The counter can be [clocked by a CPU M2 cycle or by a rising edge of PPU A12](#IRQ_Counter_Source_.28.248318.2C_write.29). When counting is enabled and the counter is not zero, the counter is incremented or decremented according to the [Mode register](#Mode_Register_.28.248100.2C_write.29)'s **D** bit. If the counter is then zero, counting is disabled, and an IRQ is raised.

## <span id="8_KiB_PRG-ROM_Bank_.28.248300-.248303.2C_write.29"></span>8 KiB PRG-ROM Bank (\$8300-\$8303, write) {#8_KiB_PRG-ROM_Bank_.28.248300-.248303.2C_write.29}

    D~[...p PPPP] A~[1... ..11 .... 00RR] (Address AND $8310)=$8300
          | ||||                      ++- Apply to which 8 KiB CPU window ($8000/$A000/$C000/$6000)
          +-++++------------------------- PRG-ROM A17/A16..A13

The **p** bit (PRG A17) has no effect on mappers **83.3** and **264**. Register \$8303 applies to the \$6000-\$7FFF window, if enabled in the Mode register; the \$E000-\$FFFF window is hard-wired to the last (inner) bank.

## <span id="1_KiB_CHR-ROM_Bank_.28.248310-.248317.2C_write.29"></span>1 KiB CHR-ROM Bank (\$8310-\$8317, write) {#1_KiB_CHR-ROM_Bank_.28.248310-.248317.2C_write.29}

    D~[CCCC CCCC] A~[1... ..11 ...1 0RRR] (Address AND $8318)=$8310
       |||| ||||                     +++- Apply to which 1 KiB PPU bank ($0000/$0400/$0800/$0C00/$1000/$1400/$1800/$1C00)
       ++++-++++------------------------- CHR-ROM A17..A10

Mappers **83.1** and **264** increase the CHR-ROM size from 256 to 512 KiB by connecting CHR-ROM A10 to PPU A10 and CHR-ROM A11..A18 to the mapper ASIC's A10..A17 outputs, similar to [INES Mapper 197](../INES_Mapper_197/). More precisely:

    Pin            Connected to
    PPU A10        CHR-ROM A10
    PPU A11        ASIC PA10
    PPU A12        ASIC PA11 and PA12
    ASIC CHR A10   CHR-ROM A11
        :             :
    ASIC CHR A17   CHR-ROM A18

causing the ASIC to see the following addresses:

    Real PPU Addr  ASIC PPU Addr
    $0000          $0000 -> reg 0
    $0400          $0000 -> reg 0
    $0800          $0400 -> reg 1
    $0C00          $0400 -> reg 1
    $1000          $1800 -> reg 6
    $1400          $1800 -> reg 6
    $1800          $1C00 -> reg 7
    $1C00          $1C00 -> reg 7

As a result, CHR-ROM is banked in 2 KiB rather than 1 KiB amounts, using on the the four CHR bank registers 0, 1, 6 and 7.

## <span id="IRQ_Counter_Source_.28.248318.2C_write.29"></span>IRQ Counter Source (\$8318, write) {#IRQ_Counter_Source_.28.248318.2C_write.29}

    D~[???? ????] A~[1... ..11 ...1 1???] (Address AND $8318)=$8318
       ++++-++++- 1=PPU A12 rising edge, 0=CPU M2

The default value \$00 selects CPU M2 cycles as the source of the IRQ counter, and a value of \$FF selects unfiltered PPU A12 rising edges as the source, clocking the counter eight times per rendered scanline. It's not known which of the eight data bits is decisive.

## <span id="Solder_pad_.28.245000.2C_read.29"></span>Solder pad (\$5000, read) {#Solder_pad_.28.245000.2C_read.29}

    D~[.... ..PP] (Address AND $F100)=$5000
              ++- Solder pad value

Several games display different names (*Street Fighter* versus *Street Blaster*) depending on the solder pad value. *Master Fighter VI* uses it to select NTSC or PAL timing routines. Yet others will just display a garbled title screen at some pad values, and the "good" value is not always 0 or 3.

## <span id="Scratch_RAM_.28.245100-.245103.2C_read.2Fwrite.29"></span>Scratch RAM (\$5100-\$5103, read/write) {#Scratch_RAM_.28.245100-.245103.2C_read.2Fwrite.29}

    D~[VVVV VVVV] A~[0101 ..01 .... 00AA] (Address AND $8310)=$8300
       |||| ||||                      ++- Scratch RAM address A1..A0
       ++++-++++------------------------- Scratch RAM value
