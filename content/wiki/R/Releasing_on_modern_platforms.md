---
title: "Releasing on modern platforms"
url: "/wiki/Releasing_on_modern_platforms/"
description: "NES game developers who qualify can increase their audience by releasing their games through modern platforms' storefronts, such as Steam, iOS App Store, PlayStation Store, and Nintendo eShop."
wikiPage: "Releasing_on_modern_platforms"
---

NES game developers who qualify can increase their audience by releasing their games through modern platforms' storefronts, such as Steam, iOS App Store, PlayStation Store, and Nintendo eShop.

## Emulation middleware {#Emulation_middleware}

It's common to port NES games written in assembly language to a modern platform using middleware consisting of an [emulator](../Emulators/) locked down to run only the included game. However, modern platforms' storefronts impose restrictions on which software licenses may be distributed. If you plan to release your NES game on such a storefront, be sure to review the distribution terms and choose an emulator with a compatible license. Copyleft licenses, such as the GNU General Public License (GPL), often tend to be incompatible. Emulators under permissive licenses, such as the BSD License, the MIT License, or the zlib License, may be used.

Some copyleft case studies:

Steam  
Valve warns the public that the Steamworks SDK license is incompatible with the GPL.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>

iOS  
Apple's App Store offers no means for users to redistribute applications received through the service. This is an unacceptable "additional restriction" per the GPL.<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

Modern Nintendo consoles  
When Atari and Majesco hired Mistic Software to port Humongous Entertainment's *Freddi Fish* and *Pajama Sam* to the Wii console, Mistic ended up using ScummVM, a middleware under the GPL without a dual license option. Nintendo's license for the Wii SDK was incompatible with the GPL. Some developers settled with Atari, after which the games were pulled from the market and Atari donated to the Free Software Foundation.<sup><span id="cite_ref-3"></span>[\[3\]](#cite_note-3)</sup>

SameBoy (MIT License) was used to port the Game Boy versions of *Aladdin* and *The Lion King* to Nintendo Switch. mGBA (Mozilla Public License) was used to port some *Medabots* games.

You should also consider the following technical aspects when evaluating emulation middleware:

- Compatibility, and [accuracy](../Accuracy/) to the extent that it increases compatibility with your software
- How much porting effort is needed to run a game in the emulated environment
- Runtime performance, particularly on web, mobile, and Nintendo Switch
- Download size
- Ability to map keyboard keys or modern platform buttons, such as by emulating a [mouse](../Mouse/) or a [controller with more buttons](../SNES_controller/)
- How the front end's user interface can be themed to blend in with the game
- Save states and rewinding to support rollback netcode in multiplayer games

Members of the community have proposed using these for porting from NES:

- [Rustico](https://github.com/zeta0134/rustico) (formerly RusticNES) by Zeta, Rust, MIT License. Includes a WebAssembly front end for showing a cut-down demo version on your website.
- [NES Bundler](https://github.com/tedsteen/nes-bundler) by Ted Steen, Rust, MIT License. Front end for Rustico that supports Windows, macOS, and desktop Linux.
- [GGVm](https://github.com/gradualgames/ggvm) by Gradual Games, Java, Unlicense. Windows and Android front ends. This emulates the CPU and a subset of PPU functionality. **Games must be ported.** Timing is entirely fake. Sound is high-level emulated through MP3 audio. It's recommended to use conditional code (such as `.if` statements in ca65) to make builds of your game that target NES and GGVm.

## High-level languages {#High-level_languages}

It's also possible to write the game logic in C, use llvm-mos or cc65 to compile it, and write front ends that handle input and output on each platform. These front ends would be written in assembly language on NES or in the preferred high-level language on modern platforms. This lets you rely on the compiler to keep the game logic identical between the NES and modern versions, while allowing the modern version to act as a remaster with high-definition graphics. Keeping the game state small also reduces the size of data that must be rewound after a received button press in rollback netcode.

Koei's NES games probably use this technique, given what is known about their bytecode interpreter.<sup><span id="cite_ref-4"></span>[\[4\]](#cite_note-4)</sup>

## Aspect ratio {#Aspect_ratio}

Modern platforms have widescreen displays with square pixels. You'll need to decide how much you're horizontally stretching the image (see [Overscan](../Overscan/)), as well as what to put in the areas off to the sides. Some suggestions:

- The Game Boy Advance ports of *Super Mario Bros. 3* and *Kirby's Adventure* take the HUD elements out of the bottom status bar and display them over the playfield. This is easier to do in an HLE or HLL environment but can be done in emulation with color key.
- *Full Quiet* and *Garbage Pail Kids* come with a selection of borders that can be applied to the sides of the screen, much as with the Super Game Boy accessory for Super NES.
- One option is to display 256x240 pixels' worth of scrolling background on the NES or 352x240 on modern platforms (stretched to 1206x720 or 1810x1080). Center the camera 48 pixels in front of the player character on NES using dual forward focus,<sup><span id="cite_ref-5"></span>[\[5\]](#cite_note-5)</sup> or center the character on modern platforms. Either way, the player can see 176 pixels ahead. Tie enemy behaviors, such as respawning, to player position rather than camera position so that turning around doesn't cause enemies to get respawned repeatedly.

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) "[Distributing Open Source Applications on Steam](https://partner.steamgames.com/doc/sdk/uploading/distributing_opensource)". *Steamworks Documentation*. Accessed 2024-03-23.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) Brian A. "[Answer to Does apple approve GPL based apps?](https://forums.developer.apple.com/forums/thread/18922?answerId=61870022#61870022)". *Apple Developer Forums*, 2015-09-23. Accessed 2024-03-23.
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) Eugene Sandulenko. "[GPL, ScummVM and violations](http://sev-notes.blogspot.com/2009/06/gpl-scummvm-and-violations.html)". *Sev's ScummVM Notes*, 2009-06-23. Accessed 2024-03-23.
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) AWJ. "[Koei bytecode](https://forums.nesdev.org/viewtopic.php?f=2&t=15931)". *NESdev BBS*, 2017-05-11. Accessed 2024-03-23.
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) Itay Keren. "[Scroll Back: The Theory and Practice of Cameras in Side-Scrollers](https://www.gamedeveloper.com/design/scroll-back-the-theory-and-practice-of-cameras-in-side-scrollers)". *Game Developer*, 2015-05-11. Accessed 2024-03-23.
