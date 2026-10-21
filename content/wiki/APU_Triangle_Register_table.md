---
title: "APU Triangle/Register table"
url: "/wiki/APU_Triangle/Register_table/"
description: "The linear counter is an extra hardware duration timer that provides higher accuracy than the length counter. It works completely seperately from length counter."
wikiPage: "APU_Triangle_Register_table"
---

### <span id="Linear_counter_control_.28.244008.2C_write.29"></span>Linear counter control (\$4008, write) {#Linear_counter_control_.28.244008.2C_write.29}

    7  bit  0
    ---- ----
    CRRR RRRR
    |||| ||||
    |+++-++++- Linear counter reload value
    +--------- Linear counter control flag and length counter halt flag

The **linear counter** is an extra hardware duration timer that provides higher accuracy than the [length counter](../../APU_Length_Counter/). It works completely seperately from length counter.

When the [APU Frame Counter](../../APU_Frame_Counter/) generates the quarter frame clock, these actions occur in order:

1.  If the [linear counter reload flag](#Length_counter_load_value_and_timer_period_upper_3_bits_.28.24400B.2C_write.29) is set the counter is reloaded with the counter reload value, otherwise, if the counter is non-zero, it is decremented.
2.  If the counter control flag is clear, the linear counter reload flag is cleared.

Note that the counter reload flag is not cleared unless the control flag is also clear, so when both are already set a value written to \$4008 will be reloaded at the next linear counter clock.

### <span id="Timer_period_lower_8_bits_.28.24400A.2C_write.29"></span>Timer period lower 8 bits (\$400A, write) {#Timer_period_lower_8_bits_.28.24400A.2C_write.29}

    7  bit  0
    ---- ----
    TTTT TTTT
    |||| ||||
    ++++-++++- Lower 8 bits of timer period value

### <span id="Length_counter_load_value_and_timer_period_upper_3_bits_.28.24400B.2C_write.29"></span>Length counter load value and timer period upper 3 bits (\$400B, write) {#Length_counter_load_value_and_timer_period_upper_3_bits_.28.24400B.2C_write.29}

    7  bit  0
    ---- ----
    LLLL LTTT
    |||| ||||
    |||| |+++- Upper 3 bits of timer period value
    ++++-+---- Length counter load value

Writing to this register sets [linear counter reload flag](#Linear_counter_control_.28.244008.2C_write.29) too.
