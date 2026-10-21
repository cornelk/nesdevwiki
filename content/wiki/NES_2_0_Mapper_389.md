---
title: "NES 2.0 Mapper 389"
url: "/wiki/NES_2.0_Mapper_389/"
description: "NES 2.0 Mapper 389 denotes the circuit board of the Caltron 9-in-1 multicart."
wikiPage: "NES_2_0_Mapper_389"
wikiCategories:
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

**NES 2.0 Mapper 389** denotes the circuit board of the *Caltron 9-in-1* multicart.

## <span id="Mirroring.2FOuter_PRG-ROM_Bank_register_.28.248000-.248FFF.2C_write.29"></span>Mirroring/Outer PRG-ROM Bank register (\$8000-\$8FFF, write) {#Mirroring.2FOuter_PRG-ROM_Bank_register_.28.248000-.248FFF.2C_write.29}

    Mask: $F000

    A~FEDC BA98 7654 3210
      -------------------
      1000 0000 .PPP P..M
                 ||| |  +- Mirroring: 0=Vertical, 1=Horizontal
                 +++-+---- PRG A15-A18

    Power-up value: $00

## <span id="PRG_Mode.2FOuter_CHR-ROM_Bank_register_.28.249000-.249FFF.2C_write.29"></span>PRG Mode/Outer CHR-ROM Bank register (\$9000-\$9FFF, write) {#PRG_Mode.2FOuter_CHR-ROM_Bank_register_.28.249000-.249FFF.2C_write.29}

    Mask: $F000

    A~FEDC BA98 7654 3210
      -------------------
      1001 0000 ..CC C.U.
                  || | +-- 0=NROM-256, 1=UNROM-064
                  ++-+---- CHR A15-A17

    Power-up value: $00

In NROM-256 mode, the Inner Bank register's PRG bits are ignored. In UNROM-064 mode, the Inner Bank register's PRG bits select one of four 16 KiB PRG-ROM banks at CPU \$8000-\$BFFF, while CPU \$C000-\$FFFF is hard-wired to inner bank \#3, similarly to [INES Mapper 081](../INES_Mapper_081/), used by the stand-alone version of the only game using this mode.

## <span id="Inner_Bank_register_.28.24A000-.24FFFF.2C_write.29"></span>Inner Bank register (\$A000-\$FFFF, write) {#Inner_Bank_register_.28.24A000-.24FFFF.2C_write.29}

    Mask: $E000

    A~FEDC BA98 7654 3210
      -------------------
      .... .... .... PPCC
                     ||++- CHR A13-A14
                     ++--- PRG A14-A15 in UNROM-064 mode when CPU A14=0

    Power-up value: $00
