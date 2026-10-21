---
title: "INES Mapper 181"
url: "/wiki/INES_Mapper_181/"
description: "iNES Mapper 181 is currently being used by FCEUX's checksum patch list to represent Seicross v2, which would have been mapper 185 except that it doesn't comply with the heuristic documented there. A…"
wikiPage: "INES_Mapper_181"
---

iNES Mapper 181 is currently being used by FCEUX's checksum patch list to represent *[Seicross v2](http://bootgod.dyndns.org:7777/profile.php?id=2332)*, which would have been [mapper 185](../CNROM/) except that it doesn't comply with the heuristic documented there. A more complex heuristic that pays attention to the addresses read from the PPU would obviate this.

A better solution would be to either elaborate the heuristic, or use submappers with mapper 185. No ROM dumps are believed to exist tagged with this mapper number, and hopefully this mapper number can still be saved for something else later.
