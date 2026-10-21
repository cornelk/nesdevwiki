---
title: "OAM internals"
url: "/wiki/OAM_internals/"
description: "OAM is subject to hardware bugs and decay that corrupt the stored data, and the behavior of these is derived from OAM's internal structure and addressing mechanism."
wikiPage: "OAM_internals"
---

OAM is subject to hardware bugs and decay that corrupt the stored data, and the behavior of these is derived from OAM's internal structure and addressing mechanism.

## Layout {#Layout}

The 256 bytes of OAM1 (primary OAM) and 32 bytes of OAM2 (secondary OAM) both exist in the same internal dynamic RAM (DRAM), made up of 32 'rows' each containing 9 bytes. The first 8 bytes are consecutive bytes of OAM1 and the 9th byte is OAM2. OAM1 uses an 8-bit address, where the high 5 bits determine the row number and the low 3 bits determine the byte number. OAM2 uses a 5-bit address that simply selects the row number.

Bytes 2 and 6 in each row, which contain OAM1 attributes, only implement 5 bits; the unused bits 4..2 do not exist as storage. These read back as 0. All other bytes are 8 bits, including the bytes used for attributes in OAM2.

## Storage {#Storage}

OAM is implemented similar to static RAM (SRAM) in that it consists of two cross-coupled inverters, so each bit of storage is actually backed by two opposite values where the order determines if the bit is 0 or 1. Unlike SRAM, OAM does not contain a pullup in each cell to maintain the value over time, and this causes a gradual loss of charge. Reading a cell restores its charge, allowing the bit to persist.

Decay occurs when enough charge is lost to make both inverters 0. Two 0's is not a valid state for the cell, and reading it in this state causes it return to one of the two valid states. The resulting bit value is random, but may have somewhat consistent behavior for a given bit in a given chip based on the specific physical properties of that bit.

## Access {#Access}

OAM is accessed a row at a time. Each inverter in a row is connected to a 'bitline', and the pair of bitlines for one bit is called a 'column'. Each bitline connects to the same inverter across all rows. When reading, bitlines are precharged to 1, and then the desired row is selected. The bitlines connected to the 0 halves of each bit are pulled to 0, making the columns take on the state of the row. The columns corresponding to the desired byte are then sent to the OAM buffer to hold that byte's value for use by the rest of the PPU.

When writing, instead of precharging the bitlines of the target byte to 1, they are precharged to the value being written. When the row is selected, this causes the target byte to take on the value in the columns. The bytes that are not being written have their bitlines precharged to 1, causing them to be harmlessly read, instead.

In short, a bit is read if it is selected while its bitlines are 11, and written if they are 01 or 10. The former sets the column to the bit's value, and the latter sets the bit to the column's value.

The PPU accesses OAM every PPU cycle. During the first half of each cycle, the columns are precharged and no row is selected. During the second half, a row is selected and a read or write occurs. Because an entire row is selected at a time, any access refreshes the charge of the entire row, not just the targeted byte.

## Row corruption {#Row_corruption}

Columns must be precharged before selecting a new row. If this does not happen, the columns contain data that was read from or written to the previous row. This functions just like a write and causes that data to be written to the new row, copying all 9 bytes from the previous row.

Unfortunately, the PPU does not synchronize address changes, and so it is possible in many circumstances to change the selected row mid-access, causing the new row to be selected without the columns having first been precharged. This typically results in the first row being copied to the second row. However, depending on the timing, it can randomly corrupt the second row, instead. It is suspected that this analog corruption occurs when the row changes very early in the access (when the columns are still partially precharged) or very late in the access (when the inverters may not have enough time left to stabilize). Because the inverters are physically much smaller than the bitlines, they change state much faster than the bitlines do, so the window for this corruption at the end of the access is likely much smaller than the start. The exact corruption behavior depends on analog factors such as process, voltage, and temperature, as well as the precise timing of factors such as CPU/PPU alignment and PPU register address decoding.
