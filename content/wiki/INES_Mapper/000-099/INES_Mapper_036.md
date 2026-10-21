---
title: "INES Mapper 036"
url: "/wiki/INES_Mapper_036/"
description: "iNES Mapper 036 is assigned to TXC's PCB 01-22000-400, featuring a 74138, a 74175, and usually a 24-pin ASIC labeled 05-00002-010. Two games are currently known to require this definition."
wikiPage: "INES_Mapper_036"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 036 is assigned to TXC's PCB **01-22000-400**, featuring a [74138](../74138/), a 74175, and usually a 24-pin ASIC labeled [05-00002-010](../TXC_05-00002-010_pinout/). [Two games](http://glukvideo.info/listado-juegos-gluk) are currently known to require this definition.

- Strike Wolf (also released as *Comando de Lobos*)
- Policeman

Even though these games can evidently be emulated as a weird kind of GNROM, the hardware does something quite different:

The CHR banking is definitely implemented using the two 7400-series ICs:

    Mask: $E200
    write $4200: [.... CCCC] - Select 8 KiB CHR bank

The PRG banking is part of the ASIC:

    Mask: $E100
     read $4100: [xxRR xxxx]
                  |||| ||||
                  ||++------ Copy internal register 'RR' to data bus.
                  ++---++++- open bus
    Mask: $E103
     write $4100: If Increment is set, internal register 'RR' <- 'RR'+1
                  Otherwise, if Invert is clear, copy internal register 'PP' to 'RR'
                             if Invert is set, copy '~PP' to 'RR'
     write $4101: [...V ....] - Invert Mode
     write $4102: [..PP ....] - Copy data bus to internal register 'PP'. Value is not yet exposed anywhere.
     write $4103: [...C ....] - Increment Mode
    Mask: $8000
     write $8000: copy internal register 'RR' to PRG banking pins

Three additional games published by Gluk use the same PCB, with no ASIC. They were also released in other regions on different hardware, and *only* Gluk's editions are compatible with the above-mentioned CHR banking.

- F-15 City War
- Volley ball
- Puzzle

Compatibility code left in means that Gluk's editions of these three games are also emulatable as [NINA-06](../NINA-003-006/).

See also:

- <http://forums.nesdev.org/viewtopic.php?p=167180#p167180> Thread about dumping Policeman starts here
- The same ASIC showed up in [iNES Mapper 132](../INES_Mapper_132/) and [iNES Mapper 173](../INES_Mapper_173/), wired differently.
