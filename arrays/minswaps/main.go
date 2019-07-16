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

func clone(arr []int) []int {
	newSlice := make([]int, len(arr))
	copy(newSlice, arr)
	return newSlice
}

func findFirstUnvisited(visitedArr []bool) (int, bool) {
	for index, val := range visitedArr {
		if !val {
			return index, true
		}
	}
	return 0, false
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int) int {
	sorted := clone(arr)
	sort.Ints(sorted)
	visited := make([]bool, len(arr))
	swaps := 0
	for {
		current, ok := findFirstUnvisited(visited)
		if !ok {
			return swaps
		}

		for {
			newCurrent := sort.SearchInts(sorted, arr[current])
			visited[current] = true
			current = newCurrent
			if visited[current] {
				break
			}
			swaps++
		}
	}
	return swaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := nTemp

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.Atoi(arrTemp[i])
		checkError(err)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
