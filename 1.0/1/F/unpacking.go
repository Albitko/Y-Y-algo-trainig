package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type rectangle struct {
	x int
	y int
}

func MinIntSlice(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	return
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func calcSquare(x1, y1, x2, y2 int) (int, int) {
	var results []int
	edges := make(map[int]rectangle)

	results = append(results, (x1+x2)*Max(y1, y2))
	edges[(x1+x2)*Max(y1, y2)] = rectangle{(x1 + x2), Max(y1, y2)}

	results = append(results, (x1+y2)*Max(y1, x2))
	edges[(x1+y2)*Max(y1, x2)] = rectangle{(x1 + y2), Max(y1, x2)}

	results = append(results, (y1+y2)*Max(x1, x2))
	edges[(y1+y2)*Max(x1, x2)] = rectangle{(y1 + y2), Max(x1, x2)}

	results = append(results, (y1+x2)*Max(x1, y2))
	edges[(y1+x2)*Max(x1, y2)] = rectangle{(y1 + x2), Max(x1, y2)}
	min := MinIntSlice(results)
	minRect := edges[min]
	return minRect.x, minRect.y
}

func problem(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	w1, _ := strconv.Atoi(input[0])
	l1, _ := strconv.Atoi(input[1])
	w2, _ := strconv.Atoi(input[2])
	l2, _ := strconv.Atoi(input[3])
	var x, y int
	if w1*l1 < w2*l2 {
		x, y = calcSquare(w2, l2, w1, l1)
	} else {
		x, y = calcSquare(w1, l1, w2, l2)
	}
	return strconv.Itoa(x) + " " + strconv.Itoa(y)
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
