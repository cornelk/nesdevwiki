---
title: "INES Mapper 006"
url: "/wiki/INES_Mapper_006/"
description: "iNES Mapper 006 denotes ROM images that have been extracted from disk images for the Front Fareast Magic Card 1M or 2M RAM cartridges. They represent games whose Doctor Header file denotes a Magic…"
wikiPage: "INES_Mapper_006"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

**iNES Mapper 006** denotes ROM images that have been extracted from disk images for the *Front Fareast Magic Card 1M* or *2M* [RAM cartridges](../RAM_cartridge/). They represent games whose [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Magic_Card_1M.2F2M.2F4M_disks) denotes a Magic Card 1M/2M disk (byte \$0 bit 7 set, bits 4 and 5 clear, byte \$7=\$00). Refer to the [Super Magic Card](../Super_Magic_Card/) article for details on bankswitching. The Super Magic Card's registers are initialized to:

    ; Play mode, WRAM bank 0, 1 KiB CHR mode disabled
    $4500 = $42 

    ; PRG memory write-protected, two-screen mirroring
    $42FF = (submapper <<5) | (verticalMirroring? 0x00: 0x10)

    ; 2M/4M banking mode disabled
    $43FF = $00 

The [NES 2.0](../NES_2.0/) Submapper field denotes the initial [latch-based banking mode](../Super_Magic_Card/#Latch-based_modes) (0-7). [NES 1.0](../INES/) files correspond to submapper 1. **iNES [Mapper 8](./)** is a synonym of Mapper 6 submapper 4.

The [iNES](../INES/) header may specify a [512-byte trainer](../INES/#Trainer) (corresponding to [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Magic_Card_1M.2F2M.2F4M_disks)'s byte \$0 bit 6 being set), which must be loaded to \$7000-\$71FF, be writable, and (on a hard reset) initialized by JSRing to \$7003 before JMPing to the game's reset vector.

Battery-saving of WRAM content is not supported by any Magic Card model. Hard-resetting a game while restoring previously-saved WRAM content in emulators interferes with the correct operation of the trainer's program.

# See also {#See_also}

- [Mapper 6](https://nesdev.org/mapper6.txt) Info on the FFE mapper. By FanWen Yang (outdated).
- [Info on various Famicom "copiers"](http://www.famicomdisksystem.com/game-doctor-copiers/)
