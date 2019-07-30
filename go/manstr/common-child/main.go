package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func key(a, b int) string {
	return fmt.Sprintf("%d$%d", a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Complete the commonChild function below.
func commonChild(s1 []rune, s2 []rune) int {
	solutions := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		solutions[i] = make([]int, len(s2)+1)
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s2[j-1] {
				solutions[i][j] = solutions[i-1][j-1] + 1
			} else {
				solutions[i][j] = max(solutions[i][j-1], solutions[i-1][j])
			}
		}
	}
	return solutions[len(s1)][len(s2)]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild([]rune(s1), []rune(s2))

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
