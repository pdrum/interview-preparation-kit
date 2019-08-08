package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Left  *Node
	Right *Node
	Val   int
	Depth int
}

type Tree struct {
	Root *Node
}

func (t *Tree) Traverse() []int {
	return t.visitNode([]int{}, t.Root)
}

func (t *Tree) visitNode(seenAlready []int, currentNode *Node) []int {
	if currentNode == nil {
		return seenAlready
	}
	result := seenAlready
	result = t.visitNode(result, currentNode.Left)
	result = append(result, currentNode.Val)
	result = t.visitNode(result, currentNode.Right)
	return result
}

func (t *Tree) SwapForK(k int) {
	t.swapForKAndNode(k, t.Root)
}

func (t *Tree) swapForKAndNode(k int, currentNode *Node) {
	if currentNode == nil {
		return
	}
	if currentNode.Depth%k == 0 {
		oldLeft := currentNode.Left
		currentNode.Left = currentNode.Right
		currentNode.Right = oldLeft
	}
	t.swapForKAndNode(k, currentNode.Left)
	t.swapForKAndNode(k, currentNode.Right)
}

func addChildren(node *Node, indices [][]int) {
	if node == nil {
		return
	}

	rowIndex := node.Val - 1
	if rowIndex < len(indices) {
		if indices[rowIndex][0] != -1 {
			node.Left = &Node{Val: indices[rowIndex][0], Depth: node.Depth + 1}
		}
		if indices[rowIndex][1] != -1 {
			node.Right = &Node{Val: indices[rowIndex][1], Depth: node.Depth + 1}
		}
	}

	addChildren(node.Left, indices)
	addChildren(node.Right, indices)
}

func NewTree(indices [][]int) Tree {
	tree := Tree{
		Root: &Node{Val: 1, Depth: 1},
	}
	addChildren(tree.Root, indices)
	return tree
}

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int, queries []int) [][]int {
	var result [][]int
	tree := NewTree(indexes)
	for _, k := range queries {
		tree.SwapForK(k)
		result = append(result, tree.Traverse())
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)

	var indexes [][]int
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int
		for _, indexesRowItem := range indexesRowTemp {
			indexesItem, err := strconv.Atoi(indexesRowItem)
			checkError(err)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != 2 {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItem, err := strconv.Atoi(readLine(reader))
		checkError(err)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
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
