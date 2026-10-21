---
title: "Mappers with battery-backed CHR-RAM"
url: "/category/Mappers_with_battery-backed_CHR-RAM/"
description: "Very few boards use a battery to retain the content of CHR-RAM. The purpose of doing so is presumably that if a game already uses CHR-RAM for other reasons and only has a few bytes of save data to…"
wikiPage: "Category_Mappers_with_battery_backed_CHR_RAM"
type: category
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

Very few boards use a battery to retain the content of CHR-RAM. The purpose of doing so is presumably that if a game already uses CHR-RAM [for other reasons](../../wiki/CHR_ROM_vs._CHR_RAM/) and only has a few bytes of save data to retain, storing these bytes in the same CHR-RAM chip (and sacrificing available tile numbers in the process) and connecting a battery to it is more economical than adding a separate WRAM chip.

[Nintendo's early MMC1 data sheet](http://forums.nesdev.org/viewtopic.php?f=2&t=16925) mentions an "SLROM" board with 8 KiB of battery-backed CHR-RAM. The actual [SLROM](../../wiki/MMC1/) board that went into production uses CHR-ROM and no PRG-RAM or CHR-RAM at all.
