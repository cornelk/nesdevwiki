---
title: "HT-2880 pinout"
url: "/wiki/HT-2880_pinout/"
description: "HT-2880 contains 8 pre-programmed audio samples and was used in Casel Zapper (only KEY1 was utilized, as rifle sound)."
wikiPage: "HT_2880_pinout"
wikiCategories:
  - name: "Audio"
    url: "/category/Audio/"
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

**HT-2880** contains 8 pre-programmed audio samples and was used in Casel Zapper (only KEY1 was utilized, as rifle sound).

              .---v---.
       GND -- |01   18| <> TEST3
       +5V -> |02   17| <> TEST2
     TEST1 <> |03   16| <- LEVEL HOLD
     AUDIO <- |04   15| -- OSC2 --[R=150k]-+
    /AUDIO <- |05   14| -- OSC1 -----------+
      KEY8 -> |06   13| <- KEY1
      KEY7 -> |07   12| <- KEY2
      KEY6 -> |08   11| <- KEY3
      KEY5 -> |09   10| <- KEY4
              `-------`
               HT-2880

## References {#References}

- [What's goin on with the Casel Zapper?](https://forums.nesdev.org/viewtopic.php?t=18519) – forums
