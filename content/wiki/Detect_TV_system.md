---
title: "Detect TV system"
url: "/wiki/Detect_TV_system/"
description: "A program can detect video output timing at runtime. There are three major categories: NTSC (Famicom and NES), PAL NES, and PAL famiclone (such as Dendy)."
wikiPage: "Detect_TV_system"
---

A program can detect video output timing at runtime. There are three major categories: NTSC (Famicom and NES), PAL NES, and PAL famiclone (such as Dendy).

## Using NMI {#Using_NMI}

The following code detects the TV system. A game can use it to compensate for differences in [clock rate](../Cycle_reference_chart/#Clock_rates) among various NES models. It has been tested on NTSC NES and Famicom, PAL NES, and PAL famiclones using the Dendy chipset, as well as emulators in NTSC, PAL, and Dendy modes.

    ;
    ; NES TV system detection code
    ; Copyright 2011 Damian Yerrick
    ;
    ; Copying and distribution of this file, with or without
    ; modification, are permitted in any medium without royalty provided
    ; the copyright notice and this notice are preserved in all source
    ; code copies.  This file is offered as-is, without any warranty.
    ;
    .export getTVSystem
    .importzp nmis

    .align 32  ; ensure that branches do not cross a page boundary

    ;;
    ; Detects which of NTSC, PAL, or Dendy is in use by counting cycles
    ; between NMIs.
    ;
    ; NTSC NES produces 262 scanlines, with 341/3 CPU cycles per line.
    ; PAL NES produces 312 scanlines, with 341/3.2 CPU cycles per line.
    ; Its vblank is longer than NTSC, and its CPU is slower.
    ; Dendy is a Russian famiclone distributed by Steepler that uses the
    ; PAL signal with a CPU as fast as the NTSC CPU.  Its vblank is as
    ; long as PAL's, but its NMI occurs toward the end of vblank (line
    ; 291 instead of 241) so that cycle offsets from NMI remain the same
    ; as NTSC, keeping Balloon Fight and any game using a CPU cycle-
    ; counting mapper (e.g. FDS, Konami VRC) working.
    ;
    ; nmis is a variable that the NMI handler modifies every frame.
    ; Make sure your NMI handler finishes within 1500 or so cycles (not
    ; taking the whole NMI or waiting for sprite 0) while calling this,
    ; or the result in A will be wrong.
    ;
    ; @return A: TV system (0: NTSC, 1: PAL, 2: Dendy; 3: unknown
    ;         Y: high byte of iterations used (1 iteration = 11 cycles)
    ;         X: low byte of iterations used
    .proc getTVSystem
        ldx #0
        ldy #0
        lda nmis
    nmiwait1:
        cmp nmis
        beq nmiwait1
        lda nmis

    nmiwait2:
        ; Each iteration takes 11 cycles.
        ; NTSC NES: 29780 cycles or 2707 = $A93 iterations
        ; PAL NES:  33247 cycles or 3022 = $BCE iterations
        ; Dendy:    35464 cycles or 3224 = $C98 iterations
        ; so we can divide by $100 (rounding down), subtract ten,
        ; and end up with 0=ntsc, 1=pal, 2=dendy, 3=unknown
        inx
        bne :+
        iny
    :
        cmp nmis
        beq nmiwait2
        tya
        sec
        sbc #10
        cmp #3
        bcc notAbove3
        lda #3
    notAbove3:
        rts
    .endproc

## Without using NMI {#Without_using_NMI}

Region timing can also be detected by polling \$2002 rather than using the NMI:

- [Re: Best way to detect NTSC or PAL](http://forums.nesdev.org/viewtopic.php?p=163258#p163258) forum post by lidnariq

## What to change {#What_to_change}

- PAL NES needs different period values from NTSC and Dendy to achieve the same music pitch.
- NTSC needs different numbers of frames per beat from PAL NES and Dendy to achieve the same tempo.
- PAL NES needs different wait times from NTSC and Dendy for raster effects.
- NTSC needs different velocity values from PAL and Dendy to achieve the same speed of scrolling and sprite movement.
- [OAM DMA](../PPU_registers/) timing within the vblank handler can be adjusted:
  - If you want to take advantage of PAL's longer vblank (e.g. to load more tiles into CHR-RAM), then OAM DMA must be done within the first 24 scanlines of vblank (due to [forced OAM refresh](../PPU_OAM/#Dynamic_RAM_decay)). Dendy needs different detection of the end of rendering in this scenario.
  - NTSC benefits from doing OAM DMA last if [using timed code to work around the DMC DMA glitch](../Controller_reading_code/#DPCM_Safety_using_OAM_DMA).
