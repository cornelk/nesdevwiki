---
title: "Battle Box"
url: "/wiki/Battle_Box/"
description: "The Battle Box is a Famicom expansion port device holding 512 bytes of EEPROM."
wikiPage: "Battle_Box"
wikiCategories:
  - name: "Hardware"
    url: "/category/Hardware/"
---

The Battle Box is a Famicom [expansion port](../Expansion_port/) device holding 512 bytes of EEPROM.

The TC89102 EEPROMs used are similar to standard microwire EEPROMs, except that the address is entered before the command. Each are configured for 7 bits of address and 16 bits of data.

Supported games:

- *Battle Stadium - Senbatsu Pro Yakyuu* (1 team slot)
- *Armadillo* (4 save slots)
- *J-League Fighting Soccer - The King of Ace Strikers* (2 team slots)
- *Seiryaku Simulation - Inbou no Wakusei - Shancara* (1 auto-save slot, manual states it will overwrite other files)

## Hardware interface {#Hardware_interface}

### Input {#Input}

#### <span id=".28.244017_Read.29"></span>(\$4017 Read) {#.28.244017_Read.29}

    7  bit  0
    ---- ----
    xxxW Rxxx
       | |
       | +---- see below
       +------ Data line to the two EEPROMs

The data line toggles after each time this is read, so the actual value is the opposite of what was last read. You cannot directly set the data output; you must read this port until it's the opposite of what you want.

#### <span id=".28.244016_Write.29"></span>(\$4016 Write) {#.28.244016_Write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Clock line to the two TC89102 EEPROMs

If "S" is high during a read from \$4017, it toggles which of two EEPROMs are enabled.

### Output {#Output}

#### <span id=".28.244017_Read.29_2"></span>(\$4017 Read) {#.28.244017_Read.29_2}

    7  bit  0
    ---- ----
    xxxW Rxxx
       | |
       | +---- bitwise NOT of the value read from the EEPROM
       +------ see above

## See also {#See_also}

- [Krzysiobal's reverse-engineered schematic and the TC89102 datasheet](https://forums.nesdev.org/viewtopic.php?t=22685)
- [Github Gist](https://gist.github.com/TakuikaNinja/9dfb280275690c33801a80376b8456f8) - Python script for listing files within Battle Box saves, with support for unscrambling data saved from Mesen2/CE.

## References {#References}

- <https://problemkaputt.de/everynes.htm#storagebattlebox>
