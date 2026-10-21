---
title: "INES Mapper 189"
url: "/wiki/INES_Mapper_189/"
description: "iNES Mapper 189 is used for unlicensed MMC3 games that bank PRG-ROM in 32 KiB amounts, like AxROM, BNROM and GNROM. Known games:"
wikiPage: "INES_Mapper_189"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

iNES Mapper 189 is used for unlicensed [MMC3](../MMC3/) games that bank PRG-ROM in 32 KiB amounts, like [AxROM](../AxROM/), [BNROM](../INES_Mapper_034/) and [GNROM](../GxROM/). Known games:

- *Thunder Warrior - 轰天至尊* from TXC Corporation with its variant *Gluk the Thunder Warrior*
- *Street Fighter II: The World Warrior* from Yoko Soft and its many variants named *Super Fighter II´*, *Master Fighter II*, *Master Fighter III*, *Mario Fighter III'*

## 32 KiB PRG-ROM Bank Register {#32_KiB_PRG-ROM_Bank_Register}

### Original Board {#Original_Board}

    D~[..PP ....] A~[010. ...1 .... ....] (Address & $E100)==$4100
         ++- PRG A16..15

The original board variants switch four 32 KiB PRG-ROM banks based on bits 4 and 5 on writes to CPU \$4100-\$5FFF (address bit 9 must be set).

- *Thunder Warrior - 轰天至尊*: TXC **01-22017-000** 60-pin PCB
- *Gluk the Thunder Warrior*: TXC **01-22018-400** 72-pin PCB
- *Street Fighter II: The World Warrior*, *Master Fighter II: The World Warrior*, *Super Fighter II´*: **FC-001EM** (DIP)/**FC-001EMC** (COB) PCB

*Thunder Warrior* uses the 28-pin version of the [PT8154](../PT8154_pinout/) ASIC, which is a stripped-down MMC3 without any IRQ functionality; *Street Fighter II: The World Warrior* however needs a fully functional MMC3 clone, as the IRQ feature is utilized, and use a combination of a 40-pin PT8154 and another 40-pin PT8159.

### Oversize Board {#Oversize_Board}

    D~[.PPP ....] A~[01.. ...1 .... ....] (Address & $C100)==$4100
        +++- PRG A17..15

The oversize board (PCB name **自動車床控制板 MACP 811023**) was only used for the hack *Mario Fighter III*, which increases the PRG-ROM size to 256 KiB, and so needs an additional bank bit. Although the address mask was broadened, so that the register now responds to writes to CPU \$4100-\$7FFF (address bit 9 must be set), the actual game still writes to the original \$4132 address.

### Generic Pirate Boards {#Generic_Pirate_Boards}

**A-37**/**810809-C**

    D~[..PP ....] A~[011. .... .... ....] (Address & $E000)==$6000, MMC3's $A001 bit 7 set
         ++- PRG A16..15

**K-3086/810836-C**

    D~[.... ....] A~[011. .... .... ..PP] (Address & $E000)==$6000, MMC3's $A001 bit 7 set
                                      ++- PRG A16..15

There are countless pirated variants of the (already unauthorized) *Street Fighter II* 8-bit backport. Most of them connect the PRG-ROM Bank Register to the MMC3's PRG-RAM interface, which simplifies address decoding. However, doing so requires modifying the PRG-ROM content to change the write address, and to enable the PRG-RAM interface by writing \$80 to the MMC3's \$A001 register. Some of these respond to data bits 4 and bits 5, just like the original PCB; others respond to address bits 0 and 1, instead.

### Generalized definition for emulation {#Generalized_definition_for_emulation}

To run all board variants without having to distinguish them, the long-standing practice in emulators has been to respond to all writes to \$4020-\$7FFF, perform a bitwise OR of the upper and lower nibble of the data byte being written, and use the result as the 32 KiB PRG-ROM bank number. This works even for the K-3086 board variant since it writes the bank number both in A0/A1 and D0/D1. To discard the generic pirate boards and only run the original and oversized board variants, respond to writes to \$4100-\$5FFF with address bit 9 set, and only use the upper nibble of the data byte being written.

## See also {#See_also}

- [First thread with board pictures](https://forums.nesdev.org/viewtopic.php?f=28&t=17124)
- [Second thread with board pictures](https://forums.nesdev.org/viewtopic.php?f=9&t=20419)
