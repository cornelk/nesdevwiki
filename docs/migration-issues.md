# Migration issues

Known problems of the content imported from the `202610-NesDevWiki.zip` backup of the old NESdev wiki.
The backup is a static XHTML snapshot, so the migration could only carry over what the snapshot contained.

Regenerate the link numbers with `make check`.

## Summary

| Area | Count | Action |
| --- | ---: | --- |
| Broken `#anchor` links (anchor never existed in the backup) | 211 refs, 65 anchors, 30 target pages | Fix the link or the heading, see below |
| Links to pages that are not in the backup | 2 | Removed from the text, see below |
| Links to dropped namespaces (File, Template, User, Talk, ...) | 136 | Turned into plain text |
| Redirect pages whose target is not migrated | 14 | Dropped, see below |
| Template syntax left in the text | 1 page | Remove by hand, see below |

Checked and not an issue: every other internal link, image and `#anchor` of the 1,109 generated pages resolves.
No anchor that existed in the backup was lost: all 211 broken anchors are absent from the original XHTML as well.

## Broken anchor links

These links point to a `#section` that does not exist on the target page. They were already broken on the old
wiki, for example when a section was renamed or removed after the link was written. The link still opens the
right page, it only does not scroll to a section. Anchors are shown decoded; the generated ids escape
characters like `(` as `.28`.

| Target page | Missing anchor | Linked from |
| --- | --- | --- |
| [APU](/wiki/APU/) | `#APU_master_control_and_status_($4015,_read/write)` | [APU_Frame_Counter](/wiki/APU_Frame_Counter/) |
| [APU](/wiki/APU/) | `#Legnth_Counter` | [APU](/wiki/APU/) |
| [APU](/wiki/APU/) | `#Register_overview#APU_master_control_and_status_($4015,_read/write)` | [APU_DMC/Register_table](/wiki/APU_DMC/Register_table/) |
| [APU](/wiki/APU/) | `#Status_($4015)` | [APU](/wiki/APU/), [CPU_ALL](/wiki/CPU_ALL/), [CPU_power_up_state](/wiki/CPU_power_up_state/), [DMA](/wiki/DMA/), [VT02+_Registers](/wiki/VT02+_Registers/) |
| [APU](/wiki/APU/) | `#Sweep` | [APU](/wiki/APU/) |
| [APU_DMC](/wiki/APU_DMC/) | `#Likely_internal_implementation_of_the_read` | [CHR_ROM_vs._CHR_RAM](/wiki/CHR_ROM_vs._CHR_RAM/) |
| [APU_DMC](/wiki/APU_DMC/) | `#Overview` | [DMA](/wiki/DMA/) |
| [APU_DMC/Register_table](/wiki/APU_DMC/Register_table/) | `#Sample_playback_unit_operation` | [APU_DMC/Register_table](/wiki/APU_DMC/Register_table/) |
| [APU_Pulse](/wiki/APU_Pulse/) | `#Muting` | [APU](/wiki/APU/) |
| [Arkanoid_controller](/wiki/Arkanoid_controller/) | `#Conversions` | [Arkanoid_controller](/wiki/Arkanoid_controller/) |
| [CNROM](/wiki/CNROM/) | `#Notes` | [M50805_pinout](/wiki/M50805_pinout/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_indexed_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_indexed_rmw` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_indexed_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_rmw` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Absolute_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Accumulator` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#BRK` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Immediate` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Implied` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Indirect` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Indirect_x_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Indirect_x_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Indirect_y_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Indirect_y_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Relative` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_indexed_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_indexed_rmw` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_indexed_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_read` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_rmw` | [Instruction_reference](/wiki/Instruction_reference/) |
| [CPU_addressing_modes](/wiki/CPU_addressing_modes/) | `#Zero_page_write` | [Instruction_reference](/wiki/Instruction_reference/) |
| [Cartridge_connector](/wiki/Cartridge_connector/) | `#Signal_Descriptions` | [Mask_ROM_pinout](/wiki/Mask_ROM_pinout/) |
| [Cycle_reference_chart](/wiki/Cycle_reference_chart/) | `#Clock_Rates` | [VRC7_audio](/wiki/VRC7_audio/) |
| [Emulator_tests](/wiki/Emulator_tests/) | `#CPU` | [CPU_interrupts](/wiki/CPU_interrupts/) |
| [Errata](/wiki/Errata/) | `#OAM_and_Sprites` | [PPU_programmer_reference](/wiki/PPU_programmer_reference/), [PPU_registers](/wiki/PPU_registers/), [PPU_sprite_evaluation](/wiki/PPU_sprite_evaluation/), [Palette_change_mid_frame](/wiki/Palette_change_mid_frame/) |
| [FDS_disk_format](/wiki/FDS_disk_format/) | `#Manufacturing_date` | [FDS_disk_format](/wiki/FDS_disk_format/) |
| [INES_Mapper_019](/wiki/INES_Mapper_019/) | `#CHR_and_NT_Select_($8000-$DFFF)_w` | [Game_bugs](/wiki/Game_bugs/) |
| [INES_Mapper_068](/wiki/INES_Mapper_068/) | `#$F000-$FFFF:_PRG_bank_at_$8000` | [INES_Mapper_188](/wiki/INES_Mapper_188/) |
| [INES_Mapper_176](/wiki/INES_Mapper_176/) | `#MMC3-compatible_registers_($8000/$8001,_$A000,_$C000/$C001,_$E000/$E001)` | [NES_2.0_Mapper_523](/wiki/NES_2.0_Mapper_523/) |
| [INES_Mapper_176](/wiki/INES_Mapper_176/) | `#RAM_Configuration_Register_($A001)` | [INES_Mapper_176](/wiki/INES_Mapper_176/) |
| [INES_Mapper_197](/wiki/INES_Mapper_197/) | `##Submapper_2` | [NES_2.0_Mapper_460](/wiki/NES_2.0_Mapper_460/) |
| [J.Y._Company_ASIC](/wiki/J.Y._Company_ASIC/) | `#Mirroring_Select_($D001)` | [J.Y._Company_ASIC](/wiki/J.Y._Company_ASIC/) |
| [MMC1](/wiki/MMC1/) | `#Hardware` | [Implementing_Mappers_In_Hardware](/wiki/Implementing_Mappers_In_Hardware/) |
| [MMC5](/wiki/MMC5/) | `#8x16_mode_enable_1_($2000_=_PPUCTRL)` | [MMC5](/wiki/MMC5/) |
| [Mirroring](/wiki/Mirroring/) | `#4-screen` | [INES](/wiki/INES/) |
| [NES_2.0](/wiki/NES_2.0/) | `#Byte_10_(RAM_size)` | [Bandai_FCG_board](/wiki/Bandai_FCG_board/), [INES_Mapper_016](/wiki/INES_Mapper_016/), [INES_Mapper_157](/wiki/INES_Mapper_157/) |
| [NES_2.0](/wiki/NES_2.0/) | `#Byte_13_(Vs._hardware)` | [PPU_variants](/wiki/PPU_variants/) |
| [NES_2.0](/wiki/NES_2.0/) | `#Byte_14_(Misc._ROMs)` | [NES_2.0_Mapper_355](/wiki/NES_2.0_Mapper_355/) |
| [NES_2.0](/wiki/NES_2.0/) | `#Byte_9_(Upper_bits_of_ROM_size)` | [Action_53_mapper](/wiki/Action_53_mapper/) |
| [NES_2.0_submappers](/wiki/NES_2.0_submappers/) | `#002,_003,_007:_UxROM,_CNROM,_AxROM` | [AxROM](/wiki/AxROM/), [UxROM](/wiki/UxROM/) |
| [NES_2.0_submappers](/wiki/NES_2.0_submappers/) | `#iNES_Mapper_032_/_Irem_G101` | [INES_Mapper_032](/wiki/INES_Mapper_032/) |
| [NTSC_video](/wiki/NTSC_video/) | `#Color_phases` | [PAL_video](/wiki/PAL_video/) |
| [PPU_OAM](/wiki/PPU_OAM/) | `#Byte_1` | [VT02+_Video_Modes](/wiki/VT02+_Video_Modes/) |
| [PPU_OAM](/wiki/PPU_OAM/) | `#Byte_2` | [PPU_palettes](/wiki/PPU_palettes/), [PPU_programmer_reference](/wiki/PPU_programmer_reference/), [PPU_sprite_priority](/wiki/PPU_sprite_priority/), [VT02+_Video_Modes](/wiki/VT02+_Video_Modes/) |
| [PPU_registers](/wiki/PPU_registers/) | `#Sprite_0` | [Don't_hardcode_OAM_addresses](/wiki/Don't_hardcode_OAM_addresses/) |
| [PPU_scrolling](/wiki/PPU_scrolling/) | `#$2006_second_write_(w_is_1)` | [PPU_glitches](/wiki/PPU_glitches/) |
| [PPU_scrolling](/wiki/PPU_scrolling/) | `#$2007_reads_and_writes` | [PPU_glitches](/wiki/PPU_glitches/), [PPU_programmer_reference](/wiki/PPU_programmer_reference/), [PPU_registers](/wiki/PPU_registers/) |
| [PPU_scrolling](/wiki/PPU_scrolling/) | `#PPU_registers` | [The_frame_and_NMIs](/wiki/The_frame_and_NMIs/) |
| [Programming_with_unofficial_opcodes](/wiki/Programming_with_unofficial_opcodes/) | `#Watermarking_instructions` | [Tricky-to-emulate_games](/wiki/Tricky-to-emulate_games/) |
| [Scanning_tables](/wiki/Scanning_tables/) | `#Efficient_Forwards_Scanning` | [Scanning_large_tables](/wiki/Scanning_large_tables/) |
| [Standard_controller](/wiki/Standard_controller/) | `#Evil_Details` | [Power_Pad](/wiki/Power_Pad/) |
| [VRC2_and_VRC4](/wiki/VRC2_and_VRC4/) | `#Mirroring_Control_($9000,_$9001,_$9002,_$9003)` | [NES_2.0_Mapper_266](/wiki/NES_2.0_Mapper_266/) |

## Links to missing pages

The backup has no page for these names, so the links were replaced by their text.

| Missing page | Links |
| --- | ---: |
| `User:Cluster` | 2 |
| `mediawikiwiki:Help:Contents` (page on mediawiki.org) | 1 |

## Dropped redirect pages

These pages only redirected to a page outside the migrated namespaces, so they were not turned into redirects.
Old links to them end in a 404.

| Old page | Redirected to |
| --- | --- |
| `Blargg_Main`, `Blargg_PPU`, `Blargg_SPC_Upload` | `User:Blargg/...` pages |
| `Help_Contents`, `Help_Editing`, `Help_Editing_pages`, `Help_Formatting`, `Help_Images`, `Help_Links`, `Help_Signatures`, `Help_Starting_a_new_page`, `Help_Talk_pages`, `Help_User_page` | the same Help pages on mediawiki.org |
| `Help_Group_rights` | `Nesdev wiki:Users` (MediaWiki permission page, not migrated) |

## Template syntax left in the text

The old wiki rendered `{{#css:...}}` blocks as style rules. The importer has no equivalent, so
[Fixed cycle delay](/wiki/Fixed_cycle_delay/) shows the raw block (`{{#css: .testtable td{padding:2px} ... }}`) in the middle of the
page. Delete it; the layout CSS already styles `testtable`.

## Not migrated

| Namespace | Pages in backup | Why |
| --- | ---: | --- |
| `File:` | 619 | Image description pages. The images themselves are kept in `static/wiki-images/`, but their credits and license notes are not. |
| `Talk:` and other `... talk:` | 316 + 75 + 6 + 4 + 3 + 3 | Discussion pages. |
| `User:` | 222 | Personal pages. |
| `Template:` | 80 | MediaWiki templates, their output is already part of the pages. |
| `MediaWiki:` | 24 | Interface messages. |
| `Nesdev wiki:` | 8 of 12 | Permission groups and spam filters. About, General disclaimer, Privacy policy and Manual of Style are kept. |

## URL changes

URLs keep the old `/wiki/<name>/` layout except for names with a colon, because Hugo cannot build a page with
a colon in its URL and Cloudflare `_redirects` treats `:name` as a placeholder.

| Old URL | New URL |
| --- | --- |
| `/wiki/Category:<name>` | `/category/<name>/` |
| `/wiki/Nesdev_wiki:<name>` and other namespaces | `/wiki/Nesdev_wiki_<name>/` (colon becomes `_`) |

No redirects exist for the old colon URLs. Add them in a Cloudflare Pages function or in the Cloudflare
dashboard if links to the old category URLs matter.
