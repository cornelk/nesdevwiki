---
title: "NES 2.0 Mapper 450"
url: "/wiki/NES_2.0_Mapper_450/"
description: "NES 2.0 Mapper 450 denotes the VRC2-based 晶太 (Jingtai) YY841157C multicart circuit board used for J.Y. Company's 1996 Super HiK 4-in-1 - 新系列瑪莉13代組合卡 multicart. The VRC2 clone's…"
wikiPage: "NES_2_0_Mapper_450"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 450** denotes the [VRC2](../VRC2_and_VRC4/)-based 晶太 (Jingtai) **YY841157C** multicart circuit board used for J.Y. Company's *1996 Super HiK 4-in-1 - 新系列瑪莉13代組合卡* multicart. The VRC2 clone's A0/A1 inputs are connected to CPU A0/A1; its µ-write interface is used to form an additional outer bank register at CPU \$6000-\$6FFF. It can be seen as the oversize VRC2 equivalent of [INES Mapper 047](../INES_Mapper_047/).

## <span id="Outer_Bank_Register_.28.246000-.246FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$6FFF, write) {#Outer_Bank_Register_.28.246000-.246FFF.2C_write.29}

    D~[.... ..BB]
              ++- PRG/CHR A18..A17
