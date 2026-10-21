---
title: "NES 2.0 Mapper 257"
url: "/wiki/NES_2.0_Mapper_257/"
description: "NES 2.0 Mapper 257 denotes two different mappers used by educational computer cartridges from 东达 (Dongda) and 鸿达 (Hongda)."
wikiPage: "NES_2_0_Mapper_257"
---

NES 2.0 Mapper 257 denotes two different mappers used by educational computer cartridges from 东达 (Dongda) and 鸿达 (Hongda).

- **Submapper 1**: Russian variant of the Dongda PEC-586 educational computer
- **Submapper 2**: Chinese variant of the Dongda PEC-586 educational computer, as well as similar cartridges from other manufacturers.

**Submapper 0** means auto-detect the variant by PRG-ROM size: submapper 1 if less than 512 KiB of PRG-ROM, submapper 2 if 512 KiB of PRG-ROM or more.

# Submapper 2 {#Submapper_2}

## Banks {#Banks}

Common:

- CPU \$6000-\$7FFF: 8 KiB of unbanked PRG-RAM
- CPU \$0000-\$1FFF: 8 KiB of unbanked CHR-RAM
- Nametable arrangement: switchable between horizontal and vertical

There are three banking modes: Scattered, 32 KiB switchable, and Mixed.

### Scattered PRG Banking Mode {#Scattered_PRG_Banking_Mode}

The entire \$8000-\$FFFF CPU address space is divided into 1 KiB banks. Each 1 KiB bank in the CPU address space is the last 1 KiB of an 8 KiB bank in the second 256 KiB half of 512 KiB PRG-ROM. For example, CPU address \$9ABC is 1 KiB bank \#6 (\$9ABC minus \$8000 divided by \$0400), so select the last 1 KiB of 8 KiB bank \#38 (6 plus 32 from being in the second 256 KiB half), and within that bank, offset \$02BC (\$9ABC AND \$03FF), or ROM offset \$4DEBC. The ROM address in a diagram:

    [01BB BBB1 11AA AAAA AAAA]
     AAAA: CPU A9..A0
     BBBB: CPU A14..A10

Example:

     $8000-$83FF: 1 KiB PRG-ROM bank# $107 (ROM offset $41C00)
     $8400-$87FF: 1 KiB PRG-ROM bank# $10F (ROM offset $43C00)
     $8800-$8BFF: 1 KiB PRG-ROM bank# $117 (ROM offset $45C00)
     $8C00-$8FFF: 1 KiB PRG-ROM bank# $11F (ROM offset $47C00)
     ...
     $F800-$FBFF: 1 KiB PRG-ROM bank# $1F7 (ROM offset $7DC00)
     $FC00-$FFFF: 1 KiB PRG-ROM bank# $1FF (ROM offset $7FC00)

### 32 KiB Switchable PRG Banking Mode {#32_KiB_Switchable_PRG_Banking_Mode}

The Mode register's lower three bits select one of eight 32 KiB banks, from the first 256 KiB half of 512 KiB PRG-ROM.

### Mixed PRG Banking Mode {#Mixed_PRG_Banking_Mode}

Five bits from the Mode register select one of thirty-two 8 KiB banks, from the second 256 KiB half of 512 KiB PRG-ROM, that are mapped to CPU \$8000-\$9FFF. For the rest of CPU address space (\$A000-\$FFFF), the Scattered PRG Banking Mode still applies. The lower three bits of the Mode Register are the same bits as those in 32 KiB Switchable PRG Banking Mode, but with different meaning!

## Registers {#Registers}

### <span id="Mode_and_PRG_Bank_.28.245000.2C_write.29"></span>Mode and PRG Bank (\$5000, write) {#Mode_and_PRG_Bank_.28.245000.2C_write.29}

    D~[CMpM pPPP] A~[0101 .000 .... ....] (Address AND $F700)=$5000
       ||+|-++++-- PRG bank (see below):
       || | |        32 KiB mode: PRG A17..A15=PPP, PRG A18=0, pp ignored
       || | |        Mixed mode:  PRG A17..A13=ppPPP, PRG A18=1 (CPU $8000-$9FFF only)
       || | +----- Nametable arrangement: 1=Vertical arrangement/Horizontal mirroring
       |+-+------- PRG banking mode (see below):
       |             0:   Scattered mode
       |             1/3: 32 KiB mode
       |             2:   Mixed mode
       +---------- CHR mode
                    0=normal
                    1=1 bits per pixel

1 bpp video mode means that CHR A3 is latched from PPU A0, and CHR A12 is latched from PPU A9, on the rising edge of PPU A13, creating one continuous 1bpp 8 KiB frame buffer. This mode is identical to that offered by [INES Mapper 164](../INES_Mapper_164/).

### <span id="Tape_output_.28.245100.2C_write.29"></span>Tape output (\$5100, write) {#Tape_output_.28.245100.2C_write.29}

    D~[.... ..T.] A~[0101 .001 .... ....] (Address AND $F700)=$5100
              +--- Data to tape

### <span id="Tape_input_.28.245300.2C_read.29"></span>Tape input (\$5300, read) {#Tape_input_.28.245300.2C_read.29}

    D~[.... ..T.] A~[0101 .011 .... ....] (Address AND $F700)=$5300
              +--- Data from tape

### Unknown writable registers {#Unknown_writable_registers}

    $5400,$5700

### Unknown readable registers {#Unknown_readable_registers}

    $5500
