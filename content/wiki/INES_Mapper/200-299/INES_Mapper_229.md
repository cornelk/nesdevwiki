---
title: "INES Mapper 229"
url: "/wiki/INES_Mapper_229/"
description: "iNES Mapper 229 aka \"BMC 31-IN-1\" is a simple mapper. Writing to anywhere in $8000-FFFF will affect both mirroring and its mapping for CHR and both PRG pages."
wikiPage: "INES_Mapper_229"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 229 aka "BMC 31-IN-1" is a simple mapper. Writing to anywhere in \$8000-FFFF will affect both mirroring and its mapping for CHR and both PRG pages.

## Banks {#Banks}

- PRG: 2x16KiB
- CHR: 1x8KiB

## Register {#Register}

    Mask $803F, range $8000-FFFF
    Address:
    %1xxx xxxx xxHP PPPP
                  C CCCC
     H = mirroring [0=vertical, 1=horizontal]
     P,C = bits for mapping both CHR and PRG, all pages

This means that both PRG pages will be loaded with the same bank. However, if all M bits are 0, then the second page of PRG gets mapped to 1 instead of 0.

Nestopia does not seem to restrict the CHR banking to the lower five bits, potentially using all 16. Otherwise, Nintendulator does the same.

## Sources {#Sources}

- [Nestopia implementation](https://github.com/rdanbrook/nestopia/blob/88d130fd083b9662ef49e8d5ef95513f4bb8759e/source/core/board/NstBoardBmc31in1.cpp)
- [Nintendulator implementation](https://sourceforge.net/p/nintendulator/code/1335/tree/mappers/trunk/src/iNES/mapper229.cpp)
