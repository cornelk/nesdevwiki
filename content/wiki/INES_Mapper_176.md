---
title: "INES Mapper 176"
url: "/wiki/INES_Mapper_176/"
description: "iNES Mapper 176 denotes the 8025 enhanced MMC3 chipset. It used by many multicarts, Chinese single-game and educational computer cartridges, and Techno Source's Intellivision X2 Plug-and-Play…"
wikiPage: "INES_Mapper_176"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with large PRG RAM"
    url: "/category/Mappers_with_large_PRG_RAM/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

**iNES Mapper 176** denotes the **8025** enhanced [MMC3](../MMC3/) chipset. It used by many multicarts, Chinese single-game and educational computer cartridges, and Techno Source's *Intellivision X2* Plug-and-Play console. Incompatible variations exist that are denoted via NES 2.0 Submapper.

<table class="wikitable">
<thead>
<tr>
<th colspan="13">Submappers</th>
</tr>
</thead>
<tbody>
<tr>
<th rowspan="2">#</th>
<th rowspan="2">PCB codes</th>
<th rowspan="2"><a href="../UNIF/">UNIF</a> MAPR</th>
<th colspan="5">MMC3</th>
<th colspan="4">Outer bank registers</th>
<th rowspan="2">Example</th>
</tr>
<tr>
<th>Ext. mode</th>
<th>PRG bits</th>
<th>$46/47</th>
<th><a href="../Mirroring/#Single-Screen">1SM</a></th>
<th>Extra WRAM</th>
<th>Address mask</th>
<th>PRG A21+</th>
<th>CHR A21+</th>
<th>CNROM latch</th>
</tr>
&#10;<tr>
<td>0</td>
<td>LP-8002KB, SFC-12B</td>
<td><strong>BMC-Super24in1SC03</strong></td>
<td>-</td>
<td>6</td>
<td>normal</td>
<td>-</td>
<td>-</td>
<td>$Fxx3</td>
<td>-</td>
<td>-</td>
<td>-</td>
<td><em>YH-xxx</em> multicarts, <em>Rockman 6-in-1</em></td>
</tr>
<tr>
<td>1</td>
<td>FK-xxx/BS-xxx</td>
<td><strong>BMC-FK23C</strong>/<strong>BMC-FK23CA</strong></td>
<td>$5FF3.1</td>
<td>8</td>
<td>normal</td>
<td>-</td>
<td>-</td>
<td>$Fxx3</td>
<td>-</td>
<td>-</td>
<td>yes</td>
<td><em>FK-xxxx</em> multicarts</td>
</tr>
<tr>
<td>2</td>
<td>FS005/FS006</td>
<td><strong>WAIXING-FS005</strong></td>
<td>$5FF3.1</td>
<td>6</td>
<td>swapped</td>
<td>yes</td>
<td>$A001.0-1</td>
<td>$Fxx3</td>
<td>$5xx0.3/7, $5xx2.6-7/5</td>
<td>-</td>
<td>-</td>
<td>Waixing 2005+ re-releases, <em>245-in-1 Real Game</em></td>
</tr>
<tr>
<td>3</td>
<td>JX9003B</td>
<td>-</td>
<td>-</td>
<td>8</td>
<td>normal</td>
<td>-</td>
<td>-</td>
<td>$Fxx7</td>
<td>$5xx5</td>
<td>$5xx6</td>
<td>-</td>
<td><em>Super Mario 160-in-1 Funny Time</em></td>
</tr>
<tr>
<td>4</td>
<td>?</td>
<td>-</td>
<td>-</td>
<td>6</td>
<td>normal</td>
<td>-</td>
<td>-</td>
<td>$Fxx3</td>
<td>$5xx2.7</td>
<td>-</td>
<td>-</td>
<td><em>GameStar Smart Genius Deluxe</em></td>
</tr>
<tr>
<td>5</td>
<td>HST-162</td>
<td>-</td>
<td>-</td>
<td>6</td>
<td>normal</td>
<td>-</td>
<td>-</td>
<td>$Fxx3</td>
<td>$4800</td>
<td>-</td>
<td>-</td>
<td><em>Game 500-in-1</em></td>
</tr>
</tbody>
</table>

- The **SFC-12B** PCB (submapper 0) mounts both CHR-ROM and CHR-RAM that are selected via \$5xx0.5.
- Eight MMC3 PRG bits mean that bit 6 and 7 of \$8000.6/7 are applied, and the fixed banks are \$FE/\$FF rather than \$3E/\$3F. This implies that the multicart's reset vector lies at address PRG-ROM address \$1FFFFC rather than \$7FFFC if the PRG-ROM is that large.

# Registers {#Registers}

**FS005 (Submapper 2)** can disable the registers in the \$5000-\$5FFF range using the [RAM Configuration Register (\$A001)](#RAM_Configuration_Register_.28.24A001.29). The same register also enables a mixed CHR-ROM/CHR-RAM mode that is similar to [INES Mapper 195](../INES_Mapper_195/).

## <span id="Mode_Register_.28.245xx0.29"></span>Mode Register (\$5xx0) {#Mode_Register_.28.245xx0.29}

Mask: \$F*xx*7 (**Submapper 3**), \$5*xx*3 (all others), *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    PCTm PMMM
    |||| ||||
    |||| |+++- Select PRG Banking Mode (ignored in Extended MMC3 Mode)
    |||| |      0: MMC3 PRG Mode, 2 MiB/512 KiB Outer PRG Bank Size
    |||| |      1: MMC3 PRG Mode, 256 KiB Outer PRG Bank Size
    |||| |      2: MMC3 PRG Mode, 128 KiB Outer PRG Bank Size
    |||| |      3: NROM-128 PRG Mode, 16 KiB PRG at $8000-$BFFF mirrored at $C000-$FFFF
    |||| |      4: NROM-256 PRG Mode, 32 KiB PRG at $8000-$FFFF
    |||| |      5: UNROM PRG Mode, 16 KiB PRG at $8000-$BFFF, inner bank #7 at $C000-$FFFF
    |||| |      6-7: Never used
    |||| +---- PRG A21 (Submapper 2 only)
    |||+------ Select Outer CHR Bank Size
    |||         0: In MMC3 CHR Mode: 256 KiB
    |||            In CNROM CHR Mode: 32 KiB (Submapper 1 only)
    |||         1: In MMC3 CHR Mode: 128 KiB
    |||            In CNROM CHR Mode: 16 KiB (Submapper 1 only)
    ||+------- Submappers 0/1 with both CHR-ROM and CHR-RAM (e.g. SFC-12B PCB):
    ||          0: Select CHR-ROM
    ||          1: Select CHR-RAM
    ||         Submapper 1 with $5xx0.6=1:
    ||          0: Enable CNROM latch (CNROM mode)
    ||          1: Disable CNROM latch (NROM mode)
    |+-------- CHR A10-12 Mode
    |           0: from MMC3
    |           1: from PPU
    +--------- PRG A22 (Submapper 2 only)

    Power-on value: $00

## <span id="PRG_Base_Register_LSB_.28.245xx1.29"></span>PRG Base Register LSB (\$5xx1) {#PRG_Base_Register_LSB_.28.245xx1.29}

Mask: \$F*xx*7 (**Submapper 3**), \$F*xx*3 (all others), *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    .PPP PPPP
     ||| ||||
     +++-++++- PRG A20..A14

    Power-on value: $00

Depending on the PRG banking mode set via \$5xx0, only the higher bits of this register are applied:

    PRG A..
    21111111
    09876543
    --------
    MMMMMMMM  Submappers 1/3 in PRG Mode 0 (MMC3 2 MiB) or Extended MMC3 Mode
    BBMMMMMM  Submappers 0/2/4 in PRG Mode 0 (MMC3 512 KiB)
    BBBMMMMM  PRG Mode 1 (MMC3 256 KiB)
    BBBBMMMM  PRG Mode 2 (MMC3 128 KiB)
    BBBBBBBC  PRG Mode 3 (NROM-128)
    BBBBBBCC  PRG Mode 4 (NROM-256)
    BBBBLLLC  PRG Mode 5 (UNROM)

    M: Bit comes from MMC3 ($8000.6/7 or fixed bank)
    B: Bit comes from PRG Base Register ($5xx1)
    L: Bit comes from UNROM Latch ($8000-$FFFF or fixed bank)
    C: Bit comes from CPU

**Submapper 5** only selects PRG A14-A18 (i.e. 512 KiB) via this register, selecting upper PRG bits via register [\$4800](#PRG_Base_Register_MSB_.28.244800.29.2C_Submapper_5_only).

## <span id="PRG_Base_Register_MSB_.28.245xx5.29.2C_Submapper_3_only"></span>PRG Base Register MSB (\$5xx5), Submapper 3 only {#PRG_Base_Register_MSB_.28.245xx5.29.2C_Submapper_3_only}

Mask: \$F*xx*7 (**Submapper 3**), *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    .... PPPP
         ||||
         ++++- PRG Base A24..A21

    Power-on value: $00

## <span id="PRG_Base_Register_MSB_.28.244800.29.2C_Submapper_5_only"></span>PRG Base Register MSB (\$4800), Submapper 5 only {#PRG_Base_Register_MSB_.28.244800.29.2C_Submapper_5_only}

Mask: \$F800

    7654 3210
    ---- ----
    ..PP PPPP
      || ||||
      ++-++++- PRG Base A24..A19

    Power-on value: $00

## <span id="CHR_Base_Register_LSB_.28.245xx2.29"></span>CHR Base Register LSB (\$5xx2) {#CHR_Base_Register_LSB_.28.245xx2.29}

Mask: \$F*xx*7 (**Submapper 3**), \$5*xx*3 (all others), *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    ccdC CCCC
    |||| ||||
    ++++-++++- CHR A20..A13
    ||+------- PRG A25 (Submapper 2 only)
    ++-------- PRG A24..A23 (Submapper 2 only)
    +--------- PRG A21 (Submapper 4)

    Power-on value: $00

Depending on the CHR banking mode set via \$5xx0, only the higher bits of this register are applied:

    CHR A..
    21111111111
    09876543210
    -----------
    BBBMMMMMMMM  MMC3 CHR Mode ($5xx0.6=0), 256 KiB ($5xx0.4=0)
    BBBBMMMMMMM  MMC3 CHR Mode ($5xx0.6=0), 128 KiB ($5xx0.4=1)
    BBBBBBLLPPP  CNROM CHR Mode ($5xx0.6=1, $5xx0.5=0), 32 KiB ($5xx0.4=0), Submapper 1 only
    BBBBBBBLPPP  CNROM CHR Mode ($5xx0.6=1, $5xx0.5=0), 16 KiB ($5xx0.4=1), Submapper 1 only
    BBBBBBBBPPP  NROM CHR Mode ($5xx0.6=1, $5xx0.5=1 or Submapper other than 1)

    M: Bit comes from MMC3 ($8000.0-5)
    B: Bit comes from CHR Base Register ($5xx2)
    L: Bit comes from CNROM Latch ($8000-$FFFF)
    P: Bit comes from PPU

## <span id="CHR_Base_Register_MSB_.28.245xx6.29.2C_Submapper_3_only"></span>CHR Base Register MSB (\$5xx6), Submapper 3 only {#CHR_Base_Register_MSB_.28.245xx6.29.2C_Submapper_3_only}

Mask: \$F*xx*7 (**Submapper 3**), *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    .... PPPP
         ||||
         ++++- CHR Base A24..A21

    Power-on value: $00

## <span id="Extended_Mode_Register_.28.245xx3.29.2C_Submappers_1-2_only"></span>Extended Mode Register (\$5xx3), Submappers 1-2 only {#Extended_Mode_Register_.28.245xx3.29.2C_Submappers_1-2_only}

Mask: \$F*xx*3, *x* determined by [solder pad setting](#Solder_Pad)

    7654 3210
    ---- ----
    .?.. .?E.
           | 
           +- Extended MMC3 Mode
               0: disable
               1: enable

    Power-on value: $00

Most multicarts write value \$44 rather than \$00 to disable Extended MMC3 Mode. Hardware tests do not indicate any difference in behavior between writing \$00 and \$44.

## <span id="Mirroring_Register_.28.24A000.29"></span>Mirroring Register (\$A000) {#Mirroring_Register_.28.24A000.29}

    Mask: $E003

    7654 3210
    ---- ----
    .... ..MM
           ++- Select nametable mirroring 
               0: Vertical
               1: Horizontal
               2: Single-screen, page 0 (Submapper 2 only)
               3: Single-screen, page 1 (Submapper 2 only)
    Power-on value: $00

Single-screen mirroring is only available when the RAM Configuration Register is enabled (\$A001.5).

## <span id="RAM_Configuration_Register_.28.24A001.29.2C_Submapper_2_only"></span>RAM Configuration Register (\$A001), Submapper 2 only {#RAM_Configuration_Register_.28.24A001.29.2C_Submapper_2_only}

Mask: \$E003

This register functions like MMC3 register \$A001 until bit 5 is set, which turns it into the RAM Configuration Register.

    7654 3210
    ---- ----
    RFE. SCWW
    |||  ||||
    |||  ||++- Select 8 KiB PRG-RAM bank at $6000-$7FFF. Ignored if Bit 5 is clear.
    |||  |+--- Select the memory type in the first 8 KiB of CHR space. Ignored if Bit 5 is clear.
    |||  |      0: First 8 KiB are CHR-ROM
    |||  |      1: First 8 KiB are CHR-RAM
    |||  +---- Unknown
    ||+------- RAM Configuration Register Enable
    ||          0: RAM Configuration Register disabled, $A001 functions as on MMC3, 8 KiB of WRAM
    ||          1: RAM Configuration Register enabled, 32 KiB of WRAM
    |+-------- Outer Bank Registers Enable. Ignored if Bit 5 is clear.
    |           0: Outer Bank Registers disabled, $5000-$5FFF maps to the second 4 KiB of the 8 KiB WRAM bank 2
    |           1: Outer Bank Registers enabled in the $5000-$5FFF range
    +--------- PRG RAM enable (0: disable, 1: enable)

    Power-on value: $00

## <span id="UNROM_latch_.28.248000-.24FFFF.29"></span>UNROM latch (\$8000-\$FFFF) {#UNROM_latch_.28.248000-.24FFFF.29}

In UNROM Mode (PRG Mode 5), writing to this address range changes the 16 KiB inner PRG bank at \$8000-\$BFFF.

## <span id="CNROM_latch_.28.248000-.24FFFF.29.2C_Submapper_1_only"></span>CNROM latch (\$8000-\$FFFF), Submapper 1 only {#CNROM_latch_.28.248000-.24FFFF.29.2C_Submapper_1_only}

In CNROM Mode, writing to this address range changes the inner CHR bank.

## <span id="MMC3-compatible_registers.2C_Extended_MMC3_Mode_.28.248000.2F.248001.2C_.24C000.2F.24C001.2C_.24E000.2F.24E001.29"></span>MMC3-compatible registers, Extended MMC3 Mode (\$8000/\$8001, \$C000/\$C001, \$E000/\$E001) {#MMC3-compatible_registers.2C_Extended_MMC3_Mode_.28.248000.2F.248001.2C_.24C000.2F.24C001.2C_.24E000.2F.24E001.29}

Mask: \$E003 (verified on real hardware)

Some multicart games depend on writes to \$9FFF not doing anything as a result of the MMC3 address mask being \$E003 rather than the standard \$E001.

If the "Extended MMC3 Mode" bit in register \$5xx3 is clear, then these registers function identically to the [MMC3](../MMC3/). On **Submappers 1 and 2**, if the "Extended MMC3 Mode" bit is set, four more bank registers become available at [\$8000/\$8001](../MMC3/#Bank_select_.28.248000-.249FFE.2C_even.29), so that the original two 2 KiB CHR banks become four 1 KiB CHR banks, and the two fixed 8 KiB PRG banks become selectable, similar to the [RAMBO-1](../RAMBO-1/). Register \$8000 if \$5xx3 bit 1 is set:

    7  bit  0
    ---- ----
    CP.. RRRR
    ||   ||||
    ||   ++++- Specify which bank register to update on next write to Bank Data register
    ||         $0: Select 1 KB CHR bank at PPU $0000-$03FF (or $1000-$13FF)
    ||         $1: Select 1 KB CHR bank at PPU $0800-$0BFF (or $1800-$1BFF)
    ||         $2: Select 1 KB CHR bank at PPU $1000-$13FF (or $0000-$03FF)
    ||         $3: Select 1 KB CHR bank at PPU $1400-$17FF (or $0400-$07FF)
    ||         $4: Select 1 KB CHR bank at PPU $1800-$1BFF (or $0800-$0BFF)
    ||         $5: Select 1 KB CHR bank at PPU $1C00-$1FFF (or $0C00-$0FFF)
    ||         $6: Select 8 KB PRG ROM bank at $8000-$9FFF (or $C000-$DFFF)
    ||         $7: Select 8 KB PRG ROM bank at $A000-$BFFF
    ||         $8: Select 8 KB PRG ROM bank at $C000-$DFFF (or $8000-$9FFF)
    ||         $9: Select 8 KB PRG ROM bank at $E000-$FFFF
    ||         $A: Select 1 KB CHR bank at PPU $0400-$07FF (or $1400-$17FF)
    ||         $B: Select 1 KB CHR bank at PPU $0C00-$0FFF (or $1C00-$1FFF)
    |+-------- Invert PRG A14
    +--------- Invert CHR A12

    Power-on values:
    * Standard MMC3 Registers $0-$7: $00, $02, $04, $05, $06, $07, $00, $01
    * Extended MMC3 Registers $8-$B: $FE, $FF, $FF, $FF

# Solder Pad {#Solder_Pad}

The address mask in the \$5000-\$5FFF range is determined by the solder pad setting:

    Pad setting  Address mask
    -----------  ------------
    0            $5013
    1            $5023
    2            $5043
    3            $5083
    4            $5103
    5            $5203
    6            $5403
    7            $5803

- A solder pad setting of zero (address mask \$5013) will produce a usable result for any ROM image.
- Some multicarts only display their menu at settings other than 0.

# <span id="Protection_.28Submapper_2_only.29"></span>Protection (Submapper 2 only) {#Protection_.28Submapper_2_only.29}

Later Waixing games (and re-releases of earlier games) use the RAM Configuration Register for copy-protection purposes:

- Write \$A1 to \$A001: Address range \$5000-\$5FFF to second half of 8 KiB WRAM bank 2, mapper registers there are disabled.
- Write three values to \$5000, \$5010 and \$5013.
- Do further initialization.
- Write \$E2 to \$A001. Mapper registers in address range \$5000-\$5FFF; WRAM at CPU \$6000-\$7FFF points to 8 KiB WRAM bank 2.
- Copy 20 bytes from \$7000 to \$6000.
- Copy and XOR bytes from \$6000, \$6010 and \$6013 to \$0100-\$0102.
- Execute code at CPU \$0100.

Hacked ROMs can be detected by them writing to \$5000/\$5010/\$5013 but then no longer jumping to \$0100.

# See also {#See_also}

- [NES 2.0 Mapper 523](../NES_2.0_Mapper_523/) is a variant of this mapper with hard-wired mirroring that connects CHR-ROM differently to produce 4/2 KiB instead of 2/1 KiB banks.
- [MMC3 with CHR ROM and CHR RAM](../MMC3_with_CHR_ROM_and_CHR_RAM/)
