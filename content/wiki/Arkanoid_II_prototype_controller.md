---
title: "Arkanoid II prototype controller"
url: "/wiki/Arkanoid_II_prototype_controller/"
description: "The Arkanoid II prototype controller is a Famicom expansion port controller that was discovered through a prototype of Arkanoid II, which expects a different interface and behavior than that of the…"
wikiPage: "Arkanoid_II_prototype_controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Arkanoid II prototype controller is a Famicom expansion port controller that was discovered through a prototype of *Arkanoid II*, which expects a different interface and behavior than that of the production [Arkanoid controller](../Arkanoid_controller/). This behavior has been inferred purely from software.

## Summary {#Summary}

The prototype controller has the following known differences compared to the production controller:

- The analog control knob data and button data for both players are returned on different register bits.
- The control knob data is 8 bits instead of 9. Reading past 8 bits returns 1.

It is suspected the prototype controller also has these differences, but they don't matter for the prototype software and thus cannot be confirmed:

- Both players' controllers may be hardwired together rather than using an additional port for daisy chaining.
- The write interface may use OUT0 to latch the shift register and OUT1 to start a conversion, rather than simply using OUT0 to start a conversion and automatically latching when finished.

## Layout {#Layout}

The prototype controller has the same knob and button controls as the production controller, so it is expected to have a similar layout. It is not known if it contains potentiometers for calibration. It is suspected that the prototype may have had two controllers connected to the same expansion port plug, rather than the daisy chain design of the production Arkanoid II controller. This is because the prototype checks for the presence of controller 1 but not controller 2, while checking for both would make more sense in a daisy chain design.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

The write interface for the prototype controller is not known for certain, but it is compatible with the [OUT0 strobe interface](../Arkanoid_controller/#Input_.28.244016_write.29) used for production Arkanoid controllers and may function the same way. However, like the final version of the game, the prototype also strobes OUT1 after it is done reading both players' inputs, which normally has no effect. Although it is impossible to verify this without access to the controller itself, it is speculated that the prototype may use OUT0 to latch the control knob data into the shift register and OUT1 to start a conversion. In contrast, the production Arkanoid II controller works the same way as the original Arkanoid controller, using OUT0 to start a conversion and automatically loading the result into the shift register when the conversion completes.

### <span id="Output_.28.244016.2F.244017_read.29"></span>Output (\$4016/\$4017 read) {#Output_.28.244016.2F.244017_read.29}

    $4016 read:
    7  bit  0
    ---- ----
    xxxx xxDx
           |
           +- Controller 1 serial control knob data (8-bit, inverted, MSb first, followed by 1's)

    $4017 read:
    7  bit  0
    ---- ----
    xxxB BxDx
       | | |
       | | +- Controller 2 serial control knob data (8-bit, inverted, MSb first, followed by 1's)
       | +--- Controller 1 fire button
       +----- Controller 2 fire button

Notably, the prototype controller appears to only expose an 8-bit control knob conversion value, not 9-bit like the production version. The production controllers connect the least significant bit of the 9-bit conversion value to the shift register's serial input, while the prototype appears to connect this to ground, causing it to return 1. This is inferred from the prototype game replacing the control knob data with \$FF if bit 8 is not 1. This was likely done to detect whether the controller is plugged in; if it isn't, this bit is 0.

### Control knob data {#Control_knob_data}

It is not known what the control knob's value range is nor whether there are potentiometers for centering or scaling. However, the *Arkanoid II* prototype expects the following value ranges:

- 1P mode: \$26-\$B9
- VS player 1: \$4C-\$D3
- VS player 2: \$5D-\$E4

Values outside of these ranges are clamped, so the game is playable even if allowing the full 8-bit range.
