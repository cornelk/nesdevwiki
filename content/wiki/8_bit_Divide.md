---
title: "8-bit Divide"
url: "/wiki/8-bit_Divide/"
description: "The following code divide two 8-bit integers (range 0...255) and outputs a 16-bit result in fixed point 8.8 format. It is only using real calculations, no lookup table, thus the size of the code is…"
wikiPage: "8_bit_Divide"
wikiCategories:
  - name: "Arithmetic"
    url: "/category/Arithmetic/"
---

The following code divide two 8-bit integers (range 0...255) and outputs a 16-bit result in fixed point 8.8 format. It is only using real calculations, no lookup table, thus the size of the code is very small.

    ;8-bit divide
    ;by Bregalad
    ;Enter with A = Dividend, Y=Divisor
    ;Output with YX = (A/Y) << 8, A = remainder
    Division
        sta Dvd     ;Stores dividend
        sty Dvs     ;Stores divisor
        lda #$00    ;A must be zero before entering the loop below.
        sta ResHi   ;Clear result, setting up a ring counter
            ldy #$01        ;by initializing the result to $0001.
            sty Res         ;When the 1 gets shifted out, we're done
    -   asl Dvd
        rol A       ;Shift in 1 bit of dividend
            bcs +           ;If carry is set, the dividend is already greater
        cmp Dvs     ;Check if fractional dividend is greater than divisor
        bcc ++
    +   sbc Dvs     ;Subtract (C is always set)
        sec     ;Necessary if we reached here via the `bcs +` above
    ++  rol Res     ;Shift result (1 if subtraction was done, 0 otherwise)
        rol ResHi
        bcc -
        ldx Res
        ldy ResHi
        rts
