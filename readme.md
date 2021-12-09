# Advent Of Code 2021

My attempt at Advent of Code 2021, written in Go using a lot of brute force and ignorance

https://adventofcode.com/2021

## Progress

<table>
  <tr>
    <td>
      <pre>
Day 1 ⭐⭐
Day 2 ⭐⭐
Day 3 ⭐⭐
Day 4 ⭐⭐
Day 5 ⭐⭐
Day 6 ⭐⭐</pre>
    </td>
  </tr>
</table>

## Running

Use make and set the `DAY` variable to what ever day you are running or testing.

All tests should pass of course, to give you that false sense you're not wasting your life. They are run on the example input rather than the full input.

```bash
make test
make test DAY=2
```

To get the real answers use `make run`, for example for day 4

```bash
make run DAY=4

🌟══════════════════════════════════🌟
║   📅🎄🎁 Advent Of Code : Day 04   ║
🌟══════════════════════════════════🌟
  🟣🟢🟠 Day 4: Giant Squid 🟠🟢🟣

First board to win was 66, with number: 66
 84   78   03   44   96
 59   86   70   80  [48]
 93  [88]  52  [43]  61
[95] [66] [46] [62] [58]
 05   25   06   85   99
 ✨ Part 1 answer: 67718

Final board to win was 31, with number: 6
[96] [59] [72] [06] [38]
[60]  70  [76] [82] [46]
[47]  53   51  [64] [98]
[44]  25  [69] [81]  33
 73  [52] [10] [74] [55]
 ✨ Part 2 answer: 1831
```
