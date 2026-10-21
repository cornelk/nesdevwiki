---
title: "Colour-emphasis games"
url: "/wiki/Colour-emphasis_games/"
description: "The following is a list of games which make use of PPU register $2001's colour emphasis bits (D7 to D5):"
wikiPage: "Colour_emphasis_games"
---

The following is a list of games which make use of [PPU register \$2001](../PPU_registers/#Mask_.28.242001.29_.3E_write)'s colour emphasis bits (D7 to D5):

## Commercial {#Commercial}

- *Bubble Bobble* for the Famicom Disk System uses red emphasis to tint the screen throughout the game, which is absent in the cartridge releases.
- *Family BASIC* allows the user to set PPU emphasis bits through the POKE command. V3 adds the "FILTER" command for this purpose.
- *The Fantastic Adventures of Dizzy* uses it in the beginning when the scroll is unfurled.
- *Final Fantasy* and *Final Fantasy II* rapidly sets and clears all 3 emphasis bits to flash the screen when going into a battle screen, also rotates colors when entering the status screen.
- *[Super Game's Aladdin](http://community.fandom.com/wiki/c:bootleggames:Aladdin_(Super_Game))* uses red emphasis to tint the screen during gameplay, and all 3 emphasis bits to dim when paused. The official version of *Aladdin*, released only in Europe, also uses the emphasis bits to tint the screen green during gameplay (as the PAL console flips the red and green emphasis bits), but does not dim the screen when paused.

The following games set all 3 emphasis bits at once to dim the screen. Contrary to what some older documentation claims, this is fairly well tolerated by TVs of the era and not cause for a game to fail certification.

- *Cliffhanger*
- *Cybernoid - The Fighting Machine* on the top part of the difficulty select and better luck next time screens. Also uses red emphasis when the screen changes in-game.
- *Dragon's Lair*
- *Felix the Cat*
- *Fun House* for fading out
- *The Immortal*
- *Indiana Jones and the Last Crusade (Taito)*, story scenes and some action scenes.
- *James Bond Jr*
- *The Jungle Book*
- *Just Breed*, whose box has a notice about incompatibility with the Sharp C-1 TV
- *Lagrange Point* uses colour emphasis, in some frames combined with grayscale bit to do a very cheap fade-out effect when switching to and from the menu on the playfield, and also when entering into battle
- *Legend of Prince Valiant, The (E)* uses full emphasis on the title screen, green emphasis for gameplay and red emphasis for status bar
- *Lethal Weapon*
- *The Lion King (E)*
- *Magician*
- *Mickey's Adventures in Numberland*
- *Muppet Adventure: Chaos at the Carnival*
- *Noah's Ark (E)* uses blue emphasis combined with grayscale mode to put part of the level underwater and all 3 emphasis bits in other cases
- *Qix*
- *Rampart* uses color emphasis on the player select screen, the options screen and some of the scrolling text. Blue emphasis is used for the score tally bar.
- *R.B.I. Baseball 3*
- *Space Shuttle Project*
- *Super Spy Hunter*, when paused. The Japanese version *Battle Formula* doesn't do this.
- *Xiao Ma Li*

## Homebrew {#Homebrew}

- [Wall](http://nesdev.org/wall.zip) by Chris Covell uses a blue emphasis to simulate the water flooding the wall.
- [Graphics editor](http://pineight.com/nes/#editor) by Damian Yerrick allows applying a `SCREEN TINT` to the picture.
- [Munchie Attack](http://nesdev.org/Munchie_Attack.zip) by Memblers dims the screen upon Game Over.
- [Super Bat Puncher Demo](http://morphcat.de/superbatpuncher/) dims the screen during pause.
- Battle Kid and Battle Kid 2 dim the screen during pause.
- [Waddles the Duck](https://cpprograms.net/classic-gaming/waddles-the-duck/) dims the screen when on a warp, and RG-emphasis when in ice-dimension, among others.
- [All Hell Unleased](https://8bitslasher.itch.io/ahu) by 8-Bit Slasher dims the screen during pause

### Commercial homebrew {#Commercial_homebrew}

- [Full Quiet](https://www.retrotainmentgames.com/collections/video-games/products/copy-of-full-quiet-regular-edition-nes-game-green-glow-cartridge-complete-in-box-cib) applies various tints over the course of its day/night cycle.
- [Project Blue](https://toggleswitch.itch.io/projectblue) sets R and G emphasis bits together on the ending screen, to achieve a warmer tan-like tint.

## <span id="Multicarts.27_menus"></span>Multicarts' menus {#Multicarts.27_menus}

- *[400 in 1](https://bootleggames.fandom.com/wiki/400-in-1_Multicart) sets all 3 bits*

## Romhacks {#Romhacks}

- *[Rockman 4 Minus Infinity](https://sites.google.com/site/rockman4mi/)* (during Brightman's stage) darkens areas farther away from Mega Man.
- *[Bisqwit's Castlevania II retranslation](http://bisqwit.iki.fi/cv2fin/)* uses the de-emphasis bits to highlight a section of a screen when scrolling the list of savegame slots.

## References {#References}

- [BBS topic](http://forums.nesdev.org/viewtopic.php?t=8623)
- [BBS topic \#2](http://forums.nesdev.org/viewtopic.php?f=16&t=11995)
