---
title: "Super Famicom NTT Data Keypad"
url: "/wiki/Super_Famicom_NTT_Data_Keypad/"
description: "The Super Famicom NTT Data Keypad (NDK10) is intended to be used in conjunction with the Super Famicom NTT Data Communication Modem. It adds 15 buttons to an otherwise-normal SNES controller."
wikiPage: "Super_Famicom_NTT_Data_Keypad"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
  - name: "Super NES"
    url: "/category/Super_NES/"
---

The Super Famicom NTT Data Keypad (NDK10) is intended to be used in conjunction with the Super Famicom NTT Data Communication Modem. It adds 15 buttons to an otherwise-normal SNES controller.

## Protocol {#Protocol}

This controller extends the [SNES controller](../SNES_controller/) protocol to 32 bits to support the additional buttons. Bits 0-15 match the [SNES controller](../SNES_controller/) except with a different signature, while bits 16-31 match the last 16 bits of a [Famicom Network Controller](../Famicom_Network_Controller/#Protocol).

     0 - B
     1 - Y
     2 - ᐊ / 前ページ (Previous page)
     3 - ᐅ / 次ページ (Next page)
     4 - Up
     5 - Down
     6 - Left
     7 - Right
     8 - A
     9 - X
    10 - L
    11 - R
    12 - (Always 0)
    13 - (Always 1)
    14 - (Always 0)
    15 - (Always 0)
    16 - 0
    17 - 1
    18 - 2
    19 - 3
    20 - 4
    21 - 5
    22 - 6
    23 - 7
    24 - 8
    25 - 9
    26 - *
    27 - #
    28 - .
    29 - C
    30 - (Always 0)
    31 - 通信終了 (End communication)

    32+ - (Always 1)

## Links {#Links}

[NTT Data Keypad](https://snes.nesdev.org/wiki/NTT_Data_Keypad) at SNESdev Wiki
