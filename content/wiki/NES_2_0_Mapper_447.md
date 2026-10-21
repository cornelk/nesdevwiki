---
title: "NES 2.0 Mapper 447"
url: "/wiki/NES_2.0_Mapper_447/"
description: "NES 2.0 Mapper 447 denotes the VRC4-based KL-06 multicart circuit board used by the 1993 New 860-in-1 Over-Valued Golden Version Games multicart. The VRC4 clone's A0/A1 inputs are connected to CPU…"
wikiPage: "NES_2_0_Mapper_447"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 447** denotes the [VRC4](../VRC2_and_VRC4/)-based **KL-06** multicart circuit board used by the *1993 New 860-in-1 Over-Valued Golden Version Games* multicart. The VRC4 clone's A0/A1 inputs are connected to CPU A2/A3, and its PRG-RAM chip-enable signal drives an outer bank register.

## <span id="Outer_Bank_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank Register (\$6000-\$7FFF, write) {#Outer_Bank_Register_.28.246000-.247FFF.2C_write.29}

    A~[0110 .... .... .NBB]
                       |++- PRG/CHR A18..A17
                       ||+- 1=Lock outer bank register
                       |+-- 1=NROM-128 mode if N=1, 0=NROM-256 mode if N=1
                       +--- 1=NROM mode

NROM mode causes the VRC4's CPU A14 line to be held low, making the VRC4's two PRG registers apply to both \$8000/\$A000 and \$C000/\$E000. NROM-256 mode additionally replaces PRG A14 with CPU A14. The 'Lock' functionality is needed because the game that is active when PRG A17=1 (*Crisis Force*) uses 2 KiB of PRG-RAM that overlap the outer bank register.
