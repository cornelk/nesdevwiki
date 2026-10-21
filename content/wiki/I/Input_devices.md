---
title: "Input devices"
url: "/wiki/Input_devices/"
description: "The NES has two general-purpose controller ports on the front of the console, as well as a (rarely used) 48-pin expansion port underneath."
wikiPage: "Input_devices"
---

The NES has two [general-purpose controller ports](../Controller_port_pinout/) on the front of the console, as well as a (rarely used) [48-pin expansion port](../Expansion_port/) underneath.

The Famicom's standard controllers are hardwired to the front of the unit, and a special [15-pin expansion port](../Expansion_port/) is commonly used for third-party controllers. The AV Famicom, however, features detachable controllers using the same ports as the NES.

The NES and Famicom have a set of I/O ports used for controllers and other peripherals, consisting of the following:

- One output port, 3 bits wide, accessible by writing the bottom 3 bits of \$4016.
  - The values latched by \$4016/write appear on the [OUT0-OUT2](../CPU_pinout/) output pins of the 2A03/07, where OUT0 is routed to the controller ports and OUT0-OUT2 to the expansion port on the NES.
- Two input ports, each 5 bits wide, accessible by reading the bottom 5 bits of \$4016 and \$4017. Reading \$4016 and \$4017 activates the [/OE1 and /OE2](../CPU_pinout/) signals, respectively, which are routed to the controller ports and the expansion port.
  - On the NES, only D0, D3, and D4 are connected to both controller ports, while all of D0-D4 are connected to the expansion port.
  - On the original Famicom, the two ports differ: \$4016 D0 and D2 and \$4017 D0 are permanently connected to both controllers, while \$4016 D1 and all of \$4017's D0-D4 are connected to the expansion port.
  - On the AV Famicom, only D0 is connected to the controller ports. The expansion port is unchanged.

## <span id="Programmer.27s_reference"></span>Programmer's reference {#Programmer.27s_reference}

- [Controller reading](../Controller_reading/)

## Hardware {#Hardware}

- [Controller port pinout](../Controller_port_pinout/)
- Controllers
  - **NES [Standard controller](../Standard_controller/)**
  - [Arkanoid controller](../Arkanoid_controller/) (also see [prototype version](../Arkanoid_II_prototype_controller/))
  - [Bandai Hyper Shot](../Bandai_Hyper_Shot/)
  - [Coconuts Pachinko](../Coconuts_Japan_Pachinko_Controller/)
  - [Doremikko Keyboard](../Doremikko_Keyboard/)
  - [Exciting Boxing Punching Bag](../Exciting_Boxing_Punching_Bag/)
  - [Family BASIC Keyboard](../Family_BASIC_Keyboard/)
  - [Four Score, NES Satellite](../Four_player_adapters/) 4-player adapters
  - [Hori 4 Players Adapter](../Four_player_adapters/) for Famicom
  - [Jissen Mahjong controller](../Jissen_Mahjong_controller/)
  - [Konami Hyper Shot](../Konami_Hyper_Shot/)
  - [Miracle Piano](../Miracle_Piano/)
  - [Mouse](../Mouse/) (SNES Mouse, Subor Mouse, Hori Track, etc.)
  - [Oeka Kids tablet](../Oeka_Kids_tablet/)
  - [Partytap](../Partytap/)
  - ~~Power Glove~~
  - [Power Pad, Family Trainer, Pokkun Moguraa Tap-tap](../Power_Pad/) mats
  - ~~RacerMate Bicycle~~
  - [SNES controller](../SNES_controller/)
  - ~~Top Rider Bike~~
  - ~~U-Force~~
  - [Virtual Boy controller](../Virtual_Boy_controller/)
  - [Zapper](../Zapper/) lightgun
- [Infrared controllers](../Infrared_controllers/)

## <span id="Other_I.2FO_devices"></span>Other I/O devices {#Other_I.2FO_devices}

- [Famicom 3D glasses](../Famicom_3D_System/)
- [Family BASIC Data Recorder](../Family_BASIC_Data_Recorder/)
- [R.O.B.](../R.O.B./)
- [Battle Box](../Battle_Box/)
- [Turbo File](../Turbo_File/)
- [Barcode Battler](../Barcode_Battler/)
- [TV-NET controller](../TV-NET_controller/)
- [TV-NET Rank 2 controller](../TV-NET_Rank_2_controller/)
- [FAM-NET Keyboard](../FAM-NET_Keyboard/)
- [Family Computer Network Adapter](../Family_Computer_Network_Adapter/)
- [Famicom Network System controller](../Famicom_Network_Controller/)
- [Super Famicom NTT Data Keypad](../Super_Famicom_NTT_Data_Keypad/)
- [Namco Serial Checker Interface](../Namco_Serial_Checker_Interface/)
- [Namco IPL Interface](../Namco_IPL_Interface/)
- [Port test controller](../Port_test_controller/)

## Usage of port pins by hardware type {#Usage_of_port_pins_by_hardware_type}

<table class="wikitable">
<tbody>
<tr>
<td style="text-align: left;">type</td>
<td colspan="3" style="border-left: 2px solid black">output</td>
<td colspan="6" style="border-left: 2px solid black">Joypad 1</td>
<td colspan="6" style="border-left: 2px solid black">Joypad 2</td>
<td style="border-left: 2px solid black">audio output</td>
</tr>
<tr>
<td style="text-align: left;">signal</td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td style="border-left: 2px solid black">AUDIO</td>
</tr>
<tr>
<td style="text-align: left;">access method</td>
<td colspan="3" style="text-align: center; border-left: 2px solid black;">write $4016</td>
<td style="border-left: 2px solid black"><sup><span id="cite_ref-1"></span><a href="#cite_note-1">[1]</a></sup></td>
<td colspan="5" style="text-align: center; border-left: 2px solid gray;">read $4016</td>
<td style="border-left: 2px solid black"><sup><span id="cite_ref-2"></span><a href="#cite_note-2">[2]</a></sup></td>
<td colspan="5" style="text-align: center; border-left: 2px solid gray;">read $4017</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left; border-top: 2px solid black;">available on these ports</td>
<td colspan="3" style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td colspan="5" style="border-left: 2px solid gray; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td colspan="5" style="border-left: 2px solid gray; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Controller port 1 (AV Famicom)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Controller port 1 (Famicom (internal))</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Controller_port_pinout/">Controller port</a> 1 (NES)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Controller port 2 (AV Famicom)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Controller port 2 (Famicom (internal))</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td>D2</td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Controller_port_pinout/">Controller port</a> 2 (NES)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Expansion_port/#Famicom">Expansion port</a> (Famicom)</td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td style="text-align: center; border-left: 2px solid black;">AUDIO</td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Expansion_port/#NES">Expansion port</a> (NES)</td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td style="text-align: center; border-left: 2px solid black;">AUDIO</td>
</tr>
<tr>
<td style="text-align: left; border-top: 2px solid black;">used by these devices</td>
<td colspan="3" style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td colspan="5" style="border-left: 2px solid gray; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
<td colspan="5" style="border-left: 2px solid gray; border-top: 2px solid black"></td>
<td style="border-left: 2px solid black; border-top: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Standard_controller/">Controller</a> (port 1)<sup><span id="cite_ref-port_3-0"></span><a href="#cite_note-port-3">[3]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Standard_controller/">Controller</a> (port 2)<sup><span id="cite_ref-port_3-1"></span><a href="#cite_note-port-3">[3]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Standard_controller/">Controller</a> (Famicom controller 2)<sup><span id="cite_ref-4"></span><a href="#cite_note-4">[4]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td>D2</td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Standard_controller/">Controller</a> (expansion port)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="text-align: center; border-left: 2px solid black;">AUDIO<sup><span id="cite_ref-5"></span><a href="#cite_note-5">[5]</a></sup></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Arkanoid_controller/">Arkanoid controller</a> (port 2)<sup><span id="cite_ref-port_3-2"></span><a href="#cite_note-port-3">[3]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Arkanoid_controller/">Arkanoid controller</a> (expansion port)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Arkanoid_controller/#Arkanoid_II_expansion_port">Arkanoid II controller</a> (2 controllers)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Bandai_Hyper_Shot/">Bandai Hyper Shot</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td></td>
<td style="text-align: center; border-left: 2px solid black;">AUDIO</td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Exciting_Boxing_Punching_Bag/">Exciting Boxing Punching Bag</a></td>
<td style="border-left: 2px solid black"></td>
<td>OUT1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../FAM-NET_Keyboard/">FAM-NET Keyboard</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Power_Pad/#Family_Trainer_Mat">Family Trainer Mat</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Family_BASIC_Keyboard/">Family BASIC Keyboard</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Famicom_3D_System/">Famicom 3D System</a></td>
<td style="border-left: 2px solid black"></td>
<td>OUT1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Famicom_Network_Controller/">Famicom Network System controller</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Four player adapter (<a href="../Four_player_adapters/">Four Score</a>)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;">Four player adapter (<a href="../Four_player_adapters/">Hori 4 Players Adapter</a>)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Hori_Track/">Hori Track</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Jissen_Mahjong_controller/">Jissen Mahjong controller</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Konami_Hyper_Shot/">Konami Hyper Shot</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Oeka_Kids_tablet/">Oeka Kids tablet</a></td>
<td style="border-left: 2px solid black"></td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td>D3</td>
<td>D2</td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Coconuts_Japan_Pachinko_Controller/">Pachinko controller</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Partytap/">Party Tap</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Power_Pad/#Family_Trainer_Mat">Pokkun Moguraa Tap-tap Mat</a></td>
<td style="border-left: 2px solid black">OUT2</td>
<td>OUT1</td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Power_Pad/">Power Pad</a> (port 2)<sup><span id="cite_ref-port_3-3"></span><a href="#cite_note-port-3">[3]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Port_test_controller/">Port test controller</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black">/OE2</td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../TV-NET_controller/">TV-NET controller</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../TV-NET_Rank_2_controller/">TV-NET Rank 2 controller</a></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td>D1</td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Zapper/">Zapper</a> (port 2)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Zapper/">Zapper</a> (expansion port)</td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray">D4</td>
<td>D3</td>
<td></td>
<td></td>
<td></td>
<td style="text-align: center; border-left: 2px solid black;"><sup><span id="cite_ref-6"></span><a href="#cite_note-6">[6]</a></sup></td>
</tr>
<tr>
<td style="text-align: left;"><a href="../Zapper/">Zapper</a> (Vs. System) (port 1)<sup><span id="cite_ref-port_3-4"></span><a href="#cite_note-port-3">[3]</a></sup></td>
<td style="border-left: 2px solid black"></td>
<td></td>
<td>OUT0</td>
<td style="border-left: 2px solid black">/OE1</td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td>D0</td>
<td style="border-left: 2px solid black"></td>
<td style="border-left: 2px solid gray"></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td style="border-left: 2px solid black"></td>
</tr>
</tbody>
</table>

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) /OE1 is activated by reading \$4016.
2.  <span id="cite_note-2"></span>[↑](#cite_ref-2) /OE2 is activated by reading \$4017.
3.  <span id="cite_note-port-3"></span>↑ <sup>[3.0](#cite_ref-port_3-0)</sup> <sup>[3.1](#cite_ref-port_3-1)</sup> <sup>[3.2](#cite_ref-port_3-2)</sup> <sup>[3.3](#cite_ref-port_3-3)</sup> <sup>[3.4](#cite_ref-port_3-4)</sup> Controllers using NES ports can be plugged into either port, using that port's /OE and data lines. However, games may expect a controller to only be in a specific port.
4.  <span id="cite_note-4"></span>[↑](#cite_ref-4) The Famicom controller 2 has a microphone that sends audio input over \$4016 D2. This is not affected by OUT0 nor /OE2.
5.  <span id="cite_note-5"></span>[↑](#cite_ref-5) A Famicom expansion controller may connect the audio output signal to a headphone jack (for example: IQ502 joypad).
6.  <span id="cite_note-6"></span>[↑](#cite_ref-6) The Casel Zapper plays audio when the trigger is pulled, but this is done entirely by the controller independent of the console's audio out.
