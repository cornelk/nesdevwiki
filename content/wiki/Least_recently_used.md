---
title: "Least recently used"
url: "/wiki/Least_recently_used/"
description: "It's often useful to post-process the output of a pseudo-random number generator (PRNG) to favor outcomes not seen recently. The most obvious way to do this is to maintain a list of all outcomes,…"
wikiPage: "Least_recently_used"
---

It's often useful to post-process the output of a pseudo-[random number generator](../Random_number_generator/) (PRNG) to favor outcomes not seen recently. The most obvious way to do this is to maintain a list of all outcomes, ordered by the **least recently used (LRU).**

## Advantages {#Advantages}

Variety  
Not seeing the same outcome too often in close succession may improve the design. For example, in a block puzzle game such as the build phase of *RHDE*, some pieces may be easier to place than others, and a good mix of simple and twisty wall pieces makes the game more enjoyable. Or in a game where the player must defend targets, such as *Thwaite*, selecting a variety of targets to defend keeps the player on his toes.

Difficulty ramping  
Not seeing particular outcomes early in a game session may improve the game's approachability. For example, in *Thwaite*, the missile silo targets are more precious than house targets, and a novice player should be able to learn how to play without having to worry about losing silos too early. So it forces early targets to be houses. And in *RHDE*, the smaller, simpler wall pieces appear earlier in the game, with the larger, twistier pieces eased in gradually.

Simple arithmetic  
Many PRNGs produce a uniform distribution of outcomes over a range that is a power of 2. The common way to produce a number of outcomes other than a power of 2 is to multiply the PRNG's output by the number of outcomes and then shift down, but multiplication can be slow on a 6502. LRU allows the total number of outcomes can be something other than a power of 2 without having to use multiplication. *Thwaite* has twelve targets: ten houses and two missile silos. *RHDE* has a total of 18 different wall pieces, with the four smallest pieces appearing twice as often as the others, for a total of 22 outcomes in a multiset.

## Implementation {#Implementation}

Once the program has loaded all outcomes into an array, there are a couple ways to update it when a random outcome is needed.

### Simple method {#Simple_method}

1.  Choose one of the items at the head of the array and save it.
2.  Copy all items after this item forward by one position.
3.  Write this item to the last space of the array, and pass it to the rest of the simulation.

### Method with fewer copies {#Method_with_fewer_copies}

If you have a large enough set of outcomes that the O(n) copy is undesirable, such as selection of non-kitten items (NKIs) in an implementation of *robotfindskitten*, an optimization to LRU reminiscent of [Alleged Rivest Cipher Four (RC4)](https://en.wikipedia.org/wiki/RC4) treats the array as a [circular buffer](https://en.wikipedia.org/wiki/circular_buffer), using an extra variable holding the index of the current head of the array.

1.  Add the RNG's output to the head index modulo the number of outcomes, producing the target index.
2.  Swap the item at the target index with the item at the head index.
3.  Pass the item swapped into the head to the rest of the simulation, and advance the head by one modulo the number of outcomes.

### Special case for consecutive repeats {#Special_case_for_consecutive_repeats}

If you are concerned only about consecutive repeats, you can save the most recent value, generate a new value 0 to *n* - 2, and if it matches the most recent value, use *n* - 1 instead. This is especially efficient if the number of outcomes is one more than a power of two. One possible application is the 9 different pills in implementations of the virus-destroying game described in U.S. Patent 5,265,888<sup><span id="cite_ref-1"></span>[\[1\]](#cite_note-1)</sup> (now expired).

## Notes {#Notes}

1.  <span id="cite_note-1"></span>[↑](#cite_ref-1) Nintendo implemented this game as *Dr. Mario*, but the sequence of pieces doesn't reflect any LRU techniques.
