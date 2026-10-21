---
title: "NES 2.0 Mapper 458"
url: "/wiki/NES_2.0_Mapper_458/"
description: "NES 2.0 Mapper 458 denotes the MMC3-based K-3102 (submapper 0) and GN-23 (submapper 1) multicart PCBs. The MMC3 is only used for CHR banking and mirroring, while all PRG banking is controlled by an…"
wikiPage: "NES_2_0_Mapper_458"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 458** denotes the [MMC3](../MMC3/)-based **K-3102** (submapper 0) and **GN-23** (submapper 1) multicart PCBs. The MMC3 is only used for CHR banking and mirroring, while all PRG banking is controlled by an extra register, similar to [NES 2.0 Mapper 259](../NES_2.0_Mapper_259/).

## <span id="Extra_Register_.28.246000-.247FFF.2C_write.29"></span>Extra Register (\$6000-\$7FFF, write) {#Extra_Register_.28.246000-.247FFF.2C_write.29}

    A~[011. .... d.DM PPPP]
                 | || ++++- PRG A17..A14
                 | || +---- 0: NROM-128, 1: NROM-256 (submapper 1)
                 | |+------ 0: NROM-128, 1: NROM-256 (submapper 0)
                 | +------- 0: Normal, 1: Replace CPU A4..0 with solder pad value (submapper 0)
                 +--------- 0: Normal, 1: Replace CPU A4..0 with solder pad value (submapper 1)
    Value on reset: $00

## Notes {#Notes}

- WRAM must be enabled in \$A001.7 before writing to the extra register.
