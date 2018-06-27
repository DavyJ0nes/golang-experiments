# Reversing Strings

## Aim

Wanted to see how standard way of reversing a string in Go worked. Haven't looked into simple array reversing before so was quite interesting to visualise its output.

## Output

We can see from output below that the loop starts from either end of the slice and works its way inwards. Is quite an efficient way of doing it because the number of iterations will always be length of input string / 2.

```shell
Input:    Golang Experiments

---- char length: 18
l: 71         [0]    (G) | r: 115        [17]    (s)
l: 111        [1]    (o) | r: 116        [16]    (t)
l: 108        [2]    (l) | r: 110        [15]    (n)
l: 97         [3]    (a) | r: 101        [14]    (e)
l: 110        [4]    (n) | r: 109        [13]    (m)
l: 103        [5]    (g) | r: 105        [12]    (i)
l: 32         [6]    ( ) | r: 114        [11]    (r)
l: 69         [7]    (E) | r: 101        [10]    (e)
l: 120        [8]    (x) | r: 112        [9]     (p)
---- iterations: 9

Output:     stnemirepxE gnaloG
```
