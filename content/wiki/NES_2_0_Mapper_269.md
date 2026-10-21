---
title: "NES 2.0 Mapper 269"
url: "/wiki/NES_2.0_Mapper_269/"
description: "NES 2.0 Mapper 269 is used for at least one plug-and-play console and two multicarts with an MMC3 clone and four additional outer bank registers:"
wikiPage: "NES_2_0_Mapper_269"
wikiCategories:
  - name: "MMC3-like mappers"
    url: "/category/MMC3-like_mappers/"
  - name: "Multicart mappers"
    url: "/category/Multicart_mappers/"
---

NES 2.0 Mapper 269 is used for at least one plug-and-play console and two multicarts with an MMC3 clone and four additional outer bank registers:

- *Games Xplosion 121-in-1*
- *15000-in-1*
- *18000-in-1*

PRG and CHR fetches are mapped to different areas of a single address space, similar to the [VTxx](../VTxx/) consoles in OneBus mode. [As the plug-and-play console's square wave duty cycles are not reversed](https://youtu.be/CFZ9yeS6984?t=31s) however, the actual hardware cannot be VTxx.

# Registers {#Registers}

## MMC3-compatible registers {#MMC3-compatible_registers}

Mask: probably \$E001

    $8000, $8001, $A000, $A001, $C000, $C001, $E000, $E001: As normal MMC3.

## <span id="Outer_bank_registers_.28.245000.29"></span>Outer bank registers (\$5000) {#Outer_bank_registers_.28.245000.29}

There are four outer bank registers, all of which are accessed at address \$5000. The first write goes to the first register, the second write to the second, and so on; the fifth write goes to the first register again.

Mask: unknown, but probably \$F008, as a few games' writes to \$5008 must be ignored

Register 0:

    7654 3210
    ---------
    CCCC CCCC
    ++++-++++- Bits 0-7 of the 1 KiB Outer CHR Bank number

Register 1:

    7654 3210
    ---------
    PPPP PPPP
    ++++-++++- Bits 0-7 of the 8 KiB Outer PRG Bank number

Register 2:

    7654 3210
    ---------
    CCCC MMMM
    |||| ++++- Outer CHR Bank mask in number of most significant bits,
    ||||       e.g. 0=0x00, 1=0x80, 2=0xC0, 3=0xE0, 4=0xF0
    ++++------ Bits 8-11 of the 1 KiB Outer CHR Bank number

Register 3:

    7654 3210
    ---------
    OOMM MMMM
    ||++-++++- Outer PRG Bank mask (bit mask, not number of bits)
    ++-------- Bits 8/9 of the 8 KiB Outer PRG Bank number, also
               Bits 12/13 of the 1 KiB Outer CHR Bank number

The final 8 KiB PRG and 1 KiB CHR bank number are therefore:

    PRGBank := (MMC3PRGBank &~OuterPRGBankMask) | (OuterPRGBank &OuterPRGBankMask);
    CHRBank := (MMC3CHRBank &~OuterCHRBankMask) | (OuterCHRBank &OuterCHRBankMask);

# CHR Scrambling {#CHR_Scrambling}

The CHR data is scrambled by changing the bit order.

    Stored bit  76543210
                --------
    Actual bit  70615243

Which means that to get from the stored data to the actual data, the following operation must be performed:

    Val =((Val &1) <<6) | ((Val &2) <<3) | ((Val &4) <<0) | ((Val &8) >>3) | ((Val &16) >>3) | ((Val &32) >>2) | ((Val &64) >>1) | ((Val &128) <<0);

# Notes {#Notes}

- Initial register values must be 0x00 for registers 0, 1 and 3, and 0x0F for register 2.
- This mapper seems to be an enhanced version of [Mapper 45](../INES_Mapper_045/).
- This mapper supports WRAM at CPU \$6000-\$7FFF, which is necessary for *18000-in-1'*s *Super Mario Bros. 3*.
- On the *Games Xplosion 121-in-1* plug-and-play, game \#48 (First Defender) enables interrupts but fails to disable the [frame IRQ](../APU_Frame_Counter/), freezing on emulators that accurately [power-on with the frame IRQ enabled](../CPU_power_up_state/) for a normal NES. On real hardware, the game runs properly.
