---
title: "NES 2.0 Mapper 347"
url: "/wiki/NES_2.0_Mapper_347/"
description: "NES 2.0 Mapper 347 is used for Kaiser's cartridge conversion of the Famicom Disk System game Yume Koujou: Doki Doki Panic. Its UNIF board name is UNL-KS7030."
wikiPage: "NES_2_0_Mapper_347"
wikiCategories:
  - name: "FDS conversion mappers"
    url: "/category/FDS_conversion_mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

NES 2.0 Mapper 347 is used for Kaiser's cartridge conversion of the Famicom Disk System game *Yume Koujou: Doki Doki Panic*. Its UNIF board name is **UNL-KS7030**.

# Banks {#Banks}

A typical example of a board designed to accomodate a single (converted) FDS game while using only 8 KiB of PRG-RAM, its memory map is quite complicated.

- CPU \$6000-\$6BFF: 3 KiB PRG-RAM bank
- CPU \$6C00-\$6FFF: 1 KiB PRG-ROM bank, selected by register at \$9000
- CPU \$7000-\$7FFF: 4 KiB PRG-ROM bank, selected by register at \$8000
- CPU \$8000-\$B7FF: 14 KiB PRG-ROM bank, hard-wired
- CPU \$B800-\$BFFF: 2 KiB PRG-RAM bank
- CPU \$C000-\$CBFF: 3 KiB PRG-ROM bank, selected by register at \$9000
- CPU \$CC00-\$D7FF: 3 KiB PRG-RAM bank
- CPU \$D800-\$FFFF: 10 KiB PRG-ROM bank, hard-wired

# Registers {#Registers}

## <span id="PRG-ROM_Bank_Select_.231.2FMirroring_Select_.28.248000-.248FFF.2C_write.29"></span>PRG-ROM Bank Select \#1/Mirroring Select (\$8000-\$8FFF, write) {#PRG-ROM_Bank_Select_.231.2FMirroring_Select_.28.248000-.248FFF.2C_write.29}

    A~FEDC BA98 7654 3210
      -------------------
      1000 .... .... MBBB
                     |+++- Select 4 KiB PRG-ROM bank at CPU $7000-$7FFF
                     +---- Select nametable mirroring type
                            0: Vertical
                            1: Horizontal

## <span id="PRG-ROM_Bank_Select_.232_.28.249000-.249FFF.2C_write.29"></span>PRG-ROM Bank Select \#2 (\$9000-\$9FFF, write) {#PRG-ROM_Bank_Select_.232_.28.249000-.249FFF.2C_write.29}

    A~FEDC BA98 7654 3210
      -------------------
      1001 .... .... BBBB
                     ++++- Select 4 KiB PRG-ROM bank,
                           the first 3 KiB of which mapped to CPU $C000-$CBFF,
                           the second 1 KiB of which mapped to CPU $6C00-$6FFF

# Notes {#Notes}

- As the [FCEUX](https://github.com/OpenEmu/FCEU-Core/blob/master/src/boards/ks7030.cpp) source code comment indicates, the actual bank order in the 128 KiB mask ROM was unknown until July 2020. Emulators expected the ROM image to be laid out like this:
  - the first 32 KiB to contain the eight banks selected by register \$8000 mapped to \$7000-\$7FFF;
  - the next 64 KiB to contain the sixteen banks selected by register \$9000, with the first 1 KiB mapped to CPU \$6C00-\$6FFF and the second 3 KiB mapped to CPU \$C000-\$CBFF;
  - the final 32 KiB mapped to CPU \$8000-\$FFFF except where replaced by RAM and the switchable PRG-ROM bank.
- The actual mask ROM layout is as follows:
  - the first 64 KiB contain the sixteen banks selected by register \$9000, with the first 3 KiB mapped to CPU \$C000-\$CBFF and the second 1 KiB mapped to CPU \$6C00-\$6FFF;
  - the next 32 KiB contain the eight banks selected by register \$8000 mapped to \$7000-\$7FFF;
  - the final 32 KiB mapped to CPU \$8000-\$FFFF except where replaced by RAM and the switchable PRG-ROM bank.
- The conversion retains all writes to the FDS sound channel registers, so an emulator could provide the expansion sound channel even though the original Kaiser cartridge did not.
- A different cartridge conversion of the same game is described by [INES Mapper 103](../INES_Mapper_103/). Its memory map is simpler, at the cost of requiring 16 KiB of WRAM instead of just 8 KiB.
