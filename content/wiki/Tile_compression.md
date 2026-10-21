---
title: "Tile compression"
url: "/wiki/Tile_compression/"
description: "Tile compression refers to techniques that allow fitting more graphics data into a smaller space. Programs using CHR ROM cannot use compressed tiles, as their tile data must be stored in the PPU's…"
wikiPage: "Tile_compression"
---

**Tile compression** refers to techniques that allow fitting more graphics data into a smaller space. Programs using [CHR ROM](../CHR_ROM_vs._CHR_RAM/) cannot use compressed tiles, as their tile data must be stored in the PPU's native format. But programs using CHR RAM can process tile data while copying it from PRG ROM to CHR RAM, and this processing allows storing more tiles in the same space.

## <span id="RLE"></span>Run-length encoding {#Run-length_encoding}

[Run-length encoding](https://en.wikipedia.org/wiki/Run-length_encoding) transforms runs of identical bytes into a shorter sequence of bytes that specifies the length of the run.

In NES tile data, byte-level run-length encoding works well when a row of 8 pixels in a tile is identical to the row above it. It also works well for [nametable](../PPU_nametables/) data because a horizontal run of blank tiles becomes a single tile.

Pixel-level run-length encoding is much slower but can achieve impressive results within a tile.

There are several different RLE data formats.

### PCX {#PCX}

The [PCX](https://en.wikipedia.org/wiki/PCX) image format became popular on PC.

| Value | Meaning                                                        |
|-------|----------------------------------------------------------------|
| 00-BF | Write this byte to the output.                                 |
| C0-FF | Read another byte, and write it to the output *n* - 192 times. |

### PackBits {#PackBits}

The [PackBits](https://en.wikipedia.org/wiki/PackBits) format was invented by Apple for MacPaint. It is also used in TIFF files and a few homebrew releases by Damian Yerrick.

| Value | Meaning                                                        |
|-------|----------------------------------------------------------------|
| 00-7F | Copy *n* + 1 bytes from input to output.                       |
| 80    | No operation                                                   |
| 81-FF | Read another byte, and write it to the output 257 - *n* times. |

### Konami RLE {#Konami_RLE}

This format is used in the U.S. version of *Contra*, and the Japanese version of *Simon's Quest*. It can be decoded and encoded with the Python program [GraveyardDuck](https://www.romhacking.net/utilities/554/). Compression ratio is more or less identical to PackBits.

| Value | Meaning                                                  |
|-------|----------------------------------------------------------|
| 00-80 | Read another byte, and write it to the output *n* times. |
| 81-FE | Copy *n* - 128 bytes from input to output.               |
| FF    | End of compressed data.                                  |

*Blades of Steel* uses a subset of this format reserving a special value for jumping to a new PPU address. See: [Blades of Steel DataCrystal reference](https://datacrystal.romhacking.net/wiki/Blades_of_Steel:ROM_map)

### Bakutoushi Patton-kun RLE {#Bakutoushi_Patton-kun_RLE}

This format is used in *Bakutoushi Patton-kun* for CHR and PRG files loaded from disk. It can be decoded with the C and Python programs found [here](https://www.romhacking.net/forum/index.php?msg=466349).

The bit-fields and terminator value were likely chosen based on frequency analysis of the source data.

| Value | Meaning |
|----|----|
| 000nnnnn | Copy *n* + 1 bytes to the output. |
| CCCnnnnn | If *C* \>= 3, copy *C* + 1 bytes to the output *n* + 1 times. Otherwise, copy *C* bytes to the output *n* + 1 times. |
| 101xxxxx | End of compressed data. |

### GBA RLUnComp {#GBA_RLUnComp}

The BIOS of the Game Boy Advance and Nintendo DS contains a decompressor for an RLE format very similar to PackBits and Konami. As described in [GBATEK](http://problemkaputt.de/gbatek.htm#biosdecompressionfunctions), it has a 4-byte size header followed by this:

| Value | Meaning                                                          |
|-------|------------------------------------------------------------------|
| 00-7F | Copy *n* + 1 bytes from input to output.                         |
| 80-FF | Read one byte from input and write it to output *n* - 125 times. |

### PB53 {#PB53}

<figure class="mw-halign-right" typeof="mw:File/Frame">
<a href="../../wiki-images/PB53_illustrated.png"><img src="../../wiki-images/PB53_illustrated.png" width="208" height="252" /></a>
<figcaption>Example of a tile encoded with PB53</figcaption>
</figure>

This codec was conceived by Damian Yerrick as an alternative to PackBits for the [Action 53](../Action_53/) multicart, and a Python packer and 6502 unpacker are included in the Action 53 menu distribution. Unlike freeform RLE formats such as Konami and PackBits, PB53 operates on 16-byte units, making it easy to divide the decompressed data into fixed-size packets to be sent to the PPU during vblank while rendering is turned on. Like [LZSS](https://en.wikipedia.org/wiki/LZSS), PB53 uses [unary coding](https://en.wikipedia.org/wiki/unary_coding) on the lengths of literal runs to save on overhead from switching between literal and run modes. This means that like LZSS, it has a worst case expansion of 12.5%, but it works fairly well on real tile data and OK on nametable data, which have shorter runs than the high-resolution files for which PackBits was designed. It also has a special mode to accommodate the layout of Shiru's NROM games *LAN Master*, *Lawn Mower*, and *Chase*, which have many identical tiles between the two pattern tables to allow tile animation.

Each tile consists of several 8-byte planes, two planes in the NES implementation. For the first plane in a tile:

| Value | Meaning |
|----|----|
| 00-7F | PB8: After this control byte, copy the following byte from input to output. Then, for each bit in the control byte from 6 to 0, if the bit is 1, repeat the previous byte; otherwise, copy another byte from the input. This is somewhat similar to how control bytes are formatted in LZSS. |
| 80 | Write eight \$00 bytes. |
| 81 | Write eight \$FF bytes. |
| 82 | Copy one tile (16 bytes) starting one tile back. (Used for a repeated tile, such as the unused tiles in many games.) |
| 83 | Copy one tile starting one segment back, usually 4096 bytes. (Used for pattern tables that share tiles, as seen in several Shiru games. The decoder switches between two instances one segment apart, each with its own input stream and output buffer.) |
| 84 | Write sixteen \$00 bytes. (Solid color 0) |
| 85 | Write eight \$FF bytes then eight \$00 bytes. (Solid color 1) |
| 86 | Write eight \$00 bytes then eight \$FF bytes. (Solid color 2) |
| 87 | Write sixteen \$FF bytes. (Solid color 3) |

For other planes:

| Value | Meaning |
|----|----|
| 00-81 | Same as first plane |
| 82 | Copy previous 8 bytes. (Used for 1-bit tiles with colors 0 and 3.) |
| 83 | Copy previous 8 bytes, bit-inverted. (Used for 1-bit tiles with colors 1 and 2.) |

PB8 is the same as PB53 except that bit 7 of the control byte is not special: it still means repeat the previous byte. It is used in all three of Retrotainment's *Haunted* games, as well as *Full Quiet*, *Garbage Pail Kids: Mad Mike and the Quest for Stale Gum*, the *Cease & Desist* demo, and the title screen of *Xcavator*. From [July 2019](https://github.com/LIJI32/SameBoy/commit/4504de828a188b520a324687a1181af6c45a7e3a) to [May 2020](https://github.com/LIJI32/SameBoy/commit/cb738190be6abca03d1cd3265440908107168a47), the workalike boot ROM included with the Game Boy Color emulator SameBoy used PB8 for the emulator's logo.

PB16 is similar to PB8 with one change: each 1 bit means a repeat of the value two bytes back. (Bit 7 is not special, unlike in PB53.) This distance of two bytes is better for Game Boy and Super NES tile data and Super NES tile maps. It is used in the Game Boy ports of 240p Test Suite and *Magic Floor*.

PB8 and PB16 inspired the creation of PB12 for the SameBoy emulator's boot ROM. So-called "PB12" by NieDzejkob (Jakub Kądziołka) is tuned to the statistics of the 3-color antialiased version of the SameBoy logo. Control bytes are interleaved with literal bytes. Control byte 00000001 is a terminator and thus must not occur in the byte stream. Otherwise, each 2 or 4 bits of the control correspond to a 4-bit word.

| Value | Meaning                                                   |
|-------|-----------------------------------------------------------|
| 00    | copy next byte from input                                 |
| 0100  | Copy 1 byte back, ORed with the same byte shifted left 1  |
| 0101  | Copy 1 byte back, ANDed with the same byte shifted left 1 |
| 0110  | Copy 1 byte back, ORed with the same byte shifted right 1 |
| 0111  | Copy 1 byte back, ANDed with the same byte shifted left 1 |
| 10    | Copy 2 bytes back                                         |
| 11    | Copy 1 byte back                                          |

### RLEINC {#RLEINC}

This RLE variant was used by Joel Yliluoma in the Simon's Quest retranslation project. It is very efficient when compressing nametables, which often contain redundancy in more complex forms than simple runs of repeating bytes. Examples include brick walls, which repeat two tiles, and complex graphics that is formed from an ascending series of successive tiles. For bitmap compression, it is slightly inferior to simpler RLE methods.

| Value | Meaning |
|----|----|
| 00–3F | LIT: Copy (n+1) bytes from input to output *backwards* |
| 40 | END: End of stream |
| 41–7F | SEQ: Read next byte b. Put b, (n−0x3F) times; add 1 to b after each iteration |
| 80–9F | DBL: Read next byte b1, and next byte b2. Put b1, (n−0x7D) times; swap b2 and b1 after each iteration |
| A0–FF | RUN: Read byte b. Put b, (0x101−n) times. |

A compressor for this scheme can be found at <http://bisqwit.iki.fi/src/nes-ppu_rleinc_compress.php.txt> (PHP), and a decompressor at <http://bisqwit.iki.fi/src/nes-ppu_rleinc_v2b.inc> (CA65).

JRoatch made [PBJ](http://forums.nesdev.org/viewtopic.php?p=161617#p161617), which adds the PB8 mode from PB53 and a common-byte mode to a modified RLEINC.

### Bit-based RLE {#Bit-based_RLE}

Most RLE schemes deal with whole bytes. There are also schemes where the compressed data forms a bitstream, that contains integers of different bit widths.

When compressing the combined tile graphics of Super Mario Bros. and Kirby's Adventure, a simple reference RLE algorithm (PackBits) gets a 12% reduction in data size. However, if the algorithm is modified as indicated below, a 21% reduction is achieved. For comparison, the graphics specialized algorithm in PB53 achieves 25% for that data set. Tokumaru compression manages to reduce the data by 41%.

| Bit sequence | Meaning                                                |
|--------------|--------------------------------------------------------|
| 0000         | End of stream.                                         |
| 0nnn         | Copy the next n×8 bits, i.e. *n* bytes, to the output. |
| 1nnn         | Read the next 8 bits, and output this byte n+2 times.  |

A possible reason why bit-based compression methods are not popular on the NES is that bit-shifts are cumbersome and slow with the 6502 CPU, as it can only shift one bit at a time. The above algorithm is still relatively simple to implement, as it operates on units of 4 bits for the input and full bytes for the output. Coincidentally, it also produced the best compression out of all bit-based RLE algorithms that were brute-force-tested for that dataset.

### NES Stripe Image RLE {#NES_Stripe_Image_RLE}

A RLE format mostly used by Nintendo for use in their Arcade ports as well as their Mario games, Also used in some homebrew games!

Format: dest, len, data, \[dest, len, data, \]\* end

dest: PPU destination address (2 bytes, big endian), to be written to \$2006

len: Length (Byte) of PPU Buffer Data:

| Length Value | Meaning |
|----|----|
| 00-3F | Literal to right: Copy *n* + 1 bytes to video memory addresses increasing by 1 |
| 40-7F | Run to right: Copy one byte *n* - 63 times to video memory addresses increasing by 1 |
| 80-BF | Literal down: Copy *n* - 127 bytes to video memory addresses increasing by 32 |
| C0-FF | Run down: Copy one byte *n* - 191 times to video memory addresses increasing by 32 |

data: PPU Data to display

end: End of Data marker. Early games use \$00; later games (particularly those with CHR RAM) may use any value with bit 7 set (\$80-\$FF) to allow writing to the first 16 tiles of video memory.

The FDS BIOS provides a routine to decode a [variant](../FDS_BIOS/#VRAM_transfer_structure) of this format with support for callable substructures.

See also [Popslide](https://forums.nesdev.org/viewtopic.php?f=22&t=15440), a video memory update buffer framework using the NES stripe image RLE format.

### SNES Stripe Image RLE {#SNES_Stripe_Image_RLE}

Same RLE format used by Nintendo as above, but for SNES.

Format: dest, len, data, end

dest: PPU Destination: \$2116 and \$2117

len: Length (Word) of PPU Buffer Data:

| Length Value | Meaning |
|----|----|
| 0000-3FFF | Copy *n* + 1 bytes to PPU buffer |
| 4000-7FFF | Copy *n* + 1 bytes to PPU buffer, with RLE |
| 8000-BFFF | Copy *n* + 1 bytes to PPU buffer, increment 32 bytes |
| C000-FFFF | Copy *n* + 1 bytes to PPU buffer, with RLE, increment 32 bytes |

data: PPU Data to display in 2-byte increments (or word increments)

end: Unlike the NES version, the end byte is \$FF or \$FFFF.

### Game Boy Printer {#Game_Boy_Printer}

This is an algorithm used by the [Game Boy Printer](https://en.wikipedia.org/wiki/Game_Boy_Printer) accessory that allows for combining runs and raw bytes.

| Value | Meaning                                                        |
|-------|----------------------------------------------------------------|
| 00-7F | Copy *n* + 1 bytes from input to output.                       |
| 80-FF | Read another byte, and write it to the output *n* - 126 times. |

## Pixel based {#Pixel_based}

### Codemasters {#Codemasters}

This is a Markov chain (predictive) algorithm that packs predictions in varying number of bits. It works on packets measured in whole tiles, and compresses mostly pixel by pixel, making it slow. [Explained on forum](http://forums.nesdev.org/viewtopic.php?p=48658#p48658).

#### <span id="Tokumaru"></span>Tokumaru {#Tokumaru}

Tokumaru discovered an improvement to the way the frequency tables are changed in Codemasters algorithm, and [released](http://forums.nesdev.org/viewtopic.php?p=54230#p54230) the compressor and decompressor as open source. And open-source rewrite of the encoder and decoder with slightly better performance can be downloaded at: <http://bisqwit.iki.fi/source/tokumaru.html>

The compressed data begins with a byte that tells how many tiles it encodes. 256 is maximum. Technically this byte can be ignored, if you already know how many tiles you are going to decode.

After the byte, any number of blocks follows. Each block begins with a *color description table*. This table tells how to encode transitions between colors in tiles belonging to this block.

There are four elements in this table, from 3 to 0, for color *n*. Each element begins with a two-bit integer *ncolors\[n\]*, that tells how many different colors that may follow a pixel of this particular color. After the number of colors, comes a pivot color *a* that is encoded as follows:

| Value | Applicable when | Meaning |
|----|----|----|
| nothing | *ncolors\[n\]=0* | No pivot color |
| 1 bit: 0 | *ncolors\[n\]\>0* | Pivot color *a* is 1 if *n \< 1*, 0 otherwise. |
| 2 bits: 10 | *ncolors\[n\]\>0* | Pivot color *a* is 2 if *n \< 2*, 1 otherwise. |
| 2 bits: 11 | *ncolors\[n\]\>0* | Pivot color *a* is 3 if *n \< 3*, 2 otherwise. |

The table of transition colors is then calculated using *n*, *a*, and the number of colors *ncolors\[n\]*. First, two other colors *b* and *c* are calculated. They are the first color indexes that are neither *n* nor *a*. E.g. if *a=2* and *n=1*, *b* and *c* become 0 and 3 respectively.

| When *ncolors\[n\]* is | Table of possible transition colors *next\[n\]* is |
|------------------------|----------------------------------------------------|
| 0                      | \[\]                                               |
| 1                      | \[a\]                                              |
| 2                      | \[b,c\]                                            |
| 3                      | \[a,b,c\]                                          |

For compression purposes, ideally *ncolors\[n\]* should be chosen to be the numbers of distinct colors that actually can follow the color *n*, based on measuring the tile data, and, if *ncolors\[n\]=3*, the pivot color *a* should be the color that is transitioned into most often from color *n*. This transition in tile data will be encoded in two bits, while the other transitions are encoded in three bits. For other values of *ncolors\[n\]*, the choice of pivot color is mandated by the actual possible colors.

After the *color description table*, comes *tile data* encoding 16 bytes, or 8×8 pixels. Each tile is comprised of eight *rows* of pixels. Each pixel row begins with 1 bit, that tells whether the row is to be copied. If the bit is set, the previously decoded row is duplicated, and no other data is encoded for this pixel row. At the start of the block, the "previously decoded row" is assumed to contain zero bytes. The previous row is not reset between different tiles, unless a new block begins. If the bit was clear, eight pixels follow for *x* coordinates 0 to 7. Each pixel is encoded as follows, depending on the color *c* of the previous pixel:

| Value | Applicable when | Meaning |
|----|----|----|
| 2 bits | *x = 0* | The color of the first pixel *c* is encoded as a 2-bit integer. |
| nothing | *ncolors\[c\]=0* | *c* does not change, and nothing is encoded. |
| 1 bit: 1 | *ncolors\[c\]\>0* | *c* does not change from previous pixel. |
| 1 bit: 0 | *ncolors\[c\]=1* | *c* becomes *next\[c\]\[0\]*. |
| 2 bits: 00 | *ncolors\[c\]\>1* | *c* becomes *next\[c\]\[0\]*. |
| 2 bits: 01 | *ncolors\[c\]=2* | *c* becomes *next\[c\]\[1\]*. |
| 3 bits: 010 | *ncolors\[c\]=3* | *c* becomes *next\[c\]\[1\]*. |
| 3 bits: 011 | *ncolors\[c\]=3* | *c* becomes *next\[c\]\[2\]*. |

After each full tile, if there are still remaining tiles to be decoded, comes one bit that tells what comes next. Value 1 means a new block start, with a new *color description table*. Value 0 means that more *tile data* will follow.

## Common byte {#Common_byte}

### Oracle common byte {#Oracle_common_byte}

This codec, used in *The Legend of Zelda: Oracle of Seasons* and *The Legend of Zelda: Oracle of Ages* according to [Dwedit](http://forums.nesdev.org/viewtopic.php?p=130355#p130355), is roughly comparable to RLE in complexity. For each 16-byte block, the compressor determines the most common byte in that block. The compressed data for each block starts with a 2-byte mask, then the most common byte, then other bytes in order that aren't the most common.

To decode a block: First read the two-byte mask. If the entire mask is zero, set the bit corresponding to the first byte to true. Then read the most common byte. For each bit in the mask, if the bit is 1, write the most common byte to output; otherwise, copy one byte.

Maximum expansion is 12.5% for any block that has 16 different bytes in it: two bytes of mask and 16 bytes of data.

## LZSS {#LZSS}

A lot of games on platforms after the NES use [LZ77](https://en.wikipedia.org/wiki/LZ77) family compression methods such as [LZSS](https://en.wikipedia.org/wiki/LZSS), which generalizes run-length encoding to allow copying data from several bytes ago, not just the previous byte. Few NES games use LZ77 because the NES's limited work RAM and limited access to video memory make it difficult to resolve back references. (Fewer still use LZW or any other LZ78-based method because of patents that subsisted through the NES, Super NES, and Nintendo 64 eras.)

In LZSS, the mask contains 8 commands, each either to copy a literal byte or to copy a back reference. determines whether the next 8 things are literal bytes or back references. Once all commands have been processed, read another mask.

Read a mask byte from input.

For each bit in the mask byte:

If the bit is 0, this is a literal:

Copy one byte from input to output.

Otherwise, this is a back-reference:

Read and decode a length and distance from input. These will be positive integers.

Copy *length* bytes from the previous output, *distance* bytes before now, to output.

LZSS flavors vary mainly in how they encode lengths and distances.

### LZ77UnComp {#LZ77UnComp}

The BIOS of the Game Boy Advance and Nintendo DS has a built-in decompressor for a straightforward LZSS flavor with 3- to 18-byte references into the previous 4096 bytes of output. The data format is described in [Martin Korth's GBATEK](http://problemkaputt.de/gbatek.htm#biosdecompressionfunctions).

Caution: In data intended for decompression directly to the GBA or DS video memory, the second byte of a 16-bit word cannot refer to the first byte of the same word. So the encoder must not write a run with distance 1 that starts at an odd offset.

### Oracle LZ {#Oracle_LZ}

This flavor of LZSS is used in *The Legend of Zelda: Oracle* games for Game Boy Color according to [Dwedit](http://forums.nesdev.org/viewtopic.php?p=130355#p130355).

An entire compressed block can be compressed with one of two subtypes of Oracle LZ: short word mode and long word mode. Short words use references of 2 to 8 bytes into the previous 32 bytes of output, and long words use references of 3 to 33 bytes into the previous 2048 bytes. (A length value of 0 means read an additional byte and use that as the length.) Only short words would work very well on NES.

### <span id="Pok.C3.A9mon_LZ"></span>Pokémon LZ {#Pok.C3.A9mon_LZ}

This compression scheme is used in the second-generation Pokémon games on the Game Boy. It is used for bitmap compression.

A byte n is read and split into two parts: code = n \>\> 5, and c = n & 0x1F. Byte 0xFF marks the end of the stream. Otherwise the *code* is interpreted as follows:

| code | Meaning |
|----|----|
| 0 | Copy the next *c* + 1 bytes to the output. |
| 1 | Read another byte, and write it to the output *c* + 1 times. |
| 2 | Read another byte b1 and byte b2, and write byte b1 to the output *c* + 1 times, swapping b1 and b2 after each write. |
| 3 | Write a zero byte (0x00) to the output *c* + 1 times. |
| 4 | Read byte b. If b \< 0x80, then read byte b2; offset is b×256+b2 bytes from the output stream beginning. Else offset = b bytes *behind* from the current output stream end. Copy *c* + 1 bytes from the output stream at offset to the output. |
| 5 | Read byte b. If b \< 0x80, then read byte b2; offset is b×256+b2 bytes from the output stream beginning. Else offset = b bytes *behind* from the current output stream end. Copy *c* + 1 bytes from the output stream at offset to the output, *reversing the bits in each byte*. |
| 6 | Read byte b. If b \< 0x80, then read byte b2; offset is b×256+b2 bytes from the output stream beginning. Else offset = b bytes *behind* from the current output stream end. Copy *c* + 1 bytes from the output stream at offset to the output, *decrementing the read position after each write* (i.e. reverse the data). |
| 7 | Read another byte b. Change code = (c \>\> 2), and change c = (c & 3) × 256 + b. Re-interpret code according to this table. |

### Chrono Trigger LZ {#Chrono_Trigger_LZ}

This compression scheme is used in Square‘s Chrono Trigger for the SNES for compression of graphics and various data.

The data consists of packets. Each packet begins with a 16-bit offset, which gives the packet ending offset relative to the beginning of compressed data. At that offset, there is always a control byte. The first control byte sets the size of offsets: If the byte is \< 0xC0, then *offsetsbits*=12. Else *offsetbits*=11. Interpreting the offsetbits is done only once. The higher-order two bits in the control bytes of all other packets are ignored. The *counter* is assigned a default value of 8.

The decompression loop goes as follows:

1.  If the packet end has been reached, a control byte is read, and *counter* is assigned the low 6 bits of that byte (i.e. *counter* = byte & 0x3F). If the *counter* is zero, the decompression is complete and ends there. If the counter was nonzero, a new 16-bit packet end offset is read.
2.  If the end of the packet has not yet been reached, a *mask* byte is read.
    - If the mask byte is zero, then next eight bytes are copied verbatim to the output. The counter is not affected.
    - If the mask byte was nonzero, its each *bit* is read, beginning from the lowest-order bit. The number of bits to be read is determined by *counter* (which is in range 1—63, i.e. it can be both smaller and greater than eight).
      - If the *bit* is zero, a single byte is copied verbatim to the output.
      - If the bit is nonzero, a 16-bit number is read from the input. *offset* becomes the lowest-order *offsetbits* from that number, and *length* is the rest of the bits plus 3. The decompressor copies *length* bytes from *offset* bytes behind to the output.
      - After all bits have been read, the *counter* is reset to the default value of 8, and the decompression loop continues.

## External links {#External_links}

- [Ad Hoc Compression Methods: RLE](https://hbfs.wordpress.com/2009/04/14/ad-hoc-compression-methods-rle/) describes various pixel-level RLE methods applied to a drawing of a Pokémon
- [Donut](https://forums.nesdev.org/viewtopic.php?t=17313) - A fast and efficient CHR compression library by JRoatch.
- [A collection of CHR files that can be used to test compression.](https://forums.nesdev.org/viewtopic.php?p=55526#p55526)
