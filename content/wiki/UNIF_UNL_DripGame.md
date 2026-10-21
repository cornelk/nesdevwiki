---
title: "UNIF/UNL-DripGame"
url: "/wiki/UNIF/UNL-DripGame/"
description: "UNL-DripGame is an FPGA mapper comparable in capability to ASIC mappers, invented by Quietust for a port of a game."
wikiPage: "UNIF_UNL_DripGame"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

UNL-DripGame is an FPGA mapper comparable in capability to ASIC mappers, invented by Quietust for a [port of a game](https://www.qmtpro.com/~nes/drip/).

[NES 2.0 Mapper 284](../../NES_2.0_Mapper_284/) is assigned to this game, as was [iNES Mapper 102](../../INES_Mapper_102/). The UNIF name is "UNL-DripGame".

## Memory map {#Memory_map}

CPU:

- \$4800-\$5FFF = Readable registers
- \$6000-\$7FFF = PRG RAM
- \$8000-\$BFFF = Switchable bank
- \$C000-\$FFFF = Fixed to the last bank

## Readable registers {#Readable_registers}

### <span id=".244800-.244FFF_.28Status.29"></span>\$4800-\$4FFF (Status) {#.244800-.244FFF_.28Status.29}

    [DSSS SSSS]
     |||| ||||
     |+++-++++---- Ready flag; $64 if ready
     +------------ DIP switch

Other registers must not be read or written until the mapper is ready. The reference implementation of the mapper is an FPGA that loads its fusemap from another on-board memory, and until the FPGA signals that it is ready, its behavior is undefined other than that the fixed PRG bank is fixed.

### <span id=".245000-.2457FF_.28Sound_channel_0_status.29"></span>\$5000-\$57FF (Sound channel 0 status) {#.245000-.2457FF_.28Sound_channel_0_status.29}

    [FE.. ....]
     ||
     |+----------- Buffer is empty
     +------------ Buffer is full

### <span id=".245800-.245FFF_.28Sound_channel_1_status.29"></span>\$5800-\$5FFF (Sound channel 1 status) {#.245800-.245FFF_.28Sound_channel_1_status.29}

Similar to \$5000 but for the second channel.

## Writable registers {#Writable_registers}

All of these registers are mirrored at \$8000-\$BFFF; the low four bits select the register number.

### <span id=".248000_.28Sound_channel_0_silence_register.29"></span>\$8000 (Sound channel 0 silence register) {#.248000_.28Sound_channel_0_silence_register.29}

Writing anything silences channel 0 and clears the buffer.

### <span id=".248001_.28Sound_channel_0_data.29"></span>\$8001 (Sound channel 0 data) {#.248001_.28Sound_channel_0_data.29}

Writing data inserts it into the buffer.

### <span id=".248002_.28Sound_channel_0_low_period.29"></span>\$8002 (Sound channel 0 low period) {#.248002_.28Sound_channel_0_low_period.29}

    [PPPP PPPP]

Specify low 8-bits of period of channel 0.

### <span id=".248003_.28Sound_channel_0_high_period.29"></span>\$8003 (Sound channel 0 high period) {#.248003_.28Sound_channel_0_high_period.29}

    [VVVV PPPP]
     |||| ||||
     |||| ++++---- High 4-bits of period
     ++++--------- Volume control

Specify high 4-bits of period of channel 0, and volume control.

### <span id=".248004_.28Sound_channel_1_silence_register.29"></span>\$8004 (Sound channel 1 silence register) {#.248004_.28Sound_channel_1_silence_register.29}

Writing anything silences channel 1 and clears the buffer.

### <span id=".248005_.28Sound_channel_1_data.29"></span>\$8005 (Sound channel 1 data) {#.248005_.28Sound_channel_1_data.29}

Writing data inserts it into the buffer.

### <span id=".248006_.28Sound_channel_1_low_period.29"></span>\$8006 (Sound channel 1 low period) {#.248006_.28Sound_channel_1_low_period.29}

    [PPPP PPPP]

Specify low 8-bits of period of channel 1.

### <span id=".248007_.28Sound_channel_1_high_period.29"></span>\$8007 (Sound channel 1 high period) {#.248007_.28Sound_channel_1_high_period.29}

    [VVVV PPPP]
     |||| ||||
     |||| ++++---- High 4-bits of period
     ++++--------- Volume control

Specify high 4-bits of period of channel 1, and volume control.

### <span id=".248008_.28Low_IRQ_register.29"></span>\$8008 (Low IRQ register) {#.248008_.28Low_IRQ_register.29}

    [XXXX XXXX]
     |||| ||||
     ++++-++++---- Low 8-bits of IRQ value

### <span id=".248009_.28High_IRQ_register.29"></span>\$8009 (High IRQ register) {#.248009_.28High_IRQ_register.29}

    [EXXX XXXX]
     |||| ||||
     |+++-++++---- High 7-bits of IRQ value
     +------------ Enable flag

IRQ is automatically disabled when it is overflows. Writes to this register additionally acknowledge the interrupt.

### <span id=".24800A_.28Control.29"></span>\$800A (Control) {#.24800A_.28Control.29}

    [.... SAMM]
          ||||
          ||++---- Nametable mirroring (00=V, 01=H, 10=1sc-A, 11=1sc-B)
          |+------ Enable extended attribute table
          +------- Enable writing to SRAM

### <span id=".24800B_.28PRG_bank.29"></span>\$800B (PRG bank) {#.24800B_.28PRG_bank.29}

    [.... PPPP]
          ||||
          ++++---- Select 16K PRG bank at $8000-$BFFF

### <span id=".24800C_.28CHR_bank_0.29"></span>\$800C (CHR bank 0) {#.24800C_.28CHR_bank_0.29}

    [.... CCCC]
          ||||
          ++++---- Select 2K CHR bank at $0000-$07FF

### <span id=".24800D_.28CHR_bank_1.29"></span>\$800D (CHR bank 1) {#.24800D_.28CHR_bank_1.29}

    [.... CCCC]
          ||||
          ++++---- Select 2K CHR bank at $0800-$0FFF

### <span id=".24800E_.28CHR_bank_2.29"></span>\$800E (CHR bank 2) {#.24800E_.28CHR_bank_2.29}

    [.... CCCC]
          ||||
          ++++---- Select 2K CHR bank at $1000-$17FF

### <span id=".24800F_.28CHR_bank_3.29"></span>\$800F (CHR bank 3) {#.24800F_.28CHR_bank_3.29}

    [.... CCCC]
          ||||
          ++++---- Select 2K CHR bank at $1800-$1FFF

## Extended attribute tables {#Extended_attribute_tables}

This memory is write-only and is mapped at \$C000-\$C3FF for first nametable, and \$C400-\$C7FF for last nametable. It is mirrored across \$C000-\$FFFF.

Only the two LSB of data are used, and they control the palette of the corresponding tile position in a nametable. (i.e. the bits written to CPU \$C523 correspond to the tile at CIRAM \$523, which if mirrored horizontally would correspond to the tile at PPU \$2923 and \$2D23)

## IRQ {#IRQ}

Data written to the low IRQ register is stored in a buffer but won't reload the counter until the high IRQ register is programmed.

IRQ is count down per cycle, and then trigger the IRQ and stop the timer when it expires.

## Audio {#Audio}

There are two channels, both of which function in the same way.

The buffers are "first in first out" buffers which store unsigned sample data with a bias of -\$80.

Each buffer stores 256 bytes which can be written in sequence. If a channel is inactive, filling the buffer with even one byte activates that channel immediately.

Changes to the period are delayed until the sample is finished playing. However, the volume control can be changed at any time.

The channel becomes inactive as soon as the buffer becomes empty.

## See also {#See_also}

- [Quietust's mapper document](https://www.qmtpro.com/~nes/drip/dripmap.txt)
