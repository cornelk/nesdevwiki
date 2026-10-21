---
title: "Visual circuit tutorial"
url: "/wiki/Visual_circuit_tutorial/"
description: "This is a crash course on making sense of the NMOS circuit displays in Visual 6502/2C02/2A03, written for people without much low-level electronics experience (like the primary author). It aims to…"
wikiPage: "Visual_circuit_tutorial"
---

This is a crash course on making sense of the [NMOS](#Terms) circuit displays in [Visual 6502](http://visual6502.org/)/2C02/2A03, written for people without much low-level electronics experience (like the primary author). It aims to present the information needed to read the diagrams at a basic level in simple language, omitting details that are less important when starting out.

You might want to read [the Visual 6502 user's guide](../Visual6502wiki/JssimUserHelp/) and the [Visual 2C02](../Visual_2C02/) page first.

## What the different colored areas are {#What_the_different_colored_areas_are}

Let's start by defining what the different colors mean:

[<img src="../../wiki-images/Vis_areas.png" width="637" height="530" alt="Vis areas.png" />](../../wiki-images/Vis_areas.png)

- Green areas are diffusion connected to ground.
- Red areas are diffusion connected to VCC (power).
- Yellow areas are diffusion that is neither connected directly to ground nor directly to VCC.
- Gray areas are metal.
- Purple areas are polysilicon (often shortened to just "poly").

At the level presented here, diffusion, metal, and polysilicon can be thought of as roughly equivalent when viewed in isolation; they all conduct current. The important difference is in how they interact with each other, which is explained below.

## Basic building blocks {#Basic_building_blocks}

### Transistors {#Transistors}

When a piece of polysilicon is sandwiched between two areas of diffusion, it acts as a *gate*, only letting current through when the polysilicon is connected to power (or, equivalently, is *on*, *conducting*, *high*, or *1*). The diffusion area from which current flows when the gate is high is called the *source*. The diffusion area into which current flows is called the *drain*. The gate together with the source and drain is what makes a [*transistor*](https://en.wikipedia.org/wiki/Field-effect_transistor).

[<img src="../../wiki-images/Vis_transistor.png" width="402" height="190" alt="Vis transistor.png" />](../../wiki-images/Vis_transistor.png)

The transistor here is an [enhancement-mode transistor](https://en.wikipedia.org/wiki/Depletion_and_enhancement_modes). All the "ordinary" selectable (see the [nodes section](#Nodes)) transistors have this type.

### Power sources {#Power_sources}

Around areas of powered diffusion we often see something like the following (note the distinctive "hook" in the polysilicon):

[<img src="../../wiki-images/Vis_power.png" width="362" height="195" alt="Vis power.png" />](../../wiki-images/Vis_power.png)

Here the polysilicon acts roughly like a resistor (or more specifically a [pull-up resistor](#Terms)). This prevents there from ever being a short from VCC to ground (through some path of high gates). In the simulators, this entire configuration is simply modeled as a power source.

The transistor here is a [depletion-mode transistor](https://en.wikipedia.org/wiki/Depletion_and_enhancement_modes), a different type of transistor compared to above (though it appears the same visually).

### An example that brings together transistors and power sources {#An_example_that_brings_together_transistors_and_power_sources}

In the following curcuit, the gate will be high and *A* powered/pulled to VCC:

                          VCC
                           |
                           |
                           |
    [power source]------[gate]
                           |
                           |
                           |
                           A

No current will ever flow from the power source to *A* (ideally). The voltage on the gate only controls whether there's a conductive path between VCC and *A*.

Similarly, the gate will be low and *A* *not* powered/pulled to VCC In the following circuit:

                          VCC
                           |
                           |
                           |
    GND-----------------[gate]
                           |
                           |
                           |
                           A

Using a switch analogy for transistors (which is common in digital electronics), the above circuits can be viewed as follows:

                                               A
                                               |
                                               |
                                               |
                                               \
    B----------[remote switch controller]       \   <- controlled switch
                                                 \
                                               |
                                               |
                                               |
                                               C

When *B* is high, the switch is closed, connecting *A* and *C*. When *B* is low, it is open (like in the figure).

## Nodes {#Nodes}

Electrically common areas are called *nodes* in Visual 6502/2C02/2A03. Clicking on a node highlights it, making it easier to see how things are connected. (Clicking on powered or grounded diffusion won't work; these only modify properties of other nodes and are not themselves nodes.) When a node is highlighted, a numeric ID unique to the node is displayed in the upper right, along with a name for the node if it has one. Node names are defined in **nodenames.js**.

Transistors can be selected separately by clicking on the gate (the part of the polysilicon between the diffusion areas). They have names that start with 't', followed by a numeric ID.

The **Find:** edit field can be used to locate nodes, either by numeric ID or by name. Numeric IDs can also be used to trace the values of nodes without an assigned name.

## Basic logic elements {#Basic_logic_elements}

### Inverters {#Inverters}

An inverter is constructed like in the image below:

[<img src="../../wiki-images/Vis_inverter.png" width="376" height="238" alt="Vis inverter.png" />](../../wiki-images/Vis_inverter.png)

Note that there is a hook in the gate to the left, meaning the left part of the circuit is a [power source](#Power_sources) instead of a "normal" transistor.

When the input gate is low, current from the power source flows into the output wire, pulling the voltage high. When the input gate is high, current from the power source instead flows into ground, driving the voltage on the output wire low. The output wire is hence the inverse of the input wire.

When one node is the inverse of another, we will say that it *inverts into* the other node.

### NOR gates {#NOR_gates}

Below is an example of a NOR gate taken from Visual 2A03, related to controlling when the first square channel is silenced:

[<img src="../../wiki-images/Vis_nor.png" width="176" height="324" alt="Vis nor.png" />](../../wiki-images/Vis_nor.png)

If any of the gates in red circles are high, the voltage of the highlighted node is pulled to ground instead of pulled high (as current will flow from the power source on the left into ground through any high gates). The value that reaches the gate in the blue circle is hence the NOR of the values on the gates in the red circles.

Note that the circles represent the only transistors in this image (except for the depletion-mode transistors on the power sources). There are polysilicon traces passing above (or in reality, below) metal traces in a few spots, but this does not form a transistor. The highlighting (which was activated by clicking on the node) shows how things are connected.

The gate in the blue circle is part of a *pass transistor*, so called because it passes current between two nodes rather than driving or grounding a node. The gate in this case is [**apu_clk1**](#APU_clock_signals), and we say that value is "buffered on **apu_clk1**".

## Storage elements {#Storage_elements}

### Wire capacitance as storage {#Wire_capacitance_as_storage}

This is the simplest form of storage, and so is covered first.

If a wire is "closed off" so that it is no longer connected to neither power nor ground, it retains its value for a while through capacitance. This is used to store some short-lived data "on the wire". As an example, here's the read buffer for the 2C02's **VBlank** flag, which lets its value be read even though reading [\$2002](../PPU_registers/) immediately clears the **VBlank** flag:

[<img src="../../wiki-images/Vis_vblbuf.png" width="119" height="294" alt="Vis vblbuf.png" />](../../wiki-images/Vis_vblbuf.png)

The circled gate is controlled by the **/read_2002_output_vblank_flag** signal (shortened to **ov** from here on). While **ov** is high, the value of **vbl_flag** (or rather **/vbl_flag** in this case) is connected to the highlighted wire. When **ov** goes low, the value on the wire is held.

A '/' denotes 'inverse', meaning the signal is the inverse of another signal with the same name but without the '/'. A '/' can also mean 'active low', meaning the signal is considered "active" when low. Visual 6502 has a slightly different convention – see [Node names in Visual 6502](#Node_names_in_Visual_6502).

While a node or wire is isolated from both VCC and ground in the above fashion, it is said to be *floating*. For bus lines, a floating line is said to be *tri-stated*, as the floating state can be viewed as a third state in addition to 0 and 1. This third state allows other devices to use the bus without interference.

Using capacitance as storage in the above fashion is an instance of [dynamic logic](https://en.wikipedia.org/wiki/Dynamic_logic_(digital_electronics)), so called since it has time-dependent behavior beyond just the input clock. Chips that make use of dynamic logic techniques tend to have a minimum clock speed at which they function correctly, as values stored via capacitance degrade to zero over time.

### <span id="Latches_.28cross-coupled_inverters.29"></span>Latches (cross-coupled inverters) {#Latches_.28cross-coupled_inverters.29}

Two cross-coupled inverters make a *latch* – an element that stores a single bit.

Below is the VBlank flag from Visual 2C02. In the left-most picture the **vbl_flag** node is highlighted, and in the middle picture its inverse (**/vbl_flag**) is highlighted. As can be seen by the two gates in gray circles, each node inverts into the other, forming two cross-coupled inverters.

[<img src="../../wiki-images/Vis_crossreg.png" width="361" height="400" alt="Vis crossreg.png" />](../../wiki-images/Vis_crossreg.png)

The gates marked *set* and *clear* set and clear the latch, respectively. To clear the latch, **vbl_flag** is driven low. To set the latch, **/vbl_flag** is driven low.

This circuit is an example of an *[SR Latch](https://en.wikipedia.org/wiki/Flip-flop_(electronics)#SR_NOR_latch)*, where *S* stands for *set* and *R* for *reset*, corresponding to the *set* and *clear* gates above. It is more specifically an SR NOR Latch, as it can be viewed as being built of NOR gates. The corresponding schematic using NOR gates is shown in the right-most picture.

### Clocked latches {#Clocked_latches}

When a latch can be set directly from the value of some line, e.g. a data bus line, an arrangement involving a clock is often used. The motivation is to avoid having to form both **data_line** and **/data_line** and route them to the *set* and *clear* terminals of the latch, which would use more logic. The clock is already routed all around the chip, so mixing it in usually isn't as much of a problem.

As an example, here's the **noi_lfsrmode** node (the *Loop noise* flag from [\$400E](../APU_Noise/)):

[<img src="../../wiki-images/Vis_clockedreg.png" width="800" height="430" alt="Vis clockedreg.png" />](../../wiki-images/Vis_clockedreg.png)

While [**apu_clk1**](#APU_clock_signals) is high, **noi_lfsrmode** will flow into the floating node (so called because it will [float](#Wire_capacitance_as_storage) when both **apu_clk1** and **w400e** are low), which then inverts into **noi\_/lfsrmode**, forming a cross-coupled inverter latch. While **apu_clk1** is low, the loop will be broken momentarily, and during this phase a new value can be copied into the latch through the gate controlled by the **w400e** signal (which goes high on writes to \$400E). The value let through by the pass transistor is the **db7** node, corresponding to the 8th bit of the data bus. (There's a [via](#Terms) between the diffusion and the metal **db7** line – easier to see if the node is highlighted.) If the loop was not broken during the write operation, the old value in the latch would interfere with setting a new value.

For another, less cluttered view of the same type of circuit, see [this image](http://www.ece.unm.edu/~jimp/vlsi/slides/chap5_2-21.gif) (substitute "**apu_clk1**" for **"/φ₁"** and "**w400e**" for **"φ₁"**).

(The circuitry in the lower-right corner is a *multiplexer*, which selects between one of two inputs depending on whether **noi_lfsrmode** or **noi\_/lfsrmode** is high; i.e., depending on whether **noi_lfsrmode** is 0 or 1. The output of the multiplexer is on the left side.)

### <span id="DRAM_.28Dynamic_RAM.29"></span>DRAM (Dynamic RAM) {#DRAM_.28Dynamic_RAM.29}

Below is an example of a DRAM cell, taken from the internal PPU [OAM](../PPU_OAM/) memory:

[<img src="../../wiki-images/Vis_dram_cell.png" width="230" height="112" alt="Vis dram cell.png" />](../../wiki-images/Vis_dram_cell.png)

In the left and right pictures the two sides of the cell are highlighted (with a different highlight color on the right due to the node being high). The two nodes are always inverses of each other, with the node highlighted in the left picture corresponding to the value held in the cell (low for 0 and high for 1).

Note that this is *not* an instance of cross-coupled inverters, as neither node is directly connected to a power source. Rather, DRAM depends on capacitance to hold the value, which will fade unless the capacitor is regularly *refreshed* (the high side recharged). This is the "dynamic" part of DRAM.

Below is a picture of the upper edge of the PPU OAM DRAM array:

[<img src="../../wiki-images/Vis_oam.png" width="777" height="278" alt="Vis oam.png" />](../../wiki-images/Vis_oam.png)

The "column" and "row" labels are conventional memory terminology; they confusingly happen to get the opposite orientation in Visual 2C02. "Row" and "column" below will refer to this terminology.

The **spr_row*x*** lines (sometimes called *word lines*) are used to connect a row of memory cells to the horizontal *bit lines* (by opening up each cell to a pair of [vias](#Terms)); this is called *opening* that row. For example, **spr_row16** opens the highlighted row, while **spr_row0** opens the row on its right side. As can be guessed from the node names, the memory layout is not as straightforward as consecutive memory locations being stored in consecutive rows. (Interestingly, we do get consecutive rows if we reverse the bits in the part of the sprite address that selects the row. It is unknown why the row selection bits were not wired to the DRAM in this "correct" configuration instead.)

On the left side of OAM we see pass transistors on the **spr_col1** and **spr_col3** lines select the bit lines from the first and second columns of the memory array, respectively (there are other, similar, lines next to them) . Each such **spr_col*x*** line is connected to eight different columns (16 bit lines), corresponding to the eight bits of the byte to be read or written (increasing bit positions are not stored in consecutive columns either). One notable exception to this pattern is that two columns only connect to **five** sets of bit lines; these columns correspond to the "flags" bytes in OAM, where the middle 3 bits don't actually exist.

#### DRAM refresh {#DRAM_refresh}

At the right side in the picture above we see [**pclk0**](#PPU_clock_signals) running down the edge of OAM, connected to [pull-up transistors](#Terms) for each bit line. During **pclk0**, these are used to *precharge* the bit lines, after which the pull-up transistors are disabled but the lines remain charged through capacitance. When the selected row is opened after **pclk0**, it will be exposed to the precharged bit lines, which has the effect of charging up the high side of the cell. On the low side of the cell, the precharge current will simply drain to ground, as the gate on that side will be driven high.

In a typical DRAM circuit, the rows are automatically and periodically refreshed to prevent values from fading. In the PPU, no such logic exists, and rows are only refreshed when explicitly accessed. The reason the PPU (usually) gets away with this is that [sprite evaluation](../PPU_sprite_evaluation/) will access the entire OAM (provided rendering is enabled), refreshing the rows as a side effect.

### <span id="SRAM_.28Static_RAM.29"></span>SRAM (Static RAM) {#SRAM_.28Static_RAM.29}

SRAM uses cross-coupled inverters for storage and is accessed using a row/column scheme similar to DRAM. Compared to DRAM, SRAM does not need to be refreshed, tends to be faster, uses more die area per memory cell, and draws more power for the [NMOS](#Terms) version.

Below is a picture of SRAM memory cells used to store the PPU's [palette](../PPU_palettes/) (in this case the rows do go horizontally):

[<img src="../../wiki-images/Vid_sram.png" width="248" height="173" alt="Vid sram.png" />](../../wiki-images/Vid_sram.png)

## Miscellaneous circuitry {#Miscellaneous_circuitry}

### Decoders and mask ROMs {#Decoders_and_mask_ROMs}

- A [decoder](https://en.wikipedia.org/wiki/Decoder) is a circuit that maps input values to output values. A decoder that maps *m* input lines to *n* output lines is called an *m-to-n decoder*.
- A [mask ROM](https://en.wikipedia.org/wiki/Mask_ROM) is a type of read-only memory constructed by masking off parts of a circuit grid.

The two elements are covered together since their implementation turns out to be similar in this case.

Pictured below is the decoder and mask ROM that act as the lookup table for initialization of the [length counters](../APU_Length_Counter/) in the APU:

[<img src="../../wiki-images/Vis_len_rom.png" width="255" height="343" alt="Vis len rom.png" />](../../wiki-images/Vis_len_rom.png)

The length is set by writing bits 7-3 of e.g. **\$4003** (in the case of the first pulse channel), so the inputs to the decoder are bits 7-3 of the data bus. The output from the decoder feeds into the mask ROM, and the output from the mask ROM is the length from the lookup table. The length is used to initialize a counter that counts down to zero before silencing the channel.

The picture below shows a zoomed-in view of the lower part of the decoder and mask ROM:

[<img src="../../wiki-images/Vis_len_pla_zoom.png" width="817" height="353" alt="Vis len pla zoom.png" />](../../wiki-images/Vis_len_pla_zoom.png)

The spots of yellow diffusion in the decoder and mask ROM are connections to the metal wires, which run horizontally in the decoder and vertically in the mask ROM. By setting the gates connected to the diffusion high, the wires can be driven low.

In the decoder (right part) the input lines and their inverses run vertically (**/db7** has been highlighted to show its connection). By looking carefully at the bottom-most horizontal row in the decoder, we see that it is powered on the right side, and that the condition for it to remain high as it passes into the mask ROM is **/db7** AND **/db6** AND **/db5** AND **/db4** AND **/db3**. Another way to put this condition is **db7**-**db3** = \$00.

Similarly, the condition for the second row from the bottom to be high is **/db7** AND **/db6** AND **/db5** AND **db4** AND **/db3**, which translates to **db7**-**db3** = \$02. The conditions for the third and fourth rows from the bottom are **db7**-**db3** = \$04 and **db7**-**db3** = \$06, respectively.

The decoder is set up so that **db*x*** and **/db*x*** will never both drive the same horizontal line low (which would make it impossible for that line to ever be high), and in this case each row has a unique bit pattern that activates it. (It would also be possible to insert a "don't care" condition in the decoder by having *neither* **db*x*** nor **/db*x*** drive the line low.)

The decoder here is a 5-to-32 decoder, with 32 rows corresponding to the 32 possible bit patterns made with five bits. This type of decoder is said to *fully decode* its inputs, and is an instance of an *n-to-2<sup>n</sup> decoder*.

In the mask ROM (this one in particular being a NOR ROM), we see that each horizontal line from the decoder when high will cause a particular pattern to appear on the **len*x*** outputs. Reading off the bottom row, this pattern is **len7**-**0** = 00001001b = 9. Reading off the remaining rows from bottom to top, we get the values 00010011b = 19, 00100111b = 39, and 01001111b = 79.

Putting together the above, we have the following incomplete map from inputs to outputs:

| Index | Length |
|-------|--------|
| \$00  | 9      |
| \$02  | 19     |
| \$04  | 39     |
| \$06  | 79     |

By checking against the [APU length counter table](../APU_Length_Counter/), we see that these indeed are the length values corresponding to those indices (minus one, due to details of how the length counter works).

To give an example of a decoder that does not feed into a mask ROM, the picture below shows the internal 2A03 address decoder for the address range **\$4000**-**\$4017**, where signals such as **r4017** (read 4017) and **w4004** (write 4004) are generated.

[<img src="../../wiki-images/Vis_addr_pla.png" width="238" height="220" alt="Vis addr pla.png" />](../../wiki-images/Vis_addr_pla.png)

The theory behind the decoder and mask ROM seen here is closely related to that of [PLAs](https://en.wikipedia.org/wiki/Programmable_logic_array) (Programmable Logic Arrays), where we could view the decoder as the AND plane and the mask ROM as the OR plane (both implemented with NOR gates). [This introduction to PLAs](http://www.cs.umd.edu/class/sum2003/cmsc311/Notes/Comb/pla.html) is helpful.

### Adders {#Adders}

Pictured below is part of the [adder](https://en.wikipedia.org/wiki/Adder_(electronics)) used by the [sweep units](../APU_Pulse/) in the 2A03 to calculate the target period for sweep period updates to the second square channel (the first square channel is identical except for a small quirk related to subtraction; see below). The pictured part calculates the second bit (bit 1) of the sum, along with the carry for that bit position.

[<img src="../../wiki-images/Vis_adder.png" width="482" height="265" alt="Vis adder.png" />](../../wiki-images/Vis_adder.png)

The adder is split into two parts. The left-most part (having four columns) calculates bit 1 of the sum. The right-most part (with three columns) calculates the carry. Both **/sum1 out** and **/carry out** are powered, and can be forced low by certain combinations of the input signals being high. (For e.g. the left-most column, this combination is **addend1** AND **carry in** AND **sq1_p1**). The essential information is captured in the following truth table:

| sq1_p1 | addend1 | carry in | /sum1 out | /carry out |
|--------|---------|----------|-----------|------------|
| 0      | 0       | 0        | 1         | 1          |
| 0      | 0       | 1        | 0         | 1          |
| 0      | 1       | 0        | 0         | 1          |
| 0      | 1       | 1        | 1         | 0          |
| 1      | 0       | 0        | 0         | 1          |
| 1      | 0       | 1        | 1         | 0          |
| 1      | 1       | 0        | 1         | 0          |
| 1      | 1       | 1        | 0         | 0          |

As expected, this corresponds to an addition operation (with the sum and carry inverted).

The same logic is used to perform subtraction, by inverting each bit of the addend (using separate logic) and setting the carry in for the zeroth bit to 1. This corresponds to the usual invert-bits-and-add-one operation for negating a number in two's complement.

For unknown reasons, the inverted carry input for the zeroth bit of the first square channel is connected to VCC instead of the inverted sweep direction flag (as it is in the other square channel), making the carry input unconditionally zero. This leads to the value *minus one* being subtracted instead on that channel.

### Barrel shifters {#Barrel_shifters}

The below circuitry forms part of a [barrel shifter](https://en.wikipedia.org/wiki/Barrel_shifter), used to shift the inputs to the adders for sweep unit period updates in this case.

[<img src="../../wiki-images/Vis_barrel_shifter.png" width="206" height="209" alt="Vis barrel shifter.png" />](../../wiki-images/Vis_barrel_shifter.png)

As a side note, the bit inversion for subtraction by the sweep units happens *before* the bits enter the barrel shifter.

### Shift registers {#Shift_registers}

*(This section might be considered "advanced" on a first reading. I just wanted an example that made more complex use of clocks.)*

The picture below shows the 16-bit [shift register](https://en.wikipedia.org/wiki/Shift_register) that holds the high bits for background tiles (see the [PPU rendering](../PPU_rendering/) page). The upper eight bits can be reloaded from PPU VRAM data bus lines, and the output is taken from the lower eight bits (in this case, the particular bit to use is selected by the **[fine x scroll](../PPU_scrolling/)**). Bits flow clockwise through the shift register.

[<img src="../../wiki-images/Vis_shift_reg.png" width="250" height="247" alt="Vis shift reg.png" />](../../wiki-images/Vis_shift_reg.png)

Below is a zoomed-in view of three bits (**tile_h15**-**13**) from the upper-left part of the shift register:

[<img src="../../wiki-images/Vis_shift_reg_zoom.png" width="446" height="488" alt="Vis shift reg zoom.png" />](../../wiki-images/Vis_shift_reg_zoom.png)

The value of each bit corresponds to the value on the **(2)** side.

#### Control signals {#Control_signals}

The following signals control the shifting and reloading of the register (the names used were invented for the article and are not standard terminology):

- The **Invert** signal corresponds to **pclk0**, which is high during the initial half-cycle of a PPU cycle (see the [Clocks](#Clocks) section).
- The **Shift** signal corresponds to **pclk1**, which is high during the second half-cycle.
- The **Parallel load** signal controls pass transistors connected to **\_db0**-**7**, used to load the upper eight bits of the shift register.

**Shift** does not always exactly mirror **pclk1**, as explained below, which is the reason for the *≈* notation.

#### Shifting {#Shifting}

Shifting the register is a two-step process:

1.  During **pclk0**, **Invert** is driven high, making the value of **(1)** flow through the pass transistor into the red-highlighted node, which causes **(1)** to invert into **(2)**.
2.  During **pclk1**, **Shift** is driven high, which causes the node marked **(2)** to invert into the node marked **(3)** (the next bit of the shift register). **Invert** is low during this phase, and the value on the red-highlighted node is held via [wire capacitance](#Wire_capacitance_as_storage), which makes this a *[dynamic](https://en.wikipedia.org/wiki/Dynamic_logic_(digital_electronics)) shift register*.

Due to the bit of powered diffusion circled in red, the default value shifted into **(1)** is 1. However, as the value is held on the inverted side **(2)**, this means that zeroes are being shifted in.

#### Parallel load {#Parallel_load}

To perform a parallel load of the register, step **(2)** from above is modified so that **Shift** remains low during **pclk1** and **Parallel load** goes high instead, causing the new value for each cell to come from the data bus lines instead of from the previous cell.

The diagram below might clarify how the control signals are related. Each row is a PPU half-cycle.

| pclk0 | Invert | Shift | Parallel load     |
|-------|--------|-------|-------------------|
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 1     | 0                 |
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 1     | 0                 |
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 0     | 1 ← Reloaded here |
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 1     | 0                 |
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 1     | 0                 |
| 1     | 1      | 0     | 0                 |
| 0     | 0      | 1     | 0                 |

### <span id="Digital-to-analog_conversion_.28DAC.29"></span>Digital-to-analog conversion (DAC) {#Digital-to-analog_conversion_.28DAC.29}

The below Visual 2A03 circuitry controls the volume on the output pin for the two square channels (the triangle, noise, and DMC channels use a separate pin). Note that each successive bit has twice the weight of the preceding one in terms of the amount of powered diffusion connected to it.

[<img src="../../wiki-images/Vis_da_conversion.png" width="772" height="203" alt="Vis da conversion.png" />](../../wiki-images/Vis_da_conversion.png)

This is an example of a [*binary-weighted DAC*](https://en.wikipedia.org/wiki/Digital-to-analog_converter#DAC_types). A different type of DAC is used for the video output from the PPU (found in the upper-left of Visual 2C02, rotated 90 degrees here):

[<img src="../../wiki-images/Vis_vid_dac.png" width="671" height="137" alt="Vis vid dac.png" />](../../wiki-images/Vis_vid_dac.png)

The upper-left end is actually connected to VCC, and the lower-right to ground. This is a [*voltage ladder*](https://en.wikipedia.org/wiki/voltage_ladder), and works by tapping the wire (which behaves as a resistor) at different points along the run to get different voltages. As the simulator is purely digital, this circuit is not directly used in the simulation, and some parts that would otherwise interfere with it have been disconnected.

### Output drivers {#Output_drivers}

These are found on pins capable of doing output, which need to be able to source (generate) and sink large currents to drive the line high or low. The polysilicon wire that would cause the pin to source current is highlighted below.

[<img src="../../wiki-images/Vis_output_driver.png" width="294" height="111" alt="Vis output driver.png" />](../../wiki-images/Vis_output_driver.png)

Large clusters of [pull-up and pull-down](#Terms) transistors like these are sometimes called *superbuffers*. They also appear in some internal circuits that need to source or sink larger currents, e.g. due to having a large [*fan-out*](https://en.wikipedia.org/wiki/Fan-out) – a large number of connections from the logic gate's output to inputs of other gates.

On lines that are capable of being tri-stated, this is done by activating neither the pull-up nor the pull-down transistors, so that the pin neither sources nor sinks current. This is also done for reads on bidirectional lines, to prevent the output driver from interfering.

### Cut-off connections {#Cut-off_connections}

Some parts of the chips, especially outside the 6502 core, were designed using a copy-and-paste process called "[standard cell](https://en.wikipedia.org/wiki/standard_cell)", leading to some seemingly nonsensical and cut-off connections. These carry no special significance. The image below contains an example.

The 6502 core inside the 2A03 is a substantially tighter block of [NMOS](#Terms) (having been [designed by hand](https://en.wikipedia.org/wiki/Full_custom)), but it still has a few cut-off connections remaining from removal of the original output drivers.

[<img src="../../wiki-images/Vis_cutoff.png" width="288" height="87" alt="Vis cutoff.png" />](../../wiki-images/Vis_cutoff.png)

## Layers {#Layers}

(This information is not essential to reading the diagrams.)

The layers that make up the chip are as follows, in order from bottom to top: [substrate](https://en.wikipedia.org/wiki/Substrate_(electronics)), diffusion, oxide (with holes for [buried contacts](#Terms) and [vias](#Terms)), polysilicon, more oxide (with holes for vias), metal, and [passivation](https://en.wikipedia.org/wiki/Passivation_(chemistry)) (or "overglass", containing holes where [bond wires](#Terms) connect).

The way diffusion is powered or grounded is through vias to large areas of metal that are either grounded or powered.

## Transistor dimensions {#Transistor_dimensions}

Visual6502, Visual2A03, and Visual2C02 are purely digital simulators, so the effects of transistor dimensions don't matter. But you will often notice locations in the simulators where transistors are different shapes.

Here's the inverter from the beginning of this tutorial, now annotated with dimensions:

[<img src="../../wiki-images/Vis_inverter_dimensions.png" width="367" height="228" alt="Vis inverter dimensions.png" />](../../wiki-images/Vis_inverter_dimensions.png)

Because the layer of substrate is uniform thickness, everything is calculable in terms of [sheet resistance](https://en.wikipedia.org/wiki/Sheet_resistance). In the above annotated picture, two transistors are shown: one significantly wider than long, and the other the opposite. The aspect ratio (length divided by width) of the depletion mode pull-up transistor on the left is approximately 3.47, while the aspect ratio of the enhancement mode pull-down transistor on the right is 0.23. As a result, the pull-down transistor is approximately 15 times more effective at sinking current (which is good: it has to be able to override the pull-up).

The 2A03 uses these analog effects in its audio path:

[<img src="../../wiki-images/Vis_2a03_pcmout012.png" width="289" height="264" alt="Vis 2a03 pcmout012.png" />](../../wiki-images/Vis_2a03_pcmout012.png)

Shown are the three least significant bits of the 2A03's APU's PCM channel. **pcmout1** drives a single transistor (with some resistance R) and **pcmout2** drives two (resulting in a resistance R÷2). To give **pcmout0** a resistance of 2·R, they either would have had to make the transistor half as wide or twice as long. Halving the width wasn't an option because the diffusion areas are already as narrow as possible using this manufacturing technology. As a result, the gate for the least significant bit is longer.

## Clocks {#Clocks}

This section lists node names for various clocks that sequence operations within the chips. Some of the 6502 pin signals might have gained a "c\_" prefix in Visual 2A03 compared to Visual 6502.

### 6502 core pins {#6502_core_pins}

clk0  
The [φ0 clock input pin](../CPU_pinout/). Goes low at the beginning of a CPU cycle.

<!-- -->

clk1out, clk2out  
The **φ1** and **φ2** output pins. **φ2** is used to form [**M2**](../CPU_pinout/) in the 2A03, which has a modified [duty cycle](../Glossary/).

### 6502 internal clock signals {#6502_internal_clock_signals}

cp1  
High during the first phase (half-cycle) of a CPU cycle. The inverse of **clk0**.

<!-- -->

cclk  
High during the second phase of a CPU cycle. Roughly equivalent to **clk0**, but modified slightly to never overlap with **cp1** (though that won't be visible in the simulators).

### APU clock signals {#APU_clock_signals}

apu_clk1  
This clock signal has a 25% duty cycle. It ticks at half the rate of the CPU clock, and is high only when **φ2** is *low*.

<!-- -->

apu\_/clk2  
Like **apu_clk1**, but ticks on the opposite phase, and is also inverted so that it has a 75% duty cycle.

<!-- -->

apu_clk2*x*, where *x* is *a*, *b*, *c*, etc.  
Inverses of **apu\_/clk2**, used internally in various components.

This clock arrangement helps to ensure that timed events (various counters being decremented or reloaded) do not conflict with writes from the CPU (which only happen when **φ2** is high).

|             |       |     |       |     |       |     |       |     |
|-------------|-------|-----|-------|-----|-------|-----|-------|-----|
| φ1          | 1     | 0   | 1     | 0   | 1     | 0   | 1     | 0   |
| φ2          | 0     | 1   | 0     | 1   | 0     | 1   | 0     | 1   |
| apu_clk1    | **1** | 0   | 0     | 0   | **1** | 0   | 0     | 0   |
| apu\_/clk2  | 1     | 1   | **0** | 1   | 1     | 1   | **0** | 1   |
| apu_clk2*x* | 0     | 0   | **1** | 0   | 0     | 0   | **1** | 0   |

### PPU clock signals {#PPU_clock_signals}

clk0  
The input clock, fed from the [master clock](../Cycle_reference_chart/#Clock_rates). Used directly in video waveform generation.

<!-- -->

\_clk0  
The inverse of **clk0**.

<!-- -->

pclk0  
The **p**ixel clock. Derived from **clk0** by dividing by four (NTSC) or five (PAL). One cycle corresponds to a rendered dot, with **pclk0** being high during the first phase (half-cycle).

<!-- -->

pclk1  
The inverse of **pclk0**. High during the second phase of a pixel clock.

### <span id="Master_clock_and_CPU.2FPPU_clock_alignment"></span>Master clock and CPU/PPU clock alignment {#Master_clock_and_CPU.2FPPU_clock_alignment}

The clock divider in the PPU is clocked on a zero-to-one transition of the master clock, while the clock divider in the CPU is clocked on a one-to-zero transition. Diagrammatically, this might look like below for NTSC (for CPU clock, "!" denotes when the 2A03's M2 line goes high but the 6502's internal clock is still low).

    Master clock    | 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 1 0 ...
    PPU pixel clock | 0 0 0 1 1 1 1 0 0 0 0 1 1 1 1 0 0 0 0 1 1 1 1 0 0 0 0 1 1 1 1 0 0 0 0 1 1 ...
    CPU clock       | 0 0 0 0 0 0 0 0 0 ! ! ! 1 1 1 1 1 1 1 1 1 1 1 1 0 0 0 0 0 0 0 0 0 ! ! ! 1 ...

## 6502 cycle and phase timing {#6502_cycle_and_phase_timing}

During each active cycle (i.e., while the RDY line of the 6502 has not been pulled low), the CPU either reads from or writes to memory; there are no "idle" cycles w.r.t. the data bus in the 6502. Each such read or write cycle is split up into two equally long phases, called **φ1** (phase 1) and **φ2** (phase 2), corresponding to the clock signals above. **φ1** takes place while the clock input is low, **φ2** while it is high.

During each cycle, the [R/W signal](../CPU_pinout/) and the address bus lines are updated during **φ1**. In the simulators we see them change right away, but in a real 6502 there will be some delay. At the end of **φ2**, values are read from or written to the data bus lines. The IRQ and NMI interrupt lines appear to be sampled on the falling edge of **φ2** (as indicated in [this thread](http://forum.6502.org/viewtopic.php?f=1&t=2532) and also from observed behavior relating to the **VBlank** flag).

[This document](http://nesdev.org/6502_cpu.txt) lists data and address bus contents during each cycle of an instruction. For more detailed timing information, see [the MOS hardware manual](http://users.telenet.be/kim1-6502/6502/hwman.html).

## Terms {#Terms}

Below are various terms you might run into:

[Bond wire](https://en.wikipedia.org/wiki/Wire_bonding)  
A wire that connects an internal pad to an external pin on the chip package; see e.g. [this](https://commons.wikimedia.org/wiki/File:Wirebond-ballbond.jpg)

Buried contact  
A connection between diffusion and polysilicon.

[NMOS](https://en.wikipedia.org/wiki/NMOS_logic)  
The technology used for the transistors in the 2A03 and 2C02. In NMOS, transistors are made by creating regions of n-doped semiconductor that become the source and drain ("n-doped" because this doping increases the mobility of electrons and their **n**egative charge). This type of transistor is good at sinking current to ground (this is what causes a 0 bit to usually "win" in [bus conflicts](../Bus_conflict/)), and worse at pulling up. The transistors used in NMOS are more precisely called *n(-type )MOSFETs*. NMOS transistors are "active" when their gate is connected to a *high* voltage (i.e. VCC). See [this YouTube video](http://www.youtube.com/watch?v=IcrBqCFLHIY) for another high-level overview of how transistors work.

[PMOS](https://en.wikipedia.org/wiki/PMOS_logic)  
The counterpart to NMOS, PMOS is an older and slower (but initially easier to manufacture) technology used for making integrated circuits. In PMOS, transistors are made by creating regions of p-doped semiconductor that become the source and drain ("p-doped" because this doping increases the mobility of electron holes and their **p**ositive charge). This type of transistor is good at sourcing current from VCC, and worse at pulling down. The transistors used in PMOS are more precisely called *p(-type )MOSFETs*. PMOS transistors are "active" when their gate is connected to a *low* voltage (i.e. GND).

[CMOS](https://en.wikipedia.org/wiki/CMOS_logic)  
The *combination* of NMOS and PMOS within a single design, CMOS chips make use of both n-type **and** p-type MOSFETs in order to form logic gates. While PMOS requires weak pull-down resistors (often permanently enabled transistors connected to GND) and NMOS requires weak pull-up resistors (originally permanently enabled transistors connected to VCC, but later replaced with [depletion loads](https://en.wikipedia.org/wiki/Depletion-load_NMOS_logic)), CMOS makes use of p-type MOSFETs as strong pull-ups and n-type MOSFETs as strong pull-downs, arranged in pairs such that exactly one transistor in each pair is active at any given moment, made even simpler by the fact that a single input signal can connect directly to both transistors (when it is high it activates the NMOS pull-down, and when it is low it activates the PMOS pull-up).

Open drain  
A type of output that works by sinking current from an external pull-up resistor instead of generating current on its own. An example is the PPU's INT pin. The pull-up resistor is denoted "RM1" in [this Famicom wiring diagram](https://www.nesdev.org/Ntd_8bit.jpg).

Pull-up resistor  
A resistor connected to power. "Pull-up" comes from pulling the wire to a high state.

Pull-up transistor  
A transistor whose gate when high causes current to flow from a power source.

Pull-down transistor  
The analogue of a pull-up transistor for sinking to ground.

Via  
A connection between polysilicon/diffusion and metal.

## Tips for working with the simulators {#Tips_for_working_with_the_simulators}

### Node names in Visual 6502 {#Node_names_in_Visual_6502}

A hash (#) or tilde (~) on a node name signifies active low or negation in Visual 6502. Due to problems passing hashes in URLs, aliases were automatically introduced that use tildes instead (hence the "automatic alias replacing hash with tilde" comments).

### Clearing highlighting {#Clearing_highlighting}

When the simulator is loaded and after it has been run with "animate during simulation" enabled, nodes that are high will be highlighted. To get rid of this highlighting, click the "clear highlighting" button.

### Local copies of the simulator {#Local_copies_of_the_simulator}

Being able to add node names to **nodenames.js** can be very helpful when figuring out a circuit. To do this, a local version of the simulator can be downloaded with e.g. **\$ wget --convert-links** on a \*nix system. Please watch the recursion level and avoid downloading data needlessly, as at least Visual 2C02 and Visual 2A03 are hosted on a limited uplink.

### Extra node names {#Extra_node_names}

Many additional node names for Visual 2C02 can be found [in this repository](https://github.com/ulfalizer/Visual-2C02-nodes). The repo is maintained separately since it is updated often and not all nodes have been confirmed.

### PPU chip layout overview {#PPU_chip_layout_overview}

A high-level overview of the layout of the PPU [can be found here](https://www.nesdev.org/w/images/default/9/99/Ppuareas.png). Another, lower-level analysis [can be found here](http://breaknes.com/files/PPU/tilemap).
