---
title: "NES 2.0 Mapper 457"
url: "/wiki/NES_2.0_Mapper_457/"
description: "NES 2.0 Mapper 457 denotes the MMC3-based 810431C multicart PCB, used for the 92 年超高K 黄金卡 35-in-1 刺激大驚險旅風侠 multicart. It is basically an oversized version of INES Mapper…"
wikiPage: "NES_2_0_Mapper_457"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 457** denotes the [MMC3](../MMC3/)-based **810431C** multicart PCB, used for the *92 年超高K 黄金卡 35-in-1 刺激大驚險旅風侠* multicart. It is basically an oversized version of [INES Mapper 047](../INES_Mapper_047/) with a selectable outer PRG and CHR bank size. Alternatively, it can be thought of as a simplified version of [INES Mapper 052](../INES_Mapper_052/) without independent PRG/CHR control.

## <span id="Extra_Register_.28.246000-.247FFF.2C_write.29"></span>Extra Register (\$6000-\$7FFF, write) {#Extra_Register_.28.246000-.247FFF.2C_write.29}

    Mask: $E000

    D~[.... SBBB]
            |+++- PRG/CHR A19..A17
            +---- 0: 128 KiB outer PRG and CHR bank size
                  1: 256 KiB outer PRG and CHR bank size
    Value on reset: $00

## Notes {#Notes}

- WRAM must be enabled in \$A001.7 before writing to the extra register.
- Nestopia Plus! has assigned mapper number **2786** for this circuit board.
