---
title: "Delay code"
url: "/wiki/Delay_code/"
description: "Code that causes a parametrised number of cycles of delay."
wikiPage: "Delay_code"
---

Code that causes a parametrised number of cycles of delay.

Note that all branch instructions are written assuming that no page wrap occurs. If you want to ensure this condition at compile time, use the bccnw/beqnw/etc. macros that are listed at [Fixed cycle delay](../Fixed_cycle_delay/).

## Inline code {#Inline_code}

### <span id="2.E2.80.943_cycles_of_delay:_delay.3Dr.2B2.3B_0_.E2.89.A4_r_.E2.89.A4_1.2C_r.E2.8A.A2Z.2C_.CE.94r_.3D_0.29"></span>2—3 cycles of delay: delay=r+2; 0 ≤ r ≤ 1, r⊢Z, Δr = 0) {#2.E2.80.943_cycles_of_delay:_delay.3Dr.2B2.3B_0_.E2.89.A4_r_.E2.89.A4_1.2C_r.E2.8A.A2Z.2C_.CE.94r_.3D_0.29}

            bne @1
    @1:

### <span id="4.E2.80.945_cycles_of_delay:_delay.3Dr.2B4.3B_0_.E2.89.A4_r_.E2.89.A4_1.2C_.CE.94r_.3D_0.29"></span>4—5 cycles of delay: delay=r+4; 0 ≤ r ≤ 1, Δr = 0) {#4.E2.80.945_cycles_of_delay:_delay.3Dr.2B4.3B_0_.E2.89.A4_r_.E2.89.A4_1.2C_.CE.94r_.3D_0.29}

            ora #0 ; use ora=A, cpx=X, cpy=Y
            bne @1
    @1:

### <span id="4.E2.80.945_cycles_of_delay:_delay.3DX.2B4.3B_0_.E2.89.A4_X_.E2.89.A4_1.29"></span>4—5 cycles of delay: delay=X+4; 0 ≤ X ≤ 1) {#4.E2.80.945_cycles_of_delay:_delay.3DX.2B4.3B_0_.E2.89.A4_X_.E2.89.A4_1.29}

            dex
            bpl @1
    @1:

### <span id="5.E2.80.947_cycles_of_delay:_delay.3DA.2B5.3B_0_.E2.89.A4_A_.E2.89.A4_2.2C_A.E2.8A.A2Z.29"></span>5—7 cycles of delay: delay=A+5; 0 ≤ A ≤ 2, A⊢Z) {#5.E2.80.947_cycles_of_delay:_delay.3DA.2B5.3B_0_.E2.89.A4_A_.E2.89.A4_2.2C_A.E2.8A.A2Z.29}

            beq @2
            lsr
    @2:     bne @3
    @3:

### <span id="5.E2.80.947_cycles_of_delay:_delay.3Dr.2B5.3B_0_.E2.89.A4_r_.E2.89.A4_2.2C_.CE.94r_.3D_0.29"></span>5—7 cycles of delay: delay=r+5; 0 ≤ r ≤ 2, Δr = 0) {#5.E2.80.947_cycles_of_delay:_delay.3Dr.2B5.3B_0_.E2.89.A4_r_.E2.89.A4_2.2C_.CE.94r_.3D_0.29}

            cmp #1 ; use cmp=A, cpx=X, cpy=Y
            bcc @3
            bne @3
    @3:

### <span id="5.E2.80.947_cycles_of_delay:_delay.3DX.2B5.3B_0_.E2.89.A4_X_.E2.89.A4_2.29"></span>5—7 cycles of delay: delay=X+5; 0 ≤ X ≤ 2) {#5.E2.80.947_cycles_of_delay:_delay.3DX.2B5.3B_0_.E2.89.A4_X_.E2.89.A4_2.29}

            dex
            bmi @3
            bne @3
    @3:

### <span id="6.E2.80.949_cycles_of_delay:_delay.3DA.2B6.3B_0_.E2.89.A4_A_.E2.89.A4_3.2C_A.E2.8A.A2Z.29"></span>6—9 cycles of delay: delay=A+6; 0 ≤ A ≤ 3, A⊢Z) {#6.E2.80.949_cycles_of_delay:_delay.3DA.2B6.3B_0_.E2.89.A4_A_.E2.89.A4_3.2C_A.E2.8A.A2Z.29}

            beq @2
            lsr
    @2:     beq @4
            bcs @4
    @4:

### <span id="7.E2.80.9410_cycles_of_delay:_delay.3DA.2B7.3B_0_.E2.89.A4_A_.E2.89.A4_3.29"></span>7—10 cycles of delay: delay=A+7; 0 ≤ A ≤ 3) {#7.E2.80.9410_cycles_of_delay:_delay.3DA.2B7.3B_0_.E2.89.A4_A_.E2.89.A4_3.29}

            lsr
            beq @3
            bpl @3
    @3:     bcs @4
    @4:

### <span id="8.E2.80.9411_cycles_of_delay:_delay.3DX.2B8.3B_0_.E2.89.A4_X_.E2.89.A4_3.29"></span>8—11 cycles of delay: delay=X+8; 0 ≤ X ≤ 3) {#8.E2.80.9411_cycles_of_delay:_delay.3DX.2B8.3B_0_.E2.89.A4_X_.E2.89.A4_3.29}

            dex
            bmi @4
            dex
            bmi @5
    @4:     bne @5
    @5:

### <span id="9.E2.80.9414_cycles_of_delay:_delay.3DA.E2.88.92242.3B_251_.E2.89.A4_A_.E2.89.A4_255.3B_C_.3D_0.29"></span>9—14 cycles of delay: delay=A−242; 251 ≤ A ≤ 255; C = 0) {#9.E2.80.9414_cycles_of_delay:_delay.3DA.E2.88.92242.3B_251_.E2.89.A4_A_.E2.89.A4_255.3B_C_.3D_0.29}

            adc #3  ;  2 2 2 2 2  FE FF 00 01 02
            bcc @4  ;  3 3 2 2 2  FE FF 00 01 02 ;bmi works too
            lsr     ;  - - 2 2 2  -- -- 00 00 01
            beq @5  ;  - - 3 3 2  -- -- 00 00 01
    @4:     lsr     ;  2 2 - - 2  7F 7F -- -- 00
    @5:     bcs @6  ;  2 3 2 3 2  7F 7F 00 00 00
    @6:

### <span id="10.E2.80.9414_cycles_of_delay:_delay.3DX.2B10.3B_0_.E2.89.A4_X_.E2.89.A4_4.29"></span>10—14 cycles of delay: delay=X+10; 0 ≤ X ≤ 4) {#10.E2.80.9414_cycles_of_delay:_delay.3DX.2B10.3B_0_.E2.89.A4_X_.E2.89.A4_4.29}

            cpx #3
            bcc @3
            bne @3
    @3:     dex
            bmi @6
            bne @6
    @6:

### <span id="9.E2.80.9414_cycles_of_delay:_delay.3DA.2B9.3B_0_.E2.89.A4_A_.E2.89.A4_5.29"></span>9—14 cycles of delay: delay=A+9; 0 ≤ A ≤ 5) {#9.E2.80.9414_cycles_of_delay:_delay.3DA.2B9.3B_0_.E2.89.A4_A_.E2.89.A4_5.29}

            lsr
            bcs @2
    @2:     beq @5
            lsr
            bcs @6 ;beq works too
    @5:     bne @6
    @6:

### <span id="9.E2.80.9416_cycles_of_delay:_delay.3DA.2B9.3B_0_.E2.89.A4_A_.E2.89.A4_7.29"></span>9—16 cycles of delay: delay=A+9; 0 ≤ A ≤ 7) {#9.E2.80.9416_cycles_of_delay:_delay.3DA.2B9.3B_0_.E2.89.A4_A_.E2.89.A4_7.29}

            lsr
            bcs @2
    @2:     beq @6
            lsr
            beq @7
            bcc @7
    @6:     bne @7
    @7:

### <span id="11.E2.80.9419_cycles_of_delay:_delay.3DA.2B11.3B_0_.E2.89.A4_A_.E2.89.A4_8.29"></span>11—19 cycles of delay: delay=A+11; 0 ≤ A ≤ 8) {#11.E2.80.9419_cycles_of_delay:_delay.3DA.2B11.3B_0_.E2.89.A4_A_.E2.89.A4_8.29}

    ;      Cycles | A | 0  0  0  0  0  0  0  0  0  | 0  1  2  3  4  5  6  7  8
            lsr       ; 2  2  2  2  2  2  2  2  2  | 0  0  1  1  2  2  3  3  4
            bcs @3    ; 2  3  2  3  2  3  2  3  2  | 0 c0  1 c1  2 c2  3 c3  4
            adc #255  ; 2  -  2  -  2  -  2  -  2  |-1  - c0  - c1  - c2  - c3
    @3:     beq @7    ; 2  3  3  2  2  2  2  2  2  |-1 c0 c0 c1 c1 c2 c2 c3 c3
            bcc @9    ; 3  -  -  2  2  2  2  2  2  |-1  -  - c1 c1 c2 c2 c3 c3 ;bmi works too
            lsr       ; -  -  -  2  2  2  2  2  2  | -  -  - c0 c0  1  1 c1 c1
            beq @9    ; -  -  -  3  3  2  2  2  2  | -  -  - c0 c0  1  1 c1 c1
    @7:     bcc @9    ; -  2  2  -  -  3  3  2  2  | - c0 c0  -  -  1  1 c1 c1
            bne @9    ; -  2  2  -  -  -  -  3  3  | - c0 c0  -  -  -  - c1 c1
    @9:       ;Total:  11 12 13 14 15 16 17 18 19

### <span id="12.E2.80.9423_cycles_of_delay:_delay.3DA.2B12.3B_0_.E2.89.A4_A_.E2.89.A4_11.29"></span>12—23 cycles of delay: delay=A+12; 0 ≤ A ≤ 11) {#12.E2.80.9423_cycles_of_delay:_delay.3DA.2B12.3B_0_.E2.89.A4_A_.E2.89.A4_11.29}

    ;      Cycles | A | 0  0  0  0  0  0  0  0  0  0  0  0  | 0  1  2  3  4  5  6  7  8  9 10 11
            lsr       ; 2  2  2  2  2  2  2  2  2  2  2  2  | 0  0  1  1  2  2  3  3  4  4  5  5
            bcs @2    ; 2  3  2  3  2  3  2  3  2  3  2  3  | 0  0  1  1  2  2  3  3  4  4  5  5
    @2:     lsr       ; 2  2  2  2  2  2  2  2  2  2  2  2  | 0  0  0  0  1  1  1  1  2  2  2  2
            bcc @5    ; 3  3  2  2  3  3  2  2  3  3  2  2  | 0  0  0  0  1  1  1  1  2  2  2  2
            bcs @5    ; -  -  3  3  -  -  3  3  -  -  3  3  | -  -  0  0  -  -  1  1  -  -  2  2 ;bpl works too
    @5:     beq @10   ; 3  3  3  3  2  2  2  2  2  2  2  2  | 0  0  0  0  1  1  1  1  2  2  2  2
            lsr       ; -  -  -  -  2  2  2  2  2  2  2  2  | -  -  -  -  0  0  0  0  1  1  1  1
            bcs @10   ; -  -  -  -  3  3  3  3  2  2  2  2  | -  -  -  -  0  0  0  0  1  1  1  1 ;beq works too
            delay_n 5 ; -  -  -  -  -  -  -  -  5  5  5  5  | -  -  -  -  -  -  -  -  1  1  1  1
    @10:      ;Total:  12 13 14 15 16 17 18 19 20 21 22 23

For delay_n 5, anything that causes 5 cycles of delay works. Examples: `inc $00`, `nop` + `cmp $C5`

### <span id="15.E2.80.94270_cycles_of_delay:_delay.3DA.2B15.3B_0_.E2.89.A4_A_.E2.89.A4_255.29"></span>15—270 cycles of delay: delay=A+15; 0 ≤ A ≤ 255) {#15.E2.80.94270_cycles_of_delay:_delay.3DA.2B15.3B_0_.E2.89.A4_A_.E2.89.A4_255.29}

This code peels slices of 5 cycles with a SBC-BCS loop, and then executes the delay code for A=251—255. The same code will appear later as a function version (which adds 12 cycles overhead due to JSR+RTS cost).

            sec     
    @L:     sbc #5  
            bcs @L  ;  6 6 6 6 6  FB FC FD FE FF
            adc #3  ;  2 2 2 2 2  FE FF 00 01 02
            bcc @4  ;  3 3 2 2 2  FE FF 00 01 02
            lsr     ;  - - 2 2 2  -- -- 00 00 01
            beq @5  ;  - - 3 3 2  -- -- 00 00 01
    @4:     lsr     ;  2 2 - - 2  7F 7F -- -- 00
    @5:     bcs @6  ;  2 3 2 3 2  7F 7F 00 00 00
    @6:

### <span id="16.E2.80.94271_cycles_of_delay:_delay.3DA.2B16.3B_0_.E2.89.A4_A_.E2.89.A4_255.29"></span>16—271 cycles of delay: delay=A+16; 0 ≤ A ≤ 255) {#16.E2.80.94271_cycles_of_delay:_delay.3DA.2B16.3B_0_.E2.89.A4_A_.E2.89.A4_255.29}

This code peels slices of 9 cycles with a CMP-BCC-SBC-BCS loop, and then executes the delay code for A=0—8.

    @L:     cmp #9          ;2
            bcc @0          ;2 (+1)
            sbc #9          ;2
            bcs @L          ;3
    ;      Cycles | A | 5  5  5  5  5  5  5  5  5  | 0  1  2  3  4  5  6  7  8
    @0:     lsr       ; 2  2  2  2  2  2  2  2  2  | 0  0  1  1  2  2  3  3  4
            bcs @3    ; 2  3  2  3  2  3  2  3  2  | 0 c0  1 c1  2 c2  3 c3  4
            adc #255  ; 2  -  2  -  2  -  2  -  2  |-1  - c0  - c1  - c2  - c3
    @3:     beq @7    ; 2  3  3  2  2  2  2  2  2  |-1 c0 c0 c1 c1 c2 c2 c3 c3
            bcc @9    ; 3  -  -  2  2  2  2  2  2  |-1  -  - c1 c1 c2 c2 c3 c3
            lsr       ; -  -  -  2  2  2  2  2  2  | -  -  - c0 c0  1  1 c1 c1
            beq @9    ; -  -  -  3  3  2  2  2  2  | -  -  - c0 c0  1  1 c1 c1
    @7:     bcc @9    ; -  2  2  -  -  3  3  2  2  | - c0 c0  -  -  1  1 c1 c1
            bne @9    ; -  2  2  -  -  -  -  3  3  | - c0 c0  -  -  -  - c1 c1
    @9:       ;Total:  16 17 18 19 20 21 22 23 24

### <span id="5.E2.80.9465285_cycles_of_delay:_delay_.3D_256.C3.97X_.2B_5"></span>5—65285 cycles of delay: delay = 256×X + 5 {#5.E2.80.9465285_cycles_of_delay:_delay_.3D_256.C3.97X_.2B_5}

Clobbers A:

    @0:     txa       ;2
            beq @10   ;3
            nop       ;2
            tya       ;2
             ldy #48  ;2
    @l:      dey      ;2×48
             bne @l   ;3×48
            tay       ;2−1
            dex       ;2
            jmp @0    ;3
    @10:

Doesn’t clobber A (2 bytes longer):

    @0:     cpx #0    ;2
            beq @10   ;3
            pha       ;3
            tya       ;2
             ldy #47  ;2
    @l:      dey      ;2×47
             bne @l   ;3×47
            tay       ;2−1
            pla       ;4
            jmp @0    ;3
    @10:

### <span id="18.E2.80.94218103813_cycles_of_delay:_delay_.3D_13.C3.97.2865536.C3.97Y_.2B_256.C3.97A_.2B_X.29_.2B_18"></span>18—218103813 cycles of delay: delay = 13×(65536×Y + 256×A + X) + 18 {#18.E2.80.94218103813_cycles_of_delay:_delay_.3D_13.C3.97.2865536.C3.97Y_.2B_256.C3.97A_.2B_X.29_.2B_18}

            iny
    @l1:    nop
            nop
    @l2:    cpx #1
            dex
            sbc #0
            bcs @l1
            dey
            bne @l2
            rts

## Callable functions {#Callable_functions}

### <span id="A_.2B_25_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V"></span>A + 25 cycles of delay, clobbers A, Z&N, C, V {#A_.2B_25_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V}

This code peels slices of 7 cycles with a CMP-BCS-SBC loop, and then executes the delay code for 9—16 cycles with A = 0—6. The reason its overhead is smaller than in the version that peels 5 cycles is because the case for A\<7 executes only two instructions instead of three. This comes at the cost that the entry point is not the first instruction. Therefore the code can only exist as a callable function and not inline code.

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays A clocks + overhead
    ; Clobbers A. Preserves X,Y.
    ; Time: A+25 clocks (including JSR)
    ;;;;;;;;;;;;;;;;;;;;;;;;
                      ;       Cycles              Accumulator         Carry flag
                      ; 0  1  2  3  4  5  6          (hex)           0 1 2 3 4 5 6
                      ;
                      ; 6  6  6  6  6  6  6   00 01 02 03 04 05 06
    :      sbc #7     ; carry set by CMP
    delay_a_25_clocks:
           cmp #7     ; 2  2  2  2  2  2  2   00 01 02 03 04 05 06   0 0 0 0 0 0 0
           bcs :-     ; 2  2  2  2  2  2  2   00 01 02 03 04 05 06   0 0 0 0 0 0 0
           lsr        ; 2  2  2  2  2  2  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
           bcs *+2    ; 2  3  2  3  2  3  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
           beq :+     ; 3  3  2  2  2  2  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
           lsr        ;       2  2  2  2  2         00 00 01 01 01       1 1 0 0 1
           beq @rts   ;       3  3  2  2  2         00 00 01 01 01       1 1 0 0 1
           bcc @rts   ;             3  3  2               01 01 01           0 0 1
    :      bne @rts   ; 2  2              3   00 00             01   0 1         0
    @rts:  rts        ; 6  6  6  6  6  6  6   00 00 00 00 01 01 01   0 1 1 1 0 0 1
    ; Total cycles:    25 26 27 28 29 30 31

### <span id="A_.2B_27_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V"></span>A + 27 cycles of delay, clobbers A, Z&N, C, V {#A_.2B_27_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V}

This function has longer overhead than delay_a_25_clocks, but it can be appended into other functions, as the execution begins from the first instruction.

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays A clocks + overhead
    ; Clobbers A. Preserves X,Y.
    ; Time: A+27 clocks (including JSR)
    ;;;;;;;;;;;;;;;;;;;;;;;;
    delay_a_27_clocks:
            sec     
    @L:     sbc #5  
            bcs @L  ;  6 6 6 6 6  FB FC FD FE FF
            adc #3  ;  2 2 2 2 2  FE FF 00 01 02
            bcc @4  ;  3 3 2 2 2  FE FF 00 01 02
            lsr     ;  - - 2 2 2  -- -- 00 00 01
            beq @5  ;  - - 3 3 2  -- -- 00 00 01
    @4:     lsr     ;  2 2 - - 2  7F 7F -- -- 00
    @5:     bcs @6  ;  2 3 2 3 2  7F 7F 00 00 00
    @6:     rts     ;

It is created by wrapping the code for 15—270 cycles of delay into a function. The JSR+RTS instructions adds 12 cycles of overhead.

### <span id="256.C3.97A_.2B_X_.2B_33_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V"></span>256×A + X + 33 cycles of delay, clobbers A, Z&N, C, V {#256.C3.97A_.2B_X_.2B_33_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V}

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays A:X clocks+overhead
    ; Time: 256*A+X+33 clocks (including JSR)
    ; Clobbers A. Preserves X,Y. Has relocations.
    ;;;;;;;;;;;;;;;;;;;;;;;;
    :   ; 5 cycles done, do 256-5 more.
        sbc #1          ; 2 cycles - Carry was set from cmp
        pha                     ; 3
         lda #(256-27 - 16)     ; 2
         jsr delay_a_27_clocks  ; 240
        pla                     ; 4
    delay_256a_x_33_clocks:
        cmp #1          ; +2
        bcs :-          ; +3 
        ; 0-255 cycles remain, overhead = 4
        txa             ; -1+2; 6; +27 = 33
        ;passthru
    <<Place the function delay_a_27_clocks immediately following here>>

Can be trivially changed to swap X, Y.

If you can clobber Y, change the part that begins with "pha" and ends with "pla" into this, for 1 byte shorter code:

       ldy #49  ; 2
    @l:     dey      ; 49*2
            bne @l   ; 49*3
            ldy $A4  ; 3-1

### <span id="256.C3.97A_.2B_16_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V"></span>256×A + 16 cycles of delay, clobbers A, Z&N, C, V {#256.C3.97A_.2B_16_cycles_of_delay.2C_clobbers_A.2C_Z.26N.2C_C.2C_V}

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays A*256 clocks + overhead
    ; Clobbers A. Preserves X,Y.
    ; Time: A*256+16 clocks (including JSR)
    ;;;;;;;;;;;;;;;;;;;;;;;;
    delay_256a_16_clocks:
        cmp #0
        bne delay_256a_11_clocks_
        rts
    delay_256a_11_clocks_:
        ;5 cycles done. Must consume 256 cycles; 251 cycles remain.
            pha                      ;3 - 3
            tya                      ;2 - 5
             ldy #46                 ;2 - 7
    @l:      dey                     ;2*46 - 99
             bne @l                  ;3*46 - 237
             nop                     ;2 - 238
            tay                      ;2 - 240
            pla                      ;4 - 244
        sec                      ;2 - 246
        sbc #1                   ;2 - 248
        jmp delay_256a_16_clocks ;3 - 251

If you can clobber Y, change the part that begins with pha and ends in pla, into this, for shorter code:

       ldy #48  ; 2
    @l:     dey      ; 49*2
            bne @l   ; 49*3
            ldy $A4  ; 3-1

### <span id="256.C3.97X_.2B_16_cycles_of_delay.2C_clobbers_X.2C_Z.26N"></span>256×X + 16 cycles of delay, clobbers X, Z&N {#256.C3.97X_.2B_16_cycles_of_delay.2C_clobbers_X.2C_Z.26N}

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays X*256 clocks + overhead
    ; Clobbers X,Y. Preserves A. Relocatable.
    ; Time: X*256+16 clocks (including JSR)
    ;;;;;;;;;;;;;;;;;;;;;;;;
    delay_256x_16_clocks:
        cpx #0
        bne delay_256x_11_clocks_
        rts
    delay_256x_11_clocks_:
        ;5 cycles done. Must consume 256 cycles; 251 cycles remain.
            pha                      ;3
            tya                      ;2
             ldy #46                 ;2
    @l:      dey                     ;2*46
             bne @l                  ;3*46
             nop                     ;2-1
             nop                     ;2
            tay                      ;2
            pla                      ;4
        dex                      ;2
        jmp delay_256x_16_clocks ;3

Can be trivially changed to swap X, Y.

If you can clobber Y, change the part that begins with pha and ends in pla, into this, for shorter code:

            ldy #50                  ;2-1
    @l: dey                      ;2*50
        bne @l                   ;3*50

### <span id="256.C3.97X_.2B_A_.2B_30_cycles_of_delay.2C_clobbers_A.2C_X.2C_Z.26N.2C_C.2C_V"></span>256×X + A + 30 cycles of delay, clobbers A, X, Z&N, C, V {#256.C3.97X_.2B_A_.2B_30_cycles_of_delay.2C_clobbers_A.2C_X.2C_Z.26N.2C_C.2C_V}

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays X*256+A clocks + overhead
    ; Clobbers A,X. Preserves Y.
    ; Depends on delay_a_25_clocks within short branch distance
    ; Time: X*256+A+30 clocks (including JSR)
    ;;;;;;;;;;;;;;;;;;;;;;;;
    delay_256x_a_30_clocks:
            cpx #0                  ;2
            beq delay_a_25_clocks   ;3
            ;4 cycles done. Must consume 256 cycles; 252 cycles remain.
            pha                             ;3
             lda #(256-4-(3+2+4+2+3))-25    ;2
             jsr delay_a_25_clocks          ;238
            pla                             ;4
            dex                             ;2
            jmp delay_256x_a_30_clocks      ;3

Can be trivially changed to swap X, Y.

Alternative version that does not depend on other delay functions, but has otherwise the same implications:

    :      sbc #7    ; carry set by CMP
    delay_256x_a_30_clocks_b:
           cmp #7    ; 2  2  2  2  2  2  2   00 01 02 03 04 05 06   0 0 0 0 0 0 0
           bcs :-    ; 2  2  2  2  2  2  2   00 01 02 03 04 05 06   0 0 0 0 0 0 0
           lsr       ; 2  2  2  2  2  2  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
           bcs @2    ; 2  3  2  3  2  3  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
    @2:    beq @6    ; 3  3  2  2  2  2  2   00 00 01 01 02 02 03   0 1 0 1 0 1 0
           lsr       ;       2  2  2  2  2         00 00 01 01 01       1 1 0 0 1
           beq @do_x ;       3  3  2  2  2         00 00 01 01 01       1 1 0 0 1
           bcc @do_x ;             3  3  2               01 01 01           0 0 1
    @6:    bne @do_x ; 2  2              3   00 00             01   0 1         0
    @do_x: txa       ;2
           beq @rts  ;3
           ;4 cycles done. Must consume 256 cycles; 252 cycles remain.
           nop       ;2
           tya       ;2
            ldy #48  ;2
    @l:     dey      ;2*48
            bne @l   ;3*48
           tay       ;2-1
           dex       ;2
           jmp @do_x ;3
    @rts:  rts

This function is constructed by concatenating delay_a_25_clocks and the inline delay code for 5—65285 cycles.

### <span id="851968.C3.97Y_.2B_3328.C3.97A_.2B_13.C3.97X_.2B_30_cycles_of_delay.2C_clobbers_A.2C_X.2C_Y.2C_Z.26N.2C_C.2C_V"></span>851968×Y + 3328×A + 13×X + 30 cycles of delay, clobbers A, X, Y, Z&N, C, V {#851968.C3.97Y_.2B_3328.C3.97A_.2B_13.C3.97X_.2B_30_cycles_of_delay.2C_clobbers_A.2C_X.2C_Y.2C_Z.26N.2C_C.2C_V}

    ;;;;;;;;;;;;;;;;;;;;;;;;
    ; Delays 30+13*(65536*Y+256*A+X) cycles including JSR.
    ; Clobbers A,X,Y.
    delay_851968y_3328a_13x_30_clocks:
            iny
    @l1:    nop
            nop
    @l2:    cpx #1
            dex
            sbc #0
            bcs @l1
            dey
            bne @l2
            rts

This is constructed by wrapping the 18—218103813 cycles inline delay code in a function.

## See also {#See_also}

- [Cycle counting](../Cycle_counting/)
- [Fixed cycle delay](../Fixed_cycle_delay/)
