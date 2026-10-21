---
title: "Partytap"
url: "/wiki/Partytap/"
description: "The Partytap is used in the following games:"
wikiPage: "Partytap"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Partytap is used in the following games:

- *Casino Derby*
- *Gimmi a Break - Shijou Saikyou no Quiz OuKetteisen*
- *Gimmi a Break - Shijou Saikyou no Quiz OuKetteisen 2*
- *Project Q*

It consists of 6 separate 1-button controllers (labelled buttons 1 to 6 below).

## Hardware interface {#Hardware_interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Controller shift register strobe

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxC BAxx
       | ||
       | |+- Serial data (Button 1, Button 4)
       | +-- Serial data (Button 2, Button 5)
       +---- Serial data (Button 3, Button 6)

The first read returns the state of buttons 1 to 3, the 2nd read gives buttons 4 to 6. The third read apparently returns a detection value (\$14).

The device supposedly requires a relatively large number of cycles between writes/reads (games wait up to 80-500 cycles in-between some operations).

## References {#References}

- [nocash's NES documentation](http://problemkaputt.de/everynes.htm#controllerspushbuttons)
- [Forums - Party Tap](https://forums.nesdev.org/viewtopic.php?t=903) (in Japanese)
