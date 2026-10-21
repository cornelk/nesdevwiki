---
title: "Enhancement"
url: "/wiki/Enhancement/"
description: "Some emulators intentionally deviate from perfect accuracy in the interest of improving a game's appearance. Some of these can be realized in hardware; others cannot."
wikiPage: "Enhancement"
---

Some [emulators](../Emulators/) intentionally deviate from perfect [accuracy](../Accuracy/) in the interest of improving a game's appearance. Some of these can be realized in hardware; others cannot.

## Overdraw {#Overdraw}

Some emulators have an option to ignore the 8 sprites per line limit. There are two ways that this can interact with [sprite evaluation](../PPU_sprite_evaluation/). One is to turn off the mapper side effects of sprite data fetching past the eighth sprite, which can't be realized with actual carts.

The other way involves rearranging the fetch sequence of the sprite unit to allow 15 sprites on a line. This requires rearranging both in-range search and pattern fetch. Cycles 1-64 are normally spent clearing secondary OAM. Doubling the sprite count would make it take twice as long, but clearing only the Y coordinate would make it four times as fast, for a net result of cycles 1-30. This would allow the actual evaluation, which takes 2 cycles per sprite and 6 more per in-range sprite, to be expanded from 176 cycles to 218 cycles. For pattern fetch, skipping seven of the dummy nametable fetches at x=265, 273, 281, 289, 297, 305, and 313 would allow a total of 15 sprites' patterns to be fetched in one horizontal blank. The first dummy fetch at x=257 lets the timer of [MMC3](../MMC3/)-family mappers trigger at the correct horizontal position, but it wouldn't work with the timer of [JY Company](../J.Y._Company_ASIC/) mapper in PA12/8 mode.

A video of the RetroUSB AVS famiclone [demonstrates its 15-sprite mode](https://www.youtube.com/watch?v=1XrjgZhG7q4) using *Teenage Mutant Ninja Turtles*.

## Control {#Control}

An emulator designed for point-and-click games such as [Videomation](../Videomation/) and Thwaite may support using a PC mouse, a touch screen, a Wii Remote, or a Kinect sensor to move a game object. This can take one of three forms:

- The emulator can directly manipulate coordinates in RAM, as if a Pro Action Replay had a controller plugged into it.
- The emulator can patch the game to read absolute coordinates from mapper ports or from the [Arkanoid controller](../Arkanoid_controller/) (if only one dimension is needed), [Oeka Kids tablet](../Oeka_Kids_tablet/), or ~~Power Glove~~.
- The emulator can patch the game to read mickeys (coordinate differences over time) from mapper ports or from the [Super NES Mouse](../Mouse/).

## Resolution {#Resolution}

The cheap way to upgrade a game's graphics from the NES's native 240p to something bigger is to apply pixel art resampling algorithms such as Super 2xSaI, Scale2x, or hq2x. A more powerful way involves creating a large graphic for each small graphic used by the game, as seen in high-resolution texture packs for HDNes, HiSMS, and several Nintendo 64 emulators. For CHR ROM games, this could be a CHR ROM with more pixels and more bit depth than the game's existing CHR ROM. Rendering would use the bigger one, while \$2007 readback and sprite 0 detection would still use the smaller one. Many CHR RAM games could use a hash table from 16-byte original tiles to larger tiles, though this would fail in games relying heavily on [dynamic drawing to CHR RAM](../CHR_ROM_vs._CHR_RAM/#CHR_RAM) (e.g. *Hatris*, *Qix*, *Videomation*, *Solstice*, *Elite*, the *Blockout* prototype and its clone *3D Block*, the second level of *Battletoads*, dialogue in *Super Bat Puncher* and *Nova the Squirrel*, menus in *RHDE*, the [Action 53](../Action_53/) menu, and a forthcoming Z-machine interpreter).

## Map logging {#Map_logging}

A game incrementally updates the nametables and then writes the new [scroll position](../PPU_scrolling/) to the PPU to make them visible. This means an emulator can't see the entire level. But an emulator could watch the [scroll registers](../PPU_registers/) and assemble newly visible areas into a map of the level. And then it could save the map for the next playthrough in order to provide a widescreen view of upcoming areas, automatic creation of an atlas for a game guide, or even a view of a full map on a high-definition screen.[\[1\]](https://gaming.stackexchange.com/a/254791/69014)

The [map maker in Nintaco](http://nintaco.com/faq.html#Emulation) does some of this.

## Sound {#Sound}

### Pop reduction {#Pop_reduction}

Some games make harsh popping noises when they begin to play sampled sound on [APU DMC](../APU_DMC/). This is caused by a write to \$4011 to set the center line of the sample instead of relying on saturation to push the sample toward the center. A few emulators support "pop reduction", which ignores all writes to \$4011, but this fails when games intentionally cause pops by using timed writes to \$4011, such as the speech in *Big Bird's Hide and Speak* and *Joshua* and a few games that use \$4011 pops as a crude kick drum.

One way to accommodate both of these cases without using CRC to detect game-specific settings involves a sixth audio channel and a heuristic to distinguish sample start pops from intentional changes. If \$4011 is written, and it is a sample start pop, add the old \$4011 value to the sixth channel and subtract the new \$4011 value. This way, the pop is exactly canceled out by the sixth channel, which [decays over the course of a few frames](https://en.wikipedia.org/wiki/Leaky_integrator) to avoid [overflow](http://forums.nesdev.org/viewtopic.php?p=115276#p115276). The most reliable heuristic to distinguish a sample start pop from other uses of \$4011, such as kick pops and the DMC sawtooth wave trick, is that a sample start pop occurs in the same frame as a rising edge on \$4015 bit 4 and at least 3 frames after the previous \$4011 write.

### Recorded music playback through the mapper {#Recorded_music_playback_through_the_mapper}

The Japanese version of *Bases Loaded* uses the [JF-13](../INES_Mapper_086/) board, which has a sampled audio playback chip. Write a value, and a sound begins to play through the mapper sound channel of the authentic Famicom console. Likewise, the TurboGrafx-16 and Sega Genesis consoles supported CD peripherals allowing playback of Red Book audio during game play. Emulators could support adding a virtual chip that plays higher quality sound. Each such game would come with .mp3 or .ogg files and a cue sheet to tell what sounds to play when what values are written. Games could be patched to allow the use of, say, OCRemix versions of songs.
