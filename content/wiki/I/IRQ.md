---
title: "IRQ"
url: "/wiki/IRQ/"
description: "IRQ (interrupt request) is a signal on the NES CPU. It is used to trigger a CPU interrupt."
wikiPage: "IRQ"
---

**IRQ** (interrupt request) is a signal on the NES CPU. It is used to trigger a [CPU interrupt](../CPU_interrupts/).

If the CPU's /IRQ input is 0 at the end of an instruction, then the CPU pushes the program counter and the processor status register, sets the I flag to ignore further IRQs, and the Program Counter takes the value read at \$fffe and \$ffff.

This behaviour is masked by the CPU's interrupt-disable [status flag](../Status_flags/). The SEI instruction disables IRQs, and the CLI instruction enables them.

/IRQ functions as an open collector input: it is normally 1, but any device on the CPU bus can pull it down to 0. An IRQ handler is expected to push any registers it uses, acknowledge the interrupt by writing to a port so that the source no longer forces /IRQ to 0, pull the registers back, and return with RTI.

Therefore if a program uses more than one source of IRQ, the priority between the conflicting interrupts should be handled in software.

Sources of IRQ on a Famicom or NES include

<table class="tabular">
<thead>
<tr>
<th>Source</th>
<th>Enable</th>
<th>Disable</th>
<th>Acknowledge</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="../APU_DMC/">APU DMC</a> finish</td>
<td>$4010 write with bit 7 = 1</td>
<td>$4010 write otherwise</td>
<td>Disable then reenable, or <a href="../APU/">APU Status</a> ($4015) write</td>
</tr>
<tr>
<td><a href="../APU_Frame_Counter/">APU Frame Counter</a></td>
<td>$4017 write with bits 7-6 = 00</td>
<td>$4017 write otherwise</td>
<td><a href="../APU/">APU Status</a> ($4015) read</td>
</tr>
<tr>
<td><a href="../MMC3/">MMC3</a></td>
<td>Write to $E001</td>
<td>Write to $E000</td>
<td>Disable then reenable</td>
</tr>
<tr>
<td><a href="../MMC5/">MMC5</a></td>
<td>Write $80 to $5204</td>
<td>Write $00 to $5204</td>
<td>Read $5204</td>
</tr>
<tr>
<td><a href="../VRC_IRQ/">VRC4/6/7</a></td>
<td colspan="3">depends on specific IC</td>
</tr>
<tr>
<td><a href="../Sunsoft_FME-7/">FME-7</a></td>
<td>write $81 to register $D</td>
<td>write even number to register $D</td>
<td>write anything to register $D</td>
</tr>
<tr>
<td><a href="../INES_Mapper_019/">Namco 163</a></td>
<td>write to $5000 and $5800</td>
<td>$5800 write with bit 7 = 0</td>
<td>write to $5000 or $5800</td>
</tr>
<tr>
<td><a href="../Family_Computer_Disk_System/">FDS</a></td>
<td>Write $02 to $4022</td>
<td>Write $00 to $4022</td>
<td>Read $4030</td>
</tr>
</tbody>
</table>

## See also {#See_also}

- [CPU interrupts](../CPU_interrupts/)
- [Status flags](../Status_flags/) used to enable/disable IRQs.
- [MMC](../Mapper/), common source of cartridge IRQs
- [NMI](../NMI/), the other interrupt signal on the CPU
