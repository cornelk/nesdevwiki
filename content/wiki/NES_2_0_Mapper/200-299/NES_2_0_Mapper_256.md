---
title: "NES 2.0 Mapper 256"
url: "/wiki/NES_2.0_Mapper_256/"
description: "NES 2.0 Mapper 256 is used for games using the bankswitching behavior of the OneBus Famiclones. The UNIF board names are UNL-OneBus, BMC-OneBus and UNL-DANCE."
wikiPage: "NES_2_0_Mapper_256"
wikiCategories:
  - name: "VTxx mappers"
    url: "/category/VTxx_mappers/"
---

**NES 2.0 Mapper 256** is used for games using the bankswitching behavior of the [OneBus](../VTxx/) Famiclones. The UNIF board names are **UNL-OneBus**, **BMC-OneBus** and **UNL-DANCE**.

## Registers {#Registers}

To emulate ROM images using mapper 256, one must support bankswitching, nametable mirroring selection and IRQs using both the VTxx native registers (see the [VT02+ PRG-ROM Bankswitching](../VT02+_PRG-ROM_Bankswitching/) and [VT02+ CHR-ROM Bankswitching](../VT02+_CHR-ROM_Bankswitching/) articles for details) as well as the [MMC3 compatible registers](../VT02+_MMC3_Compatibility_Registers/).

## <span id="Single_and_split_PRG.2FCHR_address_space"></span>Single and split PRG/CHR address space {#Single_and_split_PRG.2FCHR_address_space}

On physical hardware, bankswitching by the VTxx console itself can only be used if the cartridge uses OneBus addressing, signalled by connecting certain pins on the cartridge connector to Vcc/GND, because only in OneBus mode are the CHR address lines repurposed as higher-order address lines for the single address space. Therefore, regular OneBus ROM images will only contain PRG ROM and no CHR ROM, and PRG ROM (\$41xx) and CHR ROM (\$201x) bankswitch registers apply to the same address space.

However, ROM images converted from Waixing's WXN format will have both PRG and CHR ROM, yet use the VTxx-native bankswitch registers and/or [register address scrambling](#Bank_register_address_and_CPU_opcode_scrambling). These images cannot be run unchanged on physical hardware, and were intended to be run by Waixing's proprietary VTxx emulator found on several 32-bit plug-and-play consoles. For simplicity, mapper 256 is also used for such ROM images; to emulate them, simply apply the PRG ROM (\$41xx) and CHR ROM (\$201x) bankswitch registers to the respective separate address space.

## Bank register address and CPU opcode scrambling {#Bank_register_address_and_CPU_opcode_scrambling}

Several plug-and-play consoles use a simple protection scheme intended to prevent interchanging the ROM data between different consoles. The protection scheme works by simply swapping the addresses of the bankswitch registers. Some patterns only change the VTxx-native register addresses, others apply to the MMC3 compatibility registers as well.

Other plug-and-play consoles, in particular those by Jungletac, scramble CPU opcode bytes (and only those) by swapping bits. These consoles power-on with opcode scrambling enabled, but allow it to be disabled --- for example, to run a pirated licensed NES game without modifying it --- by writing to register **411C** (submappers 11, 12 and 14):

    D~7654 3210
      ---------
      .E.. ..F.
       |     +-- Select primary CPU opcode encryption (submapper 11 only)
       |          0: disabled
       |          1: enabled (default at power-on)
       +-------- Select CPU opcode encryption (submapper 12+14 only)
                  0: disabled
                  1: enabled (default at power-on)
                 Select secondary CPU opcode encryption (submapper 11 only)
                  0: disabled (default at power-on)
                  1: enabled

or **\$4169** (submappers 13 and 15):

    D~7654 3210
      ---------
      .... ...E
              +- Select CPU opcode encryption (submappers 13+15 only)
                  0: enabled (default at power-on)
                  1: disabled

[Submappers](../NES_2.0_submappers/#NES_2.0_Mapper_256) have been allocated to denote the type of protection scheme:

<table class="wikitable">
<thead>
<tr>
<th>Submapper #</th>
<th>Name</th>
<th colspan="12">PPU bank affected by ...</th>
<th colspan="4">CPU bank affected by ...</th>
<th>CPU opcode bytes</th>
</tr>
</thead>
<tbody>
<tr>
<th>$2012</th>
<th>$2013</th>
<th>$2014</th>
<th>$2015</th>
<th>$2016</th>
<th>$2017</th>
<th>$8000.0</th>
<th>$8000.1</th>
<th>$8000.2</th>
<th>$8000.3</th>
<th>$8000.4</th>
<th>$8000.5</th>
<th>$4107</th>
<th>$4108</th>
<th>$8000.6</th>
<th>$8000.7</th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>Normal</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>1</td>
<td>Waixing VT03</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>2</td>
<td>Power Joy Supermax</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>none</td>
</tr>
<tr>
<td>3</td>
<td>Zechess/Hummer Team</td>
<td>$0800</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1800</td>
<td>$1000</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>4</td>
<td>Qishenglong</td>
<td>$1800</td>
<td>$0800</td>
<td>$1000</td>
<td>$0000</td>
<td>$1C00</td>
<td>$1400</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>5</td>
<td>Waixing VT02</td>
<td>$1400</td>
<td>$1000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>none</td>
</tr>
<tr>
<td>11</td>
<td>Vibes</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6; D5&lt;-&gt;D4 swapped, switched via $411C.1</td>
</tr>
<tr>
<td>12</td>
<td>Cheertone</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D7&lt;-&gt;D6, D1&lt;-&gt;D2 swapped, switched via $411C.6</td>
</tr>
<tr>
<td>13</td>
<td>Taikee</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D1 and D4 swapped, switched via $4169</td>
</tr>
<tr>
<td>14</td>
<td>Karaoto</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D6 and D7 swapped, switched via $411C</td>
</tr>
<tr>
<td>15</td>
<td>Jungletac</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$0000</td>
<td>$0800</td>
<td>$0000</td>
<td>$0800</td>
<td>$1000</td>
<td>$1400</td>
<td>$1800</td>
<td>$1C00</td>
<td>$8000</td>
<td>$A000</td>
<td>$8000</td>
<td>$A000</td>
<td>D5 and D6 swapped, switched via $4169</td>
</tr>
</tbody>
</table>

## Embedded VT369 ROM {#Embedded_VT369_ROM}

The VT369 variant of the OneBus famiclones has 4 KiB of ROM embedded into the NoaC glob itself. These 4 KiB are mapped to \$1000-\$1FFF of the main CPU's address space and to \$4000-\$4FFF of the sound CPU's address space, and contain initialization code and a sound CPU program. They are included as NES 2.0 Misc ROM.
