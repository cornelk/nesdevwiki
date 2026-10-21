---
title: "Accuracy"
url: "/wiki/Accuracy/"
description: "Compatibility is how well an emulator produces the same output as the original system when running a particular known program. Accuracy refers to how well an emulator produces the same output as the…"
wikiPage: "Accuracy"
---

**Compatibility** is how well an emulator produces the same output as the original system when running a particular known program. **Accuracy** refers to how well an emulator produces the same output as the original system when running an arbitrary unknown program. An emulator can be highly compatible without being very accurate; NESticle was this way back in the 1990s, running a wide selection of popular North American commercial games.

Accuracy cannot be measured directly.<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> Compatibility with [ROMs designed to test obscure behaviors](../Emulator_tests/) and [games relying on obscure behaviors](../Tricky-to-emulate_games/) is a good (though not perfect) predictor of accuracy.

In many cases, an emulator is more forgiving than the hardware, especially in when a program is allowed to write to a register. NESticle, for instance, allowed writes to [\$2007](../PPU_registers/) at any time, while the NES only allows it during scanlines 240-260 of each frame (vertical blanking) or when \$2001 & \#\$18 == 0 (forced blanking). Most emulators allow the program to twiddle PPU registers immediately, while the NES PPU [ignores most writes for the first frame after a reset](../PPU_power_up_state/). Inaccuracies like these led to inadvertent development of [programs that the NES itself cannot run](../Program_compatibility/) during the early years of NES homebrew.

Whenever you discover a difference in behavior between the NES and best-of-breed [emulators](../Emulators/) (like Mesen, Nintendulator, and Nestopia), you can help emulators become more accurate. First reduce your program to a [minimal test case](http://sscce.org/): keep removing things while the program continues to exhibit this difference. Then characterize the behavior difference as best you can, add it to the [test suite](../Emulator_tests/), and notify emulator developers through their projects' issue trackers.

Some emulators attempt to [enhance](../Enhancement/) the games to make them appear more appealing than they did on the original NES.

## See also {#See_also}

- [Emulator tests](../Emulator_tests/)
- [Tricky-to-emulate games](../Tricky-to-emulate_games/)
- [Game bugs](../Game_bugs/) - Accurate emulation should correctly reproduce known software bugs.
- [Program compatibility](../Program_compatibility/) - Homebrew games that may suffer from compatibility problems.

## References {#References}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) Measuring an arbitrary emulator's accuracy reduces to the [halting problem](https://en.wikipedia.org/wiki/Halting_problem) over [linear bounded automata](https://en.wikipedia.org/wiki/Linear_bounded_automaton), which is intractable. It'd be easier to [formally prove](https://en.wikipedia.org/wiki/Formal_verification) equivalence to the netlist of [Visual 2A03 and 2C02](../Visual_circuit_tutorial/).

- [Topic on BBS](https://forums.nesdev.org/viewtopic.php?t=5797)
- [Different grades of accuracy on BBS](https://forums.nesdev.org/viewtopic.php?p=30784#p30784)
- [How to make a minimal test case](https://wiki.mozilla.org/QA/Minimal_Test_Cases) (focuses on HTML, but applies to any virtual machine)
