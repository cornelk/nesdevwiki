---
title: "Emulator tests"
url: "/wiki/Emulator_tests/"
description: "There are many ROMs available that test an emulator for inaccuracies."
wikiPage: "Emulator_tests"
---

There are many ROMs available that test an emulator for [inaccuracies](../Accuracy/).

## Validation ROMs {#Validation_ROMs}

|  |  |
|----|----|
| <img src="../../wiki-images/Ambox_content.png" width="40" height="40" /> | **Some of the download links below are currently dead, but many have been archived at <https://github.com/christopherpow/nes-test-roms>**. |

### CPU Tests {#CPU_Tests}

| Name | Author | Description | Resources |
|----|----|----|----|
| [branch_timing_tests](http://www.slack.net/~ant/nes-tests/branch_timing_tests.zip) | blargg | These ROMs test timing of the branch instruction, including edge cases |  |
| [cpu_dummy_reads](https://github.com/christopherpow/nes-test-roms/raw/master/cpu_dummy_reads/cpu_dummy_reads.nes) | blargg | Tests the CPU's dummy reads | [thread](https://forums.nesdev.org/viewtopic.php?p=31629) |
| [cpu_dummy_writes](http://bisqwit.iki.fi/src/nes_tests/cpu_dummy_writes.zip) | bisqwit | Tests the CPU's dummy writes | [thread](https://forums.nesdev.org/viewtopic.php?t=8738) |
| [cpu_exec_space](http://bisqwit.iki.fi/src/nes_tests/cpu_exec_space.zip) | bisqwit | Verifies that the CPU can execute code from any possible memory location, even if that is mapped as I/O space | [thread](https://forums.nesdev.org/viewtopic.php?t=8755) |
| [cpu_flag_concurrency](http://bisqwit.iki.fi/kala/test1.zip) | bisqwit | Unsure what results are meant to be, see thread for more info. | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=8757) |
| [cpu_interrupts_v2](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/cpu_interrupts_v2/cpu_interrupts.nes) | blargg | Tests the behavior and timing of CPU in the presence of interrupts, both IRQ and NMI; see CPU interrupts. | [thread](https://forums.nesdev.org/viewtopic.php?f=2&t=6510) |
| [cpu_reset](https://github.com/christopherpow/nes-test-roms/tree/master/cpu_reset) | blargg | Tests CPU registers just after power and changes during reset, and that RAM isn't changed during reset. |  |
| [cpu_timing_test6](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/cpu_timing_test6/cpu_timing_test.nes) | blargg | This program tests instruction timing for all official and unofficial NES 6502 instructions except the 8 branch instructions (Bxx) and the 12 halt instructions (HLT) | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=3831) |
| [coredump](https://forums.nesdev.org/viewtopic.php?p=157782#p157782) | jroatch | Coredump tool for displaying contents of RAM | [thread](http://forums.nesdev.org/viewtopic.php?f=22&t=11520) |
| [instr_misc](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/instr_misc/instr_misc.nes) | blargg | Tests some miscellaneous aspects of instructions, including behavior when 16-bit address wraps around, and dummy reads. |  |
| [instr_test_v5](https://github.com/christopherpow/nes-test-roms/tree/master/instr_test-v5) | blargg | Tests official and unofficial CPU instructions and lists which ones failed. It will work even if emulator has no PPU and only supports NROM, writing a copy of output to \$6000 (see readme). This more thoroughly tests instructions, but can't help you figure out what's wrong beyond what instruction(s) are failing, so it's better for testing mature CPU emulators. |  |
| [instr_timing](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/instr_timing/instr_timing.nes) | blargg | Tests timing of all instructions, including unofficial ones, page-crossing, etc. |  |
| [nestest](http://nickmass.com/images/nestest.nes) ([doc](https://www.qmtpro.com/~nes/misc/nestest.txt)) | kevtris | fairly thoroughly tests CPU operation. This is the best test to start with when getting a CPU emulator working for the first time. Start execution at \$C000 and compare execution with a [known good log](https://www.qmtpro.com/~nes/misc/nestest.log) (created using [Nintendulator](../Nintendulator/), an emulator chosen by the test's author because its CPU was verified to function correctly, aside from some minor details of the power-up state). |  |
| [ram_retain](https://forums.nesdev.org/viewtopic.php?t=13334) | rainwarrior | RAM contents test, for displaying contents of RAM at power-on or after reset | [thread](http://forums.nesdev.org/viewtopic.php?t=13334) |

### PPU Tests {#PPU_Tests}

| Name | Author | Description | Resources |
|----|----|----|----|
| [color_test](https://forums.nesdev.org/viewtopic.php?p=155593#p155593) | rainwarrior | Simple display of any chosen color full-screen | [thread](http://forums.nesdev.org/viewtopic.php?p=155593#p155593) |
| [blargg_ppu_tests_2005.09.15b](http://www.slack.net/~ant/nes-tests/blargg_ppu_tests.zip) | blargg | Miscellaneous PPU tests (palette ram, sprite ram, etc.) | [thread](https://forums.nesdev.org/viewtopic.php?t=567) |
| [full_nes_palette](http://www.slack.net/~ant/nes-hacks/full_nes_palette.nes) | blargg | Displays the full palette with all emphasis states, demonstrates direct PPU color control | [thread](https://forums.nesdev.org/viewtopic.php?p=10658) |
| [nmi_sync](http://www.slack.net/~ant/old/nes-code/nmi_sync.zip) | blargg | Verifies NMI timing by creating a specific pattern on the screen (NTSC & PAL versions) | [thread](https://forums.nesdev.org/viewtopic.php?t=6589) |
| [ntsc_torture](http://forums.nesdev.org/viewtopic.php?f=2&t=15080) | rainwarrior | NTSC Torture Test displays visual patterns to demonstrate NTSC signal artifacts | [thread](http://forums.nesdev.org/viewtopic.php?f=2&t=15080) |
| [oam_read](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/oam_read/oam_read.nes) | blargg | Tests OAM reading (\$2004), being sure it reads the byte from OAM at the current address in \$2003. | [thread](https://forums.nesdev.org/viewtopic.php?t=6424) |
| [oam_stress](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/oam_stress/oam_stress.nes) | blargg | Thoroughly tests OAM address (\$2003) and read/write (\$2004) | [thread](https://forums.nesdev.org/viewtopic.php?t=6424) |
| [oamtest3](https://forums.nesdev.org/download/file.php?id=1537) | lidnariq | Utility to upload OAM data via \$2003/\$2004 - can be used to test for the OAMADDR bug behavior | [thread 1](http://forums.nesdev.org/viewtopic.php?p=128913#p128913) [thread 2](http://forums.nesdev.org/viewtopic.php?p=128842#p128842) |
| [palette](https://forums.nesdev.org/viewtopic.php?t=13264) | rainwarrior | Palette display requiring only scanline-based palette changes, intended to demonstrate the full palette even on less advanced emulators | [thread](http://forums.nesdev.org/viewtopic.php?t=13264) |
| [ppu_open_bus](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/ppu_open_bus/ppu_open_bus.nes) | blargg | Tests behavior when reading from open-bus PPU bits/registers |  |
| [ppu_read_buffer](http://bisqwit.iki.fi/src/nes_tests/ppu_read_buffer.zip) | bisqwit | Mammoth test pack tests many aspects of the NES system, mostly centering around the PPU \$2007 read buffer | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=8847&start=0) |
| [ppu_sprite_hit](http://blargg.8bitalley.com/parodius/nes-tests/ppu_sprite_hit.zip) | blargg | Tests sprite 0 hit behavior and timing | [thread](https://forums.nesdev.org/viewtopic.php?t=626) |
| [sprite_overflow_tests](http://www.slack.net/~ant/nes-tests/sprite_overflow_tests.zip) | blargg | Tests sprite overflow behavior and timing | [thread](https://forums.nesdev.org/viewtopic.php?t=1308) |
| [ppu_vbl_nmi](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/ppu_vbl_nmi/ppu_vbl_nmi.nes) | blargg | Tests the behavior and timing of the NTSC PPU's VBL flag, NMI enable, and NMI interrupt. Timing is tested to an accuracy of one PPU clock. | [thread](https://forums.nesdev.org/viewtopic.php?t=730) |
| [scanline](http://www.qmtpro.com/~nes/demos/scanline.zip) | Quietust | Displays a test screen that will contain glitches if certain portions of the emulation are not perfect. |  |
| [sprdma_and_dmc_dma](https://github.com/christopherpow/nes-test-roms/tree/master/sprdma_and_dmc_dma) | blargg | Tests the cycle stealing behavior of the DMC DMA while running Sprite DMAs | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=6100) |
| [sprite_hit_tests_2005.10.05](http://www.slack.net/~ant/nes-tests/sprite_hit_tests.zip) | blargg | Generally the same as ppu_sprite_hit (older revision of the tests - ppu_sprite_hit is most likely better) | [thread](https://forums.nesdev.org/viewtopic.php?t=626) |
| [sprite_overflow_tests](http://www.slack.net/~ant/nes-tests/sprite_overflow_tests.zip) | blargg | Generally the same as ppu_sprite_overflow (older revision of the tests - ppu_sprite_overflow is most likely better) | [thread](https://forums.nesdev.org/viewtopic.php?t=1308) |
| [tvpassfail](http://pics.pineight.com/nes/tvpassfail.zip) | tepples | NTSC color and NTSC/PAL pixel aspect ratio test ROM (older revision of the tests - 240p Test Suite is most likely better) | [thread](https://forums.nesdev.org/viewtopic.php?t=3393) |
| [vbl_nmi_timing](https://github.com/christopherpow/nes-test-roms/tree/master/vbl_nmi_timing) | blargg | Generally the same as ppu_vbl_nmi (older revision of the tests - ppu_vbl_nmi is most likely better) | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=3953) |

### APU Tests {#APU_Tests}

| Name | Author | Description | Resources |
|----|----|----|----|
| [apu_mixer](https://github.com/christopherpow/nes-test-roms/tree/master/apu_mixer) | blargg | Verifies proper operation of the APU's sound channel mixer, including relative volumes of channels and non-linear mixing. recordings when run on NES are available for comparison, though the tests are made so that you don't really need these. | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=4911) |
| [apu_phase_reset](https://forums.nesdev.org/download/file.php?id=7496) | Rahsennor | Tests the correct square channel behavior when writing to \$4003/4007 (reset the duty cycle sequencers but not the clock dividers) | [thread](https://forums.nesdev.org/viewtopic.php?f=2&t=15346) |
| [apu_reset](https://github.com/christopherpow/nes-test-roms/tree/master/apu_reset) | blargg | Tests initial APU state at power, and the effect of reset. |  |
| [apu_test](https://github.com/christopherpow/nes-test-roms/raw/refs/heads/master/apu_test/apu_test.nes) | blargg | Tests many aspects of the APU that are visible to the CPU. Really obscure things are not tested here. |  |
| [blargg_apu_2005.07.30](http://www.slack.net/~ant/nes-tests/blargg_apu_2005.07.30.zip) | blargg | Tests APU length counters, frame counter, IRQ, etc. |  |
| [dmc_dma_during_read4](https://github.com/christopherpow/nes-test-roms/tree/master/dmc_dma_during_read4) | blargg | Tests register read/write behavior while DMA is running |  |
| [dmc_tests](https://github.com/christopherpow/nes-test-roms/tree/master/dmc_tests) | ? | ? |  |
| [dpcmletterbox](http://pics.pineight.com/nes/dpcmletterbox.zip) | tepples | Tests accuracy of the DMC channel's IRQ |  |
| [pal_apu_tests](http://www.slack.net/~ant/nes-tests/pal_apu_tests.zip) | blargg | PAL version of the blargg_apu_2005.07.30 tests |  |
| [square_timer_div2](http://blargg.8bitalley.com/parodius/nes-tests/older/square_timer_div2.zip) | blargg | Tests the square timer's period |  |
| [test_apu_2 (1-10)](https://forums.nesdev.org/download/file.php?id=1494) [(11)](https://forums.nesdev.org/download/file.php?id=5699) | x0000 | 11 tests that verify a number of behaviors with the APU (including the frame counter) | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=11174) |
| [test_apu_env](http://blargg.8bitalley.com/parodius/nes-tests/older/test_apu_env.zip) | blargg | Tests the APU envelope for correctness. |  |
| [test_apu_sweep](http://blargg.8bitalley.com/parodius/nes-tests/older/test_apu_sweep.zip) | blargg | Tests the sweep unit's add, subtract, overflow cutoff, and minimum period behaviors. |  |
| [test_apu_timers](http://blargg.8bitalley.com/parodius/nes-tests/older/test_apu_timers.zip) | blargg | Tests frequency timer of all 5 channels |  |
| [test_tri_lin_ctr](http://blargg.8bitalley.com/parodius/nes-tests/older/test_tri_lin_ctr.zip) | blargg | Tests triangle's linear counter and clocking by the frame counter |  |
| [volume_tests](http://pics.pineight.com/nes/volume_tests.zip) | tepples | Plays tones on all the APU's channels to show their relative volumes at various settings of \$4011. Package includes a recording from an NES's audio output for comparison. |  |

### Mapper-specific Tests {#Mapper-specific_Tests}

<table class="wikitable sortable">
<thead>
<tr>
<th>Name</th>
<th>Author</th>
<th>Description</th>
<th>Resources</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=13120">31_test</a></td>
<td>rainwarrior</td>
<td>Tests for mapper 31 (see thread for ROMs in various PRG sizes)</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=13120">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=2253">BNTest</a></td>
<td>tepples</td>
<td>Tests how many PRG banks are reachable in BxROM and AxROM.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=79826#p79826">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/bntest">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=12085">bxrom_512k_test</a></td>
<td>rainwarrior</td>
<td>Similar to the BxROM test in BNTest above.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=12085">thread</a></td>
</tr>
<tr>
<td><a href="http://forums.nesdev.org/download/file.php?id=10240">FdsIrqTests (v7)</a></td>
<td>Sour</td>
<td>Tests various elements of the FDS' IRQ</td>
<td><a href="http://forums.nesdev.org/viewtopic.php?f=3&amp;t=16507">thread</a></td>
</tr>
<tr>
<td><a href="https://github.com/TakuikaNinja/FDS-Mirroring-Test/releases/latest/download/FDS-Mirroring-Test.zip">FDS-Mirroring-Test</a></td>
<td>TakuikaNinja</td>
<td>Tests the FDS' nametable arrangement/mirroring switching functionality ($4025.D3 write, $4030.D3 read, $4023.D0=0 write)</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=300488#p300488">thread</a> <a href="https://github.com/TakuikaNinja/FDS-Mirroring-Test">GitHub</a></td>
</tr>
<tr>
<td><a href="https://github.com/TakuikaNinja/FDS-4023-Test/releases/latest/download/FDS-4023-Test.zip">FDS-4023-Test</a></td>
<td>TakuikaNinja</td>
<td>Displays the state of FDS registers when toggling $4023 bits</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=26313">thread</a> <a href="https://github.com/TakuikaNinja/FDS-4023-Test">GitHub</a></td>
</tr>
<tr>
<td><a href="https://github.com/TakuikaNinja/FDS-Audio-Registers/releases/latest/download/FDS-Audio-Registers.zip">FDS-Audio-Registers</a></td>
<td>TakuikaNinja</td>
<td>Displays the state of FDS audio registers when toggling $4023.D1</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=308383#p308383">thread</a> <a href="https://github.com/TakuikaNinja/FDS-Audio-Registers">GitHub</a></td>
</tr>
<tr>
<td><a href="https://github.com/TakuikaNinja/FDS-4030D1-Addr/releases/latest/download/FDS-4030D1-Addr.zip">FDS-4030D1-Addr</a></td>
<td>TakuikaNinja</td>
<td>Visualises IRQs generated by the FDS' DRAM refresh watchdog, reported in $4030.D1</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=26359">thread</a> <a href="https://github.com/TakuikaNinja/FDS-4030D1-Addr">GitHub</a></td>
</tr>
<tr>
<td><a href="http://www.qmtpro.com/~nes/demos/mmc5exram.zip">exram</a></td>
<td>Quietust</td>
<td>MMC5 exram test</td>
<td></td>
</tr>
<tr>
<td><a href="http://forums.nesdev.org/viewtopic.php?t=8639">famicom_audio_swap_tests</a></td>
<td>rainwarrior</td>
<td>Hotswap tests for Famicom expansion audio (5B, MMC5, VRC6, VRC7, N163, FDS)</td>
<td><a href="http://forums.nesdev.org/viewtopic.php?t=8639">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=2759">fme7acktest-r1</a></td>
<td>tepples</td>
<td>Checks some IRQ acknowledgment behiaviors of Sunsoft FME-7 that emulators were getting wrong in 2015.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=12436">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=2800">fme7ramtest-r1</a></td>
<td>tepples</td>
<td>Checks how much work RAM the Sunsoft FME-7 can access</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=9&amp;t=12467">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/fme7acktest">GitHub</a></td>
</tr>
<tr>
<td><a href="https://pineight.com/nes/holydiverbatman-bin-0.01.7z">Holy Mapperel</a></td>
<td>tepples</td>
<td>Detects over a dozen mappers and verifies that all PRG ROM and CHR ROM banks are reachable, that PRG RAM and CHR RAM can be written and read back without error, and that nametable mirroring, IRQ, and WRAM protection work. (Formerly Holy Diver Batman)</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=22&amp;t=10640">thread</a> <a href="https://github.com/pinobatch/holy-mapperel">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=5166">mmc3bigchrram</a></td>
<td>tepples</td>
<td>MMC3 test for large 32kb CHR RAM with NES 2.0 headers</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=13890">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/mmc3bigchrram">GitHub</a></td>
</tr>
<tr>
<td><a href="http://slack.net/~ant/old/nes-tests/mmc3_test_2.zip">mmc3_test_2</a></td>
<td>blargg</td>
<td>MMC3 scanline counter and IRQ generation tests.</td>
<td></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=261236#p261236">mmc3irqtest</a></td>
<td>N-K</td>
<td>MMC3 scanline IRQ test and $C000 glitch investigation.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=261236#p261236">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=76817#p76817">mmc5test</a></td>
<td>Drag</td>
<td>MMC5 scanline counter</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=7653">thread</a></td>
</tr>
<tr>
<td><a href="http://forums.nesdev.org/download/file.php?id=8609">mmc5test_v2</a></td>
<td>AWJ</td>
<td>MMC5 tests</td>
<td><a href="http://forums.nesdev.org/viewtopic.php?f=2&amp;t=15788">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=3753">serom</a></td>
<td>lidnariq</td>
<td>Tests the constraints of SEROM / SHROM / SH1ROM variations of the MMC1 boards.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p153298">thread</a></td>
</tr>
<tr>
<td rowspan="4">NES 2.0 Submapper Tests</td>
<td>rainwarrior</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p157804">2_test</a> - Mapper 2, Submappers 0, 1 and 2</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p157804">thread</a></td>
</tr>
<tr>
<td>rainwarrior</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p154555">3_test</a> - Mapper 3, Submappers 0, 1 and 2</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p154555">thread</a></td>
</tr>
<tr>
<td>rainwarrior</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p157804">7_test</a> - Mapper 7, Submappers 0, 1 and 2</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p157804">thread</a></td>
</tr>
<tr>
<td>rainwarrior</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p153334">34_test</a> - Mapper 34, Submappers 1 and 2</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=9350&amp;start=90#p153334">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=12185">test28</a></td>
<td>tepples</td>
<td>Tests for mapper 28</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=215345#p215345">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/test28">GitHub</a></td>
</tr>
<tr>
<td><a href="http://forums.nesdev.org/download/file.php?id=10017">vrc24test</a></td>
<td>AWJ</td>
<td>Detects and tests all VRC 2/4 variants</td>
<td><a href="http://forums.nesdev.org/viewtopic.php?f=3&amp;t=16009">thread</a></td>
</tr>
<tr>
<td><a href="http://www.mediafire.com/download/6hvuj53omv8y3fn/vrc6test.zip">vrc6test</a></td>
<td>natt</td>
<td>VRC6 mirroring tests</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=11028">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=244062#p244062">mmc5ramsize</a></td>
<td>rainwarrior</td>
<td>MMC5 large PRG-RAM support tests</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=244062#p244062">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=23619">mmc1atest</a></td>
<td>tepples</td>
<td>Characterizes behavior of <a href="../MMC1/">MMC1A</a> vs. <a href="../MMC1/">MMC1B</a></td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=23619">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/mmc1a">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=284414#p284414">n163_soundram</a></td>
<td>rainwarrior</td>
<td>Test for <a href="../Namco_163_audio/">Namco 163 audio</a> sound RAM read-back.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=284414#p284414">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=285135#p285135">n163_soundram_init</a></td>
<td>rainwarrior</td>
<td>Test for <a href="../Namco_163_audio/">Namco 163 audio</a> sound RAM contents at power-on.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=285135#p285135">thread</a></td>
</tr>
</tbody>
</table>

### Input Tests {#Input_Tests}

<table class="wikitable sortable">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Name</th>
<th>Author</th>
<th>Description</th>
<th>Resources</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="https://github.com/pinobatch/allpads-nes/releases/latest/download/allpads.nes">allpads</a></td>
<td>tepples</td>
<td>Multiple controller test supporting NES controller, Super NES controller, Famicom microphone, Four Score, Zapper, NES Arkanoid controller, and Super NES Mouse; also has raw 32-bit report mode</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=12549">thread</a> <a href="https://github.com/pinobatch/allpads-nes">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=5905">dma_sync_test_v2</a></td>
<td>Rahsennor</td>
<td>Tests DMC DMA read corruption</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=14319">thread</a></td>
</tr>
<tr>
<td><a href="https://github.com/christopherpow/nes-test-roms/raw/master/PaddleTest3/PaddleTest.nes">PaddleTest3</a></td>
<td>3gengames</td>
<td>Test for the Arkanoid controller</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=7929">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=21972">vaus</a></td>
<td>lidnariq</td>
<td>Arkanoid controller 9-bit result test</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=23801">thread</a></td>
</tr>
<tr>
<td><a href="https://github.com/christopherpow/nes-test-roms/tree/master/read_joy3">read_joy3</a></td>
<td>blargg</td>
<td>Various NES controllers tests, including read corruption due to DMC DMA</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=4124&amp;start=0">thread</a></td>
</tr>
<tr>
<td><a href="https://github.com/pinobatch/zap-ruder/releases/latest/download/ruder.nes">Zap Ruder</a></td>
<td>tepples</td>
<td>Test for the Zapper, including dual wield but not the serial Vs. variant</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=8108">thread</a> <a href="https://github.com/pinobatch/zap-ruder">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/download/file.php?id=5356">spadtest-nes</a></td>
<td>tepples</td>
<td>Super Nintendo controller test (when connected to a NES)</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=167288">thread</a></td>
</tr>
<tr>
<td><a href="http://pics.pineight.com/nes/vaus-test-0.02.zip">vaus_test</a></td>
<td>tepples</td>
<td>Another test for the Arkanoid controller</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=120455">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/vaus-test">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=231608#p231608">mset</a></td>
<td>rainwarrior</td>
<td>SNES mouse test</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=231608#p231608">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=18314&amp;p=232358#p232358">mict</a></td>
<td>rainwarrior</td>
<td>Famicom microphone test</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=3&amp;t=18314">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=22&amp;t=18998">Telling LYs?</a></td>
<td>tepples</td>
<td>Tests whether input can change on any scanline</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=22&amp;t=18998">thread</a> <a href="https://github.com/pinobatch/little-things-nes/tree/master/tellinglys">GitHub</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=19752">ctrltest</a></td>
<td>rainwarrior</td>
<td>Generic log of 16-bit report on all 5 input data lines.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=2&amp;t=19752">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=106954#p106954">raw</a>
<p><a href="https://forums.nesdev.org/viewtopic.php?t=23745">pack2</a></p></td>
<td>lidnariq</td>
<td>Immediate state of 32-bit report on all 5x2 input data lines.
<p>Immediate state of 64-bit report on all 5x2 input data lines.</p></td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=106954#p106954">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=9&amp;t=13021">zapper tests</a></td>
<td>rainwarrior</td>
<td>Simple tests for displaying output of zapper reads.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?f=9&amp;t=13021">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=24192">powerpadgesture</a></td>
<td>tepples</td>
<td>Gesture test for <a href="../Power_Pad/">Power Pad</a> displaying log of presses and releases.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?t=24192">thread</a></td>
</tr>
<tr>
<td><a href="https://www.nesdev.org/powerpd.zip">POWERPAD.NES</a></td>
<td>Tennessee Carmel-Veilleux</td>
<td>Old simple test for Power Pad.</td>
<td><a href="https://www.nesdev.org/powerpad.txt">powerpad.txt</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=294558#p294558">d34test</a></td>
<td>rainwarrior</td>
<td>Generic log of 32-bit report on D3 and D4 data lines, with 16-bit report on D0.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=294558#p294558">thread</a></td>
</tr>
<tr>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=294558#p294558">kmbtest</a></td>
<td>rainwarrior</td>
<td>Input test for a proposed USB Keyboard and Mouse interface.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=294558#p294558">thread</a></td>
</tr>
<tr>
<td><a href="https://github.com/TakuikaNinja/HoriTracker/releases/latest/download/HoriTracker.zip">Hori Tracker</a></td>
<td>TakuikaNinja</td>
<td>Input display for the <a href="../Hori_Track/">Hori Track</a>.</td>
<td><a href="https://forums.nesdev.org/viewtopic.php?p=310909#p310909">forum post</a> <a href="https://github.com/TakuikaNinja/HoriTracker">GitHub</a></td>
</tr>
</tbody>
</table>

### Misc Tests {#Misc_Tests}

| Name | Author | Description | Resources |
|----|----|----|----|
| [240pee-0.22](https://forums.nesdev.org/download/file.php?id=19130) | tepples | 240p Test Suite (an NES version of the [240p Test Suite](http://junkerhq.net/xrgb/index.php/240p_test_suite) by Artemio Urbina), including an MDFourier tone generator | [thread](https://forums.nesdev.org/viewtopic.php?t=13394) [GitHub](https://github.com/pinobatch/240p-test-mini/tree/master/nes) |
| [characterize-vs](https://forums.nesdev.org/download/file.php?id=1415) | lidnariq | VS System tests | [thread](https://forums.nesdev.org/viewtopic.php?f=3&t=10276&start=15#p127592) |
| [NEStress](http://nesdev.org/NEStress.zip) | Flubba | Old test - some of the tests are supposed to fail on real hardware. |  |
| [oc-r1a](https://forums.nesdev.org/download/file.php?id=2826) | tepples | Detects and displays accurate clock rate of the NES (since incorporated into 240p Test Suite) | [thread](https://forums.nesdev.org/viewtopic.php?f=22&t=12499) |
| [nes-audio-tests](https://github.com/bbbradsmith/nes-audio-tests) | rainwarrior | NSF and NES ROM tests for expansion audio sound, NSF behaviour, and other various audio related things. |  |

## Automated testing {#Automated_testing}

It's best if your emulator can automatically run a suite of tests at the press of a button. This allows you to re-run them every time you make a change, without any effort. Automation can be difficult, because the emulator must be able to determine success/failure without your help.

The first part of automated testing is support for a "movie" or "demo", or a list of what buttons were pressed when. An emulator makes a movie by recording presses while the user is playing, and then it plays the movie by feeding the recorded presses back through the input system. This not only helps automated testing but also makes your emulator attractive to [speedrunners](https://en.wikipedia.org/wiki/Speedrun).

To create a test case, record a movie of the player activating all tests in a ROM, take a screenshot of each result screen, and log the time and a [hash](https://en.wikipedia.org/wiki/Hash_function) of each screenshot. The simplest test ROMs won't require any button presses. ROMs that test more than one thing are more likely to require them, and an [actual game](../Tricky-to-emulate_games/) will require a playthrough. Then to run a test case, play the movie in fast-forward (no delay between frames) and take screenshots at the same times. If a screenshot's hash differs from that of the corresponding screenshot from when the test case was created, make a note of this difference in the log. Then you can compare the emulator's output frame-by-frame to that of the previous release of your emulator running the same test case.

## See also {#See_also}

- [Emulators](../Emulators/)
