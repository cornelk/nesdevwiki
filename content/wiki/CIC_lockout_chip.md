---
title: "CIC lockout chip"
url: "/wiki/CIC_lockout_chip/"
description: "The frontloading NES has a CIC Lockout Chip, a microcontroller that performs a proprietary handshake, as an anticompetitive measure. The Famicom and toploading NES consoles do not contain this chip.…"
wikiPage: "CIC_lockout_chip"
---

The frontloading NES has a **CIC Lockout Chip**, a microcontroller that performs a proprietary handshake, as an anticompetitive measure. The [Famicom](../Family_Computer/) and toploading NES consoles do not contain this chip. The abbreviation CIC is short for "Checking Integrated Circuit" according to Nintendo's patents.

## Overview {#Overview}

Both the lockout chip inside the NES and the one on the cartridge are the same IC. The one inside the NES acts as a lock and the one in the cartridge a key. The difference is how they are hooked up. The system is wired so that the output of one CIC is connected to the input of other and vice versa. LOCK/KEY is pulled to +5V inside the NES and grounded on the Cart. Both share the same 4MHZ clock on pin 6. The RESET pin on the key is connected to SLAVE CIC RESET on the lock. The lock's RESET pin is connected to the system reset bus. This can be demonstrated by inserting a game with the system already on. The NES will not work until you press the reset button which will reset the lock CIC, which in turn resets the key. /CPU & PPU RESET is not connected on the key, on the lock it is connected to the CPU and PPU reset pins. Pins 11-15 are grounded on both CIC's in an NES; these are actually used in multi-game systems so that multiple CICs may be addressed within one system. Finally VCC goes to +5V.

Once the system comes out of power-on or reset, the Lock sends the appropriate reset and initialization signals to the key. The key must then return the correct response, otherwise the lock will pull the /CPU & PPU RESET line low with a 1Hz square wave. Since both share the same clock and the lock is able to reset the key, both CIC's stay in sync with each other.

## Regions {#Regions}

One purpose of the CIC was to restrict the usage of NES consoles and games to the regions they were distributed in, in order to combat out-of-region imports. This was especially important for the PAL and Asian regions since third-parties such as Mattel, ASD, and Bandai were responsible for distributing the NES in those regions before Nintendo subsidiaries were created.

The table below summaries the CIC regions used in 72-pin cartridges. (The [Famicom Network System](../Family_Computer_Network_System/) contains CICs but does not use 72-pin connectors)

| CIC Code | CIC Region                   |
|----------|------------------------------|
| 3193     | NTSC                         |
| 6113     | NTSC                         |
| 3197     | PAL-A                        |
| 3195     | PAL-B                        |
| 3196     | Asia                         |
| 3198     | [FamicomBox](../FamicomBox/) |

### Determining cartridge CIC regions {#Determining_cartridge_CIC_regions}

It may be useful to identify the CIC region of a game without physical access to a frontloader NES, screwdriver, or the cartridge itself. To determine the CIC region of a licensed NES game:

1.  Inspect the cartridge's front label. If it contains an ID code (e.g. *NES-SM-USA* for the USA release of *Super Mario Bros.*) or only a region code, then look up the region code in the table below.
    - Some PAL games may specify PAL-A/B on the front label.
    - FamicomBox cartridges have yellow labels with no ID code, and are upside-down compared to standard NES cartridges.
2.  If the ID code is not present (such as in early first-party releases), check the rear label, box, or manual for possible region codes.
3.  Otherwise, open the cartridge using the correct screwdriver and inspect the CIC itself.

The table below maps the known region codes to their respective CIC regions.

| Region Code | Meaning | CIC Region |
|----|----|----|
| USA | United States of America | NTSC |
| CAN | Canada | NTSC |
| AUS | Australia | PAL-A |
| GBR | Great Britain | PAL-A |
| ITA | Italy | PAL-A |
| UKV | United Kingdom | PAL-A |
| EEC | European Economic Community | PAL-B |
| ESP | España (Spain) | PAL-B |
| FRA | France | PAL-B |
| FRG | Federal Republic of Germany | PAL-B |
| GPS | Multiple countries (English box with local-language sticker & manual) | PAL-B |
| HOL | Holland (Netherlands) | PAL-B |
| KOR | Korea (Hyundai Comboy) | PAL-B |
| NOE | Nintendo of Europe | PAL-B |
| SCN | Scandinavia (Denmark, Norway, Sweden, Finland, Iceland) | PAL-B |
| SWE | Sweden | PAL-B |
| ASI | Asia | Asia |
| HKG | Hong Kong | Asia |

Codes which only appear on boxes/manuals (TODO: CIC regions):

- GRE - Greece
- ISR - Israel

Codes which only appear on rear cartridge labels (usually to denote languages):

- ASD - logo for Audio Sound Distribution (distributor of early FRA releases, later merged/replaced by Bandai)
- DAS - Deutsch (German) & Spanish
- EAI - England, Australia (Oceania), Italy
- FAH - French & Hollands (Dutch)

Officially rebranded regional variants (TODO: verify console CICs):

- Playtronic (Brazil) - NTSC ([PAL-M video](https://en.wikipedia.org/wiki/PAL-M) conversion is handled by an additional board)
- Samurai (India) - Asia

## Disabling {#Disabling}

Disabling the CIC is a popular after-market modification which allows out-of-region games to boot on the frontloader NES, though game compatibility is not guaranteed. This is especially useful on PAL units due to the A/B region split.

### Two-wire method {#Two-wire_method}

Using just two wires, the CIC can be held in reset and the console's reset button can directly control the CPU and PPU's /RESET inputs instead. Compared to the pin 4 method, this has the advantage of being reversible while still being fairly simple to do. It is done by connecting the lock CIC's pin 7 (CIC RESET in) to the 74HCU04 inverter's pin 1 (oscillator), and the 74HCU04 inverter's pin 2 to the lock CIC's pin 9 (CPU and PPU /RESET in).

### Pin 4 method {#Pin_4_method}

When communication between the two CICs fails, it is only the lock, not the key, the asserts its reset output. Therefore, configuring the CIC in the console as a key prevents it from resetting. This can be done by disconnecting the console CIC's pin 4 from the board. Optionally, this disconnected pin can be tied to ground, but the CIC has internal pulldowns that make the pin read as 0 even if left floating.

## Defeating {#Defeating}

Boards made by [Camerica](https://en.wikipedia.org/wiki/Camerica), [Color Dreams](https://en.wikipedia.org/wiki/Color_Dreams), [AVE](https://en.wikipedia.org/wiki/American_Video_Entertainment), and [AGCI](https://en.wikipedia.org/wiki/American_Game_Cartridges), as well as older Famicom to NES adapters, contain a charge pump to create waveforms involving -5V, which freezes the CIC. Later runs of frontloading NES consoles have diodes to protect against out-of-spec voltages on the CIC data pins, but not on the reset pins.

In late 2006, Tengen's "Rabbit" chip was completely reverse-engineered and a [PIC](https://en.wikipedia.org/wiki/PIC_microcontroller)-based [clone](http://forums.nesdev.org/viewtopic.php?t=1219) was successfully made.

## Pinout {#Pinout}

- [CIC lockout chip pinout](../CIC_lockout_chip_pinout/)

## Reference {#Reference}

- The CIC is explained in US patents 4,799,635 and 5,070,479.
- A CIC stun circuit/IC is patented in US patent 5,004,232. Chips based on this patent were used in [NINA-001/NINA-002](../INES_Mapper_034/#NINA-001.2FNINA-002).
- <https://www.thelegendofnes.com/ntsc-pal>
- <https://www.thelegendofnes.com/id-codes>
- <https://www.thelegendofnes.com/brazil>
