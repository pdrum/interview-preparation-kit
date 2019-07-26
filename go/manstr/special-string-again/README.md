Solution to [special string again](https://www.hackerrank.com/challenges/special-palindrome-again/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings)
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

For solving this I first turned the string into a slice of struct
each member of which has a char and a count. Essentially I turned
`aaa` into `RepeatedChar{Char: a, Count: 3}`. Then I used the fact
that special polindromes are formed from two types of strings.

* `XXXXYZZZZ` strings. In these strings number of special substrings that
can be found equals number of Xs.
* `XXXX` strings. In these strings number of special substrings that
can be found equals sum of all natural numbers from 1 to k. (given the string has `k` Xs)