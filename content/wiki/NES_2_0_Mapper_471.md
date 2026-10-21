---
title: "NES 2.0 Mapper 471"
url: "/wiki/NES_2.0_Mapper_471/"
description: "NES 2.0 Mapper 471 denotes the Impact Soft IM1 circuit board, used for Haratyler (without HG or MP) and Haraforce. It is basically INES Mapper 201 with the addition of a scanline IRQ."
wikiPage: "NES_2_0_Mapper_471"
wikiCategories:
  - name: "Mappers with fixed-timing scanline IRQs"
    url: "/category/Mappers_with_fixed-timing_scanline_IRQs/"
---

**NES 2.0 Mapper 471** denotes the **Impact Soft IM1** circuit board, used for *Haratyler* (without *HG* or *MP*) and *Haraforce*. It is basically [INES Mapper 201](../INES_Mapper_201/) with the addition of a scanline IRQ.

# Banks {#Banks}

- CPU \$8000-\$FFFF: 32 KiB switchable PRG ROM bank
- PPU \$0000-\$1FFF: 8 KiB switchable CHR ROM bank
- Nametable mirroring: hard-wired, selectable via solder pad

# <span id="Address_Latch.2FAcknowledge_IRQ_.28.248000-.24FFFF.2C_write.29"></span>Address Latch/Acknowledge IRQ (\$8000-\$FFFF, write) {#Address_Latch.2FAcknowledge_IRQ_.28.248000-.24FFFF.2C_write.29}

    A~[1... .... .... ...B]
                         +- Select PRG A15 and CHR A13

Writing to this latch acknowledges a pending IRQ.

# Scanline IRQ {#Scanline_IRQ}

On every falling edge of PPU A12, an IRQ is raised that, unless interrupts are disabled in the CPU flag register, must be acknowledged by writing to the address latch. No filtering is performed; it is assumed that the interrupt handler will take long enough to execute so that it gets to be executed only once per scanline.
