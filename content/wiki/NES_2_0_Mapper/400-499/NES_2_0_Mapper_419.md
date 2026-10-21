---
title: "NES 2.0 Mapper 419"
url: "/wiki/NES_2.0_Mapper_419/"
description: "NES 2.0 Mapper 419 denotes the Taikee TK-8007 MCU circuit board and compatibles, used by several VT03-based plug 'n play consoles:"
wikiPage: "NES_2_0_Mapper_419"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "VTxx mappers"
    url: "/category/VTxx_mappers/"
---

**NES 2.0 Mapper 419** denotes the Taikee **TK-8007 MCU** circuit board and compatibles, used by several VT03-based plug 'n play consoles:

- *Game Sporz Wireless Boxing*
- *Game Sporz Wireless Duet Play Ping-Pong*
- *Game Sporz Wireless Tennis*
- *World Soccer TV Game 10-in-1*

The submapper field denotes the same kind of register/opcode scrambling as on [NES 2.0 Mapper 256](../NES_2.0_Mapper_256/):

<table class="wikitable">
<thead>
<tr>
<th>Submapper #</th>
<th>Name</th>
<th colspan="12">PPU bank affected by ...</th>
<th colspan="4">CPU bank affected by ...</th>
<th>CPU opcode bytes</th>
</tr>
</thead>
<tbody>
<tr>
<th>$2012</th>
<th>$2013</th>
<th>$2014</th>
<th>$2015</th>
<th>$2016</th>
<th>$2017</th>
<th>$8000.0</th>
<th>$8000.1</th>
<th>$8000.2</th>
<th>$8000.3</th>
<th>$8000.4</th>
<th>$8000.5</th>
<th>$4107</th>
<th>$4108</th>
<th>$8000.6</th>
<th>$8000.7</th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>Normal</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>1</td>
<td>Waixing VT03</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>2</td>
<td>Power Joy Supermax</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>none</td>
</tr>
<tr>
<td>3</td>
<td>Zechess/Hummer Team</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1000</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>4</td>
<td>Qishenglong</td>
<td>$1800</td>
<td>$0800</td>
<td>$1000</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>5</td>
<td>Waixing VT02</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>11</td>
<td>Vibes</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6; D5&lt;-&gt;D4 swapped, switched via $411C.1</td>
</tr>
<tr>
<td>12</td>
<td>Cheertone</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6</td>
</tr>
<tr>
<td>13</td>
<td>Taikee</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D1 and D4 swapped, switched via $4169</td>
</tr>
<tr>
<td>14</td>
<td>Karaoto</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D6 and D7 swapped, switched via $411C</td>
</tr>
<tr>
<td>15</td>
<td>Jungletac</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D5 and D6 swapped, switched via $4169</td>
</tr>
</tbody>
</table>

# Expansion audio {#Expansion_audio}

## <span id="I.2FO_Protocol"></span>I/O Protocol {#I.2FO_Protocol}

Instead of normal NES APU audio, the circuit board mounts an ADPCM sound chip that is used for most of the game audio. Command bytes and subsequent data bytes are sent via one of the VT03's general-purpose I/O port in nibblized form:

1.  Enable the I/O port for writing, by writing \$30 to \$410D.
2.  Write the upper nibble to \$410F, set \$4016.2, and wait for \$4017.3 to go low.
3.  Write the lower nibble to \$410F, clear \$4016.2, and wait for \$4017.3 to go high.

There is no means of distinguishing a new command byte from its data bytes. Regardless of the current chip state, sending the byte sequence \$55 \$AA via the above method will reset the chip.

Recognized command bytes:

- \$03 ll mm: Set playback period. The playback rate will be 4.09090909 MHz (45/11 MHz) divided by the playback period.
- \$04 96 bytes: Flush the 96-byte input buffer, fill it with the following 96 bytes, and reset the ADPCM decoder.
- \$06 8 bytes: Add bytes to the input buffer. Upon receiving the command byte, the chip will signal whether it is ready to receive input bytes.
- \$07: Flush the 96-byte input buffer and reset the ADPCM decoder. Sent when sample playback is aborted.

The procedure for command \$06 is as follows:

1.  Send byte \$06 via the above method.
2.  Check the READY bit (\$4017.4 set). If it is not set, stop and come back later.
3.  Send eight bytes via the above method.
4.  Go back to Step 2. As long as the READY bit is set, further data bytes (in groups of eight) *must* be sent. The chip will lower the READY signal once its 96-byte buffer is full. The READY signal is not valid at any time other than after having received the \$06 command byte or after a group of eight data bytes for it.

## ADPCM algorithm {#ADPCM_algorithm}

The input data is composed of frames. One frame is eight bytes/64 bits long and contains the data for 21 output samples, implying three input code bits per sample, decoded bottom-first. If the most significant bit of a 64-bit frame is set, the frame is silent, and all of its lower bits must be ignored.

The decoder maintains two variables of state: a "predictor" (current output) and an "index", both initialized to 0. The lower two bits of the three-bit input code specify one of four 14-element step tables into which the "index" selects a step value, as well as the change of the "index" for the next sample in the form of an index into an index table. The upper bit of the three-bit input code selects the sign of the step value:

    indexStep     [4] ={ 0,  0,  3,  5 };
    indexTable   [26] ={ 0,  0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 20, 20, 20, 20 };
    stepTable [4][21] ={
        {  0,   1,   1,   1,   1,   1,   2,   2,   2,   3,   3,   4,   5,   5,   6,   7,   8,  10,  11,  13,  15 },
        {  1,   3,   3,   3,   4,   4,   6,   6,   7,   9,  10,  12,  15,  16,  19,  22,  25,  30,  34,  40,  46 },
        {  3,   5,   5,   6,   7,   8,  10,  11,  13,  16,  18,  21,  25,  28,  32,  38,  43,  51,  58,  68,  78 },
        {  4,   7,   7,   8,  10,  11,  14,  15,  18,  22,  25,  29,  35,  39,  45,  53,  60,  71,  81,  95, 109 }
    };

    decodeSample (code) {
       predictor +=stepTable[code &3][index] *(code &4? -1: 1);
       index =indexTable[index +indexStep[code &3]];
    }
