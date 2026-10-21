---
title: "RAM cartridge"
url: "/wiki/RAM_cartridge/"
description: "Although the NES/Famicom was originally designed to only accept cartridges containing game code and data in read-only memory chips (ROM), both Nintendo and unlicensed third parties released…"
wikiPage: "RAM_cartridge"
---

Although the NES/Famicom was originally designed to only accept cartridges containing game code and data in read-only memory chips ([ROM](../ROM/)), both Nintendo and unlicensed third parties released cartridges that allowed games to be played from RAM chips, with data loaded from floppy disks or solid-state flash cards controlled by a small ROM-based BIOS.

# Licensed RAM cartridges {#Licensed_RAM_cartridges}

The [Famicom Disk System](../Family_Computer_Disk_System/)'s RAM Adapter provides 32 KiB of PRG-RAM and 8 KiB of CHR-RAM. Data is loaded from Mitsumi QuickDisk-format 2.8 inch floppy disks. Games have to be specifically written for the FDS RAM Adapter's memory map, for saving game progress to floppy disk, and if more than 32 KiB of PRG or 8 KiB of CHR data are to be used, for loading additional data from disk.

# Unlicensed RAM cartridges {#Unlicensed_RAM_cartridges}

All unlicensed RAM cartridges are primarily designed to run games that originally were released on ROM cartridges. As a consequence, they closely mimic the memory map of ROM cartridges, load the entire game code and data at boot and do not load any data afterwards.

## First generation {#First_generation}

Early unlicensed RAM cartridges had 128 or 256 KiB of PRG-RAM and 32 KiB of CHR-RAM. They could directly mimic discrete mappers such as [UNROM](../UxROM/), [CNROM](../CNROM/) and [GNROM](../GxROM/). Games that originally ran on ASIC mappers such as the Nintendo [MMC1](../MMC1/) or [MMC3](../MMC3/) had to be modified to run on these RAM cartridges. If the games used more than 32 KiB of CHR-ROM, or banked it in amounts smaller than 8 KiB, then further modification was necessary, using a technique of *caching* the most recently-used 32 KiB of CHR-ROM.

These devices have been called "copiers" in the early NES emulation scene, even though none of them actually have the capability to copy a ROM cartridge game to floppy disk.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>. Instead, users had to purchase (or copy) suitable disks from the manufacturer of the RAM cartridge or some other company close to it. Disks were usually specific to a particular RAM cartridge manufacturer and model, with later devices being backwards-compatible only to models by the same manufacturer and the original *Bung Game Doctor*, often containing detection code to specifically lock-out competing RAM cartridges.

All of these devices are connected *between* the Famicom main unit and the FDS RAM adapter. The loading process is based on the process of loading a normal FDS game: the original FDS BIOS starts loading a game once a disk has been inserted. [Disks targeting these devices](../Game_Doctor/Magic_Card_FDS_Format/) will contain a file that is loaded by the FDS BIOS to a CPU address unused by the FDS RAM adapter, but that when written to switches in the RAM cartridges's own BIOS and memory map.

- *Bung Game Doctor* (in 128 KiB and 256 KiB PRG-RAM versions)
- *Venus Game Converter*
- *Front Fareast Magicard* (adds 8 KiB PRG-"ROM" banking mode)
- *Bung Super Game Doctor* (in 256 KiB and 512 KiB versions; adds 8 KiB PRG-"ROM" banking mode and cycle-based IRQ counter)
- *Venus Turbo Game Doctor 4M* (512 KiB PRG-RAM, adds 8 KiB PRG-ROM banking mode, cycle-based IRQ counter, and real-time saving capability)

## Second generation {#Second_generation}

Later unlicensed RAM cartridges increase the PRG-RAM size to 512 KiB, and the CHR-RAM size to 256 KiB bankable in amounts smaller than 8 KiB. Games still had to be modified to run on these devices.

- *Venus Turbo Game Doctor 6M/6+*
- *Front Fareast [Super Magic Card](../Super_Magic_Card/)*. Adds PA12-based IRQ counter. Can also load games from a 3.5 inch MS-DOS-format floppy disk drive (720 KiB or 1.44 MiB), or from a parallel interface connected to a PC with CD-ROM drive.

## Third generation {#Third_generation}

The *Bung Game Master* has a Xilinx XC3042 FPGA to mimic most ASIC-based mappers, no longer requiring modifications to the games themselves. The mapper-specific fusemap is loaded from floppy disk (called "preboot disk"), just like the games themselves. It can load games from 2.8 inch QuickDisks like earlier models, or from a 3.5 inch MS-DOS-format floppy disk drive.

## Fourth generation {#Fourth_generation}

Modern RAM cartridges such as the [PowerPak](../PowerPak/) or [Everdrive N8](../Everdrive_N8/) are similar to the *Bung Game Master* in using an FPGA to mimic almost any mapper hardware, but load games and fusemaps from solid state flash cards exclusively. For this reason, they are called "flashcarts".

# See also {#See_also}

- *[Venus Mini Doctor](../Venus_Mini_Doctor/)* - a simple plug-through cartridge for loading [NROM](../NROM/)/CNROM games.
- *I2 [Souseiki Fammy](../Souseiki_Fammy/)* - capable of dumping NROM cartridges to a disk format playable on the device.
- [Info on various Famicom "copiers"](http://www.famicomdisksystem.com/game-doctor-copiers/)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) The contemporary term in their East Asian target markets was "RAM disk", while actual copiers were called "wild card".
