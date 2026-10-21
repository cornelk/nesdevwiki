---
title: "VRC2 pinout"
url: "/wiki/VRC2_pinout/"
description: "Konami VRC2 and VRC4: 0.6\" 40-pin PDIP (iNES Mappers 21, 22, 23, and 25)"
wikiPage: "VRC2_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Konami [VRC2 and VRC4](../VRC2_and_VRC4/): 0.6" 40-pin PDIP (iNES Mappers **21**, **22**, **23**, and **25**)

    f: connects to Famicom
    r: connects to ROM
    w: connects to WRAM

                          .--\/--.
           (f) CPU A13 -> |01  40| -- +5V
           (f) CPU A14 -> |02  39| -> PRG A17 (r)
           (fr) CPU Ax -> |03  38| -> PRG A15 (r)
           (fr) CPU Ay -> |04  37| <- CPU A12 (f)
           (f) PPU A12 -> |05  36| -> PRG A14 (r)
           (f) PPU A11 -> |06  35| -> PRG A13 (r)
           (f) PPU A10 -> |07  34| -> PRG A16 (r)
           (r) PRG /CE <- |08  33| <- CPU D0 (fr)
           (f) CPU R/W -> |09  32| <- CPU D1 (fr)
    (CHR /CE) (r) OR Y <- |10  31| <- CPU D2 (fr)
    (PPU A13) (f) OR B -> |11  30| <- CPU D4 (fr)
    (PPU /RD) (f) OR A -> |12  29| <- CPU D3 (fr)
         (f) CIRAM A10 <- |13  28| -> CHR A17 (r)
           (f) /ROMSEL -> |14  27| -> CHR A15 (r)
                (f) M2 -> |15  26| -> CHR A12 (r)
      VRC4 (r) CHR A18 <- |16  25| -> CHR A14 (r)
         VRC4 (f) /IRQ <- |17  24| -> CHR A13 (r)
          VRC4 /WR9003 <- |18  23| -> CHR A11 (r)
     VRC4 (w) WRAM /CE <- |19  22| -> CHR A16 (r)
                   GND -- |20  21| -> CHR A10 (r)
                          `------'

    pin 3, 4: see below
    pin 18: (VRC4): writing to $9003 pulls this pin low
    pin 16-19: in VRC2, these pins seem to have been intended for a never-seen-used EEPROM

# <span id="Pin_3.2C4_Wiring_Variants"></span>Pin 3,4 Wiring Variants {#Pin_3.2C4_Wiring_Variants}

Konami was greatly fond of making minor variations from one game to the next, presumably to make life harder on pirates. Note that no commercial games used **VRC4f** configuration.

## <span id="VRC4a.2C_PCB_352398_.28NES_2.0_Mapper_021.1.29"></span>VRC4a, PCB 352398 (NES 2.0 Mapper 021.1) {#VRC4a.2C_PCB_352398_.28NES_2.0_Mapper_021.1.29}

       (fr) CPU A2 -> |03  38| -> PRG A15 (r)
       (fr) CPU A1 -> |04  37| <- CPU A12 (f)

## <span id="VRC4b.2C_PCB_351406_.28NES_2.0_Mapper_025.1.29"></span>VRC4b, PCB 351406 (NES 2.0 Mapper 025.1) {#VRC4b.2C_PCB_351406_.28NES_2.0_Mapper_025.1.29}

       (fr) CPU A0 -> |03  38| -> PRG A15 (r)
       (fr) CPU A1 -> |04  37| <- CPU A12 (f)

## <span id="VRC4c.2C_PCB_352889_.28NES_2.0_Mapper_021.2.29"></span>VRC4c, PCB 352889 (NES 2.0 Mapper 021.2) {#VRC4c.2C_PCB_352889_.28NES_2.0_Mapper_021.2.29}

       (fr) CPU A7 -> |03  38| -> PRG A15 (r)
       (fr) CPU A6 -> |04  37| <- CPU A12 (f)

## <span id="VRC4d.2C_PCB_352400_.28NES_2.0_Mapper_025.2.29"></span>VRC4d, PCB 352400 (NES 2.0 Mapper 025.2) {#VRC4d.2C_PCB_352400_.28NES_2.0_Mapper_025.2.29}

       (fr) CPU A2 -> |03  38| -> PRG A15 (r)
       (fr) CPU A3 -> |04  37| <- CPU A12 (f)

## <span id="VRC4e.2C_PCB_352396_.28NES_2.0_Mapper_023.2.29"></span>VRC4e, PCB 352396 (NES 2.0 Mapper 023.2) {#VRC4e.2C_PCB_352396_.28NES_2.0_Mapper_023.2.29}

       (fr) CPU A3 -> |03  38| -> PRG A15 (r)
       (fr) CPU A2 -> |04  37| <- CPU A12 (f)

## <span id="VRC4f_.28NES_2.0_Mapper_023.1.29"></span>VRC4f (NES 2.0 Mapper 023.1) {#VRC4f_.28NES_2.0_Mapper_023.1.29}

       (fr) CPU A1 -> |03  38| -> PRG A15 (r)
       (fr) CPU A0 -> |04  37| <- CPU A12 (f)

## <span id="VRC2a.2C_PCB_351618.2C_.28INES_Mapper_022.29"></span>VRC2a, PCB 351618, (INES Mapper 022) {#VRC2a.2C_PCB_351618.2C_.28INES_Mapper_022.29}

       (fr) CPU A0 -> |03  38| -> PRG A15 (r)
       (fr) CPU A1 -> |04  37| <- CPU A12 (f)

## <span id="VRC2b.2C_PCBs_350603.2C_350636.2C_350926.2C_and_351179_.28NES_2.0_Mapper_023.3.29"></span>VRC2b, PCBs 350603, 350636, 350926, and 351179 (NES 2.0 Mapper 023.3) {#VRC2b.2C_PCBs_350603.2C_350636.2C_350926.2C_and_351179_.28NES_2.0_Mapper_023.3.29}

       (fr) CPU A1 -> |03  38| -> PRG A15 (r)
       (fr) CPU A0 -> |04  37| <- CPU A12 (f)

## <span id="VRC2c.2C_PCB_351948.2C_.28NES_2.0_Mapper_025.3.29"></span>VRC2c, PCB 351948, (NES 2.0 Mapper 025.3) {#VRC2c.2C_PCB_351948.2C_.28NES_2.0_Mapper_025.3.29}

       (fr) CPU A0 -> |03  38| -> PRG A15 (r)
       (fr) CPU A1 -> |04  37| <- 74'139 /Y2[1]

- The additional logic evaluates (/ROMSEL OR CPUA12), preventing bus conflicts between the Microwire interface and the PRG RAM on the board by making I/O to \$6000-\$6FFF instead appear to the VRC2 to be happening to \$7000-\$7FFF.

<!-- -->

     (f) CIRAM A10 <- |13  28| -> CHR A16 (r)
       (f) /ROMSEL -> |14  27| -> CHR A14 (r)
            (f) M2 -> |15  26| -> CHR A11 (r)
     VRC2 µWire DO -> |16  25| -> CHR A13 (r)
     VRC2 µWire DI <- |17  24| -> CHR A12 (r)
     VRC2 µWire SK <- |18  23| -> CHR A10 (r)
     VRC2 µWire CS <- |19  22| -> CHR A15 (r)
               GND -- |20  21| -> n/c

- The VRC2 µWire interface is thought to be nonfunctional.

# Pirate clone ICs {#Pirate_clone_ICs}

There exist [pirate versions](http://forums.nesdev.org/viewtopic.php?f=9&t=8569&start=15#p240820) of both VRC2 and VRC4. Pinouts are known to be the same as the genuine ICs. The source of IC pictures can be seen by clicking the image.

<table class="tabular sortable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: center;">IC</th>
<th style="text-align: center;">Clone of</th>
<th class="unsortable">Picture</th>
<th class="unsortable" style="text-align: center;">Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: center;">(nameless)</td>
<td style="text-align: center;">VRC2</td>
<td><a href="../../wiki-images/VRC2_nameless00.jpg"><img src="../../wiki-images/VRC2_nameless00.jpg" width="400" height="123" /></a></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td style="text-align: center;">23C269</td>
<td style="text-align: center;">VRC2</td>
<td><a href="../../wiki-images/23C269.jpg"><img src="../../wiki-images/23C269.jpg" width="400" height="137" /></a></td>
<td style="text-align: center;">µWire EEPROM interface <a href="https://forums.nesdev.org/viewtopic.php?p=211970#p211970">works</a>.</td>
</tr>
<tr>
<td style="text-align: center;">23C3662</td>
<td style="text-align: center;">VRC2</td>
<td><a href="../../wiki-images/23C3662.jpg"><img src="../../wiki-images/23C3662.jpg" width="400" height="118" /></a></td>
<td style="text-align: center;">Though known to be a VRC2 clone, the <a href="https://forums.nesdev.org/viewtopic.php?p=240820#p240820">mirroring register only exists at $9000</a> like VRC4.
<p>µWire EEPROM interface <a href="https://forums.nesdev.org/viewtopic.php?p=209096#p209096">works</a>.</p></td>
</tr>
<tr>
<td style="text-align: center;">AX-40G</td>
<td style="text-align: center;">VRC2</td>
<td><a href="../../wiki-images/AX-40G.png"><img src="../../wiki-images/AX-40G.png" width="400" height="132" /></a><br />
<a href="../../wiki-images/ASLIC_AX-40G.jpg"><img src="../../wiki-images/ASLIC_AX-40G.jpg" width="400" height="129" /></a></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td style="text-align: center;">AX5705</td>
<td style="text-align: center;">VRC2</td>
<td><a href="../../wiki-images/AX5705.jpg"><img src="../../wiki-images/AX5705.jpg" width="400" height="129" /></a></td>
<td style="text-align: center;">µWire EEPROM interface <a href="https://forums.nesdev.org/viewtopic.php?p=211970#p211970">works</a>. Relabeled version of 23C269?</td>
</tr>
<tr>
<td style="text-align: center;">(nameless)</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/VRC4_nameless00.jpg"><img src="../../wiki-images/VRC4_nameless00.jpg" width="400" height="129" /></a></td>
<td style="text-align: center;"><a href="../NES_2.0_Mapper_298/">$F003 (IRQ Acknowledge) does not move the 'A' control bit to the 'E' control bit.</a></td>
</tr>
<tr>
<td style="text-align: center;">23C3664</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/23C3664.png"><img src="../../wiki-images/23C3664.png" width="400" height="112" /></a></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td style="text-align: center;">AX5208C</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/AX5208C.jpg"><img src="../../wiki-images/AX5208C.jpg" width="400" height="123" /></a></td>
<td style="text-align: center;">WRAM is mapped at full $6000-$<strong>7</strong>FFF range, unlike the genuine VRC4 IC which only maps WRAM at $6000-$<strong>6</strong>FFF.
<p>/WR9003 pin is functional.</p></td>
</tr>
<tr>
<td style="text-align: center;">AX5208P</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/AX5208P.jpg"><img src="../../wiki-images/AX5208P.jpg" width="400" height="123" /></a></td>
<td style="text-align: center;">Known to work differently compared to AX5208"C". Probably regarding to WRAM. WRAM responds at full $6000-$<strong>7</strong>FFF.
<p>/WR9003 pin is functional.</p></td>
</tr>
<tr>
<td style="text-align: center;">PT8155</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/PT8155.jpg"><img src="../../wiki-images/PT8155.jpg" width="400" height="120" /></a></td>
<td style="text-align: center;">This specific clone IC seems to have PRG A17 and/or CHR A17 (and maybe CHR A18?) pins either broken, repurposed or not connected at all, because at least one <a href="http://krzysiobal.com/carts/?action=view&amp;id=183">PCB</a> with that chip (P-4073 PP-43KII, game <em>Wai Wai World 2</em>) contains additional logic to provide missing PRG A17 and CHR A17 lines.</td>
</tr>
<tr>
<td style="text-align: center;">V4</td>
<td style="text-align: center;">VRC4</td>
<td><a href="../../wiki-images/V4.jpg"><img src="../../wiki-images/V4.jpg" width="400" height="114" /></a></td>
<td style="text-align: center;"></td>
</tr>
</tbody>
</table>

Sources: [\[2\]](https://web.archive.org/web/20140920211213/http://nintendoallstars.w.interia.pl/romlab/vrcp.htm), [\[3\]](http://forums.nesdev.org/viewtopic.php?t=8569)
