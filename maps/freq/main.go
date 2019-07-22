package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Freq struct {
	numCounts map[int]int
	countFreq map[int]int
}

func NewFreq() Freq {
	return Freq{
		numCounts: map[int]int{},
		countFreq: map[int]int{},
	}
}

func (f *Freq) add(num int) {
	oldCount, _ := f.numCounts[num]
	f.numCounts[num] = oldCount + 1
	f.onNumCountChanged(oldCount, oldCount + 1)
}

func (f *Freq) delete(num int) {
	oldCount, _ := f.numCounts[num]
	if oldCount == 0 {
		return
	}
	f.numCounts[num] = oldCount - 1
	f.onNumCountChanged(oldCount, oldCount - 1)
}

func (f *Freq) onNumCountChanged(oldCount int, newCount int) {
	f.countFreq[oldCount]--

	newCountFreq, _ := f.countFreq[newCount]
	f.countFreq[newCount] = newCountFreq + 1
}

func (f *Freq) exists(count int) bool {
	freq, _ := f.countFreq[count]
	return freq > 0
}

func boolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

// Complete the freqQuery function below.
func freqQuery(queries [][]int) []int {
	freq := NewFreq()
	result := []int{}
	for _, row := range queries {
		switch row[0] {
		case 1:
			freq.add(row[1])
		case 2:
			freq.delete(row[1])
		case 3:
			result = append(result, boolToInt(freq.exists(row[1])))
		default:
			panic("Invalid command")
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	q, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)

	var queries [][]int
	for i := 0; i < q; i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int
		for _, queriesRowItem := range queriesRowTemp {
			queriesItem, err := strconv.Atoi(queriesRowItem)
			checkError(err)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
