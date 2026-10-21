---
title: "INES Mapper 111"
url: "/wiki/INES_Mapper_111/"
description: "iNES Mapper 111 denotes the unrelated GTROM circuit board and an MMC1 variant used in a Chinese homebrew translation of Ninja Ryuukenden."
wikiPage: "INES_Mapper_111"
wikiCategories:
  - name: "GNROM-like mappers"
    url: "/category/GNROM-like_mappers/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "Mappers with flash save"
    url: "/category/Mappers_with_flash_save/"
---

**iNES Mapper 111** denotes the unrelated [GTROM](../GTROM/) circuit board and an [MMC1](../MMC1/) variant used in a Chinese homebrew translation of *Ninja Ryuukenden*.

## GTROM {#GTROM}

See: [GTROM](../GTROM/)

## Non-serialized MMC1 variant {#Non-serialized_MMC1_variant}

Prior to the introduction of GTROM, Mapper 111 was assigned to a Chinese homebrew translation of *Ninja Ryukenden* (Japanese version of *Ninja Gaiden*). This translation uses a non-serialized version of [MMC1](../MMC1/) and supports 256KiB of CHR-ROM, whereas the official MMC1 is limited to 128KiB. This assignment can coexist with GTROM provided the emulator recognizes that the translation uses CHR-ROM and therefore must emulate the MMC1 variant instead of GTROM hardware.

A possibly later version of the translation uses [MMC3](../MMC3/) instead of this MMC1 variant.
