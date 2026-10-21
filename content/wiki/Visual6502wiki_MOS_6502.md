---
title: "Visual6502wiki/MOS 6502"
url: "/wiki/Visual6502wiki/MOS_6502/"
description: "We have photographs of the metal and lower layers, the polygons captured, the circuit extracted and we have published a javascript simulator. There is an FPGA project to implement the simulation in…"
wikiPage: "Visual6502wiki_MOS_6502"
---

## MOS 6502 family {#MOS_6502_family}

We have [photographs of the metal and lower layers](../Photos_of_MOS_6502D/), the polygons captured, the circuit extracted and we have published [a javascript simulator](http://visual6502.org/JSSim). There is an [FPGA project](https://github.com/pmonta/FPGA-netlist-tools) to implement the simulation in hardware. We have (as yet) unpublished simulators in python and C.

## Our Analysis {#Our_Analysis}

Here are some relatively raw materials we've collected:

- Comparative photos of the [Stack Register](../6502_Stack_Register_High_Bits/) in 6502 and 6507
- The [Decode ROM](../6507_Decode_ROM/) (describing the Atari 6507, not exactly the same as the NMOS 6502 used in the [visual6502 simulator](http://visual6502.org/JSSim)
- [All 256 6502 opcodes](../6502_all_256_Opcodes/) named and tabulated

And here is some more interpretive material from our explorations:

- Collected [observations](../6502Observations/) of 6502 layout and behaviour.
- The 6502 [datapath timing](../6502_datapath/)
- The [unsupported opcodes](../6502_Unsupported_Opcodes/)
  - a detailed explanation of the [XAA opcode](../6502_Opcode_8B_%28XAA%2C_ANE%29/) behaviour
- [Implementing a realtime netlist simulation](../6502_-_simulating_in_real_time_on_an_FPGA/) in historical systems using an FPGA

## 6502 additional information {#6502_additional_information}

See also the [links](http://visual6502.org/links.html) page on the main site.

### Primary Sources {#Primary_Sources}

- [Visual6502wiki/Photos of MOS 6502D](../Photos_of_MOS_6502D/) (also see our [website](http://visual6502.org/images/6502/index.html))
- [Visual6502wiki/Atari's 6507 Schematics](../Atari%27s_6507_Schematics/)
- [Visual6502wiki/Photos of R6502](../Photos_of_R6502/)

### Secondary Sources {#Secondary_Sources}

- [Visual6502wiki/Hanson's Block Diagram](../Hanson%27s_Block_Diagram/)
- [Visual6502wiki/Balazs' schematic and documents](../Balazs%27_schematic_and_documents/)

### Previous Analysis {#Previous_Analysis}

- [Beregnyei Balazs: 6502 Reverse Engineering](http://impulzus.sch.bme.hu/6502/letolt.php3) ([translation](http://www.downloads.reactivemicro.com/Public/Electronics/Reverse%20Engineering/))
- [Mark Ormston: 65xx Processor Data (version 0.2b)](http://anyplatform.net/media/guides/cpus/65xx%20Processor%20Data.txt)
- [Ivo van Poorten: 6502 Bugs List](http://www.textfiles.com/apple/6502.bugs.txt)
- [Neil Parker: The 6502/65C02/65C816 Instruction Set Decoded](http://www.llx.com/~nparker/a2/opcodes.html)
- [Graham: 6502/6510/8500/8502 Opcode matrix](http://www.oxyron.de/html/opcodes02.html)
- [Freddy Offenga: 6502 Undocumented Opcodes](http://members.chello.nl/taf.offenga/illopc31.txt)
- [Adam Vardy: Extra Instructions Of The 65XX Series CPU](http://www.zimmers.net/anonftp/pub/cbm/documents/chipdata/6502-NMOS.extra.opcodes)
