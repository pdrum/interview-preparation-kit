package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countInversions function below.
func countInversions(arr []int) int64 {
	if len(arr) <= 1 {
		return 0
	}
	left := arr[:len(arr)/2]
	leftInversions := countInversions(left)
	right := arr[len(arr)/2:]
	rightInversions := countInversions(right)
	var crossingInversion int64 = 0
	leftIndex := 0
	rightIndex := 0
	sorted := []int{}
	for len(sorted) < len(arr) {
		if leftIndex == len(left) {
			sorted = append(sorted, right[rightIndex])
			rightIndex++
			continue
		}
		if rightIndex == len(right) {
			sorted = append(sorted, left[leftIndex])
			leftIndex++
			continue
		}
		if left[leftIndex] <= right[rightIndex] {
			sorted = append(sorted, left[leftIndex])
			leftIndex++
		} else {
			sorted = append(sorted, right[rightIndex])
			rightIndex++
			crossingInversion += int64(len(left)) - int64(leftIndex)
		}
	}
	for i := range sorted {
		arr[i] = sorted[i]
	}
	return crossingInversion + leftInversions + rightInversions
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	t, err := strconv.Atoi(readLine(reader))
	checkError(err)

	for tItr := 0; tItr < t; tItr++ {
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

		result := countInversions(arr)

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
