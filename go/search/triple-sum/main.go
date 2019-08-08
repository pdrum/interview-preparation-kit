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

func unique(input []int) []int {
	set := map[int]interface{}{}
	for _, num := range input {
		set[num] = true
	}
	var result []int
	for num := range set {
		result = append(result, num)
	}
	return result
}

func countLte(nums []int, target int) int {
	firstLargerIndex := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	if firstLargerIndex >= 0 {
		return firstLargerIndex
	}
	return len(nums)
}

// Complete the triplets function below.
func triplets(a []int, b []int, c []int) int64 {
	var result int64
	a = unique(a)
	b = unique(b)
	c = unique(c)
	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)
	for _, num := range b {
		result += int64(countLte(a, num) * countLte(c, num))
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lena, err := strconv.Atoi(lenaLenbLenc[0])
	checkError(err)

	lenb, err := strconv.Atoi(lenaLenbLenc[1])
	checkError(err)

	lenc, err := strconv.Atoi(lenaLenbLenc[2])
	checkError(err)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int

	for i := 0; i < int(lena); i++ {
		arraItem, err := strconv.Atoi(arraTemp[i])
		checkError(err)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int

	for i := 0; i < int(lenb); i++ {
		arrbItem, err := strconv.Atoi(arrbTemp[i])
		checkError(err)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int

	for i := 0; i < int(lenc); i++ {
		arrcItem, err := strconv.Atoi(arrcTemp[i])
		checkError(err)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

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
