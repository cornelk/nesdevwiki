---
title: "Miracle Piano"
url: "/wiki/Miracle_Piano/"
description: "The Miracle Piano is a 49-key MIDI synthesizer built by Antex on behalf of The Software Toolworks with options for using RS-232 serial and a synchronous clocked-bit protocol as MIDI transport. It is…"
wikiPage: "Miracle_Piano"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Miracle Piano is a 49-key MIDI synthesizer built by Antex on behalf of The Software Toolworks with options for using RS-232 serial and a synchronous clocked-bit protocol as MIDI transport. It is part of [The Miracle Piano Teaching System](https://en.wikipedia.org/wiki/The_Miracle_Piano_Teaching_System), which had software implementations for at least the NES, SNES, Genesis, PC, and Macintosh.

The NES implementation of the software expects the "Miracle Cable" to be connected to controller port 1 and a standard controller to port 2, which is what this page will discuss.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Transfer direction/MIDI serial data to keyboard

The transfer direction is dictated by the strobe signal's duration, see protocol details below.

### <span id="Output_.28.244016_read.29"></span>Output (\$4016 read) {#Output_.28.244016_read.29}

    7  bit  0
    ---- ----
    xxxx xxxM
            |
            +- MIDI serial data from keyboard

Reading this register clocks the next bit in the current transfer for either direction.

## Transmission protocol {#Transmission_protocol}

The keyboard interface is bi-directional, and appears to be via clocked shift registers. Data is comprised of octets sent and received **MSB**-first.

Data is polled from the keyboard by raising OUT0 for 12 CPU cycles (about 6.7 μSec), lowering it, and clocking a bit in by reading \$4016.D0. May need to burn six cycles (about 3.3 μSec) before clocking the bit in. If this bit is set, there are an additional eight (inverted) bits to clock in. This process strobes OUT0 as usual, so polling the status of a controller connected to the other port is possible at this time.

Data is sent to the keyboard by raising OUT0 for about 79 CPU cycles (about 44.1 μSec) (from OUT0 raised to OUT0 set with the first data bit). After this, eight data bits are sent by setting OUT0 to each bit and clocking the bits out by reading \$4016.

Inter-bit time for transmission is about 16 cycles (about 8.9 μSec). Inter-bit time for reception is about 18 cycles (about 10 μSec).

## References {#References}

- [nocash's documentation](https://problemkaputt.de/everynes.htm#controllerspianomiraclepianocontrollerport) (contains information on MIDI messages and instrument patches)
- <http://weltenschule.de/TableHooters/Mindscape_MiraclePiano.html> (has some description of features and a review of the hardware).
- <http://www.pianoeducation.org/pnompfqn.html> (has some manual scans, schematics for two circuit boards in the piano itself, link for information on building replacement "miracle cables", and a FAQ).
