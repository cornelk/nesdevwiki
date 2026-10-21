---
title: "CPU variants"
url: "/wiki/CPU_variants/"
description: "Beyond the well-studied 2A03G, we know of the following CPU revisions, both made by Ricoh and other manufacturers:"
wikiPage: "CPU_variants"
---

Beyond the well-studied 2A03G, we know of the following CPU revisions, both made by Ricoh and other manufacturers:

## <span id="Official_.28NTSC.29"></span>Official (NTSC) {#Official_.28NTSC.29}

All official NTSC CPUs use a ÷12 clock divider.

<table class="tabular sortable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>Part</th>
<th class="unsortable">Picture</th>
<th style="text-align: center;">First Seen</th>
<th style="text-align: center;">Last Seen</th>
<th class="unsortable">Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>RP2A03 (ceramic)</td>
<td><a href="../../wiki-images/RP2A03_3G1_09.jpg"><img src="../../wiki-images/RP2A03_3G1_09.jpg" width="400" height="131" /></a></td>
<td colspan="2" style="text-align: center;">1983-06<br />
3G1 09</td>
<td>Likely similar to standard plastic RP2A03.</td>
</tr>
<tr>
<td>RP2A03</td>
<td><a href="../../wiki-images/CPU=RP2A03_3L4_27.jpg"><img src="../../wiki-images/CPU=RP2A03_3L4_27.jpg" width="400" height="105" /></a> <a href="../../wiki-images/CPU=RP2A03_3M294.jpg"><img src="../../wiki-images/CPU=RP2A03_3M294.jpg" width="400" height="125" /></a></td>
<td style="text-align: center;">1983-07<br />
3G2 ?7</td>
<td style="text-align: center;">1984-09<br />
4J2 50</td>
<td>M2 duty cycle is 17/24 instead of 15/24 <a href="https://forums.nesdev.org/viewtopic.php?p=166761#p166761">[1]</a>. Lacks tonal noise mode. Lowest noise period is 2046 instead of 4068. <a href="../APU_Frame_Counter/">APU Frame Counter</a> not restarted on reset. Has broken and disabled <a href="../RP2A03_Programmable_Interval_Timer/">programmable interval timer</a> on-die. Pin 30 connects to nothing. Other differences?</td>
</tr>
<tr>
<td>RP2A03E</td>
<td><a href="../../wiki-images/CPU=RP2A03E_5K5_26.jpg"><img src="../../wiki-images/CPU=RP2A03E_5K5_26.jpg" width="400" height="121" /></a> <a href="../../wiki-images/CPU=RP2A03E_VF4109_5H4_8534.jpg"><img src="../../wiki-images/CPU=RP2A03E_VF4109_5H4_8534.jpg" width="400" height="125" /></a></td>
<td style="text-align: center;">1984-10<br />
4K1 11</td>
<td style="text-align: center;">1986-06<br />
6F1 23</td>
<td>Pin 30 is /RDY - combined with internal signals before feeding to internal 6502 +RDY.</td>
</tr>
<tr>
<td>RP2A03G</td>
<td><a href="../../wiki-images/CPU=RP2A03G_9B3_PV_A.jpg"><img src="../../wiki-images/CPU=RP2A03G_9B3_PV_A.jpg" width="400" height="122" /></a> <a href="../../wiki-images/CPU=RP2A03G_8815AAL.jpg"><img src="../../wiki-images/CPU=RP2A03G_8815AAL.jpg" width="400" height="119" /></a></td>
<td style="text-align: center;">1987-04<br />
7D3 A0</td>
<td style="text-align: center;">1993-11<br />
3LM 5B</td>
<td>Reference model. Pin 30 enables a <a href="../CPU_Test_Mode/">CPU test mode</a>. Later runs introduced a DMC DMA bug <a href="https://forums.nesdev.org/viewtopic.php?p=275359#p275359">[2]</a>.</td>
</tr>
<tr>
<td>RP2A03H</td>
<td><a href="../../wiki-images/CPU=RP2A03H_4AM_4F.jpg"><img src="../../wiki-images/CPU=RP2A03H_4AM_4F.jpg" width="400" height="114" /></a> <a href="../../wiki-images/CPU=RP2A03H_4BM_4Z.jpg"><img src="../../wiki-images/CPU=RP2A03H_4BM_4Z.jpg" width="400" height="120" /></a></td>
<td style="text-align: center;">1993-12<br />
3MM 40</td>
<td style="text-align: center;">1999-05<br />
9EM 5B</td>
<td>No known differences from late RP2A03G.</td>
</tr>
<tr>
<td>RP2A03H (laser)</td>
<td><a href="../../wiki-images/RP2A03H_2LL_4A.jpg"><img src="../../wiki-images/RP2A03H_2LL_4A.jpg" width="400" height="124" /></a></td>
<td style="text-align: center;">2001-03<br />
1CL 42</td>
<td style="text-align: center;">2002-11<br />
2LL 4A</td>
<td></td>
</tr>
<tr>
<td>RP2A04</td>
<td><a href="../../wiki-images/CPU=RP2A04_2J_6C2_01.jpg"><img src="../../wiki-images/CPU=RP2A04_2J_6C2_01.jpg" width="400" height="113" /></a></td>
<td colspan="2" style="text-align: center;">1986-03<br />
6C2 01</td>
<td>Not actually a CPU at all, just a jumper in a 40-pin PDIP. Used in place of CPUs in <a href="../Vs._System/">Vs. System</a> boards (and thus with NTSC timing).</td>
</tr>
</tbody>
</table>

## <span id="Official_.28PAL.29"></span>Official (PAL) {#Official_.28PAL.29}

All official PAL CPUs use a ÷16 clock divider and they have different [DMC period table](../APU_DMC/) than the NTSC CPUs.

<table class="tabular sortable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>Part</th>
<th class="unsortable">Picture</th>
<th style="text-align: center;">First Seen</th>
<th style="text-align: center;">Last Seen</th>
<th class="unsortable">Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>RP2A07</td>
<td><a href="../../wiki-images/CPU=RP2A07_7C4_39.jpg"><img src="../../wiki-images/CPU=RP2A07_7C4_39.jpg" width="400" height="125" /></a> <a href="../../wiki-images/CPU=RP2A07_8F3_78.jpg"><img src="../../wiki-images/CPU=RP2A07_8F3_78.jpg" width="400" height="131" /></a></td>
<td style="text-align: center;">1987-03<br />
7C4 39</td>
<td style="text-align: center;">1990-04<br />
0DL 2G</td>
<td>Input clock divider is 16. M2 duty cycle is 19/32 <a href="https://forums.nesdev.org/viewtopic.php?p=166761#p166761">[3]</a>. Changes to noise, DPCM, frame timer tables. Fixed DPCM RDY address bus glitches. Pin 30 connects to 6502 /RDY input.</td>
</tr>
<tr>
<td>RP2A07A</td>
<td><a href="../../wiki-images/CPU=RP2A07A_1GM_36.jpg"><img src="../../wiki-images/CPU=RP2A07A_1GM_36.jpg" width="400" height="127" /></a> <a href="../../wiki-images/CPU=RP2A07A_2JM_3L.jpg"><img src="../../wiki-images/CPU=RP2A07A_2JM_3L.jpg" width="400" height="103" /></a></td>
<td style="text-align: center;">1991-06<br />
1FM 3C</td>
<td style="text-align: center;">1992-10<br />
2KM 3L</td>
<td>There are no known differences relative to 2A07letterless.</td>
</tr>
</tbody>
</table>

## Unofficial {#Unofficial}

There are many unofficial, or "clone" CPUs. Though they generally work quite well for what they are, there are many associated quirks and quality control issues. Keep these items in mind:

- Some clones have reversed duty cycles on the pulse channels, which causes the tone to sound different, but the pitch to remain correct.
- All known clone CPUs use the NTSC [DMC period table](../APU_DMC/), including those intended for PAL systems.
- The clock divider may be 12 (like official NTSC), 16 (like official PAL), or 15 (unique to clones) and this can be tested by dividing the master clock frequency by the observed M2 frequency.
- If you use the incorrect clock divider, the pitch of the sound and the speed of the gameplay will be affected. See [this](https://forums.nesdev.org/viewtopic.php?t=25679) discussion adapting to a different clock divider.

<table class="tabular">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th class="unsortable">Part</th>
<th class="unsortable">Clk<br />
Div</th>
<th class="unsortable">Picture</th>
<th class="unsortable">Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td>MG-N-501</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=MG-N-501_8933.jpg"><img src="../../wiki-images/CPU=MG-N-501_8933.jpg" width="400" height="119" /></a></td>
<td></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=154574#p154574">MG-P-501</a></td>
<td>?</td>
<td><a href="../../wiki-images/CPU=MG-P-501_9221S_415521.jpg"><img src="../../wiki-images/CPU=MG-P-501_9221S_415521.jpg" width="400" height="106" /></a></td>
<td>Micro Genius-made clone. Die has the same (UMC) © Ⓜ B6167F marking as a UA6527P.</td>
</tr>
<tr>
<td>UA6527</td>
<td class="table-GreenCell" style="text-align: center; background: #9FB; vertical-align: middle;">12</td>
<td><a href="../../wiki-images/CPU=UA6527-8909-BS.jpg"><img src="../../wiki-images/CPU=UA6527-8909-BS.jpg" width="400" height="133" /></a> <a href="../../wiki-images/CPU=UA6527-9310-CG-C12520.jpg"><img src="../../wiki-images/CPU=UA6527-9310-CG-C12520.jpg" width="400" height="122" /></a></td>
<td>UMC-made clone of 2A03G. Has swapped pulse channel duty cycles. Input clock Divider is 12.</td>
</tr>
<tr>
<td rowspan="4">UA6527P</td>
<td colspan="3">UMC-made clone of 2A03G for compatibility with NTSC software in PAL countries. Different input clock divider. Still has swapped pulse channel duty cycles. Otherwise believed same as 6527.
<p>One revision has (UMC) © Ⓜ B6167F 1989 09 on the die.</p>
<p>DMC status bit is cleared 1 APU cycle later than on RP2A03 CPUs. The cause is not known. This changes the timing for <a href="../DMA/#Bugs">DMC DMA implicit-stop glitches</a> (the sample must be started 1 APU cycle earlier to trigger the glitches), and it is suspected that it delays DMC IRQ by 1 APU cycle. Noise channel is slightly louder than others.</p></td>
</tr>
<tr>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=UA6527P_8931S.jpg"><img src="../../wiki-images/CPU=UA6527P_8931S.jpg" width="400" height="122" /></a></td>
<td>Runs hot. Revisions without "-" in the date stamp have a ÷16 CPU divider, like 6540 and 2A07</td>
</tr>
<tr>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=UA6527P_9019-BS.jpg"><img src="../../wiki-images/CPU=UA6527P_9019-BS.jpg" width="400" height="123" /></a></td>
<td>Runs hot. Revisions with "-" in the date stamp and text on the bottom may have a ÷15 CPU divider.</td>
</tr>
<tr>
<td>?</td>
<td><a href="../../wiki-images/CPU=UA6527P_9214-BS_310551.jpg"><img src="../../wiki-images/CPU=UA6527P_9214-BS_310551.jpg" width="400" height="123" /></a></td>
<td>Runs cooler</td>
</tr>
<tr>
<td rowspan="7">UA6527P<br />
(Relabeled)</td>
<td colspan="3">Note that UA6527P chips are notorious for being other chips that were relabeled. In some cases, the chip has been sanded and relabeled, but it may still be possible to figure out what it actually is based on markings on the bottom of the chip. In other cases, the chip was painted and a new label was printed on top of that. If that's the case, the original label may be revealed by removing the paint.</td>
</tr>
<tr>
<td></td>
<td><a href="../../wiki-images/CPU=UA6527P_9919-BS.jpg"><img src="../../wiki-images/CPU=UA6527P_9919-BS.jpg" width="400" height="132" /></a></td>
<td>Example top of a sanded, painted, and relabeled CPU. Note the level of freshness on the top side of the chip versus the generally filthy bottom side. Extrusion marks are not as deep as normal, or may be sanded away entirely.</td>
</tr>
<tr>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=UA6527P_9172AP_C2259_bottom_div15.jpg"><img src="../../wiki-images/CPU=UA6527P_9172AP_C2259_bottom_div15.jpg" width="400" height="137" /></a></td>
<td>This CPU has a ÷15 clock divider and correct duty cycles. These are the same bottom markings as a TA03-NP1. Pin 30 function is +TST.</td>
</tr>
<tr>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=UA6527P_TC23_bottom.jpg"><img src="../../wiki-images/CPU=UA6527P_TC23_bottom.jpg" width="400" height="141" /></a></td>
<td>This CPU has a ÷15 clock divider and reverse duty cycles. Pin 30 function is /RDY.</td>
</tr>
<tr>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=UA6527P_bottom_8th_subject.jpg"><img src="../../wiki-images/CPU=UA6527P_bottom_8th_subject.jpg" width="400" height="130" /></a></td>
<td>This CPU has a ÷15 clock divider and reverse duty cycles. Pin 30 function is /RDY.</td>
</tr>
<tr>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=UA6527P_blank_bottom_div16.jpg"><img src="../../wiki-images/CPU=UA6527P_blank_bottom_div16.jpg" width="400" height="139" /></a></td>
<td>This CPU has a ÷16 clock divider and reverse duty cycles.</td>
</tr>
<tr>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=UA6527P_bottom_7th_subject.jpg"><img src="../../wiki-images/CPU=UA6527P_bottom_7th_subject.jpg" width="400" height="134" /></a></td>
<td>This CPU has a ÷16 clock divider and reverse duty cycles. Pin 30 function is /RDY.</td>
</tr>
<tr>
<td>UA6527PQ</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=UA6527PQ_9306.JPG"><img src="../../wiki-images/CPU=UA6527PQ_9306.JPG" width="400" height="122" /></a></td>
<td></td>
</tr>
<tr>
<td>UA6540</td>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=UA6540_8833S.JPG"><img src="../../wiki-images/CPU=UA6540_8833S.JPG" width="400" height="125" /></a> <a href="../../wiki-images/CPU=UA6540-8834S.jpg"><img src="../../wiki-images/CPU=UA6540-8834S.jpg" width="400" height="126" /></a></td>
<td>UMC-made clone of 2A07 <a href="https://forums.nesdev.org/viewtopic.php?t=17257">[4]</a>. Has swapped pulse duty cycles.
<p>Subsequent research implies this is identical to the early 6527P - NTSC tuning tables, ÷16 CPU divider. <a href="https://forums.nesdev.org/viewtopic.php?p=283514#p283514">[5]</a></p></td>
</tr>
<tr>
<td>UA6547</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=UA6547_8950S.JPG"><img src="../../wiki-images/CPU=UA6547_8950S.JPG" width="400" height="136" /></a></td>
<td>Believed to be a 100% duplicate of UA6527, for use in PAL-M region.</td>
</tr>
<tr>
<td>UM6557</td>
<td>?</td>
<td><a href="../../wiki-images/UM6557.JPG"><img src="../../wiki-images/UM6557.JPG" width="400" height="151" /></a></td>
<td>Believed to be a 100% duplicate of UA6527, for use in SECAM regions.</td>
</tr>
<tr>
<td>UM6561xx-1</td>
<td class="table-GreenCell" style="text-align: center; background: #9FB; vertical-align: middle;">12</td>
<td><a href="../../wiki-images/NOAC=UM6561CF-1_0028A_E6276.jpg"><img src="../../wiki-images/NOAC=UM6561CF-1_0028A_E6276.jpg" width="200" height="153" /></a></td>
<td>NES-on-a-chip for NTSC. Revisions "xx" F, AF, BF, CF known.</td>
</tr>
<tr>
<td>UM6561xx-2</td>
<td>?</td>
<td><a href="../../wiki-images/UM6561F-2.jpg"><img src="../../wiki-images/UM6561F-2.jpg" width="248" height="203" /></a> <a href="../../wiki-images/NOAC=UM6561_AF-2_9440A_R81040.jpg"><img src="../../wiki-images/NOAC=UM6561_AF-2_9440A_R81040.jpg" width="200" height="147" /></a></td>
<td>NES-on-a-chip for PAL-B. Revisions "xx" F, AF, BF, CF known.
<p>F and AF revision pulse wave duty cycles match RP2A03, and DMC status bit is cleared 1 APU cycle later than on RP2A03 CPUs.</p>
<p>AF revision observed to have incorrect ASR #imm ($4B) behavior, but other stable illegal instructions work properly.</p></td>
</tr>
<tr>
<td>UM6569F-1</td>
<td>?</td>
<td><a href="../../wiki-images/NOAC=UM6569F-1_9241S_807680.jpg"><img src="../../wiki-images/NOAC=UM6569F-1_9241S_807680.jpg" width="200" height="155" /></a></td>
<td>NES-on-a-chip for PAL-N (Argentina). Uses external RAMs. Found on a PCB marked UM6567, implying that one's existence too.</td>
</tr>
<tr>
<td>1818N</td>
<td>?</td>
<td><a href="../../wiki-images/CPU_=_1818N_9218.png"><img src="../../wiki-images/CPU_=_1818N_9218.png" width="200" height="140" /></a></td>
<td>??-made NES-on-a-chip, NTSC timing.</td>
</tr>
<tr>
<td>T1818P</td>
<td>?</td>
<td><a href="../../wiki-images/1818P_0.jpeg"><img src="../../wiki-images/1818P_0.jpeg" width="200" height="133" /></a></td>
<td>??-made NES-on-a-chip[<a href="https://forums.nesdev.org/viewtopic.php?p=228515#p228515">[6]</a>. Requires external 2 KiB RAMs for CPU and PPU. Swapped pulse duty cycles. DMC status bit is cleared 1 APU cycle later than on RP2A03 CPUs.</td>
</tr>
<tr>
<td>TA-03N</td>
<td class="table-GreenCell" style="text-align: center; background: #9FB; vertical-align: middle;">12</td>
<td><a href="../../wiki-images/CPU=TA-03N_6527_9250.jpg"><img src="../../wiki-images/CPU=TA-03N_6527_9250.jpg" width="400" height="140" /></a>
<p><a href="../../wiki-images/CPU=TA-03N_9172N_12450820.jpg"><img src="../../wiki-images/CPU=TA-03N_9172N_12450820.jpg" width="400" height="139" /></a></p></td>
<td>??-made die-mask clone of 2A03G. Chip underside also has two codes of currently unknown purpose. Pin 30 activates CPU Test Mode like on 2A03G. Clock Divisor is 12. Illegal opcodes are the same. Early 1991 dated chips are reported to have problems with APU DMC playback, but this was corrected in 1992 onward.</td>
</tr>
<tr>
<td>TA-03NP</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">15<br />
or<br />
12</td>
<td><a href="../../wiki-images/CPU=TA-03NP_EWP0124.jpg"><img src="../../wiki-images/CPU=TA-03NP_EWP0124.jpg" width="400" height="133" /></a></td>
<td>??-made clone of 2A03G for NTSC compatibility in PAL countries. Input clock divider is 15?
<p>But not this one, this input clock divider is 12.</p></td>
</tr>
<tr>
<td>TA-03NP1</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=TA-03NP1_6527P_9231.jpg"><img src="../../wiki-images/CPU=TA-03NP1_6527P_9231.jpg" width="400" height="118" /></a>
<p><a href="../../wiki-images/CPU=TA-03NP1_6527P_9217.jpg"><img src="../../wiki-images/CPU=TA-03NP1_6527P_9217.jpg" width="400" height="125" /></a> <a href="../../wiki-images/CPU=TA-03NP1_9172P_600670.jpg"><img src="../../wiki-images/CPU=TA-03NP1_9172P_600670.jpg" width="400" height="143" /></a></p></td>
<td>??-made clone of 2A03G for NTSC compatibility in PAL countries. Input clock divider is 15. Fixed DPCM problems? Correct pulse channel duties. Noise channel is slightly louder than others. DMC status bit is cleared 1 APU cycle later than on RP2A03 CPUs.</td>
</tr>
<tr>
<td>PM03</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=PM03_HI25.jpg"><img src="../../wiki-images/CPU=PM03_HI25.jpg" width="400" height="123" /></a></td>
<td><a href="https://en.wikipedia.org/wiki/IGB_Eletr%C3%B4nica">Gradiente</a>-made clone of 2A03G. <a href="https://forums.nesdev.org/viewtopic.php?p=195175#p195175">[7]</a></td>
</tr>
<tr>
<td>GS87007</td>
<td class="table-GreenCell" style="text-align: center; background: #9FB; vertical-align: middle;">12</td>
<td><a href="../../wiki-images/CPU=GS87007_8827.jpg"><img src="../../wiki-images/CPU=GS87007_8827.jpg" width="400" height="127" /></a></td>
<td>(Goldstar??)-made clone of 2A03 - has functioning decimal mode? <a href="http://gxemu.blog67.fc2.com/blog-entry-363.html">[8]</a>
<p>Found in MicroGenius clone with PAL/NTSC switch.<sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup></p></td>
</tr>
<tr>
<td>GS87007P</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=GS87007P_8815.jpg"><img src="../../wiki-images/CPU=GS87007P_8815.jpg" width="400" height="127" /></a></td>
<td>Contains functional decimal mode, APU pulse duty cycle is swapped. <a href="https://forums.nesdev.org/viewtopic.php?p=307926#p307926">[9]</a></td>
</tr>
<tr>
<td>KC-6005</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=KC-6005.jpg"><img src="../../wiki-images/CPU=KC-6005.jpg" width="400" height="132" /></a></td>
<td>Found in MT777-DX famiclone, behaves exactly like UA6527P</td>
</tr>
<tr>
<td>6005B</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=6005B.jpg"><img src="../../wiki-images/CPU=6005B.jpg" width="400" height="145" /></a></td>
<td></td>
</tr>
<tr>
<td>2011</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=2011.jpg"><img src="../../wiki-images/CPU=2011.jpg" width="400" height="127" /></a></td>
<td></td>
</tr>
<tr>
<td>“2A03E”</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=USC_2A03E_9118S_314531.jpg"><img src="../../wiki-images/CPU=USC_2A03E_9118S_314531.jpg" width="400" height="111" /></a> <a href="../../wiki-images/CPU=2A03E_9122A_422951.jpg"><img src="../../wiki-images/CPU=2A03E_9122A_422951.jpg" width="100" height="90" /></a></td>
<td>Both with and without USC insignia</td>
</tr>
<tr>
<td>KP2B03E</td>
<td>?</td>
<td><a href="../../wiki-images/CPU=KP2B03E_DHG_44.jpg"><img src="../../wiki-images/CPU=KP2B03E_DHG_44.jpg" width="400" height="128" /></a></td>
<td></td>
</tr>
<tr>
<td>6527-21 P03</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=6527-21_P03_91240_5009.png"><img src="../../wiki-images/CPU=6527-21_P03_91240_5009.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is RDY.</td>
</tr>
<tr>
<td>6527</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=6527_9132E.png"><img src="../../wiki-images/CPU=6527_9132E.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is RDY.</td>
</tr>
<tr>
<td rowspan="2">6527P</td>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=6527P_CFE1_1PC.png"><img src="../../wiki-images/CPU=6527P_CFE1_1PC.png" width="400" height="134" /></a></td>
<td>Clock divider is /16. Pulse channel duty cycles are swapped. Pin 30 is RDY.</td>
</tr>
<tr>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=6527P_9040.jpg"><img src="../../wiki-images/CPU=6527P_9040.jpg" width="400" height="119" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is TEST.</td>
</tr>
<tr>
<td>HA6527P</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=HA6527P_CRI2.png"><img src="../../wiki-images/CPU=HA6527P_CRI2.png" width="400" height="118" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are swapped. Pin 30 is RDY.</td>
</tr>
<tr>
<td>SENITON 6527P-SS-P03</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=SENITON_6527P-SS-P03_2.png"><img src="../../wiki-images/CPU=SENITON_6527P-SS-P03_2.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is TEST.</td>
</tr>
<tr>
<td>SENITON 6527UP-8</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=SENiTON_6527UP-8.png"><img src="../../wiki-images/CPU=SENiTON_6527UP-8.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are swapped. Pin 30 is RDY.</td>
</tr>
<tr>
<td>SENITON 6527AP</td>
<td class="table-BlueCell" style="text-align: center; background: #8CF; vertical-align: middle;">16</td>
<td><a href="../../wiki-images/CPU=SENITON_6527AP_9109C.png"><img src="../../wiki-images/CPU=SENITON_6527AP_9109C.png" width="400" height="133" /></a></td>
<td>Clock divider is /16. Pulse channel duty cycles are swapped. Pin 30 is RDY.</td>
</tr>
<tr>
<td>SL/WH6527AP</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=SL-WH6527AP_9120.png"><img src="../../wiki-images/CPU=SL-WH6527AP_9120.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are swapped. Pin 30 is RDY.</td>
</tr>
<tr>
<td>SNC6527P</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=SNC6527P_9104.png"><img src="../../wiki-images/CPU=SNC6527P_9104.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is TEST.</td>
</tr>
<tr>
<td>XYZ-6783</td>
<td>?</td>
<td><a href="../../wiki-images/CPU_=_XYZ-6783_9129.jpg"><img src="../../wiki-images/CPU_=_XYZ-6783_9129.jpg" width="400" height="141" /></a></td>
<td>Lacks tonal noise mode like original RP2A03, but resets APU Frame Counter on console reset like 2A03E/2A03G. Otherwise behaves like letterless RP2A03.</td>
</tr>
<tr>
<td>6538N</td>
<td>?</td>
<td><a href="../../wiki-images/CPU_=_6538N_CW6B_3NQ.jpg"><img src="../../wiki-images/CPU_=_6538N_CW6B_3NQ.jpg" width="400" height="139" /></a></td>
<td>??-made CPU, despite the part number being similar to UMC PPU. Has inverted duty cycles like UA6527. DPCM works.</td>
</tr>
<tr>
<td>8Z01N</td>
<td class="table-GreenCell" style="text-align: center; background: #9FB; vertical-align: middle;">12</td>
<td><a href="../../wiki-images/8Z01N_CPU.jpg"><img src="../../wiki-images/8Z01N_CPU.jpg" width="400" height="140" /></a></td>
<td>Found in Family Game by NTDEC.</td>
</tr>
<tr>
<td>TECH 27</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=TECH_27.png"><img src="../../wiki-images/CPU=TECH_27.png" width="400" height="121" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is TEST.</td>
</tr>
<tr>
<td>HITEX-6527P-P03 GX</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=HITEX-6527P-P03_GX_9108.png"><img src="../../wiki-images/CPU=HITEX-6527P-P03_GX_9108.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is TEST.</td>
</tr>
<tr>
<td>WDL6527P</td>
<td class="table-RedCell" style="text-align: center; background: #FA9; vertical-align: middle;">15</td>
<td><a href="../../wiki-images/CPU=WDL6527P_9213.png"><img src="../../wiki-images/CPU=WDL6527P_9213.png" width="400" height="133" /></a></td>
<td>Clock divider is /15. Pulse channel duty cycles are correct. Pin 30 is RDY.</td>
</tr>
</tbody>
</table>

If you know of other differences or other revisions, please add them!

## See also {#See_also}

- [PPU variants](../PPU_variants/)
- <https://forums.nesdev.org/viewtopic.php?p=45889#p45889>
- <https://forums.nesdev.org/viewtopic.php?t=23916> (More Info on CPU Clones)
- <https://forums.nesdev.org/viewtopic.php?t=23682> (Lots of Images and die-shots)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) <https://forums.nesdev.org/viewtopic.php?p=241514#p241514>
