package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printSorted(a, b int) {
	if a < b {
		fmt.Println(a, b)
	} else {
		fmt.Println(b, a)
	}
}

// Complete the whatFlavors function below.
func whatFlavors(costs []int, money int) {
	costIDsMap := map[int][]int{}
	for index, c := range costs {
		costIDsMap[c] = append(costIDsMap[c], index+1)
	}
	for _, cost := range costs {
		rest := money - cost
		if rest == cost && len(costIDsMap[cost]) > 1 {
			printSorted(costIDsMap[cost][0], costIDsMap[cost][1])
			return
		}
		if rest != cost && len(costIDsMap[rest]) > 0 {
			printSorted(costIDsMap[cost][0], costIDsMap[rest][0])
			return
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	t, err := strconv.Atoi(readLine(reader))
	checkError(err)

	for tItr := 0; tItr < int(t); tItr++ {
		money, err := strconv.Atoi(readLine(reader))
		checkError(err)

		n, err := strconv.Atoi(readLine(reader))
		checkError(err)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int

		for i := 0; i < int(n); i++ {
			costItem, err := strconv.Atoi(costTemp[i])
			checkError(err)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
