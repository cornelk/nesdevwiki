---
title: "INLretro"
url: "/wiki/INLretro/"
description: "The INLretro is a multi-consoles cartridge dumper/flasher device created by InfiniteNesLives in 2016."
wikiPage: "INLretro"
wikiCategories:
  - name: "Dumper and Flasher devices"
    url: "/category/Dumper_and_Flasher_devices/"
---

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/INLretroOriginal.jpg"><img src="../../wiki-images/INLretroOriginal.jpg" width="300" height="94" /></a>
<figcaption>INLretro</figcaption>
</figure>

The INLretro is a multi-consoles cartridge dumper/flasher device created by InfiniteNesLives in 2016.

Since 2018, it replaced the [Kazzo](../Kazzo/) device.

Since 2021, the project seems to be closed.

## <span id="Hardware_.28INLretro.29"></span>Hardware (INLretro) {#Hardware_.28INLretro.29}

The board includes 6 connectors for various consoles :

- Nintendo NES (72 pins)
- Famicom (60 pins)
- Super Nintendo (SNES) / SFC (62 pins)
- Sega Genesis / MegaDrive (MD) (64 pins)
- Nintendo 64 (N64) (50 pins)
- Gameboy / GameBoyAdvance (GBA) (32 pins)

It works with a [STM32F070RB](https://www.st.com/en/microcontrollers-microprocessors/stm32f070rb.html) microcontroller and a 74HC374 chip.

It needs to be connected to an USB port of a computer (for example), so it can be used without an original console (contrary to [CopyNES](../CopyNES/)).

[STM32F070RB](https://www.st.com/en/microcontrollers-microprocessors/stm32f070rb.html) microcontroller is an Arm Cortex-M0 32-bit RISC core operating at a 48 MHz frequency with 128 Kbytes of Flash memory that contains the INLretro firmware that can be modified and updated.

## <span id="Software_.28INLretro-progdump.29"></span>Software (INLretro-progdump) {#Software_.28INLretro-progdump.29}

INLretro hardware works with the associated software ***INLretro-progdump*** which works on Windows, Linux and Mac.

The official software is available on Gitlab : <https://gitlab.com/InfiniteNesLives/INL-retro-progdump>

It uses LUA scripts for dumping or flashing cartridges.

For now, it supports essentially :

- FAMICOM and Nintendo NES cartridges
- SNES/SFC HiROM et LoROM cartridges

The explanations in the following paragraphs are based on the Linux version.

## <span id="Compilation.2FInstallation_and_Execution"></span>Compilation/Installation and Execution {#Compilation.2FInstallation_and_Execution}

Commands to execute in a terminal :

- <u>***Dependencies installation :***</u>

<!-- -->

       sudo apt-get install libusb-1.0-0-dev

- <u>***INLretro compilation / installation :***</u>

<!-- -->

       mkdir INLretro
       cd INLretro
       git clone https://gitlab.com/InfiniteNesLives/INL-retro-progdump.git
       cd INL-retro-progdump/host
       make unix

<u>***Note :***</u> Compilation warnings can be ignored.

- <u>***Execution :***</u>

<!-- -->

       su
       ./inlretro -c NES -s scripts/inlretro2.lua -m nrom -x 256 -y 128 -d dumpfile.nes

<u>***Note :***</u> For Windows, the executable is already compiled in the ***INL-retro-progdump-master/host*** directory, and the syntax is "***inlretro.exe***" in replacement of "***./inlretro***"

## Usage {#Usage}

Syntax :

     ./inlretro -c [CONSOLE] -s scripts/inlretro2.lua -m [DUMP_SCRIPT] -x [PRG_SIZE_IN_KB] -y [CHR_SIZE_IN_KB] -d dumpfile.nes
     ./inlretro -c [CONSOLE] -s scripts/inlretro2.lua -m [DUMP_SCRIPT] -x [PRG_SIZE_IN_KB] -y [CHR_SIZE_IN_KB] -p flashfile.nes

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>Parameter</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td><em><strong>[CONSOLE]</strong></em></td>
<td>This argument can be : NES, GBA, N64, GENESIS
<p>"<em><strong>NES</strong></em>" is used for NES or FAMICOM cartridges</p></td>
</tr>
<tr>
<td><em><strong>[DUMP_SCRIPT]</strong></em></td>
<td>For NES console, this argument can be : action53, bnrom, cdream, cninja, cnrom, dualport, easynsf, fme7, mapper30, mmc1, mmc3, mmc4, mmc5, nrom, unrom
<p>or another mapper that can be added in <em><strong>INL-retro-progdump-master/host/scripts/nes</strong></em> directory, and in the <em><strong>mappers</strong></em> variable in the <em><strong>nes_exec()</strong></em> function of <em><strong>INL-retro-progdump-master/host/scripts/inlretro2.lua</strong></em></p></td>
</tr>
<tr>
<td><em><strong>[PRG_SIZE_IN_KB]</strong></em></td>
<td>The size of the PRG in KBytes</td>
</tr>
<tr>
<td><em><strong>[CHR_SIZE_IN_KB]</strong></em></td>
<td>The size of the CHR in KBytes</td>
</tr>
<tr>
<td><em><strong>-d dumpfile.nes</strong></em></td>
<td>For dumping the cartridge in the file <em><strong>dumpfile.nes</strong></em></td>
</tr>
<tr>
<td><em><strong>-p flashfile.nes</strong></em></td>
<td>For flashing the file <em><strong>flashfile.nes</strong></em> in the cartridge</td>
</tr>
</tbody>
</table>

**Example :**

     ./inlretro -c NES -s scripts/inlretro2.lua -m mmc3 -x 256 -y 128 -d SuperMarioBros3.nes

This command permits to dump the cartridge "***Super Mario Bros 3***" which is ***MMC3*** mapper and have ***256*** Kb of PRG and ***128*** Kb of CHR.

## INLretro scripts {#INLretro_scripts}

INLretro-progdump uses LUA scripts to dump or flash cartridges.

LUA scripts for NES or FAMICOM cartridges are in the directory : **INL-retro-progdump-master/host/scripts/nes**

The LUA syntax is heavier and more complex than the Anago syntax, but it offers additional capabilities. An INLretro script can be 7 times larger or more than an Anago script for the same mapper.

### Add and run a new LUA dumping script {#Add_and_run_a_new_LUA_dumping_script}

Basically, a LUA dumping script has the structure :

<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>mapperXXXX.lua</th>
</tr>
</thead>
<tbody>
<tr>
<td>local mapperXXXX = {}<br />
<br />
local function create_header( file, prgKB, chrKB )<br />
end<br />
local function init_mapper( debug )<br />
end<br />
local function dump_prgrom( file, rom_size_KB, debug )<br />
end<br />
local function dump_chrrom( file, rom_size_KB, debug )<br />
end<br />
local function process( process_opts, console_opts )<br />
end<br />
<br />
mapperXXXX.process = process<br />
return mapperXXXX</td>
</tr>
</tbody>
</table>

***XXXX*** is the mapper's number.

The name of the dumping script needs to be identical than the local variable inside : ***mapperXXXX***.lua

<u>**Warning :**</u> Using lowercase letters and numbers for the variable and the filename works fine, but using capital letters isn't possible because the program is case-insensitive.

For adding a script, there is 2 steps :

1.  Save the new script with .lua extension in the directory **INL-retro-progdump-master/host/scripts/nes**
2.  Add the new script name in the **mappers** variable of **nes_exec()** function in the main application logic script **INL-retro-progdump-master/host/scripts/inlretro2.lua**  :

<!-- -->

     [...]
     function nes_exec(process_opts, console_opts)
         [...]
         local mappers = {
             action53_tsop = require "scripts.nes.action53_tsop",
             action53 = require "scripts.nes.action53",
             bnrom = require "scripts.nes.bnrom",
             [...],
             mapperXXXX = require "scripts.nes.mapperXXXX"
         }
     [...]

The new script can be run in command-line with :

     ./inlretro -c NES -s scripts/inlretro2.lua -m mapperXXXX -x 256 -y 128 -d dumpfile.nes

### Content of a LUA script {#Content_of_a_LUA_script}

Contrary to Anago, the names of the functions in the LUA scripts are not defined in the INLretro 's software code (except the *process()* function) but in the LUA script itself, so the functions can be named freely. But for convention and lisibility of scripts, it's better to use the same as those existant.

A basic LUA dumping script contains 5 functions.

Several mappers may have additional functions like *dump_wram()*, but this is uncommon.

For a basic LUA flashing script, replace *dump_prgrom()* and *dump_chrrom()* functions by the 2 functions :

       local function flash_prgrom(file, rom_size_KB, debug)
       ...
       end
       local function flash_chrrom(file, rom_size_KB, debug)
       ...
       end

A complete LUA script contains both dumping and flashing functions.

- ***Function create_header()***

This function writes the header in the dump file. The common content is something like :

     local function create_header( file, prgKB, chrKB )
         nes.write_header( file, prgKB, chrKB, mapperNumber, mirroring )
     end

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td><em><strong>file</strong></em><br />
<em><strong>prgKB</strong></em><br />
<em><strong>chrKB</strong></em></td>
<td>No modification needed on these parameters that are transmitted by the command-line.</td>
</tr>
<tr>
<td><em><strong>mapperNumber</strong></em></td>
<td>The number of the mapper to write in the header of the dump file.<br />
The value type is Integer.</td>
</tr>
<tr>
<td><em><strong>mirroring</strong></em></td>
<td>If mirroring=<em><strong>"VERT"</strong></em>, mirroring is set to vertical in the header of the dump file.<br />
Else (for example : <em><strong>"HORZ"</strong></em>), mirroring is set to horizontal in the header of the dump file.<br />
The value type is String.</td>
</tr>
</tbody>
</table>

**Example :**

     local function create_header( file, prgKB, chrKB )
         nes.write_header( file, prgKB, chrKB, 58, "VERT")
     end

In this example, the mapper is set to ***58*** and the mirroring to ***vertical*** in the header of the dump file.

- ***Function init_mapper()***

This function is only used for mappers that need to be initialized like MMC1, FME7 or others.

In many cases, this function is empty.

- ***Function dump_prgrom()***

This function contains the method for dumping PRG rom. Basically, for a cartridge whose PRG banks are ***16Kb*** and dump formula is ***0x8000 \| i*** (i : iteration counter), this function can be :

     local function dump_prgrom( file, rom_size_KB, debug )
         --PRG-ROM dump 16KB at a time
         local KB_per_read = 16
         local num_reads = rom_size_KB / KB_per_read
         local read_count = 0
         local addr_base = 0x80   -- $8000

         print("CPU dump :")
         while ( read_count < num_reads )
         do
             if debug then print( "dump PRG part ", read_count, " of ", num_reads) end

             --optional one-second pause to reduce stress on the chips--      
             local time=os.clock()+1
             while time>os.clock()
             do
             end

             local addrpart = 0x8000 | read_count
             local datapart = 0

             dict.nes("NES_CPU_WR", addrpart , datapart)
             dump.dumptofile( file, KB_per_read, addr_base, "NESCPU_PAGE", false )

             read_count = read_count + 1
         end
     end

<u>***Main functions used :***</u>

<u>*1) function nes(opcode, addrpart , datapart)*</u>

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td><em><strong>opcode</strong></em></td>
<td>Value can be <em><strong>"NES_CPU_WR"</strong></em> , <em><strong>"NES_PPU_WR"</strong></em> , <em><strong>"M2_LOW_WR"</strong></em> , <em><strong>"M2_HIGH_WR"</strong></em> , <em><strong>"NES_CPU_RD"</strong></em> , <em><strong>"NES_PPU_RD"</strong></em> ...<br />
All possible values of opcode are in <em><strong>INL-retro-progdump-master/shared/shared_dict_nes.h</strong></em><br />
The value type is String.</td>
</tr>
<tr>
<td><em><strong>addrpart</strong></em></td>
<td>Formula depending of the base address for the dump and the iteration variable <em><strong>read_count</strong></em> based on the <em><strong>CPU_Ax</strong></em> description (x=0 to 15).<br />
Example : <em><strong>0x8000 | read_count</strong></em></td>
</tr>
<tr>
<td><em><strong>datapart</strong></em></td>
<td>Formula depending of the iteration variable <em><strong>read_count</strong></em> based on the <em><strong>CPU_Dx</strong></em> description (x=0 to 7).<br />
Example : <em><strong>read_count</strong></em></td>
</tr>
</tbody>
</table>

If ***opcode = "NES_CPU_WR"***, this function is equivalent to the Anago function : ***cpu_write(d, addrpart, datapart);***.

<u>*2) function dumptofile(file, sizeKB, map, mem, debug)*</u>

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td><em><strong>file</strong></em></td>
<td>Output file.</td>
</tr>
<tr>
<td><em><strong>sizeKB</strong></em></td>
<td>Size of the read in KBytes.</td>
</tr>
<tr>
<td><em><strong>map</strong></em></td>
<td>Base address where to read datas. Cf Warning below.</td>
</tr>
<tr>
<td><em><strong>mem</strong></em></td>
<td>Size of the page to read.<br />
Value can be : <em><strong>NESCPU_PAGE</strong></em> , <em><strong>NESPPU_PAGE</strong></em> , <em><strong>NESCPU_4KB</strong></em> , <em><strong>NESPPU_1KB</strong></em><br />
All possible values of mem are in <em><strong>INL-retro-progdump-master/shared/shared_dict_buffer.h</strong></em></td>
</tr>
<tr>
<td><em><strong>debug</strong></em></td>
<td><em><strong>true</strong></em> for debug<br />
<em><strong>false</strong></em> for no debug</td>
</tr>
</tbody>
</table>

This function is equivalent to the Anago function : ***cpu_read(d, address, banksize)*** with ***map*** corresponding to ***address*** and ***sizeKB = banksize / 1024***.

<u>***Warning :***</u>  

In a lot of case, if ***address=0x8000*** with Anago, it gives ***map=0x80*** with INLretro.

But for several mappers (MMC1, FME7, ...), ***address=0x8000*** with Anago is traduced by ***map=0x08*** with INLretro.

If the dump file contains only 0xFF bytes, it can be useful to try to replace ***map=0x80*** by ***map=0x08*** for getting datas from the cartridge.

- ***Function dump_chrrom()***

This function contains the method for dumping CHR rom. Basically, for a cartridge whose CHR banks are ***8Kb*** and dump formula is ***0x8000 \| i*** (i : iteration counter), this function can be :

     local function dump_chrrom( file, rom_size_KB, debug )
         --CHR-ROM dump 8KB at a time
         local KB_per_read = 8
         local num_reads = rom_size_KB / KB_per_read
         local read_count = 0
         local addr_base = 0x00   -- $0000

         print("PPU dump :")
         while ( read_count < num_reads )
         do
             if debug then print( "dump CHR part ", read_count, " of ", num_reads) end

             --optional one-second pause to reduce stress on the chips--      
             local time=os.clock()+1
             while time>os.clock()
             do
             end

             dict.nes("NES_CPU_WR", 0x8000 | read_count , 0)
             dump.dumptofile( file, KB_per_read, addr_base, "NESPPU_PAGE", false )

             read_count = read_count + 1
         end
     end

- ***Function process()***

This function is the entry point of the script. It's like the ***main*** function in C. Its role is to run the other functions :

     local function process(process_opts, console_opts)
         [...]
         init_mapper()
         [...]
         create_header(file, prg_size, chr_size)
         [...]
         dump_prgrom(file, prg_size, false)
         dump_chrrom(file, chr_size, false)
         [...]
     end

### <span id="How_to_convert_an_Anago_script_in_an_INLretro_LUA_script_.3F"></span>How to convert an Anago script in an INLretro LUA script ? {#How_to_convert_an_Anago_script_in_an_INLretro_LUA_script_.3F}

This table describes differences and correspondences between the two syntaxes : the left part in Anago syntax can be traduced by the right part in INLretro syntax.

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>Anago</th>
<th>INLretro</th>
</tr>
</thead>
<tbody>
<tr>
<td><u>Mapper and mirroring in the <em>board</em> variable :</u><br />
<br />
board &lt;- {<br />
   ... mappernum = <em><strong>194</strong></em>, ...<br />
   ... vram_mirrorfind = <em><strong>true</strong></em><br />
};<br />
<br />
Note :<br />
   vram_mirrorfind = true is vertical mirroring<br />
   vram_mirrorfind = false is horizontal mirroring</td>
<td><u>Mapper and mirroring in the <em>create_header()</em> function :</u><br />
<br />
<br />
local function create_header( file, prgKB, chrKB )<br />
   nes.write_header( file, prgKB, chrKB, <em><strong>194</strong></em>, "<em><strong>VERT</strong></em>")<br />
end<br />
<br />
<br />
<br />
<br />
</td>
</tr>
<tr>
<td><u>PRG size and CHR size in the script and in bytes :</u><br />
<br />
board &lt;- {<br />
   ... cpu_romsize = <em><strong>262144</strong></em>, ...<br />
   ... ppu_romsize = <em><strong>131072</strong></em>, ...<br />
};<br />
<br />
Note :<br />
   262144 bytes = 256 Kbytes<br />
   131072 bytes = 128 Kbytes</td>
<td><u>PRG size and CHR size in the command-line and in Kbytes :</u><br />
<br />
<br />
<br />
./inlretro -c NES -s scripts/inlretro2.lua -m mmc3 -x <em><strong>256</strong></em> -y <em><strong>128</strong></em> -d SuperMarioBros3.nes<br />
<br />
<br />
Note:<br />
   Commonly used syntax, but it can also be set in the script with :<br />
   variables <em>prg_size</em> and <em>chr_size</em> in <em>process()</em> function.<br />
</td>
</tr>
<tr>
<td><u>CPU banksize and PPU banksize in bytes in the <em>board</em> variable:</u><br />
<br />
<br />
<br />
<br />
<br />
board &lt;- {<br />
   ... cpu_banksize = <em><strong>0x4000</strong></em>, ...<br />
   ... ppu_banksize = <em><strong>0x2000</strong></em>, ...<br />
};<br />
<br />
<br />
<br />
<br />
<br />
Note :<br />
   0x4000 = 16384 bytes = 16 Kbytes<br />
   0x2000 = 8192 bytes = 8 Kbytes<br />
</td>
<td><u>CPU banksize and PPU banksize in Kbytes in the <em>dump_prgrom()</em> and <em>dump_chrrom()</em> functions :</u><br />
<br />
...<br />
local function dump_prgrom( file, rom_size_KB, debug )<br />
   --PRG-ROM dump <em><strong>16KB</strong></em> at a time<br />
   local KB_per_read = <em><strong>16</strong></em><br />
   ...<br />
end<br />
<br />
...<br />
<br />
local function dump_chrrom( file, rom_size_KB, debug )<br />
   --CHR-ROM dump <em><strong>8KB</strong></em> at a time<br />
   local KB_per_read = <em><strong>8</strong></em><br />
   ...<br />
end<br />
...</td>
</tr>
<tr>
<td><u>Write address and data formula in <em>cpu_dump()</em> and <em>ppu_dump()</em> functions :</u><br />
<br />
<em><strong>cpu_write</strong></em>(d, <em><strong>0x8000 | i</strong></em>, <em><strong>0</strong></em>);<br />
<br />
</td>
<td><u>Write address and data formula in <em>dump_prgrom()</em> and <em>dump_chrrom()</em> functions :</u><br />
<br />
<br />
dict.nes("NES_<em><strong>CPU_WR</strong></em>", <em><strong>0x8000 | read_count</strong></em>, <em><strong>0</strong></em>)<br />
<br />
</td>
</tr>
<tr>
<td><u>Read address in <em>cpu_dump()</em> and <em>ppu_dump()</em> functions :</u><br />
<br />
<br />
cpu_read(d, <em><strong>0x8000</strong></em>, <em><strong>banksize</strong></em>);<br />
<br />
<br />
<br />
<br />
<br />
<br />
</td>
<td><u>Read address in <em>dump_prgrom()</em> and <em>dump_chrrom()</em> functions :</u><br />
<br />
...<br />
dump.dumptofile( file, <em><strong>KB_per_read</strong></em>, <em><strong>0x80</strong></em>, "NESCPU_PAGE", false )<br />
...<br />
<br />
Note :<br />
   <em>KB_per_read</em> corresponds to <em>banksize</em> in the Anago script converted in Kbytes<br />
   <em>0x80</em> corresponds to address <em>0x8000</em> in the Anago script<br />
   <em>NESCPU_PAGE</em> is commonly used because it gets all the datas in the page</td>
</tr>
<tr>
<td><u>Read 2 consecutive 16Kb banks (0x8000 + 0x4000 = 0xc000) :</u><br />
<br />
board &lt;- {<br />
   ... cpu_banksize = <em><strong>0x8000</strong></em>, ...<br />
};<br />
...<br />
cpu_read(d, <em><strong>0x8000</strong></em>, 0x4000);<br />
cpu_read(d, 0xc000, 0x4000);<br />
...<br />
<br />
Note :<br />
   0x8000 = 32768 bytes = 32 Kbytes<br />
   0x4000 = 16384 bytes = 16 Kbytes<br />
<br />
The 2 <em>cpu_read()</em> have to be separated to avoid Anago error :<br />
length range must be 0x000000 to 0x004000</td>
<td><u>Read 1 x 32Kb bank :</u><br />
<br />
...<br />
local KB_per_read = <em><strong>32</strong></em><br />
...<br />
local addr_base = <em><strong>0x80</strong></em> -- $8000<br />
...<br />
dump.dumptofile( file, <em><strong>KB_per_read</strong></em>, addr_base, "NESCPU_PAGE", false )<br />
...<br />
<br />
<br />
<br />
Note : With INLretro getting a 32Kb bank doesn't generate any error<br />
<br />
<br />
<br />
</td>
</tr>
</tbody>
</table>

## Common Errors {#Common_Errors}

### Plug error {#Plug_error}

       [ERROR] (source/usb_operations.c:159: errno: Resource temporarily unavailable) Could not find INL retro-prog USB device
       open_usb_device went to error
       freeing device list
       deinitializing libusb
       [ERROR] (source/inlprog.c:361: errno: None) Unable to open INL retro-prog usb device handle.
       Fatal error encountered, exiting.

<u>**Solution :**</u> The device isn't plugged to the USB port.

### USB permission error {#USB_permission_error}

       [ERROR] (source/usb_operations.c:110: errno: Permission denied) Unable to open USB device: Access denied (insufficient permissions)., 

       DEVICE FOUND, BUT CAN'T OPEN DEVICE, VERIFY DRIVERS ARE INSTALLED!!!
       open_usb_device went to error
       freeing device list
       deinitializing libusb
       -------------------------------------------------------
       Denied Permission is expected for initial use on Linux.
       See udev-rule-help/Readme.txt in host dir for help gaining permission.
       Try command with sudo for a cheap temporary solution.
       [ERROR] (source/inlprog.c:361: errno: None) Unable to open INL retro-prog usb device handle.
       Fatal error encountered, exiting.

<u>**Solution :**</u> This error indicates that the user doesn't have the permission for accessing to USB device.

A simple workaround is to run INLretro in ***root*** mode in a terminal, though this causes dumped ROM files to have root ownership/permissions :

       su
       ./inlretro -c NES -s scripts/inlretro2.lua -m nrom -x 256 -y 128 -d dumpfile.nes

or :

       sudo ./inlretro -c NES -s scripts/inlretro2.lua -m nrom -x 256 -y 128 -d dumpfile.nes

The proper solution is to create a [udev](https://wiki.archlinux.org/title/Udev) rule to add the INLretro device to a group with appropriate USB/serial permissions, such as *uucp* on Arch.

### Mapper error {#Mapper_error}

       Successfully found and connected to INL retro-prog
       Device firmware version: 2.3.x 
       firmware app ver request:    3
       UNSUPPORTED MAPPER:  mapperXXXX

<u>**Solution :**</u> This error indicates that the mapper isn't in the list. Go to [Add a script](./#Add_and_run_a_new_LUA_dumping_script) section to add it.

### 0xFF error {#0xFF_error}

All seem to be correct in the process but the dump file contains only 0xFF bytes.

<u>**Solutions :**</u>  

1.  Unplug the INLretro dumper from the computer for a while to let it rest. Plug it back in after several minutes.
2.  Try to replace ***addr_base = 0x80*** by ***0x08*** in ***dumptofile(file, KB_per_read, addr_base, "NESCPU_PAGE", false)*** function in the dumping script.

## Related Links {#Related_Links}

- [Original announcement of INLretro dumper](https://forums.nesdev.org/viewtopic.php?p=202107&hilit=INLretro#p202107)

<!-- -->

- [INL's Website for INLretro dumper](https://www.infiniteneslives.com/inlretro.php)

<!-- -->

- [Official INLretro-progdump software on Gitlab](https://gitlab.com/InfiniteNesLives/INL-retro-progdump)
