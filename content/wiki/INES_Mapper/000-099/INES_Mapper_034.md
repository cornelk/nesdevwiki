---
title: "INES Mapper 034"
url: "/wiki/INES_Mapper_034/"
description: "iNES Mapper 034 denotes the unrelated NINA-001/NINA-002 and BNROM circuit boards. The NES 2.0 Submapper differentiates the two:"
wikiPage: "INES_Mapper_034"
wikiCategories:
  - name: "Discrete logic mappers"
    url: "/category/Discrete_logic_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with bus conflicts"
    url: "/category/Mappers_with_bus_conflicts/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

**iNES Mapper 034** denotes the unrelated [NINA-001/NINA-002](#NINA-001.2FNINA-002) and [BNROM](#BNROM) circuit boards. The NES 2.0 Submapper differentiates the two:

- **Submapper 1** denotes the NINA-001/NINA-002 circuit board used by *Impossible Mission II*,
- **Submapper 2** denotes the I-IM/BNROM circuit board used by *Mashou*/*Deadly Towers*.

In the absence of valid submapper information, mapper 34 .NES files should be considered BNROM when the CHR-ROM size is 0-8 KiB, and NINA-001/NINA-002 when the CHR-ROM size is above 8 KiB.

# <span id="NINA-001.2FNINA-002"></span>NINA-001/NINA-002 {#NINA-001.2FNINA-002}

<div class="infobox"><div style="text-align: center"><b>NINA-001/NINA-002</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>American Video Entertainment </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=34">1 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>NINA-001, NINA-002 </td></tr><tr><td><b>PRG ROM capacity</b></td><td>64K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>8K </td></tr><tr><td><b>PRG RAM window</b></td><td>n/a </td></tr><tr><td><b>CHR capacity</b></td><td>64K </td></tr><tr><td><b>CHR window</b></td><td>4K </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>034 </td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=AVE-NINA-01">NINA-001</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=AVE-NINA-02">NINA-002</a></td></tr></tbody></table></div>

The NINA-001 PCB supports 64 KiB of PRG-ROM capacity and uses a microcontroller to disable the [CIC](../CIC_lockout_chip/). The NINA-002 PCB integrates a CIC stunner inside the PRG-ROM.

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB of unbanked PRG-RAM.
- CPU \$8000-\$FFFF: 32 KiB window into 64 KiB of PRG-ROM
- PPU \$0000-\$0FFF: 4 KiB window into 64 KiB of CHR-ROM
- PPU \$1000-\$1FFF: 4 KiB window into 64 KiB of CHR-ROM
- Nametable arrangement: Fixed horizontal arrangement/vertical mirroring without solder pad on PCB. Emulators should honor the respective [iNES header bit](../INES/#Nametable_Arrangement).

## Registers {#Registers}

The registers overlapping PRG-RAM at the same addresses mean that reading the register's address returns the last value written to the PRG-RAM, which is also the last value written to the register.

### <span id="PRG_Bank_Select_.28.247FFD.2C_write.29"></span>PRG Bank Select (\$7FFD, write) {#PRG_Bank_Select_.28.247FFD.2C_write.29}

    D~[.... ...A] A~[0111 1111 1111 1101]
               +- PRG A15 (32 KiB bank)

The power-on value is undefined; games should have a reset vector and handler in all PRG-ROM banks.

### <span id="CHR_Bank_Select_0_.28.247FFE.2C_write.29"></span>CHR Bank Select 0 (\$7FFE, write) {#CHR_Bank_Select_0_.28.247FFE.2C_write.29}

    D~[.... DCBA] A~[0111 1111 1111 1110]
            ++++- CHR A15..A12 (4 KiB bank) at PPU $0000

### <span id="CHR_Bank_Select_1_.28.247FFF.2C_write.29"></span>CHR Bank Select 1 (\$7FFF, write) {#CHR_Bank_Select_1_.28.247FFF.2C_write.29}

    D~[.... DCBA] A~[0111 1111 1111 1111]
            ++++- CHR A15..A12 (4 KiB bank) at PPU $1000

## Hardware {#Hardware}

The NINA-001 board contains a 74LS133 (13-input NAND gate), 74LS74 (dual D flip-flop), 74HCT139 (dual 1-4 decoder), 74LS00 (quad 2-input NAND gate), two 74HCT173s (4-bit tristate D flip-flop), as well as a microcontroller labelled NINA to act as a [CIC](../CIC_lockout_chip/) defeat mechanism. These 4 registers reside "on top" of PRG RAM: each write to the register goes both to the register and to the RAM location at the same address. Thus, reading the register's address returns the last value written to the RAM, which is also the last value written to the register.

One 2-NAND and one half of the 74'139 together generate /RAMSEL. Another 2-NAND inverts that to form +RAMSEL. A third 2-NAND inverts R/W to produce +WR. The 13-NAND combines +RAMSEL, +WR, and A12…A2 to generate /REGSEL. The other half of the 74'139 uses /REGSEL, A1, and A0 to generate the latching signals for the three flip-flops.

The final 2-NAND inverts PPU A12. PPU A12 and PPU /A12 are fed to the two 74'173s /OE inputs, implementing a simple quad 1-of-2 multiplexer.

A **NINA-002** board exists which [integrated a CIC stunner inside the PRG ROM](https://forums.nesdev.org/viewtopic.php?t=26661), which now has an unusual pinout. It is otherwise identical in function to the NINA-001 board.

This hardware has a trivial oversize definition that supports 8 MiB PRG and 1 MiB CHR, made by replacing the three flip-flops with 74'374s.

# BNROM {#BNROM}

<div class="infobox"><div style="text-align: center"><b>BNROM</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Irem, Nintendo </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=34">2 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>Discrete logic </td></tr><tr><td><b>Boards</b></td><td>I-IM, BNROM </td></tr><tr><td><b>PRG ROM capacity</b></td><td>128K </td></tr><tr><td><b>PRG ROM window</b></td><td>32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>8K </td></tr><tr><td><b>CHR window</b></td><td>n/a </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>Fixed H or V, controlled by solder pads </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>Yes </td></tr><tr><td><b>IRQ</b></td><td>No </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td>034, <a href="../INES_Mapper_241/">241</a>, <a href="../INES_Mapper_177/">177</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=BNROM">BNROM</a></td></tr></tbody></table></div>

## Banks {#Banks_2}

- CPU \$8000-\$FFFF: 32 KiB window into 128 KiB of PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM/-ROM
- Nametable arrangement: Fixed; solder pad selects between Horizontal and Vertical

## <span id="Register:_Bank_Select_.28.248000-.24FFFF.2C_write.29"></span>Register: Bank Select (\$8000-\$FFFF, write) {#Register:_Bank_Select_.28.248000-.24FFFF.2C_write.29}

    D~[.... ..BA] A~[1... .... .... ....]
              ++- PRG A16..A15 (32 KiB bank)

- The power-on value is undefined; games should have a reset vector and handler in all PRG-ROM banks.
- The original BNROM board is always subject to AND-type bus conflicts: the effective value is the value being written bitwise-AND'd with the PRG-ROM content at the address being written to.

## See also {#See_also}

- [iNES Mapper 241](../INES_Mapper_241/) is a variation of BxROM with 8 KiB PRG-RAM at CPU \$6000-\$7FFF, avoids bus conflicts, and optionally supports an LPC speech chip.
- [iNES Mapper 177](../INES_Mapper_177/) adds mapper-controlled nametable arrangement, in addition to having 8 KiB PRG-RAM and without bus conflicts.
