---
title: "INES Mapper 091"
url: "/wiki/INES_Mapper_091/"
description: "iNES Mapper 091 was used by the original release of Super Fighter III from an unknown publisher, but is more commonly known by a clone board from J.Y. Company."
wikiPage: "INES_Mapper_091"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with fixed-timing scanline IRQs"
    url: "/category/Mappers_with_fixed-timing_scanline_IRQs/"
---

iNES Mapper 091 was used by the original release of *Super Fighter III* from an unknown publisher, but is more commonly known by a clone board from J.Y. Company.

<table class="wikitable">
<thead>
<tr>
<th colspan="3">INES Mapper 091 submapper table</th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Board</th>
<th>Games</th>
</tr>
&#10;<tr>
<td>0</td>
<td>JY830623C</td>
<td><em>Street Fighter 3</em>, <em>Mortal Kombat II</em>, <em>Dragon Ball Z 2</em>, <em>Mario &amp; Sonic 2</em></td>
</tr>
<tr>
<td>0</td>
<td>YY840238C</td>
<td><em>1995 Super HIK 4-in-1 (JY-016)</em>, <em>1995 Super HiK 4-in-1 (JY-017)</em></td>
</tr>
<tr>
<td>1</td>
<td>EJ-006-1</td>
<td><em>Super Fighter III</em></td>
</tr>
</tbody>
</table>

Submapper 0 has a PA12-based IRQ counter and hard-wired mirroring; submapper 1 an M2-based IRQ counter and selectable mirroring.

# Banks {#Banks}

- CPU \$8000-\$9FFF: 8 KiB switchable PRG-ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG-ROM bank, hard-wired to last bank
- PPU \$0000-\$07FF: 2 KiB switchable CHR-ROM bank
- PPU \$0800-\$0FFF: 2 KiB switchable CHR-ROM bank
- PPU \$1000-\$17FF: 2 KiB switchable CHR-ROM bank
- PPU \$1800-\$1FFF: 2 KiB switchable CHR-ROM bank

# Registers {#Registers}

## <span id="CHR-ROM_Bank_Select_.28.246000-.246003.2C_write.29"></span>CHR-ROM Bank Select (\$6000-\$6003, write) {#CHR-ROM_Bank_Select_.28.246000-.246003.2C_write.29}

Mask: \$F003/\$F007 (submapper 0/1)

- \$6000: Select 2 KiB CHR-ROM bank at PPU \$0000-\$07FF
- \$6001: Select 2 KiB CHR-ROM bank at PPU \$0800-\$0FFF
- \$6002: Select 2 KiB CHR-ROM bank at PPU \$1000-\$17FF
- \$6003: Select 2 KiB CHR-ROM bank at PPU \$1800-\$1FFF

## <span id="Nametable_Mirroring_Select_.28.246004-.246005.2C_write.2C_submapper_1_only.29"></span>Nametable Mirroring Select (\$6004-\$6005, write, submapper 1 only) {#Nametable_Mirroring_Select_.28.246004-.246005.2C_write.2C_submapper_1_only.29}

Mask: \$F007

- \$6004: Select horizontal mirroring, value ignored
- \$6005: Select vertical mirroring, value ignored

## <span id="IRQ_Counter_Register_.28.246006-.246007.2C_write.2C_submapper_1_only.29"></span>IRQ Counter Register (\$6006-\$6007, write, submapper 1 only) {#IRQ_Counter_Register_.28.246006-.246007.2C_write.2C_submapper_1_only.29}

Mask: \$F007

- \$6006: Set IRQ Counter LSB
- \$6007: Set IRQ Counter MSB

## <span id="PRG-ROM_Bank_Select_.28.247000-.247001.2C_write.29"></span>PRG-ROM Bank Select (\$7000-\$7001, write) {#PRG-ROM_Bank_Select_.28.247000-.247001.2C_write.29}

Mask: \$F003/\$F007 (submapper 0/1)

- \$7000: Select 8 KiB PRG-ROM bank at CPU \$8000-\$9FFF
- \$7001: Select 8 KiB PRG-ROM bank at CPU \$A000-\$BFFF

## <span id="IRQ_Stop.2FAcknowledge_.28.247006.2C_write.29"></span>IRQ Stop/Acknowledge (\$7006, write) {#IRQ_Stop.2FAcknowledge_.28.247006.2C_write.29}

Mask: \$F003 (submapper 0)/\$F007 (submapper 1)

Value is ignored.

## <span id="IRQ_Start.2FReset_.28.247007.2C_write.29"></span>IRQ Start/Reset (\$7007, write) {#IRQ_Start.2FReset_.28.247007.2C_write.29}

Mask: \$F003/\$F007 (submapper 0/1)

Value is ignored.

## <span id="Outer_Bank_register_.28.248000-.249FFF.2C_write.2C_submapper_0_only.29"></span>Outer Bank register (\$8000-\$9FFF, write, submapper 0 only) {#Outer_Bank_register_.28.248000-.249FFF.2C_write.2C_submapper_0_only.29}

Mask: unknown

    A~FEDC BA98 7654 3210
      -------------------
      .... .... .... .PPC
                      ||+- Select outer 512 KiB CHR-ROM bank (CHR A19)
                      ++-- Select outer 128 KiB PRG-ROM bank (PRG A17-A18)

# Notes {#Notes}

- Submapper 0: the IRQ counter is a binary counter clocked by PPU A12, and simply counts a fixed number of 64 unfiltered rises of PPU A12.
- Submapper 1: the IRQ counter is a binary counter clocked by the M2 signal with a factor of 5/4, meaning counting down by five every fourth M2 cycle.
- Submapper 1's *Super Fighter III* runs a hardware test when resetting with Select+Start+B pressed. The test exists in J.Y. Company's hacked versions as well, but was not adjusted for the simpler hardware and thus will fail some tests.
- The commonly available ROM image of *Super Fighter III* is a mapper hack for [INES Mapper 197](../INES_Mapper_197/).
