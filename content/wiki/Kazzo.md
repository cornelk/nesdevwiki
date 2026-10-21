---
title: "Kazzo"
url: "/wiki/Kazzo/"
description: "The Kazzo is a Famicom/NES cartridge dumper/flasher device made by Naruko before 2009."
wikiPage: "Kazzo"
wikiCategories:
  - name: "Dumper and Flasher devices"
    url: "/category/Dumper_and_Flasher_devices/"
---

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/KazzoPrototype.jpg"><img src="../../wiki-images/KazzoPrototype.jpg" width="300" height="91" /></a>
<figcaption>Kazzo - Initial prototype (2011)</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/KazzoOriginalVersion3.0.jpg"><img src="../../wiki-images/KazzoOriginalVersion3.0.jpg" width="300" height="112" /></a>
<figcaption>Kazzo revision 3.0 (2011)</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/KazzoLast.jpg"><img src="../../wiki-images/KazzoLast.jpg" width="300" height="253" /></a>
<figcaption>Kazzo - Last version</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/KazzoAliexpress.jpg"><img src="../../wiki-images/KazzoAliexpress.jpg" width="300" height="339" /></a>
<figcaption>Kazzo copy from Aliexpress</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/AdapterNEStoFAMICOM.jpg"><img src="../../wiki-images/AdapterNEStoFAMICOM.jpg" width="300" height="232" /></a>
<figcaption>Adapter Nintendo NES 72 pins female to FAMICOM 60 pins male</figcaption>
</figure>

<figure class="mw-default-size" typeof="mw:File/Thumb">
<a href="../../wiki-images/KazzoAndAdapter.jpg"><img src="../../wiki-images/KazzoAndAdapter.jpg" width="300" height="180" /></a>
<figcaption>Kazzo with adapter Nintendo NES 72 pins female to FAMICOM 60 pins male</figcaption>
</figure>

The **Kazzo** is a Famicom/NES cartridge dumper/flasher device made by Naruko before 2009.

It is still available on marketplaces, and used to be sold by InfiniteNesLives.

The Anago software is used to operate the device.

In 2018, InfiniteNesLives created the [INLretro](../INLretro/) dumper/flasher to address the Kazzo's limitations.

## <span id="Hardware_.28Kazzo.29"></span>Hardware (Kazzo) {#Hardware_.28Kazzo.29}

The first version of Kazzo is a simple dumper composed of 3 logical chips : ATMEGA164 + 2 x 74HC175 , and only one FAMICOM connector. It needs to be connected to an USB port of a computer (for example), so it can be used without an original console (contrary to [CopyNES](../CopyNES/)). The next version added a Nintendo NES connector.

It allows to dump :

- FAMICOM cartridges
- Nintendo NES cartridges (by using an adapter Nintendo NES / FAMICOM for the first version or copies versions)

Even if the project is closed, several copies of the original Kazzo PCB can be found on MarketPlaces, like Aliexpress. Usually, these copies are built with logical chips like ATMEGA16 + 74HC574, or others. They have either an USB-B port or an USB-C port for connecting to the computer.

These copies work like the original Kazzo dumper, so they can be used with Anago to dump the content of a cartridge, or to flash a blank cartridge.

For the first version or copies versions, dumping a Nintendo NES cartridge requires :

- 1 Kazzo dumper
- 1 adapter Nintendo NES 72 pins female to FAMICOM 60 pins male

## <span id="Software_.28Anago.29"></span>Software (Anago) {#Software_.28Anago.29}

Kazzo dumper works with the Anago software, which is the result of the Unagi project. This software was initially built for Windows, but a portage has been created for Linux. It works in command-line.

AnagoWx is a graphical software for Kazzo dumper under Windows.

Two different versions of Anago using two different syntaxes for scripts exist :

- <u>*Version 0.6.2 (mainline)* :</u>

<!-- -->

       https://github.com/sharkpp/unagi_kazzo
       https://github.com/sharkpp/unagi_client

- <u>*Version 0.6.0l (legacy)* :</u>

<!-- -->

       https://github.com/zerkerX/anago

Scripts syntax differences between the two versions :

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>Anago 0.6.2 mainline</th>
<th>Anago 0.6.0l legacy</th>
</tr>
</thead>
<tbody>
<tr>
<td>cpu_rom = {size_base = 0x8000, size_max = 0x8000, banksize = 0x4000}</td>
<td>cpu_romsize = 0x8000, cpu_banksize = 0x4000</td>
</tr>
<tr>
<td>cpu_ram = {size_base = 0x2000, size_max = 0x2000, banksize = 0x2000}</td>
<td><em><strong>No equivalence</strong></em></td>
</tr>
<tr>
<td>ppu_rom = {size_base = 0x4000, size_max = 0x4000, banksize = 0x2000}</td>
<td>ppu_romsize = 0x4000, ppu_banksize = 0x2000</td>
</tr>
<tr>
<td>dofile("file.ai")</td>
<td><em><strong>No include file possible</strong></em></td>
</tr>
<tr>
<td>local ar = [i, i|2, i|4, i|5];<br />
cpu_write(d, 0x8000, ar);</td>
<td><em><strong>No array possible</strong></em></td>
</tr>
</tbody>
</table>

The legacy version (0.6.0l) is easy to compile and use but has some limitations and works only on Linux system.

The explanations in the following paragraphs are based on the Anago's legacy version.

## <span id="Compilation.2FInstallation_and_Execution"></span>Compilation/Installation and Execution {#Compilation.2FInstallation_and_Execution}

These commands work on a Linux Debian 13 system.

Linux Live versions can be run on an USB key or a CD/DVD without any installation on the computer. A Debian Live version can be downloaded at :

<https://cdimage.debian.org/mirror/cdimage/archive/>

Note : Select directories with ***-live*** suffix, and next, the architecture of the processor (amd64, ...).

Example : Common Debian 13 Live version image :

<https://cdimage.debian.org/mirror/cdimage/archive/13.3.0-live/amd64/iso-hybrid/debian-live-13.3.0-amd64-lxde.iso>

### Anago 0.6.2 mainline {#Anago_0.6.2_mainline}

Commands to execute in a terminal :

- <u>***Dependencies installation :***</u>

<!-- -->

       sudo apt-get install libsquirrel-dev libusb-dev libwxgtk3.2-dev

- <u>***Anago 0.6.2 compilation / installation :***</u>

       mkdir anago-main
       cd anago-main
       git clone https://github.com/sharkpp/unagi_kazzo.git
       git clone https://github.com/sharkpp/unagi_client.git
       mv unagi_kazzo/ kazzo
       cd unagi_client
       mkdir -p SQUIRREL2 && cd SQUIRREL2
       includeDirectory=$(ls -d /usr/include/squirrel*/ | head -n 1)
       libDirectory=$(find /usr/lib* -name "libsquirrel*.so" -print | head -n 1 | sed -e 's/^\(\/[^\/]*\/[^\/]*\)\/.*$/\1\//g')
       ln -s "${includeDirectory}" include
       ln -s "${libDirectory}" lib
       cd ..
       sed -i 's/Werror/Wno-error/g' anago/anago.mk
       libSquirrel=$(find /usr/lib* -name "libsquirrel*.so" -print | head -n 1 | sed -e 's/^\/\([^\/]*\/\)*lib\(.*\)\.so$/\2/g')
       libSqstdlib=$(find /usr/lib* -name "libsqstdlib*.so" -print | head -n 1 | sed -e 's/^\/\([^\/]*\/\)*lib\(.*\)\.so$/\2/g')
       sed -i 's/\-lsquirrel[^ \t\n\r]*/-l'"${libSquirrel}"'/g' anago/anago.mk
       sed -i 's/\-lsqstdlib[^ \t\n\r]*/-l'"${libSqstdlib}"'/g' anago/anago.mk
       sed -i 's/GetPropVal/GetAttribute/g' anago/nescartxml.cpp
       sed -i 's/sq_setprintfunc(v, print_stdout);/sq_setprintfunc(v, print_stdout, NULL);/g' anago/squirrel_wrap.c
       sed -i 's/sq_setprintfunc(v, print_other);/sq_setprintfunc(v, print_other, NULL);/g' anago/squirrel_wrap.c
       cd anago
       make -f Makefile.unix

- <u>***Execution :***</u>

<!-- -->

       su
       ./anago d nrom.ad dumpfile.nes

### Anago 0.6.0l legacy {#Anago_0.6.0l_legacy}

Commands to execute in a terminal :

- <u>***Dependencies installation :***</u>

<!-- -->

       sudo apt install libsquirrel-dev libncurses-dev libusb-dev cmake

- <u>***Anago 0.6.0l compilation / installation :***</u>

<!-- -->

       mkdir anago-legacy
       cd anago-legacy
       git clone https://github.com/zerkerX/anago.git
       cd anago
       mkdir build
       cd build
       export CFLAGS="$CFLAGS -fcommon"
       cmake ..
       make
       cp ../scripts/*.nut .

- <u>***Execution :***</u>

<!-- -->

       su
       ./anago d ../scripts/nrom.ad dumpfile.nes

## Usage {#Usage}

If the Anago executable file is installed in the ***build*** directory, the easy way is to copy the files ***dumpcore.nut***, ***flashcore.nut*** and ***flashdevice.nut*** in the same directory :

       cd build
       cp ../scripts/*.nut .

Global syntax (for Linux) :

     ./anago mode script_file target_file [...]

<u>***NB :***</u> For Windows, the syntax is "***anago.exe***" in replacement of "***./anago***"

### Dump Mode {#Dump_Mode}

Used to create iNES ROM images from carts.

Syntax :

     ./anago d?? script_file dump_file mapper

<table class="wikitable">
<thead>
<tr>
<th>Parameter</th>
<th></th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2"><em><strong>d??</strong></em></td>
<td>Dumping progress display with ncurses</td>
</tr>
<tr>
<td colspan="2"><em><strong>D??</strong></em></td>
<td>Dumping progress display is same as unagi.exe</td>
</tr>
<tr>
<td colspan="2"><em><strong>d??</strong></em> or <em><strong>D??</strong></em></td>
<td>If given, the <em><strong>??</strong></em> portion determines the length of the ROM to dump in multiples of the mapper's configured width for the program ROM portion and character ROM portion respectively. For example, UNROM boards use 128KB PRG ROMs while UOROM boards use 256KB ROMs. A dumper script is provided for UNROM. To dump a UOROM board, use the UNROM script with the d2 option.</td>
</tr>
<tr>
<td rowspan="2"><em><strong>?</strong></em></td>
<td><em><strong>2</strong></em></td>
<td>2 is applicable for unrom.ad, mmc3.ad and mmc5.ad (the first <em><strong>?</strong></em> is for CPU increase, the second <em><strong>?</strong></em> is for PPU increase)</td>
</tr>
<tr>
<td><em><strong>4</strong></em></td>
<td>4 is applicable for mmc5.ad (the first <em><strong>?</strong></em> is for CPU increase, the second <em><strong>?</strong></em> is for PPU increase)</td>
</tr>
<tr>
<td colspan="2"><em><strong>script_file</strong></em></td>
<td>Specifies <em><strong>.ad</strong></em>, <em><strong>.ae</strong></em> or <em><strong>.ag</strong></em> file for target cartridge</td>
</tr>
<tr>
<td colspan="2"><em><strong>dump_file</strong></em></td>
<td>Output file, typically ending in <em><strong>.nes</strong></em></td>
</tr>
<tr>
<td colspan="2"><em><strong>mapper</strong></em></td>
<td>If given, overrides the mapper number in the output iNES ROM image. This is useful for mapper ASICs that are assigned multiple iNES mapper numbers such as MMC3/6.</td>
</tr>
</tbody>
</table>

***<u>Example 1.1 :</u> Getting an image for UNROM***

     ./anago d unrom.ad unrom.nes

Anago gets vram mirroring connection from target cartridge.

***<u>Example 1.2</u> : Getting an image for UOROM***

     ./anago d2 unrom.ad unrom.nes

UOROM data size is twice that of UNROM. This must be specicifed by hand.

***<u>Example 2 :</u> Getting an image for Metal Slader Glory / ELROM***

     ./anago d22 mmc5.ad hal_4j.nes

In many cases, it is enough by 2M+2M to dump images for MMC5. MMC5 can manage 8M+8M. Please specify it individually according to capacity.

***<u>Example 3 :</u> Getting an image for Ys III / TKSROM***

     ./anago d mmc3.ad vfr_q8_12.nes 118

Enter digits on last commandline word to change mapper number for MMC3 variants.

### Flash Programming Mode {#Flash_Programming_Mode}

Write to Flash ROM chips on specially prepaired flash carts.

Syntax :

     ./anago f?? script_file programming_file prg_chip chr_chip
     ./anago F?? script_file programming_file prg_chip chr_chip
     ./anago a?? script_file programming_file prg_chip chr_chip

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>Parameter</th>
<th></th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2"><em><strong>f??</strong></em></td>
<td>Usual mode</td>
</tr>
<tr>
<td colspan="2"><em><strong>F??</strong></em></td>
<td>Validate written memory and abort if there is a discrepancey.</td>
</tr>
<tr>
<td colspan="2"><em><strong>a??</strong></em></td>
<td>Use <em><strong>DRIVER_DUMMY</strong></em></td>
</tr>
<tr>
<td colspan="2"><em><strong>f??</strong></em> or <em><strong>F??</strong></em></td>
<td>If given, the <em><strong>??</strong></em> portion controls the fill mode for the PRG and CHR Flash ROM chips respectively. If omitted, Fill Mode is assumed.
<p>Use <em><strong>DRIVER_KAZZO</strong></em></p></td>
</tr>
<tr>
<td rowspan="4"><em><strong>?</strong></em></td>
<td><em><strong>f</strong></em></td>
<td>Fill mode. The entire ROM chip is overwritten, duplicating the PRG or CHR ROM contents as many times as nessecary to fill the chip. This is the slowest option but may be needed under some curcumstances. If unsure, use this mode. In general, if your mapper has an unknonw power-on state or the top address lines of the ROM may float at runtime, always use this mode.</td>
</tr>
<tr>
<td><em><strong>t</strong></em></td>
<td>Only writes the ROM image once to the top-most portion of the chip. This mode is typically safe for CHR-ROM for any mapper, as the power-on state is immaterial with a proper initialization routine.</td>
</tr>
<tr>
<td><em><strong>b</strong></em></td>
<td>Only writes the ROM image once to the bottom-most portion of the chip. Use with caution, most mapper's power-on states are incompatible with this mode.</td>
</tr>
<tr>
<td><em><strong>e</strong></em></td>
<td>Write nothing for this memory device. This is equivalent to specifiying a dummy device.</td>
</tr>
<tr>
<td colspan="2"><em><strong>script_file</strong></em></td>
<td>Specifies <em><strong>.af</strong></em> or <em><strong>.ag</strong></em> file for target cartridge.</td>
</tr>
<tr>
<td colspan="2"><em><strong>programming_file</strong></em></td>
<td>Specifies the input iNES image file.</td>
</tr>
<tr>
<td colspan="2"><em><strong>prg_chip</strong></em><br />
<em><strong>chr_chip</strong></em></td>
<td>Specifies the Flash ROM chip present for each memory type in the target flash cart. Supported devices are listed in "<em><strong>flashdevice.nut</strong></em>" : <em><strong>dummy</strong></em> , <em><strong>W29C020</strong></em> , <em><strong>W29C040</strong></em> , <em><strong>W49F002</strong></em> , <em><strong>EN29F002T</strong></em> , <em><strong>AM29F040B</strong></em> , <em><strong>PM29F002T</strong></em> , <em><strong>MBM29F080A</strong></em>. "<em><strong>dummy</strong></em>" is a special device name that will transfer no data. Use this to avoid overwriting one of the memory chips.</td>
</tr>
</tbody>
</table>

***<u>Example 1.1 :</u> Tranfer 1M+1M image to mmc3 / TLROM***

     ./anago f mmc3.af tlrom_1M_1M.nes AM29F040B AM29F040B

In this case the board is configured with 4M flash ROMs. The MMC3 can only map 2M of CHR-ROM, so the high address line is assumed to be tied (either high or low). Anago would transfer the PRG portion four times to fill the PRG chip, and the CHR portion two times to fill the CHR chip.

***<u>Example 1.2 :</u> Tranfer 1M+1M image to mmc3 / TLROM***

     ./anago ftt mmc3.af tlrom_1M_1M.nes AM29F040B AM29F040B

In this example the PRG and CHR ROMs are only written once, this time to the top portion of each chip, in order to save time. The program code will need to be aware of this arrangement and only use the top-most bank numbers.

***<u>Example 1.3 :</u> Tranfer 1M+1M image to mmc3 / TLROM***

     ./anago fbt mmc3.af tlrom_1M_1M.nes AM29F040B AM29F040B

In this example the PRG ROM is written to the bottom 1M of the PRG ROM chip and the CHR ROM is written to the top 1M of the CHR ROM chip. The origional author remarks that this often works with Konami and Namcot titles.

Incidentally, mmc3 is not compatible with Namcot 109. The 109 board has hard-wired vram mirroring, mmc3 does not. Furthermore Namcot 106 is a fictitious device with well known extra sound. Don't miss 109 and 106.

- Note that the second translator does not fully understand the meaning of the above paragraph.

***<u>Example 2 :</u> Transfer UNROM(1M) image to UOROM***

     ./anago ft uorom.af unrom.nes AM29F040B

If charcter memory is RAM, charcter device can be skip. Again the program will need to only use the upper bank numbers.

***<u>Example 3 :</u> Transferring only charcter ROM image to mmc1/ SLROM***

     ./anago ftt mmc1_slrom.af skrom.nes dummy AM29F040B
     ./anago fet mmc1_slrom.af skrom.nes AM29F040B AM29F040B

Both commands result in the same behavior. Choose the one that suits your taste.

***<u>Notes :</u>***

- If the concepts of Top and Bottom transfer are not clear, and their impact and requirements on your software are not obvious, please use the full transfer mode at all times.

<!-- -->

- Anago does not have WRAM reading or writing mode at this time. Please use the Unagi utility for this purpose if needed.

## Anago scripts {#Anago_scripts}

Extensions for Anago scripts are :

|  |  |
|----|----|
| ***.ad*** | for dumping scripts (that contain functions : ***cpu_dump()*** and ***ppu_dump()*** (without ***cpu_ram_access()***)) |
| ***.ae*** | for dumping scripts (that contain function : ***cpu_ram_access()***) |
| ***.af*** | for flashing scripts (that contain functions : ***cpu_transfer()*** and ***ppu_transfer()***) |
| ***.ag*** | for both dumping and flashing scripts (that contain functions : ***cpu_dump()***, ***ppu_dump()***, ***cpu_transfer()*** and ***ppu_transfer()***) |
| ***.ai*** | for include scripts (that are called by other scripts with the ***dofile()*** function) |

### Dumping scripts {#Dumping_scripts}

An Anago dumping script is mainly composed of 3 parts :

- <u>Initialization variable "***board***" :</u>

<!-- -->

     board <- {
       mappernum = XXX,
       cpu_romsize = XXX, cpu_banksize = XXX,
       ppu_romsize = XXX, ppu_banksize = XXX,
       ppu_ramfind = false, vram_mirrorfind = false
     };

|  |  |
|----|----|
| ***mappernum*** | Mapper number |
| ***cpu_romsize*** | Size of the PRG to dump in bytes. It can be set in decimal, hexadecimal or with the constant "***mega***" (Example : cpu_romsize = 2 \* mega) |
| ***cpu_banksize*** | Size of the bank dumped at each iteration |
| ***ppu_romsize*** | Size of the CHR to dump in bytes. It can be set in decimal, hexadecimal or with the constant "***mega***" (Example : ppu_romsize = 2 \* mega) |
| ***ppu_banksize*** | Size of the bank dumped at each iteration |
| ***ppu_ramfind*** |  |
| ***vram_mirrorfind*** | Mirroring type (true = vertical mirroring / false = horizontal mirroring) |

- <u>***cpu_dump()*** function (to dump PRG part) :</u>

<!-- -->

     function cpu_dump(d, pagesize, banksize)

|  |  |
|----|----|
| ***pagesize*** | The number of iteration. The value of this parameter is the result of what has been set in the ***board*** variable : pagesize = cpu_romsize / cpu_banksize |
| ***banksize*** | The value of this parameter is the value set in the ***board*** variable for cpu_banksize |

- <u>***ppu_dump()*** function (to dump CHR part) :</u>

<!-- -->

     function ppu_dump(d, pagesize, banksize)

|  |  |
|----|----|
| ***pagesize*** | The number of iteration. The value of this parameter is the result of what has been set in the ***board*** variable : pagesize = ppu_romsize / ppu_banksize |
| ***banksize*** | The value of this parameter is the value set in the ***board*** variable for ppu_banksize |

In basic scripts, ***cpu_dump()*** and ***ppu_dump()*** functions contains a loop :

     for (local i = 0; i < pagesize; i++) {
       cpu_write(d, addressPart, dataPart);
       cpu_read(d, address, banksize);
     }

|  |  |
|----|----|
| ***addressPart*** | Formula depending of the base address for the dump and the iteration variable i (Example : 0x8000 \| i) based on the ***CPU_Ax*** description (x=0 to 15) |
| ***dataPart*** | Formula depending of the iteration variable i based on the ***CPU_Dx*** description (Example : i) (x=0 to 7) |
| ***address*** | Address where to read the bank (Example : 0x8000) |
| ***pagesize*** | Generally, the value of the parameter of the function but it can be different |
| ***banksize*** | Generally, the value of the parameter of the function but it can be different |

### Flashing scripts {#Flashing_scripts}

An Anago flashing script is mainly composed of 4 parts :

- <u>Initialization variable "***board***" :</u>

<!-- -->

     board <- {
       mappernum = XX, vram_mirrorfind = false,
       cpu = {banksize = XXXX, maxsize = XXXXXX},
       ppu = {banksize = XXXX, maxsize = XXXXXX}
     };

- <u>***program_initalize()*** function :</u>

<!-- -->

     function program_initalize(d, cpu_banksize, ppu_banksize)

- <u>***cpu_transfer()*** function (to flash PRG part) :</u>

<!-- -->

     function cpu_transfer(d, start, end, cpu_banksize)

- <u>***ppu_transfer()*** function (to flash CHR part) :</u>

<!-- -->

     function ppu_transfer(d, start, end, ppu_banksize)

## Common Errors {#Common_Errors}

### Linker error {#Linker_error}

     /usr/bin/ld : CMakeFiles/anago.dir/reader_dummy.c.o:(.data.rel.ro.local+0x0) : définitions multiples de « DRIVER_DUMMY »; CMakeFiles/anago.dir/anago.c.o:(.rodata+0x60) : défini pour la première fois ici
     /usr/bin/ld : CMakeFiles/anago.dir/reader_kazzo.c.o:(.data.rel.ro.local+0x0) : définitions multiples de « DRIVER_KAZZO »; CMakeFiles/anago.dir/anago.c.o:(.rodata+0x0) : défini pour la première fois ici
     collect2: error: ld returned 1 exit status
     make[2]: *** [CMakeFiles/anago.dir/build.make:298 : anago] Erreur 1
     make[1]: *** [CMakeFiles/Makefile2:96 : CMakeFiles/anago.dir/all] Erreur 2
     make: *** [Makefile:171 : all] Erreur 2

<u>**Solution :**</u> This error can be solved by adding the ***-fcommon*** option for the ld linker before compilation :

In the Anago root directory (***anago-master/***), run in a terminal :

       rm -rf build
       mkdir build
       cd build
       export CFLAGS="$CFLAGS -fcommon"
       cmake ..
       make
       cp ../scripts/*.nut .

### USB permission error {#USB_permission_error}

     Could not find USB device "kazzo" with vid=0x16c0 pid=0x5dc; code: 1
     reader open error

<u>**Solution :**</u> This error indicates that the user doesn't have the permission for accessing to USB device.

A simple workaround is to run Anago in ***root*** mode in a terminal, though this causes dumped ROM files to have root ownership/permissions :

       su
       ./anago d ../scripts/nrom.ad dumpfile.nes

or :

       sudo ./anago d ../scripts/nrom.ad dumpfile.nes

The proper solution is to create a [udev](https://wiki.archlinux.org/title/Udev) rule to add the Kazzo device to a group with appropriate USB/serial permissions, such as *uucp* on Arch.

### Core error {#Core_error}

     dump core script error

<u>**Solution :**</u> The files ***dumpcore.nut*** , ***flashcore.nut*** and ***flashdevice.nut*** must be into the same directory than the Anago executable or into one of the directories :

- ***.***
- ***\${HOME}***/.config/anago/
- ***CMAKE_INSTALL_PREFIX***/***SHARE_DIR***

Example : In the Anago build directory (***anago-master/build/***), run in a terminal :

       cp ../scripts/*.nut .
       ./anago d ../scripts/nrom.ad dumpfile.nes

### Script Syntax error {#Script_Syntax_error}

     AN ERROR HAS OCCURRED [the index 'cpu_romsize' does not exist]

     CALLSTACK
     *FUNCTION [dump()] dumpcore.nut line [8]

     LOCALS
     [vram] 1
     [increase_ppu] 1
     [increase_cpu] 1
     [mappernum] XXX
     [d] USERPOINTER
     [this] TABLE

<u>**Solution :**</u> Rewrite the script with the syntax corresponding to the Anago version program.

For example, if the script syntax was written for the Anago mainline version :

     board <- {
         mappernum = XXX,
         cpu_rom = {
           size_base = XXXXXX, size_max = XXXXXX, banksize = 0x4000
         },
         ppu_rom = {
           size_base = YYYYYY, size_max = YYYYYY, banksize = 0x2000
         },
         ppu_ramfind = false, vram_mirrorfind = true
       };

The valid syntax for the Anago 0.6.0l legacy version is :

     board <- {
         mappernum = XXX,
         cpu_romsize = XXXXXX, cpu_banksize = 0x4000,
         ppu_romsize = YYYYYY, ppu_banksize = 0x2000,
         ppu_ramfind = false, vram_mirrorfind = true
       };

## Related Links {#Related_Links}

- [INL's announcement of starting selling Kazzo dumpers](https://forums.nesdev.org/viewtopic.php?t=7912)

<!-- -->

- [INL's Website for Kazzo dumper](https://infiniteneslives.com/kazzo.php)

<!-- -->

- [Original Unagi project](https://web.archive.org/web/20180408011450/https://ja.osdn.net/projects/unagi/wiki/kazzo_intro_en)

<!-- -->

- [Official software for Kazzo](https://www.dropbox.com/s/f15d778hoacz0sp/kazzo.zip)

|                  |                                  |
|------------------|----------------------------------|
| md5sum kazzo.zip | b4b924d0432fdaa0e13271136d77ef05 |

Anago legacy release (Linux portage of Anago) :

- <https://github.com/zerkerX/anago>
- <https://github.com/hernan43/anago-linux>

| Releases | MD5SUM | Date | Size |
|----|----|----|----|
| [0.6.0l](https://github.com/zerkerX/anago/archive/refs/heads/master.zip) | 0df72848e948fdebe0df4cfcdd82ecad | Mar 30, 2019 |  66462 bytes |

Anago mainline release :

- <https://github.com/sharkpp/unagi_kazzo>
- <https://github.com/sharkpp/unagi_client>

Anago scripts (Github repository for getting and adding Anago scripts) :

- <https://github.com/arantius/anago-scripts>
