---
title: "Doremikko Keyboard"
url: "/wiki/Doremikko_Keyboard/"
description: "The Doremikko Keyboard is a 36-key piano keyboard released by Konami alongside the pack-in FDS game DoReMikko."
wikiPage: "Doremikko_Keyboard"
wikiCategories:
  - name: "Controllers"
    url: "/category/Controllers/"
---

The Doremikko Keyboard is a 36-key piano keyboard released by Konami alongside the pack-in [FDS](../Family_Computer_Disk_System/) game *DoReMikko*.

The keyboard uses common matrix logic similar to the [Family BASIC Keyboard](../Family_BASIC_Keyboard/) but with diodes for [N-key rollover](https://en.wikipedia.org/wiki/N-key_rollover) (as typical for musical applications). The keyboard does not contain any sound hardware.

## Keyboard map {#Keyboard_map}

     _____________________________________________
    |  _________________________________________  |
    | | U U U | U U | U U U | U U | U U U | U U | |
    | | U U U | U U | U U U | U U | U U U | U U | | Keys
    | | | | | | | | | | | | | | | | | | | | | | | |
    |_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|_|
       F G A B C D E F G A B C D E F G A B C D E    Notes
      |-------><------------><------------><----|   Octaves
        1..7        8..19    |   20..31    32..36   Key Numbers
                             |
                          Middle C

(ASCII art courtesy nocash)

There are 21 white keys and 15 black keys for a total of 36 keys spanning two + two-half octaves.

## Interface {#Interface}

### <span id="Input_.28.244016_write.29"></span>Input (\$4016 write) {#Input_.28.244016_write.29}

    7  bit  0
    ---- ----
    xxxx xxRI
           ||
           |+-- Increment row (TODO: 0->1 or 1->0 transition?)
           +--- 1: Reset to first row (TODO: does this also reset the column?)

### <span id="Output_.28.244017_read.29"></span>Output (\$4017 read) {#Output_.28.244017_read.29}

    7  bit  0
    ---- ----
    xxxK KKKx
       | |||
       +-+++--- Receive key status of currently selected row/column.

Reading this register toggles between the two columns. Any key that is held down, will read back as 1.

*DoReMikko* doesn't appear to perform long manual delays during the polling process, so there may not be any parasitic capacitances to account for compared to other devices with common matrix logic.

## Matrix {#Matrix}

This below matrix maps the row/column to their corresponding key number. Blank entries have unknown contents, and rows beyond 6 are yet to be determined.

<table class="wikitable">
<thead>
<tr>
<th></th>
<th colspan="4">Column 0</th>
<th colspan="4">Column 1</th>
</tr>
</thead>
<tbody>
<tr>
<th>$4017 bit</th>
<th>4</th>
<th>3</th>
<th>2</th>
<th>1</th>
<th>4</th>
<th>3</th>
<th>2</th>
<th>1</th>
</tr>
&#10;<tr>
<th>Row 0</th>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>2</td>
<td>1</td>
</tr>
<tr>
<th>Row 1</th>
<td>6</td>
<td>5</td>
<td>4</td>
<td>3</td>
<td></td>
<td></td>
<td>8</td>
<td>7</td>
</tr>
<tr>
<th>Row 2</th>
<td>12</td>
<td>11</td>
<td>10</td>
<td>9</td>
<td></td>
<td></td>
<td>14</td>
<td>13</td>
</tr>
<tr>
<th>Row 3</th>
<td>18</td>
<td>17</td>
<td>16</td>
<td>15</td>
<td></td>
<td></td>
<td>20</td>
<td>19</td>
</tr>
<tr>
<th>Row 4</th>
<td>24</td>
<td>23</td>
<td>22</td>
<td>21</td>
<td></td>
<td></td>
<td>26</td>
<td>25</td>
</tr>
<tr>
<th>Row 5</th>
<td>30</td>
<td>29</td>
<td>28</td>
<td>27</td>
<td></td>
<td></td>
<td>32</td>
<td>31</td>
</tr>
<tr>
<th>Row 6</th>
<td>36</td>
<td>35</td>
<td>34</td>
<td>33</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
</tbody>
</table>

## References {#References}

- [nocash's documentation](https://problemkaputt.de/everynes.htm#controllerspianokeyboards)
- [Forum thread with board photos](https://forums.nesdev.org/viewtopic.php?p=300884)
