---
title: "INES Mapper 153"
url: "/wiki/INES_Mapper_153/"
description: "iNES Mapper 153 is used for a Bandai FCG board with an LZ93D50 ASIC and 8192 bytes of battery-backed WRAM. These 8192 bytes must be denoted as PRG-NVRAM in the NES 2.0 header using byte value $70.…"
wikiPage: "INES_Mapper_153"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

iNES Mapper 153 is used for a [Bandai FCG](../Bandai_FCG_board/) board with an [LZ93D50 ASIC](../Bandai_LZ93D50_pinout/) and 8192 bytes of battery-backed WRAM. These 8192 bytes must be denoted as PRG-NVRAM in the [NES 2.0](../NES_2.0/) header using byte value \$70. Only one game, *Famicom Jump II: Saikyou no 7-nin*, uses this mapper.

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB battery-backed WRAM
- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM

# Registers {#Registers}

## <span id="Outer_PRG-ROM_Bank_Select_.28.248000-.248003_write.29"></span>Outer PRG-ROM Bank Select (\$8000-\$8003 write) {#Outer_PRG-ROM_Bank_Select_.28.248000-.248003_write.29}

    Mask: $800F

    7  bit  0
    ---- ----
    xxxx xxxP
            |
            +-- Select 256 KiB outer PRG-ROM bank at CPU $8000-$FFFF

- The same value must be written to all four registers, or the outer PRG-ROM bank would be switched as the PPU is rendering.
- Because the ASIC's PA12 and PA13 inputs are grounded, only registers \$8000-\$8003 instead of \$8000-\$8007 are relevant.
- No CHR banking is available.

## <span id="Inner_PRG-ROM_Bank_Select_.28.248008_write.29"></span>Inner PRG-ROM Bank Select (\$8008 write) {#Inner_PRG-ROM_Bank_Select_.28.248008_write.29}

## <span id="Nametable_Mirroring_Type_Select_.28.248009_write.29"></span>Nametable Mirroring Type Select (\$8009 write) {#Nametable_Mirroring_Type_Select_.28.248009_write.29}

## <span id="IRQ_Control_.28.24800A_write.29"></span>IRQ Control (\$800A write) {#IRQ_Control_.28.24800A_write.29}

## <span id="IRQ_Latch_.28.24800B-.24800C_write.29"></span>IRQ Latch (\$800B-\$800C write) {#IRQ_Latch_.28.24800B-.24800C_write.29}

These four registers function the same way as on [INES Mapper 016](../INES_Mapper_016/), submapper 5.

## <span id="PRG-RAM_Control_.28.24800D_write.29"></span>PRG-RAM Control (\$800D write) {#PRG-RAM_Control_.28.24800D_write.29}

    Mask: $800F

    7  bit  0
    ---- ----
    ..E. ....
      |
      +-------- PRG-RAM Chip Enable (1=Enable)

# Note {#Note}

When booting with WRAM filled with zeroes, *Famicom Jump II* will freeze with a black screen. Simply soft-resetting the console will then always run the game properly unless WRAM is zeroed out again.
