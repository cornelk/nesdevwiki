---
title: "FamicomBox database"
url: "/wiki/FamicomBox_database/"
description: "The FamicomBox menu cartridge contains a database of game checksums for games that do not have a Nintendo header, allowing them to be detected by the menu software as valid (assuming they have a…"
wikiPage: "FamicomBox_database"
---

The FamicomBox menu cartridge contains a database of game checksums for games that do not have a [Nintendo header](../Nintendo_header/), allowing them to be detected by the menu software as valid (assuming they have a matching 3198 CIC) and displayed in the menu with an appropriate title. This enabled some games in this list to be released without any modification to the ROM, which could be expensive and invasive.

## Format {#Format}

Each entry is 32 bytes long:

- **\$00-\$01** : Full PRG checksum. Big-endian sum of the largest region detected to contain unique data (\$E000-\$FFFF, \$C000-\$FFFF, or \$8000-\$FFFF).
- **\$02**  : Title length - 1 (1-15 = 2-16 bytes).
- **\$03-\$12** : Title of the game, in ASCII, left-justified. Normally padded with spaces (\$20).
- **\$13-\$1C** : Padding. These are always padded in valid entries with spaces (\$20).
- **\$1D-\$1E** : Upper PRG checksum. If present, big-endian sum of \$E000-\$FFFF if that is the largest region with unique data, or of \$C000-\$FFFF otherwise. This allows just the upper half of the unique data to be checked, which can work around collisions or unreliable mapper power-on state.
- **\$1F** : Upper PRG checksum present (0 = present).

Invalid entries are filled with \$FF.

## Lookup algorithm {#Lookup_algorithm}

The FamicomBox software only falls back on a database lookup if it does not find a valid [Nintendo header](../Nintendo_header/) at \$FFE0-\$FFF9 in the CPU address space. To generate a database checksum, it does the following steps:

1.  \$00 is written to the first address containing \$00 that it finds starting at \$8000; this configures some mappers, such as UNROM and GNROM, to use bank 0. If there is no \$00 value in \$8000-\$FFFF, address \$0000 is harmlessly read then written.
2.  Two checksums are produced: a "full-checksum" that includes all of the unique data and an "upper-checksum" that sums only \$E000-\$FFFF if that is the largest region with unique data or \$C000-\$FFFF otherwise.
    1.  First, it sums \$C000-\$DFFF for the full-checksum and \$E000-\$FFFF for the upper-checksum. While summing, each byte of the two regions is compared. If the regions are byte-for-byte identical, checksumming is complete and both checksums contain the sum of \$E000-\$FFFF.
    2.  Otherwise, it sums \$8000-\$BFFF for the full-checksum and \$C000-\$FFFF for the upper-checksum. If the regions are byte-for-byte identical, checksumming is complete and both checksums contain the sum of \$C000-\$FFFF.
    3.  Otherwise, it combines the full- and upper-checksums and stores the result as the full-checksum. The full-checksum contains the sum of \$8000-\$FFFF and the upper-checksum contains the sum of \$C000-\$FFFF.
3.  If the full-checksum is \$FFFF, the database search is skipped and the cartridge is marked as valid with no title, and the process is complete.
4.  The database is searched from the beginning (\$A000).
    1.  If byte \$1F of the database entry is 0, bytes \$1D-\$1E are compared against the upper-checksum. If they match, the cart is valid and we use this entry for the title. If the low bytes differ, this database entry is skipped. If only the high byte differs, or if byte \$1F isn't 0, it continues on to the full-checksum.
    2.  Bytes \$00-\$01 are compared against the full-checksum. If they match, the cart is valid and we use this entry for the title. Otherwise, we advance to the next entry.
    3.  If we advanced to \$C000, the cart is not valid.

## Entries {#Entries}

The default sorting for this table is the order they appear in the database. Said ordering appears to consist of two blocks - one block for later game revisions and another block for earlier revisions, roughly in release date order. First-party titles in particular follow the exact release order<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> - the launch titles *Donkey Kong*, *Donkey Kong Jr.*, and *Popeye* are the first in each block.

There are empty entries with checksum \$FFFF, which are noted as such. There are 82 more empty entries beyond "BABEL", the last valid entry, which are omitted from the table for brevity.

<table class="wikitable sortable">
<thead>
<tr>
<th>Title</th>
<th colspan="2">Checksum</th>
<th>Game version</th>
<th>Notes</th>
</tr>
</thead>
<tbody>
<tr>
<th>Full</th>
<th>Upper</th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>DONKEY KONG</td>
<td>$1104</td>
<td></td>
<td>Donkey Kong (World) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>DONKEY KONG JR.</td>
<td>$F97C</td>
<td></td>
<td>Donkey Kong Jr. (World) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>POPEYE</td>
<td>$2F81</td>
<td></td>
<td>Popeye (World) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>MARIO BROS.</td>
<td>$AAB1</td>
<td></td>
<td>Mario Bros. (World)</td>
<td></td>
</tr>
<tr>
<td>MAH-JONG</td>
<td>$E300</td>
<td></td>
<td>Mahjong (Japan) (Rev 1), Mahjong (Japan) (Rev 2)</td>
<td></td>
</tr>
<tr>
<td>GOMOKU-NARABE</td>
<td>$A54B</td>
<td></td>
<td>Gomoku Narabe (Japan)</td>
<td></td>
</tr>
<tr>
<td>SANSUU ASOBI</td>
<td>$02AF</td>
<td></td>
<td>Donkey Kong Jr. no Sansuu Asobi (Japan)</td>
<td></td>
</tr>
<tr>
<td>EIGO ASOBI</td>
<td>$62DB</td>
<td></td>
<td>Popeye no Eigo Asobi (Japan)</td>
<td></td>
</tr>
<tr>
<td>BASEBALL</td>
<td>$7792</td>
<td></td>
<td>Baseball (Japan)</td>
<td></td>
</tr>
<tr>
<td>PINBALL</td>
<td>$6B06</td>
<td></td>
<td>Pinball (Japan, USA)</td>
<td></td>
</tr>
<tr>
<td>TENNIS</td>
<td>$076A</td>
<td></td>
<td>Tennis (Japan, USA)</td>
<td></td>
</tr>
<tr>
<td>WILD GUNMAN</td>
<td>$1783</td>
<td></td>
<td>Wild Gunman (World) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>DUCK HUNT</td>
<td>$8261</td>
<td></td>
<td>Duck Hunt (World)</td>
<td></td>
</tr>
<tr>
<td>GOLF</td>
<td>$AF78</td>
<td></td>
<td>Golf (Japan)</td>
<td></td>
</tr>
<tr>
<td>HOGAN'S ALLEY</td>
<td>$E84A</td>
<td></td>
<td>Hogan's Alley (World)</td>
<td></td>
</tr>
<tr>
<td>DONKEY KONG 3</td>
<td>$8C5F</td>
<td></td>
<td>Donkey Kong 3 (World)</td>
<td></td>
</tr>
<tr>
<td>NUTS &amp; MILK</td>
<td>$A00E</td>
<td></td>
<td>Nuts &amp; Milk (Japan)</td>
<td></td>
</tr>
<tr>
<td>LODE RUNNER</td>
<td>$6407</td>
<td></td>
<td>Lode Runner (Japan)</td>
<td></td>
</tr>
<tr>
<td>DEVIL WORLD</td>
<td>$081A</td>
<td></td>
<td>Devil World (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>YONIN-UCHI</td>
<td>$0D81</td>
<td></td>
<td>4 Nin Uchi Mahjong (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>F1 RACE</td>
<td>$9342</td>
<td></td>
<td>F1 Race (Japan)</td>
<td></td>
</tr>
<tr>
<td>URBAN CHAMPION</td>
<td>$FB3C</td>
<td></td>
<td>Urban Champion (World)</td>
<td></td>
</tr>
<tr>
<td>CLU CLU LAND</td>
<td>$BB8A</td>
<td></td>
<td>Clu Clu Land (World)</td>
<td></td>
</tr>
<tr>
<td>EXCITEBIKE</td>
<td>$A6A0</td>
<td></td>
<td>Excitebike (Japan, USA)</td>
<td></td>
</tr>
<tr>
<td>BALOON FIGHT</td>
<td>$BDF9</td>
<td></td>
<td>Balloon Fight (Japan)</td>
<td></td>
</tr>
<tr>
<td>ICE CLIMBER</td>
<td>$E710</td>
<td></td>
<td>Ice Climber (Japan)</td>
<td></td>
</tr>
<tr>
<td>BUNGELING BAY</td>
<td>$8C9C</td>
<td></td>
<td>Raid on Bungeling Bay (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>CHAMPIONSHIP</td>
<td>$3683</td>
<td></td>
<td>Championship Lode Runner (Japan)</td>
<td></td>
</tr>
<tr>
<td>SOCCER</td>
<td>$FD07</td>
<td></td>
<td>Soccer (Japan, USA)</td>
<td></td>
</tr>
<tr>
<td>FLAPPY</td>
<td>$3809</td>
<td></td>
<td>Flappy (Japan)</td>
<td></td>
</tr>
<tr>
<td>STAR FORCE</td>
<td>$5BD4</td>
<td></td>
<td>Star Force (Japan)</td>
<td></td>
</tr>
<tr>
<td>WRECKING CREW</td>
<td>$6935</td>
<td></td>
<td>Wrecking Crew (World)</td>
<td></td>
</tr>
<tr>
<td>SPARTAN X</td>
<td>$4270</td>
<td></td>
<td>Spartan X (Japan)</td>
<td></td>
</tr>
<tr>
<td>DOOR DOOR</td>
<td>$411E</td>
<td></td>
<td>Door Door (Japan)</td>
<td></td>
</tr>
<tr>
<td>ASTRO-ROBO SASA</td>
<td>$C964</td>
<td></td>
<td>Astro Robo Sasa (Japan)</td>
<td></td>
</tr>
<tr>
<td>GEIMO</td>
<td>$A73E</td>
<td></td>
<td>Geimos (Japan)</td>
<td>The title data contains "GEIMOS", but the title length is off by 1</td>
</tr>
<tr>
<td>POOYA</td>
<td>$92F4</td>
<td></td>
<td>Pooyan (Japan)</td>
<td>The title data contains "POOYAN", but the title length is off by 1</td>
</tr>
<tr>
<td>SUPER MARIO</td>
<td>$44B9</td>
<td></td>
<td>Super Mario Bros. (World)</td>
<td></td>
</tr>
<tr>
<td>CHALLENGER</td>
<td>$AF86</td>
<td></td>
<td>Challenger (Japan)</td>
<td></td>
</tr>
<tr>
<td>PACHICON</td>
<td>$9114</td>
<td></td>
<td>Pachicom (Japan)</td>
<td></td>
</tr>
<tr>
<td>MACH RIDER</td>
<td>$13DE</td>
<td></td>
<td>Mach Rider (Japan, USA) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>PORTPIA</td>
<td>$086E</td>
<td></td>
<td>Portopia Renzoku Satsujin Jiken (Japan)</td>
<td></td>
</tr>
<tr>
<td>ONYANKO TOWN</td>
<td>$3939</td>
<td></td>
<td>Onyanko Town (Japan)</td>
<td></td>
</tr>
<tr>
<td>PENGUIN-KUN WARS</td>
<td>$832B</td>
<td></td>
<td>Penguin-kun Wars (Japan)</td>
<td></td>
</tr>
<tr>
<td>DOUGH BOY</td>
<td>$F9F9</td>
<td></td>
<td>Dough Boy (Japan)</td>
<td></td>
</tr>
<tr>
<td>LUNAR BALL</td>
<td>$5C2F</td>
<td></td>
<td>Lunar Ball (Japan)</td>
<td></td>
</tr>
<tr>
<td>EXED EXES</td>
<td>$4E76</td>
<td></td>
<td>Chou Fuyuu Yousai Exed Exes (Japan)</td>
<td></td>
</tr>
<tr>
<td>LOT LOT</td>
<td>$FB6E</td>
<td></td>
<td>Lot Lot (Japan)</td>
<td></td>
</tr>
<tr>
<td>KARATEKA</td>
<td>$B09C</td>
<td></td>
<td>Karateka (Japan)</td>
<td></td>
</tr>
<tr>
<td>1942</td>
<td>$024F</td>
<td></td>
<td>1942 (Japan, USA)</td>
<td></td>
</tr>
<tr>
<td>SON SON</td>
<td>$BDDE</td>
<td></td>
<td>Son Son (Japan)</td>
<td></td>
</tr>
<tr>
<td>VOL GUARD 2</td>
<td>$EBF4</td>
<td></td>
<td>Volguard II (Japan)</td>
<td></td>
</tr>
<tr>
<td>THEXDER</td>
<td>$B14B</td>
<td></td>
<td>Thexder (Japan)</td>
<td></td>
</tr>
<tr>
<td>ZUNO-SENKAN GARU</td>
<td>$B224</td>
<td></td>
<td>Zunou Senkan Galg (Japan)</td>
<td></td>
</tr>
<tr>
<td>BOKOSUKA WARS</td>
<td>$03FD</td>
<td></td>
<td>Bokosuka Wars (Japan)</td>
<td></td>
</tr>
<tr>
<td>BINARYLAND</td>
<td>$82A8</td>
<td></td>
<td>Binary Land (Japan)</td>
<td></td>
</tr>
<tr>
<td>HATTORI-KUN</td>
<td>$287A</td>
<td></td>
<td>Ninja Hattori-kun - Ninja wa Syugyou de Gozaru no Maki (Japan)</td>
<td></td>
</tr>
<tr>
<td>BOMBER MAN</td>
<td>$BF54</td>
<td></td>
<td>Bomberman (Japan)</td>
<td></td>
</tr>
<tr>
<td>MAGMAX</td>
<td>$58E3</td>
<td></td>
<td>Magmax (Japan)</td>
<td></td>
</tr>
<tr>
<td>CIRCUS CHARLIE</td>
<td>$178E</td>
<td></td>
<td>Circus Charlie (Japan)</td>
<td></td>
</tr>
<tr>
<td>BALTRON</td>
<td>$A8BA</td>
<td></td>
<td>Baltron (Japan)</td>
<td></td>
</tr>
<tr>
<td>HYDLIDE SPECIAL</td>
<td>$4207</td>
<td></td>
<td>Hydlide Special (Japan)</td>
<td></td>
</tr>
<tr>
<td>DRAGON QUEST</td>
<td>$34AE</td>
<td></td>
<td>Dragon Quest (Japan)</td>
<td></td>
</tr>
<tr>
<td>KEISAN-GAME 1</td>
<td>$38BC</td>
<td></td>
<td>Sansuu 1 Nen - Keisan Game (Japan)</td>
<td></td>
</tr>
<tr>
<td>KEISAN-GAME 2</td>
<td>$F339</td>
<td></td>
<td>Sansuu 2 Nen - Keisan Game (Japan)</td>
<td></td>
</tr>
<tr>
<td>KEISAN-GAME 3</td>
<td>$6D4F</td>
<td></td>
<td>Sansuu 3 Nen - Keisan Game (Japan)</td>
<td></td>
</tr>
<tr>
<td>SUPAI VS SUPAI</td>
<td>$1A81</td>
<td></td>
<td>Spy vs Spy (Japan)</td>
<td></td>
</tr>
<tr>
<td>SECROSS</td>
<td>$0AF7</td>
<td></td>
<td>Seicross (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>MIGHTY BOMB JACK</td>
<td>$C3CE</td>
<td></td>
<td>Mighty Bomb Jack (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>B-WING</td>
<td>$6673</td>
<td></td>
<td>B-Wings (Japan)</td>
<td></td>
</tr>
<tr>
<td>BIRD-WEEK</td>
<td>$84BC</td>
<td></td>
<td>Bird Week (Japan)</td>
<td></td>
</tr>
<tr>
<td>STAR SOLDIER</td>
<td>$8355</td>
<td></td>
<td>Star Soldier (Japan)</td>
<td></td>
</tr>
<tr>
<td>YIE AR KUNG-FU</td>
<td>$8C18</td>
<td></td>
<td>Yie Ar Kung-Fu (Japan) (Rev 1.2)</td>
<td></td>
</tr>
<tr>
<td>EXERION</td>
<td>$7332</td>
<td></td>
<td>Exerion (Japan)</td>
<td></td>
</tr>
<tr>
<td>ELEVATOR ACTION</td>
<td>$CDBA</td>
<td></td>
<td>Elevator Action (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>GALAGA</td>
<td>$83E4</td>
<td>$83E4</td>
<td>Galaga (Japan)</td>
<td>Specifies checksums for both $8000-$BFFF and $C000-$FFFF</td>
</tr>
<tr>
<td>ZIPPY RACE</td>
<td>$E95E</td>
<td></td>
<td>Zippy Race (Japan)</td>
<td></td>
</tr>
<tr>
<td>ANTAR. ADVENTURE</td>
<td>$7B59</td>
<td></td>
<td>Kekkyoku Nankyoku Daibouken (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>SUPER ARABIAN</td>
<td>$9F1E</td>
<td></td>
<td>Super Arabian (Japan)</td>
<td></td>
</tr>
<tr>
<td>SPACE INVADERS</td>
<td>$3001</td>
<td></td>
<td>Space Invaders (Japan)</td>
<td></td>
</tr>
<tr>
<td>XEVIOUS</td>
<td>$C6A9</td>
<td></td>
<td>Xevious (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>CHACK'N POP</td>
<td>$BC5F</td>
<td></td>
<td>Chack'n Pop (Japan)</td>
<td></td>
</tr>
<tr>
<td>DIG DUG</td>
<td>$3745</td>
<td></td>
<td>Dig Dug (Japan)</td>
<td></td>
</tr>
<tr>
<td>10-YARD FIGHT</td>
<td>$0339</td>
<td></td>
<td>10-Yard Fight (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>NINJA-KUN</td>
<td>$F6E8</td>
<td></td>
<td>Ninja-kun - Majou no Bouken (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>HYPER OLYMPIC</td>
<td>$A4A4</td>
<td></td>
<td>Hyper Olympic (Japan)</td>
<td></td>
</tr>
<tr>
<td>PAC-MAN</td>
<td>$2C69</td>
<td></td>
<td>Pac-Man (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>FIELD COMBAT</td>
<td>$F147</td>
<td></td>
<td>Field Combat (Japan)</td>
<td></td>
</tr>
<tr>
<td>FORMATION Z</td>
<td>$4656</td>
<td></td>
<td>Formation Z (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>FRONT LINE</td>
<td>$11C1</td>
<td></td>
<td>Front Line (Japan)</td>
<td></td>
</tr>
<tr>
<td>MAPPY</td>
<td>$E929</td>
<td></td>
<td>Mappy (Japan)</td>
<td></td>
</tr>
<tr>
<td>ROAD FIGHTER</td>
<td>$C7CF</td>
<td></td>
<td>Road Fighter (Japan)</td>
<td></td>
</tr>
<tr>
<td>DRUAGA</td>
<td>$83E4</td>
<td>$B88D</td>
<td>Druaga no Tou (Japan)</td>
<td>Relies on $C000-$FFFF checksum because of collision with Galaga</td>
</tr>
<tr>
<td>BATTLE CITY</td>
<td>$4218</td>
<td></td>
<td>Battle City (Japan)</td>
<td></td>
</tr>
<tr>
<td>WARPMAN</td>
<td>$08D8</td>
<td></td>
<td>Warpman (Japan)</td>
<td></td>
</tr>
<tr>
<td>HYPER SPORTS</td>
<td>$5600</td>
<td></td>
<td>Hyper Sports (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>HONSHOGI</td>
<td>$9B81</td>
<td></td>
<td>Honshougi - Naitou 9 Dan Shougi Hiden (Japan)</td>
<td></td>
</tr>
<tr>
<td>CITY CONNECTION</td>
<td>$D8EE</td>
<td></td>
<td>City Connection (Japan)</td>
<td></td>
</tr>
<tr>
<td>ROUTE-16 TURBO</td>
<td>$1876</td>
<td></td>
<td>Route-16 Turbo (Japan)</td>
<td></td>
</tr>
<tr>
<td>PAC-LAND</td>
<td>$9D5B</td>
<td></td>
<td>Pac-Land (Japan)</td>
<td></td>
</tr>
<tr>
<td>BURGER TIME</td>
<td>$9B98</td>
<td></td>
<td>BurgerTime (Japan)</td>
<td></td>
</tr>
<tr>
<td>JAJAMARU-KUN</td>
<td>$2052</td>
<td></td>
<td>Ninja Jajamaru-kun (Japan)</td>
<td></td>
</tr>
<tr>
<td>SKY DESTROYER</td>
<td>$76B0</td>
<td></td>
<td>Sky Destroyer (Japan)</td>
<td></td>
</tr>
<tr>
<td>TWIN BEE</td>
<td>$0300</td>
<td></td>
<td>TwinBee (Japan)</td>
<td></td>
</tr>
<tr>
<td>KINNIKUMAN</td>
<td>$7F7B</td>
<td></td>
<td>Kinnikuman - Muscle Tag Match (Japan) (Rev 1)</td>
<td></td>
</tr>
<tr>
<td>STAR LUSTER</td>
<td>$8742</td>
<td></td>
<td>Star Luster (Japan)</td>
<td></td>
</tr>
<tr>
<td>IKKI</td>
<td>$A2D3</td>
<td></td>
<td>Ikki (Japan)</td>
<td></td>
</tr>
<tr>
<td>OBAKE-NO Q-TARO</td>
<td>$5DD6</td>
<td></td>
<td>Obake no Q Tarou - Wanwan Panic (Japan)</td>
<td></td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td>MACROSS</td>
<td>$7BB9</td>
<td></td>
<td>Choujikuu Yousai - Macross (Japan)</td>
<td></td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td>SPELUNKER</td>
<td>$99D0</td>
<td></td>
<td>Spelunker (Japan)</td>
<td></td>
</tr>
<tr>
<td>GYRODINE</td>
<td>$1D0D</td>
<td></td>
<td>Gyrodine (Japan)</td>
<td></td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td>MAKAIMURA</td>
<td>$B7A5</td>
<td></td>
<td>Makaimura (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td>SUPER CHINESE</td>
<td>$E74A</td>
<td></td>
<td></td>
<td>This doesn't seem to match any configuration of banks for Super Chinese (Japan)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td></td>
<td>$FFFF</td>
<td></td>
<td></td>
<td>(Empty)</td>
</tr>
<tr>
<td>DONKEY KONG</td>
<td>$03EC</td>
<td></td>
<td>Donkey Kong (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>DONKEY KONG JR.</td>
<td>$E7C4</td>
<td></td>
<td>Donkey Kong Jr. (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>POPEYE</td>
<td>$2950</td>
<td></td>
<td>Popeye (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>MAH-JONG</td>
<td>$C446</td>
<td></td>
<td>Mahjong (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>WILD GUNMAN</td>
<td>$2ED6</td>
<td></td>
<td>Wild Gunman (Japan, USA) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>DEVIL WORLD</td>
<td>$07FB</td>
<td></td>
<td>Devil World (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>YONIN-UCHI</td>
<td>$11BA</td>
<td></td>
<td>4 Nin Uchi Mahjong (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>MACH RIDER</td>
<td>$0F7E</td>
<td></td>
<td>Mach Rider (Japan, USA) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>BUNGELING BEY</td>
<td>$9108</td>
<td></td>
<td>Raid on Bungeling Bay (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>CASTLE EXCELLENT</td>
<td>$BCBD</td>
<td></td>
<td>Castle Excellent (Japan)</td>
<td></td>
</tr>
<tr>
<td>MIGHTY BOMB JACK</td>
<td>$3777</td>
<td></td>
<td>Mighty Bomb Jack (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>A S</td>
<td>$8523</td>
<td></td>
<td>ASO - Armored Scrum Object (Japan)</td>
<td>The title data contains "A S O", but the title length specifies only 3 bytes</td>
</tr>
<tr>
<td>SOROMON NO KAGI</td>
<td>$1A7C</td>
<td></td>
<td>Solomon no Kagi (Japan)</td>
<td></td>
</tr>
<tr>
<td>KING' KNIGHT</td>
<td>$3652</td>
<td></td>
<td>King's Knight (Japan)</td>
<td></td>
</tr>
<tr>
<td>HOKUTO NO KEN</td>
<td>$1F3F</td>
<td></td>
<td>Hokuto no Ken (Japan)</td>
<td></td>
</tr>
<tr>
<td>TERRA CRESTA</td>
<td>$3EB7</td>
<td></td>
<td>Terra Cresta (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>BANANA</td>
<td>$801E</td>
<td></td>
<td>Banana (Japan)</td>
<td></td>
</tr>
<tr>
<td>SUPER PITFALL</td>
<td>$F7BD</td>
<td></td>
<td>Super Pitfall (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>GHOSTBUSTERS</td>
<td>$C0CA</td>
<td></td>
<td>Ghostbusters (Japan)</td>
<td></td>
</tr>
<tr>
<td>BOUKENJIMA</td>
<td>$7D7E</td>
<td></td>
<td>Takahashi Meijin no Bouken-jima (Japan)</td>
<td></td>
</tr>
<tr>
<td>SENJYO NO OHKAMI</td>
<td>$2B2F</td>
<td></td>
<td>Senjou no Ookami (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>SPACE HUNTER</td>
<td>$2B9C</td>
<td></td>
<td>Space Hunter (Japan)</td>
<td></td>
</tr>
<tr>
<td>BUGGY POPPER</td>
<td>$934B</td>
<td></td>
<td>Buggy Popper (Japan)</td>
<td></td>
</tr>
<tr>
<td>OTHELLO</td>
<td>$BC8D</td>
<td></td>
<td>Othello (Japan) (Rev 0)</td>
<td></td>
</tr>
<tr>
<td>SUPER STAR FORCE</td>
<td>$E152</td>
<td></td>
<td>Super Star Force (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>TRANSFORMER</td>
<td>$78E6</td>
<td></td>
<td>Tatakae! Chou Robot Seimeitai Transformers - Convoy no Nazo (Japan)</td>
<td></td>
</tr>
<tr>
<td>KEISAN GAME 4</td>
<td>$28D3</td>
<td></td>
<td>Sansuu 4 Nen - Keisan Game (Japan)</td>
<td></td>
</tr>
<tr>
<td>KEISAN GAME 5-6</td>
<td>$9F2B</td>
<td></td>
<td>Sansuu 5 &amp; 6 Nen - Keisan Game (Japan)</td>
<td></td>
</tr>
<tr>
<td>SUPER MONKEY</td>
<td>$9458</td>
<td></td>
<td>Ganso Saiyuuki - Super Monkey Daibouken (Japan)</td>
<td></td>
</tr>
<tr>
<td>DAIVA</td>
<td>$CF6B</td>
<td></td>
<td>Daiva - Imperial of Nirsartia (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>IKARI</td>
<td>$8C23</td>
<td></td>
<td>Ikari (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>CAT IN BOOTS</td>
<td>$1C9B</td>
<td></td>
<td>Nagagutsu o Haita Neko - Sekai Isshuu 80 Nichi Daibouken (Japan)</td>
<td></td>
</tr>
<tr>
<td>MEIKYU-KUMIKYOKU</td>
<td>$B1C0</td>
<td></td>
<td>Meikyuu Kumikyoku - Milon no Daibouken (Japan)</td>
<td></td>
</tr>
<tr>
<td>AIGINA</td>
<td>$70DC</td>
<td></td>
<td>Aigina no Yogen - Balubalouk no Densetsu Yori (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>TIGER HELI.</td>
<td>$3E3F</td>
<td></td>
<td>Tiger-Heli (Japan)</td>
<td></td>
</tr>
<tr>
<td>HOTTERMAN</td>
<td>$349E</td>
<td></td>
<td>Hottarman no Chitei Tanken (Japan)</td>
<td></td>
</tr>
<tr>
<td>CRAZY CLIMBER</td>
<td>$F13F</td>
<td></td>
<td>Crazy Climber (Japan)</td>
<td>Mapper 180 - Checksum is bank 0</td>
</tr>
<tr>
<td>TOKI NO TABIBITO</td>
<td>$8E14</td>
<td></td>
<td>Toki no Tabibito - Time Stranger (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>SHERLOCK HOLMES</td>
<td>$A4CE</td>
<td></td>
<td>Sherlock Holmes - Hakushaku Reijou Yuukai Jiken (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>SEIKIMATSU</td>
<td>$0161</td>
<td></td>
<td>Seikima II - Akuma no Gyakushuu! (Japan)</td>
<td></td>
</tr>
<tr>
<td>DORAEMON</td>
<td>$952D</td>
<td></td>
<td>Doraemon (Japan)</td>
<td>GNROM - Checksum is bank 0</td>
</tr>
<tr>
<td>TATAKAI NO BANKA</td>
<td>$4F73</td>
<td></td>
<td>Tatakai no Banka (Japan) (Rev 0)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>LAYLA</td>
<td>$0CEB</td>
<td></td>
<td>Layla (Japan)</td>
<td>UNROM - Checksum is bank 0 + bank 7</td>
</tr>
<tr>
<td>SKYKID</td>
<td>$3707</td>
<td></td>
<td>Sky Kid (Japan)</td>
<td>Mapper 206 (no banking) - Checksum is bank 0 + bank 1 + bank 2 + bank 3</td>
</tr>
<tr>
<td>SUPER CHINESE</td>
<td>$0B16</td>
<td></td>
<td>Super Chinese (Japan)</td>
<td>Mapper 206 (no banking) - Checksum is bank 0 + bank 1 + bank 2 + bank 3</td>
</tr>
<tr>
<td>BABEL</td>
<td>$F4C8</td>
<td></td>
<td>Babel no Tou (Japan)</td>
<td>Mapper 206 - Checksum is bank 0 + bank 0 + bank 2 + bank 3</td>
</tr>
</tbody>
</table>

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) <https://www.nintendo.com/jp/famicom/software/index.html>
