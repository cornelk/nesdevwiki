---
title: "Colour emphasis"
url: "/wiki/Colour_emphasis/"
description: "Colour emphasis on the NES is accomplished via bits 5, 6, and 7 of PPUMASK, for red, green, and blue respectively. These bits enable/disable an \"emphasis\" or \"tint\" effect visually on screen."
wikiPage: "Colour_emphasis"
---

Colour emphasis on the NES is accomplished via bits 5, 6, and 7 of [PPUMASK](../PPU_registers/), for red, green, and blue respectively. These bits enable/disable an "emphasis" or "tint" effect visually on screen.

- Bit 5 emphasizes red on the NTSC PPU, and green on the PAL & Dendy PPUs.
- Bit 6 emphasizes green on the NTSC PPU, and red on the PAL & Dendy PPUs.
- Bit 7 emphasizes blue on the NTSC, PAL, & Dendy PPUs.
- Each bit emphasizes 1 color while darkening the other two. Setting all three emphasis bits will darken colors \$00-\$0D, \$10-\$1D, \$20-\$2D, and \$30-\$3D (see [PPU palettes](../PPU_palettes/)). Note that \$1D black is affected by color emphasis, but \$0F black is not.
- The emphasis bits are applied independently of greyscale (bit 0 of [PPUMASK](../PPU_registers/)), so they will still tint the color of the grey image.

On Dendy and PAL PPUs, the purpose of bits 5 and 6 are reversed (red for green, and green for red).<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)<span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup>

See [NTSC video](../NTSC_video/) for a description of how bits 5-7 work on NTSC and PAL PPUs.

The [RGB PPU](../PPU_palettes/#2C03_and_2C05) used by the [Vs. System](../Vs._System/), PlayChoice-10, and some other systems treat the emphasis bits differently. Instead of darkening other RGB components, it forces one component to maximum brightness. [A few games](../Colour-emphasis_games/), which set all three tint bits to darken all colors, are unplayable on these PPUs because they instead turn the entire screen white.

## See also {#See_also}

- [NTSC video](../NTSC_video/)
- [PPU palettes](../PPU_palettes/)
- [PPU registers](../PPU_registers/)
- [Colour-emphasis games](../Colour-emphasis_games/)
- [Color \$0D games](../Color_$0D_games/)
- Talk:Myths#All_three_color_emphasis_bits_with_color_0D

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) PAL PPU swaps green and red emphasis bits: <http://forums.nesdev.org/viewtopic.php?p=131889#p13188>
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) Dendy PPU swaps green and red emphasis bits: <http://forums.nesdev.org/viewtopic.php?p=155513#p155513>
