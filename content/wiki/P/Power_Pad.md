---
title: "Power Pad"
url: "/wiki/Power_Pad/"
description: "The Power Pad (Family Fun Fitness in Europe) exercise mat for the Nintendo Entertainment System has twelve sensors arranged in a 3x4 grid, which the player steps on to control the action on the…"
wikiPage: "Power_Pad"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/NES-power-pad.jpg"><img src="../../wiki-images/NES-power-pad.jpg" width="300" height="300" /></a>
<figcaption>Side A of the NES Power Pad</figcaption>
</figure>

The **Power Pad** (**Family Fun Fitness** in Europe) exercise mat for the Nintendo Entertainment System has twelve sensors arranged in a 3x4 grid, which the player steps on to control the action on the screen. It can be used in either [controller port](../Controller_port_pinout/), though most games will only allow you to use it in controller port \#2, leaving port \#1 for a standard controller used for navigating through options.

The **Family Trainer Mat** is the equivalent exercise mat for the [Family Computer](../Family_Computer/), with the exception of its [input protocol](#Family_Trainer_Mat).

The **Tap-tap Mat** used for the game *Super Mogura Tataki!! - Pokkun Moguraa* uses the same protocol as the Family Trainer Mat.

## Pad Layouts {#Pad_Layouts}

The sensors on the Power Pad are spaced at 252 mm across by 285 mm front and back. Most games used side B, with the numbers on top. A few games turned the pad over to side A, whose markings lack numerals and lack markings for spaces 1, 4, 9, and 12 entirely (but they still send a signal).

              |                        |
    ,---------+---------.    ,---------+---------.
    | POWER PAD  side B |    | POWER PAD  side A |
    |  (1) (2) (3) (4)  |    |      (O) (O)      |
    |                   |    |                   |
    |  (5) (6) (7) (8)  |    |  (O) (X) (X) (O)  |
    |                   |    |                   |
    |  (9) (10)(11)(12) |    |      (O) (O)      |
    |                   |    |                   |
    `-------------------'    `-------------------'

The Tap-tap Mat is smaller but uses a layout equivalent to side A of the Power Pad:

     _________.-------------._________
    |         |// Tap-tap   |  igs    |
    |         |/    Mat     |         |
    |---------'============='---------|
    |  .---.   .---.   .---.   .---.  |   <--- Tap-tap numbering is identical
    | (dragn) (BANG!) (croco) (BANG!) |        to SIDE A of the Power Pad, i.e.
    |  '---'   '---'   '---'   '---'  |                 4  3  2  1
    |  .---.   .---.   .---.   .---.  |                 8  7  6  5
    | (BANG!) (monky) (BANG!) (chick) |                12 11 10  9
    |  '---'   '---'   '---'   '---'  |
    |  .---.   .---.   .---.   .---.  |
    | (frank) (BANG!) (shark) (BANG!) |
    |  '---'   '---'   '---'   '---'  |
    |_________________________________|

There is a third possible configuration, which no official game used, but which may be useful for homebrew dance simulation games in the style of *Dance Dance Revolution*: side B rotated 90 degrees anticlockwise, placing sensors 4, 8, and 12 toward the display.

    | ,---------------.           | ,---------------.
    | |  SIDE B       |           | |  SIDE DDR     |
    | |  (4) (8) (12) |           | | (Sel)    (St) |
    | |               |           | |               |
    | |  (3) (7) (11) |    ____   | |  (X) (U) (O)  |
    `-+               |    ____   `-+               |
      |  (2) (6) (10) |             |  (L)     (R)  |
      |               |             |               |
      |  (1) (5) (9)  |             |      (D)      |
      |               |             |               |
      `---------------'             `---------------'

## Pinout {#Pinout}

The Power Pad uses pins D3 and D4.

            .-               
     GND -- |O\             
     CLK <- |OO\ -- +5V    
     OUT <- |OO| <- D3
      D0 xx |OO| <- D4     
            '--'          

         NES port          

     * Directions are relative to the jack on the from of console

Standard wire colors:

| Signal | Power Pad  |
|--------|------------|
| +5V    | **White**  |
| OUT    | **Orange** |
| D3     | **Blue**   |
| D4     | **Yellow** |
| GND    | **Brown**  |
| CLK    | **Red**    |

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Controller shift register strobe

Writing 1 to this bit will record the state of each button. Writing 0 afterwards will allow the buttons to be read back, two at a time.

### <span id="Output_.28.244016.2F.244017_read.29"></span>Output (\$4016/\$4017 read) {#Output_.28.244016.2F.244017_read.29}

    7  bit  0
    ---- ----
    xxxH Lxxx
       | |
       | +---- Serial data from buttons 2, 1, 5, 9, 6, 10, 11, 7
       +------ Serial data from buttons 4, 3, 12, 8 (following 4 bits read as H=1)

The first 8 reads will indicate which buttons are pressed (1 if pressed, 0 if not pressed); all subsequent reads will return H=1 and L=1.

Remember to save *both* bits that you get from each read. If you're playing [samples](../APU_DMC/), beware of [bit deletions](../Standard_controller/#Evil_Details).

## Family Trainer Mat {#Family_Trainer_Mat}

The Family Trainer Mat looks similar to the Power Pad, but has an entirely different protocol that took advantage of the greater number of digital outputs on the [Famicom expansion port](../Expansion_port/). The Tap-tap Mat uses the same interface as the Family Trainer Mat.

### <span id="Input_.28.244016_write.29_2"></span>Input (\$4016 write) {#Input_.28.244016_write.29_2}

    7  bit  0
    ---- ----
    xxxx xABC
          |||
          ||+-- 1: Ignore bottom row, 0: include bottom row (buttons 9-12)
          |+--- 1: Ignore middle row, 0: include middle row (buttons 5-8)
          +---- 1: Ignore top row, 0: include top row (buttons 1-4)

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxD EFGx
       | |||
       | ||+-- right-most column (buttons 4,8,12)
       | |+--- right-center column (buttons 3,7,11)
       | +---- left-center column (buttons 2,6,10)
       +------ left-most column (buttons 1,5,9)

If a button is pressed in the currently selected rows, then the output returns 0. If no button is pressed, the output for that column is 1.

The matrix must be manually scanned in software; additionally it seems that there are quite large parasitic capacitances on the Family Trainer Mat which may require waiting a whole millisecond before going on to the next row.

The Family Trainer Mat does have the needed diodes for n-key-rollover. [\[1\]](https://forums.nesdev.org/viewtopic.php?t=24649)

## Programs {#Programs}

- [Controller test](https://forums.nesdev.org/viewtopic.php?p=184266#p184266) by Damian Yerrick (Supports NES version)
- [powerpadgesture](https://forums.nesdev.org/viewtopic.php?t=24192) by Damian Yerrick
- [powerpd.zip](https://www.nesdev.org/powerpd.zip) by Tennessee Carmel-Veilleux, 1999

## References {#References}

- [powerpad.txt](https://www.nesdev.org/powerpad.txt) - documentation by Tennessee Carmel-Veilleux, 2000
- <https://problemkaputt.de/everynes.htm#controllersmats> - documentation by nocash
