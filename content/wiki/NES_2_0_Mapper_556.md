---
title: "NES 2.0 Mapper 556"
url: "/wiki/NES_2.0_Mapper_556/"
description: "NES 2.0 Mapper 556 is used for the for the 超强小新2+瑪莉家族 7-in-1 (JY-215) multicart and an unknown 50-in-1 MMC3+VRC4+NROM multicart, which has VRC4 compatibility to support Crayon…"
wikiPage: "NES_2_0_Mapper_556"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multi-ASIC mappers"
    url: "/category/Multi-ASIC_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 556 is used for the for the 超强小新2+瑪莉家族 7-in-1 (JY-215) multicart and an unknown 50-in-1 MMC3+VRC4+NROM multicart, which has VRC4 compatibility to support Crayon Shin-Chan (hack of Bio Miracle Bokutte Upa) and Goal 2 (VRC4 hack of JPN version).

The multicart's game select register functions similar to [iNES Mapper 045](../INES_Mapper_045/), but is moved to \$5000 and repurposed a bit in the second write to select MMC3/VRC4 mode. The VRC4 mode seems to behave like the VRC4e variant (A2/A3 connected to VRC4 A0/A1).

## <span id="CHR-OR_LSB_.28.245000_.230.2C_write.29"></span>CHR-OR LSB (\$5000 \#0, write) {#CHR-OR_LSB_.28.245000_.230.2C_write.29}

    Mask: Unknown

    D~7654 3210
      ---------
      CCCC CCCC
      ++++-++++- Select CHR A10-A17

## <span id="PRG-OR_LSB_.28.245000_.231.2C_write.29"></span>PRG-OR LSB (\$5000 \#1, write) {#PRG-OR_LSB_.28.245000_.231.2C_write.29}

    Mask: Unknown

    D~7654 3210
      ---------
      pPPP PPPP
      |+++-++++- Select PRG A13-A19
      +--------- Select PRG A20, CHR A20

## <span id="CHR-AND.2C_CHR-OR.2FPRG-OR_MSB_.28.245000_.232.2C_write.29"></span>CHR-AND, CHR-OR/PRG-OR MSB (\$5000 \#2, write) {#CHR-AND.2C_CHR-OR.2FPRG-OR_MSB_.28.245000_.232.2C_write.29}

    Mask: Unknown

    D~7654 3210
      ---------
      MPCC cccc
      |||| ++++- Select number of CHR bits from MMC3
      ||||       ($F: 256 KiB, $E: 128 KiB ... $7-$0: 1 KiB)
      ||++------ Select CHR A18-A19
      |+-------- Select PRG A21, CHR A21
      +--------- ASIC compatibility mode
                 (0=MMC3, 1=VRC4)

## <span id="PRG-AND.2C_register_lock_.28.245000_.233.2C_write.29"></span>PRG-AND, register lock (\$5000 \#3, write) {#PRG-AND.2C_register_lock_.28.245000_.233.2C_write.29}

    Mask: Unknown

    D~7654 3210
      ---------
      L-PP PPPP
      | ++ ++++- Select PRG-AND mask (inverted)
      |          ($00: 512 KiB, $20: 256 KiB...)
      +--------- 1=Lock outer bank registers

## See Also {#See_Also}

- [Pictures and analysis for the 50-in-1 multicart](https://forums.nesdev.org/viewtopic.php?p=301711#p301711)
