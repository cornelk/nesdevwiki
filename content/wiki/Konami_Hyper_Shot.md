---
title: "Konami Hyper Shot"
url: "/wiki/Konami_Hyper_Shot/"
description: "The Konami Hyper Shot is a pair of 2-button controllers released by Konami. It is required by the following games, both released in 1985:"
wikiPage: "Konami_Hyper_Shot"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Konami Hyper Shot is a pair of 2-button controllers released by Konami. It is required by the following games, both released in 1985:

- *Hyper Olympic*
- *Hyper Sports*

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xEFx
          ||
          |+- 0=Enable Player 1 Buttons
          +-- 0=Enable Player 2 Buttons

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxD CBAx
       | |||
       | ||+-- Player 1 Run
       | |+--- Player 1 Jump
       | +---- Player 2 Run
       +------ Player 2 Jump

The Run button is on the side of the controller closer to the cord. The Jump/Run bits for a player will always be 0 if the corresponding enable bit in \$4016 is set.

## See Also {#See_Also}

- [video of outside, teardown, and game play](https://youtu.be/6PSbI03aGKc)
- [SG-1000 version](https://segaretro.org/HyperShot)
