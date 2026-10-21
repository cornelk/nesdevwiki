---
title: "INES Mapper 084"
url: "/wiki/INES_Mapper_084/"
description: "iNES Mapper 084 appears to have two conflicting allocations."
wikiPage: "INES_Mapper_084"
wikiCategories:
  - name: "Bad iNES Mappers"
    url: "/category/Bad_iNES_Mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

**iNES Mapper 084** appears to have two conflicting allocations.

## Pasofami {#Pasofami}

[neshdr20.txt by VmprHntrD](https://nesdev.org/neshdr20.txt) claims the mapper as reserved for Pasofami, an emulator made by Nobuyuki Andou. [MAME's source](https://github.com/mamedev/mame/blob/2cd762097742c10d8e784086d4ffa056d073a389/src/devices/bus/nes/nes_ines.hxx#L119) reflects this but does not include any code to handle it.

## PC-SMB2J {#PC-SMB2J}

fwNES 0.30 replaced the Pasofami mapper assignment with "PC-SMB2J" (PC = Pirate Cart) used by "SMBJ2".<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> The name is also given as "PC-Super Mario Bros 2 J Custom Mapper" and "Super Mario Bros. 2J (HK hacked)" in other mapper lists for the emulator.<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)<span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

Details are unknown, however the implementation in fwNES appears to be different from other FDS conversion mappers such as iNES Mapper [40](../INES_Mapper_040/), [50](../INES_Mapper_050/) or [142](../INES_Mapper_142/).

(TODO: [Register documentation](https://forums.nesdev.org/viewtopic.php?p=310019#p310019) & comparison with equivalent handling in Pasofami - it's possible that this is still the Pasofami mapper implementation)

At some point, FCEUX lost support for this mapper.<sup>\[***citation needed***\]</sup>

Currently a few registers of fwNES' implementation were discovered:

    Bank Map:

    CPU $6000-$7FFF - PRG-RAM Bank
    CPU $8000-$9FFF - Switchable PRG-ROM Bank
    CPU $A000-$BFFF - Switchable PRG-ROM Bank
    CPU $C000-$DFFF - Switchable PRG-ROM Bank
    CPU $E000-$FFFF - Switchable PRG-ROM Bank

    PPU $0000-$03FF - Switchable CHR-ROM Bank
    PPU $0400-$07FF - Switchable CHR-ROM Bank
    PPU $0800-$0BFF - Switchable CHR-ROM Bank
    PPU $0C00-$0FFF - Switchable CHR-ROM Bank
    PPU $1000-$13FF - Switchable CHR-ROM Bank
    PPU $1400-$17FF - Switchable CHR-ROM Bank
    PPU $1800-$1BFF - Switchable CHR-ROM Bank
    PPU $1C00-$1FFF - Switchable CHR-ROM Bank

    Registers (Write):
    Unless otherwise mentioned, these register aren't mirrored at any other addresses.

    $8000 - Set 8 KiB PRG-ROM Bank at CPU $8000-$9FFF
    $8001 - Set 8 KiB PRG-ROM Bank at CPU $A000-$BFFF
    $8002 - Set 8 KiB PRG-ROM Bank at CPU $C000-$DFFF
    $8003 - Set 8 KiB PRG-ROM Bank at CPU $E000-$FFFF

    $A000 - Set 1 KiB CHR-ROM Bank at PPU $0000-$03FF
    $A001 - Set 1 KiB CHR-ROM Bank at PPU $0400-$07FF
    $A002 - Set 1 KiB CHR-ROM Bank at PPU $0800-$0BFF
    $A003 - Set 1 KiB CHR-ROM Bank at PPU $0C00-$0FFF
    $A004 - Set 1 KiB CHR-ROM Bank at PPU $1000-$13FF
    $A005 - Set 1 KiB CHR-ROM Bank at PPU $1400-$17FF
    $A006 - Set 1 KiB CHR-ROM Bank at PPU $1800-$1BFF
    $A007 - Set 1 KiB CHR-ROM Bank at PPU $1C00-$1FFF

    $C002 - IRQ Disable (?)
    $C003 - IRQ Enable (?)

    $D000 - Set Nametable Arrangement

    $E000 - IRQ Count Enable (Bit 0 = Enable) (?)

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) WHATSNEW.TXT & FWNES98E.TXT, fwNES 0.302b, 1998/11/16, by FanWen Yang
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) <https://web.archive.org/web/20000818031458fw_/http://fwnes.vintagegaming.com/mappers.htm>
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) <https://web.archive.org/web/20000818031534fw_/http://fwnes.vintagegaming.com/faq-mppr.txt>
