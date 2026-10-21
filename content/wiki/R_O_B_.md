---
title: "R.O.B."
url: "/wiki/R.O.B./"
description: "R.O.B. (Robotic Operating Buddy) is a toy robot accessory for the Nintendo Entertainment System (NES). It was launched in July 1985 as the Family Computer Robot in Japan, and October 1985 as R.O.B.…"
wikiPage: "R_O_B_"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/NES-ROB.jpg"><img src="../../wiki-images/NES-ROB.jpg" width="200" height="253" /></a>
<figcaption>NES R.O.B.</figcaption>
</figure>

> **R.O.B.** (Robotic Operating Buddy) is a toy robot accessory for the Nintendo Entertainment System (NES). It was launched in July 1985 as the Family Computer Robot in Japan, and October 1985 as R.O.B. in North America. Its short lifespan yielded only two games in the Robot Series: *Gyromite* and *Stack-Up*.

> The R.O.B. unit's height is 24 cm (9.6 in). It has a head movement range of a 45° horizontally centered tilt. The arm movement range is 240° left and right with five stopping points, 7 cm (2.75 in) up and down with six stopping points, and 7 cm (2.75 in) between hands when open. The unit has five accessory slots around the hexagonal base, numbered clockwise, starting at the rear-left from the robot's point of view; and notches on the hands allow for specialized parts to be attached for each game. The tinted filter can be optionally attached over the eyes like sunglasses, to compensate for bright televisions or sunlight. The unit is powered by four AA batteries.

\- from [Wikipedia](https://en.wikipedia.org/wiki/R.O.B.)

R.O.B. operates by using light sensors to receive commands from supported games via screen flashes, and responds by slowly moving various parts to accomplish tasks such as:

- Spinning and placing tops to press the A/B buttons on controller 2 for in-game actions (*Gyromite*)
- Stacking circular blocks to solve puzzles provided in-game (*Stack-Up*)

## Interface {#Interface}

The protocol consists of a series of screen flashes, one bit per vertical sync. Black denotes a 0 bit, and a bright color denotes a 1 bit (official games use \$0F black and \$2A green respectively).

It takes 13 bits (or possibly 14—if Z is 1, the following bit must be 0<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup>) to send a command:

    000101w1x1y1z

Ten commands are known to exist:

| wxyz | Command |
|----|----|
| 0101 | Up 2 steps |
| 1100 | Up 1 step |
| 0010 | Down 1 step |
| 1101 | Down 2 steps |
| 0100 | Turns the body left |
| 1000 | Turns the body right |
| 0110 | Close the arms |
| 1010 | Open the arms |
| 1001 | Turn the head LED on permanently |
| 0001 | Reset = turn LED off, open arms, seek rightmost/upmost, then go back to center. |

The other six possible messages have no result.

Additionally, an endless stream of 25 or 30 Hz flashes will cause the LED to blink when R.O.B. sees light from the TV. This helps with aiming the head.

It is not known how wide R.O.B.'s field of view is, but the light sensor is the same as the [Zapper](../Zapper/)'s.

## See Also {#See_Also}

- [Controlling R.O.B. with a microcontroller](https://learn.adafruit.com/controlling-a-classic-nintendo-r-o-b-robot-using-circuit-playground-express?view=all)

## References {#References}

- [Nocash's EveryNES § R.O.B.](http://problemkaputt.de/everynes.htm#robroboticoperatingbuddy)
- [R.O.B. Firmware](../R.O.B._Firmware/) disassembly by nocash, based on a dump from [Sean Riddle](http://www.seanriddle.com/sm590/)

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) [AtariAge forum post](https://atariage.com/forums/topic/177286-any-interest-in-nes-rob-homebrews/?tab=comments#comment-2217299): Any interest in NES ROB homebrews?
