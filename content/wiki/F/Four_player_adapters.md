---
title: "Four player adapters"
url: "/wiki/Four_player_adapters/"
description: "Four player adapters are devices that plug into controller or expansion ports and provide additional ports. These ports can be used as alternatives to hardwired controllers or for multiplayer with…"
wikiPage: "Four_player_adapters"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

Four player adapters are devices that plug into controller or expansion ports and provide additional ports. These ports can be used as alternatives to hardwired controllers or for multiplayer with more than 2 simultaneous players. Contemporary adapters allow some consoles to interface with up to 8 controllers. These adapters come in multiple forms, each with their own interfaces and limitations.

# Adapters {#Adapters}

## <span id=".22Simple.22_Famicom_adapters"></span>"Simple" Famicom adapters {#.22Simple.22_Famicom_adapters}

Because the Famicom only has one expansion port, a device such as the **JoyPair** or **Twin Adapter** is required to attach multiple controllers to it. Standard expansion controllers respond on D1 of \$4016, and these passive multi-port devices redirect the second port to D1 of \$4017. These may be used as players 3 and 4. In 1- or 2-player games for the Famicom, though, it is common to OR together the D1 and D0 bits of a joypad read to allow players to use external controllers as substitutes for the hardwired ones.

The expansion ports on these multi-port devices do not pass through all of the signals and thus are limited in functionality compared to the console's expansion port. Because of this, many controllers are not compatible with these adapters.

### Pinout {#Pinout}

         JoyPair CN1 |  | Famicom EXP
    -----------------+--+-----------------
             GND  01 |--| 01  GND
            OUT0  12 |<-| 12  OUT0
    Joypad 1 /D1  13 |->| 13  Joypad 1 /D1
    Joypad 1 /OE  14 |<-| 14  Joypad 1 /OE
              5V  15 |--| 15  5V

         JoyPair CN2 |  | Famicom EXP
    -----------------+--+-----------------
             GND  01 |--| 01  GND
            OUT0  12 |<-| 12  OUT0
    Joypad 1 /D1  13 |->| 07  Joypad 2 /D1
    Joypad 1 /OE  14 |<-| 09  Joypad 2 /OE
              5V  15 |--| 15  5V

    (All other JoyPair pins not connected)

## Hori 4 Players Adapter {#Hori_4_Players_Adapter}

The **Hori 4 Players Adapter** is a Famicom expansion port device that provides 4 ports, usable in 2- and 4-controller modes selected using a switch. In 2-controller mode, ports 1 and 2 on the adapter function exactly like a "simple" Famicom adapter, and ports 3 and 4 are disabled. In 4-controller mode, ports 1 and 3 are multiplexed onto \$4016 D1 and ports 2 and 4 onto \$4017 D1. The first 8 reads from an address provide the joypad 1 D1 report from the first controller, the next 8 reads the joypad 1 D1 report from the second controller, and the next 8 reads a signature unique to that address. The read counter is reset as long as OUT0 is 1. The same connection limitations of the "simple" adapters still apply in 4-controller mode, which combined with the 8 read limit restricts the set of usable contemporary controllers to those compatible with the standard controller.

### Protocol {#Protocol}

#### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Controller read strobe

This matches the normal strobe behavior used by the [standard controller](../Standard_controller/#Input_.28.244016_write.29) and strobes all 4 at once.

#### <span id="Output_.28.244016.2F.244017_read.29"></span>Output (\$4016/\$4017 read) {#Output_.28.244016.2F.244017_read.29}

    7  bit  0
    ---- ----
    xxxx xxDx
           |
           +-- Serial controller data

#### Report {#Report}

    $4016 read D1:
      0-7  - Joypad 1 D1 bits 0-7 from controller #1
      8-15 - Joypad 1 D1 bits 0-7 from controller #3
     16-17 - (Always 0)
     18    - (Always 1)
     19-23 - (Always 0)

     24+   - (Always 1)

    $4017 read D1:
      0-7  - Joypad 1 D1 bits 0-7 from controller #2
      8-15 - Joypad 1 D1 bits 0-7 from controller #4
     16-18 - (Always 0)
     19    - (Always 1)
     20-23 - (Always 0)

     24+   - (Always 1)

Note: *Joypad 1* and *joypad 2* are names used on the controller-side [Famicom expansion port](../Expansion_port/#Famicom) pin interface, as compared to \$4016 and \$4017 on the CPU side. While joypad 1 normally connects to \$4016, this adapter routes each controller's joypad 1 /D1 to \$4016 or \$4017 as needed.

## Four Score {#Four_Score}

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/NES-Four-Score.jpg"><img src="../../wiki-images/NES-Four-Score.jpg" width="200" height="127" /></a>
<figcaption>NES Four Score</figcaption>
</figure>

The Four Score is an adapter that plugs into both NES controller ports and provides 4 ports, usable in 2- and 4-controller modes. This device is nearly identical to the Hori 4 Players Adapter, differing in that D0 is used instead of D1 and the addresses' signatures are swapped. This device also provides two turbo switches.

### Protocol {#Protocol_2}

#### <span id="Input_.28.244016_write.29_2"></span>Input (\$4016 write) {#Input_.28.244016_write.29_2}

    7  bit  0
    ---- ----
    xxxx xxxS
            |
            +- Controller read strobe

This matches the normal strobe behavior used by the [standard controller](../Standard_controller/#Input_.28.244016_write.29) and strobes all 4 at once.

#### <span id="Output_.28.244016.2F.244017_read.29_2"></span>Output (\$4016/\$4017 read) {#Output_.28.244016.2F.244017_read.29_2}

    7  bit  0
    ---- ----
    xxxx xxxD
            |
            +- Serial controller data

#### Report {#Report_2}

    $4016 read D0:
      0-7  - Joypad D0 bits 0-7 from controller #1
      8-15 - Joypad D0 bits 0-7 from controller #3
     16-18 - (Always 0)
     19    - (Always 1)
     20-23 - (Always 0)

     24+   - (Always 1)

    $4017 read D0:
      0-7  - Joypad D0 bits 0-7 from controller #2
      8-15 - Joypad D0 bits 0-7 from controller #4
     16-17 - (Always 0)
     18    - (Always 1)
     19-23 - (Always 0)

     24+   - (Always 1)

Note: *Joypad D0* refers to the data normally available on \$4016 or \$4017 bit 0. Unlike the Famicom expansion port, NES controller port wiring is relative to the port used, so while a controller normally chooses which data inputs to use, it does not choose whether those are joypad 1 or joypad 2.

### Turbo {#Turbo}

The Four Score includes separate turbo enables intended for the A and B buttons. Each of these applies universally to all attached controllers and in both 2- and 4-controller modes. These turbos force their respective bits to 0 at a fixed rate by ANDing with a mask that alternates between 0 and 1, making a held button appear as though it is being rapidly pressed and released. After a strobe, the A turbo applies to the first read of each controller's report and the B turbo to the second read. In 2-controller mode, each turbo also applies again after every 32 additional reads, causing toggles in the stream of 1's that follow the report from official controllers.

The toggle rate is based on the Four Score's internal clock and is approximately 2 frames.

## NES Satellite {#NES_Satellite}

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/NES-Satellite.jpg"><img src="../../wiki-images/NES-Satellite.jpg" width="200" height="99" /></a>
<figcaption>NES Satellite</figcaption>
</figure>

The NES Satellite is a 4-controller wireless adapter similar to the Four Score. Specific details are not yet known. Instead of a mode switch for 2 or 4 controllers, the NES Satellite has a mode switch selecting between standard controllers and the Zapper. Like the Four Score, it also features A and B turbo switches.

# Hardware {#Hardware}

Inside the Four Score and 4 Players Adapter is a 22-pin DIP labeled [FPA-S01](../FPA-PAL-S01_pinout/). This IC consists of a decoder and one set of shift register, 5-bit counter, and multiplexer per joypad address, allowing each address to return the 8-bit contents of joypads 1 and 3 or 2 and 4 in series, followed by the signature byte.

In case you want to build a four player adapter into an arcade cabinet or some other permanent installation, you don't need to use a Four Score or other adapter: you just need six total shift registers. The inputs to these shift registers are parallel, so you need separate wires for each signal, like from SMS or ZX Spectrum controllers.

# <span id="5.2B_controllers"></span>5+ controllers {#5.2B_controllers}

These adapters can be used in combination with each other on the AV Famicom for up to 8 controllers, or on the Famicom with the hardwired controllers for up to 6. Because nonstandard controllers are not compatible with 4-controller modes, using them reduces the maximum possible controller count, but if a single nonstandard controller such as a mouse is desired, it can be used with a "simple" adapter alongside up to 6 standard controllers on an AV Famicom using a Four Score in the NES controller ports, a "simple" adapter in the expansion port, and a Hori 4 Players Adapter with controllers in ports 1 and 3 in the other "simple" adapter port.

# Compatible Games {#Compatible_Games}

Games known to support 3 or more simultaneous controllers:

| Game | Hori 4-player Adapter | Four Score | Simple |
|----|----|----|----|
| A Nightmare on Elm Street | yes | yes | ? |
| Bomber Man II (J) | no | no | yes |
| Bomberman II (U) | no | yes | ? |
| Downtown Nekketsu Koushinkyoku: Soreyuke Daiundoukai | yes | yes | yes |
| Gauntlet II | no | yes | no |
| Greg Norman's Golf Power | yes | yes | ? |
| Harlem Globetrotters | no | yes | ? |
| Ike Ike! Nekketsu Hockey Bu: Subette Koronde Dai Rantou | no | ? | yes |
| Indy Heat | yes | yes | ? |
| Kings of the Beach | yes | yes | ? |
| Kunio-kun no Nekketsu Soccer League | no | unlikely | yes |
| M.U.L.E. | no | yes | ? |
| Magic Johnson's Fast Break | yes | yes | ? |
| Millionaire (Sachen) | no | no | yes |
| Moero TwinBee: Cinnamon-hakase o Sukue! | no | ? | yes |
| Monster Truck Rally | no | yes | ? |
| NES Play Action Football | yes | yes | ? |
| Nekketsu Kakutou Densetsu | no | ? | yes |
| Nekketsu Koukou Dodge Ball Bu | no | ? | yes |
| Nekketsu Street Basket: Ganbare Dunk Heroes | no | ? | yes |
| Nintendo World Cup (U) | yes | yes | ? |
| R.C. Pro-Am II | yes | yes | ? |
| Rackets and Rivals | no | yes | ? |
| Roundball: 2-on-2 Challenge | no | yes | ? |
| Smash T.V. (twin-stick for 2 players) | no | yes | ? |
| Spot | yes | yes | no |
| Super Jeopardy! | yes | yes | ? |
| Super Off Road | yes | yes | no |
| Super Spike V'Ball | everywhere but title screen | everywhere but title screen | yes |
| Swords and Serpents | no | yes | ? |
| Top Players' Tennis | yes | yes | ? |
| U.S. Championship V'Ball | only during 3+player gameplay | only during 3+player gameplay | yes |
| Wit's | yes | yes | yes |

Homebrew games:

- *Double Action Blaster Guys* becomes *Quadruple Action Blaster Guys* when Four Score is connected (NES)
- *Micro Mages* (NES and Famicom)
- *NESert Golfing* (NES and Famicom)
- *Spacey McRacey* (NES)
- *Super PakPak* (NES and Famicom)

Tech demos:

- *Eighty* (NES and Famicom)
- *allpads* (NES and Famicom)

Hacks:

- *Battle City - 4 Players v1.3 Ti* (NES and Famicom)
- *Battle City Mario - 4 players v1.0 NesDraug* (NES and Famicom)
- *Battletoads - 4 players v2.2 NakeuD2007* (emulator only)
- *Battletoads & Double Dragon - 4 players v0.9 NakeuD2007* (emulator only)
- *Super Dodge Ball - 4 Player Hack v1.00 akatranslations* (Famicom)

# Incompatible Games {#Incompatible_Games}

This section lists games that provide four player game modes without supporting a four player adapter. They are explicitly enumerated here because various lists, including one in Wikipedia, frivolously keep listing (and re-adding after an edit to remove them) these games as supporting a four player adapter.

- *Championship Bowling*: Manual clearly states "Players One and Three use Control Pad 1, and Players Two and Four use Control Pad 2".
- *Jeopardy!* series: Players 1 and 3 share controller 1.

# References {#References}

- [Forum post](https://forums.nesdev.org/viewtopic.php?p=237997#p237997): Analysis of Wit's controller reading, and speculation about the Hori 4 Players adapter.
- [NintendoAge forum thread](http://nintendoage.com/forum/messageview.cfm?catid=22&threadid=65237): Photos of Four Score PCB and hardware discussion.
- [Famicom World forum post](https://www.famicomworld.com/forum/index.php?topic=12543.15): List of Famicom games and their behaviour with Hori 4 Players adapter.
- [Forum post](https://forums.nesdev.org/viewtopic.php?p=238168#p238168): Discussion of Hori 4 Players support in *Micro Mages*.
