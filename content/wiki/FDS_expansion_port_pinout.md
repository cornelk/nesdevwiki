---
title: "FDS expansion port pinout"
url: "/wiki/FDS_expansion_port_pinout/"
description: "The FDS has an expansion port located at the rear of the RAM adapter, covered by a shutter. (Labelled as Expansion Port B on the Twin Famicom) The unreleased Famicom Network Adapter was designed to…"
wikiPage: "FDS_expansion_port_pinout"
wikiCategories:
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

The [FDS](../Family_Computer_Disk_System/) has an expansion port located at the rear of the RAM adapter, covered by a shutter. (Labelled as *Expansion Port B* on the Twin Famicom) The unreleased [Famicom Network Adapter](../Family_Computer_Network_Adapter/) was designed to use this port.

A few unlicensed devices are also known to use the port:

- *ILine-PC* (*Iライン-PC*) - to connect the FDS to a PC's printer port.
- *Super Copy System: Hacker Pro Digital* - for use with disk copier software of the same name.

## Diagram {#Diagram}

Open end view of the external connector.

             -------
             |  --  \
      +5V -- | 1||2  | <> EXT6
     EXT5 <> | 3||4  | <> EXT4
     EXT3 <> | 5||6  | <> EXT2
     EXT1 <> | 7||8  | <> EXT0
    SOUND <- | 9||10 | -- GND
             |  --  /
             -------

## Signal descriptions {#Signal_descriptions}

- **+5V** : 5V Power supply from the main voltage regulator.
- **GND** : 0V Power supply.
- **EXT0..6** : Bidirectional - outputs from \$4026.D0..D6 and inputs to \$4033.D0..D6. (open-collector with 4.7K ohm pull-ups)
  - EXT1..2 are altered when \$4025.D5 = 0 but its purpose is unknown.
- **SOUND** : Analog audio output, after the [FDS audio](../FDS_audio/) is mixed in.

## References {#References}

- <https://web.archive.org/web/20140920224500/green.ap.teacup.com/junker/119.html> ([PDF rehosted on forum](https://forums.nesdev.org/viewtopic.php?t=25219))
- [Forum thread](https://forums.nesdev.org/viewtopic.php?t=25882): I2 ILine-PC
