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

func problem(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	k1, _ := strconv.Atoi(input[0])
	n1 := 0
	m, _ := strconv.Atoi(input[1])
	k2, _ := strconv.Atoi(input[2])
	p2, _ := strconv.Atoi(input[3])
	n2, _ := strconv.Atoi(input[4])

	if p2 > 1 && k2 <= m {
		return "-1 -1"
	}

	if m == 1 {
		return "0 1"
	}

	apartmentsOnFloor := math.Ceil(float64(k2) / float64(m*(p2-1)+n2))
	apartmentsInFrontDoor := m * int(apartmentsOnFloor)
	p1 := math.Ceil(float64(k1) / float64(apartmentsInFrontDoor))
	n1 = int(math.Ceil(float64(k1-(apartmentsInFrontDoor*(int(p1)-1))) / float64(apartmentsOnFloor)))

	if float64(k2) == (float64(k2) / float64(m*(p2-1)+n2) * 2) {
		return strconv.FormatInt(int64(p1), 10) + " 0"
	}

	return strconv.FormatInt(int64(p1), 10) + " " + strconv.FormatInt(int64(n1), 10)
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
