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

func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Complete the minimumAbsoluteDifference function below.
func minimumAbsoluteDifference(arr []int) int {
	sort.Ints(arr)
	result := absDiff(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		result = findMin(result, absDiff(arr[i], arr[i-1]))
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.Atoi(arrTemp[i])
		checkError(err)
		arr = append(arr, arrItem)
	}

	result := minimumAbsoluteDifference(arr)

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
