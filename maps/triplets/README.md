[Triplets](https://www.hackerrank.com/challenges/count-triplets-1/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps)
--------
The trick to come up with a solution with acceptable time complexity is
we should iterate the array and for each item, we should check if this
item (for example `k`) was to be the middle element in a triplet, then how many
`k*r` items would be on the right of it and how many `k/r` items would
be on the left. Then we can multiple those numbers to get the number of
triplets with that item as the second element.

For doing so we have to keep two maps, one that keeps track of at which
indices each number appears and keeps indices in ascending order
and another one that does exactly the same but keeps indices in descending order.
