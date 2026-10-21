---
title: "Game bugs"
url: "/wiki/Game_bugs/"
description: "Listed are games that have been tested on NES or Famicom hardware and verified to look wrong or odd. This can be caused by NES hardware limitations, programming errors, or even intentional effects…"
wikiPage: "Game_bugs"
---

Listed are games that have been tested on NES or Famicom hardware and verified to look wrong or odd. This can be caused by NES hardware limitations, programming errors, or even intentional effects within the game. Refer to this if you're developing an emulator and find a game that looks wrong, before you look for a problem in your emulator. If you are attempting to give your emulator ["bug for bug" compatibility](../Accuracy/), you'll want to make sure that these glitches (or unusual behaviors) appear the same in your emulator as they do on the NES.

This is an incomplete list that concentrates on commercial games. For an overview of common compatibility problems in homebrew games, see [Program Compatibility](../Program_compatibility/). Sometimes, a bug slips into a game to make it rely on less-than-intentional features of the hardware; for those, see [Tricky-to-emulate games](../Tricky-to-emulate_games/).

## General bugs {#General_bugs}

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Smb3-attribute-glitch.png"><img src="../../wiki-images/Smb3-attribute-glitch.png" width="256" height="240" /></a>
<figcaption>Attribute glitch in <em>Super Mario Bros. 3</em></figcaption>
</figure>

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="width: 18%">Game</th>
<th>Problem</th>
</tr>
</thead>
<tbody>
<tr>
<td><em>The Addams Family</em></td>
<td>The in-game status bar occasionally bumps vertically by 1 pixel, caused by non-solid background pixels overlapping the sprite zero that is used for the status bar split.</td>
</tr>
<tr>
<td><em>Ai Senshi Nicol</em></td>
<td>Loading CHR-RAM can fail due to a race condition, causing level graphics to be incorrect. This occurs if an NMI lands between the two <a href="../PPU_registers/">PPUADDR</a> writes for the graphics copy, because the NMI handler reads <a href="../PPU_registers/">PPUSTATUS</a>, clearing the address write latch. Whether this happens depends on FDS disk access timing, which does not take a well-defined, fixed amount of time. It is not clear if this can happen with real-world timing variance on real hardware, but it is a concern for emulators.</td>
</tr>
<tr>
<td><em>Akumajou Densetsu</em></td>
<td><ul>
<li>When a door is opened/shut when the player goes through them, the screen shakes by one pixel.</li>
<li>During Stages 6-4 and 6-5, the game uses horizontal scrolling with <a href="../Mirroring/">vertical arrangement</a> (with one nametable reserved for the rising water) and <a href="../PPU_registers/">masks the leftmost column</a>, causing tile/attribute glitches at the left edge and clipping part of the status bar. <em>Castlevania III</em> instead uses a third nametable from the <a href="../MMC5/">MMC5</a>'s ExRAM for the rising water, preventing these issues.</li>
</ul></td>
</tr>
<tr>
<td><em>Akumajou Special</em></td>
<td>The game over screen is sometimes glitched when it occurs in bidirectional scroll stages. (Validated in Game Center CX.)</td>
</tr>
<tr>
<td><em>Armadillo</em></td>
<td><ul>
<li>The in-game status bar occasionally bumps vertically by 1 pixel.</li>
<li>Slowdowns occur frequently.</li>
<li>Crashes on Dendy consoles - fix: <a href="https://forums.nesdev.org/viewtopic.php?p=270654#p270654">[1]</a></li>
</ul></td>
</tr>
<tr>
<td><em>Asmik-kun Land</em></td>
<td>Sets the triangle period to $0003 at the end of snare/tom pitch bends, causing high-pitched (but not ultrasonic) beeps in some music tracks.</td>
</tr>
<tr>
<td><em>Aspic: Majaou no Noroi</em></td>
<td>Disk flip detection is unreliable due to a poorly structured loop for $4023 polls.<sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup> This is especially problematic on emulators (including physical drive emulators) which can swap disk sides instantly. (TODO: Determine real disk drive behaviour)</td>
</tr>
<tr>
<td><em>Atari R.B.I. Baseball</em>, <em>The Quest of Ki</em>, <em>Valkyrieの冒険</em> (all Vs.)</td>
<td>These Vs. Unisystem ports of Namco cartridge games can run on any of the RP2C03 and RP2C04-0001...RP2C04-0004 PPUs by setting the DIP switch appropriately. The look-up table they use for the RP2C04-0001 setting selects the wrong colors for white and light gray however, causing white to be displayed as light gray instead.</td>
</tr>
<tr>
<td><em>Battletoads</em></td>
<td>Clinger Winger (Level 11) is impossible with 2 players as player 1 and the Buzzball start moving before player 2 does. Fixed in the Japanese and European versions.</td>
</tr>
<tr>
<td><em>Balloon Fight</em></td>
<td><ul>
<li>The game over theme writes garbage data to the noise channel on the final note, which happens to enable the <a href="../APU_Noise/#Noise_mode_and_period_.28.24400E.2C_write.29">periodic mode</a> only present on lettered CPU revisions. Fixed in the unreleased FDS version.</li>
<li>The Japanese version fails to reset on frontloaders due to an inadequate <a href="../PPU_power_up_state/#Best_practice">wait before enabling NMIs</a>. Fixed in the NES versions and the unreleased FDS version.</li>
<li>The North American and European NES versions have bugs stemming from sound driver variables overlapping the Balloon Trip leaderboard and the warm-start detection string:
<ul>
<li>The lowest attainable rank in Balloon Trip is 47 instead of the original 50.</li>
<li>High scores are lost between resets. (warm start check always fails)</li>
<li>Resetting the game anywhere outside of the title or demo screens causes the triangle channel to play a portion of the Balloon Trip theme on the title screen.</li>
</ul></li>
</ul></td>
</tr>
<tr>
<td><em>Castlevania</em></td>
<td><ul>
<li>Revision 0 of the US version is prone to crashing when too many sprites are active (for example, when fighting Death), due to the game loading the wrong ROM bank after jumping to the NMI routine.</li>
<li>In the US versions, due to <a href="../APU_DMC/#Conflict_with_controller_and_PPU_read">DMC DMA controller corruption</a>, holding Up when Simon is hit can occasionally cause the game to pause. <em>Castlevania</em> only uses DPCM for a short grunt sample when Simon is hit, and the US versions do not have any controller re-read code, instead seemingly relying on inputs being ignored while the sample plays; Start is still processed to allow pausing, however, and if DMC DMA causes a controller bit deletion before Start is read, an Up input will be read as Start instead. The Japanese versions (FDS and Famicom) use controller re-read code, so they are exempt from this issue.</li>
</ul></td>
</tr>
<tr>
<td><em>Castlevania II - Simon's Quest</em></td>
<td>Sometimes incorrect tiles are shown on the playfield.</td>
</tr>
<tr>
<td><em>Castlevania III - Dracula's Curse</em></td>
<td><ul>
<li>The <a href="../APU_DMC/">DMC channel</a> in music sometimes mutes.</li>
<li>When pressing the B button at the exact time that Trevor falls off a block, the whip sound plays, but Trevor doesn't attack.</li>
<li>When Death reveals his second form, the floor jumps up two scanlines and the upper two lines of the floor tiles disappear. In <em>Akumajou Densetsu</em> the third line of the floor tiles also mostly disappears. This is caused by the programmer not fully understanding <a href="../PPU_scrolling/#Split_X.2FY_scroll">$2006 writes during rendering</a>, and resetting the fine Y scroll to 2 rather than the presumably intended 0.</li>
</ul></td>
</tr>
<tr>
<td><em>Chip 'n Dale Rescue Rangers</em></td>
<td>As the intro sequence is fading out, the game displays sprites with incorrect tiles in the positions for the player initials and life indicators at the top of the screen. When the screen fades back in, the same positions have the correct graphics.</td>
</tr>
<tr>
<td><em>Commando</em></td>
<td><p>Graphical issues primarily stemming from slow PPU updates:<sup><span id="cite_ref-2"></span><a href="#cite_note-2">[2]</a></sup></p>
<ul>
<li>"Uninitialized" nametable and sprite tiles are frequently visible due to staggered <a href="../PPU_pattern_tables/">CHR-RAM</a> updates.</li>
<li>Vertical scrolling is very jerky due to slow PPU VRAM updates sometimes pushing <a href="../PPU_registers/">OAM DMA</a> and <a href="../PPU_registers/">PPUSCROLL</a> updates outside of vblank, effectively offsetting the scroll.</li>
<li>The screen may briefly display garbage after a player death due to a <a href="../PPU_registers/">PPUCTRL</a> write (intended to toggle NMIs) which alters the pattern table mapping and sprite size during rendering.</li>
</ul></td>
</tr>
<tr>
<td><em>Crystalis</em></td>
<td>The scanline above the status bar and above text boxes looks wrong.</td>
</tr>
<tr>
<td><em>Cybernoid: The Fighting Machine</em></td>
<td>Relies on AND <a href="../Bus_conflict/">bus conflict</a> behaviour when enabling music (pause, then press A) due to a leftover $D418 = $0F write (C64 SID volume register).<sup><span id="cite_ref-3"></span><a href="#cite_note-3">[3]</a></sup> The location in PRG-ROM contains $50 and thus should yield a switch to CHR bank 0.</td>
</tr>
<tr>
<td><em>Devil World</em></td>
<td>JP revision 0 fails to boot if an NMI fires during the first frame, due to the game only initializing <a href="../PPU_registers/">PPUCTRL</a> (to disable NMIs on toploaders) after waiting for one vblank. Fixed in JP revision 1 and the PAL version.</td>
</tr>
<tr>
<td><em>Dizzy The Adventurer</em></td>
<td>Resets the sound phase every frame, causing a nasty 60Hz buzz.</td>
</tr>
<tr>
<td><em>Donald Land</em></td>
<td>When the player progresses too quickly by boosting off of apples, the background loads fall behind and the scroll seam becomes visible.</td>
</tr>
<tr>
<td><em>Donkey Kong</em></td>
<td>Lag frames can cause glitched tiles to appear. (E.g. by using the ladder glitch to stall on 50m) This is because the <a href="../PPU_registers/">PPUCTRL</a> write to re-enable NMIs can occur while the vblank flag is set, triggering an <a href="../PPU_registers/#Vblank_NMI">NMI late into vblank</a> and causing PPU nametable tiles to be written during rendering. The <a href="../Family_Computer_Disk_System/">FDS</a> version fixes this by waiting until the vblank flag in <a href="../PPU_registers/">PPUSTATUS</a> is clear before toggling NMIs.</td>
</tr>
<tr>
<td><em>Door Door</em></td>
<td><a href="../PPU_palettes/">Palette RAM</a> is not initialized correctly if the VBL flag (<a href="../PPU_registers/">PPUSTATUS</a> bit 7) is set on power-up/reset due to the PPU's <a href="../PPU_power_up_state/">power-up state</a>.</td>
</tr>
<tr>
<td><em>Double Dragon</em></td>
<td><ul>
<li>"Garbage sprites" (sprite 0 (for <a href="../PPU_OAM/">sprite-0 hit</a>) and sprite 1) can be seen in the lower right of the game screen.
<ul>
<li>Sprite 0 consists of tile $FF (a black tile with 2x2 non-background pixels (i.e. a tile with a 2x2 "dot" in it, visually similar to ▣ or ⚀)), and the priority bit set.</li>
<li>Sprite 1 consists of tile $FE (a tile consisting entirely of a single non-transparent colour, often palette entry $0F but varies per stage).</li>
</ul></li>
<li>The screen will sometimes shake vertically on heavy sprite usage.</li>
</ul></td>
</tr>
<tr>
<td><em>Double Dragon II</em></td>
<td>The status bar may suddenly change colors: sometimes when scrolling vertically it shows incorrectly for a couple of frames.</td>
</tr>
<tr>
<td><em>Double Dragon III</em></td>
<td>Same status bar issue as <em>Double Dragon II</em>.</td>
</tr>
<tr>
<td><em>Dr. Mario</em></td>
<td>Controller inputs can be corrupted while music is playing, due to a faulty handling of <a href="../APU_DMC/#Conflict_with_controller_and_PPU_read">DMC DMA controller corruption</a>.<sup><span id="cite_ref-4"></span><a href="#cite_note-4">[4]</a></sup> This can be easily seen by moving a pill around while holding a rotation button. The PAL 2A07 CPU does not have this conflict to begin with, so the glitch is avoided entirely in the European release.</td>
</tr>
<tr>
<td><em>DuckTales</em></td>
<td>The scroll split between the status bar and playfield is placed in the middle of the status bar's bottom-most scanline, causing it to jitter when scrolling horizontally. Additionally, the game uses different sprites for the status bar and playfield that are swapped mid-frame via a <a href="../PPU_registers/">PPUCTRL</a> write, and hides the status bar's sprites whenever the player's metasprite overlaps it, further complicating sprite 0 hit behavior and causing the bottom-most scanline to scroll across the entire screen instead.</td>
</tr>
<tr>
<td><em>Duck Maze</em></td>
<td>The original 60-pin release for famiclones relies on decimal mode working as it does in a 6502. Without decimal mode, score counting works wrong, and the game may never finish counting down at the end of a level. (Fixed for HES's 72-pin rerelease.)</td>
</tr>
<tr>
<td><em>Eliminator Boat Duel</em></td>
<td><ul>
<li>Audible buzzing due to the <a href="../APU_DMC/">DMC channel</a> being enabled in its <a href="../CPU_power_up_state/">power-on state</a> every frame.<sup><span id="cite_ref-5"></span><a href="#cite_note-5">[5]</a></sup></li>
</ul></td>
</tr>
<tr>
<td><em>Exed Exes</em></td>
<td>When pausing, the immediate next note of the music will play after the pause jingle completes.</td>
</tr>
<tr>
<td><em>Famicom Grand Prix - F-1 Race</em></td>
<td>The screen flickers on lag frames during races. This is because the game performs extra PPU transfers during forced blanking (blank area beneath the status bar), which can occur late enough into the frame to push a <a href="../PPU_registers/">PPUSTATUS</a> read and <a href="../PPU_registers/">PPUCTRL</a> write (to enable NMIs) into vblank without actually triggering an NMI for that frame.</td>
</tr>
<tr>
<td><em>F-1 Race</em></td>
<td><ul>
<li>The game displays garbled graphics and/or freezes if the Start button is held during power-up/reset.</li>
<li>Will not boot on NES-001 (frontloader) due to its <a href="../PPU_power_up_state/#Best_practice">PPU warm-up loop</a> being too short.</li>
</ul></td>
</tr>
<tr>
<td><em>F-15 City War</em></td>
<td>The last boss is glitched when playing through the game from start to finish in its original <a href="../INES_Mapper_173/">INES Mapper 173</a> and AVE "1.x" version. It is not glitched when jumping to the last boss through cheating in those versions, and never glitched in the later AVE "1.1" and Gluk Video versions.</td>
</tr>
<tr>
<td><em>Ghostbusters (J)</em></td>
<td>Loads the ending text from the wrong CHR bank, causing it to display a blank screen that eventually scrolls "りり" (hiragana "riri") on the screen.</td>
</tr>
<tr>
<td><em>Ghosts 'n Goblins</em></td>
<td><ul>
<li>Relies on the VBL flag (<a href="../PPU_registers/">PPUSTATUS</a> bit 7) being clear to <a href="../PPU_power_up_state/#Best_practice">wait before enabling NMIs</a> during the reset handler. If it is set, the write to <a href="../PPU_registers/">PPUCTRL</a> occurs too early and the game fails to boot.</li>
<li>Garbage sprites sometimes appear on-screen due to incomplete display lists being uploaded to OAM. (see <em>Strider</em>)</li>
</ul></td>
</tr>
<tr>
<td><em>Gimmick!</em></td>
<td>The game's controller polling routine relies on extra <a href="../APU_DMC/#Conflict_with_controller_and_PPU_read">DMC DMA controller corruption</a> only present on select hardware revisions (original RF, Twin, and Titler Famicoms) to detect if a controller read was corrupted. On NES consoles and the New/AV Famicom, DMC DMA glitches are not as pronounced, causing the routine to accept a corrupted read. This results in occasional spurious inputs such as pausing when holding up, or turning/inching rightwards while idle. The PAL 2A07 CPU does not have this conflict, so the glitch is avoided entirely in the European release.</td>
</tr>
<tr>
<td><em>Hebereke</em></td>
<td><ul>
<li>The triangle channel in the title screen music is silent, except after exiting from the password screen.</li>
<li>The HUD disappears when the player overlaps its vertical position, due to both using the same CHR ROM sprite bank that is swapped mid-frame.</li>
</ul></td>
</tr>
<tr>
<td><em>Hottarman no Chitei Tanken</em></td>
<td>If gameplay slows down due to excessive on-screen objects, scrolling may glitch and display garbage values.</td>
</tr>
<tr>
<td><em>The Immortal</em></td>
<td>Uses color $0D (blacker than black) as black and color $0E (normal black) as a darker gray than color $2D while setting all three deemphasis bits, causing picture instability on many television sets and TV capture cards. (see <a href="../Color_$0D_games/">Color $0D games</a> for more)</td>
</tr>
<tr>
<td><em>Ironsword: Wizards and Warriors II</em></td>
<td>Noise channel doesn't work properly, sometimes plays longer notes and sometimes mutes.</td>
</tr>
<tr>
<td><em>Jackal</em></td>
<td>Would rely on power-up state were it not for <a href="../Bus_conflict/">bus conflict</a> behavior. The identity table at $8000 is in banks 0-6 but filled with $AA in bank 7. If the '161 powers up with value $07, the game relies on AND behavior to get out of bank 7 for the first time. The instructions are <code>lda $8000,Y :: sta $8000,Y :: sta $8000,Y</code>. With Y=0 (as when trying to switch to bank 0), this would read $AA, then write $AA over $AA to switch to bank 2, then write $AA over $00 (a bus conflict) that switches to bank 0.</td>
</tr>
<tr>
<td><em>Kirby's Adventure</em></td>
<td><ul>
<li>When Kirby copies a new ability, the status bar icon may flicker or display incorrect attributes. (?)</li>
<li>Slowdowns occur frequently.</li>
<li>Jump and attack inputs are sometimes ignored.</li>
</ul></td>
</tr>
<tr>
<td><em>Kujaku Ou II</em></td>
<td>The music plays incorrect <a href="../APU_DMC/">DPCM</a> drum sounds due to the $C000-$FFFF bank (<a href="../MMC1/">MMC1</a> PRG mode 2) being swapped out with unrelated data during <a href="../DMA/#DMC_DMA">DMC DMA</a>.<sup><span id="cite_ref-6"></span><a href="#cite_note-6">[6]</a></sup> (Note that <a href="../NSF/">NSF</a> rips inadvertently fix this by stripping out unrelated banks)</td>
</tr>
<tr>
<td><em>The Legend of Zelda</em></td>
<td>Vertical scrolling between screens is very messy, jumping by 3-4 scanlines because of an incorrect fine y from the $2006 writes, two reads from $2007, and imprecise timing on these accesses that can cause a $2007 read's y increment to be lost.<sup><span id="cite_ref-7"></span><a href="#cite_note-7">[7]</a></sup> Vertical scrolling while DPCM is playing can cause the vertical scroll to momentarily jump due to conflicts between 2nd $2006 writes and automatic fine y increments. See <a href="../PPU_scrolling/">PPU scrolling</a> for details on how these registers interact with scrolling.</td>
</tr>
<tr>
<td><em>Linus Spacehead's Cosmic Crusade</em></td>
<td>The music lacks noise channel instruments due to a sound driver bug.</td>
</tr>
<tr>
<td><em>The Lord of King</em></td>
<td>Will not boot on NES-001 (frontloader) due to an inadequate <a href="../PPU_power_up_state/#Best_practice">wait before enabling NMIs</a>. Fixed in <em>Astyanax</em>.</td>
</tr>
<tr>
<td><em>Magic John</em></td>
<td>Same as <em>The Lord of King</em>. Fixed in <em>Totally Rad</em>.</td>
</tr>
<tr>
<td><em>Mega Man</em></td>
<td><ul>
<li>When pausing and unpausing the game with the Select button, the high bits of the triangle channel's period are not rewritten until Mega Man's teleporting animation finishes, causing distorted, high-pitched noise due to them being left at 0 until then.</li>
<li>If Fire Man is defeated and the stage clear orb is obtained while his ground flame remains active, it will incorrectly use tiles loaded for the stage clear text before being deleted, due to them occupying the same space in CHR RAM.</li>
</ul></td>
</tr>
<tr>
<td><em>Mega Man 3</em></td>
<td><ul>
<li>The first scanline of the weapon menu and the scanline above Shadow Man in the stage select screen are glitched due to improperly timed scroll splits.</li>
<li>In revision 0 of the PAL version, when the Wily Machine is fought, the floor heavily jitters up and down across the screen. This is because two <a href="../MMC3/">MMC3</a> scanline counter IRQs are used for the Machine and the floor, but after the first IRQ, the new reload value (116) is written to $C000 without a write to $C001, which can sometimes happen at the same time the counter is clocked from 0 by PPU A12, overwriting it with the old reload value (58) and causing the split to happen much earlier. This was fixed in later revisions by changing the timing of the $C000 write.</li>
</ul></td>
</tr>
<tr>
<td><em>Mega Man 5</em></td>
<td>In Gyro Man's stage, inside the two elevators, the floor jumps up two scanlines when the elevator starts moving, and jumps back down when the screen starts scrolling horizontally as you leave the elevator area. This bug has the same cause as the one with Death in Castlevania III.</td>
</tr>
<tr>
<td><em>Metroid</em></td>
<td>The ending song of the FDS version is <a href="https://youtu.be/NXlgXfI-6eQ?t=859">supposed to use volume envelopes</a> on the <a href="../FDS_audio/">wavetable channel</a>, yet in-game playback plays <a href="https://www.youtube.com/watch?v=YpLP8aGXcWs&amp;t=3208s">without them</a>. The reason is that the memory location that holds the value that will be written to $4080 on the next start of a note is re-used for some other purpose (routine at PC $6779 in revision 3).</td>
</tr>
<tr>
<td><em>Micro Machines</em></td>
<td>Graphical problems on PPU revisions prior to 2C02G due to <a href="../PPU_registers/">OAMDATA</a> unreadability. Resetting on Famicom and NES-101 (toploader) doesn't always work because the game fails to clear <a href="../PPU_registers/">PPUCTRL</a> and <a href="../PPU_registers/">PPUMASK</a> on reset.</td>
</tr>
<tr>
<td><em>Mitsume ga Tooru</em></td>
<td>Garbage data is visible in the upper side of the status bar by when shaken by an earthquake, due to the status bar and playfield both sharing the nametables.</td>
</tr>
<tr>
<td><em>Monopoly</em></td>
<td>Holding any direction on the Control Pad while landing on a property owned by another player when you don't have enough cash will skip paying the rent.</td>
</tr>
<tr>
<td><em>Neunzehn - 19</em></td>
<td>Disk flip detection is sluggish due to the game waiting ~512ms between $4023 reads. This is especially problematic on emulators (including physical drive emulators) which can swap disk sides instantly.</td>
</tr>
<tr>
<td><em>Othello</em></td>
<td>The FDS DV 2 and Famicom versions can crash and show garbled graphics<sup><span id="cite_ref-8"></span><a href="#cite_note-8">[8]</a></sup> during CPU matches, due to a stack overflow caused by nested subroutine calls in the CPU player logic. Seemingly fixed in FDS DV 3,5(?), leaked Rev 1 FC, and NES versions.</td>
</tr>
<tr>
<td><em>Othello</em> (Famiclone)</td>
<td>The original 60-pin release for Famiclones relies on decimal mode working as it does in a 6502. Without decimal mode, spots are counted wrong. (Fixed for HES's 72-pin rerelease.)</td>
</tr>
<tr>
<td><em>Panic Restaurant</em></td>
<td>The in-game status bar is bumped up by 1 pixel.</td>
</tr>
<tr>
<td><em>Rad Racer</em></td>
<td>Steer to the far left, and a background scanline will be seen on the road.</td>
</tr>
<tr>
<td><em>Rambo</em></td>
<td>One scanline is occasionally glitched, for the same reason as in <em>Super Mario Bros.</em><sup><span id="cite_ref-9"></span><a href="#cite_note-9">[9]</a></sup></td>
</tr>
<tr>
<td><em>Rampart</em> (Jaleco)</td>
<td>During build phase, the drums (on the noise channel) drop out fairly early.</td>
</tr>
<tr>
<td><em>Romancia</em></td>
<td>PRG1 version has a slightly shaky status bar due to a very imprecise wait-for-sprite-0-hit loop. PRG0 version is fine.</td>
</tr>
<tr>
<td><em>Snow Bros.</em></td>
<td>When you clear the stage and rise to the next floor, incorrect CHR bank switching results in glitches in the new floor's graphics.</td>
</tr>
<tr>
<td><em>Solar Jetman</em> (NTSC version)</td>
<td>Some songs use the sweep registers heavily, which are not restored after a sound effect plays that uses the sweep registers as well. The PAL version seems to have corrected this error.</td>
</tr>
<tr>
<td><em>StarTropics</em></td>
<td>The music data for two songs were not set up correctly.<sup><span id="cite_ref-10"></span><a href="#cite_note-10">[10]</a></sup>
<ul>
<li>The island map music (NSF track 1) has a problem with the second pulse channel: it is intermittently silent or playing the wrong notes after the first minute or two, because the music data was not made to fit/repeat properly.</li>
<li>The music for one of the ending cutscenes has a silent triangle channel except for a single glitched note due to an error in the music data.</li>
</ul></td>
</tr>
<tr>
<td><em>Strider</em></td>
<td>The NMI handler pushes an incomplete display list to OAM during lag frames, leading to visible garbage sprites.<sup><span id="cite_ref-11"></span><a href="#cite_note-11">[11]</a></sup></td>
</tr>
<tr>
<td><em>Super Cars</em></td>
<td>The top of the screen flickers on PPU revisions prior to 2C02G due to <a href="../PPU_registers/">OAMDATA</a> unreadability.</td>
</tr>
<tr>
<td><em>Super Mario Bros.</em></td>
<td><ul>
<li>The status bar shakes horizontally on heavy sprite usage and the music slows down. This can be seen especially in worlds 6-4, 7-4 and 8-4, where the large number of hammer objects created by Bowser's code causes the processing time to overshoot a frame. NMIs are disabled on entry to the NMI code and only reenabled when the CPU is "idle", thus when the overshoot occurs, the CPU "misses" a frame, causing general slowdown and status bar shakiness.</li>
<li>At various parts of 1-2, in certain CPU/PPU alignments, a single scanline gets glitched. This is caused by writing to <a href="../PPU_registers/">PPUCTRL</a> to reenable NMI at the exact end of the previous scanline, causing the PPU to begin that scanline from the first nametable instead of the second.<sup><span id="cite_ref-12"></span><a href="#cite_note-12">[12]</a></sup></li>
<li>When hitting a <code>?</code> block while maintaining a run at top speed, after the bounce animation finishes, the block disappears for up to two frames before settling in the background. (Updating the nametable at the <a href="../PPU_scrolling/">scroll seam</a> has priority.)</li>
</ul></td>
</tr>
<tr>
<td><em>Super Mario Bros. 3</em></td>
<td><ul>
<li>The first scanline after a scroll split is glitched. This shows up as garbage above the left side of the status bar and as incorrectly scrolled lines in the "spade" (not N-spade) bonus game.</li>
<li>Note blocks containing items become squarer for a second while the item is popping out. (This is an artifact of the <a href="../PPU_sprite_priority/">sprite priority</a> exploit that it uses.)</li>
<li>If a Hammer Bros. battle ends precisely when a note is starting, the note will freeze on an incorrect duty cycle.</li>
<li>Big fat attribute glitch on the right side of most levels, because this game uses horizontal scrolling with <a href="../Mirroring/">vertical arrangement</a>. Discussed heavily.<sup><span id="cite_ref-13"></span><a href="#cite_note-13">[13]</a></sup></li>
</ul></td>
</tr>
<tr>
<td><em>Teenage Mutant Ninja Turtles</em></td>
<td><ul>
<li>Sprite overflow (due to large numbers of enemies) causes the status bar to flicker.</li>
<li>Uses <a href="../PPU_registers/">Y scroll values</a> greater than 239, causing the PPU to read the attribute table as nametable data before looping back to the same nametable instead of rolling to the next nametable down.</li>
</ul></td>
</tr>
<tr>
<td><em>Tetris</em> (Bullet-Proof Software)<em></em></td>
<td>Sets <a href="../APU_Triangle/">APU Triangle</a> period to longest value instead of actually muting it.</td>
</tr>
<tr>
<td><em>Tetris 2 + Bombliss</em></td>
<td>The 2-player mode will become softlocked (requiring a reset) if the second player presses the start button to pause the game, as the game does not properly account for the start/select buttons being present on the second controller in this scenario. This is possible if expansion port controllers are connected, or if the game is being played on the NES or the New/AV Famicom.</td>
</tr>
<tr>
<td><em>Titanic Mystery: Ao no Senritsu</em></td>
<td>Same disk flip/swap detection problems as <em>Aspic: Majaou no Noroi</em>.</td>
</tr>
<tr>
<td><em>Topple Zip</em></td>
<td>On PPU revisions 2C02G and onwards, garbage sprites may be visible at the start of a level after power-on/reset. This is because the game accidentally reads from a mirror of <a href="../PPU_registers/">OAMDATA</a> during rendering to fetch an index to a sprite object pointer in this scenario, seemingly relying on <a href="../Open_bus_behavior/#PPU_open_bus">PPU open bus</a> from unreadable OAMDATA. If OAMDATA is readable instead, then further fetches for the sprite object use more <a href="../Mirroring/#Memory_Mirroring">PPU register mirror</a> reads during rendering, resulting in the creation of garbage sprites.<sup><span id="cite_ref-14"></span><a href="#cite_note-14">[14]</a></sup></td>
</tr>
<tr>
<td><em>Zelda II: The Adventure of Link</em></td>
<td><p>Reads from <a href="../PPU_registers/">PPUDATA</a> during the title screen twice, moving the background upward by 2 scanlines after the split point.</p></td>
</tr>
<tr>
<td><em>Zombie Nation</em></td>
<td>Same as <em>Tetris</em> (Bullet-Proof Software), above</td>
</tr>
</tbody>
</table>

## Reliance on power-on mapper state {#Reliance_on_power-on_mapper_state}

| Game | Problem |
|----|----|
| *Bad News Baseball* and *Gekitou Stadium!!* | The intro sequence is started 1 frame before the [MMC1](../MMC1/) CHR bank registers are initialized. Depending on the initial CHR bank values, the [sprite 0 hit](../PPU_OAM/) detection can fail and cause the sequence to hang (with the music still playing). |
| *Digital Devil Story: Megami Tensei II* | When powered-on with uninitialized SRAM, a "DDS II" logo is flashed from left to right before proceeding with the normal introduction. The logo's background looks garbled because the game has not fully initialized the [CHR bank select registers](../INES_Mapper_019/#CHR_and_NT_Select_.28.248000-.24DFFF.29_w) at that point. |
| *Mega Man 5* | Neglects to [disable MMC3 IRQs](../MMC3/#IRQ_disable_.28.24E000-.24FFFE.2C_even.29) on reset and executes a CLI instruction before the game has fully initialized, causing the game to not boot if an [IRQ](../IRQ/) was pending at that point. |
| *Police Academy: The Video Game* (1990, 1991 Prototypes) | Neglects to initialize the [RAMBO-1](../RAMBO-1/)'s banking registers before calling sound driver routines in the \$C000 area. |
| *Populous* (Prototype) | Neglects to initialize the [MMC1 control](../MMC1/#Control_.28internal.2C_.248000-.249FFF.29) and [CHR bank 0](../MMC1/#CHR_bank_0_.28internal.2C_.24A000-.24BFFF.29) registers. CHR-RAM contents will be filled with incorrect data at power-on if CHR bank mode = 1 (switch two separate 4 KB banks) and CHR bank 0 = 0. |

## Reliance on RAM values {#Reliance_on_RAM_values}

Several games erroneously rely on RAM areas being pre-populated with certain values at power-on, despite RAM contents not being in a consistent state on power-on. Other games rely on similar values, but in PRG-RAM (WRAM), or CHR-RAM.

Note that using power-on RAM content as a seed for *random* number generation (such as in Final Fantasy) is not a game bug, even if it makes speedruns harder to verify on console.

| Game | Problem |
|----|----|
| *Cheetahmen II* | Suspected that certain RAM values may affect being able to reach the final two levels of the game (levels 5 and 6).<sup><span id="cite_ref-15"></span>[\[15\]](#cite_note-15)</sup> |
| *Chinese Kungfu: 少林武者* | The game will not display the first self-running demo correctly if \$0707 contains the value \$FF at startup. Values \$00-\$09 will cause one of nine self-running demo sequences to play first, while values above that will cause the game to always begin with the first demo sequence. For the Western localization (*Challenge of the Dragon*, not to be confused with the Color Dreams game of the same name), the developers seemed to have noticed this problem and went out of their way to initialize this memory location with \$00. |
| *Cybernoid* | On the title screen, the default selection for the difficulty level changes based on the contents of RAM at power on. Also, the music may not start when starting a game (depending on uninitialized RAM values).<sup><span id="cite_ref-16"></span>[\[16\]](#cite_note-16)</sup> |
| *Dancing Blocks* (Sachen) | Game will not boot when addresses \$EC and \$ED are both set to \$FF.<sup><span id="cite_ref-17"></span>[\[17\]](#cite_note-17)</sup> |
| *Erika to Satoru no Yume Bouken* | Plays uninitialized sound RAM as a waveform on the title screen, resulting in a buzzy tone on some power-ons<sup><span id="cite_ref-18"></span>[\[18\]](#cite_note-18)</sup>. |
| *F-1 Race* (Beta Version only) | Title screen will be skipped if \$6B and \$70 contain non-zero values.<sup><span id="cite_ref-19"></span>[\[19\]](#cite_note-19)</sup> Game blindly reads and uses values from \$51, \$55, \$70, \$A4, and \$0200-02FF (via sprite DMA).<sup><span id="cite_ref-20"></span>[\[20\]](#cite_note-20)</sup> |
| *Famicom Jump II: 最強の7人* | If save RAM is initialized to all 00s, the game will freeze at the very first power-on with a black screen. It will work normally after a soft reset as well as subsequent power-ons. |
| *Famicom Wars* | Blindly uses many uninitialized zeropage variables during startup. In particular, the value in \$FF can potentially enable [PPU EXT out mode](../PPU_registers/#Master.2Fslave_mode_and_the_EXT_pins) (PPUCTRL bit 6 set). |
| *Gun.Smoke* (FDS version only) | The music player's RAM is not cleared before starting the title screen song, resulting in a garbage first noise channel note, with random properties, if that RAM is not zero. |
| *Huang Di* | Uses uninitialized RAM at \$0100 to determine if Cheat Mode is enabled or not. When it's zero, cheat mode is enabled, allowing infinite jumps in midair. |
| *Keroppi to Keroriinu no Splash Bomb!* | Can crash during the final boss if any of \$0680-\$069F contain any value between \$30-\$33. This is because the game's playfield data begins 2 rows down from the top of the screen, so when checking for playfield collision, the position of the falling fire is adjusted by \$20 and underflows at the top of the screen, using uninitialized RAM after the playfield for these non-existent rows. If the value indicates the fire can land there, the game uses an invalid jump table index and crashes. |
| *Layla* | The music player's RAM is not cleared before starting the title screen song, resulting in a flubbed first note, sometimes with a frequency sweep, depending on power-on RAM content. |
| *Minna no Taabou no Nakayoshi Daisakusen* | Requires that address \$11 be initialized to a value other than \$00, otherwise the game will not start.<sup><span id="cite_ref-21"></span>[\[21\]](#cite_note-21)</sup> A PRG1 version corrects the issue. |
| *Silva Saga* | When save RAM is initialized with all 0s, the game incorrectly creates 3 blank saved games which do not work properly.<sup><span id="cite_ref-22"></span>[\[22\]](#cite_note-22)</sup> |
| *Super Mario Bros.* | Blindly reads and uses variables used for the continue feature's starting world number (accessed by pressing A+Start on the title screen) and the hard mode/second quest toggle on warm starts. A cartridge swap involving *Tennis* is often used to access "glitch worlds" beyond world 8 as a result of this.<sup><span id="cite_ref-23"></span>[\[23\]](#cite_note-23)</sup> |
| *Super Mario Bros.* (bootleg versions) | Relies on portions of RAM containing \$00, otherwise player starts at world 0-1.<sup><span id="cite_ref-24"></span>[\[24\]](#cite_note-24)</sup> |
| *Terminator 2: Judgment Day* | The copyright screen is skipped if RAM is filled with \$00 (more generally, if a high score table checksum happens to be valid).<sup><span id="cite_ref-25"></span>[\[25\]](#cite_note-25)</sup> |
| *Ultima: Exodus* | Relies on PRG-RAM contents before they're initialised; a fresh/new game may see artifacts such as shaking/wobbly text during the initial text intro screens, corruption of text intro screen fonts, and possibly audio anomalies.<sup><span id="cite_ref-26"></span>[\[26\]](#cite_note-26)</sup> |

## <span id=".22Impossible.22_controller_input"></span>"Impossible" controller input {#.22Impossible.22_controller_input}

Many games behave oddly when button combinations that would be impossible (or at least very hard) to input on a [standard controller](../Standard_controller/) are pressed. This comprises pressing left+right and up+down simultaneously. Such impossible controller input should probably be prevented by default in an emulator, but they can be optionally allowed for experimental purposes.

Alternatively, [controller reading routines](../Controller_reading_code/) in most games combine inputs from both the standard ports and the Famicom's [expansion port](../Expansion_port/) (even in international versions) to support [replacement controllers](../Four_player_adapters/#Adapters), so it is possible to press one direction in a standard controller and simultaneously the opposing direction in the replacement controller to achieve the same effect.

| Game | Problem |
|----|----|
| *Battletoads* | Pressing up+down in the vertical tunnel level kills the player instantly. Additionally, pressing left+right causes the player to walk up/back even when in pure 2D stages, which can result in certain areas becoming impossible to complete. |
| *Mega Man 1* and *2* | By pressing up+down at the top of a ladder, one may enter the "climbing ladder" state briefly above the top of the ladder. This allows "zipping" through walls. |
| *Panic Restaurant* | Pressing up while crouching (by pressing down, thus pressing up+down simultaneously) the player character's sprite uses garbage data including the damage sprite. This does not occur if up is pressed before down; the player chef merely stands still. |
| *Spy vs. Spy* | The character turns into an airplane and other garbage sprites when pressing left+right or up+down. |
| *Super Mario Bros. 2* | Ladders can be climbed very fast by pressing up+down. |
| *Teenage Mutant Ninja Turtles* | If you press the attack button while you jump while pressing up+down, the player character uses a garbage sprite, and attacks will use unusual (i.e. garbage data) hitboxes. |
| *Teenage Mutant Ninja Turtles II: The Arcade Game* | When you jump with a left or right move while pressing up+down, the player character will move in unusual directions and speeds, possibly screen-wrapping. |
| *Tetris* (Nintendo) | Holding left+down+right will cause the current and next piece sprites to flicker, slowing the game down in the process. |
| *Tiny Toon Adventures* | The player can gain unusual speed when pressing left+right. |
| *Zelda II: The Adventure of Link* | Link can gain tremendous speed when pressing left+right. |

### <span id="Games_with_intentional_.22impossible.22_controller_input"></span>Games with intentional "impossible" controller input {#Games_with_intentional_.22impossible.22_controller_input}

These games are intentionally programmed with special effects on "impossible" controller inputs, usually as a debug or cheat code.

| Game | Effects |
|----|----|
| *Bad Street Brawler* | A ~~Power Glove~~ Gaming Series game which maps a unique attack to left+right. It instantly kills one enemy per stage. |
| *Predator* | Pressing left+right+A/B in normal levels, or up+down+A/B in the level number screen before 'big mode' levels, activates a level skip. |
| *Star Soldier* | Cheat code button combinations include left+right and left+right+up+down inputs, which rely on the fact that both controller inputs are ORed bitwise when polling.<sup><span id="cite_ref-27"></span>[\[27\]](#cite_note-27)</sup> |
| *Star Wars* (Namco) | Level select code requires pressing up+down on the title screen. |
| Codemasters games | Up+down+left+right+A+B+Select+Start activates a checksum check. |

## Overscan ugliness {#Overscan_ugliness}

Some games display junk tiles in the [overscan](../Overscan/) area, which is usually not seen (or is at least partially obstructed) on real TV sets. Examples include the NTSC versions of *Metal Gear* (e.g. in the jungle area when gameplay starts) and *Solstice* (on the title screen).

- [<img src="../../wiki-images/Metal-gear-junk-tiles.png" width="128" height="120" alt="Metal Gear" />](../../wiki-images/Metal-gear-junk-tiles.png)
  *Metal Gear*
- [<img src="../../wiki-images/Solstice-junk-tiles.png" width="128" height="120" alt="Solstice – The Quest for the Staff of Demnos" />](../../wiki-images/Solstice-junk-tiles.png)
  *Solstice – The Quest for the Staff of Demnos*

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) <https://forums.nesdev.org/viewtopic.php?t=16923>
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) ["The Jumpy Scrolling and Graphical Glitches of NES Commando - Behind the Code" by Displaced Gamers](https://www.youtube.com/watch?v=OauZmS0Klxw) (30 minutes)
3.  <span id="cite_note-3"></span>[↑](#cite_ref-3) <https://forums.nesdev.org/viewtopic.php?t=7273>
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) <https://forums.nesdev.org/viewtopic.php?p=296408#p296408>
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) [Eliminator Boat Duel](https://forums.nesdev.org/viewtopic.php?t=18278)
6.  <span id="cite_note-6"></span>[↑](#cite_ref-6) <https://forums.nesdev.org/viewtopic.php?p=307745#p307745>
7.  <span id="cite_note-7"></span>[↑](#cite_ref-7) [Zelda Screen Transitions are Undefined Behaviour](https://www.gridbugs.org/zelda-screen-transitions-are-undefined-behaviour/): A deep dive into Zelda scrolling that is largely good until the "Scroll Down to Scroll Up" section, which has significant factual errors.
8.  <span id="cite_note-8"></span>[↑](#cite_ref-8) ["オセロ　LEVEL３がバグる　ファミコン" ("Othello level 3 CPU bugged on Famicom") by Hiroshi Tsuchiya](https://youtu.be/28QRWfkvN7E) (Japanese, 8 minutes)
9.  <span id="cite_note-9"></span>[↑](#cite_ref-9) <http://forums.nesdev.org/viewtopic.php?p=154894#p154894>
10. <span id="cite_note-10"></span>[↑](#cite_ref-10) ["Fixing StarTropics" by Brad Smith](https://www.youtube.com/watch?v=EmUvGH2HMuc) (5 minutes)
11. <span id="cite_note-11"></span>[↑](#cite_ref-11) ["The Garbage Sprites in Strider (NES) - Behind the Code" by Displaced Gamers](https://www.youtube.com/watch?v=01aBYq91KnA) (15 minutes)
12. <span id="cite_note-12"></span>[↑](#cite_ref-12) <https://forums.nesdev.org/viewtopic.php?p=112424#p112424>
13. <span id="cite_note-13"></span>[↑](#cite_ref-13) <http://forums.nesdev.org/viewtopic.php?f=10&t=7844>
14. <span id="cite_note-14"></span>[↑](#cite_ref-14) [Topple Zip accidentally relies on unreadable OAMDATA](https://forums.nesdev.org/viewtopic.php?p=304797#p304797)
15. <span id="cite_note-15"></span>[↑](#cite_ref-15) <http://forums.nesdev.org/viewtopic.php?p=178154#p178154>
16. <span id="cite_note-16"></span>[↑](#cite_ref-16) <http://tasvideos.org/forum/viewtopic.php?p=463659&sid=7cfe55942ee3420275d8f16b2a59770a#463659>
17. <span id="cite_note-17"></span>[↑](#cite_ref-17) <http://tasvideos.org/forum/viewtopic.php?p=463659&sid=7cfe55942ee3420275d8f16b2a59770a#463659>
18. <span id="cite_note-18"></span>[↑](#cite_ref-18) [N163 Sound RAM Initialisation Behaviour](https://forums.nesdev.org/viewtopic.php?p=285135)
19. <span id="cite_note-19"></span>[↑](#cite_ref-19) <http://forums.nesdev.org/viewtopic.php?p=178168#p178168>
20. <span id="cite_note-20"></span>[↑](#cite_ref-20) <http://forums.nesdev.org/viewtopic.php?p=178015#p178015>
21. <span id="cite_note-21"></span>[↑](#cite_ref-21) <http://forums.nesdev.org/viewtopic.php?p=185663#p185663>
22. <span id="cite_note-22"></span>[↑](#cite_ref-22) <https://forums.nesdev.org/viewtopic.php?f=11&t=11045>
23. <span id="cite_note-23"></span>[↑](#cite_ref-23) ["Access Glitch Worlds in Super Mario Bros. via NES Tennis" by Retro Game Mechanics Explained](https://www.youtube.com/watch?v=hrFHNgJlJSg) (13 minutes)
24. <span id="cite_note-24"></span>[↑](#cite_ref-24) <http://forums.nesdev.org/viewtopic.php?p=178163#p178163>
25. <span id="cite_note-25"></span>[↑](#cite_ref-25) <http://forums.nesdev.org/viewtopic.php?p=180752#p180752>
26. <span id="cite_note-26"></span>[↑](#cite_ref-26) <http://forums.nesdev.org/viewtopic.php?p=185163#p185163>
27. <span id="cite_note-27"></span>[↑](#cite_ref-27) [Star Soldier cheat codes](https://taotao54321.github.io/StarSoldierResource/cheat/) (Japanese)
