---
title: "Namco Serial Checker Interface"
url: "/wiki/Namco_Serial_Checker_Interface/"
description: "The Namco Serial Checker Interface is a proprietary Famicom expansion port device used for some cartridge games developed by Namco. Supported games send and receive data using this interface to check…"
wikiPage: "Namco_Serial_Checker_Interface"
---

The **Namco Serial Checker Interface** is a proprietary Famicom [expansion port](../Expansion_port/) device used for some cartridge games developed by Namco. Supported games send and receive data using this interface to check if a cartridge self-test should be performed at power-on/reset.

As the interface hardware was never released commercially, the information on this page is based on analysis of disassembled code from supported games.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxI
            |
            +- Data to interface input

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxx Dxxx
         |
         |
         |
         +---- Data from interface output

The interface appears to consist of a 4-bit shift register (initial contents are undefined). Data sent to the interface will appear in the output after 4 reads. The received data is logically inverted from what was sent due to how the expansion port inputs work.

## Usage {#Usage}

Games using the interface do the following at some point in their reset handler (the specific code may differ):

1.  Fetch a source byte.
2.  Initialise a target byte.
3.  Read the data bit from \$4017.d3 and shift it into the target byte.
4.  Shift one bit out from the source byte and write it to \$4016.d0.
5.  Repeat steps 3-4 11 more times. (i.e. 4 "dummy" bits + 8 data bits total)
6.  Check if the target byte matches the source byte XORed with \$FF. Go to step 9 if false.
7.  Return to step 1 and repeat for however many source bytes remain.
8.  If all bytes pass the check, execute a cartridge self-test (partial PRG/CHR checksum verification).
9.  Continue with init tasks.

Note: The source data used is commonly a 32- or 64-byte string consisting of programmer credits, sometimes XORed with \$FF.

## SRAM Checker Variant {#SRAM_Checker_Variant}

A variant of this interface appears to have been used by *Family Circuit '91* and *Splatterworld* (unreleased) to verify and/or initialise battery-backed SRAM contents.

### <span id="Input_.28.244016_write.29_2"></span>Input (\$4016 write) {#Input_.28.244016_write.29_2}

    7  bit  0
    ---- ----
    xxxx xxxI
            |
            +- Data to interface input

### <span id="Output_.28.244017_read.29_2"></span>Output (\$4017 read) {#Output_.28.244017_read.29_2}

    7  bit  0
    ---- ----
    xxxx xxDx
           |
           |
           |
           +-- Data from interface output

This interface is equivalent to shorting the OUT0 and joypad 2 /D1 pins on the expansion port. Thus, the received data is logically inverted from what was sent.

## See Also {#See_Also}

- [GitHub Gist](https://gist.github.com/TakuikaNinja/cf2dab7ed4d5869c4be94a1753ffa76f) - Lua script to simulate the interface in Mesen2.
- [GitHub Gist](https://gist.github.com/TakuikaNinja/f4c3310e0000e29291aa59cf9f246dda) - Same as above, but for the SRAM variant.
- [Namco IPL Interface](../Namco_IPL_Interface/) - A similar serial interface used in [FDS](../Family_Computer_Disk_System/) games developed by Namco.

## References {#References}

- [The Cutting Room Floor](https://tcrf.net/The_Cutting_Room_Floor:Common_Things#Cartridge_Self-Tests)
- [Forum discussion](https://forums.nesdev.org/viewtopic.php?t=25809)
