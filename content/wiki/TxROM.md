---
title: "TxROM"
url: "/wiki/TxROM/"
description: "The generic designation TxROM refers to cartridge boards made by Nintendo that use the Nintendo MMC3 mapper."
wikiPage: "TxROM"
wikiCategories:
  - name: "In NesCartDB"
    url: "/category/In_NesCartDB/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Nintendo licensed mappers"
    url: "/category/Nintendo_licensed_mappers/"
---

<div class="infobox"><div style="text-align: center"><b>NESCartDB</b></div><table class="oddf0"><tbody><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=-T%25ROM">TxROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TBROM">TBROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TEROM">TEROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TFROM">TFROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TGROM">TGROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TKROM">TKROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TK1ROM">TK1ROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TKSROM">TKSROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TLROM">TLROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TL1ROM">TL1ROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TL2ROM">TL2ROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TLSROM">TLSROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TNROM">TNROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TQROM">TQROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TR1ROM">TR1ROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TSROM">TSROM</a></td></tr><tr><td><a href="https://nescartdb.com/search/advanced?unif_op=LIKE+`%25%40%25`&amp;unif=TVROM">TVROM</a></td></tr></tbody></table></div>

The generic designation **TxROM** refers to cartridge boards made by Nintendo that use the [Nintendo MMC3](../MMC3/) mapper.

## Board Types {#Board_Types}

The following TxROM boards are known to exist:

| Board | PRG ROM | PRG RAM | CHR | Comments |
|----|----|----|----|----|
| TBROM | 64 KB |  | 16 / 32 / 64 KB ROM |  |
| TEROM | 32 KB |  | 16 / 32 / 64 KB ROM | Supports fixed mirroring |
| TFROM | 128 / 256 / 512 KB |  | 16 / 32 / 64 KB ROM | Supports fixed mirroring |
| TGROM | 128 / 256 / 512 KB |  | 8 KB RAM/ROM |  |
| TKROM | 128 / 256 / 512 KB | 8 KB | 128 / 256 KB ROM |  |
| [TK1ROM](http://forums.nesdev.org/viewtopic.php?f=9&t=9891) | 128 KB | 8KB | 128KB ROM | Uses [7432](../7432/) for 28-pin CHR ROM |
| [TKSROM](../TKSROM/) | 128 / 256 / 512 KB | 8 KB | 128 KB ROM | Alternate mirroring control, Famicom only |
| TKEPROM | four 128KB ROMs | 8 KB | two 128KB ROMs | Prototyping board |
| TLROM | 128 / 256 / 512 KB |  | 128 / 256 KB ROM |  |
| TL1ROM | 128 KB |  | 128 KB | Uses [7432](../7432/) for 28-pin CHR ROM |
| TL2ROM |  |  |  | Nonstandard pinout |
| [TLBROM](https://forums.nesdev.org/viewtopic.php?p=228322#p228322) | 128 KB |  | 128 KB ROM | Uses 74541 to compensate for too-slow CHR ROM |
| [TLSROM](../TLSROM/) | 128 / 256 / 512 KB |  | 128 KB ROM | Alternate mirroring control |
| TNROM | 128 / 256, 512 KB | 8 KB | 8 KB RAM/ROM | Famicom only |
| [TQROM](../INES_Mapper_119/) | 128 KB |  | 16 / 32 / 64 KB ROM + 8 KB RAM |  |
| [TR1ROM](https://nescartdb.com/search/advanced?unif=NES-TR1ROM) | 128 / 256 / 512 KB |  | 64 KB ROM + 4 KB VRAM (4-screen [Mirroring](../Mirroring/)) | NES only |
| TSROM | 128 / 256 / 512 KB | 8 KB (no battery) | 128 / 256 KB ROM |  |
| [TVROM](https://nescartdb.com/search/advanced?unif=NES-TVROM) | 64 KB |  | 16 / 32 / 64 KB ROM + 4 KB VRAM (4-screen [Mirroring](../Mirroring/)) | NES only |

## Solder pad config {#Solder_pad_config}

### <span id="Namco_108_backwards_compatibility_.28TEROM_and_TFROM.29"></span>[Namco 108](../INES_Mapper_206/) backwards compatibility (TEROM and TFROM) {#Namco_108_backwards_compatibility_.28TEROM_and_TFROM.29}

- Normal mode: 'CL1' connected, 'CL2' connected, 'H' disconnected, 'V' disconnected.
- Backwards compatible with horizontal [mirroring](../Mirroring/): 'CL1' disconnected, 'CL2' disconnected, 'H' disconnected, 'V' connected
- Backwards compatible with vertical mirroring: 'CL1' disconnected, 'CL2' disconnected, 'H' connected, 'V' disconnected

Connecting 'CL1' enables MMC3-controlled mirroring, while connecting 'CL2' enables IRQs. However, the additional bankswitching modes available by the MMC3 that weren't available with the Namco chip used on DEROM boards are still present and activated by bits 7-6 of port \$8000.

### <span id="Battery_retention_.28TNROM.2C_TKROM_and_TKSROM.29"></span>Battery retention (TNROM, TKROM and TKSROM) {#Battery_retention_.28TNROM.2C_TKROM_and_TKSROM.29}

- PRG RAM retaining data: 'SL' disconnected, Battery, D1, D2, R1 R2 and R3 present.
- PRG RAM not retaining data: 'SL' connected, leave slots for Battery, D1, D2, R1, R2 and R3 free.

## Various notes {#Various_notes}

Boards with 4-screen mirroring uses a 8 KB SRAM chip, but only 4 KB is actually used. The 2 KB VRAM inside of the console is always disabled, and the CIRAM A10 pin of the MMC3 doesn't go to anything.

TLSROM and TKSROM boards have different mirroring control than other MMC3 boards. The mirroring is controlled directly by MMC3's CHR A17 line, and MMC3's CIRAM A10 pin doesn't go to anything. Due to their incompatibility with other MMC3 boards on a software viewpoint, they are assigned to [INES Mapper 118](../INES_Mapper_118/) instead of mapper 4.

TQROM board has both CHR ROM and RAM. Bit 6 of the bank number, which appears on MMC3's CHR A16 line, controls whenever CHR RAM or CHR-ROM is enabled. A [74HC32](../7432/) chip is used to combine this with other chip enable signals for the CHR-ROM and the CHR-RAM chips. Due to this incompatibility with the other MMC3 boards on a software viewpoint, this board is assigned to [INES Mapper 119](../INES_Mapper_119/) instead of mapper 4.
