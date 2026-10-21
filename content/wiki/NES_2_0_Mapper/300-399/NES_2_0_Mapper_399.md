---
title: "NES 2.0 Mapper 399"
url: "/wiki/NES_2.0_Mapper_399/"
description: "NES 2.0 Mapper 399 is used for BATMAP-000."
wikiPage: "NES_2_0_Mapper_399"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>BATMAP-000</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>BATMAP-000 </td></tr><tr><td><b>PRG ROM capacity</b></td><td>1024K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K + 8K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td></td></tr><tr><td><b>CHR capacity</b></td><td>32K </td></tr><tr><td><b>CHR window</b></td><td>4Kx2 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H or V, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><s class="broken-link" title="The linked page does not exist">399</s></td></tr></tbody></table></div>

**NES 2.0 Mapper 399** is used for **BATMAP-000**.

## Banks {#Banks}

- CPU \$8000-\$9FFF: fixed to first (#0) 8 KiB of PRG ROM
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM
- CPU \$8000-\$9FFF: fixed to last (#127) 8 KiB of PRG ROM
- PPU \$0000-\$0FFF: 4 KiB switchable CHR memory
- PPU \$1000-\$1FFF: 4 KiB switchable CHR memory

## Registers {#Registers}

### <span id="CHR_bank_select_.28.248000.2C_write.29"></span>CHR bank select (\$8000, write) {#CHR_bank_select_.28.248000.2C_write.29}

Mask: probably \$E001

    7654 3210
    ---------
    B... .CCC
    |     +++- 4 KiB CHR bank number
    +--------- 0: Set 4 KiB CHR bank at PPU $0000-$0FFF
               1: Set 4 KiB CHR bank at PPU $1000-$1FFF

### <span id="PRG_bank_select_.28.248001.2C_write.29"></span>PRG bank select (\$8001, write) {#PRG_bank_select_.28.248001.2C_write.29}

Mask: probably \$E001

    7654 3210
    ---------
    BPPP PPPP
    |+++ ++++- 8 KiB PRG-ROM bank number
    +--------- 0: Set 8 KiB PRG bank at CPU $A000-$BFFF
               1: Set 8 KiB PRG bank at CPU $C000-$DFFF

### <span id="Other_registers_.28.24A000-.24FFFF.29"></span>Other registers (\$A000-\$FFFF) {#Other_registers_.28.24A000-.24FFFF.29}

Mirroring (\$A000-\$BFFE even) and IRQ (\$C000-FFFF) are believed to be the same as [MMC3](../MMC3/). WRAM functionality is unknown.

## Hardware {#Hardware}

The attested game, *Star Versus* by Dustmop, uses 32 KiB of CHR RAM.

## References {#References}

- [Description by NewRisingSun](https://forums.nesdev.org/viewtopic.php?p=243816#p243816)
