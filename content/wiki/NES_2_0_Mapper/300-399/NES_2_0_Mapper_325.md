---
title: "NES 2.0 Mapper 325"
url: "/wiki/NES_2.0_Mapper_325/"
description: "NES 2.0 Mapper 325 is used for Super Mali Splash Bomb, a bootleg hack of Keroppi to Keroriinu no Splash Bomb. Its UNIF board name is UNL-MALISB. It's an MMC3 clone with the address input and output…"
wikiPage: "NES_2_0_Mapper_325"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

NES 2.0 Mapper 325 is used for *Super Mali Splash Bomb*, a bootleg hack of *Keroppi to Keroriinu no Splash Bomb*. Its UNIF board name is **UNL-MALISB**. It's an MMC3 clone with the address input and output lines connected in an unusual manner: The MMC3 clones' ...

- CPU A0 input is connected to CPU A3 and additionally to CPU A2 if CPU A14==1;
- PRG A15 (PRG bank number bit 2) is swapped with PRG A16 (PRG bank number bit 3);
- CHR A11 (CHR bank number bit 1) is swapped with CHR A15 (CHR bank number bit 5).
