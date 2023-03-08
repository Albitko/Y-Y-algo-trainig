package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func uniqueNumber(number string) string {
	number = strings.Replace(number, "-", "", -1)
	number = strings.Replace(number, "(", "", -1)
	number = strings.Replace(number, ")", "", -1)

	if len(number) == 7 {
		number = "495" + number
		return number
	}
	if string(number[0]) == "8" {
		number = number[1:]
		return number
	}
	if string(number[:2]) == "+7" {
		number = number[2:]
		return number
	}

	return ""
}

func problem(reader io.Reader) string {
	result := ""
	addedNumber := ""

	scanner := bufio.NewScanner(reader)

	var numbers []string
	for i := 0; i < 4; i++ {
		scanner.Scan()
		numbers = append(numbers, scanner.Text())
	}

	addedNumber, numbers = numbers[0], numbers[1:]

	for _, number := range numbers {
		if uniqueNumber(addedNumber) == uniqueNumber(number) {
			result += "YES\n"
		} else {
			result += "NO\n"
		}
	}

	return result
}

func main() {
	result := problem(os.Stdin)
	fmt.Println(result)
}
