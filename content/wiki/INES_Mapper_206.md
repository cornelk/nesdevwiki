---
title: "INES Mapper 206"
url: "/wiki/INES_Mapper_206/"
description: "iNES Mapper 206 is the simpler predecessor of the MMC3, and was used by Tengen and Namco. Chips used are Tengen's \"MIMIC-1\" and Namcot's \"108\", \"109\", \"118\", and \"119\", and the boards made by…"
wikiPage: "INES_Mapper_206"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>Namco 118, Tengen MIMIC-1</b><br/><b>DxROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Namco, Tengen, others </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=206">39 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>34xx, DxROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>8K + 8K + 16K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>64K </td></tr><tr><td><b>CHR window</b></td><td>2Kx2 + 1Kx4 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H/V, or 4 </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>206, <a href="../INES_Mapper_076/">76</a>, <a href="../INES_Mapper_088/">88</a>, <a href="../INES_Mapper_154/">154</a>, <a href="../INES_Mapper_095/">95</a></td></tr></tbody></table></div>

iNES Mapper 206 is the simpler predecessor of the [MMC3](../MMC3/), and was used by Tengen and Namco. Chips used are [Tengen's "MIMIC-1" and Namcot's "108", "109", "118", and "119"](../Namcot_108_family_pinout/), and the boards made by Nintendo of America that used this mapper are NES-[DxROM](../DxROM/). Many ROMS using this mapper are incorrectly listed as using MMC3, but will usually work if emulated with MMC3, and the nametable arrangement is correct, as if they were on a [TEROM or TFROM board](../TxROM/).

Compared to MMC3:

- There are no IRQs
- There is no PRG-RAM support, except for *Magical Puzzle Popils* (prototype) which expects MIMIC-1 + 8KiB battery-backed PRG-RAM.
- PRG always has the last two 8KiB banks fixed to the end.
- CHR always gives the left pattern table (\$0000-\$0FFF) the two 2KiB banks, and the right pattern table (\$1000-\$1FFF) the four 1KiB banks.
- Nametable arrangement is hardwired, one game uses 4-screen nametables (Gauntlet, DRROM).
- CHR size limit is 64KiB, PRG size limit is 128KiB.
- There are no control registers in the \$A000-\$FFFF range.
  - Writing to RAM while executing from \$8000-9FFF can cause erratic mapper writes. See [Errata](#Errata) below.

## Registers {#Registers}

Register mask: \$E001

### <span id="Bank_select_.28.248000-.249FFE.2C_even.29"></span>Bank select (\$8000-\$9FFE, even) {#Bank_select_.28.248000-.249FFE.2C_even.29}

    7  bit  0
    ---- ----
    xxxx xRRR
          |||
          +++- Specify which bank register to update on next write to Bank Data register
               0: Select 2 KB CHR bank at PPU $0000-$07FF
               1: Select 2 KB CHR bank at PPU $0800-$0FFF
               2: Select 1 KB CHR bank at PPU $1000-$13FF
               3: Select 1 KB CHR bank at PPU $1400-$17FF
               4: Select 1 KB CHR bank at PPU $1800-$1BFF
               5: Select 1 KB CHR bank at PPU $1C00-$1FFF
               6: Select 8 KB PRG ROM bank at $8000-$9FFF
               7: Select 8 KB PRG ROM bank at $A000-$BFFF

See [MMC3](../MMC3/) and note the absence of any control bits in the upper five bits of this register.

### <span id="Bank_data_.28.248001-.249FFF.2C_odd.29"></span>Bank data (\$8001-\$9FFF, odd) {#Bank_data_.28.248001-.249FFF.2C_odd.29}

    7  bit  0
    ---- ----
    xxdd DDDd
      || ||||
      ++-++++- New bank value, based on last value written to bank select register (mentioned above)

Only bits 5-1 exist for the two 2 KiB CHR banks, only bits 5-0 exist for the four 1 KiB CHR banks, and only bits 3-0 exist for the two 8 KiB PRG banks. These eight bank registers are identical to those of [MMC3](../MMC3/), except that only 128 KiB PRG and 64 KiB CHR are supported.

## Errata {#Errata}

On Namco 108, writing to \$0000-1FFF while executing from \$8000-9FFF may be seen by the mapper as a write to \$8000-9FFF, causing spurious bankswitches.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> Contemporary software avoids this problem by never executing code from this region. This is presumably why *Babel no Tou*, despite having only 32 KiB PRG, uses bankswapping and only executes bank 0 and 1 code when mapped into \$A000-BFFF. Namco 109 and MIMIC-1 are not affected by this bug. It is not known if Namco 118 and 119 are affected, but if the ten's digit indicates manufacturer, then it may be that 118 is affected and 119 is not.

## Submappers {#Submappers}

*Babel no Tou* on the **3401** PCB is the only game with 32 KiB PRG that allows—and uses!—PRG banking. All other games with 32 KiB PRG connect CPU A13 and CPU A14 directly to the PRG ROM, but fortunately they initialize their PRG registers to work on a normal board.

Submapper 0 indicates the normal PCB that allows PRG banking, while [submapper 1](../NES_2.0_submappers/#206:_1) indicate the **3407**, **3417** and **3451** PCBs with unbanked 32 KiB PRG-ROM.

## Variants {#Variants}

[Mapper 76](../INES_Mapper_076/) increases CHR to 128KiB by inflating the 1KiB CHR banks to 2KiB and making the originally-2KiB banks inaccessible.

[Mapper 88](../INES_Mapper_088/) increases CHR to 128KiB by connecting PPU's A12 line to the CHR ROM's A16 line, making tiles in \$0000 and \$1000 come from disjoint sections of ROM. Because an undersize ROM on a mapper 88 board behaves identically to mapper 206, emulators may treat these mapper numbers as synonymous.

[Mapper 154](../INES_Mapper_154/) starts with mapper 88, then adds mapper-controlled one-screen nametables.

[Mapper 95](../INES_Mapper_095/) uses the MSB to control nametable arrangement by connecting CHR A15 to CIRAM A10, much as CHR A17 controls CIRAM A10 in [TxSROM](../INES_Mapper_118/).

## References {#References}

- [FCEUX's implementation](https://sourceforge.net/p/fceultra/code/HEAD/tree/fceu/trunk/src/boards/206.cpp)
- [Enri](http://cmpslv3.stars.ne.jp/)'s reverse-engineered schematic of PCBs 3407 and 3416: <http://cmpslv3.stars.ne.jp/Famic/Fcmp206.htm>

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [Forum thread:](https://forums.nesdev.org/viewtopic.php?t=13297) Naruko's report on Namco 108 banking problems when executing from \$8000-9FFF
