---
title: "INES Mapper 207"
url: "/wiki/INES_Mapper_207/"
description: "iNES Mapper 207 represents the board used for Fudou Myouou Den."
wikiPage: "INES_Mapper_207"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 207 represents the board used for *Fudou Myouou Den*.

It modifies Taito's [X1-005](../Taito_X1-005/) ([80](../Taito_X1-005/)) boards in the exact same way that [TLSROM](../TLSROM/) ([118](../INES_Mapper_118/)) differs from [MMC3](../MMC3/) ([4](../MMC3/)), by which we mean the X1-005's CHR A17 output is connected to CIRAM A10. Thus, it achieves 1scA /1scB/H controllable mirroring instead of H/V.

     ========================
     =  Mapper 207          =
     ========================

      Registers:
     ---------------------------

       $7EF0:  [MCCC CCCC]
          M = Mirroring 0
          C = CHR Reg 0

       $7EF1:  [MCCC CCCC]
          M = Mirroring 1
          C = CHR Reg 1

       $7EF2-7EF5:  CHR Regs 2-5

       $7EF6:  was old mirroring control register, now ignored

       $7EF8-7EFF:  see iNES Mapper 080

     CHR Setup:
     ---------------------------

            $0000   $0400   $0800   $0C00   $1000   $1400   $1800   $1C00 
          +---------------+---------------+-------+-------+-------+-------+
          |    <$7EF0>    |    <$7EF1>    | $7EF2 | $7EF3 | $7EF4 | $7EF5 |
          +---------------+---------------+-------+-------+-------+-------+

     Mirroring:
     ---------------------------

       [ $7EF0 ][ $7EF0 ]
       [ $7EF1 ][ $7EF1 ]

     Mirroring bit of appropriate reg selects NTA or NTB
