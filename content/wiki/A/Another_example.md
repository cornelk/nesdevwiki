---
title: "Another example"
url: "/wiki/Another_example/"
description: "Purpose: Sets the PPU VRAM address via $2006 register."
wikiPage: "Another_example"
---

## <span id="LDA_.23.2420_.2F_STA_.242006_Low-Level_Hardware_Flow"></span>LDA \#\$20 / STA \$2006 Low-Level Hardware Flow {#LDA_.23.2420_.2F_STA_.242006_Low-Level_Hardware_Flow}

### 1. Instruction Overview {#1._Instruction_Overview}

Instruction sequence:

    LDA #$20   ; Load immediate value 0x20 into accumulator
    STA $2006  ; Store accumulator to PPUADDR register

Purpose: Sets the PPU VRAM address via \$2006 register.

Key hardware involved:

- CPU (Ricoh 2A03)
- PRG ROM
- NES Address Decoder
- PPU (Picture Processing Unit)
- System buses: Address (A0–A15), Data (D0–D7)
- Chip Enable (CE) and Read/Write (RD/WR) control signals

---

### <span id="2._LDA_.23.2420_Execution_Flow"></span>2. LDA \#\$20 Execution Flow {#2._LDA_.23.2420_Execution_Flow}

1.  CPU Fetch Opcode

- CPU places \*\*PC\*\* on address bus (A0–A15).
- NES address decoder asserts \*\*PRG ROM CE\*\*.
- PRG ROM outputs \*\*opcode \`A9\`\*\* on data bus (D0–D7).
- CPU reads opcode into instruction register.

1.  CPU Fetch Operand

- CPU increments PC → next byte address.
- PRG ROM CE active, outputs \*\*operand \`20\`\*\*.
- CPU reads operand into instruction register.

1.  Execute LDA

- CPU loads operand \`20\` into \*\*Accumulator (A)\*\*.
- No external bus activity is required during execution.

---

### <span id="3._STA_.242006_Execution_Flow"></span>3. STA \$2006 Execution Flow {#3._STA_.242006_Execution_Flow}

1.  CPU Places Address on Bus

- CPU puts \`\$2006\` on A0–A15.
- Address decoder activates \*\*PPU CE\*\*, disables other devices.

1.  CPU Drives Data Bus

- CPU puts \*\*Accumulator value (0x20)\*\* on D0–D7.
- CPU asserts \*\*WR (write)\*\* signal.

1.  PPU Latches Data

- PPU detects \*\*CE active + WR active\*\*.
- Internal \$2006 latch:

<!-- -->

     * First write → high byte of VRAM address  
     * Second write → low byte of VRAM address  

- Flip-flop tracks first/second write for scrolling & VRAM updates.

---

### 4. Cycle-Level Timing {#4._Cycle-Level_Timing}

| Instruction | CPU Cycle | Action |
|----|----|----|
| LDA \#\$20 | 1 | Fetch opcode from PRG ROM |
| LDA \#\$20 | 2 | Fetch immediate operand from PRG ROM |
| LDA \#\$20 | 3 | Load accumulator |
| STA \$2006 | 1 | Place \$2006 on address bus, assert PPU CE |
| STA \$2006 | 2 | Place data (A=20) on data bus, assert WR, PPU latches high byte |
| STA \$2006 | 3 | Next write latches low byte of VRAM address |

---

### <span id="5._NES_Bus_.26_Signal_Summary"></span>5. NES Bus & Signal Summary {#5._NES_Bus_.26_Signal_Summary}

| Step | CPU Address Bus | CPU Data Bus | CE | RD/WR | Notes |
|----|----|----|----|----|----|
| Fetch Opcode (A9) | PC | D0–D7 | PRG ROM CE active | RD | PRG ROM outputs opcode |
| Fetch Operand (20) | PC+1 | D0–D7 | PRG ROM CE active | RD | PRG ROM outputs operand |
| STA \$2006 | \$2006 | D0–D7 = Accumulator | PPU CE active | WR | PPU latches VRAM address |

---

### 6. Address Decoder Mapping {#6._Address_Decoder_Mapping}

CPU A0–A15 → NES Address Decoder → CE signals:

    $0000–$1FFF → Internal RAM
    $2000–$2007 → PPU registers
    $4000–$4017 → APU / I-O
    $8000–$FFFF → PRG ROM

- Only one device drives the data bus at a time to prevent conflicts.

---

### <span id="7._PPU_.242006_Latch_Behavior"></span>7. PPU \$2006 Latch Behavior {#7._PPU_.242006_Latch_Behavior}

- First write → high byte of VRAM address
- Second write → low byte
- Internal flip-flop tracks write sequence
- Important for scrolling, raster effects, and VRAM access timing

---

### 8. Diagram Overview {#8._Diagram_Overview}

    PRG ROM
      | (D0–D7)   <-- CE active
      v
    CPU (Accumulator, PC)
      | (A0–A15)
      | (D0–D7)
      v
    NES Address Decoder
      |----> PPU CE active -> PPUADDR latch (high/low byte)
      |----> Other CE inactive (tri-state)

---

### 9. Notes {#9._Notes}

- CPU cycles, CE activation, and bus usage are critical for timing-sensitive NES operations.
- Mirroring and PPU latch behavior affect scrolling and background updates.
- This low-level understanding is useful for NES hardware emulation, demo scene effects, and precision programming.
