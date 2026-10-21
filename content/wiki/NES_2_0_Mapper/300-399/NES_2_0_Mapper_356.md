---
title: "NES 2.0 Mapper 356"
url: "/wiki/NES_2.0_Mapper_356/"
description: "NES 2.0 Mapper 356 is used for the J.Y. Company's 7-in-1 Rockman (JY-208) multicart. It is basically INES Mapper 045 but with added functionality to switch between CHR-ROM and CHR-RAM, and between…"
wikiPage: "NES_2_0_Mapper_356"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 356 is used for the J.Y. Company's *7-in-1 Rockman* (JY-208) multicart. It is basically [INES Mapper 045](../INES_Mapper_045/) but with added functionality to switch between CHR-ROM and CHR-RAM, and between normal two-screen and four-screen mirroring.

All registers work as [INES Mapper 045](../INES_Mapper_045/), except \$6000 sequential **register 2** (third write):

    D~7654 3210
      ---------
      .4RC SSSS
       ||| ++++ CHR-AND block size (see Mapper 45 description)
       ||+----- High bit of CHR-OR (CHR A18, see Mapper 45 description)
       |+------ CHR memory type select
       |         0: CHR-RAM (8 KiB)
       |         1: CHR-ROM
       +------- Nametable memory select
                 0: CIRAM (two nametables, mirrored according to MMC3 setting)
                 1: Cartridge VRAM (four nametables)

Because unlike the original Mapper 45, there is only one high CHR-OR bit, only 512 KiB of CHR-ROM are supported.
