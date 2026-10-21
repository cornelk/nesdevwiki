---
title: "NES 2.0 Mapper 562"
url: "/wiki/NES_2.0_Mapper_562/"
description: "NES 2.0 Mapper 562 denotes ROM image files extracted from disk images for the Venus Turbo Game Doctor 4+, 6+ and 6M RAM cartridges. They represent games whose Doctor Header file denotes a Turbo Game…"
wikiPage: "NES_2_0_Mapper_562"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

**NES 2.0 Mapper 562** denotes ROM image files extracted from disk images for the *Venus Turbo Game Doctor 4+*, *6+* and *6M* [RAM cartridges](../RAM_cartridge/). They represent games whose [Doctor Header file](../Game_Doctor/Magic_Card_FDS_Format/#Venus_Turbo_Game_Doctor_disks) denotes a Turbo Game Doctor disk (byte \$0 bit 7 clear). A clone of the [Magic Card 2M](../Super_Magic_Card/), it differs in several aspects.

As with [mapper 6](../INES_Mapper_006/), the submapper field denotes the initial latch-based banking mode (0-7).

# Banks {#Banks}

- CPU \$4800-\$4FFF: BIOS offsets \$0800-\$0FFF. Some locations checked by games to lock out competing RAM cartridges from Bung and FFE.
- CPU \$6000-\$7FFF: 8 KiB WRAM
- CPU \$8000-\$FFFF: 32 KiB window into 256 KiB (Turbo Game Doctor 6+)/512 KiB (Turbo Game Doctor 4+/6M) PRG memory
- PPU \$0000-\$1FFF: 8 KiB window into 32 KiB (Turbo Game Doctor 4+)/256 KiB (Turbo Game Doctor 6+/6M) CHR memory

# Trainer {#Trainer}

Loadable to any RAM offset and of arbitrary size, Turbo Game Doctor trainers cannot use the [iNES format's trainer bit](../INES/#Trainer). Instead, its data is included as NES 2.0 [Misc. ROM](../NES_2.0/#Miscellaneous_ROM_Area) data:

- Misc. ROM offset 0: Trainer load address (16-bit little endian)
- Misc. ROM offset 2: JSR address of pre-reset-handler initialization routine, or \$0000 if none (16-bit little endian)
- Misc. ROM offset 4+: Trainer data

# Banking Modes {#Banking_Modes}

## Latch-based modes {#Latch-based_modes}

- A bus-conflict-free latch at CPU \$8000-\$FFFF switches PRG and CHR banks.
- The [1M Mode Register](#1M_banking_mode_.28.2442FC-.2442FF.2C_write-only.29) selects the latch-based mode.
- The latch is only active when PRG memory is write-protected.
- There are no CNROM-128 or NROM-256 modes; BIOS runs such games in CNROM-256 mode by duplicating CHR banks.

### 0: [UNROM](../UxROM/) {#0:_UNROM}

    D~[.... .PPP]
             +++- CPU $8000-$BFFF: switchable 16 KiB PRG bank #0-#7 via PPP
                  CPU $C000-$FFFF: fixed 16 KiB PRG bank #7
                  PPU $0000-$1FFF: fixed 8 KiB CHR bank #0, writable

### <span id="1:_UN1ROM.2BCHRSW"></span>1: [UN1ROM](../INES_Mapper_094/)+CHRSW {#1:_UN1ROM.2BCHRSW}

    D~[..BB BBCC]
         || ||++- PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC, writable
         ++-++--- CPU $8000-$BFFF: switchable 16 KiB PRG bank #0-#15 via PPP
                  CPU $C000-$FFFF: fixed 16 KiB PRG bank #7

UN1ROM+CHRSW is used for games that originally were [SKROM](../MMC1/) with 128 KiB PRG and 128 KiB CHR data. The original 128 KiB CHR data are stored in the second 128 KiB of PRG address space, with a trainer program copying the most-recentl CHR data into 32 KiB CHR memory.

### 2: [UOROM](../UxROM/) {#2:_UOROM}

    D~[.... PPPP]
            ++++- CPU $8000-$BFFF: switchable 16 KiB PRG bank #0-#15 via PPPP
                  CPU $C000-$FFFF: fixed 16 KiB PRG bank #15
                  PPU $0000-$1FFF: fixed 8 KiB CHR bank #0, writable

With no other means of masking PRG addresses, UNROM vs. UOROM are explicitly differentiated.

### <span id="3:_Reverse_UOROM.2BCHRSW"></span>3: [Reverse UOROM](../INES_Mapper_097/)+CHRSW {#3:_Reverse_UOROM.2BCHRSW}

    D~[..CC PPPP]
         || ++++- CPU $C000-$FFFF: switchable 16 KiB PRG bank #0-#15 via PPPP
         ||       CPU $8000-$7FFF: fixed 16 KiB PRG bank #15
         ++------ PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC, writable

### 4: [GNROM](../GxROM/) {#4:_GNROM}

    D~[..PP ..CC]
         ||   ++- PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC, write-protected
         ++------ CPU $8000-$FFFF: switchable 32 KiB PRG bank #0-#3 via PP

### 5: [CNROM-256](../CNROM/) {#5:_CNROM-256}

    D~[.... ..CC]
              ++- PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC, write-protected
                  CPU $8000-$FFFF: fixed 32 KiB PRG bank #3

### <span id="6:_Custom_.231"></span>6: Custom \#1 {#6:_Custom_.231}

    D~[pppp PPPP]
       |||| ++++- CPU $8000-$9FFF: switchable 8 KiB PRG bank #0-15 via PPPP
       ++++------ CPU $A000-$BFFF: switchable 8 KiB PRG bank #0-15 via pppp
                  CPU $C000-$FFFF: fixed 16 bit PRG bank #7
                  PPU $0000-$1FFF: fixed 8 KiB CHR bank #0-3 at last-selected value in another mode

This mode is used by a few games that originally ran on the [Namcot 118](../INES_Mapper_206/) ASIC.

### <span id="7:_Custom_.232"></span>7: Custom \#2 {#7:_Custom_.232}

    D~[ppp1 PPP0]
       |||| ++++- CPU $8000-$9FFF: switchable 8 KiB PRG bank #0-14 via PPP0
       ++++------ CPU $A000-$BFFF: switchable 8 KiB PRG bank #1-15 via ppp1
                  CPU $C000-$FFFF: fixed 16 bit PRG bank #7
                  PPU $0000-$1FFF: fixed 8 KiB CHR bank #0-3 at last-selected value in another mode

## 2M PRG banking mode {#2M_PRG_banking_mode}

- Four registers at \$8000-\$9FFF/\$A000-\$BFFF/\$C000-\$DFFF/\$E000-\$FFFF switch four 8 KiB PRG banks and the common 8 KiB CHR bank. They overlap with the latch at \$8000-\$FFFF and accept values even when the 2M PRG banking mode itself is not active.
- 2M PRG banking mode is enabled by the [2M Mode Register](#2M_PRG_banking_mode_.28.2443FE-.2443FF.2C_write-only.29). When active, it overrides the Latch-based mode in all aspects except CHR memory protection, which is still decided by the [1M Mode Register](#1M_banking_mode_.28.2442FC-.2442FF.2C_write-only.29).
- Only 256 KiB can be reached.
- Only the four lowest address bits can be switched on an 8 KiB level; the 128 KiB bit is switched for the entire \$8000-\$FFFF range via \$43FE.

<!-- -->

    D~[..PP PPCC]
         || ||++- PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC in $8000-$FFFF
         ++-++--- CPU $8000-$9FFF: switchable 8 KiB PRG Bank #0-31 via PPPP in $8000-$9FFF (LSB) and p (MSB) in $43FE
                  CPU $A000-$BFFF: switchable 8 KiB PRG Bank #0-31 via PPPP in $A000-$BFFF (LSB) and p (MSB) in $43FE
                  CPU $C000-$DFFF: switchable 8 KiB PRG Bank #0-31 via PPPP in $C000-$DFFF (LSB) and p (MSB) in $43FE
                  CPU $E000-$FFFF: switchable 8 KiB PRG Bank #0-31 via PPPP in $E000-$FFFF (LSB) and p (MSB) in $43FE

## 4M PRG banking mode {#4M_PRG_banking_mode}

- Identical in operation to [2M PRG Banking Mode](#2M_PRG_banking_mode), except that all six bits from the four registers at \$8000-\$9FFF/\$A000-\$BFFF/\$C000-\$DFFF/\$E000-\$FFFF are used.
- 4M PRG banking mode is enabled by the [TGD Mode Register](#Turbo_Game_Doctor_mode_.28.244411.2C_read.2Fwrite.29). When active, it overrides the Latch-based mode in all aspects except CHR memory protection, which is still decided by the [1M Mode Register](#1M_banking_mode_.28.2442FC-.2442FF.2C_write-only.29), and the 2M PRG banking mode.
- The entire 512 KiB can be reached if available.

<!-- -->

    D~[PPPP PPCC]
       |||| ||++- PPU $0000-$1FFF: switchable 8 KiB CHR bank #0-#3 via CC in $8000-$FFFF
       ++++-++--- CPU $8000-$9FFF: switchable 8 KiB PRG Bank #0-63 via PPPPPP in $8000-$9FFF
                  CPU $A000-$BFFF: switchable 8 KiB PRG Bank #0-63 via PPPPPP in $A000-$BFFF
                  CPU $C000-$DFFF: switchable 8 KiB PRG Bank #0-63 via PPPPPP in $C000-$DFFF
                  CPU $E000-$FFFF: switchable 8 KiB PRG Bank #0-63 via PPPPPP in $E000-$FFFF

## 1 KiB CHR banking mode {#1_KiB_CHR_banking_mode}

- Eight registers at \$4400-\$4407 switch eight 1 KiB CHR banks. The registers are readable as well.
- 1 KiB CHR banking mode is enabled via the [TGD Mode Register](#Turbo_Game_Doctor_mode_.28.244411.2C_read.2Fwrite.29). register. When active, it overrides the CC bits that selected 8 KiB banks from 32 KiB.
- Although the Turbo Game Doctor 4+ only has 32 KiB of CHR memory, 1 KiB CHR banking mode still exists, allowing the 32 KiB to be banked with finer granularity than on previous Bung and FFE models.

<!-- -->

    D~[CCCC CCCC]
       ++++-++++- PPU $0000-$03FF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4400
                  PPU $0400-$07FF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4401
                  PPU $0800-$0BFF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4402
                  PPU $0C00-$0FFF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4403
                  PPU $1000-$13FF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4404
                  PPU $1400-$17FF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4405
                  PPU $1800-$1BFF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4406
                  PPU $1C00-$1FFF: switchable 1 KiB CHR Bank #0-255 via CCCCCCC in $4407

# Interrupts {#Interrupts}

The Turbo Game Doctors have a cycle-based IRQ counter that operates in an unusual fashion. After resetting the counter at the start of the NMI handler, a game writes the target value at which an IRQ is to be generated. If a second IRQ is desired further down the frame, only a new target value is written while keeping the counter running. This approach avoids accumulating timing errors from delays and jitter.

# Registers {#Registers}

## <span id="1M_banking_mode_.28.2442FC-.2442FF.2C_write-only.29"></span>1M banking mode (\$42FC-\$42FF, write-only) {#1M_banking_mode_.28.2442FC-.2442FF.2C_write-only.29}

    A~FEDC BA98 7654 3210  D~7654 3210
      -------------------    ---------
      0100 0010 1111 11bM    BBBM ....
                       |+----|||+------ Set nametable mirroring type
                       |     |||         0: One-screen, page 0
                       |     |||         1: One-screen, page 1
                       |     |||         2: Vertical
                       |     |||         3: Horizontal
                       +-----|||------- 0: PRG memory is writeable, latch is disabled
                             |||        1: PRG memory is write-protected, latch is enabled
                             +++------- Select latch-based banking mode

## <span id="2M_PRG_banking_mode_.28.2443FE-.2443FF.2C_write-only.29"></span>2M PRG banking mode (\$43FE-\$43FF, write-only) {#2M_PRG_banking_mode_.28.2443FE-.2443FF.2C_write-only.29}

    A~FEDC BA98 7654 3210  D~7654 3210
      -------------------    ---------
      0100 0011 1111 111M    .p.. ..CC
                        |     |     ++- switchable 8 KiB CHR bank #0-#3 via CC
                        |     +-------- PRG A17 for CPU $8000-$FFFF
                        +- 2M PRG banking mode select
                            0: Enable
                            1: Disable

The CC bits are mirrors of those in the latch at \$8000-\$FFFF.

## <span id="Turbo_Game_Doctor_mode_.28.244411.2C_read.2Fwrite.29"></span>Turbo Game Doctor mode (\$4411, read/write) {#Turbo_Game_Doctor_mode_.28.244411.2C_read.2Fwrite.29}

    D~7654 3210
      ---------
      41p. GAWW
      |||  ||++- CPU $6000-$7FFF: switchable 8 KiB WRAM bank #0-3 via WW.
      |||  ||    Only usable in floppy disk load mode, and ignored in gameplay mode.
      |||  |+--- Map WRAM in floppy disk load mode. Ignored in gameplay mode.
      |||  +---- Game Saver attachment is active. Related to real-time saving.
      ||+------- PRG A18 in floppy disk mode. Ignored in gameplay mode.
      |+-------- 1=1 KiB CHR banking mode enabled
      +--------- 1=4M PRG banking mode enabled

## <span id="IRQ_Counter_Target_LSB.2FMSB_.28.24440D.2F.24440C.2C_write-only.29"></span>IRQ Counter Target LSB/MSB (\$440D/\$440C, write-only) {#IRQ_Counter_Target_LSB.2FMSB_.28.24440D.2F.24440C.2C_write-only.29}

- This upwards-counting signed 16-bit counter, when negative, is increased on every M2 or PA12 rise and raises an IRQ when reaching the target value, continuing to count until it reaches \$FFFF.
- Writing to either register acknowledges a pending IRQ.
- The counter is reset and disabled by writing a positive value to \$440C.

## <span id="IRQ_Counter_Value_LSB.2FMSB_.28.24440D.2F.24440C.2C_read-only.29"></span>IRQ Counter Value LSB/MSB (\$440D/\$440C, read-only) {#IRQ_Counter_Value_LSB.2FMSB_.28.24440D.2F.24440C.2C_read-only.29}

The current counter value can be read from the same addresses that the target counter value is written to.

## <span id="Current_CHR_Bank_.28.244420.2C_read-only.29"></span>Current CHR Bank (\$4420, read-only) {#Current_CHR_Bank_.28.244420.2C_read-only.29}

The value of the current 1 KiB CHR bank register can be read and seen changing as the PPU renders the screen. The BIOS uses this in test mode to determine correct operation, and some games check the result to lock out competing products from Bung and FFE.
