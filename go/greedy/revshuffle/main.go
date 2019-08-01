package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func countChars(s string) map[rune]int {
	result := map[rune]int{}
	for _, char := range []rune(s) {
		result[char]++
	}
	return result
}

func sumValues(m map[rune]int) int {
	result := 0
	for _, v := range m {
		result += v
	}
	return result
}

func countCharsInA(s string) map[rune]int {
	result := countChars(s)
	for char, cnt := range result {
		if cnt%2 == 0 {
			result[char] = cnt / 2
		} else {
			result[char] = (cnt / 2) + 1
		}
	}
	return result
}

type MinTracker struct {
	sortedUniqueRunes []rune
	minIndex          int
	charsNeededInA    map[rune]int
}

func NewMinTracker(s string) MinTracker {
	result := MinTracker{}
	result.charsNeededInA = countCharsInA(s)
	for char := range result.charsNeededInA {
		result.sortedUniqueRunes = append(result.sortedUniqueRunes, char)
	}
	sort.Slice(result.sortedUniqueRunes, func(i, j int) bool {
		return result.sortedUniqueRunes[i] < result.sortedUniqueRunes[j]
	})
	result.minIndex = 0
	return result
}

func (m *MinTracker) CurrentMin() rune {
	for {
		if m.minIndex >= len(m.sortedUniqueRunes) {
			panic("Index out of range")
		}

		possibleMin := m.sortedUniqueRunes[m.minIndex]
		allPicked := m.charsNeededInA[possibleMin] == 0
		if allPicked {
			m.minIndex++
		} else {
			return possibleMin
		}
	}
}

func (m *MinTracker) PickedForA(char rune) {
	if m.charsNeededInA[char] == 0 {
		panic("char not needed for a")
	}
	m.charsNeededInA[char]--
}

type IgnoranceTracker struct {
	charsNeededInA map[rune]int
	countLeftInS   map[rune]int
}

func NewIgnoranceTracker(s string) IgnoranceTracker {
	return IgnoranceTracker{
		charsNeededInA: countCharsInA(s),
		countLeftInS:   countChars(s),
	}
}

func (i *IgnoranceTracker) CanIgnore(char rune) bool {
	return i.countLeftInS[char]-1 >= i.charsNeededInA[char]
}

func (m *IgnoranceTracker) PickedForA(char rune) {
	if m.charsNeededInA[char] == 0 {
		panic("char not needed for a")
	}
	m.charsNeededInA[char]--
}

func (i *IgnoranceTracker) Visited(char rune) {
	if i.countLeftInS[char] == 0 {
		panic(fmt.Sprintf("All such chars already visited %v", char))
	}
	i.countLeftInS[char]--
}

type IgnoranceBreakpoint struct {
	ITracker    IgnoranceTracker
	MTracker    MinTracker
	I           int
	IgnoredChar rune
}

func cpMap(input map[rune]int) map[rune]int {
	result := map[rune]int{}
	for k, v := range input {
		result[k] = v
	}
	return result
}

func NewBreakpoint(iTracker IgnoranceTracker, mTracker MinTracker, i int, char rune) *IgnoranceBreakpoint {
	return &IgnoranceBreakpoint{
		ITracker: IgnoranceTracker{
			charsNeededInA: cpMap(iTracker.charsNeededInA),
			countLeftInS:   cpMap(iTracker.countLeftInS),
		},
		MTracker: MinTracker{
			charsNeededInA:    cpMap(mTracker.charsNeededInA),
			sortedUniqueRunes: mTracker.sortedUniqueRunes,
			minIndex:          mTracker.minIndex,
		},
		I:           i,
		IgnoredChar: char,
	}
}

// Complete the reverseShuffleMerge function below.
func reverseShuffleMerge(s string) string {
	charsNeeded := sumValues(countCharsInA(s))
	iTracker := NewIgnoranceTracker(s)
	mTracker := NewMinTracker(s)
	var result []rune
	var lastBreakpoint *IgnoranceBreakpoint
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		char := []rune(s)[i]
		if len(result) == charsNeeded {
			break
		}
		if mTracker.CurrentMin() == char {
			result = append(result, char)
			mTracker.PickedForA(char)
			iTracker.PickedForA(char)
			iTracker.Visited(char)
			lastBreakpoint = nil
		} else if !iTracker.CanIgnore(char) {
			if lastBreakpoint != nil && lastBreakpoint.IgnoredChar <= char {
				i = lastBreakpoint.I
				char = []rune(s)[i]
				iTracker = lastBreakpoint.ITracker
				mTracker = lastBreakpoint.MTracker
			}
			iTracker.Visited(char)
			result = append(result, char)
			mTracker.PickedForA(char)
			iTracker.PickedForA(char)
			lastBreakpoint = nil
		} else {
			if iTracker.charsNeededInA[char] > 0 && (lastBreakpoint == nil || (char < lastBreakpoint.IgnoredChar)) {
				lastBreakpoint = NewBreakpoint(iTracker, mTracker, i, char)
			}
			iTracker.Visited(char)
		}

	}
	return string(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := reverseShuffleMerge(s)

	fmt.Fprintf(writer, "%s\n", result)

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
