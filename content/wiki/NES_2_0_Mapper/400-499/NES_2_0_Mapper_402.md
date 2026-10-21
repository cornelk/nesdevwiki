---
title: "NES 2.0 Mapper 402"
url: "/wiki/NES_2.0_Mapper_402/"
description: "NES 2.0 Mapper 402 denotes the J-2282 circuit board, used by the 22-in-1 Olympic Games multicart with 512 KiB PRG-ROM and unbanked 8 KiB CHR-RAM."
wikiPage: "NES_2_0_Mapper_402"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 402** denotes the **J-2282** circuit board, used by the *22-in-1 Olympic Games* multicart with 512 KiB PRG-ROM and unbanked 8 KiB CHR-RAM.

    A~[1... FC.. MN.P PPPp]
            ||   || +-++++- PRG A18..A14
            ||   |+-------- 0: PRG A14=CPU A14 (NROM-256)
            ||   |          1: PRG A14=p (NROM-128)
            ||   +--------- 0: Vertical mirroring
            ||              1: Horizontal mirroring
            |+------------- 0: CHR-RAM write-protected
            |               1: CHR-RAM write-enabled
            +-------------- 0: CPU $6000-$7FFF open bus
                            1: CPU $6000-$7FFF mirror of $E000-$FFFF

The F bit is set to play the FDS-sourced version of *Volleyball*.
