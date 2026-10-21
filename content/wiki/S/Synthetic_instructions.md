---
title: "Synthetic instructions"
url: "/wiki/Synthetic_instructions/"
description: "There are several additional instructions that would be nice to have on the NES. Even though not present, they can be synthesized using existing instructions, resulting in synthetic instructions or…"
wikiPage: "Synthetic_instructions"
---

There are several additional instructions that would be nice to have on the NES. Even though not present, they can be synthesized using existing instructions, resulting in **synthetic instructions** or **pseudoinstructions**. If turned into assembler macros, they can be used almost as if they were natively supported. Being able to think of them as native instructions lightens the mental load when programming, because instructions are an important tool for abstraction. Even without making the following into macros, after reading them you will be more likely to think of one of these while coding, saying "I need a subtract-from instruction here".

## Negate A {#Negate_A}

Many processors have a native negate instruction, which subtracts the value from zero. Here we must manually calculate the two's complement of A, which involves a one's complement and increment:

    ; A = -A
    eor #$FF
    sec
    adc #0

Or if we know carry is clear coming in:

    ; A = -A
    eor #$FF
    adc #1

## Reverse subtraction {#Reverse_subtraction}

Using SBC, a value can be subtracted from A, but there's no direct way to subtract A from some value. The ARM instruction set includes instructions RSB (reverse subtract) and RSC (reverse subtract with carry) that negate the register and add the value. The 6502 can do this too:

    ; A = Value - A
    eor #$FF
    sec
    adc Value

As a special case, if we want to subtract A from 255, we can just do

    ; A = 255 - A
    eor #$FF

This also shows another way of understanding the general subtract from; we first subtract A from 255, then add one, so it's as if we subtracted from 256, which is the same as subtracting from zero, since A is only 8 bits.

## Sign-extend {#Sign-extend}

When increasing the number of bits in a signed value, the new high bits are filled with copies of the sign bit. This is called sign extension. For example, sign-extending the 8-bit value \$80 (-128) to 16 bits sets the new bits, resulting in \$FF80; sign-extending \$7F (+127) to 16 bits results in \$007F. The following sequences calculate the upper byte of the sign-extended value.

            ; calculates upper byte of sign-extension of A
            ora #$7F
            bmi neg
            lda #0
    neg:

            ; calculates upper byte of sign-extension of A, alternate version
            and #$80
            bpl pos
            lda #$ff
    pos:

This constant-time version (7 bytes, 8 cycles) destroys the carry, so don't try using it in the middle of a multi-byte addition:

      asl a     ;  cpx #$80 or cpy #$80 is also possible
      lda #$00
      adc #$FF  ; C is unchanged and A = $00 if C was set or $FF if C was clear
      eor #$FF  ; now all bits of A are set to what bit 7 was

If you're just trying to add an 8-bit delta to a 16-bit value, you could try subtracting 256 from the value by decrementing the high byte if the value is negative and then adding as if it were unsigned.

      lda delta_value
      bpl notneg
      dec value_hi
    notneg:
      clc
      adc value_lo
      sta value_lo
      lda #0
      adc value_hi
      sta value_hi

[This forum thread](http://forums.nesdev.org/viewtopic.php?p=124949#p124949) describes addition of signed deltas with clamping.

## Arithmetic shift right {#Arithmetic_shift_right}

The ARM instruction set has an arithmetic right shift, which doesn't alter the sign (top) bit. This shift is used to divide a signed value by two. But the 6502 lacks this instruction; LSR doesn't work because it shifts the sign bit to the right, then clears it.

To implement this, we need carry set to the sign bit, then we can use ROR. CMP \#\$80 performs this task; if the value is less than \$80, carry is cleared, otherwise it's set:

    ; Arithmetic shift right A
    cmp #$80
    ror a

If the operand is in memory, we just use ASL to move the sign bit into carry:

    ; Arithmetic shift right Value
    lda Value
    asl a
    ror Value

## 8-bit rotate {#8-bit_rotate}

The 65xx series rotate instructions are all 9-bit, not 8-bit as often imagined. If they really were 8-bit, then eight ROR or ROL instructions in a row would leave A with its original value. In actuality, *nine* are required to do so, since the carry acts as a ninth bit of A.

Similar to arithmetic right shift, we must set carry to the top or bottom bit in advance of the rotate. For 8-bit rotate left, it's simple:

    ; 8-bit rotate left A
    cmp #$80
    rol a

    ; alternate method
    asl a
    adc #0

For 8-bit rotate right, we must save and restore A:

    ; 8-bit rotate right A
    pha
    lsr a
    pla
    ror a

A could be saved and restored using other methods, like TAX and TXA, etc.

Branching can also be used:

        ; 8-bit rotate right A
        lsr a
        bcc skip
        adc #$80-1 ; carry is set, so will add extra 1
    skip:

If the operand is in memory:

    ; 8-bit rotate left Value
    lda Value
    asl a
    rol Value

    ; 8-bit rotate right Value
    lda Value
    lsr a
    ror Value

To rotate A left twice:<sup><span id="cite_ref-wilson_1-0"></span>[\[1\]](#cite_note-wilson-1)</sup>

    ; double 8-bit rotate left A
    asl a
    adc #$80
    rol a

## Nybble swap {#Nybble_swap}

The nybbles of A can be swapped (e.g. \$1F becomes \$F1) in 8 bytes and 12 cycles by doing the "double 8-bit rotate left A" trick above twice:<sup><span id="cite_ref-wilson_1-1"></span>[\[1\]](#cite_note-wilson-1)</sup>

    ; swap nybbles of A
    asl a
    adc #$80
    rol a
    asl a
    adc #$80
    rol a

## 16-bit increment and decrement {#16-bit_increment_and_decrement}

Incrementing/decrementing a 16-bit value involves first adjusting the low byte, then adjusting the high byte if necessary. Increment is simpler, since the high byte is adjusted when the low byte wraps around to zero; for decrement, the high byte is adjusted when the low byte wraps around to \$FF.

            ; 16-bit increment Word
            inc Word
            bne noinc
            inc Word+1
    noinc:

            ; 16-bit decrement Word
            lda Word
            bne nodec
            dec Word+1
    nodec:  dec Word

16-bit increment shows even more advantage when used to control a loop, because the 16-bit increment conveniently leaves the zero flag set at the end *only* if the entire 16-bit value is zero.

## <span id="X.2FY_as_Operand"></span>X/Y as Operand {#X.2FY_as_Operand}

*Main article: [Identity table](../Identity_table/)*

Normally X and Y cannot be used as an operand to an instruction operating on A. For example, CMP X isn't possible. Where X or Y needs to be used in such a way, they are usually saved to a temporary variable:

    ; Compare A with X
    stx temp
    cmp temp

By putting a 256-byte table in memory with each entry simply having the value of its index, X and Y can be used as operands:

    table:  .byte $00,$01,$02,$03,$04,$05,$06,$07,$08,$09,$0A,$0B,$0C,$0D,$0E,$0F
            .byte $10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$1A,$1B,$1C,$1D,$1E,$1F
            ...
            .byte $F0,$F1,$F2,$F3,$F4,$F5,$F6,$F7,$F8,$F9,$FA,$FB,$FC,$FD,$FE,$FF

    cmp table,x      ; CMP X

    eor table,x      ; EOR X

    clc
    adc table,y      ; ADC Y

## <span id="JMP_.28addr.2CX.29"></span>JMP (addr,X) {#JMP_.28addr.2CX.29}

The JMP (addr,X) instruction is present in later 65xx processors. It behaves like JMP (addr), except it fetches the 16-bit value from addr+X. The least-problematic way to implement this is using RTI:

    ; Jump to address stored at addr+X
    lda addr+1,x
    pha
    lda addr,x
    pha
    php
    rti

See [Jump table](../Jump_table/) for further explanation and alternate approaches.

## Toggle carry {#Toggle_carry}

Invert the carry flag without affecting A (N and Z flags are destroyed, however):<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

    rol a
    eor #$01
    ror a

## Set overflow flag {#Set_overflow_flag}

The overflow flag can be set by using the BIT instruction with any byte that has the second-most significant bit set. The N and Z flags are destroyed, however.<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

       bit label
       ...
    label:
       rts        ; opcode $60 (%01100000)

## Test whether A is in range {#Test_whether_A_is_in_range}

Test whether A (unsigned) is between constants *min* and *max* (inclusive). A is destroyed.<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)</sup>

Set carry flag if A is in range, otherwise clear carry:

    clc
    adc #$ff-max
    adc #max-min+1

Clear carry flag if A is in range, otherwise set carry:

    sec
    sbc #min
    sbc #max-min+1

## Count bits set {#Count_bits_set}

Count the number of bits set in *myVal*. The result will be in X. *myVal* is destroyed. (On each round of the loop, the least significant set bit is cleared.)<sup><span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup>

       ldx #0
       lda myVal
       beq end
    loop:
       inx
       dec myVal
       and myVal
       sta myVal
       bne loop
    end:

Count the number of bits set in A. The result will be in X and A. (On each round of the loop, the most significant bit is shifted out.)<sup><span id="cite_ref-6"></span>[\[6\]](#cite_note-6)</sup>

       ldx #$ff   ; Set count to -1
    incr:
       inx        ; Add one to count
    loop:
       asl        ; Shift by one bit
       bcs incr   ; Count one bits
       bne loop   ; Repeat till zero
       txa        ; Move count to A

## Set negative and zero flags without affecting carry {#Set_negative_and_zero_flags_without_affecting_carry}

Make the N and Z flags reflect the value of A:

    and #$ff  ; either this...

    ora #$00  ; ...or this...

    eor #$00  ; ...or this

(`and #$ff` is used immediately before `bpl` in *Excitebike*.)

Make the N and Z flags reflect the value of X:

    inx
    dex
    ; ...or vice versa

Make the N and Z flags reflect the value of Y:

    iny
    dey
    ; ...or vice versa

## See also {#See_also}

- [6502 assembly optimisations](../6502_assembly_optimisations/)

## External links {#External_links}

- [Pseudoinstructions in MIPS](http://www.cs.umd.edu/class/spring2003/cmsc311/Notes/Mips/pseudo.html)

## References {#References}

1.  <span id="cite_note-wilson-1"></span>↑ <sup>[1.0](#cite_ref-wilson_1-0)</sup> <sup>[1.1](#cite_ref-wilson_1-1)</sup> <http://www.6502.org/source/general/SWN.html>
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [EOR \#\$FF - 6502 Ponderables and Befuddlements](https://archive.org/details/eor6502), puzzle \$08
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) [EOR \#\$FF - 6502 Ponderables and Befuddlements](https://archive.org/details/eor6502), puzzle \$09
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) [EOR \#\$FF - 6502 Ponderables and Befuddlements](https://archive.org/details/eor6502), puzzle \$20
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) [EOR \#\$FF - 6502 Ponderables and Befuddlements](https://archive.org/details/eor6502), puzzle \$29
6.  <span id="cite_note-6"></span>[↑](#cite_ref-6) [6502.org Forum - Bit Counting](http://forum.6502.org/viewtopic.php?t=1206)
