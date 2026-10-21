---
title: "1993 Super 50 in 1 Game"
url: "/wiki/1993_Super_50_in_1_Game/"
description: "The cartridge is a multigames cartridge dated from 1993."
wikiPage: "1993_Super_50_in_1_Game"
---

## Description {#Description}

The cartridge is a multigames cartridge dated from 1993.

Versions known :

- 1993 Super 50 in 1 Game : NES version (72-pin)

## Hardware {#Hardware}

| NES PCB (Front mirrored picture) | NES PCB (Bottom picture) |
|----|----|
| [<img src="../../wiki-images/50in1-MG109-frontside-inverse.jpg" width="600" height="554" alt="50 in 1 NES PCB (Front mirrored picture)" />](../../wiki-images/50in1-MG109-frontside-inverse.jpg) | [<img src="../../wiki-images/50in1-MG109-bottomside.jpg" width="600" height="555" alt="50 in 1 NES PCB (Bottom picture)" />](../../wiki-images/50in1-MG109-bottomside.jpg) |

<table class="wikitable" data-align="center" style="margin: auto;">
<thead>
<tr>
<th colspan="2">Datas chips</th>
</tr>
</thead>
<tbody>
<tr>
<th>Chips</th>
<th>Description</th>
</tr>
&#10;<tr>
<td>Blob</td>
<td>PRG-ROM</td>
</tr>
<tr>
<td>Blob</td>
<td>CHR-ROM</td>
</tr>
<tr>
<td colspan="2">Usual fonctionnal chips</td>
</tr>
<tr>
<td>Chips</td>
<td>Description</td>
</tr>
<tr>
<td>74LS161A - 6.70</td>
<td>BCD Decade counters / 4-bit binary counters</td>
</tr>
<tr>
<td>GD74LS00 - 9223 - GoldStar</td>
<td>NAND Gate, LS Series, 4-Func, 2-Input, TTL, PDIP14</td>
</tr>
</tbody>
</table>

[<img src="../../wiki-images/CART496_50_in_1_schematic_0.png" width="1400" height="927" alt="50 in 1 Schematic" />](../../wiki-images/CART496_50_in_1_schematic_0.png)

## Software {#Software}

The 50 in 1 rom contains 64 KB of PRG and 32 KB of CHR.

The rom has been dumped and assigned to the [INES Mapper 200](../INES_Mapper_200/).

It contains 4 different games that can be beginned at different levels.

## Dumping method {#Dumping_method}

This script can dump the cartridge with the INLretro-prog dumper (Device firmware version: 2.3.x) :

<table class="wikitable" data-align="center" style="margin: auto;">
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr>
<th>dumpMG109.lua</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="../../wiki-images/DumpMG109.lua.txt">File:DumpMG109.lua.txt</a></td>
</tr>
<tr>
<td><pre><code>-- create the module&#39;s table
local dumpmg109 = {}
&#10;-- import required modules
local dict = require &quot;scripts.app.dict&quot;
local nes = require &quot;scripts.app.nes&quot;
local dump = require &quot;scripts.app.dump&quot;
local flash = require &quot;scripts.app.flash&quot;
local time = require &quot;scripts.app.time&quot;
local files = require &quot;scripts.app.files&quot;
local swim = require &quot;scripts.app.swim&quot;
local buffers = require &quot;scripts.app.buffers&quot;
&#10;-- local functions
local function create_header( file, prgKB, chrKB )
    -- Mapper : 200 ; Mirroring : 0 (Horizontal)
&#10;    nes.write_header( file, prgKB, chrKB, 200, 0)
end
&#10;--dump the PRG ROM
local function dump_prgrom( file, rom_size_KB, debug )
&#10;    --PRG-ROM dump 16KB at a time
    local KB_per_read = 16
    local num_reads = rom_size_KB / KB_per_read
    local read_count = 0
    local addr_base = 0x80  -- $8000
&#10;print(&quot;CPU dump :&quot;)
&#10;    while ( read_count &lt; num_reads )
    do
        if debug then print( &quot;dump PRG part &quot;, read_count, &quot; of &quot;, num_reads) end
&#10;        -- pause of 1 second to take care about the chips
        local time=os.clock()+1
        while time&gt;os.clock()
            do
            end
&#10;        local nobusconflicts = dict.nes(&quot;NES_CPU_RD&quot;, 0x8000 | read_count);
        dict.nes(&quot;NES_CPU_WR&quot;, 0x8000 | read_count , nobusconflicts)
&#10;        dump.dumptofile( file, KB_per_read, addr_base, &quot;NESCPU_PAGE&quot;, false )
&#10;        read_count = read_count + 1
    end
end
&#10;--dump the CHR ROM
local function dump_chrrom( file, rom_size_KB, debug )
&#10;    local KB_per_read = 8   --dump both PT at once
    local num_reads = rom_size_KB / KB_per_read
    local read_count = 0
    local addr_base = 0x00  -- $0000
&#10;
print(&quot;PPU dump :&quot;)
&#10;    while ( read_count &lt; num_reads ) do
&#10;        if debug then print( &quot;dump CHR part &quot;, read_count, &quot; of &quot;, num_reads) end
&#10;        -- pause of 1 second to take care about the chips
        local time=os.clock()+1
        while time&gt;os.clock()
            do
            end
&#10;        local nobusconflicts = dict.nes(&quot;NES_CPU_RD&quot;, 0x8000 | read_count);
        dict.nes(&quot;NES_CPU_WR&quot;, 0x8000 | read_count , nobusconflicts)
&#10;        dump.dumptofile( file, KB_per_read, addr_base, &quot;NESPPU_PAGE&quot;, false )
&#10;        read_count = read_count + 1
    end
end
&#10;--Cart should be in reset state upon calling this function
--this function processes all user requests for this specific board/mapper
--local function process( test, read, erase, program, verify, dumpfile, flashfile, verifyfile)
local function process(process_opts, console_opts)
    local test = process_opts[&quot;test&quot;]
    local read = process_opts[&quot;read&quot;]
    local erase = process_opts[&quot;erase&quot;]
    local program = process_opts[&quot;program&quot;]
    local verify = process_opts[&quot;verify&quot;]
    local dumpfile = process_opts[&quot;dump_filename&quot;]
    local flashfile = process_opts[&quot;flash_filename&quot;]
    local verifyfile = process_opts[&quot;verify_filename&quot;]
&#10;    local rv = nil
    local file
    local prg_size = console_opts[&quot;prg_rom_size_kb&quot;]
    local chr_size = console_opts[&quot;chr_rom_size_kb&quot;]
    local wram_size = console_opts[&quot;wram_size_kb&quot;]
&#10;--initialize device i/o for NES
    dict.io(&quot;IO_RESET&quot;)
    dict.io(&quot;NES_INIT&quot;)
&#10;print(&quot;\nRunning dumpMG109.lua&quot;)
&#10;--dump the cart to dumpfile
    if read then
        print(&quot;\nDumping ROM...&quot;)
&#10;        file = assert(io.open(dumpfile, &quot;wb&quot;))
&#10;        --create header: pass open &amp; empty file &amp; rom sizes
        create_header(file, prg_size, chr_size)
&#10;        --TODO find bank table to avoid bus conflicts!
        --dump cart into file
        time.start()
        --dump cart into file
        dump_prgrom(file, prg_size, false)
        dump_chrrom(file, chr_size, false)
        time.report(prg_size)
&#10;        --close file
        assert(file:close())
        print(&quot;DONE Dumping ROM&quot;)
    end
&#10;    dict.io(&quot;IO_RESET&quot;)
end
&#10;-- functions other modules are able to call
dumpmg109.process = process
&#10;-- return the module&#39;s table
return dumpmg109</code></pre></td>
</tr>
</tbody>
</table>

Under UNIX, the commands to run are :

    ./inlretro -s scripts/inlretro3.lua -c NES -m dumpmg109 -x 64 -y 32 -d 50in1-1993.nes

The NES header of the rom is to set to

## Games Menu {#Games_Menu}

There are 4 unique games and the same games at different levels :

- Duck Hunt
- Wild Gunman
- Battle City
- Mario Bros (Arcade / 1983)

<table class="wikitable" data-align="center" style="margin: auto;">
<thead>
<tr>
<th colspan="4">Cartridge Menu</th>
</tr>
</thead>
<tbody>
<tr>
<th>N°</th>
<th>Game Name</th>
<th>Original Game Name</th>
<th>Level</th>
</tr>
&#10;<tr>
<td>1</td>
<td>DUCK HUNT</td>
<td>Duck Hunt</td>
<td>Game A : 1 duck</td>
</tr>
<tr>
<td>2</td>
<td>TWO DUCKS HUNT</td>
<td>Duck Hunt</td>
<td>Game B : 2 ducks</td>
</tr>
<tr>
<td>3</td>
<td>CLAY SHOOTING</td>
<td>Duck Hunt</td>
<td>Game C : Clay Shooting</td>
</tr>
<tr>
<td>4</td>
<td>WILD GUNMAN</td>
<td>Wild Gunman</td>
<td>Game A : 1 outlaw</td>
</tr>
<tr>
<td>5</td>
<td>TWO GUNMEN</td>
<td>Wild Gunman</td>
<td>Game B : 2 outlaws</td>
</tr>
<tr>
<td>6</td>
<td>SALOON GUNMEN</td>
<td>Wild Gunman</td>
<td>Game C : Gang</td>
</tr>
<tr>
<td>7</td>
<td>BATTLE CITY 1</td>
<td>Battle City</td>
<td>Stage 1</td>
</tr>
<tr>
<td>8</td>
<td>BATTLE CITY 2</td>
<td>Battle City</td>
<td>Stage 2</td>
</tr>
<tr>
<td>9</td>
<td>BATTLE CITY 3</td>
<td>Battle City</td>
<td>Stage 3</td>
</tr>
<tr>
<td>10</td>
<td>BATTLE CITY 4</td>
<td>Battle City</td>
<td>Stage 4</td>
</tr>
<tr>
<td>11</td>
<td>BATTLE CITY 5</td>
<td>Battle City</td>
<td>Stage 5</td>
</tr>
<tr>
<td>12</td>
<td>BATTLE CITY 6</td>
<td>Battle City</td>
<td>Stage 6</td>
</tr>
<tr>
<td>13</td>
<td>BATTLE CITY 7</td>
<td>Battle City</td>
<td>Stage 7</td>
</tr>
<tr>
<td>14</td>
<td>BATTLE CITY 8</td>
<td>Battle City</td>
<td>Stage 8</td>
</tr>
<tr>
<td>15</td>
<td>BATTLE CITY 9</td>
<td>Battle City</td>
<td>Stage 9</td>
</tr>
<tr>
<td>16</td>
<td>BATTLE CITY 10</td>
<td>Battle City</td>
<td>Stage 10</td>
</tr>
<tr>
<td>17</td>
<td>BATTLE CITY 11</td>
<td>Battle City</td>
<td>Stage 11</td>
</tr>
<tr>
<td>18</td>
<td>BATTLE CITY 12</td>
<td>Battle City</td>
<td>Stage 12</td>
</tr>
<tr>
<td>19</td>
<td>BATTLE CITY 13</td>
<td>Battle City</td>
<td>Stage 13</td>
</tr>
<tr>
<td>20</td>
<td>BATTLE CITY 14</td>
<td>Battle City</td>
<td>Stage 14</td>
</tr>
<tr>
<td>21</td>
<td>BATTLE CITY 15</td>
<td>Battle City</td>
<td>Stage 15</td>
</tr>
<tr>
<td>22</td>
<td>BATTLE CITY 16</td>
<td>Battle City</td>
<td>Stage 16</td>
</tr>
<tr>
<td>23</td>
<td>BATTLE CITY 17</td>
<td>Battle City</td>
<td>Stage 17</td>
</tr>
<tr>
<td>24</td>
<td>BATTLE CITY 18</td>
<td>Battle City</td>
<td>Stage 18</td>
</tr>
<tr>
<td>25</td>
<td>BATTLE CITY 19</td>
<td>Battle City</td>
<td>Stage 19</td>
</tr>
<tr>
<td>26</td>
<td>BATTLE CITY 20</td>
<td>Battle City</td>
<td>Stage 20</td>
</tr>
<tr>
<td>27</td>
<td>BATTLE CITY 21</td>
<td>Battle City</td>
<td>Stage 21</td>
</tr>
<tr>
<td>28</td>
<td>BATTLE CITY 22</td>
<td>Battle City</td>
<td>Stage 22</td>
</tr>
<tr>
<td>29</td>
<td>BATTLE CITY 23</td>
<td>Battle City</td>
<td>Stage 23</td>
</tr>
<tr>
<td>30</td>
<td>BATTLE CITY 24</td>
<td>Battle City</td>
<td>Stage 24</td>
</tr>
<tr>
<td>31</td>
<td>BATTLE CITY 25</td>
<td>Battle City</td>
<td>Stage 25</td>
</tr>
<tr>
<td>32</td>
<td>BATTLE CITY 26</td>
<td>Battle City</td>
<td>Stage 26</td>
</tr>
<tr>
<td>33</td>
<td>BATTLE CITY 27</td>
<td>Battle City</td>
<td>Stage 27</td>
</tr>
<tr>
<td>34</td>
<td>BATTLE CITY 28</td>
<td>Battle City</td>
<td>Stage 28</td>
</tr>
<tr>
<td>35</td>
<td>BATTLE CITY 29</td>
<td>Battle City</td>
<td>Stage 29</td>
</tr>
<tr>
<td>36</td>
<td>BATTLE CITY 30</td>
<td>Battle City</td>
<td>Stage 30</td>
</tr>
<tr>
<td>37</td>
<td>BATTLE CITY 31</td>
<td>Battle City</td>
<td>Stage 31</td>
</tr>
<tr>
<td>38</td>
<td>BATTLE CITY 32</td>
<td>Battle City</td>
<td>Stage 32</td>
</tr>
<tr>
<td>39</td>
<td>BATTLE CITY 33</td>
<td>Battle City</td>
<td>Stage 33</td>
</tr>
<tr>
<td>40</td>
<td>BATTLE CITY 34</td>
<td>Battle City</td>
<td>Stage 34</td>
</tr>
<tr>
<td>41</td>
<td>BATTLE CITY 35</td>
<td>Battle City</td>
<td>Stage 35</td>
</tr>
<tr>
<td>42</td>
<td>MARIO BROS 1</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 1</td>
</tr>
<tr>
<td>43</td>
<td>MARIO BROS 2</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 2</td>
</tr>
<tr>
<td>44</td>
<td>MARIO BROS 3</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 3</td>
</tr>
<tr>
<td>45</td>
<td>MARIO BROS 4</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 4</td>
</tr>
<tr>
<td>46</td>
<td>MARIO BROS 5</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 5</td>
</tr>
<tr>
<td>47</td>
<td>MARIO BROS 6</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 6</td>
</tr>
<tr>
<td>48</td>
<td>MARIO BROS 7</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 7</td>
</tr>
<tr>
<td>49</td>
<td>MARIO BROS 8</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 8</td>
</tr>
<tr>
<td>50</td>
<td>MARIO BROS 9</td>
<td>Mario Bros (Arcade / 1983)</td>
<td>Phase 9</td>
</tr>
</tbody>
</table>
