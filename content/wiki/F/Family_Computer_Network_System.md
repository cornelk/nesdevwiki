---
title: "Family Computer Network System"
url: "/wiki/Family_Computer_Network_System/"
description: "The Family Computer Network System (HVC-050) is an add-on for the Family Computer. It facilitates communications with a proprietary dial-up modem network. A dedicated controller was bundled with the…"
wikiPage: "Family_Computer_Network_System"
wikiCategories:
  - name: "ASIC mappers"
    url: "/category/ASIC_mappers/"
  - name: "Expansion audio"
    url: "/category/Expansion_audio/"
  - name: "Mappers using $4020-$5FFF"
    url: "/category/Mappers_using_$4020-$5FFF/"
  - name: "Mappers with IRQs"
    url: "/category/Mappers_with_IRQs/"
  - name: "Pinouts"
    url: "/category/Pinouts/"
---

The Family Computer Network System (HVC-050) is an add-on for the [Family Computer](../Family_Computer/). It facilitates communications with a proprietary dial-up modem network. A [dedicated controller](../Famicom_Network_Controller/) was bundled with the FCNS to provide number pad inputs. Commercial FCNS software was released as proprietary *tsuushin cards* and could accomplish tasks including banking, stock trading, horse race betting, and playing traditional board games such as Shōgi.

The **Dataship 1200** is a standalone console version of the FCNS with the Famicom portion greatly cut down. It only accepts tsuushin cards, lacks hardwired controllers (i.e. only has an [expansion port](../Expansion_port/) to connect the FCNS controller), and uses the same port as the Game Boy link cable for [RF/AV output](#CN07:_Dataship_1200_AV_Port). It also features a 14-pin Centronics port for interfacing with a printer, which is otherwise not present on a standard FCNS.

The unreleased [Family Computer Network Adapter](../Family_Computer_Network_Adapter/) (FCN**A**) for the [FDS](../Family_Computer_Disk_System/) is the predecessor to, but not to be confused with, the FCN**S**.

# System Overview {#System_Overview}

The Famicom Network System is a complicated device with its own memory mapping system and internal CPU. The RF5C66 chip provides the main mapper functionality, delegating its own registers at \$40A0, RF5A18 Famicom registers at \$40D0, an internal Kanji ROM at \$5000, an internal 8kByte W-RAM at \$6000. It also controls the bank of a built-in 16 KiB CHR RAM (using two 8 KiB CHR RAM chips).

The RF5A18 contains CPU2, which is a 65**<u>C</u>**02 processor with its own independent CPU clock. It has a built-in 4kByte ROM. This chip is responsible for controlling the modem communications. It communicates with the Famicom CPU through four bidirectional registers at \$40Dx.

The Famicom Network System plugs into the Famicom through its cartridge connector and provides the user a ZIF style slot to insert a *tsuushin card*. The card is similar to a normal cartridge but does not have access to any PPU signals. Commercial cards are observed to have their own MMC1 memory mapper, which does not interfere with any of the registers of the Famicom Network System. CPU R/W and the CPU data bus are routed through the RF5C66 chip before making it to the card and other internal hardware. It effectively blocks the Famicom Network System from driving the data bus for certain regions of memory, and possibly also is intended to act as a bidirectional buffer / signal conditioner. It is also used for blocking writes when the lid switch is open or when the CIC fails. Older revisions of Famicom Network System also buffered the Famicom address bus with 74HC541 chips, so it is plausible that signal conditioning was a concern.

# LH5323M1 Kanji Graphic ROM {#LH5323M1_Kanji_Graphic_ROM}

The LH5323M1 is a 256kByte graphics ROM containing primarily Kanji data that is mapped at \$5000-5FFF. Each index in this range is a 32-byte space containing 16x16 1bpp graphics, usually for a single character, and each read automatically advances to the next byte in the sequence. There are 2 128kByte banks, and the low bank is default at power-on. Writing 1 to \$40B0.0 selects the high bank. Reading from \$40B0 resets to the beginning of the 32-byte sequence. Writing \$40B0 does not reset the sequence however. No values written to \$40B0 have been observed to arbitrarily change or reset the position in the sequence. Commercial software has been observed to use throwaway reads when accessing data that does not start at the beginning of the 32-byte area.

The Kanji ROM chip is connected directly to the non-buffered Famicom CPU data bus. Writes in the range \$5000-5FFF do activate Kanji ROM /CE and are subject to bus conflict.

# Expansion Audio {#Expansion_Audio}

The Famicom Network System does have expansion audio capabilities. The Famicom audio is routed through the modem module, but nowhere directly to either of the large ASICs. Dial tones have been observed through the television speakers. It is unlikely but unknown if there are other possible sources of sound.

# Disk Drive Support {#Disk_Drive_Support}

According to a block diagram with potentially dubious origins, the RF5C66 chip contains a disk drive controller. Similar design in several ways to the Famicom Disk System, it is suspected that a disk drive can be connected to the expansion port and controlled by the RF5C66. Since this feature was never used, it is unknown how to use or activate it, or even if that feature is fully implemented. The original FDS has a large DRAM that is not present as a discrete chip in the Famicom Network System. It is unknown if such a DRAM could be already integrated into the RF5C66, or could be attached externally and simply not populated, or if a special tsuushin card was to be constructed containing this RAM. All original FDS registers are notably absent and all discovered registers start immediately after where the FDS registers would normally be. There is no obvious path to produce FDS expansion audio. This remains a mystery presently.

One possibility if the RF5C66 follows a similar pinout to the FDS RP2C33:

            79 / -> Exp P3-2  (reg unknown) Serial Out (observed 95.95kHz)
           78 / <- Exp P3-3  (reg unknown) Serial In
          77 / -> Exp P3-4  (!$40A4.2)    Read / Write
         76 / -> Exp P3-5  (!$40A4.1)    Reset Transfer Timing
        75 / -> Exp P3-6  ($40A3.0)     Turn on Motor
       74 / <- Exp P3-7  ($40A5.2)     Write Protect
      73 / <- Exp P3-8  ($40A5.1)     Disk Not Ready
     72 / <- Exp P3-9  ($40A5.0)     Disk Missing
    71 / <- Exp P3-11 ($40A5.7)     Battery Health

# Memory Map {#Memory_Map}

    +================+ $0000 - Famicom Internam RAM
    | Famicom        |
    | Internal RAM   |
    +----------------+ $0800
    |   (Mirrors of  |
    |   $0000-$07FF) |
    +================+ $2000 - Famicom PPU Registers
    | Famicom PPU    |
    | Registers      |
    +----------------+ $2008
    |   (Mirrors of  |
    |   $2000-$2007) |
    +================+ $4000 - Famicom APU, IO, and Test Registers
    | FC APU and IO  |
    | Registers      |
    +----------------+ $4018
    | FC Test Mode   |
    | Registers      |
    +----------------+ $4020
    |   (Open Bus)   |
    +================+ $40A0 - Famicom Network System Internal Registers
    | Famicom Modem  |
    | RF5C66         |
    | Registers      |
    +----------------+ $40D0
    | Famicom Modem  |
    | RF5A18 (CPU2)  |
    | Registers      |
    +----------------+ $40D8
    |   (Mirror of   |
    |   $40D0-$40D7) |
    +----------------+ $40E0
    | Unused Device  |
    | Registers      |
    | (Open Bus)     |
    +----------------+ $40F0
    |   (Open Bus)   |
    +----------------+ $4100
    |   (Open Bus)   |
    +----------------+ $41A0
    |   (Mirror of   |
    |   $40A0-$40FF) |
    +----------------+ $4200
    |   (Mirrors of  |
    |   $4100-$41FF) |
    +================+ $5000 - Famicom Network System Internal Kanji ROM
    | FNS            |
    | LH5323M1       |
    | Kanji ROM      |
    +================+ $6000 - Famicom Network System Internal RAM and Tsuushin Card RAM
    | FNS            |
    | Internal RAM   |
    +================+ $8000 - Tsuushin Card ROM Space
    | Tsuushin Card  |
    | Space          |
    |                |
    +================+ $10000

## Data Bus Behavior {#Data_Bus_Behavior}

The CPU data bus is buffered in both directions through the RF5C66 chip before making it to the tsuushin card and other internal hardware, with the exception of the Kanji graphics ROM and the RF5C66's own registers, which sit directly on the CPU data bus. It appears that the buffer goes from CPU Data Bus to Card Data Bus, or Card Data Bus to CPU Data Bus only. There has not been a high-impedance state observed when testing each possible address in read and write directions on a standalone RF5C66 chip. The CPU Data Bus and Card Data Bus were always observed to be equal despite attaching opposite pull-up and pull-down resistors.

When RF5C66 pin 29 is driven low by the master CIC chip's pin 11, it forces Card R/W always high, thereby preventing most RAM and register writes in the Famicom Network System. Potentially, the data bus buffer has an additional behavior in this mode in order to prevent bus conflicts. The lid switch also imposes such a behavior when the lid is open, but it does not appear that the RF5C66 has a way of knowing when the switch is open, so the bus conflicts do not appear to be prevented in that case.

The propagation delay of the CPU data bus buffer and CPU R/W buffer are measured to be about 16 nsec.

| Address Range | Buffer Direction When Reading | Buffer Direction When Writing | Notes |
|----|:--:|:--:|----|
| \$0x00-\$4x1F | CPU -\> Card | CPU -\> Card | Famicom internal RAM and registers |
| \$4x20-\$4x9F | CPU -\> Card | CPU -\> Card | (unknown; [FDS](../Family_Computer_Disk_System/) registers exactly fit here) |
| \$4xA0-\$4xCF | CPU -\> Card | CPU -\> Card | RF5C66 registers |
| \$4xD0-\$4xDF | CPU \<- Card | CPU -\> Card | CPU 2 registers, /CE at RF5C66 pin 42 |
| \$4xE0-\$4xEF | CPU \<- Card | CPU -\> Card | Unused device registers<sup>1</sup>, /CE at RF5C66 pin 41 |
| \$4xF0-\$4xFF | CPU -\> Card | CPU -\> Card | (unknown) |
| \$5x00-\$5xFF | CPU -\> Card | CPU -\> Card<sup>2</sup> | Kanji graphic ROM, /CE at RF5C66 pin 50 |
| \$6x00-\$7xFF | CPU \<- Card | CPU -\> Card | Famicom Network System internal RAM, /CE at RF5C66 pin 40 |
| \$8x00-\$FxFF | CPU \<- Card | CPU -\> Card | Tsuushin card, /CE is /ROMSEL |

<sup>1</sup> Address range \$4xE0-\$4xEF was meant for a second device similar to CPU 2. Since that device doesn't exist, a tsuushin card could theoretically exploit those 256 addresses for its own purposes; for example, its own read/write registers.

<sup>2</sup> Writes here are subject to bus conflict on the CPU data bus.

# Known Registers {#Known_Registers}

Note: All registers available to the Famicom ignore address bits 8-11 because those bits are not physically connected to the RF5C66. Therefore, register \$4xA0 has mirrors that exist at \$40A0, \$41A0 ... \$4FA0. For simplicity, this page shows all registers as the \$40xx mirror.

<table class="wikitable" style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th>Address</th>
<th style="text-align: center;">Read<br />
Has<br />
Effect</th>
<th style="text-align: center;">Read<br />
Has<br />
Data</th>
<th style="text-align: center;">Write</th>
<th>Owner</th>
<th>Function</th>
</tr>
</thead>
<tbody>
<tr>
<td>$40A0</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>No evidence has been found that this register does or does not exist. It is possible that it is a write-only register that is not used by any known Famicom Network System card. No reads or writes to this register were ever observed to have any effect. However, its apparent absence is suspicious enough to have this placeholder here.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A1</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- Bits exist but function is unknown</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $FF with pull-downs.</li>
<li>Test code writing values $01,02,04,08,10,20,40,80 always read back $FF after each write.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A2</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>IRQ Acknowledge, similar to FDS register $4030
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- Timer Interrupt (1: an IRQ occurred)
||||||+--- Bit exists but function is unknown
||||++---- (unlikely to exist)
++++------ Bits exist but function is unknown</code></pre>
<ul>
<li>Reading this register acknowledges /IRQ.
<ul>
<li>Observed inconsistent behavior acknowledging, possibly suggesting multiple IRQ sources.</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $20 with pull-downs, $2C with pull-ups.</li>
<li>Test code writing values $01,02,04,08,10,20,40,80 always read back $20 after each write.</li>
<li>JRA-PAT:
<ul>
<li>Reads this register with a BIT op-code right before CLI op-code.</li>
<li>Reads this register and makes a decision based on D7.</li>
</ul></li>
<li>Super Mario Club reads this register with BIT op-code.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A3</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- EXP 6 = $40A3.0 (POR value = 0)
+++++++--- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Various values written to this register affect the behavior of pulses on pin 79.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
<li>JRA-PAT writes $2F to this register and appears to keep a RAM copy at $15.</li>
<li>Super Mario Club writes $2F to this register and appears to keep a RAM copy at $15.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A4</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Expansion Port Control
<pre><code>Write
76543210
|||||||+-- (unknown)
||||||+--- EXP 5 = !($40A4.1) (POR value = 0)
|||||+---- EXP 4 = !($40A4.2) (POR value = 0)
+++++----- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Note: Reading test should be repeated with pull-ups and pull-downs on expansion pins.</li>
<li>This register has not been observed read or written to by any commercial software.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A5</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>Expansion Port Input Data
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- Input value of EXP 9
||||||+--- Input value of EXP 8
|||||+---- Input value of EXP 7
|++++----- (unlikely to exist)
+--------- Input value of EXP 11</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $00 with pull-downs, $78 with pull-ups.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A6</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>M2 Cycle Counter LSB, similar to FDS register $4020
<pre><code>Write
76543210
++++++++-- Cycle counter reload value (LSB)
&#10;Read
76543210
++++++++-- Cycle counter present value (LSB)</code></pre>
<ul>
<li>Writing to this register writes to the cycle counter reload value.</li>
<li>Reading this register gives the present value of the counter.</li>
<li>Writing any value to $40A8 resets the counter to the reload value.</li>
<li>This value counts down.</li>
<li>When the value reaches $0000, the next count rolls over to $FFFF or auto-reloads depending on $40A8.0.</li>
</ul></td>
</tr>
<tr>
<td>$40A7</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>M2 Cycle Counter MSB, similar to FDS register $4021
<pre><code>Write
76543210
++++++++-- Cycle counter reload value (MSB)
&#10;Read
76543210
++++++++-- Cycle counter present value (MSB)</code></pre>
<ul>
<li>Refer to description in $40A6, this register being the MSB portion of the counter.</li>
</ul></td>
</tr>
<tr>
<td>$40A8</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>IRQ Control, similar to FDS register $4022
<pre><code>Write
76543210
|||||||+-- IRQ Repeat Flag
||||||+--- IRQ Enable
++++++---- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Writing anything to this register resets the cycle counter to the reload value.</li>
<li>Observed writing $02 makes /IRQ go low, $00 makes /IRQ go high.
<ul>
<li>Reading $40A2 with /IRQ low acknowledges it back high.</li>
<li>Acknowledging with $40A2 first before writing $02 here prevents IRQ immediately going low.</li>
</ul></li>
<li>Observed rollover of cycle counter to $FFFF with repeat flag = 0 and auto-reload when flag = 1.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
<li>JRA-PAT:
<ul>
<li>Writes $00 to this register.</li>
<li>Writes $00 again later, potentially connected to RAM $4F.</li>
<li>Later, right after having written $25 to $40A7 and $20 to $40A6, writes $02 to this register.</li>
<li>Has these various sequences hard-coded:
<ul>
<li>$25-&gt;$40A7, $20-&gt;$40A6, $02-&gt;$40A8</li>
<li>$1C-&gt;$40A7, $10-&gt;$40A6, $02-&gt;$40A8</li>
<li>$03-&gt;$40A7, $19-&gt;$40A6, $02-&gt;$40A8</li>
<li>$06-&gt;$40A7, $F1-&gt;$40A6, $02-&gt;$40A8</li>
</ul></li>
</ul></li>
<li>Super Mario Club:
<ul>
<li>Writes $00 to this register.</li>
<li>Later, right after having written $24 to $40A7 and $F8 to $40A6, writes $02 to this register.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40A9</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- Bits exist but function is unknown</code></pre>
<ul>
<li>Bench test observed value $00 with pull-ups.</li>
<li>When driving RF5C66 pin 45 high, this 8-bit value changes.</li>
<li>Pin 45 high causes the value to change continuously, as if counting cycles.</li>
<li>The value does not appear to match $40A7 or $40A6, though further testing is required to say that for sure.</li>
<li>Reading this register causes its contents to be loaded into register $40AA.</li>
</ul></td>
</tr>
<tr>
<td>$40AA</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- Bits exist but function is unknown</code></pre>
<ul>
<li>This register maintains the most recent value that was read from $40A9.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $00 with pull-ups.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40AB</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Reading this register resets the value of $40A9 back to $00.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
<li>JRA-PAT writes $00 to this register.</li>
<li>Super Mario Club writes $00 to this register.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40AC</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Reading this register prevents timed toggle on RF5C66-69 (see notes in pinout).</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test:
<ul>
<li>Observed $00 with pull-downs and $FF with pull-ups (totally open bus).</li>
</ul></li>
<li>JRA-PAT reads this register with a BIT op-code and throws away the result.</li>
<li>Super Mario Club reads this register with a BIT op-code.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40AD</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Mirroring Control
<pre><code>Write
76543210
|+++++++-- (unknown)
+--------- Mirroring (POR value = 0)
             0 = Vertical Mirroring (CIRAM A10 = PPU A10)
             1 = Horizontal Mirroring (CIRAM A10 = PPU A11)
&#10;Read
76543210
|+++++++-- (unlikely to exist)
+--------- Present value of CIRAM A10</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed $00 with pull-downs, $7F with pull-ups.</li>
<li>JRA-PAT writes $00 to this register.</li>
<li>Super Mario Club writes $80 to this register.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40AE</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Built-in RAM /CE control (POR value = 1)
|||||||      1 = Built-in RAM /CE enabled to go low for reads and writes in the range $6000-7FFF.
|||||||          Pin 5C66.28 = 1 at all address ranges.  (This pin normally n/c.)
|||||||      0 = Built-in RAM /CE is always high, preventing all reads and writes of the built-in RAM.
|||||||          Pin 5C66.28 = 0 at all address ranges.  (This pin normally n/c.)
+++++++--- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Refer also to $40C0.0 for built-in RAM enabling.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
<li>JRA-PAT:
<ul>
<li>Writes $00 to this register.</li>
<li>Later writes $01 to this register.</li>
</ul></li>
<li>Super Mario Club:
<ul>
<li>Writes $00 to this register.</li>
<li>Later writes $01 to this register.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40B0</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Kanji Graphic ROM Control
<pre><code>Write
76543210
|||||||+-- Kanji ROM Bank Select (POR value = 0)
+++++++--- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>All reads reset the Kanji auto-increment counter.</li>
<li>D0 is written with 0 or 1 in Kanji graphics-loading code depending on the Kanji character index, changing the Kanji bank.</li>
<li>For an unknown reason, RF5C66 registers can't be written on test setups. They only work in a real Famicom so far.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test found open bus when reading this register.</li>
<li>Is read with BIT and results discarded before reading Kanji data out of $5000-5FFF.</li>
<li>Super Mario Club:
<ul>
<li>Stores X to this register after incrementing X (observed value $20).</li>
<li>Reads this register and throws away the value read.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40B1</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>Modem Control
<pre><code>Write
76543210
|||||||+-- Modem Module pin 29 = $40B1.0, rises slowly, goes low fast (POR value = 1)
||||||+--- Modem Module pin 32 = $40B1.1, rises slowly, goes low fast (POR value = 1)
|||||+---- Modem Module pin 31 = $40B1.2, rises slowly, goes low fast (POR value = 1)
||||+----- Pin 5C66.68: CPU2 Reset on new revision Famicom Network Systems (POR value = 1)
||||         CPU2 /Reset = !($40B1.3)
||||         CPU2 runs when $40B1.3 = 0.
||||         See also $40C0.2.
|||+------ Exp 15 = $40B1.4, rises slowly, goes low fast (POR value = 1)
||+------- Exp 14 = $40B1.5, rises slowly, goes low fast (POR value = 1)
|+-------- Exp 13 = $40B1.6, rises slowly, goes low fast (POR value = 1)
+--------- Exp 12 = $40B1.7, rises slowly, goes low fast (POR value = 1)
&#10;Read
76543210
|||||||+-- Input value of Modem Module pin 29
||||||+--- Input value of Modem Module pin 32
|||||+---- Input value of Modem Module pin 31
||||+----- Input value of 5C66 pin 63 (normally n/c)
|||+------ Input value of EXP 15
||+------- Input value of EXP 14
|+-------- Input value of EXP 13
+--------- Input value of EXP 12</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $FF with pull-downs.</li>
<li>JRA-PAT:
<ul>
<li>Writes $F7 to this register and appears to keep a RAM copy at $17.</li>
<li>Later writes the value from $17, ORed with #$08. (Bit 3 being set to 1.)</li>
<li>Also writes the value from $17, ANDed with #$F7. (Bit 3 being set to 0.)</li>
</ul></li>
<li>Super Mario Club:
<ul>
<li>Writes $F7 to this register and appears to keep a RAM copy at $17.</li>
<li>Later writes the value from $17, ORed with #$08. (Bit 3 being set to 1.)</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40C0</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5C66</td>
<td>CIC Status, CHR Bank, and RAM Control
<pre><code>Write
76543210
|||||||+-- Pin 5C66.35 = $40C0.0 (POR value = 0)
|||||||      RAM +CE Enable (1 = enabled, 0 = disabled)
||||||+--- Pin 5C66.36 = $40C0.1 (POR value = 0)
||||||       (This pin normally n/c)
|||||+---- Pin 5C66.37 = $40C0.2 (POR value = 0)
|||||        Old Revision Famicom Network System: CPU2 /Reset (This pin n/c on newer revisions)
|||||        CPU2 runs on old revisions when $40C0.2 = 1.
|||||        See also $40B1.3.
||||+----- Pin 5C66.38 = $40C0.3 (POR value = 0)
||||         CHR RAM Bank Select (selects between two 8 KiB CHR RAM chips)
++++------ (unknown)
&#10;Read
76543210
|||||||+-- Input value of pin 5C66.31:
|||||||      From FNS CIC (Key) pin 10, or jumpered to GND on models without a CIC chip.
||||||+--- Input value of pin 5C66.32:
||||||       From FNS CIC (Key) pin 15, or jumpered to GND on models without a CIC chip.
|||||+---- Input value of pin 5C66.33:
|||||        CPU2 /Reset fed back in for all models, regardless which model-specific 5C66 pin is sending it out.
||||+----- Input value of pin 5C66.34:
||||         Selected CHR RAM Bank, fed back in from 5C66.38.
|+++------ (unlikely to exist)
+--------- Input value of pin 5C66.29:
             /CPU R/W Inhibit, from FNS CIC (Key) pin 11, or floating up to logic high on models without a CIC chip.
             1 = Writes to card, W-RAM, and $40Dx are allowed
             0 = Blocked (CPU R/W is overriden high, i.e. &quot;read&quot; for these writes)</code></pre>
<ul>
<li>RAM +CE always reflects bit 0 of this register regardless of address space.
<ul>
<li>Refer also to $40AE.0 for built-in RAM enabling.</li>
</ul></li>
<li>All examined software waits for D7 = 1 at initialization.</li>
<li>D7 is controlled by the CIC key chip, if present. If the CIC fails authentication, it latches to 0 until a power cycle.</li>
<li>D1 and D0 always observed low in all cases, regardless if the model does not have CIC key chip, or does have the chip and authentication passes or fails.
<ul>
<li>It is unknown in what scenario that the CIC chip could generate logic high on these pins.</li>
<li>Theories:
<ul>
<li>These pins may reflect detection of different lock CIC chips in the tsuushin card, for different regions or "disk drive mode" or something like that.</li>
<li>Transient role at power-on, indicating when the CIC is ready.</li>
<li>Showing internal error state of the CIC.</li>
</ul></li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bench test observed value $00 with pull-downs, $70 with pull-ups.</li>
<li>JRA-PAT:
<ul>
<li>Writes to this register from what appears to be a RAM copy at $18 ORed with #$01.</li>
<li>Also writes from $18 ANDed with #$FE.</li>
</ul></li>
<li>Super Mario Club:
<ul>
<li>Writes $00 to this register and appears to keep a RAM copy at $18.</li>
<li>Later writes the value from $18, ANDed with #$FB. (Bit 2 being set to 0.)</li>
<li>Reads this register and makes a decision using D7.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D0</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5A18</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 0
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by CPU2 from register $4123.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by CPU2 to register $4123.</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Super Mario Club:
<ul>
<li>Does multiple writes of various values to this register when opening a connection.</li>
<li>Reads this register when closing a modem connection, which reports value $80 and stores it at $701.</li>
</ul></li>
<li>Writing a value to this register does not cause the read value to change.
<ul>
<li>This shows that Famicom -&gt; CPU2 and CPU2 -&gt; Famicom directions each use separate physical registers.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D1</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5A18</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 1
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by CPU2 from register $4124.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by CPU2 to register $4124.</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Super Mario Club:
<ul>
<li>Does multiple writes of various values to this register when opening a connection.</li>
<li>Reads this register when closing a modem connection, which reports value $01 and stores it at $702.</li>
</ul></li>
<li>Writing a value to this register does not cause the read value to change.
<ul>
<li>This shows that Famicom -&gt; CPU2 and CPU2 -&gt; Famicom directions each use separate physical registers.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D2</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5A18</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 2
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by CPU2 from register $4125.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by CPU2 to register $4125.</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Super Mario Club:
<ul>
<li>Does multiple writes of various values to this register when opening a connection.</li>
<li>Reads this register when closing a modem connection, which reports value $00 and stores it at $703.</li>
</ul></li>
<li>Writing a value to this register does not cause the read value to change.
<ul>
<li>This shows that Famicom -&gt; CPU2 and CPU2 -&gt; Famicom directions each use separate physical registers.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D3</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5A18</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Acknowledge
<pre><code>Write
76543210
|||+++++-- (unknown)
+++------- 3-bit value written here can be read by CPU2 from register $4122.
&#10;Read
76543210
|||+++++-- (unlikely to exist)
+++------- 3-bit value read here was written by CPU2 to register $4122.</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>JRA-PAT writes $FF to this register and appears to keep a RAM copy at $19.</li>
<li>Super Mario Club:
<ul>
<li>Writes $FF to this register and appears to keep a RAM copy at $19.</li>
<li>Does a BIT operation on this register and makes decisions based on D7 and D6.
<ul>
<li>Probably the same: Reads this register once per frame, which reports value $E0.</li>
</ul></li>
<li>Also writes value $BF to this register, not in connection with $19.</li>
<li>Does multiple writes of various values to this register when opening and closing a connection.</li>
</ul></li>
<li>Bits 4,3,2,1,0 follow pull-ups and downs on the data bus.</li>
<li>Writing a value to this register does not cause the read value to change.
<ul>
<li>This shows that Famicom -&gt; CPU2 and CPU2 -&gt; Famicom directions each use separate physical registers.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D4</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>RF5A18</td>
<td>Tone Rx and Expansion I/O Control
<pre><code>Write
76543210
|||||||+-- RF5A18 Pin 65 (Exp 17) = $40D4.0 (open-drain)
||||||+--- RF5A18 Pin 67 (Exp 19) = $40D4.1 (push-pull)
|||||+---- RF5A18 Pin 66 (Exp 18) = $40D4.2 (open-drain)
+++++----- (unknown)
&#10;Read
76543210
|||||||+-- $40D4.0 = I/O value of RF5A18 Pin 65 (Exp 17)
||||||+--- $40D4.1 = output value of RF5A18 Pin 67 (Exp 19)
|||||+---- $40D4.2 = I/O value of RF5A18 Pin 66 (Exp 18)
||||+----- $40D4.3 = input value of RF5A18 Pin 69 (Tone Rx D1)
|||+------ $40D4.4 = input value of RF5A18 Pin 70 (Tone Rx D2)
||+------- $40D4.5 = input value of RF5A18 Pin 71 (Tone Rx D4)
|+-------- $40D4.6 = input value of RF5A18 Pin 72 (Tone Rx D8)
+--------- $40D4.7 = input value of RF5A18 Pin 73 (Tone Rx DV)</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>JRA-PAT writes $FF to this register and appears to keep a RAM copy at $1A.</li>
<li>Super Mario Club writes $FF to this register and appears to keep a RAM copy at $1A.</li>
<li>Bits 7,6,5,4,3 read back as always 0 when writing inverting values to these bits.</li>
<li>Bits 2,1,0 read back with the same value that was written.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D5</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5A18</td>
<td>Clocks
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
||++++++-- Bits exist but function is unknown
++-------- (unlikely to exist)</code></pre>
<ul>
<li>The purpose is unknown but it could potentially serve to generate random numbers.
<ul>
<li>The Famicom CPU and CPU2 have separate clocks and are asynchronous, making the value read especially random.</li>
</ul></li>
<li>Bit 5 is a 0.05% duty square wave, measured 1.200kHz (19.6608MHz CPU2 crystal / 16384)
<ul>
<li>Negative pulse width measures 406 nsec, the same as period of bit 2.</li>
</ul></li>
<li>Bit 4 is a 50% duty square wave, depending on $4114.1 and $4114.0:
<ul>
<li>(0,0) measured 19.20kHz (19.6608MHz CPU2 crystal / 1024)</li>
<li>(0,1) measured 38.40kHz (19.6608MHz CPU2 crystal / 512)</li>
<li>(1,0) measured 76.80kHz (19.6608MHz CPU2 crystal / 256)</li>
<li>(1,1) measured 153.6kHz (19.6608MHz CPU2 crystal / 128)</li>
</ul></li>
<li>Bit 3 is a 50% duty square wave, measured 307.3kHz (19.6608MHz CPU2 crystal / 64)</li>
<li>Bit 2 is a 50% duty square wave, measured 2.451MHz (19.6608MHz CPU2 crystal / 8)</li>
<li>Bit 1 is a 50% duty square wave, measured 4.902MHz (19.6608MHz CPU2 crystal / 4)
<ul>
<li>This bit toggles at the same rate as RF5A18 pin 42, though there is a 129.9 degree phase difference.</li>
</ul></li>
<li>Bit 0 is a 50% duty square wave, measured 9.804MHz (19.6608MHz CPU2 crystal / 2)</li>
<li>The phasing of these signals do not line up nicely but they do stay in perfect sync with CPU2 M2.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>$40Dx registers can be monitored in real-time by statically giving the RF5A18 the bus signals for a read from this register.
<ul>
<li>M2 = 1, Card R/W (i.e. buffered CPU R/W) = 1, Famicom /CE = 0, CPU A0,A1,A2 = x,x,x</li>
<li>Unimplemented bits will latch the Famicom data bus's last value when toggling M2.</li>
</ul></li>
<li>Bits 7,6 follow pull-ups and downs on the data bus.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$40D6</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>RF5A18</td>
<td>UART Status
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- Bit exists but function is unknown
||||||+--- $40D6.1 = !($4112.0) read by CPU2.  (0 = $4110 UART Rx buffer has a byte ready to be read.)
|||||+---- $40D6.2 = !($4112.1) read by CPU2.  (0 = $4110 UART Tx buffer can be written.)
||||+----- $40D6.3 = !($4113.1 OR $4113.2) written by CPU2.
|||+------ $40D6.4 = !($4113.7) written by CPU2.  (1 = Transmit repeat.)
||+------- $40D6.5 = !($4113.6) written by CPU2.
++-------- (unlikely to exist)</code></pre>
<p>$40D6.2 previously observed to equal !($4113.7 AND $4112.1) written by CPU2. It was later discovered that this bit also follows the $4112.1 read value (Tx buffer ready), which makes a lot more sense being beside $40D6.1 (Rx buffer ready). The previous observation can be explained because writing to those CPU2 bits both likely changed the availability of the Tx buffer.</p>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Bits 7,6 follow pull-ups and downs on the data bus.</li>
<li>Bit 0 only ever observed with value 1 regardless of various techniques writing to lots of CPU2 registers.</li>
<li>Writing a value to this register does not <em>seem to</em> cause its own read value to change.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table>

# RF5A18 Internal 65C02 CPU {#RF5A18_Internal_65C02_CPU}

The RF5A18 contains its own CPU, termed "CPU2" on this page. It is a 65C02 supporting bitwise set/clear/branch instructions. Note that CPU2 has its own parallel execution with its own address and data busses that are not available to the Famicom's CPU. CPU2 also has its own clock source, so it does not execute synchronously with the Famicom CPU. CPU2 clock speed is 2.4576 MHz (19.6608 MHz crystal / 8). This section describes CPU2's own memory mapping and its own internal registers.

CPU2 /Reset is controlled 2 different ways depending on the revision of Famicom Network System. To support all revisions when enabling CPU2, both of these bits should be written:

- \$40B1.3 = 0 (new revisions Famicom Network System, and old revisions with J1 closed)
- \$40C0.2 = 1 (old revisions Famicom Network System with J2 closed)

## CPU2 Memory Map {#CPU2_Memory_Map}

### <span id="RF5A18_Pin_26_Low_.28default.29"></span>RF5A18 Pin 26 Low (default) {#RF5A18_Pin_26_Low_.28default.29}

    +================+ $0000
    | CPU2 RAM       |
    | (U6)           |
    +================+ $2000
    |   (Returns     |
    |   last fetch)  |
    +================+ $4100
    | CPU2 Control   |
    | Registers      |
    +================+ $4140
    |   (Returns     |
    |   last fetch)  |
    +================+ $C000
    | External ROM   |
    | (not used)     |
    +================+ $E000
    | RF5A18         |
    | Internal ROM   |
    +================+ $10000

### RF5A18 Pin 26 High {#RF5A18_Pin_26_High}

    +================+ $0000
    | CPU2 RAM       |
    | (U6)           |
    +================+ $2000
    |   (Returns     |
    |   last fetch)  |
    +================+ $4100
    | CPU2 Control   |
    | Registers      |
    +================+ $4140
    |   (Returns     |
    |   last fetch)  |
    +================+ $C000
    |                |
    |                |
    | External ROM   |
    |                |
    |                |
    +================+ $10000

## CPU2 Known Registers {#CPU2_Known_Registers}

<table class="wikitable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>Address</th>
<th style="text-align: center;">Read<br />
Has<br />
Effect</th>
<th style="text-align: center;">Read<br />
Has<br />
Data</th>
<th style="text-align: center;">Write</th>
<th>Function</th>
</tr>
</thead>
<tbody>
<tr>
<td>$4100</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>NMI Timer 1 Period, low byte.
<pre><code>Write
76543210
++++++++-- NMI timer 1 period low byte.
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>Time (seconds) * 1200 = 16-bit timer 1 register value.
<ul>
<li>1 count = 2048 CPU cycles.</li>
</ul></li>
<li>This is the LSB of the 16-bit period. Refer to $4101 for the MSB.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>When tone dialing, the digit length is $0078 (~100ms) and the inter-digit gap is $0046 (~60ms).</li>
<li>Slow pulse dialing uses on-hook pulses of $004f (~67ms) and off-hook gaps of $0027 (~33ms) and the inter-digit gap is $031e (~666ms).</li>
<li>Fast pulse dialing uses on-hook pulses of $0027 (~33ms) and off-hook gaps of $0013 (~17ms) and the inter-digit gap is $027e (~533ms).
<ul>
<li>The times are approximate and assume a 1200Hz clock source, but the actual clock source is unknown.
<ul>
<li>Theory: 19.6608MHz RF5A18 crystal / 2^14 = 1200Hz.</li>
</ul></li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4101</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>NMI Timer 1 Period, high byte.
<pre><code>Write
76543210
++++++++-- NMI timer 1 period high byte.
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>Time (seconds) * 1200 = 16-bit timer 1 register value.
<ul>
<li>1 count = 2048 CPU cycles.</li>
</ul></li>
<li>This is the MSB of the 16-bit period. Refer to $4100 for the LSB.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4102</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>NMI Timer 1 Restart.
<pre><code>Write
76543210
|||||||+-- $4102.0 = timer 1 loop.  1 = Loop, 0 = One-Shot
||||||+--- $4102.1 = timer 1 restart.  1 = Restart.
++++++---- (unknown)
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>When setting bit 1 = 1, timer 1 restarts with the period specified in $4100, $4101.</li>
<li>The values of $4100, $4101 are not affected by running timer 1 and need not be rewritten each time.</li>
<li>NMI will be generated when timer 1 expires if $412F.0 = 1.</li>
<li>The behavior is untested when setting bits 0 or 1 = 1 while timer 1 is already running.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $E136 writes value #$02 to this address.</li>
<li>ROM $F3D6 writes value #$00 to this address (initialization).</li>
<li>ROM $F840 writes value #$00 to this address (Famicom command $61).</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4103</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>NMI Timer 1 Acknowledge.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- $4103.0 = Timer 1 NMI flag.  1 = Timer 1 NMI triggered.
+++++++--- (does not exist)</code></pre>
<ul>
<li>Reading from this register likely acknowledges timer 1 NMI.</li>
<li>The read value of $4103.0 reflects the timer 1 NMI flag, but does not reflect external NMI via pin 29.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $40. ($41 XOR $40 = $01 existence mask)</li>
<li>Value $40 from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $BE. ($BF XOR $BE = $01 existence mask)</li>
<li>$01 OR $01 = $01 accumulated existence mask.</li>
</ul></li>
<li>ROM $E825 reads from this register at the beginning of the NMI handler and throws away the value.</li>
<li>ROM $F3D9 reads from this register during initialization and throws away the value.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4104</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>IRQ Timer 2 Period, low byte.
<pre><code>Write
76543210
++++++++-- IRQ timer 2 period low byte.
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>Time (seconds) * 2,457,600 = 16-bit timer 2 register value.
<ul>
<li>1 count = 1 CPU cycle.</li>
</ul></li>
<li>This is the LSB of the 16-bit period. Refer to $4105 for the MSB</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $F415 writes value #$FD to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4105</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>IRQ Timer 2 Period, high byte.
<pre><code>Write
76543210
++++++++-- IRQ Timer 2 Period, high byte.
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>Time (seconds) * 2,457,600 = 16-bit timer 2 register value.
<ul>
<li>1 count = 1 CPU cycle.</li>
</ul></li>
<li>This is the MSB of the 16-bit period. Refer to $4104 for the LSB</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $F410 writes value #$2F to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4106</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>IRQ Timer 2 Restart.
<pre><code>Write
76543210
|||||||+-- $4106.0 = timer 2 loop.  1 = Loop, 0 = One-Shot
||||||+--- $4106.1 = timer 2 restart.  1 = Restart.
++++++---- (unknown)
&#10;Read
76543210
++++++++-- (does not exist)</code></pre>
<ul>
<li>Theory: When setting bit 1 = 1, timer 2 restarts with the period specified in $4100, $4101.</li>
<li>Theory: The values of $4104, $4105 are not affected by running timer 2 and need not be rewritten each time.</li>
<li>IRQ will be generated when timer 2 expires if $412F.6 = 1.</li>
<li>The behavior is untested when setting bits 0 or 1 = 1 while timer 2 is already running.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $F3DE writes value #$00 to this address.</li>
<li>ROM $F41A writes value #$03 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4107</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>IRQ Timer 2 Acknowledge.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- (unlikely to exist)
||||||+--- Likely to be the Timer 2 IRQ flag
++++++---- (unlikely to exist)</code></pre>
<ul>
<li>Reading from this register likely acknowledges timer 2 IRQ.</li>
<li>$4107.1 is probably the timer 2 IRQ flag.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 and $40 from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $BD. ($BF XOR $BD = $02 existence mask)</li>
</ul></li>
<li>ROM $E78A reads from this register directly after seeing IRQ flag $412F.6 and throws away the value.</li>
<li>ROM $F3E1 reads from this register during initialization and throws away the value.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4110</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>UART Rx/Tx Data Buffer
<pre><code>Write
76543210
++++++++-- UART Tx Data Byte
             Bits 7:0 in 8-bit mode, bits 6:0 in 7-bit mode
&#10;Read
76543210
++++++++-- UART Rx Data Byte
             Bits 7:0 in 8-bit mode, bits 6:0 in 7-bit mode</code></pre>
<ul>
<li>Writing a byte to this register automatically clears $4112.1 read value and triggers the byte to be sent on UART Tx.</li>
<li>Reading a byte from this register automatically clears $4112.0 read value, to be triggered high again next byte received on UART Rx.</li>
<li>This UART is connected to the MSM6827L TXD and RXD pins.</li>
<li>The byte is ready to read here when the IRQ flag $4112.0 is high.</li>
<li>The byte is ready to be written here when flag $4112.1 is high.</li>
<li>Writes to $4110 are specifically prevented by built-in ROM in modes 0 and 5.</li>
<li>Bench testing confirmed reads and writes of this register correspond to data read by UART Rx and written to UART Tx.
<ul>
<li>In 7-bit mode, Rx bit 7 (unused in this case) observed to be 0 but not confirmed always to be 0.</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $41 from previous fetch also observed changed to $00. ($41 XOR $00 = $41 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $00. ($40 XOR $00 = $40 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $00. ($BF XOR $00 = $BF existence mask)</li>
<li>$BE OR $41 OR $40 OR $BF = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $E69B reads from this address inside of an IRQ handler, seemingly caused by its own IRQ, with flag at $4112.0.</li>
<li>ROM $E72F writes to this address inside of an IRQ handler, which may have been caused by timer 2.</li>
<li>ROM $E754 writes to this address inside of an IRQ handler, which may have been caused by timer 2.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4111</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>UART Configuration
<pre><code>Write
76543210
|||||||+-- $4111.0 = UART Rx Enable (1 = enabled)
||||||+--- $4111.1 = UART Tx Enable (1 = enabled)
|||||+---- $4111.2 = Baudrate Scaler (0 = multiply baudrate x4)
||||+----- $4111.3 = Data bits (0 = 7 bits, 1 = 8 bits)
|||+------ $4111.4 = Tx Stop bits (0 = 1 stop bit, 1 = 2 stop bits)
||+------- $4111.5 = Parity bit (0 = don&#39;t use, 1 = append parity bit)
|+-------- $4111.6 = Parity type (0 = odd parity, 1 = even parity)
+--------- $4111.7 = Send Tx Break (1 = force UART Tx low (break), 0 = normal)
&#10;Read
76543210
++++++++-- Reads back the value written.</code></pre>
<ul>
<li>See register $4114 for baudrate table.</li>
<li>Confirmed write bits 2,3,5,6 apply to both Rx and Tx.</li>
<li>Write bit $4111.4 confirmed not to enforce 2 stop bits for UART Rx.
<ul>
<li>Consecutive Rx bytes with 1 stop bit separation can still be received when $4111.4 = 1.</li>
</ul></li>
<li>Write bit $4111.7 requires manual operation to send a break character and has no effect on Rx operation.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $00. ($41 XOR $00 = $41 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $00. ($40 XOR $00 = $40 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $00. ($BF XOR $00 = $BF existence mask)</li>
<li>$41 OR $40 OR $BF = $FF accumulated existence mask.</li>
</ul></li>
<li>The built-in ROM function at $E1BC determines the UART settings to be written to this register.
<ul>
<li>Clears bits 0,1,7.</li>
<li>Sets bit 2 (slow baudrate scaler).</li>
<li>Figures out bits 3,4,5,6 (UART bits/parity settings).</li>
<li>Returns that $4111 value in accumulator A.</li>
<li>Also sets baudrate to 1200 bits/sec by writing to $4114 before returning.</li>
</ul></li>
<li>ROM $E151 runs $E1BC, sets bits [1,0] = 0,0, and writes it to this register.</li>
<li>ROM $ECE3 runs $E1BC, sets bits [1,0] = 1,1, and writes it to this register.</li>
<li>ROM $ED20 runs $E1BC, sets bits [1,0] = 0,1, and writes it to this register.</li>
<li>ROM $ED4E reads from this address, OR's it with #$03 (i.e. sets bits [1,0] = 1,1), and writes it back to this register.</li>
<li>ROM $EE83 runs $E1BC, sets bits [1,0] = 0,1, and writes it to this register.</li>
<li>ROM $F0E8 reads from this address, AND's it with #$FE (i.e. clears bit 7), and writes it back to this register.</li>
<li>ROM $F2DC runs $E1BC, sets bits [1,0] = 0,1, and writes it to this register.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4112</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>UART Status
<pre><code>Write
76543210
|||||||+-- Bit exists but function is unknown
||||||+--- $4112.1: Tx Silence
||||||       0 = Set UART Tx directly high (idle state), 1 = Allow sending data.
||||||       This bit figures into $40D6.2 read by the Famicom. (See register $40D6.)
|+++++---- (unknown)
+--------- $4112.7 = $4110 UART Rx Data IRQ Acknowledge, 1 = Acknowledge IRQ.
&#10;Read
76543210
|||||||+-- $4112.0 = $4110 UART Rx Buffer Ready Flag (1 = $4110 buffer has a byte ready to be read.)
||||||+--- $4112.1 = $4110 UART Tx Buffer Ready Flag (1 = $4110 buffer can be written.)
|||||+---- $4112.2 = UART Tx Idle, 1 = Idle, 0 = Active.
||||+----- Bit exists but function is unknown
|||+------ $4112.4 = Rx Parity Error. (1 = error detected)
||+------- $4112.5 = Rx Framing Error. (1 = error detected)
|+-------- $4112.6 = Rx Break Received. (1 = break detected)
+--------- Bit exists but function is unknown</code></pre>
<ul>
<li>Error and Break bits persist through $4110 operations, but can be cleared by writing 1 to $4112.7.</li>
<li>Found $4112.4 read became 1 when setting in parity mode ($4111.5 = 1) and injecting a wrong parity bit.
<ul>
<li>Confirmed this flag not getting set anymore with the same test when changing parity even/odd (i.e. inverted $4111.6).</li>
</ul></li>
<li>$4112.6 = 1 (break) comes accompanied with $4112.5 = 1 (framing error).</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $04. ($41 XOR $04 = $45 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $04. ($40 XOR $04 = $44 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $04. ($BF XOR $04 = $BB existence mask)</li>
<li>$45 OR $44 OR $BB = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM implements a shadow read register at $0011.</li>
<li>ROM implements a shadow write register at $0012.</li>
<li>ROM $E159 writes value #$80 to this address when disconnecting.</li>
<li>ROM $E4D8 reads from this address in the IRQ handler and uses bit 0.</li>
<li>ROM $E4F9 writes to this address in the IRQ handler after having set bit 7 in its shadow register.
<ul>
<li>If $4112.0 read = 1, it writes a 1 to $4112.7.</li>
</ul></li>
<li>ROM $ECDB writes value #$81 to this address.</li>
<li>ROM $ED18 writes value #$81 to this address.</li>
<li>ROM $ED4B writes value #$81 to this address.</li>
<li>ROM $EE7B writes value #$81 to this address.</li>
<li>ROM $F2D4 writes value #$81 to this address.</li>
</ul>
<ul>
<li>ROM $E4F2 uses read bit 0.</li>
<li>ROM $E67C uses read bit 5.</li>
<li>ROM $E68E uses read bit 4.</li>
<li>ROM $E703 uses read bit 1.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4113</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>UART Configuration
<pre><code>Write
76543210
|||||||+-- (unknown)
||||||+--- $4113.1 figures into $40D6.3 read by the Famicom. (See register $40D6.)
|||||+---- $4113.2 figures into $40D6.3 read by the Famicom. (See register $40D6.)
||+++----- (unknown)
|+-------- $4113.6 figures into $40D6.5 read by the Famicom. (See register $40D6.)
+--------- $4113.7: Tx Repeat
             1 = Send once on pin 90 (UART Tx), 0 = Send repeatedly
             This figures into $40D6.2 and $40D6.4 read by the Famicom. (See register $40D6.)
&#10;Read
76543210
++++++++-- Bits exist but function is unknown</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Writing to Bit 4 appeared to have an effect in some tests, apparently sending only 1 bit on the UART Tx, but this was not able to be reproduced in later tests.</li>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $10. ($41 XOR $10 = $51 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $10. ($40 XOR $10 = $50 existence mask)</li>
<li>Value $04 from previous fetch observed changed to $10. ($04 XOR $10 = $14 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $10. ($BF XOR $10 = $AF existence mask)</li>
<li>$51 OR $50 OR $14 OR $AF = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM implements a shadow write register at $0013.</li>
<li>ROM $E161 writes value #$00 to this address.</li>
<li>ROM $ECD3 writes value #$C0 to this address.</li>
<li>ROM $ED11 writes value #$40 to this address.</li>
<li>ROM $ED44 writes value #$C0 to this address.</li>
<li>ROM $EE73 writes value #$40 to this address.</li>
<li>ROM $F2CD writes value #$40 to this address.</li>
<li>Bits 7 and 6 differ in these writes, suggesting that they exist.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4114</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>UART Baudrate
<pre><code>Write
76543210
||||||++-- UART Baud Rate (Rx and Tx)
||||||       This is a clock divider that also affects frequency seen on $40D5.4.
++++++---- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Baud Rate Selection:</li>
</ul>
<table class="wikitable">
<thead>
<tr>
<th>$4111.2</th>
<th>$4114.1</th>
<th>$4114.0</th>
<th>Baudrate</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td>0</td>
<td>0</td>
<td>1200</td>
</tr>
<tr>
<td>0</td>
<td>0</td>
<td>1</td>
<td>2400</td>
</tr>
<tr>
<td>0</td>
<td>1</td>
<td>0</td>
<td>4800</td>
</tr>
<tr>
<td>0</td>
<td>1</td>
<td>1</td>
<td>9600</td>
</tr>
<tr>
<td>1</td>
<td>0</td>
<td>0</td>
<td>300</td>
</tr>
<tr>
<td>1</td>
<td>0</td>
<td>1</td>
<td>600</td>
</tr>
<tr>
<td>1</td>
<td>1</td>
<td>0</td>
<td>1200</td>
</tr>
<tr>
<td>1</td>
<td>1</td>
<td>1</td>
<td>2400</td>
</tr>
</tbody>
</table>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>$40D5.4 shows these frequencies depending on $4114 bits 1 and 0:
<ul>
<li>(0,0) measured 19.20kHz (19.6608MHz CPU2 crystal / 1024)</li>
<li>(0,1) measured 38.40kHz (19.6608MHz CPU2 crystal / 512)</li>
<li>(1,0) measured 76.80kHz (19.6608MHz CPU2 crystal / 256)</li>
<li>(1,1) measured 153.6kHz (19.6608MHz CPU2 crystal / 128)</li>
</ul></li>
<li>Note: $40D5.4 is not affected by $4111.2.</li>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, $FF, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $E1DE writes value #$02 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4120</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Pin 39 (MSM6827L AD0) = $4120.0
||||||+--- Pin 38 (MSM6827L AD1) = $4120.1
|||||+---- Pin 37 (n/c) = $4120.2
||||+----- Pin 36 (MSM6827L /RD) = $4120.3
|||+------ Pin 35 (MSM6827L /WR) = $4120.4
||+------- Pin 34 (MSM6827L EXCLK) = $4120.5
|+-------- Pin 32 (MSM6827L Data): Direction = $4120.6: 1 = input, 0 = output (refer to $4121.0)
+--------- (unknown)
&#10;Read
76543210
|+++++++-- Bits exist but function is unknown
+--------- $4120.7 = Pin 33 (MSM6827L /INT)</code></pre>
<ul>
<li>When pin 32 is set as input, it is unknown if or how the value can be read.</li>
<li>Pin 32 floats in input mode, regardless of the value written to $4121.0.</li>
<li>When pin 32 is set as output, it is a push-pull output.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>Sequence:
<ul>
<li>ROM $E3BC writes value #$38 to this address.</li>
<li>ROM $E3C7 writes value #$38 &amp; #$DF = #$18 to this address.</li>
<li>ROM $E3CC writes value #$18 | #$20 = #$38 to this address.</li>
<li>ROM $E3D4 writes value #$38 &amp; #$DF = #$18 to this address.</li>
<li>ROM $E3D9 writes value #$18 &amp; #$EF = #$08 to this address.</li>
<li>ROM $E3DE writes value #$08 | #$10 = #$18 to this address.</li>
<li>ROM $E3E3 writes value #$58 to this address.</li>
</ul></li>
<li>ROM $E3E9 writes value #$5A to this address.</li>
<li>ROM $E3EE writes value #$52 to this address.</li>
<li>ROM $E3F5 writes value #$72 to this address.</li>
<li>ROM $E3FA writes value #$52 to this address.</li>
<li>ROM $E402 writes value #$72 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4121</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Pin 32 (MSM6827L Data) = $4121.0 when set as output (refer to $4120.6)
+++++++--- (unknown)
&#10;Read
76543210
++++++++-- Bits exist but function is unknown</code></pre>
<ul>
<li>Reading $4121.0 reflects the value written to $4121.0 when $4120.6 is set as output.</li>
<li>Surprisingly, reading $4121.0 is always 1 when $4120.6 is set as input regardless of:
<ul>
<li>Driving pin 32 high or low.</li>
<li>Writing 1 or 0 to $4121.0.</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $E3BF writes to this address.</li>
<li>ROM $E405 reads this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4122</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Acknowledge
<pre><code>Write
76543210
|||+++++-- (unlikely to exist)
+++------- 3-bit value written here can be read by Famicom CPU from register $40D3.
&#10;Read
76543210
|||+++++-- (unlikely to exist)
+++------- 3-bit value read here was written by Famicom CPU to register $40D3.</code></pre>
<ul>
<li>Read and write directions are represented by separate physical registers.</li>
<li>The read value can only be affected by the Famicom CPU.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $E1. ($41 XOR $E1 = $A0 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $E0. ($40 XOR $E0 = $A0 existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$A0 OR $A0 OR $40 = $E0 accumulated existence mask.</li>
</ul></li>
<li>ROM $F3E6 writes value #$FF to this address.</li>
<li>ROM $FA70 writes value #$E0 to this address.</li>
<li>ROM $FA78 writes value #$E0 to this address.</li>
<li>ROM $FA89 reads this address using BIT operation and uses BMI, indicating bit 7 used.</li>
<li>ROM $FAA7 reads this address using BIT operation and uses BPL, indicating bit 7 used.</li>
<li>ROM $FAB4 reads this address and checks if bits 7 and 6 both = 1.</li>
<li>ROM $FABF writes #$40 to this address, reads it right back, and checks if bits 7 and 6 both = 1.</li>
<li>ROM $FACF writes value #$20 to this address.</li>
<li>ROM $FAF3 writes value #$A0 to this address.</li>
<li>ROM $FB00 writes value #$20 to this address.</li>
<li>ROM $FB0D writes value #$A0 to this address.</li>
<li>ROM $FB15 writes value #$E0 to this address.</li>
<li>ROM $FB1D writes value #$20 to this address.</li>
<li>ROM $FB25 writes value #$E0 to this address.</li>
<li>ROM $FB8F writes value #$E0 to this address.</li>
<li>ROM $FC02 writes value #$E0 to this address.</li>
<li>ROM $FC13 reads this address using BIT operation and uses BMI, indicating bit 7 used.</li>
<li>ROM $FC2D reads this address using BIT operation and uses BPL, indicating bit 7 used.</li>
<li>ROM $FC3A reads this address and checks if bits 7 and 6 both = 1.</li>
<li>ROM $FC43 writes value #$60 to this address.</li>
<li>ROM $FC69 writes value #$E0 to this address.</li>
<li>ROM $FC76 writes value #$60 to this address.</li>
<li>ROM $FC83 writes value #$E0 to this address.</li>
<li>ROM $FC8D reads this address using BIT operation and uses BPL and BVC, indicating bits 7 and 6 are used.</li>
<li>ROM $FC96 writes value #$E0 to this address.</li>
<li>ROM $FCA4 writes value #$60 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4123</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 0
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by Famicom CPU from register $40D0.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by Famicom CPU to register $40D0.</code></pre>
<ul>
<li>Read and write directions are represented by separate physical registers.</li>
<li>The read value can only be affected by the Famicom CPU.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $F3E9 writes value #$FF to this address.</li>
<li>ROM $FAD9 writes to this address.</li>
<li>ROM $FB32 writes to this address.</li>
<li>ROM $FC49 reads from this address.</li>
<li>ROM $FCAE reads from this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4124</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 1
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by Famicom CPU from register $40D1.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by Famicom CPU to register $40D1.</code></pre>
<ul>
<li>Read and write directions are represented by separate physical registers.</li>
<li>The read value can only be affected by the Famicom CPU.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $F3EC writes value #$FF to this address.</li>
<li>ROM $FADF writes to this address.</li>
<li>ROM $FB38 writes to this address.</li>
<li>ROM $FC55 reads from this address and checks if it equals $00, indicating all bits exist.</li>
<li>ROM $FCB5 reads from this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4125</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Famicom CPU &lt;-&gt; CPU2 Interface, Data Byte 2
<pre><code>Write
76543210
++++++++-- 8-bit value written here can be read by Famicom CPU from register $40D2.
&#10;Read
76543210
++++++++-- 8-bit value read here was written by Famicom CPU to register $40D2.</code></pre>
<ul>
<li>Read and write directions are represented by separate physical registers.</li>
<li>The read value can only be affected by the Famicom CPU.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $F3EF writes value #$FF to this address.</li>
<li>ROM $FAE7 writes to this address.</li>
<li>ROM $FB3E writes to this address.</li>
<li>ROM $FC4F reads from this address.</li>
<li>ROM $FCBC reads from this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4126</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- $4126.0 = Pin 47 (+5V)
||||||+--- $4126.1 = !(Pin 48) (+5V)
|||||+---- $4126.2 = Pin 49 (from 5C66-69)
||||+----- $4126.3 = Pin 51 (Switch SW1-2)
|||+------ $4126.4 = Pin 52 (Switch SW1-4)
||+------- $4126.5 = Pin 53 (Modem P4-25)
|+-------- $4126.6 = Pin 54 (Modem P4-28)
+--------- $4126.7 = Pin 55 (Modem P4-23)
</code></pre>
<ul>
<li>SW1 selects between pulse and DTMF dialing.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FB. ($41 XOR $FB = $BA existence mask)</li>
<li>Value $41 from previous fetch also observed changed to $FF. ($41 XOR $FF = $BE existence mask)</li>
<li>Value $40 from previous fetch observed changed to $FF. ($40 XOR $FF = $BF existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$BA OR $BE OR $BF OR $40 = $FF accumulated existence mask.</li>
</ul></li>
<li>ROM $E482 reads this address and checks Bit 2.</li>
<li>ROM $E7BC reads this address and checks Bits 7 and 6.</li>
<li>ROM $E886 reads this address and checks Bits 7 and 6.</li>
<li>ROM $E8DC reads this address.</li>
<li>ROM $E8FC reads this address and checks Bits 7 and 6.</li>
<li>ROM $E9C0 reads this address.</li>
<li>ROM $EBAB reads this address and checks Bits 7 and 6.</li>
<li>ROM $F1B6 reads this address.</li>
<li>ROM $F4F2 reads this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4127</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Pin 56 MSM6827L +Reset = $4127.0
||||||+--- Pin 57 /Red LED = $4127.1
|||||+---- Pin 58 /Green LED = $4127.2
||||+----- Pin 59 (n/c) = $4127.3
|||+------ Pin 60 /Phone Off Hook = $4127.4
||+------- Pin 61 /DTMF Output Enable = $4127.5
|+-------- Pin 62 /Phone Audio Enable = $4127.6
+--------- Pin 63 (Modem P4-19) = $4127.7
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Pin 60 observed low corresponding to off hook and pulse dialing pulses.</li>
<li>Pin 61 observed low for the duration of DTMF numbers being dialed, but not in pulse mode.</li>
<li>Pin 62 observed low corresponding to modem sounds coming through the TV speakers.</li>
<li>Pin 63 observed always high.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, $FF, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM implements a shadow write register at $0014.</li>
<li>ROM $E237 writes the shadow register $0014 to this register, modified bit 6 (/Phone Audio Enable).</li>
<li>ROM $E24D writes the shadow register $0014 to this register, modified bit 7 (Modem P4-19).</li>
<li>ROM $E263 writes the shadow register $0014 to this register, modified bit 4 (/Phone Off Hook).</li>
<li>ROM $E27D writes the shadow register $0014 to this register, modified bit 1 (/Red LED).</li>
<li>ROM $E29A writes the shadow register $0014 to this register, modified bit 2 (/Green LED).</li>
<li>ROM $E3B0 writes the shadow register $0014 to this register, modified bit 5 (/DTMF Output Enable).</li>
<li>ROM $E987 writes the shadow register $0014 to this register, modified bit 4 (/Phone Off Hook).</li>
<li>ROM $EA83 writes the shadow register $0014 to this register, modified bit 4 (/Phone Off Hook).</li>
<li>ROM $F3BF writes value #$FF directly to this register, then clears bit 0 with value #$FE via the shadow register $0014.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4128</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Pin 68 (Tone Rx GT) = $4128.0
+++++++--- (unknown)
&#10;Read
76543210
|||||+++-- (unlikely to exist)
||||+----- $4128.3 = Pin 69 (Tone Rx D1)
|||+------ $4128.4 = Pin 70 (Tone Rx D2)
||+------- $4128.5 = Pin 71 (Tone Rx D4)
|+-------- $4128.6 = Pin 72 (Tone Rx D8)
+--------- $4128.7 = Pin 73 (Tone Rx DV)</code></pre>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $01. ($41 XOR $01 = $40 existence mask)</li>
<li>Value $41 from previous fetch also observed changed to $F9. ($41 XOR $F9 = $B8 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $F8. ($40 XOR $F8 = $B8 existence mask)</li>
<li>Value $FF from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $FF. ($BF XOR $FF = $40 existence mask)</li>
<li>$40 OR $B8 OR $B8 OR $40 = $F8 accumulated existence mask.</li>
</ul></li>
<li>ROM $E599 reads from this address.</li>
<li>ROM $E5BB reads from this address and uses bits 3,4,5,6.</li>
<li>ROM $F3F7 writes value #$00 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4129</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>P5 Expansion Port
<pre><code>Write
76543210
||||++++-- Data nybble written to device attached to P5 connector.
||++------ Used for sequencing writes to the device.
++-------- (unknown)
&#10;Read
76543210
||++++++-- Controlled by device attached to P5 connector, though the ROM code never reads it.
++-------- Not used</code></pre>
<ul>
<li>The data bus floats when reading this register, presumably to be driven by a device connected to P5.</li>
<li>Pin 23 is a /CE low when writing to this register (untested for reads).</li>
<li>Data bits 6 and 7 are not available in the P5 connector.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $FF with pull-ups on data bus.</li>
<li>Value $41 from previous fetch observed changed to $00 with pull-downs on data bus.</li>
<li>Value $41 from previous fetch observed changed to $BF, when data bus configured with pull-ups and -downs to make $BF.</li>
<li>Value $40 from previous fetch observed changed to $BF, when data bus configured with pull-ups and -downs to make $BF.</li>
<li>Value $FF from previous fetch observed changed to $BF, when data bus configured with pull-ups and -downs to make $BF.</li>
</ul></li>
<li>Sequence:
<ul>
<li>ROM $E4B6 writes a value to this address.</li>
<li>ROM $E4BB writes previous value again, OR'd with #$10.</li>
<li>ROM $E4C0 writes previous value again, AND'd with #$2F.</li>
</ul></li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$412F</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- $412F.0 = Timer 1 NMI enable.  1 = Enabled.
||||+++--- (unknown)
|||+------ $412F.4 = Pin 33 (MSM6827L /INT) IRQ enable. 1 = Enabled.
||+------- $412F.5 = Pin 73 (Tone Rx DV) IRQ enable.  1 = Enabled.
|+-------- $412F.6 = Timer 2 IRQ enable.  1 = Enabled.
+--------- $412F.7 = UART Rx IRQ enable.  1 = Enabled.
&#10;Read
76543210
|||||||+-- Bit exists but function is unknown
|||||++--- (unlikely to exist)
||||+----- Bit exists but function is unknown
|||+------ $412F.4 = !(Pin 33) (MSM6827L /INT)
||+------- $412F.5 is set when IRQ is triggered by writing 1 to $412F.5.
||           This is probably the Tone Rx DV interrupt flag.
|+-------- $412F.6 is an IRQ flag.  1 = IRQ pending.
+--------- Bit exists but function is unknown</code></pre>
<ul>
<li>ROM tends to clear $412F.0 immediately after disabling interrupts with SEI, and sets it right before enabling with CLI.
<ul>
<li>The ROM is apparently disabling/enabling both IRQs and the timer NMI this way.</li>
</ul></li>
<li>None of the write bits in this register prevent external NMI generated from pin 29 low.</li>
<li>None of the write bits prevent exteral IRQ generated from pin 28 low either.</li>
<li>Pin 33 (MSM6827L /INT) low triggers interrupt.</li>
<li>Pin 73 (Tone Rx DV) high triggers interrupt.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $08. ($41 XOR $08 = $49 existence mask)</li>
<li>Value $41 from previous fetch also observed changed to $20. ($41 XOR $20 = $61 existence mask)</li>
<li>Value $40 from previous fetch observed changed to $20. ($40 XOR $20 = $60 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $26. ($BF XOR $26 = $99 existence mask)</li>
<li>$49 OR $61 OR $60 OR $99 = $F9 accumulated existence mask.</li>
</ul></li>
<li>ROM implements a shadow read register at $0015.</li>
<li>ROM implements a shadow write register at $0016.</li>
<li>ROM $E31C clears bit 0 of this register.</li>
<li>ROM $E33B writes to this address, setting bit 0 to 1.</li>
<li>ROM $E4D3 reads from this address in the IRQ handler and uses bit 6.</li>
<li>ROM $F377 writes value #$00 to this address.</li>
<li>ROM $F40A writes value #$C1 to this address.</li>
<li>ROM $F4E1 writes value #$00 to this address.</li>
<li>ROM $F4E9 writes the shadow register to this address.</li>
<li>ROM $F99C writes value #$00 to this address.</li>
<li>Bits 7, 6, and 0 differ in values written, suggesting that they exist.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4130</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Serial bit to be written
+++++++--- (unknown, unlikely to exist)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>The buffer at RAM $0330 is written to this register 1 bit at a time, seemingly using only bit 0 of the register.</li>
<li>The most significant bit of each byte is sent first.</li>
<li>40 bytes are sent from $0330, then only the 6 most significant bits of the 41st byte.
<ul>
<li>A total of 326 bits are sent to this register.</li>
</ul></li>
<li>Note: CPU2 commands $10 and $11 fill this buffer.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, $26, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $FF2D writes to this address.</li>
<li>ROM $FF40 writes to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4131</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Serial bit to be written
+++++++--- (unknown, unlikely to exist)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>The buffer at RAM $0360 is written to this register 1 bit at a time, seemingly using only bit 0 of the register.</li>
<li>The most significant bit of each byte is sent first.</li>
<li>40 bytes are sent from $0360, then only the 6 most significant bits of the 41st byte.
<ul>
<li>A total of 326 bits are sent to this register.</li>
</ul></li>
<li>Note: CPU2 commands $10, $11, and $6A fill this buffer.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $FF51 writes to this address.</li>
<li>ROM $FF64 writes to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4132</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Serial bit to be written
+++++++--- (unknown, unlikely to exist)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>The buffer at RAM $0390 is written to this register 1 bit at a time, seemingly using only bit 0 of the register.</li>
<li>The most significant bit of each byte is sent first.</li>
<li>40 bytes are sent from $0390, then only the 6 most significant bits of the 41st byte.
<ul>
<li>A total of 326 bits are sent to this register.</li>
</ul></li>
<li>Note: CPU2 commands $10, $11, and $6A fill this buffer.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $FF75 writes to this address.</li>
<li>ROM $FF88 writes to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4133</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
|||||||+-- Serial bit to be written
+++++++--- (unknown, unlikely to exist)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>The buffer at RAM $0320 is written to this register 1 bit at a time, seemingly using only bit 0 of the register.</li>
<li>The most significant bit of each byte is sent first.</li>
<li>4 bytes are sent from $0320, then only the 4 most significant bits of the 5th byte.
<ul>
<li>A total of 36 bits are sent to this register.</li>
</ul></li>
<li>Note: CPU2 commands $10 and $11 fill this buffer.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BF from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $FF99 writes to this address.</li>
<li>ROM $FFAC writes to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4134</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- Bit exists but function is unknown
+++++++--- (unlikely to exist)</code></pre>
<ul>
<li>This register has serial data that can be read 1 bit at a time from $4134.0.
<ul>
<li>Without any other register reads/writes, the next bit is provided the next time this register is read.</li>
</ul></li>
<li>In each byte read, the most significant bit comes first.</li>
<li>CPU2 built-in ROM reads 40 bytes this way and stores them starting at RAM $03C0.
<ul>
<li>Then 6 additional bits are read and stored as the 41st byte bits 7:2, and 0 for bits 1:0.</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Value $41 from previous fetch observed changed to $40. ($41 XOR $40 = $01 existence mask)</li>
<li>Value $40 from previous fetch observed unaffected.</li>
<li>Value $BF from previous fetch observed changed to $BE. ($BF XOR $BE = $01 existence mask)</li>
<li>$01 OR $01 = $01 accumulated existence mask.</li>
</ul></li>
<li>ROM $FFB8 reads this address and uses bit 0</li>
<li>ROM $FFCD reads this address and uses bit 0</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4135</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
|||||||+-- $4135.0 = $4134 Serial Data Ready (1 = ready)
+++++++--- (unlikely to exist)</code></pre>
<ul>
<li>If $4135.0 = 1, Built-in ROM proceeds to read the serial data from $4134.</li>
<li>Note: Built-in ROM only treats this register as a flag and does not treat it as serial data.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BE from previous fetch observed unaffected.</li>
<li>Value $41 from previous fetch also observed changed to $40. ($41 XOR $40 = $01 existence mask)</li>
<li>Value $BF from previous fetch observed changed to $BE. ($BF XOR $BE = $01 existence mask)</li>
<li>$01 OR $01 = $01 accumulated existence mask.</li>
</ul></li>
<li>ROM $FD35 reads this address and uses bit 0.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4136</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Built-in ROM writes #$00 to $4137, then writes serial data to $4130, $4131, $4132, $4133, then writes #$00 to this register.</li>
<li>Theory: This is a "start" flag that starts using new data from $4130-4133 when the flag is cleared.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BE from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $F74A writes value #$00 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$4137</td>
<td class="unknown table-unknown" style="text-align: center; background: #ececec; color: #2C2C2C; font-size: smaller; vertical-align: middle;">Unknown</td>
<td class="table-no" style="text-align: center; background: #FA9; vertical-align: middle;">No</td>
<td class="table-yes" style="text-align: center; background: #9FB; vertical-align: middle;">Yes</td>
<td>Unknown Function.
<pre><code>Write
76543210
++++++++-- (unknown)
&#10;Read
76543210
++++++++-- (unlikely to exist)</code></pre>
<ul>
<li>Built-in ROM writes #$00 to this register, then writes serial data to $4130, $4131, $4132, $4133, then writes #$00 to $4136.</li>
<li>Theory: This is a "clear" flag that resets write bit counters and/or clears old data from $4130-4133 when the flag is cleared.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Read Bit Existence:
<ul>
<li>Values $41, $40, and $BE from previous fetch observed unaffected.</li>
</ul></li>
<li>ROM $F3FA writes value #$00 to this address.</li>
<li>ROM $F73B writes value #$00 to this address.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table>

# Communication Between Famicom and CPU2 {#Communication_Between_Famicom_and_CPU2}

Registers \$40D0, \$40D1, \$40D2, and \$40D3 are used for communication between Famicom and CPU2. One would tend to expect the Famicom to receive the same value from any of these four registers if read back right after writing. However, each register is actually a *separate* register in each direction. The Famicom only controls the value read by CPU2, and CPU2 only controls the value read by the Famicom. It may be that CPU2 echoes the value back in some cases, but don't be fooled.

Data packets are sent in both directions between the Famicom and CPU2 using these registers. The data flow is controlled by status and acknowledge flags in \$40D3, and data is sent 3 bytes at a time using registers \$40D0, \$40D1, and \$40D2. When CPU2 receives each 3-byte chunk, it buffers it in its RAM starting at address \$0401 until the full message has been received. The maximum message length is at most 255 bytes, possibly less.

Each message starts with a command byte, followed by a byte count. The byte after the byte count is not used and not counted towards the byte count in most commands. There are a total of 25 command bytes, which are stored in a lookup table at CPU2 ROM address \$FB52. The index of this lookup table corresponds to the index of a function pointer table at address \$FBB2, and a command mode support bitfield table at address \$FB6B. This is how CPU2 efficiently directs to a unique message handler function for each command byte. The mode bitfield checks against the mode byte at \$0051. It is probably used for enforcing the correct sequence of commands, as the command handlers themselves seem to be a major contributor changing the mode. It seems there are 6 possible modes: 0, 1, 2, 3, 4, and 5, though command \$03 is set to support modes 6 and 7 as well. The mode may actually represent a global state machine, such as 0 = disconnected, 1 = dialing, etc.

## CPU2 Commands {#CPU2_Commands}

### Commands Read by the Famicom from CPU2 {#Commands_Read_by_the_Famicom_from_CPU2}

<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>Command<br />
Byte</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>$80</td>
<td>This command is received by the Famicom in response to writing to command $00. (See next section for details.)</td>
</tr>
<tr>
<td>$81</td>
<td>Unknown function.
<pre><code>[$81] [count=$01] [$00] [connection status]</code></pre>
<ul>
<li>It is unknown what events trigger this command to be sent to the Famicom.</li>
<li>[connection status] comes from RAM $0052, see section below for enumerated values.</li>
<li>This response can only happen in mode 3 via NMI.</li>
<li>Theory: This command is a response to writing to command $01.</li>
</ul></td>
</tr>
<tr>
<td>$82</td>
<td>Unknown function.
<pre><code>[$82] [count=$01] [$00] [connection status]</code></pre>
<ul>
<li>It is unknown what events trigger this command to be sent to the Famicom.</li>
<li>[connection status] comes from RAM $0052, see section below for enumerated values.</li>
</ul>
<p>Response:</p>
<ul>
<li>[parameter] $00 means success, all other values indicate an error.
<ul>
<li>When successful, the mode changes to $02.</li>
<li>When unsuccessful, the mode changes to $00.</li>
</ul></li>
<li>[parameter] comes from RAM $0052</li>
<li>This response can only happen in mode 4 via NMI.</li>
<li>Theory: This command is a response to writing to command $02.</li>
</ul></td>
</tr>
<tr>
<td>$83</td>
<td>This command is received by the Famicom in response to writing to command $03. (See next section for details.)</td>
</tr>
<tr>
<td>$90</td>
<td>This command is received by the Famicom in response to writing to command $10. (See next section for details.)</td>
</tr>
<tr>
<td>$91</td>
<td>This command is received by the Famicom in response to writing to command $11. (See next section for details.)</td>
</tr>
<tr>
<td>$92</td>
<td>This command is received by the Famicom in response to writing to command $12. (See next section for details.)</td>
</tr>
<tr>
<td>$C0</td>
<td>UART Rx Packet
<p>There are 2 different parts of the CPU2 code that can send this message:</p>
<pre><code>[$C0] [count=N] [parameter] [...N bytes...]</code></pre>
<ul>
<li>This command can only be sent to the Famicom in mode 2.</li>
<li>[parameter] comes from RAM $0024.</li>
<li>[...N bytes...] are read from register $4110, most recent reads first, 1 byte read from $4110 per IRQ.</li>
<li>Count N does not include the parameter byte.</li>
<li>This is serial data received from the Oki MSM6827L integrated modem chip.</li>
<li>This response can only happen in mode 2 via IRQ.</li>
</ul>
<p><br />
</p>
<pre><code>[$C0] [count=$01] [$00] [parameter]</code></pre>
<ul>
<li>It is unknown what events trigger this command to be sent to the Famicom.</li>
<li>The modem abruptly disconnects when sending this command.</li>
<li>[parameter] can be #$00 or #$01.</li>
<li>This response can only happen in mode 2 via IRQ.</li>
</ul></td>
</tr>
<tr>
<td>$C1</td>
<td>Tone Rx Data Packet.
<p>When the packet only contains 1 nybble:</p>
<pre><code>[$C1] [count=$01] [$80] [tone Rx nybble]</code></pre>
<ul>
<li>[tone Rx nybble] is register ($4128 read value &gt;&gt; 3) &amp; 0x0F.</li>
</ul>
<p><br />
When there is more than 1 nybble in the packet:</p>
<pre><code>[$C1] [count=N] [parameter] [tone Rx bytes]</code></pre>
<ul>
<li>[parameter] is either $00 or $01 based on the value of RAM $007A.
<ul>
<li>Theory: [parameter] may indicate an even or odd number of nybbles in the packet.</li>
</ul></li>
<li>Each Tone Rx byte is packed with 2 nybbles.</li>
<li>Count N does not include the parameter byte.</li>
</ul></td>
</tr>
<tr>
<td>$E1</td>
<td>This command indicates that CPU2 had a failure receiving a command from the Famicom.
<p>Command from Famicom that failed due to invalid command byte, or a command not supported in the present mode:</p>
<pre><code>Message from Famicom:
[invalid command] [count] [byte 0] [byte 1] ...
&#10;Response back to Famicom:
[$E1] [count=$03] [?] [$00] [invalid command] [byte 1]</code></pre>
<ul>
<li>[?] byte appears to be open to sending an artifact from the previous response.</li>
<li>Observed JRA-PAT rev 05 read this response after attempting to write to command $01 after a failed connection.
<ul>
<li>E1 03 (00) 00 01 01</li>
</ul></li>
</ul>
<p><br />
Command from Famicom that failed due to invalid byte count:</p>
<pre><code>Message from Famicom:
[command] [invalid count] [byte 0] [byte 1] ...
&#10;Response back to Famicom:
[$E1] [count=$03] [$00] [$01] [command] [byte 1]</code></pre></td>
</tr>
<tr>
<td>$F0</td>
<td>This command is received by the Famicom in response to writing to command $7C. (See next section for details.)</td>
</tr>
</tbody>
</table>

### Commands Written by the Famicom to CPU2 {#Commands_Written_by_the_Famicom_to_CPU2}

Note that most descriptions are *incomplete*.

<table class="wikitable">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr>
<th>Command<br />
Byte</th>
<th>CPU2<br />
Handler<br />
Address</th>
<th>Byte<br />
Count<br />
Expected</th>
<th>Modes<br />
Supported</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>$F43A</td>
<td>&gt;= 0<br />
&lt;= 60</td>
<td>0</td>
<td>Hook And Dialing Sequence Configuration
<pre><code>Message Bytes:
[$00] [count=N] [xx] [...N bytes command string...]
&#10;Response:
[$80] [count=$01] [$00] [status byte]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>All observed software uses this command when opening a connection.</li>
<li>Changes to Command Mode 1 when the full message has been received from the Famicom.</li>
<li>[command string] is a series of one-byte commands executed in Command Mode 1.
<ul>
<li>$00-$1F: ignored</li>
<li>$20: wait up to 3 seconds for call progress (dial) tone to stop
<ul>
<li>Fails with status $03 if the tone never stops</li>
</ul></li>
<li>$21: switch to DTMF dialing</li>
<li>$22: switch to normal (10pps) pulse dialing</li>
<li>$23: switch to fast (20pps) pulse dialing</li>
<li>$24: turn on audio output</li>
<li>$25: turn off audio output</li>
<li>$26: disconnect phone line (on hook)</li>
<li>$27: connect phone line (off hook)</li>
<li>$28-$2F: ignored</li>
<li>$30: dial 0 (10 pulses)</li>
<li>$31-$39: dial 1-9 (1-9 pulses)</li>
<li>$3A: dial D (no pulses)</li>
<li>$3B: dial * (11 pulses)</li>
<li>$3C: dial # (12 pulses)</li>
<li>$3D: dial A (13 pulses)</li>
<li>$3E: dial B (14 pulses)</li>
<li>$3F: dial C (15 pulses)</li>
<li>$40-7F: currently unknown, but not ignored</li>
<li>$80-FF: delay 0.2-25.6 seconds (in 200ms increments)</li>
</ul></li>
</ul>
<p>Response:</p>
<ul>
<li>[connection status] comes from RAM $0052, see section below for enumerated values.
<ul>
<li>When successful, the mode changes to $02.</li>
<li>When unsuccessful, the mode changes to $00.</li>
</ul></li>
<li>This response can only happen in mode 1 via NMI.</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><ul>
<li>Starts NMI timer 1 with period $0006 (5 msec), one-shot:
<ul>
<li>Writes #$06 to $4100</li>
<li>Writes #$00 to $4101</li>
<li>Writes #$02 to $4102</li>
</ul></li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>Super Mario Club rev 09:
<ul>
<li>00 13 (00) 42 31 20 36 33 57 30 36 30 37 38 31 34 34 31 33 25 67 25 (A2 01)</li>
<li>"...B1 63W0607814413%g%¢."</li>
<li>Note: The first '%' is definitely not a '$' in this one.</li>
</ul></li>
<li>Heart no Benrikun Mini rev 01 A:
<ul>
<li>00 13 (00) 42 31 20 36 33 57 30 36 30 33 31 33 35 38 35 35 24 67 25 (37 31)</li>
<li>"...B1 63W0603135855$g%71"</li>
</ul></li>
<li>Okasan no Famicom Trade rev 01:
<ul>
<li>00 13 (00) 42 31 20 36 33 57 30 36 30 33 31 33 33 33 39 30 24 67 25 (00 00)</li>
<li>"...B1 63W0603133390$g%.."</li>
</ul></li>
<li>JRA-PAT rev 05:
<ul>
<li>00 0F (00) 88 30 42 30 20 36 34 32 35 36 34 36 36 36 24</li>
<li>"...ˆ0B0 642564666$"</li>
</ul></li>
<li>JRA-PAT rev 05, with phone number manually set to 6666666666
<ul>
<li>00 0F (00) 42 36 20 36 36 36 36 36 36 36 36 36 24 67 25</li>
<li>"...B6 666666666$g%"</li>
</ul></li>
<li>PiT Motorboat Race rev 02, with phone number manually set to 1234567890:
<ul>
<li>00 0E (00) 24 42 31 20 32 33 34 35 36 37 38 39 30 00 (00)</li>
<li>"...$B1 234567890.."</li>
</ul></li>
</ul></li>
</ul>
<p><br />
</p>
<ul>
<li>Okasan no Famicom Trade rev 01 responses observed:
<ul>
<li>80 01 00 01 after failure due to no phone line connected, reported Error 4001.</li>
<li>80 01 00 03 after failure due to dial tone not going away, reported Error 4003.</li>
<li>80 01 00 06 after failure due to server never answering, reported Error 4006.</li>
</ul></li>
<li>PiT Motorboat Race rev 02 responses observed:
<ul>
<li>80 01 00 01 after failure due to no phone line connected.</li>
<li>80 01 00 0D after failure due to server never answering.</li>
</ul></li>
<li>JRA-PAT rev 05 responses observed:
<ul>
<li>80 01 00 01 after failure due to no phone line connected, reported Error 390010019999.</li>
<li>80 01 00 03 after failure due to dial tone not going away, reported Error 390030019999.</li>
<li>80 01 00 04 after failure due to server never answering, reported Error 390040019999.</li>
</ul></li>
<li>The response [parameter] byte comes from RAM $0052.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$01</td>
<td>$F46C</td>
<td>1</td>
<td>2</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$01] [count=$01] [xx] [parameter]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>[parameter] Values:
<ul>
<li>$00 = Write #$00 to $0031</li>
<li>$01-FF = Write #$02 to $0031</li>
</ul></li>
<li>JRA-PAT and PiT Motorboat Race use this command when closing a connection.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>JRA-PAT rev 05:
<ul>
<li>01 01 (00) 01 (30 42)</li>
</ul></li>
<li>PiT Motorboat Race rev 02:
<ul>
<li>01 01 (00) 01 (02 31)</li>
</ul></li>
</ul></li>
<li>Checks if the byte received is 00 and if so, acts differently.</li>
<li>Starts NMI timer 1 with period $0006 (5 msec), one-shot:
<ul>
<li>Writes #$06 to $4100</li>
<li>Writes #$00 to $4101</li>
<li>Writes #$02 to $4102</li>
</ul></li>
<li>Changes to Command Mode 3 when complete.</li>
<li>Theory: This command generates response command $81 (see previous section).</li>
</ul></td>
</tr>
<tr>
<td>$02</td>
<td>$F48A</td>
<td>2</td>
<td>0</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$02] [count=$02] [xx] [parameter 0] [parameter 1]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>[parameter 0] gets written to $002E.</li>
<li>[parameter 1] gets written to $002B and $002C.</li>
<li>9 zero-page addresses get zeroed out; seems to be initializing something.</li>
<li>Starts NMI timer 1 with period $0006 (5 msec), one-shot:
<ul>
<li>Writes #$06 to $4100</li>
<li>Writes #$00 to $4101</li>
<li>Writes #$02 to $4102</li>
</ul></li>
<li>Changes to Command Mode 4 when complete.</li>
<li>Theory: This command generates response command $82 (see previous section).</li>
</ul></td>
</tr>
<tr>
<td>$03</td>
<td>$F4BB</td>
<td>0</td>
<td>0,1,2,3,<br />
4,5,6,7</td>
<td>System Revision and Status Information
<pre><code>Message Bytes:
[$03] [count=$00]
&#10;Response:
[$83] [count=$0A] [$00] [ROM revision] [ROM checksum MSB] [ROM checksum LSB] [byte 4] ... [byte 10]</code></pre>
<ul>
<li>Writes to $4120.</li>
<li>Writes $412F shadow register to $412F (update only, no apparent modification).</li>
<li>Response:
<ul>
<li>[byte 1] = Comes from ROM address $FFF9, which seems to be a ROM revision byte. Observed value $03.</li>
<li>[byte 2], [byte 3] = 16-bit ROM checksum. Computed at bootup with the function located at $FD7D. It adds each byte of ROM.</li>
<li>[byte 4] = Value read from register $4121 &amp; #$3F. If RAM $0023 is not equal to #$04, it also sets bit 7.</li>
<li>[byte 5] = Value read from register $4126.</li>
<li>[byte 6] = Value most recently written to $4127.</li>
<li>[byte 7] = Present mode.</li>
<li>[byte 8] = Number of 256-byte data blocks presently queued up in RAM range $1300-$1FFF.</li>
<li>[byte 9] = Value of RAM $0076.</li>
<li>[byte 10] = Value of RAM $0077.</li>
</ul></li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>Super Mario Club rev 09 at power-on:
<ul>
<li>Write: 03 00 (00)</li>
<li>Read: 83 0A 00 03 0C AF A1 FB FE 00 0D 00 00</li>
</ul></li>
<li>JRA-PAT rev 05 and 06 at power-on:
<ul>
<li>Write: 03 00 (00)</li>
<li>Read: 83 0A 00 03 0C AF 81 FB FE 00 0D 00 00</li>
</ul></li>
</ul></li>
</ul></td>
</tr>
<tr>
<td>$10</td>
<td>$F71A</td>
<td>40</td>
<td>0,2,5</td>
<td>Unknown $413x Function.
<pre><code>Message Bytes:
[$10] [count=$28] [parameter] [...40 bytes...]
&#10;Response:
[$90] [count=$28] [response parameter] [...40 bytes...]</code></pre>
<ul>
<li>This command is similar to command $11, except using pre-defined data from ROM instead of the additional chunks in command $11.</li>
<li>[parameter] is used but does not count towards the byte count.</li>
<li>[parameter] value #$00 - #$0F:
<ul>
<li>Pre-defined ROM data gets copied from tables.</li>
<li>40 bytes from ROM are copied starting at RAM address $0360. ROM table index using parameter bits 2,1,0.
<ul>
<li>$0388 (i.e. 41st byte) set to #$00.</li>
<li>See Register $4131.</li>
</ul></li>
<li>40 bytes from ROM are copied starting at RAM address $0390. ROM table index using parameter bits 1,0.
<ul>
<li>$03B8 (i.e. 41st byte) set to #$00.</li>
<li>See Register $4132.</li>
</ul></li>
<li>40 bytes from the command are copied starting at RAM address $0330.
<ul>
<li>$0358 (i.e. 41st byte) set to #$00.</li>
<li>$0330 (i.e. 1st byte) set to #$00.</li>
<li>See Register $4130.</li>
</ul></li>
<li>5 static bytes from ROM are copied starting at RAM address $0320: 01 40 50 00 60
<ul>
<li>See Register $4133.</li>
</ul></li>
</ul></li>
<li>[parameter] value #$10:
<ul>
<li>Skips the $0360 and $0390 ROM copy routines if RAM $0055 == #$00.</li>
<li>If $0055 is not #$00, it aborts the command.</li>
</ul></li>
<li>Writes #$00 to $4137</li>
<li>Writes the 4 data chunks shifting 1 bit at a time to $4130, $4131, $4132, $4133</li>
<li>Writes #$00 to $4136</li>
<li>For the response, the count has not been confirmed to include [response parameter] or not.</li>
<li>The response comes from RAM starting at $03C0. (See register $4134.)</li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>ROM Data </td>
</tr>
<tr>
<td><p>Data Loaded to $0330 (which later gets written to register $4130):</p>
<p>(These are the 40 bytes supplied by the Famicom to this command, with 1st byte and 41st byte set to $00.)</p></td>
</tr>
<tr>
<td><p>Data Loaded to $0360 (which later gets written to register $4131):</p>
<table class="wikitable">
<thead>
<tr>
<th>Parameter Bits</th>
<th>Data</th>
</tr>
</thead>
<tbody>
<tr>
<td>xxxxx000</td>
<td>00000c6bbc41351b1e0fa372963f7af704415090ee9a139ff87c60ca610dec66f74bd5b992c8868300</td>
</tr>
<tr>
<td>xxxxx001</td>
<td>0000027473468ebd1f161331ad0e5512ba56487f7b30037dac78ed9c51b308b6d8c597ca8b59cc1300</td>
</tr>
<tr>
<td>xxxxx010</td>
<td>000001c89b4630b518cbc342a45282fa0908773df56830150ec1955c9ec1abbd1100d4d878604afd00</td>
</tr>
<tr>
<td>xxxxx011</td>
<td>00000939c25f1859304000e9fdab927959a49d1acdd136da4ddb2b551b2bc942a36077d1d09dbc3500</td>
</tr>
<tr>
<td>xxxxx100</td>
<td>3cfed7ad4cd7435c3544144572d1d0e83d1859864bf8c7412cac015ced49a1cabce13fbd6262af0b00</td>
</tr>
<tr>
<td>xxxxx101</td>
<td>8aac02df739c8c0a2572493bd95755450b06ea0d77f1f08bc665636ed3cf317cb89f6f301ef03a2f00</td>
</tr>
<tr>
<td>xxxxx110</td>
<td>39b07c4f12614dadb4eb214af3a6a0dc046b47d10c184b04d16c91d9eeb5e0c21a052529738c6f0100</td>
</tr>
<tr>
<td>xxxxx111</td>
<td>36a6a31fa79ba49a9f1d1985c640ad9a6c79a4195d06ece8726fc006e1d347efa5d129cdb79017b900</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td><p>Data Loaded to $0390 (which later gets written to register $4132):</p>
<table class="wikitable">
<thead>
<tr>
<th>Parameter Bits</th>
<th>Data</th>
</tr>
</thead>
<tbody>
<tr>
<td>xxxxxx00</td>
<td>882a14099f3f32b2419d0b09237a252ece6534274700628043a6df86657a84c0c4463fa72b49ad4900</td>
</tr>
<tr>
<td>xxxxxx01</td>
<td>bbb57552f19d8a71d6ed68fd037188b700eab270071acedc624ff4c64b53ed1c4e1046ed89e775fd00</td>
</tr>
<tr>
<td>xxxxxx10</td>
<td>cfb8e78ec54b81a1100f02f524acf0e05a45cc36a98f161c115a41082abb587ecabad5453511877100</td>
</tr>
<tr>
<td>xxxxxx11</td>
<td>aa189432166db82500fffbb1107ced2469b12f0959b263c63a0332bb4a8e6ea378a0a47150cb296100</td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td><p>Data Loaded to $0320 (which later gets written to register $4133):</p>
<table class="wikitable">
<thead>
<tr>
<th>Parameter Bits</th>
<th>Data</th>
</tr>
</thead>
<tbody>
<tr>
<td>xxxxxxxx</td>
<td>0140500060</td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$11</td>
<td>$F75C</td>
<td>127</td>
<td>0,2,5</td>
<td>Unknown $413x Function.
<pre><code>Message Bytes:
[$11] [count=$7F] [...5 byte chunk...] [...41 byte chunk 0...]
                  [...41 byte chunk 1...] [...41 byte chunk 2...]
&#10;Response:
[$91] [count=$29] [response parameter] [...41 bytes...]</code></pre>
<ul>
<li>Note that the first byte of the 5 byte chunk is not ignored in this command like it is in most others, though it still does not count towards the byte count.</li>
<li>The 5 byte chunk gets written to RAM starting at $0320 with lots of bit shifting and byte swapping. (See Register $4133.)</li>
<li>The 41 byte chunk 0 gets copied to RAM starting at $0330. (See Register $4130.)</li>
<li>The 41 byte chunk 1 gets copied to RAM starting at $0360. (See Register $4131.)</li>
<li>The 41 byte chunk 2 gets copied to RAM starting at $0390. (See Register $4132.)</li>
<li>Shares the remaining code in common with command $10.</li>
<li>For the response, the count has not been confirmed to include [response parameter] or not.</li>
<li>The response comes from RAM starting at $03C0. (See register $4134.)</li>
</ul></td>
</tr>
<tr>
<td>$12</td>
<td>$F773</td>
<td>&gt;= 3<br />
&lt;= 252</td>
<td>0,2,5</td>
<td>CRC Calculator
<pre><code>Message Bytes:
[$12] [count=N] [initial] [polynomial 1] [polynomial 2] [...N-2 bytes...]
&#10;Response:
[$92] [count=$02] [$00] [CRC 2] [CRC 1]</code></pre>
<ul>
<li>The [initial] byte does not count towards the byte count.</li>
<li>See next section for information about CRC keys.</li>
<li>Doesn't seem to write to any registers.</li>
</ul></td>
</tr>
<tr>
<td>$40</td>
<td>$F7AF</td>
<td>&gt;= 1<br />
&lt;= 252</td>
<td>2</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$40] [count=N] [xx] [...N-1 bytes...]</code></pre>
<ul>
<li>[xx] byte is ignored, but <em>does</em> count towards the byte count.</li>
<li>Copies the N-1 bytes data directly into memory starting at the location +2 pointed to by $0000:0001.</li>
<li>Copies the byte count to the location +1 pointed to by $0000:0001.</li>
<li>Writes value $FF directly to the location pointed to by $0000:0001.
<ul>
<li>Not known what sets up that pointer.</li>
</ul></li>
<li>This command is tricky to disassemble because it actually does copy the [xx] byte, then writes over it with the byte count, ultimately not using it.</li>
<li>Theory: This command sends a UART Tx Packet to the Oki MSM6827L, and connected with response command $C0.</li>
</ul></td>
</tr>
<tr>
<td>$41</td>
<td>$F7D0</td>
<td>&gt;= 1<br />
&lt;= 100</td>
<td>5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$41] [count=N] [xx] [...N bytes...]</code></pre>
<ul>
<li>[xx] byte is ignored and does not count towards the byte count.</li>
<li>Writes ((Byte count + 2) * 2) to the location +1 pointed to by $0000:0001.</li>
<li>Writes value $00 to the location +2 pointed to by $0000:0001.</li>
<li>Writes value $0F to the location +3 pointed to by $0000:0001.</li>
<li>For each of N bytes:
<ul>
<li>Writes upper nybble of the byte (&gt;&gt; 4) to the next byte pointed to by $0000:0001.</li>
<li>Writes the lower nybble of the byte (&amp; $0F) to the next byte pointed to by $0000:0001.</li>
</ul></li>
<li>Appends the value of $007C similarly as 2 nybbles at the end.</li>
<li>Writes value $FF directly to the location pointed to by $0000:0001.
<ul>
<li>Not known what sets up that pointer.</li>
</ul></li>
<li>Theory: This command is something to do with the Tone Rx chip, and connected with response command $C1.</li>
</ul></td>
</tr>
<tr>
<td>$60</td>
<td>$F829</td>
<td>6</td>
<td>0</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$60] [count=$06] [...6 data bytes...]</code></pre>
<ul>
<li>Note that the first data byte is not ignored in this command, and it <em>does</em> count towards the byte count.</li>
<li>The 6 data bytes written by this command get copied to CPU2 RAM starting at address $84.</li>
<li>PiT Motorboat Race uses this command when opening a connection.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>PiT Motorboat Race rev 02:
<ul>
<li>60 06 00 01 80 00 00 00 (00)</li>
</ul></li>
</ul></li>
</ul></td>
</tr>
<tr>
<td>$61</td>
<td>$F839</td>
<td>0</td>
<td>1,3,4</td>
<td>Disconnect
<pre><code>Message Bytes:
[$61] [count=$00]</code></pre>
<ul>
<li>This command abruptly disconnects the modem.</li>
<li>Writes #$00 to register $4102. (Stops NMI timer 1.)</li>
<li>Changes mode to 0 when complete.</li>
</ul></td>
</tr>
<tr>
<td>$62</td>
<td>$F849</td>
<td>1</td>
<td>0,2,5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$62] [count=$01] [xx] [parameter]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>Disables interrupts while executing this command.</li>
<li>Valid values of the [parameter] byte are $00 and $01.
<ul>
<li>Other values abort the command.</li>
</ul></li>
<li>Appears to write to $4120</li>
</ul></td>
</tr>
<tr>
<td>$63</td>
<td>$F887</td>
<td>0</td>
<td>0,1,2,3,<br />
4,5</td>
<td>NOP Command
<pre><code>Message Bytes:
[$63] [count=$00]</code></pre>
<ul>
<li>This command only validates the byte count = #$00 and is then complete.</li>
</ul></td>
</tr>
<tr>
<td>$64</td>
<td>$F88D</td>
<td>1</td>
<td>0,2,5</td>
<td>Modem Off Hook Control
<pre><code>Message Bytes:
[$64] [count=$01] [xx] [Modem Off Hook mode]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>[Modem Off Hook mode] Values:
<ul>
<li>$00 = Don't update but save mode as $00.</li>
<li>$01 = Set $4127.4 (/Modem Off Hook) = 0 (off hook). (Also stores #$00 to $0049)</li>
<li>$02-FF = Set $4127.4 (/Modem Off Hook) = 1 (hang up). (Also stores #$00 to $0049)</li>
<li>Value gets saved to $004E.</li>
</ul></li>
<li>PiT Motorboat Race uses this command when closing a connection.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>PiT Motorboat Race rev 02:
<ul>
<li>64 01 (00) 02 (A9 31)</li>
</ul></li>
</ul></li>
<li>Note: This description is fairly complete.</li>
</ul></td>
</tr>
<tr>
<td>$65</td>
<td>$F89D</td>
<td>2</td>
<td>0,2,5</td>
<td>Modem Audio Enable and P4-19 Control
<pre><code>Message Bytes:
[$65] [count=$02] [xx] [Modem Audio Enable mode] [P4-19 mode]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>[Modem Audio Enable mode] Values:
<ul>
<li>$00 = Don't update but save mode as $00.</li>
<li>$01 = Set $4127.6 (/Modem Audio Enable) = 0 (enabled to TV speakers).</li>
<li>$02-FF = Set $4127.6 (/Modem Audio Enable) = 1 (disabled).</li>
<li>Value gets saved to $004F.</li>
</ul></li>
<li>[P4-19 mode] Values:
<ul>
<li>$00 = Don't update but save mode as $00.</li>
<li>$01 = Set $4127.7 (Modem P4-19) = 0.</li>
<li>$02-FF = Set $4127.7 (Modem P4-19) = 1.</li>
<li>Value gets saved to $0050.</li>
</ul></li>
<li>Heart no Benrikun Mini rev 01 A uses this command when closing a connection.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>Heart no Benrikun Mini rev 01 A:
<ul>
<li>65 02 (00) 02 00 (20)</li>
</ul></li>
</ul></li>
<li>Note: This description is fairly complete.</li>
</ul></td>
</tr>
<tr>
<td>$66</td>
<td>$F8B7</td>
<td>2</td>
<td>0,2,5</td>
<td>LED Control
<pre><code>Message Bytes:
[$66] [count=$02] [xx] [red LED mode] [green LED mode]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>[red LED mode] Values:
<ul>
<li>$00 = Don't update but save mode as $00.</li>
<li>$01 = Set Red LED on.</li>
<li>$02-7F = Set Red LED off.</li>
<li>$80-FF = Set Red LED to be controlled automatically by network status.</li>
<li>Value gets saved to $004C.</li>
</ul></li>
<li>[green LED mode] Values:
<ul>
<li>$00 = Don't update but save mode as $00.</li>
<li>$01 = Set Green LED on.</li>
<li>$02-7F = Set Green LED off.</li>
<li>$80-FF = Set Green LED to be controlled automatically by network status.</li>
<li>Value gets saved to $004D.</li>
</ul></li>
<li>PiT Motorboat Race uses this command when closing a connection.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>PiT Motorboat Race rev 02:
<ul>
<li>66 02 (00) 02 02 (31)</li>
</ul></li>
</ul></li>
<li>Note: This description is fairly complete.</li>
</ul></td>
</tr>
<tr>
<td>$67</td>
<td>$F8D1</td>
<td>2</td>
<td>0,5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$67] [count=$02] [xx] [parameter 0] [parameter 1]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>If [parameter 0] is $00:
<ul>
<li>Abruptly disconnects the modem.</li>
<li>Zeroes out RAM $0076 and $0077.</li>
<li>Changes to mode 0 when complete.</li>
</ul></li>
<li>Else if [parameter 0] is not $00:
<ul>
<li>Writes [parameter 0] to $74.</li>
<li>Writes [parameter 1] to $75.</li>
<li>Refreshes Green LED based on its mode.</li>
<li>Refreshes P4-27 (Phone Off Hook) based on its mode.</li>
<li>Zeroes out RAM $76,77,37,39,38.</li>
<li>Refreshes P4-21 (Phone Audio Enable) based on its mode.</li>
<li>Zeroes out all queued up outgoing modem Tx messages at RAM $1300-1FFF.
<ul>
<li>Each message has space $100 and number of queued messages is at RAM $0100.</li>
</ul></li>
<li>Zeroes out Tone Rx incoming data buffer at RAM $0500-05FF.</li>
<li>Changes to mode 5 when complete.</li>
</ul></li>
</ul></td>
</tr>
<tr>
<td>$68</td>
<td>$F912</td>
<td>1</td>
<td>0,5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$68] [count=$01] [xx] [parameter]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>The code has 3 different paths depending on parameter.
<ul>
<li>[parameter] &lt; #$10</li>
<li>[parameter] == #$FF</li>
<li>All other values of [parameter]</li>
</ul></li>
</ul></td>
</tr>
<tr>
<td>$69</td>
<td>$F951</td>
<td>10</td>
<td>0</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$69] [count=$0A] [xx] [...10 data bytes...]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>Observed Messages (unused bytes in parentheses):
<ul>
<li>PiT Motorboat Race rev 02 when opening a connection:
<ul>
<li>69 0A (00) 00 04 14 1D 28 80 00 30 26 21 (00 00)</li>
</ul></li>
<li>JRA-PAT rev 05 and 06 at power-on:
<ul>
<li>69 0A (00) 00 04 14 1D 28 30 00 30 26 21 (00 00)</li>
</ul></li>
</ul></li>
<li>The 10 data bytes are copied to CPU2 RAM starting at $8A.</li>
<li>If the 6th data byte (location $8F) is &gt; #$FC, it gets overwritten by #$01.</li>
</ul></td>
</tr>
<tr>
<td>$6A</td>
<td>$F96F</td>
<td>80</td>
<td>0,2,5</td>
<td>Unknown $413x Function.
<pre><code>Message Bytes:
[$6A] [count=$50] [xx] [...40 byte chunk 0...] [...40 byte chunk 1...]</code></pre>
<ul>
<li>[xx] byte is ignored.</li>
<li>Copies the 40 byte chunk 0 to RAM starting at $0360. (See Register $4131.)</li>
<li>Copies the negative (2's complement) of each byte in 40 byte chunk 1 to RAM starting at $0390. (See Register $4132.)</li>
<li>Writes #$00 to $0388 and $03B8. (i.e. 41st byte of each.)</li>
<li>Writes #$10 to $0055.</li>
<li>Observations:
<ul>
<li>No updates were made to buffers at $0320, $0330 from this command. (i.e. register $4133, $4130 data)</li>
<li>The command does not appear to initiate any operations to the $413x registers.</li>
</ul></li>
<li>Note: This description is fairly complete.</li>
</ul></td>
</tr>
<tr>
<td>$7B</td>
<td>$F994</td>
<td>No<br />
Check</td>
<td>0,1,2,3,<br />
4,5</td>
<td>Software Reset
<pre><code>Message Bytes:
[$7B] [count=xx] [xx]</code></pre>
<ul>
<li>Does not care about any bytes beyond the command byte.</li>
<li>Refreshes the green LED.</li>
<li>Disables interrupts.</li>
<li>Writes #$00 to $412F.</li>
<li>Sorts through some memory comparing values then generates a software reset.
<ul>
<li>Stores #$09 as the reset type code into $00FF.</li>
</ul></li>
</ul></td>
</tr>
<tr>
<td>$7C</td>
<td>$F9AC</td>
<td>5</td>
<td>0,1,2,3,<br />
4,5</td>
<td>Arbitrary Memory Read
<pre><code>Message Bytes:
[$7C] [count=$05] [xx] [CRC 1] [CRC 2] [address MSB] [address LSB] [read count N]
&#10;Response:
[$F0] [read count N] [$00] ...N bytes from memory...</code></pre>
<ul>
<li>[xx] byte does not count towards the byte count.</li>
<li>See next section for information about CRC bytes.</li>
<li>If [address MSB] &gt;= #$80, exits early with no response. (Prevent ROM dump?? :) )</li>
<li>If [read count N] &gt;= #$39, exits early with no response.</li>
</ul></td>
</tr>
<tr>
<td>$7D</td>
<td>$F9D9</td>
<td>&gt;= 5</td>
<td>0,1,2,3,<br />
4,5</td>
<td>Arbitrary Memory Write
<pre><code>Message Bytes:
[$7D] [count=N] [xx] [CRC 1] [CRC 2] [address MSB] [address LSB] [...N-4 bytes...]</code></pre>
<ul>
<li>[xx] byte does not count towards the byte count.</li>
<li>See next section for information about CRC bytes.</li>
<li>count - 4 is the number of bytes to write to memory.</li>
<li>Writes $00 to $006E.</li>
<li>Super Mario Club rev 09 uses this command at the welcome screen.
<ul>
<li>7D 1A (00) 07 60 01 04 09 01 06 E0 16 A2 90 A5 29 F0 06 C9 FF D0 04 A2 01 86 29 4C 20 F4 (30)</li>
<li>CRC bytes: 07 60</li>
<li>Starting address: $0104 (farthest reaches of the stack)</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Disassembly </td>
</tr>
<tr>
<td><pre><code>       .org $0104
       .dd2 T0109   ;Wrote over Main Super-loop Function Pointer, now executes code below.
       .dd2 TE006   ;Wrote over IRQ Handler Function Pointer, now immediately exits IRQ.
       .dd1 $16     ;Satisfy periodic checksum of bytes $102-109.
      &#10;T0109  ldx  #$90    ;Arbitrary code inserted.
       lda  $29
       beq  L0115
       cmp  #$ff
       bne  L0117
       ldx  #$01
L0115  stx  $29
L0117  jmp  $f420   ;Continue to original main super-loop.
&#10;       ; Code Below Existing In CPU2 ROM:
       .org $e006
TE006  jmp  TE502
      &#10;       .org $e502
TE502  ply          ;Return from IRQ.
       plx
       pla
       rti</code></pre></td>
</tr>
</tbody>
</table>
<ul>
<li>JRA-PAT rev 05 and 06 use this command at power-on.
<ul>
<li>7D 29 (00) EC AD 01 02 1A 01 09 01 06 E0 DB A2 90 A5 29 F0 06 C9 FF D0 04 A2 01 86 29 4C 20 F4 AD 32 00 C9 05 D0 03 EE 32 00 4C 39 E8 (00)</li>
<li>CRC bytes: EC AD</li>
<li>Starting address: $0102 (farthest reaches of the stack)</li>
</ul></li>
</ul>
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Disassembly </td>
</tr>
<tr>
<td><pre><code>       .org $0102
       .dd2 T011A   ;Wrote over NMI function pointer, now executes code below.
       .dd2 T0109   ;Wrote over Main Super-loop Function Pointer, now executes code below.
       .dd2 TE006   ;Wrote over IRQ Handler Function Pointer, now immediately exits IRQ.
       .dd1 $db     ;Satisfy periodic checksum of bytes $102-109.
      &#10;T0109  ldx  #$90    ;Arbitrary code inserted.
       lda  $29
       beq  L0115
       cmp  #$ff
       bne  L0117
       ldx  #$01
L0115  stx  $29
L0117  jmp  $f420   ;Continue to original main super-loop.
      &#10;T011A  lda  $0032   ;Arbitrary NMI handler code inserted.
       cmp  #$05
       bne  L0124
       inc  $0032
L0124  jmp  TE839
      &#10;       ; Code Below Existing In CPU2 ROM:
       .org $e006
TE006  jmp  TE502
      &#10;       .org $e502
TE502  ply          ;Return from IRQ.
       plx
       pla
       rti
      &#10;       .org $e839
TE839  ply          ;Return from NMI.
       plx
       pla
       rti</code></pre></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$7E</td>
<td>$F9FE</td>
<td>5</td>
<td>0,1,2,3,<br />
4,5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$7E] [count=$05] [xx] [CRC 1] [CRC 2] [parameter 0] [parameter 1] [parameter 2]</code></pre>
<ul>
<li>[xx] byte does not count towards the byte count.</li>
<li>See next section for information about CRC bytes.</li>
<li>Writes a lot of stuff to $4120 and $4121</li>
<li>Writes to $4127 bit 5 (/DTMF Output Enable)</li>
</ul></td>
</tr>
<tr>
<td>$7F</td>
<td>$FA16</td>
<td>5</td>
<td>0,1,2,3,<br />
4,5</td>
<td>Unknown Function.
<pre><code>Message Bytes:
[$7F] [count=$05] [parameter 0] [CRC 1] [CRC 2] [parameter 1] [parameter 2] [parameter 3]</code></pre>
<ul>
<li>[parameter 0] byte does not count towards the byte count.</li>
<li>See next section for information about CRC bytes.</li>
<li>If [parameter 0] is $00-7F, write value $0C out the P5 Expansion Port.</li>
<li>If [parameter 0] is $80-FF, dump RAM $0240-$0280 out the P5 Expansion Port.
<ul>
<li>The P5 Expansion port sequences out 1 nybble at a time.</li>
</ul></li>
<li>[parameter 1] gets written to $00B7.</li>
<li>[parameter 2] gets written to $00B6.</li>
<li>[parameter 3] gets written to $00B8.</li>
<li>[parameter 0] gets written to $00B5.</li>
</ul></td>
</tr>
</tbody>
</table>

### CPU2 Commands with CRC key bytes {#CPU2_Commands_with_CRC_key_bytes}

CPU2 command \$12 calculates the CRC-16 of its message using \[polynomial 1\] and \[polynomial 2\] as the CRC polynomial, with \[initial\] placed in both bytes of the initial value, and performs the calculation LSB-first or "reflected". This command requires the polynomial to be specified in reciprocal form. For example, to calculate the standard CRC-16 with polynomial \$8005, \[polynomial 1\] would be set to \$40, \[polynomial 2\] would be set to \$03, and \[initial\] would be set to \$00.

CPU2 Commands \$7C, \$7D, \$7E, and \$7F use CRC key bytes \[CRC 1\] and \[CRC 2\] to validate the message. These commands calculate a CRC-16 across the message backwards, starting from the last byte and ending at \[CRC 1\]. The result must be 0 or the command is ignored. The CRC-16 polynomial is \$8385 and the initial value is \$35AC. The calculation is done LSB-first, or "reflected".

### Connection Status Byte {#Connection_Status_Byte}

Response commands \$80, \$81, and \$82 send an enumerated value indicating the result when attempting to make a telephone connection. These enumerations are used clearly in the CPU2 ROM: 00, 01, 02, 03, 04, 05, 06, 07, 08, 0B, 0C, 0D, 0E. CPU2 ROM keeps track of this value at RAM location \$0052. \$00 indicates a successful connection, and all other values indicate a failure. Super Mario Club shows these failures with error code 40xx, xx directly reflecting this byte. Super Mario Club's manual gives troubleshooting information for most of these values. That information was used to create the table below.

<table class="wikitable">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>Status<br />
Byte</th>
<th>Error<br />
Code</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>$00</td>
<td>(n/a)</td>
<td>Connection was successful.</td>
</tr>
<tr>
<td>$01</td>
<td>4001</td>
<td>No dial tone was detected.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>電話線がはずれています。<br />
The telephone line is off.</p></td>
<td><ul>
<li>通信アダプタの電話回線の差し込みへ電話線を接続してください。</li>
<li>Connect the telephone line to the telephone line plug of the communication adapter.</li>
</ul>
<p><br />
</p>
<ul>
<li>回線テストのテスト1(29ページ) を行ってください。</li>
<li>Perform Test 1 (page 29) of the line test.</li>
</ul></td>
</tr>
<tr style="vertical-align: top;">
<td><p>既に電話が使われています。<br />
The phone is already in use.</p></td>
<td><p>通話が終わるまで待ってください。<br />
Please wait until the call ends.</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>ホー厶テレホン用のモジュラーコンセントなどにテレホンスイッチを取りつけていませんか?<br />
Is the telephone switch attached to a modular outlet for the Ho 厶 telephone?</p></td>
<td><p>テレホンスイッチがNTT の局線に直接取りつけられるように工事を行ってください。<br />
Make sure that the telephone switch can be attached directly to the NTT station line.</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>電話線は正常に接続されていますか？<br />
Is the telephone line connected normally?<br />
<br />
自宅の電話が使用中ではありませんか？<br />
Is your home phone in use?</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$02</td>
<td>4002</td>
<td>An incoming phone call is being received.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>電話がかかってきました。<br />
I got a call.</p></td>
<td><p>受話器をお取りください。<br />
Please pick up the handset.</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>もう一度ご利用ください。<br />
Please use it again.</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$03</td>
<td>4003</td>
<td>Pulse vs. Tone setting is incorrect.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>プッシュ回線とダイヤル回線の設定を誤っています。<br />
The push line and dial line settings are incorrect.</p></td>
<td><p>通信アダプタ裏の切換スイッチを回線に合わせてください。<br />
Set the changeover switch on the back of the communication adapter to match the line.<br />
<br />
</p>
<p>(番号をかけた時にプッシュ回線の場合はピ・ポ・パと鳴りますのでPBにセットしてください。それ以外はダイヤル回線ですので、20にセットしてください。</p>
<ul>
<li>部の地域では 10 にセットしていただく場合があります。)<br />
</li>
</ul>
<p>(If you use a push line when you dial a number, you will hear a beep, so set it to PB. Other than that, it is a dial line, so set it to 20.</p>
<ul>
<li>Set it to 10 in some areas. You may be asked to do so.)</li>
</ul></td>
</tr>
<tr style="vertical-align: top;">
<td><p>テレホンスイッチを新電電アダプタ等の後に取りつけていませんか?<br />
Is the telephone switch installed after the new electric adapter, etc.?</p></td>
<td><p>テレホンスイッチを新電電アダプタより前(NTTの局線側)に取りつけてください。<br />
Install the telephone switch in front of the new electric adapter (NTT station line side).</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>ご利用電話の種類の設定が誤っています。<br />
The phone type setting is incorrect.<br />
<br />
通信アダプタ裏の電話回線切換スイッチをセットしください。<br />
Set the telephone line selector switch on the back of the communication adapter.<br />
<br />
任天堂の通信アダプタセット取扱説明書をご覧下さい。<br />
Please refer to the Nintendo communication adapter set instruction manual.</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$04</td>
<td>4004</td>
<td>"0" Outgoing call setting is incorrect.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>0発信の設定が誤っています。<br />
The 0 outgoing call setting is incorrect.<br />
<br />
(0発信電話なのに0発信の設定がなされていません。 又は直通電話なのに0発信の設定がされています。)<br />
(0 outgoing call is not set even though it is a 0 outgoing call, or 0 outgoing call is set even though it is a direct call.)</p></td>
<td><p>登録内容を見直してください。(訂正する場合は、28ページをご参照ください。)<br />
Please review the registered contents. (Refer to page 28 for corrections.)<br />
<br />
</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>テレホンスイッチの取りつけが誤っています<br />
Wrong installation of telephone switch</p></td>
<td><p>通信機器の接続 (7,8ページ) をご確認ください。<br />
Check the connection of communication devices (pages 7 and 8).</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>電話回線は利用可能となっていますか？<br />
Is the telephone line available?<br />
<br />
ご利用の電話は「ゼロ発信」ではありませんか？<br />
Isn't your phone "zero call"?</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$05</td>
<td>4005</td>
<td>DDX-TP connection completed earlier than expected.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>DDX-TPの開通 (工事) が完了していないと思われます。<br />
It seems that the opening (construction) of DDX-TP has not been completed.<br />
<br />
</p></td>
<td><ul>
<li>回線テストのテスト2 (29ページ) を行ってください。(DDX-TP網IDの場合)</li>
<li>Perform Test 2 (page 29) of the line test. (For DDX-TP network ID)</li>
</ul>
<p><br />
</p>
<ul>
<li>NTTの113番に電話して確認してください。</li>
<li>Please call NTT 113 to confirm.</li>
</ul></td>
</tr>
<tr style="vertical-align: top;">
<td><p>テレホンスイッチの取りつけが誤っています。<br />
The telephone switch is installed incorrectly.</p></td>
<td><p>通信機器の接続(7, 8ページ) をご確認ぐださい。<br />
Check the connection of communication devices (pages 7 and 8).</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><ul>
<li>113に電話して、NTTにDDXの工事完了を確認して下さい。</li>
<li>Call 113 to confirm with NTT that the DDX construction has been completed.</li>
</ul></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$06</td>
<td>4006</td>
<td>Unknown.</td>
</tr>
<tr>
<td>$07</td>
<td>4007</td>
<td>Unknown.</td>
</tr>
<tr>
<td>$08</td>
<td>4008</td>
<td>Server was busy, disconnected, or rejected the password.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>センター側が混雑しています。<br />
The center side is crowded.</p></td>
<td><p>しばらく経ってから再度お試しください。<br />
Please try again after a while.</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>センターが停止しています。<br />
The center is stopped.</p></td>
<td><p>サービス時間を確認してください。<br />
Please check the service time.</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>DDX-TP の開通 (工事) が完了していないと思われます。<br />
It seems that the opening (construction) of DDX-TP has not been completed.</p></td>
<td><ul>
<li>回線テストのテスト2(29ページ) を行ってください。(DDX- TP綱IDの場合)</li>
<li>Perform Test 2 (page 29) of the line test. (For DDX-TP rope ID)</li>
</ul>
<p><br />
</p>
<ul>
<li>NTTの113番に電話して確認してください。</li>
<li>Please call NTT 113 to confirm.</li>
</ul></td>
</tr>
<tr style="vertical-align: top;">
<td><p>テレホンスイッチの取りつけが誤っています。<br />
The telephone switch is installed incorrectly.</p></td>
<td><p>通信機器の接続(7,8ページ) をご確認ください。<br />
Check the connection of communication devices (pages 7 and 8).</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>DDX-TPパスワードで、登録に誤りがあります。<br />
There is an error in the registration of the DDX-TP password.</p></td>
<td><p>パスワード番号とパスワード登録電話番号をご確認ください。<br />
Please check your password number and password registration phone number.</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:<br />
Same as 4005.</p></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$09</td>
<td>4009</td>
<td>Server was busy.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:<br />
Same as 4008.<br />
<br />
</p>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>相手方が話中です。しばらくして、再度かけ返して下さい。<br />
The other party is busy. Please call back again after a while.<br />
<br />
</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$0A</td>
<td>400A</td>
<td rowspan="3">DDX-TP registration incorrect or server disconnected.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>DDX-TPの登録内容に誤りがあります。<br />
There is an error in the registered contents of DDX-TP.</p></td>
<td><p>NTTの113番に電話して登録内容の確認をしてください。<br />
Please call NTT 113 to confirm your registration details.</p></td>
</tr>
<tr style="vertical-align: top;">
<td><p>センターが停止しています。<br />
The center is stopped.</p></td>
<td><p>サービス時間を確認してください。(15ページ)<br />
Please check the service time. (Page 15)</p></td>
</tr>
</tbody>
</table>
<p>Kangyou Sumimaru no Famicom Trade Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>原 因•対応方法 (Cause • Countermeasure)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>再度、お試し下さい。<br />
Please try again.<br />
<br />
</p>
<p>頻繁に発生する場合にはファミコントレード事務局へお電話下さい。<br />
If it occurs frequently, please call the NES Trade Secretariat.</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$0B</td>
<td>400B</td>
</tr>
<tr>
<td>$0C</td>
<td>400C</td>
</tr>
<tr>
<td>$0D</td>
<td>400D</td>
<td>Problem with telephone line connection.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:</p>
<table class="wikitable">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th>おもな原因 (Main Cause)</th>
<th>対処方法 (Workaround)</th>
</tr>
</thead>
<tbody>
<tr style="vertical-align: top;">
<td><p>通信アダプタから電話線がはずれました。<br />
The telephone line has been disconnected from the communication adapter.</p></td>
<td><p>通信アダプタと電話線の接続を確認してください。<br />
Check the connection between the communication adapter and the telephone line.</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table></td>
</tr>
<tr>
<td>$0E</td>
<td>400E</td>
<td>Problem with telephone line connection or server.
<table class="wikitable mw-collapsible mw-collapsed" role="presentation">
<colgroup>
<col style="width: 100%" />
</colgroup>
<tbody>
<tr>
<td>Reference Data </td>
</tr>
<tr>
<td><p>Super Mario Club Manual:<br />
Same as 400D.<br />
<br />
</p>
<p>Kangyou Sumimaru no Famicom Trade Manual:<br />
Same as 400A.</p></td>
</tr>
</tbody>
</table></td>
</tr>
</tbody>
</table>

# Pinouts {#Pinouts}

## RF5C66 Mapper and Disk Drive Controller {#RF5C66_Mapper_and_Disk_Drive_Controller}

                                                           _____
                                                          /     \
                                               CPU A0 -> / 1 100 \ -- +5Vcc
                                              CPU A1 -> / 2    99 \ -- n/c
                                             CPU A2 -> / 3      98 \ <> CPU D0
                                            CPU A3 -> / 4        97 \ <> CPU D1
                                           CPU A4 -> / 5          96 \ <> CPU D2
                                          CPU A5 -> / 6            95 \ <> CPU D3
                                         CPU A6 -> / 7              94 \ <> CPU D4
                                        CPU A7 -> / 8                93 \ <> CPU D5
                                      CPU A12 -> / 9                  92 \ <> CPU D6
                                     CPU A13 -> / 10                   91 \ <> CPU D7
                                    CPU A14 -> / 11                     90 \ -- GND
                                   /ROMSEL -> / 12                       89 \ <> Card D0
                                  CPU R/W -> / 13                         88 \ <> Card D1
                                      M2 -> / 14                           87 \ <> Card D2
              P6-1 Lid Switch, Card R/W <- / 15                             86 \ <> Card D3
              (20k resistor to 5Vcc) ? -> / 16                               85 \ <> Card D4
                                 /IRQ <- / 17                                 84 \ <> Card D5
                               +5Vcc -- / 18                                   83 \ <> Card D6
                                n/c -- / 19                                     82 \ <> Card D7
                  21.47727MHz Xtal -- / 20                                       81 \ -- +5Vcc
                             Xtal -- / 21                                            \
                             n/c -- / 22                                     O       /
                            GND -- / 23                                          80 / -- n/c
            (n/c) Xtal Osc Out <- / 24                                          79 / -> Exp P3-2
                          n/c -- / 25                                          78 / <- Exp P3-3
       ToneRx Xin, CIC Clock <- / 26              Nintendo RF5C66             77 / -> Exp P3-4
                        n/c -- / 27       Package QFP-100, 0.65mm pitch      76 / -> Exp P3-5
             (n/c) $40AE.0 <- / 28                                          75 / -> Exp P3-6
     CIC /CPU R/W Inhibit -> / 29                Mapper and                74 / <- Exp P3-7
          Key CIC-12 (?) <- / 30           Disk Drive Controller          73 / <- Exp P3-8
                           /       O                                     72 / <- Exp P3-9
                           \                                            71 / <- Exp P3-11
          Key CIC-10 (?) -> \ 31                                       70 / -- GND
           Key CIC-15 (?) -> \ 32                                     69 / -> 5A18-49
               CPU2 /Reset -> \ 33                                   68 / -> CPU2 /Reset (new rev)
        CHR RAM /CE (input) -> \ 34                                 67 / <> Exp P3-12           Orientation:
                     RAM +CE <- \ 35                               66 / <> Exp P3-13            --------------------
                (n/c) $40C0.1 <- \ 36                             65 / <> Exp P3-14                 80         51
     CPU2 /Reset (old rev: J2) <- \ 37                           64 / <> Exp P3-15                   |         |
                    CHR RAM /CE <- \ 38                         63 / <- $40B1.3 (n/c)               .-----------.
                             GND -- \ 39                       62 / <> Modem P4-31               81-|O Nintendo |-50
    Built-in RAM /CE ($6000-7FFF) <- \ 40                     61 / <> Modem P4-32                   |  RF5C66   |
          (n/c) ? /CE ($4xE0-4xEF) <- \ 41                   60 / <> Modem P4-29                100-|  GCD 4R  O|-31
           5A18-85 /CE ($4xD0-4xDF) <- \ 42                 59 / -- +5Vcc                           \-----------'
                             (GND) ? -> \ 43               58 / -- n/c                               |         |
                              (GND) ? -> \ 44             57 / -> Kanji ROM A17                     01         30
                               (GND) ? -> \ 45           56 / -> Kanji ROM A4
                                (GND) ? -> \ 46         55 / -> Kanji ROM A3          Legend:
                               CIRAM A10 <- \ 47       54 / -> Kanji ROM A2           ------------------------------
                                  PPU A11 -> \ 48     53 / -> Kanji ROM A1            --[RF5C66]-- Power
                                   PPU A10 -> \ 49   52 / -> Kanji ROM A0             ->[RF5C66]<- RF5C66 input
                 Kanji ROM /CE ($5000-5FFF) <- \ 50 51 / -- n/c                       <-[RF5C66]-> RF5C66 output
                                                \     /                               <>[RF5C66]<> Bidirectional
                                                 \   /                                    f      Famicom connection
                                                  \ /                                     r      ROM chip connection
                                                   V                                      R      RAM chip connection
    Notes:
    - +5Vcc pins 18, 59, 81, 100 are all connected together internally.
    - GND pins 23, 39, 70, 90 are all connected together internally.
    - 43, 44, 45, 46 are GND on the PCB, but have internal protection diodes from GND, suggesting they are logic pins.
      - Pins 45-46, when pulled high, causes oscillation on pin 56.
    - 24, 28, 36, 37, 41, 63 are n/c on the PCB, but function as noted.
    - /CE Pins 40, 41, 42, 50 behaviors:
      - Pin 40 (Built-in RAM /CE) activates low in range $6000-7FFF, regardless of CPU R/W, but only when M2 is high.
      - Pin 41 (Unknown /CE) always activates low in range $4xE0-4xEF, regardless of CPU R/W and M2.
      - Pin 42 (RF5A18 CPU2 /CE) always activates low in range $4xD0-4xDF, regardless of CPU R/W and M2.
      - Pin 50 (Kanji ROM /CE) activates low in range $5000-5FFF, regardless of CPU R/W.  (It appears to go low only when M2 is high but not specifically proven.)
    - Pin 29:
      - If the FNS has a CIC chip, CIC-11 drives this pin high after 24.8msec, and remains high as long as the CIC authentication is successful.  There is a low-pass filter in this case.
      - If the FNS does not have a CIC chip, the pin floats high.  There may be a pull-up resistor somewhere.  The delay to go high (if any) has not been measured.
      - When this pin is low, it resets pins 52-57 low and possibly lots of other things.
      - Card R/W is always high when this pin is low.
    - Pin 31 (CIC-10) and Pin 32 (CIC-15):
      - Both are jumpered direct to GND if the FNS does not have a CIC chip.
      - If it does have a CIC chip, these signals are both always low with or without tsuushin card inserted.
      - In no observed case are these signals ever high.
    - Pin 16 Pull-up of 20k to 5V is also required in order to avoid triggering reset.
    - Pin 16 seems to be related to pin 29.  With pin 29 floating and pin 16 pulled high at power on, the chip runs for 5 seconds, then enters reset.
    - Tested 10k instead of 20k (per original PCB) on pin 16, found no difference in time or function.
    - Pin 69 has a high pulse of 11.9085 usec at any time that register $4xAC has not been read for 12.4892 seconds.
      - Each additional 12.4892 seconds generates another pulse.
      - It has very repeatable precision, at least 6 figures on each.
      - It is not synchronized to M2 or any other inputs.
      - Note that 12.4892 sec * 21.47727 MHz = 2^28, with an error of 0.075%. (Nominal would be 12.4986 sec.)
      - Note that 11.9085 usec * 21.47727 MHz = 2^8, with an error of 0.093%. (Nominal would be 11.9196 usec.)
    - Pins 52-56 drive the address pins of the Kanji ROM.  (See notes below the LH5323M1 pinout.)
    - Pin 15 (Card R/W) is a non-inverted buffer of CPU R/W.  This signal connects through the lid switch.
    - Pin 26 puts out a 3.58 MHz square wave, ~50% duty.  This corresponds to 21.47727 MHz / 6.
    - Pin 79 (Exp 2) puts out a 95.95 kHz square wave, 93.7% duty.  Clock source unknown.
      - Note that this seems similar to FDS serial bitrate.
      - Standalone chip can get into a 341.2 kHz mode when touching pin 80, though pulling 80 high or low doesn't correlate.
      - Either frequency, the negative pulse width is 650 nsec.
    - Pins 71-79 appear strikingly similar to an FDS interface.
    - CIRAM A10 follows PPU A10 by default.

## <span id="RF5A18_CPU2_.2F_Modem_Controller"></span>RF5A18 CPU2 / Modem Controller {#RF5A18_CPU2_.2F_Modem_Controller}

                                                         _____
                                                        /     \
                                            CPU2 A0 <- / 1 100 \ -- GND
                                           CPU2 A1 <- / 2    99 \ <> CPU2 D0         
                                          CPU2 A2 <- / 3      98 \ <> CPU2 D1         
                                         CPU2 A3 <- / 4        97 \ <> CPU2 D2         
                                        CPU2 A4 <- / 5          96 \ <> CPU2 D3         
                                       CPU2 A5 <- / 6            95 \ <> CPU2 D4         
                                      CPU2 A6 <- / 7              94 \ <> CPU2 D5         
                                     CPU2 A7 <- / 8                93 \ <> CPU2 D6         
                                      +5Vcc -- / 9                  92 \ <> CPU2 D7         
                                   CPU2 A8 <- / 10                   91 \ -- +5Vcc
                                  CPU2 A9 <- / 11                     90 \ -> UART Tx (MSM6827L TXD)
                                CPU2 A10 <- / 12                       89 \ <- UART Rx (MSM6827L RXD)
                               CPU2 A11 <- / 13                         88 \ <- CPU A2
                              CPU2 A12 <- / 14                           87 \ <- CPU A1
                       (n/c) CPU2 A13 <- / 15                             86 \ <- CPU A0
                      (n/c) CPU2 A14 <- / 16                               85 \ <- /CE (5C66-42)
                     (n/c) CPU2 A15 <- / 17                                 84 \ <- P6-1 Lid Switch, Card R/W
                               GND -- / 18                                   83 \ <- M2
       (2.4576 MHz) (n/c) CPU2 M2 <- / 19                                     82 \ <> Card D7
                        CPU2 R/W <- / 20                                       81 \ <> Card D6
           RAM /CE ($0000-1FFF) <- / 21                                            \
    (n/c) ROM /CE ($C000-xFFF) <- / 22                                     O       /
          P5 /CE ($4129 Only) <- / 23                                          80 / <> Card D5
           (GND) CPU2 +Reset -> / 24                                          79 / <> Card D4
                    (GND) ? -> / 25                                          78 / -- GND
      /Internal ROM Enable -> / 26             Nintendo RF5A18              77 / <> Card D3
    (5C66-68) CPU2 /Reset -> / 27      Package QFP-100, 0.65mm pitch       76 / <> Card D2
      (10k up) CPU2 /IRQ -> / 28                                          75 / <> Card D1
     (10k up) CPU2 /NMI -> / 29             Modem Controller             74 / <> Card D0
                   n/c -- / 30                   CPU2                   73 / <- Tone Rx DV
                         /       O                                     72 / <- Tone Rx D8
                         \                                            71 / <- Tone Rx D4
                 +5Vcc -- \ 31                                       70 / <- Tone Rx D2
          MSM6827L DATA <> \ 32                                     69 / <- Tone Rx D1
           MSM6827L /INT -> \ 33                                   68 / -> Tone Rx GT
             MSM6827L /RD <- \ 34                                 67 / -> Exp P3-19
              MSM6827L /WR <- \ 35                               66 / <> Exp P3-18
             MSM6827L EXCLK <- \ 36                             65 / <> Exp P3-17
               (n/c) $4120.2 <- \ 37                           64 / -- +5Vcc                  Orientation:
                 MSM6827L AD1 <- \ 38                         63 / -> Modem P4-19             --------------------
                  MSM6827L AD0 <- \ 39                       62 / -> /Phone Audio Enable          80         51
                  (n/c) CPU2 D0 <- \ 40                     61 / -> /DTMF Output Enable            |         |
                         (n/c) ? <- \ 41                   60 / -> /Phone Off Hook                .-----------.
           (4.9152 MHz) Exp P3-16 <- \ 42                 59 / -> $4127.3 (n/c)                81-|O  RF5A18  |-50
                               GND -- \ 43               58 / -> /Green LED                       |  Nintendo |
                    19.6608MHz Xtal -- \ 44             57 / -> /Red LED                      100-|  GCD 8C  O|-31
                          1k to Xtal -- \ 45           56 / -> MSM6827L +Reset                    \-----------'
                                  GND -- \ 46         55 / <- Audio from phone line                |         |
                       (+5Vcc) $4126.0 -> \ 47       54 / <- Modem P4-28                          01         30
                     (+5Vcc) !($4126.1) -> \ 48     53 / <- Modem P4-25             
                                 5C66-69 -> \ 49   52 / <- Switch SW1-4             Legend:
                                      n/c -- \ 50 51 / <- Switch SW1-2              ------------------------------
                                              \     /                               --[RF5A18]-- Power, n/a
                                               \   /                                ->[RF5A18]<- RF5A18 input
                                                \ /                                 <-[RF5A18]-> RF5A18 output
                                                 V                                  <>[RF5A18]<> Bidirectional
    Notes:
    - This chip contains its very own 65C02 CPU, with built-in ROM.
    - +5Vcc pins 9, 31, 64, 91 are all connected together internally.
    - GND pins 18, 43, 46, 78, 100 are all connected together internally.
    - 24, 26 are GND on the PCB, but function as noted.
    - 25 is GND on the PCB, but has internal protection diode from GND, suggesting it is a logic pin.
    - 47, 48 are +5Vcc on the PCB, but function as noted.
    - 15, 16, 17, 19, 22, 37, 40, 59 are n/c on the PCB, but function as noted.
    - 41 is n/c on the PCB, but has protection diode from GND, suggesting it may have a function.
    - Pin 42 (Exp 16) puts out a 4.92 MHz square wave, ~50% duty.  This is 19.6608 MHz / 4.
    - CPU2 /Reset comes from RF5C66 pin 68 on new revisions and selectable with J1, J2 on old revisions:
      - J2 closed = RF5C66 pin 37 (default)
      - J1 closed = RF5C66 pin 68
    - Pin 24 prevents CPU2 functioning when held high at power-on.  If the pin is then driven low, the reset vector is then fetched after that.
      - Pin 24 can be freely used as a +reset at any time this way.
    - Pin 25 low at any time causes address bus to go to $FFFF and data bus shows a toggle on bits 2,5,6,7: period 208.7 usec, low for 7.93 usec.
      - Other data bits always low.
      - Shortly after applying power, the toggle has a lot of variations for a period of about 1.5 seconds, including a 225 msec gap where the bits are low.
      - The mentioned data bits all appear to have the same data.
      - Held low at power-on will fetch the reset vector later when driven high.
      - Held high at power-on, driven low later, enters the data bus toggle mode but:
        - Does not appear to fetch the reset vector when driven high after that.
        - Does execute code, possibly resuming from where it left off.
    - When pin 26 is set low (default case: This pin is tied directly to GND on the PCB):
      - Internal ROM is enabled in CPU2 range $E000-FFFF.
      - Open Bus in CPU2 range $C000-DFFF.
      - Pin 22 (ROM /CE) is enabled low in range $C000-DFFF
      - This mode allows ROM expansion at $C000-DFFF, with internal ROM (and its vector table) in place.
    - When pin 26 is set high:
      - Internal ROM is disabled, leaving open bus in CPU2 range $E000-FFFF.
      - Open Bus in CPU2 range $C000-DFFF.
      - Pin 22 (ROM /CE) is enabled low in the entire range $C000-FFFF.
      - This mode allows a totally custom external CPU2 ROM with its own interrupt vector table.

## LH5323M1 Kanji Graphic ROM {#LH5323M1_Kanji_Graphic_ROM_2}

                                  _____  Note: Flat spot does not correspond to pin 1.
                                 /     \
                         n/c -- / 12 11 \ -- n/c
               (5C66-52) A0 -> / 13   10 \ -- n/c
                    CPU D0 <> / 14      9 \ <- A1 (5C66-53)
                   CPU D1 <> / 15        8 \ <- A2 (5C66-54)
                  CPU D2 <> / 16          7 \ <- A3 (5C66-55)
                    GND -- / 17            6 \ -- GND
                CPU D3 <> / 18              5 \ <- A5 (CPU A0)
               CPU D4 <> / 19                4 \ -- n/c
              CPU D5 <> / 20                  3 \ <- A6 (CPU A1)
             CPU D6 <> / 21                    2 \ <- A7 (CPU A2)
            CPU D7 <> / 22  Nintendo LH5323M1   1 \ -- n/c
                     /        Package QFP-44       \
                     \         0.8mm pitch         /
               n/c -- \ 23                     44 / <- A8 (CPU A3)
                n/c -- \ 24   Kanji Graphic   43 / <- A13 (CPU A8)
           (GND) /OE -- \ 25       ROM       42 / <- A16 (CPU A11)
         (CPU A6) A11 -> \ 26               41 / <- A4 (5C66-56)         Orientation:
         (5C66-50) /CE -> \ 27             40 / -- n/c                   --------------------
                    GND -- \ 28           39 / -- n/c                        33         23
            (CPU A7) A12 -> \ 29         38 / -- +5Vcc                        |         |
             (CPU A5) A10 -> \ 30       37 / <- A17 (5C66-57 Bankswitch)     .-----------.
                       n/c -- \ 31     36 / <- A15 (CPU A10)              34-| Nintendo O|-22
                        n/c -- \ 32   35 / -- n/c                            |  CCR-01   |
                 (CPU A4) A9 -> \ 33 34 / <- A14 (CPU A9)                    | LH5323M1  |
                                 \     /                                  44-|O 9528 D   |-12
                                  \   /                                      '-----------/
                                   \ /                                        |         |
                                    V                                        01         11

    Notes:
    - 6 & 28 are connected together internally.
    - 17 has no measurable connection to 6 & 28.
    - All logic pins have protection diode from pin 17, suggesting this is the true GND.
    - Pin 25 also appears as a logic pin with respect to pin 17.
    - When pins 25 and 27 are both driven low, the data bus becomes an output.  Otherwise it is hi-z.
    - Pins 13, 9, 8, 7, 41, 37 are controlled by the RF5C66.
      - Pins 13, 9, 8, 7, 41 are controlled with auto-increment function.
      - The value of these pins increments each M2 falling edge when the CPU is in range $5000-5FFF.
      - Pin 37 is a bankswitch, controlled by register $40B0.0
      - At reset and when reading from register $40B0, these pins reset to 0.
      - The conditions resetting or maintaining the bankswitch pin to 1 are still unknown.

## 8633 Famicom Network System CIC Key Chip {#8633_Famicom_Network_System_CIC_Key_Chip}

Unlike the NES console, the Famicom Network System appears to have the CIC key.

                                     _______   _______
                                     |      \_/      |
     (To Card CIC Pin 2) Data Out <- | 1          18 | -- +5Vcc
    (From Card CIC Pin 1) Data In -> | 2  O       17 | -- n/c
                              n/c -- | 3   8633   16 | -- n/c
                              n/c -- | 4          15 | -> ? (5C66-32) always observed low
                              n/c -- | 5    CIC   14 | -- n/c
                              n/c -- | 6    Key   13 | -- n/c
                            Clock -> | 7          12 | <- ? (5C66-30)
    (From Card CIC Pin 11) +Reset -> | 8    U8    11 | -> /CPU R/W Inhibit (5C66-29)
                              GND -- | 9          10 | -> ? (5C66-31) always observed low
                                     |_______________|

- When the CIC key drives pin 11 low, this stops operation of the Famicom Network System by means of holding Card R/W high.
- The clock is 3.58 MHz, coming from RF5C66 pin 26.

## 8634A Tsuushin Card CIC Lock Chip {#8634A_Tsuushin_Card_CIC_Lock_Chip}

Unlike the NES cartridge, the tsuushin card appears to have the CIC lock.

                                     _______   _______
                                     |      \_/      |
      (To FNS CIC Pin 2) Data Out <- | 1          18 | -- +5Vcc
     (From FNS CIC Pin 1) Data In -> | 2  O       17 | -- n/c
                              n/c -- | 3   8634A  16 | ?? GND
                              n/c -- | 4          15 | -- n/c
                              n/c -- | 5    CIC   14 | -- n/c
                              n/c -- | 6   Lock   13 | ?? +5V
                            Clock -> | 7          12 | ?? Card-33, n/c in Famicom Network System
               (Cap to 5V) +Reset -> | 8          11 | -> CIC Key +Reset (To FNS CIC Pin 8)
                              GND -- | 9          10 | -- n/c
                                     |_______________|

- +Reset is connected with a ceramic capacitor to 5V. This gives a momentary positive pulse at power-on.
- The clock is 3.58 MHz, coming from RF5C66 pin 26.
- Note: Some assumptions made on CIC chips based on similarity to F411A from Super NES.

## 8kByte CHR RAM {#8kByte_CHR_RAM}

                       _______   _______
                       |      \_/      |
               n/c? -- | 1          28 | -- +5Vcc
            PPU A12 -> | 2  O       27 | <- PPU /WR
             PPU A7 -> | 3          26 | <- +CE: U3=RF5C66 34/38, U4=PPU /A13
             PPU A6 -> | 4          25 | <- PPU A8
             PPU A5 -> | 5  LH5268  24 | <- PPU A9
             PPU A4 -> | 6    CHR   23 | <- PPU A11
             PPU A3 -> | 7    RAM   22 | <- /OE: PPU /RD
             PPU A2 -> | 8   U3/U4  21 | <- PPU A10
             PPU A1 -> | 9          20 | <- /CE: U3=PPU A13, U4=RF5C66 34/38
             PPU A0 -> | 10         19 | <> PPU D7
             PPU D0 <> | 11         18 | <> PPU D6
             PPU D1 <> | 12         17 | <> PPU D5
             PPU D2 <> | 13         16 | <> PPU D4
                GND -- | 14         15 | <> PPU D3
                       |_______________|

## 8kByte W-RAM {#8kByte_W-RAM}

                       _______   _______
                       |      \_/      |
               n/c? -- | 1          28 | -- +5Vcc
            CPU A12 -> | 2  O       27 | <- /WR: Card R/W (P6-2 Lid Switch)
             CPU A7 -> | 3          26 | <- +CE: RAM +CE
             CPU A6 -> | 4          25 | <- CPU A8
             CPU A5 -> | 5  LH5268  24 | <- CPU A9
             CPU A4 -> | 6 Built-in 23 | <- CPU A11
             CPU A3 -> | 7  W-RAM   22 | <- /OE: GND
             CPU A2 -> | 8    U5    21 | <- Card A10
             CPU A1 -> | 9          20 | <- /CE: Built-in RAM /CE
             CPU A0 -> | 10         19 | <> Card D7
            Card D0 <> | 11         18 | <> Card D6
            Card D1 <> | 12         17 | <> Card D5
            Card D2 <> | 13         16 | <> Card D4
                GND -- | 14         15 | <> Card D3
                       |_______________|

## 8kByte CPU2 RAM {#8kByte_CPU2_RAM}

                       _______   _______
                       |      \_/      |
               n/c? -- | 1          28 | -- +5Vcc
           CPU2 A12 -> | 2  O       27 | <- /WR: CPU2 R/W
            CPU2 A7 -> | 3          26 | <- +CE: +5Vcc
            CPU2 A6 -> | 4          25 | <- CPU2 A8
            CPU2 A5 -> | 5  LH5268  24 | <- CPU2 A9
            CPU2 A4 -> | 6   CPU2   23 | <- CPU2 A11
            CPU2 A3 -> | 7    RAM   22 | <- /OE: GND
            CPU2 A2 -> | 8    U6    21 | <- CPU2 A10
            CPU2 A1 -> | 9          20 | <- /CE: CPU2 RAM /CE
            CPU2 A0 -> | 10         19 | <> CPU2 D7
            CPU2 D0 <> | 11         18 | <> CPU2 D6
            CPU2 D1 <> | 12         17 | <> CPU2 D5
            CPU2 D2 <> | 13         16 | <> CPU2 D4
                GND -- | 14         15 | <> CPU2 D3
                       |_______________|

## P4: Modem Module Edge Connector {#P4:_Modem_Module_Edge_Connector}

           Famicom     | Modem  |    Famicom
        Network System | Module | Network System
                       __________
                       |        |
              +5Vcc -- | 1   19 | <- 5A18-63
    MSM6827L +Reset -> | 2   20 | <- Tone Rx GT
       MSM6827L AD0 <> | 3   21 | <- /Phone Audio Enable
                GND -- | 4   22 | -> Tone Rx DV
       MSM6827L AD1 <> | 5   23 | -> 5A18-55, Audio from phone line
       MSM6827L RXD <- | 6   24 | <- Tone Rx Xin, from 5C66-26
      MSM6827L DATA <- | 7   25 | -> 5A18-53
       MSM6827L TXD -> | 8   26 | -- GND
       MSM6827L /WR -> | 9   27 | <- /Phone Off Hook
     MSM6827L EXCLK -> | 10  28 | -> 5A18-54
       MSM6827L /RD -> | 11  29 | <> 5C66-60
              +5Vcc -- | 12  30 | <- /DTMF Output Enable
         Tone Rx D1 <- | 13  31 | <> 5C66-62           __________________________
      MSM6827L /INT <- | 14  32 | <> 5C66-61           | Modem Module           |
              +5Vcc -- | 15  33 | <- Audio from 2A03   | Orientation            |
         Tone Rx D2 <- | 16  34 | -> Audio to RF       |                        |
         Tone Rx D8 <- | 17  35 | -- GND               |    19 _____________ 36 |
         Tone Rx D4 <- | 18  36 | -- GND               |     1 |___________| 18 |
                       |________|                      |________________________|

    Note: The modem module uses modem chip Oki MSM6827L and Dual Tone Receiver MC14LC5436P.

## P2: Tsuushin Card Connector {#P2:_Tsuushin_Card_Connector}

Note that the tsuushin card may appear to have a metric 1mm pin pitch, but in fact it has an imperial 0.040" (40 thousandths) pin pitch.

    Card |  | Famicom Network System
    -----+--+-------------------------
      1  |--| +5Vcc
      2  |--| +5Vcc
      3  |??| n/c in JRA-PAT card, n/c in FNS
      4  |??| n/c in JRA-PAT card, FNS has 10k pull-up only.
      5  |<>| Card D0
      6  |<>| Card D1
      7  |<>| Card D2
      8  |<>| Card D3
      9  |<>| Card D4
     10  |<>| Card D5
     11  |<>| Card D6
     12  |<>| Card D7
     13  |<-| Card R/W (P6-2 Lid Switch)
     14  |<-| M2
     15  |<-| /ROMSEL
     16  |<-| CPU A0
     17  |<-| CPU A1
     18  |<-| CPU A2
     19  |<-| CPU A3
     20  |<-| CPU A4
     21  |<-| CPU A5
         |  |
         |  |
     22  |<-| CPU A6
     23  |<-| CPU A7
     24  |<-| CPU A8
     25  |<-| CPU A9
     26  |<-| CPU A10
     27  |<-| CPU A11
     28  |<-| CPU A12
     29  |<-| CPU A13
     30  |<-| CPU A14
     31  |??| n/c in JRA-PAT card, FNS has 10k pull-up only.
     32  |??| n/c in JRA-PAT card, FNS has 10k pull-up only.
     33  |??| connected to Card Lock CIC-12 in JRA-PAT, n/c in FNS
     34  |??| n/c in JRA-PAT card, n/c in FNS
     35  |->| CIC Key Reset (Card Lock CIC-11 -> FNS Key CIC-8)
     36  |->| CIC Lock-to-Key Data (Card Lock CIC-1 -> FNS Key CIC-2)
     37  |<-| CIC Key-to-Lock Data (Card Lock CIC-2 <- FNS Key CIC-1)
     38  |<-| CIC Clock (3.58 MHz, from 5C66.26)
     39  |??| n/c in JRA-PAT card, FNS has 10k pull-up only.
     40  |<-| RAM +CE (n/c in JRA-PAT card)
     41  |--| GND
     42  |--| GND

## P3: Expansion Connector {#P3:_Expansion_Connector}

                     Outside    |  FNS  |    Outside                   _____________________
                                _________                             / Orientation        /|
                                |       |                            /____________________/ |
                        /IRQ -> | 1  20 | -- +5Vcc                   |o|_| ==      |_||_|o|/
    (95.94kHz Clock) 5C66-79 <- | 2  19 | -> 5A18-67                 \_  _  _ _||_  _  _ _/|
                     5C66-78 -> | 3  18 | -> 5A18-66                  |-| |    || HVC-050| |
                     5C66-77 <- | 4  17 | -> 5A18-65                  |-|_|    ||        | |
                     5C66-76 <- | 5  16 | -> 5A18-42 (4.92MHz Clock)  |        ||        | |
                     5C66-75 <- | 6  15 | <> 5C66-64                  |  20 __/__\__ 11  | |
                                |       |                             |o  1 |______| 10 o| |
                     5C66-74 -> | 7  14 | <> 5C66-65                  | ________________ | |
                     5C66-73 -> | 8  13 | <> 5C66-66                  |/_______________/|| |
                     5C66-72 -> | 9  12 | <> 5C66-67                  ||______________|/ | |
                         GND -- | 10 11 | <- 5C66-71                  |                  | |
                                |_______|                             |      ______      | |
                                                                      |o    | |    |    o| |
                                                                      \_____|/     |_____|/

## P5: Expansion Connector {#P5:_Expansion_Connector}

Note: This connector only exists on old revisions of Famicom Network System. Expansion P5 /CE is activated low specifically at CPU2 address \$4129.

    Outside |  | Famicom Network System
    --------+--+-------------------------
        9   |--| GND
        8   |<>| CPU2 D5
        7   |<>| CPU2 D4
        6   |<>| CPU2 D3
        5   |<>| CPU2 D2
        4   |<>| CPU2 D1
        3   |<>| CPU2 D0
        2   |<-| Expansion P5 /CE
        1   |--| +5V

## CN07: Dataship 1200 AV Port {#CN07:_Dataship_1200_AV_Port}

    Outside    |FNS|    Outside
                ___
               /   |
      +5Vcc -- |1|2| -> audio
        n/c -- |3|4| -> video
        GND -- |5|6| -- GND
               \___|

The Dataship 1200 uses the same port for AV as the Game Boy link cable. The official AV cable uses pins 2, 4, and 5. The official RF cable uses pins 1, 2, 4, 5, and 6. Note that pin 4 is frequently missing from link cables, so care must be taken to find a compatible aftermarket cable when making a reproduction AV cable.

# See Also {#See_Also}

- [Famicom Network Controller](../Famicom_Network_Controller/)
- <https://forums.nesdev.org/viewtopic.php?f=9&t=18530>
- <https://niwanetwork.org/wiki/Tsuushin_Cartridge>
- <https://niwanetwork.org/wiki/List_of_Famicom_Network_System_software>
