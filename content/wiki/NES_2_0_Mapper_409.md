---
title: "NES 2.0 Mapper 409"
url: "/wiki/NES_2.0_Mapper_409/"
description: "NES 2.0 Mapper 409 denotes the retroUSB DPCMcart circuit board, used for the homebrew music cartridge \"A Winner is You\" released by retroUSB.. It is basically an extremely oversized version of UNROM…"
wikiPage: "NES_2_0_Mapper_409"
wikiCategories:
  - name: "Mappers with CHR RAM"
    url: "/category/Mappers_with_CHR_RAM/"
---

**NES 2.0 Mapper 409** denotes the **retroUSB DPCMcart** circuit board, used for the [homebrew music cartridge "A Winner is You"](https://sergioelisondo.bandcamp.com/album/a-winner-is-you) released by [retroUSB.](https://retrousb.com/product_info.php?cPath=36&products_id=143). It is basically an extremely oversized version of [UNROM](../UxROM/) that latches twelve bits of the address, rather than the data, bus for selecting banks from a 64 MiB PRG address space.

Despite being called "DPCMcart", it cannot actually bank DPCM samples because of its fixed bank. Instead, *A Winner is You* uses the large ROM for PCM sample storage, played back with timed CPU code.

# Banks {#Banks}

- CPU \$8000-\$BFFF: 16 KiB PRG bank, switchable
- CPU \$C000-\$FFFF: 16 KiB PRG bank, fixed to last bank
- PPU \$0000-\$1FFF. 8 KiB CHR RAM, unbanked
- Mirroring: Hard-wired

# Registers {#Registers}

## <span id="Address_Latch_.28.24C000-.24CFFF.2C_write.29"></span>Address Latch (\$C000-\$CFFF, write) {#Address_Latch_.28.24C000-.24CFFF.2C_write.29}

    Register Mask: $8000 ?
       Value Mask: $0FFF

    A~FEDC BA98 7654 3210
      -------------------
      1... BBBB BBBB BBBB
           ++++-++++-++++- Selects bank via PRG A25..A14 at CPU $8000-BFFF (CPU A14=0)

This register may respond to all of \$8000-FFFF, but it is unknown whether it does. *A Winner is You* only uses the \$C000-CFFF range.

# See also {#See_also}

- <https://forums.nesdev.org/viewtopic.php?t=15188>
- <https://web.archive.org/web/20191024003230/nintendoage.com/forum/messageview.cfm?catid=34&threadid=160050>
