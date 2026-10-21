---
title: "INES Mapper 254"
url: "/wiki/INES_Mapper_254/"
description: "iNES Mapper 254 has been allocated for Pikachu Y2K of crypted ROMs which is one of the pirate hacks of Felix the Cat."
wikiPage: "INES_Mapper_254"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 254 has been allocated for *Pikachu Y2K of crypted ROMs* which is one of the pirate hacks of Felix the Cat.

This is basically [MMC3](../MMC3/) but WRAM read is XOR'ed in certain condition. (**TODO**: what is the source *hardware*?)

# Registers {#Registers}

These registers are overlaid on regular [MMC3](../MMC3/) registers and works in parallel.

## <span id=".248000-.249FFE.2C_even_.28write.29"></span>\$8000-\$9FFE, even (write) {#.248000-.249FFE.2C_even_.28write.29}

According to [MesenCE emulation source](https://github.com/nesdev-org/MesenCE/blob/master/Core/NES/Mappers/Mmc3Variants/MMC3_254.h), the PRG-RAM read value is XOR'ed with [\$A001](../MMC3/#PRG_RAM_protect_.28.24A001-.24BFFF.2C_odd.29) until the game writes to [\$8000](../MMC3/#Bank_select_.28.248000-.249FFE.2C_even.29) at least once.

It is unknown if the original hardware uses reset-detection for clearing the flag when reset. However there should be reset detection otherwise the game won't work.

## <span id=".24A001-.24BFFF.2C_odd_.28write.29"></span>\$A001-\$BFFF, odd (write) {#.24A001-.24BFFF.2C_odd_.28write.29}

This register sets the XOR mask value to apply to PRG-RAM read value, applied until the game writes to [\$8000](../MMC3/#Bank_select_.28.248000-.249FFE.2C_even.29) at least once.
