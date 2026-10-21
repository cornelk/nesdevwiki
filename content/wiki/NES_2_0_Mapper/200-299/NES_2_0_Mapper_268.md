---
title: "NES 2.0 Mapper 268"
url: "/wiki/NES_2.0_Mapper_268/"
description: "NES 2.0 Mapper 268 denotes cartridges using the AA6023 ASICs, labelled SMD132/SMD133 when packaged, MMC3 clones that can address up to 32 MiB of PRG-ROM, 256 KiB of CHR-ROM or -RAM, and 8 KiB of…"
wikiPage: "NES_2_0_Mapper_268"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
---

**NES 2.0 Mapper 268** denotes cartridges using the **AA6023** ASICs, labelled **SMD132**/**SMD133** when packaged, [MMC3](../MMC3/) clones that can address up to 32 MiB of PRG-ROM, 256 KiB of CHR-ROM or -RAM, and 8 KiB of WRAM. The ASIC provides six outer bank registers. Their address is selected by a solder pad labelled "5/6K", whose setting is denoted by NES 2.0 Submapper bit 0. One variant can even address 64 MiB of PRG-ROM.

| Setting | Address Range | NES 2.0 Submapper bit 0 | UNIF MAPR |
|----|----|----|----|
| right | \$6000-\$6FFF (submappers except 2/3)/\$7000-\$7FFF (submappers 2/3) | 0 | **COOLBOY** (without prefix) |
| left | \$5000-\$5FFF | 1 | **MINDKIDS** (without prefix) |

NES 2.0 Submapper bits 1-3 distinguish between incompatible revisions and wiring variants:

<table class="wikitable sortable">
<thead>
<tr>
<th>NES 2.0 Submappers</th>
<th>PCB Name</th>
<th>Notes</th>
</tr>
</thead>
<tbody>
<tr>
<td rowspan="4">0/1</td>
<td><strong>COOLBOY</strong>/<strong>MINDKIDS</strong>/<strong>KT-008</strong></td>
<td>CHR-RAM only, max. 32 MiB PRG-ROM</td>
</tr>
<tr>
<td><strong>YH2018A</strong></td>
<td>CHR-RAM only, max. 64 MiB PRG-ROM</td>
</tr>
<tr>
<td><strong>JTH-813</strong></td>
<td>CHR-ROM only, max. 8 MiB PRG-ROM, max. 1 MiB CHR-ROM</td>
</tr>
<tr>
<td><strong>KT-???</strong></td>
<td>CHR-ROM+CHR-RAM; uses the ASIC's mixed CHR-ROM/RAM functionality selected by register <a href="#.24xxx4">$xxx4</a></td>
</tr>
<tr>
<td>2/3</td>
<td><strong>GH2009_V01</strong>/<strong>SMD173C_60</strong>/<strong>SMD173C_L1</strong></td>
<td>Later revision of 0/1. Die version <strong>AA6023B</strong>, chip still labelled as <strong>SMD133</strong></td>
</tr>
<tr>
<td>4/5</td>
<td><strong>KP-6022</strong>,<strong>LD622D</strong></td>
<td>Max. 4 MiB PRG-ROM</td>
</tr>
<tr>
<td>6/7</td>
<td><strong>J-852C</strong></td>
<td>Max. 128 KiB of CHR-RAM, as CHR A17 selects between two PRG-ROM chips</td>
</tr>
<tr>
<td>8/9</td>
<td><strong>SMD72A_V5S_V01</strong></td>
<td>Max. 2 MiB PRG-ROM, Max. 256 KiB CHR-RAM that can be write-protected</td>
</tr>
<tr>
<td>10/11</td>
<td><strong>SMD172C-L1</strong></td>
<td>Max. 8 MiB PRG-ROM, with single-screen mirroring</td>
</tr>
</tbody>
</table>

## Registers {#Registers}

In the \$6000-\$6FFF/\$7000-\$7FFF setting, the ASIC's registers registers overlay, but are functionally independent from, any PRG-RAM present.

Masks: \$F007

### <span id=".24xxx0"></span>\$xxx0 {#.24xxx0}

Submappers 0/1/2/3:

     7  bit  0
     ---- ----
     ABCC DEEE
     |||| ||||
     |||| |+++-- PRG offset (PRG A19, A18, A17)
     |||| +----- Alternate CHR A17
     ||++------- PRG offset (PRG A24, A23)
     ||++------- CHR offset (CHR A19, A18) on CHR-ROM-having boards
     |+--------- PRG mask (PRG A17 from 0: MMC3; 1: offset)
     +---------- CHR mask (CHR A17 from 0: MMC3; 1: alternate)

Submappers 4/5:

     7  bit  0
     ---- ----
     ABCC DEEE
     |||| ||||
     |||| |+++-- PRG offset (PRG A19, A18, A17)
     |||| |++--- CHR offset (A21, A20)
     |||| +----- Alternate CHR A17
     ||++------- PRG offset (PRG A21, A20)
     ||++------- CHR offset (CHR A19, A18) on CHR-ROM-having boards
     |+--------- PRG mask (PRG A17 from 0: MMC3; 1: offset)
     +---------- CHR mask (CHR A17 from 0: MMC3; 1: alternate)

Submappers 6/7:

     7  bit  0
     ---- ----
     ABCC DEEE
     |||| ||||
     |||| |+++-- PRG offset (PRG A19, A18, A17)
     |||| +----- PRG chip select
     ||++------- PRG offset (PRG A24, A23), CHR offset (CHR A19, A18)
     |+--------- PRG mask (PRG A17 from 0: MMC3; 1: offset)
     +---------- PRG chip select source (0: from MMC3 CHR A17; 1: from D bit)

Both PRG chips always have the same size, so that the first half of a submapper 6/7 ROM file represents the content of the first, and the second half of the second chip.

Submappers 8/9:

     7  bit  0
     ---- ----
     AB.C DEEE
     || | ||||
     || | |+++-- PRG offset (PRG A19, A18, A17)
     || | +----- Alternate CHR A17
     || +------- 1=Write-protect CHR-RAM
     |+--------- PRG mask (PRG A17 from 0: MMC3; 1: offset)
     +---------- CHR mask (CHR A17 from 0: MMC3; 1: alternate)

Submappers 10/11:

     7  bit  0
     ---- ----
     ABSC DEEE
     |||| ||||
     |||| |+++-- PRG offset (PRG A19, A18, A17)
     |||| +----- Alternate CHR A17
     |||+------- CIRAM A10 if S=0
     ||+-------- 0=1-Screen mirroring via bit C
     ||          1=Normal mirroring via MMC3 register A000
     |+--------- PRG mask (PRG A17 from 0: MMC3; 1: offset)
     +---------- CHR mask (CHR A17 from 0: MMC3; 1: alternate)

### <span id=".24xxx1"></span>\$xxx1 {#.24xxx1}

Submappers 0/1/6/7/10/11:

     7  bit  0
     ---- ----
     GHIJ KKLS
     |||| |||+-- SC0 bit (see below)
     |||| ||+--- GNROM mode bank PRG size (1: 32 KiB bank, PRG A14=CPU A14; 0: 16 KiB bank, PRG A14=offset A14)
     |||+-++---- PRG offset (in order: PRG A20, A22, A21)
     ||+-------- PRG mask (PRG A20 from 0: offset; 1: MMC3)
     |+--------- PRG mask (PRG A19 from 0: offset; 1: MMC3)
     +---------- PRG mask (PRG A18 from 0: MMC3; 1: offset)

Submappers 2/3:

     7  bit  0
     ---- ----
     GHIL JKKx
     |||| |||
     |||| +++--- PRG offset (in order: A20, A21, A22)
     |||+------- GNROM mode bank PRG size (0: 32 KiB bank, PRG A14=CPU A14; 1: 16 KiB bank, PRG A14=offset A14)
     ||+-------- PRG mask (PRG A20 from 0: offset; 1: MMC3)
     |+--------- PRG mask (PRG A19 from 0: offset; 1: MMC3)
     +---------- PRG mask (PRG A18 from 0: MMC3; 1: offset)

Note that both meaning and position of the L bit is flipped between submappers 0/1 and 2/3, and the order of bits A20-A22.

Submappers 4/5/8/9:

     7  bit  0
     ---- ----
     GHIx xxLx
     |||    |
     |||    +--- GNROM mode bank PRG size (1: 32 KiB bank, PRG A14=CPU A14; 0: 16 KiB bank, PRG A14=offset A14)
     ||+-------- PRG mask (PRG A20 from 0: offset; 1: MMC3)
     |+--------- PRG mask (PRG A19 from 0: offset; 1: MMC3)
     +---------- PRG mask (PRG A18 from 0: MMC3; 1: offset)

### <span id=".24xxx2"></span>\$xxx2 {#.24xxx2}

     7  bit  0
     ---- ----
     STTT MMMM
     |||| ||||
     |||| ++++-- CHR offset for GNROM mode (CHR A16, A15, A14, A13)
     |+++------- GNROM CHR bitmask (ANDed with CHR A16, A15, A14 in MMMM)
     +---------- GNROM CHR bitmask lock (1: prevents subsequent updates to S or T bits)

### <span id=".24xxx3"></span>\$xxx3 {#.24xxx3}

     7  bit  0
     ---- ----
     NPZP QQRS
     |||| |||+-- SC1 bit (see below)
     |||| +++--- PRG offset for GNROM mode (PRG A16, A15, A14)
     |||+------- 1: GNROM mode; 0: MMC3 mode
     ||||         (1: PRG A16...13 from QQ, L, R, CPU A14, A13 + CHR A16...10 from MMMM, PPU A12...10;
     ||||          0: PRG A16...13 from MMC3 + CHR A16...A10 from MMC3 )
     ||+-------- 1: Also enable PRG RAM in $5000-$5FFF
     |+-+------- Banking mode
     |+--------- "Weird MMC3 mode"
     +---------- Lockout (prevent further writes to all registers but the one at $xxx2, only works in MMC3 mode)

### <span id=".24xxx4"></span>\$xxx4 {#.24xxx4}

     7  bit  0
     ---- ----
     UUUU UUUV
     |||| ||||
     |||| |||+--- 0: Select CHR-ROM only, 1: Select mixed CHR-RAM/CHR-ROM mode
     ++++-+++---- MMC3 bank register value D1-D7 that select CHR-RAM rather than CHR-ROM

This register is only meaningful on cartridges that mount both CHR-ROM and CHR-RAM. By writing the appropriate value to D1-D7, the functionality of [INES Mapper 074](../INES_Mapper_074/) and [INES Mapper 194](../INES_Mapper_194/) is replicated.

### <span id=".24xxx5"></span>\$xxx5 {#.24xxx5}

     7  bit  0
     ---- ----
     xxxW XXYY
        | ||||
        | ||++--- Functionality unclear without extra hardware on PCB
        | ++----- Functionality unclear without extra hardware on PCB (CHR)
        +-------- 0: Normal MMC3 oversize mode
                  1: mega-UNROM mode; uses CHR A16 and A17 bits as PRG A18 and A19

## Register interpretation {#Register_interpretation}

The following things are always true, regardless of mode:

- The CCKK bits are always connected to PRG A24..A21.
- The PRG and CHR mask controls (I, H, G, B, & A) always control whether PRG A20..A17 and CHR A17 are connected to the internal MMC3 registers, or the literal values in other controls (J, EEE, & D).

### <span id="Banking_mode_.3D_.2400"></span>Banking mode = \$00 {#Banking_mode_.3D_.2400}

“Normal” oversize MMC3 mode.

Supports MMC3 anywhere between 128 KiB and 2 MiB, otherwise identical to ordinary [MMC3](../MMC3/). Goofy mask values will produce goofy memory layouts.

The combined register of CCKKJEEE could be thought of as a way to select one of 256 different 128 KiB PRG MMC3 games, or CCKK selecting one of 16 different 2 MiB MMC3 games.

Lockout works in this mode.

### <span id="Banking_mode_.3D_.2410"></span>Banking mode = \$10 {#Banking_mode_.3D_.2410}

“Normal” GNROM mode.

PRG A16 and A15 come from the QQ controls.

When in 16 KiB PRG mode (L control), PRG A14 comes from the R control. Otherwise, it comes from CPU A14.

CHR A16, A15, A14, and A13 come from MMMM controls. CHR A12, A11, and A10 come from PPU A12, A11, and A10.

The combined register of CCKKJEEEQQR could be thought of as a way to select one of 2048 different 16 KiB NROM games.

Note: The CHR and PRG mask registers STILL AFFECT THIS MODE, allowing weird splicing of MMC3 banking (with 128 KiB granularity) overlaying GNROM-style banking.

Lockout DOES NOT work in this mode.

### <span id="Banking_mode_.3D_.2440"></span>Banking mode = \$40 {#Banking_mode_.3D_.2440}

“Weird” MMC3 mode.

If [MMC3 PRG ROM bank mode](../MMC3/#Bank_select_.28.248000-.249FFE.2C_even.29) is "normal" (8+8+16F), then the PRG banks at 0xC000 and 0xE000 are fixed to 8 KiB bank 0, instead of 0xFE and 0xFF. Supposedly this effect is completely ignored when MMC3 PRG ROM bank mode is inverted (8F+8+8+8F).

The second half of each of the 2 KiB CHR banks is replaced with bank 0, and the LSB of MMC3 registers 0 and 1 work (instead of being replaced by PPU A10).

Otherwise, behaves as “normal” MMC3 mode above.

Supposedly this was intended be used with extra hardware to enable an MMC4-like mode, perhaps similar to [iNES Mapper 165](../INES_Mapper_165/).

### <span id="Banking_mode_.3D_.2450"></span>Banking mode = \$50 {#Banking_mode_.3D_.2450}

“Weird” GNROM-ish mode.

Combine the rules mentioned above for “Normal” GNROM mode with “Weird” MMC3 mode. (These descriptions should not be contradictory.)

Lockout DOES NOT work in this mode.

# 64 MiB mapping {#64_MiB_mapping}

The **YH2018A** PCB uses the SC0 (\$6001 bit 0) and SC1 (\$6001 bit 1) bits (see above) to address two halves of 32 MiB for a total of 64 MiB. To do so, the ASIC's SL0 input pin is connected to GND, and the ASIC's OA0 output is connected to the ROM chip's PRG A25 input. This results in the following behaviors:

    $xxx1.0  $xxx3.0  Meaning
    0        0        PRG A25=PRG A0, i.e. bytes at even addresses come from the first 32 MiB and bytes at odd addresses come from the second 32 MiB half.
    0        1        PRG A25=1, i.e. second 32 MiB half
    1        x        PRG A25=0, i.e. first 32 MiB half

After reset, both bits are cleared, so either the code at the reset vector must be carefully placed with bytes alternating between the two 32 MiB halves, or (as the existing multicarts do) must be duplicated in both 32 MiB halves.

# WRAM usage {#WRAM_usage}

MMC3-mode operation with WRAM in the \$6000-\$6FFF (submapper 0) setting involves setting up the extra registers with the MMC3's WRAM disabled, setting the Lockout bit, then enabling WRAM again:

    LDA #$00
    STA $A001   ; Disable WRAM
    ; initialize PRG/CHR offsets and masks
    ; ...
    LDA #$80
    STA $6003   ; Enable Lockout
    STA $A001   ; Enable WRAM

Because Lockout does not work in GNROM mode, it is impossible to have GNROM game use WRAM in the \$6000-\$6FFF (submapper 0) setting, as any WRAM write would inadvertently change the outer bank registers. The one known Chinese RPG (楚留香新传) running in GNROM mode with WRAM uses the \$5000-\$5FFF (submapper 1) setting.

# Game List {#Game_List}

Multicarts, mostly using the "Coolboy" or "Mindkids" circuit board:

- *218-in-1 Real Game*
- *MegaMan 8-in-1 (PL-0006)*
- *Data East All-Star Collection*
- *PocketGames 150-in-1*
- *Super 208-in-1*
- *Super 340-in-1*
- *Super 360-in-1*
- *Super 402-in-1*
- *Super Game 143-in-1*, a.k.a. *The Best Games of NES*
- *深圳市仁顺科技 400-in-1 Real Game*
- *霸王小子 500-in-1*

Chinese RPGs with more than 512 KiB of PRG-ROM:

- *大话水浒* (Dàhuà Shuǐhǔ)
- *楚留香新传* (Chǔliúxiāng xīn Zhuàn)
- *魔道劫* (Módào Jié)
- *口袋妖怪: 叶绿 (vol. 1-3)* (Kǒudài Yāoguài: Yè Lǜ), a.k.a. *Pokémon Leaf-Green*
- *最终幻想 1 - 黑暗篇* (Zuìzhōng Huànxiǎng 1 - Hēi'àn Piān), a.k.a. *Final Fantasy X, Part 1*
- *最终幻想 2 - 光明篇* (Zuìzhōng Huànxiǎng 2 - Guāngmíng Piān), a.k.a. *Final Fantasy X, Part 2*, original release only
- *最终幻想 3 - 黑暗篇* (Zuìzhōng Huànxiǎng 3 - Zhōngjié Piān), a.k.a. *Final Fantasy X, Part 3*
- *神奇宝贝* (Shénqí Bǎobèi), a.k.a. *Pokémon Diamond*

Other games that do not merely use the ASIC as a simple MMC3 clone:

- *Gravity Trooper Metal Storm* (Retro-Bit release)
- *Holy Diver* (Retro-Bit release)
- *Quest Forge* (Piko Interactive release), although the underlying game is simply NROM

A great number of Chinese cartridges use the ASIC as a simple MMC3 clone and thus are usually found set to [INES Mapper 004](../MMC3/). They can be recognized by writing \$80 to \$6003 in their initialization code.

# Official documentation {#Official_documentation}

- [File:AA6023Av1 1 2-OK.pdf](../../wiki-images/AA6023Av1_1_2-OK.pdf) - datasheet for AA6023A
- [File:AA6023Bv1.0-ok.pdf](../../wiki-images/AA6023Bv1.0-ok.pdf) - datasheet for AA6023B
- [File:YBB52LQF064006 SMD132 AA6023A.pdf](../../wiki-images/YBB52LQF064006_SMD132_AA6023A.pdf) - pinout for SMD132
- [SMD133 (AA6023) pinout](../SMD133_%28AA6023%29_pinout/)
- [File:SMD129.pcb](../../wiki-images/SMD129.pcb) - PCB layout of one of "COOLBOY" cartridges (for Altium Designer)

# Notes {#Notes}

- No cart has been discovered so far that makes use of the "weird" mode.
- Some new cartridges from Spritemind (also known by their brand name Mindkids) have /WE and /OE pins connected to mapper, which allows you to rewrite flash memory without soldering. This also allows console to write data to the cartridge. No cart has been discovered so far that makes use of this feature, but this can be used for homebrew.
- AA6023's IRQ operator follows MMC3 "NEC" behavior, i.e. a latch value of zero disables IRQ.
- See also: [COOLGIRL](../NES_2.0_Mapper_342/)
