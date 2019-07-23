[Sherlock and Anagrams](https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps)
---------------------
The thing I noticed first is small number of inputs. Therefore
I am not going to optimize time complexity. Next I found out that
each two pairs of anagrams look the same if I sort the characters forming them.

For solving this problem I found all substrings of given string.
For each substring I create a string which is made by taking that
substring and sorting its charcters.
I call it a sorted form. For each substring I create a sorted form
and keep a map from each sorted form to number of substrings with that
sorted form.

Then I iterate over values of that map. Number of pairs of anagrams
with each sorted form equals number of ways I can choose 2 from
total number of substrings with that sorted form. (i.e. `(cnt + cnt - 1) / 2`)