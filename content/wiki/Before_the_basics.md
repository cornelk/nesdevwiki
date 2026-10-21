---
title: "Before the basics"
url: "/wiki/Before_the_basics/"
description: "A digital computer, such as a PC or NES, is filled with millions of tiny switches that can be turned on or off. Some switches store information for short time; these are called memory. Other switches…"
wikiPage: "Before_the_basics"
---

## Background {#Background}

A digital computer, such as a PC or NES, is filled with millions of tiny switches that can be turned on or off. Some switches store information for short time; these are called memory. Other switches calculate things based on the output of other switches; these are called logic. Each switch can be on or off, representing true or false, or the numbers 1 or 0. Multiple switches grouped together can represent larger values. A CPU reads machine code from a memory and treats it as instructions to perform arithmetic and logic operations on other values in memory or to perform input and output. Some of the input and output involve user interaction: the user gives commands to the computer, and the computer displays the result.

If you're confused, now is the opportunity to become familiar with the following basic concepts from computer engineering and computer science. These articles from Wikipedia are in English, and many have been translated into other languages.

- [Computer](https://en.wikipedia.org/wiki/Computer) (it's not just a PC)
- [Bit](https://en.wikipedia.org/wiki/Bit)
  - [Boolean logic](https://en.wikipedia.org/wiki/Boolean_algebra)
- [Binary number](https://en.wikipedia.org/wiki/Binary_number)
  - [Byte](https://en.wikipedia.org/wiki/Byte)
  - [Arithmetic logic unit](https://en.wikipedia.org/wiki/Arithmetic_logic_unit)
  - [Binary prefix](https://en.wikipedia.org/wiki/Binary_prefix) (KiB, MiB)
- [Random-access memory](https://en.wikipedia.org/wiki/Random-access_memory) (RAM)
- [Read-only memory](https://en.wikipedia.org/wiki/Read-only_memory) (ROM)
- [Central processing unit](https://en.wikipedia.org/wiki/Central_processing_unit) (CPU)
  - [MOS Technology 6502](https://en.wikipedia.org/wiki/MOS_Technology_6502)
  - [Clock rate](https://en.wikipedia.org/wiki/Clock_rate)
  - [Processor register](https://en.wikipedia.org/wiki/Processor_register)
- [Computer program](https://en.wikipedia.org/wiki/Computer_program)
  - [Programming language](https://en.wikipedia.org/wiki/Programming_language)
  - [Machine code](https://en.wikipedia.org/wiki/Machine_code)
  - [Assembly language](https://en.wikipedia.org/wiki/Assembly_language) (asm)
  - [Source code](https://en.wikipedia.org/wiki/Source_code)
  - [Object file](https://en.wikipedia.org/wiki/Object_file)
  - [Parallel computing](https://en.wikipedia.org/wiki/Parallel_computing)
  - [Thread](https://en.wikipedia.org/wiki/Thread_(computing))
- [Input/output](https://en.wikipedia.org/wiki/Input/output) (I/O)
  - [Memory-mapped I/O](https://en.wikipedia.org/wiki/Memory-mapped_I/O)
  - [Serial port](https://en.wikipedia.org/wiki/Serial_port)
  - [Interrupt](https://en.wikipedia.org/wiki/Interrupt)
  - [Serial Peripheral Interface Bus](https://en.wikipedia.org/wiki/Serial_Peripheral_Interface_Bus) (SPI)
  - [Direct Memory Access](https://en.wikipedia.org/wiki/Direct_memory_access) (DMA)
- [Raster graphics](https://en.wikipedia.org/wiki/Raster_graphics)
  - [Pixel](https://en.wikipedia.org/wiki/Pixel)
  - [Raster scan](https://en.wikipedia.org/wiki/Raster_scan)
  - [Bitmap](https://en.wikipedia.org/wiki/Bitmap)
  - [Planar bitmap](https://en.wikipedia.org/wiki/Planar_(computer_graphics))
  - [Progressive scan](https://en.wikipedia.org/wiki/Progressive_scan)
  - [Vertical blanking interval](https://en.wikipedia.org/wiki/Vertical_blanking_interval) (vblank)
  - [Text mode](https://en.wikipedia.org/wiki/Text_mode)
  - [Text-based user interface](https://en.wikipedia.org/wiki/Text-based_user_interface)
  - [ASCII art](https://en.wikipedia.org/wiki/ASCII_art)
  - [Sprite](https://en.wikipedia.org/wiki/Sprite_(computer_graphics))
  - [Display list](https://en.wikipedia.org/wiki/Display_list)
- [Command-line interface](https://en.wikipedia.org/wiki/Command-line_interface)
  - [Shell](https://en.wikipedia.org/wiki/Shell_(computing))
  - [Windows Command Prompt](https://en.wikipedia.org/wiki/cmd.exe)
  - [Batch file](https://en.wikipedia.org/wiki/Batch_file)
  - [Plug-in](https://en.wikipedia.org/wiki/Plug-in_(computing))
  - [Filename extension](https://en.wikipedia.org/wiki/Filename_extension)
  - [Zip file](https://en.wikipedia.org/wiki/Zip_(file_format))
  - [Environment variable](https://en.wikipedia.org/wiki/Environment_variable)

### More advanced topics {#More_advanced_topics}

It has been suggested that understanding these topics is not necessary to program low-complexity games in assembly language without sound, such as a workalike of Magnavox Odyssey tennis. They are held here for forthcoming sorting into a separate page.

- Games larger than one screen
  - [Scrolling](https://en.wikipedia.org/wiki/Scrolling)
- Audio
  - [Digital signal](https://en.wikipedia.org/wiki/Digital_signal)
  - [Digital audio](https://en.wikipedia.org/wiki/Digital_audio)
  - [Signal generator](https://en.wikipedia.org/wiki/Signal_generator)
  - [Delta modulation](https://en.wikipedia.org/wiki/Delta_modulation)
  - [Digital frequency divider](https://en.wikipedia.org/wiki/Frequency_divider#Digital_dividers) (also [on Wikia](http://community.fandom.com/wiki/c:electronicmusic:Pulse_divider))
- Mapper
  - [Logic gate](https://en.wikipedia.org/wiki/Logic_gate)
  - [Linear feedback shift register](https://en.wikipedia.org/wiki/Linear_feedback_shift_register)
  - [Ring counter](https://en.wikipedia.org/wiki/Ring_counter)
  - [Bus](https://en.wikipedia.org/wiki/Bus_(computing))
  - [Bank switching](https://en.wikipedia.org/wiki/Bank_switching)
  - [Integrated circuit](https://en.wikipedia.org/wiki/Integrated_circuit) (IC)
  - [7400 series](https://en.wikipedia.org/wiki/7400_series)
- High-level languages
  - [Compiler](https://en.wikipedia.org/wiki/Compiler)

## Introducing the NES {#Introducing_the_NES}

The Nintendo Entertainment System has the following components:

- 2A03 CPU IC made by Ricoh
  - CPU based on the MOS Technology 6502 8-bit microprocessor
  - serial input for game controllers
  - audio output comprising four tone generators and a delta modulation playback device
- 2 KiB of RAM for use by the CPU
- 2C02 PPU (picture processing unit)
  - tile-based background image
  - 64 sprites (individually moving objects)
  - 25 colors out of 53
  - 256x240 pixel progressive picture generator
  - NTSC color encoder
- 2 KiB of RAM for use by the PPU

Cartridges have the following components:

- 16 KiB or more PRG ROM, for use by the CPU
- 8 KiB or more CHR ROM or CHR RAM, for use by the PPU (and in some cases by the CPU)
- (optional) Bank switching hardware for the ROMs
- (optional) Logic to generate interrupts

## Important Considerations {#Important_Considerations}

The NES is not an easy system to develop games for anybody who does not have a good information and knowledge about the low-level workings of a computer. Developing games for NES can be tricky for developers who are used to tools like Unreal Engine, Unity, Godot, and game libraries and frameworks like Raylib, OpenGL, etc.

A few of such (at first) hard-to-grasp concepts and difficulties that one might face during their journey are:

- Memory-mapped I/O, memory mirroring, incomplete address decoding and mapping and mappers can be confusing to understand and deal with, at first.
- Issues like race conditions can pop up during the journey, such as not being able to write to video memory while the PPU is busy rendering the screen (you have to wait until the Vertical Blanking interval begins, of which the PPU can notify using an interrupt).
- One cannot have "pre-initialized" global variables unless coded manually (due to program residing entirely in the ROM).

If you do not have a good programming knowledge, then you should firstly consider getting programming experience. Knowledge of the 6502 assembly language is important and knowledge and experience in C programming language is extremely helpful, as it is often used and compiled and linked together with the 6502 assembly.

Not understanding concepts clearly such as memory-management and not having enough technical knowledge about the hardware can give one a hard time developing tools or games on the NES, such as if you don't understand the memory-range of the hardware and you allocate memory outside the range of what hardware can access, you might get bugs and problems popping up in your code. Hence, having a good grasp of assembly is essential as it allows you to gain experience to the memory mappings of the hardware and management of the same.

And most importantly, **one should not expect much from the NES hardware**: its processor only runs at 1.8 MHz (compared to a modern processor like the Intel i5-13400 which runs at 2.5 GHz), and all instructions take multiple cycles to execute (so it's only running between 300,000 and 900,000 instructions per second). The PPU on the NES spends exactly *1 clock cycle per pixel*. Another important thing is that **the PPU and CPU do not share any memory** - both have their own buses, and the ability to transfer data between them has significant limits.

## External links {#External_links}

- [Easy 6502](https://skilldrick.github.io/easy6502/)
