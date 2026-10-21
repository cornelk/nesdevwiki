---
title: "CPU unofficial opcodes"
url: "/wiki/CPU_unofficial_opcodes/"
description: "Unofficial opcodes, sometimes called illegal opcodes or undocumented opcodes, are CPU instructions that were officially left unused by the original design. The 6502 family datasheet from MOS…"
wikiPage: "CPU_unofficial_opcodes"
---

**Unofficial opcodes**, sometimes called **illegal opcodes** or **undocumented opcodes**, are [CPU instructions](../6502_instructions/) that were officially left unused by the original design. The 6502 family datasheet from MOS Technology does not specify or document their function, but they actually do perform various operations.

Some of these instructions are useful; some are not predictable; some do nothing but burn cycles; some halt the CPU until reset. Most NMOS 6502 cores interpret them the same way, although there are slight differences with the less stable instructions. CMOS variants of the 6502 handle them completely differently, and later CPUs in the same family (e.g. 65C02, HuC6280, 65C816) were free to implement new instructions in the place of the unused ones.

An [accurate](../Accuracy/) NES emulator must implement all instructions, not just the official ones. A small number of games use them (see [below](#Games_using_unofficial_opcodes)).

## Arrangement {#Arrangement}

The microcode of the 6502 is compressed into a 130-entry decode ROM. Instead of 256 entries telling how to process each separate opcode, it's encoded as combinational logic post-processing the output of a "sparse" ROM that acts in some ways like a programmable logic array (PLA). Each entry in the ROM means "if these bits are on, and these bits are off, do things on these six cycles."<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

Many instructions activate multiple lines of the decode ROM at once. Often this is on purpose, such as one line for the addressing mode and one for the opcode part. But many of the unofficial opcodes simultaneously trigger parts of the ROM that were intended for completely unrelated instructions.

Perhaps the pattern is easier to see by shuffling the 6502's opcode matrix. This table lists all 6502 opcodes, 32 columns per row. The columns are colored by bits 1 and 0: 00 red, 01 green, 10 blue, and 11 gray.

<table style="width:100%;">
<colgroup>
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
</colgroup>
<thead>
<tr>
<th></th>
<th>+00</th>
<th>+01</th>
<th>+02</th>
<th>+03</th>
<th>+04</th>
<th>+05</th>
<th>+06</th>
<th>+07</th>
<th>+08</th>
<th>+09</th>
<th>+0A</th>
<th>+0B</th>
<th>+0C</th>
<th>+0D</th>
<th>+0E</th>
<th>+0F</th>
<th>+10</th>
<th>+11</th>
<th>+12</th>
<th>+13</th>
<th>+14</th>
<th>+15</th>
<th>+16</th>
<th>+17</th>
<th>+18</th>
<th>+19</th>
<th>+1A</th>
<th>+1B</th>
<th>+1C</th>
<th>+1D</th>
<th>+1E</th>
<th>+1F</th>
</tr>
</thead>
<tbody>
<tr data-valign="top">
<td>00</td>
<td style="background: #FCC">BRK<br />
</td>
<td style="background: #CFC">ORA<br />
(d,x)</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>SLO</strong><br />
(d,x)</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #CFC">ORA<br />
d</td>
<td style="background: #CCF">ASL<br />
d</td>
<td style="background: #DDD"><strong>SLO</strong><br />
d</td>
<td style="background: #FCC">PHP<br />
</td>
<td style="background: #CFC">ORA<br />
#i</td>
<td style="background: #CCF">ASL<br />
</td>
<td style="background: #DDD"><strong>ANC</strong><br />
#i</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a</td>
<td style="background: #CFC">ORA<br />
a</td>
<td style="background: #CCF">ASL<br />
a</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a</td>
<td style="background: #FCC">BPL<br />
*+d</td>
<td style="background: #CFC">ORA<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>SLO</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">ORA<br />
d,x</td>
<td style="background: #CCF">ASL<br />
d,x</td>
<td style="background: #DDD"><strong>SLO</strong><br />
d,x</td>
<td style="background: #FCC">CLC<br />
</td>
<td style="background: #CFC">ORA<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">ORA<br />
a,x</td>
<td style="background: #CCF">ASL<br />
a,x</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>20</td>
<td style="background: #FCC">JSR<br />
a</td>
<td style="background: #CFC">AND<br />
(d,x)</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>RLA</strong><br />
(d,x)</td>
<td style="background: #FCC">BIT<br />
d</td>
<td style="background: #CFC">AND<br />
d</td>
<td style="background: #CCF">ROL<br />
d</td>
<td style="background: #DDD"><strong>RLA</strong><br />
d</td>
<td style="background: #FCC">PLP<br />
</td>
<td style="background: #CFC">AND<br />
#i</td>
<td style="background: #CCF">ROL<br />
</td>
<td style="background: #DDD"><strong>ANC</strong><br />
#i</td>
<td style="background: #FCC">BIT<br />
a</td>
<td style="background: #CFC">AND<br />
a</td>
<td style="background: #CCF">ROL<br />
a</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a</td>
<td style="background: #FCC">BMI<br />
*+d</td>
<td style="background: #CFC">AND<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>RLA</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">AND<br />
d,x</td>
<td style="background: #CCF">ROL<br />
d,x</td>
<td style="background: #DDD"><strong>RLA</strong><br />
d,x</td>
<td style="background: #FCC">SEC<br />
</td>
<td style="background: #CFC">AND<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">AND<br />
a,x</td>
<td style="background: #CCF">ROL<br />
a,x</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>40</td>
<td style="background: #FCC">RTI<br />
</td>
<td style="background: #CFC">EOR<br />
(d,x)</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>SRE</strong><br />
(d,x)</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #CFC">EOR<br />
d</td>
<td style="background: #CCF">LSR<br />
d</td>
<td style="background: #DDD"><strong>SRE</strong><br />
d</td>
<td style="background: #FCC">PHA<br />
</td>
<td style="background: #CFC">EOR<br />
#i</td>
<td style="background: #CCF">LSR<br />
</td>
<td style="background: #DDD"><strong>ALR</strong><br />
#i</td>
<td style="background: #FCC">JMP<br />
a</td>
<td style="background: #CFC">EOR<br />
a</td>
<td style="background: #CCF">LSR<br />
a</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a</td>
<td style="background: #FCC">BVC<br />
*+d</td>
<td style="background: #CFC">EOR<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>SRE</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">EOR<br />
d,x</td>
<td style="background: #CCF">LSR<br />
d,x</td>
<td style="background: #DDD"><strong>SRE</strong><br />
d,x</td>
<td style="background: #FCC">CLI<br />
</td>
<td style="background: #CFC">EOR<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">EOR<br />
a,x</td>
<td style="background: #CCF">LSR<br />
a,x</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>60</td>
<td style="background: #FCC">RTS<br />
</td>
<td style="background: #CFC">ADC<br />
(d,x)</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>RRA</strong><br />
(d,x)</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #CFC">ADC<br />
d</td>
<td style="background: #CCF">ROR<br />
d</td>
<td style="background: #DDD"><strong>RRA</strong><br />
d</td>
<td style="background: #FCC">PLA<br />
</td>
<td style="background: #CFC">ADC<br />
#i</td>
<td style="background: #CCF">ROR<br />
</td>
<td style="background: #DDD"><strong>ARR</strong><br />
#i</td>
<td style="background: #FCC">JMP<br />
(a)</td>
<td style="background: #CFC">ADC<br />
a</td>
<td style="background: #CCF">ROR<br />
a</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a</td>
<td style="background: #FCC">BVS<br />
*+d</td>
<td style="background: #CFC">ADC<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>RRA</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">ADC<br />
d,x</td>
<td style="background: #CCF">ROR<br />
d,x</td>
<td style="background: #DDD"><strong>RRA</strong><br />
d,x</td>
<td style="background: #FCC">SEI<br />
</td>
<td style="background: #CFC">ADC<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">ADC<br />
a,x</td>
<td style="background: #CCF">ROR<br />
a,x</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>80</td>
<td style="background: #FCC"><strong>NOP</strong><br />
#i</td>
<td style="background: #CFC">STA<br />
(d,x)</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #DDD"><strong>SAX</strong><br />
(d,x)</td>
<td style="background: #FCC">STY<br />
d</td>
<td style="background: #CFC">STA<br />
d</td>
<td style="background: #CCF">STX<br />
d</td>
<td style="background: #DDD"><strong>SAX</strong><br />
d</td>
<td style="background: #FCC">DEY<br />
</td>
<td style="background: #CFC"><strong>NOP</strong><br />
#i</td>
<td style="background: #CCF">TXA<br />
</td>
<td style="background: #DDD"><strong>XAA</strong><br />
#i</td>
<td style="background: #FCC">STY<br />
a</td>
<td style="background: #CFC">STA<br />
a</td>
<td style="background: #CCF">STX<br />
a</td>
<td style="background: #DDD"><strong>SAX</strong><br />
a</td>
<td style="background: #FCC">BCC<br />
*+d</td>
<td style="background: #CFC">STA<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>AHX</strong><br />
(d),y</td>
<td style="background: #FCC">STY<br />
d,x</td>
<td style="background: #CFC">STA<br />
d,x</td>
<td style="background: #CCF">STX<br />
d,y</td>
<td style="background: #DDD"><strong>SAX</strong><br />
d,y</td>
<td style="background: #FCC">TYA<br />
</td>
<td style="background: #CFC">STA<br />
a,y</td>
<td style="background: #CCF">TXS<br />
</td>
<td style="background: #DDD"><strong>TAS</strong><br />
a,y</td>
<td style="background: #FCC"><strong>SHY</strong><br />
a,x</td>
<td style="background: #CFC">STA<br />
a,x</td>
<td style="background: #CCF"><strong>SHX</strong><br />
a,y</td>
<td style="background: #DDD"><strong>AHX</strong><br />
a,y</td>
</tr>
<tr data-valign="top">
<td>A0</td>
<td style="background: #FCC">LDY<br />
#i</td>
<td style="background: #CFC">LDA<br />
(d,x)</td>
<td style="background: #CCF">LDX<br />
#i</td>
<td style="background: #DDD"><strong>LAX</strong><br />
(d,x)</td>
<td style="background: #FCC">LDY<br />
d</td>
<td style="background: #CFC">LDA<br />
d</td>
<td style="background: #CCF">LDX<br />
d</td>
<td style="background: #DDD"><strong>LAX</strong><br />
d</td>
<td style="background: #FCC">TAY<br />
</td>
<td style="background: #CFC">LDA<br />
#i</td>
<td style="background: #CCF">TAX<br />
</td>
<td style="background: #DDD"><strong>LAX</strong><br />
#i</td>
<td style="background: #FCC">LDY<br />
a</td>
<td style="background: #CFC">LDA<br />
a</td>
<td style="background: #CCF">LDX<br />
a</td>
<td style="background: #DDD"><strong>LAX</strong><br />
a</td>
<td style="background: #FCC">BCS<br />
*+d</td>
<td style="background: #CFC">LDA<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>LAX</strong><br />
(d),y</td>
<td style="background: #FCC">LDY<br />
d,x</td>
<td style="background: #CFC">LDA<br />
d,x</td>
<td style="background: #CCF">LDX<br />
d,y</td>
<td style="background: #DDD"><strong>LAX</strong><br />
d,y</td>
<td style="background: #FCC">CLV<br />
</td>
<td style="background: #CFC">LDA<br />
a,y</td>
<td style="background: #CCF">TSX<br />
</td>
<td style="background: #DDD"><strong>LAS</strong><br />
a,y</td>
<td style="background: #FCC">LDY<br />
a,x</td>
<td style="background: #CFC">LDA<br />
a,x</td>
<td style="background: #CCF">LDX<br />
a,y</td>
<td style="background: #DDD"><strong>LAX</strong><br />
a,y</td>
</tr>
<tr data-valign="top">
<td>C0</td>
<td style="background: #FCC">CPY<br />
#i</td>
<td style="background: #CFC">CMP<br />
(d,x)</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #DDD"><strong>DCP</strong><br />
(d,x)</td>
<td style="background: #FCC">CPY<br />
d</td>
<td style="background: #CFC">CMP<br />
d</td>
<td style="background: #CCF">DEC<br />
d</td>
<td style="background: #DDD"><strong>DCP</strong><br />
d</td>
<td style="background: #FCC">INY<br />
</td>
<td style="background: #CFC">CMP<br />
#i</td>
<td style="background: #CCF">DEX<br />
</td>
<td style="background: #DDD"><strong>AXS</strong><br />
#i</td>
<td style="background: #FCC">CPY<br />
a</td>
<td style="background: #CFC">CMP<br />
a</td>
<td style="background: #CCF">DEC<br />
a</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a</td>
<td style="background: #FCC">BNE<br />
*+d</td>
<td style="background: #CFC">CMP<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>DCP</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">CMP<br />
d,x</td>
<td style="background: #CCF">DEC<br />
d,x</td>
<td style="background: #DDD"><strong>DCP</strong><br />
d,x</td>
<td style="background: #FCC">CLD<br />
</td>
<td style="background: #CFC">CMP<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">CMP<br />
a,x</td>
<td style="background: #CCF">DEC<br />
a,x</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>E0</td>
<td style="background: #FCC">CPX<br />
#i</td>
<td style="background: #CFC">SBC<br />
(d,x)</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #DDD"><strong>ISC</strong><br />
(d,x)</td>
<td style="background: #FCC">CPX<br />
d</td>
<td style="background: #CFC">SBC<br />
d</td>
<td style="background: #CCF">INC<br />
d</td>
<td style="background: #DDD"><strong>ISC</strong><br />
d</td>
<td style="background: #FCC">INX<br />
</td>
<td style="background: #CFC">SBC<br />
#i</td>
<td style="background: #CCF">NOP<br />
</td>
<td style="background: #DDD"><strong>SBC</strong><br />
#i</td>
<td style="background: #FCC">CPX<br />
a</td>
<td style="background: #CFC">SBC<br />
a</td>
<td style="background: #CCF">INC<br />
a</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a</td>
<td style="background: #FCC">BEQ<br />
*+d</td>
<td style="background: #CFC">SBC<br />
(d),y</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #DDD"><strong>ISC</strong><br />
(d),y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #CFC">SBC<br />
d,x</td>
<td style="background: #CCF">INC<br />
d,x</td>
<td style="background: #DDD"><strong>ISC</strong><br />
d,x</td>
<td style="background: #FCC">SED<br />
</td>
<td style="background: #CFC">SBC<br />
a,y</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a,y</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">SBC<br />
a,x</td>
<td style="background: #CCF">INC<br />
a,x</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a,x</td>
</tr>
</tbody>
</table>

Key: a is a 16-bit absolute address, and d is an 8-bit zero page address. Entries in bold represent unofficial opcodes.

But if we rearrange it so that columns with the same bits 1-0 are close together, correlations become easier to see:

<table style="width:100%;">
<colgroup>
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
<col style="width: 3%" />
</colgroup>
<thead>
<tr>
<th></th>
<th>+00</th>
<th>+04</th>
<th>+08</th>
<th>+0C</th>
<th>+10</th>
<th>+14</th>
<th>+18</th>
<th>+1C</th>
<th>+01</th>
<th>+05</th>
<th>+09</th>
<th>+0D</th>
<th>+11</th>
<th>+15</th>
<th>+19</th>
<th>+1D</th>
<th>+02</th>
<th>+06</th>
<th>+0A</th>
<th>+0E</th>
<th>+12</th>
<th>+16</th>
<th>+1A</th>
<th>+1E</th>
<th>+03</th>
<th>+07</th>
<th>+0B</th>
<th>+0F</th>
<th>+13</th>
<th>+17</th>
<th>+1B</th>
<th>+1F</th>
</tr>
</thead>
<tbody>
<tr data-valign="top">
<td>00</td>
<td style="background: #FCC">BRK<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #FCC">PHP<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a</td>
<td style="background: #FCC">BPL<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">CLC<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">ORA<br />
(d,x)</td>
<td style="background: #CFC">ORA<br />
d</td>
<td style="background: #CFC">ORA<br />
#i</td>
<td style="background: #CFC">ORA<br />
a</td>
<td style="background: #CFC">ORA<br />
(d),y</td>
<td style="background: #CFC">ORA<br />
d,x</td>
<td style="background: #CFC">ORA<br />
a,y</td>
<td style="background: #CFC">ORA<br />
a,x</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ASL<br />
d</td>
<td style="background: #CCF">ASL<br />
</td>
<td style="background: #CCF">ASL<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ASL<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">ASL<br />
a,x</td>
<td style="background: #DDD"><strong>SLO</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>SLO</strong><br />
d</td>
<td style="background: #DDD"><strong>ANC</strong><br />
#i</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a</td>
<td style="background: #DDD"><strong>SLO</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>SLO</strong><br />
d,x</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a,y</td>
<td style="background: #DDD"><strong>SLO</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>20</td>
<td style="background: #FCC">JSR<br />
a</td>
<td style="background: #FCC">BIT<br />
d</td>
<td style="background: #FCC">PLP<br />
</td>
<td style="background: #FCC">BIT<br />
a</td>
<td style="background: #FCC">BMI<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">SEC<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">AND<br />
(d,x)</td>
<td style="background: #CFC">AND<br />
d</td>
<td style="background: #CFC">AND<br />
#i</td>
<td style="background: #CFC">AND<br />
a</td>
<td style="background: #CFC">AND<br />
(d),y</td>
<td style="background: #CFC">AND<br />
d,x</td>
<td style="background: #CFC">AND<br />
a,y</td>
<td style="background: #CFC">AND<br />
a,x</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ROL<br />
d</td>
<td style="background: #CCF">ROL<br />
</td>
<td style="background: #CCF">ROL<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ROL<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">ROL<br />
a,x</td>
<td style="background: #DDD"><strong>RLA</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>RLA</strong><br />
d</td>
<td style="background: #DDD"><strong>ANC</strong><br />
#i</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a</td>
<td style="background: #DDD"><strong>RLA</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>RLA</strong><br />
d,x</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a,y</td>
<td style="background: #DDD"><strong>RLA</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>40</td>
<td style="background: #FCC">RTI<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #FCC">PHA<br />
</td>
<td style="background: #FCC">JMP<br />
a</td>
<td style="background: #FCC">BVC<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">CLI<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">EOR<br />
(d,x)</td>
<td style="background: #CFC">EOR<br />
d</td>
<td style="background: #CFC">EOR<br />
#i</td>
<td style="background: #CFC">EOR<br />
a</td>
<td style="background: #CFC">EOR<br />
(d),y</td>
<td style="background: #CFC">EOR<br />
d,x</td>
<td style="background: #CFC">EOR<br />
a,y</td>
<td style="background: #CFC">EOR<br />
a,x</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">LSR<br />
d</td>
<td style="background: #CCF">LSR<br />
</td>
<td style="background: #CCF">LSR<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">LSR<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">LSR<br />
a,x</td>
<td style="background: #DDD"><strong>SRE</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>SRE</strong><br />
d</td>
<td style="background: #DDD"><strong>ALR</strong><br />
#i</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a</td>
<td style="background: #DDD"><strong>SRE</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>SRE</strong><br />
d,x</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a,y</td>
<td style="background: #DDD"><strong>SRE</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>60</td>
<td style="background: #FCC">RTS<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d</td>
<td style="background: #FCC">PLA<br />
</td>
<td style="background: #FCC">JMP<br />
(a)</td>
<td style="background: #FCC">BVS<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">SEI<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">ADC<br />
(d,x)</td>
<td style="background: #CFC">ADC<br />
d</td>
<td style="background: #CFC">ADC<br />
#i</td>
<td style="background: #CFC">ADC<br />
a</td>
<td style="background: #CFC">ADC<br />
(d),y</td>
<td style="background: #CFC">ADC<br />
d,x</td>
<td style="background: #CFC">ADC<br />
a,y</td>
<td style="background: #CFC">ADC<br />
a,x</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ROR<br />
d</td>
<td style="background: #CCF">ROR<br />
</td>
<td style="background: #CCF">ROR<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">ROR<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">ROR<br />
a,x</td>
<td style="background: #DDD"><strong>RRA</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>RRA</strong><br />
d</td>
<td style="background: #DDD"><strong>ARR</strong><br />
#i</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a</td>
<td style="background: #DDD"><strong>RRA</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>RRA</strong><br />
d,x</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a,y</td>
<td style="background: #DDD"><strong>RRA</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>80</td>
<td style="background: #FCC"><strong>NOP</strong><br />
#i</td>
<td style="background: #FCC">STY<br />
d</td>
<td style="background: #FCC">DEY<br />
</td>
<td style="background: #FCC">STY<br />
a</td>
<td style="background: #FCC">BCC<br />
*+d</td>
<td style="background: #FCC">STY<br />
d,x</td>
<td style="background: #FCC">TYA<br />
</td>
<td style="background: #FCC"><strong>SHY</strong><br />
a,x</td>
<td style="background: #CFC">STA<br />
(d,x)</td>
<td style="background: #CFC">STA<br />
d</td>
<td style="background: #CFC"><strong>NOP</strong><br />
#i</td>
<td style="background: #CFC">STA<br />
a</td>
<td style="background: #CFC">STA<br />
(d),y</td>
<td style="background: #CFC">STA<br />
d,x</td>
<td style="background: #CFC">STA<br />
a,y</td>
<td style="background: #CFC">STA<br />
a,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #CCF">STX<br />
d</td>
<td style="background: #CCF">TXA<br />
</td>
<td style="background: #CCF">STX<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">STX<br />
d,y</td>
<td style="background: #CCF">TXS<br />
</td>
<td style="background: #CCF"><strong>SHX</strong><br />
a,y</td>
<td style="background: #DDD"><strong>SAX</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>SAX</strong><br />
d</td>
<td style="background: #DDD"><strong>XAA</strong><br />
#i</td>
<td style="background: #DDD"><strong>SAX</strong><br />
a</td>
<td style="background: #DDD"><strong>AHX</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>SAX</strong><br />
d,y</td>
<td style="background: #DDD"><strong>TAS</strong><br />
a,y</td>
<td style="background: #DDD"><strong>AHX</strong><br />
a,y</td>
</tr>
<tr data-valign="top">
<td>A0</td>
<td style="background: #FCC">LDY<br />
#i</td>
<td style="background: #FCC">LDY<br />
d</td>
<td style="background: #FCC">TAY<br />
</td>
<td style="background: #FCC">LDY<br />
a</td>
<td style="background: #FCC">BCS<br />
*+d</td>
<td style="background: #FCC">LDY<br />
d,x</td>
<td style="background: #FCC">CLV<br />
</td>
<td style="background: #FCC">LDY<br />
a,x</td>
<td style="background: #CFC">LDA<br />
(d,x)</td>
<td style="background: #CFC">LDA<br />
d</td>
<td style="background: #CFC">LDA<br />
#i</td>
<td style="background: #CFC">LDA<br />
a</td>
<td style="background: #CFC">LDA<br />
(d),y</td>
<td style="background: #CFC">LDA<br />
d,x</td>
<td style="background: #CFC">LDA<br />
a,y</td>
<td style="background: #CFC">LDA<br />
a,x</td>
<td style="background: #CCF">LDX<br />
#i</td>
<td style="background: #CCF">LDX<br />
d</td>
<td style="background: #CCF">TAX<br />
</td>
<td style="background: #CCF">LDX<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">LDX<br />
d,y</td>
<td style="background: #CCF">TSX<br />
</td>
<td style="background: #CCF">LDX<br />
a,y</td>
<td style="background: #DDD"><strong>LAX</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>LAX</strong><br />
d</td>
<td style="background: #DDD"><strong>LAX</strong><br />
#i</td>
<td style="background: #DDD"><strong>LAX</strong><br />
a</td>
<td style="background: #DDD"><strong>LAX</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>LAX</strong><br />
d,y</td>
<td style="background: #DDD"><strong>LAS</strong><br />
a,y</td>
<td style="background: #DDD"><strong>LAX</strong><br />
a,y</td>
</tr>
<tr data-valign="top">
<td>C0</td>
<td style="background: #FCC">CPY<br />
#i</td>
<td style="background: #FCC">CPY<br />
d</td>
<td style="background: #FCC">INY<br />
</td>
<td style="background: #FCC">CPY<br />
a</td>
<td style="background: #FCC">BNE<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">CLD<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">CMP<br />
(d,x)</td>
<td style="background: #CFC">CMP<br />
d</td>
<td style="background: #CFC">CMP<br />
#i</td>
<td style="background: #CFC">CMP<br />
a</td>
<td style="background: #CFC">CMP<br />
(d),y</td>
<td style="background: #CFC">CMP<br />
d,x</td>
<td style="background: #CFC">CMP<br />
a,y</td>
<td style="background: #CFC">CMP<br />
a,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #CCF">DEC<br />
d</td>
<td style="background: #CCF">DEX<br />
</td>
<td style="background: #CCF">DEC<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">DEC<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">DEC<br />
a,x</td>
<td style="background: #DDD"><strong>DCP</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>DCP</strong><br />
d</td>
<td style="background: #DDD"><strong>AXS</strong><br />
#i</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a</td>
<td style="background: #DDD"><strong>DCP</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>DCP</strong><br />
d,x</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a,y</td>
<td style="background: #DDD"><strong>DCP</strong><br />
a,x</td>
</tr>
<tr data-valign="top">
<td>E0</td>
<td style="background: #FCC">CPX<br />
#i</td>
<td style="background: #FCC">CPX<br />
d</td>
<td style="background: #FCC">INX<br />
</td>
<td style="background: #FCC">CPX<br />
a</td>
<td style="background: #FCC">BEQ<br />
*+d</td>
<td style="background: #FCC"><strong>NOP</strong><br />
d,x</td>
<td style="background: #FCC">SED<br />
</td>
<td style="background: #FCC"><strong>NOP</strong><br />
a,x</td>
<td style="background: #CFC">SBC<br />
(d,x)</td>
<td style="background: #CFC">SBC<br />
d</td>
<td style="background: #CFC">SBC<br />
#i</td>
<td style="background: #CFC">SBC<br />
a</td>
<td style="background: #CFC">SBC<br />
(d),y</td>
<td style="background: #CFC">SBC<br />
d,x</td>
<td style="background: #CFC">SBC<br />
a,y</td>
<td style="background: #CFC">SBC<br />
a,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
#i</td>
<td style="background: #CCF">INC<br />
d</td>
<td style="background: #CCF">NOP<br />
</td>
<td style="background: #CCF">INC<br />
a</td>
<td style="background: #CCF"><strong>STP</strong><br />
</td>
<td style="background: #CCF">INC<br />
d,x</td>
<td style="background: #CCF"><strong>NOP</strong><br />
</td>
<td style="background: #CCF">INC<br />
a,x</td>
<td style="background: #DDD"><strong>ISC</strong><br />
(d,x)</td>
<td style="background: #DDD"><strong>ISC</strong><br />
d</td>
<td style="background: #DDD"><strong>SBC</strong><br />
#i</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a</td>
<td style="background: #DDD"><strong>ISC</strong><br />
(d),y</td>
<td style="background: #DDD"><strong>ISC</strong><br />
d,x</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a,y</td>
<td style="background: #DDD"><strong>ISC</strong><br />
a,x</td>
</tr>
</tbody>
</table>

Thus the 00 (red) block is mostly control instructions, 01 (green) is ALU operations, and 10 (blue) is read-modify-write (RMW) operations and data movement instructions involving X. The RMW instructions (all but row 80 and A0) in columns +06, +0E, +16, and +1E have the same addressing modes as the corresponding ALU operations.

The 11 (gray) block is unofficial opcodes combining the operations of instructions from the ALU and RMW blocks. all of them having the same addressing mode as the corresponding ALU opcode. The RMW+ALU instructions that affect memory are easiest to understand because their RMW part completes during the opcode and the ALU part completes during the next opcode's fetch. Column +0B, on the other hand, has no extra cycles; everything completes during the next opcode's fetch. This causes instructions to have strange mixing properties. Some even [differ based on analog effects](../Visual6502wiki/6502_Opcode_8B_%28XAA%2C_ANE%29/).

## Games using unofficial opcodes {#Games_using_unofficial_opcodes}

The use of unofficial opcodes is rare in NES games. It appears to occur mostly in late or unlicensed titles:

- *Beauty and the Beast* (E) (1994) uses `$80` (a 2-byte NOP).<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>
- *Disney's Aladdin* (E) (December 1994) uses `$07` (SLO). This is Virgin's port of the Game Boy game, itself a port of the Genesis game, not any of the [pirate originals](https://bootleggames.fandom.com/wiki/Special:PrefixIndex/Aladdin).
- *Dynowarz: Destruction of Spondylus* (April 1990) uses 1-byte NOPs `$DA` and `$FA` on the first level when your dino throws his fist.
- *F-117A Stealth Fighter* uses `$89` (a 2-byte NOP).
- *文字广场+排雷* (romanized in GoodNES as Cantonese "Gaau Hok Gwong Cheung (Ch)"): After selecting the left game (排雷) from this 2-in-1 multicart, a glitchy [32 KiB bankswitch](../INES_Mapper_242/) causes the CPU to get lost in non-code ROM space that only with correct emulation of unofficial opcodes, including \$8B (XAA immediate), will have it eventually reach a BRK instruction that properly branches to that game's reset handler.
- *Infiltrator* uses `$89` (a 2-byte NOP).
- *Ninja Jajamaru-kun* uses `$04` (a 2-byte NOP) when your ninja collides with an enemy, as a consequence of branching into the middle of an instruction.
- *Puzznic* (all regions) (US release November 1990) uses `$89` (a 2-byte NOP).
- *Super Cars* (U) (February 1991) uses `$B3` (LAX).

### Homebrew games {#Homebrew_games}

- The MUSE music engine, used in *Driar* and *STREEMERZ: Super Strength Emergency Squad Zeta*, uses `$8F` (SAX), `$B3` (LAX), and `$CB` (AXS).<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>
- *Attribute Zone* uses `$0B` (ANC), `$2F` (RLA), `$4B` (ALR), `$A7` (LAX), `$B3` (LAX), `$CB` (AXS), `$D3` (DCP) and `$DB` (DCP).
- The port of *Zork* to the Famicom uses a few unofficial opcodes.
- *Eyra, the Crow Maiden* uses several unofficial opcodes.

## See also {#See_also}

- [Programming with unofficial opcodes](../Programming_with_unofficial_opcodes/)

## External links {#External_links}

- [6502 opcode matrix including unofficial opcodes](http://www.oxyron.de/html/opcodes02.html)
- [65C02](http://www.oxyron.de/html/opcodesc02.html) and [65816](http://www.oxyron.de/html/opcodes816.html)
- [Illegal opcodes](https://en.wikipedia.org/wiki/Illegal_opcode) at Wikipedia.
- [65xx Processor Data](http://www.romhacking.net/documents/318/)
- [6502_cpu.txt](http://nesdev.org/6502_cpu.txt)
- [No More Secrets v1.00](https://csdb.dk/release/?id=258111)

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) Michael Steil. "[How MOS 6502 Illegal Opcodes really work](https://www.pagetable.com/?p=39)". *Pagetable*, 2008-07-29. Accessed 2019-11-10.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [puNES 0.20 changelog](http://forums.nesdev.org/viewtopic.php?f=3&t=6928) indicating \$80 opcode in *Beauty and the Beast*.
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) <http://forums.nesdev.org/viewtopic.php?p=100957#p100957>
