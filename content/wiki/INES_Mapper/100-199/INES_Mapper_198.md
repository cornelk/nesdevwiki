---
title: "INES Mapper 198"
url: "/wiki/INES_Mapper_198/"
description: "iNES Mapper 198 denotes the circuit board used for Xiangfeng Cartoon's Chinese translation of Capcom's 天地を喰らう II: 諸葛孔明伝 (Destiny of an Emperor II), entitled 吞食天地:…"
wikiPage: "INES_Mapper_198"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**iNES Mapper 198** denotes the circuit board used for [Xiangfeng Cartoon's](https://bootleggames.fandom.com/wiki/Xianfeng_Cartoon) Chinese translation of Capcom's 天地を喰らう II: 諸葛孔明伝 (Destiny of an Emperor II), entitled 吞食天地: 三国外传 (Tūnshí Tiāndì - Sānguó Wàizhuàn). In addition to normal 8 KiB of CHR-RAM and 8 KiB of battery-backed WRAM at \$6000-\$7FFF, it has an additional non-backed 4 KiB of WRAM at \$5000-\$5FFF. 640 KiB of PRG-ROM data is spread over two chips, 512 KiB providing the data for 8 KiB PRG-ROM banks \$00-\$3F, and 128 KiB providing the data for 8 KiB PRG-ROM banks \$40-\$4F, for a total of 640 KiB of PRG-ROM. Special consideration is needed for the two fixed MMC3 banks, which are located in the second chip, meaning that they are bank numbers \$4E and \$4F, respectively.

[INES Mapper 199](../INES_Mapper_199/) provides the same amount of RAM with more conventional (oversize) PRG-ROM banking.
