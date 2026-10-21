---
title: "NES 2.0 Mapper 328"
url: "/wiki/NES_2.0_Mapper_328/"
description: "NES 2.0 mapper 328 is used for the Test Ver. 1.01 Dlya Proverki TV Pristavok test cartridge. Its UNIF board name is UNL-RT-01. There is only 16K of PRG and 2K of CHR and mirroring is hard-wired."
wikiPage: "NES_2_0_Mapper_328"
---

NES 2.0 mapper 328 is used for the *Test Ver. 1.01 Dlya Proverki TV Pristavok* test cartridge. Its UNIF board name is **UNL-RT-01**. There is only 16K of PRG and 2K of CHR and mirroring is hard-wired.

## <span id="Random_number_generator_or_Protection_.28.24CE80-.24CEFF.2C_.24FE80-.24FEFF.2C_read.29"></span>Random number generator or Protection (\$CE80-\$CEFF, \$FE80-\$FEFF, read) {#Random_number_generator_or_Protection_.28.24CE80-.24CEFF.2C_.24FE80-.24FEFF.2C_read.29}

    Mask: unknown. If what CaH4e3 wrote (quoted below) is accurate, $BFFF.

    D~7654 3210
      ---------
      1111 RR1R
           ++-+- Random bits

CaH4e3 wrote that the "PRG EPROM has copy protected areas with "weak bits", which is tested at some points of the program. \[The random numbers are t\]rying to simulate "weak bits" behaviour"

The relevant program calculates a 16-bit checksum of the memory in each of these regions and loops until the checksum *changes*.
