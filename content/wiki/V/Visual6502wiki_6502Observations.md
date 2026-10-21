---
title: "Visual6502wiki/6502Observations"
url: "/wiki/Visual6502wiki/6502Observations/"
description: "We've found some interesting things on the 6502, from the layout level, up through circuit level to the programmer visible level."
wikiPage: "Visual6502wiki_6502Observations"
---

We've found some interesting things on the 6502, from the layout level, up through circuit level to the programmer visible level.

# Programmer Visible {#Programmer_Visible}

Notes here on bugs and undocumented behaviour.

- [BRK, the B bit](../6502_BRK_and_B_bit/), and other interrupts
- [Timing of Interrupt Handling](../6502_Timing_of_Interrupt_Handling/) noting that a taken branch delays interrupt handling, also that CLI/PLP allow one further instruction to execute, unlike RTI.
- [Unsupported or undocumented opcodes](../6502_Unsupported_Opcodes/) such as SAX and XAA
- [The ROR bug](https://www.pagetable.com/?p=406) which is found only in rare early devices

See also [our catalogue of 6502 test programs](../6502TestPrograms/), useful to verify simulators or emulators.

# Circuit and Logic {#Circuit_and_Logic}

Notes here on timing fixes and non-digital circuit techniques, and departures from NMOS design style orthodoxy.

- [Signs of a fix](../6502_datapath_control_timing_fix/) to datapath control timing

# Layout {#Layout}

Notes here on the traces of bug fixes, and remnants of the original 6501 layout.

- [Traces in the layout](../6502_traces_of_6501/) of the original 6501 part which was withdrawn after legal wrangling
