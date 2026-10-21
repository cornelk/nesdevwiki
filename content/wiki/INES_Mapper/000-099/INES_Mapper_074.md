---
title: "INES Mapper 074"
url: "/wiki/INES_Mapper_074/"
description: "iNES Mapper 074 denotes the 43-393/43-406/860908C PCB, used by several games from Waixing:"
wikiPage: "INES_Mapper_074"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**iNES Mapper 074** denotes the **43-393**/**43-406**/**860908C** PCB, used by several games from Waixing:

- *機甲戰士* (Jījiǎ Zhànshì, Chinese translation of Data East's *Metal Max*)
- *甲A - China Soccer League for Division A*
- *第四次: 机器人大战 - Robot War IV*
- *風雲 - Traitor Legend* (original 1997 version)

The circuit board mounts an [MMC3](../MMC3/) clone together with a 74LS138 and 74LS139 to redirect 1 KiB CHR-ROM banks \#8 and \#9 to 2 KiB of CHR-RAM.

## Notes {#Notes}

- SMYNES and older versions of VirtuaNES EX/Plus use mapper 74 to denote *any* MMC3-based mixed CHR-ROM/CHR-RAM game by simply making *all* of CHR-ROM writable, even if those games actually use [191](../INES_Mapper_191/), [192](../INES_Mapper_192/), [194](../INES_Mapper_194/), or [195](../INES_Mapper_195/). This approach fails in the case of *甲A - China Soccer League for Division A*, as the game writes all zeros to whatever CHR-ROM banks are mapped at power-on, overwriting CHR-ROM data in the process. As a result, the publicly available dump of that game (PRG-ROM CRC32 0x78FB3ED6) has been modified to remove those writes.
- Nestopia emulates all of Waixing's MMC3 clones with a slight variation on the mirroring control register at \$A000, additionally treating a written value of 2 or 3 as 1scA and 1scB mirroring respectively.
- iNES Mapper [252](../INES_Mapper_252/) and [253](../INES_Mapper_253/), also used for Waixing's localizations, are similar but use a VRC4 clone instead.

## See also {#See_also}

- [MMC3 with CHR ROM and CHR RAM](../MMC3_with_CHR_ROM_and_CHR_RAM/)
