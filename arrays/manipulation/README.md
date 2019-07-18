Array Manipulation
------------------
This is solution of [array manipulation](https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays)
problem on hackerrank. There is a straight forward `O(n^2)` answer to the
problem but based on size of input a linear solution is desired.

I approached the problem like the problem of checking matching paranthesis.
I created an array of length `n`, then when I should add k to all
numbers between indexes `a` and `b`, what I do instead is I keep in
mind when I later iterate over the array and pass over index `a` I should
add `k` to the running sum. Then when I pass over index `b+1` I should
subtract `k` from the running sum. Max value of that running sum is
the solution to the problem. Ofcourse because `a` and `b` in input
are one based, in code I first made them zero based.
