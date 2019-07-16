package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	converted := []int{}
	for _, num := range q {
		converted = append(converted, int(num))
	}
	result, err := recursiveBribes(converted)
	if err != nil {
		os.Stdout.WriteString("Too chaotic\n")
		return
	}
	os.Stdout.WriteString(strconv.Itoa(result) + "\n")
}

func recursiveBribes(queue []int) (int, error) {
	if len(queue) == 0 {
		return 0, nil
	}
	if len(queue) == queue[len(queue)-1] {
		return recursiveBribes(queue[:len(queue)-1])
	}
	if len(queue) == queue[len(queue)-2] {
		swap(queue, len(queue)-1, len(queue)-2)
		rest, err := recursiveBribes(queue[:len(queue)-1])
		if err != nil {
			return 0, err
		}
		return rest + 1, nil
	}
	if len(queue) == queue[len(queue)-3] {
		swap(queue, len(queue)-3, len(queue)-2)
		swap(queue, len(queue)-2, len(queue)-1)
		rest, err := recursiveBribes(queue[:len(queue)-1])
		if err != nil {
			return 0, err
		}
		return rest + 2, nil
	}
	return 0, errors.New("invalid queue")
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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
