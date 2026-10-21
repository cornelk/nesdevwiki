---
title: "Souseiki Fammy"
url: "/wiki/Souseiki_Fammy/"
description: "The Souseiki Fammy (創世機ファミー) is an unlicensed video game add-on made by I2 (アイ・ツー) for the Famicom. It primarily acts as a middleman between the Famicom Disk System and the…"
wikiPage: "Souseiki_Fammy"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "Hardware"
    url: "/category/Hardware/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
---

The **Souseiki Fammy** (創世機ファミー) is an unlicensed video game add-on made by I2 (アイ・ツー) for the Famicom. It primarily acts as a middleman between the [Famicom Disk System](../Family_Computer_Disk_System/) and the Famicom to facilitate transfers between cartridges and disks.

The following disk software was released for the Souseiki Fammy:

- *ROM-QD for 256K+64K* - dumps [NROM](../NROM/) cartridges to a disk format playable on the Fammy.
- *B.B RAM Plus* - dumps/writes battery-backed RAM contents for cartridges without PRG-RAM protections.
- *Disk BASIC Generator Kit* - converts a *Family BASIC v2* cartridge to a standalone disk, with support for disk saving.
- *Slow Plus for ROM* - applies slow-down to cartridge games.
- *Slow Plus for QD* - applies slow-down to disk games.
- *Slow Plus for ROM-QD* - applies slow-down to disks created using ROM-QD.

## Hardware {#Hardware}

The Souseiki Fammy is a cartridge middleman device with slots for the FDS RAM adapter and a standard game cartridge. A Disk/ROM switch on the front panel controls whether the console will boot from/can access the FDS or the game cartridge. Discrete logic is used to control the memory mapping between the two slots. A reset button, 4-input DIP switch, and LED are also located on the front panel.

An expansion port is located at the rear of the Fammy but no software or hardware is known to use it. TODO: Expansion port pinout?

The Souseiki Fammy contains 8KiB of on-board PRG-RAM. A 4KiB half of it is always mapped to \$5000-\$5FFF - its full range is only accessible when mirrored to \$E000-\$FFFF, see the register descriptions below.

## Banks {#Banks}

- PPU \$0000-\$1FFF: 8KiB FDS CHR-RAM or cartridge CHR-ROM/RAM
- CPU \$5000-\$5FFF: 4KiB upper half of Fammy PRG-RAM
- CPU \$6000-\$7FFF: 8KiB FDS PRG-RAM or cartridge PRG-RAM
- CPU \$8000-\$DFFF: 24KiB FDS PRG-RAM or cartridge PRG-ROM
- CPU \$E000-\$EFFF: 4KiB FDS BIOS PRG-ROM, cartridge PRG-ROM, or upper half of Fammy PRG-RAM mirror
- CPU \$F000-\$FFFF: 4KiB FDS BIOS PRG-ROM, cartridge PRG-ROM, or lower half of Fammy PRG-RAM mirror

## Registers {#Registers}

All registers are write-only.

These descriptions assume that the 4 DIP switches are all set to the ON position, as no software is currently known to rely on alternate DIP switch settings.

TODO: Describe DIP switch settings for completeness.

### <span id="Memory_mapping_control_.28.244C00-.244DFF.29"></span>Memory mapping control (\$4C00-\$4DFF) {#Memory_mapping_control_.28.244C00-.244DFF.29}

    7  bit  0
    ---------
    LRDW FEPM
    |||| ||||
    |||| |||+- Disk/ROM mode (0 = use current switch setting, 1 = invert switch setting)
    |||| ||+-- PPU CIRAM /CE, A10, A13 mode (0 = use current Disk/ROM mode, 1 = invert Disk/ROM mode)
    |||| |+--- Fammy PRG-RAM mirroring (0 = $5000-$5FFF only, 1 = map a mirror at $E000-$FFFF)
    |||| +---- Fammy PRG-RAM $E000 mirror mapping (0 = $E000-$FFFF, 1 = $F000-$FFFF only)
    |||+------ Disk CPU/PPU write control (0 = enable, 1 = disable writes to CPU $8000-$FFFF and PPU $0000-$1FFF for Disk slot)
    ||+------- Disk M2 control (0 = use current Disk/ROM mode, 1 = invert Disk/ROM mode)
    |+-------- ROM M2 control (0 = use current Disk/ROM mode, 1 = invert Disk/ROM mode)
    +--------- On-board LED (0 = off, 1 = on)

Notes:

- This register's contents are cleared with all 0s at power-on/reset. The Fammy's reset button also clears this register.
- The Disk/ROM mode (bit 0) controls whether or not the physical switch setting is inverted before being used by other controls.
- The PPU CIRAM /CE, A10, A13 mode flips which slot uses the mappings for [PPU pattern tables](../PPU_pattern_tables/), [nametables](../PPU_nametables/), and its [arrangement](../Mirroring/#Nametable_Mirroring).
- The 4KiB upper half of the Fammy PRG-RAM is always mapped to \$5000-\$5FFF. Its full range is only accessible when mirrored to \$E000-\$FFFF.
  - \$6000-\$7FFF is mapped to either the FDS PRG-RAM or the cartridge's PRG-RAM (if present and not protected). The latter mapping is required by *B.B RAM Plus*.
- Disabling Disk CPU/PPU writes causes [bus conflict](../Bus_conflict/) behaviour, since writes are forced into reads. This is for compatibility with *ROM-QD* disks of games which expect PRG-ROM/CHR-ROM writes to have no effect.
  - \$5000-\$7FFF is unaffected, therefore it will most likely break games which rely on [open bus](../Open_bus_behavior/) for this area.
- The M2 controls flip whether or not the M2 signal is sent to the corresponding slot (forced low when disabled). This effectively controls CPU access for mapper registers and/or the \$6000-\$7FFF region of a given slot, independent of /ROMSEL (M2 NAND CPU A15).
  - *ROM-QD* boots with the switch in Disk mode, then writes \$E1 (%11100001) before dumping a portion of cartridge ROM data to \$6000-\$67FF. Ignoring bit 7 (LED on), bit 0 inverts the switch setting to become ROM mode (so ROM gets /ROMSEL and PPU signals), bit 5 enables M2 on the Disk slot in order to access the FDS registers plus \$6000-\$7FFF, and bit 6 disables M2 for the ROM slot in order to avoid bus conflicts with the Disk slot.
- The LED is primarily used by supported software to indicate transfers between the FDS and game cartridge. The *Slow Plus* disks additionally flash the LED to indicate the "slow level" (i.e. game speed).

Some signals are not controlled by the Fammy and thus will cause conflicts between slots (irrelevant for the intended RAM adapter + NROM combination):

- PPU /A13, for cartridges which map their own nametable memory.
- CPU /IRQ, for cartridges other than the RAM adapter which generate [IRQs](../IRQ/).
- Audio from the 2A03 [APU](../APU/) is sent to both slots, and Audio to RF from both slots are mixed before being returned to the console. Therefore, the [expansion audio](../List_of_games_with_expansion_audio/) level is halved relative to the APU.

### <span id="Unused_.28.244800-.244BFF.29"></span>Unused (\$4800-\$4BFF) {#Unused_.28.244800-.244BFF.29}

There is logic present to route \$4800-\$4BFF writes to the expansion port located at the rear of the Fammy. It may have been reserved for registers in unreleased expansion hardware.

## See Also {#See_Also}

- Japanese video demonstrations by Ootakeke (おおたけけチャンネル):
  - [Fammy overview & *ROM-QD for 256K+64K*](https://www.youtube.com/watch?v=WxhiE1M3baI) (*Donkey Kong*)
  - [*ROM-QD for 256K+64K*](https://www.youtube.com/watch?v=z66Kbno-0y8) (*F1 Race*)
  - [*Disk BASIC Generator Kit*](https://www.youtube.com/watch?v=8XWxWEZ2NaY) (creation)
  - [*Disk BASIC Generator Kit*](https://www.youtube.com/watch?v=drIvX5ijUtQ) (creation & disk saving)

## References {#References}

- [Forum discussion](https://forums.nesdev.org/viewtopic.php?t=26064)
- [GitHub repository](https://github.com/edorfaus/souseiki-fammy-kicad): KiCad schematic & analysis
