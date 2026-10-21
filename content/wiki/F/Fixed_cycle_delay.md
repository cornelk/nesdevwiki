---
title: "Fixed cycle delay"
url: "/wiki/Fixed_cycle_delay/"
description: "Shortest possible CPU code that creates N cycles of delay, depending on constraints."
wikiPage: "Fixed_cycle_delay"
---

Shortest possible CPU code that creates N cycles of delay, depending on constraints.

## Code {#Code}

All code samples are written for CA65.

Assumptions:

- No page wrap occurs during any branch instruction. If a page wrap occurs, it adds +1 cycle for each loop, completely thwarting the accurate delay.
- No interrupt / NMI occurs during the delay code.

It is permissible for DMA to steal cycles during the loops. If you are expecting that to happen, you have to manually adjust the delay cycle count (and it is in fact possible to do so) in order to get the correct delay.

### Explanations on the requirements {#Explanations_on_the_requirements}

- @rts12 means you know a memory address that contains byte \$60 (`RTS`).

cycle instruction that fits your constraints (such as `LDA $00`), followed by `RTS`.

### <span id="Instructions.2C_addressing_modes.2C_byte_counts.2C_cycle_counts_and_notes"></span>Instructions, addressing modes, byte counts, cycle counts and notes {#Instructions.2C_addressing_modes.2C_byte_counts.2C_cycle_counts_and_notes}

| Addressing mode | Instruction type | Bytes | Cycle count | Example instruction | Notes |
|----|----|----|----|----|----|
| Implied | Inter-register | 1 | 2 | `TAX` | `NOP` has no side effects. Flag-manipulations like `CLC`, and `SEC``CLV` are used when their effects are desired. |
| Implied | Stack push | 1 | 3 | `PHA` | `PHP` is only paired with `PLP`. |
| Implied | Stack pop | 1 | 4 | `PLA` |  |
| Implied | Return | 1 | 6 | `RTS` | Used indirectly when paired with `JSR`. Similarly for `RTI`. |
| Immediate |  | 2 | 2 | `CMP #$C5` | Includes instructions like `LDA`, `LDX` and `LDY`. Other ALU instructions are used in more complex situations. |
| Relative | Branch | 2 | 2—4 | `BCC *+2` | Branch takes 3 cycles when taken, 2 otherwise. A page crossing adds +1 cycle when branch is taken, but because of difficulties setting that up, we don't use it. |
| Zeropage | Read, write | 2 | 3 | `LDA $A5` |  |
| Zeropage | RMW | 2 | 5 | `INC @zptemp` | Writing to zeropage is only permitted when @zptemp is available. Technically we could save @zptemp into register and restore at end, but it is bytewise inferior to other techniques. |
| Zeropage indexed | Read, write | 2 | 4 | `LDA $EA,X` | Inferior to 2 × `NOP`, but useful for hiding additional code to be executed in a loop. |
| Zeropage indexed | RMW | 2 | 6 | `INC @zptemp,X` | Only doable when X is known to be 0, or when entire zeropage can be clobbered. |
| Indexed indirect | Read, write | 2 | 6 | `STA (@ptrtemp,X)` | Only doable when X is known to be 0. |
| Indexed indirect | RMW | 2 | 8 | `SLO (@ptrtemp,X)` | The most cost-effective instruction. Only doable when X is known to be 0, lest we write to a random address. All instructions in this category are unofficial. |
| Indirect indexed | Read | 2 | 5—6 | `LDA (@ptrtemp),Y` | Never used by this code. |
| Indirect indexed | Write | 2 | 6 | `STA (@ptrtemp),Y` | Only doable when Y is known to be 0. |
| Indirect indexed | RMW | 2 | 8 | `SLO (@ptrtemp),Y` | All instructions in this category are unofficial. |
| Absolute | Jump | 3 | 3 | `JMP *+3` |  |
| Absolute | Read, write | 3 | 4 | `LDA $2808` | Inferior to 2 × `NOP`, but can be used carefully to hide additional code to be executed in a loop. |
| Absolute | RMW | 3 | 6 | `INC $4018` | Inferior to 3 × `NOP`. |
| Absolute indexed | Read | 3 | 4—5 | `LDA $0200,X` | Inferior to shorter alternatives. |
| Absolute indexed | Write | 3 | 5 | `STA $0200,X` | Inferior to shorter alternatives. |
| Absolute indexed | RMW | 3 | 7 | `INC $4018,X` | Only doable when writing into the given address is harmless considering the possible values of X. |
| Absolute indirect | Jump | 3 | 5 | `JMP (@ptrtemp)` | Inferior to shorter alternatives. |

{{#css:

     .testtable td{padding:2px} .testtable td pre{padding:2px;margin:2px}

}}

### 2 cycles {#2_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">1 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA       NOP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- All instructions cost at least 2 cycles. There is no way to do 1 cycle of delay (though *−1 cycles* may sometimes appear in branch cost calculations).

### 3 cycles {#3_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">2 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>C5 C5    CMP $C5</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>24 24    BIT $24</code></pre></td>
<td>Clobbers Z&amp;N, and V</td>
</tr>
<tr>
<td><pre><code>A5 A5    LDA $A5</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A4 A4    LDY $A4</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">3 bytes</td>
</tr>
<tr>
<td><pre><code>4C xx xx JMP *+3</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- Not relocatable means that the target address is hardcoded into the code. In ROM hacking, it sometimes makes sense to move code blobs around, and a hardcoded address makes it difficult to relocate the code. This restriction does not apply to branches, which use relative addressing. It is also assumed to not apply to `JSR` instructions, as chances are the JSR target is outside the code being relocated.

### 4 cycles {#4_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">2 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 2</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- zp-indexed modes such as `LDA $00,X` also do 4 cycles, but having side effects, these two-byte instructions are inferior to a simple 2 × `NOP`.
- There is also an unofficial opcode `NOP $00,X` (34 00), but there is no reason to use this instruction when the official equivalent has the same performance.

### 5 cycles {#5_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">3 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
90 00    BCC *+2</code></pre></td>
<td>Clobbers C</td>
</tr>
<tr>
<td><pre><code>B8       CLV
50 00    BVC *+2</code></pre></td>
<td>Clobbers V</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A5 A5    LDA $A5</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A6 A6    LDX $A6</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A4 A4    LDY $A4</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">4 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
4C xx xx JMP *+3</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- abs-indexed modes such as `LDA $1234,X` cause 4 or 5 cycles of delay, depending whether a page wrap occurred. Because you need extra setup code to make sure that a wrap does occur, you do not see this mode in these samples, outside situations where circumstances permit.

### 6 cycles {#6_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">3 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 3</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- zp-indexed RMW instructions such as `INC @zptemp,X` do 6 cycles. Unless we know the value of X, it might write into any address between \$00-\$FF. This option is only useful if the entire range of \$00-\$FF is free for clobbering with random data, or if X has a known value.
- ix instructions like `LDA ($00,X)` do 6 cycles, but the value of X decides where a pointer is read from, and said pointer can point anywhere. We only do that when the value of X is known.
- iy instructions like `LDA ($00),Y` also do 5-6 cycles, but in addition to the note above, we cannot predict whether a wrap occurs or not. So we don't use this mode.
- Absolute RMW instructions like `INC $4018` do 6 cycles, but weighing 3 bytes with side-effects it would be inferior to 3 × `NOP`.

### 7 cycles {#7_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">2 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- `PHP-PLP` is very efficient for 7 cycles of delay, but it does modify stack contents. S register remains unchanged though.
- `PLA-PHA` does not overwrite any bytes in stack. It just writes back the same byte. But it does clobber A and Z+N. It is not interrupt-unsafe either: If an interrupt happens, the stack byte does get temporarily clobbered, but the value is still in A when the interrupt exits, and gets written back in stack.
- abs-indexed RMW instructions such as `INC abs,X` do 7 cycles. We only do this when either we know the value of X (for instance, `INC $4018,X` is safe when X is 0—7, or when the entire 256-byte page can be safely overwritten with random data.

### 8 cycles {#8_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 4</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- Unofficial ix and iy RMW instructions such as `SLO ($00,X)` or `SLO ($00),Y` would do 8 cycles for 2 bytes of code. We only do that if we know X or Y to be zero, and we have a known pointer to safely rewritable data.

### 9 cycles {#9_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">3 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA       NOP
08       PHP
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- Jumping into the middle of another instruction and thereby reusing code is a very efficient way of reducing code size. Note that all code samples using branches on this page require that no page wrap occurs.

### 10 cycles {#10_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP
C5 C5    CMP $C5
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- The `ROL-ROR` sequence preserves the original value of the memory address. Carry is also preserved.

### 11 cycles {#11_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 12 cycles {#12_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">3 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>20 xx xx JSR @rts12</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">4 bytes</td>
</tr>
<tr>
<td><pre><code>36 36    ROL $36,X
76 36    ROR $36,X</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td colspan="2">5 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
90 00    BCC *+2
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

- `JSR-RTS` causes 12 cycles of delay. But it does write a function return address in the stack, which may be unwanted in some applications. S is not modified.
- Again, `ROL-ROR` does not have side effects (as long as an interrupt does not happen in the middle), except for Z+N.

### 13 cycles {#13_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 3
08       PHP
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 14 cycles {#14_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP       \ × 2
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 15 cycles {#15_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP
BA       TSX
28       PLP
9A       TXS
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>C5 C5    CMP $C5
20 xx xx JSR @rts12</code></pre></td>
<td>Clobbers Z&amp;N, and C; and requires @rts12</td>
</tr>
<tr>
<td><pre><code>24 24    BIT $24
20 xx xx JSR @rts12</code></pre></td>
<td>Clobbers Z&amp;N, and V; and requires @rts12</td>
</tr>
<tr>
<td><pre><code>A5 A5    LDA $A5
20 xx xx JSR @rts12</code></pre></td>
<td>Clobbers A, and Z&amp;N; and requires @rts12</td>
</tr>
<tr>
<td><pre><code>A4 A4    LDY $A4
20 xx xx JSR @rts12</code></pre></td>
<td>Clobbers Y, and Z&amp;N; and requires @rts12</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
28       PLP
EA   ... NOP      × 4</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 16 cycles {#16_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA       NOP
08       PHP       \ × 2
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 17 cycles {#17_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP
48       PHA
A5 A5    LDA $A5
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 18 cycles {#18_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP       \ × 2
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 19 cycles {#19_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>08       PHP
28       PLP
20 xx xx JSR @rts12</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
36 36    ROL $36,X
76 36    ROR $36,X
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 20 cycles {#20_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
08       PHP       \ × 2
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 21 cycles {#21_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 04    LDY #4
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP       \ × 3
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 22 cycles {#22_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 02    LDX #2
EA       NOP
CA       DEX
10 FC    BPL *-2</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 03    LDY #3
EA       NOP
88       DEY
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
BA       TSX
08       PHP
28   ... PLP      × 2
9A       TXS
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
C5 C5    CMP $C5
28       PLP
20 xx xx JSR @rts12</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP       \ × 2
28       PLP       /
EA   ... NOP      × 4</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 23 cycles {#23_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18   ... CLC      × 2
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP       \ × 3
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 24 cycles {#24_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>20 xx xx JSR @rts12× 2</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A4 A4    LDY $A4
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
C5 C5    CMP $C5
28       PLP       \ × 2
08       PHP       /
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 25 cycles {#25_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">7 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>98       TYA
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP       \ × 3
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 26 cycles {#26_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 04    LDX #4
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 05    LDY #5
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
20 xx xx JSR @rts12× 2</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
36 36    ROL $36,X
76 36    ROR $36,X
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 27 cycles {#27_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A5 A5    LDA $A5
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>24 2C    BIT &lt;$2C ;hides &#39;BIT $FDA2&#39;
A2 FD    LDX #253
E8       INX
D0 FA    BNE *-4</code></pre></td>
<td>Clobbers X, Z&amp;N, and V</td>
</tr>
<tr>
<td><pre><code>24 2C    BIT &lt;$2C ;hides &#39;BIT $FDA0&#39;
A0 FD    LDY #253
C8       INY
D0 FA    BNE *-4</code></pre></td>
<td>Clobbers Y, Z&amp;N, and V</td>
</tr>
<tr>
<td><pre><code>A4 AC    LDY &lt;$AC ;hides &#39;LDY $82A2&#39;
A2 82    LDX #130
CA       DEX
30 FA    BMI *-4</code></pre></td>
<td>Clobbers X, Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>24 24    BIT $24
20 xx xx JSR @rts12× 2</code></pre></td>
<td>Clobbers Z&amp;N, and V; and requires @rts12</td>
</tr>
<tr>
<td><pre><code>20 xx xx JSR @rts12
08       PHP
BA       TSX
28       PLP
9A       TXS
28       PLP</code></pre></td>
<td>Clobbers X; and requires @rts12</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
08       PHP       \ × 3
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 28 cycles {#28_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>38   ... SEC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A2 04    LDX #4
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP       \ × 4
28       PLP       /</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 29 cycles {#29_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA       NOP
90 FC    BCC *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 04    LDX #4
EA       NOP
CA       DEX
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 04    LDY #4
EA       NOP
88       DEY
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 02    LDX #2
EA       NOP
CA       DEX
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 03    LDY #3
EA       NOP
88       DEY
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
28       PLP
08       PHP
C5 C5    CMP $C5
28       PLP
20 xx xx JSR @rts12</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
C5 C5    CMP $C5
28       PLP
08       PHP
36 36    ROL $36,X
76 36    ROR $36,X
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 30 cycles {#30_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">7 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A2 04    LDX #4
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18   ... CLC      × 2
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18   ... CLC      × 2
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 6A    LDA #$6A ;hides &#39;ROR A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 31 cycles {#31_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 05    LDX #5
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 06    LDY #6
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
28       PLP
20 xx xx JSR @rts12× 2</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
A6 A6    LDX $A6
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A4 A4    LDY $A4
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
36 36    ROL $36,X \ × 2
76 36    ROR $36,X /
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 32 cycles {#32_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A2 05    LDX #5 ;hides &#39;ORA zp&#39;
CA       DEX ;first loop only
CA       DEX
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 05    LDY #5 ;hides &#39;ORA zp&#39;
88       DEY ;first loop only
88       DEY
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA   ... NOP      × 3
10 FA    BPL *-4</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 04    LDX #4
EA       NOP
CA       DEX
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A4 A4    LDY $A4
A0 04    LDY #4
EA       NOP
88       DEY
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
98       TYA
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
98       TYA
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
A8       TAY
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A2 04    LDX #4
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18   ... CLC      × 2
A9 6A    LDA #$6A ;hides &#39;ROR A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 33 cycles {#33_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>18   ... CLC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 04    LDX #4
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
28       PLP
20 xx xx JSR @rts12× 2</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP       \ × 2
28       PLP       /
08       PHP
36 36    ROL $36,X
76 36    ROR $36,X
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 34 cycles {#34_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A0 88    LDY #136 ;hides &#39;DEY&#39;
88       DEY
30 FC    BMI *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>C5 C5    CMP $C5
48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A5 A5    LDA $A5
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 35 cycles {#35_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
08       PHP
28       PLP
10 FB    BPL *-3</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 F8    LDX #248 ;hides &#39;SED&#39;
E8   ... INX      × 2
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers X, Z&amp;N, and D</td>
</tr>
<tr>
<td><pre><code>A0 88    LDY #136 ;hides &#39;DEY&#39;
88   ... DEY      × 2
30 FB    BMI *-3</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>98       TYA
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
38   ... SEC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
38   ... SEC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A2 04    LDX #4
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 36 cycles {#36_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 E9    LDA #$E9 ;hides &#39;SBC #$2A&#39;
2A       ROL A ;first loop only
B0 FC    BCS *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, C, and V</td>
</tr>
<tr>
<td><pre><code>A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 06    LDY #6
88       DEY
10 FD    BPL *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>38       SEC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
38       SEC
10 FC    BPL *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA       NOP
90 FC    BCC *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA       NOP
90 FC    BCC *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 04    LDX #4
EA       NOP
CA       DEX
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 04    LDY #4
EA       NOP
88       DEY
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>20 xx xx JSR @rts12× 3</code></pre></td>
<td>Requires @rts12</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
38       SEC
10 FC    BPL *-2
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 37 cycles {#37_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">7 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A5 A5    LDA $A5
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 04    LDX #4
EA   ... NOP      × 2
CA       DEX
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 04    LDY #4
EA   ... NOP      × 2
88       DEY
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
98       TYA
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A2 04    LDX #4
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18   ... CLC      × 2
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 38 cycles {#38_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>38       SEC
A9 69    LDA #$69 ;hides &#39;ADC #$EA&#39;
EA       NOP ;first loop only
B0 FC    BCS *-2</code></pre></td>
<td>Clobbers A, Z&amp;N, C, and V</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 06    LDY #6
88       DEY
10 FD    BPL *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 39 cycles {#39_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A4 A4    LDY $A4
A0 06    LDY #6
88       DEY
10 FD    BPL *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>98       TYA
A0 88    LDY #136 ;hides &#39;DEY&#39;
88   ... DEY      × 2
30 FB    BMI *-3
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 05    LDX #5 ;hides &#39;ORA zp&#39;
CA       DEX ;first loop only
CA       DEX
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers A, and X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 05    LDY #5 ;hides &#39;ORA zp&#39;
88       DEY ;first loop only
88       DEY
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers A, and Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA   ... NOP      × 3
10 FA    BPL *-4
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA   ... NOP      × 3
10 FA    BPL *-4
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
48       PHA
98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
A6 A6    LDX $A6
A2 04    LDX #4
EA       NOP
CA       DEX
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A4 A4    LDY $A4
A0 04    LDY #4
EA       NOP
88       DEY
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">11 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
98       TYA
A0 04    LDY #4
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 40 cycles {#40_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A2 05    LDX #5 ;hides &#39;ORA zp&#39;
EA       NOP
CA       DEX
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 05    LDY #5 ;hides &#39;ORA zp&#39;
EA       NOP
88       DEY
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>98       TYA
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A0 06    LDY #6
88       DEY
10 FD    BPL *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
18   ... CLC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
18   ... CLC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 41 cycles {#41_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>38       SEC
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 08    LDX #8
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 08    LDY #8
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 88    LDY #136 ;hides &#39;DEY&#39;
88       DEY
30 FC    BMI *-2
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
A6 A6    LDX $A6
A2 05    LDX #5
CA       DEX
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A5 A5    LDA $A5
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 42 cycles {#42_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A5 A5    LDA $A5
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A2 05    LDX #5 ;hides &#39;ORA zp&#39;
EA       NOP
CA       DEX
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 05    LDY #5 ;hides &#39;ORA zp&#39;
EA       NOP
88       DEY
D0 FB    BNE *-3</code></pre></td>
<td>Clobbers A, Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
08       PHP
28       PLP
10 FB    BPL *-3
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
08       PHP
28       PLP
10 FB    BPL *-3
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 F8    LDX #248 ;hides &#39;SED&#39;
E8   ... INX      × 2
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 88    LDY #136 ;hides &#39;DEY&#39;
88   ... DEY      × 2
30 FB    BMI *-3
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
98       TYA
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
38   ... SEC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
10 FD    BPL *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 43 cycles {#43_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>38   ... SEC      × 2
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A2 05    LDX #5
EA       NOP
CA       DEX
10 FC    BPL *-2</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 06    LDY #6
EA       NOP
88       DEY
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 E9    LDA #$E9 ;hides &#39;SBC #$2A&#39;
2A       ROL A ;first loop only
B0 FC    BCS *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, C, and V</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 E9    LDA #$E9 ;hides &#39;SBC #$2A&#39;
2A       ROL A ;first loop only
B0 FC    BCS *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
38       SEC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
38       SEC
10 FC    BPL *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA       NOP
90 FC    BCC *-2
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 44 cycles {#44_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
EA   ... NOP      × 2
10 FB    BPL *-3</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A0 88    LDY #136 ;hides &#39;DEY&#39;
EA       NOP
88       DEY
30 FB    BMI *-3</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 08    LDX #8
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>C5 C5    CMP $C5
48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A5 A5    LDA $A5
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 04    LDX #4
EA   ... NOP      × 2
CA       DEX
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 04    LDY #4
EA   ... NOP      × 2
88       DEY
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>EA       NOP
48       PHA
98       TYA
A0 06    LDY #6
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td colspan="2">11 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
98       TYA
A0 05    LDY #5
88       DEY
D0 FD    BNE *-1
A8       TAY
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 45 cycles {#45_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">7 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>98       TYA
A0 08    LDY #8
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A2 08    LDX #8
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
A0 08    LDY #8
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
38       SEC
A9 69    LDA #$69 ;hides &#39;ADC #$EA&#39;
EA       NOP ;first loop only
B0 FC    BCS *-2
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, C, and V</td>
</tr>
<tr>
<td><pre><code>08       PHP
38       SEC
A9 69    LDA #$69 ;hides &#39;ADC #$EA&#39;
EA       NOP ;first loop only
B0 FC    BCS *-2
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA       NOP
08       PHP
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18       CLC
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 46 cycles {#46_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">5 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A2 08    LDX #8
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 09    LDY #9
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">6 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
A6 A6    LDX $A6
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A4 A4    LDY $A4
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
98       TYA
A0 88    LDY #136 ;hides &#39;DEY&#39;
88   ... DEY      × 2
30 FB    BMI *-3
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td colspan="2">11 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
EA   ... NOP      × 3
10 FA    BPL *-4
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 47 cycles {#47_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">8 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>98       TYA
A0 06    LDY #6
EA       NOP
88       DEY
D0 FC    BNE *-2
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
A2 08    LDX #8
CA       DEX
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 05    LDX #5 ;hides &#39;ORA zp&#39;
EA       NOP
CA       DEX
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers A, and X</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 3
A0 08    LDY #8
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 05    LDY #5 ;hides &#39;ORA zp&#39;
EA       NOP
88       DEY
D0 FB    BNE *-3
28       PLP</code></pre></td>
<td>Clobbers A, and Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
98       TYA
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
A8       TAY
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N</td>
</tr>
<tr>
<td><pre><code>08       PHP
98       TYA
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
A8       TAY
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A2 07    LDX #7
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>EA   ... NOP      × 2
08       PHP
A0 06    LDY #6
88       DEY
10 FD    BPL *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
18   ... CLC      × 2
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
90 FD    BCC *-1
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 48 cycles {#48_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>EA       NOP
A2 08    LDX #8
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>EA       NOP
A0 09    LDY #9
88       DEY
D0 FD    BNE *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
38       SEC
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
38       SEC
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 08    LDX #8
CA       DEX
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 08    LDY #8
88       DEY
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 0A    LDA #$0A ;hides &#39;ASL A&#39;
18       CLC
10 FC    BPL *-2
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 49 cycles {#49_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">4 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A0 88    LDY #136 ;hides &#39;DEY&#39;
30 FD    BMI *-1</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>18       CLC
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
08       PHP
28       PLP
90 FB    BCC *-3</code></pre></td>
<td>Clobbers A, Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>A6 A6    LDX $A6
A2 08    LDX #8
CA       DEX
10 FD    BPL *-1</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>C5 C5    CMP $C5
48       PHA
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
A5 A5    LDA $A5
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td colspan="2">10 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 2A    LDA #$2A ;hides &#39;ROL A&#39;
08       PHP
28       PLP
10 FB    BPL *-3
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

### 50 cycles {#50_cycles}

<table class="wikitable testtable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2">6 bytes</th>
</tr>
</thead>
<tbody>
<tr>
<td><pre><code>A9 E9    LDA #$E9 ;hides &#39;SBC #$2A&#39;
2A       ROL A ;first loop only
EA       NOP
B0 FB    BCS *-3</code></pre></td>
<td>Clobbers A, Z&amp;N, C, and V</td>
</tr>
<tr>
<td><pre><code>A2 07    LDX #7
EA       NOP
CA       DEX
D0 FC    BNE *-2</code></pre></td>
<td>Clobbers X, and Z&amp;N</td>
</tr>
<tr>
<td><pre><code>A0 06    LDY #6
EA       NOP
88       DEY
10 FC    BPL *-2</code></pre></td>
<td>Clobbers Y, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">7 bytes</td>
</tr>
<tr>
<td><pre><code>98       TYA
A0 09    LDY #9
88       DEY
D0 FD    BNE *-1
A8       TAY</code></pre></td>
<td>Clobbers A, and Z&amp;N</td>
</tr>
<tr>
<td colspan="2">8 bytes</td>
</tr>
<tr>
<td><pre><code>48       PHA
38   ... SEC      × 2
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
68       PLA</code></pre></td>
<td>Clobbers Z&amp;N, and C</td>
</tr>
<tr>
<td><pre><code>08       PHP
38   ... SEC      × 2
A9 4A    LDA #$4A ;hides &#39;LSR A&#39;
D0 FD    BNE *-1
28       PLP</code></pre></td>
<td>Clobbers A</td>
</tr>
<tr>
<td><pre><code>08       PHP
A2 05    LDX #5
EA       NOP
CA       DEX
10 FC    BPL *-2
28       PLP</code></pre></td>
<td>Clobbers X</td>
</tr>
<tr>
<td><pre><code>08       PHP
A0 06    LDY #6
EA       NOP
88       DEY
D0 FC    BNE *-2
28       PLP</code></pre></td>
<td>Clobbers Y</td>
</tr>
<tr>
<td colspan="2">9 bytes</td>
</tr>
<tr>
<td><pre><code>08       PHP
48       PHA
A9 E9    LDA #$E9 ;hides &#39;SBC #$2A&#39;
2A       ROL A ;first loop only
B0 FC    BCS *-2
68       PLA
28       PLP</code></pre></td>
<td>No requirements</td>
</tr>
</tbody>
</table>

## Sanity checks {#Sanity_checks}

It is possible to verify on compile time that no page wrap occurs, by replacing all branches with these macros:

    .macro branch_check opc, dest
        opc dest
        .assert >* = >(dest), warning, "branch_check: failed, crosses page"
    .endmacro
    .macro bccnw dest
            branch_check bcc, dest
    .endmacro
    .macro bcsnw dest
            branch_check bcs, dest
    .endmacro
    .macro beqnw dest
            branch_check beq, dest
    .endmacro
    .macro bnenw dest
            branch_check bne, dest
    .endmacro
    .macro bminw dest
            branch_check bmi, dest
    .endmacro
    .macro bplnw dest
            branch_check bpl, dest
    .endmacro
    .macro bvcnw dest
            branch_check bvc, dest
    .endmacro
    .macro bvsnw dest
            branch_check bvs, dest
    .endmacro

## See also {#See_also}

- [Cycle counting](../Cycle_counting/)
- [Delay code](../Delay_code/) for functions that produce runtime-determined amount of delay
- Bisqwit’s “vending machine” for producing a ca65-compatible delay_n macro for arbitrary number of cycles, with more fine-grained configurable constraints: <http://bisqwit.iki.fi/utils/nesdelay.php> The samples on this page are excerpts from files generated by this online tool.
