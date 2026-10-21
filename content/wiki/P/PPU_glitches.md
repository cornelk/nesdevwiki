---
title: "PPU glitches"
url: "/wiki/PPU_glitches/"
description: "When the CPU writes a value, it signals that a write is happening before driving data onto the bus, leaving a brief window at the start of the write where the bus contents are open bus.…"
wikiPage: "PPU_glitches"
---

# Early Writes {#Early_Writes}

When the CPU writes a value, it signals that a write is happening before driving data onto the bus, leaving a brief window at the start of the write where the bus contents are [open bus](../Open_bus_behavior/). Unfortunately, the PPU assumes the data is valid for the entire write, and for many of its registers the PPU will briefly use this open bus value. This can cause visible glitches.

In most cases, the early write value seen by the PPU is the high byte of the PPU register address, because the CPU normally fetches the high byte of the address from the instruction operand on the cycle before doing the write. Therefore, the early write value is normally \$20.

## Mitigations {#Mitigations}

Early write bugs are usually inconsequential because most PPU writes occur outside of rendering, preventing them from causing visible artifacts. Even writes that occur during rendering may only matter during specific vulnerable dots. However, for writes that may be affected, there are mitigations that work by ensuring the open bus value at the start of the write matches the value being written.

Note that the early write issue mostly matters in the case where the old and new value of a register are the same. When they are the same, an early write can cause a brief window where the value is different. However, when the two values are different, early writes only change the timing of the value transition, making it either slightly earlier or later depending on whether the early write matches the old value or new value. This means, though, that bits within the write may change at different times because different bits may match one or the other.

### PPU open bus {#PPU_open_bus}

The most general solution is to prime the bus using [PPU open bus](../Open_bus_behavior/#PPU_open_bus). When a write-only PPU register is read, it will return the last value present on the PPU's internal bus, which is the last value that the CPU read from or wrote to it. This value eventually decays over long time scales of at least 1 ms, but is reliable on shorter timescales. By putting the intended value into PPU open bus and then reading this on the cycle before the write, the CPU open bus value at the start of the write will match the actual write value, preventing the PPU from briefly seeing a different value.

A PPU open bus read can be inserted between the operand read and value write by using an indexed write. These 5-cycle instructions perform a read from the target address on the 4th cycle before writing to it on the 5th cycle, in order to handle page crossings that may result in the address being off by 1 page on the 4th cycle. When writing to a write-only PPU register, this 4th cycle read will read PPU open bus. For example, to avoid the early write issue when writing to [PPUMASK](../PPU_registers/):

    LDA ppu_mask_value
    LDX #$00
    STA PPUSTATUS  ; Set the PPU open bus value by writing to a read-only register.
    STA PPUMASK,X  ; Read PPU open bus from the target register before writing it.

or more generally:

    ; A = your value
    LDX #$00
    STA PPUSTATUS
    STA write_only_ppu_register,X  ; Any of PPUCTRL, PPUMASK, OAMADDR, PPUSCROLL, PPUADDR

Note that this should not be done when writing to [PPUDATA](../PPU_registers/), which does not have early write issues, because this register is both readable and writable, has side effects on read, and takes longer than 1 CPU cycle to handle CPU accesses.

### Register mirrors {#Register_mirrors}

Because the value on the bus is normally the high byte of the PPU register address, the bus can also be primed by writing to a PPU register mirror that matches the value being written. Because mirrors only exist in the range \$2000-\$3FFF, this only mitigates the issue for the low 5 bits; the upper 3 bits will always be %001. However, if those bits are unaffected or inconsequential, this approach can be sufficient. For example, to safely enable rendering mid-screen:

    LDA #$1E   ; Fully enables rendering
    STA $3E01  ; Mirror of PPUMASK: (PPUMASK | ($1E << 8)) 

For code where the value is not fixed, all 8 bits matter, or certain register bits must change at the same time, the PPU open bus approach is likely superior.

## Hardware explanation {#Hardware_explanation}

The 6502, and thus also the 2A03, guarantee that R/W and address bus are stable while φ2 (or M2) are high, but do **not** guarantee the data bus is stable.

Here is a timing diagram for the 2A03G:

     (10ns) 0  40  80 120 160 200 240 280 320 360 400 440 480 520 560 600 640
         M2 \____________________/¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\______
    /ROMSEL __/¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\__________________________________/¯¯¯¯ (when relevant)
    /PPUSEL ____/¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\__________________________________/¯¯ (when relevant)
        R/W ¯¯¯¯¯\_________________________________________________________ (read to write cycle)
         D0 ======ZZZZZZZZZZZZZZZZZZZZZZZZZZZ=============================Z
        R/W ____________/¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯ (write to read cycle)

(TODO: Add relative timing of address bus, if different from R/W)

## Affected registers {#Affected_registers}

### [PPUCTRL](../PPU_registers/) {#PPUCTRL}

In the 2C02G, the \$C3 bits are processed in an asynchronous manner, and this can cause various problems:

- \$01s bit: On every active scanline, a write to PPUCTRL at the exact wrong time ([write starting on dot 257](../PPU_scrolling/#At_dot_257_of_each_scanline), ending on dot 258) can cause the left nametable to be drawn for the upcoming scanline. [\[1\]](https://forums.nesdev.org/viewtopic.php?t=10104) This is because the contents of [open bus](../Open_bus_behavior/) - \$20 - are used by the PPU on dot 257, setting the "base nametable" to the left one instead of the intended one. As a work-around, write to the PPU address mirror where the bottom 2 bits of the upper byte of the address match the data that will be written. [\[2\]](https://forums.nesdev.org/viewtopic.php?p=230450#p230450)
- \$02s bit: On every active field, a write to PPUCTRL at the exact wrong time ([starting on prerender scanline dot 304](../PPU_scrolling/#During_dots_280_to_304_of_the_pre-render_scanline_.28end_of_vblank.29), ending on dot 305) can cause the top nametable to be drawn for the upcoming field. Workaround is same as above.
- \$40s bit: Any write to PPUCTRL during the active field can temporarily disable "output colors on EXT pins" for one pixel. This is believed to be the cause of certain bugs reported in the HDNES.
- \$80s bit: Any write to PPUCTRL during the vertical blanking interval can cause the NMI output to be asserted, or deasserted, for about 80ns. However, this glitch is invisible, because the 6502 ignores its NMI input during this time.

In the 2C02A, it's known that the \$18 bits in PPUMASK are also processed in an asynchronous manner, and suspected that all the other bits do also:

- \$04s bit: On the 2C02A, it's believed that this will have no effect, because the write to PPUCTRL would have to occur right in the middle of incrementing the PPU address while rendering is disabled.
- \$08s bit: On the 2C02A, it's believed that a write during horizontal blanking could cause exactly one bitplane of one sliver of one sprite to be fetched from the wrong pattern table.
- \$10s bit: On the 2C02A, it's believed that a write during active redraw could cause exactly one bitplane of one sliver of one background tile to be fetched from the wrong pattern table.
- \$20s bit: On the 2C02A, it's believed that a write at any time could cause one sprite sliver could be drawn incorrectly, specifics are unclear.

### [PPUMASK](../PPU_registers/) {#PPUMASK}

On the 2C02G, the \$81 bits are processed in an asynchronous manner and this can cause unimportant glitches:

- \$01s bit: Any write to PPUMASK can, at any time, turn off the "monochrome" flag for one pixel. [\[3\]](https://forums.nesdev.org/viewtopic.php?p=256431#p256431)
- \$80s bit: Any write to PPUMASK can, regardless of subpixel phase, turn off "blue emphasis" for **half** of one pixel. [\[4\]](https://forums.nesdev.org/viewtopic.php?p=256737#p256737)

In the 2C02A and 2C03B, it's known that the \$38 bits in PPUMASK are also processed in an asynchronous manner, and suspected that all the other bits do also:

- \$18 bits: On the 2C02A, it's known that any write to PPUMASK during active redraw can turn off rendering for one pixel, causing all the documented hazards with disabling rendering.
- \$06 bits: on the 2C02A, the effects of this will be invisible under the above disabling.
- \$20 bit: on the 2C03B, it's known that any write to PPUMASK can turn on "red" emphasis off for one whole pixel (a whole pixel due to how the RGB PPU works)
- \$60 bits: on the 2C02A and 2C02B, it's believed to be able to turn on red emphasis and green emphasis off for one half pixel

### [OAMADDR](../PPU_registers/) {#OAMADDR}

On the 2C02G, changes to the OAM address can cause a copy from the first address' row of memory to the second address'. Writing OAMADDR can trigger this in multiple ways, one of which is through early writes. Unfortunately, while mitigating the early write issue can make OAMADDR writes work reliably on some CPU/PPU alignments, it is not sufficient to make them work correctly in all CPU/PPU alignments.

### <span id="PPUADDR.2C_PPUSCROLL"></span>[PPUADDR](../PPU_registers/), [PPUSCROLL](../PPU_registers/) {#PPUADDR.2C_PPUSCROLL}

On the 2C02G, any write starting on dot 257 and ending on dot 258 that updates coarse X in "t" can cause the same symptoms as writes to PPUCTRL. [\[5\]](https://forums.nesdev.org/viewtopic.php?p=230503#p230503)

### [PPUADDR](../PPU_registers/) {#PPUADDR}

Any second write to PPUADDR will immediately change update the lower three bits of coarse X and five bits of coarse Y to the value of open bus, and will then (on the next pixel, except as covered by the dot 257/258 glitch mentioned above) write the correct value. This is usually invisible but can cause an incorrect sliver.

### [PPUSCROLL](../PPU_registers/) {#PPUSCROLL}

Any first write to PPUSCROLL will immediately change fine X to the value of open bus, and will then (on the next pixel) correct fine X.

# PPU-internal bus conflicts {#PPU-internal_bus_conflicts}

Any time the PPU both increments v and reloads v from t on the same dot, a bus conflict occurs that sets both v and t to the bitwise AND of the two v inputs. This bus conflict affects just the components that were being incremented (X and/or Y).

This can happen (at least) two different ways:

- A [second write of PPUADDR](../PPU_scrolling/#.242006_second_write_.28w_is_1.29) (loading a new scroll location) so that the t to v copy occurs at the same time that the [Y bits are incremented](../PPU_scrolling/#At_dot_256_of_each_scanline) (dot 256) or the [coarse X bits are incremented](../PPU_scrolling/#Between_dot_328_of_a_scanline.2C_and_256_of_the_next_scanline) (the last dot of every set of 4 background memory fetches). [\[6\]](https://forums.nesdev.org/viewtopic.php?t=18092)
- A [read or write of PPUDATA](../PPU_scrolling/#.242007_reads_and_writes) so that the resulting fine Y and coarse X increments occur at the same time the PPU naturally reloads v from t (dot 257 on each rendering scanline, or dot 304 of the prerender scanline).
