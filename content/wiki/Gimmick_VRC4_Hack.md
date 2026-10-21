---
title: "Gimmick VRC4 Hack"
url: "/wiki/Gimmick_VRC4_Hack/"
description: "It is regular VRC4e with PRG-ROM mapped also at $6000-$7fff and one additional register at $6000-$7fff (bits of address are latched)"
wikiPage: "Gimmick_VRC4_Hack"
wikiCategories:
  - name: "Mappers"
    url: "/category/Mappers/"
---

[<img src="../../wiki-images/Gimmick_vrc4_components.jpg" width="300" height="191" alt="Gimmick vrc4 components.jpg" />](../../wiki-images/Gimmick_vrc4_components.jpg)

[](../../wiki-images/Gimmick_vrc4_components.jpg)

[](../../wiki-images/Gimmick_vrc4_components.jpg)

[](../../wiki-images/Gimmick_vrc4_components.jpg)

[](../../wiki-images/Gimmick_vrc4_components.jpg)[<img src="../../wiki-images/Gimmick_vrc4_top.png" width="300" height="178" alt="Gimmick vrc4 top.png" />](../../wiki-images/Gimmick_vrc4_top.png)

[](../../wiki-images/Gimmick_vrc4_top.png)

[](../../wiki-images/Gimmick_vrc4_top.png)

[](../../wiki-images/Gimmick_vrc4_top.png)

[](../../wiki-images/Gimmick_vrc4_top.png)[<img src="../../wiki-images/Gimmick_vrc4_bottom.png" width="300" height="171" alt="Gimmick vrc4 bottom.png" />](../../wiki-images/Gimmick_vrc4_bottom.png)

[](../../wiki-images/Gimmick_vrc4_bottom.png)

[](../../wiki-images/Gimmick_vrc4_bottom.png)

[](../../wiki-images/Gimmick_vrc4_bottom.png)

[](../../wiki-images/Gimmick_vrc4_bottom.png)[<img src="../../wiki-images/Gimmick_vrc4_sch.png" width="300" height="283" alt="Gimmick vrc4 sch.png" />](../../wiki-images/Gimmick_vrc4_sch.png)

[](../../wiki-images/Gimmick_vrc4_sch.png)

[](../../wiki-images/Gimmick_vrc4_sch.png)

[It is regular](../../wiki-images/Gimmick_vrc4_sch.png) [VRC4e](../VRC2_and_VRC4/) with PRG-ROM mapped also at \$6000-\$7fff and one additional register at \$6000-\$7fff (bits of address are latched)

    A~[011. .... .... PPPP] $6000-$7fff
                      ||||
                      ++++- 8kB PRG-ROM bank at $6000-$7fff (only banks 0-15 can be swapped)

This modification was needed because original Gimmick, which is based on mapper FME-7 also uses capability of mapping PRG at \$6000-\$7fff.

## Hardware implementation {#Hardware_implementation}

It consists of AX5208C (pirate VRC4), 4 bit latch (74\*161) and two multiplexers (74\*157). VRC's lines controlling WRAM at \$6000-\$6fff are tied to latch. Because VRC4 decodes RAM only at \$6000-\$6fff, when CPU accesses \$0000-\$7fff, one of muxes fools VRC by feeding GND at his A12, so that when CPU access \$6000-\$6fff or \$7000-\$7fff, VRC thinks for both it is \$6000-\$6fff and enables WRAM control line.

## Trivia {#Trivia}

For unknown reason, grounds of 74\*157 are not connected to the ground of rest cartridge.

Contrary to previous opinion, this *is* the true description of [iNES Mapper 183](../INES_Mapper_183/).
