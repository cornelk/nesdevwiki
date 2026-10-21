---
title: "Serial Cable Construction"
url: "/wiki/Serial_Cable_Construction/"
description: "This page describes how to construct a serial cable for connecting a NES/Famicom to a PC. See Serial Cable Specification for more about pin assignments and rationale."
wikiPage: "Serial_Cable_Construction"
---

This page describes how to construct a serial cable for connecting a NES/Famicom to a PC. See [Serial Cable Specification](../Serial_Cable_Specification/) for more about pin assignments and rationale.

# Overview {#Overview}

RS-232 serial communication uses +12 and -12V signals, while the NES/Famicom console uses +5V and 0V, commonly referred to as TTL. Connecting RS-232 directly would damage the console, so the RS-232 levels must be translated to/from TTL. RS-232 is also inverted from normal, so a converter inverts the signals as well. Thus, a converter translates +12V to/from 0V, and -12V to/from +5V. There are a few ways to do this conversion, described below. There are also a few ways to connect the converted signals to the NES/Famicom consoles, described separately. A cable is a combination of a conversion method and connection method, so there are even more possibilities.

RS-232 signals are named TxD, RxD, CTS, and GND. Inverted TTL-level signals are named TX, RX, /CTS, GND, and +5V. The NES should *never* be connected directly to any RS-232 signals, except GND.

# RS-232-TTL level conversion {#RS-232-TTL_level_conversion}

## FTDI TTL-232R {#FTDI_TTL-232R}

The FTDI TTL-232R cable plugs into a USB port on almost any computer and provides the necessary TTL signals for the NES. This cable is by far the simplest to use and has high-quality drivers for modern operating systems (Windows, OS X, Linux). It is well worth its cost. You'll also need a male header, 0.1 inch spacing, with 5 pins in a row (if more, you can cut off the extras). Cutting and stripping the wires is not recommended, because you might want to use the FTDI cable for other devices which have a header connector as well, like the popular Arduino microcontroller boards.

*Do not* connect the red or green wires to anything. Also, *avoid* the TTL-232R-3V3 cable which uses 3.3V signals and will not work with the NES.

          ___
         | _ |
     GND ||_|| Black
    /CTS ||_|| Brown
         ||_|| Red
      TX ||_|| Orange
      RX ||_|| Yellow
         |'_'| Green
         |___|
     FTDI TTL-232R

TODO: mention the cheaper PL2303-based cables, even though they are nearly identical, just different wire colors.

TODO: mention 1K resistor between TX and NES to reduce current when NES is off but serial is plugged into USB.

## MAX232 {#MAX232}

The MAX232 level converter chip internally generates the RS-232 +/-12V, ensuring compatibility with any serial port. Pay close attention to the counter-intuitive capacitor polarities. If possible, also connect pins 1, 4, and 6 of the RS-232 connector together, which helps the PC determine that something is connected.

[<img src="../../wiki-images/Max232.png" width="265" height="220" alt="MAX232 RS-232-TTL schematic" />](../../wiki-images/Max232.png)

| Quantity | Part | Comments |
|----|----|----|
| 1 | MAX232 | Could use variations, like MAX232A etc. See \[[\| datasheet](http://www.maxim-ic.com/datasheet/index.mvp/id/1798)\] |
| 5 | 1 μF 16V or greater capacitor | Polarized electrolytic or tantalum |

    GND     TxD   RxD
    .-------------------.
    \ 5   4   3   2   1 /
     \                 /
      \ o   8   o   6 /
       `-------------'
           CTS

    9-pin RS-232 cable (female)

## Cheap {#Cheap}

Two transistor-based inverters convert well enough to work with many current serial ports.

For a one-way PC to NES cable, construct an inverter that converts -12V to +5V and +12V to 0V (left-hand schematic). A diode prevents a -12V input from exceeding the transistor's reverse breakdown voltage, which would damage the transistor over time (note diode's orientation). If possible, also connect pins 1, 4, and 6 of the RS-232 connector together, which helps the PC determine that something is connected.

For a two-way cable that allows NES to PC as well, additionally construct a second inverter on the NES output (right-hand schematic). While this only goes down to 0V instead of -12V, it works with many modern serial ports. This has worked on everything I've tried so far, including a few PL2303-based USB serial adaptors, and the built-in serial port on a Dell PC. It's only the NES to PC direction that doesn't follow the spec exactly; PC to NES will work with all serial ports.

[<img src="../../wiki-images/Rs232_to_ttl.png" width="198" height="183" alt="RS-232 to TTL schematic" />](../../wiki-images/Rs232_to_ttl.png)

[<img src="../../wiki-images/Ttl_to_rs232.png" width="190" height="183" alt="TTL to RS-232 schematic" />](../../wiki-images/Ttl_to_rs232.png)

| Quantity | Part | Comments |
|----|----|----|
| 2 | 2N3904 transistor | Any general-purpose NPN transistor will do, like the 2N2222, etc. |
| 2 | 1k resistor | 1/4 or 1/8 watt, brown-black-red |
| 2 | 4k7 resistor | 1/4 or 1/8 watt, yellow-violet-red |
| 1 | 1N914 diode | Other small-signal silicon diode will also work, like the 1N4148, etc. |

# <span id="Connection_to_NES.2FFamicom"></span>Connection to NES/Famicom {#Connection_to_NES.2FFamicom}

The TTL signals from one of the above circuits can be connected to the NES via a controller cable, or the NES/Famicom via the expansion port. For clarity, unused pins are not numbered or labeled.

## NES controller cable {#NES_controller_cable}

The simplest method of connecting to a NES is through a normal controller cable. Colors are for standard NES controller cable, not third-party cables. Verify pin connections if possible.

                  _
                 /1| GND (brown)
    (white) +5V |5.| 
                |.3| RX and /CTS (orange)
                |.4| TX (yellow)
                '--'
          NES controller cable

## Famicom expansion cable {#Famicom_expansion_cable}

The expansion port is the only way to connect on the Famicom. If desired, an external controller can still be connected at the same time as serial, as serial uses D4. A Neo-Geo controller extension cable can be used, as it has the same connector. A PC joystick connector *will not work*, due to having insufficient depth.

                       TX         GND
    -----------------------------------
    \  o   o   o   o   4   o   o   1  /
     \                               /
      \  15  o   o  12   o   o   o  /
       `---------------------------'
        +5V      RX & /CTS  

    Famicom female expansion cable

## NES expansion port {#NES_expansion_port}

Connecting to the expansion port avoids disruption of the second controller entirely, but connectors are hard to find, if they exist at all. Another option is to solder directly to the circuit board, and put the serial circuit inside the NES.

              ------- 
             |  ___  | 
         +5V | 1  48 | +5V
         GND | 2  47 | GND
             | [   ] |
             | [   ] |
             | [   ] |
             | [  43 | RX & /CTS
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
    (back)   | [   ] |   (front of NES)
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
          TX | 20  ] |
             | [   ] |
             | [   ] |
             | [   ] |
             | [   ] |
             |  ---  |
              -------
    NES expansion port on bottom of console
