---
title: "INES Mapper 142"
url: "/wiki/INES_Mapper_142/"
description: "iNES Mapper 142 represents a PCB used for several of Kaiser's unlicensed ports of FDS games. Its UNIF board name is UNL-KS7032, and it features the KS202 ASIC, which is an upgrade to Konami's VRC3."
wikiPage: "INES_Mapper_142"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

iNES Mapper 142 represents a PCB used for several of Kaiser's unlicensed ports of FDS games. Its UNIF board name is **UNL-KS7032**, and it features the [KS202 ASIC](../VRC3_pinout/), which is an upgrade to Konami's [VRC3](../VRC3/).

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KB switchable PRG ROM bank
- CPU \$8000-\$9FFF: 8 KB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KB PRG ROM bank, fixed to the last bank.
- PPU \$0000-\$1FFF: 8 KB CHR RAM, unbanked

## Registers {#Registers}

### <span id="IRQ_Reload_value_.28.248000-.24BFFF.29"></span>IRQ Reload value (\$8000-\$BFFF) {#IRQ_Reload_value_.28.248000-.24BFFF.29}

Writes set IRQ reload value, in four four-bit words, same as [VRC3](../VRC3/)

### <span id="IRQ_Enable_.28.24C000-.24CFFF.29"></span>IRQ Enable (\$C000-\$CFFF) {#IRQ_Enable_.28.24C000-.24CFFF.29}

Kaiser's SMB2J writes \$02 or \$05 to enable, and \$00 to disable, indicating that bits 1 and 2 may function similarly to the VRC3.

### <span id="IRQ_Acknowledge_.28.24D000-.24DFFF.29"></span>IRQ Acknowledge (\$D000-\$DFFF) {#IRQ_Acknowledge_.28.24D000-.24DFFF.29}

Any write to this register will acknowledge the pending IRQ, like VRC3. Unlike a VRC3, \$C000 bit 0 must not be moved to bit 1, otherwise SMB2J will freeze.

### <span id="Bank_select_.28.24E000-.24EFFF.29"></span>Bank select (\$E000-\$EFFF) {#Bank_select_.28.24E000-.24EFFF.29}

Select one of four bank registers to update on next write to \$F000.

    7  bit  0
    ---- ----
    .... .RRR
          |||
          +++- Specify which bank register to update on next write to Bank Data register
               1: Select 8 KB PRG bank at CPU $8000-$9FFF
               2: Select 8 KB PRG bank at CPU $A000-$BFFF
               3: Select 8 KB PRG bank at CPU $C000-$DFFF
               4: Select 8 KB PRG bank at CPU $6000-$7FFF
               6: Writes to $F000 have no effect
               0,5,7: unknown

### <span id="Bank_data_.28.24F000-.24FFFF.29"></span>Bank data (\$F000-\$FFFF) {#Bank_data_.28.24F000-.24FFFF.29}

    7  bit  0
    ---- ----
    .... DDDD
         ||||
         ++++- New bank value, based on last value written to Bank select register (mentioned above)

## See also {#See_also}

- [PCB photos](https://forums.nesdev.org/viewtopic.php?p=242685#p242685)
- [iNES Mapper 056](../INES_Mapper_056/) augments this ASIC with some additional hardware
