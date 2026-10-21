---
title: "APU Pulse/Register table"
url: "/wiki/APU_Pulse/Register_table/"
description: "Register $4000, $4001, $4002 and $4003 correspond to the first Pulse channel, $4004, $4005, $4006, $4007 correspond to the second Pulse channel."
wikiPage: "APU_Pulse_Register_table"
---

Register [\$4000](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29), [\$4001](#Sweep_control_.28.244001_.2F_.244005.2C_write.29), [\$4002](#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29) and [\$4003](#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29) correspond to the first Pulse channel, [\$4004](#Pulse_settings_.28.244000_.2F_.244004.2C_write.29), [\$4005](#Sweep_control_.28.244001_.2F_.244005.2C_write.29), [\$4006](#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29), [\$4007](#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29) correspond to the second Pulse channel.

### <span id="Pulse_settings_.28.244000_.2F_.244004.2C_write.29"></span>Pulse settings (\$4000 / \$4004, write) {#Pulse_settings_.28.244000_.2F_.244004.2C_write.29}

    7  bit  0
    ---- ----
    DDLC VVVV
    |||| ||||
    |||| ++++- When C is set: channel volume value
    ||||       When C is clear: envelope divider value
    |||+------ Constant volume flag (override envelope with V)
    ||+------- Length counter halt flag and envelope loop flag
    ++-------- Pulse width
                   0 = 12.5%
                   1 = 25.0%
                   2 = 50.0%
                   3 = 75.0%

### <span id="Sweep_control_.28.244001_.2F_.244005.2C_write.29"></span>Sweep control (\$4001 / \$4005, write) {#Sweep_control_.28.244001_.2F_.244005.2C_write.29}

    7  bit  0
    ---- ----
    EPPP NSSS
    |||| ||||
    |||| |+++- Bit shift count
    |||| +---- Negate flag
    ||||           0: add to period, sweeping toward lower frequencies
    ||||           1: subtract from period, sweeping toward higher frequencies
    |+++------ Clock counter divider value
    +--------- Enable flag

Writing to this register also sets sweep reload flag.

See [APU Sweep](../).

### <span id="Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29"></span>Timer period lower 8 bits (\$4002 / \$4006, write) {#Timer_period_lower_8_bits_.28.244002_.2F_.244006.2C_write.29}

    7  bit  0
    ---- ----
    TTTT TTTT
    |||| ||||
    ++++-++++- Lower 8 bits of timer period value

### <span id="Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29"></span>Length counter load value and timer period upper 3 bits (\$4003 / \$4007, write) {#Length_counter_load_value_and_timer_period_upper_3_bits_.28.244003_.2F_.244007.2C_write.29}

    7  bit  0
    ---- ----
    LLLL LTTT
    |||| ||||
    |||| |+++- Upper 3 bits of timer period value
    ++++-+---- Length counter load value

Writing to this register also resets the sequencer index (i.e., the pulse phase) to the first value, restart [envelope](../../APU_Envelope/).
