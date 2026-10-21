---
title: "NES-EVENT"
url: "/wiki/NES-EVENT/"
description: "NES-EVENT is the MMC1-based board used for Nintendo World Championships, a multicart including timed versions of Super Mario Bros., Rad Racer, and Tetris. Like SUROM and SXROM, it repurposes the MMC1…"
wikiPage: "NES_EVENT"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>NES-EVENT</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Nintendo </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=105">1 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>NES-EVENT </td></tr><tr><td><b>PRG ROM capacity</b></td><td>256K </td></tr><tr><td><b>PRG ROM window</b></td><td>16K + 16K fixed or 32K </td></tr><tr><td><b>PRG RAM capacity</b></td><td>None </td></tr><tr><td><b>CHR capacity</b></td><td>8K RAM <p><br/></p></td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>Yes </td></tr><tr><td><b>Audio</b></td><td>No </td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">105</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?ines=105">iNES 105</a></td></tr></tbody></table></div>

**NES-EVENT** is the [MMC1](../MMC1/)-based board used for *Nintendo World Championships*, a multicart including timed versions of *Super Mario Bros.*, *Rad Racer*, and *Tetris*. Like [SUROM](../MMC1/) and [SXROM](../MMC1/), it repurposes the MMC1 CHR address bits, and MMC1 wired this way is [mapper 105](./).

## Parts {#Parts}

- 256K of PRG ROM as two 27C010 UV EPROMs
- 8K of CHR RAM
- 8K of PRG RAM
- Several 4000-series and 7400-series discrete ICs

## Register {#Register}

MMC1 register \$A000 is mapped like this:

    43210
    |||||
    ||||+- Not used (ordinarily selects bottom or top half of CHR RAM)
    |+++-- 0-3: Select 32 KiB bank in $8000-$FFFF from lower 128KB ROM
    |      4-7: Normal MMC1 behavior from upper 128KB ROM
    +----- 0: Run timer
           1: Reset timer

The first 32 KiB is hardwired until the timer is started (write 0 then 1 to bit 4) for the first time.

The other ICs implement a 30-bit up counter clocked by M2 that fires an IRQ when it reaches a high enough value. DIP switches set bits 28-25 of the counter's initial value. This results in a range of roughly 5 to 10 minutes on an NTSC console.

## Disch Notes {#Disch_Notes}

     Here are Disch's original notes:  
     ========================
     =  Mapper 105          =
     ========================

     aka
     --------------------------
     NES-EVENT

     Example Game:
     --------------------------
     Nintendo World Championships 1990

     Notes:
     ---------------------------
     This mapper is an MMC1 with crazy wiring and a huge 30-bit CPU cycle driven IRQ counter.  Registers are all
     internal and not directly accessable -- and the latch must be written to 1 bit at a time -- just like on a
     normal MMC1.  For details on how regs are written to, see mapper 001.

     This mapper has 8k CHR-RAM, and it is not swappable.

     Registers:
     ---------------------------

     Note that like a normal MMC1, registers are internal and not accessed directly.

       $8000-9FFF:   [.... PSMM]  Same as MMC1 (but CHR mode bit isn't used)

       $A000-BFFF:   [...I OAA.]
            I = IRQ control / initialization toggle
            O = PRG Mode/Chip select
            A = PRG Reg 'A'

       $C000-DFFF:   [.... ....]  Unused

       $E000-FFFF:   [...W BBBB]
            W = WRAM disable (same as MMC1)
            B = PRG Reg 'B'

     Powerup / Reset / Initialization:
     ---------------------------

       On powerup and reset, the first 32k of PRG (from the first PRG chip) is selected at $8000 *no matter what*.
     PRG cannot be swapped until the mapper has been "initialized" by setting the 'I' bit to 0, then to '1'.  This
     toggling will "unlock" PRG swapping on the mapper.

       Note 'I' also controls the IRQ counter (see below)

     PRG Setup:
     ---------------------------

       There are 2 PRG chips, each 128k.  The 'O' bit selects between the chips, and also determines which PRG Reg
     is used to select the page.

       O=0:  Use first PRG chip (first 128k), use 'A' PRG Reg, 32k swap
       O=1:  Use second PRG chip (second 128k), use 'B' PRG Reg, MMC1 style swap

       In addition, if the mapper has not been "unlocked", the first 32k of the first chip is always selected
     regardless (as if $A000 contained $00).

       Modes as listed below:

                       $8000   $A000   $C000   $E000
                     +-------------------------------+
     Uninitialized:  |             { 0 }             |  <-- use first 128k
                     +-------------------------------+
     O=0:            |             $A000             |  <-- use first 128k
                     +-------------------------------+
     O=1, P=0:       |            <$E000>            |  <-- use second 128k
                     +-------------------------------+
     O=1, P=1, S=0:  |     { 0 }     |     $E000     |  <-- use second 128k
                     +---------------+---------------+
     O=1, P=1, S=1:  |     $E000     |     {$07}     |  <-- use second 128k
                     +---------------+---------------+

     IRQ Counter:
     ---------------------------

       The 'I' bit in $A000 controls the IRQ counter.  When cleared, the IRQ counter counts up every cycle.  When
     set, the IRQ counter is reset to 0 and stays there (does not count), and the pending IRQ is acknowledged.

       The cart has 4 dipswitches which control how high the counter must reach for an IRQ to be generated.

       The IRQ counter is 30 bits wide.. when it reaches the following value, an IRQ is fired:

       [1D CBAx xxxx xxxx xxxx xxxx xxxx xxxx]
         ^ ^^^
         | |||
         either 0 or 1, depending on the corresponding dipswitch.

     So if all dipswitches are open (use '0' above), the counter must reach $20000000.
     If all dipswitches are closed (use '1' above), the counter must reach $3E000000.
     etc

       In the official tournament, 'C' was closed, and the others were open, so the counter had to reach $2800000.

Note that:

- PPU A12 is still connected to the MMC1 (so the register at \$C000 exists, and the bit in \$8000 to control CHR banking mode still exists, and writing a value with that bit set would make the register at \$C000 relevant), and
- MMC1 PRG A14 is connected to both PRG ROMs (so writes to \$8000 could select either 16+16 banking mode, and then writes to \$E000 could cause either the first or second 16K of the first ROM to be mapped at both \$8000 and \$C000 before the timer was unlocked)

## References {#References}

- [Kevin Horton's reverse engineering of NES-EVENT](http://kevtris.org/mappers/nes_custom/NES_EVENT.html)
