---
title: "NES 2.0 Mapper 379"
url: "/wiki/NES_2.0_Mapper_379/"
description: "NES 2.0 Mapper 379 denotes the 35BH-1 circuit board, used for a simple 35-in-1 NROM multicart. It's basically identical to INES Mapper 038 except that the bank register is in the $8000-$FFFF rather…"
wikiPage: "NES_2_0_Mapper_379"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 379** denotes the **35BH-1** circuit board, used for a simple *35-in-1* NROM multicart. It's basically identical to [INES Mapper 038](../INES_Mapper_038/) except that the bank register is in the \$8000-\$FFFF rather than the \$7000-\$7FFF range.

## <span id="Outer_Bank_register_.28.248000-.24FFFF.2C_write.29"></span>Outer Bank register (\$8000-\$FFFF, write) {#Outer_Bank_register_.28.248000-.24FFFF.2C_write.29}

    Mask: $8000

    D~7654 3210
      ---------
      .... CCPP
           ||++- Select 32 KiB PRG-ROM bank
           ++--- Select 8 KiB CHR-ROM bank
