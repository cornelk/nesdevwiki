---
title: "INES Mapper 016/Submapper table"
url: "/wiki/INES_Mapper_016/Submapper_table/"
description: "If there are 256 bytes of serial EEPROM, they will be denoted as PRG-NVRAM in the NES 2.0 header (byte value $20). See: INES Mapper 016"
wikiPage: "INES_Mapper_016_Submapper_table"
---

<table class="wikitable">
<thead>
<tr>
<th colspan="3">INES Mapper 016 submapper table</th>
</tr>
</thead>
<tbody>
<tr>
<th>Submapper #</th>
<th>Meaning</th>
<th>Note</th>
</tr>
&#10;<tr>
<td>0</td>
<td>Unspecified</td>
<td>Emulate both FCG-1/2 and LZ93D50 chips in their respective CPU address ranges.</td>
</tr>
<tr>
<td><em>1</em></td>
<td><em>LZ93D50 with 128 byte serial EEPROM (24C01)</em></td>
<td><em>Deprecated, use <a href="../../INES_Mapper_159/">INES Mapper 159</a> instead.</em></td>
</tr>
<tr>
<td><em>2</em></td>
<td><em>Datach Joint ROM System</em></td>
<td><em>Deprecated, use <a href="../../INES_Mapper_157/">INES Mapper 157</a> instead.</em></td>
</tr>
<tr>
<td><em>3</em></td>
<td><em>8 KiB of WRAM instead of serial EEPROM</em></td>
<td><em>Deprecated, use <a href="../../INES_Mapper_153/">INES Mapper 153</a> instead.</em></td>
</tr>
<tr>
<td>4</td>
<td>FCG-1/2</td>
<td>Responds only in the CPU $6000-$7FFF address range; IRQ counter is not latched.</td>
</tr>
<tr>
<td>5</td>
<td>LZ93D50 with no or 256-byte serial EEPROM (24C02)</td>
<td>Responds only in the CPU $8000-$FFFF address range; IRQ counter is latched.</td>
</tr>
</tbody>
</table>

If there are 256 bytes of serial EEPROM, they will be denoted as PRG-NVRAM in the [NES 2.0](../../NES_2.0/) header (byte value \$20). See: [INES Mapper 016](../)
