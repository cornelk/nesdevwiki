---
title: "INES Mapper 116"
url: "/wiki/INES_Mapper_116/"
description: "iNES Mapper 116 denotes Gouder's SOMARI-P circuit board bearing the Huang-1 or Huang-2 ASICs and a PAL, used by releases from Gouder, also known as 哥德 (Gēdé):"
wikiPage: "INES_Mapper_116"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
---

**iNES Mapper 116** denotes Gouder's **SOMARI-P** circuit board bearing the Huang-1 or Huang-2 ASICs and a PAL, used by releases from Gouder, also known as 哥德 (Gēdé):

- Huang-1 chip (**submappers 0/1**):
  - *A.V. 美少女戰士 Girl Fighting* (A.V. Měi Shàonǚ Zhànshì)
  - *餓狼伝説 Special* (Garō Densetsu Special, original Gouder release only)
  - *Kart Fighter*
  - *Somari the Adventurer* (original Gouder and Kǎshèng releases, two revisions)
- Huang-2 chip (**submapper 2**):
  - *A.V. 究極麻將 II* (A.V. Jiūjí Májiàng, original Gouder release only)
- Reset-based multicart with Huang-1 chip (**submapper 3**):
  - *瑪莉家族大集合 5-in-1* (Mario 5-in-1)

This mapper can simulate three different ASICs: the Nintendo [MMC1](../MMC1/), [MMC3](../MMC3/) and Konami [VRC2](../VRC2_and_VRC4/) (A0/A1, i.e. VRC2b). All three simulated chips have their separate registers, so that when switching between modes using the supervisor register, the entire mapper state changes, including PRG/CHR-ROM switchable and fixed banks as well as IRQ state.

## <span id="Supervisor_Register_.28.244100.2C_write.29"></span>Supervisor Register (\$4100, write) {#Supervisor_Register_.28.244100.2C_write.29}

    D~[.... .CMM] (Address & $E100)==$4100    
             |++- Mapper mode
             |     0: VRC2
             |     1: MMC3
             |     2: MMC1
             |     3: same as 2
             +--- Outer 256 KiB CHR-ROM bank (CHR A18)

Power-on value: \$01 (MMC3 mode)

## Notes {#Notes}

- **Submapper 0** denotes the -W configuration of the PAL that resets the MMC1's serial shift register when switching *to* MMC1 mode, equivalent to writing to \$8000-\$FFFF with bit 7 set. All single-game cartridges except for the GK2 version of *Somari the Adventurer* behave like that.
- **Submapper 1** denotes the -P configuration of the PAL that does not reset the MMC1's serial shift register when switching *to* MMC1 mode. Only the GK2 version of *Somari the Adventurer* behaves like that, and that game does not seem to rely on it.
- **Submapper 2** denotes the Huang-2 version of the ASIC. Its MMC1 addresses 8 KiB rather than 16 KiB PRG-ROM banks, meaning that the PRG bank register value is shifted by one bit to the left compared to what one would write to a normal MMC1.
- **Submapper 3** denotes a reset-based multicart with the first game having 256 KiB of PRG- and CHR-ROM each, the remaining ones 128 KiB of PRG- and CHR-ROM each.
- The VRC2 CHR-ROM registers are initialized on power-up to \$FF. One version of *Somari the Adventurer* relies on this.
- *Garō Densetsu Special* relies on being able to write to [mirrors](../Mirroring/#Memory_Mirroring) of \$8xxx, \$9xxx, and \$Axxx in VRC2 mode.
- The Kǎshèng rereleases of *A.V. 究極麻將 II* and *Garō Densetsu Special* use [INES Mapper 115](../INES_Mapper_115/) instead.

## Similar mappers {#Similar_mappers}

- [INES Mapper 014](../INES_Mapper_014/) and [INES Mapper 238](../INES_Mapper_238/) use the same ASIC but with the PAL programmed differently to respond in other ways to other addresses.
