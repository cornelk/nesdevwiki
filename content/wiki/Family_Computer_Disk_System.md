---
title: "Family Computer Disk System"
url: "/wiki/Family_Computer_Disk_System/"
description: "The Family Computer Disk System (Famicom Disk System, FDS for short) is a video game add-on made by Nintendo for the Famicom and sold in Japan and Hong Kong. It was designed to reduce the cost of…"
wikiPage: "Family_Computer_Disk_System"
wikiCategories:
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/Nintendo-Famicom-Disk-System.jpg"><img src="../../wiki-images/Nintendo-Famicom-Disk-System.jpg" width="300" height="266" /></a>
<figcaption>Famicom, FDS RAM adapter, and FDS disk drive</figcaption>
</figure>

<div class="infobox"><div style="text-align: center"><b>FDS</b></div><table class="oddf0 floatright"><tbody><tr><td><b>Company</b></td><td>Nintendo </td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Pinout</b></td><td><a href="../RP2C33_pinout/">RP2C33 pinout</a></td></tr><tr><td><b>BIOS PRG ROM size</b></td><td>8K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>32K </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>Disk capacity</b></td><td>~64K per side </td></tr><tr><td><b>Mirroring</b></td><td>H or V, switchable </td></tr><tr><td><b>Bus conflicts</b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>Yes </td></tr></tbody></table></div>

The **Family Computer Disk System** (**Famicom Disk System**, **FDS** for short) is a video game add-on made by Nintendo for the Famicom and sold in Japan and Hong Kong. It was designed to reduce the cost of making copies of software by switching from mask [ROM](../ROM/) chips to Nintendo's proprietary Disk Card, a storage medium based on Mitsumi's Quick Disk (QD).

Unfortunately for Nintendo, it also reduced the pirates' cost of making copies of software. Software is stored on one or multiple disk sides. The [FDS BIOS](../FDS_BIOS/) is responsible for controlling the disk drive to load data from disks to PRG RAM or VRAM before giving control to the software. Additional hardware features include a timer [IRQ](../IRQ/) and a [wavetable channel](../FDS_audio/).

## Hardware {#Hardware}

The FDS comes in two parts: The disk drive (HVC-022) and the RAM adapter (HVC-023).

The RAM adapter is a special shaped cartridge that contains the RAM chips and an ASIC with DRAM controller, IRQ hardware, sound generation hardware, serial interface for the disk drive, and parallel port. The disk drive has to be powered separately and is only connected to the console via a [serial cable](../FDS_RAM_adaptor_cable_pinout/) to the RAM adapter.

The disk drive is a modified version of the Mistumi Quick Disk (QD) drive. Conventional floppy disk drives contain two motors: a spindle motor that spins the disk at a specific speed, and a stepper motor which moves the read/write head between each circular data track. By comparison, QD drives only contain a single motor which does both at once, so it instead stores the data in a single spiral-shaped track. There is a mechanism that detects when the head reaches the end of the disk and makes it return to the start (making an audible click). Because of this limitation, random access to the disk is impossible, making QD drive data access behave similarly to a reel of tape (but *much* faster). Data can only be accessed by spinning the disk, waiting for the head to reach the inner edge, then waiting again until the desired data file is reached. A complete cycle through an entire disk side takes about 7 seconds. A red LED is located at the front of the drive and turns on while the drive motor is active.

The QD drive only contains basic electronics, there is no "intelligence" in it; therefore, the serial interface almost directly represents what is stored on the disk.

### Protections {#Protections}

Some protections were added to the FDS drive in an attempt to thwart unlicensed software (primarily disk copiers):

- The FDS disk drive contains a physical lockout in the form of molds which only fit disks with recesses spelling "NINTENDO", in order to prevent regular QDs from being used (also a similar trademark protection as the [license screen](../FDS_BIOS/#Approval_check)). This can be bypassed by using disks with custom recesses, including adapters which can be fitted onto regular QDs.
- Later revisions of the drive controller power board/IC contain protections against long data write operations performed by disk copiers. This can be bypassed through hardware modifications.
  - TODO: Measure/determine how much data can be written on protected drives.

### Disks {#Disks}

The FDS Disk Card (HVC-021) is a modified version of the Mitsumi Quick Disk. The recesses spelling "NINTENDO" is designed to fit with molds on the FDS disk drive (see protections above).

See:

- [FDS disk format](../FDS_disk_format/) - the disk data format and file structure.
- [FDS file format](../FDS_file_format/) (**.FDS**) - an archival file format for storing and emulating FDS disks.
- [QD file format](../QD/) (**.QD**) - an archival file format used in emulators by Nintendo.

## Banks {#Banks}

All Banks are fixed

- PPU \$0000-\$1FFF: 8k CHR RAM
- CPU \$6000-\$DFFF: 32k PRG RAM
- CPU \$E000-\$FFFF: 8k BIOS PRG ROM

## Registers {#Registers}

|  |  |
|----|----|
| <img src="../../wiki-images/Edit-clear.svg" width="40" height="40" /> | This article section is pending updates to include new research and to verify existing documentation. |

\$402x registers are write-only, \$403x registers are read-only

### <span id="Timer_IRQ_reload_value_low_.28.244020.29"></span>Timer IRQ reload value low (\$4020) {#Timer_IRQ_reload_value_low_.28.244020.29}

    7  bit  0
    ---------
    LLLL LLLL
    |||| ||||
    ++++-++++- 8 LSB of timer IRQ reload value

### <span id="Timer_IRQ_reload_value_high_.28.244021.29"></span>Timer IRQ reload value high (\$4021) {#Timer_IRQ_reload_value_high_.28.244021.29}

    7  bit  0
    ---------
    LLLL LLLL
    |||| ||||
    ++++-++++- 8 MSB of timer IRQ reload value

Unlike \$4022, \$4020 and \$4021 are not affected by the \$4023.0 (disk registers enabled) flag - the reload value can be altered even when disk registers are disabled.

### <span id="Timer_IRQ_control_.28.244022.29"></span>Timer IRQ control (\$4022) {#Timer_IRQ_control_.28.244022.29}

    7  bit  0
    ---------
    xxxx xxER
           ||
           |-- Timer IRQ Repeat Flag
           +-- Timer IRQ Enabled

When \$4022 is written to with bit 1 (IRQ enabled) set, the reload value is copied into the IRQ's counter. Each CPU clock cycle the counter is decremented by one if the enable flag is set.

When the counter's value is 0 and the IRQ enable flag is on, the following happens on every CPU cycle:

- An IRQ is generated.
- The IRQ counter is reset to its reload value (contained in \$4020+\$4021)
- If the IRQ repeat flag is NOT set, the IRQ enabled flag is cleared and the counter stops.

Notes:

- This register is affected by the \$4023.0 (Enable disk I/O registers) flag - if disk registers are disabled, it is impossible to start the IRQ counter (writing to \$4022 has no effect).
- Clearing \$4023.0 will immediately stop the IRQ counter and acknowledge any pending timer IRQs.
- Writing to \$4022 with bit 1 (IRQ enabled) cleared will stop the IRQ counter and acknowledge any pending timer IRQs.
- Enabling timer IRQs when the reload value is set to 0 will cause an IRQ immediately. Doing this with the repeat flag enabled will cause an infinite loop of IRQs on every CPU cycle.
- Since the disk transfer routine also uses IRQs, it's very important to disable timer IRQs before doing any access to the disk.

There are only 3 known ways to acknowledge the timer IRQ:

- Read \$4030
- Disable timer IRQs by writing to \$4022
- Disable disk registers by writing to \$4023

### <span id="Master_I.2FO_enable_.28.244023.29"></span>Master I/O enable (\$4023) {#Master_I.2FO_enable_.28.244023.29}

    7  bit  0
    ---------
    ?xxx xxSD
    |      ||
    |      |+- Enable disk I/O registers
    |      +-- Enable sound I/O registers
    +--------- Unknown (TODO)

The FDS BIOS writes \$00, then \$83 to it during reset. The purpose of bit 7 is unknown.

Disabling disk registers disables both disk and timer IRQs.

### <span id="Write_data_register_.28.244024.29"></span>Write data register (\$4024) {#Write_data_register_.28.244024.29}

The data that this register is programmed with will be the next 8-bit quantity to load into the shift register (next time the byte transfer flag raises), and to be shifted out and appear on pin 5 of the RAM adapter cable (2C33 pin 52).

Writing to this register acknowledges disk IRQs.<sup>\[***citation needed***\]</sup>

### <span id="FDS_Control_.28.244025.29"></span>FDS Control (\$4025) {#FDS_Control_.28.244025.29}

    7  bit  0
    ---------
    IE1C MRDS
    |||| ||||
    |||| |||+- Scan Disk (0: reset, 1: scan)
    |||| ||+-- Drive Motor Control (0: start, 1: stop)
    |||| |+--- Transfer Mode (0: write; 1: read)
    |||| +---- Nametable Arrangement
    ||||         0: Horizontal ("Vertical Mirroring")
    ||||         1: Vertical ("Horizontal Mirroring")
    |||+------ CRC Transfer Control (1: transfer CRC value)
    ||+------- Unknown, always set to '1'
    |+-------- CRC Enabled (0: disable/reset, 1: enable)
    +--------- Interrupt Enabled
                 1: Generate an IRQ every time the byte transfer flag is raised.

Notes:

- The BIOS and commercial software are known to always write 1 to bit 5. Writing 0 to it alters bits 1 and 2 of the external connector (\$4026/\$4033) but its purpose is currently unknown. (TODO: this supposedly also affects byte transfer IRQs?)
- CRC values are [stored on the disk](../FDS_disk_format/#CRCs) and verified when reading/writing a given file.
- Disabling the CRC resets its state. The FDS BIOS disables the CRC between file blocks, then enables it before accessing each file block to calculate/verify their CRC values.
- To change the nametable arrangement on the fly, a read-modify-write of its [mirror variable](../FDS_BIOS/#Zero-page_variables) should be done to prevent altering unrelated bits.

Writing to this register acknowledges disk IRQ.

### <span id="External_connector_.28.244026.29"></span>External connector (\$4026) {#External_connector_.28.244026.29}

Output of [expansion terminal](../FDS_expansion_port_pinout/) where there's a shutter on the back of the RAM adapter. The outputs of \$4026 (open-collector with 4.7K ohm pull-ups (except on bit 7)), are shared with the inputs on \$4033.

TODO: Setting \$4025.D5 to 0 alters bits 1 and 2 of the output, which are also reflected in \$4033. Its purpose is currently unknown.

### <span id="Disk_Status_register_.28.244030.29"></span>Disk Status register (\$4030) {#Disk_Status_register_.28.244030.29}

    7  bit  0
    ---------
    TExB Mx?D
    || | | ||
    || | | |+- Timer Interrupt (1: an IRQ occurred)
    || | | +-- Unknown Interrupt source (TODO)
    || | +---- Nametable Arrangement (from $4025.D3)
    || +------ CRC control (0: CRC passed; 1: CRC error)
    |+-------- End of Head (1 when disk head is on the most inner track)
    +--------- Byte transfer flag. Set every time 8 bits have been transferred between the RAM adaptor & disk drive (service $4024/$4031). 
               Reset when $4024, $4031, or $4030 has been serviced.

Reading this register acknowledges timer and disk IRQs.

### <span id="Read_data_register_.28.244031.29"></span>Read data register (\$4031) {#Read_data_register_.28.244031.29}

This register is loaded with the contents of an internal shift register every time the byte transfer flag raises. The shift register receives its serial data via pin 9 of the RAM adapter cable (2C33 pin 51).

Reading this register acknowledges disk IRQs.

### <span id="Disk_drive_status_register_.28.244032.29"></span>Disk drive status register (\$4032) {#Disk_drive_status_register_.28.244032.29}

    7  bit  0
    ---------
    xxxx xPRS
          |||
          ||+- Disk flag  (0: Disk inserted; 1: Disk not inserted)
          |+-- Ready flag (0: Disk readу; 1: Disk not ready)
          +--- Protect flag (0: Not write protected; 1: Write protected or disk ejected)

Notes:

- The Ready flag corresponds to the drive head's position. It is set to 1 when the head reaches the end of the disk, and cleared once it returns to the beginning of the disk.
- The Protect flag corresponds to the write protect tab present on the upper-left corner of the inserted disk side. It is set to 1 if the tab is broken.

Reading this register acknowledges disk IRQs.<sup>\[***citation needed***\]</sup>

### <span id="External_connector_read_.28.244033.29"></span>External connector read (\$4033) {#External_connector_read_.28.244033.29}

    7  bit  0
    ---------
    BIII IIII
    |||| ||||
    |+++-++++- Input from expansion terminal where there's a shutter on the back of the RAM adapter.
    +--------- Drive power/battery status (0: Voltage is low; 1: Good).

When a bit is clear in \$4026 port it will read back as '0' here (excluding battery bit) because of how open collector input works. The drive power/battery bit must be checked while the motor is on (\$4025.D1 = 0), otherwise it will always be read as 0.

### <span id="Sound_.28.244040-.244092.29"></span>Sound (\$4040-\$4092) {#Sound_.28.244040-.244092.29}

For details on sound information, see [FDS audio](../FDS_audio/).

## BIOS {#BIOS}

The FDS contains a fixed 8KB BIOS at \$E000-FFFF. This controls the Famicom at power-on and reset, dispatches the NMI and IRQ, and offers an API for accessing the [data on disk](../FDS_disk_format/). Routines for common tasks including controller reading and PPU handling are also provided for programmer convenience.

See: [FDS BIOS](../FDS_BIOS/)

## See Also {#See_Also}

- [FDS BIOS](../FDS_BIOS/)
- [FDS disk format](../FDS_disk_format/)
- [FDS file format](../FDS_file_format/) (**.FDS**)
- [QD file format](../QD/) (**.QD**)
- [FDS audio](../FDS_audio/)
- [FDS RAM adaptor cable pinout](../FDS_RAM_adaptor_cable_pinout/)
- [FDS expansion port pinout](../FDS_expansion_port_pinout/)
- [RP2C33 pinout](../RP2C33_pinout/)
- [FdsIrqTests(v7)](https://forums.nesdev.org/download/file.php?id=10240) by Sour - Test program for various elements of the FDS' IRQ
- [FDS-Mirroring-Test](https://forums.nesdev.org/download/file.php?id=28216) by TakuikaNinja - Test program for the FDS' nametable arrangement/mirroring switching functionality
- [iNES mapper 20](../INES_Mapper_020/) - Reserved for FDS dumps, but not widely used for it.
- [TNES](../TNES/) - Nintendo 3DS Virtual Console ROM format with support for .QD images.
- [GitHub repository:](https://github.com/bbbradsmith/NES-ca65-example/tree/fds) Simple FDS example for ca65
- [GitHub repository:](https://github.com/TakuikaNinja/asm6f-fds-example) Simple FDS example for asm6f
- [Forum post:](https://forums.nesdev.org/viewtopic.php?p=194826#p194826) Skipping the FDS license screen
- [FDS List](https://www.chrismcovell.com/software.html) by ccovell - command line utility to inspect FDS disk image contents.
- [FDS Lister](https://www.chrismcovell.com/fds-lister.html) by ccovell - utility to inspect FDS disk contents that runs on an FDS.
- [Triton Quick Disk Drive](https://en.wikipedia.org/wiki/Triton_Quick_Disk_Drive) - a QD drive peripheral which was sold for various home computers.

## References {#References}

- <https://www.nintendo.com/jp/famicom/hardware/disksystem.html> (Japanese)
- [FDS technical reference.txt](https://nesdev.org/FDS%20technical%20reference.txt) by Brad Taylor (old/outdated)
- [Enri's Famicom Disk System page](http://cmpslv3.stars.ne.jp/Famic/Famdis.htm) (Japanese) - Descriptions for \$4025 and \$4030 are inaccurate.
- [Enri's Famicom Disk System page](https://web.archive.org/web/20091023182159/http://www2.odn.ne.jp/~haf09260/Famic/Famdis.htm) (Japanese) (old/outdated)
- [fds-nori.txt](https://nesdev.org/fds-nori.txt) - FDS reference in Japanese by Nori (old/outdated)
- [Forum post](https://forums.nesdev.org/viewtopic.php?p=194867#p194867): .fds format: Can checksums be heuristically detected? - Includes a CRC implementation in C.
- [Forum post](https://forums.nesdev.org/viewtopic.php?f=3&t=16507): FDS IRQ reload flag/value

<!-- -->

- <https://www.famicomdisksystem.com/disks/>
- <https://famicomworld.com/workshop/tech/famicom-disk-system-fd3206-write-mod/>
- <https://famicomworld.com/workshop/tech/fds-power-board-modifications/>
