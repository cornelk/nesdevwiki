---
title: "TNS-HFC"
url: "/wiki/TNS-HFC/"
description: "The TNS-HFC is a series of hardware NSF players that was produced starting in the early 2000s by Terra Network Systems. Production seems to have stopped by 2018, due to the creator's sickness. These…"
wikiPage: "TNS_HFC"
---

The TNS-HFC is a series of hardware [NSF](../NSF/) players that was produced starting in the early 2000s by Terra Network Systems. Production seems to have stopped by 2018, due to the creator's sickness. These are best known for being the only produced devices to successfully play back expansion sound NSF files on all known expansion chips as well as offer playback on multiple expansion chips at once. They don't require that the host Famicom be connected to a TV or video display; all NSFs can be played and controlled using the standard 1P controller and the buttons on the device itself. Audio is routed through a headphone jack on the cartridge.

These devices are incompatible with unmodified NTSC and PAL NES systems because they don't [wait for the PPU to finish resetting](../PPU_power_up_state/#Best_practice) ([NMIs](../NMI/) are used to tick the play routine). An NES can be modified to support them by tying the PPU reset pin high, mimicking the Famicom reset circuit.

## Known models {#Known_models}

### TNS-HFC1 {#TNS-HFC1}

Earliest known model. Unlike later models it's most unique for its PCB shape not conforming to the form factor of any known cartridge, so it can't be placed inside a cartridge shell for protection. By default, it only supports 2A03-only 32KB non-bankswitched NSFs due to its basic operation; An Atmel ATMEGA8L controls the address bus to load an NSF into the 32KB SRAM present on the cartridge, and once loaded the system waits for the RESET button being pressed to hand off bus control to the 2A03 so it can play back the loaded NSF.

A spot is allocated for allowing playback on external sound cartridges, however only MMC5 and VRC7 audio were initially supported with the v1.05 firmware, and then later N163 audio in the 1.50 firmware update. However correct playback could not be guaranteed due to a lack of full understanding of the proper access times such as those on VRC7's audio registers.

### TNS-HFC1-E {#TNS-HFC1-E}

Hardware revision of the TNS-HFC1. Otherwise is restricted to the same non-banked 32KB restriction as the standard HFC1.

It was given a form factor allowing it to fit inside of a "large shell" Famicom cartridge shell with some slight modifications.

Much like its prior incarnation, the end user must supply their own 60-pin connector if they want to play NSFs using external sound cartridges. Additionally, they must alter the configuration of two hardware jumpers in order to correctly route 2A03 audio to the expansion cart. Later runs included the connector, but still required the jumper configuration.

In both this model and the prior HFC1, NSFs ripped from games, even those without bankswitching, were often unsupported, often cited for their unusual loading address locations. In essence, this meant that the device was more focused towards modern chiptuning and compositions that were specifically made for the device, and not simply for listening to legacy soundtrack rips.

The original v0.98 firmware revision likely inherited VRC7 cartridge timing writes akin to the prior model, and later firmware updates were aimed at making the device more usable for general NSF use rather than specific chiptuning applications.

### TNS-HFC2 {#TNS-HFC2}

One of the most popular models, this unit by itself only supports 2A03 sound, but introduces NSF bankswitching, which would become standard across later models. Its shorter form factor allows it to fit into a modified standard Famicom game cartridge shell.

Without an addon device, this device only supports bankswitched NSFs with a maximum size of 512KB.

### TNS-HFC3 {#TNS-HFC3}

Another rather popular model, modeling itself after the past TNS-HFC1-E's form factor.

This device supports 1024KB/1MB NSFs that are bankswitched, but does not support FDS audio without an addon device.

Two DIP switches are present on the front side of the PCB, being an MMC5 write protect and a switch to swap the A0 and A1 lines, but only when using VRC6 audio. By default, it's configured to the "alternative" configuration of Esper Dream 2 and Madara.

Next to the serial code region on the PCB is silkscreen options indicating a 512KB option as well as the external sound slot being an optional addition rather than included. It's unknown if any TNS-HFC3 unit was ever produced with these options.

### TNS-HFC4 {#TNS-HFC4}

Models itself after the form factor of the Famicom Disk System's RAM adapter. Supports 1024KB/1MB NSFs just like the TNS-HFC3. Has an onboard RP2C33 for playback of FDS audio, as well as an optional adjustable mixer, and is one of the first models to support CHR-RAM, for NSFs that want to use visuals alongside sound.

It retains the external sound cartridge feature of the TNS-HFC3, except it loses the VRC6 addressing swap functionality, but it can be restored with a hardware modification.

### TNS-HFC5 {#TNS-HFC5}

Uses the TNS-HFC2 form factor, but retains the memory upgrades and CHR-RAM functionality of the TNS-HFC4. Has a socket for an optional YMZ294, but otherwise does not support expansion sound without an addon device. Swaps the dual analog volume wheels for the Left and Right headphone outputs with digital button controls.

Additionally also supports an optional small LCD for displaying NSF metadata information.

### TNS-HFP1 {#TNS-HFP1}

Last known released model. Unlike all prior units which required an existing Famicom, this one is standalone by including a CPU/PPU within the system, and is powered either by an internal battery or external DC. Otherwise, it is basically like a TNS-HFC5 albeit standalone and portable.

## Official Addons {#Official_Addons}

Items produced and sold by Terra Network Systems itself.

### TNS-HFE4 {#TNS-HFE4}

Sometimes referred to as "The Stack" by some of its users, it is a multislot sound expansion device, intended for use with the TNS-HFC2 (although likely compatible with later models). Includes a socket for a YMZ294 to simulate Sunsoft 5B audio, though the initial release used a 3.58 MHz crystal resulting in a recall to add a divider module for correct operation.

Likely also formed the basis for the design of the unofficial APH-001 addon.

### TNS-HFS1 {#TNS-HFS1}

Sound cartridge simply containing a YMZ294-D SSG to stand in for Sunsoft 5B audio and a XILINX CPLD for simulation of Namco 163 audio (erroneously referred to as "N106" on the PCB). Additionally contains a DIP switch to control which expansion audio option to use (both cannot be used simultaneously), and to enable the expansion cart's CHR-RAM capability, and also has two customizable trimpots to control the mixing level of both the incoming 2A03 audio and the expansion audio, before mixing both.

Although likely compatible with any TNS-HFC model with an external sound slot, it was likely meant for the TNS-HFC3, which would give it CHR-RAM like the later TNS-HFC4 that succeeded it.

### TNS-HFX4 {#TNS-HFX4}

Supports playback of nearly every common sound expansion at once. One slot too few prevents all expansions from being simultaneously usable unless a TNS-HFC4 (for 2C33) is used to free up the final slot. It additionally includes an on-board YM2413 and YMZ294 to substitute VRC7 and 5B sound cartridges, genuine 5B cartridges are non-operational. Outputs audio on its own headphone jack, but a modification can be made to route the audio back into the host TNS-HFC unit's headphone jack.

The companion TNS-HFE1C riser board has silk screening that references a TNS-HFM that is not known to have ever been produced.

## Unofficial Addons {#Unofficial_Addons}

Items for the TNS-HFC standard not produced or sold by Terra Network Systems.

### <span id="APH-001_.22Nishizumi.22"></span>APH-001 "Nishizumi" {#APH-001_.22Nishizumi.22}

Produced by [Applesorce](https://x.com/applesorce). Currently out of production, but schematics are [publically available on Github](https://github.com/applesorce/nishizumi). Works similarly to the TNS-HFX4, but supports all expansion chips at once as well as individual mixing levels of each chip, but does not provide an alternative VRC7 or Sunsoft 5B sound option. A firmware upgrade to the TNS-HFC3 means any device running the upgrade is incompatible with this addon.

Its boxart, a rendition of Miho Nishizumi from the anime Girls und Panzer, was drawn by [RIKI](https://x.com/riki2riki).

### <span id="APH-002_.22Akiyama.22"></span>APH-002 "Akiyama" {#APH-002_.22Akiyama.22}

Produced by [Applesorce](https://x.com/applesorce). Simply a card edge conversion for RAM Adapter PCBs from flat to upright, intended for use with Nishizumi as the spacing between slots is too narrow.

### HVC-EMX {#HVC-EMX}

Designed by forple, PCB produced by Lockster. A small interposer board that replaces a cartridge's own expansion audio mixing with an externally controllable one.

Primarily was designed to allow Namco 129 cartridges the ability to be used alongside 2A03 audio, however its design also meant that Namco 163/160 titles that don't use expansion sound could also use expansion sound (bypassing their "shorted" audio loop), and other expansion sound cartridges gain controllable mixing.

### TNS-ESG1 {#TNS-ESG1}

Designed by forple and Lockster. A sound cartridge compatible with the Sunsoft 5B audio registers, it exchanges the 5B for allowing AY8930 and YM2149 chips to be used instead. It additionally uses an inverter amp circuit with a 74HCU04 to mimic the real Sunsoft 5B's inverter circuit. Design is not finalized.

In a sense, it can be considered a Sunsoft 5B audio-focused variant of the TNS-HFS1.

## Playback Examples {#Playback_Examples}

TNS-HFC1(-E):

- <https://www.nicovideo.jp/watch/sm5249681>

TNS-HFC4:

- <https://youtu.be/kqb6FbYzorM>

TNS-HFC5:

- <https://youtu.be/AEtSmswGbnk>

TNS-HFC + Nishizumi:

- [Introduction video](https://www.nicovideo.jp/watch/sm21501631)
- <https://www.nicovideo.jp/watch/sm19653477>
- <https://www.nicovideo.jp/watch/sm23737913>

TNS-HFX4:

- <https://www.nicovideo.jp/watch/sm28809008>

## External Links {#External_Links}

- [Official Website](https://www2s.biglobe.ne.jp/~tns/)
- [TNS-HFC Series Register Map (English)](https://carol6502.neocities.org/TNSREGISTERMAP_ENG)
- [APH-001 "Nishizumi" Github](https://github.com/applesorce/nishizumi)
