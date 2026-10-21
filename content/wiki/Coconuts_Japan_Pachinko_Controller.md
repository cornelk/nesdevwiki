---
title: "Coconuts Japan Pachinko Controller"
url: "/wiki/Coconuts_Japan_Pachinko_Controller/"
description: "The Coconuts Japan Pachinko Controller (CJPC-102) is a Famicom expansion port controller designed by Hori Electric Co. that combines a standard controller with a large analog dial with a range of…"
wikiPage: "Coconuts_Japan_Pachinko_Controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Coconuts Japan Pachinko Controller (CJPC-102) is a Famicom expansion port controller designed by Hori Electric Co. that combines a standard controller with a large analog dial with a range of about 90 degrees. The dial uses a spring to return it to its rest position. This controller was used in the following games:

- *Pachinko Daisakusen*
- *Pachinko Daisakusen 2*
- *Pachio Kun 4*
- *Pachio Kun 5*

## Layout {#Layout}

        |
    ,--------.   \-,-------.
    |        |    |         |
    |       |    |           |
    |       |   /|           |
    |   +   |  /_|   B   A   /
    |        |    |         /
    |         `--'         /
    |                     /
    |  Se St             /
    `-------------------'

The Pachinko Controller is a thick controller with a front surface that slopes downward. Its button layout matches that of the standard controller, but with the select and start buttons positioned for use with the left hand rather than a more-ambidextrous central location. A large gap runs down most of the center of the controller, and the right half is rounded. On the right half, sandwiched between the front and back plastic shells of the controller is a dial with a slightly larger radius than the controller. On the left of the dial are two angled ridges. The controller is intended to be stabilized or held with the left hand and the dial manipulated with the right hand, using a palm grip that places the right hand's thumb below the larger ridge, index finger to the left of the smaller ridge, and other fingers placed on the dial, allowing clockwise rotation against the spring and counterclockwise rotation with the spring. The dial can also be used with the right index finger when holding the controller as a standard controller, but not as effectively.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Strobe controller shift register

This matches the normal strobe behavior used by the [standard controller](../Standard_controller/#Input_.28.244016_write.29).

### <span id="Output_.28.244016_read.29"></span>Output (\$4016 read) {#Output_.28.244016_read.29}

    7  bit  0
    ---- ----
    xxxx xxEx
           |
           +-- Controller status bit

The controller status is returned in a 16-bit serial report:

     Bits
     0-7  - Standard controller
     8    - (Always 1)
     9-15 - 7-bit analog value, inverted (high bits first)

     16+  - (Always 1)

Note that the analog value is provided to the CPU inverted. This article will discuss the raw, non-inverted values.

Unlike the [Arkanoid controller](../Arkanoid_controller/), which starts an analog-to-digital conversion in response to a strobe, the Pachinko Controller converts constantly. The controller has a 7-bit counter that is always counting up at a rate of approximately 20 kHz. A voltage based on all 7 bits of the counter is compared against a voltage based on the dial's potentiometer. When the counter voltage exceeds the potentiometer voltage, the 7 counter bits are simultaneously latched. When the controller is strobed, it is this latched conversion that is latched for the report, so the controller always provides a single complete conversion in its report.

The counter wraps every 128 counts, which takes about 6.4 ms, so conversions will latch at this rate on average. However, conversions will occur faster or slower than this as the dial is turned because the conversion occurs when exceeding the potentiometer voltage, which varies. In fact, if the potentiometer voltage increases faster than the counter voltage, multiple conversions could occur on a single pass, but this may be infeasible in practice because of the limits of human movement speed and capacitors that restrict how quickly the voltages can change.

At power-on, if the dial is in its rest position, the latched value is usually 0, but a conversion may occur immediately, typically producing a value of \$01 or \$02. Once the dial has been used, its minimum value is \$03 and its maximum value is approximately \$72-74, depending on mechanical factors. The maximum value is restricted by the plastic components and thus may vary by controller and by elastic deformation of the plastic. On the other hand, the minimum value is restricted by the conversion circuit, and there is a small deadzone around the rest position of the dial where turning does not change the value.

When the potentiometer voltage is sufficiently low, the counter voltage always exceeds it, preventing the conversion latches from taking a new value. Thus, if it is below the voltage that would produce the minimum value of \$03, no conversions occur and the stale value continues to be latched. If the dial is returned to rest quickly enough, this means that larger values than the minimum can be latched until the dial is turned far enough to trigger conversions again. For this reason, it is common in normal use for the at-rest value to be \$04 rather than the actual minimum of \$03, and values as large as \$06 have been observed by allowing the spring to freely return the dial to rest from the maximum position. Software should extend the deadzone to accommodate this quirk.

Contemporary software expects a range of \$00-63, clamping values above \$63. If bit 8 of the report does not match the expected 1, these games ignore the analog value. Note that the controller cannot report values less than \$03 under normal operation, so ball launching cannot be completely stopped.

## Hardware {#Hardware}

The controller contains two PCBs, with these chips:

- Two 4021 shift registers
- A 4024 7 bit ripple counter
- A 74HC574 8-bit D register
- A µPC393 dual comparator

One half of the µPC393 is used to form the 20kHz clock, nominally 75kΩ·470pF·2·ln(2). This clock feeds the 4024. The 4024's RESET input is always disabled. The output of the 4024 goes into a binary-weighted DAC, forming a sawtooth wave. The dial angle sets a voltage using a potentiometer. The other half of the µPC393 is used to detect when the sawtooth wave exceeds the voltage from the potentiometer. This latches the current count from the 4024 in the 74HC574. The shift registers are daisy-chained to form a single 16-bit report, like in the early SNES controllers or Virtual Boy controller.

## Images {#Images}

### PCBs {#PCBs}

[<img src="../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_1.jpg" width="400" height="297" alt="Coconuts Japan Pachinko Controller PCB 1.jpg" />](../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_1.jpg) [<img src="../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_2.jpg" width="400" height="307" alt="Coconuts Japan Pachinko Controller PCB 2.jpg" />](../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_2.jpg) [<img src="../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_3.jpg" width="400" height="233" alt="Coconuts Japan Pachinko Controller PCB 3.jpg" />](../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_3.jpg) [<img src="../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_4.jpg" width="400" height="236" alt="Coconuts Japan Pachinko Controller PCB 4.jpg" />](../../wiki-images/Coconuts_Japan_Pachinko_Controller_PCB_4.jpg)

## See Also {#See_Also}

- [Video demonstrating how the dial maps to gameplay](https://youtu.be/qWbw-9C0fGE?t=606)
