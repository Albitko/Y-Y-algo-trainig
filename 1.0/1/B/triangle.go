package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func problem(reader io.Reader) string {
	result := "YES"
	scanner := bufio.NewScanner(reader)

	var edges []int
	for i := 0; i < 3; i++ {
		scanner.Scan()
		edge, _ := strconv.Atoi(scanner.Text())
		edges = append(edges, edge)
	}

	for i := 0; i < 3; i++ {
		if edges[i] >= edges[(i+1)%3]+edges[(i+2)%3] {
			result = "NO"
		}
	}

	return result
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
