---
title: "Programming NROM"
url: "/wiki/Programming_NROM/"
description: "NROM and the other boards that make up mapper 0 are the simplest of all NES cartridge boards. All address decoding and chip enable handling are handled by the NES hardware; the only integrated…"
wikiPage: "Programming_NROM"
---

[NROM](../NROM/) and the other boards that make up **mapper 0** are the simplest of all NES cartridge boards. All address decoding and chip enable handling are handled by the NES hardware; the only integrated circuits on the board are the PRG-ROM, CHR-ROM, and (in 72-pin carts) the [CIC](../CIC_lockout_chip/). This simplicity lends itself well to small-scale, beginner-level tutorials and projects.

## Configuration {#Configuration}

NROM has two major board configurations:

- NROM-256 with 32KiB PRG-ROM and 8KiB CHR-ROM
- NROM-128 with 16KiB PRG-ROM and 8KiB CHR-ROM

Your program is mapped into \$8000-\$FFFF (NROM-256) or [both](../Mirroring/) \$8000-\$BFFF and \$C000-\$FFFF (NROM-128). Most NROM-128 games actually run in \$C000-\$FFFF rather than \$8000-\$BFFF because it makes the program easier to assemble and link. Some kinds of data used by the NES CPU, such as the interrupt vectors and [DPCM samples](../APU_DMC/), have to be in \$C000-\$FFFF, and it simplifies the linker script if everything is in the same memory region. There are probably a few games that rely on the mirroring, but experiments with [a multicart engine](../Forbidden_Four/) show that most can run with garbage in \$8000-\$BFFF.

## NES 2.0 header {#NES_2.0_header}

Below is an example [NES 2.0](../NES_2.0/) header for NROM boards. It should be backward-compatible with emulators supporting only the older [iNES](../INES/) header format, but they may emulate [extra RAM at \$6000-\$7FFF](../PRG_RAM_circuit/), whereas official boards (except for *Family BASIC*) have [open bus](../Open_bus_behavior/) instead.

    .segment "HEADER"
      .byte "NES", $1A
      .byte 2         ; PRG-ROM size: 1 or 2 x 16KiB for NROM-128 or NROM-256 respectively
      .byte 1         ; CHR-ROM size: 1 x 8KiB
      .byte $00       ; Mapper 0; No battery; Vertical arrangement ("Horizontal mirroring")
      .byte $08       ; Mapper 0; NES 2.0 identifier
      .byte $00       ; No submapper
      .byte $00       ; PRG-ROM not 4 MiB or larger
      .byte $00       ; No PRG-RAM
      .byte $00       ; No CHR-RAM
      .byte $00       ; CPU/PPU timing: NTSC
      .byte $00       ; Not a Vs. System or extended console type
      .byte $00       ; No misc. ROMs
      .byte $00       ; Default expansion device: Unspecified

To use horizontal arrangement ("vertical [mirroring](../Mirroring/)") instead of vertical arrangement ("horizontal mirroring"), change the first `.byte $00` after the CHR-ROM size field to `.byte $01`.

## Common problems {#Common_problems}

Here are some common problems encountered in NROM (or similar discrete logic mapper) projects and their associated workarounds:

- No on-board [IRQ](../IRQ/) source for scroll splits or other raster effects - Use [Sprite-0 hit](../PPU_OAM/). [DMC IRQ](../APU_DMC/#Usage_of_DMC_for_syncing_to_video), [Sprite overflow](../PPU_registers/#Sprite_overflow_flag), and/or [cycle-timed code](../Cycle_counting/) can also work but are trickier to get right.
- No CHR-ROM banking for tile animations - Use [PPUCTRL](../PPU_registers/) bits 4-3 as a "poor man's [CNROM](../CNROM/)" by swapping the pattern table access between PPU \$0000 and \$1000.
- No game save capabilities (outside of *Family BASIC*) - Use a password system.

More advanced requirements, such as fine-grained PRG/CHR bankswitching, nametable arrangement switching at runtime, or IRQ generation, will necessitate a mapper change (sometimes to an ASIC mapper). The project scope and budget should be evaluated beforehand if this is the case, especially if [publishers and boardmakers](../Publishers_and_boardmakers/) are involved.

## Example programs {#Example_programs}

CA65:

- <https://github.com/pinobatch/nrom-template> - Contains linker configurations for both NROM-128 and NROM-256.
