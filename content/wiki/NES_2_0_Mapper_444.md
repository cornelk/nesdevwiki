---
title: "NES 2.0 Mapper 444"
url: "/wiki/NES_2.0_Mapper_444/"
description: "NES 2.0 Mapper 444 denotes the NC7000M multicart circuit board. Its UNIF board name is BMC-BS-110, which is the name of a different PCB whose name was mixed-up."
wikiPage: "NES_2_0_Mapper_444"
wikiCategories:
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
  - name: "NES 2.0 mappers with submappers"
    url: "/category/NES_2.0_mappers_with_submappers/"
---

[<img src="../../wiki-images/NC7000M.jpg" width="300" height="123" />](../../wiki-images/NC7000M.jpg)

[](../../wiki-images/NC7000M.jpg)

[Front view](../../wiki-images/NC7000M.jpg)

[](../../wiki-images/NC7000M.jpg)

**NES 2.0 Mapper 444** denotes the **NC7000M** multicart circuit board. Its UNIF board name is **BMC-BS-110**, which is the name of a different PCB whose name was mixed-up.

The top-right solder pads are relevant for bankswitching and are denoted by Submapper bits. A clear bit represents the *upper* pad connection:

- Submapper bit 0 (top-rightmost pad): CHR-ROM A17 connected to \$6000.0 (clear) or to the MMC3 CHR A17 output (set). Selects whether the MMC3 can address 128 or 256 KiB of CHR data.
- Submapper bit 1 (the top-second-rightmost pad): CHR-ROM A18 connected to \$6000.1 (clear) or to \$6000.4 (set). Since PRG A18 is also determined by \$6000.1, it thus selects whether CHR A18 can be determined separately or not.

Known cartridges:

- *7000-in-1* (CA-004): both bits set (submapper 3)
- *Super 8000-in-1*: both bits clear (submapper 0)

[With both pads in the position represented by submapper 0, NC7000M's definition has a subset in the](../../wiki-images/NC7000M.jpg) [**K-3043** PCB](https://forums.nesdev.org/viewtopic.php?p=277283).

## <span id="Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29"></span>Outer Bank and Mode Register (\$6000-\$7FFF, write) {#Outer_Bank_and_Mode_Register_.28.246000-.247FFF.2C_write.29}

    A~[011. .... ...C MMBA]
                    | ||++- PRG A18..17
                    | |||+- CHR A17 if submapper bit 0 clear
                    | ||+-- CHR A18 if submapper bit 1 clear
                    | ++--- Select PRG-ROM mode
                    |        0: MMC3 mode
                    |        1: NROM-256 mode (PRG A13..14=CPU A13..14, PRG A15..16=MMC3 A15..16)
                    |        2: Read solder pads' value in D0/D1 (leftmost pads in picture)
                    |        3: NROM-128 mode (PRG A13=CPU A13, PRG A14..16=MMC3 A14..16)
                    +------ CHR A18 if submapper bit 1 set

- WRAM must be enabled in \$A001.7 before writing to this register.
- The inner PRG bank is restricted to 128 KiB.
- NROM mode forces MMC3's CPU A13 and A14 inputs to GND to make MMC3 bank register 6 apply to the entire CPU \$8000-\$FFFF range.

## MMC3-compatible registers {#MMC3-compatible_registers}

Mask: \$E001

See [MMC3](../MMC3/).
