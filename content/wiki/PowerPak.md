---
title: "PowerPak"
url: "/wiki/PowerPak/"
description: "The PowerPak is a Flash Cartridge made by RetroUSB. It uses an FPGA to emulate a wide variety of mappers, allowing the user to store a large collection of ROMs on a single Compact Flash card and run…"
wikiPage: "PowerPak"
wikiCategories:
  - name: "Flash Cartridge"
    url: "/category/Flash_Cartridge/"
---

The **PowerPak** is a [Flash Cartridge](../../category/Flash_Cartridge/) made by ~~RetroUSB~~. It uses an ~~FPGA~~ to emulate a wide variety of [mappers](../Mapper/), allowing the user to store a large collection of ROMs on a single Compact Flash card and run them on an NES. It is widely used by homebrew NES developers to test their software. It's also compatible with the [CopyNES](../CopyNES/).

In addition to NES ROMs, the PowerPak is able to play [FDS](../Family_Computer_Disk_System/) disk images, as well as [NSF](../NSF/) music files.

Famicom expansion audio is supported, and output on the EXP 6 expansion pin on the [cartridge connector](../Cartridge_connector/). A simple modification to the NES allows the expansion audio to be mixed with its output.

Specifications:

- PRG size: 512 KB (252 KB for NSF)
- CHR size: 512 KB
- Auxiliary PRG-RAM size: 32 KB

Product page: <http://www.retrousb.com/product_info.php?products_id=34>

## Mapper Compatibility {#Mapper_Compatibility}

The PowerPak mappers have undergone several revisions, gradually improving compatibility. After official development ceased in 2010, Loopy and TheFox have each created a supplemental set of PowerPak mappers to improve its capabilities.

The commonly recommended setup is:

- Begin with the Official mapper set.
- Add Loopy's mapper set on top, replacing files.
- Add the PowerMapper set on top if you want savestate support (see its readme).
- Add any of the additional single mappers if needed.
- Replace the N.MAP loader with a [patched version](https://forums.nesdev.org/viewtopic.php?p=283943#p283943) so newer ROMs with [header byte 15 set](../NES_2.0/#Default_Expansion_Device) can be loaded.

### PowerMappers {#PowerMappers}

TheFox created a set of revised PowerPak mappers to supplement or augment the existing ones, most notably adding a savestate feature for the mappers it contains, but removing the Game Genie feature.

Download: <http://fo.aspekt.fi>

- [NROM](../NROM/) ([0](../NROM/))
- [MMC1](../MMC1/) ([1](../MMC1/))
- [UxROM](../UxROM/) ([2](../UxROM/))
- [CxROM](../CxROM/) ([3](../CNROM/))
- [MMC3](../MMC3/) ([4](../MMC3/))
- [AxROM](../AxROM/) ([7](../AxROM/))
- [MMC2](../MMC2/) ([9](../MMC2/))
- [MMC4](../MMC4/) ([10](../MMC4/))
- [ColorDreams](../Color_Dreams/) ([11](../Color_Dreams/))
- [BxROM](../INES_Mapper_034/) ([34](../INES_Mapper_034/))
- [GxROM](../GxROM/) ([66](../GxROM/))
- [FME-7](../Sunsoft_FME-7/) ([69](../Sunsoft_FME-7/), no expansion audio)
- [Codemasters](../INES_Mapper_071/) ([71](../INES_Mapper_071/))
- MMC3/[TxSROM](../INES_Mapper_118/) ([118](../INES_Mapper_118/))
- MMC3/[TQROM](../INES_Mapper_119/) ([119](../INES_Mapper_119/))

Known problems:

- Mapper 4 IRQ is delayed by 2 cycles like [RAMBO-1](../RAMBO-1/) to work around noise issues with PPU A12. This does not significantly affect most games.
- Mapper 69 does not support the [Sunsoft 5B expansion audio](../Sunsoft_5B_audio/) used in Gimmick!

### <span id="Loopy.27s_Mappers"></span>Loopy's Mappers {#Loopy.27s_Mappers}

Loopy released a set of revised PowerPak mappers in 2011, adding fixes and additional support for several mappers:

Download: <http://3dscapture.com/NES/powerpak_loopy.zip>

- [CNROM](../CNROM/) ([3](../CNROM/))
- [MMC3](../MMC3/) ([4](../MMC3/))
- [MMC5](../MMC5/) ([5](../MMC5/), no expansion audio)
- [N163](../INES_Mapper_019/) ([19](../INES_Mapper_019/))
- [VRC4](../VRC2_and_VRC4/) ([21](../VRC2_and_VRC4/), [23](../VRC2_and_VRC4/), [25](../VRC2_and_VRC4/))
- [VRC6](../VRC6/) ([24](../VRC6/), [26](../VRC6/))
- [BxROM](../INES_Mapper_034/) ([34](../INES_Mapper_034/))
- [Sunsoft-5B](../Sunsoft_FME-7/) ([69](../Sunsoft_FME-7/))
- [Codemasters](../INES_Mapper_071/) ([71](../INES_Mapper_071/))
- [JY Company](../J.Y._Company_ASIC/) ([90](../J.Y._Company_ASIC/), partial)
- [FDS](../Family_Computer_Disk_System/)
- [NSF](../NSF/) (no MMC5 audio, no VRC7 audio)

Known problems:

- Mapper 4 IRQ has reliability issues due to PPU A12 noise, causing status bars etc. to jitter up and down on some systems. (A fixed version is available: [forum thread](https://forums.nesdev.org/viewtopic.php?t=24229).)

Notes:

- Mapper 3 now supports unlicensed oversize variants (e.g. used by Panesian games).
- Mapper 4 now supports Startropics.
- Mapper 5 does not support the [MMC5 expansion audio](../MMC5_audio/).

### <span id="Myask.27s_Mappers_.28WIP.29"></span>Myask's Mappers (WIP) {#Myask.27s_Mappers_.28WIP.29}

Myask made a few mappers not yet covered by others.

Download: [here](http://forums.nesdev.org/download/file.php?id=5989). BBS Thread: [here](http://forums.nesdev.org/viewtopic.php?f=9&t=14373).

- Irem G-101 ([032](../INES_Mapper_032/), but *Major League* hack/submapper not implemented)
- Irem H3001 ([065](../INES_Mapper_065/))
- Taito TC0190 ([033](../INES_Mapper_033/))
  - Note that some mapper 48 roms have been mislabeled as mapper 33.
- Taito TC0690, TC0190+PAL16R4 ([048](../INES_Mapper_048/), very buggy)
- Magicseries Corp ([107](../INES_Mapper_107/))

### Miscellaneous {#Miscellaneous}

- Action 53 mapper ([028](../Action_53_mapper/)): [forum thread](https://forums.nesdev.org/viewtopic.php?p=102718#p102718)
- [GTROM](../GTROM/) ([111](../INES_Mapper_111/)): [forum thread](https://forums.nesdev.org/viewtopic.php?f=9&t=16631)
- Magic Kid GooGoo ([190](../INES_Mapper_190/)): [forum thread](https://forums.nesdev.org/viewtopic.php?f=9&t=15682)
- [UNROM 512](../UNROM_512/) ([030](../UNROM_512/)): [forum post](https://forums.nesdev.org/viewtopic.php?p=236715#p236715)
- ~~UXROM 512~~ Support up to 512k PRG: [forum post](https://forums.nesdev.org/viewtopic.php?t=24230)
- [SXROM](../MMC1/) ([001](../MMC1/)): [packaged with PR8](https://neilbaldwin.github.io/nes-audio/), but created by bunnyboy
- NES 2.0 [Default Expansion Device](../NES_2.0/#Default_Expansion_Device) header incompatibility fix: [forum thread](https://forums.nesdev.org/viewtopic.php?p=283943#p283943)
- Nova's Rad PowerPak Menu: [forum thread](https://forums.nesdev.org/viewtopic.php?t=12091) (alternative PowerPak menu)
- Rainwarriror's MMC3 IRQ timing accuracy fixes: [forum thread](https://forums.nesdev.org/viewtopic.php?t=24229)

### Offical Mappers V1.35b {#Offical_Mappers_V1.35b}

The last official release of mappers was in 2010. It supports a wide variety of popular mappers.

Download: [here](https://www.retrousb.com/product_info.php?products_id=34)

Supported mappers:

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

- Mapper 2 limited to 256k PRG. (An oversized version exists for 512k PRG: [forum thread](https://forums.nesdev.org/viewtopic.php?t=24230).)
- Mapper 3 limited to CNROM support, excluding unlicensed oversize variants (e.g. used by Panesian games).
- Mapper 4 does not support Startropics. (See [mapper 4 and MMC6](../MMC3/#iNES_Mapper_004_and_MMC6).)
- Mapper 5 (MMC5) is incomplete, and fails to run most MMC5 games.
- Mapper 23 (VRC2/4 variants) is listed as buggy.
- Mapper 92 (Jaleco-JF variant) is listed as buggy.
- Mapper 95 (Namcot-3425) is listed as buggy.
- Mapper 96 (Oeka Kids) is listed as buggy.

## Software development limitations {#Software_development_limitations}

Aside from mapper incompatibility, there are minor differences between running NES programs on the PowerPak versus a traditional single-game cartridge.

- The PowerPak does not accurately simulate power-on state. Because power-on always boots the PowerPak menu, RAM and various registers will be initialized to a consistent state before any NES ROM is chosen to run. (Reset state, however, is not affected by this problem.)
- [Open bus behavior](../Open_bus_behavior/) may be different in several memory regions that are used by the PowerPak, but would not be connected on a regular cartridge. ([forum post](http://forums.nesdev.org/viewtopic.php?f=2&t=12549))

## Utilities {#Utilities}

- [make_sram](../Make_sram/) : a program written in Python to create PowerPak save files for all NES ROMs in a folder or on a CF card

## PowerPak development resources {#PowerPak_development_resources}

- [PowerPak Menu](../PowerPak_Menu/) - information about the organization of the PowerPak's operating system.
- Collection of information and photos, source code for menu and loader: <http://nespowerpak.com/>
- Source code for CNROM example (Xilinx ISE Webpack 9.1): [powerpakdev1.zip](http://www.nespowerpak.com/powerpakdev1.zip)
- Source code for some of Loopy's mappers (Verilog): [powerpak_loopy_src.zip](http://forums.nesdev.org/viewtopic.php?p=173302#p173302) / [powerpak_loopy.zip](http://3dscapture.com/NES/powerpak_loopy.zip)
- Source code for NSF player: [powerpak_nsf_src.zip](http://3dscapture.com/NES/powerpak_nsf_src.zip)
- Source code for Mapper 190 (Magic Kid GooGoo): [forum thread](https://forums.nesdev.org/viewtopic.php?p=191417)
- Source code for Mapper 111 (GTROM): [forum thread](https://forums.nesdev.org/viewtopic.php?p=206883)
