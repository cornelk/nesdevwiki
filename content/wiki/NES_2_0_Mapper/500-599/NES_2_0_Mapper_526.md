---
title: "NES 2.0 Mapper 526"
url: "/wiki/NES_2.0_Mapper_526/"
description: "NES 2.0 Mapper 526 is used for a bootleg version of Namco's 三国志: 中原の覇者 (Sangokushi: Chūgen no Hasha). Its UNIF board name is UNL-BJ-56. Mirroring seems to be hard-wired (to vertical)."
wikiPage: "NES_2_0_Mapper_526"
wikiCategories:
  - name: "Mappers with fixed-timing cycle IRQs"
    url: "/category/Mappers_with_fixed-timing_cycle_IRQs/"
---

NES 2.0 Mapper 526 is used for a bootleg version of Namco's 三国志: 中原の覇者 (*Sangokushi: Chūgen no Hasha*). Its [UNIF](../UNIF/) board name is **UNL-BJ-56**. Mirroring seems to be hard-wired (to vertical).

# Registers {#Registers}

## <span id="CHR-ROM_bank_select_.28.248000-.248007.29"></span>CHR-ROM bank select (\$8000-\$8007) {#CHR-ROM_bank_select_.28.248000-.248007.29}

    Mask: unknown, probably $800F
    $8000: Set 1 KiB CHR-ROM bank at PPU $0000-$03FF
    $8001: Set 1 KiB CHR-ROM bank at PPU $0400-$07FF
    $8002: Set 1 KiB CHR-ROM bank at PPU $0800-$0BFF
    $8003: Set 1 KiB CHR-ROM bank at PPU $0C00-$0FFF
    $8004: Set 1 KiB CHR-ROM bank at PPU $1000-$13FF
    $8005: Set 1 KiB CHR-ROM bank at PPU $1400-$17FF
    $8006: Set 1 KiB CHR-ROM bank at PPU $1800-$1BFF
    $8007: Set 1 KiB CHR-ROM bank at PPU $1C00-$1FFF

## <span id="PRG-ROM_bank_select_.28.248008-.24800B.29"></span>PRG-ROM bank select (\$8008-\$800B) {#PRG-ROM_bank_select_.28.248008-.24800B.29}

    Mask: unknown, probably $800F
    $8008: Set 8 KiB PRG-ROM bank at CPU $8000-$9FFF
    $8009: Set 8 KiB PRG-ROM bank at CPU $A000-$BFFF
    $800A: Set 8 KiB PRG-ROM bank at CPU $C000-$DFFF
    $800B: Set 8 KiB PRG-ROM bank at CPU $E000-$FFFF

## <span id="IRQ_counter_.28.24800D.2F.24800F.29"></span>IRQ counter (\$800D/\$800F) {#IRQ_counter_.28.24800D.2F.24800F.29}

Mask: unknown, probably \$800F. One of these two acknowledges a pending IRQ, and the other resets to IRQ counter to zero. Because they are always written to one after the other, it's not clear which one does which.

# IRQ Counter Operation {#IRQ_Counter_Operation}

The IRQ counter is increased on every M2 cycle. The IRQ line is asserted when bit 12 is set and released when bit 12 is cleared. In other words, an IRQ is generated after 4096 M2 cycles and self-acknowledged after further 4096 M2 cycles.
