---
title: "VT02+ MMC3 Compatibility Registers"
url: "/wiki/VT02+_MMC3_Compatibility_Registers/"
description: "In OneBus mode, VT02+ consoles combine the Famicom cartridge connector's CPU and PPU address lines into one 32 MiB address space, with separate CPU and PPU bankswitch registers pointing to…"
wikiPage: "VT02__MMC3_Compatibility_Registers"
---

In OneBus mode, VT02+ consoles combine the Famicom cartridge connector's CPU and PPU address lines into one 32 MiB address space, with separate CPU and PPU bankswitch registers pointing to appropriately-placed code and picture data. The bankswitching scheme is based on, and indeed backwards-compatible to, the Nintendo MMC3's. Backwards compatibility is realized by way of forwarding CPU writes to \$8000-\$FFFF to the appropriate VT02+ register. V.R. Technology's data sheets call this the "old compatible mode". Forwarding can be disabled by setting register \$401B bit 3 (FWEN).

When \$410B bit 3 (FWEN) is cleared, the [MMC3](../MMC3/)'s registers are forwarded to VT02+ registers as follows:

    MMC3 register   VT02+ register          Meaning
    -------------   --------------          -------
    $8000           $4105, except bit 5     MMC3 Index Register ("Pointer"), PRG A14 inversion, CHR A12 inversion
    $8001  reg
            0       $2016                   1 KiB CHR Bank Number at PPU $0000-$03FF (AND $FE) and PPU $0400-$07FF (OR $01)
            1       $2017                   1 KiB CHR Bank Number at PPU $0400-$07FF (AND $FE) and PPU $0800-$0FFF (OR $01)
            2       $2012                   1 KiB CHR Bank Number at PPU $1000-$13FF
            3       $2013                   1 KiB CHR Bank Number at PPU $1400-$17FF
            4       $2014                   1 KiB CHR Bank Number at PPU $1800-$1BFF
            5       $2015                   1 KiB CHR Bank Number at PPU $1C00-$1FFF
            6       $4107                   8 KiB PRG Bank Number at CPU $8000-$9FFF
            7       $4108                   8 KiB PRG Bank Number at CPU $A000-$BFFF
    $A000           $4106, only bit 0       Mirroring select
    $A001           not forwarded           WRAM Enable/Write Protect
    $C000           $4101                   IRQ Latch
    $C001           $4102                   Reload IRQ Counter
    $E000           $4103                   Disable/Acknowledge IRQ
    $E001           $4104                   Enable IRQ

Notes:

- The table disregards mirrored MMC3 addresses.
- Writing a register number to the lower three bits of \$4105 also decides which register is affected when writing data to \$8001, as \$4105 is indeed a mirror of \$8000 (excluding bit 5) and vice-versa. EmuVT does not take this into account, resulting in glitches when running games such as *Mighty Bomb Jack* on the *D-CAT8* handheld's compilation(s).
- The JungleTac game *Bolt Fighter* on the *Classic Max Lite* handheld's compilation writes garbage data to the \$8000-\$FFFF range during gameplay. This would lead to inadvertent CHR bank changes, were it not for the fact that the menu sets FWEN to 1 to disable forwarding.
- As the VT02+ consoles themselves have no WRAM at \$6000-\$7FFF, there is no register to which a \$A001 write could be forwarded.
