---
title: "Extended NES I/O"
url: "/wiki/Extended_NES_I/O/"
description: "The Extended NES I/O (ENIO) is a homebrew device by Chykn consisting of an EXP board and a CPU board. The EXP board plugs into the NES-001 expansion port and provides expansion audio and various…"
wikiPage: "Extended_NES_I_O"
---

The Extended NES I/O (ENIO) is a homebrew device by Chykn consisting of an EXP board and a CPU board. The EXP board plugs into the NES-001 expansion port and provides expansion audio and various headers, while the CPU board attaches to the EXP board and provides networking, storage, and USB. The EXP board was produced and sold in small quantities, while the CPU board did not make it beyond prototypes and was not generally sold. As such, much of this information has been sourced from archives of [public documentation](https://web.archive.org/web/20161225130448/http://enio.chykn.com/wiki/index.php/Main_Page).

# EXP board {#EXP_board}

The EXP board enables expansion audio by connecting EXP6 to the expansion port's audio in with a 47K resistor. Headers are included breaking out the various expansion port signals, one of them providing Famicom expansion port functionality via an included cable.

## <span id="Pinout_.28v1.2.29"></span>Pinout (v1.2) {#Pinout_.28v1.2.29}

### J1: NES EXP connector {#J1:_NES_EXP_connector}

See [NES expansion port](../../Expansion_port/#Pinout_2). Note that pins 1 and 48 (both +5V) are not connected on the EXP board.

### <span id="J2:_Famicom_EXP_Pin_Compatible_Header_.2816-pin_R.2FA_IDC.29"></span>J2: Famicom EXP Pin Compatible Header (16-pin R/A IDC) {#J2:_Famicom_EXP_Pin_Compatible_Header_.2816-pin_R.2FA_IDC.29}

                                +------+
                         GND -- |01  09| -> Joypad 2 /OE
    Audio In (not connected) -> |02  10| -> OUT2
                        /IRQ ?? |03  11| -> OUT1
                 Joypad 2 D4 ?? |04  12  -> OUT0
                 Joypad 2 D3 ?? |05  13  <- Joypad 1 D1
                 Joypad 2 D2 -> |06  14| -> Joypad 1 /OE
                 Joypad 2 D1 -> |07  15| -> +5V (J1.1)
                 Joypad 2 D0 ?? |08  16| -- NC
                                +------+ 

Notes:

- ?? directionality depends on the presence of other hardware, such as a joypad or a cartridge that always drives /IRQ. If conflicting hardware is not present, these can be used as inputs to the console.
- **J2.2 Audio In** : On the Famicom, this pin is used for audio out, though can accept audio input. Here, it connects to J6.2 and is intended to connect to the NES audio in via a 47K resistor, which appears to normally not be populated, leaving this pin not connected.
- Pin order differs on this header compared to J3 and J4.
- The EXP board came with an adapter cable that connects to this header and provides a Famicom-compatible male DA15 port.

### <span id="J3:_ENIO_CPU_Board_Header_.2826-pin_R.2FA_IDC.29"></span>J3: ENIO CPU Board Header (26-pin R/A IDC) {#J3:_ENIO_CPU_Board_Header_.2826-pin_R.2FA_IDC.29}

                       +------+
         +5V (J1.1) -- |01  02| -- GND
                A15 <- |03  04| -> OUT2
               OUT1 <- |05  06| -> OUT0
               EXP9 ?? |07  08| -> /CE (EXP8)
         R/W (EXP7) <- |09  10| ?? EXP5
       Joypad 2 /OE <- |11  12| ?? Joypad 2 D0
        Joypad 2 D1 -> |13  14  <- Joypad 2 D2
        Joypad 2 D3 ?? |15  16| ?? Joypad 2 D4
             CPU D0 <> |17  18| <> CPU D1
             CPU D2 <> |19  20| <> CPU D3
             CPU D4 <> |21  22| <> CPU D5
             CPU D6 <> |23  24| <> CPU D7
    Unregulated Vdd -- |25  26| ?? /IRQ
                       +------+

Notes:

- ?? directionality depends on the presence of other hardware, such as a joypad or a cartridge that always drives /IRQ. If conflicting hardware is not present, these can be used as inputs to the console.

### <span id="J4:_Misc_Signal_Header_.2816-pin_unpopulated.29"></span>J4: Misc Signal Header (16-pin unpopulated) {#J4:_Misc_Signal_Header_.2816-pin_unpopulated.29}

         +5V (J1.48) -- |01  02| -- GND
                EXP0 ?? |03  04| <> /NMI
                EXP2 ?? |05  06| ?? EXP1
                EXP4 ?? |07  08| ?? EXP3
         Joypad 1 D4 ?? |09  10| -> Joypad 1 /OE
         Joypad 1 D2 -> |11  12| ?? Joypad 1 D3
         Joypad 1 D0 ?? |13  14| -> Joypad 1 D1
    4.00 MHz CIC CLK <- |15  16| ?? /IRQ

Notes:

- ?? directionality depends on the presence of other hardware, such as a joypad or a cartridge that always drives /IRQ. If conflicting hardware is not present, these can be used as inputs to the console.

### <span id="J5:_Audio_and_Video_Out_.283-pin_unpopulated.29"></span>J5: Audio and Video Out (3-pin unpopulated) {#J5:_Audio_and_Video_Out_.283-pin_unpopulated.29}

    |1| -- GND
    |2| -> Audio Out
    |3| -> Video Out

### <span id="J6:_Audio_In_.282-pin_unpopulated.29"></span>J6: Audio In (2-pin unpopulated) {#J6:_Audio_In_.282-pin_unpopulated.29}

    |1| -- GND
    |2| <- Audio In (not connected)

Notes:

- J6.2 Audio In: Connects to J2.2, but does not connect to the console. See [J2.2](#J2:_Famicom_EXP_Pin_Compatible_Header_.2816-pin_R.2FA_IDC.29) above.

### <span id="J7:_Power_.283-pin_Molex_KK_254.29"></span>J7: Power (3-pin Molex KK 254) {#J7:_Power_.283-pin_Molex_KK_254.29}

    1  -- Unregulated Vdd
    2| -- +5V (J1.48)
    3  -- GND

# CPU board {#CPU_board}

The CPU board contains the following components:

- PIC32
- 100 Mb ethernet port
- MicroSD cart slot (for bootloading or game data storage)
- USB A port (such as for keyboards)
- MiniUSB port (for power if not connected to the console)
- 26-pin IDC header (see [EXP board J3](#J3:_ENIO_CPU_Board_Header_.2826-pin_R.2FA_IDC.29) above)
- Xilinx CPLD
- 3.3V switching regulator (selectable between 5V and unregulated Vdd, and 5V selectable between NES and miniUSB)

The CPLD enables communication between the NES and PIC, raising an interrupt on the PIC on access, and also acts as a level shifter. The PIC establishes network connections, transfers data, and buffers the CGP (Classic Gaming Protocol)<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> packets sent between the ENIO and NES.

## Addressing {#Addressing}

The CPU board can be accessed 8 bits at a time in two ways.

### Compatibility Mode {#Compatibility_Mode}

This can be used with any cartridge hardware, but is not compatible with a second controller. This uses:

- \$4017R for reads
- Any address below \$8000 for writes
- OUT1 as R/W

### Direct Addressing Mode {#Direct_Addressing_Mode}

This uses the cartridge to map reads and writes to \$4700, supplying R/W and /CE to the CPU board via EXP pins:

    EXP7: R/W
    EXP8: /CE

# References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [Classic Gaming Protocol proposal](https://web.archive.org/web/20150817135809/http://enio.chykn.com/wiki/index.php/CGP)
