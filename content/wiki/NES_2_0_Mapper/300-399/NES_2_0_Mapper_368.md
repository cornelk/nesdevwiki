---
title: "NES 2.0 Mapper 368"
url: "/wiki/NES_2.0_Mapper_368/"
description: "NES 2.0 Mapper 368 denotes the Yung-08 cartridge conversion PCB of the disk system game Super Mario Bros. 2 (Japan). It is the \"parent\" of the N-32 and Mr. Mary 2 conversions, as both retain…"
wikiPage: "NES_2_0_Mapper_368"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with fixed-timing cycle IRQs"
    url: "/category/Mappers_with_fixed-timing_cycle_IRQs/"
---

[<img src="../../wiki-images/YUNG8-front.jpg" width="300" height="225" />](../../wiki-images/YUNG8-front.jpg)

[](../../wiki-images/YUNG8-front.jpg)

[PCB front](../../wiki-images/YUNG8-front.jpg)

[](../../wiki-images/YUNG8-front.jpg)

[](../../wiki-images/YUNG8-front.jpg)

[](../../wiki-images/YUNG8-front.jpg)[<img src="../../wiki-images/YUNG8-back.jpg" width="300" height="400" />](../../wiki-images/YUNG8-back.jpg)

[](../../wiki-images/YUNG8-back.jpg)

[PCB back (mirrored, pin 31 on the left)](../../wiki-images/YUNG8-back.jpg)

[](../../wiki-images/YUNG8-back.jpg)

[NES 2.0 Mapper 368 denotes the Yung-08 cartridge conversion PCB of the disk system game *Super Mario Bros. 2 (Japan)*. It is the "parent" of the](../../wiki-images/YUNG8-back.jpg) [N-32](../INES_Mapper_050/) and [Mr. Mary 2](../NES_2.0_Mapper_357/) conversions, as both retain (jumped-over) [protection-reading code](#Latch_read_.28.244122.2C_read.29) from this version.

There are two 32 KiB PRG-ROM chips, one 2 KiB PRG-ROM, and one 8 KiB CHR-ROM chip. The ROM image's 128 KiB PRG-ROM part is arranged as follows:

1.  Offset \$00000: First PRG-ROM chip (32 KiB)
2.  Offset \$08000: Second PRG-ROM chip (32 KiB)
3.  Offset \$10000: Third PRG-ROM chip, repeated 32 times (64 KiB)

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB PRG-ROM bank, fixed to \#2
- CPU \$8000-\$9FFF: 8 KiB PRG-ROM bank, fixed to \#1
- CPU \$A000-\$BFFF: 8 KiB PRG-ROM bank, fixed to \#0
- CPU \$C000-\$DFFF: 8 KiB PRG-ROM bank, switchable
- CPU \$E000-\$FFFF: 8 KiB PRG-ROM bank, fixed to \#8 (2 KiB PRG-ROM repeated four times)
- PPU \$0000-\$1FFF: unbanked 8 KiB CHR-ROM

# Registers {#Registers}

## <span id="PRG_Bank_Select_.28.244022.2C_write.29"></span>PRG Bank Select (\$4022, write) {#PRG_Bank_Select_.28.244022.2C_write.29}

    Mask: $F1FF

    Bit 7654 3210
        ---------
        .... .CCC
              +++- Select 8 KiB PRG-ROM bank at CPU $C000-$DFFF

The actual bank number is:

    Value  Bank#
    ------------
    0      4
    1      3
    2      5
    3      3
    4      6
    5      3
    6      7
    7      3

## <span id="Latch_write.2FIRQ_Control_.28.244122.2C_write.29"></span>Latch write/IRQ Control (\$4122, write) {#Latch_write.2FIRQ_Control_.28.244122.2C_write.29}

    Mask: $F1FF

    D~7654 3210
      ---------
      ..LL .L.I
        ||  | +- 0: Acknowledge and disable IRQ, reset counter
        ||  | |  1: Enable IRQ
        ++--+-+- Bits to be read-back

When enabled, the 12-bit IRQ counter increases on every M2 cycle until it overflows, upon which an IRQ is fired.

## <span id="Latch_read_.28.244122.2C_read.29"></span>Latch read (\$4122, read) {#Latch_read_.28.244122.2C_read.29}

    Mask: $F1FF

    D~7654 3210
      ---------
      10LL 1L1L
        ++--+-+- Latched bits

When holding the A button while pressing START to begin in World A-1, the game checks the latch's functionality and returns to the title screen if the check fails.
