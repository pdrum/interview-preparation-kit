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

// Complete the maxMin function below.
func maxMin(k int, arr []int) int {
	sort.Ints(arr)
	minValue := -1
	for minNumIndex := range arr {
		maxNumIndex := minNumIndex + k - 1
		if maxNumIndex >= len(arr) {
			return minValue
		}
		if minValue == -1 || arr[maxNumIndex] - arr[minNumIndex] < minValue {
			minValue = arr[maxNumIndex] - arr[minNumIndex]
		}
	}
	return minValue
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)

	k, err := strconv.Atoi(readLine(reader))
	checkError(err)

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.Atoi(readLine(reader))
		checkError(err)
		arr = append(arr, arrItem)
	}

	result := maxMin(k, arr)

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
