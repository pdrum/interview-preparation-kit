package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the pairs function below.
func pairs(k int, arr []int) int {
	set := map[int]interface{}{}
	for _, num := range arr {
		set[num] = true
	}
	result := 0
	for _, num := range arr {
		if _, ok := set[num-k]; ok {
			result++
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	n, err := strconv.Atoi(nk[0])
	checkError(err)

	k, err := strconv.Atoi(nk[1])
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.Atoi(arrTemp[i])
		checkError(err)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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
