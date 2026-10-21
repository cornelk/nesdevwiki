---
title: "Sachen 8259"
url: "/wiki/Sachen_8259/"
description: "Sachen 8259 describes a series of boards by Sachen. The following information was gathered primarily from FCEU-mm and while it appears to be a good enough behavioral description to run many ROMs, it…"
wikiPage: "Sachen_8259"
---

Sachen 8259 describes a series of boards by Sachen. The following information was gathered primarily from FCEU-mm and while it appears to be a good enough behavioral description to run many ROMs, it is likely incomplete.

## Overview {#Overview}

- PRG ROM size: up to 128 KiB
- PRG ROM bank size: 32 KiB
- PRG RAM: No
- CHR capacity: up to 512KiB ROM, depending on variant
- CHR bank size: 2 KiB
- Nametable mirroring: Mapper controlled
- Subject to [bus conflicts](../Bus_conflict/): No

Also see [Sachen SA8259A pinout](../Sachen_SA8259A_pinout/).

## Variants {#Variants}

There are four known variants of the board. The B variant is described below, as well as the A and C variants which alter behavior by rewiring CHR banks. The D variant behaves much differently and is described in [iNES Mapper 137](../INES_Mapper_137/).

| iNES Mapper | PCB Label | UNIF Name | CHR Shift |
|----|----|----|----|
| 138 | *unknown* | UNL-Sachen-8259B | 0 |
| 141 | TC-A003-72 | UNL-Sachen-8259A | 1 |
| 139 | *unknown* | UNL-Sachen-8259C | 2 |
| [137](../INES_Mapper_137/) (separate page) | *unknown* | UNL-Sachen-8259D | -1 (approximately) |

## Banks {#Banks}

- CPU \$8000-\$FFFF: 32KiB switchable PRG ROM bank
- PPU \$0000-\$07FF: 2KiB switchable CHR ROM bank
- PPU \$0800-\$0FFF: 2KiB switchable CHR ROM bank
- PPU \$1000-\$17FF: 2KiB switchable CHR ROM bank
- PPU \$1800-\$1FFF: 2KiB switchable CHR ROM bank

## Registers {#Registers}

Registers should be masked with \$C101.

### <span id="Reg_Select_.28.244100.29"></span>Reg Select (\$4100) {#Reg_Select_.28.244100.29}

    7  bit  0
    ---- ----
    .... .rrr
          |||
          +++- Select register number for next data write

### <span id="Reg_Data_.28.244101.29"></span>Reg Data (\$4101) {#Reg_Data_.28.244101.29}

    7  bit  0
    ---- ----
    .... .ddd
          |||
          +++- Output data to register selected by $4100.

### <span id="Chr_Select_.28Internal_0-3.29"></span>Chr Select (Internal 0-3) {#Chr_Select_.28Internal_0-3.29}

    7  bit  0
    ---- ----
    .... .ccc
          |||
          +++- Select lower 3 bits of 2KiB CHR block for PPU $0000/$0800/$1000/$1800

### <span id="Chr_Outer_Bank_.28Internal_4.29"></span>Chr Outer Bank (Internal 4) {#Chr_Outer_Bank_.28Internal_4.29}

    7  bit  0
    ---- ----
    .... .bbb
          |||
          +++- Select upper 3 bits of all four 2KiB CHR blocks

### <span id="Prg_Bank_.28Internal_5.29"></span>Prg Bank (Internal 5) {#Prg_Bank_.28Internal_5.29}

    7  bit  0
    ---- ----
    .... .ppp
          |||
          +++- Select 32KiB PRG bank for CPU $8000

### <span id="Mode_and_Mirroring_Select_.28Internal_7.29"></span>Mode and Mirroring Select (Internal 7) {#Mode_and_Mirroring_Select_.28Internal_7.29}

    7  bit  0
    ---- ----
    .... .mms
          |||
          ||+- Enable "simple" mode.
          ++-- Select mirroring (0 = V, 1 = H, 2 = (0,1,1,1), 3 = 1scA)

When simple mode is enabled, mirroring is fixed to V, ignoring the m bits, and all 4 CHR banks use the CHR bank 0 register.

Note that variant D (mapper [137](../INES_Mapper_137/)) differs because of the CHR shift, and is described on its own page.

## <span id="CHR_Rewiring_.28A_and_C_variants.29"></span>CHR Rewiring (A and C variants) {#CHR_Rewiring_.28A_and_C_variants.29}

The B variant described above has 6 bit CHR regs covering 2KiB blocks, giving a CHR ROM total of up to 128KiB. The A and C variants rewire the board to increase this limit to 256KiB and 512KiB respectively. This is done by shifting up the CHR ROM address outputs on the mapper and connecting the lower lines directly to the PPU. To be specific:

    B: Passes PPU A[10:0] through to CHR ROM A[10:0]
       Uses PPU A[12:11] to choose chr reg
       Passes 6 bits of chr reg to CHR ROM A[16:11]
    A: Passes PPU A[11:0] through to CHR ROM A[11:0]
       Uses PPU A[12:11] to choose chr reg
       Passes 6 bits of chr reg to CHR ROM A[17:12]
    C: Passes PPU A[12:0] through to CHR ROM A[12:0]
       Uses PPU A[12:11] to choose chr reg
       Passes 6 bits of chr reg to CHR ROM A[18:13]

This can be emulated by shifting the CHR regs left 1 or 2 bits, and inserting the appropriate 1 or 2 bits of PPU address in the low bits. Note that when s = 1, CHR reg 0 is used for all banks, but the extra address line(s) still leak through on the A and C boards. This means that on a C board, when s = 1, the board can be considered to have a single 8KiB CHR ROM bank for pattern tables.

## Errata {#Errata}

- The game *Q-Boy*, using a PCB labelled **2M-RAM-COB**, replaces CHR-ROM with 8 KiB of unbanked CHR-RAM, and is usually set to mapper 141. Emulators must therefore distinguish a banked CHR-ROM variant with a CHR shift of 1, and an unbanked CHR-RAM variant, of mapper 141.
- kevtris had originally assigned mapper 141 to be used just for *Q-Boy*, and mapper **135** instead of 141 for the 8259A CHR-ROM variant with a CHR shift of 1.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) NES/fami Core Release Notes for Analogue NT Mini Jailbreak firmware
