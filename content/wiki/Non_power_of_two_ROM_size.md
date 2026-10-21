---
title: "Non-power-of-two ROM size"
url: "/wiki/Non-power-of-two_ROM_size/"
description: "Because of the cost of soldering multiple DIP chips to a board, it wasn't common to have more than one PRG ROM or more than one CHR ROM on an NES cartridge board. Below is a list of such boards:"
wikiPage: "Non_power_of_two_ROM_size"
---

Because of the cost of soldering multiple DIP chips to a board, it wasn't common to have more than one PRG ROM or more than one CHR ROM on an NES cartridge board. Below is a list of such boards:

    Board Name        | Mpr | Type                | What switches | Games                                 | PCB
    ------------------+-----+---------------------+---------------+---------------------------------------+------
    STROM             | 0   | PRG 2x8k            | Transistor    | Baseball                              | [1]
                      |     |                     |               | Pinball                               | [2]
    ------------------+-----+---------------------+---------------+---------------------------------------+-------
    ?                 | ?   | PRG 1M + 256k       | PAL16L8       | PEGASUS 5 in 1                        | [3]
    B-5064            | ?   | PRG 512k + 3x256k   | PAL16L8       | PEGASUS 5 in 1                        | [4]
    ------------------+-----+---------------------+---------------+---------------------------------------+-------
    SMROM             | 1   | PRG 2x128k          | ROM +CE (*)   | Hokkaidou Rensa Satsujin: Okhotsu ... | [5]
    023-N507          | 228 | PRG 3x512k          | PAL16L8       | Action 52                             | [6]
    JUMP A            | 16  | PRG 2x128k          | 74139         | Famicom Jump: Eiyuu Retsuden          | [7]
    TENGEN-800042     | 68  | CHR 2x128k          | 7402?         | After Burner                          | [8]
    UNK-SUNSOFT-AFB   | 68  | CHR 2x128k          | ROM +CE (*)   | After Burner                          | [9]
    AKIRA-A           | 33  | CHR 2x128k          | 74139         | Akira                                 | [10]
    TH2180-2          | 193 | CHR 2x128k          | NTDEC TC-112  | Fighting Hero                         | [11]
    FB-N-128-02       | 0   | PRG 2x16k           | 7420          | Family BASIC                          | [12]
    RTROM             | 0   | PRG 2x8k            | Transistor    | Excitebike                            | [13]
    UNK-HVC-023-PROTO | 2   | PRG 4x32k           | 74139         | Chester Field: Ankoku Shin ... (Proto)| [14]
    NES-EVENT-02      | 105 | PRG 2x128k          | 7404          | Nintendo World Championships 1990     | [15]
    NES-SNWEPROM      | 1   | PRG 2x128           | Transistor?   | Final Fantasy II (Proto)              | [16]
    BLUE TRAIN        | 16  | CHR 2x128           | 74139         | Nishimura Kyoutarou Mystery: Blue T...| [17]
    SAP-E301          | 69  | CHR 2x128           | 7404          | Batman: Return of the Joker (Proto)   | [18]
    ------------------+-----+---------------------+---------------+---------------------------------------+-------
    NES-TKEPROM-01    | 4   | CHR 2x128           | 74139         | Days of Thunder (Proto)               | [19]
    NES-TKEPROM-01    | 4   | CHR 2x128           | 74139         | Gremlins 2: The New Batch (Proto)     | [20]
    NES-TKEPROM-01    | 4   | PRG 2x128           | 74139         | Mega Man 3 (Prototype)                | [21]
    NES-TKEPROM-01    | 4   | CHR 2x128           | 74139         | Where in Time Is Carmen Sandiego?     | [22]
    NES-TKEPROM-02    | 4   | PRG 4x128, CHR 2x128| 74139         | Kirby's Adventure (Proto)             | [23]
    NES-TKEPROM-02    | 4   | PRG 2x128           | 74139         | Tecmo Super Bowl (Proto)              | [24]
    NES-TKEPROM-02    | 4   | PRG 2x128, CHR 2x128| 74139         | Wario's Woods (Proto)                 | [25]
    ------------------+-----+---------------------+---------------+---------------------------------------+-------
    NoSA-8802B        | 513 | PRG 6x256k          | 7400          | Princess Maker                        | [26]

    (*) "ROM +CE" means the circuitry is effectively built into the ROMs.  One of two memories connected
        in parallel has a positive chip enable input while the other has negative.

    This list does not include many pirate versions of official cartriges or multicarts, which often had PRG/CHR
    split into many 128k chips, for example: [27]

Even then, because there were almost always two ROMs of the same size on the PRG or CHR side, the PRG ROM size and the CHR ROM size were each almost always a power of two. The most notable exception on NES is *[Action 52](../INES_Mapper_228/)*, whose mapper is defined to support up to four 512Kx8 (4 Mbit) PRG ROMs. Three are populated on the board, totaling 12 Mbit (1.5 MiB).

Non-power-of-two ROM size became more common in the 16-bit era, when games contained only a PRG ROM. Games of size 10, 12, 20, 24, and 48 Mbit (1.25, 1.5, 2.5, 3, and 6 MiB) exist. In the memory space, the larger ROM is at a lower address, and the smaller ROM is [mirrored](../Mirroring/) to appear as large as the larger ROM. A 12 Mbit ROM usually shows up as one copy of the first 8 Mbit and two copies of the last 4 Mbit (ABCC). A 20 Mbit ROM is one copy of the first 16 Mbit and four copies of the last 4 Mbit (ABCDEEEE). This interpretation matches the checksum values in Super NES ROMs' internal header.

The following doubling algorithm should produce reasonable results for most ROMs with non-power-of-two sizes:

    While ROM size is not a power of two:
        Find the place value of the least significant 1 bit in ROM size
        Double up that many bytes at the end

For example, 10 Mbit is 1310720 (\$140000) bytes. Doubling the last \$40000 bytes produces 12 Mbit, or 1572864 (\$180000) bytes. Further doubling the last \$80000 bytes of this produces 16 Mbit, or 2097152 (\$200000) bytes.

For *Action 52*, this produces reasonable results from an emulation perspective. Though ROM slots 0, 1, and 3 are populated on the actual board (ABxC), this algorithm additionally populates slot 2 with a copy of slot 3's data (ABCC).
