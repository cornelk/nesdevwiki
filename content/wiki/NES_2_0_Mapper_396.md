---
title: "NES 2.0 Mapper 396"
url: "/wiki/NES_2.0_Mapper_396/"
description: "NES 2.0 Mapper 396 denotes the 晶太 YY850437C and Realtec GN-51 multicart PCBs, basically an UNROM mapper with an outer bank register. Its UNIF MAPR is BMC-830752C."
wikiPage: "NES_2_0_Mapper_396"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 396** denotes the 晶太 **YY850437C** and Realtec *GN-51* multicart PCBs, basically an [UNROM](../UxROM/) mapper with an outer bank register. Its UNIF MAPR is **BMC-830752C**.

- *1995 Super 8-in-1* (JY-050 rev0)
- *Super 8-in-1 Gold Card Series* (JY-085)
- *Super 8-in-1 Gold Card Series* (JY-086)
- *2-in-1* (Realtec PG-07, GN-51 PCB)

Submappers:

- **Submapper 0**: Outer Bank Register at \$A000-\$BFFF, Nametable Arrangement via D5 or D6 (combined submapper 1 and 2)
- **Submapper 1**: Outer Bank Register at \$A000-\$BFFF, Nametable Arrangement via D6 (J.Y. **YY850437C** PCB variant)
- **Submapper 2**: Outer Bank Register at \$A000-\$BFFF, Nametable Arrangement via D5 (Realtec **GN-51** PCB variant)
- **Submapper 3**: Outer Bank Register at \$8000-\$BFFF, Nametable Arrangement via D5 (Realtec **8030** PCB variant)

## <span id="Outer_Bank_Register_.28.248000.2F.24A000-.24BFFF.29"></span>Outer Bank Register (\$8000/\$A000-\$BFFF) {#Outer_Bank_Register_.28.248000.2F.24A000-.24BFFF.29}

    D~[.Mm. .ppp]  (Address AND $E000) = $A000 (submappers 1,2)
        ||   |||   (Address AND $C000) = $8000 (submapper 3)
        ||   +++- PRG A19..A17
        |+------- Submapper 2/3 nametable arrangement, 1=Vertical
        +-------- Submapper 1 nametable arrangement, 1=Vertical

## <span id="Inner_Bank_Register_.28.248000-.249FFF.2C_.24C000-.24FFFF.29"></span>Inner Bank Register (\$8000-\$9FFF, \$C000-\$FFFF) {#Inner_Bank_Register_.28.248000-.249FFF.2C_.24C000-.24FFFF.29}

    D~[.... .ppp] (Address AND $E000) != $A000 (submappers 1,2)
             |||  (Address And $C000)  = $C000 (submapper 3)
             +++- PRG A16..A14 if CPU A14=0
