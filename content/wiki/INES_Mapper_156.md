---
title: "INES Mapper 156"
url: "/wiki/INES_Mapper_156/"
description: "iNES Mapper 156 is assigned with unlicensed Korean boards using the DAOU ROM Controller DIS23C01 DAOU 245 ASIC from Daou Infosys."
wikiPage: "INES_Mapper_156"
wikiCategories:
  - name: "INES Mappers"
    url: "/category/INES_Mappers/"
---

***iNES Mapper 156*** is assigned with unlicensed Korean boards using the **DAOU ROM Controller DIS23C01 DAOU 245** ASIC from [Daou Infosys](http://community.fandom.com/wiki/c:bootleggames:Daou_Infosys).

Example games:

- Metal Force
- General's Son
- Koko Adventure (Buzz & Waldog)

## Banks {#Banks}

- CPU \$6000-\$7FFF: 8 KiB RAM
- CPU \$8000-\$BFFF: 16 KiB switchable PRG ROM bank
- CPU \$C000-\$FFFF: 16 KiB PRG ROM bank, fixed to the last bank
- PPU \$0000-\$03FF: 1 KiB switchable CHR ROM bank
- PPU \$0400-\$07FF: 1 KiB switchable CHR ROM bank
- PPU \$0800-\$0BFF: 1 KiB switchable CHR ROM bank
- PPU \$0C00-\$0FFF: 1 KiB switchable CHR ROM bank
- PPU \$1000-\$13FF: 1 KiB switchable CHR ROM bank
- PPU \$1400-\$17FF: 1 KiB switchable CHR ROM bank
- PPU \$1800-\$1BFF: 1 KiB switchable CHR ROM bank
- PPU \$1C00-\$1FFF: 1 KiB switchable CHR ROM bank

## Registers {#Registers}

### <span id="CHR-ROM_Bank_Select_.28.24C000-.24C00F.29"></span>CHR-ROM Bank Select (\$C000-\$C00F) {#CHR-ROM_Bank_Select_.28.24C000-.24C00F.29}

      Upper       Lower
    7  bit  0   7  bit  0
    ---------   ---------
    .... ...H   LLLL LLLL
            |   |||| ||||
            +---++++-++++- Select 9-bit bank number for 1 KiB CHR bank

<table class="wikitable">
<thead>
<tr>
<th colspan="2">Write to CPU addresses</th>
<th>1KB CHR bank affected</th>
</tr>
</thead>
<tbody>
<tr>
<th>(high 1 bit)</th>
<th>(low 8 bits)</th>
<th></th>
</tr>
&#10;<tr>
<td>$C004</td>
<td>$C000</td>
<td>$0000-$03FF</td>
</tr>
<tr>
<td>$C005</td>
<td>$C001</td>
<td>$0400-$07FF</td>
</tr>
<tr>
<td>$C006</td>
<td>$C002</td>
<td>$0800-$0BFF</td>
</tr>
<tr>
<td>$C007</td>
<td>$C003</td>
<td>$0C00-$0FFF</td>
</tr>
<tr>
<td>$C00C</td>
<td>$C008</td>
<td>$1000-$13FF</td>
</tr>
<tr>
<td>$C00D</td>
<td>$C009</td>
<td>$1400-$17FF</td>
</tr>
<tr>
<td>$C00E</td>
<td>$C00A</td>
<td>$1800-$1BFF</td>
</tr>
<tr>
<td>$C00F</td>
<td>$C00B</td>
<td>$1C00-$1FFF</td>
</tr>
</tbody>
</table>

### <span id="PRG-ROM_Bank_Select_.28.24C010.29"></span>PRG-ROM Bank Select (\$C010) {#PRG-ROM_Bank_Select_.28.24C010.29}

    7  bit  0
    ---- ----
    .... PPPP
         ||||
         ++++-- Select 16 KiB PRG-ROM bank at CPU $8000-$BFFF   

### <span id="Nametable_Mirroring_Type_Select_.28.24C014.29"></span>Nametable Mirroring Type Select (\$C014) {#Nametable_Mirroring_Type_Select_.28.24C014.29}

    7  bit  0
    ---- ----
    .... ..MM
           ||
           ++-- Select nametable mirroring type
                 00: Vertical
                 01: Horizontal
                 1x: One-screen, page 0

Mirroring defaults to one-screen.

## Notes {#Notes}

- The current description of it is based on FCEUX's implementation. It appears to fully-decode them, but that's probably inaccurate.
- This mapper is similar to [INES Mapper 016](../INES_Mapper_016/), but with no IRQ and EEPROM support and supports 512KB CHR.
