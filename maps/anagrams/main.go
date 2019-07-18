package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int {
	substrings := findSubstrings(s)
	sortedCharsOccurrenceMap := map[string]int{}
	for _, sub := range substrings {
		cnt, _ := sortedCharsOccurrenceMap[sub.sortedRunes()]
		sortedCharsOccurrenceMap[sub.sortedRunes()] = cnt + 1
	}
	sum := 0
	for _, occurrenceCnt := range sortedCharsOccurrenceMap {
		sum += pairCountForSortedRune(occurrenceCnt)
	}
	return sum
}

func pairCountForSortedRune(totalCount int) int {
	return totalCount * (totalCount - 1) / 2
}

type SubString struct {
	StartIndex int
	Runes      []rune
}

func (ss SubString) String() string {
	return fmt.Sprintf("(start=%d string=%s)", ss.StartIndex, string(ss.Runes))
}

func (s SubString) sortedRunes() string {
	sortedRunes := []rune{}
	sortedRunes = append(sortedRunes, s.Runes...)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func findSubstrings(s string) []SubString {
	lenSubStrings := findSubstringsTillLen([]rune(s), len(s))
	result := []SubString{}
	for _, substrings := range lenSubStrings {
		result = append(result, substrings...)
	}
	return result
}

func findSubstringsTillLen(originalString []rune, length int) map[int][]SubString {
	if length == 1 {
		substrings := []SubString{}
		for index, char := range originalString {
			substrings = append(substrings, SubString{
				StartIndex: index,
				Runes:      []rune{char},
			})
		}
		return map[int][]SubString{1: substrings}
	}
	result := findSubstringsTillLen(originalString, length-1)
	prevLenSubstrings := result[length-1]
	result[length] = []SubString{}
	for _, subStr := range prevLenSubstrings {
		newCharIndex := subStr.StartIndex + len(subStr.Runes)
		if newCharIndex >= len(originalString) {
			continue
		}
		result[length] = append(result[length], SubString{
			StartIndex: subStr.StartIndex,
			Runes:      append(subStr.Runes, originalString[newCharIndex]),
		})
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
