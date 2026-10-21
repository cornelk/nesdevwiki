---
title: "VT02+ Registers"
url: "/wiki/VT02+_Registers/"
description: "When TSYNEN is 0, this register sets \"the number of AD12 switching high low\". When TSYNEN is 1, this register sets \"the number of HSYNC switching high low\" instead. (See $410B.)"
wikiPage: "VT02__Registers"
---

## System Registers {#System_Registers}

### <span id="Program_Bank_1.2C_Video_Bank_2_.28.244100.29_.3E_write"></span>Program Bank 1, Video Bank 2 (\$4100) \> write {#Program_Bank_1.2C_Video_Bank_2_.28.244100.29_.3E_write}

    7  bit  0
    ---- ----
    PPPP VVVV
    |||| ||||
    |||| ++++- Video Bank 2   (VA24, VA23, VA22, VA21)
    ++++------ Program Bank 1 (PA24, PA23, PA22, PA21)

### <span id="Timer_Interrupt_Preload_Times_.28.244101.29_.3E_write"></span>Timer Interrupt Preload Times (\$4101) \> write {#Timer_Interrupt_Preload_Times_.28.244101.29_.3E_write}

When TSYNEN is 0, this register sets "the number of AD12 switching high low". When TSYNEN is 1, this register sets "the number of HSYNC switching high low" instead. (See \$410B.)

### <span id="Load_Timer_Interrupt_Value_.28.244102.29_.3E_write"></span>Load Timer Interrupt Value (\$4102) \> write {#Load_Timer_Interrupt_Value_.28.244102.29_.3E_write}

Writing any value to this register will set the timer's value and start counting.

### <span id="Timer_Interrupt_Disable_.28.244103.29_.3E_write"></span>Timer Interrupt Disable (\$4103) \> write {#Timer_Interrupt_Disable_.28.244103.29_.3E_write}

Writing any value to this register will disable the timer interrupt.

### <span id="Timer_Interrupt_Enable_.28.244104.29_.3E_write"></span>Timer Interrupt Enable (\$4104) \> write {#Timer_Interrupt_Enable_.28.244104.29_.3E_write}

Writing any value to this register will enable the timer interrupt.

### <span id="V_Bank_0_decode_type.2C_P_Bank_0_decode_type.2C_Inter_Char_VRAM_.28.244105.29_.3E_write"></span>V Bank 0 decode type, P Bank 0 decode type, Inter Char VRAM (\$4105) \> write {#V_Bank_0_decode_type.2C_P_Bank_0_decode_type.2C_Inter_Char_VRAM_.28.244105.29_.3E_write}

    7  bit  0
    ---- ----
    VPIx xxxx
    |||| ||||
    |||+-++++- unused
    ||+------- Internal VRAM as CHR RAM (0: disabled; 1: enabled)
    |+-------- Program Bank 0 Decoding (see below table)
    +--------- Video Bank 0 Decoding (see below table)

#### Program Bank 0 Decoding {#Program_Bank_0_Decoding}

<table class="wikitable">
<thead>
<tr>
<th>PQ2EN ($410B)</th>
<th>COMR6</th>
<th>A[14:13] (CPU)</th>
<th>TPA20</th>
<th>TPA19</th>
<th>TPA18</th>
<th>TPA17</th>
<th>TPA16</th>
<th>TPA15</th>
<th>TPA14</th>
<th>TPA13</th>
</tr>
</thead>
<tbody>
<tr>
<td>0</td>
<td colspan="2">0</td>
<td>PQ07</td>
<td>PQ06</td>
<td>PQ05</td>
<td>PQ04</td>
<td>PQ03</td>
<td>PQ02</td>
<td>PQ01</td>
<td>PQ00</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">1</td>
<td>PQ17</td>
<td>PQ16</td>
<td>PQ15</td>
<td>PQ14</td>
<td>PQ13</td>
<td>PQ12</td>
<td>PQ11</td>
<td>PQ10</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">2</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>0</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">3</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">4</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>0</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">5</td>
<td>PQ17</td>
<td>PQ16</td>
<td>PQ15</td>
<td>PQ14</td>
<td>PQ13</td>
<td>PQ12</td>
<td>PQ11</td>
<td>PQ10</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">6</td>
<td>PQ07</td>
<td>PQ06</td>
<td>PQ05</td>
<td>PQ04</td>
<td>PQ03</td>
<td>PQ02</td>
<td>PQ01</td>
<td>PQ00</td>
</tr>
<tr>
<td>0</td>
<td colspan="2">7</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">0</td>
<td>PQ07</td>
<td>PQ06</td>
<td>PQ05</td>
<td>PQ04</td>
<td>PQ03</td>
<td>PQ02</td>
<td>PQ01</td>
<td>PQ00</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">1</td>
<td>PQ17</td>
<td>PQ16</td>
<td>PQ15</td>
<td>PQ14</td>
<td>PQ13</td>
<td>PQ12</td>
<td>PQ11</td>
<td>PQ10</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">2</td>
<td>PQ27</td>
<td>PQ26</td>
<td>PQ25</td>
<td>PQ24</td>
<td>PQ23</td>
<td>PQ22</td>
<td>PQ21</td>
<td>PQ20</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">3</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">4</td>
<td>PQ27</td>
<td>PQ26</td>
<td>PQ25</td>
<td>PQ24</td>
<td>PQ23</td>
<td>PQ22</td>
<td>PQ21</td>
<td>PQ20</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">5</td>
<td>PQ17</td>
<td>PQ16</td>
<td>PQ15</td>
<td>PQ14</td>
<td>PQ13</td>
<td>PQ12</td>
<td>PQ11</td>
<td>PQ10</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">6</td>
<td>PQ07</td>
<td>PQ06</td>
<td>PQ05</td>
<td>PQ04</td>
<td>PQ03</td>
<td>PQ02</td>
<td>PQ01</td>
<td>PQ00</td>
</tr>
<tr>
<td>1</td>
<td colspan="2">7</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
<td>1</td>
</tr>
</tbody>
</table>

#### Video Bank 0 Decoding {#Video_Bank_0_Decoding}

<table class="wikitable">
<thead>
<tr>
<th>COMR7</th>
<th>AD[12:10]</th>
<th>TVA17</th>
<th>TVA16</th>
<th>TVA15</th>
<th>TVA14</th>
<th>TVA13</th>
<th>TVA12</th>
<th>TVA11</th>
<th>TVA10</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2">0,1,C,D</td>
<td>RV47</td>
<td>RV46</td>
<td>RV45</td>
<td>RV44</td>
<td>RV43</td>
<td>RV42</td>
<td>RV41</td>
<td>AD10</td>
</tr>
<tr>
<td colspan="2">2,3,E,F</td>
<td>RV57</td>
<td>RV56</td>
<td>RV55</td>
<td>RV54</td>
<td>RV53</td>
<td>RV52</td>
<td>RV51</td>
<td>AD10</td>
</tr>
<tr>
<td colspan="2">4,8</td>
<td>RV07</td>
<td>RV06</td>
<td>RV05</td>
<td>RV04</td>
<td>RV03</td>
<td>RV02</td>
<td>RV01</td>
<td>RV00</td>
</tr>
<tr>
<td colspan="2">5,9</td>
<td>RV17</td>
<td>RV16</td>
<td>RV15</td>
<td>RV14</td>
<td>RV13</td>
<td>RV12</td>
<td>RV11</td>
<td>RV10</td>
</tr>
<tr>
<td colspan="2">6,A</td>
<td>RV27</td>
<td>RV26</td>
<td>RV25</td>
<td>RV24</td>
<td>RV23</td>
<td>RV22</td>
<td>RV21</td>
<td>RV20</td>
</tr>
<tr>
<td colspan="2">7,B</td>
<td>RV37</td>
<td>RV36</td>
<td>RV35</td>
<td>RV34</td>
<td>RV33</td>
<td>RV32</td>
<td>RV31</td>
<td>RV30</td>
</tr>
</tbody>
</table>

### <span id="Horizontal.2FVertical_Scrolling_Selector_.28.244106.29_.3E_write"></span>Horizontal/Vertical Scrolling Selector (\$4106) \> write {#Horizontal.2FVertical_Scrolling_Selector_.28.244106.29_.3E_write}

    7  bit  0
    ---- ----
    VPIx xxxx
    |||| ||||
    |||| |||+- H/V scrolling selector (0: Horizontal; 1: Vertical)
    ++++-+++-- unused

### <span id="Program_Bank_0_Register_0_.28.244107.29_.3E_write"></span>Program Bank 0 Register 0 (\$4107) \> write {#Program_Bank_0_Register_0_.28.244107.29_.3E_write}

    7  bit  0
    ---- ----
    PPPP PPPP
    |||| ||||
    |||| |||+- PQ00
    |||| ||+-- PQ01
    |||| |+--- PQ02
    |||| +---- PQ03
    |||+------ PQ04
    ||+------- PQ05
    |+-------- PQ06
    +--------- PQ07

### <span id="Program_Bank_0_Register_1_.28.244108.29_.3E_write"></span>Program Bank 0 Register 1 (\$4108) \> write {#Program_Bank_0_Register_1_.28.244108.29_.3E_write}

    7  bit  0
    ---- ----
    PPPP PPPP
    |||| ||||
    |||| |||+- PQ10
    |||| ||+-- PQ11
    |||| |+--- PQ12
    |||| +---- PQ13
    |||+------ PQ14
    ||+------- PQ15
    |+-------- PQ16
    +--------- PQ17

### <span id="Program_Bank_0_Register_2_.28.244109.29_.3E_write"></span>Program Bank 0 Register 2 (\$4109) \> write {#Program_Bank_0_Register_2_.28.244109.29_.3E_write}

    7  bit  0
    ---- ----
    PPPP PPPP
    |||| ||||
    |||| |||+- PQ20
    |||| ||+-- PQ21
    |||| |+--- PQ22
    |||| +---- PQ23
    |||+------ PQ24
    ||+------- PQ25
    |+-------- PQ26
    +--------- PQ27

### <span id="Program_Bank_0_Register_3_.28.24410A.29_.3E_write"></span>Program Bank 0 Register 3 (\$410A) \> write {#Program_Bank_0_Register_3_.28.24410A.29_.3E_write}

    7  bit  0
    ---- ----
    PPPP PPPP
    |||| ||||
    |||| |||+- PQ30
    |||| ||+-- PQ31
    |||| |+--- PQ32
    |||| +---- PQ33
    |||+------ PQ34
    ||+------- PQ35
    |+-------- PQ36
    +--------- PQ37

### <span id="Timer_interrupt_clock_selector.2C_Program_Bank_0_register_2_enable.2Fdisable.2C_RS232_enable.2Fdisable.2C_Bus_output_normal.2Ftristate.2C_Program_Bank_0_selector_.28.24410B.29_.3E_write"></span>Timer interrupt clock selector, Program Bank 0 register 2 enable/disable, RS232 enable/disable, Bus output normal/tristate, Program Bank 0 selector (\$410B) \> write {#Timer_interrupt_clock_selector.2C_Program_Bank_0_register_2_enable.2Fdisable.2C_RS232_enable.2Fdisable.2C_Bus_output_normal.2Ftristate.2C_Program_Bank_0_selector_.28.24410B.29_.3E_write}

    7  bit  0
    ---- ----
    TpRB XPPP
    |||| ||||
    |||| |+++- Program Bank 0 Selector
    |||| +---- XRWB switch (0: Writes to $6000-$FFFF will not activate XRWB; 1: writes to $6000-$FFFF will activate XRWB)
    |||+------ Bus output control (0: Normal bus output; 1: Tri-state bus)
    ||+------- RS232 (0: disabled; 1: enabled)
    |+-------- Program Bank 0 Register 2 enable (0: disabled; 1: enabled)
    +--------- Timer interrupt clock select (0: AD12; 1: HSYNC)

### <span id="I.2FO_Port_Control_.28.24410D.29_.3E_write"></span>I/O Port Control (\$410D) \> write {#I.2FO_Port_Control_.28.24410D.29_.3E_write}

    7  bit  0
    ---- ----
    DdCc BbAa
    |||| ||||
    |||| |||+- I/O port 0 mode (0: input; 1: output)
    |||| ||+-- I/O port 0 enable (0: disabled; 1:enabled)
    |||| |+--- I/O port 1 mode (0: input; 1: output)
    |||| +---- I/O port 1 enable (0: disabled; 1:enabled)
    |||+------ I/O port 2 mode (0: input; 1: output)
    ||+------- I/O port 2 enable (0: disabled; 1:enabled)
    |+-------- I/O port 3 mode (0: input; 1: output)
    +--------- I/O port 3 enable (0: disabled; 1:enabled)

- Bits d0-d3 must be set to \$A if using flash memory in 16-bit mode.
- External SRAM is not available if using flash memory in 16-bit mode.

### <span id="I.2FO_Port_Output_Data.2C_Ports_0_and_1_.28.24410E.29_.3C.3E_read.2Fwrite"></span>I/O Port Output Data, Ports 0 and 1 (\$410E) \<\> read/write {#I.2FO_Port_Output_Data.2C_Ports_0_and_1_.28.24410E.29_.3C.3E_read.2Fwrite}

    7  bit  0
    ---- ----
    DDDD dddd
    |||| ||||
    |||| ++++- I/O port 0 data (XVD3, XVD2, XVD1, XVD0)
    ++++------ I/O port 1 data (XVD7, XVD6, XVD5, XVD4)

### <span id="I.2FO_Port_Output_Data.2C_Ports_2_and_3_.28.24410F.29_.3C.3E_read.2Fwrite"></span>I/O Port Output Data, Ports 2 and 3 (\$410F) \<\> read/write {#I.2FO_Port_Output_Data.2C_Ports_2_and_3_.28.24410F.29_.3C.3E_read.2Fwrite}

    7  bit  0
    ---- ----
    DDDD dddd
    |||| ||||
    |||| ++++- I/O port 2 data (XAD12, XAD11, XAD10, XRA10)
    ++++------ I/O port 3 data (XVRW, VXOEB, XRCB, XRC)

### RS232 Timer {#RS232_Timer}

"In PAL system, CK21M is 26.601712MHz, in NTSC system is 21.47727MHz. RS232T=#4115,#4114 data. Baud rate will be CK21M/((RS232T+2)\*2). For example, In PAL system, the baud rate 9600, RS232T=0567."

#### <span id="RS232_Timer_Low_Byte_.28.244114.29_.3E_write"></span>RS232 Timer Low Byte (\$4114) \> write {#RS232_Timer_Low_Byte_.28.244114.29_.3E_write}

Write the low byte of the RS232 timer.

#### <span id="RS232_Timer_High_Byte_.28.244115.29_.3E_write"></span>RS232 Timer High Byte (\$4115) \> write {#RS232_Timer_High_Byte_.28.244115.29_.3E_write}

Write the high byte of the RS232 timer.

### <span id="RS232_Register_.28.244119.29_.3E_write"></span>RS232 Register (\$4119) \> write {#RS232_Register_.28.244119.29_.3E_write}

    7  bit  0
    ---- ----
    xxBx xxxT
    |||| ||||
    |||| |||+- TX bit 8
    |||+-+++-- unused
    ||+------- Bit 8 enable (0: 10 bits mode including start, end bits, and bit 7-0 data; 1: 11 bits mode including start, end bits, bit 8 and bit7-0 data)
    ++-------- unused

### <span id="RS232_Flags_.28.244119.29_.3C_read"></span>RS232 Flags (\$4119) \< read {#RS232_Flags_.28.244119.29_.3C_read}

    7  bit  0
    ---- ----
    rsTF PxER
    |||| ||||
    |||| |||+- RX bit 8
    |||| ||+-- RERRF (If 1, receiving error occurred)
    |||| |+--- unused
    |||| +---- XPORN (0: NTSC; 1: PAL)
    |||+------ XF5OR6 (0: 60Hz; 1: 50Hz)
    ||+------- RINGF (If 1, currently receiving data)
    |+-------- TIFLAG (If 1, completed sending data)
    +--------- RIFLAG (If 1, completed receiving data)

### <span id="RS232_TX_data_.28.24411A.29_.3E_write"></span>RS232 TX data (\$411A) \> write {#RS232_TX_data_.28.24411A.29_.3E_write}

Write RS232 data (bits 7-0).

### <span id="RS232_RX_data_.28.24411B.29_.3C_read"></span>RS232 RX data (\$411B) \< read {#RS232_RX_data_.28.24411B.29_.3C_read}

Read RS232 data (bits 7-0).

## Graphics Registers {#Graphics_Registers}

### <span id="PPU_Control_.28.242000.29_.3E_write"></span>PPU Control (\$2000) \> write {#PPU_Control_.28.242000.29_.3E_write}

    7  bit  0
    ---- ----
    VxHB SINN
    |||| ||||
    |||| ||++- Base nametable address
    |||| ||    (0 = $2000; 1 = $2400; 2 = $2800; 3 = $2C00)
    |||| |+--- VRAM address increment (0: +1/horizontal; 1: +32/vertical)
    |||| +---- Sprite pattern table address (0: $0000; 1: $1000)
    |||+------ Background pattern table address (0: $0000; 1: $1000)
    ||+------- Sprite Size (0: Small Sprite (8x8 or 16x8); 1: Big Sprite (8x16 or 16x16))
    |+-------- unused ("Test pin control I/O")
    +--------- NMI control (0: enabled; 1: disabled)

### <span id="PPU_Mask_.28.242001.29_.3E_write"></span>PPU Mask (\$2001) \> write {#PPU_Mask_.28.242001.29_.3E_write}

    7  bit  0
    ---- ----
    xxxs bMmG
    |||| ||||
    |||| |||+- Grayscale (0: Color; 1: Grayscale)
    |||| ||+-- "Background initial coordinate" (0: "righter"; 1: "lefter")
    |||| |+--- "Sprite initial coordinate" (0: "righter"; 1: "lefter")
    |||| +---- Show background (0: disabled; 1: enabled)
    |||+------ Show sprites (0: disabled; 1: enabled)
    +++------- unused

### <span id="PPU_Status_.28.242002.29_.3C_read"></span>PPU Status (\$2002) \< read {#PPU_Status_.28.242002.29_.3C_read}

    7  bit  0
    ---- ----
    VSOx xxxx
    |||| ||||
    |||+-++++- unused
    ||+------- Sprite overload indicator
    |+-------- "Priority indicator"
    +--------- Blanking indicator (0: display; 1: blanking)

### <span id="OAM_Address_.28.242003.29_.3E_write"></span>OAM Address (\$2003) \> write {#OAM_Address_.28.242003.29_.3E_write}

"Set sprite pool counter initial data by this register."

    7  bit  0
    ---- ----
    AAAA AAAA
    |||| ||||
    ++++-++++- Initial sprite data address

### <span id="OAM_Data_.28.242004.29_.3E_write"></span>OAM Data (\$2004) \> write {#OAM_Data_.28.242004.29_.3E_write}

"Write data into sprite pool and increment sprite counter."

    7  bit  0
    ---- ----
    DDDD DDDD
    |||| ||||
    ++++-++++- Data to write

### <span id="PPU_Scroll_.28.242005.29_.3E.3E_write_x2"></span>PPU Scroll (\$2005) \>\> write x2 {#PPU_Scroll_.28.242005.29_.3E.3E_write_x2}

### <span id="PPU_Address_.28.242006.29_.3E.3E_write_x2"></span>PPU Address (\$2006) \>\> write x2 {#PPU_Address_.28.242006.29_.3E.3E_write_x2}

### <span id="PPU_Data_.28.242007.29_.3C.3E_read.2Fwrite"></span>PPU Data (\$2007) \<\> read/write {#PPU_Data_.28.242007.29_.3C.3E_read.2Fwrite}

### <span id="Extended_Graphics_Control_1_.28.242010.29_.3E_write"></span>Extended Graphics Control 1 (\$2010) \> write {#Extended_Graphics_Control_1_.28.242010.29_.3E_write}

"Old color compatible, Background/Sprite address Extension enable, Sprite 16 colors or 16 pixels enable, Background 16 colors enable, Sprite 16 colors or 16 pixels selector."

    7  bit  0
    ---- ----
    Cxxb sSBM
    |||| ||||
    |||| |||+- Sprite mode (0: 16 colors; 1: 16 pixels)
    |||| ||+-- 16 color backgrounds (0: disabled, 1: enabled)
    |||| |+--- 16 color/pixel sprites (0: disabled, 1: enabled)
    |||| +---- Sprite address extension (0: disabled; 1: enabled)
    |||+------ Background address extension (0: disabled; 1: enabled)
    |++------- unused
    +--------- Color Mode (0: old/compatible; 1: new color mapping)

#### New Color Mapping Mode {#New_Color_Mapping_Mode}

<table class="wikitable">
<tbody>
<tr>
<th>RC=1</th>
<td colspan="6">$3F80,$3F81,$3F82...</td>
<td colspan="6">$3F00,$3F01,$3F02...</td>
</tr>
<tr>
<th>Data</th>
<td>D5</td>
<td>D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
<td>D5</td>
<td>D4</td>
<td>D3</td>
<td>D2</td>
<td>D1</td>
<td>D0</td>
</tr>
<tr>
<th>Function</th>
<td>SAT3</td>
<td>SAT2</td>
<td>SAT1</td>
<td>SAT0</td>
<td>LUM3</td>
<td>LUM2</td>
<td>LUM1</td>
<td>LUM0</td>
<td>PHA3</td>
<td>PHA2</td>
<td>PHA1</td>
<td>PHA0</td>
</tr>
</tbody>
</table>

> 4 <= LUM[3:0] X 2 + SAT[3:0] <= 1F
>     If you set LUM = F, SAT must be <= 1.
>     LUM = E, SAT must be <= 3.
>     .
>     .
>     .
>     LUM = 3, SAT must be <= 2.
>     LUM = 2, SAT must be = 0.

### <span id="Extended_Graphics_Control_2_.28.242011.29_.3E_write"></span>Extended Graphics Control 2 (\$2011) \> write {#Extended_Graphics_Control_2_.28.242011.29_.3E_write}

"Option of Vertical line number of LCD display, B/W 2 color mode, Composted Video DA Enable, Video Extension Address EVA12 selector, Enable the internal VRAM or not"

    7  bit  0
    ---- ----
    xxVV IBCE
    |||| ||||
    |||| |||+- Video Extension Address EVA selector (0: reg. BKPAGE; 1: HV)
    |||| ||+-- Composited Video DA (0: enabled; 1: disabled)
    |||| |+--- Black/White two color mode (0: disabled; 1: enabled)
    |||| +---- Internal VRAM enable (0: enabled; 1: disabled)
    ||++------ LCD display vertical lines (0: 240 lines; 1: 160 lines; 2: 120 lines; 3: 80 lines)
    ++-------- unused

### <span id="Video_Bank_0_Register_0_.28.242012.29_.3E_write"></span>Video Bank 0 Register 0 (\$2012) \> write {#Video_Bank_0_Register_0_.28.242012.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV00
    |||| ||+-- RV01
    |||| |+--- RV02
    |||| +---- RV03
    |||+------ RV04
    ||+------- RV05
    |+-------- RV06
    +--------- RV07

### <span id="Video_Bank_0_Register_1_.28.242013.29_.3E_write"></span>Video Bank 0 Register 1 (\$2013) \> write {#Video_Bank_0_Register_1_.28.242013.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV10
    |||| ||+-- RV11
    |||| |+--- RV12
    |||| +---- RV13
    |||+------ RV14
    ||+------- RV15
    |+-------- RV16
    +--------- RV17

### <span id="Video_Bank_0_Register_2_.28.242014.29_.3E_write"></span>Video Bank 0 Register 2 (\$2014) \> write {#Video_Bank_0_Register_2_.28.242014.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV20
    |||| ||+-- RV21
    |||| |+--- RV22
    |||| +---- RV23
    |||+------ RV24
    ||+------- RV25
    |+-------- RV26
    +--------- RV27

### <span id="Video_Bank_0_Register_3_.28.242015.29_.3E_write"></span>Video Bank 0 Register 3 (\$2015) \> write {#Video_Bank_0_Register_3_.28.242015.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV30
    |||| ||+-- RV31
    |||| |+--- RV32
    |||| +---- RV33
    |||+------ RV34
    ||+------- RV35
    |+-------- RV36
    +--------- RV37

### <span id="Video_Bank_0_Register_4_.28.242016.29_.3E_write"></span>Video Bank 0 Register 4 (\$2016) \> write {#Video_Bank_0_Register_4_.28.242016.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV40
    |||| ||+-- RV41
    |||| |+--- RV42
    |||| +---- RV43
    |||+------ RV44
    ||+------- RV45
    |+-------- RV46
    +--------- RV47

### <span id="Video_Bank_0_Register_5_.28.242017.29_.3E_write"></span>Video Bank 0 Register 5 (\$2017) \> write {#Video_Bank_0_Register_5_.28.242017.29_.3E_write}

    7  bit  0
    ---- ----
    7654 3210
    |||| ||||
    |||| |||+- RV50
    |||| ||+-- RV51
    |||| |+--- RV52
    |||| +---- RV53
    |||+------ RV54
    ||+------- RV55
    |+-------- RV56
    +--------- RV57

### <span id="Video_Bank_1_Register.2C_BKPAGE.2C_Video_RW_Bank_.28.242018.29_.3E_write"></span>Video Bank 1 Register, BKPAGE, Video RW Bank (\$2018) \> write {#Video_Bank_1_Register.2C_BKPAGE.2C_Video_RW_Bank_.28.242018.29_.3E_write}

    7  bit  0
    ---- ----
    xRRR BVVV
    |||| ||||
    |||| |+++- Video bank when accessing video data
    |||| +---- Reg. BKPAGE address is EVA12 when EVA12S=0
    |+++------ Video Bank 1 Register
    +--------- unused

### <span id="Gun_Port_Reset_.28.242019.29_.3E_write"></span>Gun Port Reset (\$2019) \> write {#Gun_Port_Reset_.28.242019.29_.3E_write}

Writing any value to this register will reset the X and Y coordinates of the guns in ports 1 and 2.

### <span id="Video_Bank_0_Register_6.2C_Video_Bank_0_Selector_.28.24201A.29_.3E_write"></span>Video Bank 0 Register 6, Video Bank 0 Selector (\$201A) \> write {#Video_Bank_0_Register_6.2C_Video_Bank_0_Selector_.28.24201A.29_.3E_write}

    7  bit  0
    ---- ----
    RRRR RBBB
    |||| ||||
    |||| |+++- Video Bank 0 selector
    ++++-+---- Video Bank 0 register 6

### <span id="Gun_Port_1_X_Coordinate_.28.24201C.29_.3C_read"></span>Gun Port 1 X Coordinate (\$201C) \< read {#Gun_Port_1_X_Coordinate_.28.24201C.29_.3C_read}

    7  bit  0
    ---- ----
    XXXX XXXX
    |||| ||||
    ++++-++++- Gun port 1 X coordinate

### <span id="Gun_Port_1_Y_Coordinate_.28.24201D.29_.3C_read"></span>Gun Port 1 Y Coordinate (\$201D) \< read {#Gun_Port_1_Y_Coordinate_.28.24201D.29_.3C_read}

    7  bit  0
    ---- ----
    YYYY YYYY
    |||| ||||
    ++++-++++- Gun port 1 Y coordinate

### <span id="Gun_Port_2_X_Coordinate_.28.24201E.29_.3C_read"></span>Gun Port 2 X Coordinate (\$201E) \< read {#Gun_Port_2_X_Coordinate_.28.24201E.29_.3C_read}

    7  bit  0
    ---- ----
    XXXX XXXX
    |||| ||||
    ++++-++++- Gun port 2 X coordinate

### <span id="Gun_Port_2_Y_Coordinate_.28.24201F.29_.3C_read"></span>Gun Port 2 Y Coordinate (\$201F) \< read {#Gun_Port_2_Y_Coordinate_.28.24201F.29_.3C_read}

    7  bit  0
    ---- ----
    YYYY YYYY
    |||| ||||
    ++++-++++- Gun port 2 Y coordinate

## Sound Registers {#Sound_Registers}

## Miscellaneous Registers {#Miscellaneous_Registers}

### <span id="DMA_Source_Address_High_Byte_.28.244014.29_.3E_write"></span>DMA Source Address High Byte (\$4014) \> write {#DMA_Source_Address_High_Byte_.28.244014.29_.3E_write}

Two bytes are required to set a DMA source address. This register sets the high byte of the address (\$\[XX\]X0). Writing to this register also starts the DMA. The VT03 is capable of DMA for both video and sprite data; see register \$4034 for more details.

    7  bit  0
    ---- ----
    AAAA AAAA
    |||| ||||
    ++++-++++- DMA source address high byte

### <span id="DMA_Settings_.28.244034.29_.3E_write"></span>DMA Settings (\$4034) \> write {#DMA_Settings_.28.244034.29_.3E_write}

    7  bit  0
    ---- ----
    SSSS LLLT
    |||| ||||
    |||| |||+- DMA type (0: Sprite data, using $2004; 1: Video data, using $2007)
    |||| +++-- DMA length (000: 256 bytes; 100: 16 bytes; 101: 32 bytes; 110: 64 bytes; 111: 128 bytes)
    +++------- DMA source address low byte (d7-d4)

"Under 64byte mode, VT03 cut the memory into 4 pieces. If you want to access complete 64 bytes. The low byte of address must be 00H, 40H, 80H or C0H, because VT03 will stop accessing when address counted to 3FH, 7FH, BFH or FFH respectively. Under 16 byte mode, VT03 cut the memory into 16 pieces. Under 128 byte mode, VT03 cut the memory into 2 pieces."

### <span id="Enable.2FDisable_XOP1_.26_DWS_IRQ_.28.244015.29_.3E_write"></span>Enable/Disable XOP1 & DWS IRQ (\$4015) \> write {#Enable.2FDisable_XOP1_.26_DWS_IRQ_.28.244015.29_.3E_write}

Similar to [\$4015 on the NES APU](../APU/#Status_.28.244015.29).

    7  bit  0
    ---- ----
    xxxD NT21
    |||| ||||
    |||| |||+- Rhythm/Pulse A (0: stop; 1: start)
    |||| ||+-- Rhythm/Pulse B (0: stop; 1: start)
    |||| |+--- Envelope/Triangle (0: stop; 1: start)
    |||| +---- Noise (0: stop; 1: start)
    |||+------ DWS/PCM (0: stop; 1: start)
    +++------- unused

### <span id="Read_XOP1_Flag_.28.244015.29_.3C_read"></span>Read XOP1 Flag (\$4015) \< read {#Read_XOP1_Flag_.28.244015.29_.3C_read}

Similar to \[[on the NES APU](http://wiki.nesdev.org/w/index.php/APU#Status_.28.244015.29%7C$4015)\].

    7  bit  0
    ---- ----
    IFxD NT21
    |||| ||||
    |||| |||+- Rhythm/Pulse A (0: ended; 1: playing)
    |||| ||+-- Rhythm/Pulse B (0: ended; 1: playing)
    |||| |+--- Envelope/Triangle (0: ended; 1: playing)
    |||| +---- Noise (0: ended; 1: playing)
    |||+------ DWS/PCM (0: ended; 1: playing)
    ||+------- unused
    |+-------- IRQ (0: inactive; 1: active)
    +--------- DWS/PCM IRQ (0: inactive; 1: active)

### <span id="Enable.2FDisable_XOP2_.28.244035.29_.3E_write"></span>Enable/Disable XOP2 (\$4035) \> write {#Enable.2FDisable_XOP2_.28.244035.29_.3E_write}

    7  bit  0
    ---- ----
    xxxx NT21
    |||| ||||
    |||| |||+- Rhythm/Pulse A (0: stop; 1: start)
    |||| ||+-- Rhythm/Pulse B (0: stop; 1: start)
    |||| |+--- Envelope/Triangle (0: stop; 1: start)
    |||| +---- Noise (0: stop; 1: start)
    ++++------ unused

### <span id="Read_XOP2_Flag_.28.244035.29_.3C_read"></span>Read XOP2 Flag (\$4035) \< read {#Read_XOP2_Flag_.28.244035.29_.3C_read}

    7  bit  0
    ---- ----
    xxxx NT21
    |||| ||||
    |||| |||+- Rhythm/Pulse A (0: ended; 1: playing)
    |||| ||+-- Rhythm/Pulse B (0: ended; 1: playing)
    |||| |+--- Envelope/Triangle (0: ended; 1: playing)
    |||| +---- Noise (0: ended; 1: playing)
    ++++------ unused

### <span id="Set_Output_Pin_XQ.5B2:0.5D_.28.244016.29_.3E_write"></span>Set Output Pin XQ\[2:0\] (\$4016) \> write {#Set_Output_Pin_XQ.5B2:0.5D_.28.244016.29_.3E_write}

    7  bit  0
    ---- ----
    xxxx x210
    |||| ||||
    |||| |||+- Set output pin XQ0
    |||| ||+-- Set output pin XQ1
    |||| |+--- Set output pin XQ2
    ++++-+---- unused

### <span id="Read_peripheral_data_.28.244016.29_.3C_read"></span>Read peripheral data (\$4016) \< read {#Read_peripheral_data_.28.244016.29_.3C_read}

    7  bit  0
    ---- ----
    xxxx xMFJ
    |||| ||||
    |||| |||+- First joystick data
    |||| ||+-- Floppy disk data
    |||| |+--- Microphone data
    ++++-+---- unused

### <span id="APU_Frame_Counter_.28.244017.29_.3E_write"></span>APU Frame Counter (\$4017) \> write {#APU_Frame_Counter_.28.244017.29_.3E_write}

    7  bit  0
    ---- ----
    MIxx xxxx
    |||| ||||
    ||++-++++- unused
    |+-------- Clock IRQ (0: enabled, 60Hz; 1: disabled)
    +--------- BLCK1/BLCK2 (0: BLCK1=250Hz, BLCK2=120Hz; 1: BLCK1=200Hz, BLCK2=100Hz)

### <span id="Read_peripheral_data_.28.244017.29_.3C_read"></span>Read peripheral data (\$4017) \< read {#Read_peripheral_data_.28.244017.29_.3C_read}

    7  bit  0
    ---- ----
    xxxF FFFJ
    |||| ||||
    |||| |||+- Second joystick data
    |||+-+++-- Floppy disk data
    +++------- unused

## References {#References}

- [Datasheets](http://www.vrt.com.tw/datasheet.htm)
- [Downloads](http://www.vrt.com.tw/download.htm) (including EmuVT)
