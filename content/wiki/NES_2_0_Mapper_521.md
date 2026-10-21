---
title: "NES 2.0 Mapper 521"
url: "/wiki/NES_2.0_Mapper_521/"
description: "NES 2.0 Mapper 521 is used for 장두진 바둑교실: 입문편, commonly known as Korean Igo. Its UNIF board name is DREAMTECH01, without prefix."
wikiPage: "NES_2_0_Mapper_521"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**NES 2.0 Mapper 521** is used for *장두진 바둑교실: 입문편*, commonly known as *Korean Igo*. Its UNIF board name is **DREAMTECH01**, without prefix.

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB switchable PRG-ROM bank
- CPU \$C000-\$FFFF: 16 KiB fixed PRG-ROM bank (last bank in ROM file)
- PPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM.

Mirroring is hard-wired.

Since the fixed PRG-ROM bank comes from a separate PRG-ROM chip, the ROM file can have a non-power-of-two size.

# <span id="Bank_Register_.28.245020.2C_write.29"></span>Bank Register (\$5020, write) {#Bank_Register_.28.245020.2C_write.29}

    Mask: $F030
    D~[.... BBBB]
            ++++- PRG A18..14 when CPU A14=0

Note that writes to \$5030 are ignored.

# See also {#See_also}

[PCB images and circuit diagram](https://forums.nesdev.org/viewtopic.php?p=296603)
