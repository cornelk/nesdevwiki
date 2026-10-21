---
title: "Bandai Hyper Shot"
url: "/wiki/Bandai_Hyper_Shot/"
description: "The Bandai Hyper Shot is a Famicom lightgun peripheral released by Bandai alongside the pack-in game Space Shadow. It combines a Zapper and standard controller into a single expansion port device…"
wikiPage: "Bandai_Hyper_Shot"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Bandai Hyper Shot is a Famicom lightgun peripheral released by Bandai alongside the pack-in game *Space Shadow*. It combines a [Zapper](../Zapper/) and [standard controller](../Standard_controller/) into a single [expansion port](../Expansion_port/) device shaped like a machine gun. The device also mounts a rumble/vibration system and a speaker.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xABS
          |||
          ||+- Controller shift register strobe
          |+-- Body Vibration System/Gun Move (0: off; 1: on)
          +--- Sound System (0: off; 1: on)

This controls the device's vibration system (the only component powered by batteries) and speaker output (routed from the expansion port) in addition to the controller strobe.

### <span id="Controller_report_.28.244016_read.29"></span>Controller report (\$4016 read) {#Controller_report_.28.244016_read.29}

    7  bit  0
    ---- ----
    xxxx xxEx
           |
           +-- Embedded controller status bit

This is a standard controller report through joypad 1 /D1, which is required by *Space Shadow* for ingame actions (e.g. B/sub-trigger to throw grenade). Note that the A button can never be registered as pressed by this device despite the box claiming that the main trigger is mapped to it.

### <span id="Lightgun_report_.28.244017_read.29"></span>Lightgun report (\$4017 read) {#Lightgun_report_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxT Wxxx
       | |   
       | +---- Light sensed at the current scanline (0: detected; 1: not detected)
       +------ Trigger (0: released; 1: fully pulled)

This is identical to the standard Zapper report *except* for the trigger status; the Bandai Hyper Shot registers 1 when the trigger is *fully pulled* in order to support continuous fire by holding down the trigger. *Space Shadow* and other lightgun games by Bandai appear to rely on this trigger behavior.

Normal Zapper games which expect the half-pulled behavior (0 -\> 1 -\> 0 transition) will register the trigger for this device when *releasing* the trigger instead.

## References {#References}

- [nocash's documentation](https://problemkaputt.de/everynes.htm#controllerslightgunszapper)
