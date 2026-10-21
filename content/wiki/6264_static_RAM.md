---
title: "6264 static RAM"
url: "/wiki/6264_static_RAM/"
description: "The 6264 is an 8 KiB static RAM, available in 70 to 200 nanosecond access time variants. It can function on both the NES PPU's 8080 style bus (separate /WE and /OE strobes) or on the CPU's 6500 style…"
wikiPage: "6264_static_RAM"
wikiCategories:
  - name: "Integrated circuits"
    url: "/category/Integrated_circuits/"
---

The **6264** is an 8 KiB static RAM, available in 70 to 200 nanosecond access time variants. It can function on both the NES PPU's 8080 style bus (separate /WE and /OE strobes) or on the CPU's 6500 style bus (by grounding /OE and connecting the R/W signal to /WE).

## 6116 {#6116}

Certain Konami VRC boards take a 24-pin 2Kx8 bit SRAM:

          ,---\/---.
     A7 - | 1    24| - VCC
     A6 - | 2    23| - A8
     A5 - | 3    22| - A9
     A4 - | 4    21| - /WE
     A3 - | 5    20| - /OE
     A2 - | 6    19| - A10
     A1 - | 7    18| - /CS1
     A0 - | 8    17| - D7
     D0 - | 9    16| - D6
     D1 - |10    15| - D5
     D2 - |11    14| - D4
    GND - |12    13| - D3
          `--------'

(Source: [MB8416](http://matthieu.benoit.free.fr/cross/data_sheets/MB8416-20.pdf); [6116](https://www.idt.com/document/dst/6116sala-data-sheet))

## <span id="6264_Pinout_.2862256_in_parentheses.29"></span>6264 Pinout (62256 in parentheses) {#6264_Pinout_.2862256_in_parentheses.29}

The 6264 has been used for work RAM and save RAM in numerous NES games, commonly [decoded at \$6000-\$7FFF](../PRG_RAM_circuit/).

The 62256 is used in *RacerMate Challenge II*, *Romance of the Three Kingdoms II*, and a few Famicom games.

Note that 62256 does not have second enable line.

- 6264: 8Kx8 bit SRAM
- 62256: 32Kx8 bit SRAM

<!-- -->

                ,----\/----.
     (A14) NC - | 1      28| - VCC
          A12 - | 2      27| - /WE
           A7 - | 3      26| - CS2 (A13) 
           A6 - | 4      25| - A8      
           A5 - | 5      24| - A9      
           A4 - | 6      23| - A11       
           A3 - | 7      22| - /OE       
           A2 - | 8      21| - A10       
           A1 - | 9      20| - /CS1    
           A0 - |10      19| - D7      
           D0 - |11      18| - D6      
           D1 - |12      17| - D5      
           D2 - |13      16| - D4      
          GND - |14      15| - D3      
                `----------'

## <span id="621024_Pinout_.2862512_in_parentheses.29"></span>621024 Pinout (62512 in parentheses) {#621024_Pinout_.2862512_in_parentheses.29}

In the NES, these chips might be used in "EPROM emulators", or devices into which a game is loaded from an external source (such as a connected personal computer or a NAND flash) before it runs.

Note that 62512 does not have second enable line.

- 621024: 128Kx8 bit SRAM
- 62512: 512Kx8 bit SRAM (breaks name pattern)

<!-- -->

                .----\/----.
     (A18) NC - | 1      32| - VCC
          A16 - | 2      31| - A15
          A14 - | 3      30| - CS2 (A17)
          A12 - | 4      29| - /WE
           A7 - | 5      28| - A13
           A6 - | 6      27| - A8      
           A5 - | 7      26| - A9      
           A4 - | 8      25| - A11       
           A3 - | 9      24| - /OE       
           A2 - |10      23| - A10       
           A1 - |11      22| - /CS1    
           A0 - |12      21| - D7      
           D0 - |13      20| - D6      
           D1 - |14      19| - D5      
           D2 - |15      18| - D4      
          GND - |16      17| - D3      
                `----------'

## Signal descriptions {#Signal_descriptions}

A0-A17  
address lines

D0-D7  
data lines

/CS1 and (when present) CS2 (chip selects 1 and 2)  
when /CS1 is low and CS2 is high, the chip is selected and will accept input or provide output to the data bus.

/WE (write enable)  
when the chip is selected, if this pin is low, the eight-bit data D will be written to the address inside the RAM specified by A.

/OE (output enable)  
when the chip is selected, if this pin is low and /WE is high, the eight bit data in the RAM at address A will be output onto the pins D.

The difference between /OE and the chip selects is that the 6264 responds much faster to /OE, but it draws less power when the chip selects are deasserted.

Note that when the RAM was selected via /CS1 (and CS2) and both /OE and /WE are pulled low, /WE overrides /OE so this works for 2A03's 6500-style bus as well. In other words, if /OE is tied to GND always, /WE is essentially works as R/W signal.

A circuit to generate /CS1 for decoding a 6264 RAM at \$6000-\$7FFF is described at [PRG RAM circuit](../PRG_RAM_circuit/). It can also be used to decode a 62256 if you have latches available to control A13 and A14.
