---
title: "Consistent frame synchronization"
url: "/wiki/Consistent_frame_synchronization/"
description: "This page describes a method for consistently synchronizing with the PPU every frame from inside an NMI handler, without having to cycle-time everything. This method allows synchronization just as…"
wikiPage: "Consistent_frame_synchronization"
---

## Introduction {#Introduction}

This page describes a method for consistently synchronizing with the [PPU](../PPU/) every frame from inside an [NMI](../NMI/) handler, without having to cycle-time everything. This method allows synchronization just as good as is possible with completely cycle-timed code. At the beginning, the PPU is precisely synchronized with, ensuring that the code behaves the same every time it's run, every time the NES is powered up or reset. It's fully predictable.

Currently only [PAL](../PAL_video/) version is covered, since the PAL PPU's frame timing is simpler. The NTSC version operates in a similar manner, and will be covered eventually.

## PAL timing {#PAL_timing}

*See also: [Cycle reference chart](../Cycle_reference_chart/)*

The PAL NES has a master clock shared by the PPU and CPU. The CPU divides the master clock by 16 to get its instruction cycle clock, which we'll call *cycle* for simplicity. For example, a NOP instruction takes 2 cycles. The PPU divides the master clock by 5 to get its pixel clock, which we'll call *pixel* for simplicity. There are 16/5 = 3.2 pixels per cycle.

A video frame consists of 312 scanlines, each 341 pixels long. Unlike NTSC, there are no short frames, and rendering being enabled or disabled has no effect on frame length. Thus, *every* frame is exactly 312\*341 = 106396 pixels = 33247.5 cycles long. We'll have pixel 0 refer to the first pixel of a frame, and pixel 106395 refer to the last pixel of a frame.

A frame begins with the [vertical blanking interval (VBL)](../The_frame_and_NMIs/#VBlank.2C_Rendering_Time.2C_and_NMIs), then the visible scanlines. The notation VBL+N refers to N cycles after the cycle that VBL began within, VBL+0. To talk about pixels since VBL, we simply refer to pixel P, where pixel 0 is the beginning of VBL, and pixel 106395 is the last pixel in the frame.

## Basic synchronization {#Basic_synchronization}

If we're going to write at a particular pixel, we must first synchronize the CPU to the beginning of a frame, so that pixel 0 begins at the beginning of a cycle, and we know how many cycles ago that was. Reading [\$2002](../PPU_registers/#Status_.28.242002.29_.3C_read) gives the current status of the VBL flag in bit 7, then clears it. The VBL flag is set at pixel 0 of each frame, and cleared around when VBL ends. We can use the VBL flag to achieve synchronization.

A frame is 33247.5 cycles long. If we could somehow read \$2002 every 33247.5 cycles, we'd read at the same point in each frame. But if we read \$2002 every 33248 cycles, we'll be reading 0.5 cycles (1.6 pixels) later each successive frame. If we have a loop do this until it finds the VBL flag set, it will synchronize with the PPU. Each time through, it will read later in the frame, until it reads just as the VBL flag for the next frame is set.

            ; Fine synchronize
    :       delay 33241
            bit $2002
            bpl :-

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th>PPU</th>
<th>CPU</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td></td>
<td></td>
</tr>
<tr>
<td>1</td>
<td></td>
<td></td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>33246</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>33246.5</td>
<td></td>
<td></td>
</tr>
<tr>
<td>33247</td>
<td></td>
<td></td>
</tr>
<tr>
<td>33247.5</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>66494</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>66494.5</td>
<td></td>
<td></td>
</tr>
<tr>
<td>66495</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>99742</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>99742.5</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>132990</td>
<td>Set VBL flag</td>
<td>Read $2002 = $80</td>
</tr>
</tbody>
</table>

Looking at it relative to each frame, we more clearly see how the CPU effectively reads later by half a cycle each frame.

| Cycle | Frame 1 | Frame 2 | Frame 3 | Frame 4 | Event        |
|-------|---------|---------|---------|---------|--------------|
| -1.5  | read    |         |         |         |              |
| -1.0  |         | read    |         |         |              |
| -0.5  |         |         | read    |         |              |
| 0     |         |         |         | read    | VBL flag set |

The loop must be started so that the first \$2002 read is slightly before the end of the frame, otherwise it might start out reading well after the flag has been set. We can do this by starting with a simpler coarse synchronization loop.

    sync_ppu:
            ; Coarse synchronize
            bit $2002
    :       bit $2002
            bpl :-

            delay 33231
            jmp first

            ; Fine synchronize
    :       delay 33241
    first:  bit $2002
            bpl :-

            rts

The coarse synchronization loop might read \$2002 just as the VBL flag was set, or read it nearly 7 cycles after it was set. Then, in the fine synchronization loop, \$2002 is read 33240 to 33247 cycles later. In most cases, this will be slightly before the VBL flag is set, so the loop will delay and read \$2002 again 33248 cycles later, etc.

Once done, the CPU will have executed two cycles after the final \$2002 read that found the VBL flag just set.

## Writing to a particular pixel {#Writing_to_a_particular_pixel}

In order to achieve some graphical effect, we want to write to the PPU at a particular pixel every frame. As an example, we'll write to \$2006 at pixel 30400, which is near the upper-center of the screen. To simplify things, we'll not care what value we write. This requires that we write to \$2006 at VBL+9500.

            ; Synchronize to PPU
            jsr sync_ppu

            ; Delay almost a full frame, so that the code below begins on
            ; a frame.
            delay 33238

    vbl:    ; VBL begins in this cycle

            delay 9497
            sta $2006

<table class="tabular">
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th>Pixel</th>
<th>Cycle</th>
<th>Event</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td>0</td>
<td>VBL begins<br />
delay 9497</td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td></td>
<td>9497</td>
<td>STA $2006</td>
</tr>
<tr>
<td></td>
<td>9498</td>
<td></td>
</tr>
<tr>
<td></td>
<td>9499</td>
<td></td>
</tr>
<tr>
<td>30400</td>
<td>9500</td>
<td>$2006 write</td>
</tr>
</tbody>
</table>

If we try to make this write to the same pixel each frame, we run into a problem: the frame length isn't a whole number of cycles. We'll count frames and treat odd frames as being 33247 cycles long, and even frames 33248 cycles long, which will average to the correct 33247.5 cycles per frame.

            ; Synchronize to PPU
            jsr sync_ppu

            ; Delay almost a full frame, so that the code below begins on
            ; a frame.
            delay 33233

            ; We were on frame 1 after sync_ppu, but vbl will begin on frame 2
            lda #2
            sta frame_count

    vbl:    ; VBL begins in this cycle

            delay 9497
            sta $2006

            delay 23731

            ; Delay extra cycle on even frames
            lda frame_count
            and #$01
            beq extra
    extra:  inc frame_count

            jmp vbl

Now our write time doesn't drift, but it still doesn't write to the same pixel each frame. Since even frames begin in the middle of a cycle, our write is half a cycle/1.6 pixels earlier.

<table class="tabular">
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th>Odd frame pixel</th>
<th>Even frame pixel</th>
<th>Cycle</th>
<th>Event</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td></td>
<td>0</td>
<td rowspan="2">VBL begins<br />
delay 9497</td>
</tr>
<tr>
<td></td>
<td>0</td>
<td>0.5</td>
</tr>
<tr>
<td colspan="4" style="text-align: center;">...</td>
</tr>
<tr>
<td></td>
<td></td>
<td>9497</td>
<td>STA $2006</td>
</tr>
<tr>
<td></td>
<td></td>
<td>9498</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>9499</td>
<td></td>
</tr>
<tr>
<td>30400</td>
<td>30398.4</td>
<td>9500</td>
<td>$2006 write</td>
</tr>
</tbody>
</table>

Our write will thus fall on pixel 30400 on odd frames, and pixel 30398.4 on even frames. That's the best we can do, regardless of how we write our code, as this is a hardware limitation.

Another similar limitation is that when the NES is powered up or reset, the CPU and PPU master clock dividers start in random states, adding up to 1.6 additional pixels of variance. This offset doesn't change until the NES is powered off or reset.

## Ideal NMI {#Ideal_NMI}

Above, all the code had to be cycle-timed to ensure that each write occurred at the correct time. This isn't practical in most programs, which instead use NMI for synchronizing roughly to VBL. In these programs, timing-critical code is at the beginning of the NMI handler, followed by code that isn't carefully timed. Thus, such code relies on NMI occurring shortly after VBL, and not being delayed.

Ideally, NMI would begin a fixed number of cycles after VBL, without waiting for the current instruction to finish. If that were the case, we'd have it nearly as easy as before. Here, we'll imagine NMI always occurs at VBL+2. NMI takes 7 cycles to vector to our NMI handler, so that our NMI handler begins at VBL+9. To simplify the code and timing diagrams, we won't bother saving any registers as we'd normally do in an NMI handler.

    nmi:    ; VBL+9
            delay 9488
            sta $2006         ; write at VBL+9500

<table class="tabular">
<thead>
<tr>
<th>Even frame pixel</th>
<th>Odd frame pixel</th>
<th>Cycle</th>
<th>Event</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td></td>
<td>0</td>
<td rowspan="2">VBL begins</td>
</tr>
<tr>
<td></td>
<td>0</td>
<td>0.5</td>
</tr>
<tr>
<td></td>
<td></td>
<td>1</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>2</td>
<td>NMI vectored</td>
</tr>
<tr>
<td></td>
<td></td>
<td>3</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>4</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>5</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>6</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>7</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>8</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>9</td>
<td>delay 9488</td>
</tr>
<tr>
<td colspan="4" style="text-align: center;">...</td>
</tr>
<tr>
<td></td>
<td></td>
<td>9497</td>
<td>STA $2006</td>
</tr>
<tr>
<td></td>
<td></td>
<td>9498</td>
<td></td>
</tr>
<tr>
<td></td>
<td></td>
<td>9499</td>
<td></td>
</tr>
<tr>
<td>30400</td>
<td>30398.4</td>
<td>9500</td>
<td>$2006 write</td>
</tr>
</tbody>
</table>

## NMI delay {#NMI_delay}

In reality, NMI waits until the current instruction completes before vectoring to the NMI handler, adding an extra delay as compared to the ideal NMI described above. Also, sometimes the NES powers up with the PPU and CPU dividers such that the NMI occurs an additional cycle later.

By ensuring that a short instruction is executing when VBL occurs, we can minimize the delay before NMI is vectored. For example, if we have a series of NOP instructions executing when VBL occurs, NMI will occur from 2 to 4 cycles after VBL. The table shows the four possible timings, with each column titled with the time NMI vectoring begins.

            nop
            nop
            nop

| Cycle | VBL + 2      | VBL + 3      | VBL + 4      | Event      |
|-------|--------------|--------------|--------------|------------|
| -1    |              | NOP          |              |            |
| 0     | NOP          |              | NOP          | VBL begins |
| 1     |              | NOP          |              |            |
| 2     | NMI vectored |              | NOP          |            |
| 3     |              | NMI vectored |              |            |
| 4     |              |              | NMI vectored |            |

So, at best, we have 2 to 4 cycles of delay between VBL and our NMI handler.

Using a long sequence of NOP instructions isn't practical, because it requires either a large number of NOP instructions, or that we know how long the code before them takes so that we can delay entry into the NOP sequence until NMI is about to occur. If we instead have a simple infinite loop made of a single JMP instruction, we only increase the maximum delay by one cycle, to 5.

     loop:   jmp loop

| Cycle | VBL + 2      | VBL + 3      | VBL + 4      | VBL + 5      | Event      |
|-------|--------------|--------------|--------------|--------------|------------|
| -1    | JMP          |              |              | JMP          |            |
| 0     |              | JMP          |              |              | VBL begins |
| 1     |              |              | JMP          |              |            |
| 2     | NMI vectored |              |              | JMP          |            |
| 3     |              | NMI vectored |              |              |            |
| 4     |              |              | NMI vectored |              |            |
| 5     |              |              |              | NMI vectored |            |

## Compensating for NMI delay {#Compensating_for_NMI_delay}

With a JMP loop to wait for NMI, we have 2 to 5 cycles of delay between VBL and our NMI handler. We want to compensate for this delay D by delaying an additional 5-D cycles. Here, we have the NOP always begin at VBL+12. We can't actually do this, but it shows what we must do the equivalent of.

    nmi:   delay 5-D
        nop

| Cycle | VBL + 2      | VBL + 3      | VBL + 4      | VBL + 5        | Event      |
|-------|--------------|--------------|--------------|----------------|------------|
| -1    | JMP          |              |              | JMP            |            |
| 0     |              | JMP          |              |                | VBL begins |
| 1     |              |              | JMP          |                |            |
| 2     | NMI vectored |              |              | JMP            |            |
| 3     |              | NMI vectored |              |                |            |
| 4     |              |              | NMI vectored |                |            |
| 5     |              |              |              | NMI vectored   |            |
| 6     |              |              |              |                |            |
| 7     |              |              |              |                |            |
| 8     |              |              |              |                |            |
| 9     | delay 3      |              |              |                |            |
| 10    |              | delay 2      |              |                |            |
| 11    |              |              | delay 1      |                |            |
| 12    | NOP          | NOP          | NOP          | NOP (no delay) |            |

We just have to find out *how* to determine the number of cycles of delay to add.

## Sprite DMA always ends on even cycle {#Sprite_DMA_always_ends_on_even_cycle}

When sprite DMA (\$4014) is written to, the next instruction always begins on an odd cycle. If the \$4014 write is on an odd cycle, it pauses the CPU for an additional 513 cycles, otherwise 514 cycles. We can use this aspect to partially compensate for NMI's variable delay.

    nmi:    lda #$07          ; sprites at $700
            sta $4014
            nop

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th style="text-align: center;">VBL + 2</th>
<th style="text-align: center;">VBL + 3</th>
<th style="text-align: center;">VBL + 4</th>
<th style="text-align: center;">VBL + 5</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td colspan="4" style="text-align: center;">VBL begins</td>
</tr>
<tr>
<td>1</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>2</td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>3</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>4</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>5</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
</tr>
<tr>
<td>6</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>8</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>9</td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>10</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>11</td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>12</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
</tr>
<tr>
<td>13</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>14</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
</tr>
<tr>
<td>15</td>
<td style="text-align: center;">514-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>16</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">513-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>17</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">514-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
</tr>
<tr>
<td>18</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">513-cycle DMA</td>
</tr>
<tr>
<td colspan="5" style="text-align: center;">...</td>
</tr>
<tr>
<td>527</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>528</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>529</td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>530</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;">DMA finishes</td>
</tr>
<tr>
<td>531</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;">NOP</td>
</tr>
</tbody>
</table>

This reduces the number of different delays from four to two. The NOP always executes at either VBL+529 or VBL+531. This is an improvement. We just need a way to determine which time DMA finished at, and delay two extra cycles if it was the earlier one.

## VBL flag cleared at end of VBL {#VBL_flag_cleared_at_end_of_VBL}

The VBL flag is cleared near the end of VBL. If we read \$2002 around the time the flag is cleared, we can determine whether the read occurred before or after the flag was cleared. We will have to avoid reading \$2002 elsewhere in the NMI handler, since reading \$2002 clears the flag.

The VBL flag is cleared around pixel 23869, sometimes one less, so we want to read \$2002 at VBL+7458 or VBL+7460. It works out nicely that sprite DMA leaves two cycles between the possible ending times, as this ensures that our \$2002 read is several pixels before or after when the flag is cleared, giving us a good margin for error. If we find the flag set, we know we are on the earlier of the two DMA ending times, so we delay an extra two cycles.

    nmi:    lda #$07          ; sprites at $700
            sta $4014
            delay 6926
            bit $2002         ; read at VBL+7458 or VBL+7460
            bpl skip
            bit 0
    skip:   nop

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th style="text-align: center;">VBL + 2</th>
<th style="text-align: center;">VBL + 3</th>
<th style="text-align: center;">VBL + 4</th>
<th style="text-align: center;">VBL + 5</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td colspan="4" style="text-align: center;">VBL begins</td>
</tr>
<tr>
<td>1</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>2</td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>3</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>4</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>5</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
</tr>
<tr>
<td>6</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>8</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>9</td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>10</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>11</td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>12</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
</tr>
<tr>
<td>13</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>14</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
</tr>
<tr>
<td>15</td>
<td style="text-align: center;">514-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>16</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">513-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>17</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">514-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
</tr>
<tr>
<td>18</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">513-cycle DMA</td>
</tr>
<tr>
<td colspan="5" style="text-align: center;">...</td>
</tr>
<tr>
<td>527</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>528</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>529</td>
<td style="text-align: center;">delay 6926</td>
<td style="text-align: center;">delay 6926</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>530</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;">DMA finishes</td>
</tr>
<tr>
<td>531</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">delay 6926</td>
<td style="text-align: center;">delay 6926</td>
</tr>
<tr>
<td colspan="5" style="text-align: center;">...</td>
</tr>
<tr>
<td>7455</td>
<td style="text-align: center;">BIT $2002</td>
<td style="text-align: center;">BIT $2002</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7456</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7457</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">BIT $2002</td>
<td style="text-align: center;">BIT $2002</td>
</tr>
<tr>
<td>7458</td>
<td style="text-align: center;">$2002 read = $80</td>
<td style="text-align: center;">$2002 read = $80</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7459</td>
<td style="text-align: center;">BPL not taken</td>
<td style="text-align: center;">BPL not taken</td>
<td style="text-align: center;">VBL cleared</td>
<td style="text-align: center;">VBL cleared</td>
</tr>
<tr>
<td>7460</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">$2002 read = 0</td>
<td style="text-align: center;">$2002 read = 0</td>
</tr>
<tr>
<td>7461</td>
<td style="text-align: center;">BIT 0</td>
<td style="text-align: center;">BIT 0</td>
<td style="text-align: center;">BPL taken</td>
<td style="text-align: center;">BPL taken</td>
</tr>
<tr>
<td>7462</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7463</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>7464</td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;">NOP</td>
<td style="text-align: center;">NOP</td>
</tr>
</tbody>
</table>

This achieves our goal, but not in all cases.

## VBL begins on odd cycles {#VBL_begins_on_odd_cycles}

Unfortunately, VBL doesn't always begin during an even cycle, as we've so far assumed. When VBL begins during an odd cycle, our code doesn't work so well:

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th style="text-align: center;">VBL + 2</th>
<th style="text-align: center;">VBL + 3</th>
<th style="text-align: center;">VBL + 4</th>
<th style="text-align: center;">VBL + 5</th>
</tr>
</thead>
<tbody>
<tr>
<td>1</td>
<td colspan="4" style="text-align: center;">VBL begins</td>
</tr>
<tr>
<td>2</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>3</td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>4</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>5</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>6</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">NMI</td>
</tr>
<tr>
<td>7</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>8</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>9</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>10</td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>11</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>12</td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>13</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">LDA #$07</td>
</tr>
<tr>
<td>14</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>15</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">STA $4014</td>
</tr>
<tr>
<td>16</td>
<td style="text-align: center;">513-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>17</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">514-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>18</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">513-cycle DMA</td>
<td style="text-align: center;">$4014 write</td>
</tr>
<tr>
<td>19</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">514-cycle DMA</td>
</tr>
<tr>
<td colspan="5" style="text-align: center;">...</td>
</tr>
<tr>
<td>527</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>528</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>529</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>530</td>
<td style="text-align: center;"></td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;">DMA finishes</td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>531</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
</tr>
<tr>
<td>532</td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;"></td>
<td style="text-align: center;">DMA finishes</td>
</tr>
</tbody>
</table>

Now DMA ends at *three* different times, covering a wider range than the original NMI times did, thus making things worse!

We need to keep track of when VBL begins during an odd cycle, and compensate *before* we begin DMA. After our PPU synchronization routine finishes, the last \$2002 read it makes will have just found the VBL flag set. In the following table, that is cycle 0.

| Pixel  | Cycle    | Frame |
|--------|----------|-------|
| 0      | 0        | 1     |
| 106392 | 33247.5  | 2     |
| 212784 | 66495    | 3     |
| 319176 | 99742.5  | 4     |
| 425568 | 132990   | 5     |
| 531960 | 166237.5 | 6     |
| 638352 | 199485   | 7     |
| 744744 | 232732.5 | 8     |

Looking at which cycle each frame begins on, we see they follow a four-frame pattern: even, odd, odd, even. So we'll just have a variable that starts out at 1 and increments every frame, then examine bit 1 and delay an extra cycle if it's clear. This extra code takes 8 cycles on frames where VBL begins during an even cycle, and 7 cycles otherwise.

But we also need to insert a complementary delay *after* DMA, before the \$2002 read, since on frames where VBL begins during an odd cycle we'll need to read \$2002 one cycle later after DMA than for even frames.

    nmi:    lda frame_count
            and #$02
            beq even
    even:   lda #$07          ; sprites at $700
            sta $4014
            delay 6911
            lda frame_count
            and #$02
            bne odd
    odd:    bit $2002
            bpl skip
            bit 0
    skip:   inc frame_count
            delay 2028
            sta $2006

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th colspan="4">Frames 1, 4, 5, 8 ...</th>
<th colspan="4">Frames 2, 3, 6, 7 ...</th>
</tr>
</thead>
<tbody>
<tr>
<th>VBL + 2</th>
<th>VBL + 3</th>
<th>VBL + 4</th>
<th>VBL + 5</th>
<th>VBL + 2</th>
<th>VBL + 3</th>
<th>VBL + 4</th>
<th>VBL + 5</th>
<th></th>
</tr>
&#10;<tr>
<td>0</td>
<td>VBL</td>
<td>VBL</td>
<td>VBL</td>
<td>VBL</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>1</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>VBL</td>
<td>VBL</td>
<td>VBL</td>
<td>VBL</td>
</tr>
<tr>
<td>2</td>
<td>NMI</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>3</td>
<td></td>
<td>NMI</td>
<td></td>
<td></td>
<td>NMI</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>4</td>
<td></td>
<td></td>
<td>NMI</td>
<td></td>
<td></td>
<td>NMI</td>
<td></td>
<td></td>
</tr>
<tr>
<td>5</td>
<td></td>
<td></td>
<td></td>
<td>NMI</td>
<td></td>
<td></td>
<td>NMI</td>
<td></td>
</tr>
<tr>
<td>6</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>NMI</td>
</tr>
<tr>
<td>7</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>8</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>9</td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>10</td>
<td></td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>11</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td></td>
<td></td>
</tr>
<tr>
<td>12</td>
<td>AND #$02</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td></td>
</tr>
<tr>
<td>13</td>
<td></td>
<td>AND #$02</td>
<td></td>
<td></td>
<td>AND #$02</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
</tr>
<tr>
<td>14</td>
<td>BEQ taken</td>
<td></td>
<td>AND #$02</td>
<td></td>
<td></td>
<td>AND #$02</td>
<td></td>
<td></td>
</tr>
<tr>
<td>15</td>
<td></td>
<td>BEQ taken</td>
<td></td>
<td>AND #$02</td>
<td>BEQ not taken</td>
<td></td>
<td>AND #$02</td>
<td></td>
</tr>
<tr>
<td>16</td>
<td></td>
<td></td>
<td>BEQ taken</td>
<td></td>
<td></td>
<td>BEQ not taken</td>
<td></td>
<td>AND #$02</td>
</tr>
<tr>
<td>17</td>
<td>LDA #$07</td>
<td></td>
<td></td>
<td>BEQ taken</td>
<td>LDA #$07</td>
<td></td>
<td>BEQ not taken</td>
<td></td>
</tr>
<tr>
<td>18</td>
<td></td>
<td>LDA #$07</td>
<td></td>
<td></td>
<td></td>
<td>LDA #$07</td>
<td></td>
<td>BEQ not taken</td>
</tr>
<tr>
<td>19</td>
<td>STA $4014</td>
<td></td>
<td>LDA #$07</td>
<td></td>
<td>STA $4014</td>
<td></td>
<td>LDA #$07</td>
<td></td>
</tr>
<tr>
<td>20</td>
<td></td>
<td>STA $4014</td>
<td></td>
<td>LDA #$07</td>
<td></td>
<td>STA $4014</td>
<td></td>
<td>LDA #$07</td>
</tr>
<tr>
<td>21</td>
<td></td>
<td></td>
<td>STA $4014</td>
<td></td>
<td></td>
<td></td>
<td>STA $4014</td>
<td></td>
</tr>
<tr>
<td>22</td>
<td>$4014 write</td>
<td></td>
<td></td>
<td>STA $4014</td>
<td>$4014 write</td>
<td></td>
<td></td>
<td>STA $4014</td>
</tr>
<tr>
<td>23</td>
<td>514-cycle DMA</td>
<td>$4014 write</td>
<td></td>
<td></td>
<td>514-cycle DMA</td>
<td>$4014 write</td>
<td></td>
<td></td>
</tr>
<tr>
<td>24</td>
<td></td>
<td>513-cycle DMA</td>
<td>$4014 write</td>
<td></td>
<td></td>
<td>513-cycle DMA</td>
<td>$4014 write</td>
<td></td>
</tr>
<tr>
<td>25</td>
<td></td>
<td></td>
<td>514-cycle DMA</td>
<td>$4014 write</td>
<td></td>
<td></td>
<td>514-cycle DMA</td>
<td>$4014 write</td>
</tr>
<tr>
<td>26</td>
<td></td>
<td></td>
<td></td>
<td>513-cycle DMA</td>
<td></td>
<td></td>
<td></td>
<td>513-cycle DMA</td>
</tr>
<tr>
<td colspan="9" style="text-align: center;">...</td>
</tr>
<tr>
<td>535</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>536</td>
<td>DMA finishes</td>
<td>DMA finishes</td>
<td></td>
<td></td>
<td>DMA finishes</td>
<td>DMA finishes</td>
<td></td>
<td></td>
</tr>
<tr>
<td>537</td>
<td>delay 6911</td>
<td>delay 6911</td>
<td></td>
<td></td>
<td>delay 6911</td>
<td>delay 6911</td>
<td></td>
<td></td>
</tr>
<tr>
<td>538</td>
<td></td>
<td></td>
<td>DMA finishes</td>
<td>DMA finishes</td>
<td></td>
<td></td>
<td>DMA finishes</td>
<td>DMA finishes</td>
</tr>
<tr>
<td>539</td>
<td></td>
<td></td>
<td>delay 6911</td>
<td>delay 6911</td>
<td></td>
<td></td>
<td>delay 6911</td>
<td>delay 6911</td>
</tr>
<tr>
<td colspan="9" style="text-align: center;">...</td>
</tr>
<tr>
<td>7448</td>
<td>LDA frame_count</td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td>LDA frame_count</td>
<td></td>
<td></td>
</tr>
<tr>
<td>7449</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7450</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td>LDA frame_count</td>
<td></td>
<td></td>
<td>LDA frame_count</td>
<td>LDA frame_count</td>
</tr>
<tr>
<td>7451</td>
<td>AND #$02</td>
<td>AND #$02</td>
<td></td>
<td></td>
<td>AND #$02</td>
<td>AND #$02</td>
<td></td>
<td></td>
</tr>
<tr>
<td>7452</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7453</td>
<td>BNE not taken</td>
<td>BNE not taken</td>
<td>AND #$02</td>
<td>AND #$02</td>
<td>BNE taken</td>
<td>BNE taken</td>
<td>AND #$02</td>
<td>AND #$02</td>
</tr>
<tr>
<td>7454</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7455</td>
<td>BIT $2002</td>
<td>BIT $2002</td>
<td>BNE not taken</td>
<td>BNE not taken</td>
<td></td>
<td></td>
<td>BNE taken</td>
<td>BNE taken</td>
</tr>
<tr>
<td>7456</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>BIT $2002</td>
<td>BIT $2002</td>
<td></td>
<td></td>
</tr>
<tr>
<td>7457</td>
<td></td>
<td></td>
<td>BIT $2002</td>
<td>BIT $2002</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7458</td>
<td>$2002 read = $80</td>
<td>$2002 read = $80</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>BIT $2002</td>
<td>BIT $2002</td>
</tr>
<tr>
<td>7459</td>
<td>BPL not taken</td>
<td>BPL not taken</td>
<td>VBL cleared</td>
<td>VBL cleared</td>
<td>$2002 read = $80</td>
<td>$2002 read = $80</td>
<td></td>
<td></td>
</tr>
<tr>
<td>7460</td>
<td></td>
<td></td>
<td>$2002 read = 0</td>
<td>$2002 read = 0</td>
<td>BPL not taken</td>
<td>BPL not taken</td>
<td>VBL cleared</td>
<td>VBL cleared</td>
</tr>
<tr>
<td>7461</td>
<td>BIT 0</td>
<td>BIT 0</td>
<td>BPL taken</td>
<td>BPL taken</td>
<td></td>
<td></td>
<td>$2002 read = 0</td>
<td>$2002 read = 0</td>
</tr>
<tr>
<td>7462</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>BIT 0</td>
<td>BIT 0</td>
<td>BPL taken</td>
<td>BPL taken</td>
</tr>
<tr>
<td>7463</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7464</td>
<td>INC frame_count</td>
<td>INC frame_count</td>
<td>INC frame_count</td>
<td>INC frame_count</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7465</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>INC frame_count</td>
<td>INC frame_count</td>
<td>INC frame_count</td>
<td>INC frame_count</td>
</tr>
<tr>
<td>7466</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7467</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7468</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7469</td>
<td>delay 2028</td>
<td>delay 2028</td>
<td>delay 2028</td>
<td>delay 2028</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>7470</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>delay 2028</td>
<td>delay 2028</td>
<td>delay 2028</td>
<td>delay 2028</td>
</tr>
<tr>
<td colspan="9" style="text-align: center;">...</td>
</tr>
<tr>
<td>9497</td>
<td>STA $2006</td>
<td>STA $2006</td>
<td>STA $2006</td>
<td>STA $2006</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>9498</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>STA $2006</td>
<td>STA $2006</td>
<td>STA $2006</td>
<td>STA $2006</td>
</tr>
<tr>
<td>9499</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>9500</td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr>
<tr>
<td>9501</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
<td>$2006 write at VBL+9500</td>
</tr>
</tbody>
</table>

The \$2006 write is done at VBL+9500 in all cases. Remember that the right four columns have VBL beginning on cycle 1 (an odd cycle), which is why the final writes appear to be one cycle later than the others.

## Synchronizing with even CPU cycle {#Synchronizing_with_even_CPU_cycle}

Since our final synchronization method relies on knowing whether a given frame begins during an even or odd cycle, we must initially ensure that our PPU synchronization routine's final \$2002 read is also during an even cycle. Since the fine synchronization loop takes an even number of cycles, we merely need to ensure that the first time through that the \$2002 read is on an even cycle. We can do this by initiating sprite DMA before the fine synchronization loop.

    sync_ppu:
            ; Coarse synchronize
            bit $2002
    :       bit $2002
            bpl :-

            sta $4014
            delay 32713
            jmp first

            ; Fine synchronize
    :       delay 33241
    first:  bit $2002
            bpl :-

            ; NMI won't be fired until frame 2
            lda #2
            sta frame_count

            rts

The STA \$4014 takes up to 518 cycles, so we subtract that from the initial delay. After the STA \$4014, the delay begins on an odd cycle. Since also it's an odd number of cycles until the \$2002 read, it will occur on an even cycle, as desired.

## Simpler synchronization routine {#Simpler_synchronization_routine}

The PPU synchronization routine is pretty short, but it requires use of the delay macro, which takes a fair amount of code to implement. It's possible to eliminate that without any negative impact.

The fine synchronization loop needs to read \$2002 every 33248 cycles, so it can find when the VBL flag is set just before the read. This seems to require a long delay between reads. Until the final iteration, it must not find the VBL flag set. If it were like the coarse loop and read the VBL flag every 7 cycles, it would clearly stop somewhere near the beginning of the first frame, but rarely right at the beginning. It might read \$2002 one cycle before the VBL flag is set, loop, then read it 7 cycles later and find it now set. This isn't what we want. If we read it slightly more often, like every 33248/2 = 16624 cycles, it would still work, since the VBL flag is automatically cleared near the end of VBL.

    sync_ppu:
            ; Coarse synchronize
            bit $2002
    :       bit $2002
            bpl :-

            sta $4014
            delay 16089
            jmp first

            ; Fine synchronize
    :       delay 16617
    first:  bit $2002
            bpl :-

            rts

| Cycle    | PPU            | CPU                |
|----------|----------------|--------------------|
| 0        | Set VBL flag   |                    |
| 7459     | Clear VBL flag |                    |
| 16622    |                | Read \$2002 = 0    |
| 33246    |                | Read \$2002 = 0    |
| 33247.5  | Set VBL flag   |                    |
| 40706.5  | Clear VBL flag |                    |
| 49870    |                | Read \$2002 = 0    |
| 66494    |                | Read \$2002 = 0    |
| 66495    | Set VBL flag   |                    |
| 73954    | Clear VBL flag |                    |
| 83118    |                | Read \$2002 = 0    |
| 99742    |                | Read \$2002 = 0    |
| 99742.5  | Set VBL flag   |                    |
| 107201.5 | Clear VBL flag |                    |
| 116366   |                | Read \$2002 = 0    |
| 132990   | Set VBL flag   | Read \$2002 = \$80 |

That works, but reducing the delays doesn't eliminate the need for them. The important thing is that the \$2002 read only be able to happen just after the VBL flag is set, rather than many cycles after it was set. Rather than rely on the PPU to clear the VBL flag, we can clear it ourselves. 16 is a factor of 33248, so we can have the loop take only 16 cycles and still synchronize properly.

    sync_ppu:
            ; Coarse synchronize
            bit $2002
    :       bit $2002
            bpl :-

            sta $4014
            bit <0

            ; Fine synchronize
    :       bit <0
            nop
            bit $2002
            bit $2002
            bpl :-

            rts

<table class="tabular">
<thead>
<tr>
<th>Cycle</th>
<th>PPU</th>
<th>CPU</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td>10</td>
<td></td>
<td>Dummy read $2002 = $80</td>
</tr>
<tr>
<td>14</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>26</td>
<td></td>
<td>Dummy read $2002 = 0</td>
</tr>
<tr>
<td>30</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>33242</td>
<td></td>
<td>Dummy read $2002 = 0</td>
</tr>
<tr>
<td>33246</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>33247.5</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td>33258</td>
<td></td>
<td>Dummy read $2002 = $80</td>
</tr>
<tr>
<td>33262</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>66490</td>
<td></td>
<td>Dummy read $2002 = 0</td>
</tr>
<tr>
<td>66494</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>66495</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td>66506</td>
<td></td>
<td>Dummy read $2002 = $80</td>
</tr>
<tr>
<td>66510</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>99738</td>
<td></td>
<td>Dummy read $2002 = 0</td>
</tr>
<tr>
<td>99742</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td>99742.5</td>
<td>Set VBL flag</td>
<td></td>
</tr>
<tr>
<td>99754</td>
<td></td>
<td>Dummy read $2002 = $80</td>
</tr>
<tr>
<td>99758</td>
<td></td>
<td>Read $2002 = 0</td>
</tr>
<tr>
<td colspan="3" style="text-align: center;">...</td>
</tr>
<tr>
<td>132986</td>
<td></td>
<td>Dummy read $2002 = 0</td>
</tr>
<tr>
<td>132990</td>
<td>Set VBL flag</td>
<td>Read $2002 = $80</td>
</tr>
</tbody>
</table>

Essentially there's a four-cycle window that the second \$2002 read in the loop is watching for the VBL flag to be set within. On entry to the loop, we ensure that the flag will never be set within this window. Every 33248/16 = 2078 iterations, the second \$2002 read is half a cycle later in the frame, just like the original version. On every other iteration, the dummy \$2002 read four cycles before has ensured that the VBL flag is cleared.
