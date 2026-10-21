---
title: "French NES RGB output pinout"
url: "/wiki/French_NES_RGB_output_pinout/"
description: "Uniquely, the French NES contained a PAL-to-RGB conversion chip in the metal box and no RF nor composite output."
wikiPage: "French_NES_RGB_output_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

Uniquely, the French NES contained a PAL-to-RGB conversion chip in the metal box and no RF nor composite output.

This is the view of the back of the console:

                UP                DOWN
                      +-----\
    Composite Sync <- | F  A | -> Mono audio
              Blue <- | G  B | -- Ground
             Green <- | H  C | -- Ground
               Red <- | I  D | -- Ground
    SCART "SELECT" <- | J  E | -> SCART "STATUS"
                      +-----/

SCART pin 16 "Select" means the video is encoded as RGB. Here it's 5V via a 100Ω resistor.

SCART pin 8 "Status" high requests displaying the video from this connector. Here it's 12V via a half-wave rectifier, 1kΩ resistor, and shunt zener.
