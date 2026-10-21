---
title: "FDS RAM adaptor cable pinout"
url: "/wiki/FDS_RAM_adaptor_cable_pinout/"
description: "The FDS RAM adapter requires a cable to be connected to the disk drive (physical or emulated) in order to facilitate disk transfers."
wikiPage: "FDS_RAM_adaptor_cable_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

|  |  |
|----|----|
| <img src="../../wiki-images/Edit-clear.svg" width="40" height="40" /> | This article is pending updates to verify the pin functions (based on documentation for other devices using QD drives) and the FDS register mappings. |

The [FDS](../Family_Computer_Disk_System/) RAM adapter requires a cable to be connected to the disk drive (physical or emulated) in order to facilitate disk transfers.

## Diagram {#Diagram}

Open-end view of the RAM adapter's disk drive connector.

     ████████████████████████████████
     ████████████████████████████████
     ███  1   3   5   7   9   11  ███
     ███                          ███
     ███  2   4   6   8   10  12  ███
     \██████████████████████████████/
      \████████████████████████████/

## Pin meanings {#Pin_meanings}

| pin \# | 2C33 register | \*[2C33 pin](../RP2C33_pinout/) | \*RAM pins | I/O | signal description |
|----|----|----|----|----|----|
| 1 | \$4025.2 | 50 | 5 (green) | Output | /write gate (TODO: active low or high?) |
| 2 | \- | 64 | 12 (cyan) | Output | VCC (+5VDC) |
| 3 | \$4025.1 | 48 | 7 (violet) | Output | /motor on |
| 4 | \- | 32 | 1 (brown) | Output | VEE (ground) |
| 5 | \$4024<sup><span id="cite_ref-shiftreg_1-0"></span>[\[1\]](#cite_note-shiftreg-1)</sup> | 52 | 3 (orange) | Output | /write data (TODO: active low or high?) |
| 6 | \$4033.7 | 37 | 11 (pink) | Input | busy |
| 7 | \$4032.2 | 47 | 8 (grey) | Input | write protect (TODO: active low or high?) |
| 8 | \- | \- | \- | Input | motor VCC† |
| 9 | \$4031<sup><span id="cite_ref-shiftreg_1-1"></span>[\[1\]](#cite_note-shiftreg-1)</sup> | 51 | 4 (yellow) | Input | read data |
| 10 | \$4032.0 | 45 | 10 (black) | Input | /media set |
| 11 | \$4032.1 | 46 | 9 (white) | Input | /ready |
| 12 | \$4025.0 | 49 | 6 (blue) | Output | /reset |

    notes on symbols
    ----------------
    I/O: input/output
    / : Indicates a signal which is active on a low (0) condition.
    * : These are corresponding pinouts for the 2C33 I/O chip, and the other  end  
    of the RAM adaptor cable, which both are located inside the RAM adaptor.
    † : The RAM adaptor does not use this signal (there is no wire in the cable 
    to carry the signal). An electronically controlled 5-volt power supply 
    inside the disk drive unit generates the power that appears here. This power 
    is also shared with the drive's internal electric motor. Therefore, the 
    motor only comes on when there is voltage on this pin.

## Signal Descriptions {#Signal_Descriptions}

### <span id="Pin_1_.28Output.29_.2Fwrite_gate_.28.244025.2.29"></span>Pin 1 (Output) /write gate (\$4025.2) {#Pin_1_.28Output.29_.2Fwrite_gate_.28.244025.2.29}

While active, this signal indicates that data appearing on the "/write data" pin is to be written to the storage media.

### <span id="Pin_3_.28Output.29_.2Fmotor_on_.28.244025.1.29"></span>Pin 3 (Output) /motor on (\$4025.1) {#Pin_3_.28Output.29_.2Fmotor_on_.28.244025.1.29}

This signal controls the drive motor by default.

TODO: If "/write gate" is active, \$4025.D0 ("/reset"?) seems to forcefully stop the motor and disable the write logic.

### <span id="Pin_5_.28Output.29_.2Fwrite_data"></span>Pin 5 (Output) /write data {#Pin_5_.28Output.29_.2Fwrite_data}

This is the serial data the RAM adaptor issues to be written to the storage media on the "/write gate" condition.

### <span id="Pin_6_.28Input.29_busy_.28.244033.7.29"></span>Pin 6 (Input) busy (\$4033.7) {#Pin_6_.28Input.29_busy_.28.244033.7.29}

After the RAM adaptor issues a "/motor on" signal, the BIOS will check the status of this input to see if the disk drive motor has turned on. If this input is found to be inactive, the BIOS interprets this as the disk drive's batteries having failed. Essentially, this signal's operation is identical to the above mentioned "motor VCC" signal, except that this is a TTL signal version of it.

TODO: This signal may be controlling the drive LED as well?

### <span id="Pin_7_.28Input.29_write_protect_.28.244032.2.29"></span>Pin 7 (Input) write protect (\$4032.2) {#Pin_7_.28Input.29_write_protect_.28.244032.2.29}

When active, this signal indicates to the RAM adaptor that the current media is write protected.

### <span id="Pin_9_.28Input.29_read_data"></span>Pin 9 (Input) read data {#Pin_9_.28Input.29_read_data}

When "/motor on" is active, data that is progressively read off the storage media is expected to appear here.

### <span id="Pin_10_.28Input.29_.2Fmedia_set_.28.244032.0.29"></span>Pin 10 (Input) /media set (\$4032.0) {#Pin_10_.28Input.29_.2Fmedia_set_.28.244032.0.29}

Also called "/media sw (switch)". When active, this signal indicates the presence of storage media.

### <span id="Pin_11_.28Input.29_.2Fready_.28.244032.1.29"></span>Pin 11 (Input) /ready (\$4032.1) {#Pin_11_.28Input.29_.2Fready_.28.244032.1.29}

The falling edge of this signal would indicate to the RAM adaptor that the disk drive has acknowledged the "/motor on" signal, and the disk drive head is currently at the beginning of the disk (most outer track). While this signal remains active, this indicates that the disk head is advancing across the disk's surface, and appropriate data can be transferred to/from the disk. This signal would then go inactive if the head advances to the end of the disk (most inner track), or if the "/motor on" signal becomes inactive.

### <span id="Pin_12_.28Output.29_.2Freset_.28.244025.0.29"></span>Pin 12 (Output) /reset (\$4025.0) {#Pin_12_.28Output.29_.2Freset_.28.244025.0.29}

While active, this instructs the storage media pointer to be reset (and stay reset) at the beginning of the media. When inactive, the media pointer is to be advanced at a constant rate, and data progressively transferred to/from the media (via the media pointer).

TODO: The "/reset" logic specified in devices like the Yamaha MDF1 seems simpler than this "media pointer" nonsense. It probably just stops the motor, blocks the write logic, and makes "/ready" go inactive.

## Notes {#Notes}

1.  <span id="cite_note-shiftreg-1"></span>↑ <sup>[1.0](#cite_ref-shiftreg_1-0)</sup> <sup>[1.1](#cite_ref-shiftreg_1-1)</sup> Parallel/serial conversion is done via a pair of shift registers inside 2C33

## References {#References}

- [Forum thread](https://forums.nesdev.org/viewtopic.php?t=26678): FDS/Mitsumi Quick Disk Drive
