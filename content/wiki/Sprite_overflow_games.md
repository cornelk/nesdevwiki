---
title: "Sprite overflow games"
url: "/wiki/Sprite_overflow_games/"
description: "The following is a list of games which rely on putting more than 8 sprites on a scanline."
wikiPage: "Sprite_overflow_games"
---

The following is a list of games which rely on putting more than 8 sprites on a scanline.

## Use of sprite overflow flag {#Use_of_sprite_overflow_flag}

The sprite overflow flag is rarely used, mainly due to bugs when exactly 8 sprites are present on a scanline. No games rely on the buggy behavior. See [sprite overflow bug](../PPU_sprite_evaluation/#Sprite_overflow_bug) for more details.

Nonetheless, games can intentionally place 9 or more sprites in a scanline to trigger the overflow flag consistently, as long as no previous scanlines have exactly 8 sprites.

### Commercial {#Commercial}

- *Bee 52*: At the title screen, the game splits the screen with sprite overflow (at scanline 165), then splits the screen with a sprite 0 hit (at scanline 207). If sprite overflow is not emulated, the game will crash at a solid blue-purple screen.

### Homebrew {#Homebrew}

- blargg's [sprite overflow test ROMs](https://github.com/christopherpow/nes-test-roms/tree/master/sprite_overflow_tests): tests behavior of sprite overflow, including the buggy behavior.
- *City Trouble* uses both the sprite overflow flag and the sprite 0 flag to make two scroll splits.

## Use of excess sprites for masking effects {#Use_of_excess_sprites_for_masking_effects}

Some games intentionally place multiple blank sprites early in the [OAM](../PPU_OAM/) at the same Y position so that other sprites on those scanlines are hidden.

### Commercial {#Commercial_2}

- *Castlevania II: Simon's Quest* (a.k.a. *Dracula 2*): When Simon enters a swamp, the lower half of his body should be hidden.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>
- *Felix the Cat*: When entering or exiting a bag.
- *Gimmick!*: When entering a level. Also used to keep extra sprites out of the status bar.
- *Gremlins 2 - The New Batch*: Uses multiple blank sprites to mask rows during cutscenes.
- *Majou Densetsu II: Daimashikyou Galious*: When entering a doorway, Popolon's body should gradually disappear (to imitate walking down stairs).<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>
- *Ninja Gaiden 1, 2 and 3*: All sprites in all cutscenes should be confined inside the black background borders.
- *The Legend of Zelda* (a.k.a. *Zelda 1*, *Zeruda no Densetsu*): On the top or bottom of dungeon screens.

### Homebrew {#Homebrew_2}

- *Lizard* hides sprites overlapping the pause overlay while the game is paused.

### Detecting masking effects {#Detecting_masking_effects}

Games will place 8 consecutive sprites with the same Y coordinate and same tile number. If you see this, then that is a sign that the game is using a masking effect, and the 8-sprite limit should be enforced for that area.

## Misc {#Misc}

### Commercial {#Commercial_3}

- *Solstice*: during the intro cutscene, there are stray sprites on the screen beyond the 8 per scanline, but the NES won't display the excess sprites. This is not a masking effect, it is merely the hardware covering up a mistake that wasn't caught by the original programmers.

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [BBS topic](http://forums.nesdev.org/viewtopic.php?p=136304#p136304) with screenshots of *Castlevania II: Simon's Quest*.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Youtube video](https://www.youtube.com/watch?v=9XHXYhb_gDo) demonstrating the effect on *Majou Densetsu II: Daimashikyou Galious*.
