package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func problem(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var line string
	var n int

	scanner.Scan()
	line = scanner.Text()
	n, _ = strconv.Atoi(line)

	prices := make([]int, n)

	scanner.Scan()
	row := scanner.Text()

	values := strings.Split(row, " ")

	var value int
	for i := 0; i < n; i++ {
		value, _ = strconv.Atoi(values[i])
		prices[i] = value
	}

	profit := 0
	for i := 1; i < n; i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	problem(os.Stdin)
}
