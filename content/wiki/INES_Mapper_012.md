---
title: "INES Mapper 012"
url: "/wiki/INES_Mapper_012/"
description: "Even as FFE 4M was assigned earlier (as \"FFE F6xxx\") than SL-5020B, the latter was emulated earlier and so is assigned submapper 0."
wikiPage: "INES_Mapper_012"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**iNES Mapper 012** denotes ...

- the 哥德 (Gouder) **SL-5020B** circuit board (**submapper 0**);
- ROM images that have been extracted from disk images for the Front Fareast Magic Card 4M [RAM cartridge](../RAM_cartridge/) (**submapper 1**).

Even as FFE 4M was assigned earlier (as "FFE F6xxx") than SL-5020B, the latter was emulated earlier and so is assigned submapper 0.

# <span id="SL-5020B_.28submapper_0.29"></span>SL-5020B (submapper 0) {#SL-5020B_.28submapper_0.29}

**Submapper 0'**s circuit board mounts a chip-on-board variant of the same Huang-1 ASIC as mappers [14](../INES_Mapper_014/), [116](../INES_Mapper_116/) and [238](../INES_Mapper_238/) plus GAL16V8 programmable logic. It is used for two games:

- Rex Soft's *Dragon Ball Z 5* cartrige;
- Version 7.5 of the homebrew hack *Ultimate Mortal Kombat III*.

The Huang-1 ASIC functions exclusively as an MMC3A clone on this PCB; [MMC3A-specific IRQ behavior](../MMC3/#IRQ_Specifics) is necessary for running *Dragon Ball Z 5*. The GAL16V8 provides an outer CHR bank register at \$4132 to support 512 KiB of CHR-ROM:

    Mask: $E100

    D~7654 3210 (write)
      ---------
      ...B ...A
         |    +- CHR A18 when PPU A12=0
         +------ CHR A18 when PPU A12=1

Because this register is implemented outside the ASIC, it is not affected by MMC3 \$8000 bit 7, and writes take effect immediately. FCEUX instead emulates it as a temporary latch register that takes effect on the next write to an MMC3 CHR bank register. The homebrew hack *Ultimate Mortal Kombat III* relies on FCEUX' behavior and would therefore not work on the original PCB.

*Dragon Ball Z 5* reads the same address to select between Chinese and English text.

    Mask: $E100

    D~7654 3210 (read)
      ---------
      .... ...A
              +- Language-determining bit value

It was previously thought that the value of this bit reflects a jumper or solder pad. Pictures of the circuit board however show no sign of either. Either such a solder pad is hidden underneath the GAL16V8, or the GAL16V8 has the value hard-wired (to Chinese text in all known copies).

See also: [Forum thread about adding parts to an MMC3A board to make this mapper](https://forums.nesdev.org/viewtopic.php?p=160268#p160268)

- [<img src="../../wiki-images/SL-5020B-front.jpg" width="120" height="54" alt="PCB front" />](../../wiki-images/SL-5020B-front.jpg)

  PCB front

- [<img src="../../wiki-images/SL-5020B-back.jpg" width="120" height="54" alt="PCB back" />](../../wiki-images/SL-5020B-back.jpg)

  PCB back

# <span id="Front_Fareast_Magic_Card_4M_.28submapper_1.29"></span>Front Fareast Magic Card 4M (submapper 1) {#Front_Fareast_Magic_Card_4M_.28submapper_1.29}

**Submapper 1** denotes ROM images that have been extracted from disk images for the *Front Fareast Magic Card 4M* [RAM cartridge](../RAM_cartridge/). It represents games whose [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Magic_Card_1M.2F2M.2F4M_disks) denotes a Magic Card 4M disk (byte \$0 bit 7 set, bits 4 or 5 set, byte \$7=\$00). Refer to the [Super Magic Card](../Super_Magic_Card/) article for details on bankswitching. The Super Magic Card's registers are initialized to:

    ; Play mode, WRAM bank 0, 1 KiB CHR mode disabled
    $4500 = $42 

    ; PRG memory write-protected, two-screen mirroring
    $42FF = $20 | (verticalMirroring? 0x00: 0x10)

    ; 4M banking mode enabled
    $43FC = $00

    ; Initial PRG register content
    $4504 = Number of 8 KiB PRG banks -4
    $4505 = Number of 8 KiB PRG banks -3
    $4506 = Number of 8 KiB PRG banks -2
    $4507 = Number of 8 KiB PRG banks -1

The [iNES](../INES/) header may specify a [512-byte trainer](../INES/#Trainer) (corresponding to [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Front_Fareast_Magic_Card_1M.2F2M.2F4M_disks)'s byte \$0 bit 6 being set), which must be loaded to \$7000-\$71FF, be writable, and (on a hard reset) initialized by JSRing to \$7003 before JMPing to the game's reset vector.

Battery-saving of WRAM content is not supported by any Magic Card model. Hard-resetting a game while restoring previously-saved WRAM content in emulators interferes with the correct operation of the trainer's program.

Because the Magic Card 4M only had 32 KiB of CHR-RAM, CHR data is loaded into *PRG* address space at offset \$40000, i.e. the second half of its 512 KiB of PRG memory, so that the trainer program can copy portions of it to 32 KiB CHR-RAM as needed. The iNES header specifies it as CHR-ROM nonetheless so that the initial PRG registers can be correctly initialized to point to the game's reset vector.
