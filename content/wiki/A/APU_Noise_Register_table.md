---
title: "APU Noise/Register table"
url: "/wiki/APU_Noise/Register_table/"
description: "The timer period is set to entry P of the following:"
wikiPage: "APU_Noise_Register_table"
---

### <span id="Noise_envelope_settings_.28.24400C.2C_write.29"></span>Noise envelope settings (\$400C, write) {#Noise_envelope_settings_.28.24400C.2C_write.29}

    7  bit  0
    ---- ----
    ..LC VVVV
      || ||||
      || ++++- When C is set: channel volume value
      ||       When C is clear: envelope divider reload value
      |+------ Constant volume flag (override envelope with V)
      +------- Length counter halt flag and envelope loop flag

### <span id="Noise_mode_and_period_.28.24400E.2C_write.29"></span>Noise mode and period (\$400E, write) {#Noise_mode_and_period_.28.24400E.2C_write.29}

    7  bit  0
    ---- ----
    M... PPPP
    |    ||||
    |    ++++- Noise period
    +--------- Periodic noise mode

The timer period is set to entry `P` of the following:

| Period (`P`) | NTSC | PAL |
|----|----|----|
| \$0 | 4 | 4 |
| \$1 | 8 | 8 |
| \$2 | 16 | 14 |
| \$3 | 32 | 30 |
| \$4 | 64 | 60 |
| \$5 | 96 | 88 |
| \$6 | 128 | 118 |
| \$7 | 160 | 148 |
| \$8 | 202 | 188 |
| \$9 | 254 | 236 |
| \$A | 380 | 354 |
| \$B | 508 | 472 |
| \$C | 762 | 708 |
| \$D | 1016 | 944 |
| \$E | 2034 | 1890 |
| \$F | 4068 (2046 in [early RP2A03 revisions](../../CPU_variants/#Official_.28NTSC.29)) | 3778 |

The period determines how many CPU cycles happen between shift register clocks. These periods are all even numbers because there are 2 CPU cycles in an APU cycle.

### <span id="Length_counter_load_value_.28.24400F.2C_write.29"></span>Length counter load value (\$400F, write) {#Length_counter_load_value_.28.24400F.2C_write.29}

    7  bit  0
    ---- ----
    LLLL L...
    |||| |
    ++++-+---- Length counter load value

Writing to this register also restarts [envelope](../../APU_Envelope/).
