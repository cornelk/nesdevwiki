---
title: "INES Mapper 241"
url: "/wiki/INES_Mapper_241/"
description: "iNES Mapper 241 denotes a BxROM-like circuit board with the addition of 8 KiB of WRAM at CPU $6000-$7FFF that can be battery-backed. INES Mapper 034 is not commonly used for BNROM with WRAM, because…"
wikiPage: "INES_Mapper_241"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**iNES Mapper 241** denotes a [BxROM](../INES_Mapper_034/)-like circuit board with the addition of 8 KiB of WRAM at CPU \$6000-\$7FFF that can be battery-backed. [INES Mapper 034](../INES_Mapper_034/) is not commonly used for BNROM with WRAM, because the WRAM's address space is partially overlaid by registers from the NINA-001 part of the mapper specification (although that could be alleviated by disabling the NINA-001 registers in the absence of CHR-ROM).

BxROM with WRAM is typically used natively by games from Hengge Technology (Hénggé Diàn​zǐ, a.k.a. "Henggedianzi") with hard-wired mirroring and a few educational computer cartridges, but is more commonly used as a target for mapper hacks of games that originally used mappers [164](../INES_Mapper_164/), [178](../INES_Mapper_178/) and [227](../INES_Mapper_227/). Because those mapper hacks rarely take bus conflicts into consideration, mapper 241 should be emulated without bus conflicts. As a result of being mapper hacks, such games may also access registers in the \$4800-\$5FFF range.

Mirroring is hard-wired to horizontal or vertical; [INES Mapper 177](../INES_Mapper_177/) denotes a variant with mapper-controlled two-screen mirroring.

# LPC Speech Chip {#LPC_Speech_Chip}

A number of educational computer cartridges, English- as well as Chinese-language, additionally mount a [TMS5220C-data-compatible LPC speech chip](https://en.wikipedia.org/wiki/Texas_Instruments_LPC_Speech_Chips). The bitstream data is in the TMS5220C-PE LPC format and is sent four bits at a time. Refer to the [TMS5220 data sheet](https://www.sprow.co.uk/bbc/hardware/speech/tms5220.pdf), U.S. Patents [4,209,844](https://patents.google.com/patent/US4829573A/en), [4,331,836](https://patents.google.com/patent/US4331836/en) and [4,335,277](https://patents.google.com/patent/US4335277/en) as well as [MAME's source code](https://github.com/mamedev/mame/blob/master/src/devices/sound/tms5220.cpp) for details on the LPC data format and LPC decoding. Note that the hardware interface is different from that of a regular TMS5220.

## <span id="LPC_Chip_Status_.28.245000-5FFF.2C_read.29"></span>LPC Chip Status (\$5000-5FFF, read) {#LPC_Chip_Status_.28.245000-5FFF.2C_read.29}

    D~[.R.. ....]
        +--------- 1=Ready to receive data

## <span id="LPC_Data_.28.245000-5FFF.2C_write.29"></span>LPC Data (\$5000-5FFF, write) {#LPC_Data_.28.245000-5FFF.2C_write.29}

    D~[.R.R DDDD]
        | | ++++-- LPC data
        +-+------- Both set: normal operation
                   Both clear: reset chip

# Errata {#Errata}

FCEUX implements this mapper in an idiosyncratic fashion by OR'ing the 32 KiB bank number with 8 when bit 7 of the latch register is set. This serves to run an overdump of the Subor cartridge *小霸王 2合1꞉ 仓库世家 & 动脑筋*. Its proper dump runs as [NES 2.0 Mapper 481](../NES_2.0_Mapper_481/).
