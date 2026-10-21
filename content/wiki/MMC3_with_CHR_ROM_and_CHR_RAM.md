---
title: "MMC3 with CHR ROM and CHR RAM"
url: "/wiki/MMC3_with_CHR_ROM_and_CHR_RAM/"
description: "The following mappers used an uncommon combination: MMC3 with CHR ROM and CHR RAM."
wikiPage: "MMC3_with_CHR_ROM_and_CHR_RAM"
wikiCategories:
  - name: "MMC3 with CHR ROM and CHR RAM"
    url: "/category/MMC3_with_CHR_ROM_and_CHR_RAM/"
---

The following mappers used an uncommon combination: **MMC3 with CHR ROM and CHR RAM**.

Apart from *Pinbot* and *High Speed*, which used the Nintendo-made [TQROM](../INES_Mapper_119/) board, this combination is primarily used for games in the Chinese language using an [MMC3](../MMC3/)-clone mapper, which need to mix [logographic characters](https://en.wikipedia.org/wiki/Chinese_characters) arbitrarily. No mapper other than [MMC5](../MMC5/) has CHR-ROM banks fine-grained enough to allow showing dialogue using Chinese characters with only CHR ROM without prerendering each sentence, and unlike Japanese, Chinese cannot use the common shortcut of using only [phonetic characters](https://en.wikipedia.org/wiki/kana). These games are thought to follow a similar strategy to *Faxanadu* and *Super Bat Puncher*, reserving some CHR RAM for software-rendered text boxes.

## List of mappers {#List_of_mappers}

All MMC3-like iNES mappers known to be used with CHR RAM are listed below:

<table class="tabular">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Mapper</th>
<th>First RAM bank</th>
<th>Size of RAM<br />
(1024-byte banks)</th>
<th>RAM mirrored?</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="../MMC3/">004</a></td>
<td>0</td>
<td>0 or 8*</td>
<td>Yes</td>
</tr>
<tr>
<td><a href="../INES_Mapper_074/">074</a></td>
<td>8</td>
<td>2</td>
<td>No</td>
</tr>
<tr>
<td><a href="../INES_Mapper_119/">119</a></td>
<td>64, 192</td>
<td>8</td>
<td>Yes</td>
</tr>
<tr>
<td><a href="../INES_Mapper_176/">176</a></td>
<td>0</td>
<td>8</td>
<td>No</td>
</tr>
<tr>
<td><a href="../INES_Mapper_191/">191</a></td>
<td>128</td>
<td>2</td>
<td>Yes</td>
</tr>
<tr>
<td><a href="../INES_Mapper_192/">192</a></td>
<td>8</td>
<td>4</td>
<td>No</td>
</tr>
<tr>
<td><a href="../INES_Mapper_194/">194</a></td>
<td>0</td>
<td>2</td>
<td>No</td>
</tr>
<tr>
<td><a href="../INES_Mapper_195/">195</a></td>
<td>0</td>
<td>4</td>
<td>No</td>
</tr>
</tbody>
</table>

\* In iNES format, mapper 4 has CHR RAM only if CHR ROM is not present.

## Behaviour {#Behaviour}

Because the default size is ignored, these mappers collapse into the following distinct behaviors. Some of these mappers differ only in the amount of CHR RAM that replaces CHR ROM. ROMs with the [NES 2.0](../NES_2.0/) header specify the size of CHR RAM separately.

004, 194, 195  
CHR RAM starts at bank 0 and is mirrored only if there is no CHR ROM

074, 192  
CHR RAM starts at bank 8 and is not mirrored

119  
CHR A16 switches between CHR ROM and CHR RAM

176  
\$A001 bit 2 switches between CHR ROM and CHR RAM in the first 8 KiB

191  
CHR A17 switches between CHR ROM and CHR RAM
