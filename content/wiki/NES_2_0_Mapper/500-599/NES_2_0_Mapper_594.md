---
title: "NES 2.0 Mapper 594"
url: "/wiki/NES_2.0_Mapper_594/"
description: "NES 2.0 Mapper 594 denotes the combination of Rinco Talkman 3-in-1 multicart (PSG code FSG2) using the enhanced audio features of Rinco's Talkman RS-7600 famiclone console."
wikiPage: "NES_2_0_Mapper_594"
wikiCategories:
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "Mappers with scanline IRQs"
    url: "/category/Mappers_with_scanline_IRQs/"
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 594** denotes the combination of *Rinco Talkman 3-in-1* multicart (PSG code **FSG2**) using the enhanced audio features of Rinco's *Talkman RS-7600* famiclone console.

# Hardware overview {#Hardware_overview}

- 3 MiB of PRG-ROM, spread across a 2 MiB and a 1 MiB chip
- 1 MiB of CHR-ROM
- CPU \$8000-\$FFFF mapped by [MMC3](../MMC3/) clone from 512 KiB window into the 3 MiB of PRG-ROM
- CPU \$6000-\$7FFF mapped by discrete chips into the entire PRG-ROM address space
- PPU \$0000-\$1FFF mapped either like [INES Mapper 187](../INES_Mapper_187/) (512 KiB window) or like normal MMC3 (256 KiB window)
- The 6502 program sends ADPCM data to an Advanced Micro Devices Am7202A 1024-byte FIFO until it is at least half full
- The FIFO serves an OKI MSM6585 ADPCM IC

# Registers {#Registers}

These registers exist in addition or, where their address ranges would overlap, in lieu of the standard MMC3 registers.

## <span id="Extra_Bank_.28.249000.2C_write.29"></span>Extra Bank (\$9000, write) {#Extra_Bank_.28.249000.2C_write.29}

    D~[BBBB BBBB] (Address AND $F001)=$9000
       ++++-++++-- PRG A20..A13

Selects the 8 KiB bank number that is mapped to CPU \$6000-\$7FFF. Up to an entire 2 MiB are accessible, depending on the register \$B000 configuration, while the MMC3 can only address 512 KiB. The respective 2 MiB and 512 KiB address spaces are set by register \$B000.

## <span id="Unknown_Register_.28.249001.2C_write.29"></span>Unknown Register (\$9001, write) {#Unknown_Register_.28.249001.2C_write.29}

    D~[...? ....] (Address AND $F001)=$9001

The function of this register is not known; the three games always write \$10 to \$9001 before writing to \$9000.

## <span id="Address_Space_and_Mode_.28.24B000.2C_write.29"></span>Address Space and Mode (\$B000, write) {#Address_Space_and_Mode_.28.24B000.2C_write.29}

    D~[BB.. ....] (Address AND $F000)=$B000
       ++--------- PRG and CHR address space and mode

Meanining of the BB bits:

    00: Extra Bank addresses PRG $000000-$1FFFFF
        MMC3 addresses PRG $000000-$07FFFF
                   and CHR $00000-$3FFFF (PPU $0000-$0FFF)
                   and CHR $40000-$7FFFF (PPU $1000-$1FFF)
    40: Extra Bank addresses PRG $180000-$1FFFFF
        MMC3 addresses PRG $180000-$1FFFFF
                   and CHR $80000-$BFFFF (PPU $0000-$1FFF)
    80: Extra Bank addresses PRG $200000-$2FFFFF
        MMC3 addresses PRG $200000-$27FFFF
                   and CHR $C0000-$FFFFF (PPU $0000-$1FFF)

## <span id="FIFO_Data_.28.245000.2C_write.29"></span>FIFO Data (\$5000, write) {#FIFO_Data_.28.245000.2C_write.29}

    D~[DDDD DDDD] (Address AND $F001)=$5000
       ++++-++++-- ADPCM data byte

Writes a data byte into the 1024-byte FIFO. Probably only a good idea if the FIFO isn't full.

## <span id="FIFO_Reset_and_ADPCM_Playback_Rate.28.245001.2C_write.29"></span>FIFO Reset and ADPCM Playback Rate(\$5001, write) {#FIFO_Reset_and_ADPCM_Playback_Rate.28.245001.2C_write.29}

    D~[DD.. ..] (Address AND $F001)=$5001
       ++--------- Playback Rate
                   0: 4 kHz
                   1: 8 kHz
                   2: 16 kHz
                   3: 32 kHz

Writing to this register empties the FIFO and sets the sampling rate of the OKI MSM6585.

## <span id="FIFO_Status_.28.245001.2C_read.29"></span>FIFO Status (\$5001, read) {#FIFO_Status_.28.245001.2C_read.29}

    D~[?D?? .....] (Address AND $F001)=$5001
        +---------- FIFO half-full status
                    0=FIFO at least half-full
                    1=FIFO not at least half full

Games keep writing bytes to the FIFO as long as bit 6 is set. Since this register exists in the Rinco console, and a normal console returns open bus on reads with bit 6 set, games will never exit the "fill-FIFO" loop on a normal console.

# Further Reading {#Further_Reading}

- [Advanced Micro Devices Am7202A datasheet](https://rocelec.widen.net/view/pdf/fpsxe16aba/AMDIS02260-1.pdf?t.download=true&u=5oefqw)
- [OKI MSM6585 datasheet](https://datasheet.datasheetarchive.com/originals/library/Datasheet-027/DSA00482771.pdf)
