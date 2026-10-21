---
title: "Famicom Network Controller"
url: "/wiki/Famicom_Network_Controller/"
description: "The Famicom Network Controller (HVC-051) is a 23-button controller for use with the Famicom Network System. It is essentially a standard controller that adds a number pad, and can be used with normal…"
wikiPage: "Famicom_Network_Controller"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

<figure class="mw-halign-right" typeof="mw:File/Thumb">
<a href="../../wiki-images/Nintendo-Famicom-Modem-Controller.jpg"><img src="../../wiki-images/Nintendo-Famicom-Modem-Controller.jpg" width="400" height="216" /></a>
<figcaption>Famicom Network Controller</figcaption>
</figure>

The Famicom Network Controller (HVC-051) is a 23-button controller for use with the [Famicom Network System](../Family_Computer_Network_System/). It is essentially a [standard controller](../Standard_controller/) that adds a number pad, and can be used with normal games.

## Layout {#Layout}

     .------------------------------------------.
     | (<)   (>)  (1) (2) (3) (*) (C)           |
     | SEL _ STA                        (END)   |
     |   _| |_    (4) (5) (6) (#) (.)           |
     |  |_   _|                                 |
     |    |_|     (7) (8) (9) (  0  )  (B)  (A) |
     '------------------------------------------'

(ASCII art courtesy Nocash)

- The Select button has the additional text 前ぺージ = Previous Page
- The Start button has the additional text 次ページ = Next Page
- The B button has the additional text 目次 = Table of Contents
- The A button has the additional text 実行 = Execute
- The button marked "END" instead has the text 通信終了 = End of Communication

## Protocol {#Protocol}

The protocol looks like a standard Famicom [expansion port](../Expansion_port/) controller, returning its result via reads from \$4016 D1. After the initial eight reads, the following sixteen reads return the following bits in order:

     0-7 - see Standard controller
     8 - 0
     9 - 1
    10 - 2
    11 - 3
    12 - 4
    13 - 5
    14 - 6
    15 - 7
    16 - 8
    17 - 9
    18 - *
    19 - #
    20 - .
    21 - C
    22 - (Always 0)
    23 - 通信終了

Like the standard controller, reads after the first 24 read as logic 1.

## See also {#See_also}

- [Super Famicom NTT Data Keypad](../Super_Famicom_NTT_Data_Keypad/) - The equivalent Super Famicom controller with a number pad.

## Sources {#Sources}

- [Raphnet](https://www.raphnet.net/divers/famicom_network_controller_hvc_051/index_en.php)
- [astro187 on the forum](https://forums.nesdev.org/viewtopic.php?p=247107#p247107)
