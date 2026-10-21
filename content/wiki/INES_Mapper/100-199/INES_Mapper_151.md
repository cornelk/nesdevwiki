---
title: "INES Mapper 151"
url: "/wiki/INES_Mapper_151/"
description: "iNES Mapper 151 represents the VRC1 on the Vs. System. Given that the iNES format has a bit to mark \"For Vs. System\", this assignment was erroneous and any games marked thus should be mapper 75…"
wikiPage: "INES_Mapper_151"
wikiCategories:
  - name: "Bad iNES Mappers"
    url: "/category/Bad_iNES_Mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

iNES Mapper 151 represents the [VRC1](../VRC1/) on the [Vs. System](../Vs._System/). Given that the iNES format [has a bit to mark](../INES/#Flags_7) "For Vs. System", this assignment was erroneous and any games marked thus should be **mapper 75** instead.

Several emulators treat this as "mapper 75 with hardwired mirroring", rather than letting the "alternative nametables" or "Vs. System" flags override mapper 75's normal ASIC-controlled mirroring. Since Vs. System *by definition* has 4K of NTRAM, and no games have hardwired mirroring without 4K of NTRAM, this separation is artificial.

FCEUX-2.1.5's and Nestopia-1.4.0's implementations of mapper 151 do not attach the [5th bit of CHR bank](../VRC1/#Mirroring_Control.2C_CHR_bits_.28.249000-.249FFF.29) onto the lower 4 and instead emulate all registers as 8 bits wide because no Vs. System VRC1-using games have more than 64KiB of CHR.

Pictures of relevant Vs. System modules ([here](http://www.johnsarcade.com/nintendo_vs_ppu_info.php) or [here](http://playchoice.riemen.net/vs_list.html); scroll down to "Vs. Gradius" or "Vs. The Goonies") show the same VRC1 as in [NesCartDB](https://nescartdb.com/profile/view/3040/)
