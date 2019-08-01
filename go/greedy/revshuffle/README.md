Solution to [Reverse Shuffle](https://www.hackerrank.com/challenges/reverse-shuffle-merge/problem) problem
----------------------------------------------------------------------------------------------------------
First I notice that I can simply come up with count of each character
in response. If a character has length `n` there is `n/2` of that
character in `A` if `n` is even and `(n/2)+1` if it is odd.

I iterate the string in reverse direction. For each character, I check
if I can skip that character or not. A character can be skipped if it
is not the lexographically smallest character to be added to result
and also we have enough of it left in remaining of the string. If I can't
skip a character I backtarck to the (first occurrence of)
lexographically smallest character I have
seen between last addition to `A` and this character and continue the
process from there.
