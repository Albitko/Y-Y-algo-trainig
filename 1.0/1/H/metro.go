package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func problem(reader io.Reader) string {
	var a, b, n, m int
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	var firstPlatform []int
	var secondPlatform []int

	firstPlatmformMax := a*(n+1) + n
	firstPlatform = append(firstPlatform, firstPlatmformMax)
	firstPlatformMin := n + a*(n-1)
	firstPlatform = append(firstPlatform, firstPlatformMin)

	secondPlatmformMax := b*(m+1) + m
	secondPlatform = append(secondPlatform, secondPlatmformMax)
	secondPlatformMin := m + b*(m-1)
	secondPlatform = append(secondPlatform, secondPlatformMin)

	minFirst, maxFirst := MinMax(firstPlatform)
	minSecond, maxSecond := MinMax(secondPlatform)

	max := math.Min(float64(maxFirst), float64(maxSecond))
	min := math.Max(float64(minFirst), float64(minSecond))
	if min > max {
		return "-1"
	}
	return fmt.Sprint(min) + " " + fmt.Sprint(max)
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
