---
title: "Pulse Channel frequency chart"
url: "/wiki/Pulse_Channel_frequency_chart/"
description: "Using the MIDI Tuning Standard with A=440 Hz, we can derive the equivalent period timer value for the two Pulse channels. Due to limitations on the 11-bit period's range, the lowest standard MIDI…"
wikiPage: "Pulse_Channel_frequency_chart"
---

## MIDI Tuning Table {#MIDI_Tuning_Table}

Using the [MIDI Tuning Standard](https://en.wikipedia.org/wiki/MIDI_tuning_standard) with A=440 Hz, we can derive the equivalent period timer value for the two [Pulse](../APU_Pulse/) channels. Due to limitations on the 11-bit period's range, the lowest standard MIDI frequency that can be reproduced with acceptable tuning is A-1 (55 Hz) on NTSC, and G#1 (51.91 Hz) on PAL. Here we show the desired frequency for each MIDI index, the closest period value, and the resulting tuning error (in cents) from the standard.

**Note**: The [Triangle](../APU_Triangle/) channel can make use of the same 11-bit period values. Due to its longer 32-step sequence, it will sound exactly one octave lower (half of the apparent frequency) compared to the equivalent period on Pulse.

<table class="wikitable" style="text-align: right">
<tbody>
<tr>
<td colspan="3" data-bgcolor="DDDDDD" style="text-align: center; font-weight: bold;">MIDI</td>
<td colspan="3" data-bgcolor="FFDDDD" style="text-align: center; font-weight: bold;">NTSC APU Pulse</td>
<td colspan="3" data-bgcolor="DDDDFF" style="text-align: center; font-weight: bold;">PAL APU Pulse</td>
</tr>
<tr>
<td style="text-align: center;">Index</td>
<td style="text-align: center;">Note</td>
<td style="text-align: center;">Frequency</td>
<td style="text-align: center;">11bit Period</td>
<td style="text-align: center;">Frequency</td>
<td style="text-align: center;">Tuning</td>
<td style="text-align: center;">11bit Period</td>
<td style="text-align: center;">Frequency</td>
<td style="text-align: center;">Tuning</td>
</tr>
<tr>
<td style="text-align: center;">31</td>
<td style="text-align: center;">G-1</td>
<td style="text-align: center;">49.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">-</td>
<td style="text-align: center;">-</td>
<td style="text-align: center;">-</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$7FF</td>
<td style="text-align: center;">50.74</td>
<td style="text-align: center;">+60c</td>
</tr>
<tr>
<td style="text-align: center;">32</td>
<td style="text-align: center;">G#1</td>
<td style="text-align: center;">51.91</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$7FF</td>
<td style="text-align: center;">54.62</td>
<td style="text-align: center;">+88c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$7D1</td>
<td style="text-align: center;">51.90</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">33</td>
<td style="text-align: center;">A-1</td>
<td style="text-align: center;">55.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$7F1</td>
<td style="text-align: center;">55.00</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$760</td>
<td style="text-align: center;">55.01</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">34</td>
<td style="text-align: center;">A#1</td>
<td style="text-align: center;">58.27</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$77F</td>
<td style="text-align: center;">58.26</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$6F6</td>
<td style="text-align: center;">58.28</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">35</td>
<td style="text-align: center;">B-1</td>
<td style="text-align: center;">61.74</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$713</td>
<td style="text-align: center;">61.73</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$692</td>
<td style="text-align: center;">61.74</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">36</td>
<td style="text-align: center;">C-2</td>
<td style="text-align: center;">65.41</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$6AD</td>
<td style="text-align: center;">65.42</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$634</td>
<td style="text-align: center;">65.40</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">37</td>
<td style="text-align: center;">C#2</td>
<td style="text-align: center;">69.30</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$64D</td>
<td style="text-align: center;">69.31</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$5DB</td>
<td style="text-align: center;">69.28</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">38</td>
<td style="text-align: center;">D-2</td>
<td style="text-align: center;">73.42</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$5F3</td>
<td style="text-align: center;">73.40</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$586</td>
<td style="text-align: center;">73.44</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">39</td>
<td style="text-align: center;">D#2</td>
<td style="text-align: center;">77.78</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$59D</td>
<td style="text-align: center;">77.79</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$537</td>
<td style="text-align: center;">77.78</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">40</td>
<td style="text-align: center;">E-2</td>
<td style="text-align: center;">82.41</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$54C</td>
<td style="text-align: center;">82.43</td>
<td style="text-align: center;">+1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$4EC</td>
<td style="text-align: center;">82.41</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">41</td>
<td style="text-align: center;">F-2</td>
<td style="text-align: center;">87.31</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$500</td>
<td style="text-align: center;">87.32</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$4A5</td>
<td style="text-align: center;">87.32</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">42</td>
<td style="text-align: center;">F#2</td>
<td style="text-align: center;">92.50</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$4B8</td>
<td style="text-align: center;">92.52</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$462</td>
<td style="text-align: center;">92.53</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">43</td>
<td style="text-align: center;">G-2</td>
<td style="text-align: center;">98.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$474</td>
<td style="text-align: center;">98.04</td>
<td style="text-align: center;">+1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$423</td>
<td style="text-align: center;">98.03</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">44</td>
<td style="text-align: center;">G#2</td>
<td style="text-align: center;">103.83</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$434</td>
<td style="text-align: center;">103.86</td>
<td style="text-align: center;">+1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$3E8</td>
<td style="text-align: center;">103.81</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">45</td>
<td style="text-align: center;">A-2</td>
<td style="text-align: center;">110.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$3F8</td>
<td style="text-align: center;">109.99</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$3B0</td>
<td style="text-align: center;">109.96</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">46</td>
<td style="text-align: center;">A#2</td>
<td style="text-align: center;">116.54</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$3BF</td>
<td style="text-align: center;">116.52</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$37B</td>
<td style="text-align: center;">116.49</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">47</td>
<td style="text-align: center;">B-2</td>
<td style="text-align: center;">123.47</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$389</td>
<td style="text-align: center;">123.47</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$349</td>
<td style="text-align: center;">123.41</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">48</td>
<td style="text-align: center;">C-3</td>
<td style="text-align: center;">130.81</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$356</td>
<td style="text-align: center;">130.83</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$319</td>
<td style="text-align: center;">130.87</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">49</td>
<td style="text-align: center;">C#3</td>
<td style="text-align: center;">138.59</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$326</td>
<td style="text-align: center;">138.61</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$2ED</td>
<td style="text-align: center;">138.55</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">50</td>
<td style="text-align: center;">D-3</td>
<td style="text-align: center;">146.83</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$2F9</td>
<td style="text-align: center;">146.80</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$2C3</td>
<td style="text-align: center;">146.77</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">51</td>
<td style="text-align: center;">D#3</td>
<td style="text-align: center;">155.56</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$2CE</td>
<td style="text-align: center;">155.58</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$29B</td>
<td style="text-align: center;">155.56</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">52</td>
<td style="text-align: center;">E-3</td>
<td style="text-align: center;">164.81</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$2A6</td>
<td style="text-align: center;">164.74</td>
<td style="text-align: center;">-1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$275</td>
<td style="text-align: center;">164.94</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">53</td>
<td style="text-align: center;">F-3</td>
<td style="text-align: center;">174.61</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$280</td>
<td style="text-align: center;">174.51</td>
<td style="text-align: center;">-1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$252</td>
<td style="text-align: center;">174.64</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">54</td>
<td style="text-align: center;">F#3</td>
<td style="text-align: center;">185.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$25C</td>
<td style="text-align: center;">184.89</td>
<td style="text-align: center;">-1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$231</td>
<td style="text-align: center;">184.90</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">55</td>
<td style="text-align: center;">G-3</td>
<td style="text-align: center;">196.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$23A</td>
<td style="text-align: center;">195.90</td>
<td style="text-align: center;">-1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$211</td>
<td style="text-align: center;">196.06</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">56</td>
<td style="text-align: center;">G#3</td>
<td style="text-align: center;">207.65</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$21A</td>
<td style="text-align: center;">207.53</td>
<td style="text-align: center;">-1c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$1F3</td>
<td style="text-align: center;">207.83</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">57</td>
<td style="text-align: center;">A-3</td>
<td style="text-align: center;">220.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$1FB</td>
<td style="text-align: center;">220.20</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$1D7</td>
<td style="text-align: center;">220.15</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">58</td>
<td style="text-align: center;">A#3</td>
<td style="text-align: center;">233.08</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$1DF</td>
<td style="text-align: center;">233.04</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$1BD</td>
<td style="text-align: center;">232.99</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">59</td>
<td style="text-align: center;">B-3</td>
<td style="text-align: center;">246.94</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$1C4</td>
<td style="text-align: center;">246.93</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$1A4</td>
<td style="text-align: center;">246.82</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">60</td>
<td style="text-align: center;">C-4</td>
<td style="text-align: center;">261.63</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$1AB</td>
<td style="text-align: center;">261.36</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$18C</td>
<td style="text-align: center;">261.75</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">61</td>
<td style="text-align: center;">C#4</td>
<td style="text-align: center;">277.18</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$193</td>
<td style="text-align: center;">276.88</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$176</td>
<td style="text-align: center;">277.10</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">62</td>
<td style="text-align: center;">D-4</td>
<td style="text-align: center;">293.66</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$17C</td>
<td style="text-align: center;">293.60</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$161</td>
<td style="text-align: center;">293.54</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">63</td>
<td style="text-align: center;">D#4</td>
<td style="text-align: center;">311.13</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$167</td>
<td style="text-align: center;">310.72</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$14D</td>
<td style="text-align: center;">311.12</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">64</td>
<td style="text-align: center;">E-4</td>
<td style="text-align: center;">329.63</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$152</td>
<td style="text-align: center;">329.97</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$13A</td>
<td style="text-align: center;">329.88</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">65</td>
<td style="text-align: center;">F-4</td>
<td style="text-align: center;">349.23</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$13F</td>
<td style="text-align: center;">349.57</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$129</td>
<td style="text-align: center;">348.70</td>
<td style="text-align: center;">-3c</td>
</tr>
<tr>
<td style="text-align: center;">66</td>
<td style="text-align: center;">F#4</td>
<td style="text-align: center;">369.99</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$12D</td>
<td style="text-align: center;">370.40</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$118</td>
<td style="text-align: center;">369.80</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">67</td>
<td style="text-align: center;">G-4</td>
<td style="text-align: center;">392.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$11C</td>
<td style="text-align: center;">392.49</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$108</td>
<td style="text-align: center;">392.12</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">68</td>
<td style="text-align: center;">G#4</td>
<td style="text-align: center;">415.30</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$10C</td>
<td style="text-align: center;">415.84</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0F9</td>
<td style="text-align: center;">415.65</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">69</td>
<td style="text-align: center;">A-4</td>
<td style="text-align: center;">440.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0FD</td>
<td style="text-align: center;">440.40</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0EB</td>
<td style="text-align: center;">440.31</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">70</td>
<td style="text-align: center;">A#4</td>
<td style="text-align: center;">466.16</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0EF</td>
<td style="text-align: center;">466.09</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0DE</td>
<td style="text-align: center;">465.98</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">71</td>
<td style="text-align: center;">B-4</td>
<td style="text-align: center;">493.88</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0E1</td>
<td style="text-align: center;">494.96</td>
<td style="text-align: center;">+4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0D1</td>
<td style="text-align: center;">494.82</td>
<td style="text-align: center;">+3c</td>
</tr>
<tr>
<td style="text-align: center;">72</td>
<td style="text-align: center;">C-5</td>
<td style="text-align: center;">523.25</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0D5</td>
<td style="text-align: center;">522.71</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0C6</td>
<td style="text-align: center;">522.18</td>
<td style="text-align: center;">-4c</td>
</tr>
<tr>
<td style="text-align: center;">73</td>
<td style="text-align: center;">C#5</td>
<td style="text-align: center;">554.37</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0C9</td>
<td style="text-align: center;">553.77</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0BA</td>
<td style="text-align: center;">555.68</td>
<td style="text-align: center;">+4c</td>
</tr>
<tr>
<td style="text-align: center;">74</td>
<td style="text-align: center;">D-5</td>
<td style="text-align: center;">587.33</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0BD</td>
<td style="text-align: center;">588.74</td>
<td style="text-align: center;">+4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0B0</td>
<td style="text-align: center;">587.08</td>
<td style="text-align: center;">-1c</td>
</tr>
<tr>
<td style="text-align: center;">75</td>
<td style="text-align: center;">D#5</td>
<td style="text-align: center;">622.25</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0B3</td>
<td style="text-align: center;">621.45</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$0A6</td>
<td style="text-align: center;">622.23</td>
<td style="text-align: center;">+0c</td>
</tr>
<tr>
<td style="text-align: center;">76</td>
<td style="text-align: center;">E-5</td>
<td style="text-align: center;">659.26</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$0A9</td>
<td style="text-align: center;">658.00</td>
<td style="text-align: center;">-3c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$09D</td>
<td style="text-align: center;">657.68</td>
<td style="text-align: center;">-4c</td>
</tr>
<tr>
<td style="text-align: center;">77</td>
<td style="text-align: center;">F-5</td>
<td style="text-align: center;">698.46</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$09F</td>
<td style="text-align: center;">699.13</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$094</td>
<td style="text-align: center;">697.40</td>
<td style="text-align: center;">-3c</td>
</tr>
<tr>
<td style="text-align: center;">78</td>
<td style="text-align: center;">F#5</td>
<td style="text-align: center;">739.99</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$096</td>
<td style="text-align: center;">740.80</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$08B</td>
<td style="text-align: center;">742.24</td>
<td style="text-align: center;">+5c</td>
</tr>
<tr>
<td style="text-align: center;">79</td>
<td style="text-align: center;">G-5</td>
<td style="text-align: center;">783.99</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$08E</td>
<td style="text-align: center;">782.24</td>
<td style="text-align: center;">-4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$084</td>
<td style="text-align: center;">781.30</td>
<td style="text-align: center;">-6c</td>
</tr>
<tr>
<td style="text-align: center;">80</td>
<td style="text-align: center;">G#5</td>
<td style="text-align: center;">830.61</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$086</td>
<td style="text-align: center;">828.60</td>
<td style="text-align: center;">-4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$07C</td>
<td style="text-align: center;">831.30</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">81</td>
<td style="text-align: center;">A-5</td>
<td style="text-align: center;">880.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$07E</td>
<td style="text-align: center;">880.79</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$075</td>
<td style="text-align: center;">880.62</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">82</td>
<td style="text-align: center;">A#5</td>
<td style="text-align: center;">932.33</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$077</td>
<td style="text-align: center;">932.17</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$06E</td>
<td style="text-align: center;">936.15</td>
<td style="text-align: center;">+7c</td>
</tr>
<tr>
<td style="text-align: center;">83</td>
<td style="text-align: center;">B-5</td>
<td style="text-align: center;">987.77</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$070</td>
<td style="text-align: center;">989.92</td>
<td style="text-align: center;">+4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$068</td>
<td style="text-align: center;">989.65</td>
<td style="text-align: center;">+3c</td>
</tr>
<tr>
<td style="text-align: center;">84</td>
<td style="text-align: center;">C-6</td>
<td style="text-align: center;">1046.50</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$06A</td>
<td style="text-align: center;">1045.43</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$062</td>
<td style="text-align: center;">1049.63</td>
<td style="text-align: center;">+5c</td>
</tr>
<tr>
<td style="text-align: center;">85</td>
<td style="text-align: center;">C#6</td>
<td style="text-align: center;">1108.73</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$064</td>
<td style="text-align: center;">1107.53</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$05D</td>
<td style="text-align: center;">1105.46</td>
<td style="text-align: center;">-5c</td>
</tr>
<tr>
<td style="text-align: center;">86</td>
<td style="text-align: center;">D-6</td>
<td style="text-align: center;">1174.66</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$05E</td>
<td style="text-align: center;">1177.48</td>
<td style="text-align: center;">+4c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$057</td>
<td style="text-align: center;">1180.83</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">87</td>
<td style="text-align: center;">D#6</td>
<td style="text-align: center;">1244.51</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$059</td>
<td style="text-align: center;">1242.90</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$052</td>
<td style="text-align: center;">1251.96</td>
<td style="text-align: center;">+10c</td>
</tr>
<tr>
<td style="text-align: center;">88</td>
<td style="text-align: center;">E-6</td>
<td style="text-align: center;">1318.51</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$054</td>
<td style="text-align: center;">1316.01</td>
<td style="text-align: center;">-3c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$04E</td>
<td style="text-align: center;">1315.35</td>
<td style="text-align: center;">-4c</td>
</tr>
<tr>
<td style="text-align: center;">89</td>
<td style="text-align: center;">F-6</td>
<td style="text-align: center;">1396.91</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$04F</td>
<td style="text-align: center;">1398.26</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$049</td>
<td style="text-align: center;">1404.23</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">90</td>
<td style="text-align: center;">F#6</td>
<td style="text-align: center;">1479.98</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$04B</td>
<td style="text-align: center;">1471.85</td>
<td style="text-align: center;">-10c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$045</td>
<td style="text-align: center;">1484.47</td>
<td style="text-align: center;">+5c</td>
</tr>
<tr>
<td style="text-align: center;">91</td>
<td style="text-align: center;">G-6</td>
<td style="text-align: center;">1567.98</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$046</td>
<td style="text-align: center;">1575.50</td>
<td style="text-align: center;">+8c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$041</td>
<td style="text-align: center;">1574.44</td>
<td style="text-align: center;">+7c</td>
</tr>
<tr>
<td style="text-align: center;">92</td>
<td style="text-align: center;">G#6</td>
<td style="text-align: center;">1661.22</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$042</td>
<td style="text-align: center;">1669.56</td>
<td style="text-align: center;">+9c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$03E</td>
<td style="text-align: center;">1649.41</td>
<td style="text-align: center;">-12c</td>
</tr>
<tr>
<td style="text-align: center;">93</td>
<td style="text-align: center;">A-6</td>
<td style="text-align: center;">1760.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$03F</td>
<td style="text-align: center;">1747.83</td>
<td style="text-align: center;">-12c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$03A</td>
<td style="text-align: center;">1761.24</td>
<td style="text-align: center;">+1c</td>
</tr>
<tr>
<td style="text-align: center;">94</td>
<td style="text-align: center;">A#6</td>
<td style="text-align: center;">1864.66</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$03B</td>
<td style="text-align: center;">1864.35</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$037</td>
<td style="text-align: center;">1855.59</td>
<td style="text-align: center;">-8c</td>
</tr>
<tr>
<td style="text-align: center;">95</td>
<td style="text-align: center;">B-6</td>
<td style="text-align: center;">1975.53</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$038</td>
<td style="text-align: center;">1962.47</td>
<td style="text-align: center;">-11c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$034</td>
<td style="text-align: center;">1960.62</td>
<td style="text-align: center;">-13c</td>
</tr>
<tr>
<td style="text-align: center;">96</td>
<td style="text-align: center;">C-7</td>
<td style="text-align: center;">2093.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$034</td>
<td style="text-align: center;">2110.58</td>
<td style="text-align: center;">+14c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$031</td>
<td style="text-align: center;">2078.26</td>
<td style="text-align: center;">-12c</td>
</tr>
<tr>
<td style="text-align: center;">97</td>
<td style="text-align: center;">C#7</td>
<td style="text-align: center;">2217.46</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$031</td>
<td style="text-align: center;">2237.22</td>
<td style="text-align: center;">+15c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$02E</td>
<td style="text-align: center;">2210.91</td>
<td style="text-align: center;">-5c</td>
</tr>
<tr>
<td style="text-align: center;">98</td>
<td style="text-align: center;">D-7</td>
<td style="text-align: center;">2349.32</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$02F</td>
<td style="text-align: center;">2330.43</td>
<td style="text-align: center;">-14c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$02B</td>
<td style="text-align: center;">2361.66</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">99</td>
<td style="text-align: center;">D#7</td>
<td style="text-align: center;">2489.02</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$02C</td>
<td style="text-align: center;">2485.80</td>
<td style="text-align: center;">-2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$029</td>
<td style="text-align: center;">2474.12</td>
<td style="text-align: center;">-10c</td>
</tr>
<tr>
<td style="text-align: center;">100</td>
<td style="text-align: center;">E-7</td>
<td style="text-align: center;">2637.02</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$029</td>
<td style="text-align: center;">2663.35</td>
<td style="text-align: center;">+17c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$026</td>
<td style="text-align: center;">2664.43</td>
<td style="text-align: center;">+18c</td>
</tr>
<tr>
<td style="text-align: center;">101</td>
<td style="text-align: center;">F-7</td>
<td style="text-align: center;">2793.83</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$027</td>
<td style="text-align: center;">2796.52</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$024</td>
<td style="text-align: center;">2808.46</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">102</td>
<td style="text-align: center;">F#7</td>
<td style="text-align: center;">2959.96</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$025</td>
<td style="text-align: center;">2943.71</td>
<td style="text-align: center;">-10c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$022</td>
<td style="text-align: center;">2968.94</td>
<td style="text-align: center;">+5c</td>
</tr>
<tr>
<td style="text-align: center;">103</td>
<td style="text-align: center;">G-7</td>
<td style="text-align: center;">3135.96</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$023</td>
<td style="text-align: center;">3107.24</td>
<td style="text-align: center;">-16c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$020</td>
<td style="text-align: center;">3148.88</td>
<td style="text-align: center;">+7c</td>
</tr>
<tr>
<td style="text-align: center;">104</td>
<td style="text-align: center;">G#7</td>
<td style="text-align: center;">3322.44</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$021</td>
<td style="text-align: center;">3290.02</td>
<td style="text-align: center;">-17c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$01E</td>
<td style="text-align: center;">3352.03</td>
<td style="text-align: center;">+15c</td>
</tr>
<tr>
<td style="text-align: center;">105</td>
<td style="text-align: center;">A-7</td>
<td style="text-align: center;">3520.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$01F</td>
<td style="text-align: center;">3495.65</td>
<td style="text-align: center;">-12c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$01D</td>
<td style="text-align: center;">3463.76</td>
<td style="text-align: center;">-28c</td>
</tr>
<tr>
<td style="text-align: center;">106</td>
<td style="text-align: center;">A#7</td>
<td style="text-align: center;">3729.31</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$01D</td>
<td style="text-align: center;">3728.69</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$01B</td>
<td style="text-align: center;">3711.18</td>
<td style="text-align: center;">-8c</td>
</tr>
<tr>
<td style="text-align: center;">107</td>
<td style="text-align: center;">B-7</td>
<td style="text-align: center;">3951.07</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$01B</td>
<td style="text-align: center;">3995.03</td>
<td style="text-align: center;">+19c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$019</td>
<td style="text-align: center;">3996.65</td>
<td style="text-align: center;">+20c</td>
</tr>
<tr>
<td style="text-align: center;">108</td>
<td style="text-align: center;">C-8</td>
<td style="text-align: center;">4186.01</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$01A</td>
<td style="text-align: center;">4142.99</td>
<td style="text-align: center;">-18c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$018</td>
<td style="text-align: center;">4156.52</td>
<td style="text-align: center;">-12c</td>
</tr>
<tr>
<td style="text-align: center;">109</td>
<td style="text-align: center;">C#8</td>
<td style="text-align: center;">4434.92</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$018</td>
<td style="text-align: center;">4474.43</td>
<td style="text-align: center;">+15c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$016</td>
<td style="text-align: center;">4517.95</td>
<td style="text-align: center;">+32c</td>
</tr>
<tr>
<td style="text-align: center;">110</td>
<td style="text-align: center;">D-8</td>
<td style="text-align: center;">4698.64</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$017</td>
<td style="text-align: center;">4660.87</td>
<td style="text-align: center;">-14c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$015</td>
<td style="text-align: center;">4723.32</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">111</td>
<td style="text-align: center;">D#8</td>
<td style="text-align: center;">4978.03</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$015</td>
<td style="text-align: center;">5084.58</td>
<td style="text-align: center;">+37c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$014</td>
<td style="text-align: center;">4948.24</td>
<td style="text-align: center;">-10c</td>
</tr>
<tr>
<td style="text-align: center;">112</td>
<td style="text-align: center;">E-8</td>
<td style="text-align: center;">5274.04</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$014</td>
<td style="text-align: center;">5326.71</td>
<td style="text-align: center;">+17c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$013</td>
<td style="text-align: center;">5195.65</td>
<td style="text-align: center;">-26c</td>
</tr>
<tr>
<td style="text-align: center;">113</td>
<td style="text-align: center;">F-8</td>
<td style="text-align: center;">5587.65</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$013</td>
<td style="text-align: center;">5593.04</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$012</td>
<td style="text-align: center;">5469.10</td>
<td style="text-align: center;">-37c</td>
</tr>
<tr>
<td style="text-align: center;">114</td>
<td style="text-align: center;">F#8</td>
<td style="text-align: center;">5919.91</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$012</td>
<td style="text-align: center;">5887.41</td>
<td style="text-align: center;">-10c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$011</td>
<td style="text-align: center;">5772.94</td>
<td style="text-align: center;">-44c</td>
</tr>
<tr>
<td style="text-align: center;">115</td>
<td style="text-align: center;">G-8</td>
<td style="text-align: center;">6271.93</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$011</td>
<td style="text-align: center;">6214.49</td>
<td style="text-align: center;">-16c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$010</td>
<td style="text-align: center;">6112.53</td>
<td style="text-align: center;">-45c</td>
</tr>
<tr>
<td style="text-align: center;">116</td>
<td style="text-align: center;">G#8</td>
<td style="text-align: center;">6644.88</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$010</td>
<td style="text-align: center;">6580.05</td>
<td style="text-align: center;">-17c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00F</td>
<td style="text-align: center;">6494.56</td>
<td style="text-align: center;">-40c</td>
</tr>
<tr>
<td style="text-align: center;">117</td>
<td style="text-align: center;">A-8</td>
<td style="text-align: center;">7040.00</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00F</td>
<td style="text-align: center;">6991.30</td>
<td style="text-align: center;">-12c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00E</td>
<td style="text-align: center;">6927.53</td>
<td style="text-align: center;">-28c</td>
</tr>
<tr>
<td style="text-align: center;">118</td>
<td style="text-align: center;">A#8</td>
<td style="text-align: center;">7458.62</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00E</td>
<td style="text-align: center;">7457.39</td>
<td style="text-align: center;">+0c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00D</td>
<td style="text-align: center;">7422.35</td>
<td style="text-align: center;">-8c</td>
</tr>
<tr>
<td style="text-align: center;">119</td>
<td style="text-align: center;">B-8</td>
<td style="text-align: center;">7902.13</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00D</td>
<td style="text-align: center;">7990.06</td>
<td style="text-align: center;">+19c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00C</td>
<td style="text-align: center;">7993.30</td>
<td style="text-align: center;">+20c</td>
</tr>
<tr>
<td style="text-align: center;">120</td>
<td style="text-align: center;">C-9</td>
<td style="text-align: center;">8372.02</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00C</td>
<td style="text-align: center;">8604.68</td>
<td style="text-align: center;">+47c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00B</td>
<td style="text-align: center;">8659.41</td>
<td style="text-align: center;">+58c</td>
</tr>
<tr>
<td style="text-align: center;">121</td>
<td style="text-align: center;">C#9</td>
<td style="text-align: center;">8869.84</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00C</td>
<td style="text-align: center;">8604.68</td>
<td style="text-align: center;">-53c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00B</td>
<td style="text-align: center;">8659.41</td>
<td style="text-align: center;">-42c</td>
</tr>
<tr>
<td style="text-align: center;">122</td>
<td style="text-align: center;">D-9</td>
<td style="text-align: center;">9397.27</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00B</td>
<td style="text-align: center;">9321.73</td>
<td style="text-align: center;">-14c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$00A</td>
<td style="text-align: center;">9446.63</td>
<td style="text-align: center;">+9c</td>
</tr>
<tr>
<td style="text-align: center;">123</td>
<td style="text-align: center;">D#9</td>
<td style="text-align: center;">9956.06</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00A</td>
<td style="text-align: center;">10169.16</td>
<td style="text-align: center;">+37c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$009</td>
<td style="text-align: center;">10391.29</td>
<td style="text-align: center;">+74c</td>
</tr>
<tr>
<td style="text-align: center;">124</td>
<td style="text-align: center;">E-9</td>
<td style="text-align: center;">10548.08</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$00A</td>
<td style="text-align: center;">10169.16</td>
<td style="text-align: center;">-63c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$009</td>
<td style="text-align: center;">10391.29</td>
<td style="text-align: center;">-26c</td>
</tr>
<tr>
<td style="text-align: center;">125</td>
<td style="text-align: center;">F-9</td>
<td style="text-align: center;">11175.30</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$009</td>
<td style="text-align: center;">11186.08</td>
<td style="text-align: center;">+2c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$008</td>
<td style="text-align: center;">11545.88</td>
<td style="text-align: center;">+56c</td>
</tr>
<tr>
<td style="text-align: center;">126</td>
<td style="text-align: center;">F#9</td>
<td style="text-align: center;">11839.82</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$008</td>
<td style="text-align: center;">12428.98</td>
<td style="text-align: center;">+84c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">$008</td>
<td style="text-align: center;">11545.88</td>
<td style="text-align: center;">-44c</td>
</tr>
<tr>
<td style="text-align: center;">127</td>
<td style="text-align: center;">G-9</td>
<td style="text-align: center;">12543.85</td>
<td style="text-align: center;" data-bgcolor="FFDDDD">$008</td>
<td style="text-align: center;">12428.98</td>
<td style="text-align: center;">-16c</td>
<td style="text-align: center;" data-bgcolor="DDDDFF">-</td>
<td style="text-align: center;">-</td>
<td style="text-align: center;">-</td>
</tr>
</tbody>
</table>

## Code Friendly Tables {#Code_Friendly_Tables}

Here are the period values from the above table in ca65 syntax, split into low/high byte halves for convenient indexing.

**Note**: These include every MIDI index from 0-127. This means the first ~32 entries in the table contain nonsense values, as the equivalent periods are larger than the range of the 11-bit timer. Programmers may wish to trim these early entries to save space.

### NTSC Periods {#NTSC_Periods}

    ntsc_period_low:
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $f1, $7f, $13, $ad, $4d, $f3, $9d
      .byte $4c, $00, $b8, $74, $34, $f8, $bf, $89
      .byte $56, $26, $f9, $ce, $a6, $80, $5c, $3a
      .byte $1a, $fb, $df, $c4, $ab, $93, $7c, $67
      .byte $52, $3f, $2d, $1c, $0c, $fd, $ef, $e1
      .byte $d5, $c9, $bd, $b3, $a9, $9f, $96, $8e
      .byte $86, $7e, $77, $70, $6a, $64, $5e, $59
      .byte $54, $4f, $4b, $46, $42, $3f, $3b, $38
      .byte $34, $31, $2f, $2c, $29, $27, $25, $23
      .byte $21, $1f, $1d, $1b, $1a, $18, $17, $15
      .byte $14, $13, $12, $11, $10, $0f, $0e, $0d
      .byte $0c, $0c, $0b, $0a, $0a, $09, $08, $08

    ntsc_period_high:
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $06, $06, $05, $05
      .byte $05, $05, $04, $04, $04, $03, $03, $03
      .byte $03, $03, $02, $02, $02, $02, $02, $02
      .byte $02, $01, $01, $01, $01, $01, $01, $01
      .byte $01, $01, $01, $01, $01, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00

### PAL Periods {#PAL_Periods}

    pal_period_low:
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $ff, $ff, $ff, $ff, $ff, $ff, $ff, $ff
      .byte $d1, $60, $f6, $92, $34, $db, $86, $37
      .byte $ec, $a5, $62, $23, $e8, $b0, $7b, $49
      .byte $19, $ed, $c3, $9b, $75, $52, $31, $11
      .byte $f3, $d7, $bd, $a4, $8c, $76, $61, $4d
      .byte $3a, $29, $18, $08, $f9, $eb, $de, $d1
      .byte $c6, $ba, $b0, $a6, $9d, $94, $8b, $84
      .byte $7c, $75, $6e, $68, $62, $5d, $57, $52
      .byte $4e, $49, $45, $41, $3e, $3a, $37, $34
      .byte $31, $2e, $2b, $29, $26, $24, $22, $20
      .byte $1e, $1d, $1b, $19, $18, $16, $15, $14
      .byte $13, $12, $11, $10, $0f, $0e, $0d, $0c
      .byte $0b, $0b, $0a, $09, $09, $08, $08, $07

    pal_period_high:
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $07, $07, $07, $07, $07, $07
      .byte $07, $07, $06, $06, $06, $05, $05, $05
      .byte $04, $04, $04, $04, $03, $03, $03, $03
      .byte $03, $02, $02, $02, $02, $02, $02, $02
      .byte $01, $01, $01, $01, $01, $01, $01, $01
      .byte $01, $01, $01, $01, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00
      .byte $00, $00, $00, $00, $00, $00, $00, $00

## See also {#See_also}

- [APU period table](../APU_period_table/)

<table class="collapsible collapsed" style="width: 100%; border: 1px solid #000; padding: 0">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th colspan="2" style="background-color: #CCF; font-size: 120%">V | <del>T</del>
<a href="../APU/">APU</a></th>
</tr>
</thead>
<tbody>
<tr>
<td style="background-color: #DDF"><strong>Concepts</strong></td>
<td><a href="../APU/">Registers</a> | <a href="../APU_basics/">Basics</a> | <a href="../APU_period_table/">Period table</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Channels</strong></td>
<td><a href="../APU_Pulse/">Pulse 1 and Pulse 2</a> | <a href="../APU_Triangle/">Triangle</a> | <a href="../APU_Noise/">Noise</a> | <a href="../APU_DMC/">DMC</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Components</strong></td>
<td><a href="../APU_Envelope/">Envelope</a> | <a href="../APU_Pulse/">Sweep</a> | <a href="../APU_Frame_Counter/">Frame Counter</a> | <a href="../APU_Length_Counter/">Length Counter</a> | <a href="../DMA/">DMA</a> | <a href="../APU_Mixer/">Mixer</a></td>
</tr>
<tr>
<td style="background-color: #DDF"><strong>Expansion audio</strong></td>
<td><a href="../FDS_audio/">2C33 (Disk System)</a> | <a href="../MMC5_audio/">MMC5</a> | <a href="../VRC6_audio/">VRC6</a> | <a href="../VRC7_audio/">VRC7</a> | <a href="../Namco_163_audio/">Namco 163</a> | <a href="../Sunsoft_5B_audio/">Sunsoft 5B</a> | <a href="../INES_Mapper_086/#Registers">µPD7756</a></td>
</tr>
</tbody>
</table>
