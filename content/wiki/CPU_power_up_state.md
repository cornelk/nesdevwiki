---
title: "CPU power up state"
url: "/wiki/CPU_power_up_state/"
description: "Initial tests on the power-up/reset state of the CPU/APU and RAM contents were done using an NTSC front-loading NES from 1988 with a RP2A03G CPU on the NES-CPU-07 board revision."
wikiPage: "CPU_power_up_state"
---

Initial tests on the power-up/reset state of the CPU/APU and RAM contents were done using an NTSC front-loading NES from 1988 with a RP2A03G CPU on the NES-CPU-07 board revision.

Countless bugs in [commercial](../Game_bugs/) and [homebrew](../Program_compatibility/) games exist because of a reliance on the initial system state. An NES programmer should not rely on the state of CPU/APU registers and RAM contents not guaranteed at power-up/reset.

## CPU {#CPU}

| Register | At Power | After Reset |
|----|----|----|
| A, X, Y | 0 | unchanged |
| PC | (\$FFFC) | (\$FFFC) |
| [S](../Stack/)<sup><span id="cite_ref-reset-stack-push_1-0"></span>[\[1\]](#cite_note-reset-stack-push-1)</sup> | \$00 - 3 = \$FD | S -= 3 |
| [C](../Status_flags/#C:_Carry) | 0 | unchanged |
| [Z](../Status_flags/#Z:_Zero) | 0 | unchanged |
| [I](../Status_flags/#I:_Interrupt_Disable) | 1 | 1 |
| [D](../Status_flags/#D:_Decimal) | 0 | unchanged |
| [V](../Status_flags/#V:_Overflow) | 0 | unchanged |
| [N](../Status_flags/#N:_Negative) | 0 | unchanged |

Initial [CPU Register](../CPU_registers/) Values

## APU {#APU}

| Register | At Power | After Reset |
|----|----|----|
| [Pulses](../APU_Pulse/) (\$4000-\$4007) | 0 | unchanged? |
| [Triangle](../APU_Triangle/) (\$4008-\$400B) | 0 | unchanged? |
| [Triangle](../APU_Triangle/) phase | ? | 0 (output = 15) |
| [Noise](../APU_Noise/) (\$400C-\$400F) | 0 | unchanged? |
| [Noise](../APU_Noise/) 15-bit LFSR | \$0000 (all 0s, first clock shifts in a 1)<sup><span id="cite_ref-2"></span>[\[2\]](#cite_note-2)</sup> | unchanged? |
| [DMC](../APU_DMC/) flags and rate (\$4010)<sup><span id="cite_ref-ebd_3-0"></span>[\[3\]](#cite_note-ebd-3)</sup> | 0 | unchanged |
| [DMC](../APU_DMC/) direct load (\$4011)<sup><span id="cite_ref-ebd_3-1"></span>[\[3\]](#cite_note-ebd-3)</sup> | 0 | \[\$4011\] &= 1 |
| [DMC](../APU_DMC/) sample address (\$4012)<sup><span id="cite_ref-ebd_3-2"></span>[\[3\]](#cite_note-ebd-3)</sup> | 0 | unchanged |
| [DMC](../APU_DMC/) sample length (\$4013)<sup><span id="cite_ref-ebd_3-3"></span>[\[3\]](#cite_note-ebd-3)</sup> | 0 | unchanged |
| [DMC](../APU_DMC/) LFSR | 0? (revision-dependent?) | ? (revision-dependent?) |
| [Status](../APU/#Status_.28.244015.29) (\$4015) | 0 (all channels disabled) | 0 (all channels disabled) |
| [Frame Counter](../APU_Frame_Counter/) (\$4017) | 0 (**frame IRQ enabled**) | unchanged |
| [Frame Counter](../APU_Frame_Counter/) LFSR<sup><span id="cite_ref-.244017_lfsr_4-0"></span>[\[4\]](#cite_note-.244017_lfsr-4)</sup> | \$7FFF (all 1s) | revision-dependent |

Initial [APU](../APU/) Register Values

### Revision-dependent Register Values {#Revision-dependent_Register_Values}

| Register | At Power | After Reset |
|----|----|----|
| [DMC](../APU_DMC/) LFSR | 0? | ? |
| [Frame Counter](../APU_Frame_Counter/) LFSR<sup><span id="cite_ref-.244017_lfsr_4-1"></span>[\[4\]](#cite_note-.244017_lfsr-4)</sup> | \$7FFF (all 1s) | unchanged |

2A03 letterless

| Register | At Power | After Reset |
|----|----|----|
| [DMC](../APU_DMC/) LFSR | 0? | ? |
| [Frame Counter](../APU_Frame_Counter/) LFSR<sup><span id="cite_ref-.244017_lfsr_4-2"></span>[\[4\]](#cite_note-.244017_lfsr-4)</sup> | \$7FFF (all 1s) | \$7FFF (all 1s) |

2A03E, 2A03G, 2A07, various clones

## RAM contents {#RAM_contents}

Internal RAM (\$0000-\$07FF) and cartridge RAM (usually \$6000–\$7FFF, depends on mapper) have an unreliable state on power-up and is unchanged after a reset. Some machines may have consistent RAM contents at power-up, but others may not. Emulators often implement a consistent RAM startup state (e.g. all \$00 or \$FF, or a particular pattern), and [flashcarts](../../category/Flash_Cartridge/) may partially or fully initialize RAM before starting a program.

Battery-backed save RAM and other types of SRAM/NVRAM have an unreliable state on the first power-up and is generally unchanged after subsequent resets and power-ups. However, there is an added chance of data corruption due to loss of power or other external factors (bugs, cheats, etc). Emulators and flashcarts may initialize save files with a consistent state (much like other sections of RAM) and persist this data without corruption after closing or reloading a game.

Because of these factors, an NES programmer must be careful not to blindly trust the initial contents of RAM.

## Best practices {#Best_practices}

- Configure the emulator so it provides a random system state and random RAM contents on power-up.
  - [Mesen](https://www.mesen.ca/) provides a set of such emulation options recommended for developers, along with a debugger setting to break execution on all reads from uninitialized RAM.
- Refer to the [init code](../Init_code/) article when setting up the reset handler. The sample implementation is a good point to start from.
  - If you are using an [audio driver](../Audio_drivers/), make sure to call its initialization routine in the reset handler before playing any sound.
- If some RAM state is intended to persist across resets, ensure that the checks used to do so are robust against random initial RAM contents. (e.g. unique multi-byte signatures, checksum calculations, etc)
- Validate any data read from potentially unreliable sources before using it. For example, the stats of an RPG character could be checked against valid ranges when loading them from a save.

## See also {#See_also}

- [PPU power up state](../PPU_power_up_state/)

## References {#References}

1.  <span id="cite_note-reset-stack-push-1"></span>[↑](#cite_ref-reset-stack-push_1-0) RESET uses the logic shared with NMI, IRQ, and BRK that would push PC and P. However, like [some but not all 6502s](../Visual6502wiki/6502_BRK_and_B_bit/#masking_of_the_stack_writes_during_RESET), the 2A03 prohibits writes during reset. [This test](https://forums.nesdev.org/viewtopic.php?p=184247#p184247) relies on open bus being precharged by these reads. See [27c3: Reverse Engineering the MOS 6502 CPU (en)](https://www.youtube.com/watch?v=fWqBmmPQP40&t=41m45s) from 41:45 onward for details
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) [Noise channel init log](https://forums.nesdev.org/viewtopic.php?p=172797#p172797)
3.  <span id="cite_note-ebd-3"></span>↑ <sup>[3.0](#cite_ref-ebd_3-0)</sup> <sup>[3.1](#cite_ref-ebd_3-1)</sup> <sup>[3.2](#cite_ref-ebd_3-2)</sup> <sup>[3.3](#cite_ref-ebd_3-3)</sup> [DMC power-up state manifests as buzzing in Eliminator Boat Duel](https://forums.nesdev.org/viewtopic.php?p=231773#p231773)
4.  <span id="cite_note-.244017_lfsr-4"></span>↑ <sup>[4.0](#cite_ref-.244017_lfsr_4-0)</sup> <sup>[4.1](#cite_ref-.244017_lfsr_4-1)</sup> <sup>[4.2](#cite_ref-.244017_lfsr_4-2)</sup> [2A03letterless is missing transistor to set frame counter LFSR on reset](https://forums.nesdev.org/viewtopic.php?p=214939#p214939)
