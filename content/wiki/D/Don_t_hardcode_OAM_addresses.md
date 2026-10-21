---
title: "Don't hardcode OAM addresses"
url: "/wiki/Don't_hardcode_OAM_addresses/"
description: "To display sprites (or moving objects), a program builds a display list in a 256-byte page of CPU RAM, sometimes called \"shadow OAM\". Then the program copies it to OAM in the PPU during the next…"
wikiPage: "Don_t_hardcode_OAM_addresses"
---

To display sprites (or moving objects), a program builds a display list in a 256-byte page of CPU RAM, sometimes called "shadow OAM". Then the program copies it to OAM in the PPU during the next [vertical blanking period](../The_frame_and_NMIs/) by writing the high byte of the display list's address to [\$4014](../PPU_registers/#Sprite_0). For example, if it places the display list at \$0200-\$02FF, it writes \$02 to \$4014.

Occasionally, programmers get the idea to define variables so as to hardcode the position of a particular actor's sprites in the display list, using code similar to the following:

    OAM = $0200  ; start of shadow OAM page

    player_oam_base = OAM + $04      ; skip sprite 0
    player_y = player_oam_base + $00
    player_tile = player_oam_base + $01
    player_palette = player_oam_base + $02
    player_x = player_oam_base + $03

In all but the simplest cases, hardcoding OAM addresses of actors is an anti-pattern because it limits the ability to do several things:

- Separate the motion of the camera from motion of game objects
- Change the size of a cel later on, in case some cels need more sprites than others
- Draw a cel half-offscreen
- Reuse movement logic among multiple actor types
- Turn objectionable dropout when more than 8 sprites are briefly on a line into less objectionable flicker by varying from frame to frame which sprite appears earlier in OAM

(A "cel", also called a "metasprite", is a group of sprites that represent one frame of a character's animation.)

A better way:

1.  Keep separate variables for each actor's position in world space that are stored outside of the display list. Also keep variables for the camera's position.
2.  Each frame, initialize a variable denoting how many bytes of OAM have been filled so far. This would usually be 0 or 4; it would be 4 if sprite 0 is reserved for [sprite 0 hit](../PPU_registers/#Sprite_0).
3.  Each frame, fill shadow OAM from start to finish based on the position of the actor, plus the position of the individual sprites within the actor's cel, minus the position of the camera.
4.  Replace the Y coordinates of unused sprites with \$FF. (Any value below the bottom of the screen, at least \$EF, will work.)

That way, you can perform flicker by varying the order in which actors are drawn during step 3. And if you skip drawing a sprite because it is offscreen, you can skip increasing the OAM used variable.

For example, [`player.s` in nrom-template](https://github.com/pinobatch/nrom-template/blob/master/src/player.s) uses a variable called `oam_used` to store an index into shadow OAM, separate `player_xhi` and `player_yhi` variables to store the position of the player character, `player_frame` to store which cel to use, and [`ppuclear.s`](https://github.com/pinobatch/nrom-template/blob/master/src/ppuclear.s) demonstrates filling the remainder of the display list after `oam_used` with \$FF values.
