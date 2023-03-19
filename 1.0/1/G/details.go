package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcDetails(n, k, m int) (int, int) {
	kCount := int(math.Floor(float64(n) / float64(k)))
	nRemainder := n - (k * kCount)
	if kCount <= 0 {
		return 0, 0
	}
	mCount := int(math.Floor(float64(k) / float64(m)))
	if mCount <= 0 {
		return 0, 0
	}
	kRemainder := (k - (m * mCount)) * kCount
	totalDetails := mCount * kCount
	totalRemainders := nRemainder + kRemainder
	return totalDetails, totalRemainders
}

func problem(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(input[0])
	k, _ := strconv.Atoi(input[1])
	m, _ := strconv.Atoi(input[2])
	totalDetails, iterDetails := 0, 0
	for int(math.Floor(float64(n)/float64(k))) > 0 {
		iterDetails, n = calcDetails(n, k, m)
		totalDetails += iterDetails
	}

	return strconv.Itoa(totalDetails)
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
