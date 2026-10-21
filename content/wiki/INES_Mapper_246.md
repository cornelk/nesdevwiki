---
title: "INES Mapper 246"
url: "/wiki/INES_Mapper_246/"
description: "iNES Mapper 246 denotes the G0151-1 circuit board, used only for the Taiwanese game 封神榜 (Fēngshénbǎng: Fúmó Sān Tàizǐ)."
wikiPage: "INES_Mapper_246"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

**iNES Mapper 246** denotes the **G0151-1** circuit board, used only for the Taiwanese game [封神榜 (*Fēngshénbǎng: Fúmó Sān Tàizǐ*)](http://community.fandom.com/wiki/c:bootleggames:Feng_Shen_Bang).

# Banks {#Banks}

- CPU \$6800-\$6FFF: 2 KiB PRG-RAM, battery-banked
- CPU \$8000-\$9FFF: 8 KiB PRG-ROM bank, switchable via \$6000
- CPU \$A000-\$BFFF: 8 KiB PRG-ROM bank, switchable via \$6001
- CPU \$C000-\$DFFF: 8 KiB PRG-ROM bank, switchable via \$6002
- CPU \$E000-\$FFFF: 8 KiB PRG-ROM bank, switchable via \$6003
- PPU \$0000-\$07FF: 2 KiB CHR-ROM bank, switchable via \$6004
- PPU \$0800-\$0FFF: 2 KiB CHR-ROM bank, switchable via \$6005
- PPU \$1000-\$17FF: 2 KiB CHR-ROM bank, switchable via \$6006
- PPU \$1800-\$1FFF: 2 KiB CHR-ROM bank, switchable via \$6007

Nametable mirroring is hard-wired and set by the iNES header.

# Registers {#Registers}

## <span id="PRG-ROM_bank_select_.28.246000-.246003.29.2C_write"></span>PRG-ROM bank select (\$6000-\$6003), write {#PRG-ROM_bank_select_.28.246000-.246003.29.2C_write}

    A~[0110 0000 000. .0rr] D~[..bb bbbb]
                        ||       ++-++++- PRG A18..A13 of bank to select
                        ++--------------- CPU A14..A13 bank to switch

- When reading from CPU address \$FFE4-\$FFE7, \$FFEC-\$FFEF, \$FFF4-\$FFF7, or \$FFFC-\$FFFF, PRG A17 is forced high, as if register \$6003 were OR'd with \$10.
- Register \$6003 is set to \$FF at power-on only.

## <span id="CHR-ROM_bank_select_.28.246004-.246007.29.2C_write"></span>CHR-ROM bank select (\$6004-\$6007), write {#CHR-ROM_bank_select_.28.246004-.246007.29.2C_write}

    A~[0110 0000 000. .1rr] D~[bbbb bbbb]
                        ||     ++++-++++- CHR A18..A11 of bank to select
                        ++--------------- PPU A12..A11 bank to switch

# See also {#See_also}

- <http://forums.nesdev.org/viewtopic.php?t=13969> Why does Fong Shen Bang only work in Famiclones?
