---
title: "Text compression"
url: "/wiki/Text_compression/"
description: "Text compression refers to techniques that allow fitting more text data into a smaller space. General-purpose compression formats that work well on a platform with relatively plentiful RAM, such as a…"
wikiPage: "Text_compression"
---

**Text compression** refers to techniques that allow fitting more text data into a smaller space. General-purpose compression formats that work well on a platform with relatively plentiful RAM, such as a PC, Nintendo 64, or Nintendo DS, may not work on platforms whose RAM is much smaller, such as the NES.

## Dictionary compression and DTE {#Dictionary_compression_and_DTE}

Dictionary compression is a technique where part of the character set is reserved to denote references to a "dictionary". If the byte falls into this range, a string is copied from the dictionary rather than the byte being copied verbatim. It could be considered the textual equivalent of metatiles in scrolling maps. As this compression technique does not require knowledge of past data, it is very easy to implement on machines having little memory like the NES.

Sometimes the compression may be applied recursively, where the dictionary string itself may contain references to other dictionary strings.

Dual-tile encoding, or DTE for short, is a special case of dictionary compression. It is also known as [byte-pair encoding](https://en.wikipedia.org/wiki/Byte_pair_encoding), or digram coding. In this case, the dictionary strings are all two bytes long.

Example implementations:

### <span id="Simon.27s_Quest_.28FDS.2FNES.29"></span>Simon's Quest (FDS/NES) {#Simon.27s_Quest_.28FDS.2FNES.29}

In *Dracula II: Noroi no Fuuin* and *Castlevania II: Simon's Quest*, the font is 252 symbols long, although only a small part of that is actual text symbols used in dialog text.

| Value | Meaning |
|----|----|
| 00–FA | Print this symbol. |
| FB | Denotes the end of a substring. Restores the string pointer saved by opcode FC. |
| FC | Save current string pointer. The next byte determines the string number; this number will be used to calculate the new string pointer. Interpreting will continue from that address. |
| FE | Newline. |
| FF | End of text. If used in a substring, will not return to the main string. A string that ends in FC can omit the trailing FF, if the substring ends in FF. |

Dictionary strings are arbitrary length. There is room for only one saved string pointer, so substrings can not refer to other substrings, unless it is to terminate the entire string.

The substring mechanism is heavily utilized in the FDS version for common words/phrases. (E.g. スイショウ = crystal) The cartridge versions keep this mechanism, though none of the English text utilizes it.

### <span id="Chrono_Trigger_.28SNES.29"></span>Chrono Trigger (SNES) {#Chrono_Trigger_.28SNES.29}

In *Chrono Trigger* (SNES) (all versions), the font is 768 symbols long, but a significant number of those symbols can not be printed.

| Value | Meaning |
|----|----|
| 00 | End of string. |
| 01 | Read next byte; print symbol *byte+0x100*. |
| 02 | Read next byte; print symbol *byte+0x200*. |
| 03–20 | Various text effects, references to item tables, and references to party member names. |
| 21–xx | Reference to a dictionary string. xx is a compile-time constant that determines the length of the dictionary. This number is 0x9F in the USA version and 0x3F in the Japanese version. |
| xx+1–FF | Print this symbol. |

Dictionary strings have a length limit of 255 bytes. They are not applied recursively. The dictionary strings are stored in length-data format without an end delimiter.

### <span id="Joel_Yliluoma.27s_ppu_read_buffer_test_.28NES.29"></span>Joel Yliluoma's ppu_read_buffer test (NES) {#Joel_Yliluoma.27s_ppu_read_buffer_test_.28NES.29}

In Joel Yliluoma's emulator test ROM ppu_read_buffer ([\[1\]](http://forums.nesdev.org/viewtopic.php?t=8847)) the font is 128 symbols long, and in addition to the alphabet it includes some pre-rendered substrings in variable-width font. The ROM uses a combination of DTE and dictionary for text compression. For DTE, the compression is applied recursively, in both DTE bytes. Additionally the string may contain a jump label to another string, which forms a [DAFSA](https://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton) that allows the reuse of the same string suffix in different test descriptions.

| Value | Meaning |
|----|----|
| 00 | End of string. |
| 01–80 | Print this symbol. After printing, if the stack pointer is wrong, pop a byte and interpret it rather than loading the next byte from the string. |
| 81–FE | Push DTE_TABLE1\[n-0x81\] in stack, and interpret DTE_TABLE0\[n-0x81\]. |
| FF | A 16-bit word follows. This word is loaded as a new string pointer and loading continues from that address. The old string pointer is not saved. |

### <span id="Damian_Yerrick.27s_robotfindskitten_.28NES.29"></span>Damian Yerrick's robotfindskitten (NES) {#Damian_Yerrick.27s_robotfindskitten_.28NES.29}

[This implementation](http://forums.nesdev.org/viewtopic.php?f=22&t=11943) of [robotfindskitten](https://en.wikipedia.org/wiki/robotfindskitten) contains a compressor written in Python and a 6502 decompressor. Comments in the compressor (`dte.py`) refer to the method as "digram tree encoding" for two reasons: to emphasize its recursive nature and because code units aren't "tiles", especially when displayed with a proportional font. The decompressor first copies the compressed string into RAM, into the *end* of a fixed size buffer, where the first portion of the buffer (and also the string itself) doubles as a stack for the recursive DTE references. The string is then interpreted from the buffer as follows.

| Value | Meaning |
|----|----|
| 00 | End of string. This code is only interpreted when copying the compressed string. |
| 00–7F | Print this symbol, and increase the reading position in the compressed string. If the reading position is now at the end of the buffer, the string has ended. |
| 80–FF | *Replace* the current character with bpe_table\[n×2+1\] and decrease the reading position in the compressed string. Read bpe_table\[n×2+0\] and interpret it. The multiplication by 2 discards the carry. The replaced character will be interpreted next, when the decompressor again reads bytes from the compressed string. |

Decoding recursive DTE references with the hardware stack, rather than a software stack, would permit a scheme where a subroutine is repeatedly called to decompress one character at a time. (E.g. TakuikaNinja's [Dreamful Diskmag](https://github.com/TakuikaNinja/FDS-diskmag/blob/3c9044ca4698f516142bea8e9426fecd34d9a9a8/main.asm#L480) uses such a scheme to update scrolling text) The game however does not do that; it decompresses the entire string first, so that it can then perform operations like word-wrap on the decompressed text.

Ignoring the special meaning of byte 00 when decompressing means that the string itself could include raw 00 bytes, if hidden inside the bytepair tables. They could be used e.g. as newlines. The game however does not do that. Byte 0A acts as a newline, as is traditional on Linux and UNIX systems.

[The NES and Game Boy ports of 240p Test Suite](https://github.com/pinobatch/240p-test-mini) also use DTE for help text.

## Bitrate reduction methods {#Bitrate_reduction_methods}

### Fixed-bit encoding {#Fixed-bit_encoding}

When the character set is small, such as 64 characters at most, strings could be encoded in a bitstream that packs 6 bits per character rather than 8 bits per character. This results in 20 % reduction of data size.

Main article: [Fixed Bit Length Encoding](../Fixed_Bit_Length_Encoding/)

### Variable-bit encodings {#Variable-bit_encodings}

In variable-bit encodings different symbols are stored in different number of bits.

Example assembler code for reading a MSB-first bit stream (adopted from [Tokumaru's tile decompressor](../Tile_compression/#Tokumaru)):

    InitBitReader:
            ; Input:   InputStream = pointer to beginning of bitstream
            ; Output:  Y = buffer position
            ; Clobber: A,ZN
            lda #$80
            sta BitBuffer
            ldy #0
            rts

    ReadBits:
            ; Input:   X = number of bits to read (1..8)
            ;          Y = buffer position (modified)
            ; Output:  A = integer read
            ; Clobber: C,ZN
            lda #0
    @loop:  jsr ReadBit
            rol a
            dex
            bne @loop
            rts

    ReadBit:
            ; Input:   Y = buffer position (modified)
            ; Output:  C = bit
            ; Clobber: ZN
            asl BitBuffer
            bne @ret
            pha
             lda (InputStream), y
             iny
             beq @wrap
    @ret0:   rol a
             sta BitBuffer
            pla
    @ret:   rts
    @wrap:  inc InputStream+1
            bne @ret0

Ideally, you would store common symbols using few bits and infrequent symbols using more bits. Which brings us to…

#### Huffman coding {#Huffman_coding}

A special case of variable-bit encodings is [Huffman coding](https://en.wikipedia.org/wiki/Huffman_coding). Huffman coding defines the optimal coding for a given set of symbols based on a table of frequencies that each symbol is used.

When combined with static dictionary coding, the technique is called [Huffword](https://pineight.com/mw/?title=Huffword).

#### Arithmetic coding {#Arithmetic_coding}

Huffman coding is also a special case of arithmetic coding. In arithmetic coding, each symbol is represented by a range of binary fractions rather than a particular number of bits. As arithmetic coding was covered by a number of patents up to 1993, and it is calculation intensive, chances are no NES game uses it. However, on the Super NES, the S-DD1 and SPC7110 coprocessors implement a mathematical model that approximates arithmetic coding under license. And by 1999, a design-around called "range coding" was discovered, leading to LZMA compression.

## LZ based methods {#LZ_based_methods}

LZ77 operates based on references to previous decompressed data. Decompression requires that either the previous decompressed data be readable or that a sliding window of previous data be kept in RAM. This isn't very efficient on the NES for two reasons: the CPU is connected to only 2K of RAM (plus whatever is in the cartridge), and VRAM can be accessed only during vblank.

Lempel-Ziv methods are incapable of efficient random access on a low-RAM system. Random access to an LZ77 or LZ78 stream works in one of three ways:

- Decompress from the beginning to retrieve a substring. This is time-inefficient for the decompressor.
- Compress each substring independently. As LZ77 relies on correlation within a string, this makes the compressed data larger.
- Buffer the entire decompressed data in RAM. This requires more memory in the decompressor, but the tradeoff may work well on a platform with more RAM, such as the Commodore 64, Genesis, Super NES, or Game Boy Color.

*See also: [Tile compression#LZSS](../Tile_compression/#LZSS)*

## See also {#See_also}

- [Compression](../Compression/)
