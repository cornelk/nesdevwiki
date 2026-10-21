---
title: "Namco IPL Interface"
url: "/wiki/Namco_IPL_Interface/"
description: "The Namco IPL Interface is a proprietary Famicom expansion port device used for an Initial Program Loader (IPL.MAIN) present in Famicom Disk System games developed by Namco. The IPL and its interface…"
wikiPage: "Namco_IPL_Interface"
---

The Namco IPL Interface is a proprietary Famicom [expansion port](../Expansion_port/) device used for an Initial Program Loader (`IPL.MAIN`) present in [Famicom Disk System](../Family_Computer_Disk_System/) games developed by Namco. The IPL and its interface allowed developers to upload PRG-RAM and [PPU nametable](../PPU_nametables/) data to the system, with an option to save the data to disk.

As the interface hardware was never released commercially, the information on this page is based on analysis of disassembled code from the IPL.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- IPL interface strobe

The strobe is only ever performed once by the IPL on start-up.

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxD SDDx
       | ||| 
       | ||+-- Data bit
       | |+--- Data bit
       | +---- Status bit
       +------ Data bit

The IPL checks for the interface's presence by reading the status bit before and after a single strobe. If the status bit changes (0-\>1 or 1-\>0) after the strobe, the IPL initialises a data transfer. Otherwise, the IPL displays a fake [license screen](../FDS_BIOS/#Approval_check) and boots the loaded game instead. The state of the status bit is unknown after successive polls (i.e. while the data bits are being polled).

It is unknown if all data bits contain identical data. The IPL simply ANDs the \$4017 read with `%00010110` and checks if any bits were set in the result. The bit order is `%76543210`.

## Data transfer format {#Data_transfer_format}

The interface transfers data to the Famicom using a variant of the [Intel HEX](https://en.wikipedia.org/wiki/Intel_HEX) format described as the following:

    offset size    description
    ------  ----    -----------
    0   1   start code, raw $3A byte (":")
    1   1   ASCII-encoded data length (0: end data transfer)
    2   2   ASCII-encoded destination address (big endian)
    4   1   ASCII-encoded record type (must be $00, data type)
    5   x   ASCII-encoded data
    5+x 1   ASCII-encoded checksum complement
    6+x     -

ASCII-encoded bytes are sent by the interface as byte pairs encoding the hexadecimal digits in ASCII (E.g. \$5F would be encoded as `'5', 'F' -> $35, $46`), which are then decoded by the IPL. This means that the true sizes for each ASCII-encoded field in the above table are doubled.

An example record expressed as a list of ASCII strings:

`":", "02", "6942", "00", "BEEF", "A6"`

If the data length is 0, the IPL reads 10 more raw bytes before exiting. The most straightforward way to format the final record would be to append an extra ASCII-encoded byte to a traditional EOF record, as the record type is not checked in this scenario:

`":", "00", "0000", "01", "FF", "00"`

The destination address determines whether the data is transferred to PRG-RAM or the PPU. If the destination address is below \$6000, the IPL adds \$2000 to it and sends the data to the PPU (likely intended for [PPU nametable](../PPU_nametables/) data). Otherwise, the destination address is used as-is and the data is sent to PRG-RAM instead.

The checksum operation is performed on the decoded bytes with the carry always cleared, starting from the data length field. The complement must result in a final checksum of 0.

## Hardware {#Hardware}

The interface hardware was likely connected using an RS232-style [serial cable](../Serial_Cable_Construction/) with a baud rate of 38400 (~46.60 CPU cycles per bit, the IPL waits 47 cycles). The polling routine in the IPL checks for a stop + start bit (0-\>1-\>0 transition), then waits 62 CPU cycles (~1.5 periods) before polling the data bits. Parity bits are not checked by the IPL.

USB-to-TTL parts such as the PL2303 or FT232 may be able to recreate the interface on modern hardware, using something similar to [this cable specification](../Serial_Cable_Specification/). A key difference here is that the TX signal must be *logically inverted* (so the console "un-inverts" the incoming signal).

An [Arduino-based implementation](https://forums.nesdev.org/viewtopic.php?p=303177#p303177) of the interface hardware was used to successfully upload payloads to original hardware.

## See Also {#See_Also}

- [Namco Serial Checker Interface](../Namco_Serial_Checker_Interface/) - A similar serial interface used for some cartridge games developed by Namco.
- [Serial Bootloader](../Serial_Bootloader/) - A minimal bootloader which also uses serial transfers.
- [Serial Cable Routines](../Serial_Cable_Routines/) - Code examples for serial transfers, targeting a baud rate of 57600.

## References {#References}

- [GitHub repository:](https://github.com/TakuikaNinja/IPL-MAIN) IPL.MAIN documentation and disassembly. Also includes a Lua script and Arduino sketch to simulate the interface.
- [GitHub repository:](https://github.com/TakuikaNinja/IPL-demo) IPL-demo program, to be loaded using the IPL.
- [Forum discussion](https://forums.nesdev.org/viewtopic.php?t=24983)
