---
title: "M50805 pinout"
url: "/wiki/M50805_pinout/"
description: "Mitsubishi M50805: 22-pin 0.4\" PDIP (Used in one mapper 3 game)"
wikiPage: "M50805_pinout"
wikiCategories:
  - name: "Audio"
    url: "/category/Audio/"
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Mitsubishi M50805: 22-pin 0.4" PDIP (Used in [one](https://nescartdb.com/profile/view/3953/family-trainer-3-aerobics-studio) [mapper 3](../CNROM/#Notes) game)

                     .--\/--.
               NC -- |01  22| -- NC
               NC -- |02  21| -- NC
               NC -- |03  20| -- NC
        DREQ/TEST <- |04  19| <- /D0
      MODE \  MS1 -> |05  18| <- /D1
    SELECT /  MS0 -> |06  17| <- /D2,DTIN
             XOUT <- |07  16| -> BUSY
              XIN -> |08  15| <- /SYNC
              /CE -> |09  14| -> DA1 \ DA
           /RESET -> |10  13| -> DA0 / OUTPUT
              +5V -- |11  12| -- GND
                     '------'

Sources: [\[1\]](https://forums.nesdev.org/viewtopic.php?f=6&t=16129), [\[2\]](https://forums.nesdev.org/viewtopic.php?f=9&t=9449)

The [die shot](http://seanriddledecap.blogspot.com/2017/07/blog-post.html) shows that pins 1-3 and 20-22 are genuinely no connection - there's nothing to be wire bonded to.
