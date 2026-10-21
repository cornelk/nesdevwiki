---
title: "Programming MMC3"
url: "/wiki/Programming_MMC3/"
description: "The MMC3 was Nintendo's third ASIC mapper for the NES. It features scanline IRQs, fine-grained PRG/CHR bankswitching, nametable arrangement switching, and optional PRG-RAM support. The trade-off…"
wikiPage: "Programming_MMC3"
---

|  |  |
|----|----|
| <img src="../../wiki-images/Edit-clear.svg" width="40" height="40" /> | This is a **work-in-progress** programming guide. |

The [MMC3](../MMC3/) was Nintendo's third ASIC mapper for the NES. It features scanline [IRQs](../IRQ/), fine-grained PRG/CHR bankswitching, [nametable arrangement](../Mirroring/#Nametable_Mirroring) switching, and optional PRG-RAM support. The trade-off between its feature set and complexity/cost (especially in comparison to the [MMC5](../MMC5/)) made it a popular choice for games released late into the commercial era. It is still a popular choice for modern homebrew development due to abundant documentation and after-market reproductions.

## NES 2.0 header {#NES_2.0_header}

TODO:

- Example header
- Oversize definition?

## Init code {#Init_code}

Appropriate [init code](../Init_code/) and [interrupt](../CPU_interrupts/) vectors/handlers must be located in the MMC3's fixed bank at \$E000-\$FFFF. The following should be initialized for the MMC3:

- Scanline IRQ disable (\$E000 write) and counter state (see section below)
- Nametable arrangement (\$A000 write), prior to initialising [nametables](../PPU_nametables/)
- PRG-ROM bank mode and CHR A12 inversion (\$8000 write)
  - This should either be saved to a mirror variable at runtime, or set as a constant at build-time.
- PRG-RAM protection (\$A001 write) and PRG-RAM contents (\$6000-\$7FFF)
  - Some emulators may ignore the protection state due to an [iNES mapper conflation with MMC6](../MMC3/#iNES_Mapper_004_and_MMC6). Set the [NES 2.0 submapper](../NES_2.0_submappers/#004:_MMC3) to correctly specify MMC3 behavior.
  - If the PRG-RAM is battery-backed, verifying data integrity before using it is recommended.

Afterwards, the game can initialize the bank registers and proceed with logic located outside of the fixed bank.

### Scanline counter init {#Scanline_counter_init}

Commercial [MMC3](../MMC3/) games contain code to manually clock the scanline counter in the reset handler (after [waiting for the PPU to warm up](../PPU_power_up_state/#Best_practice)) and before exiting forced blanking (rendering disabled during bulk PPU transfers) by toggling PPU A12. This prevents undesirable behaviour caused by reloading the counter on consecutive scanlines.

This snippet is present verbatim in most commercial games:

      bit $2002
      lda #$10
      tax
    @loop:
      sta $2006
      sta $2006
      eor #$10
      dex
      bne @loop

The following snippet does the minimum amount of toggles required, and uses a ring counter to reduce the code size and cycle count. A12 is set to 0, 1, 0, 1, 0.

      bit $2002
      lda #%00001010
    @loop:
      sta $2006
      sta $2006
      asl a
      bcc @loop

This unrolled snippet does the minimum amount of toggles required (similar code appears in *Super Mario Bros. 3*):

      bit $2002
      lda #$00
      sta $2006
      sta $2006
      lda #$10
      sta $2006
      sta $2006
      lda #$00
      sta $2006
      sta $2006
      lda #$10
      sta $2006
      sta $2006

## Bankswitch interface {#Bankswitch_interface}

Bankswitching on the MMC3 is accomplished by first specifying the desired destination register (and other mapper configurations) with a write to \$8000, then writing the desired data to \$8001.

Example based on the [UNROM](../Programming_UNROM/#Bankswitching) bankswitch routine, using mirror variables for thread safety:

    ; Build-time constants for PRG & CHR banking modes
    PRG_MODE = 1
    CHR_A12 = 0

    ; The mapper is write-only; need to track its state separately
    .segment "ZEROPAGE":
    current_register: .res 1
    current_bank: .res 1

    .segment "CODE"
    ; A = bank register, Y = bank number
    bankswitch_ay:
      sta current_register                    ; interrupt handlers should restore these values before returning
      sty current_bank
    bankswitch_nosave:
      ora #((CHR_A12 << 7) | (PRG_MODE << 6)) ; preserve the current mapper config (a mirror variable can also be used)
      sta $8000                               ; set register
      sty $8001                               ; write data
      rts

With this setup, switching banks works as follows:

      lda #$06 ; set register 6 (PRG)
      ldy #$0B ; to bank 11
      jsr bankswitch_ay

If you switch banks during interrupt handlers (NMI/IRQ), do not write to `current_register` or `current_bank`. Instead, use `bankswitch_nosave` and restore the bank state at the end of the handler just before pulling registers:

      lda current_register
      ldy current_bank
      jsr bankswitch_nosave

Or alternatively, if the interrupt handler is guaranteed to not change registers/banks used by the main code:

      lda current_register
      ora #((CHR_A12 << 7) | (PRG_MODE << 6)) ; preserve the current mapper config
      sta $8000 ; we only need to account for interrupts between the $8000 and $8001 writes

It may be preferable to bake the PRG/CHR mode bits into the register select values to remove the need for bitwise OR operations.

For far-call interfaces (where a PRG bank is temporarily swapped in to call specific routines), the bankswitch and restore process can be simplified by including the bank number at a known location in each bank (e.g. `.bank` in CA65), pushing it to the stack before performing the far-call, and restoring the bank once it returns. (Effectively simulating the 65c816's `JSL`/`RTL` instructions)

### Banking modes {#Banking_modes}

The MMC3 has 8 internal banking registers (labelled here as R0 through R7), and the banking modes determine how they map PRG and CHR banks.

#### PRG banking {#PRG_banking}

The PRG ROM bank mode selects between the following:

- 0 = \$8000-\$9FFF swappable using R6, \$C000-\$DFFF fixed to second-last bank. Example use case: [UxROM](../UxROM/)-style bankswitching for \$8000-\$BFFF
- 1 = \$C000-\$DFFF swappable using R6, \$8000-\$9FFF fixed to second-last bank. Example use case: bankswitching [DPCM](../APU_DMC/) samples in \$C000-\$DFFF

Or in other terms, flipping the PRG ROM bank mode swaps around the banks mapped to \$8000-\$9FFF and \$C000-\$DFFF. The \$A000-\$BFFF area is always swappable using R7.

#### CHR banking {#CHR_banking}

The CHR A12 inversion bit selects between the following:

- 0 = two 2 KiB banks at \$0000-\$0FFF controlled by R0..R1, four 1 KiB banks at \$1000-\$1FFF controlled by R2..R5
- 1 = two 2 KiB banks at \$1000-\$1FFF controlled by R0..R1, four 1 KiB banks at \$0000-\$0FFF controlled by R2..R5

The 2KiB banks may only select even numbered CHR banks. (The lowest bit is ignored.) R0 and R1 always control the 2KiB banks, and R2 through R5 always control the 1KiB banks.

Keeping in mind the pattern table mapping for safe scanline IRQ operation (explained below), this effectively controls whether the sprites or the nametables get 1KiB CHR banks respectively.

In terms of tile counts, 2KiB = 128 tiles and 1KiB = 64 tiles. Thus, the CHR banking mode is a trade-off between fine-grained banking and bulk tile updates:

- Sprites could benefit from fine-grained banking for arbitrary combinations of individual object animations, and nametables could benefit from bulk background animations without updating tiles in an [NMI thread](../NMI_thread/).
- On the other hand, nametables could benefit from fine-grained banking for arbitrary combinations of stage elements, and sprites could benefit from bulk object animations without updating an [OAMDMA](../PPU_registers/) buffer.

## Scanline IRQ {#Scanline_IRQ}

The MMC3's scanline IRQ is primarily used for synchronizing raster effects to PPU rendering, such as placing a static status bar near the bottom of the picture (which would otherwise be very inefficient to handle using traditional [sprite-0 hits](../PPU_OAM/), and tricky to implement using [DMC IRQ](../APU_DMC/#Usage_of_DMC_for_syncing_to_video)).

### Pattern table setup {#Pattern_table_setup}

In order to have the MMC3 count scanlines correctly, the [PPU pattern tables](../PPU_pattern_tables/) must be used in a specific way by configuring [PPUCTRL](../PPU_registers/):

- [Nametable](../PPU_nametables/) tiles must always come from \$0000-\$0FFF ("left table")
- [Sprite](../PPU_OAM/) tiles must always come from \$1000-\$1FFF ("right table"), **even in 8x16 mode!**
  - In 8x16 sprite mode, unused sprites should be placed at a Y coordinate of \$F0 or greater (not \$EF). This is because the PPU still fetches the pattern data on the [pre-render scanline](../PPU_rendering/#Pre-render_scanline_.28-1_or_261.29) for sprites left in secondary OAM from the previous frame, which may cause extra clocks of the scanline counter.

Deviating from this safe pattern table mapping results in inconsistent or unreliable scanline counting. E.g. *Wario's Woods* has a flickering scanline on some PPU alignments.

### Operation {#Operation}

The broad approach for handling scanline IRQs seen in commercial games is as follows (where N and M are latch values for a given scanline split):

1.  Disable IRQ generation beforehand, then CLI. (\$E000 write)
2.  After performing vblank tasks in the NMI handler, *enable rendering* and set the latch value to N (\$C000 write).
3.  Request a counter reload (\$C001 write), then enable IRQ generation (\$E001 write). The IRQ will occur at scanline N+1 of the rendered frame (i.e. excluding vblank).
4.  In the IRQ handler, disable IRQ generation (\$E000 write) to ack the IRQ. The counter has already automatically reloaded the same latch value of N at this point, so the next IRQ will occur in N+1 scanlines.
5.  If further scanline IRQs are desired, re-enable IRQ generation (\$E001 write) for the next split (the one that was just reloaded with N) and set the latch value to M (\$C000 write) for the split after that.
    - Note the lack of a manual reload in this approach. This means the latch value for a given split must be written 2 splits in advance, with the first latch in vblank being applied twice.
6.  Handle the desired raster effect.
7.  Repeat from step 4 until all desired IRQs have been completed.

Steps 5-6 are typically implemented using a [jump table](../Jump_table/) to select new latch values and handle raster effects based on a "mode" value. The NMI handler can also use this value to determine if IRQs should be set up in the first place.

An alternative approach is to always perform the "latch, manual reload, enable IRQ" sequence in step 5, which simplifies the IRQ handler by avoiding the need to schedule latch values 2 splits in advance.

#### Notes {#Notes}

- N = 0 results in version specific behavior (fire once, or fire on every scanline). Even with the [NES 2.0 submapper](../NES_2.0_submappers/#004:_MMC3) set to denote specific behavior, it should not be relied upon.
  - Thus, the earliest scanline an IRQ can be safely fired on for all revisions is scanline 2 (N = 1), which is also the shortest reliable duration between IRQs. Any raster effects requiring changes on an individual scanline basis should be done by manually [counting cycles](../Cycle_counting/) after the IRQ.
- Reloading the counter on consecutive scanlines can result in poorly understood "pathological" behavior. The aforementioned scanline counter init code is presumed to be for avoiding this behavior on reset.

## See also {#See_also}

- [MMC3](../MMC3/) technical reference
