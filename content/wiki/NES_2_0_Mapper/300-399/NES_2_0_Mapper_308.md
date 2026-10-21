---
title: "NES 2.0 Mapper 308"
url: "/wiki/NES_2.0_Mapper_308/"
description: "NES 2.0 Mapper 308 is used for a bootleg version of the Sunsoft game Batman. It uses a VRC2 clone (VRC2b, CPU A0=chip A0, CPU A1=chip A1, i.e. similar to Mapper 23 Submapper 3) with custom IRQ…"
wikiPage: "NES_2_0_Mapper_308"
wikiCategories:
  - name: "Mappers with cycle IRQs"
    url: "/category/Mappers_with_cycle_IRQs/"
---

NES 2.0 Mapper 308 is used for a bootleg version of the Sunsoft game *Batman*. It uses a VRC2 clone (VRC2b, CPU A0=chip A0, CPU A1=chip A1, i.e. similar to Mapper 23 Submapper 3) with custom IRQ functionality. Its UNIF board name is **UNL-TH2131-1**.

# Registers {#Registers}

## <span id="VRC2-compatible_registers_.28.248000-.24FFFF.29"></span>VRC2-compatible registers (\$8000-\$FFFF) {#VRC2-compatible_registers_.28.248000-.24FFFF.29}

See description at [VRC2](../VRC2_and_VRC4/).

## <span id="IRQ_Acknowledge_and_Reset_.28.24F000.29"></span>IRQ Acknowledge and Reset (\$F000) {#IRQ_Acknowledge_and_Reset_.28.24F000.29}

Mask: \$F003

Acknowledge IRQ and disable IRQ counting; reset low counter to zero.

## <span id="IRQ_Counter_Enable_.28.24F001.29"></span>IRQ Counter Enable (\$F001) {#IRQ_Counter_Enable_.28.24F001.29}

Mask: \$F003

Enable IRQ counting.

## <span id="IRQ_High_Counter_Value_.28.24F003.29"></span>IRQ High Counter Value (\$F003) {#IRQ_High_Counter_Value_.28.24F003.29}

    Mask: $F003

    D~7654 3210
      ---------
      VVVV ....
      ++++----- Value for high counter

# IRQ Operation {#IRQ_Operation}

If IRQ counting is enabled, the 12 bit low counter is increased on every M2 cycle. Upon reaching the value 2048, the high counter is decreased. If the high counter is zero and the low counter is below 2048, an IRQ is asserted.

# See also {#See_also}

[PCB image and discussion](https://forums.nesdev.org/viewtopic.php?f=28&t=13270)
