---
title: "NES 2.0 Mapper 312"
url: "/wiki/NES_2.0_Mapper_312/"
description: "NES 2.0 Mapper 312 is used for Kaiser's bootleg version of Highway Star. Its UNIF board name is UNL-KS7013B."
wikiPage: "NES_2_0_Mapper_312"
---

NES 2.0 Mapper 312 is used for Kaiser's bootleg version of *Highway Star*. Its UNIF board name is **UNL-KS7013B**.

# Banks {#Banks}

- CPU \$8000-\$BFFF: Switchable 16 KiB PRG-ROM bank
- CPU \$C000-\$FFFF: Fixed 16 KiB PRG-ROM bank \#7
- PPU \$0000-\$1FFF: Unbanked 8 KiB of CHR-RAM

# Registers {#Registers}

## <span id="PRG-ROM_Select_.28.246000-.247FFF.29"></span>PRG-ROM Select (\$6000-\$7FFF) {#PRG-ROM_Select_.28.246000-.247FFF.29}

    Mask: $E000

    D~7654 3210
      ---------
      .... .PPP
            +++- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF

## <span id="Nametable_Mirroring_Type_Select_.28.248000-.24FFFF.29"></span>Nametable Mirroring Type Select (\$8000-\$FFFF) {#Nametable_Mirroring_Type_Select_.28.248000-.24FFFF.29}

    Mask: $8000

    D~7654 3210
      ---------
      .... ...M
              +- Select nametable mirroring type
                  0: Vertical
                  1: Horizontal

## See also {#See_also}

- [PCB photos](https://forums.nesdev.org/viewtopic.php?p=256930#p256930)
- [Krzysiobal's reverse engineering from above photos](https://forums.nesdev.org/viewtopic.php?p=256961#p256961)
