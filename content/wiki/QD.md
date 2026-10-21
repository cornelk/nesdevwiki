---
title: "QD"
url: "/wiki/QD/"
description: "The QD file format is an image of Quick Disk media and is the format used for FDS games in emulators by Nintendo."
wikiPage: "QD"
wikiCategories:
  - name: "File formats"
    url: "/category/File_formats/"
---

The **QD** file format is an image of Quick Disk media and is the format used for [FDS](../Family_Computer_Disk_System/) games in emulators by Nintendo.

## QD file format {#QD_file_format}

The QD format (filename suffix `.qd`) is a way to store [Famicom Disk System](../Family_Computer_Disk_System/) disk data. It consists of:

1.  Disk data (**65536** \* *x* bytes)

### Disk data {#Disk_data}

The disk data follows the [FDS disk format](../FDS_disk_format/), and *includes* CRCs after each block. Gaps are not included in the .QD image.

Most games are an even number of sides. Ports from NROM were one side. No commercial FDS game had an odd number of sides greater than 1. Disk sides come in the following order:

- Disk 1 Side A
- Disk 1 Side B
- Disk 2 Side A
- Disk 2 Side B
- etc...

After the CRC value of the last file block, fill a side with all 0 so that the disk side reaches exactly **65536** bytes.

## Differences from FDS file format {#Differences_from_FDS_file_format}

The QD format has significant differences from the popular [FDS file format](../FDS_file_format/) (`.fds`):

- No header. This means that the number of disk sides must always be calculated based on the file size.
- **65536** bytes per disk side instead of 65500 bytes per disk side.
- CRC values are *included* after each file block. Emulators should treat these as bytes that can always be read from/written to the disk, regardless of if they are being used for CRC calculation/verification.
  - Some Virtual Console releases are known to have patched the .QD image without updating the CRCs, so images extracted from these may trigger unexpected load errors. ([BIOS error code](../FDS_BIOS/#Error_list) = 27)

## See also {#See_also}

- [TNES](../TNES/) - 3DS Virtual Console ROM format, which can also encapsulate .QD images.
- [FDS file format](../FDS_file_format/) (**.FDS**)
- [GitHub Gist](https://gist.github.com/einstein95/6545066905680466cdf200c4cc8ca4f0) - Python script to convert between QD/FDS formats.

## References {#References}

- [FDS .qd format](https://forums.nesdev.org/viewtopic.php?f=2&t=15792)
