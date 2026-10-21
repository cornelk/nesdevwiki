---
title: "INES Mapper 120"
url: "/wiki/INES_Mapper_120/"
description: "iNES Mapper 120 denotes Whirlwind Manu's LH15 cartridge conversion of Square's FDS game とびだせ大作戦 (Tobidase Daisakusen)."
wikiPage: "INES_Mapper_120"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**iNES Mapper 120** denotes Whirlwind Manu's **LH15** cartridge conversion of Square's [FDS](../Family_Computer_Disk_System/) game とびだせ大作戦 (Tobidase Daisakusen).

# Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG ROM bank
- CPU \$8000-\$FFFF: 32 KiB hard-wired PRG ROM bank \#2
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR RAM
- Nametable mirroring: hard-wired (to vertical)

# Registers {#Registers}

## <span id="PRG_ROM_switch_.28.2441FF.2C_write.29"></span>PRG ROM switch (\$41FF, write) {#PRG_ROM_switch_.28.2441FF.2C_write.29}

    Address Mask: $E100

    D~[.... .PPP]
             +++- Select 8 KiB PRG ROM bank at CPU $6000-$7FFF
