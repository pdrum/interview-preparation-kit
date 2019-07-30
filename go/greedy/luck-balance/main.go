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

func sum(values []int) int {
	result := 0
	for _, num := range values {
		result += num
	}
	return result
}

// Complete the luckBalance function below.
func luckBalance(k int, contests [][]int) int {
	importants := []int{}
	unimportants := []int{}
	for _, row := range contests {
		luck := row[0]
		important := row[1]
		if important == 1 {
			importants = append(importants, luck)
		} else {
			unimportants = append(unimportants, luck)
		}
	}
	if len(importants) <= k {
		return sum(unimportants) + sum(importants)
	}
	sort.Ints(importants)
	sort.Ints(unimportants)
	return sum(unimportants) + sum(importants[len(importants) - k:]) - sum(importants[:len(importants) - k])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nk := strings.Split(readLine(reader), " ")

	n, err := strconv.Atoi(nk[0])
	checkError(err)

	k, err := strconv.Atoi(nk[1])
	checkError(err)

	var contests [][]int
	for i := 0; i < n; i++ {
		contestsRowTemp := strings.Split(readLine(reader), " ")

		var contestsRow []int
		for _, contestsRowItem := range contestsRowTemp {
			contestsItem, err := strconv.Atoi(contestsRowItem)
			checkError(err)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

	fmt.Fprintf(writer, "%d\n", result)

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
