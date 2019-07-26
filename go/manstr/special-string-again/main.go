package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type RepeatedChar struct {
	Char  rune
	Count int
}

func sumTill(num int) int64 {
	return int64(num) * int64(num+1) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Complete the substrCount function below.
func substrCount(s string) int64 {
	var repeatedChars []RepeatedChar
	for _, char := range []rune(s) {
		if len(repeatedChars) == 0 || repeatedChars[len(repeatedChars)-1].Char != char {
			repeatedChars = append(repeatedChars, RepeatedChar{
				Char:  char,
				Count: 1,
			})
			continue
		}
		repeatedChars[len(repeatedChars)-1].Count++
	}

	var specials int64 = 0
	for _, rc := range repeatedChars {
		specials += sumTill(rc.Count)
	}
	for i := 1; i < len(repeatedChars)-1; i++ {
		left := repeatedChars[i-1]
		right := repeatedChars[i+1]
		if left.Char != right.Char {
			continue
		}
		current := repeatedChars[i]
		if left.Char == current.Char || right.Char == current.Char || current.Count != 1 {
			continue
		}
		specials += int64(min(left.Count, right.Count))
	}
	return specials
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	_, err = strconv.Atoi(readLine(reader))
	checkError(err)

	s := readLine(reader)

	result := substrCount(s)

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
