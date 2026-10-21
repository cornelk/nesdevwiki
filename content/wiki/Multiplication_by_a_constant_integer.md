---
title: "Multiplication by a constant integer"
url: "/wiki/Multiplication_by_a_constant_integer/"
description: "In binary arithmetic, multiply by 2^n is equivalent to shifting left n times (technically this is true even if n is negative). For this reason, it is recommended that you design your NES project in a…"
wikiPage: "Multiplication_by_a_constant_integer"
wikiCategories:
  - name: "Arithmetic"
    url: "/category/Arithmetic/"
---

## Multiply by a power of two {#Multiply_by_a_power_of_two}

In binary arithmetic, multiply by 2^n is equivalent to shifting left n times (technically this is true even if n is negative). For this reason, it is recommended that you design your NES project in a way that takes advantage of this fact.

For signed numbers tough, its required you save the sign first, and restore it at the end, and this will be true for all multiplications discussed in this article, which focuses on unsigned numbers.

## Muliply by a constant {#Muliply_by_a_constant}

When you want to multiply 2 numbers, if one of them is constant, it is possible take advantage of this to optimize code (as opposed to use the general purpose [multiplication by 2 variables](../8-bit_Multiply/)).

First thing is to decompose the constant number into sum-of-power-of-two (i.e. write it in binary form). Take the variable number, shift it and add the content of the original variable after a shift whenever the corresponding constant bit is set. For an example here I'll multiply by 13 = 0b1101

    ;Multiplication by 13
         lda Var      ;'1'
         asl A
         adc Var      ;'1'
         asl A        ;'0'
         asl A
         adc Var      ;'1'

Of course the result will only be correct if smaller than 8-bit in the above example (as it relies on the ASL A instruction to clear carry between ads), if you expect a larger number of bits you should use other variables such as in this example of a multiplication by 49 = 0b110001:

    ;Multiplication by 49
         lda #$00
         sta ResH     ;Init high bits of result
         lda Var      ;'1'
         asl A
         rol ResH
         adc Var      ;'1'
         asl A
         rol ResH     ;'0'
         asl A
         rol ResH     ;'0'
         asl A
         rol ResH     ;'0'
         asl A
         rol ResH
         adc Var      ;'1'
         ldy ResH

Note that multiplications of 2^n numbers can easily be seen as a particular case of the above.

## Chain multiply by a variable {#Chain_multiply_by_a_variable}

If you want to do an algorithm that does a really great number of multiplication by a variable, but that the value of the variable is constant for the whole algorithm, it can be faster to write a code that generate the above code in RAM and execute it that way instead of doing the slower variable x variable code.
