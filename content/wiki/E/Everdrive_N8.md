---
title: "Everdrive N8"
url: "/wiki/Everdrive_N8/"
description: "The Everdrive N8 is a Flash Cartridge made by Krikzz. It uses an FPGA to emulate a wide variety of mappers, allowing the user to store a large collection of ROMs on a single SD card and run them on…"
wikiPage: "Everdrive_N8"
wikiCategories:
  - name: "Flash Cartridge"
    url: "/category/Flash_Cartridge/"
  - name: "Pages with broken file links"
    url: "/category/Pages_with_broken_file_links/"
---

The **Everdrive N8** is a [Flash Cartridge](../../category/Flash_Cartridge/) made by Krikzz. It uses an FPGA to emulate a wide variety of mappers, allowing the user to store a large collection of ROMs on a single SD card and run them on an NES or Famicom.

In addition to NES ROMs, the Everdrive N8 is able to play FDS disk images.

Famicom expansion audio is supported, and on the NES version is output on the EXP 6 expansion pin on the cartridge connector as used by the [PowerPak](../PowerPak/). A simple modification to the NES allows the expansion audio to be mixed with its output.

Specifications:

- PRG size: 512 KB
- CHR size: 512 KB

NES product: <https://krikzz.com/store/home/31-everdrive-n8-nes.html>

Famicom product: <https://krikzz.com/store/home/32-everdrive-n8-famicom.html>

This flashcart was eventually discontinued and succeeded by the much newer [Everdrive N8 Pro](../Everdrive_N8_Pro/).

## Hardware {#Hardware}

[<img src="../../wiki-images/Everdriven8-top.jpg" width="300" height="186" />](../../wiki-images/Everdriven8-top.jpg)

[](../../wiki-images/Everdriven8-top.jpg)

[PCB - top side](../../wiki-images/Everdriven8-top.jpg)

[](../../wiki-images/Everdriven8-top.jpg)

[](../../wiki-images/Everdriven8-top.jpg)

[](../../wiki-images/Everdriven8-top.jpg)[<img src="../../wiki-images/Everdriven8-bottom.jpg" width="300" height="186" />](../../wiki-images/Everdriven8-bottom.jpg)

[](../../wiki-images/Everdriven8-bottom.jpg)

[PCB - bottom side](../../wiki-images/Everdriven8-bottom.jpg)

[](../../wiki-images/Everdriven8-bottom.jpg)

[](../../wiki-images/Everdriven8-bottom.jpg)

[](../../wiki-images/Everdriven8-bottom.jpg)[<img src="../../wiki-images/Everdriven8-sch.png" width="300" height="148" />](../../wiki-images/Everdriven8-sch.png)

[](../../wiki-images/Everdriven8-sch.png)

[schematics](../../wiki-images/Everdriven8-sch.png)

[](../../wiki-images/Everdriven8-sch.png)

~~File:Everdriven8-shell.jpg~~

cartridge shell

[](../../wiki-images/Everdriven8-sch.png)

[](../../wiki-images/Everdriven8-sch.png)[<img src="../../wiki-images/Everdriven8-menu.jpg" width="300" height="169" />](../../wiki-images/Everdriven8-menu.jpg)

[](../../wiki-images/Everdriven8-menu.jpg)

[cartridge menu](../../wiki-images/Everdriven8-menu.jpg)

[](../../wiki-images/Everdriven8-menu.jpg)

Cartridge consists of:

- Altera Cyclone II FPGA (EP2C5T144), which is reprogrammed with mapper of preselected ROM
- Altera Max II CPLD (EPM240T100C5), which is glue logic
- 2 x CY7C1049CV33T (PRG-ROM, CHR-RAM/ROM)
- 29W160 (2MB) Flash (BIOS)
- IS52LV1024 (128kB PRG-RAM)
- SN74LVCR162245A 3.3V-5V buffers
- an optional FT245RL, that is responsible for communication between USB port and the CPLD

The software part consist of:

- CPLD - logic code inside EPM240T100C5 that is pre-programmed at factory and cannot be updated
- BIOS - program code inside 29W160 that is executed just after powering up the console (by default, it cannot be updated, unless a jumper at back side of the cartridge is closed)
- OS - program code that is stored on the flash card

## Mapper compatibility {#Mapper_compatibility}

[As of the OS update v1.26 in 2021, the Everdrive N8 supports the following](../../wiki-images/Everdriven8-menu.jpg) [mappers](../Mapper/): [\[1\]](https://krikzz.com/pub/support/everdrive-n8/original-series/OS/changelist.txt)

|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| [000](../NROM/) | [001](../MMC1/) | [002](../UxROM/) | [003](../CNROM/) | [004](../MMC3/) | [005](../MMC5/) | [006](../INES_Mapper_006/) | [007](../AxROM/) | [008](../INES_Mapper_006/) | [009](../MMC2/) | [010](../MMC4/) | [011](../Color_Dreams/) | [012](../INES_Mapper_012/) | [013](../CPROM/) | [014](../INES_Mapper_014/) | [015](../INES_Mapper_015/) |
| [016](../INES_Mapper_016/) | [017](../INES_Mapper_017/) | [018](../INES_Mapper_018/) | [019](../INES_Mapper_019/) | [020](../INES_Mapper_020/) | [021](../VRC2_and_VRC4/) | [022](../VRC2_and_VRC4/) | [023](../VRC2_and_VRC4/) | [024](../VRC6/) | [025](../VRC2_and_VRC4/) | [026](../VRC6/) | [027](../INES_Mapper_027/) | [028](../Action_53_mapper/) | [029](../INES_Mapper_029/) | [030](../UNROM_512/) | [031](../INES_Mapper_031/) |
| [032](../INES_Mapper_032/) | [033](../INES_Mapper_033/) | [034](../INES_Mapper_034/) | [035](../J.Y._Company_ASIC/) | [036](../INES_Mapper_036/) | [037](../INES_Mapper_037/) | [038](../INES_Mapper_038/) | [039](../INES_Mapper_039/) | [040](../INES_Mapper_040/) | [041](../INES_Mapper_041/) | [042](../INES_Mapper_042/) | [043](../INES_Mapper_043/) | [044](../INES_Mapper_044/) | [045](../INES_Mapper_045/) | [046](../INES_Mapper_046/) | [047](../INES_Mapper_047/) |
| [048](../INES_Mapper_048/) | [049](../INES_Mapper_049/) | [050](../INES_Mapper_050/) | [051](../INES_Mapper_051/) | [052](../INES_Mapper_052/) | [053](../INES_Mapper_053/) | [054](../INES_Mapper_054/) | [055](../INES_Mapper_055/) | [056](../INES_Mapper_056/) | [057](../INES_Mapper_057/) | [058](../INES_Mapper_058/) | [059](../INES_Mapper_059/) | [060](../INES_Mapper_060/) | [061](../INES_Mapper_061/) | [062](../INES_Mapper_062/) | [063](../INES_Mapper_063/) |
| [064](../RAMBO-1/) | [065](../INES_Mapper_065/) | [066](../GxROM/) | [067](../INES_Mapper_067/) | [068](../INES_Mapper_068/) | [069](../Sunsoft_FME-7/) | [070](../INES_Mapper_070/) | [071](../INES_Mapper_071/) | [072](../INES_Mapper_072/) | [073](../VRC3/) | [074](../INES_Mapper_074/) | [075](../VRC1/) | [076](../INES_Mapper_076/) | [077](../INES_Mapper_077/) | [078](../INES_Mapper_078/) | [079](../NINA-003-006/) |
| [080](../Taito_X1-005/) | [081](../INES_Mapper_081/) | [082](../Taito_X1-017/) | [083](../INES_Mapper_083/) | [084](../INES_Mapper_084/) | [085](../VRC7/) | [086](../INES_Mapper_086/) | [087](../INES_Mapper_087/) | [088](../INES_Mapper_088/) | [089](../INES_Mapper_089/) | [090](../J.Y._Company_ASIC/) | [091](../INES_Mapper_091/) | [092](../INES_Mapper_092/) | [093](../INES_Mapper_093/) | [094](../INES_Mapper_094/) | [095](../INES_Mapper_095/) |
| [096](../INES_Mapper_096/) | [097](../INES_Mapper_097/) | [098](../INES_Mapper_098/) | [099](../INES_Mapper_099/) | [100](../INES_Mapper_100/) | [101](../INES_Mapper_101/) | [102](../INES_Mapper_102/) | [103](../INES_Mapper_103/) | [104](../PEGASUS_5_IN_1/) | [105](../NES-EVENT/) | [106](../INES_Mapper_106/) | [107](../INES_Mapper_107/) | [108](../INES_Mapper_108/) | [109](../INES_Mapper_109/) | [110](../INES_Mapper_110/) | [111](../INES_Mapper_111/) |
| [112](../INES_Mapper_112/) | [113](../INES_Mapper_113/) | [114](../INES_Mapper_114/) | [115](../INES_Mapper_115/) | [116](../INES_Mapper_116/) | [117](../INES_Mapper_117/) | [118](../INES_Mapper_118/) | [119](../INES_Mapper_119/) | [120](../INES_Mapper_120/) | [121](../INES_Mapper_121/) | [122](../INES_Mapper_122/) | [123](../INES_Mapper_123/) | [124](../INES_Mapper_124/) | [125](../INES_Mapper_125/) | [126](../NES_2.0_Mapper_534/) | [127](../INES_Mapper_127/) |
| [128](../NES_2.0_Mapper_265/) | [129](../INES_Mapper_129/) | [130](../INES_Mapper_130/) | [131](../INES_Mapper_131/) | [132](../INES_Mapper_132/) | [133](../INES_Mapper_133/) | [134](../INES_Mapper_134/) | [135](../Sachen_8259/) | [136](../INES_Mapper_136/) | [137](../INES_Mapper_137/) | [138](../Sachen_8259/) | [139](../Sachen_8259/) | [140](../INES_Mapper_140/) | [141](../Sachen_8259/) | [142](../INES_Mapper_142/) | [143](../INES_Mapper_143/) |
| [144](../INES_Mapper_144/) | [145](../INES_Mapper_145/) | [146](../NINA-003-006/) | [147](../INES_Mapper_147/) | [148](../INES_Mapper_148/) | [149](../INES_Mapper_149/) | [150](../INES_Mapper_150/) | [151](../INES_Mapper_151/) | [152](../INES_Mapper_152/) | [153](../INES_Mapper_153/) | [154](../INES_Mapper_154/) | [155](../MMC1/) | [156](../INES_Mapper_156/) | [157](../INES_Mapper_157/) | [158](../INES_Mapper_158/) | [159](../INES_Mapper_159/) |
| [160](../INES_Mapper_160/) | [161](../INES_Mapper_161/) | [162](../INES_Mapper_162/) | [163](../INES_Mapper_163/) | [164](../INES_Mapper_164/) | [165](../INES_Mapper_165/) | [166](../INES_Mapper_167/) | [167](../INES_Mapper_167/) | [168](../INES_Mapper_168/) | [169](../INES_Mapper_169/) | [170](../INES_Mapper_170/) | [171](../INES_Mapper_171/) | [172](../INES_Mapper_172/) | [173](../INES_Mapper_173/) | [174](../INES_Mapper_174/) | [175](../INES_Mapper_175/) |
| [176](../INES_Mapper_176/) | [177](../INES_Mapper_177/) | [178](../INES_Mapper_178/) | [179](../INES_Mapper_179/) | [180](../INES_Mapper_180/) | [181](../INES_Mapper_181/) | [182](../INES_Mapper_114/) | [183](../INES_Mapper_183/) | [184](../INES_Mapper_184/) | [185](../CNROM/) | [186](../INES_Mapper_186/) | [187](../INES_Mapper_187/) | [188](../INES_Mapper_188/) | [189](../INES_Mapper_189/) | [190](../INES_Mapper_190/) | [191](../INES_Mapper_191/) |
| [192](../INES_Mapper_192/) | [193](../INES_Mapper_193/) | [194](../INES_Mapper_194/) | [195](../INES_Mapper_195/) | [196](../INES_Mapper_196/) | [197](../INES_Mapper_197/) | [198](../INES_Mapper_198/) | [199](../INES_Mapper_199/) | [200](../INES_Mapper_200/) | [201](../INES_Mapper_201/) | [202](../INES_Mapper_202/) | [203](../INES_Mapper_203/) | [204](../INES_Mapper_204/) | [205](../INES_Mapper_205/) | [206](../INES_Mapper_206/) | [207](../INES_Mapper_207/) |
| [208](../INES_Mapper_208/) | [209](../J.Y._Company_ASIC/) | [210](../INES_Mapper_210/) | [211](../J.Y._Company_ASIC/) | [212](../INES_Mapper_212/) | [213](../INES_Mapper_213/) | [214](../INES_Mapper_214/) | [215](../INES_Mapper_215/) | [216](../INES_Mapper_216/) | [217](../INES_Mapper_217/) | [218](../INES_Mapper_218/) | [219](../INES_Mapper_219/) | [220](../INES_Mapper_220/) | [221](../INES_Mapper_221/) | [222](../INES_Mapper_222/) | [223](../INES_Mapper_223/) |
| [224](../INES_Mapper_224/) | [225](../INES_Mapper_225/) | [226](../INES_Mapper_226/) | [227](../INES_Mapper_227/) | [228](../INES_Mapper_228/) | [229](../INES_Mapper_229/) | [230](../INES_Mapper_230/) | [231](../INES_Mapper_231/) | [232](../INES_Mapper_232/) | [233](../INES_Mapper_233/) | [234](../INES_Mapper_234/) | [235](../INES_Mapper_235/) | [236](../INES_Mapper_236/) | [237](../INES_Mapper_237/) | [238](../INES_Mapper_238/) | ~~239~~ |
| [240](../INES_Mapper_240/) | [241](../INES_Mapper_241/) | [242](../INES_Mapper_242/) | [243](../INES_Mapper_243/) | [244](../INES_Mapper_244/) | [245](../INES_Mapper_245/) | [246](../INES_Mapper_246/) | ~~247~~ | [248](../INES_Mapper_248/) | [249](../T9552/) | [250](../INES_Mapper_250/) | [251](../INES_Mapper_251/) | [252](../INES_Mapper_252/) | [253](../INES_Mapper_253/) | [254](../INES_Mapper_254/) | [255](../INES_Mapper_255/) |

Known problems:

- Mapper 71 only supports the memory controller used by FireHawk, all other Mapper 71 games must be reassigned to Mapper 2 to work correctly.

## Obsolete Mappers {#Obsolete_Mappers}

A few mappers have been created by others to supplement the Everdrive's provided set, but have since been integrated into or replaced by official mappers:

- [UNROM 512](../UNROM_512/) ([030](../UNROM_512/)): [forum post](https://forums.nesdev.org/viewtopic.php?p=235457#p235457) (lacks flash save)
- [GTROM](../GTROM/) ([iNES Mapper 111](../INES_Mapper_111/)): [forum post](https://forums.nesdev.org/viewtopic.php?f=9&t=18841) (lacks flash save)

## Software development limitations {#Software_development_limitations}

Aside from mapper incompatibility, there are minor differences between running NES programs on the Everdrive versus a traditional single-game cartridge.

- The Everdrive does not accurately simulate power-on state. Because power-on always boots the Everdrive menu, RAM and various registers will be initialized to a consistent state before any NES ROM is chosen to run. (Reset state, however, is not affected by this problem.)
- [Open bus behavior](../Open_bus_behavior/) may be different in several memory regions that are used by the Everdrive, but would not be connected on a regular cartridge. ([forum post](http://forums.nesdev.org/viewtopic.php?f=2&t=12549))
- The Everdrive is incompatible with an NES that has the [CopyNES](../CopyNES/) modification installed, due to a bus conflict with its boot code.

## Everdrive development resources {#Everdrive_development_resources}

- Software tools and example mapper source code: <http://krikzz.com/pub/support/everdrive-n8/original-series/development/>
