package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func boolToStr(input bool) string {
	if input {
		return "YES"
	}
	return "NO"
}

func areAllTheSame(numbers []int) bool {
	for _, num := range numbers {
		if num != numbers[0] {
			return false
		}
	}
	return true
}

func maxNumIndex(numbers []int) int {
	max := -1
	maxIndex := -1
	for i, num := range numbers {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return maxIndex
}

func validWithMaxFreqRemoval(countList []int) bool {
	var clone []int
	clone = append(clone, countList...)
	clone[maxNumIndex(clone)]--
	if areAllTheSame(clone) {
		return true
	}
	return false
}

func validWithFreq1Removal(countList []int) bool {
	oneCount := 0
	for _, cnt := range countList {
		if cnt == 1 {
			oneCount++
		}
	}
	if oneCount != 1 {
		return false
	}
	oneRemoved := []int{}
	for _, cnt := range countList {
		if cnt == 1 {
			continue
		}
		oneRemoved = append(oneRemoved, cnt)
	}
	return areAllTheSame(oneRemoved)
}

// Complete the isValid function below.
func isValid(s string) string {
	charCntMap := map[rune]int{}
	for _, char := range []rune(s) {
		cnt, _ := charCntMap[char]
		charCntMap[char] = cnt + 1
	}
	countList := []int{}
	for _, cnt := range charCntMap {
		countList = append(countList, cnt)
	}
	if areAllTheSame(countList) {
		return boolToStr(true)
	}
	if validWithMaxFreqRemoval(countList) {
		return boolToStr(true)
	}
	if validWithFreq1Removal(countList) {
		return boolToStr(true)
	}
	return boolToStr(false)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := isValid(s)

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
