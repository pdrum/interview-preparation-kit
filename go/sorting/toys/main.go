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

// Complete the maximumToys function below.
func maximumToys(prices []int, k int) int {
	sort.Ints(prices)
	count := 0
	currentSum := 0
	for _, price := range prices {
		currentSum += price
		if currentSum < k {
			count ++
		} else {
			break
		}
	}
	return count
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

	pricesTemp := strings.Split(readLine(reader), " ")

	var prices []int

	for i := 0; i < int(n); i++ {
		pricesItem, err := strconv.Atoi(pricesTemp[i])
		checkError(err)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

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
