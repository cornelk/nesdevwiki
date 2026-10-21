---
title: "Sunsoft FME-7"
url: "/wiki/Sunsoft_FME-7/"
description: "The Sunsoft FME-7 is a mapper IC used by Sunsoft in their several games. It's nearly identical to the Sunsoft 5A and Sunsoft 5B mapper chips used only in Famicom games, with the 5B notably having…"
wikiPage: "Sunsoft_FME_7"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
  - name: "Mappers with large PRG RAM"
    url: "/category/Mappers_with_large_PRG_RAM/"
  - name: "Mappers with single-screen mirroring"
    url: "/category/Mappers_with_single-screen_mirroring/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>FME-7</b><br/><b>Sunsoft 5A/5B</b></div><table class="oddf0"><tbody><tr><td><b>Company</b></td><td>Sunsoft </td></tr><tr><td><b>Games</b></td><td><a href="https://nescartdb.com/search/advanced?ines=69">9 in NesCartDB</a></td></tr><tr><td><b>Complexity</b></td><td>ASIC </td></tr><tr><td><b>Boards</b></td><td>JLROM, JSROM,<br/>NES-BTR, others </td></tr><tr><td><b>Pinout</b></td><td><a href="../Sunsoft_5_pinout/">Sunsoft 5 pinout</a></td></tr><tr><td><b>PRG ROM capacity</b></td><td>512K (FME-7)<br/>256K (5A/5B) </td></tr><tr><td><b>PRG ROM window</b></td><td>8K×4 + 8K fixed </td></tr><tr><td><b>PRG RAM capacity</b></td><td>512K (FME-7)<br/>256K (5A/5B) </td></tr><tr><td><b>PRG RAM window</b></td><td>8K </td></tr><tr><td><b>CHR capacity</b></td><td>256K </td></tr><tr><td><b>CHR window</b></td><td>1Kx8 </td></tr><tr><td><b><a href="../Mirroring/#Nametable_Mirroring">Nametable arrangement</a></b></td><td>H, V, or 1, switchable </td></tr><tr><td><b><a href="../Bus_conflict/">Bus conflicts</a></b></td><td>No </td></tr><tr><td><b>IRQ</b></td><td>CPU cycle counter </td></tr><tr><td><b>Audio</b></td><td><a href="../Sunsoft_5B_audio/">5B only</a></td></tr><tr><td><b>iNES <a href="../Mapper/">mappers</a></b></td><td><a href="./">069</a></td></tr></tbody></table></div>

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=SUNSOFT-FME-7">Sunsoft FME-7</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=SUNSOFT-5A">Sunsoft 5A</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=SUNSOFT-5B">Sunsoft 5B</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?ines=69">iNES 069</a></td></tr></tbody></table></div>

The Sunsoft FME-7 is a mapper IC used by Sunsoft in their several games. It's nearly identical to the **Sunsoft 5A** and **Sunsoft 5B** mapper chips used only in Famicom games, with the 5B notably having expansion audio (see [Sunsoft 5B audio](../Sunsoft_5B_audio/)).

The FME-7, 5A and 5B are grouped together as **iNES Mapper 69**.

All FME-7, 5A and 5B ICs exist as a 44 pin TQFP chip: [pinout](../Sunsoft_5_pinout/)

In Europe, boards using the FME-7 were labeled as [JSROM and JLROM](../JxROM/). The FME-7 mapper was used in only one game released in the US, *Batman: Return of the Joker*. Many Japanese releases by Sunsoft used the FME-7: *Gimmick!*, *Hebereke*, *Gremlins 2* (but not in the US version), *Barcode World*, and others.

## <span id="Differences_between_FME-7.2C_5A_and_5B"></span>Differences between FME-7, 5A and 5B {#Differences_between_FME-7.2C_5A_and_5B}

Other features that aren't listed here are common through all three chips. The differences about pin layout is described in the [pinout](../Sunsoft_5_pinout/) page.

### 5A {#5A}

- Max 256 KiB PRG-ROM/PRG-RAM
- No expansion audio

### 5B {#5B}

- Max 256 KiB PRG-ROM/PRG-RAM
- Functional [expansion audio](../Sunsoft_5B_audio/)

### FME-7 {#FME-7}

- Max 512 KiB PRG-ROM/PRG-RAM
- No expansion audio

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB switchable PRG ROM or PRG RAM bank
- CPU \$8000-\$9FFF: 8 KiB switchable PRG ROM bank
- CPU \$A000-\$BFFF: 8 KiB switchable PRG ROM bank
- CPU \$C000-\$DFFF: 8 KiB switchable PRG ROM bank
- CPU \$E000-\$FFFF: 8 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR bank

## Registers {#Registers}

Configuration of the FME-7 is accomplished by first writing the command number to the Command Register at \$8000-9FFF, then writing the command's parameter byte to the Parameter Register at \$A000-BFFF.

There are 16 commands:

- **\$0-7** control CHR banking
- **\$8-B** control PRG banking
- **\$C** controls nametable mirroring
- **\$D-F** controls IRQ

On the 5B variant, there are **two additional registers** at \$C000-DFFF and \$E000-FFFF that control the audio expansion. See: [Sunsoft 5B audio](../Sunsoft_5B_audio/)

### <span id="Command_Register_.28.248000-.249FFF.29"></span>Command Register (\$8000-\$9FFF) {#Command_Register_.28.248000-.249FFF.29}

    7  bit  0
    ---- ----
    .... CCCC
         ||||
         ++++- The command number to invoke when writing to the Parameter Register

### <span id="Parameter_Register_.28.24A000-.24BFFF.29"></span>Parameter Register (\$A000-\$BFFF) {#Parameter_Register_.28.24A000-.24BFFF.29}

    7  bit  0
    ---- ----
    PPPP PPPP
    |||| ||||
    ++++-++++- The parameter to use for this command. Writing to this register invokes the command in the Command Register

## Commands {#Commands}

### <span id="CHR_Bank_0-7_.28.240-7.29"></span>CHR Bank 0-7 (\$0-7) {#CHR_Bank_0-7_.28.240-7.29}

    7  bit  0
    ---- ----
    BBBB BBBB
    |||| ||||
    ++++-++++- The bank number to select for the specified bank.

- Bank \$0 - PPU \$0000-\$03FF
- Bank \$1 - PPU \$0400-\$07FF
- Bank \$2 - PPU \$0800-\$0BFF
- Bank \$3 - PPU \$0C00-\$0FFF
- Bank \$4 - PPU \$1000-\$13FF
- Bank \$5 - PPU \$1400-\$17FF
- Bank \$6 - PPU \$1800-\$1BFF
- Bank \$7 - PPU \$1C00-\$1FFF

### <span id="PRG_Bank_0_.28.248.29"></span>PRG Bank 0 (\$8) {#PRG_Bank_0_.28.248.29}

    7  bit  0
    ---- ----
    ERbB BBBB
    |||| ||||
    ||++-++++- The bank number to select at CPU $6000 - $7FFF
    |+-------- RAM / ROM Select Bit
    |              0 = PRG ROM
    |              1 = PRG RAM
    +--------- RAM Enable Bit (6264 +CE line)
                   0 = PRG RAM Disabled
                   1 = PRG RAM Enabled

The FME-7 has up to 6 bits for PRG banking (512 KiB), though this was never used in a game. The 5A and 5B, however, support only 5 (256 KiB)—hence the lowercase 'b' above. The extra address line is instead an audio expansion line, or unused.

It is [confirmed](http://forums.nesdev.org/viewtopic.php?f=9&t=12467) that the FME-7 outputs the bank number during accesses to \$6000-\$7FFF even if RAM is enabled. Though Sunsoft never took advantage of this, it would allow making a cartridge that bank switches up to 256 KiB of PRG RAM. The FME-7 mapper in Loopy's [PowerPak](../PowerPak/) mappers, for example, supports 32 KiB.

Open bus occurs if the RAM / ROM Select Bit is 1 (RAM selected), but the RAM Enable Bit is 0 (disabled), i.e. any value in the range \$40-\$7F. This is a limited form of WRAM write protection on power-up.

NOTE: the enable bit is NOT functional with 2048×8, 32768×8, and 524288×8 RAMs, because those RAMs don't have a +CE input.

### <span id="PRG_Bank_1-3_.28.249-B.29"></span>PRG Bank 1-3 (\$9-B) {#PRG_Bank_1-3_.28.249-B.29}

    7  bit  0
    ---- ----
    ..bB BBBB
      || ||||
      ++-++++- The bank number to select for the specified bank.

- Bank \$9 - CPU \$8000-\$9FFF
- Bank \$A - CPU \$A000-\$BFFF
- Bank \$B - CPU \$C000-\$DFFF

The FME-7 has up to 6 bits for PRG banking, but the 5A and 5B support only 5.

### <span id="Nametable_Arrangement_.28.24C.29"></span>Nametable Arrangement (\$C) {#Nametable_Arrangement_.28.24C.29}

    7  bit  0
    ---- ----
    .... ..MM
           ||
           ++- Nametable arrangement
              (0: Horizontal (A10), 1: Vertical (A11),
               2: 1scA (GND),       3: 1scB (VCC))

### <span id="IRQ_Control_.28.24D.29"></span>IRQ Control (\$D) {#IRQ_Control_.28.24D.29}

    7  bit  0
    ---- ----
    C... ...T
    |       |
    |       +- IRQ Enable
    |           0 = Do not generate IRQs
    |           1 = Do generate IRQs
    +--------- IRQ Counter Enable
                0 = Disable Counter Decrement
                1 = Enable Counter Decrement

All writes to this register acknowledge an active IRQ.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> It is not yet known what will happen if this register is written to at the same time as an IRQ would have been generated.

On FME-7, the "IRQ Enable" signal permits its /EXT IRQ input pin to also assert 2A03 /IRQ; no commercial hardware used this.

### <span id="IRQ_Counter_Low_Byte_.28.24E.29"></span>IRQ Counter Low Byte (\$E) {#IRQ_Counter_Low_Byte_.28.24E.29}

    7  bit  0
    ---- ----
    LLLL LLLL
    |||| ||||
    ++++-++++- The low eight bits of the IRQ counter

### <span id="IRQ_Counter_High_Byte_.28.24F.29"></span>IRQ Counter High Byte (\$F) {#IRQ_Counter_High_Byte_.28.24F.29}

    7  bit  0
    ---- ----
    HHHH HHHH
    |||| ||||
    ++++-++++- The high eight bits of the IRQ counter

Writes to the IRQ counter registers directly set the lower or upper eight bits of the counter. Unlike on MMC3, there is no separate reload latch.

## IRQ Operation {#IRQ_Operation}

The IRQ feature of FME-7 is a CPU cycle counting IRQ generator. When enabled the 16-bit IRQ counter is decremented once per CPU cycle. When the IRQ counter is decremented from \$0000 to \$FFFF an IRQ is generated. The IRQ line is held low until it is acknowledged.

### How to Use the IRQ Generator {#How_to_Use_the_IRQ_Generator}

1.  Set the counter to the desired number of cycles minus one.
2.  Enable the IRQ generator by turning on both the IRQ Enable and IRQ Counter Enable flags of the IRQ Control command.
3.  Within the IRQ handler, write to the IRQ Control command to acknowledge the IRQ.
4.  Optional: Go back to the first step for further IRQ.

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) Test performed in 2015 by Oliveira using [IRQ acknowledge test ROM on NESdev BBS](http://forums.nesdev.org/viewtopic.php?p=142243#p142243)

## See also {#See_also}

- [Sunsoft Mapper](http://nesdev.org/sunsoft.txt) by goroh, translated by Sgt. Bowhack.
- [NES Mapper List](http://www.romhacking.net/documents/362/) by Disch
- [Comprehensive NES Mapper Document](http://nesdev.org/mappers.zip) by \Firebug\\
