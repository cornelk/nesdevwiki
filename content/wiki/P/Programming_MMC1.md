---
title: "Programming MMC1"
url: "/wiki/Programming_MMC1/"
description: "MMC1 was Nintendo's first ASIC mapper for the NES."
wikiPage: "Programming_MMC1"
---

[MMC1](../MMC1/) was Nintendo's first ASIC mapper for the NES.

## Quick setup for UNROM style {#Quick_setup_for_UNROM_style}

If you are using the [SGROM](../MMC1/) or [SNROM](../MMC1/) board to provide an environment similar to [UNROM](../Programming_UNROM/), with 8 KB of CHR RAM, a fixed PRG ROM bank at \$C000, and a 16 KB switchable PRG ROM bank at \$8000, do this in your init code after the mapper has been reset:

      lda #$0E   ; vertical mirroring, fixed $C000, 8 KB CHR pages
      sta $8000  ; (use $0F instead for horizontal mirroring)
      lsr a
      sta $8000
      lsr a
      sta $8000
      lsr a
      sta $8000
      lsr a
      sta $8000

Games that use CHR RAM switch to another PRG bank before they [copy tile data into CHR RAM](../CHR_ROM_vs._CHR_RAM/#Switching_to_CHR_RAM).

## PRG banks {#PRG_banks}

Some revisions of the MMC1 IC might power up in a mode other than fixed-\$C000, requiring that the vectors and the start of the [init code](../Init_code/) be placed in all banks, much as in [BxROM](../INES_Mapper_034/) or [AxROM](../AxROM/) or [GxROM](../GxROM/). Other revisions guarantee that the fixed bank is loaded at power on. To make sure your code works on all MMC1 revisions, put the following code in the last 16 bytes of each 16384 byte bank. (*Barbie* uses almost identical code.)

    reset_stub:
      sei
      ldx #$FF
      txs        ; set the stack pointer
      stx $8000  ; reset the mapper
      jmp reset  ; must be in $C000-$FFED
      .addr nmiHandler, reset_stub, irqHandler

Then to switch PRG ROM banks, load the bank number (0-15) into A and call this subroutine:

    mmc1_load_prg_bank:
      sta $E000
      lsr a
      sta $E000
      lsr a
      sta $E000
      lsr a
      sta $E000
      lsr a
      sta $E000
      rts

## Interrupts {#Interrupts}

If an NMI or IRQ can interrupt a series of writes, it is not easy to know what state the serial register was in before the interruption.

One technique for coping with this problem involves using a flag variable to indicate that a serial write should be retried<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>:

1.  The normal serial write routine should have an "interrupted" flag:
    - Clear the flag before beginning the series of writes.
    - When completed, check the flag. If an interrupt is indicated, reset the serial register, clear the interrupted flag and begin the serial writes again.
2.  If the NMI or IRQ needs to switch banks, it should reset the serial register and set the "interrupted" flag to indicate that it has done this.

Another technique involves saving the bank number and resetting the mapper before every bankswitch to guarantee the correct bank is switched in after an interrupt. However, this forces the mode to be fixed-\$C000<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>:

1.  The normal serial write routine should save the current bank number and reset the mapper before writing to the shift register.
2.  The interrupt handler should do one of the following:
    - A. Save the current bank number on entry. Restore the current bank number, bankswitch, then reset the mapper before exiting.
    - B. If interrupts are guaranteed to never be nested, it can bankswitch without saving the bank number (but still resetting the mapper). Load the current bank number and bankswitch before exiting.

## See also {#See_also}

- [MMC1](../MMC1/) technical reference

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [forum post](https://forums.nesdev.org/viewtopic.php?p=132577#p132577): thefox describes MMC1 interrupt technique used in STREEMERZ
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [forum post](https://forums.nesdev.org/viewtopic.php?t=25500): foobles' thread-safe MMC1 bankswap without checks

## External links {#External_links}

- [SNROM project template](https://github.com/pinobatch/snrom-template) by Damian Yerrick on GitHub
