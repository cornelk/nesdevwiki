---
title: "Music"
url: "/wiki/Music/"
description: "Music is available for the NES. It is most often stored in NSF files, which package the data for one or more songs along with 6502 machine code to play it back. NSF files can be played on a NES or on…"
wikiPage: "Music"
wikiCategories:
  - name: "Audio"
    url: "/category/Audio/"
---

Music is available for the NES. It is most often stored in NSF files, which package the data for one or more songs along with 6502 machine code to play it back. NSF files can be played on a NES or on a PC with an appropriate player which emulates the NES sound hardware and CPU.

The music from a large number of NES games has been "ripped" into NSF format, and people are actively creating [new NES music](http://www.2a03.org/) with sequencing tools.

Refer to the [Music tools](../Tools/#Music_tools) wiki page for details.

Some of these general-purpose music engines are designed for making NSFs where the only thing executing is the music player. This means they can be fairly large, and in an [NROM](../NROM/) game, their features might not justify their size.

If you want to write your own music engine that targets the NES [APU](../APU/), you'll first need a lookup table for the note periods. This means you'll need the frequencies that correspond to various pitches. See [APU period table](../APU_period_table/), [Wikipedia:Pitch (music)](https://en.wikipedia.org/wiki/Pitch_(music)), [Wikipedia:Semitone](https://en.wikipedia.org/wiki/Semitone), and [Wikipedia:Equal temperament](https://en.wikipedia.org/wiki/Equal_temperament).

## NES Audio {#NES_Audio}

- [NES APU](../APU/)
