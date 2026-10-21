---
title: "NES 2.0 Mapper 304"
url: "/wiki/NES_2.0_Mapper_304/"
description: "NES 2.0 Mapper 304 is used for several ROM cartridge conversions of Famicom Disk Systems for which, along with the entirely-different INES Mapper 043 and NES 2.0 Mapper 311, UNIF MAPR UNL-SMB2J is…"
wikiPage: "NES_2_0_Mapper_304"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

NES 2.0 Mapper 304 is used for several ROM cartridge conversions of Famicom Disk Systems for which, along with the entirely-different [INES Mapper 043](../INES_Mapper_043/) and [NES 2.0 Mapper 311](../NES_2.0_Mapper_311/), UNIF MAPR **UNL-SMB2J** is used.

- *Super Mario Bros. 2 (J)* (Whirlwind Manu LE10, [hangs after completing level 4-4 on real hardware](https://archive.org/details/youtube-M2-ovBQW978))
- *Volleyball (J)* (Whirlwind Manu LE08)
- *Zanac (J)* (Whirlwind Manu LF11)

FCEUX' source code says that the PCB for both the CHR-ROM and CHR-RAM variants is named *09-034A*.

# Banks {#Banks}

- CPU \$6000-\$7FFF: Switchable 8 KiB PRG-ROM bank (#4-#5)
- CPU \$8000-\$FFFF: Fixed 32 KiB PRG-ROM bank \#0
- PPU \$0000-\$1FFF: Unbanked 8 KiB CHR-ROM or CHR-RAM.

# Registers {#Registers}

## <span id="Unknown_Registers_.28.244042-.244055.2C_read.29"></span>Unknown Registers (\$4042-\$4055, read) {#Unknown_Registers_.28.244042-.244055.2C_read.29}

Reading from these addresses must return \$FF rather than open bus, otherwise SMB2J will freeze.

## <span id="PRG_Bank_Select_.28.244027.2C_write.29"></span>PRG Bank Select (\$4027, write) {#PRG_Bank_Select_.28.244027.2C_write.29}

Mask: Unknown

    Bit 7654 3210
        ---------
        .... ...B
                +- 0: CPU $6000-$7FFF to 8 KiB bank #4
                   1: CPU $6000-$7FFF to 8 KiB bank #5

On games with only 32 KiB ROM, \$6000-\$7FFF will (as usual) wrap to banks \#0 and \#1.

## <span id="IRQ_Control_.28.244068.2C_write.29"></span>IRQ Control (\$4068, write) {#IRQ_Control_.28.244068.2C_write.29}

Mask: Unknown

    Bit 7654 3210
        ---------
        .... ...I
                +- 0: Disable and Acknowledge IRQ, and reset counter
                   1: Enable IRQ

According to [FCEUX](https://github.com/TASVideos/fceux/blob/master/src/boards/09-034a.cpp), when enabled, IRQ counter increases on every M2 cycle until it reaches 5750, upon which an IRQ is fired. While this does work nicely with the ROM image of LE10, the [hardware video recording](https://archive.org/details/youtube-M2-ovBQW978) seems to indicate that the actual hardware does not have a functioning IRQ counter.
