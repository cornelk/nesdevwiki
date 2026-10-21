---
title: "Venus Mini Doctor"
url: "/wiki/Venus_Mini_Doctor/"
description: "The Venus Mini Doctor is a plug-through cartridge to play NROM and CNROM games loaded from 2.8\" FDS diskettes without modification to the games' code, priced below a simple Game Doctor 1M. It is a…"
wikiPage: "Venus_Mini_Doctor"
---

The **Venus Mini Doctor** is a plug-through cartridge to play NROM and CNROM games loaded from 2.8" [FDS](../Family_Computer_Disk_System/) diskettes without modification to the games' code, priced below a simple [Game Doctor 1M](../RAM_cartridge/#First_generation). It is a plug-through device that is inserted into the cartridge port of a Nintendo Famicom, and takes the FDS' RAM Adapter into its own cartridge port. The games come on two types of FDS disks, those specific to the Venus Mini Doctor, and regular Game Doctor disks containing NROM or CNROM games.

# Hardware {#Hardware}

The Venus Mini Doctor consists of a PAL (Models V40 and V64) and optional 32 KiB of CHR-RAM (Model V64 only). The PAL operates in two modes:

In **Load** mode, CPU address space is unaltered; all CPU blocks are loaded to DRAM at \$6000-\$DFFF, just as with normal FDS games, and CHR blocks are either loaded to the FDS RAM Adapter's write-enabled CHR-RAM (CHR bank 0) or to the Mini Doctor V64's 32 KiB of CHR-RAM (CHR banks 1-3). The CHR bank is selected by D0-D1 of the Mini Doctor's register at \$43xx. For some odd reason, the CHR bank automatically resets to 0 after a few M2 cycles in Load mode, so \$43xx must be written-to before every write to CHR-RAM banks 1-3.

The Mini Doctor is switched into **Play** mode by a write to \$41xx. In Play mode, the PAL modifies the address from the CPU such that the \$E000-\$FFFF bank is mapped to \$6000-\$7FFF. CHR-RAM is write-protected, a CNROM latch that responds to writes at \$8000-\$FFFF selects either the FDS RAM Adapter's CHR-RAM (latch=0) or three banks of the Mini Doctor V64's 32 KiB of CHR-RAM, if present (latch=1-3). On the model V40, selecting a CHR-RAM bank above 0 results in CHR pattern data coming from open bus, as the 32 KiB CHR-RAM chip is not installed.

# Disk Layouts {#Disk_Layouts}

The simplest native Mini Doctor disks containing NROM games only contain blocks to populate 32 KiB of PRG-RAM, 8 KiB of CHR-RAM, a block that writes a single byte to the [FDS' Nametable Arrangement register](../Family_Computer_Disk_System/#FDS_Control_.28.244025.29), and a final block that writes to address \$41xx to put the Mini Doctor into Play mode. No license screen data or bypass setup is necessary, as the \$41xx write automatically removes the FDS BIOS that would perform the [approval check](../FDS_BIOS/#Approval_check) from CPU address space. A soft reset is then required for the CPU to follow the game's reset vector, now mapped to CPU \$FFFC.

Most native Mini Doctor disks contain their own loader that takes over reading the disk from the BIOS, reading PRG-RAM data as well as up to 32 KiB of CHR-RAM data, performing the necessary \$43xx CHR bank switches as necessary, switching the Mini Doctor into Play mode, and finally jumping to the game's reset handler without needing the user to soft-reset the console.

A utility named **MiniLoad** (Loader for Venus Mini Doctor) allows playing NROM and CNROM games from disks in the standard [Game Doctor/Magic Card FDS Format](../Game_Doctor/Magic_Card_FDS_Format/). Trying to use the utility with a native Mini Doctor disk in the aforementioned format will produce an "incompatible disk" message, as the utility is only for regular Game Doctor/Magic Card disks.
