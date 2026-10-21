---
title: "Palette change mid frame"
url: "/wiki/Palette_change_mid_frame/"
description: "Changing palette entries in the middle of a frame on the NES is difficult, but marginally possible."
wikiPage: "Palette_change_mid_frame"
---

Changing palette entries in the middle of a frame on the NES is difficult, but marginally possible.

A few games are known to change palette entries for a status bar region:

- *Day Dreamin' Davey*
- *Fantastic Adventures of Dizzy*
- *Startropics*
- *Wizards & Warriors*
- *Wizards & Warriors 3*

*Indiana Jones and the Last Crusade* changes its background color many times across the title screen for a gradient sky effect:

[<img src="../../wiki-images/Indiana_Jones_and_the_Last_Crusade_palette_changes.png" width="341" height="262" />](../../wiki-images/Indiana_Jones_and_the_Last_Crusade_palette_changes.png)

Other games using mid frame palette changes:

- *Ivan Ironman Stewart's Super Off-Road* - on the "Speed Shop" screen, swaps colors for the 4 panels below.

## Details {#Details}

To write to the palette in the middle of the screen, the following sequence of operations need to be done:

1.  Disable rendering (\$2001)
2.  Set PPU address (\$2006)
3.  Write new palette (\$2007)
4.  Restore PPU address (\$2006)
5.  Re-enable rendering (\$2001)

Ideally, these should be done within or mostly within horizontal blanking to minimize the appearance of visual errors. There are several problems:

- NTSC horizontal blank only lasts 21 cycles, which is not enough time to fit all of these steps at once.
  - Half of the address may be written to \$2006 early.
  - Three values can be loaded into the A, X, Y registers in advance of blanking to avoid extra load cycles.
  - Most IRQ or \$2002 polling techniques for timing are subject to jitter, varying by about 7 cycles, which cuts into the usable horizontal blank time significantly.
- Because sprites are evaluated one line in advance, turning off rendering will cause the first line of sprites to be corrupted after re-enabling.
  - The write to disable rendering (\$2001) should be done after pixel 240 on a scanline to avoid corrupting all sprites on the next frame (see: [Errata](../Errata/#OAM_and_Sprites)).
  - After the change, one line with background rendering enabled but sprites hidden can be used to flush the invalid sprite evaluation buffer. (With background rendering on, sprite evaluation will resume but remain hidden.)
- While rendering is turned off, and the PPU address is pointing at a palette entry at the same time, this will become the current background color,
  - If the background color is unchanged and stored redundantly at PPU \$3F00/4/8/C... BIT \$2007 can be used to skip over the background color, followed by 3 writes to \$2007, returning the PPU address to point at the background color again. This can be used to change 3 colors in a single horizontal blank, and if rendering is kept off 3 more colors may be updated in each subsequent line. An extra empty line is needed after the palette writes are finished, since there is not enough time in horizontal blank to write 3 colors *and* restore the PPU address and return to rendering.

Because of these issues, it is best not to use this techniques where sprites may overlap the affected lines. *Indiana Jones* hides some potential visual errors by only working where the background is already blank, and only changing the background color. Games that change colors for a status bar can use blank lines at the top of the status bar to similarly avoid visual errors.

**PPUADDR Pointer technique:** While rendering is turned off, it may be marginally useful to point the PPUADDR to one of the 16 palette entries to have the background rendered as that colour. This can serve to create lines of colours not otherwise available; or available otherwise only available to the sprite plane. Once done, you should restore the scrolling position and nametable selection before turning rendering on again. In practice, you prepare the palettes as per normal during vblank and then have the PPU look up any palette entry as if it was the backdrop colour; mid-frame.

## See Also {#See_Also}

- [PPU palettes](../PPU_palettes/)
- [PPU scrolling](../PPU_scrolling/)
- [PPU rendering](../PPU_rendering/)
- [PPU sprite evaluation](../PPU_sprite_evaluation/)
- [Consistent frame synchronization](../Consistent_frame_synchronization/)
- [Errata](../Errata/)

## References {#References}

- [Re: The quest for mid-screen palette changes!](https://forums.nesdev.org/viewtopic.php?p=209711#p209711) - useful summary post, discussion thread
- [Re: would screen splitting give extra 13 colours for tiles?](https://forums.nesdev.org/viewtopic.php?p=139925#p139925) - annotated analysis of *Indiana Jones*
- [Palette swap within the game region?](https://forums.nesdev.org/viewtopic.php?t=16679) - discussion thread
- [Hblank - Palette swap mid frame, etc....](https://forums.nesdev.org/viewtopic.php?t=16299) - discussion thread
- [Update of my Window demo with mid-frame palette writes](https://forums.nesdev.org/viewtopic.php?t=4448) - ROM/src demonstration of background color changes
- [Another palette test](https://forums.nesdev.org/viewtopic.php?t=13264) - ROM/src displays the full NES palette on a single screen, updating the colors in horizontal blank
