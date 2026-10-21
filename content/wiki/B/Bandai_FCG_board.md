---
title: "Bandai FCG board"
url: "/wiki/Bandai_FCG_board/"
description: "Bandai FCG boards are used largely by Bandai for Dragon Ball and Gundam games, as well as a few others. One Irem title uses it as well."
wikiPage: "Bandai_FCG_board"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

Bandai FCG boards are used largely by Bandai for *Dragon Ball* and *Gundam* games, as well as a few others. One Irem title uses it as well.

All of these games were originally assigned to [INES Mapper 016](../INES_Mapper_016/) before the subtle differences, that are nonetheless relevant for fully functional emulation, became known. Please refer to the [INES Mapper 016](../INES_Mapper_016/) page for a general description of all registers, and to the [INES Mapper 153](../INES_Mapper_153/), [INES Mapper 157](../INES_Mapper_157/) and [INES Mapper 159](../INES_Mapper_159/) pages for information on the differences of those particular mappers compared to mapper 16.

- [INES Mapper 016](../INES_Mapper_016/) submapper 4: FCG-1/2 ASIC, no serial EEPROM, banked CHR-ROM
- [INES Mapper 016](../INES_Mapper_016/) submapper 5: LZ93D50 ASIC and no or 256-byte serial EEPROM, banked CHR-ROM
- [INES Mapper 159](../INES_Mapper_159/): LZ93D50 ASIC and 128-byte serial EEPROM, banked CHR-ROM
- [INES Mapper 153](../INES_Mapper_153/): LZ93D50 ASIC and 8 KiB battery-backed WRAM, unbanked CHR-RAM
- [INES Mapper 157](../INES_Mapper_157/) (Datach Joint ROM System): LZ93D50 ASIC and 256-byte serial EEPROM on Datach Main Unit, optional additional 128-byte serial EEPROM on game cartridge, unbanked CHR-RAM, barcode reader

| Name | ASIC | CHR Memory | Save Data | Mapper.Submapper | [NES 2.0 Byte 10](../NES_2.0/#Byte_10_.28RAM_size.29) |
|----|----|----|----|----|----|
| *Akuma-kun: Makai no Wana* | FCG-2 | CHR-ROM | \- | 16.4 | \$00 |
| *Crayon Shin-chan: Ora to Poi Poi* | LZ93D50 | CHR-ROM | \- | 16.5 | \$00 |
| *Dragon Ball: Daimaou Fukkatsu* | FCG-1 | CHR-ROM | \- | 16.4 | \$00 |
| *Dragon Ball 3: Gokuu Den* | FCG-2 | CHR-ROM | \- | 16.4 | \$00 |
| *Dragon Ball Z II: Gekishin Freezer!!* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *Dragon Ball Z III: Ressen Jinzou Ningen* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *Dragon Ball Z Gaiden: Saiya-jin Zetsumetsu Keikaku* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *Famicom Jump: Hero Retsuden* | FCG-2 | CHR-ROM | \- | 16.4 | \$00 |
| *Meimon! Dai-3 Yakyuu-bu* | FCG-1 | CHR-ROM | \- | 16.4 | \$00 |
| *Nishimura Kyoutarou Mystery: Blue Train Satsujin Jiken* | FCG-1 | CHR-ROM | \- | 16.4 | \$00 |
| *Rokudenashi Blues* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *Sakigake!! Otoko Juku: Shippu 1-gou Sei* | FCG-1 | CHR-ROM | \- | 16.4 | \$00 |
| *SD Gundam Gaiden - Knight Gundam Monogatari 2: Hikari no Kishi* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *SD Gundam Gaiden - Knight Gundam Monogatari 3: Densetsu no Kishidan* | LZ93D50 | CHR-ROM | 24C02 | 16.5 | \$20 |
| *Dragon Ball Z: Kyoushuu! Saiya-jin* | LZ93D50 | CHR-ROM | 24C01 | 159 | \$10 |
| *Magical Taruruuto-kun: Fantastic World!!* | LZ93D50 | CHR-ROM | 24C01 | 159 | \$10 |
| *Magical Taruruuto-kun 2: Mahou Daibouken* | LZ93D50 | CHR-ROM | 24C01 | 159 | \$10 |
| *SD Gundam Gaiden - Knight Gundam Monogatari* | LZ93D50 | CHR-ROM | 24C01 | 159 | \$10 |
| *Famicom Jump II: Saikyou no 7-nin* | LZ93D50 | CHR-RAM | 8 KiB WRAM | 153 | \$70 |
| *Datach Crayon Shin-chan: Ora to Poi Poi* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *Dragon Ball Z: Gekitou Tenkaichi Budoukai* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *J-League Super Top Players* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *SD Gundam Wars* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *Ultraman Club: Spokon Fight!!* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *Yuu Yuu Hakusho - Bakutou Ankoku Bujutsu-kai* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02 | 157 | \$00 |
| *Battle Rush: Build up Robot Tournament* | LZ93D50 | CHR-RAM | Datach Main Unit's 24C02+24C01 on cartridge | 157 | \$10 |
