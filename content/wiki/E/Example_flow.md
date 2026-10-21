---
title: "Example flow"
url: "/wiki/Example_flow/"
wikiPage: "Example_flow"
---

## <span id="LDA_.23.2420_.2F_STA_.242006_Low-Level_Hardware_Flow"></span>LDA \#\$20 / STA \$2006 Low-Level Hardware Flow {#LDA_.23.2420_.2F_STA_.242006_Low-Level_Hardware_Flow}

### <span id="1._Instruction:_LDA_.23.2420"></span>1. Instruction: LDA \#\$20 {#1._Instruction:_LDA_.23.2420}

- Opcode:\* A9
- Operand:\* 20

1.  \*\*CPU Fetch Opcode\*\*
2.  CPU places \*\*Program Counter (PC)\*\* on address bus (A0–A15).
3.  NES address decoder asserts \*\*CE (chip enable)\*\* for \*\*PRG ROM\*\*.
4.  PRG ROM drives \*\*data bus D0–D7\*\* with opcode \`A9\`.
5.  CPU reads opcode into instruction register.

<!-- -->

1.  \*\*CPU Fetch Operand\*\*
2.  CPU increments PC → address bus points to next byte.
3.  PRG ROM CE active, data bus outputs operand \`20\`.
4.  CPU reads operand → Accumulator (A) will later be loaded with \`20\`.

<!-- -->

1.  \*\*Execute LDA\*\*
2.  CPU loads operand into \*\*Accumulator (A)\*\*.
3.  No external bus activity required for execution.

### <span id="2._Instruction:_STA_.242006"></span>2. Instruction: STA \$2006 {#2._Instruction:_STA_.242006}

- Target:\* PPUADDR (PPU register \$2006)

1.  \*\*CPU Places Address on Bus\*\*
2.  CPU puts \`\$2006\` on address bus (A0–A15).
3.  NES address decoder activates \*\*PPU CE\*\*, deactivates other devices.

<!-- -->

1.  \*\*CPU Drives Data Bus\*\*
2.  CPU places \*\*Accumulator value \`20\`\*\* on data bus D0–D7.
3.  CPU asserts \*\*WR (write)\*\* signal.

<!-- -->

1.  \*\*PPU Latches Data\*\*
2.  PPU detects \*\*CE active + WR active\*\*.
3.  \`\$2006\` first write → high byte of VRAM address latched.
4.  \`\$2006\` second write → low byte latched.
5.  PPU internal flip-flop tracks first/second write.

### <span id="3._Bus_.26_Signal_Summary"></span>3. Bus & Signal Summary {#3._Bus_.26_Signal_Summary}

| Step | CPU Address Bus | CPU Data Bus | CE | RD/WR | Notes |
|----|----|----|----|----|----|
| Fetch Opcode (A9) | PC | D0–D7 | PRG ROM CE active | RD | PRG ROM outputs opcode |
| Fetch Operand (20) | PC+1 | D0–D7 | PRG ROM CE active | RD | PRG ROM outputs operand |
| STA \$2006 | \$2006 | D0–D7 = Accumulator | PPU CE active | WR | PPU latches VRAM address |

### 4. Notes {#4._Notes}

- Only one device drives the data bus at a time to prevent conflicts.
- PPU registers \$2000–\$2007 are memory-mapped; accessed via address decoder.
- Internal PPU latches track high/low byte for \$2006 VRAM address writes.
