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

func revert(numbers []int) []int {
	result := make([]int, len(numbers))
	for i := range result {
		result[i] = numbers[len(numbers) - 1 - i]
	}
	return result
}

// Complete the getMinimumCost function below.
func getMinimumCost(k int, prices []int) int {
	sort.Ints(prices)
	prices = revert(prices)
	total := 0
	for i := 0; i < len(prices); i++ {
		total += prices[i] * ((i / k) + 1)
	}
	return total
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

	cTemp := strings.Split(readLine(reader), " ")

	var c []int

	for i := 0; i < n; i++ {
		cItem, err := strconv.Atoi(cTemp[i])
		checkError(err)
		c = append(c, cItem)
	}

	minimumCost := getMinimumCost(k, c)

	fmt.Fprintf(writer, "%d\n", minimumCost)

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
