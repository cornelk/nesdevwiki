---
title: "INES Mapper 177"
url: "/wiki/INES_Mapper_177/"
description: "iNES Mapper 177 is used for games from 恒格电子 (Hénggé Diànzǐ). It's basically oversize BNROM without bus conflicts, with an additional mirroring control, and with 8-32 KiB of battery-backed…"
wikiPage: "INES_Mapper_177"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

iNES Mapper 177 is used for games from 恒格电子 (Hénggé Diànzǐ). It's basically oversize [BNROM](../INES_Mapper_034/) without bus conflicts, with an additional mirroring control, and with 8-32 KiB of battery-backed WRAM at CPU \$6000-\$7FFF. **Submapper 1** defines a variant used by the *Nikita PC Vivaz* educational computer cartridge, in which bit 5 gains an alternative function. Some cartridges also have a TMS5220-compatibel LPC speech chip at CPU \$5000-\$5FFF.

                7  bit  0
                ---------
    $8000-FFFF: WWMP PPPP
                |||+-++++- Select 32 KiB PRG-ROM bank at CPU $8000-$FFFF
                ||+------- Submapper 0: Select nametable mirroring
                ||         0: Vertical
                ||         1: Horizontal
                ||         Submapper 1: Enable 16 KiB CHR-RAM mode
                ++-------- Select 8 KiB PRG-RAM bank at CPU $6000-$7FFF

- Some Hengedianzi games use hard-wired mirroring and do not set the M bit correctly; they are set to [INES Mapper 241](../INES_Mapper_241/) instead.
- [NES 2.0 Mapper 329](../NES_2.0_Mapper_329/) was originally defined as mapper 177 with more than 8 KiB of PRG-RAM, but this is redundant, as specifying only 8 KiB of PRG-RAM in the NES 2.0 header will effectively render the WW bits nonfunctional.
- In 16 KiB CHR-RAM mode, PPU \$0000-\$0FFF is automatically switched between four 4 KiB CHR-RAM banks based on PPU A8-A9 during nametable accesses, while PPU \$1000-\$1FFF is always mapped to the last 4 KiB CHR-RAM bank.

# LPC Speech Chip {#LPC_Speech_Chip}

One educational computer cartridge (*英语发声词霸 - ABC 声霸*) additionally mounts a [TMS5220C-data-compatible LPC speech chip](https://en.wikipedia.org/wiki/Texas_Instruments_LPC_Speech_Chips). The bitstream data is in the TMS5220C-PE LPC format and is sent four bits at a time. Refer to the [TMS5220 data sheet](https://www.sprow.co.uk/bbc/hardware/speech/tms5220.pdf), U.S. Patents [4,209,844](https://patents.google.com/patent/US4829573A/en), [4,331,836](https://patents.google.com/patent/US4331836/en) and [4,335,277](https://patents.google.com/patent/US4335277/en) as well as [MAME's source code](https://github.com/mamedev/mame/blob/master/src/devices/sound/tms5220.cpp) for details on the LPC data format and LPC decoding. Note that the hardware interface is different from that of a regular TMS5220.

## <span id="LPC_Chip_Status_.28.245000-5FFF.2C_read.29"></span>LPC Chip Status (\$5000-5FFF, read) {#LPC_Chip_Status_.28.245000-5FFF.2C_read.29}

    D~[R... ....]
       +--------- 1=Ready to receive data

## <span id="LPC_Data_.28.245000-5FFF.2C_write.29"></span>LPC Data (\$5000-5FFF, write) {#LPC_Data_.28.245000-5FFF.2C_write.29}

    D~[.R.R DDDD]
        | | ++++-- LPC data
        +-+------- Both set: normal operation
                   Both clear: reset chip
