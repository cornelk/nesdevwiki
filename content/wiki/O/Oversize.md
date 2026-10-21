---
title: "Oversize"
url: "/wiki/Oversize/"
description: "Oversize refers to a ROM image that is larger than the largest image used for the licensed boards associated with the same mapper."
wikiPage: "Oversize"
---

*This article is about ROM images with more unique data than the original versions of a cartridge board allowed. For ROM images that contain redundant copies of data, see ~~Overdump~~.*

**Oversize** refers to a [ROM](../ROM/) image that is larger than the largest image used for the licensed boards associated with the same mapper.

As mask ROM capacities increased, Nintendo produced official oversize versions of some cartridge boards: [UNROM (mapper 2) became UOROM](../UxROM/) and [ANROM (mapper 7) became AOROM](../AxROM/). This was done in cases where Nintendo could get away with simply routing more pins of the [74HC161](../74161/) already on the board to the ROM's address inputs. Third parties also produced clones of existing boards with more capabilities. Panesian used an oversize clone of [CNROM (mapper 3)](../CNROM/) for its erotic games *Bubble Bath Babes*, *Peek-a-Boo Poker*, and *Hot Slots*. Likewise, it is believed that the second *Battle Kid* game uses an oversize clone of UxROM.

The iNES mapper corresponding to a particular [mapper chip](../Mapper/) may support larger ROMs than the actual ASIC did. Extra address lines might have been left out of the mapper to cut costs. Pins are one of the most expensive parts of an ASIC, and manufacturers left off extra pins to save money. Each pin for an extra address line costs money, as does the die area occupied by each bit of each register, and upper address lines unlikely to ever see use within the mapper's life would have been a waste of money. For example, mapper 4 theoretically supports up to 2048 KiB of PRG ROM with 8 bank bits, while the actual [MMC3](../MMC3/) has only 6-bit bank registers, hence a 512 KiB limit. This affects discrete mappers as well, as Nintendo was using 74HC161s as 4-bit registers rather than [74HC377s for 8-bit registers](../74377/). AOROM was limited to 256 KiB because three bank bits and one nametable selection bit ate up the whole '161.

Sometimes the existing uses for a mapper's bank bits may get in the way of going too far oversize. [TQROM (mapper 119)](../INES_Mapper_119/) can't go past 64 KiB of CHR ROM because CHR bank bit 6 is the ROM/RAM selection bit. Mapper 7 can't go past 512 KiB of PRG ROM because PRG bank bit 4 is the nametable selection bit.

Some emulators and copiers refuse to correctly load oversize ROMs. For example, early versions of [PowerPak](../PowerPak/) mappers supported only the two least significant bits for CNROM and [BxROM (mapper 34)](../INES_Mapper_034/) and couldn't play larger unlicensed ROMs such as the Panesian games, [Action 53](../Action_53/) multicart builds using BNROM, or *Lizard*. [BNTest](http://forums.nesdev.org/viewtopic.php?p=79826#p79826) tests support for oversize mappers 7 and 34.
