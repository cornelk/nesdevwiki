---
title: "Barcode Battler"
url: "/wiki/Barcode_Battler/"
description: "The Barcode Battler is a handheld electronic game released by Epoch Co. which uses a barcode scanner to generate ingame characters and items. A successor called the Barcode Battler II would later be…"
wikiPage: "Barcode_Battler"
---

The **[Barcode Battler](https://en.wikipedia.org/wiki/Barcode_Battler)** is a handheld electronic game released by Epoch Co. which uses a barcode scanner to generate ingame characters and items. A successor called the **Barcode Battler II** would later be released in Japan with an output port to interface with video game consoles such as the Famicom and Super Famicom. (This is the version released internationally as the Barcode Battler)

A Famicom [expansion port](../Expansion_port/) cable was bundled with the game *Barcode World* in order to interface with the device, which is what this page will describe.

## <span id="Input_.28.244017_Read.29"></span>Input (\$4017 Read) {#Input_.28.244017_Read.29}

    7  bit  0
    ---- ----
    xxxx xDxx
          |
          +---- Serial data

The transmission format appears to be level-shifted RS232 serial at 1200 baud:

- 1 Start bit (must be 1=LOW)
- 8 Data bits (LSB first, 1=LOW=Zero, 0=HIGH=One)
- 1 Stop bit (must be 0=HIGH)

## Data format {#Data_format}

The data format is a 20-byte ASCII string consisting of:

- A 13-byte barcode formatted as:
  - [EAN-13](https://en.wikipedia.org/wiki/EAN-13)
  - [EAN-8](https://en.wikipedia.org/wiki/EAN-8) (*leading* space padding)
  - [UPC-A](https://en.wikipedia.org/wiki/UPC-A)? (unknown padding)
  - [UPC-E](https://en.wikipedia.org/wiki/UPC-E)? (unknown padding)
  - `ERROR` (*trailing* space padding, denotes scanning error)
- A 7-byte ID:
  - `EPOCH\r\n` - sent by hardware
  - `SUNSOFT` - alternate string accepted by *Barcode World*

## References {#References}

- [nocash's documentation](https://problemkaputt.de/everynes.htm#controllersbarcodebattlerbarcodereader) (covers FC & SFC protocols)
