---
title: "MMC3 + NROM multicart"
url: "/wiki/MMC3_+_NROM_multicart/"
description: "All MMC3's PRG \u0026 CHR registers are available at $8000-$ffff for the whole time. There is no WRAM at $6000-$7fff but additional register (which can be protected from writes using MMC3's $a001 just…"
wikiPage: "MMC3___NROM_multicart"
wikiCategories:
  - name: "Mappers"
    url: "/category/Mappers/"
---

[<img src="../../wiki-images/MMC3+NROM_farid_components.jpg" width="300" height="181" alt="MMC3+NROM farid components.jpg" />](../../wiki-images/MMC3+NROM_farid_components.jpg)

[](../../wiki-images/MMC3+NROM_farid_components.jpg)

[](../../wiki-images/MMC3+NROM_farid_components.jpg)

[](../../wiki-images/MMC3+NROM_farid_components.jpg)

[](../../wiki-images/MMC3+NROM_farid_components.jpg)[<img src="../../wiki-images/MMC3+NROM_farid_top.jpg" width="300" height="178" alt="MMC3+NROM farid top.jpg" />](../../wiki-images/MMC3+NROM_farid_top.jpg)

[](../../wiki-images/MMC3+NROM_farid_top.jpg)

[](../../wiki-images/MMC3+NROM_farid_top.jpg)

[](../../wiki-images/MMC3+NROM_farid_top.jpg)

[](../../wiki-images/MMC3+NROM_farid_top.jpg)[<img src="../../wiki-images/MMC3+NROM_farid_bottom.jpg" width="300" height="178" alt="MMC3+NROM farid bottom.jpg" />](../../wiki-images/MMC3+NROM_farid_bottom.jpg)

[](../../wiki-images/MMC3+NROM_farid_bottom.jpg)

[](../../wiki-images/MMC3+NROM_farid_bottom.jpg)

[](../../wiki-images/MMC3+NROM_farid_bottom.jpg)

[](../../wiki-images/MMC3+NROM_farid_bottom.jpg)[<img src="../../wiki-images/MMC3+NROM_farid_schematics.png" width="300" height="152" alt="MMC3+NROM farid schematics.png" />](../../wiki-images/MMC3+NROM_farid_schematics.png)

[](../../wiki-images/MMC3+NROM_farid_schematics.png)

[](../../wiki-images/MMC3+NROM_farid_schematics.png)

## Overview {#Overview}

- PRG: 512kB ROM1 + 32kB ROM2
- CHR: 512kB ROM
- Mirroring: H or V (selected by software)
- Bus conflicts: no

All MMC3's PRG & CHR registers are available at \$8000-\$ffff for the whole time. There is no WRAM at \$6000-\$7fff but additional register (which can be protected from writes using MMC3's \$a001 just like it protects WRAM).

After power up & soft/hard reset, this is cleared and ROM2 is selected (for \$8000-\$ffff).

    [011. ...R .wMG GPPP] REG $6000-$7fff (address bits are latched)
             |  ||| ||||
             |  ||| |+++- selects inner 16KB PRG when in NROM mode
             |  ||+-+---- selects outer 128KB PRG & 128KB CHR (in any mode)
             |  |+------- 0: NROM mode, 1: MMC3 mode
             |  +-------- 1: protect whole $6000-$7fff register from further changes
             +----------- 0: selects ROM2, 1: selects ROM1 and protects this single bit from further changes

## PRG banking {#PRG_banking}

    condition         | $8000-$9fff | $a000-$bfff | $c000-$dfff | $e000-$ffff | notes
    ------------------+-------------+-------------+-------------+-------------+--------
    R=0               |  one 32 KB PRG from ROM2 (mapped to first 32 KB)      | NROM32 
    ------------------+-------------+-------------+-------------+-------------+--------
    R=1 & M=1         |   four CPU banks just as in MMC3 (from outer bank GG) | MMC3
    ------------------+-------------+-------------+-------------+-------------+--------
    R=1 & M=0 & GG=0  |        GGPPPP             |         GGPPPP            | NROM16
    ------------------+-------------+-------------+-------------+-------------+--------
    R=1 & M=0 & GG!=0 |         GGPPP0            |         GGPPP1            | NROM32
    ------------------+-------------+-------------+-------------+-------------+--------

## CHR banking {#CHR_banking}

CHR banking works in the same manner like in MMC3.

## Mirroring {#Mirroring}

Mirroring works in the same manner like in MMC3.

## Notes {#Notes}

More info about this multicart can be found here: <http://forums.nesdev.org/viewtopic.php?t=7926&start=15>
