---
title: ".pal"
url: "/wiki/.pal/"
description: ".pal is an extension used by palette files for NES emulators. Palette files are required for emulation of the RGB PPUs of the VS Unisystem, PlayChoice 10, and Famicom Titler systems. Before NTSC…"
wikiPage: "_pal"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
---

**.pal** is an extension used by palette files for NES emulators. Palette files are required for emulation of the RGB PPUs of the VS Unisystem, PlayChoice 10, and Famicom Titler systems. Before [NTSC video](../NTSC_video/) was fully understood, emulators also used palette files to translate the [PPU's color values](../PPU_palettes/) into sRGB values for the PC display.

Most palette files used by NES emulators are 192 bytes in size, consisting of 64 3-byte entries, one for each of 64 NES color values.

Each entry consists of three unsigned 8-bit bytes, in this order:

1.  red intensity
2.  green intensity
3.  blue intensity

Intensity is most commonly ranged from 0 (black) to 255 (full intensity), presuming the [sRGB](https://en.wikipedia.org/wiki/sRGB) color space used by most monitors.

Most emulators will ignore any entries past the first 64, but an optional extension of the .pal format (currently only known to be supported by Nintendulator 0.965 and later, and FCEUX 2.3.0) lengthens the file to 1536 bytes in order to store each palette value when colour (de-)emphasis bits are applied. This allows customization of the emphasis behaviour, and is particularly useful for representing the alternative emphasis behaviour of the RGB PPU used in the [Vs. System](../Vs._System/) and PlayChoice.

## See Also {#See_Also}

- [PPU palettes](../PPU_palettes/)

## External links {#External_links}

- [VGA Palette on ModdingWiki](http://www.shikadi.net/moddingwiki/VGA_Palette)
- [Editing custom NES palettes in Linux](https://forums.libretro.com/t/editing-custom-nes-palettes-in-linux/6733/6) instructions on editing .pal files in GIMP
