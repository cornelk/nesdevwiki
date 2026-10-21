---
title: "Exciting Boxing Punching Bag"
url: "/wiki/Exciting_Boxing_Punching_Bag/"
description: "This punching bag/doll is used in a single game:"
wikiPage: "Exciting_Boxing_Punching_Bag"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

This punching bag/doll is used in a single game:

- *Exciting Boxing*

It contains 8 sensors, each returned as a single bit when reading \$4017:

- Left/Right Hook
- Left/Right Jab
- Left/Right Move
- Straight
- Body

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxAx
           |
           +- Select output sensors (see below)

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxE DCBx
       | |||
       | ||+-- Left Hook (A = 0), Left Jab (A = 1)
       | |+--- Move Right (A = 0), Body (A = 1)
       | +---- Move Left (A = 0), Right Jab (A = 1)
       +------ Right Hook (A = 0), Straight (A = 1)

## See Also {#See_Also}

- [image and English description](https://famicomworld.com/system/controllers/exciting-boxing/)
