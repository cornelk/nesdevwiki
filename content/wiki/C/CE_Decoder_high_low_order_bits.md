---
title: "CE Decoder high/low order bits"
url: "/wiki/CE_Decoder_high/low_order_bits/"
description: "The NES motherboard has multiple devices sharing the same CPU bus:"
wikiPage: "CE_Decoder_high_low_order_bits"
---

## <span id="NES_CPU_Memory_.26_CE_Decoder_Cheat_Sheet"></span>NES CPU Memory & CE Decoder Cheat Sheet {#NES_CPU_Memory_.26_CE_Decoder_Cheat_Sheet}

### 1. Purpose of the Address Decoder for CE {#1._Purpose_of_the_Address_Decoder_for_CE}

The NES motherboard has multiple devices sharing the same CPU bus:

- PRG ROM
- CPU RAM
- PPU registers
- APU / I/O

To prevent bus conflicts, only one device should drive the data bus at any time. The \*\*address decoder\*\* handles this by:

- Observing the CPU address pins (A0–A15)
- Determining which device should respond
- Activating its \*\*Chip Enable (CE)\*\* while tri-stating all other devices (high-impedance on their data pins)

### 2. How CE Logic Works for the PPU {#2._How_CE_Logic_Works_for_the_PPU}

For example, \*\*PPU registers\*\* are mapped to \$2000–\$3FFF:

- The decoder checks the high address pins (A15–A13).
- If the address falls within \$2000–\$3FFF, CE for the PPU is activated.
- The lower address pins (A0–A2) are passed to the PPU to select a specific register (\$2000 → \$2007).
- All other devices see CE = inactive → their outputs are tri-stated and do not drive the data bus.

🔹 Tri-state behavior

- Tri-state = effectively disconnected from the bus.
- Only the selected device drives D0–D7 at any given cycle.
- This prevents electrical conflicts and ensures the CPU reads/writes the correct device.

### 3. CPU Memory Ranges {#3._CPU_Memory_Ranges}

| CPU Address Range | Size  | Device        | Notes / Mirrors                    |
|-------------------|-------|---------------|------------------------------------|
| \$0000–\$07FF     | 2 KB  | CPU RAM       | Base RAM                           |
| \$0800–\$0FFF     | 2 KB  | RAM mirror    | Mirror 1                           |
| \$1000–\$17FF     | 2 KB  | RAM mirror    | Mirror 2                           |
| \$1800–\$1FFF     | 2 KB  | RAM mirror    | Mirror 3                           |
| \$2000–\$2007     | 8 B   | PPU registers | Internal A0–A2 select PPU register |
| \$2008–\$3FFF     | 8 KB  | PPU mirror    | Mirrors every 8 bytes              |
| \$4000–\$401F     | 32 B  | APU / I/O     | Controllers, sound registers       |
| \$4020–\$5FFF     | 8 KB  | Expansion ROM | Optional cartridge area            |
| \$6000–\$7FFF     | 8 KB  | SRAM          | Battery-backed RAM                 |
| \$8000–\$FFFF     | 32 KB | PRG ROM       | Game code, banks vary by cartridge |

### 4. High-Order Bits Observed by CE Decoder {#4._High-Order_Bits_Observed_by_CE_Decoder}

| Component | CPU Address Range | High Bits A15–A13 (or used) | CE Line (74LS138 output) |
|----|----|----|----|
| RAM | \$0000–\$1FFF | 000 / 001 (mirrors) | Y0 |
| PPU | \$2000–\$3FFF | 001 | Y1 |
| Expansion / I/O | \$4000–\$5FFF | 010 | Y2 |
| SRAM | \$6000–\$7FFF | 011 | Y3 |
| PRG ROM | \$8000–\$FFFF | 100 / 101 / 110 / 111 | Y4–Y7 (depending on bank) |

Mirrors are handled by ignoring lower bits in the decoder. Each component has unique high-order bits → prevents bus conflicts.

### 5. 74LS138 CE Logic {#5._74LS138_CE_Logic}

- Inputs:
  - A2 (pin 15) ← CPU highest observed bit (e.g., A15)
  - A1 (pin 14) ← CPU next high bit (e.g., A14)
  - A0 (pin 13) ← CPU lowest high bit (e.g., A13)
  - Enable: G1=HIGH, G2A & G2B=LOW

<!-- -->

- Outputs: Y0–Y7 → active LOW, connected to CE of devices

| Input A2 A1 A0 | Binary | Output LOW | NES Device     |
|----------------|--------|------------|----------------|
| 0 0 0          | 000    | Y0         | CPU RAM        |
| 0 0 1          | 001    | Y1         | PPU            |
| 0 1 0          | 010    | Y2         | Expansion/I/O  |
| 0 1 1          | 011    | Y3         | SRAM           |
| 1 0 0          | 100    | Y4         | PRG ROM bank 0 |
| 1 0 1          | 101    | Y5         | PRG ROM bank 1 |
| 1 1 0          | 110    | Y6         | PRG ROM bank 2 |
| 1 1 1          | 111    | Y7         | PRG ROM bank 3 |

Only one output is active LOW at a time; all others are tri-stated.

### 6. PPU Internal Register Selection {#6._PPU_Internal_Register_Selection}

- CPU writes to \$2000–\$2007
- CPU A0–A2 lines select one of 8 registers

| CPU A0–A2 | PPU Register |
|-----------|--------------|
| 000       | \$2000       |
| 001       | \$2001       |
| 010       | \$2002       |
| 011       | \$2003       |
| 100       | \$2004       |
| 101       | \$2005       |
| 110       | \$2006       |
| 111       | \$2007       |

Mirrors occur every 8 bytes up to \$3FFF.

### 7. High-order bits in CPU address decoder for CE {#7._High-order_bits_in_CPU_address_decoder_for_CE}

The NES motherboard looks at CPU address lines (A13–A15, sometimes A12) to decide whether the PPU chip should be enabled.

- This is purely for chip selection — to activate CE and tri-state all other devices.
- The decoder doesn’t know or care about VRAM addresses inside the PPU; it just asks: “Does this CPU address fall in the PPU register range (\$2000–\$3FFF)?”
- All addresses in that range look identical to the CE logic, because it only observes the high-order CPU address bits.

### 8. Key Takeaways {#8._Key_Takeaways}

- Each component has unique high-order CPU address bits → CE lines are never conflicted.
- Mirrors are ignored by the decoder using don’t-care bits.
- 74LS138 maps 3 high-order bits → 8 CE lines (only 3–4 used on NES motherboard).
- Tri-state bus logic ensures only the selected device drives the data bus.
- PPU register selection is handled internally by A0–A2 of the CPU address.
