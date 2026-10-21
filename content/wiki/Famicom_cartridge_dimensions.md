---
title: "Famicom cartridge dimensions"
url: "/wiki/Famicom_cartridge_dimensions/"
description: "This page documents the dimensions of common Famicom cartridges, including their cartridge shell and their corresponding PCB type. The source files are open-source and available on Github."
wikiPage: "Famicom_cartridge_dimensions"
wikiCategories:
  - name: "Hardware"
    url: "/category/Hardware/"
---

This page documents the dimensions of common Famicom cartridges, including their cartridge shell and their corresponding PCB type. The source files are open-source and available on [Github](https://github.com/Gumball2415/NES-Famicom-Cartridge-Dimensions).

*TODO: identify more cartridge shell forms*

## Famicom cartridge shell outline {#Famicom_cartridge_shell_outline}

### Small cartridge form {#Small_cartridge_form}

#### HVC-TxROM {#HVC-TxROM}

*TODO: measure and publish dimensions*

#### <span id="HVC-NROM.2FCNROM"></span>HVC-NROM/CNROM {#HVC-NROM.2FCNROM}

*TODO: measure and publish dimensions*

### Large cartridge form {#Large_cartridge_form}

#### HVC-ExROM {#HVC-ExROM}

*TODO: measure and publish dimensions*

### Miscellaneous cartridge forms {#Miscellaneous_cartridge_forms}

*TODO: document and measure other forms*

## Famicom cartridge PCB outline {#Famicom_cartridge_PCB_outline}

These outlines of different cartridge PCB types are based on the most space-filling variant, so that designers may choose to cut off unused portions that their designs may not need. Some outlines of cartridges with no maximally space filling PCB are thus modified to be such, with measurements being based on adjacent types made by the same producer.

Each cartridge type varies in how the 60 pin card edge is implemented. The pin pitch is 2.54mm, with 2 rows of 30 pads. Each pad may be uniform in dimensions, or slightly thicker and longer for power/ground.

### <span id="HVC-Ax.2FFx.2FPx.2FSx.2FTxROM_.28MMC1-4_mapper_boards.29"></span>HVC-Ax/Fx/Px/Sx/TxROM (MMC1-4 mapper boards) {#HVC-Ax.2FFx.2FPx.2FSx.2FTxROM_.28MMC1-4_mapper_boards.29}

This outline is based on measurements from a HVC-TGROM-01 type cartridge PCB, compatible with any cartridge shell that supports any HVC-Ax/Fx/Px/Sx/TxROM PCB or similar/bigger. These measurements are measured with vernier calipers with 0.05 mm accuracy. Note that this outline also details where the soldermask ends for the card edge.

[<img src="../../wiki-images/HVC-TGROM-01-board.svg" width="843" height="303" />](../../wiki-images/HVC-TGROM-01-board.svg)

Technical drawing of the outlines of an HVC-TGROM-01 type cartridge PCB.

<table class="wikitable" data-align="center" style="margin: auto;">
<thead>
<tr>
<th colspan="3">Dimensions</th>
</tr>
</thead>
<tbody>
<tr>
<td>Designator</td>
<td>Dimensions in mm (±0.05 mm)</td>
<td>Description</td>
</tr>
<tr>
<td>A</td>
<td>90</td>
<td>Board width</td>
</tr>
<tr>
<td>A.1</td>
<td>78.4</td>
<td>Card edge width</td>
</tr>
<tr>
<td>A.2</td>
<td>81</td>
<td>Notched width</td>
</tr>
<tr>
<td>A.2.1</td>
<td>43</td>
<td>Inset board width 1</td>
</tr>
<tr>
<td>A.2.2</td>
<td>33</td>
<td>Inset board width 2</td>
</tr>
<tr>
<td>B</td>
<td>46.1</td>
<td>Board height</td>
</tr>
<tr>
<td>B.1</td>
<td>23.2</td>
<td>Inset board height 1</td>
</tr>
<tr>
<td>B.2</td>
<td>20.9</td>
<td>Inset board height 2</td>
</tr>
<tr>
<td>B.3</td>
<td>10.7</td>
<td>Card edge depth</td>
</tr>
<tr>
<td>C</td>
<td>1.2</td>
<td>Board thickness</td>
</tr>
<tr>
<td>D.1</td>
<td>3</td>
<td>Card edge base to pad top</td>
</tr>
<tr>
<td>D.2</td>
<td>4</td>
<td>Card edge base to soldermask keepout</td>
</tr>
<tr>
<td>D.3</td>
<td>6.7</td>
<td>Power pad height</td>
</tr>
<tr>
<td>D.4</td>
<td>5.7</td>
<td>Signal pad height</td>
</tr>
<tr>
<td>E.1</td>
<td>0.57</td>
<td>Power pad to card edge side</td>
</tr>
<tr>
<td>E.2</td>
<td>1.57</td>
<td>Signal pad to card edge side</td>
</tr>
<tr>
<td>E.3</td>
<td>2.6</td>
<td>Power pad width</td>
</tr>
<tr>
<td>E.4</td>
<td>1.6</td>
<td>Signal pad width</td>
</tr>
<tr>
<td>E.5</td>
<td>0.94</td>
<td>Space between pads</td>
</tr>
<tr>
<td>E.6</td>
<td>2.54</td>
<td>Pad spacing</td>
</tr>
</tbody>
</table>

### HVC-ExROM {#HVC-ExROM_2}

*TODO: measure and publish dimensions*

### <span id="HVC-N.2FCx.2FGx.2FUxROM_01-05_.28discrete_mapper_boards.29"></span>HVC-N/Cx/Gx/UxROM 01-05 (discrete mapper boards) {#HVC-N.2FCx.2FGx.2FUxROM_01-05_.28discrete_mapper_boards.29}

This outline is based on measurements from a HVC-CNROM-256K-01 type cartridge PCB, augmented with dimensions from a HVC-TGROM-01 type PCB. This is compatible with any cartridge shell that supports any HVC-N/Cx/Gx/UxROM 01-05 PCB or similar/bigger. These measurements are measured with vernier calipers with 0.05 mm accuracy. Note that this outline also details where the soldermask ends for the card edge.

[<img src="../../wiki-images/HVC-CNROM-256K-01-board.svg" width="843" height="303" />](../../wiki-images/HVC-CNROM-256K-01-board.svg)

Technical drawing of the outlines of an augmented HVC-CNROM-256K-01 type cartridge PCB.

<table class="wikitable" data-align="center" style="margin: auto;">
<thead>
<tr>
<th colspan="3">Dimensions</th>
</tr>
</thead>
<tbody>
<tr>
<td>Designator</td>
<td>Dimensions in mm (±0.05 mm)</td>
<td>Description</td>
</tr>
<tr>
<td>A</td>
<td>90</td>
<td>Board width</td>
</tr>
<tr>
<td>A.1</td>
<td>78.4</td>
<td>Card edge width</td>
</tr>
<tr>
<td>A.2</td>
<td>68</td>
<td>Notched width</td>
</tr>
<tr>
<td>A.2.1</td>
<td>31.5</td>
<td>Inset board width 1</td>
</tr>
<tr>
<td>A.2.2</td>
<td>26.5</td>
<td>Inset board width 2</td>
</tr>
<tr>
<td>A.3</td>
<td>39.5</td>
<td>Hole 1 to board edge</td>
</tr>
<tr>
<td>A.4</td>
<td>1/2 of A</td>
<td>Hole 2 to board edge</td>
</tr>
<tr>
<td>A.5</td>
<td>3</td>
<td>Hole 1 diameter</td>
</tr>
<tr>
<td>A.6</td>
<td>5</td>
<td>Hole 2 diameter</td>
</tr>
<tr>
<td>B</td>
<td>46.1</td>
<td>Board height</td>
</tr>
<tr>
<td>B.1</td>
<td>23.2</td>
<td>Inset board height 1</td>
</tr>
<tr>
<td>B.2</td>
<td>20.9</td>
<td>Inset board height 2</td>
</tr>
<tr>
<td>B.3</td>
<td>10.7</td>
<td>Card edge depth</td>
</tr>
<tr>
<td>B.4</td>
<td>19.8</td>
<td>Hole 1 center height</td>
</tr>
<tr>
<td>B.5</td>
<td>10.3</td>
<td>Hole 2 center height</td>
</tr>
<tr>
<td>C</td>
<td>1.2</td>
<td>Board thickness</td>
</tr>
<tr>
<td>D.1</td>
<td>3</td>
<td>Card edge base to pad top</td>
</tr>
<tr>
<td>D.2</td>
<td>4</td>
<td>Card edge base to soldermask keepout</td>
</tr>
<tr>
<td>D.3</td>
<td>6.7</td>
<td>Power pad height</td>
</tr>
<tr>
<td>D.4</td>
<td>5.7</td>
<td>Signal pad height</td>
</tr>
<tr>
<td>E.1</td>
<td>0.57</td>
<td>Power pad to card edge side</td>
</tr>
<tr>
<td>E.2</td>
<td>1.57</td>
<td>Signal pad to card edge side</td>
</tr>
<tr>
<td>E.3</td>
<td>2.6</td>
<td>Power pad width</td>
</tr>
<tr>
<td>E.4</td>
<td>1.6</td>
<td>Signal pad width</td>
</tr>
<tr>
<td>E.5</td>
<td>0.94</td>
<td>Space between pads</td>
</tr>
<tr>
<td>E.6</td>
<td>2.54</td>
<td>Pad spacing</td>
</tr>
</tbody>
</table>
