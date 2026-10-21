---
title: "INES Mapper 212"
url: "/wiki/INES_Mapper_212/"
description: "According to FCEUX's source, iNES Mapper 212 is yet another simple discrete logic pirate multicart mapper."
wikiPage: "INES_Mapper_212"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with overlapping PRG and CHR registers"
    url: "/category/Mappers_with_overlapping_PRG_and_CHR_registers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

According to FCEUX's source, iNES Mapper 212 is yet another simple discrete logic pirate multicart mapper.

Nestopia calls it "BMC Super HiK 300-in-1".

    Write $8000-$FFFF:
     A~[1o.. .... .... MBBb]
         |             ||||
         |             |+++-- Combined PRG and CHR bank number
         |             +----- Nametable mirroring (0:V/A10, 1:H/A11)
         +------------------- Banking style

When Banking style is 0, BBb specifies a 16 KiB PRG bank at both CPU \$8000 and \$C000. When it's 1, BB is 32 KiB PRG bank at CPU \$8000.

Regardless of the value of the "o" bit, BBb is 8 KiB CHR bank at PPU \$0000.

Equivalently, the latched copy of A2 through A0 goes to PRG A16 through A14 and CHR A15 through A13.

    Mask: $E010
    Read: $6000: [1... PPPP]
                       ++++-- Solder pad value, used for menu selection
    Read: $6010: [0... PPPP]
                       ++++-- Solder pad value, used for menu selection

Reading from \$6000-\$7FFF must return the inverse of A4 in D7, meaning when reading \$6000, data bit 7 must be set, and when reading \$6010, data bit 7 must be clear.

This is reminiscent of [iNES Mapper 107](../INES_Mapper_107/).
