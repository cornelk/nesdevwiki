---
title: "Visual6502wiki/6502 Opcode 8B (XAA, ANE)"
url: "/wiki/Visual6502wiki/6502_Opcode_8B_(XAA,_ANE)/"
description: "Of all the unsupported opcodes, 8B has had a lot of attention because it seems unpredictable. Even the same computer has been seen to act differently even with the same inputs."
wikiPage: "Visual6502wiki_6502_Opcode_8B__XAA__ANE_"
---

Of all the unsupported opcodes, 8B has had a lot of attention because it seems unpredictable. Even the same computer has been seen to act differently even with the same inputs.

## Explanation {#Explanation}

The reason is that this opcode connects the A register to SB (the Special Bus) at both input and output: in a sense, A is both read from and written to at the same time. Unlike the stack pointer, the A register is not designed to do that, and the result is a circuit configuration which behaves in an interesting way.

Note that our switch-level simulation tends to produce wired-AND behaviour: if two logic gates both drive the same wire, then either of them can drive it low. A real 6502 usually does the same, which is why 8B - often called XAA - will more or less AND together the three inputs: the X register, the A register, and the immediate operand.

Why more or less? Two reasons: the A register is fed back on itself, and because of an interaction with the RDY input.

The A register drives the SB directly, and bits 0 and 4 read SB directly. The other 6 bits read SB through the Decimal Adjust logic, which doesn't affect the logic value but does affect the timing, the logic thresholds and the drive strengths. Exactly what happens is an analogue problem, not a digital one, so it will depend on the exact model of CPU, the variations of chip manufacture, the power supply and the temperature. We can't even model this without knowing the transistor strengths and having some idea of the transistor parameters - which we can only guess at.

The RDY input is a more digital influence on the outcome. RDY is intended to stall the CPU during read accesses, so it can read from slow memory. As it happens, the 6502 samples the databus on every falling clock edge, and loads the IDL (Input Data Latch), and then drives into the target register. Normally, the final cycle is the one which counts, overwriting the stray external values. In some computers, RDY is used to stall the CPU while the bus is used for DMA, which means the bus contains data such as video data for several cycles, except the last. In the case of XAA, every cycle's data is ANDed into A, and this is why the final value of A changes even for the same values of operand, X and A.

## Circuit Diagram {#Circuit_Diagram}

Here's an abridged circuit diagram. Note that bits 0 and 4 have direct A feedback whereas the other bits have indirect feedback. Note that phi1 is when A is written, but the preceding phi2 is when the operand is loaded and the two busses precharged high.

[<img src="../../../wiki-images/6502-XAA-Idb-sb.png" width="945" height="476" />](../../../wiki-images/6502-XAA-Idb-sb.png)

(Logic gate pullups shown as resistors, although in NMOS logic pullups are not usually depletion-mode transistors. They pull up to the positive rail. The pass transistors and precharges cannot pull up to the rail: they drop a threshold voltage. These considerations will affect an analogue analysis.)

## Testing this opcode {#Testing_this_opcode}

This opcode has 3 bytes of input, supposing that we're not allowing RDY to stall the machine and add more operands. We have a test program which tests 256^3 combinations of inputs and compares the final A and the two affected flags against a model. We also have a few specific combinations we've used to characterise different chips.

- \<describe or define the programs here\>
- \<also mention the Java simulation which tests the robustness of the switch simulator results (against the order of evaluation)\>

## Modelling this opcode {#Modelling_this_opcode}

\<mention and link to an emulator code fragment\>

The base formula for XAA seems to be:

    A = (A | magic) & X & imm

"magic" defines which bits of A shine through.

## Tested CPUs {#Tested_CPUs}

We collect here some results of testing this opcode on various CPUs from different manufacturers and in various computers.

<table class="wikitable sortable" style="text-align: center; font-size: 85%; width: auto; table-layout: fixed;">
<colgroup>
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
<col style="width: 7%" />
</colgroup>
<thead>
<tr>
<th>manufacturer</th>
<th>type</th>
<th>YYWW</th>
<th>country</th>
<th>markings</th>
<th>on back</th>
<th>device tested in</th>
<th>tester</th>
<th>magic</th>
<th>RDY clears #4</th>
<th>stable*</th>
<th>N,Z flags OK**</th>
<th>notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>MOS</td>
<td>6502</td>
<td>7551</td>
<td>USA?</td>
<td>MOS<br />
MCS 6502<br />
5175</td>
<td>?</td>
<td>KIM-1</td>
<td>Michael</td>
<td>FF</td>
<td>?</td>
<td>?</td>
<td>?</td>
<td>only minimal testing done</td>
</tr>
<tr>
<td>MOS</td>
<td>6502</td>
<td>8402</td>
<td>Philippines</td>
<td>MOS<br />
6502<br />
0284</td>
<td>PHILIPPINES<br />
IH434564</td>
<td>CBM1541</td>
<td>Michael</td>
<td>EE</td>
<td>?</td>
<td>yes</td>
<td>?</td>
<td>this is the chip that came with this disk drive</td>
</tr>
<tr>
<td>MOS</td>
<td>6502B</td>
<td>8207</td>
<td>Korea</td>
<td>MOS<br />
C01437706<br />
0782</td>
<td>6502<br />
KOREA<br />
5231 07<br />
03-82</td>
<td>CBM1541</td>
<td>Michael</td>
<td>EE</td>
<td>?</td>
<td>yes</td>
<td>?</td>
<td>from my Atari 800</td>
</tr>
<tr>
<td>Rockwell</td>
<td>6502</td>
<td>8228</td>
<td>Mexico</td>
<td>R6502P<br />
R6502-11<br />
8228</td>
<td>R6502F<br />
MEXICO<br />
0737</td>
<td>CBM1541</td>
<td>Michael</td>
<td>FF</td>
<td>?</td>
<td>yes</td>
<td>?</td>
<td>Simon's; spare part bought from retailer</td>
</tr>
<tr>
<td>MOS</td>
<td>6510</td>
<td>8337</td>
<td>?</td>
<td>MOS<br />
6510CBM<br />
3783</td>
<td>?</td>
<td>C64<br />
NTSC<br />
250407/REV.B</td>
<td>Michael</td>
<td>FF</td>
<td>no</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>MOS</td>
<td>6510</td>
<td>8431</td>
<td>Hong Kong</td>
<td>MOS<br />
6510CBM<br />
3184</td>
<td>HONG KONG<br />
HH265111</td>
<td>C64<br />
NTSC<br />
250407/REV.A</td>
<td>Michael</td>
<td>FF</td>
<td>no</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>MOS</td>
<td>8500</td>
<td>8551</td>
<td>?</td>
<td>MOS<br />
8500R3<br />
5185</td>
<td>?</td>
<td>C64<br />
PAL<br />
250425/REV.B</td>
<td>Michael</td>
<td>FE</td>
<td>yes</td>
<td>yes</td>
<td>?</td>
<td><em>very</em> early 8500</td>
</tr>
<tr>
<td>MOS</td>
<td>8500</td>
<td>9009</td>
<td>Hong Kong</td>
<td>CSG<br />
8500<br />
0990 24</td>
<td>HONG KONG<br />
HH096205<br />
MP150SG</td>
<td>C64<br />
PAL<br />
250469/REV.B</td>
<td>Michael</td>
<td>FE</td>
<td>yes</td>
<td>yes</td>
<td>?</td>
<td>very late 6502-like CPU</td>
</tr>
<tr>
<td>MOS</td>
<td>6502AD</td>
<td>8521</td>
<td>?</td>
<td>MOS<br />
6502AD<br />
2185</td>
<td>?</td>
<td>CBM1541</td>
<td>Michael</td>
<td>FF</td>
<td>?</td>
<td>no</td>
<td>?</td>
<td>bit #3 of X input gets treated as "bit #3 of X &amp; bit #4 of X" most of the time (depends on A though)<br />
very unstable<br />
1 MHz mode tested, can also do 2 MHz; chip is from a VC1571</td>
</tr>
<tr>
<td>Pravetz</td>
<td>CM630P</td>
<td>8744</td>
<td>Hungary</td>
<td>(symbol)<br />
CM630P<br />
8744</td>
<td>(none)</td>
<td>CBM1541</td>
<td>Michael</td>
<td>FE</td>
<td>?</td>
<td>no</td>
<td>yes</td>
<td>bit #4 anomalies</td>
</tr>
<tr>
<td>Synertek</td>
<td>6502</td>
<td>????</td>
<td>?</td>
<td>SY<br />
6502<br />
TODO</td>
<td>?</td>
<td>CBM1541</td>
<td>Michael</td>
<td></td>
<td></td>
<td></td>
<td>?</td>
<td>Simon's; yet to test</td>
</tr>
<tr>
<td>MOS</td>
<td>8502</td>
<td>????</td>
<td>?</td>
<td>MOS<br />
8502<br />
TODO</td>
<td>?</td>
<td>C128D</td>
<td>Michael</td>
<td></td>
<td></td>
<td></td>
<td>?</td>
<td>yet to test; can do 1 MHz and 2 MHz</td>
</tr>
<tr>
<td>MOS</td>
<td>6502</td>
<td>????</td>
<td>?</td>
<td>MOS<br />
6502<br />
TODO</td>
<td>?</td>
<td>VC1581</td>
<td>Michael</td>
<td></td>
<td></td>
<td></td>
<td>?</td>
<td>yet to test; can do 1 MHz and 2 MHz</td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8323</td>
<td>?</td>
<td>C014806-03<br />
8323</td>
<td>?</td>
<td>Atari 800XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>yes</td>
<td>yes</td>
<td></td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8320</td>
<td>?</td>
<td>C014806-03<br />
8320</td>
<td>?</td>
<td>Atari 800XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>almost</td>
<td>?</td>
<td>40 errors in 256^3 full test<br />
sometimes bit 3 was set</td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8408</td>
<td>?</td>
<td>C014806-03<br />
8408</td>
<td>?</td>
<td>Atari 800XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>no</td>
<td>yes</td>
<td>~150k - 450k errors (1% - 2.7%) in full test<br />
sometimes bit 3 set, for example A=03 X=FF imm=FF results either in 03 or 0B in repeated tests</td>
</tr>
<tr>
<td>Rockwell</td>
<td>SALLY</td>
<td>8322</td>
<td>?</td>
<td>C014806-12<br />
11151-12<br />
8322</td>
<td>?</td>
<td>Atari 800XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>no</td>
<td>almost</td>
<td>~30k - 80k errors (0.2% - 0.5%) in full test<br />
sometimes bit 3 is set, but also bit 2 and 5 were set sometimes<br />
for example A=5F or A=87 resulted in a set bit 3 (quite frequently), bit 5 (less frequently) or bit 2 (least frequent)<br />
only flipping from 0 to 1 observed, no flipping from 1 to 0<br />
flags were wrong 115 times (~7ppm)</td>
</tr>
<tr>
<td>NCR</td>
<td>SALLY</td>
<td>8737</td>
<td>?</td>
<td>NCR C014806C-29<br />
F826948 S8737</td>
<td>?</td>
<td>Atari 800XE</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>?</td>
<td>SALLY</td>
<td>?</td>
<td>?</td>
<td>C014806-35<br />
(C) ATARI 1980</td>
<td>?</td>
<td>Atari 65XE</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>no</td>
<td>no</td>
<td>This one is highly unstable and the formula seems to be more like A &amp; X &amp; (imm | 6E)<br />
when the CPU is cold A=FF X=FF imm=00 result in 46, later 66 and then 6E (when the CPU is warm)<br />
bit 0 often flips from 0 to 1, for example A=01 X=01 imm=0C results in 00 or 01 (01 occurring more frequently when the CPU is warm)<br />
Also bit 3 flipping from 1 to 0 was observed with A=09 X=E5 and imm=05 or 41 (result: 00 instead of 08)<br />
also the Z flag is often incorrectly set to 1 when the result is non-zero. N flag seems to be OK.</td>
</tr>
<tr>
<td>Rockwell</td>
<td>SALLY</td>
<td>8328</td>
<td>?</td>
<td>C014806-12<br />
11151-12<br />
0579 8328</td>
<td>?</td>
<td>Atari 130XE</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8324</td>
<td>?</td>
<td>C014806-03<br />
8324</td>
<td>?</td>
<td>Atari 600XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8321</td>
<td>?</td>
<td>C014806-03<br />
8321</td>
<td>?</td>
<td>Atari 600XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>no</td>
<td>?</td>
<td>~95k errors (0.6%) in full test, sometimes bit 3 was set</td>
</tr>
<tr>
<td>Synertek</td>
<td>SALLY</td>
<td>8407</td>
<td>?</td>
<td>C014806-03<br />
8407</td>
<td>?</td>
<td>Atari 800XL</td>
<td>Hias</td>
<td>00</td>
<td>-</td>
<td>yes</td>
<td>?</td>
<td></td>
</tr>
<tr>
<td>Rockwell</td>
<td>?</td>
<td>8402</td>
<td>Mexico?</td>
<td>R6502AP<br />
R6502-13<br />
8407</td>
<td>?</td>
<td>BBC Model B</td>
<td>EdS</td>
<td>?</td>
<td>?</td>
<td>?</td>
<td>?</td>
<td>?</td>
</tr>
</tbody>
</table>

(\*)Note: "stable" means that the formula, the "magic" value and the potential \#4 clearing by RDY fully describe the behavior.  
(\*\*)Note: N and Z flags are set according to the result of XAA

## Resources {#Resources}

- For a list of all opcodes and some explanation of what they do, see [Visual6502wiki/6502 all 256 Opcodes](../6502_all_256_Opcodes/).
- For notes on other opcodes we've explored in our simulations, see [Visual6502wiki/6502 Unsupported Opcodes](../6502_Unsupported_Opcodes/).
- VICE notes: [64doc](http://www.commodore.ca/download/commodore/Vice_C64_Notes.txt) by John West and Marko Mäkelä
- 2004 [forum thread](http://plus4world.powweb.com/forum.php?postid=6775) on plus/4 world and [followup thread](http://plus4world.powweb.com/forum/10336#10511)
- 2006 [forum thread](http://noname.c64.org/csdb/forums/?roomid=11&topicid=30951&showallposts=1) on CSDb
- [This issue](http://sourceforge.net/tracker/?func=detail&aid=2110948&group_id=223021&atid=1057617) in the VICE bugtracker on sourceforge
- [This simulation](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=a955a2338b0feaea&steps=13&loglevel=2&logmore=adl,adh,idl,idb,alua,alub,alu,dasb,sb,dpc2_XSB,dpc9_DBADD,dpc11_SBADD,dpc23_SBAC,dpc24_ACSB,dpc25_SBDB,dpc43_DL/DB,DPControl) on visual6502 stops at the appropriate step and traces the appropriate busses and control signals
- [This simulation](http://visual6502.org/JSSim/expert.html?graphics=f&a=0&d=a955a2338b0feaea&steps=17&loglevel=2&logmore=ir,rdy,idl,idb,dasb,sb,dpc2_XSB,dpc23_SBAC,dpc24_ACSB,dpc25_SBDB,dpc43_DL/DB&rdy0=10&rdy1=14&time=12&databus=c4) stalls the 6502 using RDY to show the influence of the databus
