---
title: "INES Mapper 167"
url: "/wiki/INES_Mapper_167/"
description: "iNES Mapper 167 denotes the circuit board for Subor's Chinese 小霸王 中英文电脑学习机 IV (\"Subor Chinese and English Computer Learning Machine IV\") and Russian Сюбор…"
wikiPage: "INES_Mapper_167"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**iNES Mapper 167** denotes the circuit board for Subor's Chinese **小霸王 中英文电脑学习机 IV** ("Subor Chinese and English Computer Learning Machine IV") and Russian **Сюбор Обучающий Компьютер** ("Subor Educational Computer") cartridges.

# Overview {#Overview}

- CPU \$6000-\$7FFF: 8 KiB unbanked PRG-RAM
- CPU \$8000-\$BFFF: 16 KiB fixed or switchable window into 1 MiB PRG-ROM
- CPU \$C000-\$FFFF: 16 KiB fixed or switchable window into 1 MiB PRG-ROM
- PPU \$0000-\$1FFF: 8 KiB unbanked CHR-RAM
- Nametable arrangement: Switchable between Horizontal and Vertical

# Registers {#Registers}

There are four writable registers in the CPU \$8000-\$FFFF address range.

**Register 1: Responds to writes to CPU \$8000-\$9FFF**

    D~[...F ...N] A~[100. .... .... ....]
          |    +- Nametable arrangement
          |       0: Horizontal
          |       1: Vertical
          +------ PRG A19 for switchable window (XOR'd with "f" bit from register 2)

**Register 2: Responds to writes to CPU \$A000-\$BFFF**

    D~[...f MM..] A~[101. .... .... ....]
          | ++--- PRG-ROM bankswitching mode
          |       0: UNROM-512 with fixed bank number $20
          |       1: Reverse UNROM-512 with fixed bank number $1F
          |       2: NROM-256 with PRG A14 being inverted from CPU A14
          |       3: same as mode 2
          +------ PRG A19 for switchable window (XOR'd with "F" bit from register 1)

**Register 3: Responds to writes to CPU \$C000-\$DFFF**

    D~[...E DCBA] A~[110. .... .... ....]
          +-++++- PRG A18..A14 for switchable window (XOR'd with "edcba" bits from register 4)

**Register 4: Responds to writes to CPU \$E000-\$FFFF**

    D~[...e dcba] A~[111. .... .... ....]
          +-++++- PRG A18..A14 for switchable window (XOR'd with "EDCBA" bits from register 3)

Each bank number bit for the switchable window exists twice. The final bank number is the XOR'd value of both bits. This provides the necessary flexibility for including the various mini-games, that come from various sources, on the cartridges. The combination of mode and XOR'd bank bits results in the following three possible layouts:

<table class="wikitable">
<thead>
<tr>
<th>Mode</th>
<th colspan="2">16 KiB bank number (binary) for CPU address range ...</th>
</tr>
</thead>
<tbody>
<tr>
<th>8000-BFFF</th>
<th>C000-FFFF</th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>FE'DCBA XOR fe'dcba</td>
<td>10'0000</td>
</tr>
<tr>
<td>1</td>
<td>01'1111</td>
<td>FE'DCBA XOR fe'dcba</td>
</tr>
<tr>
<td>2/3</td>
<td>FE'DCB1 XOR fe'dcb0</td>
<td>FE'DCB0 XOR fe'dcb0</td>
</tr>
</tbody>
</table>

# Errata {#Errata}

[INES Mapper 166](./) had been used at one time to denote mapper 167 ROM image files with an incorrect bank order, namely, a fixed bank \$07 rather than \$20 in mode 0, and a non-inverted PRG A14 in modes 2 and 3.
