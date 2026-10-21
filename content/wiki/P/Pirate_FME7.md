---
title: "Pirate FME7"
url: "/wiki/Pirate_FME7/"
description: "This Pirate FME7 (used at least for game Batman, which has identical PRG \u0026 CHR compared to licensed version) is built entirely using discrete chips with addition of one PAL16L8 for generating clocks…"
wikiPage: "Pirate_FME7"
wikiCategories:
  - name: "Mappers"
    url: "/category/Mappers/"
---

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Pirate_FME7_components.jpg"><img src="../../wiki-images/Pirate_FME7_components.jpg" width="300" height="190" /></a>
<figcaption>Front with components</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Pirate_FME7_front.jpg"><img src="../../wiki-images/Pirate_FME7_front.jpg" width="300" height="200" /></a>
<figcaption>Front without components</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Pirate_FME7_back.jpg"><img src="../../wiki-images/Pirate_FME7_back.jpg" width="300" height="193" /></a>
<figcaption>Back</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/Pirate_FME7_scheme.png"><img src="../../wiki-images/Pirate_FME7_scheme.png" width="300" height="224" /></a>
</figure>

This Pirate FME7 (used at least for game Batman, which has identical PRG & CHR compared to licensed version) is built entirely using discrete chips with addition of one PAL16L8 for generating clocks and enable signals. It is almost the same like original [FME7](../Sunsoft_FME-7/) with the following exceptions:

- max PRG-ROM 128 kB (but can be increased up to 4MBit by adding one 74\*670),
- hardwired vertical mirroring,
- no WRAM at \$6000-\$7fff but PRG-ROM,
- IRQ enable & IRQ counter enable combined into one bit.

# Capabilities {#Capabilities}

- PRG-ROM: 128 kB
- CHR-ROM: 256 kB
- Bus conflicts: no
- Mirroring: vertical (hardwired)
- One presettable 16 bit counter which decreases at beginning of cpu cycle (rising edge of M2) and triggers interrupt when decreasing from \$0000 to \$ffff

# Memory map {#Memory_map}

- CPU \$6000-\$7fff: 8kB switchable PRG bank
- CPU \$8000-\$9fff: 8kB switchable PRG bank
- CPU \$a000-\$bfff: 8kB switchable PRG bank
- CPU \$c000-\$dfff: 8kB switchable PRG bank
- CPU \$e000-\$ffff: 8kB PRG bank, fixed to the last bank
- PPU \$0000-\$03ff: 1kB switchable CHR bank
- PPU \$0400-\$07ff: 1kB switchable CHR bank
- PPU \$0800-\$0bff: 1kB switchable CHR bank
- PPU \$0c00-\$0fff: 1kB switchable CHR bank
- PPU \$1000-\$13ff: 1kB switchable CHR bank
- PPU \$1400-\$17ff: 1kB switchable CHR bank
- PPU \$1800-\$1bff: 1kB switchable CHR bank
- PPU \$1c00-\$1fff: 1kB switchable CHR bank

# Registers {#Registers}

## Register select {#Register_select}

    [.... RRRR] $8000-$9fff: register select
          ||||
          ++++- select register to update at next write to $a000-$bfff (0-15)

## CHR registers {#CHR_registers}

    [CCCC CCCC] $a000-$bfff (when RRRR=$00)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $0000-$03ff

    [CCCC CCCC] $a000-$bfff (when RRRR=$01)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $0400-$07ff

    [CCCC CCCC] $a000-$bfff (when RRRR=$02)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $0800-$0bff

    [CCCC CCCC] $a000-$bfff (when RRRR=$03)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $0c00-$0fff

    [CCCC CCCC] $a000-$bfff (when RRRR=$04)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $1000-$13ff

    [CCCC CCCC] $a000-$bfff (when RRRR=$05)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $1400-$17ff

    [CCCC CCCC] $a000-$bfff (when RRRR=$06)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $1800-$1bff

    [CCCC CCCC] $a000-$bfff (when RRRR=$07)
     |||| ||||
     ++++-++++- select 1kb CHR-ROM bank at $1d00-$1fff

## PRG registers {#PRG_registers}

    [.... PPPP] $a000-$bfff (when RRRR=$08)
          ||||
          ++++- select 8kb PRG-ROM bank at $6000-$7fff

    [.... PPPP] $a000-$bfff (when RRRR=$09)
          ||||
          ++++- select 8kb PRG-ROM bank at $8000-$9fff

    [.... PPPP] $a000-$bfff (when RRRR=$0a)
          ||||
          ++++- select 8kb PRG-ROM bank at $a000-$bfff

    [.... PPPP] $a000-$bfff (when RRRR=$0b)
          ||||
          ++++- select 8kb PRG-ROM bank at $c000-$dfff

## IRQ and counter {#IRQ_and_counter}

    [.... ....] $a000-$bfff (when RRRR=$0c) - nothing

    [E... ....] $a000-$bfff (when RRRR=$0d)
     |        
     +--------- counter decrements at each CPU cycle (1:enabled, 0:disabled)
     +--------- irq acknowledgement (0: acknowledge)

    [LLLL LLLL] $a000-$bfff (when RRRR=$0e)
     |||| ||||
     ++++-++++- set low 8 bits of counter value

    [HHHH HHHH] $a000-$bfff (when RRRR=$0f)
     |||| ||||
     ++++-++++- set high 8 bits of counter value

# Trivia {#Trivia}

On bottom PCB layer there was one patch cable connecting IC3 pin3 (top third left) and IC4 pin 11 (fourth top left) which is desoldered on the photo.
