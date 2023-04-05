package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func problem(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	values := scanner.Text()
	var lastValue int

	valuesSlice := strings.Split(values, " ")
	for i, value := range valuesSlice {
		currentValue, _ := strconv.Atoi(value)
		if i == 0 {
			lastValue = currentValue
		} else {
			if currentValue <= lastValue {
				return "NO"
			}
			lastValue = currentValue
		}
	}

	return "YES"
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
