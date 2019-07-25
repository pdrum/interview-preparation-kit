package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Tail struct {
	Counts             []int
	SortedStartIndices []int
}

func NewTail() Tail {
	return Tail{
		Counts:             make([]int, 201),
		SortedStartIndices: make([]int, 201),
	}
}

func (t *Tail) resetSortedStartIndices() {
	t.SortedStartIndices[0] = 0
	for num := range t.Counts {
		if num == 0 {
			continue
		}
		t.SortedStartIndices[num] = t.SortedStartIndices[num-1] + t.Counts[num-1]
	}
}

func (t *Tail) Add(num int) {
	t.Counts[num]++
	t.resetSortedStartIndices()
}

func (t *Tail) Remove(num int) {
	t.Counts[num]--
	t.resetSortedStartIndices()
}

func (t Tail) NumAtIndex(targetIndex int) int {
	for i := len(t.SortedStartIndices) - 1; i >= 0; i-- {
		if t.SortedStartIndices[i] <= targetIndex && t.Counts[i] > 0 {
			return i
		}
	}
	panic("WAT")
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int, d int) int {
	tail := NewTail()
	notifs := 0
	for i, value := range expenditure {
		if i < d {
			tail.Add(value)
			continue
		}
		var twiceTheMedian int
		if d%2 == 1 {
			twiceTheMedian = tail.NumAtIndex(d/2) * 2
		} else {
			twiceTheMedian = tail.NumAtIndex(d/2) + tail.NumAtIndex((d/2)-1)
		}
		if value >= twiceTheMedian {
			notifs++
		}
		tail.Remove(expenditure[i-d])
		tail.Add(value)
	}
	return notifs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	n, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)

	d, err := strconv.Atoi(nd[1])
	checkError(err)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int

	for i := 0; i < int(n); i++ {
		expenditureItem, err := strconv.Atoi(expenditureTemp[i])
		checkError(err)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
