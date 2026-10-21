---
title: "Sachen SA8259A pinout"
url: "/wiki/Sachen_SA8259A_pinout/"
description: "Most of Sachen's mapper ICs have names that collide with other well-known ICs, presumably as an initial round of indirection to confuse would-be mass copiers. The 8259A is canonically the original…"
wikiPage: "Sachen_SA8259A_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

[Sachen SA8259A](../Sachen_8259/): 28-pin 0.6" DIP

Most of Sachen's mapper ICs have names that collide with other well-known ICs, presumably as an initial round of indirection to confuse would-be mass copiers. The 8259A is canonically the original IBM PC's [Programmable Interrupt Controller](https://en.wikipedia.org/wiki/Programmable_Interrupt_Controller).

                  .---\/---.              
       PRG A15 <- | 01  28 | -- Vcc
     CIRAM A10 <- | 02  27 | -> r42
       PPU A11 -> | 03  26 | <- CPU A8 (+CE)
       PPU A13 -> | 04  25 | -> rX0
       PPU /RD -> | 05  24 | <- R/W
             ? ?? | 06  23 | <- M2
       PRG A16 <- | 07  22 | <- PPU A12
       CPU A14 -> | 08  21 | <- PPU A10
       CPU  A0 -> | 09  20 | <- /ROMSEL
           rX1 <- | 10  19 | -> r41
       CPU  D0 -> | 11  18 | <- CPU D2
       CPU  D1 -> | 12  17 | -> r60
       PRG A17 <- | 13  16 | -> rX2
           GND -- | 14  15 | -> r40
                  '--------'

The exact connectivity of the pins labelled r?? depends on the corresponding mapper number or name allocated.

<table class="datatable">
<thead>
<tr>
<th>pin \ mapper</th>
<th><a href="../INES_Mapper_137/">137</a></th>
<th>138</th>
<th>139</th>
<th>141</th>
</tr>
</thead>
<tbody>
<tr>
<td>rX0</td>
<td>CHR A10</td>
<td>CHR A11</td>
<td>CHR A13</td>
<td>CHR A12</td>
</tr>
<tr>
<td>rX1</td>
<td>CHR A11</td>
<td>CHR A12</td>
<td>CHR A14</td>
<td>CHR A13</td>
</tr>
<tr>
<td>rX2</td>
<td>CHR A12</td>
<td>CHR A13</td>
<td>CHR A15</td>
<td>CHR A14</td>
</tr>
<tr>
<td>r40</td>
<td>CHR A14¹</td>
<td>CHR A14</td>
<td>CHR A16</td>
<td>CHR A15</td>
</tr>
<tr>
<td>r41</td>
<td>CHR A14¹</td>
<td>CHR A15</td>
<td>CHR A17</td>
<td>CHR A16</td>
</tr>
<tr>
<td>r42</td>
<td>CHR A14¹</td>
<td>CHR A16</td>
<td>CHR A18</td>
<td>CHR A17</td>
</tr>
<tr>
<td>r60</td>
<td>CHR A13¹</td>
<td>n/c</td>
<td>n/c</td>
<td>n/c</td>
</tr>
<tr>
<td>PPU A10</td>
<td>PPU A11</td>
<td colspan="3" style="text-align: center;">PPU A10</td>
</tr>
<tr>
<td>PPU A11</td>
<td>PPU A10</td>
<td colspan="3" style="text-align: center;">PPU A11</td>
</tr>
<tr>
<td>PPU A12</td>
<td>PPU A11</td>
<td colspan="3" style="text-align: center;">PPU A12</td>
</tr>
</tbody>
</table>

¹: r40, r41, r42, and r60 go through a 74LS253 dual 4 to 1 multiplexer and a 74LS257 quad 2 to 1 multiplexer to produce 4x1+4F CHR banking

Reference: Cah4e3 ( <http://cah4e3.shedevr.org.ru/%5Blst%5D-sachen-mappers.txt> )

- [Picture of box, cart, and PCB of *Super Pang II*](https://www.flickr.com/photos/153392699@N08/sets/72157679901951752)
