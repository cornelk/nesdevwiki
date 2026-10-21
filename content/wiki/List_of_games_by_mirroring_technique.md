---
title: "List of games by mirroring technique"
url: "/wiki/List_of_games_by_mirroring_technique/"
description: "This article categorizes games by their rendering and scrolling methods, and how that relates to mirroring."
wikiPage: "List_of_games_by_mirroring_technique"
---

This article categorizes games by their rendering and [scrolling](../PPU_scrolling/) methods, and how that relates to [mirroring](../Mirroring/).

<table class="wikitable" style="width:100%;">
<colgroup>
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
<col style="width: 14%" />
</colgroup>
<thead>
<tr>
<th>Scrolling Type</th>
<th>Mirroring – With Status Bar</th>
<th>Mirroring – With Status Bar and Parallax Scrolling Effect</th>
<th>Mirroring – With Parallax Scrolling Effect</th>
<th>Mirroring – Without Status Bar and Parallax Scrolling Effect</th>
<th>Games example</th>
<th>Comment</th>
</tr>
</thead>
<tbody>
<tr>
<td>1-screen fixed Only</td>
<td>Any</td>
<td>Vertical, Horizontal<sup>2</sup></td>
<td>Vertical, Horizontal<sup>2</sup></td>
<td>Any</td>
<td><em>Donkey Kong</em></td>
<td></td>
</tr>
<tr>
<td>Horizontal Only</td>
<td>Vertical</td>
<td>Vertical, Horizontal<sup>2</sup></td>
<td>Vertical, Horizontal<sup>2</sup></td>
<td>Vertical</td>
<td><em>Castlevania</em>, <em>Cheetahmen II</em>, <em>Gimmick!</em>, <em>Gradius</em>, <em>Super Mario Bros.</em></td>
<td></td>
</tr>
<tr>
<td>Vertical Only</td>
<td>Vertical, Horizontal, Single Screen, 3-screen V/H<sup>3</sup></td>
<td>Vertical, Horizontal<sup>2</sup>, Single Screen<sup>2</sup>, 3-screen V/D<sup>3</sup></td>
<td>Horizontal<sup>2</sup>, Vertical<sup>1</sup>, Diagonal, 3-screen V/D<sup>3</sup>, 4-screen<sup>4</sup></td>
<td>Horizontal</td>
<td><em>Gun.Smoke</em>, <em>Ice Climber</em>, <em>Recca</em></td>
<td>Horizontal mirroring with a status bar works best with as status bar will have to change address when scrolling vertically.
<p>If you are using a parallax scrolling effect on the entire screen with a status bar, such as the <em>Recca</em>, it is best to use a Single Screen, Horizontal, or the Vertical mirroring with a scanline counter IRQ.</p></td>
</tr>
<tr>
<td>Horizontal Only, field limited horizontally</td>
<td>Vertical</td>
<td>Vertical</td>
<td>Vertical</td>
<td>Vertical</td>
<td><em>Bomberman</em>, <em>Lode Runner</em></td>
<td>No data has to be loaded at all when scrolling horizontally, but the area is limited to two screens.
<p>This screen type, some of parallax scrolling effect does not work.</p></td>
</tr>
<tr>
<td>Vertical Only, field limited vertically</td>
<td>Horizontal</td>
<td>Horizontal</td>
<td>Horizontal</td>
<td>Horizontal</td>
<td><em>Ninja Kun</em>, <em>Wrecking Crew</em></td>
<td>No data has to be loaded at all when scrolling vertically, but the area is limited to two screens (less the status bar if present).</td>
</tr>
<tr>
<td>Horizontal/Vertical aligned screens</td>
<td>Vertical, Horizontal<sup>2</sup>, Diagonal, 3-screen V/D<sup>3</sup>, 4-screen<sup>4</sup></td>
<td>Vertical, Horizontal<sup>2</sup>, Single Screen<sup>2</sup>, 3-screen V/D<sup>3</sup>, 4-screen<sup>4</sup></td>
<td>Vertical, Horizontal<sup>2</sup>, Diagonal, 3-screen V/D<sup>3</sup>, 4-screen<sup>4</sup></td>
<td>Alternate H/V, L-shaped, Diagonal, 3-screen D<sup>3</sup>, 4-screen<sup>4</sup></td>
<td><em>Air Fortress</em>, <em>Duck Tales</em>, <em>Jackie Chan</em>, <em>Mega Man</em> series, <em>Metroid</em>, <em>TMNT</em></td>
<td>For use if scrolling direction changes only on a screen-based pattern, as in <em>Metroid</em> and <em>Mega Man</em> 3, 4, 5, 6.
<p>Horizontal mirroring with a status bar works best with as status bar will have to change address when scrolling vertically (Also work as Diagonal mirroring) or playfield will have to be written twice to memory. (<em>Panic Restaurant</em>, <em>TMNT</em>)</p>
<p>If you are using a parallax scrolling effect on the entire screen with a status bar, such as the <em>Recca</em>, it is best to use a Single Screen, Horizontal, or the Vertical mirroring with a scanline counter IRQ.</p></td>
</tr>
<tr>
<td>1-screen switching by the Horizontal/Vertical</td>
<td>Alternate H/V, Vertical, Diagonal</td>
<td>Alternate H/V, Vertical, Diagonal</td>
<td>Alternate H/V, Diagonal</td>
<td>Alternate H/V, Diagonal</td>
<td><em>Guardian Legend</em>, <em>Hydlide 3</em>, <em>Legend of Zelda</em></td>
<td>Depending on the direction of the pattern to scroll in a single screen, it is the best choice for switching the mirroring by the mapper.</td>
</tr>
<tr>
<td>Horizontal/Vertical/Bidirectional aligned screens</td>
<td>Vertical &lt;-&gt; Single Screen<sup>2</sup>, Horizontal<sup>2</sup>, Vertical, 4-screen<sup>4</sup></td>
<td>Vertical &lt;-&gt; Single Screen<sup>2</sup>, Single Screen<sup>2</sup>, Horizontal<sup>2</sup>, Vertical, 4-screen<sup>4</sup></td>
<td>Vertical<sup>1</sup>, Horizontal<sup>2</sup>, Diagonal, 4-screen<sup>4</sup></td>
<td>Vertical<sup>1</sup>, Horizontal<sup>2</sup>, Diagonal, 4-screen<sup>4</sup></td>
<td><em>Akumajou Special</em>, <em>Rockman 4 Minus Infinity</em>, <em>TMNT</em> II/III</td>
<td>Horizontal or 4-screen mirroring with a status bar works best with as status bar will have to change address when scrolling vertically or playfield will have to be written twice to memory. (<em>TMNT III: The Manhattan Project</em>)
<p>Vertical mirroring Only with a status bar works best with a scanline counter IRQ. (<em>TMNT II: The Arcade Game</em>)</p>
<p>If you are using a parallax scrolling effect on the entire screen with a status bar, such as the <em>Recca</em>, it is best to use a Single Screen, Horizontal, or the Vertical mirroring with a scanline counter IRQ.</p></td>
</tr>
<tr>
<td>Bidirectional, field limited vertically</td>
<td>Horizontal<sup>2</sup></td>
<td>Horizontal<sup>2</sup></td>
<td>Horizontal<sup>2</sup></td>
<td>Horizontal<sup>2</sup></td>
<td><em>Super Mario Bros. 3</em>, <em>Tiny Toon Adventures</em></td>
<td>No data has to be loaded at all when scrolling vertically, but the area is limited to two screens (less the status bar if present).</td>
</tr>
<tr>
<td>Bidirectional, field limited horizontally</td>
<td>Vertical</td>
<td>Vertical</td>
<td>Vertical<sup>1</sup></td>
<td>Vertical<sup>1</sup></td>
<td><em>Fire Emblem</em></td>
<td>No data has to be loaded at all when scrolling horizontally, but the area is limited to two screens.
<p>Vertical mirroring with a status bar works best with a scanline counter IRQ.</p>
<p>This screen type, some of parallax scrolling effect does not work.</p></td>
</tr>
<tr>
<td>Bidirectional, field limited 4-screen</td>
<td>4-screen<sup>4</sup></td>
<td>4-screen<sup>4</sup></td>
<td>4-screen<sup>4</sup></td>
<td>4-screen<sup>4</sup></td>
<td><em>Gauntlet</em></td>
<td>No data has to be loaded at all when scrolling 4-screen, but the area is limited to four screens (less the status bar if present).</td>
</tr>
<tr>
<td>Free bidirectional</td>
<td>Single Screen<sup>2</sup>, Horizontal<sup>2</sup>, Vertical, 4-screen<sup>4</sup></td>
<td>Single Screen<sup>2</sup>, Horizontal<sup>2</sup>, Vertical, 4-screen<sup>4</sup></td>
<td>Vertical<sup>1</sup>, Horizontal<sup>2</sup>, Diagonal, 4-screen<sup>4</sup></td>
<td>Vertical<sup>1</sup>, Horizontal<sup>2</sup>, Diagonal, 4-screen<sup>4</sup></td>
<td><em>Battletoads</em>, <em>Castlevania II</em>, <em>Hebereke</em>, <em>Kirby's Adventure</em>, countless RPG games</td>
<td>If using horizontal or 4-screen mirroring with a status bar there are 2 possibilities:
<ol>
<li>The status bar will have to change address when scrolling vertically (<em>Conquest of Crystal Palace</em>, <em>Double Dragon</em> series, <em>Gradius II</em>)</li>
<li>The playfield will have to be written twice to memory (<em>Incredible Crash Dummmies</em>, <em>Kirby's Adventure</em>, <em>Little Nemo: The Dream Master</em>), which only works if the status's bar size remains constant.</li>
</ol>
<p>Vertical mirroring with a status bar works best with a scanline counter IRQ. (<em>Jungle Book</em>, <em>Mickey's Safari in Letterland</em>)</p></td>
</tr>
<tr>
<td>Depth screen</td>
<td>Vertical, Horizontal, 4-screen<sup>4</sup></td>
<td>Vertical, Horizontal, 4-screen<sup>4</sup></td>
<td>Vertical, Horizontal, 4-screen<sup>4</sup></td>
<td>Vertical, Horizontal, 4-screen<sup>4</sup></td>
<td><em>After Burner</em>, <em>F-1 Race</em>, <em>Mach Rider</em>, <em>Rad Racer</em></td>
<td>Bidirectional Pseudo 3D depth perspective works best with a CHR bank switching ROM nametables. (<em>After Burner</em>)</td>
</tr>
</tbody>
</table>

- <sup>1</sup>: Vertical glitches on PAL screens and on NTSC that doesn't overscan.
- <sup>2</sup>: Horizontal glitches will be unavoidable when scrolling (see above).
- <sup>3</sup>: 3-screen mirroring, you can use only specific mapper of MMC5.
- <sup>4</sup>: 4-screen mirroring, and can be used only when you added an additional RAM to the cartridge board.
