---
title: "Open Source Cartridge Reader"
url: "/wiki/Open_Source_Cartridge_Reader/"
description: "Sanni Open Source Cartridge Reader (OSCR) is a multi-consoles cartridge dumper/flasher device appeared in 2022."
wikiPage: "Open_Source_Cartridge_Reader"
wikiCategories:
  - name: "Dumper and Flasher devices"
    url: "/category/Dumper_and_Flasher_devices/"
---

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/OSCR.jpg"><img src="../../wiki-images/OSCR.jpg" width="300" height="255" /></a>
<figcaption>OSCR</figcaption>
</figure>

Sanni Open Source Cartridge Reader (OSCR) is a multi-consoles cartridge dumper/flasher device appeared in 2022.

It is an easy to build and easy to modify cartridge device, licensed under the GNU General Public License v3.0.

A turnkey version can be found on Marketplaces like Aliexpress, or on the [store](https://www.cartreader.net/).

## Hardware {#Hardware}

The device is a shield for the Arduino Mega that can back up video game cartridges.

Its main purpose is to dump a video game's ROM and save file to an SD card without the need of a computer.

Supported cartridges :

- NES / Famicom / Family Basic
- Super Nintendo / Super Famicom (including SF Memory, Satellaview, Sufami Turbo, and Game Processor RAM Cassettes)
- Nintendo 64 (including Controller Pak, Gameshark, and Xplorer 64)
- Game Boy Color (including GB Memory, Codebreaker, and Gameshark)
- Game Boy Advance
- Sega Mega Drive / Genesis
- Sega Master System

## Software {#Software}

Even if the device isn't connected to a computer, it compounds an internal software which allow an automatic mapper configuration.

It works by calculating two CRC32 checksums over 512 bytes of data each found in the PRG ROM of the NES cartridge.

The first at offset 0x8000 and the second at 0xE000 as seen by the Cart Reader with no mapper selected. The Cart Reader calculates checksums of two areas because by default some games have the first bank and some games have the last bank mapped while the other is just random data. The OSCR then compares these checksums to the second number in the nes.txt database.

## Related Links {#Related_Links}

- [OSCR's official repository](https://github.com/sanni/cartreader)

<!-- -->

- [OSCR's Wiki](https://github.com/sanni/cartreader/wiki)

<!-- -->

- [OSCR's store](https://www.cartreader.net/)
