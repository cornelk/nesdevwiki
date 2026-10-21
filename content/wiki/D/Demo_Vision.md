---
title: "Demo Vision"
url: "/wiki/Demo_Vision/"
description: "The Demo Vision, also branded as Demo Boy II, is a specialized NES console used in Game Boy kiosks. It does not have a cartridge slot nor controllers. Instead, an attached Game Boy supplies video and…"
wikiPage: "Demo_Vision"
---

The Demo Vision, also branded as **Demo Boy II**, is a specialized NES console used in Game Boy kiosks. It does not have a cartridge slot nor controllers. Instead, an attached Game Boy supplies video and audio to the system.

This system has not yet been thoroughly reverse engineered.

## Design {#Design}

The Demo Vision uses two pairs of standard, 8 KiB CHR-RAM (32 KiB total) to hold Game Boy screen data. Each pair is used as a screen buffer. The active pair is read by the PPU to render the screen, while Game Boy screen data is written into the inactive pair. Their roles can be swapped at any time by firmware software running on the NES CPU. An additional 8 KiB of CHR-ROM is present, used to hold the screen border graphics.

An MMC5 or MMC5A mapper is built into the system, allowing any of the active 24 KiB of CHR data to be used anywhere on the screen. The active screen buffer is available as the first 16 KiB of CHR, followed immediately by the CHR-ROM. The inactive screen buffer is not mapped.

The firmware is contained in 8 KiB of PRG-ROM. This software sets up the mapper, nametables, attributes, and palettes for displaying the Game Boy screen data and border, and is responsible for controlling /reset in the Game Boy unit, which it enables after completing setup. The firmware is also responsible for switching screen buffers. This is done in an interrupt handler, but rather than using NMIs at the NES frame rate, an IRQ at the Game Boy frame rate is used, ensuring the active buffer always contains a complete Game Boy frame. However, there is still visible tearing because the swap can happen at any time in the NES frame.

The CPU address space features 8 KiB of RAM at \$0000-1FFF, rather than the standard 2 KiB.

There are 6 DIP switches available for software use. The firmware uses switches 1-3 for a gameplay timer. If configured with a time, the firmware will reset the Game Boy after the timer expires by toggling its /reset input. The other 3 are unused.

Demo Vision systems were available in 6 variations: domestic, Europe, and UK, each in single and dual varieties. Domestic versions use an NTSC CPU and PPU and Europe versions use a PAL CPU and PPU. The UK variant has not been documented. The dual systems are believed to use two identical Demo Vision boards with no logical connections between them except some kind of shared 2-pin reset cable.

While initially available with just the original Game Boy, kiosks could later be used with Game Boy Pockets. Functionally, these differ in the behavior of the handheld's power switch. The power switch on the Game Boy controls power for the handheld as normal, while the Game Boy Pocket's switch controls power for just the screen, allowing the kiosk to continue to function with the handheld's switch in the off position.

## Registers {#Registers}

### <span id=".244016_write"></span>\$4016 write {#.244016_write}

    7  bit  0
    ---- ----
    .... ..BP
           ||
           |+- Game Boy /reset (0 = hold in reset, 1 = run)
           +-- Active screen buffer

- Game Boy /reset should be set to 1 once the system is configured and ready. The Game Boy can be reset by clearing this to 0 for at least 2 PPU frames and then setting it back to 1.
- Active screen buffer swaps which 16 KiB of CHR-RAM is accessible to the PPU as the first 16 KiB of CHR and which is being filled with Game Boy screen data. It is not known if the active CHR-RAM can be written by the PPU.

### <span id=".244016_read"></span>\$4016 read {#.244016_read}

    7  bit  0
    ---- ----
    ..65 4321
      || ||||
      ++-++++- DIP switches 6-1 (0 = off/up, 1 = on/down)

## IRQ {#IRQ}

The CPU /IRQ is believed to be asserted at some point during every Game Boy vblank. It is asserted for a limited time and does not require an ack. The firmware software's IRQ handler includes a 1281 cycle wait loop, and the handler takes 1322 cycles in total, which should be taken as an upper bound on the IRQ length.

## DIP switches {#DIP_switches}

There are 6 DIP switches. They are software-readable and have no other hardware function. They read as 0 in the off/up position and 1 in the on/down position.

### Software {#Software}

- 1: Gameplay time
- 2: Gameplay time
- 3: Gameplay time
- 4: Unused
- 5: Unused
- 6: Unused

<table class="wikitable">
<thead>
<tr>
<th data-valign="top">SW3</th>
<th data-valign="top">SW2</th>
<th data-valign="top">SW1</th>
<th colspan="2">Gameplay Time</th>
</tr>
</thead>
<tbody>
<tr>
<th>Demo Boy II</th>
<th>Demo Vision</th>
<th></th>
<th></th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>0</td>
<td>0</td>
<td>1 minute</td>
<td>1 minute</td>
</tr>
<tr>
<td>0</td>
<td>0</td>
<td>1</td>
<td>3 minutes</td>
<td>3 minutes</td>
</tr>
<tr>
<td>0</td>
<td>1</td>
<td>0</td>
<td>5 minutes</td>
<td>5 minutes</td>
</tr>
<tr>
<td>0</td>
<td>1</td>
<td>1</td>
<td>7 minutes</td>
<td>7 minutes</td>
</tr>
<tr>
<td>1</td>
<td>0</td>
<td>0</td>
<td>12 minutes</td>
<td>12 minutes</td>
</tr>
<tr>
<td>1</td>
<td>0</td>
<td>1</td>
<td>18374 seconds*</td>
<td>Unlimited</td>
</tr>
<tr>
<td>1</td>
<td>1</td>
<td>0</td>
<td>Unlimited</td>
<td>Unlimited</td>
</tr>
<tr>
<td>1</td>
<td>1</td>
<td>1</td>
<td>65481 seconds*</td>
<td>Unlimited</td>
</tr>
</tbody>
</table>

\*Note: The Demo Boy II revision A firmware allows the timer to run for switch values %101 and %111, but does not define timer values for these settings. Instead, the code following the timer value table is interpreted as timer values.

If the Demo Vision firmware detects that switches 1-3 have changed, it resets the Game Boy and sets the timer with the new time. The Demo Boy II rev A firmware only reads the DIP switch value on reset.

## Firmware {#Firmware}

### Power-on {#Power-on}

The Demo Boy II revision A firmware requires one or two resets to reliably enter stable operation, and Demo Boy II documentation instructs that 2 resets should be performed. This issue is fixed in the Demo Vision firmware. The Demo Vision firmware duplicates the reset code, running it twice with no other significant changes compared to the Demo Boy II. The reset code disables rendering, waits for at least 4 frames, sets up the palettes, nametables, attributes, MMC5, scroll, and timer, releases the Game Boy from reset, and enables rendering. It is not currently known why executing this twice prevents the problem.

The observed behavior is that onscreen graphics are corrupted, either displaying a blank border and garbage in the Game Boy screen area, or showing correct graphics but with nametable corruption exhibiting a checkboard pattern. It is suspected that MMC5 ExRAM is being initialized properly, but CIRAM is not.

### Unused code {#Unused_code}

The Demo Boy II revision A firmware contains unused code that is incompatible with the finalized hardware. A joypad-reading function is present that targets standard controllers, but keeps OUT1 set, suggesting that OUT1 may have originally controlled Game Boy reset. Another function reads Game Boy screen pixel data from \$8000-A3FF, separates it into bitplanes, and copies it into CHR-RAM, but in the finalized hardware, all of this is automated, there are two screen buffers, and the Game Boy screen data is not believed to be mapped into the CPU address space.

## Pinouts {#Pinouts}

### <span id="Power_connector_.28P1.29"></span>Power connector (P1) {#Power_connector_.28P1.29}

    Demo Vision
       +---+  
       | 2 | -- GND
       | 1 | -- +5V
       +---+

### <span id="Game_Boy_connector_.28P2.29"></span>Game Boy connector (P2) {#Game_Boy_connector_.28P2.29}

              Demo Vision
               +-------+
               |       |
          S -> |01   02| <- CPL
         ST -> |03   04| <- CP 
        LD1 ->  05   06  <- LD0
    speaker -> |07   08| -> /reset
        +5V -- |09   10| -- GND
               |       |
               +-------+

Pinout is as viewed into the port, pin 1 on the top right. Note that while this is keyed the same on both sides, the PCB partially blocks the keyhole on the bottom side, preventing the cable from being fully inserted upside-down.

#### <span id="Game_Boy_.28DMG.29_connections"></span>Game Boy (DMG) connections {#Game_Boy_.28DMG.29_connections}

    cable |  | Game Boy
    --------------------------
       01 |<-| P3.12 (S)
       02 |<-| P3.18 (CPL)
       03 |<-| P3.17 (ST)
       04 |<-| P3.14 (CP)
       05 |<-| P3.15 (LD1)
       06 |<-| P3.16 (LD0)
       07 |<-| P3.20 (speaker)
       08 |->| SW1.3 (/reset)
       09 |--| SW1.1 (+5V)
       10 |--| P3.01 (GND)

#### <span id="Game_Boy_Pocket_.28MGB.29_connections"></span>Game Boy Pocket (MGB) connections {#Game_Boy_Pocket_.28MGB.29_connections}

    cable |  | Game Boy Pocket
    --------------------------
       01 |<-| S
       02 |<-| CPL
       03 |<-| ST
       04 |<-| P2-CP
       05 |<-| LD1
       06 |<-| P2-LD0
       07 |<-| U3.3 (speaker)
       08 |->| P1.30 RES
       09 |--| SW1-VCC
       10 |--| P1.32 GND

## Hardware configurations {#Hardware_configurations}

Demo Boy II:

- Board:
  - Rev A (green PCB)
  - Rev B (green PCB)
- Chipsets:
  - RP2A03G / RP2C02G-0
  - RP2A07A / RP2C07-0 (Europe)

Demo Vision:

- Board:
  - Rev B (blue PCB)
- Chipsets:
  - RP2A03G / RP2C02G-0 (domestic)
  - RP2A07A / RP2C07-0 (Europe)

The Demo Boy II rev B and Demo Vision boards are suspected to be the same aside from color.
