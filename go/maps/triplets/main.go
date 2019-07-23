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

func revert(input []int) []int {
	result := make([]int, len(input))
	for index, val := range input {
		result[len(input)-index-1] = val
	}
	return result
}

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	valueAscIndexMap := map[int64][]int{}
	for index, val := range arr {
		indices, _ := valueAscIndexMap[val]
		valueAscIndexMap[val] = append(indices, index)
	}
	valueDescIndexMap := map[int64][]int{}
	for num, indices := range valueAscIndexMap {
		valueDescIndexMap[num] = revert(indices)
	}

	var tripletsCnt int64 = 0
	for val, indices := range valueAscIndexMap {
		if val%r != 0 {
			continue
		}
		for _, index := range indices {
			nextNumIndices, _ := valueAscIndexMap[val*r]
			firstRight := sort.Search(len(nextNumIndices), func(i int) bool {
				return nextNumIndices[i] > index
			})
			if firstRight < 0 {
				continue
			}
			lenRight := len(nextNumIndices) - firstRight

			prevNumIndices, _ := valueDescIndexMap[val/r]
			lastLeft := sort.Search(len(prevNumIndices), func(i int) bool {
				return prevNumIndices[i] < index
			})
			if lastLeft < 0 {
				continue
			}
			lenLeft := len(prevNumIndices) - lastLeft
			tripletsCnt += int64(lenLeft * lenRight)
		}

	}
	return tripletsCnt
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
