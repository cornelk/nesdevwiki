---
title: "INES Mapper 014"
url: "/wiki/INES_Mapper_014/"
description: "iNES Mapper 014 denotes the 哥德 (Gouder) SL-1632 PCB, used on the 8-character version of Samurai Spirits by Rex Soft. Like INES Mapper 116, it uses the Huang-1 ASIC together with a PAL that…"
wikiPage: "INES_Mapper_014"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
---

**iNES Mapper 014** denotes the 哥德 (Gouder) **SL-1632** PCB, used on the 8-character version of *Samurai Spirits* by Rex Soft. Like [INES Mapper 116](../INES_Mapper_116/), it uses the **Huang-1** ASIC together with a PAL that provides a supervisor register for selecting CHR A18 and for switching between [MMC3](../MMC3/) and [VRC2](../VRC2_and_VRC4/) modes (A0/A1, VRC2b), but moves the supervisor register address to \$A131, overlapping the MMC3 and VRC2 registers there. Its UNIF board name is **UNL-SL1632**.

According to Nestopia's mapper source code, mirroring can only be controlled in VRC2 mode, and the VRC2 PRG and mirroring registers do not respond to address mirrors. These behaviors contradict other emulators and have not been verified on original hardware.

## <span id="Supervisor_Register_.28.24A131.2C_write.29"></span>Supervisor Register (\$A131, write) {#Supervisor_Register_.28.24A131.2C_write.29}

    Mask: $FFFF?

    D~7654 3210
      ---------
      A.B. C.M.
      | |  | +-- Mapper mode
      | |  |      0: VRC2
      | |  |      1: MMC3
      | |  +---- CHR A18 for PPU $0000-$0FFF
      | +------- CHR A18 for PPU $1000-$17FF
      +--------- CHR A18 for PPU $1800-$1FFF
